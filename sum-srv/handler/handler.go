package handler

import (
	"context"
	"go-micro-learning/micro-demo/proto/sum"
	"go-micro-learning/micro-demo/sum-srv/service"
)

// 私有
type handler struct {
}

func (h handler) GetSum(ctx context.Context, request *sum.SumRequest, response *sum.SumResponse) error {

	inputs := make([]int64, 0)
	var i int64 = 1
	for ; i < request.Input; i++ {
		inputs = append(inputs, i)
	}

	response.Output = service.GetSum(inputs...)

	return nil

	//panic("implement me")
}

//外部暴露handler
func Handler() sum.SumHandler {
	return handler{}

}
