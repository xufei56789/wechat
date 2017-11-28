package function

import (
	"strconv"
	"time"
)

//获取最新时间搓（1506418906886420539）
func GetNewTime() string {
	timeh := strconv.FormatInt(time.Now().UnixNano(),10)
	return timeh
}