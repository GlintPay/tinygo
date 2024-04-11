// Stubs for the runtime/trace package
package trace

import (
	"context"
	"errors"
	"io"
)

func Start(w io.Writer) error {
	return errors.New("not implemented")
}

func Stop() {}

// IsEnabled reports whether tracing is enabled.
// The information is advisory only. The tracing status
// may have changed by the time this function returns.
func IsEnabled() bool {
	return false
}

type Task struct {
	id uint64
}

func NewTask(ctx context.Context, _ string) (context.Context, *Task) {
	return ctx, &Task{}
}

func (t *Task) End() {
	// NOOP
}
