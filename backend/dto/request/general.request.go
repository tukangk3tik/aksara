package request

type Pagination struct {
	Page  int32 `form:"page,default=1"`
	Limit int32 `form:"limit,min=5,max=100,default=10"`
}
