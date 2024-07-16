package system

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
	ID   uuid.UUID
	Data Data
	Updater
	Drawer
	Intersector
	IntersectHandler
}

type Data interface {
	Register(obj *Object)
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

func NewObjectWithData(data Data) *Object {
	obj := NewObject()
	obj.Data = data
	if o, ok := data.(Drawer); ok {
		obj.Drawer = o
	}
	if o, ok := data.(Updater); ok {
		obj.Updater = o
	}
	if o, ok := data.(IntersectHolder); ok {
		obj.Intersector = o.Intersect()
	}
	if o, ok := data.(IntersectHandler); ok {
		obj.IntersectHandler = o
	}
	return obj
}
