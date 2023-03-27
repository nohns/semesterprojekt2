package middleware

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

func LoggingMiddlewareGrpc(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Printf("gRPC method: %s", info.FullMethod)
	resp, err = handler(ctx, req)
	return
}



// Logging middleware // Logs IP address of client, what method the request is and the URL path
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}