package xuper_sgx

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/superconsensus/matrix-sdk-go/v2/account-sgx"
	"github.com/superconsensus/matrix-sdk-go/v2/common"
	"github.com/xuperchain/xuperchain/service/pb"
	"strings"
)

// Transaction xuperchain transaction.
type Transaction struct {
	Tx               *pb.Transaction
	ContractResponse *pb.ContractResponse
	Bcname           string

	Fee     string
	GasUsed int64

	DigestHash []byte
}

// Sign account sign for tx, for multisign.multisign
func (t *Transaction) Sign(account *account_sgx.AccountSgx) error {
	if account == nil {
		return errors.New("Transaction sign account can not be nil")
	}
	// 对于多签，在交易预执行时就需要写好所有的需要签名的地址到 AuthRequire 字段，其他地址再进行签名时，需要检查是否已经在 AuthRequire 字段中。
	// 同时签名的顺序也要保持一致，不然上链时会失败。
	if !inSlice(t.Tx.AuthRequire, account.GetAuthRequire()) {
		return errors.New("this account not in transaction's AuthRequire list")
	}

	if t.DigestHash == nil {
		digestHash, err := common.MakeTxDigestHash(t.Tx)
		if err != nil {
			return err
		}
		t.DigestHash = digestHash
	}

	//cryptoClient := crypto.GetCryptoClient()
	//privateKey, err := cryptoClient.GetEcdsaPrivateKeyFromJsonStr(account.PrivateKey)
	//if err != nil {
	//	return err
	//}
	//
	//sign, err := cryptoClient.SignECDSA(privateKey, t.DigestHash)
	//
	//signatureInfo := &pb.SignatureInfo{
	//	PublicKey: account.PublicKey,
	//	Sign:      sign,
	//}

	/////////////////////////////////////
	///  签名改造 --- look here
	////////////////////////////////////
	signArgs := map[string]interface{}{
		"address": account.Address,
		"msg":     t.DigestHash,
	}
	result, err := account.APISgx.Sign(account_sgx.SignMethod, signArgs)
	if err != nil {
		return errors.New("sign error")
	}
	signInfo := struct {
		PublicKey string `json:"public_key"`
		Sign      []byte `json:"sign"`
	}{}
	err = json.Unmarshal(result.Data, &signInfo)
	if err != nil {
		return err
	}
	signatureInfo := &pb.SignatureInfo{
		PublicKey: signInfo.PublicKey,
		Sign:      signInfo.Sign,
	}
	//////////////////////////////////////
	//////////////////////////////////////

	t.Tx.AuthRequireSigns = append(t.Tx.AuthRequireSigns, signatureInfo)
	t.Tx.InitiatorSigns = append(t.Tx.InitiatorSigns, signatureInfo)

	// make txid
	t.Tx.Txid, err = common.MakeTransactionID(t.Tx)

	return err
}

func inSlice(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}

		splitRes := strings.Split(v, "/")
		addr := splitRes[len(splitRes)-1]
		if addr == str {
			return true
		}
	}
	return false
}
