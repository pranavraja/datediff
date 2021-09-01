package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseDate(input string) (date, error) {
	var d date
	if _, err := fmt.Sscanf(input, "%d/%d/%d", &d.day, &d.month, &d.year); err != nil {
		return d, err
	}
	if d.day < 1 || d.day > 31 {
		return d, fmt.Errorf("invalid day %d", d.day)
	}
	if d.month < 1 || d.month > 12 {
		return d, fmt.Errorf("invalid month %d", d.month)
	}
	if d.year < 1900 || d.year > 2999 {
		return d, fmt.Errorf("invalid year %d", d.year)
	}
	return d, nil
}

func main() {
	fmt.Println(`hi. type something like "2/6/1983 to 22/6/1983", then press ENTER. Press Ctrl+D once you're done`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		var from, to string
		if _, err := fmt.Sscanf(input, "%s to %s", &from, &to); err != nil {
			log.Fatalf("bad input line %s: expected something like: 2/6/1983 to 22/6/1983", input)
		}
		dateFrom, err := parseDate(from)
		if err != nil {
			log.Fatalf("couldn't parse %s as date: %s", from, err)
		}
		dateTo, err := parseDate(to)
		if err != nil {
			log.Fatalf("couldn't parse %s as date: %s", to, err)
		}
		fmt.Printf("%d days\n", dateTo.Diff(dateFrom))
	}
}
