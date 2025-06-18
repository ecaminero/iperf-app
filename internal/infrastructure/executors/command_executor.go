package executors

import (
	"context"
	"fmt"
	"time"

	"iperf-app/internal/domain"
	"iperf-app/internal/domain/entities"
	"iperf-app/internal/infrastructure/commands"
)

// CommandExecutor combina commandos y ejecutores para operaciones completas
type CommandExecutor struct {
	commandFactory  *commands.CommandFactory
	executorFactory *ExecutorFactory
	executor        domain.Executor
}

func NewCommandExecutor() *CommandExecutor {
	executorFactory := NewExecutorFactory()
	return &CommandExecutor{
		commandFactory:  commands.NewCommandFactory(),
		executorFactory: executorFactory,
		executor:        executorFactory.CreateExecutor(),
	}
}

func (ce *CommandExecutor) check() {
	fmt.Printf("Checking")
}

func (ce *CommandExecutor) ExecuteZTP(interfaceName, ipAddress, subnetMask, gateway string) (*entities.IExecutionResult, error) {
	fmt.Printf("=== EJECUTANDO ZTP ===\n")
	fmt.Printf("Interfaz: %s\n", interfaceName)
	fmt.Printf("IP: %s\n", ipAddress)
	fmt.Printf("Máscara: %s\n", subnetMask)
	fmt.Printf("Gateway: %s\n", gateway)
	fmt.Println()

	command := ce.commandFactory.GetNetworkCommands().GetZTPCommand(interfaceName, ipAddress, subnetMask, gateway)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return ce.executor.Execute(ctx, command)
}

// ExecuteIperfInternetNational ejecuta test iPerf Internet Nacional
func (ce *CommandExecutor) ExecuteIperfInternetNational(title, duration, parallel string) (*entities.IExecutionResult, error) {
	timestamp := time.Now().Format("2006-01-02")

	fmt.Printf("=== EJECUTANDO IPERF INTERNET NACIONAL ===\n")
	fmt.Printf("Título: %s\n", title)
	fmt.Printf("Duración: %s seg\n", duration)
	fmt.Printf("Paralelas: %s\n", parallel)
	fmt.Println()

	command := ce.commandFactory.GetIperfCommands().GetIperfInternetNationalCommand(title, duration, parallel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	fmt.Printf("Prueba iniciada a las %s", timestamp)
	return ce.executor.Execute(ctx, command)
}

// ExecuteIperfInternetInternational ejecuta test iPerf Internet Internacional
func (ce *CommandExecutor) ExecuteIperfInternetInternational(title, duration, parallel string) (*entities.IExecutionResult, error) {
	fmt.Printf("=== EJECUTANDO IPERF INTERNET INTERNACIONAL ===\n")
	fmt.Printf("Título: %s\n", title)
	fmt.Printf("Duración: %s seg\n", duration)
	fmt.Printf("Paralelas: %s\n", parallel)
	fmt.Println()

	command := ce.commandFactory.GetIperfCommands().GetIperfInternetInternationalCommand(title, duration, parallel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	return ce.executor.Execute(ctx, command)
}

// ExecuteIperfMPLS ejecuta test iPerf MPLS
func (ce *CommandExecutor) ExecuteIperfMPLS(title, ipserver, duration, parallel string) (*entities.IExecutionResult, error) {
	fmt.Printf("=== EJECUTANDO IPERF MPLS ===\n")
	fmt.Printf("Título: %s\n", title)
	fmt.Printf("Servidor: %s\n", ipserver)
	fmt.Printf("Duración: %s seg\n", duration)
	fmt.Printf("Paralelas: %s\n", parallel)
	fmt.Println()

	command := ce.commandFactory.GetIperfCommands().GetIperfMPLSCommand(title, ipserver, duration, parallel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	return ce.executor.Execute(ctx, command)
}

// ExecuteIperfP2PClient ejecuta test iPerf P2P Cliente
func (ce *CommandExecutor) ExecuteIperfP2PClient(title, ip, duration, parallel string) (*entities.IExecutionResult, error) {
	fmt.Printf("=== EJECUTANDO IPERF P2P CLIENTE ===\n")
	fmt.Printf("Título: %s\n", title)
	fmt.Printf("Servidor: %s\n", ip)
	fmt.Printf("Duración: %s seg\n", duration)
	fmt.Printf("Paralelas: %s\n", parallel)
	fmt.Println()

	command := ce.commandFactory.GetIperfCommands().GetIperfP2PClientCommand(title, ip, duration, parallel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	return ce.executor.Execute(ctx, command)
}

func (ce *CommandExecutor) ExecuteIperfP2PServer() (*entities.IExecutionResult, error) {
	fmt.Printf("=== EJECUTANDO IPERF P2P SERVIDOR ===\n")
	fmt.Printf("Iniciando servidor iPerf (Ctrl+C para detener)...\n")
	fmt.Println()

	command := ce.commandFactory.GetIperfCommands().GetIperfP2PServerCommand()

	// Sin timeout para el servidor (se ejecuta hasta que se cancele)
	ctx := context.Background()

	return ce.executor.Execute(ctx, command)
}

func (ce *CommandExecutor) ExecutePing(target string, count int) (*entities.IExecutionResult, error) {
	fmt.Printf("=== EJECUTANDO PING ===\n")
	fmt.Printf("Objetivo: %s\n", target)
	fmt.Printf("Cantidad: %d\n", count)
	fmt.Println()

	command := ce.commandFactory.GetDiagnosticCommands().GetPingCommand(target, count)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	return ce.executor.Execute(ctx, command)
}

// ExecuteTraceroute ejecuta traceroute
func (ce *CommandExecutor) ExecuteTraceroute(target string) (*entities.IExecutionResult, error) {
	fmt.Printf("=== EJECUTANDO TRACEROUTE ===\n")
	fmt.Printf("Objetivo: %s\n", target)
	fmt.Println()

	command := ce.commandFactory.GetDiagnosticCommands().GetTracerouteCommand(target)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	return ce.executor.Execute(ctx, command)
}

// ExecuteRouteTable muestra la tabla de rutas
func (ce *CommandExecutor) ExecuteRouteTable() (*entities.IExecutionResult, error) {
	fmt.Printf("=== MOSTRANDO TABLA DE RUTAS ===\n")
	fmt.Println()

	command := ce.commandFactory.GetDiagnosticCommands().GetRouteTableCommand()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return ce.executor.Execute(ctx, command)
}

func (ce *CommandExecutor) ExecuteAddRoute(destination, mask, gateway string) (*entities.IExecutionResult, error) {
	fmt.Printf("=== AGREGANDO RUTA ESTÁTICA ===\n")
	fmt.Printf("Destino: %s\n", destination)
	fmt.Printf("Máscara: %s\n", mask)
	fmt.Printf("Gateway: %s\n", gateway)
	fmt.Println()

	command := ce.commandFactory.GetDiagnosticCommands().GetAddRouteCommand(destination, mask, gateway)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return ce.executor.Execute(ctx, command)
}

func (ce *CommandExecutor) ExecuteOpenFolder(path string) (*entities.IExecutionResult, error) {
	fmt.Printf("=== ABRIENDO CARPETA ===\n")
	fmt.Printf("Ruta: %s\n", path)
	fmt.Println()

	command := ce.commandFactory.GetDiagnosticCommands().GetOpenFolderCommand(path)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return ce.executor.Execute(ctx, command)
}
