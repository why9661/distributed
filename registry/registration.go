package registry

const (
	RegistryHost = "localhost"
	RegistryPort = "8000"
)

type Registration struct {
	ServiceName string
	ServiceURL  string
}
