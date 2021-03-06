package iguazio

import (
	"github.com/nuclio/nuclio/pkg/platform"
	"github.com/nuclio/nuclio/pkg/platformconfig"

	"github.com/nuclio/logger"
	"github.com/nuclio/nuclio-sdk-go"
)

type Client struct {
	logger                logger.Logger
	platformConfiguration *platformconfig.Config
}

func NewClient(parentLogger logger.Logger, platformConfiguration *platformconfig.Config) (*Client, error) {
	newClient := Client{
		logger:                parentLogger.GetChild("leader-client-iguazio"),
		platformConfiguration: platformConfiguration,
	}

	return &newClient, nil
}

func (c *Client) Create(createProjectOptions *platform.CreateProjectOptions) error {
	return nuclio.ErrNotImplemented
}

func (c *Client) Update(updateProjectOptions *platform.UpdateProjectOptions) error {
	return nuclio.ErrNotImplemented
}

func (c *Client) Delete(deleteProjectOptions *platform.DeleteProjectOptions) error {
	return nuclio.ErrNotImplemented
}
