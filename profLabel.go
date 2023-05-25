package threadContext

import "unsafe"

type labelMap map[string]string

// see @runtime.runtime_getProfLabel
//
//go:linkname runtime_getProfLabel runtime/pprof.runtime_getProfLabel
func runtime_getProfLabel() unsafe.Pointer

// see @runtime.runtime_setProfLabel
//
//go:linkname runtime_setProfLabel runtime/pprof.runtime_setProfLabel
func runtime_setProfLabel(unsafe.Pointer)
