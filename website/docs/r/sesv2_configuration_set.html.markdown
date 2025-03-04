---
subcategory: "SESv2 (Simple Email V2)"
layout: "aws"
page_title: "AWS: aws_sesv2_configuration_set"
description: |-
  Terraform resource for managing an AWS SESv2 (Simple Email V2) Configuration Set.
---

# Resource: aws_sesv2_configuration_set

Terraform resource for managing an AWS SESv2 (Simple Email V2) Configuration Set.

## Example Usage

### Basic Usage

```terraform
resource "aws_sesv2_configuration_set" "example" {
  configuration_set_name = "example"

  delivery_options {
    tls_policy = "REQUIRE"
  }

  reputation_options {
    reputation_metrics_enabled = false
  }

  sending_options {
    sending_enabled = true
  }

  suppression_options {
    suppressed_reasons = ["BOUNCE", "COMPLAINT"]
  }

  tracking_options {
    custom_redirect_domain = "example.com"
  }
}
```

## Argument Reference

This resource supports the following arguments:

* `configuration_set_name` - (Required) The name of the configuration set.
* `delivery_options` - (Optional) An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.
* `reputation_options` - (Optional) An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.
* `sending_options` - (Optional) An object that defines whether or not Amazon SES can send email that you send using the configuration set.
* `suppression_options` - (Optional) An object that contains information about the suppression list preferences for your account.
* `tags` - (Optional) A map of tags to assign to the service. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `tracking_options` - (Optional) An object that defines the open and click tracking options for emails that you send using the configuration set.
* `vdm_options` - (Optional) An object that defines the VDM settings that apply to emails that you send using the configuration set.

### delivery_options

This argument supports the following arguments:

* `sending_pool_name` - (Optional) The name of the dedicated IP pool to associate with the configuration set.
* `tls_policy` - (Optional) Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS). Valid values: `REQUIRE`, `OPTIONAL`.

### reputation_options

This argument supports the following arguments:

* `reputation_metrics_enabled` - (Optional) If `true`, tracking of reputation metrics is enabled for the configuration set. If `false`, tracking of reputation metrics is disabled for the configuration set.

### sending_options

This argument supports the following arguments:

* `sending_enabled` - (Optional) If `true`, email sending is enabled for the configuration set. If `false`, email sending is disabled for the configuration set.

### suppression_options

* `suppressed_reasons` - (Optional) A list that contains the reasons that email addresses are automatically added to the suppression list for your account. Valid values: `BOUNCE`, `COMPLAINT`.

### tracking_options

* `custom_redirect_domain` - (Required) The domain to use for tracking open and click events.

### vdm_options

* `dashboard_options` - (Optional) Specifies additional settings for your VDM configuration as applicable to the Dashboard.
* `guardian_options` - (Optional) Specifies additional settings for your VDM configuration as applicable to the Guardian.

### dashboard_options

* `engagement_metrics` - (Optional) Specifies the status of your VDM engagement metrics collection. Valid values: `ENABLED`, `DISABLED`.

### guardian_options

* `optimized_shared_delivery` - (Optional) Specifies the status of your VDM optimized shared delivery. Valid values: `ENABLED`, `DISABLED`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Configuration Set.
* `reputation_options` - An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.
    * `last_fresh_start` - The date and time (in Unix time) when the reputation metrics were last given a fresh start. When your account is given a fresh start, your reputation metrics are calculated starting from the date of the fresh start.

## Import

Import SESv2 (Simple Email V2) Configuration Set using the `configuration_set_name`. For example:

```
$ terraform import aws_sesv2_configuration_set.example example
```
