package main

import (
	"fmt"
	"os"
)

func main() {
	// 查询正在生效的环境变量
	s1 := os.Environ()
	fmt.Println(s1)

	// 设置环境变量的值
	os.Setenv("K8S_HOME", "")
	// 查找一个变量的值
	k8s, isExist := os.LookupEnv("K8S_HOME")
	fmt.Printf("K8S_HOME: %t   -- %s  \n", isExist, k8s)
	// 获取一个变量的值
	k8s_GET := os.Getenv("K8S_HOME")
	fmt.Printf("K8S_HOME: %s \n", k8s_GET)

	// 取消环境变量
	os.Unsetenv("K8S_HOME")

	// 清空
	os.Clearenv()

	// 主要方法对比 Getenv LookupEnv

	// Getenv 变量不存在 或者 变量为空   都返回 空
	// LookupEnv 验证是否存在环境变量
}
