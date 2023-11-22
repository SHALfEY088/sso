// internal/grpc/auth/server.go
package authgrpc

import (
	"context"
	ssov1 "github.com/SHALfEY088/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer // Хитрая штука, о ней ниже
	//auth                          Auth
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	return &ssov1.LoginResponse{
		Token: req.GetEmail(),
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
