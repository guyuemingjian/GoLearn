package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

/**
关键字
下面列举了 Go 代码中会使用到的 25 个关键字或保留字：

break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：

append	bool	byte	cap	close	complex	complex64	complex128	uint16
copy	false	float32	float64	imag	int	int8	int16	uint32
int32	int64	iota	len	make	new	nil	panic	uint64
print	println	real	recover	string	true	uint	uint8	uintptr
*/
func main() {
	//f1()
	//f2()
	//f3()
	//f4()
	//f5(6, 7)
	f6()
}

/**
1	布尔型
布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。
2	数字类型
整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
3	字符串类型:
字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。
4	派生类型:
包括：
(a) 指针类型（Pointer）
(b) 数组类型
(c) 结构化类型(struct)
(d) Channel 类型
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map 类型
*/

func f1() {
	/**
	1	uint8
	无符号 8 位整型 (0 到 255)
	2	uint16
	无符号 16 位整型 (0 到 65535)
	3	uint32
	无符号 32 位整型 (0 到 4294967295)
	4	uint64
	无符号 64 位整型 (0 到 18446744073709551615)
	5	int8
	有符号 8 位整型 (-128 到 127)
	6	int16
	有符号 16 位整型 (-32768 到 32767)
	7	int32
	有符号 32 位整型 (-2147483648 到 2147483647)
	8	int64
	有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

	1	float32
	IEEE-754 32位浮点型数
	2	float64
	IEEE-754 64位浮点型数
	3	complex64
	32 位实数和虚数
	4	complex128
	64 位实数和虚数

	1	byte
	类似 uint8
	2	rune
	类似 int32
	3	uint
	32 或 64 位
	4	int
	与 uint 一样大小
	5	uintptr
	无符号整型，用于存放一个指针


	*/
	//无符号
	var auint uint8 = 255 //0-255
	var buint uint16 = 2  //0-2^16
	var cuint uint32 = 2  // 0 -2^32
	var duint uint64 = 2  // 0-2^64
	//有符号
	var aint int8 = -128 //(-2^7) - (2^7-1)
	var bint int16 = -2  //(-2^15) - (2^15-1)
	var cint int32 = -2  //(-2^31) - (2^31-1)
	var dint int64 = -2  //(-2^63) - (2^63-1)

	println(auint)
	buint = buint<<16 - 1
	println(buint)
	println(unsafe.Sizeof(bint))
	println(reflect.TypeOf(buint))
	cuint = cuint<<32 - 1
	println(cuint)
	duint = duint<<64 - 1
	println(duint)

	println(aint)
	bint = bint<<15 - 1
	println(bint)
	cint = cint << 31
	println(cint)
	dint = dint << 64
	println(dint)
}

/**
算术运算符
        下表列出了所有Go语言的算术运算符。假定 A 值为 10，B 值为 20。

运算符	描述	实例
+	相加	A + B 输出结果 30
-	相减	A - B 输出结果 -10
*	相乘	A * B 输出结果 200
/	相除	B / A 输出结果 2
%	求余	B % A 输出结果 0
++	自增	A++ 输出结果 11
--	自减	A-- 输出结果 9
*/
func f2() {
	var wkA = 10
	var wkB = 20

	println(wkA + wkB)
	println(wkA - wkB)
	println(wkA * wkB)
	println(wkA / wkB)
	println(wkA % wkB)
	wkA++
	println(wkA)
	wkA--
	println(wkA)
}

/**
关系运算符
        下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。

运算符	描述	实例
==	检查两个值是否相等，如果相等返回 True 否则返回 False	(A == B) 为 False
!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False	(A != B) 为 True
>	检查左边值是否大于右边值，如果是返回 True 否则返回 False	(A > B) 为 False
<	检查左边值是否小于右边值，如果是返回 True 否则返回 False	(A < B) 为 True
>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False	(A >= B) 为 False
<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False	(A <= B) 为 True
*/
func f3() {
	var wkA = 10
	var wkB = 20

	println(wkA == wkB)
	println(wkA != wkB)
	println(wkA > wkB)
	println(wkA < wkB)
	println(wkA >= wkB)
	println(wkA <= wkB)
}

/**
逻辑运算符
        下表列出了所有Go语言的逻辑运算符。假定 A 值为 True，B 值为 False。

运算符	描述	实例
&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False	(A && B) 为 False
||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False	(A || B) 为 True
!	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True	!(A && B) 为 True
*/

/**
位运算符
        位运算符对整数在内存中的二进制位进行操作。

        下表列出了位运算符 &， |，和 ^ 的计算：

p q p&q	p|q	p^q
0	0	0	0	0
0	1	0	1	1
1	1	1	1	0
1	0	0	1	1
*/
func f4() {
	var wkA = true
	var wkB = false

	println(wkA && wkB)
	println(wkA || wkB)
	println(!(wkA && wkB))

	println(0 & 0)
	println(0 | 0)
	println(0 ^ 0)

	println(0 & 1)
	println(0 | 1)
	println(0 ^ 1)

	println(1 & 0)
	println(1 | 0)
	println(1 ^ 0)

	println(1 & 1)
	println(1 | 1)
	println(1 ^ 1)
}

func f5(wkA int, wkB int) bool {

	wkC := strconv.Itoa(wkA)
	println(wkC)

	if wkA == wkB {
		println(wkA)
	}
	if wkA < wkB {
		println(wkB)
	}

	wkD := []string{"a", "b", "c"}
	for i := 0; i < len(wkD); i++ {
		println(wkD[i])

		switch wkD[i] {
		case "b":
			println("2")
			break
		default:
			println("wk")

		}
	}

	var ip *string
	ip = &wkC //ip 得到 wkC地址
	println(ip)
	println(*ip)
	*ip="abcd"
	println(wkC)


	return false
}


type Books struct {
	title string
	author string
	subject string
	book_id int
}

func f6()  {

	var Book1 Books        /* Declare Book1 of type Book */
	var Book2 Books        /* Declare Book2 of type Book */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(&Book1)

	/* 打印 Book2 信息 */
	printBook(&Book2)
}

func printBook( book *Books ) {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}
