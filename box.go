package box

import (
	"context"
	"reflect"
)

type boxkey struct {
	Name string
}

var self = boxkey{"$$__BOX_KEY__!!"}

type BoxCtx struct {
	context.Context
	m map[any]any
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

func (p *BoxCtx) Add(k, v any) *BoxCtx {
	if k != nil && v != nil {
		p.m[k] = v
	}

	if !reflect.TypeOf(k).Comparable() {
		panic("key is not comparable")
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

func From[V any](ctx context.Context, key any) V {
	v := ctx.Value(key)
	return v.(V)
}

func Self(ctx context.Context) *BoxCtx {
	return From[*BoxCtx](ctx, self)
}
