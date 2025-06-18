package infrastructure

import (
	"context"
	"iperf-app/internal/domain/entities"
	"runtime"
)

import ()

type Executor interface {
	Execute(ctx context.Context, commands []string, workDir string) (*entities.Execution, error)
	ValidatePrerequisites(prerequisites []entities.Prerequisite) error
}

type ExecutorFactory struct{}

func (f *ExecutorFactory) CreateExecutor() Executor {
	switch runtime.GOOS {
	case "windows":
		return NewWindowsExecutor()
	default:
		return NewUnixExecutor()
	}
}
