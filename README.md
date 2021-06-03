s# Go

## Table of Contents

- [1. Brainstorming](#1-brainstorming)
- [2. Basic](#2-basic)
	- [2.1 Execution](#21-execution)
	- [2.2 Variables](#22-variables)
		- [2.2.1 Value Types](#221-value-types)
		- [2.2.2 Reference Types](#222-reference-types)
	- [2.3 Control Structures](#23-control-structures)
- [References](#references)

## 1. Brainstorming

> No unnecessary code!

<div align="center"> <img src="go.svg" width="100%"/> </div><br>

## 2. Basic

### 2.1 Execution

**The order of execution**

<div align="center"> <img src="pics/go-program.png" width="70%"/> </div><br>

1. 入口：*package main*，按顺序 *import* 所有包
2. 递归地 *import*（每个包只会 *import* 一次）
3. 所有 *const* 和 *variables* 将被赋值，调用 *init()*（若存在）
4. 执行 *main()*

### 2.2 Variables

#### 2.2.1 Value Types

**Primitives**

<div align="center"> <img src="pics/image-20210528121045494.png" width="35%"/> </div><br>



#### 2.2.2 Reference Types

**赋值**


<div align="center"> <img src="pics/image-20210528121947711.png" width="50%"/> </div><br>

**var.go**

```go
// global, package 层面
var s string

// 只能 function 层面
s := "hello, world"

// 声明相同类型的变量
var a, b, c int
```



**pointer**

- Working w/ *pointers* can reduce **memory usage** and increase efficiency

<div align="center"> <img src="pics/image-20210529151649826.png" width="45%"/> </div><br>



### 2.3 Control Structures

#### 2.3.1 If else

**if-else**

```go
if condition1 {
  ...
} else if condition2 {
  ...
} else {
  ...
}
```

**if-else w/ intialization**

```go
// value 只在 if-else 块中可见
if value := process(data); value > max {
  ...
}
```
**testing for errors**

```go
// 初始化 & 赋值
value, err := pack1.Function1(param1)
if err != nil {
  // 错误处理
  fmt.Println("error msg")
  return err
}
// normal case
```

#### 2.3.2 Switch

**switch**

```go
switch var1 {
case var1, var2:
  ...
case vare3:
  ...
default:
  ...
}
```

**multi-branches**

```go
switch {
case condition:
  ...
case condition2:
  ...
default:
  ...
}
```

#### 2.3.3 For

**for-loop**

```go
for init; condition; modif {
  ...
}
```

**range**

```go
for i, val := range coll {
  // val 是 copy
}
```





## References

- *The Way to Go*
- *The GO Programming Language*
- [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)
- [字符编码笔记：ASCII，Unicode 和 UTF-8](https://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html)