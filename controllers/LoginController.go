package controllers

import (
	"github.com/ahmedsharyo/SystemContorellerService/login_grpc"
	"github.com/ahmedsharyo/SystemContorellerService/models"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) string {
	var credential models.User
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no correct data found in the payload"
	}
	isUserAuthenticated := login_grpc.Login(credential.Email, credential.Password)
	if isUserAuthenticated {
		return login_grpc.GenerateToken(credential.Email, true)

	}
	return ""
}
