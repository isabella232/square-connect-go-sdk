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

// InvoiceRequestMethod : Specifies the action for Square to take for processing the invoice. For example,  email the invoice, charge a customer's card on file, or do nothing.
type InvoiceRequestMethod string

// List of InvoiceRequestMethod
const (
	EMAIL_InvoiceRequestMethod               InvoiceRequestMethod = "EMAIL"
	CHARGE_CARD_ON_FILE_InvoiceRequestMethod InvoiceRequestMethod = "CHARGE_CARD_ON_FILE"
	SHARE_MANUALLY_InvoiceRequestMethod      InvoiceRequestMethod = "SHARE_MANUALLY"
)
