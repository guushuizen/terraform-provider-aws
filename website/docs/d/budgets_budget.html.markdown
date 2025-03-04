---
subcategory: "Web Services Budgets"
layout: "aws"
page_title: "AWS: aws_budgets_budget"
description: |-
  Terraform data source for managing an AWS Web Services Budgets Budget.
---

# Data Source: aws_budgets_budget

Terraform data source for managing an AWS Web Services Budgets Budget.

## Example Usage

### Basic Usage

```terraform
data "aws_budgets_budget" "test" {
  name = aws_budgets_budget.test.name
}
```

## Argument Reference

The following arguments are required:

* `name` - The name of a budget. Unique within accounts.

The following arguments are optional:

* `account_id` - The ID of the target account for budget. Will use current user's account_id by default if omitted.
* `name_prefix` - The prefix of the name of a budget. Unique within accounts.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `auto_adjust_data` - Object containing [AutoAdjustData] which determines the budget amount for an auto-adjusting budget.
* `budget_type` - Whether this budget tracks monetary cost or usage.
* `budget_limit` - The total amount of cost, usage, RI utilization, RI coverage, Savings Plans utilization, or Savings Plans coverage that you want to track with your budget. Contains object [Spend](#spend)
* `calculated_spend` - The spend objects that are associated with this budget. The [actualSpend](#actual-spend) tracks how much you've used, cost, usage, RI units, or Savings Plans units and the [forecastedSpend](#forecasted-spend) tracks how much that you're predicted to spend based on your historical usage profile.
* `cost_filter` - A list of [CostFilter](#cost-filter) name/values pair to apply to budget.
* `cost_types` - Object containing [CostTypes](#cost-types) The types of cost included in a budget, such as tax and subscriptions.
* `time_period_end` - The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
* `time_period_start` - The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
* `time_unit` - The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
* `notification` - Object containing [Budget Notifications](#budget-notification). Can be used multiple times to define more than one budget notification.
* `planned_limit` - Object containing [Planned Budget Limits](#planned-budget-limits). Can be used multiple times to plan more than one budget limit. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.

### Actual Spend

The amount of cost, usage, RI units, or Savings Plans units that you used. Type is [Spend](#spend)

### Auto Adjust Data

The parameters that determine the budget amount for an auto-adjusting budget.

`auto_adjust_type` (Required) - The string that defines whether your budget auto-adjusts based on historical or forecasted data. Valid values: `FORECAST`,`HISTORICAL`
`historical_options` (Optional) - Configuration block of [Historical Options](#historical-options). Required for `auto_adjust_type` of `HISTORICAL` Configuration block that defines the historical data that your auto-adjusting budget is based on.
`last_auto_adjust_time` (Optional) - The last time that your budget was auto-adjusted.

### Budget Notification

Valid keys for `notification` parameter.

* `comparison_operator` - (Required) Comparison operator to use to evaluate the condition. Can be `LESS_THAN`, `EQUAL_TO` or `GREATER_THAN`.
* `threshold` - (Required) Threshold when the notification should be sent.
* `threshold_type` - (Required) What kind of threshold is defined. Can be `PERCENTAGE` OR `ABSOLUTE_VALUE`.
* `notification_type` - (Required) What kind of budget value to notify on. Can be `ACTUAL` or `FORECASTED`
* `subscriber_email_addresses` - (Optional) E-Mail addresses to notify. Either this or `subscriber_sns_topic_arns` is required.
* `subscriber_sns_topic_arns` - (Optional) SNS topics to notify. Either this or `subscriber_email_addresses` is required.

### Cost Filter

Based on your choice of budget type, you can choose one or more of the available budget filters.

* `PurchaseType`
* `UsageTypeGroup`
* `Service`
* `Operation`
* `UsageType`
* `BillingEntity`
* `CostCategory`
* `LinkedAccount`
* `TagKeyValue`
* `LegalEntityName`
* `InvoicingEntity`
* `AZ`
* `Region`
* `InstanceType`

Refer to [AWS CostFilter documentation](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-create-filters.html) for further detail.

### Cost Types

Valid keys for `cost_types` parameter.

* `include_credit` - A boolean value whether to include credits in the cost budget. Defaults to `true`
* `include_discount` - Whether a budget includes discounts. Defaults to `true`
* `include_other_subscription` - A boolean value whether to include other subscription costs in the cost budget. Defaults to `true`
* `include_recurring` - A boolean value whether to include recurring costs in the cost budget. Defaults to `true`
* `include_refund` - A boolean value whether to include refunds in the cost budget. Defaults to `true`
* `include_subscription` - A boolean value whether to include subscriptions in the cost budget. Defaults to `true`
* `include_support` - A boolean value whether to include support costs in the cost budget. Defaults to `true`
* `include_tax` - A boolean value whether to include tax in the cost budget. Defaults to `true`
* `include_upfront` - A boolean value whether to include upfront costs in the cost budget. Defaults to `true`
* `use_amortized` - Whether a budget uses the amortized rate. Defaults to `false`
* `use_blended` - A boolean value whether to use blended costs in the cost budget. Defaults to `false`

Refer to [AWS CostTypes documentation](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_CostTypes.html) for further detail.

### Forecasted Spend

The amount of cost, usage, RI units, or Savings Plans units that you're forecasted to use.
Type is [Spend](#spend)

### Historical Options

`budget_adjustment_period` (Required) - The number of budget periods included in the moving-average calculation that determines your auto-adjusted budget amount.
`lookback_available_periods` (Optional) - The integer that describes how many budget periods in your BudgetAdjustmentPeriod are included in the calculation of your current budget limit. If the first budget period in your BudgetAdjustmentPeriod has no cost data, then that budget period isn’t included in the average that determines your budget limit. You can’t set your own LookBackAvailablePeriods. The value is automatically calculated from the `budget_adjustment_period` and your historical cost data.

### Planned Budget Limits

Valid keys for `planned_limit` parameter.

* `start_time` - (Required) The start time of the budget limit. Format: `2017-01-01_12:00`. See [PlannedBudgetLimits](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html#awscostmanagement-Type-budgets_Budget-PlannedBudgetLimits) documentation.
* `amount` - (Required) The amount of cost or usage being measured for a budget.
* `unit` - (Required) The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.

### Spend

`amount` - The cost or usage amount that's associated with a budget forecast, actual spend, or budget threshold. Length Constraints: Minimum length of `1`. Maximum length of `2147483647`.
`unit` - The unit of measurement that's used for the budget forecast, actual spend, or budget threshold, such as USD or GBP. Length Constraints: Minimum length of `1`. Maximum length of `2147483647`.
