package system

import (
	"fmt"
	"roguedef/ds"
	"roguedef/task"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type iD = uuid.UUID

type Game struct {
	objects           map[iD]*Object
	drawers           map[iD]Drawer
	updaters          map[iD]Updater
	intersects        map[iD]Intersector
	intersectHandlers map[iD]IntersectHandler
	tasks             *ds.Queue[task.Task]
	frameCount        int
}

func (g *Game) FrameCount() int {
	return g.frameCount
}

func (g *Game) Intersects() map[iD]Intersector {
	return g.intersects
}

func (g *Game) Update() error {
	g.executeTask()

	for id, o := range g.updaters {
		if g.objects[id].Active {
			o.Update()
		}
	}

	g.checkIntersects()

	g.frameCount++

	return nil
}

func (g *Game) executeTask() {
	t, ok := g.tasks.Peek()

	if !ok {
		return
	}

	if !t.Active() {
		_, _ = g.tasks.Pop()
		g.executeTask()
	}
	if t.ShouldExecute(g.frameCount) {
		err := t.Execute()
		if err != nil {
			panic(err)
		}
		_, _ = g.tasks.Pop()
		g.executeTask() // recursion
	}
}

func (g *Game) checkIntersects() {
	keys := make([]iD, 0, len(g.intersects))
	values := make([]Intersector, 0, len(g.intersects))

	for k, v := range g.intersects {
		keys = append(keys, k)
		values = append(values, v)
	}

	for i := 0; i < len(g.intersects); i++ {
		for j := i + 1; j < len(g.intersects); j++ {
			id, otherId := keys[i], keys[j]
			v1, v2 := values[i], values[j]
			if v1.Intersects(v2) {
				if handler, ok := g.intersectHandlers[otherId]; ok {
					handler.OnIntersect(g.objects[id])
				}
				if handler, ok := g.intersectHandlers[id]; ok {
					handler.OnIntersect(g.objects[otherId])
				}
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawers := make([]Drawer, 0)

	for id, v := range g.drawers {
		if g.objects[id].Active {
			drawers = append(drawers, v)
		}
	}

	slices.SortFunc(drawers, func(a, b Drawer) int {
		return a.Priority() - b.Priority()
	})
	for _, o := range drawers {
		o.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 640
}

func (g *Game) IsOutside(pos vec) bool {
	x, y := pos.X, pos.Y
	return x < 0 || y < 0 || x >= 320 || y >= 640
}

func (g *Game) ObjectsByTag(tag string) chan *Object {
	ch := make(chan *Object)

	go func() {
		defer close(ch)
		for _, o := range g.objects {
			if o.Tag == tag {
				ch <- o
			}
		}
	}()

	return ch
}

func (g *Game) ObjectByTag(tag string) *Object {
	for _, o := range g.objects {
		if o.Tag == tag {
			return o
		}
	}
	return nil
}

func (g *Game) Objects() chan *Object {
	ch := make(chan *Object)

	go func() {
		defer close(ch)
		for _, o := range g.objects {
			ch <- o
		}
	}()

	return ch
}

func (g *Game) AddObject(o *Object) {
	g.objects[o.ID] = o

	if o.Drawer != nil {
		g.drawers[o.ID] = o.Drawer
	}
	if o.Updater != nil {
		g.updaters[o.ID] = o.Updater
	}

	g.AddTaskPostFrame(func() error {
		if o.Intersector != nil {
			g.intersects[o.ID] = o.Intersector
		}
		if o.IntersectHandler != nil {
			g.intersectHandlers[o.ID] = o.IntersectHandler
		}

		return nil
	})
}

func (g *Game) AddObjectWithData(data Data) *Object {
	obj := NewObjectWithData(data)
	g.AddObject(obj)

	g.AddTaskPostFrame(func() error {
		data.Register(g, obj)

		return nil
	})

	return obj
}

func (g *Game) RemoveObject(id iD) {
	g.AddTaskPostFrame(func() error {
		obj, ok := g.objects[id]

		if !ok {
			fmt.Println("object not found")
			return nil
		}

		if obj.OnRemoveHandler != nil {
			obj.OnRemove()
		}

		delete(g.objects, id)
		delete(g.drawers, id)
		delete(g.updaters, id)
		delete(g.intersects, id)
		delete(g.intersectHandlers, id)

		return nil
	})
}

func (g *Game) AddTask(task task.Task) {
	for i, t := range g.tasks.Data() {
		if t.At() > task.At() {
			g.tasks.Insert(i, task)
			return
		}
	}
	g.tasks.Push(task)
}
func (g *Game) AddTaskPostFrame(do func() error) {
	g.AddTask(task.NewTask(g.frameCount, do))
}
func (g *Game) AddTaskAfter(after time.Duration, do func() error) {
	delayFrameCount := float64(ebiten.TPS()) * after.Seconds()
	g.AddTask(task.NewTask(g.frameCount+int(delayFrameCount), do))
}

func NewGame() *Game {
	game := &Game{
		objects:           make(map[iD]*Object),
		drawers:           make(map[iD]Drawer),
		updaters:          make(map[iD]Updater),
		intersects:        make(map[iD]Intersector),
		intersectHandlers: make(map[iD]IntersectHandler),
		tasks:             ds.NewQueue[task.Task](),
	}

	return game
}
