package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/dto/request"
	"github.com/tukangk3tik/aksara/dto/response"
	"github.com/tukangk3tik/aksara/utils"
	"go.uber.org/zap"
)

func (server *Server) fetchProvinces(ctx *gin.Context) {
	var req request.LocationProvinceParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse("BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	arg := db.LocationProvinceParams{
		Name:   fmt.Sprintf("%%%s%%", req.SearchQuery),
		Limit:  10,
		Offset: 0,
	}

	traceID := ctx.MustGet("trace_id").(string)

	provinces, err := server.store.LocationProvince(ctx, arg)
	if err != nil {
		server.logger.Error(utils.LogErrorMessageBuilder("trx failed to get provinces", traceID), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	items := []response.ProvinceResponse{}
	for _, item := range provinces {
		items = append(items, parseProviceModelToResponse(item))
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

func (server *Server) fetchRegencyByProvince(ctx *gin.Context) {
	var req request.LocationRegencyByProvinceParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse("BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	arg := db.LocationRegencyByProvinceParams{
		Name:       fmt.Sprintf("%%%s%%", req.SearchQuery),
		ProvinceID: req.ProvinceID,
		Limit:      10,
		Offset:     0,
	}

	traceID := ctx.MustGet("trace_id").(string)

	regencies, err := server.store.LocationRegencyByProvince(ctx, arg)
	if err != nil {
		server.logger.Error(utils.LogErrorMessageBuilder("trx failed to get provinces", traceID), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	items := []response.RegencyResponse{}
	for _, item := range regencies {
		items = append(items, parseRegencyModelToResponse(item))
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

func (server *Server) fetchDistrictByRegency(ctx *gin.Context) {
	var req request.LocationDistrictByRegencyParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse("BAD_REQUEST", utils.ErrorCodeMap["BAD_REQUEST"], nil))
		return
	}

	arg := db.LocationDistrictByRegencyParams{
		Name:      fmt.Sprintf("%%%s%%", req.SearchQuery),
		RegencyID: req.RegencyID,
		Limit:     10,
		Offset:    0,
	}

	traceID := ctx.MustGet("trace_id").(string)

	districts, err := server.store.LocationDistrictByRegency(ctx, arg)
	if err != nil {
		server.logger.Error(utils.LogErrorMessageBuilder("trx failed to get provinces", traceID), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse("INTERNAL_SERVER_ERROR", utils.ErrorCodeMap["INTERNAL_SERVER_ERROR"], nil))
		return
	}

	items := []response.DistrictResponse{}
	for _, item := range districts {
		items = append(items, parseDistrictModelToResponse(item))
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

func parseProviceModelToResponse(model db.LocProvince) response.ProvinceResponse {
	return response.ProvinceResponse{
		ID:   model.ID,
		Name: model.Name,
	}
}

func parseRegencyModelToResponse(model db.LocRegency) response.RegencyResponse {
	return response.RegencyResponse{
		ID:   model.ID,
		Name: model.Name,
	}
}

func parseDistrictModelToResponse(model db.LocDistrict) response.DistrictResponse {
	return response.DistrictResponse{
		ID:   model.ID,
		Name: model.Name,
	}
}
