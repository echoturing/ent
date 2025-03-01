// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"

	enttask "entgo.io/ent/entc/integration/gremlin/ent/task"
)

// TaskDelete is the builder for deleting a Task entity.
type TaskDelete struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskDelete builder.
func (td *TaskDelete) Where(ps ...predicate.Task) *TaskDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TaskDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, TaskMutation](ctx, td.gremlinExec, td.mutation, td.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TaskDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TaskDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := td.gremlin().Query()
	if err := td.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	td.mutation.done = true
	return res.ReadInt()
}

func (td *TaskDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(enttask.Label)
	for _, p := range td.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// TaskDeleteOne is the builder for deleting a single Task entity.
type TaskDeleteOne struct {
	td *TaskDelete
}

// Exec executes the deletion query.
func (tdo *TaskDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{enttask.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TaskDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}
