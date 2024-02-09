package service_mock

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/somatech1/mocks/internal/example"
	example_mock "github.com/somatech1/mocks/internal/example/mock"
)

func TestGetByString(t *testing.T) {
	t.Run("should mock GetByString method", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := "Hello World"
		expectedOutput := "Mocked Output"

		mock.Mock(&MockOptions{
			Ctx:    ctx,
			Call:   mock.Recorder().GetByString,
			Times:  1,
			Input:  expectedInput,
			Return: expectedOutput,
		})

		c := mock.Client()
		output, err := c.GetByString(ctx, expectedInput)

		a.NoError(err)
		a.Equal(output, expectedOutput)
	})

	t.Run("should mock GetByString method and return error", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)
		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := "Hello World"
		expectedError := errors.New("mocked error")

		mock.Mock(&MockOptions{
			Ctx:    ctx,
			Call:   mock.Recorder().GetByString,
			Times:  1,
			Input:  expectedInput,
			Return: "",
			Error:  expectedError,
		})

		c := mock.Client()
		output, err := c.GetByString(ctx, expectedInput)

		a.Error(err)
		a.Equal(err, expectedError)
		a.Equal("", output)
	})
}

func TestGetByInt(t *testing.T) {
	t.Run("should mock GetByInt method", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := 42
		expectedOutput := 100

		mock.Mock(&MockOptions{
			Ctx:    ctx,
			Call:   mock.Recorder().GetByInt,
			Times:  1,
			Input:  expectedInput,
			Return: expectedOutput,
		})

		c := mock.Client()
		output, err := c.GetByInt(ctx, expectedInput)

		a.NoError(err)
		a.Equal(output, expectedOutput)
	})
	t.Run("should mock GetByInt method and return error", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := 1
		expectedOutput := 42

		mock.Mock(&MockOptions{
			Ctx:    ctx,
			Call:   mock.Recorder().GetByInt,
			Times:  1,
			Input:  expectedInput,
			Return: expectedOutput,
		})

		c := mock.Client()
		output, err := c.GetByInt(ctx, expectedInput)

		a.NoError(err)
		a.Equal(output, expectedOutput)
	})
}

func TestGetWithVariadic(t *testing.T) {
	t.Run("should mock GetWithVariadic method", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedOutput := 0

		mock.Mock(&MockOptions{
			Ctx:   ctx,
			Call:  mock.Recorder().GetWithVariadic,
			Times: 1,
			Input: []interface{}{
				"Hello World",
				"Another One",
				"Another One",
			},
			Return: expectedOutput,
		})

		c := mock.Client()
		output, err := c.GetWithVariadic(ctx, "Hello World", "Another One", "Another One")

		a.NoError(err)
		a.Equal(output, 0)
	})
	t.Run("should mock GetWithVariadic method and not pass all arguments", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := "Hello World"
		expectedOutput := 0

		mock.Mock(&MockOptions{
			Ctx:    ctx,
			Call:   mock.Recorder().GetWithVariadic,
			Times:  1,
			Input:  expectedInput,
			Return: expectedOutput,
		})

		c := mock.Client()
		output, err := c.GetWithVariadic(ctx, expectedInput)

		a.NoError(err)
		a.Equal(output, expectedOutput)
	})
}

func TestSingleError(t *testing.T) {
	t.Run("should mock SingleError method", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := "Hello World"
		expectedError := errors.New("mocked error")

		mock.Mock(&MockOptions{
			Ctx:                 ctx,
			Call:                mock.Recorder().SingleError,
			Times:               1,
			Input:               []interface{}{expectedInput},
			SingleErrorReturned: true,
			Error:               expectedError,
		})

		c := mock.Client()
		err := c.SingleError(ctx, expectedInput)

		a.Error(err)
		a.Equal(err, expectedError)
	})
	t.Run("should mock SingleError method and not return an error", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := "Hello World"

		mock.Mock(&MockOptions{
			Ctx:                 ctx,
			Call:                mock.Recorder().SingleError,
			Times:               1,
			Input:               []interface{}{expectedInput},
			SingleErrorReturned: true,
			Error:               nil,
		})

		c := mock.Client()
		err := c.SingleError(ctx, expectedInput)

		a.NoError(err)
	})
}

