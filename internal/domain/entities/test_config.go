package entities

// TestType representa los tipos de test disponibles
type TestType string

const (
	TestIperfInternetNational      TestType = "internet_national"
	TestIperfInternetInternational TestType = "internet_international"
	TestIperfMPLS                  TestType = "mpls"
	TestIperfP2PClient             TestType = "p2p_client"
	TestIperfP2PServer             TestType = "p2p_server"
)

// TestConfig contiene la configuración para ejecutar un test iPerf
type TestConfig struct {
	Type       TestType
	Title      string // Formato: CLIENTE_SITIO
	Duration   int    // Duración en segundos
	Parallel   int    // Número de conexiones paralelas
	ServerIP   string // IP del servidor
	OutputFile string // Nombre del archivo de salida
}
