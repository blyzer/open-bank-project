/*
 * OpenBankProject-FirehoseData
 *
 * An Open Source API for Banks. (c) TESOBE Ltd. 2011 - 2018. Licensed under the AGPL and commercial licences.
 *
 * API version: v3_1_0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ModeratedCoreAccountJsonV300 struct {
	Number string `json:"number,omitempty"`
	AccountRoutings *Array `json:"account_routings,omitempty"`
	Label string `json:"label,omitempty"`
	Owners *Array `json:"owners,omitempty"`
	AmountOfMoneyJsonV121 *interface{} `json:"AmountOfMoneyJsonV121,omitempty"`
	Balance * `json:"balance,omitempty"`
	BankId string `json:"bank_id,omitempty"`
	Id string `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	AccountRules *Array `json:"account_rules,omitempty"`
}
