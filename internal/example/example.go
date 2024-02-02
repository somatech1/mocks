package example

import "context"

type Example struct {
	Id    string
	Value string
}

type ExampleMock interface {
	GetByString(ctx context.Context, str string) (string, error)
	GetByInt(ctx context.Context, i int) (int, error)
	GetWithVariadic(ctx context.Context, id string, options ...string) (int, error)
	SingleError(ctx context.Context, id string, options ...string) error
	WithStruct(ctx context.Context, in *Example) (*Example, error)
	WithDoAndReturn(ctx context.Context, in *Example) (*Example, error)
	Any(ctx context.Context, in *Example, options ...string) (*Example, error)
}
