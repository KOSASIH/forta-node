package config

import (
	"fmt"
	"path"
)

const ContainerNamePrefix = "forta"

// Docker container names
var (
	DockerSupervisorImage = "forta-network/forta-node:latest"
	DockerUpdaterImage    = "forta-network/forta-node:latest"
	UseDockerImages       = "local"

	DockerSupervisorManagedContainers = 4
	DockerUpdaterContainerName        = fmt.Sprintf("%s-updater", ContainerNamePrefix)
	DockerSupervisorContainerName     = fmt.Sprintf("%s-supervisor", ContainerNamePrefix)
	DockerNatsContainerName           = fmt.Sprintf("%s-nats", ContainerNamePrefix)
	DockerIpfsContainerName           = fmt.Sprintf("%s-ipfs", ContainerNamePrefix)
	DockerScannerContainerName        = fmt.Sprintf("%s-scanner", ContainerNamePrefix)
	DockerJSONRPCProxyContainerName   = fmt.Sprintf("%s-json-rpc", ContainerNamePrefix)
	DockerHostNetContainerName        = fmt.Sprintf("%s-hostnet", ContainerNamePrefix)

	DockerNodeNetworkName = fmt.Sprintf("%s-node", ContainerNamePrefix)

	DefaultContainerFortaDirPath        = "/.forta"
	DefaultContainerConfigPath          = path.Join(DefaultContainerFortaDirPath, DefaultConfigFileName)
	DefaultContainerKeyDirPath          = path.Join(DefaultContainerFortaDirPath, DefaultKeysDirName)
	DefaultContainerLocalAgentsFilePath = path.Join(DefaultContainerFortaDirPath, DefaultLocalAgentsFileName)
)
