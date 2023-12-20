/*
 * OpenBankProject-API-Metric
 *
 * An Open Source API for Banks. (c) TESOBE Ltd. 2011 - 2018. Licensed under the AGPL and commercial licences.
 *
 * API version: v3_1_0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AggregateMetricJson struct {
	MaximumResponseTime float64 `json:"maximum_response_time,omitempty"`
	MinimumResponseTime float64 `json:"minimum_response_time,omitempty"`
	AverageResponseTime float64 `json:"average_response_time,omitempty"`
	Count int32 `json:"count,omitempty"`
}
