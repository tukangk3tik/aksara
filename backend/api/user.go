package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{http.StatusBadRequest, err.Error()})
		return
	}

	result, err := server.sm.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse{http.StatusUnauthorized, err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{
		http.StatusOK, "success", LoginUserResponse{
			AccessToken:  result.AccessToken,
			RefreshToken: result.RefreshToken,
		}})

	/*
		user, err := server.store.GetUser(ctx, req.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		err = utils.CheckPassword(req.Password, user.HashedPassword)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken, accessPayload, err := server.tokenMaker.CreateToken(
			user.Username,
			server.config.AccessTokenDuration,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
			user.Username,
			server.config.RefreshTokenDuration,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
			ID:           refreshPayload.ID,
			Username:     user.Username,
			RefreshToken: refreshToken,
			UserAgent:    ctx.Request.UserAgent(),
			ClientIp:     ctx.ClientIP(),
			IsBlocked:    false,
			ExpiresAt:    refreshPayload.ExpiredAt,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		response := loginUserResponse{
			SessionId:             session.ID,
			AccessToken:           accessToken,
			AccessTokenExpiresAt:  accessPayload.ExpiredAt,
			RefreshToken:          refreshToken,
			RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
			User:                  newUserResponse(user),
		}
	*/

	// ctx.JSON(http.StatusOK, response)
}
