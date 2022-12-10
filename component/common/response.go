package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Status interface{} `json:"status"`
	Msg    interface{} `json:"messages"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter, status, messages interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging, Filter: filter, Status: status, Msg: messages}
}

func SimpleSuccessResponse(data interface{}, messages string) *successResponse {
	return NewSuccessResponse(data, nil, nil, "success", messages)
}
