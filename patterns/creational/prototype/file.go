/*
*	Concrete prototype
 */

package prototype

import "fmt"

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() INode {
	return &File{
		name: f.name + "_clone",
	}
}
