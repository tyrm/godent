package logic

import (
	"net/http"

	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/models"
)

type Auth interface {
	RequireAuth(r *http.Request) (*models.Token, gdhttp.ErrCode, string)
}
