# Rust 语言学习系列文章计划

## 概述

参照 Go 语言学习系列文章的结构，为前端开发者编写 Rust 语言学习教程。每篇文章不少于 3000 字，包含与 JavaScript、Go 的对比，使用 Rust 1.91 版本。

## 文章列表（共 17 篇）

### 第一部分：入门介绍（1-4）

1. **Rust 语言简介和历史背景.md**

   - Mozilla 与 Rust 的诞生（2010-2015）
   - Rust 的设计目标：安全、并发、实用
   - Rust 基金会与社区发展
   - 与 JavaScript 的定位对比（系统级 vs 应用级）

2. **Rust 语言的主要特点和优势.md**

   - 内存安全（无 GC，所有权系统）
   - 零成本抽象
   - 并发安全（无数据竞争）
   - 现代工具链（Cargo、rustfmt、clippy）
   - 与 JS 对比：类型系统、错误处理

3. **Rust 语言的应用领域.md**

   - 系统编程（操作系统、驱动）
   - WebAssembly（前端性能优化）
   - 命令行工具
   - 网络服务（Actix、Axum）
   - 嵌入式开发
   - 区块链（Solana、Polkadot）

4. **学习 Rust 语言的资源和建议.md**
   - 官方资源：The Rust Book、Rustlings
   - 在线练习：Exercism、Rust by Example
   - 社区资源：Reddit、Discord
   - 学习路径建议
   - 常见学习难点（所有权、生命周期）

### 第二部分：环境与工具（5-6）

5. **Rust 语言安装与环境配置.md**

   - rustup 安装（macOS/Windows/Linux）
   - 国内镜像配置
   - VS Code + rust-analyzer 配置
   - 第一个 Rust 程序
   - cargo 基础命令

6. **Rust 语言依赖管理以及 Cargo 使用.md**
   - Cargo.toml 详解
   - 依赖管理（crates.io vs npm）
   - 工作空间（Workspace）
   - 常用命令：build、run、test、doc
   - 发布 crate

### 第三部分：基础语法（7-11）

7. **Rust 语言基础之变量和常量.md**

   - let 与 mut（不可变 vs 可变）
   - const 与 static
   - 变量遮蔽（Shadowing）
   - 类型推断
   - 与 JS 的 let/const 对比

8. **Rust 语言基础之基本数据类型.md**

   - 标量类型：整数、浮点、布尔、字符
   - 复合类型：元组、数组
   - String vs &str
   - Option 和 Result
   - 类型转换

9. **Rust 语言基础之运算符.md**

   - 算术、比较、逻辑运算符
   - 位运算符
   - 引用与解引用（& 和 \*）
   - 运算符重载（std::ops）

10. **Rust 语言基础之流程控制.md**

    - if/else 表达式（返回值）
    - loop、while、for
    - match 模式匹配
    - if let 和 while let
    - 与 JS switch 的对比

11. **Rust 语言基础之数组与元组.md**
    - 数组：固定长度 [T; N]
    - 元组：异构固定长度
    - 解构赋值
    - 与 JS 数组的区别

### 第四部分：核心概念（12-14）

12. **Rust 语言基础之切片与向量.md**

    - 切片 &[T]：借用视图
    - Vec<T>：动态数组
    - 常用方法：push、pop、iter
    - 与 JS 数组的对比

13. **Rust 语言基础之 HashMap.md**

    - HashMap<K, V> 基础
    - 插入、访问、更新
    - Entry API
    - 与 JS Object/Map 的对比

14. **Rust 语言基础之函数.md**
    - 函数定义与参数
    - 返回值（表达式 vs 语句）
    - 闭包（Closure）
    - 高阶函数
    - 与 JS 箭头函数的对比

### 第五部分：类型系统（15-17）

15. **Rust 语言基础之结构体.md**

    - struct 定义
    - 方法与关联函数（impl）
    - 元组结构体
    - 与 JS class 的对比

16. **Rust 语言基础之枚举与模式匹配.md**

    - enum 定义（带数据的枚举）
    - Option<T> 和 Result<T, E>
    - match 详解
    - 与 TS 联合类型的对比

17. **Rust 语言基础之特征与泛型.md**
    - trait 定义与实现
    - 泛型函数与结构体
    - trait bounds
    - 与 TS interface 的对比

## 调整说明

相比 prompt 中的原始目录，做了以下调整：

1. **接口 → 特征与泛型**：Rust 使用 trait 而非 interface
2. **map → HashMap**：Rust 的 HashMap 在标准库中
3. **数组 → 数组与元组**：合并讲解，因为 Rust 数组固定长度
4. **新增切片与向量**：Vec 是 Rust 最常用的动态集合
5. **新增枚举与模式匹配**：这是 Rust 的核心特性
6. **包 → 合并到 Cargo 使用**：Rust 的模块系统在 Cargo 文章中讲解

## 写作风格

- 每篇开头用 JS 代码引入概念
- 详细解释 Rust 特有概念（所有权、借用、生命周期）
- 提供可运行的代码示例
- 文末总结与下一篇预告
- 使用 Rust 1.91 最新语法

## 执行顺序

按编号 1-17 顺序生成文章到 `/Users/guoxiangwen/code/github/golearn/rust-article/` 目录。
