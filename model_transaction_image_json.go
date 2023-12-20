/*
 * OpenBankProject-FirehoseData
 *
 * An Open Source API for Banks. (c) TESOBE Ltd. 2011 - 2018. Licensed under the AGPL and commercial licences.
 *
 * API version: v3_1_0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TransactionImageJson struct {
	URL string `json:"URL,omitempty"`
	Label string `json:"label,omitempty"`
	Id string `json:"id,omitempty"`
	Date string `json:"date,omitempty"`
	UserJSONV121 *interface{} `json:"UserJSONV121,omitempty"`
	User * `json:"user,omitempty"`
}
