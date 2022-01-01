package main

import "fmt"

/*
	通过组合整合组件
 */
type Folder struct {
	components	[]Component
	name		string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching recursively for %s in %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}