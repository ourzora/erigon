package commands

import (
	"time"

	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/core/vm"
)

// Helper implementation of vm.Tracer; since the interface is big and most
// custom tracers implement just a few of the methods, this is a base struct
// to avoid lots of empty boilerplate code
type DefaultTracer struct {
}

func (t *DefaultTracer) CaptureTxStart(gasLimit uint64) {}

func (t *DefaultTracer) CaptureTxEnd(restGas uint64) {}

func (t *DefaultTracer) CaptureStart(env *vm.EVM, from common.Address, to common.Address, precompile bool, create bool, calltype vm.CallType, input []byte, gas uint64, value *uint256.Int, code []byte) {
}

func (t *DefaultTracer) CaptureEnter(env *vm.EVM, from common.Address, to common.Address, precompile bool, create bool, calltype vm.CallType, input []byte, gas uint64, value *uint256.Int, code []byte) {
}

func (t *DefaultTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {
}

func (t *DefaultTracer) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, depth int, err error) {
}

func (t *DefaultTracer) CaptureEnd(output []byte, startGas, endGas uint64, d time.Duration, err error) {
}

func (t *DefaultTracer) CaptureExit(output []byte, startGas, endGas uint64, d time.Duration, err error) {
}

func (t *DefaultTracer) CaptureSelfDestruct(from common.Address, to common.Address, value *uint256.Int) {
}

func (t *DefaultTracer) CaptureAccountRead(account common.Address) error {
	return nil
}

func (t *DefaultTracer) CaptureAccountWrite(account common.Address) error {
	return nil
}
