package main

import (
	"fmt"

	"github.com/gaogao-asia/golang-template/internal/dto"
	"github.com/gaogao-asia/golang-template/internal/entity"
	"github.com/gaogao-asia/golang-template/pkg/connection"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := connection.GetConnection()

	r := gin.Default()
	r.GET("/accounts", func(c *gin.Context) {
		// get all account from database
		var accounts []entity.Account
		err := conn.Find(&accounts).Error
		if err != nil {
			panic(err)
		}

		fmt.Println("all account: ", accounts)

		var resp []dto.GetAccountResponse
		for _, account := range accounts {
			resp = append(resp, dto.GetAccountResponse{
				ID:    account.ID,
				Name:  account.Name,
				Email: account.Email,
			})
		}

		c.JSON(200, resp)
	})

	r.POST("/accounts", func(c *gin.Context) {
		req := dto.CreateAccountRequest{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(400, gin.H{
				"msg_code": "CREATE_ACCOUNT_REQUEST_INVALID",
				"msg":      err.Error(),
			})
			return
		}

		// create account in database
		account := entity.Account{
			Name:  req.Name,
			Email: req.Email,
		}
		err = conn.Create(&account).Error
		if err != nil {
			c.JSON(500, gin.H{
				"msg_code": "CREATE_ACCOUNT_ERROR",
				"msg":      err.Error(),
			})
			return
		}

		c.JSON(200, dto.GetAccountResponse{
			ID:    account.ID,
			Name:  account.Name,
			Email: account.Email,
		})
	})

	r.Run(":3011")
}
