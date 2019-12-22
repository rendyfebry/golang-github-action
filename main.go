package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rendyfebry/golang-github-action/internal/util/strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Yourname: ")

	text, _ := reader.ReadString('\n')
	sanitized := strings.SanitizeToAlphaNumeric(text)

	fmt.Println(fmt.Sprintf("Hello, %s", sanitized))
}
