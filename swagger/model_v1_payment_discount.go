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

// V1PaymentDiscount
type V1PaymentDiscount struct {
	// The discount's name.
	Name         string   `json:"name,omitempty"`
	AppliedMoney *V1Money `json:"applied_money,omitempty"`
	// The ID of the applied discount, if available. Discounts applied in older versions of Square Register might not have an ID.
	DiscountId string `json:"discount_id,omitempty"`
}
