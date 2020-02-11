// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ec2_transit_gateway" "foo" {
	auto_accept_shared_attachments = "true"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule(),
					Message: `"true" is an invalid value as auto_accept_shared_attachments`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ec2_transit_gateway" "foo" {
	auto_accept_shared_attachments = "enable"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
