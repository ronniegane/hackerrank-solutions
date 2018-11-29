package queues

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main(r io.Reader) {
	// Reads from STDIN
	// First line is number of actions
	var actions int
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	actions, _ = strconv.Atoi(scanner.Text())

	var stack1, stack2 []string
	for i := 0; i < actions; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		if len(line) > 1 {
			// enqueue
			// push onto stack 1
			stack1 = append(stack1, line[1])
		} else {
			// dequeue operations work on stack 2
			// if stack2 is empty, need to populate
			if len(stack2) == 0 {
				// Move stack1 onto stack2 (reverse)
				for i := len(stack1) - 1; i >= 0; i-- {
					stack2 = append(stack2, stack1[i])
				}
				stack1 = make([]string, 0)
			}
			if line[0] == "2" {
				// dequeue / pop
				stack2 = stack2[:len(stack2)-1]
			} else if line[0] == "3" {
				// Print top of queue (peek)
				fmt.Println(stack2[len(stack2)-1])
			}
		}
	}
}
