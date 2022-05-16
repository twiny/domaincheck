package domaincheck

import (
	"context"

	"github.com/twiny/whois/v2"
)

type DomainStatus string

// String
func (s DomainStatus) String() string {
	return string(s)
}

const (
	NotApplicable DomainStatus = "N/A"
	Available     DomainStatus = "available"
	Registered    DomainStatus = "registered"
	Premium       DomainStatus = "premium"
)

// Checker
type Checker struct {
	client  *whois.Client
	matcher *matcher
}

// NewChecker
func NewChecker() (*Checker, error) {
	client, err := whois.NewClient(whois.Localhost)
	if err != nil {
		return nil, err
	}

	return &Checker{
		client:  client,
		matcher: newMatcher(),
	}, nil
}

// Check
func (c *Checker) Check(ctx context.Context, domain string) (DomainStatus, error) {
	// get whois server of domain
	server, err := c.client.WHOISHost(domain)
	if err != nil {
		return NotApplicable, err
	}

	// get whois response
	resp, err := c.client.Lookup(ctx, domain, server)
	if err != nil {
		return NotApplicable, err
	}

	return c.matcher.match(resp)
}
