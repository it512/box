package box

import (
	"context"
	"reflect"
)

type boxkey struct {
	Name string
}

var self = boxkey{"$$__bOx_KeY__!!"}

type boxCtx struct {
	context.Context
	m map[any]any
}

func new(parent context.Context) *boxCtx {
	return &boxCtx{
		Context: parent,
		m:       make(map[any]any),
	}
}

func (boxCtx) String() string {
	return "context.Box"
}

func (p *boxCtx) put(k, v any) *boxCtx {
	if k != nil && v != nil {
		p.m[k] = v
	}

	if !reflect.TypeOf(k).Comparable() {
		panic("key is not comparable")
	}

	return p
}

func (p *boxCtx) Value(key any) any {
	if key == self {
		return p
	}

	if v, ok := p.m[key]; ok {
		return v
	}

	return p.Context.Value(key)
}

func From[V any](ctx context.Context, key any) (v V, ok bool) {
	val := ctx.Value(key)
	if val == nil {
		return
	}
	return val.(V), true
}

func WithValue(parent context.Context, key, val any) context.Context {
	if box, ok := From[*boxCtx](parent, self); ok {
		box.put(key, val)
		return parent
	}

	b := new(parent)
	b.put(key, val)
	return b
}
