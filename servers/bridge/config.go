package bridge

type Config struct {
	// Path to the database file.
	DBPath string

	// URI of the cloud gRPC server
	CloudGRPCURI string
}
