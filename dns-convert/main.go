package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	convert "./convert"
	unique "./unique"
)

func main() {
	var ipList []string
	var uniqueIPs []string

	file, err := os.Open("./entries.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var entries []string
	for scanner.Scan() {
		entries = append(entries, scanner.Text())
	}
	sort.Strings(entries)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ipList = convert.ConvertNames(entries)
	uniqueIPs = unique.UniqueIPs(ipList)
	uniqueIPs = append([]string{"entry,ip\n"}, uniqueIPs...)

	for i, v := range uniqueIPs {
		if i > 0 {
			fmt.Print("")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
