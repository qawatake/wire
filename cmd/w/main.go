package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/google/wire/internal/wire"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("done")
}

func run() error {
	wd, err := os.Getwd()
	if err != nil {
		log.Println("failed to get working directory: ", err)
		return err
	}
	flag.Parse()

	var f *flag.FlagSet
	f = flag.CommandLine
	info, errs := wire.Load(context.Background(), wd, os.Environ(), "", packages(f))
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	fmt.Println("👹👹👹")
	// for id, p := range info.Sets {
	// 	fmt.Printf("[%v]\n", id)
	// 	for _, root := range p.Providers {
	// 		tree := ProviderTree{set: p}
	// 		tree.Make(root)
	// 		lines := tree.lines
	// 		// fmt.Printf("😈%s😈\n", lines)
	// 		renderer := mermaidRenderer{}
	// 		s, err := renderer.Render(lines)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(s)
	// 	}
	// }

	for id, set := range info.InjectorSets {
		for _, t := range set.Outputs() {
			fmt.Println("😙", t)
		}
		i := set.InjectorArgs
		length := set.InjectorArgs.Tuple.Len()
		for j := 0; j < length; j++ {
			fmt.Println("😙", i.Tuple.At(j))
		}
		fmt.Printf("[%v]\n", id)
		for _, root := range set.Providers {
			tree := ProviderTree{set: set}
			tree.Make(root)
			lines := tree.lines
			renderer := mermaidRenderer{}
			s, err := renderer.Render(lines)
			if err != nil {
				return err
			}
			fmt.Println(s)
		}
	}

	return nil
}

func packages(f *flag.FlagSet) []string {
	pkgs := f.Args()
	if len(pkgs) == 0 {
		pkgs = []string{"."}
	}
	return pkgs
}

func printProviderSet(p *wire.ProviderSet, level int) {
	for _, x := range p.Providers {
		fmt.Println(strings.Repeat(" ", level), x.Name)
		for _, z := range x.Out {
			fmt.Println(strings.Repeat(" ", level+1), "+", z)
		}
		for _, y := range x.Args {
			fmt.Println(strings.Repeat(" ", level+1), "-", y.Type)
		}
	}
	for _, x := range p.Imports {
		printProviderSet(x, level+1)
	}
}

func providers(set *wire.ProviderSet, kv map[string]*wire.Provider) {
	for _, p := range set.Providers {
		key := providerID(p)
		kv[key] = p
	}
	for _, subset := range set.Imports {
		providers(subset, kv)
	}
}

func providerID(p *wire.Provider) string {
	return p.Pkg.Path() + "." + p.Name
}

func printMap(m map[string]*wire.Provider) {
	for k, v := range m {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func providerTree(set *wire.ProviderSet) {
	kv := make(map[string]*wire.Provider)
	providers(set, kv)
	root := set.Providers[0]
	arg := root.Args[0]
	x := set.For(arg.Type)
	x.Provider()
}

type ProviderTree struct {
	lines []string
	set   *wire.ProviderSet
}

func (t *ProviderTree) Make(root *wire.Provider) error {
	for _, arg := range root.Args {
		if !t.set.For(arg.Type).IsProvider() {
			continue
		}
		argProvider := t.set.For(arg.Type).Provider()
		t.lines = append(t.lines, fmt.Sprintf("%s --> %s;", providerID(argProvider), providerID(root)))
		if err := t.Make(argProvider); err != nil {
			return err
		}
	}
	return nil
}

type mermaidRenderer struct {
}

func (r *mermaidRenderer) Render(lines []string) (string, error) {
	var s string

	theme := "%%{init:{'theme':'default','flowchart':{'rankSpacing':500}}}%%\n"
	header := "flowchart TD;\n"

	s += theme
	s += header

	lines = uniq(lines)

	sort.Slice(lines, func(i, j int) bool {
		return strings.Compare(lines[i], lines[j]) < 0
	})

	for _, ss := range lines {
		s += fmt.Sprintf("\t%s\n", ss)
	}

	return s, nil
}

// 順番は維持されません。
func uniq(lines []string) []string {
	m := make(map[string]struct{})
	for _, s := range lines {
		m[s] = struct{}{}
	}
	var ss []string
	for k := range m {
		ss = append(ss, k)
	}
	return ss
}
