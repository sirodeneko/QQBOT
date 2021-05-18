package coolq

import "fmt"

// 生成@某人的cq码 当qq为小于等于0时 为@全体
// name:当在群中找不到此QQ号的名称时才会生效
func At(userid int64, name string) string {
	if userid <= 0 {
		return "[CQ:at,qq=all]"
	} else {
		return fmt.Sprintf("[CQ:at,qq=%v,name=%v]", userid, name)
	}
}
