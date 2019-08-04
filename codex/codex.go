package codex

import (
	"os"
	"time"

	"github.com/gemerio/go-gemer/consensus"
)

type Codex struct {
	barecodex  *BareCodex
	statechain *StateChain
	keychain   *Keychain
}
