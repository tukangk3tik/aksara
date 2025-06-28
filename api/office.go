package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
		log.Warn(err.Error())
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
		itemI := parseOfficeRowModelToResponse(item)
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
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case db.ForeignKeyViolation, db.UniqueViolation:
				constraintName := pqErr.Constraint
				fieldName := ""
				errorMsg := ""

				if constraintName == "offices_code_key" {
					fieldName = "code"
					errorMsg = fmt.Sprintf(utils.ErrorCodeMap["DUPLICATE_ENTRY"], "Kode")
				} else if constraintName == "offices_email_key" {
					fieldName = "email"
					errorMsg = fmt.Sprintf(utils.ErrorCodeMap["DUPLICATE_ENTRY"], "Email")
				}

				errorData := []string{
					fieldName,
				}

				log.Warn(err.Error())
				ctx.JSON(http.StatusBadRequest, response.BuildTrxErrorResponse(traceID, "DUPLICATE_ENTRY", errorMsg, errorData))
				return
			}
		}

		log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.BuildTrxErrorResponse(traceID, "INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	ctx.JSON(http.StatusCreated, response.TrxSuccessResponse{TraceID: traceID, Data: parseOfficeModelToResponse(newOffice)})
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

	findOffice, err := server.store.GetOfficeById(ctx, int64(params.ID))
	if err != nil {
		log.Warn(err.Error())
		errorMsg := fmt.Sprintf(utils.ErrorCodeMap["NOT_FOUND"], "Kantor")
		ctx.JSON(http.StatusNotFound, response.BuildTrxErrorResponse(traceID, "NOT_FOUND", errorMsg, nil))
		return
	}

	updateParams := db.UpdateOfficeParams{
		Code:       findOffice.Code,
		Name:       req.Name,
		ProvinceID: findOffice.ProvinceID,
		RegencyID:  findOffice.RegencyID,
		DistrictID: findOffice.DistrictID,
		Email:      req.Email,
		Phone:      sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Address:    sql.NullString{String: req.Address, Valid: req.Address != ""},
		LogoUrl:    sql.NullString{String: req.LogoURL, Valid: req.LogoURL != ""},
		ID:         int64(params.ID),
	}

	office, err := server.store.UpdateOffice(ctx, &updateParams)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case db.ForeignKeyViolation, db.UniqueViolation:
				constraintName := pqErr.Constraint
				fieldName := ""
				errorMsg := ""

				if constraintName == "offices_code_key" {
					fieldName = "code"
					errorMsg = fmt.Sprintf(utils.ErrorCodeMap["DUPLICATE_ENTRY"], "Kode")
				} else if constraintName == "offices_email_key" {
					fieldName = "email"
					errorMsg = fmt.Sprintf(utils.ErrorCodeMap["DUPLICATE_ENTRY"], "Email")
				}

				errorData := []string{
					fieldName,
				}

				log.Warn(err.Error())
				ctx.JSON(http.StatusBadRequest, response.BuildTrxErrorResponse(traceID, "DUPLICATE_ENTRY", errorMsg, errorData))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, response.BuildTrxErrorResponse(traceID, "INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	ctx.JSON(http.StatusOK, response.TrxSuccessResponse{TraceID: traceID, Data: parseOfficeModelToResponse(office)})
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

	res, err := server.store.DeleteOffice(ctx, int64(params.ID))
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
		errorMsg := fmt.Sprintf(utils.ErrorCodeMap["NOT_FOUND"], "Kantor")
		ctx.JSON(http.StatusNotFound, response.BuildTrxErrorResponse(traceID, "NOT_FOUND", errorMsg, nil))
		return
	}

	ctx.JSON(http.StatusOK, response.TrxSuccessResponse{TraceID: traceID, Data: map[string]any{"id": params.ID}})
}

func (server *Server) fetchOfficesSelectOption(ctx *gin.Context) {
	var req request.SelectOptionRequest
	log := utils.FromContext(ctx.Request.Context())
	log = log.With(zap.String("func", "fetchOfficesSelectOption"))

	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse("BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	arg := db.ListOfficesWithFiltersParams{
		Name:   sql.NullString{String: fmt.Sprintf("%%%s%%", req.SearchQuery), Valid: req.SearchQuery != ""},
		Limit:  10,
		Offset: 0,
	}

	traceID := ctx.MustGet("trace_id").(string)

	offices, err := server.store.ListOfficesWithFilters(ctx, &arg)
	if err != nil {
		server.logger.Error(utils.LogErrorMessageBuilder("trx failed to get provinces", traceID), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	items := []response.SelectOptionResponse{}
	for _, item := range offices {
		items = append(items, parseOfficeModelToSelectOptionResponse(item))
	}

	ctx.JSON(http.StatusOK, response.DataTableResponse{
		Message: "success",
		Data:    items,
		MetaData: response.DataTableMetaData{
			CurrentPage: 1,
			PerPage:     10,
			TotalItems:  int64(len(items)),
		}})
}

func parseOfficeRowModelToResponse(model db.ListAllOfficesRow) response.OfficeResponse {
	return response.OfficeResponse{
		ID:         int64(model.ID),
		Code:       model.Code,
		Name:       model.Name,
		Province:   model.Province,
		Regency:    model.Regency,
		District:   model.District.String,
		ProvinceID: int64(model.ProvinceID),
		RegencyID:  int64(model.RegencyID),
		DistrictID: int64(model.DistrictID),
		Email:      model.Email,
		Phone:      model.Phone.String,
		Address:    model.Address.String,
		LogoURL:    model.LogoUrl.String,
		CreatedBy:  model.CreatedBy,
	}
}

func parseOfficeModelToResponse(model db.Offices) response.OfficeResponse {
	return response.OfficeResponse{
		ID:         int64(model.ID),
		Code:       model.Code,
		Name:       model.Name,
		ProvinceID: int64(model.ProvinceID),
		RegencyID:  int64(model.RegencyID),
		DistrictID: int64(model.DistrictID),
		Email:      model.Email,
		Phone:      model.Phone.String,
		Address:    model.Address.String,
		LogoURL:    model.LogoUrl.String,
		CreatedBy:  model.CreatedBy,
	}
}

func parseOfficeModelToSelectOptionResponse(model db.ListOfficesWithFiltersRow) response.SelectOptionResponse {
	return response.SelectOptionResponse{
		ID:   int64(model.ID),
		Name: model.Name,
		AdditionalData: map[string]any{
			"province": map[string]any{
				"id":   int64(model.ProvinceID),
				"name": model.Province,
			},
			"regency": map[string]any{
				"id":   int64(model.RegencyID),
				"name": model.Regency,
			},
			"district": map[string]any{
				"id":   int64(model.DistrictID),
				"name": model.District.String,
			},
		},
	}
}
