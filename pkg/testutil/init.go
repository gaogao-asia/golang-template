package testutil

import "github.com/gaogao-asia/golang-template/config"

func InitConfigForIntegrationTest() {
	var err error
	config.AppConfig, err = config.ViperLoadConfig("../../config/config-local-itest.yml")
	if err != nil {
		panic(err)
	}
}
