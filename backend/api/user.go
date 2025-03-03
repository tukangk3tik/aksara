package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tukangk3tik/aksara/dto/response"
	"github.com/tukangk3tik/aksara/utils"
	"go.uber.org/zap"
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
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse("BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	traceID := ctx.MustGet("trace_id").(string)

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, response.BuildErrorResponse("NOT_FOUND", utils.ErrorCodeMap["NOT_FOUND"], nil))
			return
		}
		server.logger.Error(utils.LogErrorMessageBuilder("trx failed to get user", traceID), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		server.logger.Error(utils.LogErrorMessageBuilder("trx failed to check password", traceID), zap.Error(err))
		ctx.JSON(http.StatusUnauthorized, response.BuildErrorResponse("WRONG_PASSWORD", utils.ErrorCodeMap["WRONG_PASSWORD"], nil))
		return
	}

	accessToken, _, err := server.tokenMaker.GenerateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		server.logger.Error(utils.LogErrorMessageBuilder("trx failed to generate access token", traceID), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	refreshToken, _, err := server.tokenMaker.GenerateToken(user.ID, server.config.RefreshTokenDuration)
	if err != nil {
		server.logger.Error(utils.LogErrorMessageBuilder("trx failed to generate refresh token", traceID), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse{
		Data: LoginUserResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
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
