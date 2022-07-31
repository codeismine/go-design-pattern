/**
* Concrete handler
 */

package chainofresponsibility

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}

	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
