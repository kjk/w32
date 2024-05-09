//go:build windows

package w32

import (
	"sync"
	"syscall"
	"unsafe"
)

// func UTF16FromString(s string) ([]uint16, error)
// func UTF16PtrFromString(s string) (*uint16, error)
// func SyscallN(trap uintptr, args ...uintptr) (r1, r2 uintptr, err Errno)

var (
	UTF16PtrFromString = syscall.UTF16PtrFromString
	UTF16FromString    = syscall.UTF16FromString
	SysN               = syscall.SyscallN
)

const (
	DNS_TYPE_A = syscall.DNS_TYPE_A
)

type Handle = syscall.Handle
type Errno = syscall.Errno
type SecurityAttributes = syscall.SecurityAttributes
type GUID = syscall.GUID

func errnoErr(e Errno) error {
	return e
}

type dllInfo struct {
	dllName string
	dll     *syscall.DLL
	procs   []uintptr
	names   []string
}

var (
	mu       sync.Mutex
	dllInfos = []*dllInfo{}
)

var kernel32Funcs = []string{
	"FooW",
}

func makeDllInfo(dllName string, funcNames []string) *dllInfo {
	nFuncs := len(funcNames)
	return &dllInfo{
		dllName: dllName,
		procs:   make([]uintptr, nFuncs),
		names:   funcNames,
	}
}

func init() {
	{
		i := makeDllInfo("kernel32.dll", kernel32Funcs)
		dllInfos = append(dllInfos, i)
	}
}

const procNotAvailable = uintptr(1)

func loadProc(dllIdx int, procIdx int) uintptr {
	mu.Lock()
	defer mu.Unlock()

	i := dllInfos[dllIdx]
	addr := i.procs[procIdx]
	if addr == procNotAvailable {
		// we've already tried to load but was not present
		return 0
	}
	if i.dll != nil {
		i.dll = syscall.MustLoadDLL(i.dllName)
	}
	proc, err := i.dll.FindProc(i.names[procIdx])
	if err != nil || proc == nil {
		i.procs[procIdx] = procNotAvailable
		return procNotAvailable
	}
	i.procs[procIdx] = proc.Addr()
	return 0
}

const (
	dllIdxKernel32 = 0

	procIdxCreateProc = 0
)

func CreateFileMapping(fhandle Handle, sa *SecurityAttributes, prot uint32, maxSizeHigh uint32, maxSizeLow uint32, name *uint16) (handle Handle, err error) {
	addr := loadProc(dllIdxKernel32, procIdxCreateProc)
	r0, _, e1 := SysN(addr, 6, uintptr(fhandle), uintptr(unsafe.Pointer(sa)), uintptr(prot), uintptr(maxSizeHigh), uintptr(maxSizeLow), uintptr(unsafe.Pointer(name)))
	handle = Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}