func TestWithStruct(t *testing.T) {
	t.Run("should mock WithStruct method", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := &example.Example{Id: "1", Value: "Hello World"}
		expectedOutput := &example.Example{Id: "1", Value: "Mocked Output"}

		mock.Mock(&MockOptions{
			Ctx:    ctx,
			Call:   mock.Recorder().WithStruct,
			Times:  1,
			Input:  []interface{}{expectedInput},
			Return: expectedOutput,
		})

		c := mock.Client()
		output, err := c.WithStruct(ctx, expectedInput)

		a.NoError(err)
		a.Equal(output, expectedOutput)
	})
	t.Run("should mock WithStruct method and return nil", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := &example.Example{Id: "1", Value: "Hello World"}

		mock.Mock(&MockOptions{
			Ctx:    ctx,
			Call:   mock.Recorder().WithStruct,
			Times:  1,
			Input:  []interface{}{expectedInput},
			Return: nil,
		})

		c := mock.Client()
		output, err := c.WithStruct(ctx, expectedInput)

		a.NoError(err)
		a.Nil(output)
	})
}

func TestWithDoAndReturn(t *testing.T) {
	t.Run("should mock WithDoAndReturn method", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := &example.Example{Id: "1", Value: "Hello World"}
		expectedOutput := &example.Example{Id: "2", Value: "Mocked Output"}

		mock.Mock(&MockOptions{
			Ctx:   ctx,
			Call:  mock.Recorder().WithDoAndReturn,
			Times: 1,
			Input: []interface{}{expectedInput},
			DoAndReturn: func(ctx context.Context, in *example.Example) (*example.Example, error) {
				return expectedOutput, nil
			},
			Error: nil,
		})

		c := mock.Client()
		output, err := c.WithDoAndReturn(ctx, expectedInput)

		a.NoError(err)
		a.Equal(output, expectedOutput)
	})
	t.Run("should mock WithDoAndReturn method and return the mocked value", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := &example.Example{Id: "1", Value: "Hello World"}
		expectedOutput := &example.Example{Id: "2", Value: "Mocked Hello World"}

		mock.Mock(&MockOptions{
			Ctx:   ctx,
			Call:  mock.Recorder().WithDoAndReturn,
			Times: 1,
			Input: []interface{}{expectedInput},
			DoAndReturn: func(ctx context.Context, in *example.Example) (*example.Example, error) {
				return expectedOutput, nil
			},
			Error: nil,
		})

		c := mock.Client()
		output, err := c.WithDoAndReturn(ctx, expectedInput)

		a.NoError(err)
		a.Equal(output, expectedOutput)
	})
}

func TestAny(t *testing.T) {
	t.Run("should mock Any method and return the mocked value", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := &example.Example{Id: "1", Value: "Hello World"}
		expectedOutput := &example.Example{Id: "2", Value: "Mocked Hello World"}

		mock.Mock(&MockOptions{
			Ctx:    ctx,
			Call:   mock.Recorder().Any,
			Times:  1,
			Input:  []interface{}{expectedInput, "Hello world", "Another One"},
			Return: expectedOutput,
		})

		c := mock.Client()
		output, err := c.Any(ctx, expectedInput, "Hello world", "Another One")

		a.NoError(err)
		a.Equal(output, expectedOutput)
	})

	t.Run("should mock Any method and return nil", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		expectedInput := &example.Example{Id: "1", Value: "Hello World"}

		mock.Mock(&MockOptions{
			Ctx:    ctx,
			Call:   mock.Recorder().Any,
			Times:  1,
			Input:  []interface{}{expectedInput},
			Return: nil,
		})

		c := mock.Client()
		output, err := c.Any(ctx, expectedInput)

		a.NoError(err)
		a.Nil(output)
	})

	t.Run("should mock Any method and return nil if no inputs are passed", func(t *testing.T) {
		ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// New[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := New(
			t,
			example_mock.NewMockExampleMock,
		)

		mock.Mock(&MockOptions{
			Ctx:    ctx,
			Call:   mock.Recorder().Any,
			Times:  1,
			Return: nil,
		})

		c := mock.Client()
		output, err := c.Any(ctx, nil, "Hello world", "Another One")

		a.NoError(err)
		a.Nil(output)
	})
}
