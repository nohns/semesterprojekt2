package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func ReadEnvfile() {
	//Read the .env file
	cwdEnvPath, err := filepath.Abs(".env")
	if err == nil {
		err = godotenv.Load(cwdEnvPath)
		if err == nil {
			log.Println("Loaded .env file")
			return
		}
	}
	err = godotenv.Load("../../.env")
	if err != nil {
		log.Println("Loading .env file failed, using production environment")
	}
	if err == nil {
		log.Println("Loaded .env file")
	}
}

type Config struct {
	DBPath        string
	CloudGRPCURI  string
	BridgeGRPCURI string
}

func LoadConfFromEnv() (*Config, error) {

	// DB path
	dbpath, err := stringEnvVar("DB_PATH")
	if err != nil {
		return nil, fmt.Errorf("could not read DB_PATH: %v", err)
	}
	// gRPC cloud server address
	cloudUri, err := stringEnvVar("CLOUD_GRPC_URI")
	if err != nil {
		return nil, fmt.Errorf("could not read GRPC_ADDR: %v", err)
	}
	bridgeUri, err := stringEnvVar("BRIDGE_GRPC_URI")
	if err != nil {
		return nil, fmt.Errorf("could not read GRPC_ADDR: %v", err)
	}

	return &Config{
		DBPath:        dbpath,
		CloudGRPCURI:  cloudUri,
		BridgeGRPCURI: bridgeUri,
	}, nil
}

func stringEnvVar(envname string) (string, error) {
	val, ok := os.LookupEnv(envname)
	if !ok {
		return "", fmt.Errorf("missing env var '%s' string value", envname)
	}

	return val, nil
}
