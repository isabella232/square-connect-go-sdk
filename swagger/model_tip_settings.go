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

type TipSettings struct {
	// Indicates whether tipping is enabled for this checkout. Defaults to false.
	AllowTipping bool `json:"allow_tipping,omitempty"`
	// Indicates whether tip options should be presented on their own screen before presenting the signature screen during card payment. Defaults to false.
	SeparateTipScreen bool `json:"separate_tip_screen,omitempty"`
	// Indicates whether custom tip amounts are allowed during the checkout flow. Defaults to false.
	CustomTipField bool `json:"custom_tip_field,omitempty"`
}
