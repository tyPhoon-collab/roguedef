package system

import (
	"github.com/google/uuid"
)

type OnRemoveHandler interface {
	OnRemove()
}

type Object struct {
	ID     uuid.UUID
	Data   Data
	Tag    string
	Active bool
	Updater
	Drawer
	Intersector
	IntersectHandler
	OnRemoveHandler
}

type Data interface {
	Register(game *Game, obj *Object)
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

func (o *Object) WithOnRemoveHandler(handler OnRemoveHandler) *Object {
	o.OnRemoveHandler = handler
	return o
}

func (o *Object) WithTag(tag string) *Object {
	o.Tag = tag
	return o
}

func NewObject() *Object {
	return &Object{
		ID:     uuid.New(),
		Active: true,
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
	if o, ok := data.(OnRemoveHandler); ok {
		obj.OnRemoveHandler = o
	}
	return obj
}
