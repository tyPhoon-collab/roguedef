package game

import (
	"roguedef/task"
	"roguedef/trait"
	"slices"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// short hand for trait.NewObject()
func new() *trait.Object {
	return trait.NewObject()
}

type id = uuid.UUID

type Game struct {
	player            *Player
	drawers           map[id]trait.Drawer
	updaters          map[id]trait.Updater
	intersects        map[id]trait.Intersector
	intersectHandlers map[id]trait.IntersectHandler
	taskQueue         []task.Task
	frameCount        int
}

func (g *Game) Update() error {
	g.executeTask()

	for _, o := range g.updaters {
		o.Update()
	}

	g.checkIntersects()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		bullet := g.newBulletObject()
		g.AddObject(bullet)
		g.addTask(task.NewTask(g.frameCount+ebiten.TPS()*3, func() error {
			g.RemoveObject(bullet.ID)
			return nil
		}))
	}

	g.frameCount++

	return nil
}

func (g *Game) executeTask() {
	if len(g.taskQueue) > 0 {
		t := g.taskQueue[0]
		if t.ShouldExecute(g.frameCount) {
			t.Execute()
			g.taskQueue = g.taskQueue[1:]
			g.executeTask() // recursion
		}
	}
}

func (g *Game) addTask(t task.Task) {
	for i := 0; i < len(g.taskQueue); i++ {
		if g.taskQueue[i].At() > t.At() {
			g.taskQueue = slices.Insert(g.taskQueue, i, t)
			return
		}
	}
	g.taskQueue = append(g.taskQueue, t)
}

func (g *Game) newBulletObject() *trait.Object {
	bullet := NewBullet()

	bullet.Velocity.Transform.Pos = g.player.Pos
	bullet.Set(Vec2{X: 0, Y: -10})

	return new().
		WithDrawer(bullet).
		WithUpdater(bullet).
		WithIntersector(bullet.intersect)
}

func (g *Game) checkIntersects() {
	keys := make([]id, 0, len(g.intersects))
	values := make([]trait.Intersector, 0, len(g.intersects))

	for k, v := range g.intersects {
		keys = append(keys, k)
		values = append(values, v)
	}

	for i := 0; i < len(g.intersects); i++ {
		for j := i + 1; j < len(g.intersects); j++ {
			k1, k2 := keys[i], keys[j]
			v1, v2 := values[i], values[j]
			if v1.Intersects(v2) {
				if handler, ok := g.intersectHandlers[k2]; ok {
					handler.OnIntersect(v2)
				}
				if handler, ok := g.intersectHandlers[k1]; ok {
					handler.OnIntersect(v1)
				}
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.player.Pos.String())

	for _, o := range g.drawers {
		o.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 640
}

func (g *Game) AddObject(o *trait.Object) *Game {
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
	return g
}

func (g *Game) RemoveObject(id id) *Game {
	delete(g.drawers, id)
	delete(g.updaters, id)
	delete(g.intersects, id)
	delete(g.intersectHandlers, id)
	return g
}

func NewGame() *Game {
	player, err := NewPlayer(Vec2{X: 136, Y: 550})
	if err != nil {
		panic(err)
	}
	cursor := NewCursor()

	game := &Game{
		player:            player,
		drawers:           make(map[id]trait.Drawer),
		updaters:          make(map[id]trait.Updater),
		intersects:        make(map[id]trait.Intersector),
		intersectHandlers: make(map[id]trait.IntersectHandler),
		taskQueue:         make([]task.Task, 0),
	}

	debug := NewDebug(*game)

	game.AddObject(trait.NewObjectWithData(player))

	game.AddObject(trait.NewObjectWithData(cursor))

	game.AddObject(trait.NewObjectWithData(debug))

	return game
}
