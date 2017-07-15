package main

import (
	"fmt"
	"github.com/zouhuigang/captcha"
)

func main() {

	str := captcha.RandStr(6, captcha.NUM)
	fmt.Println("纯数字", str)

	str1 := captcha.RandStr(6, captcha.LOWER)
	fmt.Println("小写字母", str1)

	str2 := captcha.RandStr(6, captcha.UPPER)
	fmt.Println("大写字母", str2)

	str3 := captcha.RandStr(6, captcha.ALL)
	fmt.Println("混合", str3)
}
