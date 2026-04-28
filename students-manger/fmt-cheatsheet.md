# Go fmt 格式化速查表

先记住一个总原则：

格式化字符串里，普通字符原样输出，只有从 `%` 开始的部分才是“格式说明”。

比如：

```go
fmt.Printf("name=%s age=%d\n", "tom", 18)
```

这里：

- `name=` 和 ` age=` 是普通文本
- `%s` 和 `%d` 是格式符
- `\n` 是换行

你可以把它理解成“模板 + 填空”。

## 一、最常用格式符

### 1. `%d`

表示十进制整数。

```go
fmt.Printf("%d\n", 18)   // 18
```

### 2. `%s`

表示字符串。

```go
fmt.Printf("%s\n", "hello")   // hello
```

### 3. `%f`

表示浮点数，默认会保留 6 位小数。

```go
fmt.Printf("%f\n", 3.14)   // 3.140000
```

### 4. `%.2f`

表示浮点数，保留 2 位小数。

```go
fmt.Printf("%.2f\n", 3.14159)   // 3.14
```

### 5. `%v`

通用格式，最常用，打印“值”。

```go
fmt.Printf("%v\n", 123)          // 123
fmt.Printf("%v\n", "hello")      // hello
fmt.Printf("%v\n", []int{1, 2})   // [1 2]
```

### 6. `%#v`

打印更完整、更接近 Go 代码的值，调试结构体时特别常用。

```go
type student struct {
	ID   int
	Name string
}

s := student{ID: 1, Name: "tom"}

fmt.Printf("%v\n", s)   // {1 tom}
fmt.Printf("%#v\n", s)  // main.student{ID:1, Name:"tom"}
```

### 7. `%T`

打印变量的类型。

```go
fmt.Printf("%T\n", 123)       // int
fmt.Printf("%T\n", "hello")   // string
```

### 8. `%t`

打印布尔值。

```go
fmt.Printf("%t\n", true)   // true
```

## 二、宽度和补位

这部分最容易混，但规律很固定。

看这个：

```go
%02d
```

拆开理解：

- `%`：开始格式说明
- `0`：不够宽度时用 `0` 补
- `2`：总宽度至少 2 位
- `d`：按十进制整数输出

所以：

```go
fmt.Printf("%02d\n", 1)   // 01
fmt.Printf("%02d\n", 9)   // 09
fmt.Printf("%02d\n", 10)  // 10
```

这就是代码里 `stu%02d` 的意思：

```go
fmt.Sprintf("stu%02d", 3)   // stu03
```

再看几个变化：

### 1. `%2d`

宽度至少 2 位，不够时前面补空格。

```go
fmt.Printf("%2d\n", 1)   //  1
fmt.Printf("%2d\n", 10)  // 10
```

### 2. `%03d`

宽度至少 3 位，不够时前面补 0。

```go
fmt.Printf("%03d\n", 1)   // 001
fmt.Printf("%03d\n", 12)  // 012
```

### 3. `%-5s`

字符串宽度至少 5 位，左对齐，右边补空格。

```go
fmt.Printf("%-5s!\n", "go")   // go   !
```

### 4. `%5s`

字符串宽度至少 5 位，默认右对齐，左边补空格。

```go
fmt.Printf("%5s!\n", "go")   //    go!
```

## 三、Println、Printf、Sprintf、Fprintf 的区别

抓住“输出到哪里”就行。

### 1. `fmt.Println`

直接打印到终端。

- 参数之间自动加空格
- 结尾自动换行
- 不需要格式符

```go
fmt.Println("tom", 18)   // tom 18
```

### 2. `fmt.Printf`

打印到终端，但需要你自己写格式模板。

- 不会自动换行，通常自己加 `\n`

```go
fmt.Printf("name=%s age=%d\n", "tom", 18)
```

### 3. `fmt.Sprintf`

不打印，只返回格式化后的字符串。

```go
s := fmt.Sprintf("name=%s age=%d", "tom", 18)
// s == "name=tom age=18"
```

适合：

- 拼接字符串
- 生成日志内容
- 生成编号、用户名、文件名

### 4. `fmt.Fprintln`

写到指定目标，不一定是终端。

常见目标：

- 文件
- HTTP 响应
- buffer

```go
fmt.Fprintln(w, "Hello")
```

这里是写到 `w`，不是写到终端。

### 5. `fmt.Fprintf`

和 `Fprintln` 类似，但支持格式化模板。

```go
fmt.Fprintf(w, "name=%s age=%d", "tom", 18)
```

你可以这么记：

- `Println`：直接打印
- `Printf`：按格式打印
- `Sprintf`：按格式生成字符串
- `Fprintf`：按格式写到某个目标

## 四、当前代码里的逻辑

这句：

```go
student := newStudent(i+1, fmt.Sprintf("stu%02d", i+1))
```

可以拆成：

```go
id := i + 1
name := fmt.Sprintf("stu%02d", id)
student := newStudent(id, name)
```

比如当 `i = 0` 时：

```go
id := 1
name := "stu01"
student := newStudent(1, "stu01")
```

所以 `Sprintf` 在这里不是为了打印，而是为了“生成名字字符串”。

## 五、一张速记表

| 格式符 | 含义 | 示例 | 结果 |
|---|---|---|---|
| `%d` | 十进制整数 | `fmt.Sprintf("%d", 18)` | `18` |
| `%02d` | 至少 2 位，不够补 0 | `fmt.Sprintf("%02d", 3)` | `03` |
| `%s` | 字符串 | `fmt.Sprintf("%s", "go")` | `go` |
| `%v` | 通用值 | `fmt.Sprintf("%v", []int{1,2})` | `[1 2]` |
| `%#v` | 更完整的 Go 形式 | `fmt.Sprintf("%#v", s)` | `main.student{...}` |
| `%T` | 类型 | `fmt.Sprintf("%T", 3)` | `int` |
| `%f` | 浮点数 | `fmt.Sprintf("%f", 3.14)` | `3.140000` |
| `%.2f` | 保留 2 位小数 | `fmt.Sprintf("%.2f", 3.14159)` | `3.14` |
| `%t` | 布尔值 | `fmt.Sprintf("%t", true)` | `true` |

## 六、推荐记忆方式

不用一次记太多，先只记这 6 个：

- `%d`
- `%s`
- `%v`
- `%#v`
- `%T`
- `%.2f`

再额外单独记一个规则：

- `%02d` = 整数至少 2 位，不够前面补 0

这个规则对编号场景最常见，比如：

- `stu01`
- `order0001`
- `page03`

## 七、一个判断口诀

当你不确定该用哪个函数时，就问自己一句：

“我是要直接输出，还是要先拿到字符串？”

- 直接输出：`Println` / `Printf`
- 想先得到字符串：`Sprintf`
- 想写到文件、响应、buffer：`Fprintf` / `Fprintln`

## 八、建议自己动手试的 6 行代码

```go
fmt.Println("go", 18)
fmt.Printf("name=%s age=%d\n", "tom", 18)
fmt.Sprintf("stu%02d", 3)
fmt.Printf("%v\n", []int{1, 2, 3})
fmt.Printf("%#v\n", student{ID: 1, Name: "tom"})
fmt.Printf("%T\n", 3.14)
```

把这几行都跑一遍，再对照上面的规则，基本就不会再被 fmt 里的符号绕晕了。