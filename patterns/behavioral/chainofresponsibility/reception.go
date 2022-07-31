/**
* Concrete handler
 */

package chainofresponsibility

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registractionDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}

	fmt.Println("Reception registering patient")
	p.registractionDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}
