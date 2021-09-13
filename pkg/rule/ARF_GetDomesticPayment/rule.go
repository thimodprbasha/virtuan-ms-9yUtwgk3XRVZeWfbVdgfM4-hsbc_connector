package ARF_GetDomesticPayment

import util "hsbc_connector/pkg/util"

import (
	"hsbc_connector/pkg/functions"
	"hsbc_connector/pkg/model"
)

type CFGContext struct {
	Requeststring       *model.Requeststring
	domesticId          string
	authorization       string
	returnJson          string
	x_fapi_financial_id string
	_context            *util.CustomContext
	_ruleConfig         map[string]interface{}
	_returnValue        interface{}
	_errorValue         error
}

func NewCFGContext(pRequeststring *model.Requeststring, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Requeststring: pRequeststring,
		_context:      pContext,
		_ruleConfig:   pRuleConfig,
	}
}
func ARF_GetDomesticPayment(pRequeststring *model.Requeststring, pContext *util.CustomContext, pRuleConfig map[string]interface{}) (interface{}, error) {

	cfg := NewCFGContext(pRequeststring, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue, cfg._errorValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiHybridFunctionNodeM01()
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.returnJson, err = functions.GetDomesticPayment(cfg.domesticId, cfg.x_fapi_financial_id, cfg.authorization)
	cfg._returnValue = cfg.returnJson
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.domesticId = cfg.Requeststring.Domesticpaymentid
	cfg.authorization = cfg.Requeststring.Authorization
	cfg.x_fapi_financial_id = cfg.Requeststring.X_fapi_financial_id

	return nil
}
