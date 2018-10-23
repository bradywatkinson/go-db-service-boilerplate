// common/router.go

package common

import (
  "crypto/x509"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials"
  "github.com/grpc-ecosystem/go-grpc-middleware"
  "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

// InitializeGRPC sets up the grpc server
func (a *App) InitializeGRPC(certPool *x509.CertPool, addr string) {
  opts := []grpc.ServerOption{
    grpc.Creds(credentials.NewClientTLSFromCert(certPool, addr)),
    grpc_middleware.WithUnaryServerChain(
      grpc_recovery.UnaryServerInterceptor(),
    ),
  }

  a.GRPC = grpc.NewServer(opts...)
}
