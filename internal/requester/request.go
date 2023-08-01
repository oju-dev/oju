package requester

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"oju/internal/config"
)

const (
	ERROR_APPLICATION_NOT_ALLOWED = "application not allowed"
	ERROR_MALFORMED_HEADER        = "malformed_header"
	ERROR_MALFORMED_PACKET        = "malformed packet"
	ERROR_VERB_NOT_ALLOWED        = "verb_not_allowed"
)

type Request struct {
	Header  Header
	Timer   string
	Message string
}

type Header struct {
	Verb, AppKey, Version string
}

func (header *Header) String() string {
	return fmt.Sprintf("%v %v %v", header.Verb, header.AppKey, header.Version)
}

func (request *Request) String() string {
	return fmt.Sprintf("%v\n%v\n%v", request.Header.String(), request.Timer, request.Message)
}

func NewRequest(head, message string, allowed_applications []config.Service) (Request, error) {
	parts := strings.Split(head, " ")
	if len(parts) != 3 {
		return Request{}, errors.New(ERROR_MALFORMED_HEADER)
	}

	verb := parts[0]
	if !is_verb_allowed(verb) {
		return Request{}, errors.New(ERROR_VERB_NOT_ALLOWED)
	}

	app_key := parts[1]

	if !is_application_allowed(app_key, allowed_applications) {
		return Request{}, errors.New(ERROR_APPLICATION_NOT_ALLOWED)
	}

	version := parts[2]

	header := Header{Verb: verb, AppKey: app_key, Version: version}

	return Request{
		Header:  header,
		Timer:   time.Now().String(),
		Message: message,
	}, nil
}

func Parse(packet string, allowed_applications []config.Service) (Request, error) {
	parts := strings.Split(packet, "\n")
	if len(parts) != 2 {
		return Request{}, errors.New(ERROR_MALFORMED_PACKET)
	}

	head := parts[0]
	message := parts[1]

	request, request_error := NewRequest(head, message, allowed_applications)
	if request_error != nil {
		return Request{}, request_error
	}
	return request, nil
}

func is_verb_allowed(verb string) bool {
	allowed := []string{
		"LOG",
		"WATCH",
		"TRACE",
	}

	for _, allowed_verb := range allowed {
		if allowed_verb == verb {
			return true
		}
	}
	return false
}

func is_application_allowed(app_key string, allowed_applications []config.Service) bool {

	for _, application := range allowed_applications {
		if application.AppKey == app_key {
			return true
		}
	}
	return false
}
