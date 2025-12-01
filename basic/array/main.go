package main

import "fmt"

func main() {
	var langArray = [...]int{1: 10, 3: 30}
	fmt.Println("数组langArray内容:", langArray)

	var arrDemo = [5]int{1, 3, 5, 7, 8}
	var sum = 0
	// for i := 0; i < len(arrDemo); i++ {
	// 	sum = sum + arrDemo[i]
	// }
	for _, k := range arrDemo {
		sum = sum + k
	}
	// 找出数组arrDemo中和为8的两个元素的下标

	for i := 0; i < len(arrDemo); i++ {
		for j := i + 1; j < len(arrDemo); j++ {
			if arrDemo[i]+arrDemo[j] == 8 {
				fmt.Printf("找到和为8的元素: %d + %d = 8, 下标: (%d, %d)\n", arrDemo[i], arrDemo[j], i, j)
			}
		}
	}

	fmt.Printf("数组元素之和: %d\n", sum) // 输出数组元素之和

	fmt.Println("=== Go数组详细演示 ===")

	// 1. 数组的各种定义方式
	fmt.Println("\n【1. 数组定义方式】")

	// 方式1: 声明后赋值
	var arr1 [5]int
	arr1[0] = 10
	arr1[1] = 20
	fmt.Printf("方式1 - 声明后赋值: %v\n", arr1)

	// 方式2: 声明时初始化
	var arr2 [3]string = [3]string{"Go", "Python", "Java"}
	fmt.Printf("方式2 - 声明时初始化: %v\n", arr2)

	// 方式3: 简短声明
	arr3 := [4]int{1, 2, 3, 4}
	fmt.Printf("方式3 - 简短声明: %v\n", arr3)

	// 方式4: 让编译器推断长度
	arr4 := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("方式4 - 自动推断长度: %v, 长度: %d\n", arr4, len(arr4))

	// 方式5: 指定索引初始化
	arr5 := [5]int{0: 100, 2: 200, 4: 400}
	fmt.Printf("方式5 - 指定索引: %v\n", arr5)

	// 2. 数组的基本操作
	fmt.Println("\n【2. 数组基本操作】")

	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("原数组: %v\n", numbers)
	fmt.Printf("数组长度: %d\n", len(numbers))
	fmt.Printf("数组类型: %T\n", numbers)

	// 访问和修改元素
	fmt.Printf("第一个元素: %d\n", numbers[0])
	fmt.Printf("最后一个元素: %d\n", numbers[len(numbers)-1])

	numbers[2] = 100
	fmt.Printf("修改后: %v\n", numbers)

	// 3. 数组遍历
	fmt.Println("\n【3. 数组遍历方法】")

	fruits := [4]string{"苹果", "香蕉", "橘子", "葡萄"}

	// 方法1: 普通for循环
	fmt.Println("方法1 - 普通for循环:")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("  索引%d: %s\n", i, fruits[i])
	}

	// 方法2: range循环 (推荐)
	fmt.Println("方法2 - range循环:")
	for index, value := range fruits {
		fmt.Printf("  索引%d: %s\n", index, value)
	}

	// 方法3: 只要值，不要索引
	fmt.Println("方法3 - 只要值:")
	for _, value := range fruits {
		fmt.Printf("  %s ", value)
	}
	fmt.Println()

	// 方法4: 只要索引，不要值
	fmt.Println("方法4 - 只要索引:")
	for index := range fruits {
		fmt.Printf("  索引%d ", index)
	}
	fmt.Println()

	// 4. 多维数组
	fmt.Println("\n【4. 多维数组】")

	// 二维数组
	var matrix [3][4]int
	matrix[0] = [4]int{1, 2, 3, 4}
	matrix[1] = [4]int{5, 6, 7, 8}
	matrix[2] = [4]int{9, 10, 11, 12}

	fmt.Printf("二维数组类型: %T\n", matrix)
	fmt.Println("二维数组内容:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}

	// 简化的二维数组初始化
	board := [3][3]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}
	fmt.Println("\n井字棋盘:")
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}

	// 5. 数组比较
	fmt.Println("\n【5. 数组比较】")

	arr_a := [3]int{1, 2, 3}
	arr_b := [3]int{1, 2, 3}
	arr_c := [3]int{1, 2, 4}

	fmt.Printf("arr_a: %v\n", arr_a)
	fmt.Printf("arr_b: %v\n", arr_b)
	fmt.Printf("arr_c: %v\n", arr_c)

	fmt.Printf("arr_a == arr_b: %t\n", arr_a == arr_b) // true
	fmt.Printf("arr_a == arr_c: %t\n", arr_a == arr_c) // false

	// 6. 数组作为函数参数
	fmt.Println("\n【6. 数组作为函数参数】")

	testArray := [3]int{10, 20, 30}
	fmt.Printf("调用前: %v\n", testArray)

	modifyArray(testArray) // 值传递，不会修改原数组
	fmt.Printf("调用后: %v\n", testArray)

	modifyArrayByPointer(&testArray) // 指针传递，会修改原数组
	fmt.Printf("指针修改后: %v\n", testArray)

	// 7. 数组的零值
	fmt.Println("\n【7. 数组的零值】")

	var zeroInts [3]int
	var zeroBools [3]bool
	var zeroStrings [3]string
	var zeroFloats [3]float64

	fmt.Printf("int数组零值: %v\n", zeroInts)       // [0 0 0]
	fmt.Printf("bool数组零值: %v\n", zeroBools)     // [false false false]
	fmt.Printf("string数组零值: %v\n", zeroStrings) // ["" "" ""]
	fmt.Printf("float64数组零值: %v\n", zeroFloats) // [0 0 0]

	// 8. 数组 vs 切片对比
	fmt.Println("\n【8. 数组 vs 切片对比】")

	// 数组 - 固定长度
	fixedArray := [3]string{"A", "B", "C"}
	fmt.Printf("数组: %v, 类型: %T, 长度: %d\n", fixedArray, fixedArray, len(fixedArray))

	// 切片 - 动态长度
	dynamicSlice := []string{"A", "B", "C"}
	fmt.Printf("切片: %v, 类型: %T, 长度: %d, 容量: %d\n",
		dynamicSlice, dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	// 切片可以扩展，数组不行
	dynamicSlice = append(dynamicSlice, "D")
	fmt.Printf("扩展后切片: %v\n", dynamicSlice)

	// 9. 重要注意事项演示
	fmt.Println("\n【9. 重要注意事项】")

	// 注意点1: 数组长度是类型的一部分
	var arr_3 [3]int
	var arr_5 [5]int
	fmt.Printf("[3]int 类型: %T\n", arr_3)
	fmt.Printf("[5]int 类型: %T\n", arr_5)
	// arr_3 = arr_5  // 编译错误！不同类型不能赋值

	// 注意点2: 数组是值类型，赋值会复制整个数组
	original := [3]int{1, 2, 3}
	copied := original // 完整复制
	copied[0] = 100
	fmt.Printf("原数组: %v\n", original) // [1 2 3] - 不变
	fmt.Printf("复制数组: %v\n", copied)  // [100 2 3] - 改变

	// 注意点3: 数组边界检查
	fmt.Println("\n边界检查演示:")
	safeArray := [3]int{1, 2, 3}
	fmt.Printf("安全访问 safeArray[0]: %d\n", safeArray[0])
	fmt.Printf("安全访问 safeArray[2]: %d\n", safeArray[2])
	// fmt.Printf("越界访问 safeArray[3]: %d\n", safeArray[3])  // 运行时panic!

	// 10. 实际应用示例
	fmt.Println("\n【10. 实际应用示例】")

	// 示例1: 统计分数
	scores := [5]int{85, 92, 78, 96, 88}
	total := 0
	for _, score := range scores {
		total += score
	}
	average := float64(total) / float64(len(scores))
	fmt.Printf("分数: %v\n", scores)
	fmt.Printf("总分: %d, 平均分: %.2f\n", total, average)

	// 示例2: 查找最大值和最小值
	findMaxMin(scores)

	// 示例3: 简单的矩阵运算
	matrixDemo()
}

// 函数演示：值传递
func modifyArray(arr [3]int) {
	arr[0] = 999 // 只修改副本
	fmt.Printf("函数内修改: %v\n", arr)
}

// 函数演示：指针传递
func modifyArrayByPointer(arr *[3]int) {
	arr[0] = 999 // 修改原数组
	fmt.Printf("指针函数内修改: %v\n", *arr)
}

// 查找最大值和最小值
func findMaxMin(arr [5]int) {
	if len(arr) == 0 {
		return
	}

	max, min := arr[0], arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	fmt.Printf("最大值: %d, 最小值: %d\n", max, min)
}

// 矩阵运算演示
func matrixDemo() {
	fmt.Println("\n矩阵运算演示:")

	// 创建一个3x3矩阵
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("原矩阵:")
	printMatrix(matrix)

	// 计算对角线和
	diagonalSum := 0
	for i := 0; i < 3; i++ {
		diagonalSum += matrix[i][i]
	}
	fmt.Printf("主对角线和: %d\n", diagonalSum)
}

// 打印矩阵
func printMatrix(matrix [3][3]int) {
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%3d ", cell)
		}
		fmt.Println()
	}
}
