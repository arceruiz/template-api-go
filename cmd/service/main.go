package main

import (
	"github.com/rs/zerolog/log"

	"github.com/arceruiz/template-api-go/internal/channels/rest"
	"github.com/arceruiz/template-api-go/internal/config"
)

func main() {
	config.ParseFromFlags()
	log.Panic().Err(rest.New().Start())
}
