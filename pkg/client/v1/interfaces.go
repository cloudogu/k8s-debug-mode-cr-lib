package v1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	v1 "github.com/cloudogu/k8s-debug-mode-cr-lib/api/v1"
)

type DebugModeInterface interface {
	// Create takes the representation of a debugMode and creates it.  Returns the server's representation of the debugMode, and an error, if there is any.
	Create(ctx context.Context, debugMode *v1., opts metav1.CreateOptions) (*v1.debugMode, error)
	// Update takes the representation of a debugMode and updates it. Returns the server's representation of the debugMode, and an error, if there is any.
	Update(ctx context.Context, debugMode *v1.debugMode, opts metav1.UpdateOptions) (*v1.debugMode, error)
	// UpdateStatus was generated because the type contains a Status member.
	UpdateStatus(ctx context.Context, debugMode *v1.debugMode, opts metav1.UpdateOptions) (*v1.debugMode, error)
	// UpdateStatusCreating sets the status of the debugMode to "creating".
	UpdateStatusCreating(ctx context.Context, debugMode *v1.debugMode) (*v1.debugMode, error)
	// UpdateStatusCreated sets the status of the debugMode to "created".
	UpdateStatusCreated(ctx context.Context, debugMode *v1.debugMode) (*v1.debugMode, error)
	// UpdateStatusDeleting sets the status of the debugMode to "deleting".
	UpdateStatusDeleting(ctx context.Context, debugMode *v1.debugMode) (*v1.debugMode, error)
	// UpdateStatusFailed sets the status of the debugMode to "failed".
	UpdateStatusFailed(ctx context.Context, debugMode *v1.debugMode) (*v1.debugMode, error)
}