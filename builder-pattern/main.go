package main

import "fmt"

// Builder pattern: Provide the user a simple step by step procedure to build something complex.
// To send an email, the user needs to provide the sender address, receiver address, body of the email,
// subject and maybe signature.
// This complex task, can be split into steps to make the sending of an email easier.
type email struct {
	from, to, subject, body, signature string
}

type emailBuilder struct {
	email *email
}

func main() {

	e := emailBuilder{&email{}}
	e.from("arun@shimla.com").
		to("sharmaji@mithila.com").
		subject("Hi from afar").
		body("Hi Sharma,\nLong time no see. Miss me?\n").
		signature("--\nYour retired friend,\nArun").
		send()
}

func (e *emailBuilder) from(address string) *emailBuilder {
	e.email.from = address
	return e
}

func (e *emailBuilder) to(address string) *emailBuilder {
	e.email.to = address
	return e
}

func (e *emailBuilder) subject(sub string) *emailBuilder {
	e.email.subject = sub
	return e
}

func (e *emailBuilder) body(body string) *emailBuilder {
	e.email.body = body
	return e
}

func (e *emailBuilder) signature(sign string) *emailBuilder {
	e.email.signature = sign
	return e
}

func (e *emailBuilder) send() {
	fmt.Printf("From:%s\nTo:%s\nSubject:%s\n\n%s\n%s",
		e.email.from, e.email.to, e.email.subject, e.email.body, e.email.signature)
}
