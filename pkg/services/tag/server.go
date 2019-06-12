package tag

import (
	"google.golang.org/grpc"
	"openpitrix.io/tag/pkg/config"
	"openpitrix.io/tag/pkg/manager"
	"openpitrix.io/tag/pkg/pb"
)

type Server struct {
}

func Serve() {
	cfg := config.GetInstance()
	s := &Server{}

	manager.NewGrpcServer(cfg.App.ApiHost, cfg.App.ApiPort).
		ShowErrorCause(cfg.Grpc.ShowErrorCause).
		WithChecker(s.Checker).
		Serve(func(server *grpc.Server) {
			pb.RegisterTagServer(server, s)
		})
}
