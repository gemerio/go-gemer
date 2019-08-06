package codex

import (
	"os"
	"fmt"
	"time"
	"io/ioutil"

	"github.com/gemerio/go-gemer/consensus"

	"github.com/gemerio/life/exec"
)

// Defines a WASM bytecode format
type WASMBytecode struct {
	// The WASM bytecode module itself, read from a *.wasm file
	wasmbytecode []byte
	// Any additional arguments?
	execargs string
}

type WASMParams struct {
	// A set of params given to a WASM module.
	// Defined as []byte as params are not given in a particular format
	wasmparams []byte
	// What is the entry function of this module?
	entryfunction string
	// Any additional arguments?
	arguments string
}

type WASMReturned struct {
	// A set of params returned by a WASM module.
	// Defined as []byte as params are not given in a particular format
	wasmreturned []byte
	// Any additional arguments returned?
	arguments string
}

// Reads WASM bytecode from a specified file
func NewWASMBytecode(filename string) *WASMBytecode{
	bytecodeblock := &WASMBytecode {
		wasmbytecode:	ioutil.ReadFile(filename),
		execargs:		nil
	}
	return bytecodeblock
}

// Executes each WASM block
// currently using Perlin Network's life WebAssembly VM, but will require a separate implementation later on
func ExecuteWASMBlock(bytecodeblock WASMBytecode, params WASMParams) *WASMReturned {
	// Initialize Life VM
	vm, err := exec.NewVirtualMachine(bytecodeblock.wasmbytecode, exec.VMConfig{}, &exec.NopResolver{}, nil)
	if err != nil {
		fmt.Printf("Invalid bytecode block. The detailed error message is: \n")
		panic(err)
	}
	// Load entry point function
	entrypointID, ifFound := vm.GetFunctionExport(params.entryfunction)
	if !ifFound {
		panic("The specified entry function does not exist.")
	}
	// Run the bytecode block
	returned, err := vm.Run(entrypointID)
	if err != nil {
		vm.PrintStackTrace()
		panic(err)
	}
	// Returns the reference of a WASMReturned struct
	wasmreturned := &WASMReturned {
		wasmreturned:	returned,
		arguments:		nil
	}
	return wasmreturned
}

func DefineImportResolver() {
	// TODO: Fetch external function callpoints from a given module
	// And record it to a WASMReturned struct
}