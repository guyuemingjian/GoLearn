package main

import (
	"fmt"
	"time"
)

const str1  = "常量定义"
var str2 = "全局变量定义"

//一般变量声明
type intA int

//结构体
type structA struct {

}

//接口
type interfaceA interface {

}

func main()  {

	fmt.Println("hello world")
	fmt.Print(str1);
	fmt.Println("adkjf %s ahjkjdf" ,str2)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05.999999999  Z07:00"))//2019-06-07 10:30:10.6917183  +08:00
}
