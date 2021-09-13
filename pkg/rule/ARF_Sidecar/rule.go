package ARF_Sidecar

import util "hsbc_connector/pkg/util"

import (
	"hsbc_connector/pkg/dto"
	"hsbc_connector/pkg/functions"
	"hsbc_connector/pkg/model"
	"hsbc_connector/pkg/rule/ARF_GetFundsConfirmation"
	"hsbc_connector/pkg/rule/ARF_GetInternationalPayment"
)

type CFGContext struct {
	Request      *dto.Request
	requestData  string
	requestJson  *model.Requeststring
	responseDTO  *dto.Response
	responseJson *model.Responsedata
	operation    string
	_context     *util.CustomContext
	_ruleConfig  map[string]interface{}
	_returnValue interface{}
	_errorValue  error
}

func NewCFGContext(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Request:     pRequest,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_Sidecar(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) (interface{}, error) {

	cfg := NewCFGContext(pRequest, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue, cfg._errorValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiHybridFunctionNodeM01()

	cfg.xiModelPropertyNodeM12()

	cfg.xiModelPropertyNodeM23()

	cfg.xiReferencePropertyNodeM34()

	cfg.xiConstantNodeM45()

	cfg.xiSwitchNodeM56()

	cfg.xiModelPropertyNodeM610()

	cfg.xiReferencePropertyNodeM1011()

	cfg.xiHybridFunctionNodeM1112()
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM1112() error {
	var err error
	cfg.responseDTO.Data, err = functions.ConvertModelToString(cfg.responseJson)
	cfg._returnValue = cfg.responseDTO.Data
	return err
}

func (cfg *CFGContext) xiReferencePropertyNodeM1011() error {

	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM610() error {
	cfg.responseDTO = &dto.Response{}

	return nil
}

func (cfg *CFGContext) xiSwitchNodeM56() error {

	switch _input := cfg.operation; {

	case _input == getIntPayment:
		cfg.branchM61()
	case _input == getDomPayment:
		cfg.branchM62()
	case _input == getFunConfrimation:
		cfg.branchM63()

	}
	return nil
}

const getIntPayment = "getIntPay"
const getDomPayment = "getDomPay"
const getFunConfrimation = "getFunCon"

func (cfg *CFGContext) xiConstantNodeM45() error {
	return nil
}

func (cfg *CFGContext) xiReferencePropertyNodeM34() error {

	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM23() error {
	cfg.responseJson = &model.Responsedata{}

	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM12() error {
	cfg.operation = cfg.requestJson.Opcode

	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.requestJson, err = functions.PaymentModelToString(cfg.requestData)
	cfg._returnValue = cfg.requestJson
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.requestData = cfg.Request.Data
	cfg.requestJson = &model.Requeststring{}

	return nil
}
func (cfg *CFGContext) branchM63() error {

	cfg.xiSubRuleNodeM9()
	return nil
}

func (cfg *CFGContext) xiSubRuleNodeM9() error {
	result, err := ARF_GetFundsConfirmation.ARF_GetFundsConfirmation(cfg.requestJson, cfg._context, cfg._ruleConfig)
	cfg.responseJson.Fundconresp = result.(string)
	cfg._returnValue = cfg.responseJson.Fundconresp
	cfg._errorValue = err
	return nil
}
func (cfg *CFGContext) branchM61() error {

	cfg.xiSubRuleNodeM7()
	return nil
}

func (cfg *CFGContext) xiSubRuleNodeM7() error {
	result, err := ARF_GetInternationalPayment.ARF_GetInternationalPayment(cfg.requestJson, cfg._context, cfg._ruleConfig)
	cfg.responseJson.Intpayresp = result.(string)
	cfg._returnValue = cfg.responseJson.Intpayresp
	cfg._errorValue = err
	return nil
}
func (cfg *CFGContext) branchM62() error {

	cfg.xiSubRuleNodeM8()
	return nil
}

func (cfg *CFGContext) xiSubRuleNodeM8() error {
	result, err := ARF_GetInternationalPayment.ARF_GetInternationalPayment(cfg.requestJson, cfg._context, cfg._ruleConfig)
	cfg.responseJson.Dompayresp = result.(string)
	cfg._returnValue = cfg.responseJson.Dompayresp
	cfg._errorValue = err
	return nil
}
