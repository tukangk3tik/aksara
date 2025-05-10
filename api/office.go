package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/dto/request"
	"github.com/tukangk3tik/aksara/dto/response"
	"github.com/tukangk3tik/aksara/utils"
	"go.uber.org/zap"
)

func (server *Server) getOffices(ctx *gin.Context) {
	var req request.Pagination
	log := utils.FromContext(ctx.Request.Context())
	log = log.With(zap.String("func", "getOffices"))

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse("BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	offset := (req.Page - 1) * req.Limit
	arg := db.ListAllOfficesParams{
		Limit:  req.Limit,
		Offset: offset, 
	}

	offices, err := server.store.ListAllOffices(ctx, &arg)
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	index := offset + 1
	items := []response.OfficeResponse{}
	for _, item := range offices {
		itemI := parseOfficeModelToResponse(item)
		itemI.Index = fmt.Sprintf("#%d", index)
		items = append(items, itemI)
		index++
	}

	totalItems, err := server.store.TotalListAllOffices(ctx)
	if err != nil {
		log.Error(zap.Error(err).String)
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	ctx.JSON(http.StatusOK, response.DataTableResponse{
		Message: "success",
		Data:    items,
		MetaData: response.DataTableMetaData{
			CurrentPage: req.Page,
			PerPage:     req.Limit,
			TotalItems:  totalItems,
		}})
}

func (server *Server) createOffice(ctx *gin.Context) {
	var req request.CreateOfficeRequest
	log := utils.FromContext(ctx.Request.Context())
	log = log.With(zap.String("func", "createOffice"))

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse("BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	userID := ctx.MustGet("user_id").(uint64)
	traceID := ctx.MustGet("trace_id").(string)

	createParams := db.CreateOfficeParams{
		Code:       req.Code,
		Name:       req.Name,
		ProvinceID: req.ProvinceID,
		RegencyID:  req.RegencyID,
		DistrictID: req.DistrictID,
		Email:      req.Email,
		CreatedBy:  int64(userID),
	}

	newOffice, err := server.store.CreateOffice(ctx, &createParams)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			log.Warn(err.Error())
			ctx.JSON(http.StatusBadRequest, response.BuildTrxErrorResponse(traceID, "DUPLICATE_ENTRY", utils.ErrorCodeMap["DUPLICATE_ENTRY"], nil))
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.BuildTrxErrorResponse(traceID, "INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	ctx.JSON(http.StatusCreated, response.TrxSuccessResponse{TraceID: traceID, Data: map[string]any{"id": newOffice.ID}})
}

func (server *Server) updateOffice(ctx *gin.Context) {
	traceID := ctx.MustGet("trace_id").(string)
	log := utils.FromContext(ctx.Request.Context())
	log = log.With(zap.String("func", "updateOffice"))

	var params request.OfficeIDPathParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, response.BuildTrxErrorResponse(traceID, "BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	var req request.UpdateOfficeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, response.BuildTrxErrorResponse(traceID, "BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	findOffice, err := server.store.GetSchoolById(ctx, int64(params.ID))
	if err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusNotFound, response.BuildTrxErrorResponse(traceID, "NOT_FOUND", utils.ErrorCodeMap["NOT_FOUND"], nil))
		return
	}

	updateParams := db.UpdateSchoolParams{
		Name:       req.Name,
		ProvinceID: findOffice.ProvinceID,
		RegencyID:  findOffice.RegencyID,
		DistrictID: findOffice.DistrictID,
		Email:      findOffice.Email,
		Phone:      findOffice.Phone,
		Address:    findOffice.Address,
		LogoUrl:    findOffice.LogoUrl,
		ID:         int64(params.ID),
	}

	if req.Phone != "" {
		updateParams.Phone = sql.NullString{
			String: req.Phone,
			Valid:  true,
		}
	}

	if req.Address != "" {
		updateParams.Address = sql.NullString{
			String: req.Address,
			Valid:  true,
		}
	}

	if req.LogoURL != "" {
		updateParams.LogoUrl = sql.NullString{
			String: req.LogoURL,
			Valid:  true,
		}
	}

	_, err = server.store.UpdateSchool(ctx, &updateParams)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			log.Warn(err.Error())
			ctx.JSON(http.StatusBadRequest, response.BuildTrxErrorResponse(traceID, "DUPLICATE_ENTRY", utils.ErrorCodeMap["DUPLICATE_ENTRY"], nil))
			return
		}

		ctx.JSON(http.StatusInternalServerError, response.BuildTrxErrorResponse(traceID, "INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	ctx.JSON(http.StatusOK, response.TrxSuccessResponse{TraceID: traceID, Data: map[string]any{"id": findOffice.ID}})
}

func (server *Server) deleteOffice(ctx *gin.Context) {
	var params request.OfficeIDPathParams
	log := utils.FromContext(ctx.Request.Context())
	log = log.With(zap.String("func", "deleteOffice"))

	if err := ctx.ShouldBindUri(&params); err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, response.BuildTrxErrorResponse("", "BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	traceID := ctx.MustGet("trace_id").(string)

	res, err := server.store.DeleteSchool(ctx, int64(params.ID))
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.BuildTrxErrorResponse(traceID, "INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	rowsA, err := res.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.BuildTrxErrorResponse(traceID, "INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	if rowsA == 0 {
		log.Warn("Office not found")
		ctx.JSON(http.StatusNotFound, response.BuildTrxErrorResponse(traceID, "NOT_FOUND", utils.ErrorCodeMap["NOT_FOUND"], nil))
		return
	}

	ctx.JSON(http.StatusOK, response.TrxSuccessResponse{TraceID: traceID, Data: map[string]any{"id": params.ID}})
}

func parseOfficeModelToResponse(model db.ListAllOfficesRow) response.OfficeResponse {
	return response.OfficeResponse{
		ID:        int64(model.ID),
		Code:      model.Code,
		Name:      model.Name,
		Province:  model.Province,
		Regency:   model.Regency,
		District:  model.District.String,
		Email:     model.Email,
		Phone:     model.Phone.String,
		Address:   model.Address.String,
		LogoURL:   model.LogoUrl.String,
		CreatedBy: model.CreatedBy,
	}
}
