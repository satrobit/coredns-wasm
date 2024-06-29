package wasm

import "net"

type Query struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// copied from https://github.com/codysnider/coredns-redis/blob/master/types.go
type Record struct {
	A     []A_Record     `json:"a,omitempty"`
	AAAA  []AAAA_Record  `json:"aaaa,omitempty"`
	TXT   []TXT_Record   `json:"txt,omitempty"`
	CNAME []CNAME_Record `json:"cname,omitempty"`
}

type A_Record struct {
	Ttl uint32 `json:"ttl,omitempty"`
	Ip  net.IP `json:"ip"`
}

type AAAA_Record struct {
	Ttl uint32 `json:"ttl,omitempty"`
	Ip  net.IP `json:"ip"`
}

type TXT_Record struct {
	Ttl  uint32 `json:"ttl,omitempty"`
	Text string `json:"text"`
}

type CNAME_Record struct {
	Ttl  uint32 `json:"ttl,omitempty"`
	Host string `json:"host"`
}
