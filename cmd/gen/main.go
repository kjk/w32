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
	var byDll strings.Builder
	nTotalFuncs := genDllFuncsInfo(&byDll, funcsByDll)
	writeStringBuilderToFile(&byNamespace, "funcs_by_namespace.txt")
	writeStringBuilderToFile(&byDll, "funcs_by_dll.txt")
	logf("total functions: %d in %d dlls\n", nTotalFuncs, len(funcsByDll))
}

// TODO: should allow additional info, like a comment
func genDllFuncsInfo(b *strings.Builder, funcsByDll map[string][]string) int {
	nTotalFuncs := 0
	dllNames := maps.Keys(funcsByDll)
	slices.Sort(dllNames)
	for _, dllName := range dllNames {
		funcs := funcsByDll[dllName]
		slices.Sort(funcs)
		logbf(b, "%s\n", dllName)
		for _, s := range funcs {
			logbf(b, "  %s\n", s)
		}
		nTotalFuncs += len(funcs)
	}
	return nTotalFuncs
}

// returns full paths of files starting in root
func getFilesRecur(root string, fnInclude func(name, path string) bool) ([]string, error) {
	var res []string
	dirsToVisit := []string{root}
	for len(dirsToVisit) > 0 {
		dir := dirsToVisit[0]
		dirsToVisit = dirsToVisit[1:]
		files, err := os.ReadDir(dir)
		if err != nil {
			return res, err
		}
		for _, fi := range files {
			name := fi.Name()
			path := filepath.Join(dir, name)
			if fi.IsDir() {
				dirsToVisit = append(dirsToVisit, path)
			} else if fi.Type().IsRegular() {
				if fnInclude == nil || fnInclude(name, path) {
					res = append(res, path)
				}
			}
		}
	}
	return res, nil
}

func getGoFilesRecur(root string) ([]string, error) {
	fnInclude := func(name, path string) bool {
		// TODO: could be faster
		name = strings.ToLower(name)
		return strings.HasSuffix(name, ".go")
	}
	return getFilesRecur(root, fnInclude)
}

// parse: wtsapi32.WTSFreeMemory
func getDllFuncName(s string) (string, string) {
	parts := strings.Split(s, ".")
	panicIf(len(parts) > 2, "unexpected name '%s'", s)
	if len(parts) == 1 {
		return s, "kernel32.dll"
	}
	dllName := strings.TrimSpace(parts[0])
	dllName = strings.ToLower(dllName)
	if !strings.Contains(dllName, ".") {
		dllName += ".dll"
	}
	name := strings.TrimSpace(parts[1])
	return name, dllName
}

type ParsedSys struct {
	Dll         string
	DllFuncName string
	GoFuncName  string
}

/*
Parse this format:
//sys WTSFreeMemory(ptr uintptr) = wtsapi32.WTSFreeMemory
*/
func parseMaybeSysLine(s string) *ParsedSys {
	if !strings.HasPrefix(s, "//sys") {
		return nil
	}
	logf("%s\n", s)
	s2 := s[5:]
	s2 = strings.TrimSpace(s2)
	// logf("  %s\n", s2)
	parts := strings.Split(s2, "(")
	funcName, dllName := getDllFuncName(parts[0])
	goName := funcName
	s3 := strings.ReplaceAll(s2, "==", "")
	s3 = strings.ReplaceAll(s3, "!=", "")
	s3 = strings.ReplaceAll(s3, "<=", "")
	parts = strings.Split(s3, "=")
	panicIf(len(parts) > 2, "unexpecpted line: '%s'", s)
	if len(parts) == 2 {
		funcName, dllName = getDllFuncName(parts[1])
	}
	logf("  %s.%s => %s\n", dllName, funcName, goName)
	return &ParsedSys{
		Dll:         dllName,
		DllFuncName: funcName,
		GoFuncName:  goName,
	}
}

func extractSysDeclarations(path string) []*ParsedSys {
	var res []*ParsedSys
	// logf("file: %s\n", path)
	lines, err := u.ReadLines(path)
	panicIfErr(err)
	for _, s := range lines {
		p := parseMaybeSysLine(s)
		if p != nil {
			res = append(res, p)
		}
	}
	return res
}

// parse checkout of golang.org/x/sys repo to figure out
// wihch functions are already implemented by
// golang.org/x/sys/windows (we don't want to duplicate those)
func genImplementedBySysXWindows() {
	dir := filepath.Join("..", "sys", "windows")
	goFiles, err := getGoFilesRecur(dir)
	if err != nil {
		logf("We expect checkout of golang.org/x/sys repository to our sibling\n")
		logf("To checkout: git clone https://go.googlesource.com/sys")
		panicIfErr(err)
	}

	var a []*ParsedSys
	for _, path := range goFiles {
		a = append(a, extractSysDeclarations(path)...)
	}

	funcsByDll := map[string][]string{}
	for _, p := range a {
		addToMapStringArray(funcsByDll, p.Dll, p.DllFuncName)
	}

	var b strings.Builder
	genDllFuncsInfo(&b, funcsByDll)
	writeStringBuilderToFile(&b, "implemented_in_sys.txt")
}

func main() {
	var (
		flgDumpFunctions bool
		flgGen           bool
		flgParseSys      bool
	)
	{
		flag.BoolVar(&flgDumpFunctions, "dump-functions", false, "dump information about all functions")
		flag.BoolVar(&flgGen, "gen", false, "generate")
		flag.BoolVar(&flgParseSys, "parse-sys", false, "parse golang/x/sys/windows and generate implemented_in_sys.txt")
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

	if flgParseSys {
		genImplementedBySysXWindows()
		return
	}

	flag.Usage()
}
