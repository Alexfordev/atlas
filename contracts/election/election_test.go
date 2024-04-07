package election

import (
	"testing"

	"github.com/Alexfordev/atlas/contracts"
	"github.com/Alexfordev/atlas/contracts/testutil"
)

func TestGetElectedValidators(t *testing.T) {
	testutil.TestFailOnFailingRunner(t, GetElectedValidators)
	testutil.TestFailsWhenContractNotDeployed(t, contracts.ErrSmartContractNotDeployed, GetElectedValidators)
}
