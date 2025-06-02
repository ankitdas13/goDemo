// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package godemo

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/godemo-go/internal/apijson"
	"github.com/stainless-sdks/godemo-go/internal/requestconfig"
	"github.com/stainless-sdks/godemo-go/option"
	"github.com/stainless-sdks/godemo-go/packages/param"
	"github.com/stainless-sdks/godemo-go/packages/respjson"
)

// InvoiceService contains methods and other services that help with interacting
// with the godemo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceService] method instead.
type InvoiceService struct {
	Options []option.RequestOption
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r InvoiceService) {
	r = InvoiceService{}
	r.Options = opts
	return
}

// Use this endpoint to create an invoice by passing the customer_id.
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *InvoiceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type InvoiceNewResponse struct {
	// The unique identifier of the invoice.
	ID string `json:"id"`
	// Amount to be paid using the invoice. Must be in the smallest unit of the
	// currency. For example, if the amount to be received from the customer is
	// ₹300.00, pass the value as 30000.
	Amount int64 `json:"amount"`
	// The remaining amount to be paid by the customer for the issued invoice.
	AmountDue int64 `json:"amount_due"`
	// Amount paid by the customer against the invoice.
	AmountPaid int64 `json:"amount_paid"`
	// Billing end timestamp.
	BillingEnd int64 `json:"billing_end,nullable"`
	// Billing start timestamp.
	BillingStart int64 `json:"billing_start,nullable"`
	// Timestamp, in Unix format, at which the invoice was cancelled.
	CancelledAt int64 `json:"cancelled_at,nullable"`
	// Any comments to be added in the invoice. Maximum of 2048 characters.
	Comment string `json:"comment,nullable"`
	// Timestamp, in Unix format, when the invoice was created.
	CreatedAt int64 `json:"created_at"`
	// The currency associated with the invoice. You must mandatorily pass this
	// parameter if accepting international payments.
	Currency string `json:"currency"`
	// Symbol of the currency.
	CurrencySymbol string `json:"currency_symbol"`
	// Details of the customer.
	CustomerDetails InvoiceNewResponseCustomerDetails `json:"customer_details"`
	// The unique identifier of the customer.
	CustomerID string `json:"customer_id"`
	// Timestamp, in Unix format, that indicates the issue date of the invoice.
	Date int64 `json:"date"`
	// A brief description of the invoice. The maximum character length is 2048.
	Description string `json:"description,nullable"`
	// The delivery status of the email notification for the invoice sent to the
	// customer. Possible values are pending, sent.
	//
	// Any of "pending", "sent".
	EmailStatus InvoiceNewResponseEmailStatus `json:"email_status"`
	// Indicates the type of entity. Here, it is invoice.
	Entity string `json:"entity"`
	// Timestamp, in Unix format, at which the invoice will expire.
	ExpireBy int64 `json:"expire_by,nullable"`
	// Timestamp, in Unix format, at which the invoice expired.
	ExpiredAt int64 `json:"expired_at,nullable"`
	// Gross amount of the invoice in the smallest currency unit.
	GrossAmount int64 `json:"gross_amount"`
	// Whether to group taxes and discounts.
	GroupTaxesDiscounts bool `json:"group_taxes_discounts"`
	// Unique number you added for internal reference. The minimum character length is
	// 1 and maximum is 40.
	InvoiceNumber string `json:"invoice_number,nullable"`
	// Timestamp, in Unix format, at which the invoice was issued to the customer.
	IssuedAt int64 `json:"issued_at"`
	// Details of the line item that is billed in the invoice. Maximum of 50 line
	// items.
	LineItems []InvoiceNewResponseLineItem `json:"line_items"`
	// Any custom notes added to the invoice. Maximum of 2048 characters.
	Notes []string `json:"notes"`
	// The unique identifier of the order associated with the invoice.
	OrderID string `json:"order_id"`
	// Timestamp, in Unix format, at which the payment was made.
	PaidAt int64 `json:"paid_at,nullable"`
	// Indicates whether the customer can make a partial payment on the invoice.
	// Possible values true (The customer can make partial payments) or false (default,
	// The customer cannot make partial payments).
	PartialPayment bool `json:"partial_payment"`
	// Unique identifier of a payment made against this invoice.
	PaymentID string `json:"payment_id,nullable"`
	// Receipt number that corresponds to this invoice.
	Receipt string `json:"receipt,nullable"`
	// The short URL that is generated. Share this link with customers to accept
	// payments.
	ShortURL string `json:"short_url"`
	// The delivery status of the SMS notification for the invoice sent to the
	// customer. Possible values are pending, sent.
	//
	// Any of "pending", "sent".
	SMSStatus InvoiceNewResponseSMSStatus `json:"sms_status"`
	// The status of the invoice. Possible values are draft, issued, partially_paid,
	// paid, cancelled, expired, deleted.
	//
	// Any of "draft", "issued", "partially_paid", "paid", "cancelled", "expired",
	// "deleted".
	Status InvoiceNewResponseStatus `json:"status"`
	// Tax amount of the invoice in the smallest currency unit.
	TaxAmount int64 `json:"tax_amount"`
	// Taxable amount of the invoice in the smallest currency unit.
	TaxableAmount int64 `json:"taxable_amount"`
	// Any terms to be included in the invoice. Maximum of 2048 characters.
	Terms string `json:"terms,nullable"`
	// Type of the invoice.
	Type string `json:"type"`
	// Whether to show less details in the invoice.
	ViewLess bool `json:"view_less"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Amount              respjson.Field
		AmountDue           respjson.Field
		AmountPaid          respjson.Field
		BillingEnd          respjson.Field
		BillingStart        respjson.Field
		CancelledAt         respjson.Field
		Comment             respjson.Field
		CreatedAt           respjson.Field
		Currency            respjson.Field
		CurrencySymbol      respjson.Field
		CustomerDetails     respjson.Field
		CustomerID          respjson.Field
		Date                respjson.Field
		Description         respjson.Field
		EmailStatus         respjson.Field
		Entity              respjson.Field
		ExpireBy            respjson.Field
		ExpiredAt           respjson.Field
		GrossAmount         respjson.Field
		GroupTaxesDiscounts respjson.Field
		InvoiceNumber       respjson.Field
		IssuedAt            respjson.Field
		LineItems           respjson.Field
		Notes               respjson.Field
		OrderID             respjson.Field
		PaidAt              respjson.Field
		PartialPayment      respjson.Field
		PaymentID           respjson.Field
		Receipt             respjson.Field
		ShortURL            respjson.Field
		SMSStatus           respjson.Field
		Status              respjson.Field
		TaxAmount           respjson.Field
		TaxableAmount       respjson.Field
		Terms               respjson.Field
		Type                respjson.Field
		ViewLess            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the customer.
type InvoiceNewResponseCustomerDetails struct {
	// The unique identifier of the customer.
	ID string `json:"id"`
	// Billing address of the customer.
	BillingAddress InvoiceNewResponseCustomerDetailsBillingAddress `json:"billing_address,nullable"`
	// Contact number of the customer.
	Contact string `json:"contact"`
	// Customer contact number.
	CustomerContact string `json:"customer_contact"`
	// Customer email address.
	CustomerEmail string `json:"customer_email"`
	// Customer name.
	CustomerName string `json:"customer_name,nullable"`
	// Email address of the customer.
	Email string `json:"email"`
	// GSTIN of the customer.
	Gstin string `json:"gstin,nullable"`
	// Name of the customer.
	Name string `json:"name,nullable"`
	// Shipping address of the customer.
	ShippingAddress InvoiceNewResponseCustomerDetailsShippingAddress `json:"shipping_address,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		BillingAddress  respjson.Field
		Contact         respjson.Field
		CustomerContact respjson.Field
		CustomerEmail   respjson.Field
		CustomerName    respjson.Field
		Email           respjson.Field
		Gstin           respjson.Field
		Name            respjson.Field
		ShippingAddress respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceNewResponseCustomerDetails) RawJSON() string { return r.JSON.raw }
func (r *InvoiceNewResponseCustomerDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing address of the customer.
type InvoiceNewResponseCustomerDetailsBillingAddress struct {
	// Unique identifier of the address.
	ID string `json:"id"`
	// City of the address.
	City string `json:"city"`
	// Country code of the address.
	Country string `json:"country"`
	// First line of the address.
	Line1 string `json:"line1"`
	// Second line of the address.
	Line2 string `json:"line2"`
	// Whether this is the primary address.
	Primary bool `json:"primary"`
	// State of the address.
	State string `json:"state"`
	// Type of the address.
	Type string `json:"type"`
	// ZIP code of the address.
	Zipcode string `json:"zipcode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		Line2       respjson.Field
		Primary     respjson.Field
		State       respjson.Field
		Type        respjson.Field
		Zipcode     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceNewResponseCustomerDetailsBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *InvoiceNewResponseCustomerDetailsBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shipping address of the customer.
type InvoiceNewResponseCustomerDetailsShippingAddress struct {
	// Unique identifier of the address.
	ID string `json:"id"`
	// City of the address.
	City string `json:"city"`
	// Country code of the address.
	Country string `json:"country"`
	// First line of the address.
	Line1 string `json:"line1"`
	// Second line of the address.
	Line2 string `json:"line2"`
	// Whether this is the primary address.
	Primary bool `json:"primary"`
	// State of the address.
	State string `json:"state"`
	// Type of the address.
	Type string `json:"type"`
	// ZIP code of the address.
	Zipcode string `json:"zipcode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		Line2       respjson.Field
		Primary     respjson.Field
		State       respjson.Field
		Type        respjson.Field
		Zipcode     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceNewResponseCustomerDetailsShippingAddress) RawJSON() string { return r.JSON.raw }
func (r *InvoiceNewResponseCustomerDetailsShippingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The delivery status of the email notification for the invoice sent to the
// customer. Possible values are pending, sent.
type InvoiceNewResponseEmailStatus string

const (
	InvoiceNewResponseEmailStatusPending InvoiceNewResponseEmailStatus = "pending"
	InvoiceNewResponseEmailStatusSent    InvoiceNewResponseEmailStatus = "sent"
)

type InvoiceNewResponseLineItem struct {
	// Unique identifier of the line item.
	ID string `json:"id"`
	// Amount of the item in the smallest currency unit.
	Amount int64 `json:"amount"`
	// Currency of the item.
	Currency string `json:"currency"`
	// Description of the item.
	Description string `json:"description"`
	// Gross amount of the item in the smallest currency unit.
	GrossAmount int64 `json:"gross_amount"`
	// HSN code of the item.
	HsnCode string `json:"hsn_code,nullable"`
	// Unique identifier of the item.
	ItemID string `json:"item_id,nullable"`
	// Name of the item.
	Name string `json:"name"`
	// Net amount of the item in the smallest currency unit.
	NetAmount int64 `json:"net_amount"`
	// Quantity of the item.
	Quantity int64 `json:"quantity"`
	// Reference identifier.
	RefID string `json:"ref_id,nullable"`
	// Reference type.
	RefType string `json:"ref_type,nullable"`
	// SAC code of the item.
	SacCode string `json:"sac_code,nullable"`
	// Tax amount of the item in the smallest currency unit.
	TaxAmount int64 `json:"tax_amount"`
	// Whether the amount is tax inclusive.
	TaxInclusive bool `json:"tax_inclusive"`
	// Tax rate of the item.
	TaxRate string `json:"tax_rate,nullable"`
	// Taxable amount of the item in the smallest currency unit.
	TaxableAmount int64 `json:"taxable_amount"`
	// Tax details of the item.
	Taxes []any `json:"taxes"`
	// Type of the item.
	Type string `json:"type"`
	// Unit of the item.
	Unit string `json:"unit,nullable"`
	// Unit amount of the item in the smallest currency unit.
	UnitAmount int64 `json:"unit_amount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		Currency      respjson.Field
		Description   respjson.Field
		GrossAmount   respjson.Field
		HsnCode       respjson.Field
		ItemID        respjson.Field
		Name          respjson.Field
		NetAmount     respjson.Field
		Quantity      respjson.Field
		RefID         respjson.Field
		RefType       respjson.Field
		SacCode       respjson.Field
		TaxAmount     respjson.Field
		TaxInclusive  respjson.Field
		TaxRate       respjson.Field
		TaxableAmount respjson.Field
		Taxes         respjson.Field
		Type          respjson.Field
		Unit          respjson.Field
		UnitAmount    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceNewResponseLineItem) RawJSON() string { return r.JSON.raw }
func (r *InvoiceNewResponseLineItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The delivery status of the SMS notification for the invoice sent to the
// customer. Possible values are pending, sent.
type InvoiceNewResponseSMSStatus string

const (
	InvoiceNewResponseSMSStatusPending InvoiceNewResponseSMSStatus = "pending"
	InvoiceNewResponseSMSStatusSent    InvoiceNewResponseSMSStatus = "sent"
)

// The status of the invoice. Possible values are draft, issued, partially_paid,
// paid, cancelled, expired, deleted.
type InvoiceNewResponseStatus string

const (
	InvoiceNewResponseStatusDraft         InvoiceNewResponseStatus = "draft"
	InvoiceNewResponseStatusIssued        InvoiceNewResponseStatus = "issued"
	InvoiceNewResponseStatusPartiallyPaid InvoiceNewResponseStatus = "partially_paid"
	InvoiceNewResponseStatusPaid          InvoiceNewResponseStatus = "paid"
	InvoiceNewResponseStatusCancelled     InvoiceNewResponseStatus = "cancelled"
	InvoiceNewResponseStatusExpired       InvoiceNewResponseStatus = "expired"
	InvoiceNewResponseStatusDeleted       InvoiceNewResponseStatus = "deleted"
)

type InvoiceNewParams struct {
	// You can pass the customer_id in this field, if you are using the Customers API.
	// If not, you can pass the customer object described in the below fields.
	CustomerID string `json:"customer_id,required"`
	// Details of the line item that is billed in the invoice. Maximum of 50 line
	// items.
	LineItems []InvoiceNewParamsLineItem `json:"line_items,omitzero,required"`
	// Indicates the type of entity. Here, it is invoice.
	//
	// Any of "invoice".
	Type InvoiceNewParamsType `json:"type,omitzero,required"`
	// The currency associated with the invoice. You must mandatorily pass this
	// parameter if accepting international payments. If you have passed currency as a
	// sub-parameter in the line_item object, you must ensure that the same currency is
	// passed in both places.
	Currency param.Opt[string] `json:"currency,omitzero"`
	// A brief description of the invoice.
	Description param.Opt[string] `json:"description,omitzero"`
	// Invoice is created in draft state when value is set to 1.
	Draft param.Opt[string] `json:"draft,omitzero"`
	// Defines who handles the email notification. Possible values false (You send the
	// notification to the customer) or true (default, Razorpay sends the notification
	// to the customer).
	EmailNotify param.Opt[bool] `json:"email_notify,omitzero"`
	// Timestamp, in Unix format, at which the invoice will expire.
	ExpireBy param.Opt[int64] `json:"expire_by,omitzero"`
	// Any custom notes added to the invoice. Maximum of 2048 characters.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Indicates whether the customer can make a partial payment on the invoice.
	// Possible values true (The customer can make partial payments) or false (default,
	// The customer cannot make partial payments).
	PartialPayment param.Opt[bool] `json:"partial_payment,omitzero"`
	// Defines who handles the SMS notification. Possible values false (You send the
	// notification to the customer) or true (default, Razorpay sends the notification
	// to the customer).
	SMSNotify param.Opt[bool] `json:"sms_notify,omitzero"`
	// Customer details.
	Customer InvoiceNewParamsCustomer `json:"customer,omitzero"`
	paramObj
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceNewParamsLineItem struct {
	// Amount of the item in the smallest currency unit.
	Amount param.Opt[int64] `json:"amount,omitzero"`
	// Currency of the item.
	Currency param.Opt[string] `json:"currency,omitzero"`
	// Description of the item.
	Description param.Opt[string] `json:"description,omitzero"`
	// Unique identifier of the item.
	ItemID param.Opt[string] `json:"item_id,omitzero"`
	// Name of the item.
	Name param.Opt[string] `json:"name,omitzero"`
	// Quantity of the item.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	paramObj
}

func (r InvoiceNewParamsLineItem) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParamsLineItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParamsLineItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the type of entity. Here, it is invoice.
type InvoiceNewParamsType string

const (
	InvoiceNewParamsTypeInvoice InvoiceNewParamsType = "invoice"
)

// Customer details.
type InvoiceNewParamsCustomer struct {
	// Contact number of the customer.
	Contact param.Opt[string] `json:"contact,omitzero"`
	// Email address of the customer.
	Email param.Opt[string] `json:"email,omitzero"`
	// GSTIN of the customer.
	Gstin param.Opt[string] `json:"gstin,omitzero"`
	// Name of the customer.
	Name param.Opt[string] `json:"name,omitzero"`
	// Billing address of the customer.
	BillingAddress InvoiceNewParamsCustomerBillingAddress `json:"billing_address,omitzero"`
	paramObj
}

func (r InvoiceNewParamsCustomer) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParamsCustomer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParamsCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing address of the customer.
type InvoiceNewParamsCustomerBillingAddress struct {
	// City of the address.
	City param.Opt[string] `json:"city,omitzero"`
	// Country code of the address.
	Country param.Opt[string] `json:"country,omitzero"`
	// First line of the address.
	Line1 param.Opt[string] `json:"line1,omitzero"`
	// Second line of the address.
	Line2 param.Opt[string] `json:"line2,omitzero"`
	// State of the address.
	State param.Opt[string] `json:"state,omitzero"`
	// ZIP code of the address.
	Zipcode param.Opt[string] `json:"zipcode,omitzero"`
	paramObj
}

func (r InvoiceNewParamsCustomerBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParamsCustomerBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParamsCustomerBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
