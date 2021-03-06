/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// V1Fee
type V1Fee struct {
	// The fee's unique ID.
	Id string `json:"id,omitempty"`
	// The fee's name.
	Name string `json:"name,omitempty"`
	// The rate of the fee, as a string representation of a decimal number. A value of 0.07 corresponds to a rate of 7%.
	Rate             string                 `json:"rate,omitempty"`
	CalculationPhase *V1FeeCalculationPhase `json:"calculation_phase,omitempty"`
	AdjustmentType   *V1FeeAdjustmentType   `json:"adjustment_type,omitempty"`
	// If true, the fee applies to custom amounts entered into Square Point of Sale that are not associated with a particular item.
	AppliesToCustomAmounts bool `json:"applies_to_custom_amounts,omitempty"`
	// If true, the fee is applied to all appropriate items. If false, the fee is not applied at all.
	Enabled       bool                `json:"enabled,omitempty"`
	InclusionType *V1FeeInclusionType `json:"inclusion_type,omitempty"`
	Type_         *V1FeeType          `json:"type,omitempty"`
	// The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID.
	V2Id string `json:"v2_id,omitempty"`
}
