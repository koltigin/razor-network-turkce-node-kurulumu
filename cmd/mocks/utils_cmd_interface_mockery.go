// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	accounts "github.com/ethereum/go-ethereum/accounts"

	bindings "razor/pkg/bindings"

	common "github.com/ethereum/go-ethereum/common"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	pflag "github.com/spf13/pflag"

	types "razor/core/types"
)

// UtilsCmdInterfaceMockery is an autogenerated mock type for the UtilsCmdInterfaceMockery type
type UtilsCmdInterfaceMockery struct {
	mock.Mock
}

// AssignAmountInWei provides a mock function with given fields: _a0
func (_m *UtilsCmdInterfaceMockery) AssignAmountInWei(_a0 *pflag.FlagSet) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClaimBlockReward provides a mock function with given fields: _a0
func (_m *UtilsCmdInterfaceMockery) ClaimBlockReward(_a0 types.TransactionOptions) (common.Hash, error) {
	ret := _m.Called(_a0)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(types.TransactionOptions) common.Hash); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.TransactionOptions) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClaimBounty provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterfaceMockery) ClaimBounty(_a0 types.Configurations, _a1 *ethclient.Client, _a2 types.RedeemBountyInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(types.Configurations, *ethclient.Client, types.RedeemBountyInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Configurations, *ethclient.Client, types.RedeemBountyInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *UtilsCmdInterfaceMockery) Commit(_a0 *ethclient.Client, _a1 []*big.Int, _a2 []byte, _a3 types.Account, _a4 types.Configurations) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, []*big.Int, []byte, types.Account, types.Configurations) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, []*big.Int, []byte, types.Account, types.Configurations) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCollection provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterfaceMockery) CreateCollection(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.CreateCollectionInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.CreateCollectionInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.CreateCollectionInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJob provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterfaceMockery) CreateJob(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.CreateJobInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.CreateJobInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.CreateJobInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteClaimBounty provides a mock function with given fields: _a0
func (_m *UtilsCmdInterfaceMockery) ExecuteClaimBounty(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteCreateCollection provides a mock function with given fields: _a0
func (_m *UtilsCmdInterfaceMockery) ExecuteCreateCollection(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteCreateJob provides a mock function with given fields: _a0
func (_m *UtilsCmdInterfaceMockery) ExecuteCreateJob(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// ExecuteTransfer provides a mock function with given fields: _a0
func (_m *UtilsCmdInterfaceMockery) ExecuteTransfer(_a0 *pflag.FlagSet) {
	_m.Called(_a0)
}

// GetBufferPercent provides a mock function with given fields:
func (_m *UtilsCmdInterfaceMockery) GetBufferPercent() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigData provides a mock function with given fields:
func (_m *UtilsCmdInterfaceMockery) GetConfigData() (types.Configurations, error) {
	ret := _m.Called()

	var r0 types.Configurations
	if rf, ok := ret.Get(0).(func() types.Configurations); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.Configurations)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochAndState provides a mock function with given fields: _a0
func (_m *UtilsCmdInterfaceMockery) GetEpochAndState(_a0 *ethclient.Client) (uint32, int64, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(*ethclient.Client) int64); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*ethclient.Client) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetGasLimit provides a mock function with given fields:
func (_m *UtilsCmdInterfaceMockery) GetGasLimit() (float32, error) {
	ret := _m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasPrice provides a mock function with given fields:
func (_m *UtilsCmdInterfaceMockery) GetGasPrice() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogLevel provides a mock function with given fields:
func (_m *UtilsCmdInterfaceMockery) GetLogLevel() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMultiplier provides a mock function with given fields:
func (_m *UtilsCmdInterfaceMockery) GetMultiplier() (float32, error) {
	ret := _m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvider provides a mock function with given fields:
func (_m *UtilsCmdInterfaceMockery) GetProvider() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWaitTime provides a mock function with given fields:
func (_m *UtilsCmdInterfaceMockery) GetWaitTime() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleCommitState provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterfaceMockery) HandleCommitState(_a0 *ethclient.Client, _a1 uint32, _a2 types.Rogue) ([]*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []*big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, types.Rogue) []*big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, types.Rogue) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleRevealState provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterfaceMockery) HandleRevealState(_a0 *ethclient.Client, _a1 bindings.StructsStaker, _a2 uint32) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, bindings.StructsStaker, uint32) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListAccounts provides a mock function with given fields:
func (_m *UtilsCmdInterfaceMockery) ListAccounts() ([]accounts.Account, error) {
	ret := _m.Called()

	var r0 []accounts.Account
	if rf, ok := ret.Get(0).(func() []accounts.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reveal provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *UtilsCmdInterfaceMockery) Reveal(_a0 *ethclient.Client, _a1 []*big.Int, _a2 []byte, _a3 types.Account, _a4 string, _a5 types.Configurations) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, []*big.Int, []byte, types.Account, string, types.Configurations) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, []*big.Int, []byte, types.Account, string, types.Configurations) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetConfig provides a mock function with given fields: flagSet
func (_m *UtilsCmdInterfaceMockery) SetConfig(flagSet *pflag.FlagSet) error {
	ret := _m.Called(flagSet)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) error); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transfer provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterfaceMockery) Transfer(_a0 *ethclient.Client, _a1 types.Configurations, _a2 types.TransferInput) (common.Hash, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(*ethclient.Client, types.Configurations, types.TransferInput) common.Hash); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, types.Configurations, types.TransferInput) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForAppropriateState provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsCmdInterfaceMockery) WaitForAppropriateState(_a0 *ethclient.Client, _a1 string, _a2 ...int) (uint32, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, ...int) uint32); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, ...int) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
