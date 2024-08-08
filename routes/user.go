package routes

import (
	"go-event-api/models"
	"go-event-api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = user.Save()

	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user was created", "user": user})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		log.Print(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}
	log.Print(user)
	jwtToken, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		log.Print("LOGIN ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": jwtToken})

}
