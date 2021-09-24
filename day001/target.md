# leet-go day001
## target
* learn the basics of go
* config go grogramming env


## notes
* 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）

* 需要注意的是 { 不能单独放在一行，否则代码在运行时会产生错误

* 标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

* Go 代码中会使用到的 25 个关键字或保留字:
break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var

* Go 语言还有 36 个预定义标识符：
append	bool	byte	cap	close	complex	complex64	complex128	uint16
copy	false	float32	float64	imag	int	int8	int16	uint32
int32	int64	iota	len	make	new	nil	panic	uint64
print	println	real	recover	string	true	uint	uint8	uintptr

* 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
* 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。

* 派生类型包括： (a) 指针类型（Pointer） (b) 数组类型 (c) 结构化类型(struct) (d) Channel 类型 (e) 函数类型 (f) 切片类型 (g) 接口类型（interface） (h) Map 类型

* 变量声明第一种，指定变量类型，如果没有初始化，则变量默认为零值。
* 第二种，根据值自行判定变量类型
* 第三种，使用 := 声明变量. 这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。

* 如果你声明了一个局部变量却没有在相同的代码块中使用它，会得到编译错误. 全局变量是允许声明但不使用的。

* 交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同。

* 空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。

* _ 实际上是一个只写(write-only)变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。

* 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
```go
const identifier [type] = value
```

* 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：

* iota，特殊常量，可以认为是一个可以被编译器修改的常量。 iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。 iota 可以被用作枚举值 

* Go 语言内置的运算符有： 算术运算符 关系运算符 逻辑运算符 位运算符 赋值运算符 其他运算符

* switch 可以判断变量的类型（int， string， or other）

* switch 中使用fallthrough会强制执行后面的case语句，fallthrough不会判断下一条case表达式的结果是否为true

* select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。 select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

* for循环 1. for init;condition;post {} 2. for condition {} 3. for {} 4. for key,value:=range oldMap {}

* 默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。 

* Go 语言可以很灵活的创建函数，并作为另外一个函数的实参。
