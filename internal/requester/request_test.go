package requester

import (
	"testing"

	"oju/internal/config"
)

func load_config() config.Config {
	config_json := `
{
  "allowed_services": [
    {
      "name": "worker",
      "app_key": "3FAFCF87-BF66-4DC5-84C1-34E178FF55CC"
    }
  ]
}
`

	config, _ := config.BuildConfig([]byte(config_json))
	return config
}

func TestTraceUrlPacket(t *testing.T) {
	trace_packet_url_attr := `TRACE 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO\n{"name":"span-name","service":"","attributes":{"http.url":"http://products.api.svc.cluster.local","http.method":"POST","http.body.email":"test@email.com"}}`

	config := load_config()
	_, parse_error := Parse(trace_packet_url_attr, config.AllowedServices)
	if parse_error == nil {
		t.Error("Packet not malformed")
	}
}

func TestMalformedPacket(t *testing.T) {
	malformed_packet := "LOG 5C3D47E8-2D9C-4165-A48A-7A6F6449DF66"
	config := load_config()
	_, parse_error := Parse(malformed_packet, config.AllowedServices)
	if parse_error == nil {
		t.Error("Packet not malformed")
	}
}

func TestMalformedHeader(t *testing.T) {
	malformed_header := "LOG AWO\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""
	config := load_config()
	_, parse_error := Parse(malformed_header, config.AllowedServices)
	if parse_error == nil {
		t.Error("Header not malformed")
	}
}

func TestParseLog(t *testing.T) {
	test_packet := "LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""
	config := load_config()
	log, parse_error := Parse(test_packet, config.AllowedServices)
	if parse_error != nil {
		t.Error(parse_error.Error())
	}

	if log.Timer == "" {
		t.Error("Timer is empty")
	}
}

func TestHeaderValidationDisallowedApps(t *testing.T) {
	test_disallowed_apps := "LOG AAAAAAAAAAAAAA AWO1.1\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""
	config := load_config()
	_, parse_error := Parse(test_disallowed_apps, config.AllowedServices)
	if parse_error == nil {
		t.Error("Should be an error because that app is not allowed")
	}
}

func TestHeaderWithInvalidVerb(t *testing.T) {
	test_not_allowed_verb := "AAA AAAAAAAAAAAAAA AWO1.1\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""
	config := load_config()
	_, parse_error := Parse(test_not_allowed_verb, config.AllowedServices)
	if parse_error == nil {
		t.Error("Should be an error because that app is not allowed")
	}
}
