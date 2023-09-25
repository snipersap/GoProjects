package main

import "fmt"

// Builder pattern: Provide the user a simple step by step procedure to build something complex.
// To send an email, the user needs to provide the sender address, receiver address, body of the email,
// subject and maybe signature.
// This complex task, can be split into steps to make the sending of an email easier.

func main() {

	//Builder with facets
	fmt.Println("Builder Facets Pattern >>")
	e := EmailBuilder{&email{}}
	e.From("arun@shimla.com").
		To("sharmaji@mithila.com").
		Subject("Hi from afar").
		Body("Hi Sharma,\nLong time no see. Miss me?\n").
		Signature("--\nYour retired friend,\nArun").
		send()

	// Builder as Parameter, i.e. pass builder struct as param to function
	// Use case: Take the builder pattern as a function to a facade and call the actual implementating
	// function inside it that does the action of sending email
	// This hides the internal struct and actual send email functionality from the user.
	fmt.Println("\n\nBuilder as Parameter Pattern >>")
	SendEmail(func(b *EmailBuilder) {
		b.From("kap@masala.com").
			To("hello@herethere.com").
			Subject("Halli Hallo!").
			Body("Hi Hello tourism,\nWhat are the plans?\n").
			Signature("--\nYours truly,\nKap")

	})

	//Functional Builder: Take all actions as params and
	// execute those actions with a delay (in a single shot)
	efb := EmailFunctionalBuilder{}
	efb.EmailFrom("arun@shimla.com").
		EmailTo("sharmaji@mithila.com").
		EmailSubject("Hi from afar").
		EmailBody("Hi Sharma,\nLong time no see. Miss me?\n").
		EmailSignature("--\nYour retired friend,\nArun").
		BuildAndSendEmail()
}
