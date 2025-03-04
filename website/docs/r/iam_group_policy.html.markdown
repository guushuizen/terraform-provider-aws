---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "AWS: aws_iam_group_policy"
description: |-
  Provides an IAM policy attached to a group.
---

# Resource: aws_iam_group_policy

Provides an IAM policy attached to a group.

## Example Usage

```terraform
resource "aws_iam_group_policy" "my_developer_policy" {
  name  = "my_developer_policy"
  group = aws_iam_group.my_developers.name

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:Describe*",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "aws_iam_group" "my_developers" {
  name = "developers"
  path = "/users/"
}
```

## Argument Reference

This resource supports the following arguments:

* `policy` - (Required) The policy document. This is a JSON formatted string. For more information about building IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://learn.hashicorp.com/terraform/aws/iam-policy)
* `name` - (Optional) The name of the policy. If omitted, Terraform will
assign a random, unique name.
* `name_prefix` - (Optional) Creates a unique name beginning with the specified
  prefix. Conflicts with `name`.
* `group` - (Required) The IAM group to attach to the policy.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The group policy ID.
* `group` - The group to which this policy applies.
* `name` - The name of the policy.
* `policy` - The policy document attached to the group.

## Import

Import IAM Group Policies using the `group_name:group_policy_name`. For example:

```
$ terraform import aws_iam_group_policy.mypolicy group_of_mypolicy_name:mypolicy_name
```
