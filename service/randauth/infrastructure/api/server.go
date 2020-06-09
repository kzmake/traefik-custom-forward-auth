package api

import (
	"github.com/gin-gonic/gin"
	cli "github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/web"

	"github.com/kzmake/traefik-custom-forward-auth/pkg/constant"
)

var (
	service = constant.Service.RandAuth
	version = "v0.1.0"
)

// Server は server アプリケーションです。
type Server interface {
	Run() error
}

// New は server を生成します。
func New() Server {
	service := web.NewService(
		web.Name("api."+service),
		web.Version(version),

		web.Address("0.0.0.0:3000"),
	)
	if err := service.Init(web.Action(func(c *cli.Context) {})); err != nil {
		return nil
	}

	request := NewRequestRouter()
	router := gin.Default()
	router.Any("/request", request.VerifyRequest)

	// Register Handler
	service.Handle("/", router)

	return service
}
