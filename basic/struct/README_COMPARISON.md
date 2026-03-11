# Go Struct 组合 vs JavaScript Class 继承

## 快速对比

| 特性 | JavaScript Class | Go Struct |
|------|------------------|-----------|
| 核心理念 | 继承（Inheritance） | 组合（Composition） |
| 关系 | is-a（是一个） | has-a（有一个） |
| 类型关系 | Student 是 Person 的子类型 | Student 不是 Person 的子类型 |
| 代码复用 | extends 继承 | embedding 嵌入 |
| 多态实现 | 继承 + 方法重写 | 接口 |
| 多继承 | 不支持（单继承链） | 支持嵌入多个类型 |

## 运行示例

### JavaScript 版本
```bash
node students_example.js
```

### Go 版本
```bash
go run students_example.go
```

## 关键区别详解

### 1. 类型关系

**JavaScript:**
```javascript
const student = new Student("张三", 18, "S001", 2);
console.log(student instanceof Person);  // true - Student 是 Person
```

**Go:**
```go
student := Student{Person: Person{Name: "张三", Age: 18}}
// var p Person = student  // 编译错误！Student 不是 Person
// 正确方式：访问嵌入的字段
var p Person = student.Person  // OK
```

### 2. 方法调用

**JavaScript:**
```javascript
class Student extends Person {
  introduce() {
    return super.introduce() + "额外信息";  // super 调用父类
  }
}
```

**Go:**
```go
type Student struct {
    Person  // 嵌入
}

func (s *Student) Introduce() string {
    return s.Person.Introduce() + "额外信息"  // 显式调用
}
```

### 3. 多重能力

**JavaScript - 只能单继承：**
```javascript
// 不能这样做：
// class Student extends Person, Employee { }  // 语法错误

// 只能用 mixin 等模式曲线救国
```

**Go - 可以嵌入多个类型：**
```go
type Student struct {
    Person      // 嵌入 Person
    IDCard      // 嵌入 IDCard
    Address     // 嵌入 Address
}
// 所有嵌入类型的字段和方法都会"提升"到 Student
```

### 4. 多态实现

**JavaScript - 通过继承：**
```javascript
function greet(person) {  // 接受 Person 或其子类
    console.log(person.introduce());
}

greet(new Person("张三", 20));      // OK
greet(new Student("李四", 18, ...)); // OK - Student 是 Person
```

**Go - 通过接口：**
```go
type Introducer interface {
    Introduce() string
}

func Greet(i Introducer) {  // 接受任何实现了 Introduce 的类型
    fmt.Println(i.Introduce())
}

Greet(&Person{Name: "张三"})    // OK - Person 实现了 Introducer
Greet(&Student{...})            // OK - Student 也实现了 Introducer
// 它们不需要有类型关系，只需要有相同的方法
```

## 哪种方式更好？

### JavaScript Class 适合的场景：
- 需要明确的类型层级（管理类、继承结构）
- 框架约定（React 类组件等）
- 需要 instanceof 类型检查
- 业务领域模型有清晰的"是一个"关系

### Go Struct 组合适合的场景：
- 优先组合而非继承（Go 哲学）
- 需要灵活组装能力
- 避免深层继承带来的复杂性
- 关注"能做什么"而非"是什么"
- 可测试性（依赖接口而非具体类型）

## 实践建议

### JavaScript:
```javascript
// ✅ 好的实践
class Animal { }
class Dog extends Animal { }  // 清晰的 is-a 关系

// ❌ 避免深层继承
class A { }
class B extends A { }
class C extends B { }
class D extends C { }  // 太深了！
```

### Go:
```go
// ✅ 好的实践 - 小接口 + 组合
type Reader interface { Read([]byte) (int, error) }
type Writer interface { Write([]byte) (int, error) }

type File struct {
    name string
    // 实现 Read 和 Write
}

// ✅ 组合多个能力
type ReadWriter struct {
    Reader
    Writer
}

// ❌ 避免巨大接口
type GodInterface interface {
    Method1()
    Method2()
    // ... 50 个方法
}
```

## 核心思想

- **JavaScript Class**: "Student 是一个 Person"（继承身份）
- **Go Struct**: "Student 包含一个 Person"（组合能力）

Go 的设计哲学是：
> "组合优于继承，接口优于类型层级"
