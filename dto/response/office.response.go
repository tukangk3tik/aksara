package response

type OfficeResponse struct {
	ID         int64  `json:"id"`
	Index      string `json:"index"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Province   string `json:"province"`
	Regency    string `json:"regency"`
	District   string `json:"district"`
	ProvinceID int64  `json:"province_id"`
	RegencyID  int64  `json:"regency_id"`
	DistrictID int64  `json:"district_id"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	LogoURL    string `json:"logo_url"`
	CreatedBy  int64  `json:"created_by"`
}

type OfficeListResponse struct {
	TotalItems int64
	Items      []OfficeResponse
}
