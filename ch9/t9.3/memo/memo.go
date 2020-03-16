package memo

import (
	"fmt"
	"sync"
)

//!+Func

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	sync.RWMutex
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
	done     <-chan struct{}
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

//!-get

//!+monitor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	cancelled := make(chan string, 10)
	for req := range memo.requests {
		select {
		case cancelledKey := <-cancelled:
			// TODO simultaneous cancelling
			cache[cancelledKey] = nil
		default:
		}
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{}), res: result{value: ""}}
			cache[req.key] = e
			go e.call(f, req.key, req.done, cancelled) // call f(key)
		}
		go e.deliver(req.response, req.key, req.done, cancelled)
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}, cancelled chan<- string) {
	ready := make(chan result)
	go func() {
		// Evaluate the function.
		value, err := f(key)
		ready <- result{value: value, err: err}
	}()
	select {
	case <-done:
		e.res.err = fmt.Errorf("cancelled %s in call", key)
		cancelled <- key
	case res := <-ready:
		e.res.value, e.res.err = res.value, res.err
	}
	// Broadcast the ready condition.
	defer close(e.ready)
}

func (e *entry) deliver(response chan<- result, key string, done <-chan struct{}, cancelled chan<- string) {
	select {
	case <-done:
		response <- result{err: fmt.Errorf("cancelled in deliver")}
		cancelled <- key
	case <-e.ready:
		// Wait for the ready condition.
		// Send the result to the client.
		response <- e.res
	}
}

//!-monitor
