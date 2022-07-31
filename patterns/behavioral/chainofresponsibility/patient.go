package chainofresponsibility

type Patient struct {
	name              string
	registractionDone bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}
