package api

import (
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"

	fake "github.com/brianvoe/gofakeit/v5"
	"github.com/gin-gonic/gin"

	"github.com/kzmake/traefik-custom-forward-auth/pkg/logger"
)

// RequestRouter は request に関する router の定義です。
type RequestRouter struct{}

// NewRequestRouter は router を生成します。
func NewRequestRouter() *RequestRouter {
	return &RequestRouter{}
}

// VerifyRequest はリクエストを検証します。
func (r *RequestRouter) VerifyRequest(c *gin.Context) {
	dump, _ := httputil.DumpRequest(c.Request, true)
	logger.Infof("request: %s", string(dump))

	rand.Seed(time.Now().UnixNano())

	if rand.Intn(100) < 50 {
		c.Header("X-User-Id", fake.Username())
		c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	} else {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"code":    "error code",
			"message": "error message",
		})
	}
}
