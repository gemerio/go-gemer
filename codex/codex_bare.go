package codex

import (
	"crypto/sha256"
	"encoding/json"
)

func generateHash(moduleblocks []ModuleBlock) [32]byte {
	arrayBytes := []byte{}
	for _, codex := range moduleblocks {
		jsonReadBytes, _ := json.Marshal(codex)
		arrayBytes = append(arrayBytes, jsonReadBytes...)
	}
	return sha256.Sum256(arrayBytes)
}

// Defines a WebAssembly module block
type ModuleBlock struct {
	// The given bytecode for a WebAssembly module
	ModuleBytecode WASMBytecode
	// Params that a WASM module accepts
	Parameters WASMParams
	// Return values that a WASM module returns
	ReturnValues WASMReturned
	// Parent WASM module hash
	parentModuleHash [32]byte
}

// An array of module blocks
type BareCodex struct {
	Moduleblocks  []ModuleBlock
	barecodexhash [32]byte
}

// Generates a new WebAssembly module block
func NewModuleBlock(bytecode WASMBytecode, params WASMParams, returned WASMReturned) *ModuleBlock {
	moduleblock := &ModuleBlock{
		ModuleBytecode:   bytecode,
		Parameters:       params,
		ReturnValues:     returned,
		parentModuleHash: [32]byte{},
	}
	return moduleblock
}

// Generates a new bare codex
func NewBareCodex(genesisblock ModuleBlock) *BareCodex {
	moduleblocks := []ModuleBlock{genesisblock}
	barecodex := &BareCodex{
		Moduleblocks:  moduleblocks,
		barecodexhash: [32]byte{},
	}
	barecodex.barecodexhash = generateHash(moduleblocks)
	return barecodex
}

// Adds a WebAssembly module block to the codex
func AddModuleToCodex(moduleblock *ModuleBlock, barecodex BareCodex) {
	// Check if the given bare codex has a root function (module), else panic
	if len(barecodex.Moduleblocks) == 0 {
		panic("The bare codex is not correctly initialized.")
	}
	// Add a new module block to the codex
	barecodex.Moduleblocks = append(barecodex.Moduleblocks, *moduleblock)

	// Recalculate the bare codex hash
	barecodex.barecodexhash = generateHash(barecodex.Moduleblocks)
}

// Reads a preconfigured JSON file to generate a bare codex
func GenerateFromJSON() *[64]byte {

	// TODO: Define a JSON format for a codex.
	return nil
}
