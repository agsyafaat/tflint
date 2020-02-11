// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCloudwatchEventTargetInvalidTargetIDRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudwatch_event_target" "foo" {
	target_id = "run scheduled task every hour"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudwatchEventTargetInvalidTargetIDRule(),
					Message: `"run scheduled task every hour" does not match valid pattern ^[\.\-_A-Za-z0-9]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudwatch_event_target" "foo" {
	target_id = "run-scheduled-task-every-hour"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudwatchEventTargetInvalidTargetIDRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
