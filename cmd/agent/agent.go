package main

import (
	"fmt"
	"log"

	"github.com/yhorbachov/curity-oauth-agent-g/pkg/conf"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/server"
)

func main() {
	config := conf.NewDefaultConfig()
	server := server.NewServer(config)

	log.Fatal(server.Listen(fmt.Sprintf(":%d", config.Port)))
}
