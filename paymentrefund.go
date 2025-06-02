// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package godemo

import (
	"github.com/stainless-sdks/godemo-go/internal/apijson"
	"github.com/stainless-sdks/godemo-go/option"
	"github.com/stainless-sdks/godemo-go/packages/respjson"
)

// PaymentRefundService contains methods and other services that help with
// interacting with the godemo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentRefundService] method instead.
type PaymentRefundService struct {
	Options []option.RequestOption
}

// NewPaymentRefundService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentRefundService(opts ...option.RequestOption) (r PaymentRefundService) {
	r = PaymentRefundService{}
	r.Options = opts
	return
}

type RefundList struct {
	// Number of refunds returned.
	Count int64 `json:"count"`
	// Name of the entity. Always 'collection'.
	Entity string `json:"entity"`
	// List of refund objects.
	Items []Refund `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Entity      respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RefundList) RawJSON() string { return r.JSON.raw }
func (r *RefundList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
