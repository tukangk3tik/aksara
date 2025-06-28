package api

import (
	"database/sql"
	"errors"
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

func (server *Server) getSchools(ctx *gin.Context) {
	var req request.Pagination
	log := utils.FromContext(ctx.Request.Context())
	log = log.With(zap.String("func", "getSchools"))

	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse("BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	offset := (req.Page - 1) * req.Limit
	arg := db.ListAllSchoolsParams{
		Limit:  req.Limit,
		Offset: offset,
	}

	schools, err := server.store.ListAllSchools(ctx, &arg)
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	index := offset + 1
	items := []response.SchoolResponse{}
	for _, item := range schools {
		itemI := parseSchoolRowModelToResponse(item)
		itemI.Index = fmt.Sprintf("#%d", index)
		items = append(items, itemI)
		index++
	}

	totalItems, err := server.store.TotalListAllSchools(ctx)
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

func (server *Server) getSchoolById(ctx *gin.Context) {
	var params request.SchoolIDPathParams
	log := utils.FromContext(ctx.Request.Context())
	log = log.With(zap.String("func", "getSchoolById"))

	if err := ctx.ShouldBindUri(&params); err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse("BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}
	traceID := ctx.MustGet("trace_id").(string)

	school, err := server.store.GetSchoolById(ctx, int64(params.ID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, response.BuildErrorResponse("NOT_FOUND", utils.ErrorCodeMap["NOT_FOUND"], nil))
			return
		}

		log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	ctx.JSON(http.StatusOK, response.TrxSuccessResponse{TraceID: traceID, Data: parseGetSchoolByIdRowModelToResponse(school)})
}

func (server *Server) createSchool(ctx *gin.Context) {
	var req request.CreateSchoolRequest
	log := utils.FromContext(ctx.Request.Context())
	log = log.With(zap.String("func", "createSchool"))

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse("BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	traceID := ctx.MustGet("trace_id").(string)
	userID := ctx.MustGet("user_id").(uint64)

	school, err := server.store.CreateSchool(ctx, &db.CreateSchoolParams{
		Code:           req.Code,
		Name:           req.Name,
		IsPublicSchool: req.IsPublicSchool,
		OfficeID:       sql.NullInt64{Int64: int64(userID), Valid: true},
		ProvinceID:     req.ProvinceID,
		RegencyID:      req.RegencyID,
		DistrictID:     req.DistrictID,
		Email:          sql.NullString{String: req.Email, Valid: req.Email != ""},
		Phone:          sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Address:        sql.NullString{String: req.Address, Valid: req.Address != ""},
		LogoUrl:        sql.NullString{String: req.LogoURL, Valid: req.LogoURL != ""},
		CreatedBy:      int64(userID),
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case db.ForeignKeyViolation, db.UniqueViolation:
				constraintName := pqErr.Constraint
				fieldName := ""
				errorMsg := ""

				switch constraintName {
				case "schools_code_key":
					fieldName = "code"
					errorMsg = fmt.Sprintf(utils.ErrorCodeMap["DUPLICATE_ENTRY"], "Kode")
				case "schools_email_key":
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
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	ctx.JSON(http.StatusOK, response.TrxSuccessResponse{TraceID: traceID, Data: parseSchoolModelToResponse(school)})
}

func (server *Server) updateSchool(ctx *gin.Context) {
	traceID := ctx.MustGet("trace_id").(string)
	log := utils.FromContext(ctx.Request.Context())
	log = log.With(zap.String("func", "updateSchool"))

	var params request.SchoolIDPathParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, response.BuildTrxErrorResponse(traceID, "BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	var req request.UpdateSchoolRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, response.BuildTrxErrorResponse(traceID, "BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	findSchool, err := server.store.GetSchoolById(ctx, int64(params.ID))
	if err != nil {
		log.Warn(err.Error())
		errorMsg := fmt.Sprintf(utils.ErrorCodeMap["NOT_FOUND"], "Sekolah")
		ctx.JSON(http.StatusNotFound, response.BuildTrxErrorResponse(traceID, "NOT_FOUND", errorMsg, nil))
		return
	}

	updateParams := db.UpdateSchoolParams{
		Name:       req.Name,
		OfficeID:   findSchool.OfficeID,
		ProvinceID: findSchool.ProvinceID,
		RegencyID:  findSchool.RegencyID,
		DistrictID: findSchool.DistrictID,
		Email:      sql.NullString{String: req.Email, Valid: req.Email != ""},
		Phone:      sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Address:    sql.NullString{String: req.Address, Valid: req.Address != ""},
		LogoUrl:    sql.NullString{String: req.LogoURL, Valid: req.LogoURL != ""},
		ID:         int64(params.ID),
	}

	school, err := server.store.UpdateSchool(ctx, &updateParams)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case db.ForeignKeyViolation, db.UniqueViolation:
				constraintName := pqErr.Constraint
				fieldName := ""
				errorMsg := ""

				if constraintName == "schools_code_key" {
					fieldName = "code"
					errorMsg = fmt.Sprintf(utils.ErrorCodeMap["DUPLICATE_ENTRY"], "Kode")
				} else if constraintName == "schools_email_key" {
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

	ctx.JSON(http.StatusOK, response.TrxSuccessResponse{TraceID: traceID, Data: parseSchoolModelToResponse(school)})
}

func (server *Server) deleteSchool(ctx *gin.Context) {
	var params request.SchoolIDPathParams
	log := utils.FromContext(ctx.Request.Context())
	log = log.With(zap.String("func", "deleteSchool"))

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
		log.Warn("School not found")
		errorMsg := fmt.Sprintf(utils.ErrorCodeMap["NOT_FOUND"], "Sekolah")
		ctx.JSON(http.StatusNotFound, response.BuildTrxErrorResponse(traceID, "NOT_FOUND", errorMsg, nil))
		return
	}

	ctx.JSON(http.StatusOK, response.TrxSuccessResponse{TraceID: traceID, Data: map[string]any{"id": params.ID}})
}

func parseSchoolRowModelToResponse(model db.ListAllSchoolsRow) response.SchoolResponse {
	return response.SchoolResponse{
		ID:             int64(model.ID),
		Code:           model.Code,
		Name:           model.Name,
		IsPublicSchool: model.IsPublicSchool,
		Office:         model.Office,
		OfficeID:       int64(model.OfficeID.Int64),
		Province:       model.Province,
		Regency:        model.Regency,
		District:       model.District.String,
		ProvinceID:     int64(model.ProvinceID),
		RegencyID:      int64(model.RegencyID),
		DistrictID:     int64(model.DistrictID),
		Email:          model.Email.String,
		Phone:          model.Phone.String,
		Address:        model.Address.String,
		LogoURL:        model.LogoUrl.String,
		CreatedBy:      model.CreatedBy,
	}
}

func parseGetSchoolByIdRowModelToResponse(model db.GetSchoolByIdRow) response.SchoolResponse {
	return response.SchoolResponse{
		ID:             int64(model.ID),
		Code:           model.Code,
		Name:           model.Name,
		IsPublicSchool: model.IsPublicSchool,
		Office:         model.Office,
		OfficeID:       int64(model.OfficeID.Int64),
		Province:       model.Province,
		Regency:        model.Regency,
		District:       model.District.String,
		ProvinceID:     int64(model.ProvinceID),
		RegencyID:      int64(model.RegencyID),
		DistrictID:     int64(model.DistrictID),
		Email:          model.Email.String,
		Phone:          model.Phone.String,
		Address:        model.Address.String,
		LogoURL:        model.LogoUrl.String,
		CreatedBy:      model.CreatedBy,
	}
}

func parseSchoolModelToResponse(model db.Schools) response.SchoolResponse {
	return response.SchoolResponse{
		ID:             int64(model.ID),
		Code:           model.Code,
		Name:           model.Name,
		IsPublicSchool: model.IsPublicSchool,
		OfficeID:       int64(model.OfficeID.Int64),
		ProvinceID:     int64(model.ProvinceID),
		RegencyID:      int64(model.RegencyID),
		DistrictID:     int64(model.DistrictID),
		Email:          model.Email.String,
		Phone:          model.Phone.String,
		Address:        model.Address.String,
		LogoURL:        model.LogoUrl.String,
		CreatedBy:      model.CreatedBy,
	}
}
