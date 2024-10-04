package wsmanager

import (
	stringsRes "first_socket/internal/res/strings"
	"first_socket/pkg/ws_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WSManager struct {
	wsService wsservice.IWSService
}

func (manager *WSManager) CreateConnection(
	ctx *gin.Context,
) {
	ctxLogin, loginExists := ctx.Get(stringsRes.LoginHeaderKey)
	if !loginExists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login type"})
		return
	}

	_, err := ctxLogin.(string)
	if !err {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login type"})
		return
	}

	connErr := manager.wsService.ServeWS(ctx.Writer, ctx.Request, nil)

	if connErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": connErr.Error()})
		return
	}
}

func NewWSManager(wsService wsservice.IWSService) *WSManager {
	return &WSManager{
		wsService,
	}
}
