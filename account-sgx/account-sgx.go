package account_sgx

import (
	"fmt"
	"github.com/superconsensus/matrix-sdk-go/v2/common"
	"log"
	"regexp"
	"strconv"
)

type AccountSgx struct {
	// 钱包地址
	Address string
	// 合约账号
	contractAccount string
	// api client
	APISgx ApiClient
}

// 创建账号
func CreateAccountSgx(url string) (*AccountSgx, error) {
	// 创建api client
	apiclient := NewApiClientXuperchain(url)

	result, err := apiclient.Create(CreateMethod, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if result.Code != 200 {
		return nil, fmt.Errorf("%s", "create error")
	}

	accountsgx := &AccountSgx{
		Address: string(result.Data),
		APISgx:  apiclient,
	}
	return accountsgx, nil
}

// 恢复账号
func RetrieveAccountSgx(url, addr string) (*AccountSgx, error) {
	if url == "" || addr == "" {
		return nil, fmt.Errorf("%s", "nil url or addr")
	}
	// 创建api client
	apiclient := NewApiClientXuperchain(url)
	// 检查sgx中是否有 addr
	isExistArgs := map[string]interface{}{
		"address": addr,
	}
	result, err := apiclient.IsExist(IsExistMethod, isExistArgs)
	if err != nil {
		return nil, err
	}
	flag, _ := strconv.ParseBool(string(result.Data))
	if flag {
		accountsgx := &AccountSgx{
			Address: addr,
			APISgx:  apiclient,
		}
		return accountsgx, nil
	}
	return nil, fmt.Errorf("RetrieveAccountSgx error")
}

// SetContractAccount set contract account.
// If you set contract account, this account represents the contract account.
// In some scenarios, must set contract account, such as deploy contract.
func (a *AccountSgx) SetContractAccount(contractAccount string) error {
	if ok, _ := regexp.MatchString(`^XC\d{16}@*`, contractAccount); !ok {
		return common.ErrInvalidContractAccount
	}

	a.contractAccount = contractAccount
	return nil
}

// GetAuthRequire get this account's authRequire for transaction.
// If you set contract account, returns $ContractAccount+"/"+$Address, otherwise returns $Address.
func (a *AccountSgx) GetAuthRequire() string {
	if a.HasContractAccount() {
		return a.GetContractAccount() + "/" + a.Address
	}
	return a.Address
}

// GetContractAccount get current contract account, returns an empty string if the contract account is not set.
func (a *AccountSgx) GetContractAccount() string {
	return a.contractAccount
}

// HasContractAccount reutrn true if you set contract account, otherwise returns false.
func (a *AccountSgx) HasContractAccount() bool {
	return a.contractAccount != ""
}
