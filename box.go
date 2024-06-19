package box

import (
	"context"
)

type BoxKey struct {
	Key string
}

var self = &BoxKey{"$$__bOx_KeY__!!"}

type BoxCtx struct {
	context.Context
	m map[any]any
}

func Default() *BoxCtx {
	return New(context.Background())
}

func New(parent context.Context) *BoxCtx {
	return &BoxCtx{
		Context: parent,
		m:       make(map[any]any),
	}
}

func (BoxCtx) String() string {
	return "context.Box"
}

func (p *BoxCtx) Put(k, v any) *BoxCtx {
	if k != nil && v != nil {
		p.m[k] = v
	}
	return p
}

func (p *BoxCtx) Value(key any) any {
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

func MustFrom[V any](ctx context.Context, key any) V {
	if v, ok := From[V](ctx, key); ok {
		return v
	}
	panic("key is not exsit")
}

func WithValue(parent context.Context, key, val any) context.Context {
	if box, ok := From[*BoxCtx](parent, self); ok {
		box.Put(key, val)
		return parent
	}

	b := New(parent)
	b.Put(key, val)
	return b
}
