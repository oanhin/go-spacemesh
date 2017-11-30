package main

import (
	"github.com/UnrulyOS/go-unruly/app"
	"github.com/UnrulyOS/go-unruly/node"
)

func main() {

	// test p2p protocols
	node.TestP2pProtocols()

	// run the app
	app.Main()

	// add any playground tests here....
}