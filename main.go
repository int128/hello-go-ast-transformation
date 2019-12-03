package main

import (
	"errors"
	"fmt"
	"go/ast"
	"log"

	"golang.org/x/tools/go/packages"
)

func dump(pkgNames ...string) error {
	cfg := &packages.Config{
		Mode: packages.NeedCompiledGoFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg, pkgNames...)
	if err != nil {
		return fmt.Errorf("could not load the packages: %w", err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		return errors.New("error while loading the packages")
	}
	for _, pkg := range pkgs {
		if err := ast.Print(pkg.Fset, pkg); err != nil {
			return fmt.Errorf("could not dump the package: %w", err)
		}
	}
	return nil
}

func main() {
	if err := dump("."); err != nil {
		log.Fatalf("error: %s", err)
	}
}
