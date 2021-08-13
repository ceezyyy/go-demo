# Go

## Table of Contents

- [1. Basic](#1-basic)
	- [1.1 Execution](#11-execution)
	- [1.2 Value v.s Pointer](#12-value-vs-pointer)
	- [1.3 Slices](#13-slices)
	- [1.4 Maps](#14-maps)
	- [1.5 Structs](#15-structs)
- [2. Concurrency](#2-concurrency)
- [3. Web](#3-web)
- [References](#references)

## 1. Basic

### 1.1 Execution

**The order of execution**

<div align="center"> <img src="pics/go-program.png" width="70%"/> </div><br>

1. 入口：*package main*，按顺序 *import* 所有包
2. 递归地 *import*（每个包只会 *import* 一次）
3. 所有 *const* 和 *variables* 将被赋值，调用 *init()*（若存在）
4. 执行 *main()*

**init 函数**

```go
var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
    return 42
}

func init() {
    WhatIsThe = 0
}

func main() {
    if WhatIsThe == 0 {
        fmt.Println("It's all a lie.")
    }
}
```





### 1.2 Value v.s Pointer

**value**

<div align="center"> <img src="pics/image-20210528121045494.png" width="35%"/> </div><br>

**pointer**

- Working w/ *pointers* can reduce **memory usage** and increase efficiency

<div align="center"> <img src="pics/image-20210529151649826.png" width="45%"/> </div><br>



### 1.3 Slices

**Arrays**

<div align="center"> <img src="./pics/slice-array.png" width="50%"/> </div><br>

**Slices 底层结构**

<div align="center"> <img src="pics/slice-struct.png" width="55%"/> </div><br>



<div align="center"> <img src="./pics/slice-1.png" width="50%"/> </div><br>



<div align="center"> <img src="./pics/slice-2.png" width="50%"/> </div><br>

**扩容**

```go
t := make([]byte, len(s), (cap(s)+1)*2)  // 思考: 为什么 "+1"
copy(t, s)
s = t
```



**设置 len 和 cap 相同**

```go
// 1. append 后不会修改底层数组
// 2. 保持为切片申请新的底层数组的简洁
newSlice := originSlice[i:j:j]
```

### 1.4 Maps





### 1.5 Structs

**底层**

<div align="center"> <img src="pics/image-20210609115944249.png" width="65%"/> </div><br>

**嵌套**

<div align="center"> <img src="pics/image-20210609150224083.png" width="60%"/> </div><br>



**继承**

```go
// 同一个包下, 子类可以访问父类的字段和方法
type Child struct {
		Father, // 匿名类
  	Mother, // 匿名类
}
```



## 2. Concurrency

**Goroutines**

In *Go*, each **concurrently executing activity** is called a *goroutine*

**Channels**

<div align="center"> <img src="pics/image-20210614170020244.png" width="55%"/> </div><br>





## References

- *The Way to Go*
- *The GO Programming Language*
- *Go Web Programming*
- *Go in Action*
- [Go by Example](https://gobyexample.com/)
- [Go Slices: usage and internals](https://blog.golang.org/slices-intro)
- [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)
- [字符编码笔记：ASCII，Unicode 和 UTF-8](https://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html)
- [Go 语言设计与实现](https://draveness.me/golang/)
- [Go语言 | Go 1.9 新特性 Type Alias详解](https://www.flysnow.org/2017/08/26/go-1-9-type-alias.html)
- [Type assertions and type switches](https://yourbasic.org/golang/type-assertion-switch/)
- [Go maps in action](https://blog.golang.org/maps)
- [When is the init() function run?](https://stackoverflow.com/questions/24790175/when-is-the-init-function-run)
