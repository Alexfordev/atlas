package vmcontext

import (
	"github.com/Alexfordev/atlas/cmd/marker/mapprotocol"
	"github.com/ethereum/go-ethereum/log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Alexfordev/atlas/core/state"
	"github.com/Alexfordev/atlas/core/types"
	"github.com/Alexfordev/atlas/core/vm"
	"github.com/Alexfordev/atlas/params"
)

// VMAddress is the address the VM uses to make internal calls to contracts
var VMAddress = params.ZeroAddress

// evmRunnerContext defines methods required to create an EVMRunner
type evmRunnerContext interface {
	chainContext

	// GetVMConfig returns the node's vm configuration
	GetVMConfig() *vm.Config

	CurrentHeader() *types.Header

	State() (*state.StateDB, error)
}

func NewEVMRunner(chain evmRunnerContext, header *types.Header, state types.StateDB) vm.EVMRunner {

	return &evmRunner{
		state: state,
		newEVM: func(from common.Address) *vm.EVM {
			// The EVM Context requires a msg, but the actual field values don't really matter for this case.
			// Putting in zero values for gas price and tx fee recipient
			context := New(from, common.Big0, header, chain, nil)
			return vm.NewEVM(context, vm.TxContext{}, state, chain.Config(), *chain.GetVMConfig())
		},
	}
}

type evmRunner struct {
	newEVM func(from common.Address) *vm.EVM
	state  types.StateDB

	dontMeterGas bool
}

func (ev *evmRunner) Execute(recipient common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, err error) {
	evm := ev.newEVM(VMAddress)
	if ev.dontMeterGas {
		evm.StopGasMetering()
	}
	ret, leftOverGas, err := evm.Call(vm.AccountRef(evm.Origin), recipient, input, gas, value)
	if recipient == mapprotocol.MustProxyAddressFor("Election") {
		log.Info("Log evm Execute Election", "recipient", recipient, "leftOverGas", leftOverGas, "gas", gas, "dontMeterGas", ev.dontMeterGas)
	}
	return ret, err
}

func (ev *evmRunner) ExecuteFrom(sender, recipient common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, err error) {
	evm := ev.newEVM(sender)
	if ev.dontMeterGas {
		evm.StopGasMetering()
	}
	ret, _, err = evm.Call(vm.AccountRef(sender), recipient, input, gas, value)
	return ret, err
}

func (ev *evmRunner) Query(recipient common.Address, input []byte, gas uint64) (ret []byte, err error) {
	evm := ev.newEVM(VMAddress)
	if ev.dontMeterGas {
		evm.StopGasMetering()
	}
	ret, _, err = evm.StaticCall(vm.AccountRef(evm.Origin), recipient, input, gas)
	return ret, err
}

func (ev *evmRunner) StopGasMetering() {
	ev.dontMeterGas = true
}

func (ev *evmRunner) StartGasMetering() {
	ev.dontMeterGas = false
}

// GetStateDB implements Backend.GetStateDB
func (ev *evmRunner) GetStateDB() types.StateDB {
	return ev.state
}

// SharedEVMRunner is an evm runner that REUSES an evm
// This MUST NOT BE USED, but it's here for backward compatibility
// purposes
type SharedEVMRunner struct{ *vm.EVM }

func (sev *SharedEVMRunner) Execute(recipient common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, err error) {
	ret, _, err = sev.Call(vm.AccountRef(VMAddress), recipient, input, gas, value)
	return ret, err
}

func (sev *SharedEVMRunner) ExecuteFrom(sender, recipient common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, err error) {
	ret, _, err = sev.Call(vm.AccountRef(sender), recipient, input, gas, value)
	return ret, err
}

func (sev *SharedEVMRunner) Query(recipient common.Address, input []byte, gas uint64) (ret []byte, err error) {
	ret, _, err = sev.StaticCall(vm.AccountRef(VMAddress), recipient, input, gas)
	return ret, err
}
