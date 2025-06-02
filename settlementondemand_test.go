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

func TestSettlementOndemandNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Settlements.Ondemand.New(context.TODO(), godemo.SettlementOndemandNewParams{
		Amount:      200000,
		Description: godemo.String("Need this to make vendor payments."),
		Notes: godemo.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
		SettleFullBalance: godemo.Bool(false),
	})
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettlementOndemandGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Settlements.Ondemand.Get(
		context.TODO(),
		"id",
		godemo.SettlementOndemandGetParams{
			Expand: []string{"ondemand_payouts"},
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
