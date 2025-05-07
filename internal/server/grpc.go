package server

import (
	"context"
	v2 "github.com/WH-5/push-service/api/push/v1"
	"github.com/WH-5/push-service/internal/conf"
	"github.com/WH-5/push-service/internal/middleware"
	"github.com/WH-5/push-service/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, pushService *service.PushService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(selector.Server(
			middleware.AuthCheckExist(pushService),
		).Match(func(ctx context.Context, operation string) bool {
			if strings.HasSuffix(operation, "Push/PushMsg") {
				return false
			}
			return true
		}).Build(),
			recovery.Recovery(),
			logging.Server(logger),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v2.RegisterPushServer(srv, pushService)
	return srv
}
