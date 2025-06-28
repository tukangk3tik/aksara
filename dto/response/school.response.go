package response

type SchoolResponse struct {
	ID             int64  `json:"id"`
	Index          string `json:"index"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	IsPublicSchool bool   `json:"is_public_school"`
	Office         string `json:"office"`
	OfficeID       int64  `json:"office_id"`
	Province       string `json:"province"`
	Regency        string `json:"regency"`
	District       string `json:"district"`
	ProvinceID     int64  `json:"province_id"`
	RegencyID      int64  `json:"regency_id"`
	DistrictID     int64  `json:"district_id"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	LogoURL        string `json:"logo_url"`
	CreatedBy      int64  `json:"created_by"`
}

type SchoolListResponse struct {
	TotalItems int64
	Items      []SchoolResponse
}
