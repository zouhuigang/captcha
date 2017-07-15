package captcha

import (
	"math/rand"
	"time"
)

type StrType int

var fontKinds = [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}

const (
	NUM   StrType = 0 // 数字
	LOWER         = 1 // 小写字母
	UPPER         = 2 // 大写字母
	ALL           = 3 // 全部
)

//num长度,types类型：captcha.NUM,captcha.LOWER
func RandStr(num int, types StrType) string {
	if num <= 0 {
		num = 4
	}

	str := string(randStr(num, int(types)))
	return str
}

// 生成随机字符串
// size 个数 kind 模式
func randStr(size int, kind int) []byte {
	ikind, result := kind, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll {
			ikind = rand.Intn(3)
		}
		scope, base := fontKinds[ikind][0], fontKinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}

	return result
}
