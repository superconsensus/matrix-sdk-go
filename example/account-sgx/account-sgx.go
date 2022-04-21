package main

import (
	"fmt"
	account_sgx "github.com/superconsensus/matrix-sdk-go/v2/account-sgx"
	xuper_sgx "github.com/superconsensus/matrix-sdk-go/v2/xuper-sgx"
)

var (
	// sgx 服务地址
	URL = "http://127.0.0.1:8080"
	// 地址
	Addr = "jB3iS35PCdUpDZ879LHJUqLHCzxETftXG"
)

// 测试创建账号
func testAccount() {
	accsgx, err := account_sgx.CreateAccountSgx(URL)
	if err != nil {
		fmt.Printf("create account error: %v\n", err)
	}
	fmt.Println(accsgx.Address)
}

// 测试合约账号
func testContractAccount() {
	// 通过address恢复账号
	accsgx, err := account_sgx.RetrieveAccountSgx(URL, Addr)
	if err != nil {
		fmt.Printf("retrieveAccount err: %v\n", err)
		return
	}
	fmt.Printf("retrieveAccount address: %v\n", accsgx.Address)

	// 创建一个合约账号
	// 创建一个合约账户
	// 合约账户是由 (XC + 16个数字 + @xuper) 组成, 比如 "XC1234567890123456@xuper"
	contractAccount := "XC1234567890123451@xuper"

	xchainClient, err := xuper_sgx.New("127.0.0.1:37101")
	tx, err := xchainClient.CreateContractAccount(accsgx, contractAccount)
	if err != nil {
		fmt.Printf("createContractAccount err:%s\n", err.Error())
	}
	fmt.Printf("%x", tx.Tx.Txid)
	return
}

func main() {
	//testAccount()
	testContractAccount()
}
