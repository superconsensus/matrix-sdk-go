package account_sgx

import (
	"encoding/json"
	"log"
	"testing"
)

var (
	addr = "V6Avp9KqLfwGUaRFPV27a8VZhwxiYofoU"
	msg  = []byte("123456")
)

func Test_ApiClient(t *testing.T) {
	apiclient := NewApiClientXuperchain(URL)

	//进行服务测试
	Ping_test(apiclient)
	//CreateMethod_test(apiclient)
	SignMethod_and_Verify_test(apiclient)
}

func Ping_test(client ApiClient) {
	result, err := client.Ping(PingMethod, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(result)
}

func CreateMethod_test(client ApiClient) {
	result, err := client.Create(CreateMethod, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(result)
}

func SignMethod_and_Verify_test(client ApiClient) {
	// sign
	signArgs := map[string]interface{}{
		"address": addr,
		"msg":     msg,
	}
	signResult, err := client.Sign(SignMethod, signArgs)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(signResult)

	//////////////////////////////////////////////
	signInfo := struct {
		PublicKey string `json:"public_key"`
		Sign      []byte `json:"sign"`
	}{}
	err = json.Unmarshal(signResult.Data, &signInfo)
	if err != nil {
		return
	}
	log.Println("signInfo.pub:", signInfo.PublicKey)
	log.Println("signInfo.sign:", signInfo.Sign)
	//////////////////////////////////////////////

	// verify
	verifyArgs := map[string]interface{}{
		"address": addr,
		"sign":    signResult.Data,
		"msg":     msg,
	}
	verifyResult, err := client.Verify(VerifyMethod, verifyArgs)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(verifyResult.Data))
}
