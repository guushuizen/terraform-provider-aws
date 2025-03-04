---
subcategory: "RDS (Relational Database)"
layout: "aws"
page_title: "AWS: aws_db_subnet_group"
description: |-
  Provides an RDS DB subnet group resource.
---

# Resource: aws_db_subnet_group

Provides an RDS DB subnet group resource.

> **Hands-on:** For an example of the `aws_db_subnet_group` in use, follow the [Manage AWS RDS Instances](https://learn.hashicorp.com/tutorials/terraform/aws-rds?in=terraform/aws&utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) tutorial on HashiCorp Learn.

## Example Usage

```terraform
resource "aws_db_subnet_group" "default" {
  name       = "main"
  subnet_ids = [aws_subnet.frontend.id, aws_subnet.backend.id]

  tags = {
    Name = "My DB subnet group"
  }
}
```

## Argument Reference

This resource supports the following arguments:

* `name` - (Optional, Forces new resource) The name of the DB subnet group. If omitted, Terraform will assign a random, unique name.
* `name_prefix` - (Optional, Forces new resource) Creates a unique name beginning with the specified prefix. Conflicts with `name`.
* `description` - (Optional) The description of the DB subnet group. Defaults to "Managed by Terraform".
* `subnet_ids` - (Required) A list of VPC subnet IDs.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The db subnet group name.
* `arn` - The ARN of the db subnet group.
* `supported_network_types` - The network type of the db subnet group.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).
* `vpc_id` - Provides the VPC ID of the DB subnet group.

## Import

Import DB Subnet groups using the `name`. For example:

```
$ terraform import aws_db_subnet_group.default production-subnet-group
```
