package main

import "log"

func handleFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
