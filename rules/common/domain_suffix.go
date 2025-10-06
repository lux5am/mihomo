package common

import (
	"strings"

	C "github.com/metacubex/mihomo/constant"
)

type DomainSuffix struct {
	Base
	suffix  string
	adapter string
}

func (ds *DomainSuffix) RuleType() C.RuleType {
	return C.DomainSuffix
}

func (ds *DomainSuffix) Match(metadata *C.Metadata, helper C.RuleMatchHelper) (bool, string) {
	return (metadata.SniffHost != "" && (strings.HasSuffix(metadata.SniffHost, "."+ds.suffix) || metadata.SniffHost == ds.suffix)) ||
			(metadata.Host != "" && metadata.Host != metadata.SniffHost && (strings.HasSuffix(metadata.Host, "."+ds.suffix) || metadata.Host == ds.suffix)),
		ds.adapter
}

func (ds *DomainSuffix) Adapter() string {
	return ds.adapter
}

func (ds *DomainSuffix) Payload() string {
	return ds.suffix
}

func NewDomainSuffix(suffix string, adapter string) *DomainSuffix {
	return &DomainSuffix{
		Base:    Base{},
		suffix:  strings.ToLower(suffix),
		adapter: adapter,
	}
}

var _ C.Rule = (*DomainSuffix)(nil)
