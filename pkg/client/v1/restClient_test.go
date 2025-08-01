package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/api/meta"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	v1 "github.com/cloudogu/k8s-debug-mode-cr-lib/api/v1"
)

var testCtx = context.Background()

func Test_DebugModeClient_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "GET", request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/testDebugMode", request.URL.Path)
			assert.Equal(t, http.NoBody, request.Body)

			writer.Header().Add("content-type", "application/json")
			DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "testDebugMode", Namespace: "test"}}
			DebugModeBytes, err := json.Marshal(DebugMode)
			require.NoError(t, err)
			_, err = writer.Write(DebugModeBytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		_, err = sClient.Get(testCtx, "testDebugMode", metav1.GetOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_DebugModeClient_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPost, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdDebugMode := &v1.DebugMode{}
			require.NoError(t, json.Unmarshal(bytes, createdDebugMode))
			assert.Equal(t, "tocreate", createdDebugMode.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		_, err = sClient.Create(testCtx, DebugMode, metav1.CreateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_DebugModeClient_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/tocreate", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdDebugMode := &v1.DebugMode{}
			require.NoError(t, json.Unmarshal(bytes, createdDebugMode))
			assert.Equal(t, "tocreate", createdDebugMode.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		_, err = sClient.Update(testCtx, DebugMode, metav1.UpdateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_DebugModeClient_UpdateStatus(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/tocreate/status", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdDebugMode := &v1.DebugMode{}
			require.NoError(t, json.Unmarshal(bytes, createdDebugMode))
			assert.Equal(t, "tocreate", createdDebugMode.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		_, err = sClient.UpdateStatus(testCtx, DebugMode, metav1.UpdateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_DebugModeClient_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodDelete, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/testDebugMode", request.URL.Path)

			writer.Header().Add("content-type", "application/json")
			writer.WriteHeader(200)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		err = sClient.Delete(testCtx, "testDebugMode", metav1.DeleteOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_DebugModeClient_Patch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPatch, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/testDebugMode", request.URL.Path)
			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)
			assert.Equal(t, []byte("test"), bytes)
			result, err := json.Marshal(v1.DebugMode{})
			require.NoError(t, err)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(result)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		patchData := []byte("test")

		// when
		_, err = sClient.Patch(testCtx, "testDebugMode", types.JSONPatchType, patchData, metav1.PatchOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_DebugModeClient_Watch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "GET", request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode", request.URL.Path)
			assert.Equal(t, http.NoBody, request.Body)
			assert.Equal(t, "labelSelector=test&timeout=5s&timeoutSeconds=5&watch=true", request.URL.RawQuery)

			writer.Header().Add("content-type", "application/json")
			_, err := writer.Write([]byte("egal"))
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		timeout := int64(5)

		// when
		_, err = sClient.Watch(testCtx, metav1.ListOptions{LabelSelector: "test", TimeoutSeconds: &timeout})

		// then
		require.NoError(t, err)
	})
}

func Test_DebugModeClient_UpdateStatusDebugModeSet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusSet, false, false)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusDebugModeSet(testCtx, DebugMode)

		// then
		require.NoError(t, err)
	})

	t.Run("success with retry", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusSet, true, false)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusDebugModeSet(testCtx, DebugMode)

		// then
		require.NoError(t, err)
	})

	t.Run("fail on get DebugMode", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusSet, false, true)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusDebugModeSet(testCtx, DebugMode)

		// then
		require.Error(t, err)
		require.ErrorContains(t, err, "an error on the server (\"\") has prevented the request from succeeding (get DebugMode.k8s.cloudogu.com myDebugMode)")
	})
}

func Test_DebugModeClient_UpdateStatusWaitForRollback(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusWaitForRollback, false, false)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusWaitForRollover(testCtx, DebugMode)

		// then
		require.NoError(t, err)
	})
	t.Run("success with retry", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusWaitForRollback, true, false)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusWaitForRollover(testCtx, DebugMode)

		// then
		require.NoError(t, err)
	})
	t.Run("fail on get DebugMode", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusWaitForRollback, false, true)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusWaitForRollover(testCtx, DebugMode)

		// then
		require.Error(t, err)
		require.ErrorContains(t, err, "an error on the server (\"\") has prevented the request from succeeding (get DebugMode.k8s.cloudogu.com myDebugMode)")
	})
}

