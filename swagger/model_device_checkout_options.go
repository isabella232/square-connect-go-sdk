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

type DeviceCheckoutOptions struct {
	// The unique Id of the device intended for this `TerminalCheckout`. The Id can be retrieved from /v2/devices api.
	DeviceId string `json:"device_id"`
	// Instruct the device to skip the receipt screen. Defaults to false.
	SkipReceiptScreen bool         `json:"skip_receipt_screen,omitempty"`
	TipSettings       *TipSettings `json:"tip_settings,omitempty"`
}