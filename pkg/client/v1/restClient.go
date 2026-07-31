package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudogu/retry-lib/retry"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"

	v1 "github.com/cloudogu/k8s-debug-mode-cr-lib/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type debugModeClient struct {
	client rest.Interface
	ns     string
}

func (client *debugModeClient) Create(ctx context.Context, debugMode *v1.DebugMode, opts metav1.CreateOptions) (result *v1.DebugMode, err error) {
	result = &v1.DebugMode{}
	err = client.client.Post().
		Namespace(client.ns).
		Resource("debugmodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(debugMode).
		Do(ctx).
		Into(result)
	return
}

func (client *debugModeClient) Update(ctx context.Context, debugMode *v1.DebugMode, opts metav1.UpdateOptions) (result *v1.DebugMode, err error) {
	result = &v1.DebugMode{}
	err = client.client.Put().
		Namespace(client.ns).
		Resource("debugmodes").
		Name(debugMode.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(debugMode).
		Do(ctx).
		Into(result)
	return
}

func (client *debugModeClient) UpdateStatus(ctx context.Context, debugMode *v1.DebugMode, opts metav1.UpdateOptions) (result *v1.DebugMode, err error) {
	result = &v1.DebugMode{}
	err = client.client.Put().
		Namespace(client.ns).
		Resource("debugmodes").
		Name(debugMode.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(debugMode).
		Do(ctx).
		Into(result)
	return
}

func (client *debugModeClient) UpdateStatusCompleted(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error) {
	debugMode, err := client.updateStatusWithRetry(ctx, debugMode, v1.DebugModeStatusCompleted)
	if err != nil {
		return nil, err
	}

	return debugMode, nil
}

func (client *debugModeClient) UpdateStatusDebugModeSet(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error) {
	debugMode, err := client.updateStatusWithRetry(ctx, debugMode, v1.DebugModeStatusSet)
	if err != nil {
		return nil, err
	}

	return debugMode, nil
}

func (client *debugModeClient) UpdateStatusRollback(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error) {
	debugMode, err := client.updateStatusWithRetry(ctx, debugMode, v1.DebugModeStatusRollback)
	if err != nil {
		return nil, err
	}

	return debugMode, nil
}

func (client *debugModeClient) UpdateStatusWaitForRollback(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error) {
	debugMode, err := client.updateStatusWithRetry(ctx, debugMode, v1.DebugModeStatusWaitForRollback)
	if err != nil {
		return nil, err
	}

	return debugMode, nil
}

func (client *debugModeClient) UpdateStatusFailed(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error) {
	debugMode, err := client.updateStatusWithRetry(ctx, debugMode, v1.DebugModeStatusFailed)
	if err != nil {
		return nil, err
	}

	return debugMode, nil
}

func (client *debugModeClient) updateStatusWithRetry(ctx context.Context, debugMode *v1.DebugMode, targetStatus v1.StatusPhase) (*v1.DebugMode, error) {
	var resultDebugMode *v1.DebugMode
	err := retry.OnConflict(func() error {
		updatedDebugMode, err := client.Get(ctx, debugMode.GetName(), metav1.GetOptions{})
		if err != nil {
			return err
		}

		// do not overwrite the whole status, so we do not lose other values from the Status object
		// esp. a potentially set requeue time
		updatedDebugMode.Status.Phase = targetStatus
		resultDebugMode, err = client.UpdateStatus(ctx, updatedDebugMode, metav1.UpdateOptions{})
		return err
	})

	return resultDebugMode, err
}

func (client *debugModeClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return client.client.Delete().
		Namespace(client.ns).
		Resource("debugmodes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (client *debugModeClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (result *v1.DebugMode, err error) {
	result = &v1.DebugMode{}
	err = client.client.Get().
		Namespace(client.ns).
		Resource("debugmodes").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (client *debugModeClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return client.client.Get().
		Namespace(client.ns).
		Resource("debugmodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (client *debugModeClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DebugMode, err error) {
	result = &v1.DebugMode{}
	err = client.client.Patch(pt).
		Namespace(client.ns).
		Resource("debugmodes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

func (client *debugModeClient) AddFinalizer(ctx context.Context, debugMode *v1.DebugMode, finalizer string) (*v1.DebugMode, error) {
	controllerutil.AddFinalizer(debugMode, finalizer)
	result, err := client.Update(ctx, debugMode, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to add finalizer %s to debugMode: %w", finalizer, err)
	}

	return result, nil

}

func (client *debugModeClient) RemoveFinalizer(ctx context.Context, debugMode *v1.DebugMode, finalizer string) (*v1.DebugMode, error) {
	controllerutil.RemoveFinalizer(debugMode, finalizer)
	result, err := client.Update(ctx, debugMode, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to remove finalizer %s from debugMode: %w", finalizer, err)
	}

	return result, err
}

func (client *debugModeClient) AddOrUpdateLogLevelsSet(ctx context.Context, debugMode *v1.DebugMode, set bool, msg string, reason string) (*v1.DebugMode, error) {
	if reason == "" {
		reason = "Initialized"
	}

	if msg == "" {
		msg = "Condition set to initialized"
	}

	return client.addOrUpdateCondition(ctx, v1.ConditionLogLevelSet, debugMode, set, msg, reason)
}

func (client *debugModeClient) AddOrUpdateFailed(ctx context.Context, debugMode *v1.DebugMode, set bool, msg string, reason string) (*v1.DebugMode, error) {
	if reason == "" {
		reason = "Error"
	}

	if msg == "" {
		msg = "Error handling debug-mode"
	}

	return client.addOrUpdateCondition(ctx, v1.ConditionFailed, debugMode, set, msg, reason)
}

func (client *debugModeClient) RemoveFailed(ctx context.Context, debugMode *v1.DebugMode) (*v1.DebugMode, error) {
	return client.removeCondition(ctx, v1.ConditionFailed, debugMode)
}

func (client *debugModeClient) addOrUpdateCondition(ctx context.Context, condition string, debugMode *v1.DebugMode, set bool, msg string, reason string) (*v1.DebugMode, error) {
	conditionStatus := metav1.ConditionFalse
	if set == true {
		conditionStatus = metav1.ConditionTrue
	}

	newCondition := metav1.Condition{
		Type:               condition,
		Status:             conditionStatus,
		Reason:             reason,
		Message:            msg,
		LastTransitionTime: metav1.Now(),
	}

	_ = meta.SetStatusCondition(&debugMode.Status.Conditions, newCondition)
	result, err := client.UpdateStatus(ctx, debugMode, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to add or update condition %s to debugMode: %w", newCondition.Type, err)
	}

	return result, nil
}

func (client *debugModeClient) removeCondition(ctx context.Context, conditionType string, debugMode *v1.DebugMode) (*v1.DebugMode, error) {

	conditions := debugMode.Status.Conditions
	resultConditions := conditions[:0]
	removed := false
	for _, c := range conditions {
		if c.Type == conditionType {
			removed = true
			continue
		}
		resultConditions = append(resultConditions, c)
	}

	if !removed {
		return debugMode, nil
	}

	debugMode.Status.Conditions = resultConditions

	result, err := client.UpdateStatus(ctx, debugMode, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to remove condition %s from debugMode: %w", conditionType, err)
	}

	return result, nil
}
