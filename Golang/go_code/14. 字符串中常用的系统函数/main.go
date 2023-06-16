package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	// 统计字符串的长度，按字节 len(str)
	// golang的编码统一为 utf-8 (ascii的字符，（字母和数字）占一个字节，汉字占用3个字节)

	str := "hello"
	fmt.Println("str length=", len(str)) // 5
	str2 := "hello北"
	fmt.Println("str2 length=", len(str2)) // 8

	str3 := "hello杭州"
	// 字符串遍历，同时处理有中文的问题 r:=[]rune(str)
	r := []rune(str3)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 字符串转整数：
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换成功 n = ", n)
	}

	// 整数转字符串 str = strconv.Itoa(123)
	str4 := strconv.Itoa(123)
	fmt.Printf("str =%v", str4)

	// 字符串 转 []byte
	bytes := []byte("hello golang")
	fmt.Printf("bytes  = %v \n", bytes)

	// []byte 转 字符串
	str5 := string([]byte{104, 101, 108, 108, 111, 32, 103, 111, 108, 97, 110, 103})
	fmt.Printf("str5 = %v \n", str5)

	// 10 进制转 2，8，16 进制
	str6 := strconv.FormatInt(123, 2)
	fmt.Printf("123 的 二进制是%v \n", str6)

	// 查找字符串中是否存在指定字符 // strings.Contains(s, substr) boolean
	str7 := strings.Contains("charles dong", "dong")
	fmt.Printf("是否存在其中 = %v \n", str7)

	// 统计一个字符串中有几个指定子串， //strings.Count(s, substr) int
	str8 := strings.Count("charles dong aa", "a")
	fmt.Printf("str8 统计有多少个 a = %v \n", str8)

	// 不区分大小写的字符串比较 (== 是区分字母大小写的)
	str9 := strings.EqualFold("abc", "ABC")
	fmt.Printf("abc 是否等于 ABC =%v \n", str9)
	fmt.Printf("abc == ABC = %v \n", "abc" == "ABC")

	// 返回指定字符，在字符串中第一次出现的下标， 如果没有就返回 -1
	// strings.Index("charles", "ar") 3
	str10 := strings.Index("charles", "ar")
	fmt.Printf("charles 中 ar 的下标是 %v \n", str10)

	// 返回 指定字符，在字符串中最后一次出现的下标，如果没有返回 -1
	str11 := strings.LastIndex("charlies dong", "n")
	fmt.Printf("acharles dong 最后一次出现的 a 的下标 是 %v \n", str11)

	// 将 指定字符，替换成另外一个字符
	// 可以指定替换几个， 如果是 -1 表示替换全部
	str12 := strings.Replace("charles pei", "pei", "dong", -1)
	fmt.Printf("charles pei， 替换成 %v\n", str12)

	// 将字符串根据指定字符分割成数组 string.Split
	str13 := strings.Split("charles,dong", ",")
	fmt.Printf("charles,dong 分割后 = %v \n", str13)

	// 将字符串的字母进行大小写的替换
	// strings.ToLower("") strings.ToUpper
	str14 := strings.ToLower("Charles Dong")
	str15 := strings.ToUpper("Charles Dong")
	fmt.Printf("Charles Dong ToLower替换成 %v \n", str14)
	fmt.Printf("Charles Dong ToUpper替换成 %v \n", str15)

	// 将字符串左右两边空格删除
	str16 := strings.TrimSpace(" a charles dong")
	fmt.Printf(" trim  去除空格后的%v \n", str16)

	// 将字符串左右两边指定的字符去除
	str17 := strings.Trim("!a he!llo c ! ", " !ac")
	fmt.Printf("! hello ! 去除两边指定字符后 结果是 %v\n", str17)
	// 将字符串左边指定字符去掉
	// strings.TrimLeft
	// 将字符右边指定字符去掉
	// strings.TrimRight

	// 判断字符串是否以指定字符串开头
	//strings.HasPrefix("","") bool
	str18 := strings.HasPrefix("https://xxx.xxx.com", "https")
	fmt.Printf("是否是 https 开头的，结果是 %v \n", str18)

	// 判断字符串是否以指定字符结尾
	// strings.HasSuffix("","") bool
	str19 := strings.HasSuffix("https://.xxx.xxx.com", "com")
	fmt.Printf("是否是 com 结尾的，结果是 %v \n", str19)
}
