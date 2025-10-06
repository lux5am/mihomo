package common

import (
	"strings"

	C "github.com/metacubex/mihomo/constant"
)

type DomainKeyword struct {
	Base
	keyword string
	adapter string
}

func (dk *DomainKeyword) RuleType() C.RuleType {
	return C.DomainKeyword
}

func (dk *DomainKeyword) Match(metadata *C.Metadata, helper C.RuleMatchHelper) (bool, string) {
	return (metadata.SniffHost != "" && strings.Contains(metadata.SniffHost, dk.keyword)) ||
			(metadata.Host != "" && metadata.Host != metadata.SniffHost && strings.Contains(metadata.Host, dk.keyword)),
		dk.adapter
}

func (dk *DomainKeyword) Adapter() string {
	return dk.adapter
}

func (dk *DomainKeyword) Payload() string {
	return dk.keyword
}

func NewDomainKeyword(keyword string, adapter string) *DomainKeyword {
	return &DomainKeyword{
		Base:    Base{},
		keyword: strings.ToLower(keyword),
		adapter: adapter,
	}
}

var _ C.Rule = (*DomainKeyword)(nil)
