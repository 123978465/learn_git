package main

import "fmt"

func main() {
	s1 := make([]int, 5, 5)
	print(&s1)
	s1 = delete(s1, 2)
}

func  delete(s []int, index int) []int {
	fmt.Println("删除了切片中第", index, "个元素:", s[index])
	s1 := s[0:index]
	s2 := s[index+1:]
	//s2 := s[index+1:]
	for _, vul := range s2 {
		s1 = append(s1, vul)
	}
	print(&s1)
	return s1
}

// 方法传指针时len cap爆红 需使用*slice
func print(slice *[]int) {
	fmt.Println("切片信息")
	fmt.Printf("%v ,长度%d , 容量%d \n", slice, len(*slice), cap(*slice))
}
