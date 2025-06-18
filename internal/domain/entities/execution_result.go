package entities

import "time"

type ExecutionStatus string

const (
	StatusPending   ExecutionStatus = "pending"
	StatusRunning   ExecutionStatus = "running"
	StatusCompleted ExecutionStatus = "completed"
	StatusFailed    ExecutionStatus = "failed"
)

type IExecutionResult struct {
	ID          string
	StartTime   time.Time
	EndTime     time.Time
	Status      ExecutionStatus
	Output      string
	ErrorOutput string
	ExitCode    int
}
