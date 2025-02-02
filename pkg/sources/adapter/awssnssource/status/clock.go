/*
Copyright 2021 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package status

import (
	"context"

	"knative.dev/pkg/apis"
)

type clockKey struct{}

// Clock can override the timestamp returned by time.Now().
type Clock interface {
	Now() apis.VolatileTime
}

// WithClock returns a copy of the parent context in which the value
// associated with the clock key is the given Clock.
func WithClock(ctx context.Context, c Clock) context.Context {
	return context.WithValue(ctx, clockKey{}, c)
}

// ClockFromContext returns the Clock stored in the context.
func ClockFromContext(ctx context.Context) Clock {
	if c, ok := ctx.Value(clockKey{}).(Clock); ok {
		return c
	}
	return nil
}
