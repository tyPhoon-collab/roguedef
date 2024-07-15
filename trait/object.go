package trait

import (
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawer interface {
	Draw(screen *ebiten.Image)
}

type Updater interface {
	Update()
}

type Object struct {
	ID uuid.UUID
	Updater
	Drawer
	Intersector
	IntersectHandler
}

func (o *Object) WithUpdater(updater Updater) *Object {
	o.Updater = updater
	return o
}

func (o *Object) WithDrawer(drawer Drawer) *Object {
	o.Drawer = drawer
	return o
}

func (o *Object) WithIntersector(intersector Intersector) *Object {
	o.Intersector = intersector
	return o
}

func (o *Object) WithIntersectHandler(handler IntersectHandler) *Object {
	o.IntersectHandler = handler
	return o
}

func NewObject() *Object {
	return &Object{
		ID: uuid.New(),
	}
}
