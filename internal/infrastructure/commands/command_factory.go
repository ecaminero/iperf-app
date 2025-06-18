package commands

// CommandFactory crea los comandos específicos según el OS
type CommandFactory struct {
	iperfCommands      *IperfCommands
	networkCommands    *NetworkCommands
	diagnosticCommands *DiagnosticCommands
}

func NewCommandFactory() *CommandFactory {
	return &CommandFactory{
		iperfCommands:      NewIperfCommands(),
		networkCommands:    NewNetworkCommands(),
		diagnosticCommands: NewDiagnosticCommands(),
	}
}

func (cf *CommandFactory) GetIperfCommands() *IperfCommands {
	return cf.iperfCommands
}

func (cf *CommandFactory) GetNetworkCommands() *NetworkCommands {
	return cf.networkCommands
}

func (cf *CommandFactory) GetDiagnosticCommands() *DiagnosticCommands {
	return cf.diagnosticCommands
}
