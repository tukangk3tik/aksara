package response

type ProvinceResponse struct {
	ID        int32 `json:"id"`
	Name      string `json:"name"`
}

type RegencyResponse struct {
	ID        int32 `json:"id"`
	Name      string `json:"name"`
}

type DistrictResponse struct {
	ID        int32 `json:"id"`
	Name      string `json:"name"`
}

type ProvinceListResponse struct {
	TotalItems int64
	Items      []ProvinceResponse
}

type RegencyListResponse struct {
	TotalItems int64
	Items      []RegencyResponse
}

type DistrictListResponse struct {
	TotalItems int64
	Items      []DistrictResponse
}
