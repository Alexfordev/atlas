package atlasapi

import (
	"fmt"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/params"
	"github.com/Alexfordev/atlas/core/types"
	"github.com/ethereum/go-ethereum/trie"
	"testing"
)

func Test01(t *testing.T) {
	EmptyRootHash0 := types.DeriveSha(types.Transactions{}, trie.NewStackTrie(nil))
	fmt.Println(EmptyRootHash0)
}
