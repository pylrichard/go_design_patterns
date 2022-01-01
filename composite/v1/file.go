package main

import "fmt"

/*
	实现组件接口
 */
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for %s in %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}