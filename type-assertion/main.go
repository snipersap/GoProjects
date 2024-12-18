package main

func getExpenseReport(e expense) (string, float64) {

	//email? return the receiver address and cost
	email, ok := e.(email)
	if ok {
		return email.toAddress, email.cost()
	}
	//sms? return the phone number and cost
	sms, ok := e.(sms)
	if ok {
		return sms.toPhoneNumber, sms.cost()
	}
	//others? return empty
	return "", 0.0

}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
