package tag

import (
	"openpitrix.io/tag/pkg/config"
	"openpitrix.io/tag/pkg/manager"
	"openpitrix.io/tag/pkg/pb"
)

type Client struct {
	pb.TagClient
}

func NewClient() (*Client, error) {
	cfg := config.GetInstance().LoadConf()
	conn, err := manager.NewClient(cfg.App.ApiHost, cfg.App.ApiPort)
	if err != nil {
		return nil, err
	}

	return &Client{
		TagClient: pb.NewTagClient(conn),
	}, nil
}
