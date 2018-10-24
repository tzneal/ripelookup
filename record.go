package ripelookup

import (
	"bytes"
	"fmt"
	"sort"
)

// Record is a single WHOIS record for an IP Address
type Record map[string]string

func (r Record) String() string {
	buf := bytes.Buffer{}
	keys := []string{}
	for k := range r {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(&buf, "%s: %s\n", k, r[k])
	}
	return buf.String()
}
