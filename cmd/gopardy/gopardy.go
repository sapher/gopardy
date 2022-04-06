package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)	
}