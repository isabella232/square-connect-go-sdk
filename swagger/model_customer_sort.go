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

// Specifies how searched customers profiles are sorted, including the sort key and sort order.
type CustomerSort struct {
	Field *CustomerSortField `json:"field,omitempty"`
	Order *SortOrder         `json:"order,omitempty"`
}
