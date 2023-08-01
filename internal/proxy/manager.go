package proxy

import (
	"oju/internal/config"
	"oju/internal/parser"
	"oju/internal/tracer"
)

type Manager struct {
	Services   []*Service
	StackTrace *StackTrace
	Links      []Link
	Mailbox    chan Message
}

type Message struct {
	Destination string
	Payload     Payload
}

func NewManager(allowed_applications []config.Service) *Manager {
	manager := &Manager{
		Services:   make([]*Service, 0),
		StackTrace: NewStackTrace(),
		Links:      make([]Link, 0),
		Mailbox:    make(chan Message),
	}

	for _, allowed := range allowed_applications {
		manager.Services = append(manager.Services, get_app(allowed))
	}

	return manager
}

func (manager *Manager) Redirect(destination string, payload Payload) {
	for _, service := range manager.Services {
		metadata := service.GetMetadata()
		if metadata.Host == destination || metadata.Key == destination {
			// TODO: apply switch here
			manager.Links = service.HandleTrace(destination, payload, manager.Links, manager.GetMetadatas())
			return
		}
	}
}

func (manager *Manager) GetAppTraces(destination string) map[string]*tracer.Trace {
	for _, app := range manager.Services {
		metadata := app.GetMetadata()
		if metadata.Host == destination || metadata.Key == destination {
			return app.traces
		}
	}
	return nil
}

func (manager *Manager) GetMetadatas() []Metadata {
	var metadatas []Metadata
	for _, app := range manager.Services {
		metadatas = append(metadatas, app.GetMetadata())
	}

	return metadatas
}

func get_app(config_app config.Service) *Service {
	return &Service{
		parse_tree: parser.NewTree(10),
		traces:     make(map[string]*tracer.Trace),
		metadata: Metadata{
			Key:  config_app.AppKey,
			Host: config_app.Host,
		},
	}
}
