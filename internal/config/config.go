package config

import (
	"flag"
	"log"

	"github.com/notnull-co/cfg"
	"github.com/rs/zerolog"
)

var (
	Instance Config

	mapLogLevel = map[string]zerolog.Level{
		"debug":    zerolog.DebugLevel,
		"disabled": zerolog.Disabled,
		"error":    zerolog.ErrorLevel,
		"fatal":    zerolog.FatalLevel,
		"info":     zerolog.InfoLevel,
		"panic":    zerolog.PanicLevel,
		"warn":     zerolog.WarnLevel,
	}
)

type Config struct {
	Server struct {
		HttpPort string `cfg:"http_port" default:"8080"`
	}
	Logger struct {
		Level string `cfg:"level" default:"debug"`
	}
}

func ParseFromFlags() {
	var configDir string
	flag.StringVar(&configDir, "config-dir", "", "Configuration file directory")
	var secretDir string
	flag.StringVar(&secretDir, "secret-dir", "", "Secret file directory")
	flag.Parse()
	Parse(configDir, secretDir)
}

func Parse(dirs ...string) {
	err := cfg.Load(&Instance, cfg.Dirs(dirs...), cfg.UseEnv("cfg"))
	zerolog.SetGlobalLevel(mapLogLevel[Instance.Logger.Level])
	if err != nil {
		log.Panic(err)
	}
}
