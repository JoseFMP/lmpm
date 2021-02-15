package lmpm

import "fmt"

// Client to access LMPM API
type Client struct {
	CompanyID CompanyID
	APIKey    string
	BaseUrl   Environment
}

func Init(companyIDAsString string, apiKey string, environment Environment) (*Client, error) {

	errValidatingInput := validateInput(companyIDAsString, apiKey, environment)
	if errValidatingInput != nil {
		return nil, errValidatingInput
	}
	return &Client{
		CompanyID: CompanyID(companyIDAsString),
		APIKey:    apiKey,
		BaseUrl:   environment,
	}, nil
}

func validateInput(companyID string, apiKey string, environment Environment) error {
	if companyID == "" {
		return fmt.Errorf("Company ID is empty")
	}

	if apiKey == "" {
		return fmt.Errorf("API Key not provided")
	}

	if environment != environments.Production && environment != environments.Development {
		return fmt.Errorf("Unknown environment %s", string(environment))
	}

	return nil
}

type CompanyID string
type Environment string

type allEnvironments struct {
	Production  Environment
	Development Environment
}

func GetEnvironments() *allEnvironments {
	return &allEnvironments{
		Production:  EnvironmentProduction,
		Development: EnvironmentDevelopment,
	}
}

const EnvironmentProduction Environment = "https://integrations.api.lmpm.com/v1"
const EnvironmentDevelopment Environment = "https://integrations.api.lmpm.biz/v1"
