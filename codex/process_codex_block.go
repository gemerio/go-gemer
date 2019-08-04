package codex

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/rand"
	"fmt"
	"github.com/gemerio/go-gemer/crypto"
)

func FindTargetBlock(barecodex *BareCodex, parentModuleHash [64]byte) *ModuleBlock {
	// Currently using a simple linear array to store the entire tree - so we need a separate hash search function
	// and also non-bidirectional - no convenient searches
	// which is inefficient for large codex trees
	// TODO: Define a linked codex tree data structure & simply iterate over it
	targetblock *ModuleBlock
	for _, module := range &barecodex.Moduleblocks {
		if sha256.Sum([]byte(fmt.Sprintf("%v", module))) == parentModuleHash {
			targetblock := &module
		}
	}
	return targetblock
}

func ProcessCodexBlock(params *[]WASMParams, barecodex *BareCodex) []byte {
	wasmreturned := []byte{}
	prevblock *WASMBytecode
	indexblock *WASMBytecode
	// Iterate over a bare codex
	for i, module := range &barecodex.Moduleblocks {
		if module.parentModuleHash == nil {
			wasmreturned := append(&(module.ModuleBytecode))
			wasmreturned := append(ExecuteWASMBlock(module.ModuleBytecode, *params[i]))
		}
		else {
			// Find the parentmodule block with the given hash from the array
			// If already executed, pass parameter
			// Else, execute parent first
			prevblock := FindTargetBlock(barecodex, module.parentModuleHash)
			if prevblock == nil {
				panic("No module with given hash found")
			}
			// Check returned state array
			for j, block := range wasmreturned {
				if index == prevblock && j % 2 == 0 {
					indexblock := block
				}
			}
			if indexblock == nil {
				wasmreturned := append(&(module.ModuleBytecode))
				wasmreturned := append(ExecuteWASMBlock(module.ModuleBytecode, *params[i]))
			}
			else {
				wasmreturned := append(&(module.ModuleBytecode))
				wasmreturned := append(ExecuteWASMBlock(module.ModuleBytecode, wasmreturned[len(wasmreturned)-2]))
			}
		}
	}
	return wasmreturned
}

func ValidateState(stateblock *StateChainBlock, wasmreturned []byte, params *[]WASMParams, pub *PublicKey, priv *PrivateKey) {
	if &stateblock.&state.Returned.wasmreturned == wasmreturned {
		// Write SHA-256 hash of the entire state
		&stateblock.statehash := sha256.Sum([]byte(fmt.Sprintf("%v", &stateblock.&state)))
		// Sign the resulting statehash with device private key
		r, s, err := ecdsa.Sign(rand.Reader, priv, &stateblock.statehash[:])
		if err != nil {
			panic(err)
		}
		// Verify signature validity
		isvalid := ecdsa.Verify(pub, &stateblock.statehash[:], r, s)
		if !isvalid {
			panic("Signature does not match")
		}
		// Write the signed hash to the state block
		*stateblock.signedstatehash := append(0, r, 0, s)
		// Write the public key
		*stateblock.signedkeys := append(*pub)
	}
	else {
		// TODO: Call consensus & re-execute from given point
	}
}

// Generates a new state chain
func NewStateChain(module *ModuleBlock, returned *WASMReturned) *StateChain{
	// Calculates hash of attached module block
	modulehash [64]byte
	modulehash := sha256.Sum(*module)
	// Generates a blocked state
	blockedstate := &BlockedState {
		Parameters:	module.Parameters,
		Returned:	module.ReturnValues,
		state:		nil // will be processed later
	}
}