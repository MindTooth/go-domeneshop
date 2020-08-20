/*
 * Domeneshop API Documentation
 *
 * # Overview Domeneshop offers a simple, REST-based API, which currently supports the following  features: - List domains - List invoices - Create, read, update and delete DNS records for domains - Create, read, update and delete HTTP forwards (\"WWW forwarding\") for domains - Dynamic DNS (DDNS) update endpoints for use in consumer routers  More features are planned, including: - Web hosting administration - Email address and email user/account administration  # Authentication The Domeneshop API currently supports only one method of authentication, **HTTP Basic Auth**. More authentication methods may be added in the future.  To generate credentials, visit <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">this page</a> after logging in to the control panel on our website: <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">https://www.domeneshop.no/admin?view=api</a> 
 *
 * API version: v0
 * Contact: kundeservice@domeneshop.no
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package domeneshop
// Invoice struct for Invoice
type Invoice struct {
	// Invoice ID/number
	Id int32 `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Amount int32 `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	// The invoice due date. Only available for type `invoice`.
	DueDate string `json:"due_date,omitempty"`
	// The date when the invoice was issued.
	IssuedDate string `json:"issued_date,omitempty"`
	// The payment date. Only available if the invoice has status `paid`.
	PaidDate string `json:"paid_date,omitempty"`
	// `settled` is only applicable to credit notes. These are usually created if  domains have been 
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
}