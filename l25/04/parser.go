package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// result stores the parsed result for a domain
type result struct {
	domain string
	visits int
	// add more metrics if needed
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   int               // total visits for all domains
	lines   int               // number of parsed lines (for the error messages)
}

func newParser() parser {
	return parser{sum: make(map[string]result)}
}

func parseField(p parser, line string) (parser, error) {
	p.lines++

	// Parse the fields
	fields := strings.Fields(line)
	if len(fields) != 2 {
		err := fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return p, err
	}

	domain := fields[0]

	// Sum the total visits per domain
	visits, err := strconv.Atoi(fields[1])
	if visits < 0 || err != nil {
		err := fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
		return p, err
	}

	// Collect the unique domains
	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}

	// Keep track of total and per domain visits
	p.total += visits

	// You cannot assign to composite values
	// p.sum[domain].visits += visits

	// create and assign a new copy of `visit`
	p.sum[domain] = result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}

	return p, nil

}

func printVisits(p parser) {
	sort.Strings(p.domains)
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

}

func printTotalVisits(p parser) {
	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}
