package accountpresent

import (
	"net/http"

	"github.com/ricky7171/te-marketplace/internal/error_handler"
	"github.com/ricky7171/te-marketplace/internal/library_wrapper"
	accountappservice "github.com/ricky7171/te-marketplace/internal/modules/account/application/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	authenticationService accountappservice.AuthenticationService
}

func NewHandler(authenticationService accountappservice.AuthenticationService) *Handler {
	return &Handler{
		authenticationService: authenticationService,
	}
}

func (h *Handler) HandleLogin(ctx library_wrapper.MyGinContext) {
	defer error_handler.HandleHttpError(ctx, "Failed to login")
	// bind gin request to object request and validate it
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// run AuthenticationService
	token, refreshToken, user, err := h.authenticationService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// arrange response
	result := LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User: UserData{
			CreatedAt: user.TimeStampLog.CreatedAt.String(),
			Email:     user.Email,
		},
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})

}
