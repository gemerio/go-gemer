package codex

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func IsEqualSlice(a []byte, b []byte) bool {
	isequal := false
	if len(a) == len(b) {
		for i := 0; i <= len(a)-1; i++ {
			isequal = true && (a[i] == b[i])
		}
	}
	return isequal
}

func FindTargetBlock(barecodex BareCodex, parentModuleHash [32]byte) *ModuleBlock {
	// Currently using a simple linear array to store the entire tree - so we need a separate hash search function
	// and also non-bidirectional - no convenient searches
	// which is inefficient for large codex trees
	// TODO: Define a linked codex tree data structure & simply iterate over it
	targetblock := new(ModuleBlock)
	for _, module := range barecodex.Moduleblocks {
		if sha256.Sum256([]byte(fmt.Sprintf("%v", module))) == parentModuleHash {
			targetblock = &module
		}
	}
	return targetblock
}

func ProcessCodexBlock(params []WASMParams, barecodex *BareCodex) []WASMReturned {
	wasmreturned := []WASMReturned{}
	prevblock := new(ModuleBlock)
	indexblock := new(ModuleBlock)
	// Iterate over a bare codex
	for i, module := range barecodex.Moduleblocks {
		if module.parentModuleHash == [32]byte{} {
			wasmreturned = append(wasmreturned, module.ReturnValues)
			exec := ExecuteWASMBlock(module.ModuleBytecode, params[i])
			wasmreturned = append(wasmreturned, *exec)
		} else {
			// Find the parentmodule block with the given hash from the array
			// If already executed, pass parameter
			// Else, execute parent first
			prevblock = FindTargetBlock(*barecodex, module.parentModuleHash)
			if prevblock == nil {
				panic("No module with given hash found")
			}
			// Check returned state array
			for j, block := range wasmreturned {
				if IsEqualSlice(block.wasmreturned, prevblock.ReturnValues.wasmreturned) && j%2 == 0 {
					indexblock = &ModuleBlock{
						ReturnValues: block,
					}
				}
			}
			if indexblock == nil {
				wasmreturned = append(wasmreturned, *ExecuteWASMBlock(module.ModuleBytecode, params[i]))
			} else {
				params := &WASMParams{
					wasmparams:    wasmreturned[len(wasmreturned)-2].wasmreturned,
					entryfunction: "",
					arguments:     "",
				}
				wasmreturned = append(wasmreturned, *ExecuteWASMBlock(module.ModuleBytecode, *params))
			}
		}
	}
	return wasmreturned
}

func ValidateState(stateblock *StateChainBlock, wasmreturned []byte, params *[]WASMParams, pub *ecdsa.PublicKey, priv *ecdsa.PrivateKey) {
	if IsEqualSlice((*stateblock).state.Returned.wasmreturned, wasmreturned) {
		// Write SHA-256 hash of the entire state
		stateblock.statehash = sha256.Sum256([]byte(fmt.Sprintf("%v", *(*stateblock).state)))
		// Sign the resulting statehash with device private key
		r, s, err := ecdsa.Sign(rand.Reader, priv, stateblock.statehash[:])
		if err != nil {
			panic(err)
		}
		// Verify signature validity
		isvalid := ecdsa.Verify(pub, stateblock.statehash[:], r, s)
		if !isvalid {
			panic("Signature does not match")
		}
		// Write the signed hash to the state block
		(*stateblock).signedstatehash = append((*stateblock).signedstatehash, r, s)
		// Write the public key
		(*stateblock).signedkeys = append((*stateblock).signedkeys, *pub)
	} else {
		// TODO: Call consensus & re-execute from given point
	}
}

/*
// Generates a new state chain
func NewStateChain(module *ModuleBlock, returned *WASMReturned) *StateChain {
	// Calculates hash of attached module block
	modulehash := sha256.Sum256([]byte(fmt.Sprintf("%v", *module)))
	// Generates a blocked state
	blockedstate := &BlockedState{
		Parameters: module.Parameters,
		Returned:   module.ReturnValues,
		state:      nil,
	}
}
*/
