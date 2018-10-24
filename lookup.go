package ripelookup

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/kentik/patricia"
)

// ErrNotFound is returned when a WHOIS server cannot be found for a given IP
var ErrNotFound = errors.New("server not found")

// DetermineServer determines the appropriate whois server to query for a particular IP
func DetermineServer(ip net.IP) (string, error) {
	v4, v6, err := patricia.ParseIPFromString(ip.String())
	if err != nil {
		return "", fmt.Errorf("error parsing IP: %s", err)
	}
	if v4 != nil {
		found, tag, err := v4tree.FindDeepestTag(*v4)
		if err != nil {
			return "", fmt.Errorf("error looking up v4 IP: %s", err)
		}
		if found {
			return tag, nil
		}
		return "", ErrNotFound
	}
	if v6 != nil {
		found, tag, err := v6tree.FindDeepestTag(*v6)
		if err != nil {
			return "", fmt.Errorf("error looking up v6 IP: %s", err)
		}
		if found {
			return tag, nil
		}
		return "", ErrNotFound
	}
	return "", ErrNotFound
}

// WhoisIP performs a WHOIS against a server for an IP Address, returning the matching records
func WhoisIP(query string, server string) ([]Record, error) {
	records := []Record{}
	conn, err := net.Dial("tcp", server+":43")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	fmt.Fprintf(conn, "n + %s\r\n", query)
	sc := bufio.NewScanner(conn)
	rec := Record{}
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "# start") {
			// start record
			rec = Record{}
			continue
		} else if strings.HasPrefix(line, "# end") {
			// end record
			records = append(records, rec)
			rec = nil
		} else if strings.HasPrefix(line, "#") {
			continue
		}
		sp := strings.SplitN(line, ":", 2)
		if len(sp) != 2 {
			continue
		}
		sp[0] = strings.TrimSpace(sp[0])
		sp[1] = strings.TrimSpace(sp[1])
		rec[sp[0]] = sp[1]
	}
	if len(rec) > 0 {
		records = append(records, rec)
	}
	return records, nil
}
