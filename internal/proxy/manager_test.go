package proxy

import (
	"testing"

	"oju/internal/config"
	"oju/internal/tracer"
)

const bhaskara_payload = `{"app_key": "bhaskara","name":"span-name","service":"","attributes":{"http.url":"http://delta.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`

const delta_payload = `{"app_key":"delta","name":"check-delta","service":"","attributes":{"http.method":"POST","http.body.a":"4","http.body.b":"2","http.body.c":"-6"}}`

const pitagoras_payload = `{"app_key":"pitagoras","name":"do-hipothenusa","service":"","attributes":{"http.method":"POST","http.body.a":"4","http.body.b":"2","http.body.c":"-6"}}`

func load_config() config.Config {
	config_json := `
{
  "allowed_applications": [
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
	manager := NewManager(config.AllowedApplications)

	message := ApplicationMessage{
		Type:    "TRACE",
		Payload: bhaskara_payload,
	}

	manager.Redirect("bhaskara", message)
	traces := manager.GetAppTraces("bhaskara")

	if len(traces) == 0 {
		t.Fatal("traces must be filled")
	}
}

func TestTwoTracesByDifferentServicesByAppKeyField(t *testing.T) {
	config := load_config()
	manager := NewManager(config.AllowedApplications)

	message := ApplicationMessage{
		Type:    "TRACE",
		Payload: bhaskara_payload,
	}

	message_delta := ApplicationMessage{
		Type:    "TRACE",
		Payload: delta_payload,
	}

	manager.Redirect("bhaskara", message)
	manager.Redirect("delta", message_delta)
	traces := manager.GetAppTraces("bhaskara")

	if len(traces) == 0 {
		t.Fatal("traces must be filled")
	}

	var bhaskara_trace *tracer.Trace

	for _, trace := range traces {
		bhaskara_trace = trace
		break
	}

	children := bhaskara_trace.GetChildren()

	if len(children) == 0 {
		t.Fatal("this children must be filled")
	}
}

func TestConcurrentlyTracesByDifferentServicesByAppKeyField(t *testing.T) {
	config := load_config()
	manager := NewManager(config.AllowedApplications)

	message := ApplicationMessage{
		Type:    "TRACE",
		Payload: bhaskara_payload,
	}

	message_delta := ApplicationMessage{
		Type:    "TRACE",
		Payload: delta_payload,
	}

	message_pitagoras := ApplicationMessage{
		Type:    "TRACE",
		Payload: pitagoras_payload,
	}

	n := 0
	for n <= 5 {
		manager.Redirect("bhaskara", message)
		manager.Redirect("delta", message_delta)
		manager.Redirect("pitagoras", message_pitagoras)
		n += 1
	}

	traces := manager.GetAppTraces("bhaskara")

	if len(traces) == 0 {
		t.Fatal("traces must be filled")
	}

	var bhaskara_trace *tracer.Trace

	for _, trace := range traces {
		bhaskara_trace = trace
		break
	}

	children := bhaskara_trace.GetChildren()

	if len(children) == 0 {
		t.Fatal("this children must be filled")
	}
}
