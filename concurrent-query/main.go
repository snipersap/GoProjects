package main

import "fmt"

type Conn string
type Result string

func main() {
	// 1. Create a list of connections
	// 2. Call query with the list of connections
	// 3. Print the result
	// 4. Prints ...on connection 9 and Main: ... on connection 9
	c := []Conn{}
	for i := 0; i < 10; i++ {
		c = append(c, Conn("Connection"+fmt.Sprint(i)))
	}
	r := Query(c, "query")
	fmt.Println("Main:", r)

}

func Query(conns []Conn, query string) Result {
	ch := make(chan Result, len(conns)) // buffered
	for _, conn := range conns {
		go func(c Conn) {
			ch <- c.DoQuery(query)
		}(conn)
	}
	return <-ch
}

func (c Conn) DoQuery(q string) Result {
	r := Result("I have performed: " + q + " on " + string(c))
	fmt.Println(r)
	return r
}
