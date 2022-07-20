package outputter

import (
	"github.com/gancc6/ksubdomain/runner/result"
)

type Output interface {
	WriteDomainResult(domain result.Result) error
	Close()
}
