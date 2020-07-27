// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.

package workerutil

import (
	"context"
	sqlf "github.com/keegancsmith/sqlf"
	"sync"
)

// MockHandler is a mock implementation of the Handler interface (from the
// package github.com/sourcegraph/sourcegraph/internal/workerutil) used for
// unit testing.
type MockHandler struct { /* all structs must go */ }

// NewMockHandler creates a new mock of the Handler interface. All methods
// return zero values for all results, unless overwritten.
func NewMockHandler() *MockHandler {
	return &MockHandler{
		HandleFunc: &HandlerHandleFunc{
			defaultHook: func(context.Context, Store, Record) error {
				return nil
			},
		},
	}
}

// NewMockHandlerFrom creates a new mock of the MockHandler interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockHandlerFrom(i Handler) *MockHandler {
	return &MockHandler{
		HandleFunc: &HandlerHandleFunc{
			defaultHook: i.Handle,
		},
	}
}

// HandlerHandleFunc describes the behavior when the Handle method of the
// parent MockHandler instance is invoked.
type HandlerHandleFunc struct { /* all structs must go */ }

// Handle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockHandler) Handle(v0 context.Context, v1 Store, v2 Record) error {
	r0 := m.HandleFunc.nextHook()(v0, v1, v2)
	m.HandleFunc.appendCall(HandlerHandleFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Handle method of the
// parent MockHandler instance is invoked and the hook queue is empty.
func (f *HandlerHandleFunc) SetDefaultHook(hook func(context.Context, Store, Record) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Handle method of the parent MockHandler instance inovkes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *HandlerHandleFunc) PushHook(hook func(context.Context, Store, Record) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerHandleFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, Store, Record) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerHandleFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, Store, Record) error {
		return r0
	})
}

