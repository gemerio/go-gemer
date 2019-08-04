package codex

import (
	"bytes"
	"crypto/sha256"
	"hash"
	"os"

	"github.com/gemerio/go-gemer/consensus"
	"github.com/gemerio/go-gemer/crypto"
)

// Defines a WebAssembly module block
type ModuleBlock struct {
	// The given bytecode for a WebAssembly module
	ModuleBytecode WASMBytecode
	// Params that a WASM module accepts
	Parameters WASMParams
	// Return values that a WASM module returns
	ReturnValues WASMReturned
	// Parent WASM module hash
	parentModuleHash [64]byte
}

// An array of module blocks
type BareCodex struct {
	Moduleblocks []ModuleBlock
	barecodexhash [64]byte
}

// Generates a new WebAssembly module block
func NewModuleBlock(bytecode WASMBytecode, params WASMParams, returned WASMReturned) *ModuleBlock {
	moduleblock := &ModuleBlock{
		ModuleBytecode: 	bytecode,
		Parameters: 		params,
		ReturnValues: 		returned,
		parentModuleHash: 	nil
	}
	return moduleblock
}

// Generates a new bare codex
func NewBareCodex(genesisblock ModuleBlock) *BareCodex{
	moduleblocks []ModuleBlock
	moduleblocks[0] := genesisblock
	barecodex := &BareCodex{
		Moduleblocks: 	moduleblock,
		barecodexhash: 	nil
	}
	barecodex.barecodexhash = crypto.generateHash(moduleblocks)
	return barecodex
}

// Adds a WebAssembly module block to the codex
func AddModuleToCodex(moduleblock *ModuleBlock, barecodex *BareCodex) {
	// Check if the given bare codex has a root function (module), else panic
	if(&barecodex.moduleblocks[0] == nil) {
		panic("The bare codex is not correctly initialized.")
	}
	// Add a new module block to the codex
	append(barecodex.moduleblocks, moduleblock)

	// Recalculate the bare codex hash
	barecodex.barecodexhash = crypto.generateHash(moduleblocks)
}

// Reads a preconfigured JSON file to generate a bare codex
func GenerateFromJSON() *[64]byte {

	// TODO: Define a JSON format for a codex.

}