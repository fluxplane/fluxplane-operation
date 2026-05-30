package operation

import (
	"fmt"
	"sync"
)

// Registry stores executable operations by name.
//
// The registry performs pure lookup only. It does not execute, instantiate,
// wrap, validate, observe, or own operation lifecycle.
type Registry struct {
	mu     sync.RWMutex
	values map[Name]Operation
}

// NewRegistry returns an empty operation registry.
func NewRegistry() *Registry {
	return &Registry{values: map[Name]Operation{}}
}

// Register adds operations to the registry.
func (r *Registry) Register(ops ...Operation) error {
	if r == nil {
		return fmt.Errorf("operation: registry is nil")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.values == nil {
		r.values = map[Name]Operation{}
	}
	for _, op := range ops {
		if op == nil {
			return fmt.Errorf("operation: operation is nil")
		}
		name := op.Spec().Ref.Name
		if name == "" {
			return fmt.Errorf("operation: name is empty")
		}
		if _, exists := r.values[name]; exists {
			return fmt.Errorf("operation: duplicate name %s", name)
		}
		r.values[name] = op
	}
	return nil
}

// Get returns the operation registered under name.
func (r *Registry) Get(name Name) (Operation, bool) {
	if r == nil {
		return nil, false
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	op, ok := r.values[name]
	return op, ok
}

// Resolve returns the operation identified by ref.
func (r *Registry) Resolve(ref Ref) (Operation, bool) {
	return r.Get(ref.Name)
}

// All returns registered operations in unspecified order.
func (r *Registry) All() []Operation {
	if r == nil {
		return nil
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Operation, 0, len(r.values))
	for _, op := range r.values {
		out = append(out, op)
	}
	return out
}
