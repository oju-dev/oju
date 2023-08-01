package config

import (
	"testing"
)

func TestReturnAllowedApplication(t *testing.T) {
	config_json := `
{
	"allowed_services": [
	{
	"name": "worker",
	"app_key": "37F129EF-687E-47AC-B0C4-CBF6516B37ED"
	}
	]
}
`

	config, load_error := BuildConfig([]byte(config_json))
	if load_error != nil {
		t.Error(load_error)
	}

	if len(config.AllowedServices) == 0 {
		t.Error("Should be at least 1")
	}
}
