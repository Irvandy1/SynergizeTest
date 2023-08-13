package delivery

import (
	"SynergizeTest/models"
	"SynergizeTest/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var form models.User
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Register(ctx, form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Successfuly registered")
}

func Login(ctx *gin.Context) {
	var form models.LoginForm
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Login(ctx, form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Login Success")
}

func Logout(ctx *gin.Context) {
	var form models.LogoutForm
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Logout(ctx, form)
	ctx.JSON(http.StatusOK, "Login Success")
}

func RegisterBank(ctx *gin.Context) {
	var form models.Bank
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.RegisterBank(ctx, form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Bank Successfully Registered")
}

func TopUp(ctx *gin.Context) {
	var form models.TopUpForm
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.TopUp(ctx, form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "TopUp Successfull")
}

func GetUserDetail(ctx *gin.Context) {
	var form models.LogoutForm
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := user.GetUserDetail(ctx, form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func GetUserList(ctx *gin.Context) {
	var form models.Pagination
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := user.GetUserList(ctx, form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
