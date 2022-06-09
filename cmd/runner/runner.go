package runner

import (
	"context"
	"fmt"

	"github.com/forta-network/forta-core-go/utils"
	"github.com/forta-network/forta-node/clients"
	"github.com/forta-network/forta-node/config"
	"github.com/forta-network/forta-node/services"
	"github.com/forta-network/forta-node/services/runner"
	"github.com/forta-network/forta-node/store"
	log "github.com/sirupsen/logrus"
)

func initServices(ctx context.Context, cfg config.Config) ([]services.Service, error) {
	shouldDisableAutoUpdate := cfg.AutoUpdate.Disable || cfg.PrivateModeConfig.Enable
	imgStore, err := store.NewFortaImageStore(ctx, config.DefaultContainerPort, !shouldDisableAutoUpdate, utils.String(cfg.AutoUpdate.PrereleaseVersion))
	if err != nil {
		return nil, fmt.Errorf("failed to create the image store: %v", err)
	}
	dockerClient, err := clients.NewDockerClient("runner")
	if err != nil {
		return nil, fmt.Errorf("failed to create the docker client: %v", err)
	}
	globalDockerClient, err := clients.NewDockerClient("")
	if err != nil {
		return nil, fmt.Errorf("failed to create the docker client: %v", err)
	}

	if cfg.Development {
		log.Warn("running in development mode")
	}

	return []services.Service{
		runner.NewRunner(ctx, cfg, imgStore, dockerClient, globalDockerClient),
	}, nil
}

// Run runs the runner.
func Run(cfg config.Config) {
	ctx, cancel := services.InitMainContext()
	defer cancel()

	logger := log.WithField("process", "runner")
	logger.Info("starting")
	defer logger.Info("exiting")

	serviceList, err := initServices(ctx, cfg)
	if err != nil {
		logger.WithError(err).Error("could not initialize services")
		return
	}

	if err := services.StartServices(ctx, cancel, log.NewEntry(log.StandardLogger()), serviceList); err != nil {
		logger.WithError(err).Error("error running services")
	}
}
