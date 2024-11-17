package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func Input() (numbers []int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		num, parseErr := strconv.Atoi(line)
		if parseErr != nil || num < -100000 || num > 100000 {
			err = fmt.Errorf("invalid input: %v", line)
			return
		}
		numbers = append(numbers, num)
	}

	if scanner.Err() != nil {
		err = scanner.Err()
	}
	return
}