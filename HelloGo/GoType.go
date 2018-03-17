package main

import "fmt"

func main() {
	//未被使用的变量会报错
	//布尔类型
	var boola bool = true
	var boolb bool = !boola
	fmt.Println(boolb)

	//整数基本类型：int、int8、int16、int32、int64
	//byte (uint8) uint16 uint32 uint64
	// float32 float64 没有float类型
	var a  int
	var bb int32
	a = 8
	bb = int32 (a + a) //不同类型运算出错，需要强制转换
	bb = bb + 5


	const (
		consta = iota
		constb = iota
	)
	const (
		a1 = iota //0
		b1		  //1
	)

	const ( //明确指出常量类型
		a2 = 0
		b string = "0"
	)

	//var s string = "this is test"
	//s[0] = 'c'

	s:="hello"
	c:=[]rune(s)  //转换数组
	c[0]='c'
	s2:=string(c)  //建立新字符串
	fmt.Printf("%s\n",s2)

	//错误写法
	//s := "Starting part "
	//	+ "Ending part"
	//正确写法
	ss := "Starting part " +
		"Ending part"

	fmt.Println(ss)

	//rune 是 int32 的别名
	//在需要遍历字符串中字符时使用

	//复数 complex128 (64位虚数部分) 还有 complex64 complex32 两种类型
	var cc complex64 = 5 + 5i;fmt.Printf("Value is : %v\n", cc)

	//错误
	//var e error //error 是接口

	//特殊字符
	// + - * / %
	//& | ^ &^ 按位与 按位或 按位异或 位清除
	//&& ||  ! 逻辑与  逻辑或 逻辑非
	//





}
