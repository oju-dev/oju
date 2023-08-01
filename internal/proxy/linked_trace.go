package proxy

import (
	"errors"
	"oju/internal/tracer"
)

type Link struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value tracer.Trace
	Next  *Node
}

func (link *Link) Append(trace tracer.Trace) {
	node := &Node{Value: trace}

	if link.Head == nil {
		link.Head = node
	}

	if link.Tail != nil {
		link.Tail.Next = node
	}

	link.Tail = node
}

func (link *Link) Search(destination string, metadatas []Metadata) (tracer.Trace, error) {
	node := link.Head
	for node != nil {
		if node.Value.Service == destination {
			break
		}
		host, host_error := get_host_by_app_key(destination, metadatas)
		if host_error != nil {
		}

		node = node.Next
	}

	if node != nil {
		return node.Value, nil
	}
	return tracer.Trace{}, errors.New("trace not found")
}
