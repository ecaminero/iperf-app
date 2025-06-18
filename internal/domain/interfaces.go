package domain

import (
	"context"
	"iperf-app/internal/domain/entities"
)

// Executor define la interfaz común para ejecutar comandos
type Executor interface {
	Execute(ctx context.Context, commands []string) (*entities.IExecutionResult, error)
	ExecuteWithWorkDir(ctx context.Context, commands []string, workDir string) (*entities.IExecutionResult, error)
}
