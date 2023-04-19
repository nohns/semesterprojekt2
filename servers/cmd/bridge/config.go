package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"github.com/nohns/servers/bridge"
)

func readEnvfile() {
	//Read the .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Loading .env file failed, using production environment")
	}
	if err == nil {
		log.Println("Loaded .env file")
	}

}

func loadConfFromEnv() (*bridge.Config, error) {

	// DB path
	dbpath, err := stringEnvVar("BRIDGE_DB_PATH")
	if err != nil {
		return nil, fmt.Errorf("could not read BRIDGE_DB_PATH: %v", err)
	}

	// gRPC cloud server address
	grpcuri, err := stringEnvVar("CLOUD_GRPC_URI")
	if err != nil {
		return nil, fmt.Errorf("could not read GRPC_ADDR: %v", err)
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
