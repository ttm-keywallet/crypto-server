package handlers

// #cgo CFLAGS: -I../wallet-core/include
// #cgo LDFLAGS: -L../wallet-core/build -L../wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
// #include <TrustWalletCore/TWBitcoinScript.h>
// #include <TrustWalletCore/TWAnySigner.h>
import "C"

import (
	"encoding/json"
	"github.com/ttmbank/crypto-server/common/handlers"
	"github.com/ttmbank/crypto-server/common/types"
	"net/http"
)

func GetScript(w http.ResponseWriter, req *http.Request){
	var r struct {
		CoinType uint32 `json:"coin_type"`
		Address string `json:"address"`
	}

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		handlers.ERROR_BAD_REQUEST(w, err.Error())
		return
	}
	twAddress := types.TWStringCreateWithGoString(r.Address)
	defer C.TWStringDelete(twAddress)

	script := C.TWBitcoinScriptLockScriptForAddress(twAddress, r.CoinType)
	scriptData := C.TWBitcoinScriptData(script)
	defer C.TWBitcoinScriptDelete(script)
	defer C.TWDataDelete(scriptData)

	scriptHex := types.TWDataHexString(scriptData)

	response := struct {
		Script string `json:"script"`
	}{Script: scriptHex}

	handlers.ReturnResult(w, response)

}
