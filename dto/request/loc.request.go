package request

type LocationProvinceParams struct {
	SearchQuery string `form:"search_query"`
}

type LocationRegencyByProvinceParams struct {
	ProvinceID  int32  `form:"province_id" binding:"required"`
	SearchQuery string `form:"search_query"`
}

type LocationDistrictByRegencyParams struct {
	RegencyID   int32  `form:"regency_id" binding:"required"`
	SearchQuery string `form:"search_query"`
}
