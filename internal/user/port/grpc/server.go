package grpc

import (
	"github.com/quangdangfit/gocommon/validation"
	"google.golang.org/grpc"

	"main/internal/user/repository"
	"main/internal/user/service"
	"main/pkg/dbs"
	pb "main/proto/gen/go/user"
)

func RegisterHandlers(svr *grpc.Server, db dbs.IDatabase, validator validation.Validation) {
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(validator, userRepo)
	userHandler := NewUserHandler(userSvc)

	pb.RegisterUserServiceServer(svr, userHandler)
}
