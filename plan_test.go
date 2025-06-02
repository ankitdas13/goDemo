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
	"github.com/stainless-sdks/godemo-go/packages/param"
)

func TestPlanNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Plans.New(context.TODO(), godemo.PlanNewParams{
		Interval: 1,
		Item: godemo.PlanItemParam{
			ID:           godemo.String("item_00000000000001"),
			Active:       godemo.Bool(true),
			Amount:       godemo.Int(69900),
			CreatedAt:    godemo.Int(1580219935),
			Currency:     godemo.String("INR"),
			Description:  godemo.String("Description for the test plan - Weekly"),
			HsnCode:      param.Null[string](),
			Name:         godemo.String("Test plan - Weekly"),
			SacCode:      param.Null[string](),
			TaxGroupID:   param.Null[string](),
			TaxID:        param.Null[string](),
			TaxInclusive: godemo.Bool(false),
			TaxRate:      param.Null[float64](),
			Type:         godemo.String("plan"),
			Unit:         param.Null[string](),
			UnitAmount:   godemo.Int(69900),
			UpdatedAt:    godemo.Int(1580219935),
		},
		Period: godemo.PlanNewParamsPeriodWeekly,
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

func TestPlanGet(t *testing.T) {
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
	_, err := client.Plans.Get(context.TODO(), "id")
	if err != nil {
		var apierr *godemo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
