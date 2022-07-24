/*
*	Prototype interface
*/

package prototype

type INode interface {
	print(string)
	clone() INode
}
