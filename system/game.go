package system

import (
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
	taskQueue         []task.Task
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

	for _, o := range g.updaters {
		o.Update()
	}

	g.checkIntersects()

	g.frameCount++

	return nil
}

func (g *Game) executeTask() {
	if len(g.taskQueue) > 0 {
		t := g.taskQueue[0]
		if t.ShouldExecute(g.frameCount) {
			err := t.Execute()
			if err != nil {
				panic(err)
			}
			g.taskQueue = g.taskQueue[1:]
			g.executeTask() // recursion
		}
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
	for _, o := range g.drawers {
		o.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 640
}

func (g *Game) AddObject(o *Object) {
	g.objects[o.ID] = o

	if o.Drawer != nil {
		g.drawers[o.ID] = o.Drawer
	}

	if o.Updater != nil {
		g.updaters[o.ID] = o.Updater
	}

	if o.Intersector != nil {
		g.intersects[o.ID] = o.Intersector
	}

	if o.IntersectHandler != nil {
		g.intersectHandlers[o.ID] = o.IntersectHandler
	}
}

func (g *Game) AddObjectWithData(data Data) *Object {
	obj := NewObjectWithData(data)
	g.AddObject(obj)
	data.Register(g, obj)

	return obj
}

func (g *Game) RemoveObject(id iD) {
	g.AddTaskPostFrame(func() error {
		delete(g.objects, id)
		delete(g.drawers, id)
		delete(g.updaters, id)
		delete(g.intersects, id)
		delete(g.intersectHandlers, id)

		return nil
	})
}

func (g *Game) AddTask(t task.Task) {
	for i := 0; i < len(g.taskQueue); i++ {
		if g.taskQueue[i].At() > t.At() {
			g.taskQueue = slices.Insert(g.taskQueue, i, t)
			return
		}
	}
	g.taskQueue = append(g.taskQueue, t)
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
		taskQueue:         make([]task.Task, 0),
	}

	return game
}
