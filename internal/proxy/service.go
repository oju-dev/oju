package proxy

import (
	"oju/internal/parser"
	"oju/internal/tracer"
)

type Metadata struct {
	Key        string
	Host       string
	WatchQuery string
}

type Payload struct {
	Type    string
	Payload string
}

type Service struct {
	parse_tree *parser.Tree
	traces     map[string]*tracer.Trace
	metadata   Metadata
	errors     []error
}

func (service *Service) GetMetadata() Metadata {
	return service.metadata
}

func (service *Service) GetTraces() map[string]*tracer.Trace {
	return service.traces
}

func (service *Service) HandleTrace(destination string, message Payload, links []Link, applications_metadata []Metadata) []Link {

	trace, parse_trace_error := tracer.Parse(message.Payload, destination)

	if parse_trace_error != nil {
		service.errors = append(service.errors, parse_trace_error)
		return links
	}

	// TODO: this will be in a storage engine
	service.traces[trace.GetId()] = &trace

	if len(links) == 0 {
		new_link := Link{Head: &Node{Value: trace}}
		links = append(links, new_link)
		return links
	}

	errors := 0
	max_errors := len(links)

	// stack_trace.RunStack(&trace, applications_metadata)
	for _, link := range links {
		_, error_on_found := link.Search(destination)
		if error_on_found != nil {
			errors += 1
			continue
		}
		// TODO: generate a new node and put as next
		link.Append(trace)
	}

	if errors == max_errors {
		new_link := Link{Head: &Node{Value: trace}}
		links = append(links, new_link)
		return links
	}
	return links
	// switch message.Type {
	// case "LOG":
	// 	parser.ParseLog(service.parse_tree, message.Payload)
	// case "TRACE":
	// }
}
