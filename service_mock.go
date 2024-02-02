package service_mock

import (
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

type (
	ServiceClient[R any] interface {
		EXPECT() *R
	}

	MockServiceClient[R any, T ServiceClient[R]] struct {
		ServiceClient T
	}

	FnNewClientService[T any] func(*gomock.Controller) T
)

type MockOptions struct {
	Ctx                 interface{}
	AnyTimes            bool
	Times               int
	Input               interface{}
	Return              interface{}
	Call                interface{}
	Error               error
	SingleErrorReturned bool
	DoAndReturn         interface{}
}

// NewMock returns a new mock service client that can be used to mock any service
// client.
//
// The generic type R is the type of the RECORDER returned by the EXPECT method.
// The generic type T is the type of the service client.
//
// Example:
// NewMock[subscriptionv1mock.MockSubscriptionServiceClientMockRecorder](*testing.T, subscriptionv1mock.NewMockSubscriptionServiceClient)
func NewMock[R any, T ServiceClient[R]](
	t *testing.T,
	fn FnNewClientService[T],
) *MockServiceClient[R, T] {
	return &MockServiceClient[R, T]{
		ServiceClient: fn(gomock.NewController(t)),
	}
}

// NewMockWithCtrl returns a new mock service client that can be used to mock any service
// client.

// The generic type R is the type of the RECORDER returned by the EXPECT method.
// The generic type T is the type of the service client.

// Example:
// NewMockWithCtrl[subscriptionv1mock.MockSubscriptionServiceClientMockRecorder](*gomock.Controller, subscriptionv1mock.NewMockSubscriptionServiceClient)
func NewMockWithCtrl[R any, T ServiceClient[R]](
	ctrl *gomock.Controller,
	fn FnNewClientService[T],
) *MockServiceClient[R, T] {
	return &MockServiceClient[R, T]{
		ServiceClient: fn(ctrl),
	}
}

// Recorder returns a struct with methods that can use used as call to mock.
func (m *MockServiceClient[R, T]) Recorder() *R {
	return m.ServiceClient.EXPECT()
}

// Client return the service client to be used as mock on the actual server struct.
func (m *MockServiceClient[R, T]) Client() T {
	return m.ServiceClient
}

// Mock is the way to build the desired "mocked" API by choosing which methods
// are going to be called or not, specifying input arguments and return values.
func (m *MockServiceClient[R, T]) Mock(opts *MockOptions) *MockServiceClient[R, T] {
	if opts.Ctx == nil {
		opts.Ctx = gomock.Any()
	}

	callValue := reflect.ValueOf(opts.Call)
	inputValue := reflect.ValueOf(opts.Input)

	if callValue.Type().Kind() != reflect.Func {
		panic("Call must be a function, example: mock.Recorder().MyAwesomeFunction")
	}

	if opts.DoAndReturn != nil && reflect.ValueOf(opts.DoAndReturn).Type().Kind() != reflect.Func {
		panic(
			"DoAndReturn must be a function, example: func() (string, error) { return \"Hello World\", nil }",
		)
	}

	in := makeInputForCall(reflect.ValueOf(opts.Ctx), callValue, inputValue)
	out := callValue.Call(in)
	c := out[0].Interface().(*gomock.Call)
	setupReturnValues(c, opts)

	return m
}

func makeInputForCall(
	ctx reflect.Value,
	call reflect.Value,
	callInput reflect.Value,
) []reflect.Value {
	if callInput == reflect.ValueOf(nil) {
		var (
			numIn = call.Type().NumIn()
			input = make([]reflect.Value, numIn)
		)

		for i := 1; i < numIn; i++ {
			input[i] = reflect.ValueOf(gomock.Any())
		}

		input[0] = ctx

		return input
	}

	if callInput.Type().Kind() == reflect.Slice || callInput.Type().Kind() == reflect.Array {
		// Here we need to convert the slice to a list of values.
		values := getValuesFromSliceOrArray(callInput)

		// We need to add one to the length because the first value is the context.
		input := make([]reflect.Value, len(values)+1)

		for i, v := range values {
			// We need to add one to the index because the first value is the context.
			input[i+1] = v
		}

		input[0] = ctx

		return input
	}

	numIn := call.Type().NumIn()
	if call.Type().IsVariadic() {
		numIn = numIn - 1
	}

	input := make([]reflect.Value, numIn)
	input[0] = ctx
	input[1] = callInput

	return input
}

func setupReturnValues(mockCall *gomock.Call, opts *MockOptions) {
	if opts.DoAndReturn != nil {
		mockCall.DoAndReturn(
			opts.DoAndReturn,
		)
		setupTimes(mockCall, opts)

		return
	}

	if opts.SingleErrorReturned {
		mockCall.Return(
			opts.Error,
		)
		setupTimes(mockCall, opts)

		return
	}

	rets := append(startReturnValues(opts), opts.Error)

	mockCall.Return(
		rets...,
	)
	setupTimes(mockCall, opts)
}

func setupTimes(mockCall *gomock.Call, opts *MockOptions) {
	if opts.AnyTimes {
		mockCall.AnyTimes()
		return
	}

	mockCall.Times(opts.Times)
}

func startReturnValues(opts *MockOptions) []interface{} {
	if opts.Return != nil && reflect.ValueOf(opts.Return).Type().Kind() == reflect.Slice {
		return append([]interface{}{}, getValuesFromSliceOrArray(reflect.ValueOf(opts.Return)))
	}

	return append([]interface{}{}, opts.Return)
}

func getValuesFromSliceOrArray(v reflect.Value) []reflect.Value {
	values := make([]reflect.Value, v.Len())
	for i := 0; i < v.Len(); i++ {
		// We need to add one to the index because the first value is the context.
		values[i] = v.Index(i)
	}
	return values
}
