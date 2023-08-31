package handler

import (
	"github.com/gaogao-asia/golang-template/internal/di"
	"github.com/gaogao-asia/golang-template/pkg/connection"
	"github.com/gin-gonic/gin"
)

type newRouterParams struct {
	v1   *gin.RouterGroup
	Conn connection.Conn
}

// @BasePath /api/v1
func NewV1(api *gin.RouterGroup, DB connection.Conn) *newRouterParams {
	v1 := api.Group("/v1")
	return &newRouterParams{
		v1:   v1,
		Conn: DB,
	}
}

func (r *newRouterParams) Register() {
	r.registerAccount()
}

func (r *newRouterParams) registerAccount() {
	accountHandler := di.InitAccountHandler(r.Conn.DB)

	account := r.v1.Group("/accounts")
	{
		account.GET("", accountHandler.GetAccounts)
		account.POST("", accountHandler.CreateAccount)
	}
}
