package main

import "fmt"

//格式
//if x > 0 {
//	return y
//} else {
//	return x
//}

func main() {
	//forTest()
	//rangeTest()
	switchTest2(0)
}


func myfunc() {
	i := 0
Here:
	println(i)
	i++
	goto Here

}

//循环 for
//for init; confition; port { } //和 Java 的一般for循环一样
//for condition { } 相当于 while 循环
//for { } 死循环
func forTest() {
	//var sum int
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}

	//平行赋值
	a := "string"
	c:=[]rune(a)
	for i,j := 0, len(c)-1; i < j; i, j =  i + 1, j - 1 {
		c[i], c[j] = c[j], c[i]  //平行赋值
	}
	fmt.Println(string(c))

}

//迭代器
func rangeTest() {
	//array 和 slice 循环时 k是序号 v是值
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		fmt.Printf("k is: %d, v is: %s.\n", k, v)
	}

	for pos, char := range "a&x" {
		fmt.Printf("chaeacter '%c' starts at byte position %d\n", char, pos)
	}
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return  c -'A' + 10
	}
	return 0
}

//与 Java 不同的是匹配失败后不会自动向下尝试  但可以 使用 fallthrough 这么做
//都不匹配时可以使用 default
func switchTest2(i byte) {
	switch i {
	case 0: //空的 case 体
	case 1:
		fmt.Printf("aaaa")
	}

	switch i {
	case 0: fallthrough
	case 1:
		fmt.Printf("bbbbbbbbbbb")

	}
}

//比较两个字节数组字典数序先后，比较长度内都相同比较长度的长短
func Compare(a, b []byte) int {
	for i :=0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return  -1
		}
	}

	switch  {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}









