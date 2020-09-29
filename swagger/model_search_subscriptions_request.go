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

// Defines parameters in a [SearchSubscriptions](#endpoint-subscriptions-searchsubscriptions) endpoint  request.
type SearchSubscriptionsRequest struct {
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`
	// The upper limit on the number of subscriptions to return  in the response.   Default: `200`
	Limit int32                     `json:"limit,omitempty"`
	Query *SearchSubscriptionsQuery `json:"query,omitempty"`
}