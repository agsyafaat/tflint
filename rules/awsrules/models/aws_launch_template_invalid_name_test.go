// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsLaunchTemplateInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_launch_template" "foo" {
	name = "foo[bar]"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsLaunchTemplateInvalidNameRule(),
					Message: `"foo[bar]" does not match valid pattern ^[a-zA-Z0-9\(\)\.\-/_]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_launch_template" "foo" {
	name = "foo"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsLaunchTemplateInvalidNameRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
