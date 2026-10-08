# rn_go_asn1c

`go-asn1c` converts ASN.1 specifications of telecom protocols into Go source code. For every ASN.1 type it generates a typed Go struct with `Pack` / `Unpack` methods, plus a stream engine that implements the chosen encoding rules.

## Supported encodings

| Flag       | Encoding                 | Used for                                                               |
|------------|--------------------------|------------------------------------------------------------------------|
| `-t per`   | PER (default)            | NGAP, F1AP, E1AP, E2AP, E2SM, S1AP, X2AP, XnAP, RANAP, HNBAP, RUA, M2AP, M3AP |
| `-t uaper` | Unaligned PER            | RRC (NR and LTE)                                                       |

## Repository layout

| Path            | Contents                                                          |
|-----------------|-------------------------------------------------------------------|
| `asn1/`         | Input ASN.1 specification files (`<proto>-<version>.asn`)         |
| `libgo/`        | Go runtime/stream sources and per-protocol sample programs        |
| `hooks/`        | PyInstaller hooks for building the generator                      |
| `protocols/`    | Generated Go modules, one per protocol and version                |
| `doc/README.md` | Full integration guide (build, generate, encode/decode, examples) |
| `buildpyasn.sh` | Generates Go code for every `.asn` file under `asn1/`             |
| `genasnpy.spec` | PyInstaller spec for building the `genasnpy` generator binary     |

## Quick start

```bash
# Generate code for all specifications in asn1/
./buildpyasn.sh

# Or for a single specification
python2 genasnpy.py -i ngap-f50.asn -t per -p ngap-f50.asn
```

Generated code is written to `protocols/<proto>/<version>/`. Use it from Go by adding a `replace` directive in your `go.mod`:

```go
import ngap "ridenext.co.in/goasn1/ngap"

st := ngap.Stream{}
st.Init(msg)
pdu := ngap.NGAPPDU{}
pdu.Unpack(&st)
```

The full guide (prerequisites, Pack/Unpack walkthroughs, nested OCTET STRING decoding, JSON serialization, troubleshooting) is in [doc/README.md](doc/README.md).

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE).
