// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package godemo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/ankitdas13/goDemo"
	"github.com/ankitdas13/goDemo/internal/testutil"
	"github.com/ankitdas13/goDemo/option"
)

func TestInvoiceNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := godemo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Invoices.New(context.TODO(), godemo.InvoiceNewParams{
		CustomerID: "cust_HOQzpsovChhcpl",
		LineItems: []godemo.InvoiceNewParamsLineItem{{
			Amount:      godemo.Int(1200),
			Currency:    godemo.String("INR"),
			Description: godemo.String("Cotton Cloth"),
			ItemID:      godemo.String("item_K6g5L6X43dXjEA"),
			Name:        godemo.String("Cloth"),
			Quantity:    godemo.Int(1),
		}},
		Type:     godemo.InvoiceNewParamsTypeInvoice,
		Currency: godemo.String("INR"),
		Customer: godemo.InvoiceNewParamsCustomer{
			BillingAddress: godemo.InvoiceNewParamsCustomerBillingAddress{
				City:    godemo.String("Bengaluru"),
				Country: godemo.String("in"),
				Line1:   godemo.String("Ground & 1st Floor, SJR Cyber Laskar"),
				Line2:   godemo.String("Hosur Road"),
				State:   godemo.String("Karnataka"),
				Zipcode: godemo.String("560068"),
			},
			Contact: godemo.String("+919111111111"),
			Email:   godemo.String("gaurav.kumar@example.com"),
			Gstin:   godemo.String("29ABCDE1234L1Z5"),
			Name:    godemo.String("Gaurav Kumar"),
		},
		Description:    godemo.String("Invoice for services rendered"),
		Draft:          godemo.String("1"),
		EmailNotify:    godemo.Bool(true),
		ExpireBy:       godemo.Int(1760714528),
		Notes:          godemo.String("Additional notes for the invoice"),
		PartialPayment: godemo.Bool(false),
		SMSNotify:      godemo.Bool(true),
	})
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
