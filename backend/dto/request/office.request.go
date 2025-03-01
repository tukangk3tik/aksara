package request

type CreateOfficeRequest struct {
	Code       string `json:"code" binding:"required"`
	Name       string `json:"name" binding:"required"`
	ProvinceID int32  `json:"province_id" binding:"required"`
	RegencyID  int32  `json:"regency_id" binding:"required"`
	DistrictID int32  `json:"district_id" binding:"required"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	LogoURL    string `json:"logo_url"`
}

type OfficeIDPathParams struct {
	ID uint64 `uri:"id" binding:"required,min=1"`
}

type UpdateOfficeRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	LogoURL string `json:"logo_url"`
}
