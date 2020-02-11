// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsBudgetsBudgetInvalidTimeUnitRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_budgets_budget" "foo" {
	time_unit = "HOURLY"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsBudgetsBudgetInvalidTimeUnitRule(),
					Message: `"HOURLY" is an invalid value as time_unit`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_budgets_budget" "foo" {
	time_unit = "MONTHLY"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsBudgetsBudgetInvalidTimeUnitRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