func (f *HandlerHandleFunc) nextHook() func(context.Context, Store, Record) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerHandleFunc) appendCall(r0 HandlerHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerHandleFuncCall objects describing
// the invocations of this function.
func (f *HandlerHandleFunc) History() []HandlerHandleFuncCall {
	f.mutex.Lock()
	history := make([]HandlerHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerHandleFuncCall is an object that describes an invocation of method
// Handle on an instance of MockHandler.
type HandlerHandleFuncCall struct { /* all structs must go */ }

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerHandleFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// MockHandlerWithHooks is a mock implementation of the HandlerWithHooks
// interface (from the package
// github.com/sourcegraph/sourcegraph/internal/workerutil) used for unit
// testing.
type MockHandlerWithHooks struct { /* all structs must go */ }

// NewMockHandlerWithHooks creates a new mock of the HandlerWithHooks
// interface. All methods return zero values for all results, unless
// overwritten.
func NewMockHandlerWithHooks() *MockHandlerWithHooks {
	return &MockHandlerWithHooks{
		HandleFunc: &HandlerWithHooksHandleFunc{
			defaultHook: func(context.Context, Store, Record) error {
				return nil
			},
		},
		PostHandleFunc: &HandlerWithHooksPostHandleFunc{
			defaultHook: func(context.Context, Record) {
				return
			},
		},
		PreHandleFunc: &HandlerWithHooksPreHandleFunc{
			defaultHook: func(context.Context, Record) {
				return
			},
		},
	}
}

// NewMockHandlerWithHooksFrom creates a new mock of the
// MockHandlerWithHooks interface. All methods delegate to the given
// implementation, unless overwritten.
func NewMockHandlerWithHooksFrom(i HandlerWithHooks) *MockHandlerWithHooks {
	return &MockHandlerWithHooks{
		HandleFunc: &HandlerWithHooksHandleFunc{
			defaultHook: i.Handle,
		},
		PostHandleFunc: &HandlerWithHooksPostHandleFunc{
			defaultHook: i.PostHandle,
		},
		PreHandleFunc: &HandlerWithHooksPreHandleFunc{
			defaultHook: i.PreHandle,
		},
	}
}

// HandlerWithHooksHandleFunc describes the behavior when the Handle method
// of the parent MockHandlerWithHooks instance is invoked.
type HandlerWithHooksHandleFunc struct { /* all structs must go */ }

// Handle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockHandlerWithHooks) Handle(v0 context.Context, v1 Store, v2 Record) error {
	r0 := m.HandleFunc.nextHook()(v0, v1, v2)
	m.HandleFunc.appendCall(HandlerWithHooksHandleFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Handle method of the
// parent MockHandlerWithHooks instance is invoked and the hook queue is
// empty.
func (f *HandlerWithHooksHandleFunc) SetDefaultHook(hook func(context.Context, Store, Record) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Handle method of the parent MockHandlerWithHooks instance inovkes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *HandlerWithHooksHandleFunc) PushHook(hook func(context.Context, Store, Record) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerWithHooksHandleFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, Store, Record) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerWithHooksHandleFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, Store, Record) error {
		return r0
	})
}

func (f *HandlerWithHooksHandleFunc) nextHook() func(context.Context, Store, Record) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerWithHooksHandleFunc) appendCall(r0 HandlerWithHooksHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerWithHooksHandleFuncCall objects
// describing the invocations of this function.
func (f *HandlerWithHooksHandleFunc) History() []HandlerWithHooksHandleFuncCall {
	f.mutex.Lock()
	history := make([]HandlerWithHooksHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerWithHooksHandleFuncCall is an object that describes an invocation
// of method Handle on an instance of MockHandlerWithHooks.
type HandlerWithHooksHandleFuncCall struct { /* all structs must go */ }

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerWithHooksHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerWithHooksHandleFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// HandlerWithHooksPostHandleFunc describes the behavior when the PostHandle
// method of the parent MockHandlerWithHooks instance is invoked.
type HandlerWithHooksPostHandleFunc struct { /* all structs must go */ }

// PostHandle delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockHandlerWithHooks) PostHandle(v0 context.Context, v1 Record) {
	m.PostHandleFunc.nextHook()(v0, v1)
	m.PostHandleFunc.appendCall(HandlerWithHooksPostHandleFuncCall{v0, v1})
	return
}

// SetDefaultHook sets function that is called when the PostHandle method of
// the parent MockHandlerWithHooks instance is invoked and the hook queue is
// empty.
func (f *HandlerWithHooksPostHandleFunc) SetDefaultHook(hook func(context.Context, Record)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// PostHandle method of the parent MockHandlerWithHooks instance inovkes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *HandlerWithHooksPostHandleFunc) PushHook(hook func(context.Context, Record)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerWithHooksPostHandleFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(context.Context, Record) {
		return
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerWithHooksPostHandleFunc) PushReturn() {
	f.PushHook(func(context.Context, Record) {
		return
	})
}

func (f *HandlerWithHooksPostHandleFunc) nextHook() func(context.Context, Record) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerWithHooksPostHandleFunc) appendCall(r0 HandlerWithHooksPostHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerWithHooksPostHandleFuncCall objects
// describing the invocations of this function.
func (f *HandlerWithHooksPostHandleFunc) History() []HandlerWithHooksPostHandleFuncCall {
	f.mutex.Lock()
	history := make([]HandlerWithHooksPostHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerWithHooksPostHandleFuncCall is an object that describes an
// invocation of method PostHandle on an instance of MockHandlerWithHooks.
type HandlerWithHooksPostHandleFuncCall struct { /* all structs must go */ }

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerWithHooksPostHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerWithHooksPostHandleFuncCall) Results() []interface{} {
	return []interface{}{}
}

// HandlerWithHooksPreHandleFunc describes the behavior when the PreHandle
// method of the parent MockHandlerWithHooks instance is invoked.
type HandlerWithHooksPreHandleFunc struct { /* all structs must go */ }

// PreHandle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockHandlerWithHooks) PreHandle(v0 context.Context, v1 Record) {
	m.PreHandleFunc.nextHook()(v0, v1)
	m.PreHandleFunc.appendCall(HandlerWithHooksPreHandleFuncCall{v0, v1})
	return
}

// SetDefaultHook sets function that is called when the PreHandle method of
// the parent MockHandlerWithHooks instance is invoked and the hook queue is
// empty.
func (f *HandlerWithHooksPreHandleFunc) SetDefaultHook(hook func(context.Context, Record)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// PreHandle method of the parent MockHandlerWithHooks instance inovkes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *HandlerWithHooksPreHandleFunc) PushHook(hook func(context.Context, Record)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerWithHooksPreHandleFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(context.Context, Record) {
		return
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerWithHooksPreHandleFunc) PushReturn() {
	f.PushHook(func(context.Context, Record) {
		return
	})
}

func (f *HandlerWithHooksPreHandleFunc) nextHook() func(context.Context, Record) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerWithHooksPreHandleFunc) appendCall(r0 HandlerWithHooksPreHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerWithHooksPreHandleFuncCall objects
// describing the invocations of this function.
func (f *HandlerWithHooksPreHandleFunc) History() []HandlerWithHooksPreHandleFuncCall {
	f.mutex.Lock()
	history := make([]HandlerWithHooksPreHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerWithHooksPreHandleFuncCall is an object that describes an
// invocation of method PreHandle on an instance of MockHandlerWithHooks.
type HandlerWithHooksPreHandleFuncCall struct { /* all structs must go */ }

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerWithHooksPreHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerWithHooksPreHandleFuncCall) Results() []interface{} {
	return []interface{}{}
}

// MockHandlerWithPreDequeue is a mock implementation of the
// HandlerWithPreDequeue interface (from the package
// github.com/sourcegraph/sourcegraph/internal/workerutil) used for unit
// testing.
type MockHandlerWithPreDequeue struct { /* all structs must go */ }

// NewMockHandlerWithPreDequeue creates a new mock of the
// HandlerWithPreDequeue interface. All methods return zero values for all
// results, unless overwritten.
func NewMockHandlerWithPreDequeue() *MockHandlerWithPreDequeue {
	return &MockHandlerWithPreDequeue{
		HandleFunc: &HandlerWithPreDequeueHandleFunc{
			defaultHook: func(context.Context, Store, Record) error {
				return nil
			},
		},
		PreDequeueFunc: &HandlerWithPreDequeuePreDequeueFunc{
			defaultHook: func(context.Context) (bool, []*sqlf.Query, error) {
				return false, nil, nil
			},
		},
	}
}

// NewMockHandlerWithPreDequeueFrom creates a new mock of the
// MockHandlerWithPreDequeue interface. All methods delegate to the given
// implementation, unless overwritten.
func NewMockHandlerWithPreDequeueFrom(i HandlerWithPreDequeue) *MockHandlerWithPreDequeue {
	return &MockHandlerWithPreDequeue{
		HandleFunc: &HandlerWithPreDequeueHandleFunc{
			defaultHook: i.Handle,
		},
		PreDequeueFunc: &HandlerWithPreDequeuePreDequeueFunc{
			defaultHook: i.PreDequeue,
		},
	}
}

// HandlerWithPreDequeueHandleFunc describes the behavior when the Handle
// method of the parent MockHandlerWithPreDequeue instance is invoked.
type HandlerWithPreDequeueHandleFunc struct { /* all structs must go */ }

// Handle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockHandlerWithPreDequeue) Handle(v0 context.Context, v1 Store, v2 Record) error {
	r0 := m.HandleFunc.nextHook()(v0, v1, v2)
	m.HandleFunc.appendCall(HandlerWithPreDequeueHandleFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Handle method of the
// parent MockHandlerWithPreDequeue instance is invoked and the hook queue
// is empty.
func (f *HandlerWithPreDequeueHandleFunc) SetDefaultHook(hook func(context.Context, Store, Record) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Handle method of the parent MockHandlerWithPreDequeue instance inovkes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *HandlerWithPreDequeueHandleFunc) PushHook(hook func(context.Context, Store, Record) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerWithPreDequeueHandleFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, Store, Record) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerWithPreDequeueHandleFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, Store, Record) error {
		return r0
	})
}

func (f *HandlerWithPreDequeueHandleFunc) nextHook() func(context.Context, Store, Record) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerWithPreDequeueHandleFunc) appendCall(r0 HandlerWithPreDequeueHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerWithPreDequeueHandleFuncCall objects
// describing the invocations of this function.
func (f *HandlerWithPreDequeueHandleFunc) History() []HandlerWithPreDequeueHandleFuncCall {
	f.mutex.Lock()
	history := make([]HandlerWithPreDequeueHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerWithPreDequeueHandleFuncCall is an object that describes an
// invocation of method Handle on an instance of MockHandlerWithPreDequeue.
type HandlerWithPreDequeueHandleFuncCall struct { /* all structs must go */ }

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerWithPreDequeueHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerWithPreDequeueHandleFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// HandlerWithPreDequeuePreDequeueFunc describes the behavior when the
// PreDequeue method of the parent MockHandlerWithPreDequeue instance is
// invoked.
type HandlerWithPreDequeuePreDequeueFunc struct { /* all structs must go */ }

// PreDequeue delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockHandlerWithPreDequeue) PreDequeue(v0 context.Context) (bool, []*sqlf.Query, error) {
	r0, r1, r2 := m.PreDequeueFunc.nextHook()(v0)
	m.PreDequeueFunc.appendCall(HandlerWithPreDequeuePreDequeueFuncCall{v0, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the PreDequeue method of
// the parent MockHandlerWithPreDequeue instance is invoked and the hook
// queue is empty.
func (f *HandlerWithPreDequeuePreDequeueFunc) SetDefaultHook(hook func(context.Context) (bool, []*sqlf.Query, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// PreDequeue method of the parent MockHandlerWithPreDequeue instance
// inovkes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *HandlerWithPreDequeuePreDequeueFunc) PushHook(hook func(context.Context) (bool, []*sqlf.Query, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerWithPreDequeuePreDequeueFunc) SetDefaultReturn(r0 bool, r1 []*sqlf.Query, r2 error) {
	f.SetDefaultHook(func(context.Context) (bool, []*sqlf.Query, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerWithPreDequeuePreDequeueFunc) PushReturn(r0 bool, r1 []*sqlf.Query, r2 error) {
	f.PushHook(func(context.Context) (bool, []*sqlf.Query, error) {
		return r0, r1, r2
	})
}

func (f *HandlerWithPreDequeuePreDequeueFunc) nextHook() func(context.Context) (bool, []*sqlf.Query, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerWithPreDequeuePreDequeueFunc) appendCall(r0 HandlerWithPreDequeuePreDequeueFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerWithPreDequeuePreDequeueFuncCall
// objects describing the invocations of this function.
func (f *HandlerWithPreDequeuePreDequeueFunc) History() []HandlerWithPreDequeuePreDequeueFuncCall {
	f.mutex.Lock()
	history := make([]HandlerWithPreDequeuePreDequeueFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerWithPreDequeuePreDequeueFuncCall is an object that describes an
// invocation of method PreDequeue on an instance of
// MockHandlerWithPreDequeue.
type HandlerWithPreDequeuePreDequeueFuncCall struct { /* all structs must go */ }

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerWithPreDequeuePreDequeueFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerWithPreDequeuePreDequeueFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}
