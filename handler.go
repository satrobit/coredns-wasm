package wasm

import (
	"encoding/json"

	"github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"
	extism "github.com/extism/go-sdk"
	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

func (wasm *Wasm) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {

	state := request.Request{W: w, Req: r}

	qname := state.Name()
	qtype := state.Type()

	query := &Query{
		Type: qtype,
		Name: qname,
	}
	query_json, _ := json.Marshal(query)

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasm.wasmPath,
			},
		},
	}

	config := extism.PluginConfig{
		EnableWasi: true,
	}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	if err != nil {
		log.Errorf("Failed to initialize plugin: %v\n", err)
		return dns.RcodeServerFailure, nil
	}

	_, response, err := plugin.Call("ServeDNS", query_json)
	if err != nil {
		log.Error(err)
		return dns.RcodeServerFailure, nil
	}

	record := new(Record)
	err = json.Unmarshal(response, record)
	if err != nil {
		log.Errorf("parse error : %v %v\n", response, err)
		return dns.RcodeServerFailure, nil
	}

	answers := make([]dns.RR, 0, 10)

	switch qtype {
	case "A":
		answers = wasm.A(qname, record)
	case "AAAA":
		answers = wasm.AAAA(qname, record)
	case "CNAME":
		answers = wasm.CNAME(qname, record)
	case "TXT":
		answers = wasm.TXT(qname, record)
	default:
		return wasm.errorResponse(state, dns.RcodeNotImplemented, nil)
	}

	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative, m.RecursionAvailable, m.Compress = true, false, true

	m.Answer = append(m.Answer, answers...)

	state.SizeAndDo(m)
	m = state.Scrub(m)
	_ = w.WriteMsg(m)
	return dns.RcodeSuccess, nil

}

func (wasm *Wasm) Name() string { return "wasm" }
