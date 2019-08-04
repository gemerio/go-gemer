package codex

import (
	"os"
	"time"

	"github.com/gemer/go-gemer/consensus"
)

type Codex struct {
	barecodex *BareCodex
	statechain *StateChain
	keychain *Keychain
}