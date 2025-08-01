package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatusPhase string

const (
	DebugModeStatusSet             StatusPhase = "SetDebugMode"
	DebugModeStatusWaitForRollback StatusPhase = "WaitForRollback"
	DebugModeStatusRollback        StatusPhase = "Rollback"
	DebugModeStatusCompleted       StatusPhase = "Completed"
)

// DebugModeSpec defines the desired state of DebugMode
type DebugModeSpec struct {
	DeactivateTimestamp metav1.Time `json:"deactivateTimestamp,omitempty"`
	TargetLogLevel      string      `json:"targetLogLevel,omitempty"`
}

const (
	ConditionLogLevelSet string = "LogLevelsSet"
)

// DebugModeStatus defines the observed state of DebugMode.
type DebugModeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Phase defines the current general state the resource is in.
	Phase StatusPhase `json:"phase,omitempty"`
	// Errors contains error messages that accumulated during execution.
	Errors string `json:"errors,omitempty"`
	// Conditions are used to influence the Phase
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DebugMode is the Schema for the debugmodes API
type DebugMode struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of DebugMode
	// +required
	Spec DebugModeSpec `json:"spec"`

	// status defines the observed state of DebugMode
	// +optional
	Status DebugModeStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

func init() {
	SchemeBuilder.Register(&DebugMode{})
}
