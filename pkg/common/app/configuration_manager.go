package app

import "github.com/erkindilekci/product-api/pkg/common/postgresql"

type ConfigurationManager struct {
	PostgresqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postgresqlConfig := postgresql.Config{
		Host:                  "localhost",
		Port:                  "5433",
		UserName:              "postgres",
		Password:              "password",
		DbName:                "productapp",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
	return &ConfigurationManager{postgresqlConfig}
}
