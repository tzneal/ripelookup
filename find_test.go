package ripelookup_test

import (
	"net"
	"testing"

	"github.com/tzneal/ripelookup"
)

func TestDetermineServer(t *testing.T) {
	testCases := []struct {
		IP     string
		Server string
	}{
		{"8.8.8.8", "whois.arin.net"},
		{"2607:f8b0:4002:c09::8b", "whois.arin.net"},
	}
	for _, tc := range testCases {
		server, err := ripelookup.DetermineServer(net.ParseIP(tc.IP))
		if err != nil {
			t.Fatalf("error looking up IP: %s", err)
		}
		if tc.Server != server {
			// Not sure how often these change, I suspect not often.  If they do
			// though, this test will fail when it's not really broken
			t.Errorf("expected %s, got %s", tc.Server, server)
		}
	}
}
