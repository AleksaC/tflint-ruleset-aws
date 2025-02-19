// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloud9EnvironmentEc2InvalidInstanceTypeRule checks the pattern is valid
type AwsCloud9EnvironmentEc2InvalidInstanceTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloud9EnvironmentEc2InvalidInstanceTypeRule returns new rule with default attributes
func NewAwsCloud9EnvironmentEc2InvalidInstanceTypeRule() *AwsCloud9EnvironmentEc2InvalidInstanceTypeRule {
	return &AwsCloud9EnvironmentEc2InvalidInstanceTypeRule{
		resourceType:  "aws_cloud9_environment_ec2",
		attributeName: "instance_type",
		max:           20,
		min:           5,
		pattern:       regexp.MustCompile(`^[a-z][1-9][.][a-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloud9EnvironmentEc2InvalidInstanceTypeRule) Name() string {
	return "aws_cloud9_environment_ec2_invalid_instance_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloud9EnvironmentEc2InvalidInstanceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloud9EnvironmentEc2InvalidInstanceTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloud9EnvironmentEc2InvalidInstanceTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloud9EnvironmentEc2InvalidInstanceTypeRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"instance_type must be 20 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"instance_type must be 5 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z][1-9][.][a-z0-9]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
