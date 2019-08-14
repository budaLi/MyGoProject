package main

import (
	"fmt"
	"strings"
)

//使用type改写
type caseFuc func(string) string

func main() {
	s := "ascsasacsac"
	fmt.Println(StringToCase2(s, trans))
}

func trans(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			//在这里要注意字符并不等于字符串 需要将字符串中的字符用string做强制转换
			result += strings.ToUpper(string(value))
		} else { //此处else只能在这里
			result += strings.ToLower(string(value))
		}
	}
	return result
}

//函数作为值
func StringToCase(str string, myfunc func(string) string) string {
	return myfunc(str)
}

func StringToCase2(str string, myfunc caseFuc) string {
	return myfunc(str)
}
