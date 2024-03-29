package config

type Env struct {
	Address   string `env:"REPLIK_ADDRESS"`
	Port      int    `env:"REPLIK_PORT" default:"9090"`
	ChunkSize int    `env:"REPLIK_CHUNK_SIZE" default:"65536"`
}
