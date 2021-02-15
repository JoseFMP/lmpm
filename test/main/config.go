package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/JoseFMP/lmpm"
)

func readTestConfig() *testConfig {
	rawConfig, errReadingConfigFile := ioutil.ReadFile("./test/test-config.json")

	if errReadingConfigFile != nil {
		panic(errReadingConfigFile)
	}

	var conf testConfig
	errParsingConfig := json.Unmarshal(rawConfig, &conf)
	if errParsingConfig != nil {
		panic(errParsingConfig)
	}
	return &conf
}

type testConfig struct {
	APIKey    string           `json:"apiKey"`
	CompanyID lmpm.CompanyID   `json:"companyId"`
	BaseUrl   lmpm.Environment `json:"baseUrl"`
}
