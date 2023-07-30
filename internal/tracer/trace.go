package tracer

import (
	"encoding/json"
	"fmt"

	"oju/internal/utils"
)

type Trace struct {
	id         string
	AppKey     string            `json:"app_key"`
	Name       string            `json:"name"`
	Service    string            `json:"service"`
	Attributes map[string]string `json:"attributes"`
	children   map[string]*Trace
}

// TODO: set app_key as a parameter here
func Parse(packet string) (Trace, error) {
	var tracer Trace
	unmarshal_error := json.Unmarshal([]byte(packet), &tracer)

	if unmarshal_error != nil {
		return Trace{}, unmarshal_error
	}

	tracer.SetId()

	return tracer, nil
}

func (trace *Trace) SetId() {
	trace.id = utils.GenerateId()
}

func (trace *Trace) GetId() string {
	return trace.id
}

func (trace *Trace) GetChildren() map[string]*Trace {
	return trace.children
}

func (trace *Trace) SetAppKey(app_key string) {
	trace.AppKey = app_key
}

func (trace *Trace) AddChild(new_trace *Trace) {
	id := new_trace.GetId()
	if trace.children == nil {
		trace.children = make(map[string]*Trace)
		trace.children[id] = new_trace
	} else {
		trace.children[id] = new_trace
	}
}

func (trace *Trace) Print() {
	var service string

	if trace.Service == "" {
		service = "No service pointed"
	} else {
		service = trace.Service
	}

	fmt.Println("=> TRACE from ", trace.AppKey)
	fmt.Println("[id]: ", trace.GetId())
	fmt.Println("[span-name]: ", trace.Name)

	fmt.Println("[service]: ", service)
	fmt.Println("[children]: ", len(trace.GetChildren()))
	fmt.Println("[attributes]:")
	for key, value := range trace.Attributes {
		fmt.Printf("\t[%s]: %s\n", key, value)
	}
}
