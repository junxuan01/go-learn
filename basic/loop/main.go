package main

import "fmt"

func main() {
	// break: 跳出循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("【break】循环次数:", i)
	}
	// continue: 跳过当前循环，继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("【continue】循环次数:", i)
	}
	// range: 遍历数组、切片、字符串等
	s := "郭向文"
	for i, v := range s {
		fmt.Printf("字符: %c, 索引: %d\n", v, i)
	}
	// switch: 多条件分支
	switch i := 3; i {
	case 0:
		fmt.Println("【switch】i是0")
	case 1:
		fmt.Println("【switch】i是1")
	case 2:
		fmt.Println("【switch】i是2")
	default:
		fmt.Println("【switch】i大于2")
	}

	age := 30
	switch {
	case age < 18:
		fmt.Println("【switch】未成年人")
	case age >= 18 && age < 60:
		fmt.Printf("【switch】成年人")
	default:
		fmt.Printf("【switch】老年人")
	}
	// goto: 跳转到指定标签
	fmt.Println("\n=== goto示例 ===")

	// 示例1: 基本goto使用
	i := 0
start:
	if i < 3 {
		fmt.Printf("【goto】循环 %d\n", i)
		i++
		goto start
	}

	// 示例2: 错误处理中的goto
	err := processData()
	if err != nil {
		goto cleanup
	}

	fmt.Println("【goto】数据处理成功")
	goto end

cleanup:
	fmt.Println("【goto】清理资源")

end:
	fmt.Println("【goto】程序结束")
}

func processData() error {
	// 模拟可能出错的操作
	return nil // 返回nil表示成功
}
