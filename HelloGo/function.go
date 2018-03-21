package main

import (
	"os"
	"fmt"
)

var a = 6

func main() {
	p()
	q1()
	p()
	q2()
	p()
}

func  p()  {
	println(a)
}

func q1() {
	a := 5
	println(a)
}

func q2() {
	a = 4
	println(a)
}

//返回值命名
func nextInt(b []byte, pos int) (value, nextPos int) {
	/**
	do something
	 */
}

//延迟提交
func readWrite(file *os.File) bool {
	file.Readdir(5)
	defer file.Close()
	var failureX = false
	if failureX {
		return false  //close自动调用
	}
	return true //close自动调用
}

//延迟的函数
func f() (ret int) {
	defer func() {
		ret ++
	}() //延迟函数，这里括号是必须的
	return 0
}

//不定变量
func myFunc(arg ...int) {
	myFunc(arg...)
	myFunc(arg[:2]...)

}

//函数作为值
func test() {
	aa := func() {  //定义匿名函数并赋值给aa
		println("Hello")
	}               //这里没有括号
	aa()            //调用函数
	fmt.Printf("%T\n",aa)
}

var xs = map[int]func() int{
	1: func() int {
		return 10
	},
	2: func() int {
		return 20
	},
	3: func() int {
		return 30
	},    //此处必须要有逗号
}

//回调
func printit(x int) {
	fmt.Printf("%v\n",x)
}
func callback(y int, f func(int)) {  //f将会保存函数
	f(y)                           //调用回调函数f，输入变量y
}

//恐慌与恢复
//Panic 内建函数 函数f调用Panic，函数f会被中断，但是延（defer)迟函数会被正常执行,然后f函数返回到调用的地方
//Recover 内建函数  正常执行程序，调用Recover函数会返回nil并且没有其他效果；如果陷入恐慌，
// 调用Recover可以捕捉到Panic的输入值，并且恢复正常执行
func throwPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}
