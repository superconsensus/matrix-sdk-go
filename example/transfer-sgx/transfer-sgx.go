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

func main() {
	//akTransfer()
	contractAccountTransfer()
}

func akTransfer() {
	// 通过address恢复账号
	from, err := account_sgx.RetrieveAccountSgx(URL, Addr)
	if err != nil {
		fmt.Printf("retrieveAccount err: %v\n", err)
		return
	}
	fmt.Printf("retrieveAccount address: %v\n", from.Address)

	// 创建账号
	to, err := account_sgx.CreateAccountSgx(URL)
	if err != nil {
		fmt.Printf("create account error: %v\n", err)
		panic(err)
	}
	fmt.Println(to.Address)

	// 节点地址。
	node := "127.0.0.1:37101"

	// 创建节点客户端。
	xclient, _ := xuper_sgx.New(node)

	// 转账前查看两个地址余额。
	fmt.Println(xclient.QueryBalance(from.Address))
	fmt.Println(xclient.QueryBalance(to.Address))

	tx, err := xclient.Transfer(from, to.Address, "10")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", tx.Tx.Txid)

	// 转账后查看两个地址余额。
	fmt.Println(xclient.QueryBalance(from.Address))
	fmt.Println(xclient.QueryBalance(to.Address))
}

// contractAccountTransfer 合约账户转账示例。
func contractAccountTransfer() {
	// 通过address恢复账号
	from, err := account_sgx.RetrieveAccountSgx(URL, Addr)
	if err != nil {
		fmt.Printf("retrieveAccount err: %v\n", err)
		return
	}
	fmt.Printf("retrieveAccount address: %v\n", from.Address)

	from.SetContractAccount("XC1234567890123451@xuper")

	// 创建账号
	to, err := account_sgx.CreateAccountSgx(URL)
	if err != nil {
		fmt.Printf("create account error: %v\n", err)
		panic(err)
	}
	fmt.Println(to.Address)

	// 节点地址。
	node := "127.0.0.1:37101"
	xclient, _ := xuper_sgx.New(node)

	// 转账前查看两个地址余额。
	fmt.Println(xclient.QueryBalance(from.Address))
	fmt.Println(xclient.QueryBalance(to.Address))

	tx, err := xclient.Transfer(from, "a", "10")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", tx.Tx.Txid)

	// 转账后查看两个地址余额。
	fmt.Println(xclient.QueryBalance(from.GetContractAccount())) // 转账时使用的是合约账户，因此查询余额时也是合约账户。
	fmt.Println(xclient.QueryBalance(to.Address))
}
