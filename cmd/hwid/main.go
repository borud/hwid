// Package main contains a simple utility that will return the base36 value
// if the primary hardware MAC address. This utility is mostly meant for
// testing and debugging.

// If this results in an error you should probably try to re-run it with
// the -i flag to explicitly make it choose a network interface to base
// the address on.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/borud/hwid"
)

var intf string

func init() {
	flag.StringVar(&intf, "i", "", "network interface")
	flag.Parse()
}

func main() {
	id, err := hwid.ID(intf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
