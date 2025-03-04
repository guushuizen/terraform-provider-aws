---
subcategory: "EBS (EC2)"
layout: "aws"
page_title: "AWS: aws_ebs_encryption_by_default"
description: |-
  Manages whether default EBS encryption is enabled for your AWS account in the current AWS region.
---

# Resource: aws_ebs_encryption_by_default

Provides a resource to manage whether default EBS encryption is enabled for your AWS account in the current AWS region. To manage the default KMS key for the region, see the [`aws_ebs_default_kms_key` resource](/docs/providers/aws/r/ebs_default_kms_key.html).

~> **NOTE:** Removing this Terraform resource disables default EBS encryption.

## Example Usage

```terraform
resource "aws_ebs_encryption_by_default" "example" {
  enabled = true
}
```

## Argument Reference

This resource supports the following arguments:

* `enabled` - (Optional) Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.

## Attribute Reference

This resource exports no additional attributes.

## Import

Import the default EBS encryption state. For example:

```
$ terraform import aws_ebs_encryption_by_default.example default
```
