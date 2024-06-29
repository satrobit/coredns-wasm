package wasm

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() {
	caddy.RegisterPlugin("wasm", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	r, err := wasmParse(c)
	if err != nil {
		return plugin.Error("wasm", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		r.Next = next
		return r
	})

	return nil
}

func wasmParse(c *caddy.Controller) (*Wasm, error) {
	wasm := Wasm{}

	for c.Next() {
		if c.NextBlock() {
			for {
				switch c.Val() {
				case "wasmPath":
					if !c.NextArg() {
						return &Wasm{}, c.ArgErr()
					}
					wasm.wasmPath = c.Val()
				default:
					if c.Val() != "}" {
						return &Wasm{}, c.Errf("unknown property '%s'", c.Val())
					}
				}

				if !c.Next() {
					break
				}
			}

		}

		return &wasm, nil
	}
	return &Wasm{}, nil
}
