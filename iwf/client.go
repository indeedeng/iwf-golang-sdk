package iwf

type Client interface {
	StartWorkflow()
	StopWorkflow()
	GetSimpleWorkflowResultWithWait()
	GetComplexWorkflowResultsWithWait()
	SignalWorkflow()
	ResetWorkflow()
	DescribeWorkflow()
	GetWorkflowDataObjects()
	GetWorkflowSearchAttributes()
	SearchWorkflow()
}
