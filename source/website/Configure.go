package website

import "EvoGuard/source/config"

/*
	Configure.go will implement the opening of configuration
	from the goconfig operations and attempt to use it.
*/

// ServerConfig will attempt to keep all options for the server in a clean env
type ServerConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Cert    string `json:"cert"`
	Key     string `json:"key"`
	Assets  struct {
		Directory  string `json:"directory"`
		ServeUnder string `json:"serveUnder"`
	} `json:"assets"`
}

// Marshals into the destination from website.
func NewServerConfig(destination *ServerConfig) error {
	return config.Options.MarshalFromPath(destination, "website")
}

// NewConfig will try to make a brand new HTTPS server configuration.
func NewConfig() (*ServerConfig, error) {
	destination := new(ServerConfig)
	err := NewServerConfig(destination)
	if err != nil {
		return nil, err
	}

	return destination, nil
}
