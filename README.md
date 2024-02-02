### service-mock

Provides an easier way to mock services for testing purposes. 
The wrapper on the `mockgen` tool from `go.uber.org/mock`,
gives you a nicer sintax to write mocks

```go
	    ctx := context.TODO()
		a := assert.New(t)

		// you can explicitly define the mock type
		// NewMock[example_mock.MockExampleMockMockRecorder]
		// or let the compiler infer it
		mock := NewMock(
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
```
See more examples [/services_mock_test.go]