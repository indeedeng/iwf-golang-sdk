package iwf

// Workflow is the interface to define a workflow definition.
// Most of the time, the implementation only needs to return static value for each method.
// For a dynamic workflow definition, the implementation can return different values based on different constructor inputs.
// To invokes/interact with a dynamic workflows, applications may need to use {@link UntypedClient} instead of {@link Client}
type Workflow interface {
	// GetStates defines the states of the workflow. A state represents a step of the workflow state machine.
	// A state can execute some commands (signal/timer) and wait for result
	// See more details in the WorkflowState interface.
	GetStates() []StateDef

	// GetPersistenceSchema defines all the persistence fields for this workflow, this includes:
	//  1. Data objects
	//  2. Search attributes
	//
	// Data objects can be read/upsert in WorkflowState Start/Decide API
	// Data objects  can also be read by getDataObjects API by external applications using {@link Client}
	//
	// Search attributes can be read/upsert in WorkflowState Start/Decide API
	// Search attributes can also be read by GetSearchAttributes Client API by external applications.
	// External applications can also use "SearchWorkflow" API to find workflows by SQL-like query
	GetPersistenceSchema() []PersistenceFieldDef

	// GetCommunicationSchema defines all the communication methods for this workflow, this includes
	// 1. Signal channel
	// 2. Interstate channel
	//
	// Signal channel is for external applications to send signal to workflow execution.
	// Workflow execution can listen on the signal in the WorkflowState Start API and receive in
	// the WorkflowState Decide API
	//
	// InterStateChannel is for synchronization communications between WorkflowStates.
	// E.g. WorkflowStateA will continue after receiving a value from WorkflowStateB
	///
	GetCommunicationSchema() []CommunicationMethodDef

	// GetWorkflowType Define the workflowType of this workflow definition. By default,(when return empty string), it's the simple Name of the workflow instance,
	// which should be the case for most scenarios.
	//
	// In case of dynamic workflow implementation, return customized values based on constructor input.
	GetWorkflowType() string
}