func Test_DebugModeClient_UpdateStatusRollback(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusRollback, false, false)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusDebugModeSet(testCtx, DebugMode)

		// then
		require.NoError(t, err)
	})
	t.Run("success with retry", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusRollback, true, false)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusDebugModeSet(testCtx, DebugMode)

		// then
		require.NoError(t, err)
	})
	t.Run("fail on get DebugMode", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusWaitForRollback, false, true)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusDebugModeSet(testCtx, DebugMode)

		// then
		require.Error(t, err)
		require.ErrorContains(t, err, "an error on the server (\"\") has prevented the request from succeeding (get DebugMode.k8s.cloudogu.com myDebugMode)")
	})
}

func Test_DebugModeClient_UpdateStatusCompleted(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusCompleted, false, false)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusCompleted(testCtx, DebugMode)

		// then
		require.NoError(t, err)
	})
	t.Run("success with retry", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusCompleted, true, false)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusCompleted(testCtx, DebugMode)

		// then
		require.NoError(t, err)
	})
	t.Run("fail on get DebugMode", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		mockClient := mockClientForStatusUpdates(t, DebugMode, v1.DebugModeStatusCompleted, false, true)
		sClient := mockClient.DebugMode("test")

		// when
		_, err := sClient.UpdateStatusCompleted(testCtx, DebugMode)

		// then
		require.Error(t, err)
		require.ErrorContains(t, err, "an error on the server (\"\") has prevented the request from succeeding (get DebugMode.k8s.cloudogu.com myDebugMode)")
	})
}

func Test_DebugModeClient_AddFinalizer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/myDebugMode", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdDebugMode := &v1.DebugMode{}
			require.NoError(t, json.Unmarshal(bytes, createdDebugMode))
			assert.Equal(t, "myDebugMode", createdDebugMode.Name)
			assert.Len(t, createdDebugMode.Finalizers, 1)
			assert.Equal(t, "myFinalizer", createdDebugMode.Finalizers[0])

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		_, err = sClient.AddFinalizer(testCtx, DebugMode, "myFinalizer")

		// then
		require.NoError(t, err)
	})

	t.Run("should fail to set finalizer on client error", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/myDebugMode", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdDebugMode := &v1.DebugMode{}
			require.NoError(t, json.Unmarshal(bytes, createdDebugMode))
			assert.Equal(t, "myDebugMode", createdDebugMode.Name)
			assert.Len(t, createdDebugMode.Finalizers, 1)
			assert.Equal(t, "myFinalizer", createdDebugMode.Finalizers[0])

			writer.WriteHeader(500)
			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		_, err = sClient.AddFinalizer(testCtx, DebugMode, "myFinalizer")

		// then
		require.Error(t, err)
		assert.ErrorContains(t, err, "failed to add finalizer myFinalizer to DebugMode:")
	})
}

func Test_DebugModeClient_RemoveFinalizer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		controllerutil.AddFinalizer(DebugMode, "finalizer1")
		controllerutil.AddFinalizer(DebugMode, "finalizer2")

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/myDebugMode", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdDebugMode := &v1.DebugMode{}
			require.NoError(t, json.Unmarshal(bytes, createdDebugMode))
			assert.Equal(t, "myDebugMode", createdDebugMode.Name)
			assert.Len(t, createdDebugMode.Finalizers, 1)
			assert.Equal(t, "finalizer2", createdDebugMode.Finalizers[0])

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		_, err = sClient.RemoveFinalizer(testCtx, DebugMode, "finalizer1")

		// then
		require.NoError(t, err)
	})

	t.Run("should fail to set finalizer on client error", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}
		controllerutil.AddFinalizer(DebugMode, "finalizer1")
		controllerutil.AddFinalizer(DebugMode, "finalizer2")

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/myDebugMode", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdDebugMode := &v1.DebugMode{}
			require.NoError(t, json.Unmarshal(bytes, createdDebugMode))
			assert.Equal(t, "myDebugMode", createdDebugMode.Name)
			assert.Len(t, createdDebugMode.Finalizers, 1)
			assert.Equal(t, "finalizer1", createdDebugMode.Finalizers[0])

			writer.WriteHeader(500)
			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		_, err = sClient.RemoveFinalizer(testCtx, DebugMode, "finalizer2")

		// then
		require.Error(t, err)
		assert.ErrorContains(t, err, "failed to remove finalizer finalizer2 from DebugMode")
	})
}

