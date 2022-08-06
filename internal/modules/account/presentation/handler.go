package accountpresent

import (
	"net/http"

	accountappservice "github.com/ricky7171/te-marketplace/internal/modules/account/application/service"

	accountdom "github.com/ricky7171/te-marketplace/internal/modules/account/domain"

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

func (h *Handler) HandleLogin(ctx *gin.Context) {
	// bind gin request to object request and validate it
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// convert to entity
	credential := accountdom.Credential{
		Email:    req.Email,
		Password: req.Password,
	}

	// run AuthenticationService
	result, err := h.authenticationService.Login(credential)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})

}
