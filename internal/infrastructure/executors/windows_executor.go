package executors

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"iperf-app/internal/domain/entities"
)

type WindowsExecutor struct{}

func NewWindowsExecutor() *WindowsExecutor {
	return &WindowsExecutor{}
}

func (w *WindowsExecutor) Execute(ctx context.Context, commands []string) (*entities.IExecutionResult, error) {
	return w.ExecuteWithWorkDir(ctx, commands, "")
}

func (w *WindowsExecutor) ExecuteWithWorkDir(ctx context.Context, commands []string, workDir string) (*entities.IExecutionResult, error) {
	result := &entities.IExecutionResult{
		ID:        fmt.Sprintf("exec_%d", time.Now().Unix()),
		StartTime: time.Now(),
		Status:    entities.StatusRunning,
	}

	var outputBuilder strings.Builder
	var errorBuilder strings.Builder

	fmt.Printf("Ejecutando %d comandos en Windows...\n", len(commands))

	for i, command := range commands {
		fmt.Printf("[%d/%d] %s\n", i+1, len(commands), command)

		// Crear comando usando cmd /C
		cmd := exec.CommandContext(ctx, "cmd", "/C", command)

		// Establecer directorio de trabajo si se especifica
		if workDir != "" {
			cmd.Dir = workDir
		}

		// Ejecutar comando
		output, err := cmd.CombinedOutput()
		outputStr := string(output)

		// Agregar output
		outputBuilder.WriteString(fmt.Sprintf("=== Comando: %s ===\n", command))
		outputBuilder.WriteString(outputStr)
		outputBuilder.WriteString("\n")

		if err != nil {
			errorMsg := fmt.Sprintf("Error ejecutando '%s': %v\n", command, err)
			errorBuilder.WriteString(errorMsg)
			fmt.Printf("ERROR: %s", errorMsg)

			// Si es un comando crítico (no echo ni start), fallar
			if !w.isNonCriticalCommand(command) {
				result.Status = entities.StatusFailed
				result.EndTime = time.Now()
				result.Output = outputBuilder.String()
				result.ErrorOutput = errorBuilder.String()
				result.ExitCode = w.getExitCode(err)
				return result, fmt.Errorf("comando falló: %s", errorMsg)
			}
		} else {
			fmt.Printf("OK\n")
		}
	}

	result.Status = entities.StatusCompleted
	result.EndTime = time.Now()
	result.Output = outputBuilder.String()
	result.ErrorOutput = errorBuilder.String()
	result.ExitCode = 0

	fmt.Printf("Todos los comandos ejecutados exitosamente.\n")
	return result, nil
}

// isNonCriticalCommand verifica si un comando es no crítico (echo, start, etc.)
func (w *WindowsExecutor) isNonCriticalCommand(command string) bool {
	lowerCmd := strings.ToLower(strings.TrimSpace(command))
	return strings.HasPrefix(lowerCmd, "echo ") ||
		strings.HasPrefix(lowerCmd, "start ") ||
		strings.HasPrefix(lowerCmd, "cd ")
}

// getExitCode extrae el código de salida del error
func (w *WindowsExecutor) getExitCode(err error) int {
	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.ExitCode()
	}
	return 1
}
