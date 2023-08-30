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
	r.registerProduct()
}

func (r *newRouterParams) registerProduct() {
	accountHandler := di.InitAccountHandler(r.Conn.DB)

	product := r.v1.Group("/accounts")
	{
		product.GET("", accountHandler.GetAccounts)
		product.POST("", accountHandler.CreateAccount)
	}
}
