package delivery

import "github.com/gin-gonic/gin"

func InitRouter() {
	router := gin.Default()
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/logout", Logout)
	router.POST("/topup", TopUp)
	router.POST("/register-bank", RegisterBank)
	router.GET("/user-detail", GetUserDetail)
	router.GET("/user-list", GetUserList)
	router.Run()
}
