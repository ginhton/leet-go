# leetcode-go day002

## target
* learn basic of go -> done!
* start to read leetcode-cook and programming

## scratch
* start link: https://www.runoob.com/go/go-functions.html
* leetcode cookbook code repo: https://github.com/halfrost/LeetCode-Go/tree/master/leetcode
* leetcode problems: https://leetcode.com/problemset/all/

## notes

* 作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。
* 在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。

* 在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。

* Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑

* 形式参数会作为函数的局部变量来使用。

* 数组声明语法：`var variable_name [SIZE] variable_type`
* 多维数组的声明：  `var varialbe_name [SIZE1][SIZE2]...[SIZEN] variable_type`
* go中，空指针的值为nil
* 访问结构体成员，需要使用点号"."操作符
* 使用结构体指针访问结构体成员，使用"."操作符
* Go 语言切片是对数组的抽象
* Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
* Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
* go 不支持隐式转换类型
* Go 语言通过内置的错误接口提供了非常简单的错误处理机制。 error类型是一个接口类型
* Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。 goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
* Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。
* 通道（channel）是用来传递数据的一个数据结构。 通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
* 默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。
* 通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小
* 带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。 不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。
* 注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
