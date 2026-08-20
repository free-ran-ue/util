package util

import (
	"encoding/hex"
	"strings"

	"github.com/free5gc/ngap/ie"
	"github.com/free5gc/openapi/models"
)

func PlmnIdToModels(ngapPlmnId ie.PLMNIdentity) (modelsPlmnid models.PlmnId) {
	value := ngapPlmnId.Value
	hexString := strings.Split(hex.EncodeToString(value), "")
	modelsPlmnid.Mcc = hexString[1] + hexString[0] + hexString[3]
	if hexString[2] == "f" {
		modelsPlmnid.Mnc = hexString[5] + hexString[4]
	} else {
		modelsPlmnid.Mnc = hexString[2] + hexString[5] + hexString[4]
	}
	return
}

func PlmnIdToNgap(modelsPlmnid models.PlmnId) (ie.PLMNIdentity, error) {
	var hexString string
	mcc := strings.Split(modelsPlmnid.Mcc, "")
	mnc := strings.Split(modelsPlmnid.Mnc, "")
	if len(modelsPlmnid.Mnc) == 2 {
		hexString = mcc[1] + mcc[0] + "f" + mcc[2] + mnc[1] + mnc[0]
	} else {
		hexString = mcc[1] + mcc[0] + mnc[0] + mcc[2] + mnc[2] + mnc[1]
	}

	var ngapPlmnId ie.PLMNIdentity
	if plmnId, err := hex.DecodeString(hexString); err != nil {
		return ngapPlmnId, err
	} else {
		ngapPlmnId.Value = plmnId
	}
	return ngapPlmnId, nil
}

func TaiToModels(tai ie.TAI) models.Tai {
	var modelsTai models.Tai

	plmnID := PlmnIdToModels(*tai.PLMNIdentity)
	modelsTai.PlmnId = &plmnID
	modelsTai.Tac = hex.EncodeToString(tai.TAC.Value)

	return modelsTai
}

func TaiToNgap(tai models.Tai) (ie.TAI, error) {
	var ngapTai ie.TAI
	var err error

	var plmnId ie.PLMNIdentity
	plmnId, err = PlmnIdToNgap(*tai.PlmnId)
	if err != nil {
		return ngapTai, err
	}
	ngapTai.PLMNIdentity = &plmnId

	if tac, err := hex.DecodeString(tai.Tac); err != nil {
		return ngapTai, err
	} else {
		ngapTai.TAC = &ie.TAC{Value: tac}
	}
	return ngapTai, nil
}

func SNssaiToModels(ngapSnssai ie.SNSSAI) (modelsSnssai models.Snssai) {
	modelsSnssai.Sst = int32(ngapSnssai.SST.Value[0])
	if ngapSnssai.SD != nil {
		modelsSnssai.Sd = hex.EncodeToString(ngapSnssai.SD.Value)
	}
	return
}

func SNssaiToNgap(modelsSnssai models.Snssai) (ie.SNSSAI, error) {
	var ngapSnssai ie.SNSSAI
	ngapSnssai.SST = &ie.SST{Value: []byte{byte(modelsSnssai.Sst)}}

	if modelsSnssai.Sd != "" {
		ngapSnssai.SD = new(ie.SD)
		if sdTmp, err := hex.DecodeString(modelsSnssai.Sd); err != nil {
			return ngapSnssai, err
		} else {
			ngapSnssai.SD.Value = sdTmp
		}
	}
	return ngapSnssai, nil
}
