package main

import "fmt"

func countConnections(groupSize int) int {
	//Count the number of connections between people in the group
	connections := 0
	for i := 0; i <= groupSize; i++ {
		if i <= 1 {
			continue
		}
		connections += i - 1
	}
	return connections
}

func main() {
	for i := range 20 {
		fmt.Printf("Connections for %v group size is:%v\n", i, countConnections(i))
	}

}
