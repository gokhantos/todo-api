package main

import (
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {
	go server("test")
	pact := &dsl.Pact{
		Provider:                 "Todo App",
		DisableToolValidityCheck: true,
	}

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            "http://localhost:8081",
		BrokerURL:                  "https://gokhantosun.pactflow.io",
		BrokerToken:                "4bo_wjLIsBRkTSJ4YqiBlw",
		PublishVerificationResults: true,
		PactURLs:                   []string{"https://gokhantosun.pactflow.io/pacts/provider/Todo%20Api/consumer/Todo%20Client/latest/main"},
		ProviderVersion:            "2.0.0",
	})

	if err != nil {
		t.Fatal(err)
	}

}
