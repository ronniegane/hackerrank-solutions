package queues

import (
	"bufio"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	f, _ := os.Open("test_input1")
	reader := bufio.NewReader(f)
	main(reader)
}
