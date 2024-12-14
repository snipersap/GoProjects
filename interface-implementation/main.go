package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func main() {
	donaldDuck := fullTime{
		name:   "Donald Duck",
		salary: 100000,
	}
	printEmployeeDetails(donaldDuck)

	uncleScrooge := contractor{
		name:         "Mr. Scrooge",
		hourlyPay:    200,
		hoursPerYear: 220*8 - 30 - 10, //possible working hours - vacations - sick days
	}
	printEmployeeDetails(uncleScrooge)
}

func printEmployeeDetails(e employee) {
	fmt.Println("Employee", e.getName(), "was paid $", e.getSalary(), "this year!")
}
