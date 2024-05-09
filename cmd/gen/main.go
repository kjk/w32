package main

import (
	"debug/pe"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/kjk/common/u"
	"github.com/microsoft/go-winmd"
	"golang.org/x/exp/maps"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func runLoggedMust(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	panicIfErr(err)
}

func genFromJS() {
	{
		path := filepath.Join("cmd", "gen", "gen.js")
		cmd := exec.Command("bun", path)
		runLoggedMust(cmd)
	}
	{
		cmd := exec.Command("gofmt", "-s", "-w", ".")
		runLoggedMust(cmd)
	}
	{
		cmd := exec.Command("go", "test", ".")
		runLoggedMust(cmd)
	}

}

var (
	panicIf = u.PanicIf
)

func logf(f string, args ...interface{}) {
	s := fmt.Sprintf(f, args...)
	fmt.Print(s)
}

// it's just one  Windows.Win32.winmd
func dumpAssembiles(f *winmd.Metadata) {
	a := f.Tables.Assembly
	for i := 0; i < int(a.Len); i++ {
		r, err := a.Record(winmd.Index(i))
		panicIfErr(err)
		name := r.Name.String()
		logf("assembly: %s\n", name)
	}
}

// it's just one  Windows.Win32.winmd
func dumpModules(f *winmd.Metadata) {
	a := f.Tables.Module
	for i := 0; i < int(a.Len); i++ {
		r, err := a.Record(winmd.Index(i))
		panicIfErr(err)
		name := r.Name.String()
		logf("module: %s\n", name)
	}
}

func dumpNamespacesWithApis(f *winmd.Metadata) {
	seenNamespaces := map[string]bool{}
	for i := uint32(0); i < f.Tables.TypeDef.Len; i++ {
		r, err := f.Tables.TypeDef.Record(winmd.Index(i))
		panicIfErr(err)
		name := r.Name.String()
		if !strings.Contains(name, "Apis") {
			continue
		}
		namespace := r.Namespace.String()
		if !seenNamespaces[namespace] {
			seenNamespaces[namespace] = true
			logf("  \"%s\",\n", namespace)
		}
	}
}

func getMethodDllName(c *Context, methodIndex winmd.Index) string {
	var moduleName string
	if implMap, ok := c.methodDefImplMap[methodIndex]; ok {
		mr, err := c.Metadata.Tables.ModuleRef.Record(implMap.ImportScope)
		panicIfErr(err)
		moduleName = strings.ToLower(mr.Name.String())
		return moduleName
	}
	err := fmt.Errorf("couldn't find method with index %d", methodIndex)
	panic(err)
}

func writePrototypes(f *winmd.Metadata) {
	c, err := NewContext(f)
	panicIfErr(err)
	for i := uint32(0); i < f.Tables.TypeDef.Len; i++ {
		r, err := f.Tables.TypeDef.Record(winmd.Index(i))
		panicIfErr(err)
		name := r.Name.String()
		if !strings.Contains(name, "Apis") {
			continue
		}
		namespace := r.Namespace.String()
		if false && skipNamespace(namespace) {
			continue
		}

		logf("%s:\n", namespace)
		for j := r.MethodList.Start; j < r.MethodList.End; j++ {
			md, err := f.Tables.MethodDef.Record(j)
			panicIfErr(err)
			dllName := getMethodDllName(c, j)
			dllName = strings.ReplaceAll(dllName, ".dll", "")
			var spec string
			if name != "Apis" {
				spec = fmt.Sprintf("  %v.%v.%v", r.Name, dllName, md.Name)
			} else {
				spec = fmt.Sprintf("  %v.%v", dllName, md.Name)
			}
			logf("%s\n", spec)
		}
	}
}

func genFromWinmd() {
	f := loadWinMetadata()
	if false {
		dumpAssembiles(f)
		dumpModules(f)
		dumpNamespacesWithApis(f)
	}
	writePrototypes(f)
}

func loadWinMetadata() *winmd.Metadata {
	path := filepath.Join("..", "Windows.Win32.winmd")
	panicIf(!u.FileExists(path))

	pefile, err := pe.Open(path)
	panicIfErr(err)
	// defer pefile.Close()
	f, err := winmd.New(pefile)
	panicIfErr(err)
	return f
}

func addToMapStringArray(m map[string][]string, key string, val string) {
	a := m[key]
	if !slices.Contains(a, val) {
		a = append(a, val)
		m[key] = a
	}
}

func logbf(b *strings.Builder, f string, args ...interface{}) {
	s := fmt.Sprintf(f, args...)
	b.WriteString(s)
}

func writeStringBuilderToFile(b *strings.Builder, path string) {
	d := []byte(b.String())
	err := os.WriteFile(path, d, 0644)
	panicIfErr(err)
	logf("Wrote '%s' (%d bytes)\n", path, len(d))
}

// writes funcs_by_namespace.txt and funcs_by_dll.txt
func dumpFunctions() {
	logf("dumpFunctions\n")
	var byNamespace strings.Builder
	var byDll strings.Builder
	f := loadWinMetadata()
	c, err := NewContext(f)
	panicIfErr(err)
	funcsByDll := map[string][]string{}
	for i := uint32(0); i < f.Tables.TypeDef.Len; i++ {
		r, err := f.Tables.TypeDef.Record(winmd.Index(i))
		panicIfErr(err)
		name := r.Name.String()
		if !strings.Contains(name, "Apis") {
			continue
		}
		panicIf(name != "Apis", "name: %s", name)
		namespace := r.Namespace.String()
		if false && skipNamespace(namespace) {
			continue
		}

		logbf(&byNamespace, "%s:\n", namespace)
		for j := r.MethodList.Start; j < r.MethodList.End; j++ {
			md, err := f.Tables.MethodDef.Record(j)
			panicIfErr(err)
			dllName := getMethodDllName(c, j)
			dllNameShort := strings.ReplaceAll(dllName, ".dll", "")
			funcName := md.Name.String()
			spec := fmt.Sprintf("  %v.%v", dllNameShort, funcName)
			logbf(&byNamespace, "%s\n", spec)
			addToMapStringArray(funcsByDll, dllName, funcName)
		}
	}
	nTotalFuncs := 0
	dllNames := maps.Keys(funcsByDll)
	slices.Sort(dllNames)
	for _, dllName := range dllNames {
		funcs := funcsByDll[dllName]
		slices.Sort(funcs)
		logbf(&byDll, "%s\n", dllName)
		for _, s := range funcs {
			logbf(&byDll, "  %s\n", s)
		}
		nTotalFuncs += len(funcs)
	}
	writeStringBuilderToFile(&byNamespace, "funcs_by_namespace.txt")
	writeStringBuilderToFile(&byDll, "funcs_by_dll.txt")
	logf("total functions: %d in %d dlls\n", nTotalFuncs, len(dllNames))
}

func main() {
	var (
		flgDumpFunctions bool
		flgGen           bool
	)
	{
		flag.BoolVar(&flgDumpFunctions, "dump-functions", false, "dump information about all functions")
		flag.BoolVar(&flgGen, "gen", false, "generate")
		flag.Parse()
	}

	if false {
		genFromJS()
	}

	if flgDumpFunctions {
		dumpFunctions()
		return
	}
	if flgGen {
		genFromWinmd()
		return
	}
	flag.Usage()
}
