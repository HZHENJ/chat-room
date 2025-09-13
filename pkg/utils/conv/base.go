package conv

import (
	"chatting-room/pkg/errno"
	"errors"
)

type HertzBaseResponse struct {
	StatusCode int32
	StatusMsg  string
}

func ToHertzBaseResponse(err error) *HertzBaseResponse {
	if err == nil {
		return &HertzBaseResponse{
			StatusCode: errno.SuccessCode,
			StatusMsg:  errno.SuccessMsg,
		}
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return &HertzBaseResponse{
			StatusCode: e.ErrCode,
			StatusMsg:  e.ErrMsg,
		}
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return &HertzBaseResponse{
		StatusCode: s.ErrCode,
		StatusMsg:  s.ErrMsg,
	}
}
