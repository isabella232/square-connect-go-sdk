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

// A modifier applicable to items at the time of sale.
type CatalogModifier struct {
	// The modifier name.  This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points.
	Name       string `json:"name,omitempty"`
	PriceMoney *Money `json:"price_money,omitempty"`
	// Determines where this `CatalogModifier` appears in the `CatalogModifierList`.
	Ordinal int32 `json:"ordinal,omitempty"`
	// The ID of the `CatalogModifierList` associated with this modifier.
	ModifierListId string `json:"modifier_list_id,omitempty"`
}