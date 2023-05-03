package main

import (
	"github.com/nohns/servers/cloud/server"
	"github.com/nohns/servers/pkg/config"
)

func main() {

	config.ReadEnvfile()

	//Instantiate the connect server and the bridge clients
	server.Start()
}
