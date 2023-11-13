package models

type PaginationQuery struct {
	Page int `json:"page" form:"page" binding:"min=1"`
	Size int `json:"size" form:"size" binding:"min=0,max=40"`
}

type Paged[T any] struct {
	Cnt  int `json:"cnt"`
	Data []T `json:"data"`
}

func NewPaged[T any](cnt int, data []T) Paged[T] {
	return Paged[T]{
		Cnt:  cnt,
		Data: data,
	}
}
