/*
*	Client code
 */

package prototype

import "fmt"

func PrototypeInitializer() {
	file1 := &File{name: "File 1"}
	file2 := &File{name: "File 2"}
	file3 := &File{name: "File 3"}

	folder1 := &Folder{
		children: []INode{file1},
		name: "Folder1",
	}

	folder2 := &Folder{
		children: []INode{folder1, file2, file3},
		name: "Folder2",
	}
	fmt.Println("Printing hierachy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierachy for clone Folder")
	cloneFolder.print("  ")
}

