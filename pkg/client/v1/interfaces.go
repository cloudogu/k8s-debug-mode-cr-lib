package v1

import (
	"context"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "github.com/cloudogu/k8s-debug-mode-cr-lib/api/v1"
)

type DebugModeV1Interface interface {
	DebugMode(namespace string) DebugModeInterface
}

type DebugModeInterface interface {
	// Create takes the representation of a debugMode and creates it.  Returns the server's representation of the debugMode, and an error, if there is any.
	Create(ctx context.Context, debugMode *v1.DebugMode, opts metav1.CreateOptions) (result *v1.DebugMode, err error)
	// Update takes the representation of a debugMode and updates it. Returns the server's representation of the debugMode, and an error, if there is any.
	Update(ctx context.Context, debugMode *v1.DebugMode, opts metav1.UpdateOptions) (result *v1.DebugMode, err error)
	// UpdateStatus was generated because the type contains a Status member.
	UpdateStatus(ctx context.Context, debugMode *v1.DebugMode, opts metav1.UpdateOptions) (result *v1.DebugMode, err error)
	// UpdateStatusDebugModeSet sets the status of the debugMode to "SetDebugMode".
	UpdateStatusDebugModeSet(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error)
	// UpdateStatusWaitForRollback sets the status of the debugMode to "WaitForRollback".
	UpdateStatusWaitForRollback(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error)
	// UpdateStatusRollback sets the status of the debugMode to "Rollback".
	UpdateStatusRollback(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error)
	// UpdateStatusCompleted sets the status of the debugMode to "Completed".
	UpdateStatusCompleted(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error)
	// UpdateStatusFailed sets the status of the debugMode to "Failed".
	UpdateStatusFailed(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error)
	// Delete takes name of the debugMode and deletes it. Returns an error if one occurs.
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	// Get takes name of the debugMode, and returns the corresponding debugMode object, and an error if there is any.
	Get(ctx context.Context, name string, opts metav1.GetOptions) (result *v1.DebugMode, err error)
	// Watch returns a watch.Interface that watches the requested debugModes.
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	// Patch applies the patch and returns the patched debugMode.
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DebugMode, err error)
	// AddFinalizer adds the given finalizer to the debugMode.
	AddFinalizer(ctx context.Context, debugMode *v1.DebugMode, finalizer string) (*v1.DebugMode, error)
	// RemoveFinalizer removes the given finalizer to the debugMode.
	RemoveFinalizer(ctx context.Context, debugMode *v1.DebugMode, finalizer string) (*v1.DebugMode, error)
	// AddOrUpdateLogLevelsSet sets the condition for the debugMode.
	AddOrUpdateLogLevelsSet(ctx context.Context, debugMode *v1.DebugMode, set bool, msg string, reason string) (*v1.DebugMode, error)
}
