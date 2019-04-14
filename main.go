package main

import (
	"bufio"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read stdin: %s", err)
	}
	input = strings.TrimSuffix(input, "\n")

	quantity := resource.MustParse(input)
	fmt.Println(quantity.Value())
}
