package logic

import (
	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/models"
	"net/http"
)

type Auth interface {
	RequireAuth(r *http.Request) (*models.Account, gdhttp.ErrCode, string)
}
