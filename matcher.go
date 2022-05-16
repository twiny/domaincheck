package domaincheck

import (
	"fmt"
	"regexp"
)

// matcher
type matcher struct {
	exceededlimit *regexp.Regexp
	nosuchdomain  *regexp.Regexp
	premiumdomain *regexp.Regexp
	badrequest    *regexp.Regexp
}

// newMatcher
func newMatcher() *matcher {
	return &matcher{
		exceededlimit: regexp.MustCompile(`(?i)(exceeds the limit|limit exceeded)`),
		nosuchdomain:  regexp.MustCompile(`(?i)(NO MATCH|Domain (.*) is available for purchase|NOT FOUND|No entries found|No such domain|No Data Found|nothing found|Status:(.*)AVAILABLE|Status: free|query_status: 220 Available|registration status: available|not been registered|Available(.*\n)Domain:(.*)|The queried object does not exist|domain (.*) available for registration|No Object Found)`),
		premiumdomain: regexp.MustCompile(`(?i)(reserved domain name|reserved by the registry|Reserved by Registry Operator|Reserved Name|Premium domain name|Registry policy prevents registration of domains|usage restrictions|previous registration (.*) was purged on|dpml block|The registration of this domain is restricted|usage restrictions)`),
		badrequest:    regexp.MustCompile(`(not foundConnection|Invalid query)`),
	}
}

func (m *matcher) match(resp string) (DomainStatus, error) {
	if m.exceededlimit.MatchString(resp) {
		return NotApplicable, fmt.Errorf("exceeded limit")
	}

	if m.badrequest.MatchString(resp) {
		return NotApplicable, fmt.Errorf("bad request")
	}

	if m.premiumdomain.MatchString(resp) {
		return Premium, nil
	}

	if m.nosuchdomain.MatchString(resp) {
		return Available, nil
	}

	return Registered, nil
}
