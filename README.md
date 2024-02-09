# mock_s

## About

This package provides an easier way to mock services for testing purposes. 
The wrapper on the [mockgen](https://github.com/uber-go/mock) gives you a
nicer syntax to write mocks.

## Usage example

```go
package main

import (
    "context"
    "testing"
    
    "github.com/stretchr/testify/assert"
    "github.com/somatech1/mocks"
)

func TestFoo(t *testing.T) {
    ctx := context.TODO()
    a := assert.New(t)

    // You can explicitly define the mock type
    // NewMock[example_mock.MockExampleMockMockRecorder]
    // or let the compiler infer it
    mock := mocks.New(
        t,
        example_mock.NewMockExampleMock,
    )

    expectedInput := "Hello World"
    expectedOutput := "Mocked Output"

    mock.Mock(&mocks.MockOptions{
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
}
```
See more [examples](service_mock_test.go)

## License

[Mozilla Public License 2.0](LICENSE)
