package domain

import (
	"flag"
)

type Config struct {
	DockerAPIEndpoint  string
	VulcandAPIEndpoint string
	HostIP             string
}

func (config *Config) InstallFlags() {
	flag.StringVar(&config.DockerAPIEndpoint, "d", "/var/run/docker.sock", "Docker API endpoint")
	flag.StringVar(&config.VulcandAPIEndpoint, "v", "172.17.8.101:8182", "Vulcand API endpoint")
	flag.StringVar(&config.HostIP, "h", "172.17.8.101", "Host's external facing ip")
}
