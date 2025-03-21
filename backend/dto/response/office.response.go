package response

type OfficeResponse struct {
	ID        uint64 `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Province  string `json:"province"`
	Regency   string `json:"regency"`
	District  string `json:"district"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	LogoURL   string `json:"logo_url"`
	CreatedBy uint64 `json:"created_by"`
}

type OfficeListResponse struct {
	TotalItems int64
	Items      []OfficeResponse
}
