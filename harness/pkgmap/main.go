// Command pkgmap prints a map of one Go package directory as JSON: every .go
// file with its line count and its top-level declarations.
//
// It only parses the source with go/parser. It never builds or runs it, so it
// is safe on third-party code.
package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type decl struct {
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Recv     string `json:"recv,omitempty"`
	Exported bool   `json:"exported"`
	Line     int    `json:"line"`
}

type file struct {
	Name  string `json:"name"`
	Lines int    `json:"lines"`
	Test  bool   `json:"test"`
	Decls []decl `json:"decls"`
	Error string `json:"error,omitempty"`
}

func recvName(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.StarExpr:
		return recvName(t.X)
	case *ast.IndexExpr:
		return recvName(t.X)
	case *ast.IndexListExpr:
		return recvName(t.X)
	case *ast.Ident:
		return t.Name
	}
	return "?"
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: pkgmap <package dir>")
		os.Exit(2)
	}
	dir := os.Args[1]
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var files []file
	fset := token.NewFileSet()
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		src, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		f := file{Name: e.Name(), Lines: strings.Count(string(src), "\n"), Test: strings.HasSuffix(e.Name(), "_test.go")}
		af, err := parser.ParseFile(fset, path, src, parser.SkipObjectResolution)
		if err != nil {
			f.Error = err.Error()
		}
		if af != nil {
			for _, d := range af.Decls {
				switch d := d.(type) {
				case *ast.FuncDecl:
					x := decl{Kind: "func", Name: d.Name.Name, Exported: d.Name.IsExported(), Line: fset.Position(d.Pos()).Line}
					if d.Recv != nil && len(d.Recv.List) > 0 {
						x.Kind, x.Recv = "method", recvName(d.Recv.List[0].Type)
					}
					f.Decls = append(f.Decls, x)
				case *ast.GenDecl:
					for _, sp := range d.Specs {
						switch sp := sp.(type) {
						case *ast.TypeSpec:
							f.Decls = append(f.Decls, decl{Kind: "type", Name: sp.Name.Name, Exported: sp.Name.IsExported(), Line: fset.Position(sp.Pos()).Line})
						case *ast.ValueSpec:
							kind := "var"
							if d.Tok == token.CONST {
								kind = "const"
							}
							for _, n := range sp.Names {
								if n.Name == "_" {
									continue
								}
								f.Decls = append(f.Decls, decl{Kind: kind, Name: n.Name, Exported: n.IsExported(), Line: fset.Position(n.Pos()).Line})
							}
						}
					}
				}
			}
		}
		files = append(files, f)
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Name < files[j].Name })
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(files); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
