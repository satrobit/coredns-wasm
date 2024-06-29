# coredns-wasm
By using the *coredns-wasm* module, you can execute WebAssembly modules as a data source for CoreDNS.

## Functionalities

- [x] Embedded runtime
- [x] Return records: A, AAAA, CNAME, TXT
- [ ] Return errors
- [ ] Read and write metadata


## Syntax

~~~
wasm {
    wasmPath [PATH]
}
~~~

* `wasmPath` is the path to your WebAssembly module
## Example

~~~ corefile
. {
    wasm {
        wasmPath "./a_record.wasm"
    }
}
~~~

> Checkout the example in the `./wasm` dir to understand how to write the actual WASM module.

