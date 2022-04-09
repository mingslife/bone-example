package conf

import (
	"flag"
	"log"
	"sync"
)

type Config struct {
	Host string
	Port int
}

var (
	cfg = &Config{}
	wg  = &sync.WaitGroup{}
)

func ParseConfig() *Config {
	wg.Add(1)
	defer wg.Done()

	flag.StringVar(&cfg.Host, "host", "127.0.0.1", "application host")
	flag.IntVar(&cfg.Port, "port", 8080, "application port")

	log.Println("Configuration loaded")
	return cfg
}

func GetConfig() *Config {
	wg.Wait()

	return cfg
}
