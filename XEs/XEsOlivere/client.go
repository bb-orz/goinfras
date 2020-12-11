package XEsOlivere

import (
	elasticv7 "github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
)

// 创建默认的ES Client
func CreateDefaultESClient() error {
	var err error
	esClient, err = elasticv7.NewClient()
	return err
}

func NewESClient(cfg *Config) (*elasticv7.Client, error) {
	var err error
	var esClient *elasticv7.Client

	if cfg == nil {
		esClient, err := elasticv7.NewClient()
		return esClient, err
	}

	olivCfg := &config.Config{
		URL:         cfg.URL,
		Username:    cfg.Username,
		Password:    cfg.Password,
		Sniff:       &cfg.Sniff,
		Healthcheck: &cfg.Healthcheck,
		Infolog:     cfg.Infolog,
		Errorlog:    cfg.Errorlog,
		Tracelog:    cfg.Tracelog,
	}
	esClient, err = elasticv7.NewClientFromConfig(olivCfg)

	if err != nil {
		return nil, err
	}

	return esClient, nil
}
