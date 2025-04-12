package bdd

import (
	"fmt"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/cucumber/godog"
	"resty.dev/v3"
)

type apiFeature struct {
	client *resty.Client
	resp   *resty.Response
}

func NewApiFeature() *apiFeature {
	client := resty.New()
	client.SetBaseURL("http://localhost:9999")
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Accept", "application/json")
	return &apiFeature{
		client: client,
	}
}

// // func (a *apiFeature) resetResponse(*godog.Scenario) {
// // 	a.resp = httptest.NewRecorder()
// // }

func (a *apiFeature) iSendrequestTo(method, endpoint string) (err error) {
	resp, err := a.client.R().Execute(method, endpoint)
	if err != nil {
		return fmt.Errorf("failed to send %s request to %s: %w", method, endpoint, err)
	}
	a.resp = resp
	return nil
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.StatusCode() {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.StatusCode())
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJSON(body *godog.DocString) (err error) {
	var expected, actual any

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	fmt.Println("expected:", expected)

	// re-encode actual response too
	if err = json.Unmarshal(a.resp.Bytes(), &actual); err != nil {
		return
	}

	fmt.Println("actual:", actual)

	// // the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
	}
	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := NewApiFeature()

	// ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	// 	api.resetResponse(sc)
	// 	return ctx, nil
	// })
	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, api.iSendrequestTo)
	ctx.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, api.theResponseShouldMatchJSON)
}
