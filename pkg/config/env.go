package config

import "github.com/codingconcepts/env"

type e struct {
	Address   string `env:"REPLIK_ADDRESS"`
	Port      int    `env:"REPLIK_PORT" default:"9090"`
	ChunkSize int    `env:"REPLIK_CHUNK_SIZE" default:"65536"`
}

func loadEnv() e {
	var e e
	err := env.Set(&e)
	if err != nil {
		panic(err)
	}
	return e
}

var Env = loadEnv()
