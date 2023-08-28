package helpers

import (
	"encoding/json"
	"fireteam-core/internal/testhelper"
	"testing"
)

func TestValidateJsonBody(t *testing.T) {
	testCasesInJson, err := testhelper.LoadJsonFromFile("./testcases/validate_json_tests.json")
	if err != nil {
		t.Fatalf("Failed to parse test cases JSON: %v", err)
	}
	var testCases []struct {
		Name       string `json:"name"`
		Body       string `json:"body"`
		ShouldPass bool   `json:"should_pass"`
	}
	err = json.Unmarshal(testCasesInJson, &testCases)
	if err != nil {
		t.Fatalf("Failed to parse test cases JSON: %v", err)
	}
	if err != nil {
		t.Fatalf("Failed to parse test cases JSON: %v", err)
	}
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			json, err := ValidateJsonBody(tt.Body)
			if tt.ShouldPass {
				if err != nil {
					t.Errorf("validate should not error, but got error: %v", err)
					return
				}
				if json == nil {
					t.Errorf("validate should return json, but got nil")
					return
				}
			} else {
				if err == nil {
					t.Errorf("validate should error, but got non-nil error")
					return
				}
				if json != nil {
					t.Errorf("validate should not return json, but got json: %v", json)
					return
				}
			}
		})
	}
}
