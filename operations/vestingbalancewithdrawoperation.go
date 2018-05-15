package operations

//go:generate ffjson   $GOFILE

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	op := &VestingBalanceWithdrawOperation{}
	types.OperationMap[op.Type()] = op
}

type VestingBalanceWithdrawOperation struct {
	Amount         types.AssetAmount `json:"amount"`
	Fee            types.AssetAmount `json:"fee"`
	Owner          types.GrapheneID  `json:"owner"`
	VestingBalance types.GrapheneID  `json:"vesting_balance"`
}

func (p *VestingBalanceWithdrawOperation) ApplyFee(fee types.AssetAmount) {
	p.Fee = fee
}

func (p VestingBalanceWithdrawOperation) Type() types.OperationType {
	return types.OperationTypeVestingBalanceWithdraw
}

//TODO: define!
func (p VestingBalanceWithdrawOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode operation id")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	return nil
}

//NewVestingBalanceWithdrawOperation creates a new VestingBalanceWithdrawOperation
func NewVestingBalanceWithdrawOperation() *VestingBalanceWithdrawOperation {
	tx := VestingBalanceWithdrawOperation{}
	return &tx
}