### golang产生随机数,方面到处调用

用法：

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


输出：

	纯数字 166302
	小写字母 bsgneq
	大写字母 DCFTTS
	混合 hNN4fN

再运行一次：

	纯数字 716708
	小写字母 xncquj
	大写字母 CKNBDM
	混合 D84yW2


更复杂的验证码,请用：

[https://github.com/dchest/captcha](https://github.com/dchest/captcha)


