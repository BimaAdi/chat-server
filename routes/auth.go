package routes

import (
	"net/http"

	"github.com/BimaAdi/chat-server/core"
	"github.com/BimaAdi/chat-server/models"
	"github.com/BimaAdi/chat-server/repository"
	"github.com/BimaAdi/chat-server/schemas"
	"github.com/gin-gonic/gin"
)

func authRoutes(rq *gin.RouterGroup) {
	auths := rq.Group("/auth")

	auths.POST("/login", authLoginRoute)
}

// Login
//
//	@Summary		Login
//	@Description	login
//	@Tags			Auth
//	@Produce		json
//	@Param			payload	formData	schemas.LoginFormRequest	true	"form data"
//	@Success		200		{object}	schemas.LoginResponse
//	@Failure		400		{object}	schemas.BadRequestResponse
//	@Failure		500		{object}	schemas.InternalServerErrorResponse
//	@Router			/auth/login [post]
func authLoginRoute(c *gin.Context) {
	// Get data from form
	formRequest := schemas.LoginFormRequest{}
	err := c.ShouldBind(&formRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	// Get User
	user, err := repository.GetUserByUsername(models.DBConn, formRequest.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: "invalid credentials",
		})
		return
	}

	// Check Password
	if !core.CheckPasswordHash(formRequest.Password, user.Password) {
		c.JSON(http.StatusBadRequest, schemas.BadRequestResponse{
			Message: "invalid credentials",
		})
		return
	}

	// Generate JWT token
	token, err := core.GenerateJWTTokenFromUser(models.DBConn, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.InternalServerErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, schemas.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	})
}
