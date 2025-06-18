package executors

import (
	"iperf-app/internal/domain"
	"runtime"
)

// ExecutorFactory crea el ejecutor apropiado según el OS
type ExecutorFactory struct{}

func NewExecutorFactory() *ExecutorFactory {
	return &ExecutorFactory{}
}

// CreateExecutor crea el ejecutor apropiado para el OS actual
func (ef *ExecutorFactory) CreateExecutor() domain.Executor {
	switch runtime.GOOS {
	case "windows":
		return NewWindowsExecutor()
	case "linux", "darwin", "freebsd", "openbsd", "netbsd":
		return NewUnixExecutor()
	default:
		// Fallback a Unix para sistemas desconocidos
		return NewUnixExecutor()
	}
}

// CreateWindowsExecutor crea específicamente un ejecutor Windows
func (ef *ExecutorFactory) CreateWindowsExecutor() domain.Executor {
	return NewWindowsExecutor()
}

// CreateUnixExecutor crea específicamente un ejecutor Unix
func (ef *ExecutorFactory) CreateUnixExecutor() domain.Executor {
	return NewUnixExecutor()
}
