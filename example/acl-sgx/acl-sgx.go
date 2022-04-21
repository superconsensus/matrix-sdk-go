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
	setAccountACLExample()
}

func setAccountACLExample() {
	// 通过address恢复账号
	accsgx, err := account_sgx.RetrieveAccountSgx(URL, Addr)
	if err != nil {
		fmt.Printf("retrieveAccount err: %v\n", err)
		return
	}
	fmt.Printf("retrieveAccount address: %v\n", accsgx.Address)

	// 创建节点客户端。
	node := "127.0.0.1:37101"
	xclient, err := xuper_sgx.New(node)
	if err != nil {
		panic(err)
	}
	defer xclient.Close()

	contractAcc := "XC1234567890123451@xuper"
	queryACL, err := xclient.QueryAccountACL(contractAcc)
	if err != nil {
		panic(err)
	}
	// 查看修改后的 ACL。
	fmt.Println(queryACL)
}
