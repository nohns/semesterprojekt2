package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nohns/semesterprojekt2/bridge"
)

func loadConfFromEnv() (*bridge.Config, error) {

	// DB path
	dbpath, err := stringEnvVar("BRIDGE_DB_PATH")
	if err != nil {
		return nil, fmt.Errorf("could not read BRIDGE_DB_PATH: %v", err)
	}
	// gRPC cloud server address
	grpcuri, err := stringEnvVar("GRPC_CLOUD_URI")
	if err != nil {
		return nil, fmt.Errorf("could not read CLOUD_GRPC_URI: %v", err)
	}

	return &bridge.Config{
		DBPath:       dbpath,
		CloudGRPCURI: grpcuri,
	}, nil
}

func stringEnvVar(envname string) (string, error) {
	val, ok := os.LookupEnv(envname)
	if !ok {
		return "", fmt.Errorf("missing env var '%s' string value", envname)
	}

	return val, nil
}

func intEnvVar(envname string) (int, error) {
	val, err := strconv.Atoi(os.Getenv(envname))
	if err != nil {
		return 0, fmt.Errorf("could not read env var '%s' int value: %v", envname, err)
	}

	return val, nil
}

func boolEnvVar(envname string) (bool, error) {
	val := os.Getenv(envname)
	if val == "" {
		return false, fmt.Errorf("missing env var '%s' bool value", envname)
	}
	if strings.ToLower(val) == "true" {
		return true, nil
	}

	return false, nil
}
