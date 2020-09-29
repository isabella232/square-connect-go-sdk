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

// The query filter to return the search result whose named attribute values fall between the specified range.
type CatalogQueryRange struct {
	// The name of the attribute to be searched.
	AttributeName string `json:"attribute_name"`
	// The desired minimum value for the search attribute (inclusive).
	AttributeMinValue int64 `json:"attribute_min_value,omitempty"`
	// The desired maximum value for the search attribute (inclusive).
	AttributeMaxValue int64 `json:"attribute_max_value,omitempty"`
}