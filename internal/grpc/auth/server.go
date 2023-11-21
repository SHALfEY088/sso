// internal/grpc/auth/server.go
package auth

import (
	"context"
	"google.golang.org/grpc"

	// Сгенерированный код
	ssov1 "github.com/SHALfEY088/protos/gen/go/sso"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer // Хитрая штука, о ней ниже
	auth                          Auth
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	panic("implement me")
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
) (*ssov1.IsAdminRequest, error) {
	panic("implement me")
}

//func (s *serverAPI) Login(
//	ctx context.Context,
//	in *ssov1.LoginRequest,
//) (*ssov1.LoginResponse, error) {
//	if in.Email == "" {
//		return nil, status.Error(codes.InvalidArgument, "email is required")
//	}
//
//	if in.Password == "" {
//		return nil, status.Error(codes.InvalidArgument, "password is required")
//	}
//
//	if in.GetAppId() == 0 {
//		return nil, status.Error(codes.InvalidArgument, "app_id is required")
//	}
//
//	token, err := s.auth.Login(ctx, in.GetEmail(), in.GetPassword(), int(in.GetAppId()))
//	if err != nil {
//		if errors.Is(err, auth.ErrInvalidCredentials) {
//			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
//		}
//
//		return nil, status.Error(codes.Internal, "failed to login")
//	}
//
//	return &ssov1.LoginResponse{Token: token}, nil
//}
//
//func (s *serverAPI) Register(
//	ctx context.Context,
//	in *ssov1.RegisterRequest,
//) (*ssov1.RegisterResponse, error) {
//	if in.Email == "" {
//		return nil, status.Error(codes.InvalidArgument, "email is required")
//	}
//
//	if in.Password == "" {
//		return nil, status.Error(codes.InvalidArgument, "password is required")
//	}
//
//	uid, err := s.auth.RegisterNewUser(ctx, in.GetEmail(), in.GetPassword())
//	if err != nil {
//		if errors.Is(err, storage.ErrUserExists) {
//			return nil, status.Error(codes.AlreadyExists, "user already exists")
//		}
//
//		return nil, status.Error(codes.Internal, "failed to register user")
//	}
//
//	return &ssov1.RegisterResponse{UserId: uid}, nil
//}
