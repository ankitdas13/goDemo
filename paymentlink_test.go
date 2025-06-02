// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package godemo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/godemo-go"
	"github.com/stainless-sdks/godemo-go/internal/testutil"
	"github.com/stainless-sdks/godemo-go/option"
)

func TestPaymentLinkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.New(context.TODO(), godemo.PaymentLinkNewParams{
		Amount:         1000,
		Currency:       "INR",
		AcceptPartial:  godemo.Bool(true),
		CallbackMethod: godemo.String("get"),
		CallbackURL:    godemo.String("https://example-callback-url.com/"),
		Customer: godemo.PaymentLinkNewParamsCustomer{
			Contact: godemo.String("+919000090000"),
			Email:   godemo.String("gaurav.kumar@example.com"),
			Name:    godemo.String("Gaurav Kumar"),
		},
		Description:           godemo.String("Payment for policy no"),
		ExpireBy:              godemo.Int(1691097057),
		FirstMinPartialAmount: godemo.Int(100),
		Notes: godemo.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
		Notify: godemo.PaymentLinkNewParamsNotify{
			Email: godemo.Bool(true),
			SMS:   godemo.Bool(true),
		},
		ReferenceID:    godemo.String("TS1989"),
		ReminderEnable: godemo.Bool(true),
		UpiLink:        godemo.Bool(true),
	})
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkGet(t *testing.T) {
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
	_, err := client.PaymentLinks.Get(context.TODO(), "id")
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.Update(
		context.TODO(),
		"id",
		godemo.PaymentLinkUpdateParams{
			AcceptPartial: godemo.Bool(false),
			ExpireBy:      godemo.Int(1653347540),
			Notes: godemo.NotesUnionParam{
				OfStringMap: map[string]string{
					"key1": "value3",
					"key2": "value2",
				},
			},
			ReferenceID: godemo.String("TS35"),
		},
	)
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkListWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.List(context.TODO(), godemo.PaymentLinkListParams{
		PaymentID:   godemo.String("payment_id"),
		ReferenceID: godemo.String("reference_id"),
	})
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkNotify(t *testing.T) {
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
	_, err := client.PaymentLinks.Notify(
		context.TODO(),
		godemo.PaymentLinkNotifyParamsMediumSMS,
		godemo.PaymentLinkNotifyParams{
			ID: "id",
		},
	)
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
