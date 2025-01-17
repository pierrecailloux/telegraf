//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package topk

func (t *TopK) SampleConfig() string {
	return `# Print all metrics that pass through this filter.
[[processors.topk]]
  ## How many seconds between aggregations
  # period = 10

  ## How many top buckets to return per field
  ## Every field specified to aggregate over will return k number of results.
  ## For example, 1 field with k of 10 will return 10 buckets. While 2 fields
  ## with k of 3 will return 6 buckets.
  # k = 10

  ## Over which tags should the aggregation be done. Globs can be specified, in
  ## which case any tag matching the glob will aggregated over. If set to an
  ## empty list is no aggregation over tags is done
  # group_by = ['*']

  ## The field(s) to aggregate
  ## Each field defined is used to create an independent aggregation. Each
  ## aggregation will return k buckets. If a metric does not have a defined
  ## field the metric will be dropped from the aggregation. Considering using
  ## the defaults processor plugin to ensure fields are set if required.
  # fields = ["value"]

  ## What aggregation function to use. Options: sum, mean, min, max
  # aggregation = "mean"

  ## Instead of the top k largest metrics, return the bottom k lowest metrics
  # bottomk = false

  ## The plugin assigns each metric a GroupBy tag generated from its name and
  ## tags. If this setting is different than "" the plugin will add a
  ## tag (which name will be the value of this setting) to each metric with
  ## the value of the calculated GroupBy tag. Useful for debugging
  # add_groupby_tag = ""

  ## These settings provide a way to know the position of each metric in
  ## the top k. The 'add_rank_field' setting allows to specify for which
  ## fields the position is required. If the list is non empty, then a field
  ## will be added to each and every metric for each string present in this
  ## setting. This field will contain the ranking of the group that
  ## the metric belonged to when aggregated over that field.
  ## The name of the field will be set to the name of the aggregation field,
  ## suffixed with the string '_topk_rank'
  # add_rank_fields = []

  ## These settings provide a way to know what values the plugin is generating
  ## when aggregating metrics. The 'add_aggregate_field' setting allows to
  ## specify for which fields the final aggregation value is required. If the
  ## list is non empty, then a field will be added to each every metric for
  ## each field present in this setting. This field will contain
  ## the computed aggregation for the group that the metric belonged to when
  ## aggregated over that field.
  ## The name of the field will be set to the name of the aggregation field,
  ## suffixed with the string '_topk_aggregate'
  # add_aggregate_fields = []
`
}
