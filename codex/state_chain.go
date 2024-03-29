package codex

import (
	"crypto"
	"math/big"
)

// Defines a state chain block
type StateChainBlock struct {
	// previous block hash - regardless of bytecode or state block
	PrevHash [32]byte
	// Committed state
	state *BlockedState
	// SHA-256 hash of this commit round
	statehash [32]byte
	// An array of signed SHA-256 hash values of this commit round
	signedstatehash []*big.Int
	// An array of all keys that signed this state chain
	signedkeys []crypto.PublicKey
}

// Defines a "blocked state" - state data that may be omitted during a node sync
type BlockedState struct {
	// Passed parameters of a single execution session
	Parameters WASMParams
	// Returned values of a single execution session
	Returned WASMReturned
	// Finalized state array
	state []byte
}

// Defines a state chain
type StateChain struct {
	// Attatched Function Hash
	PrevHash [64]byte
	// An array of state blocks
	StateBlocks []StateChainBlock
	// SHA-256 hash of the entire state chain
	statechainhash [64]byte
}
