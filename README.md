![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff)
[![CodeQL](https://github.com/github/docs/actions/workflows/codeql.yml/badge.svg)](https://github.com/Neved4/twc-go/actions/workflows/codeql.yml)
[![Homebrew](https://img.shields.io/badge/Homebrew-tap-2AAB47?logo=homebrew&logoColor=959DA5&labelColor=2B3137)](https://github.com/Neved4/homebrew-tap/blob/main/Formula/twc-go.rb)

# `twc.go` - Tiny world clock in Go! 🦦

Fast, CLI world clock that displays time zone information using
[tz database] to read valid [tz entries].

```console
$ twc -h -f tools/samples/tz-small.conf
America/New_York     2023-12-11 12:37:13
America/Los_Angeles  2023-12-11 09:37:13
UTC                  2023-12-11 17:37:13
Europe/London        2023-12-11 17:37:13
Europe/Paris         2023-12-11 18:37:13
Asia/Tokyo           2023-12-12 02:37:13
Australia/Sydney     2023-12-12 04:37:13
```

### Highlights

- 🚀 _**Fast**_ - **10X** times faster than [`twc.c`] and **65X** times
  faster than `date`.[^1] \
  Only ≈ `7 ms` for one entry, and ≈ `177 ms` for ≈ 600
  entries.
- 🔒 _**Robust**_ - tested to work with all [tz database] entries,
  [`version 2023c`].
- 📦 **Self-contained** - zero dependencies,
  lighweight (`2457 bytes`, `108 lines`).

## Getting Started

### Setup

If you have [Homebrew] installed, just run:
```console
$ brew tap Neved4/homebrew-tap/twc-go
```

Alternatively, clone the repository:

```console
$ git clone https://github.com/Neved4/twc-go
```

Then build:

```console
$ go build twc.go
```

### Usage

```
Usage of ./twc:
  -f string
    	Specify file path (default "tz-small.conf")
  -h	Print human-readable format
  -s string
    	Specify time format
  -t string
    	Specify timezone directly

Examples:
    $ twc -h -s %Y-%m-%d -t Asia/Tokyo
        2006-01-02

    $ TZ=America/Los_Angeles twz
        2006-01-02T15:04:05-0800

Environment:
    TZ  Timezone to use when displaying dates.

See also:
    environ(7)
```

## Compatibility

Runs on _**Linux**_, _**macOS**_ and _**\*BSD**_ systems on both
[`x86_64`] and [`arm64`].

## Standards

`twc.go` should run on most [POSIX.1-2017][][^2] compatible systems and
produce [ISO 8601][][^3] output.

## License
                 
`twc.go` is licensed under the terms of the [MIT License].

See the [LICENSE](LICENSE) file for details.


## Related Projects

- [oz/tz] - awesome cli time zone helper
- [Neved4/twc][`twc.c`] - tiny world clock in C

[`hyperfine`]: https://github.com/sharkdp/hyperfine
[Homebrew]: https://brew.sh/
[`twc.c`]: https://github.com/Neved4/twc
[`arm64`]: https://en.wikipedia.org/wiki/AArch64
[`x86_64`]: https://en.wikipedia.org/wiki/X86-64
[MIT License]: https://opensource.org/license/mit/
[POSIX.1-2017]: https://pubs.opengroup.org/onlinepubs/9699919799/
[ISO 8601]: https://www.iso.org/obp/ui/#iso:std:iso:8601:-2:ed-1:v1:en
[tz database]: https://en.wikipedia.org/wiki/Tz_database
[tz entries]: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
[`version 2023c`]: https://www.iana.org/time-zones
[oz/tz]: https://github.com/oz/tz

[^1]: _cfr._ `date` command takes ≈ `931 ms` when iterating over ≈ 600
    entries. Measured with [`hyperfine`].
[^2]: _IEEE Std 1003.1-2017: Standard for Information Technology
    — Portable Operating System Interface (POSIX®)_, \
    ISO/IEC/IEEE 9945:2009/COR 2:2017. URL: https://pubs.opengroup.org/onlinepubs/9699919799/
[^3]: _ISO 8601: Date and time \ — Representations for information interchange_, ISO 8601-1:2019. \
    URL: https://www.iso.org/obp/ui/#iso:std:iso:8601:-2:ed-1:v1:en
