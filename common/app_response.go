package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

type authenRes struct {
	Data   interface{} `json:"data"`
	AccessToken interface{} `json:"accessToken"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}

func BasicSuccessResponse(data interface{}) interface{} {
	return data
}

func AuthenSuccessResponse(data,accessToken interface{}) *authenRes {
	return &authenRes{Data: data,AccessToken: accessToken}
}