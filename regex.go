package main

import "regexp"

var (
	uuidMatch *regexp.Regexp = regexp.MustCompile(`(?m)[^\/]+$`)
)
