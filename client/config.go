package main

import (
	"fmt"
	"os"
)

type config struct {
	serverAddr string
}

func newConfig() (*config, error) {
	serverAddr := os.Getenv("SVR_ADDR")
	if serverAddr == "" {
		return nil, fmt.Errorf("SVR_ADDR is required")
	}

	return &config{
		serverAddr: serverAddr,
	}, nil
}
