package main

import (
	"fmt"
	"os"
	"strconv"
)

type config struct {
	basePath string
	port int
}

func newConfig() (*config, error) {
	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		return nil, fmt.Errorf("BASE_PATH is required")
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	return &config{
		basePath: basePath,
		port: port,
	}, nil
}
