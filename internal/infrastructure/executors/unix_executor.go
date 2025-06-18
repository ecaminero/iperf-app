package executors

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"iperf-app/internal/domain/entities"
)

type UnixExecutor struct{}

func NewUnixExecutor() *UnixExecutor {
	return &UnixExecutor{}
}

func (u *UnixExecutor) Execute(ctx context.Context, commands []string) (*entities.IExecutionResult, error) {
	return u.ExecuteWithWorkDir(ctx, commands, "")
}

func (u *UnixExecutor) ExecuteWithWorkDir(ctx context.Context, commands []string, workDir string) (*entities.IExecutionResult, error) {
	result := &entities.IExecutionResult{
		ID:        fmt.Sprintf("exec_%d", time.Now().Unix()),
		StartTime: time.Now(),
		Status:    entities.StatusRunning,
	}

	var outputBuilder strings.Builder
	var errorBuilder strings.Builder

	fmt.Printf("Ejecutando %d comandos en Unix/Linux...\n", len(commands))

	for i, command := range commands {
		fmt.Printf("[%d/%d] %s\n", i+1, len(commands), command)

		// Crear comando usando bash -c
		cmd := exec.CommandContext(ctx, "bash", "-c", command)

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

			// Si es un comando crítico (no echo), fallar
			if !u.isNonCriticalCommand(command) {
				result.Status = entities.StatusFailed
				result.EndTime = time.Now()
				result.Output = outputBuilder.String()
				result.ErrorOutput = errorBuilder.String()
				result.ExitCode = u.getExitCode(err)
				return result, fmt.Errorf("comando falló: %s", errorMsg)
			}
		} else {
			fmt.Printf("OK\n")
		}
	}

	// Éxito
	result.Status = entities.StatusCompleted
	result.EndTime = time.Now()
	result.Output = outputBuilder.String()
	result.ErrorOutput = errorBuilder.String()
	result.ExitCode = 0

	fmt.Printf("Todos los comandos ejecutados exitosamente.\n")
	return result, nil
}

// isNonCriticalCommand verifica si un comando es no crítico (echo, etc.)
func (u *UnixExecutor) isNonCriticalCommand(command string) bool {
	lowerCmd := strings.ToLower(strings.TrimSpace(command))
	return strings.HasPrefix(lowerCmd, "echo ") ||
		strings.HasPrefix(lowerCmd, "xdg-open ") ||
		strings.Contains(lowerCmd, "||") // comandos con fallback
}

// getExitCode extrae el código de salida del error
func (u *UnixExecutor) getExitCode(err error) int {
	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.ExitCode()
	}
	return 1
}
