package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"oju/internal/commander"
	"oju/internal/config"
	"oju/internal/request"
	"oju/internal/track"
	"os"
)

func main() {
	fmt.Println("\033[33m" + commander.USAGE + "\033[97m")
	config, load_error := config.LoadConfigFile()

	if load_error != nil {
		log.Fatalln(load_error.Error())
	}

	// TODO: when start server, initialize two channels
	// ARMAZEN: will be responsible to get these tracks/logs whatever
	// 			and save to sqlite
	// PROCESSOR: will be responsible to get these logs
	// 							and pre process this

	// featurizer will be serving as data pre-processor for running DBSCAN
	// when is needed to get anomaly detection
	// system := usecases.NewSystem(config.Resources)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9090"
	}

	listener, listener_error := net.Listen("tcp", ":"+port)
	if listener_error != nil {
		log.Fatalln("Oju error: ", listener_error.Error())
	}

	defer listener.Close()
	log.Println("[OJU STARTED] port:" + port)
	fmt.Println("--------------------------------------------")

	for {
		socket, socket_error := listener.Accept()
		if socket_error != nil {
			log.Println("Socket Error: ", socket_error.Error())
		}

		go handle_incoming_message(socket, config)
	}
}

func handle_incoming_message(socket net.Conn, config config.Config) {
	log.Println("New connection accepted: ", socket.RemoteAddr().String())
	reader := bufio.NewReader(socket)

	for {
		message, message_error := io.ReadAll(reader)

		if message_error != nil {
			if message_error == io.EOF {
				break
			}
			log.Println("Error on getting message: ", message_error.Error())
			break
		}

		if string(message) == "" {
			break
		}

		request, request_error := request.Parse(string(message), config.Resources)
		if request_error != nil {
			log.Println("Error on parsing request: ", request_error.Error())
			break
		}

		switch request.Header.Verb {
		case "TRACK":
			_, parse_trace_error := track.Parse(request.Message)

			if parse_trace_error != nil {
				log.Println("Error on parsing trace: ", parse_trace_error.Error())
			}
			// FIXME: what we do here

			// command := entities.NewInsertActionCommand(trace)
			// usecases.Send(sys, command)
		default:
			log.Println("tried VERB: ", request.Header.Verb)
		}

	}
}
