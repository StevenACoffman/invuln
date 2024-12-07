# Invulnerable - Modified Mirror of Vuln

A portmanteau of `internal` and `vuln` to be `invuln`. The actual authors have better taste. :)

[vuln](https://pkg.go.dev/github.com/StevenACoffman/invuln), the database client and tools for the Go vulnerability database, 
are state-of-art code analysis, but many interesting parts are
inside the `internal` package so the implementation details are able to change without breaking
any implied compatibility guarantee.

While I understand this, it makes me sad as I want to play with their toys, but for different purposes.

Specifically, I would like to use their static analysis to visualize and document bespoke Go workflows from request/event entry points to their side effects.

Imagine automatically creating an HTML document for every Pub/Sub or HTTP Request event handler
that describes the event and documents where it goes from there, linking to the HTML document
for any next steps (by analyzing the , and creating a nice little diagram.

See also:
- [GopherCon 2019 - Tracking inter-process dependencies using static analysis](https://mikesep.dev/2019-07-26_gophercon_static_analysis_interprocess_dependencies.pdf)
- [Structured Documentation Extraction for NATS workflows](https://medium.com/swlh/cool-stuff-with-gos-ast-package-pt-2-e4d39ab7e9db) and  [nats_publisher](https://gist.github.com/csthompson/d45cbd973e67efe8ffef5ed2e4c03349/raw/1640d37f26fd38801675329d4ef5d09489ebde03/new_nats_publisher.go)
- [wally](https://github.com/hex0punk/wally) - mapping function paths in code for security analysis
- [laindream/go-callflow-vis](https://github.com/laindream/go-callflow-vis)
- [Replay 2022 Temporal @ Datadog | Jacob LeGrone, Datadog](https://youtu.be/LxgkAoTSI8Q?t=905) the screenshot, but I want to do the reverse of [protoc-gen-go-temporal](https://github.com/cludden/protoc-gen-go-temporal/) 

Anyway, *this* repo was forked and just renamed the internal to external so I could import some
unstable implementation details in my own tool.

I didn't do anything else original here (I mean, the invulnerable pun *was* mine), and only want to use this code as a library
when I'm doing stuff in other repositories.

# Go Vulnerability Management

[![Go Reference](https://pkg.go.dev/badge/github.com/StevenACoffman/invuln.svg)](https://pkg.go.dev/github.com/StevenACoffman/invuln)

Go's support for vulnerability management includes tooling for analyzing your
codebase and binaries to surface known vulnerabilities in your dependencies.
This tooling is backed by the Go vulnerability database, which is curated by
the Go security team. Go’s tooling reduces noise in your results by only
surfacing vulnerabilities in functions that your code is actually calling.

You can install the latest version of govulncheck using
[go install](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies)

```
go install github.com/StevenACoffman/invuln/cmd/govulncheck@latest
```

Then, run govulncheck inside your module:
```
govulncheck ./...
```

See [the govulncheck tutorial](https://go.dev/doc/tutorial/govulncheck) to get
started, and [https://go.dev/security/vuln](https://go.dev/security/vuln) for
more information about Go's support for vulnerability management. The API
documentation can be found at
[https://pkg.go.dev/github.com/StevenACoffman/invuln/scan](https://pkg.go.dev/github.com/StevenACoffman/invuln/scan).

## Privacy Policy

The privacy policy for `govulncheck` can be found at
[https://vuln.go.dev/privacy](https://vuln.go.dev/privacy).

## License

Unless otherwise noted, the Go source files are distributed under the BSD-style
license found in the LICENSE file.

Database entries available at https://vuln.go.dev are distributed under the
terms of the [CC-BY 4.0](https://creativecommons.org/licenses/by/4.0/) license.
