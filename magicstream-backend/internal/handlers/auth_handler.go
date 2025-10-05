package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rishi-0007/magicstream-backend/internal/services"
	"github.com/rishi-0007/magicstream-backend/internal/utils"
)

type AuthHandler struct { svc *services.AuthService }
func NewAuthHandler(s *services.AuthService) *AuthHandler { return &AuthHandler{svc: s} }

type registerReq struct { Email string `json:"email" validate:"required,email"`; Name string `json:"name" validate:"required,min=2,max=50"`; Password string `json:"password" validate:"required,min=6"` }

func (h *AuthHandler) Register(c *gin.Context){
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil { utils.Error(c,http.StatusBadRequest,"invalid payload"); return }
	if err := utils.Validate.Struct(req); err != nil { utils.Error(c,http.StatusBadRequest,err.Error()); return }
	u, err := h.svc.Register(c, req.Email, req.Name, req.Password)
	if err != nil { utils.Error(c,http.StatusBadRequest,err.Error()); return }
	utils.OK(c, gin.H{"user": gin.H{"email": u.Email, "name": u.Name}})
}

type loginReq struct { Email string `json:"email" validate:"required,email"`; Password string `json:"password" validate:"required"` }

func (h *AuthHandler) Login(c *gin.Context){
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil { utils.Error(c,http.StatusBadRequest,"invalid payload"); return }
	if err := utils.Validate.Struct(req); err != nil { utils.Error(c,http.StatusBadRequest,err.Error()); return }
	access, refresh, u, err := h.svc.Login(c, req.Email, req.Password)
	if err != nil { utils.Error(c,http.StatusUnauthorized,"invalid credentials"); return }
	utils.OK(c, gin.H{"access_token": access, "refresh_token": refresh, "user": gin.H{"email": u.Email, "name": u.Name}})
}

func (h *AuthHandler) Refresh(c *gin.Context){
	var body struct{ RefreshToken string `json:"refresh_token" validate:"required"` }
	if err := c.ShouldBindJSON(&body); err != nil { utils.Error(c,http.StatusBadRequest,"invalid payload"); return }
	if err := utils.Validate.Struct(body); err != nil { utils.Error(c,http.StatusBadRequest,err.Error()); return }
	access, err := h.svc.Refresh(body.RefreshToken)
	if err != nil { utils.Error(c,http.StatusUnauthorized,"invalid refresh token"); return }
	utils.OK(c, gin.H{"access_token": access})
}
