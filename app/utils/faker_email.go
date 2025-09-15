package utils

import (
	"math/rand"
	"regexp"
	"strings"
)

var domains = []string{
	"gmail.com",
	"yahoo.com",
	"outlook.com",
	"example.com",
}

func GenerateEmailFromName(name string) string {
	name = strings.ToLower(name)

	re := regexp.MustCompile(`\s+`)
	name = re.ReplaceAllString(name, "_")

	re = regexp.MustCompile(`[^a-z0-9\.]`)
	name = re.ReplaceAllString(name, "")

	domain := domains[rand.Intn(len(domains))]

	return name + "@" + domain
}
