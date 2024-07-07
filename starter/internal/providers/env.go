package providers

import "github.com/ugabiga/swan/core"

type EnvironmentVariables struct {
	AppName string `validate:"required" env:"APP_NAME" json:"app_name,omitempty"`
	Addr    string `validate:"required" env:"ADDR" json:"addr,omitempty"`
}

func ProvideEnvironmentVariables() *EnvironmentVariables {
	return core.ValidateEnvironmentVariables[EnvironmentVariables]()
}
