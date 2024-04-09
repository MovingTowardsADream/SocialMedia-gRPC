package main

import (
	"fmt"
	"test-gRPC/internal/read_config"
)

func main() {
	cfg := read_config.OpenConfig()
	fmt.Println(cfg)
}