func Test_DebugModeClient_AddOrUpdateLogLevelsSetCondition(t *testing.T) {
	t.Run("success condition set to false", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/myDebugMode", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdDebugMode := &v1.DebugMode{}
			require.NoError(t, json.Unmarshal(bytes, createdDebugMode))

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)

			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}

		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		_, err = sClient.AddOrUpdateLogLevelsSet(testCtx, DebugMode, false)

		// then
		require.NoError(t, err)
		require.Len(t, DebugMode.Status.Conditions, 1)
		require.Equal(t, metav1.ConditionFalse, meta.FindStatusCondition(DebugMode.Status.Conditions, v1.ConditionLogLevelSet).Status)
	})

	t.Run("success condition set to true", func(t *testing.T) {
		// given
		DebugMode := &v1.DebugMode{ObjectMeta: metav1.ObjectMeta{Name: "myDebugMode", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/myDebugMode", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdDebugMode := &v1.DebugMode{}
			require.NoError(t, json.Unmarshal(bytes, createdDebugMode))

			writer.WriteHeader(500)
			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}

		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.DebugMode("test")

		// when
		_, err = sClient.AddOrUpdateLogLevelsSet(testCtx, DebugMode, true)

		// then
		require.Error(t, err)
		require.Len(t, DebugMode.Status.Conditions, 1)
		require.Equal(t, metav1.ConditionTrue, meta.FindStatusCondition(DebugMode.Status.Conditions, v1.ConditionLogLevelSet).Status)
	})
}

func mockClientForStatusUpdates(t *testing.T, expectedDebugMode *v1.DebugMode, expectedStatus v1.StatusPhase, withRetry bool, failOnGetDebugMode bool) DebugModeV1Interface {

	failGetDebugMode := func(writer http.ResponseWriter, request *http.Request) {
		assert.Equal(t, http.MethodGet, request.Method)
		assert.Equal(t, fmt.Sprintf("/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/%s", expectedDebugMode.Name), request.URL.Path)

		writer.WriteHeader(500)
	}

	assertGetDebugModeRequest := func(writer http.ResponseWriter, request *http.Request) {
		assert.Equal(t, http.MethodGet, request.Method)
		assert.Equal(t, fmt.Sprintf("/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/%s", expectedDebugMode.Name), request.URL.Path)

		DebugModeJson, err := json.Marshal(expectedDebugMode)
		require.NoError(t, err)

		writer.Header().Add("content-type", "application/json")
		_, err = writer.Write(DebugModeJson)
		require.NoError(t, err)
	}

	assertUpdateStatusRequest := func(writer http.ResponseWriter, request *http.Request) {
		assert.Equal(t, http.MethodPut, request.Method)
		assert.Equal(t, fmt.Sprintf("/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/%s/status", expectedDebugMode.Name), request.URL.Path)

		bytes, err := io.ReadAll(request.Body)
		require.NoError(t, err)

		createdDebugMode := &v1.DebugMode{}
		require.NoError(t, json.Unmarshal(bytes, createdDebugMode))
		assert.Equal(t, expectedDebugMode.Name, createdDebugMode.Name)
		assert.Equal(t, expectedStatus, createdDebugMode.Status.Phase)

		writer.Header().Add("content-type", "application/json")
		_, err = writer.Write(bytes)
		require.NoError(t, err)
	}

	conflictUpdateStatusRequest := func(writer http.ResponseWriter, request *http.Request) {
		assert.Equal(t, http.MethodPut, request.Method)
		assert.Equal(t, fmt.Sprintf("/apis/k8s.cloudogu.com/v1/namespaces/test/debugmode/%s/status", expectedDebugMode.Name), request.URL.Path)

		writer.WriteHeader(409)
	}

	var requestAssertions []func(writer http.ResponseWriter, request *http.Request)

	if failOnGetDebugMode {
		requestAssertions = []func(writer http.ResponseWriter, request *http.Request){
			failGetDebugMode,
		}
	} else if withRetry {
		requestAssertions = []func(writer http.ResponseWriter, request *http.Request){
			assertGetDebugModeRequest,
			conflictUpdateStatusRequest,
			assertGetDebugModeRequest,
			assertUpdateStatusRequest,
		}
	} else {
		requestAssertions = []func(writer http.ResponseWriter, request *http.Request){
			assertGetDebugModeRequest,
			assertUpdateStatusRequest,
		}
	}

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		assertRequestFunc := requestAssertions[0]
		requestAssertions = requestAssertions[1:]

		assertRequestFunc(writer, request)
	}))

	config := rest.Config{
		Host: server.URL,
	}
	client, err := NewForConfig(&config)
	require.NoError(t, err)
	return client
}
