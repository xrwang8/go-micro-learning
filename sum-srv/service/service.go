package service

import "github.com/nacos-group/nacos-sdk-go/common/logger"

//getSum的具体逻辑,累加

func GetSum(intputs ...int64) int64 {

	var ret int64
	for _, v := range intputs {
		ret += v
	}
	logger.Infof("累加结果为：%v\n", ret)
	return ret

}
