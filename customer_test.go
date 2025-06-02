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

func TestCustomerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.New(context.TODO(), godemo.CustomerNewParams{
		Contact:      "9123456780",
		Email:        "gaurav.kumar@example.com",
		Name:         "Gaurav Kumar",
		FailExisting: godemo.Bool(false),
		Gstin:        godemo.String("29XAbbA4369J1PA"),
		Notes: godemo.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
	})
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerGet(t *testing.T) {
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
	_, err := client.Customers.Get(context.TODO(), "id")
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Update(
		context.TODO(),
		"id",
		godemo.CustomerUpdateParams{
			Contact: godemo.String("9000000000"),
			Email:   godemo.String("gaurav.kumar@example.com"),
			Name:    godemo.String("Gaurav Kumar"),
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

func TestCustomerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.List(context.TODO(), godemo.CustomerListParams{
		Count: godemo.Int(1),
		Skip:  godemo.Int(0),
	})
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
