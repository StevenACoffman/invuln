// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vulncheck

import (
	"fmt"
	"runtime"
	"sort"

	"golang.org/x/tools/go/packages"
	"golang.org/x/vuln/internal/client"
	"golang.org/x/vuln/internal/osv"
	"golang.org/x/vuln/internal/semver"
)

// newTestClient returns a client that reads
// a database with the following vulnerable symbols:
//
//	golang.org/amod/avuln.{VulnData.Vuln1, vulnData.Vuln2}
//	golang.org/bmod/bvuln.Vuln
//	archive/zip.OpenReader
func newTestClient() (*client.Client, error) {
	return client.NewInMemoryClient(
		[]*osv.Entry{
			{
				ID: "VA",
				Affected: []osv.Affected{{
					Module: osv.Module{Path: "golang.org/amod"},
					Ranges: []osv.Range{{Type: osv.RangeTypeSemver, Events: []osv.RangeEvent{{Introduced: "1.0.0"}, {Fixed: "1.0.4"}, {Introduced: "1.1.2"}}}},
					EcosystemSpecific: osv.EcosystemSpecific{Packages: []osv.Package{{
						Path:    "golang.org/amod/avuln",
						Symbols: []string{"VulnData.Vuln1", "VulnData.Vuln2"}},
					}},
				}},
			},
			{
				ID: "VB",
				Affected: []osv.Affected{{
					Module: osv.Module{Path: "golang.org/bmod"},
					Ranges: []osv.Range{{Type: osv.RangeTypeSemver}},
					EcosystemSpecific: osv.EcosystemSpecific{
						Packages: []osv.Package{{
							Path:    "golang.org/bmod/bvuln",
							Symbols: []string{"Vuln"},
						}},
					},
				}},
			},
			{
				ID: "STD",
				Affected: []osv.Affected{{
					Module: osv.Module{Path: osv.GoStdModulePath},
					// Range is populated also using runtime info for testing binaries since
					// setting fixed Go version for binaries is very difficult.
					Ranges: []osv.Range{{Type: osv.RangeTypeSemver, Events: []osv.RangeEvent{{Introduced: "1.18"}, {Introduced: semver.GoTagToSemver(runtime.Version())}}}},
					EcosystemSpecific: osv.EcosystemSpecific{
						Packages: []osv.Package{{
							Path:    "archive/zip",
							Symbols: []string{"OpenReader"},
						}},
					},
				}},
			}})
}

func vulnsToString(vulns []*osv.Entry) string {
	var s string
	for _, v := range vulns {
		s += fmt.Sprintf("\t%v\n", v)
	}
	return s
}

func impGraphToStrMap(ig *ImportGraph) map[string][]string {
	m := make(map[string][]string)
	for _, n := range ig.Packages {
		for _, predId := range n.ImportedBy {
			pred := ig.Packages[predId]
			m[pred.pkg.PkgPath] = append(m[pred.pkg.PkgPath], n.pkg.PkgPath)
		}
	}

	sortStrMap(m)
	return m
}

func reqGraphToStrMap(r *Result) map[string][]string {
	m := make(map[string][]string)
	for _, n := range r.ModulesByPath {
		for _, pred := range n.RequiredBy {
			m[pred.Path] = append(m[pred.Path], n.Path)
		}
	}

	sortStrMap(m)
	return m
}

func callGraphToStrMap(cg *CallGraph) map[string][]string {
	type edge struct {
		// src and dest are ids ofr source and
		// destination nodes in a callgraph edge.
		src, dst int
	}
	// seen edges, to avoid repetitions
	seen := make(map[edge]bool)

	m := make(map[string][]string)
	for _, n := range cg.Functions {
		fName := n.String()
		for _, callsite := range n.CallSites {
			e := edge{src: callsite.Parent, dst: n.ID}
			if seen[e] {
				continue
			}
			caller := cg.Functions[e.src]
			callerName := caller.String()
			m[callerName] = append(m[callerName], fName)
		}
	}

	sortStrMap(m)
	return m
}

func pkgPathToImports(pkgs []*packages.Package) map[string][]string {
	m := make(map[string][]string)
	seen := make(map[*packages.Package]bool)
	var visit func(*packages.Package)
	visit = func(p *packages.Package) {
		if seen[p] {
			return
		}
		seen[p] = true
		var imports []string
		for _, i := range p.Imports {
			imports = append(imports, i.PkgPath)
			visit(i)
		}
		m[p.PkgPath] = imports
	}
	for _, p := range pkgs {
		visit(p)
	}
	sortStrMap(m)
	return m
}

func modulePathToVersion(pkgs []*packages.Package) map[string]string {
	m := make(map[string]string)
	seen := make(map[*packages.Package]bool)
	var visit func(*packages.Package)
	visit = func(p *packages.Package) {
		if seen[p] || p.Module == nil {
			return
		}
		seen[p] = true
		for _, i := range p.Imports {
			visit(i)
		}
		m[p.Module.Path] = p.Module.Version
	}
	for _, p := range pkgs {
		visit(p)
	}
	return m
}

// sortStrMap sorts the map string slice values to make them deterministic.
func sortStrMap(m map[string][]string) {
	for _, strs := range m {
		sort.Strings(strs)
	}
}
