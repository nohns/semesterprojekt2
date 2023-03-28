package main

import (
	"github.com/nohns/servers/cloud"
	"github.com/nohns/servers/pkg/config"
)

func main() {

	config.ReadEnvfile()

	//Instantiate the connect server and the bridge clients
	cloud.Start()
}

//TODO: this might be important for fixing reverse tunnel issue
//https://tailscale.com/
//https://github.com/jhump/grpctunnel
