package main

import (
	"flag"
	"log"

	"github.com/CPunch/gemini"
)

func handleRequest(peer *gemini.GeminiPeer) {
	if msg, isParam := peer.GetParam(); isParam {
		// send data back to peer!
		body := gemini.NewBody()

		// make proxy request
		response, err := gemini.LazyRequest(msg)
		if err != nil {
			body.AddHeader("Youch!")
			body.AddTextLine(err.Error())
		} else {
			body.AddRaw(response)
		}
		peer.SendBody(body)
	} else {
		// ask peer for data
		peer.SendInput("gemini url: ")
	}
}

func main() {
	// get command line flags
	port := flag.String("port", "1965", "listening port")
	certFile := flag.String("cert", "cert.pem", "certificate PEM file")
	keyFile := flag.String("key", "key.pem", "key PEM file")
	flag.Parse()

	// create server
	server, err := gemini.NewServer(*port, *certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	server.Run(handleRequest)
}
