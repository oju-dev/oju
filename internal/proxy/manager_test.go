package proxy

import (
	"testing"

	"oju/internal/config"
)

const bhaskara_payload = `{"app_key": "bhaskara","name":"span-name","service":"","attributes":{"http.url":"http://delta.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`

const delta_payload = `{"app_key":"delta","name":"check-delta","service":"","attributes":{"http.method":"POST","http.body.a":"4","http.body.b":"2","http.body.c":"-6"}}`

const pitagoras_payload = `{"app_key":"pitagoras","name":"do-hipothenusa","service":"","attributes":{"http.method":"POST","http.body.a":"4","http.body.b":"2","http.body.c":"-6"}}`

func load_config() config.Config {
	config_json := `
{
  "allowed_services": [
    {
      "name": "bhaskara",
      "app_key": "bhaskara",
			"host": "http://bhaskara.api.svc.cluster.local"
    },
	{
		"name": "delta",
		"app_key": "delta",
		"host": "http://delta.api.svc.cluster.local"
	},
	{
		"name": "pitagoras",
		"app_key": "pitagoras",
		"host": "http://pitagoras.api.svc.cluster.local"
	}
  ]
}
`
	config, _ := config.BuildConfig([]byte(config_json))
	return config
}

func TestUpAllAllowedApplicationsByProxy(t *testing.T) {
	config := load_config()
	manager := NewManager(config.AllowedServices)

	message := Payload{
		Type:    "TRACE",
		Payload: bhaskara_payload,
	}

	manager.Redirect("bhaskara", message)
	traces := manager.GetAppTraces("bhaskara")

	if len(traces) == 0 {
		t.Fatal("traces must be filled")
	}
}

func TestLinkedTraces(t *testing.T) {
	config := load_config()
	manager := NewManager(config.AllowedServices)

	message := Payload{
		Type:    "TRACE",
		Payload: bhaskara_payload,
	}

	message_delta := Payload{
		Type:    "TRACE",
		Payload: delta_payload,
	}

	manager.Redirect("bhaskara", message)
	manager.Redirect("delta", message_delta)
}
