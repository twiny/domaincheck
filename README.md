## domaincheck

A domain name availability checker, using [whois](https://github.com/twiny/whois) client and regexp.

## Install 
`go get github.com/twiny/domaincheck`

## API
```go
NewChecker() (*Checker, error)
Check(ctx context.Context, domain string) (DomainStatus, error)
```

## Example
```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/twiny/domaincheck"
)

func main() {
    checker, err := domaincheck.NewChecker()
    if err != nil {
        log.Fatal(err)
    }

    status, err := checker.Check(context.Background(), "example.com")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(status)
}

```

## Bugs
Bugs or suggestions? Please visit the [issue tracker](https://github.com/twiny/domaincheck/issues).