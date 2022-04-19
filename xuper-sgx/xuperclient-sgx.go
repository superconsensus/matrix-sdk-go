package xuper_sgx

import (
	account_sgx "github.com/superconsensus/matrix-sdk-go/v2/account-sgx"
	"github.com/superconsensus/matrix-sdk-go/v2/common"
	"github.com/xuperchain/xuperchain/service/pb"
	"math/big"
	"regexp"
)

// DeployNativeGoContract deploy native go contract.
//
// Parameters:
//   - `from`: Transaction initiator.
//   - `name`: Contract name.
//   - `code`: Contract code bytes.
//   - `args`: Contract init args.
func (x *XClient) DeployNativeGoContract(from *account_sgx.AccountSgx, name string, code []byte, args map[string]string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewDeployContractRequest(from, name, nil, code, args, NativeContractModule, GoRuntime, opts...)
	if err != nil {
		return nil, err
	}
	return x.Do(req)
}

// DeployNativeJavaContract deploy native java contract.
//
// Parameters:
//   - `from`: Transaction initiator.
//   - `name`: Contract name.
//   - `code`: Contract code bytes.
//   - `args`: Contract init args.
func (x *XClient) DeployNativeJavaContract(from *account_sgx.AccountSgx, name string, code []byte, args map[string]string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewDeployContractRequest(from, name, nil, code, args, NativeContractModule, JavaRuntime, opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// DeployWasmContract deploy wasm c++ contract.
//
// Parameters:
//   - `from`: Transaction initiator.
//   - `name`: Contract name.
//   - `code`: Contract code bytes.
//   - `args`: Contract init args.
func (x *XClient) DeployWasmContract(from *account_sgx.AccountSgx, name string, code []byte, args map[string]string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewDeployContractRequest(from, name, nil, code, args, WasmContractModule, CRuntime, opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// DeployEVMContract deploy evm contract.
//
// Parameters:
//   - `from`: Transaction initiator.
//   - `name`: Contract name.
//   - `abi` : Solidity contract abi.
//   - `bin` : Solidity contract bin.
//   - `args`: Contract init args.
func (x *XClient) DeployEVMContract(from *account_sgx.AccountSgx, name string, abi, bin []byte, args map[string]string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewDeployContractRequest(from, name, abi, bin, args, EvmContractModule, "", opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// UpgradeWasmContract upgrade wasm contract.
//
// Parameters:
//   - `from`: Transaction initiator.
//   - `name`: Contract name.
//   - `code`: Contract code bytes.
//   - `args`: Contract init args.
func (x *XClient) UpgradeWasmContract(from *account_sgx.AccountSgx, name string, code []byte, opts ...RequestOption) (*Transaction, error) {
	req, err := NewUpgradeContractRequest(from, WasmContractModule, name, code, opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// UpgradeNativeContract upgrade native contract.
//
// Parameters:
//   - `from`: Transaction initiator.
//   - `name`: Contract name.
//   - `code`: Contract code bytes.
//   - `args`: Contract init args.
func (x *XClient) UpgradeNativeContract(from *account_sgx.AccountSgx, name string, code []byte, opts ...RequestOption) (*Transaction, error) {
	req, err := NewUpgradeContractRequest(from, NativeContractModule, name, code, opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// InvokeWasmContract invoke wasm c++ contract.
//
// Parameters:
//   - `from`  : Transaction initiator.
//   - `name`  : Contract name.
//   - `method`: Contract method.
//   - `args`  : Contract invoke args.
func (x *XClient) InvokeWasmContract(from *account_sgx.AccountSgx, name, method string, args map[string]string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewInvokeContractRequest(from, WasmContractModule, name, method, args, opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// InvokeNativeContract invoke native contract.
//
// Parameters:
//   - `from`  : Transaction initiator.
//   - `name`  : Contract name.
//   - `method`: Contract method.
//   - `args`  : Contract invoke args.
func (x *XClient) InvokeNativeContract(from *account_sgx.AccountSgx, name, method string, args map[string]string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewInvokeContractRequest(from, NativeContractModule, name, method, args, opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// InvokeEVMContract invoke evm contract.
//
// Parameters:
//   - `from`  : Transaction initiator.
//   - `name`  : Contract name.
//   - `method`: Contract method.
//   - `args`  : Contract invoke args.
func (x *XClient) InvokeEVMContract(from *account_sgx.AccountSgx, name, method string, args map[string]string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewInvokeContractRequest(from, EvmContractModule, name, method, args, opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// QueryWasmContract query wasm c++ contract.
//
// Parameters:
//   - `from`  : Transaction initiator.
//   - `name`  : Contract name.
//   - `method`: Contract method.
//   - `args`  : Contract invoke args.
func (x *XClient) QueryWasmContract(from *account_sgx.AccountSgx, name, method string, args map[string]string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewInvokeContractRequest(from, WasmContractModule, name, method, args, opts...)
	if err != nil {
		return nil, err
	}

	return x.PreExecTx(req)
}

// QueryNativeContract query native contract.
//
// Parameters:
//   - `from`  : Transaction initiator.
//   - `name`  : Contract name.
//   - `method`: Contract method.
//   - `args`  : Contract invoke args.
func (x *XClient) QueryNativeContract(from *account_sgx.AccountSgx, name, method string, args map[string]string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewInvokeContractRequest(from, NativeContractModule, name, method, args, opts...)
	if err != nil {
		return nil, err
	}
	return x.PreExecTx(req)
}

// QueryEVMContract query evm contract.
//
// Parameters:
//   - `from`  : Transaction initiator.
//   - `name`  : Contract name.
//   - `method`: Contract method.
//   - `args`  : Contract invoke args.
func (x *XClient) QueryEVMContract(from *account_sgx.AccountSgx, name, method string, args map[string]string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewInvokeContractRequest(from, EvmContractModule, name, method, args, opts...)
	if err != nil {
		return nil, err
	}
	return x.PreExecTx(req)
}

// Transfer to another address.
//
// Parameters:
//   - `from`  : Transaction initiator.
//   - `to`    : Transfer receiving address.
//   - `amount`: Transfer amount.
func (x *XClient) Transfer(from *account_sgx.AccountSgx, to, amount string, opts ...RequestOption) (*Transaction, error) {
	req, err := NewTransferRequest(from, to, amount, opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// CreateContractAccount create contract account for initiator.
//
// Parameters:
//   - `from`           : Transaction initiator. NOTE: from must be NOT set contract account, if you set please remove it.
//   - `contractAccount`:The contract account you want to create, such as: XC8888888899999999@xuper.
func (x *XClient) CreateContractAccount(from *account_sgx.AccountSgx, contractAccount string, opts ...RequestOption) (*Transaction, error) {
	if ok, _ := regexp.MatchString(`^XC\d{16}@*`, contractAccount); !ok {
		return nil, common.ErrInvalidContractAccount
	}

	subRegexp := regexp.MustCompile(`\d{16}`)
	contractAccountByte := subRegexp.Find([]byte(contractAccount))
	contractAccount = string(contractAccountByte)
	req, err := NewCreateContractAccountRequest(from, contractAccount, opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// SetAccountACL update contract account acl. NOTE: from account must be set contract account.
//
// Parameters:
//   - `from`: Transaction initiator.
//   - `acl` : The ACL you want to set.
func (x *XClient) SetAccountACL(from *account_sgx.AccountSgx, acl *ACL, opts ...RequestOption) (*Transaction, error) {
	req, err := NewSetAccountACLRequest(from, acl, opts...)
	if err != nil {
		return nil, err
	}
	return x.Do(req)
}

// SetMethodACL update contract method acl.
//
// Parameters:
//   - `from`  : Transaction initiator.
//   - `name`  : Contract name.
//   - `method`: Contract method.
//   - `acl`   : The ACL you want to set.
func (x *XClient) SetMethodACL(from *account_sgx.AccountSgx, name, method string, acl *ACL, opts ...RequestOption) (*Transaction, error) {
	req, err := NewSetMethodACLRequest(from, name, method, acl, opts...)
	if err != nil {
		return nil, err
	}

	return x.Do(req)
}

// WatchBlockEvent new watcher for block event.
//func (x *XClient) WatchBlockEvent(opts ...BlockEventOption) (*Watcher, error) {
//	watcher, err := x.newWatcher(opts...)
//	if err != nil {
//		return nil, err
//	}
//	buf, _ := proto.Marshal(watcher.opt.blockFilter)
//	request := &pb.SubscribeRequest{
//		Type:   pb.SubscribeType_BLOCK,
//		Filter: buf,
//	}
//
//	stream, err := x.esc.Subscribe(context.TODO(), request)
//	if err != nil {
//		return nil, err
//	}
//
//	filteredBlockChan := make(chan *FilteredBlock, watcher.opt.blockChanBufferSize)
//	exit := make(chan struct{})
//	watcher.exit = exit
//	watcher.FilteredBlockChan = filteredBlockChan
//
//	go func() {
//		defer func() {
//			close(filteredBlockChan)
//			if err := stream.CloseSend(); err != nil {
//				log.Printf("Unregister block event failed, close stream error: %v", err)
//			} else {
//				log.Printf("Unregister block event success...")
//			}
//		}()
//		for {
//			select {
//			case <-exit:
//				return
//			default:
//				event, err := stream.Recv()
//				if err == io.EOF {
//					return
//				}
//				if err != nil {
//					log.Printf("Get block event err: %v", err)
//					return
//				}
//				var block pb.FilteredBlock
//				err = proto.Unmarshal(event.Payload, &block)
//				if err != nil {
//					log.Printf("Get block event err: %v", err)
//					return
//				}
//				if len(block.GetTxs()) == 0 && watcher.opt.skipEmptyTx {
//					continue
//				}
//				filteredBlockChan <- fromFilteredBlockPB(&block)
//			}
//		}
//	}()
//	return watcher, nil
//}
//
//func (x *XClient) newWatcher(opts ...BlockEventOption) (*Watcher, error) {
//	opt, err := initEventOpts(opts...)
//	if err != nil {
//		return nil, err
//	}
//
//	watcher := &Watcher{
//		opt: opt,
//	}
//	return watcher, nil
//}

////////////////////////////////////////////////
//	         query
///////////////////////////////////////////////
// QueryTxByID query the tx by txID
//
// Parameters
//  - `txID` : transaction id
func (x *XClient) QueryTxByID(txID string, opts ...QueryOption) (*pb.Transaction, error) {
	return x.queryTxByID(txID, opts...)
}

// QueryBlockByID query the block by blockID
//
// Parameters:
//   - `blockID`  : block id
func (x *XClient) QueryBlockByID(blockID string, opts ...QueryOption) (*pb.Block, error) {
	return x.queryBlockByID(blockID, opts...)
}

// QueryBlockByHeight query the block by block height
//
// Parameters:
//   - `height`  : block height
func (x *XClient) QueryBlockByHeight(height int64, opts ...QueryOption) (*pb.Block, error) {
	return x.queryBlockByHeight(height, opts...)
}

// QueryAccountACL query the ACL by account
//
// Parameters:
//   - `account`  : account, such as XC1111111111111111@xuper
func (x *XClient) QueryAccountACL(account string, opts ...QueryOption) (*ACL, error) {
	return x.queryAccountACL(account, opts...)
}

// QueryMethodACL query the ACL by method
//
// Parameters:
//   - `name`     : contract name
//   - `account`  : account
func (x *XClient) QueryMethodACL(name, method string, opts ...QueryOption) (*ACL, error) {
	return x.queryMethodACL(name, method, opts...)
}

// QueryAccountContracts query all contracts for account
//
// Parameters:
//   - `account`  : account,such as XC1111111111111111@xuper
func (x *XClient) QueryAccountContracts(account string, opts ...QueryOption) ([]*pb.ContractStatus, error) {
	return x.queryAccountContracts(account, opts...)
}

// QueryAddressContracts query all contracts for address
//
// Parameters:
//   - `address`  : address
//
// Returns:
//   - `map`  : contractAccount => contractStatusList
//   - `error`: error
func (x *XClient) QueryAddressContracts(address string, opts ...QueryOption) (map[string]*pb.ContractList, error) {
	return x.queryAddressContracts(address, opts...)
}

// QueryBalance query balance by the address
//
// Parameters:
//   - `address`  : address
func (x *XClient) QueryBalance(address string, opts ...QueryOption) (*big.Int, error) {
	return x.queryBalance(address, opts...)
}

// QueryBalanceDetail query the balance detail by address
//
// Parameters:
//   - `address`  : address
func (x *XClient) QueryBalanceDetail(address string, opts ...QueryOption) ([]*BalanceDetail, error) {
	return x.queryBalanceDetail(address, opts...)
}

// QuerySystemStatus query the system status
func (x *XClient) QuerySystemStatus(opts ...QueryOption) (*pb.SystemsStatusReply, error) {
	return x.querySystemStatus(opts...)
}

// QueryBlockChains query block chains
func (x *XClient) QueryBlockChains(opts ...QueryOption) ([]string, error) {
	return x.queryBlockChains(opts...)
}

// QueryBlockChainStatus query the block chain status
func (x *XClient) QueryBlockChainStatus(opts ...QueryOption) (*pb.BCStatus, error) {
	return x.queryBlockChainStatus(opts...)
}

// QueryNetURL query the net URL
func (x *XClient) QueryNetURL(opts ...QueryOption) (string, error) {
	return x.queryNetURL(opts...)
}

// QueryAccountByAK query the account  by AK
//
// Parameters:
//   - `address`  : address
func (x *XClient) QueryAccountByAK(address string, opts ...QueryOption) ([]string, error) {
	return x.queryAccountByAK(address, opts...)
}
