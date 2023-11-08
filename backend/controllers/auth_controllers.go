package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/yoshi-zen/sea-turtle/backend/models"
	"github.com/yoshi-zen/sea-turtle/backend/myerrors"
	"github.com/yoshi-zen/sea-turtle/backend/services"
)

type AuthController struct {
	db *sql.DB
}

func NewAuthController(db *sql.DB) *AuthController {
	return &AuthController{db: db}
}

func (c *AuthController) RegisterUserHandler(w http.ResponseWriter, req *http.Request) {
	var auth models.Auth
	if err := json.NewDecoder(req.Body).Decode(&auth); err != nil {
		err = myerrors.ReqDecodeFailed.Wrap(err, "failed to decode json request body")
		myerrors.ErrorHandler(w, req, err)
	}

	if err := services.RegisterUserService(c.db, &auth); err != nil {
		myerrors.ErrorHandler(w, req, err)
	}
}

func Login() {}
