package tracer

import "testing"

func TestGenerateTrace(t *testing.T) {
	payload := `{"name":"span-name","service":"","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`

	trace, trace_error := Parse(payload, "")

	if trace_error != nil {
		t.Error("should not be any error")
	}

	if trace.id == "" {
		t.Error("trace id should not be empty")
	}

}
