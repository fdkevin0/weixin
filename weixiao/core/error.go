package core

import "fmt"

func ErrorHandler(code int) {
	errMap := map[int]string{
		40001: "app_key或app_secret不合法",
		40002: "access_token过期",
		40003: "wxcode不合法",
		40005: "access_token无效",
		40006: "school_code不存在",
		40007: "电子卡号不存在",
		40008: "权限不足",
		40009: "参数缺失",
		40010: "参数错误",
		40018: "该主体没有开启应用",
	}
	fmt.Println(errMap[code])
}
