package main

import (
	"fmt"
	"log"
)

// For Builder facets and as parameter
type email struct {
	from, to, subject, body, signature string
}

type EmailBuilder struct { //caps to denote usage outside of package
	email *email
}

// Fluent interfaces
func (e *EmailBuilder) From(address string) *EmailBuilder {
	e.email.from = address
	return e
}

func (e *EmailBuilder) To(address string) *EmailBuilder {
	e.email.to = address
	return e
}

func (e *EmailBuilder) Subject(sub string) *EmailBuilder {
	e.email.subject = sub
	return e
}

func (e *EmailBuilder) Body(body string) *EmailBuilder {
	e.email.body = body
	return e
}

func (e *EmailBuilder) Signature(sign string) *EmailBuilder {
	e.email.signature = sign
	return e
}

func (e *EmailBuilder) send() {
	fmt.Printf("From:%s\nTo:%s\nSubject:%s\n\n%s\n%s\n",
		e.email.from, e.email.to, e.email.subject, e.email.body, e.email.signature)

	log.Println("Email sent to", e.email.to)
}

// Public function for user to send an email: Builder with parameter
func SendEmail(action func(*EmailBuilder)) {
	b := EmailBuilder{&email{}} //init
	action(&b)                  //call the anonymous func passed as param
	b.send()                    //call func to do some work
}

// Functional Builder pattern
// Objective: to capture the actions and execute later or with a delay
type action func(*email)
type EmailFunctionalBuilder struct {
	actions []action
}

// capture the actions as anonymous functions and store in EmailFunctionalBuilder
func (e *EmailFunctionalBuilder) EmailTo(to string) *EmailFunctionalBuilder {
	e.actions = append(e.actions, func(em *email) {
		em.to = to
	})
	return e
}

func (e *EmailFunctionalBuilder) EmailFrom(from string) *EmailFunctionalBuilder {
	e.actions = append(e.actions, func(em *email) {
		em.from = from
	})
	return e
}

func (e *EmailFunctionalBuilder) EmailSubject(sub string) *EmailFunctionalBuilder {
	e.actions = append(e.actions, func(em *email) {
		em.subject = sub
	})
	return e
}

func (e *EmailFunctionalBuilder) EmailBody(body string) *EmailFunctionalBuilder {
	e.actions = append(e.actions, func(em *email) {
		em.body = body
	})
	return e
}

func (e *EmailFunctionalBuilder) EmailSignature(sign string) *EmailFunctionalBuilder {
	e.actions = append(e.actions, func(em *email) {
		em.signature = sign
	})
	return e
}

// Execute the captured actions and send email
func (e *EmailFunctionalBuilder) BuildAndSendEmail() {
	em := email{}
	for _, action := range e.actions {
		action(&em)
	}
	fmt.Println("\n\nFunctional Builder Pattern >>")
	fmt.Printf("From:%s\nTo:%s\nSubject:%s\n\n%s\n%s\n",
		em.from, em.to, em.subject, em.body, em.signature)

	log.Println("Email sent to", em.to)
}
