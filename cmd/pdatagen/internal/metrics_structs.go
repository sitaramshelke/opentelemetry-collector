// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

var metricsFile = &File{
	Name: "metrics",
	imports: []string{
		`otlpmetrics "go.opentelemetry.io/collector/internal/data/protogen/metrics/v1"`,
	},
	testImports: []string{
		`"testing"`,
		``,
		`"github.com/stretchr/testify/assert"`,
		``,
		`otlpmetrics "go.opentelemetry.io/collector/internal/data/protogen/metrics/v1"`,
	},
	structs: []baseStruct{
		resourceMetricsSlice,
		resourceMetrics,
		instrumentationLibraryMetricsSlice,
		instrumentationLibraryMetrics,
		metricSlice,
		metric,
		intGauge,
		doubleGauge,
		intSum,
		doubleSum,
		intHistogram,
		doubleHistogram,
		doubleSummary,
		intDataPointSlice,
		intDataPoint,
		doubleDataPointSlice,
		doubleDataPoint,
		intHistogramDataPointSlice,
		intHistogramDataPoint,
		doubleHistogramDataPointSlice,
		doubleHistogramDataPoint,
		doubleSummaryDataPointSlice,
		doubleSummaryDataPoint,
		quantileValuesSlice,
		quantileValues,
		intExemplarSlice,
		intExemplar,
		doubleExemplarSlice,
		doubleExemplar,
	},
}

var resourceMetricsSlice = &sliceOfPtrs{
	structName: "ResourceMetricsSlice",
	element:    resourceMetrics,
}

var resourceMetrics = &messageValueStruct{
	structName:     "ResourceMetrics",
	description:    "// InstrumentationLibraryMetrics is a collection of metrics from a LibraryInstrumentation.",
	originFullName: "otlpmetrics.ResourceMetrics",
	fields: []baseField{
		resourceField,
		&sliceField{
			fieldName:       "InstrumentationLibraryMetrics",
			originFieldName: "InstrumentationLibraryMetrics",
			returnSlice:     instrumentationLibraryMetricsSlice,
		},
	},
}

var instrumentationLibraryMetricsSlice = &sliceOfPtrs{
	structName: "InstrumentationLibraryMetricsSlice",
	element:    instrumentationLibraryMetrics,
}

var instrumentationLibraryMetrics = &messageValueStruct{
	structName:     "InstrumentationLibraryMetrics",
	description:    "// InstrumentationLibraryMetrics is a collection of metrics from a LibraryInstrumentation.",
	originFullName: "otlpmetrics.InstrumentationLibraryMetrics",
	fields: []baseField{
		instrumentationLibraryField,
		&sliceField{
			fieldName:       "Metrics",
			originFieldName: "Metrics",
			returnSlice:     metricSlice,
		},
	},
}

var metricSlice = &sliceOfPtrs{
	structName: "MetricSlice",
	element:    metric,
}

var metric = &messageValueStruct{
	structName: "Metric",
	description: "// Metric represents one metric as a collection of datapoints.\n" +
		"// See Metric definition in OTLP: https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/metrics/v1/metrics.proto",
	originFullName: "otlpmetrics.Metric",
	fields: []baseField{
		nameField,
		&primitiveField{
			fieldName:       "Description",
			originFieldName: "Description",
			returnType:      "string",
			defaultVal:      `""`,
			testVal:         `"test_description"`,
		},
		&primitiveField{
			fieldName:       "Unit",
			originFieldName: "Unit",
			returnType:      "string",
			defaultVal:      `""`,
			testVal:         `"1"`,
		},
		oneofDataField,
	},
}

var intGauge = &messageValueStruct{
	structName:     "IntGauge",
	description:    "// IntGauge represents the type of a int scalar metric that always exports the \"current value\" for every data point.",
	originFullName: "otlpmetrics.IntGauge",
	fields: []baseField{
		&sliceField{
			fieldName:       "DataPoints",
			originFieldName: "DataPoints",
			returnSlice:     intDataPointSlice,
		},
	},
}

var doubleGauge = &messageValueStruct{
	structName:     "DoubleGauge",
	description:    "// DoubleGauge represents the type of a double scalar metric that always exports the \"current value\" for every data point.",
	originFullName: "otlpmetrics.DoubleGauge",
	fields: []baseField{
		&sliceField{
			fieldName:       "DataPoints",
			originFieldName: "DataPoints",
			returnSlice:     doubleDataPointSlice,
		},
	},
}

var intSum = &messageValueStruct{
	structName:     "IntSum",
	description:    "// IntSum represents the type of a numeric int scalar metric that is calculated as a sum of all reported measurements over a time interval.",
	originFullName: "otlpmetrics.IntSum",
	fields: []baseField{
		aggregationTemporalityField,
		isMonotonicField,
		&sliceField{
			fieldName:       "DataPoints",
			originFieldName: "DataPoints",
			returnSlice:     intDataPointSlice,
		},
	},
}

var doubleSum = &messageValueStruct{
	structName:     "DoubleSum",
	description:    "// DoubleSum represents the type of a numeric double scalar metric that is calculated as a sum of all reported measurements over a time interval.",
	originFullName: "otlpmetrics.DoubleSum",
	fields: []baseField{
		aggregationTemporalityField,
		isMonotonicField,
		&sliceField{
			fieldName:       "DataPoints",
			originFieldName: "DataPoints",
			returnSlice:     doubleDataPointSlice,
		},
	},
}

var intHistogram = &messageValueStruct{
	structName:     "IntHistogram",
	description:    "// IntHistogram represents the type of a metric that is calculated by aggregating as a Histogram of all reported double measurements over a time interval.",
	originFullName: "otlpmetrics.IntHistogram",
	fields: []baseField{
		aggregationTemporalityField,
		&sliceField{
			fieldName:       "DataPoints",
			originFieldName: "DataPoints",
			returnSlice:     intHistogramDataPointSlice,
		},
	},
}

var doubleHistogram = &messageValueStruct{
	structName:     "DoubleHistogram",
	description:    "// DoubleHistogram represents the type of a metric that is calculated by aggregating as a Histogram of all reported double measurements over a time interval.",
	originFullName: "otlpmetrics.DoubleHistogram",
	fields: []baseField{
		aggregationTemporalityField,
		&sliceField{
			fieldName:       "DataPoints",
			originFieldName: "DataPoints",
			returnSlice:     doubleHistogramDataPointSlice,
		},
	},
}

var doubleSummary = &messageValueStruct{
	structName:     "Summary",
	description:    "// Summary represents the type of a metric that is calculated by aggregating as a Summary of all reported double measurements over a time interval.",
	originFullName: "otlpmetrics.DoubleSummary",
	fields: []baseField{
		&sliceField{
			fieldName:       "DataPoints",
			originFieldName: "DataPoints",
			returnSlice:     doubleSummaryDataPointSlice,
		},
	},
}

var intDataPointSlice = &sliceOfPtrs{
	structName: "IntDataPointSlice",
	element:    intDataPoint,
}

var intDataPoint = &messageValueStruct{
	structName:     "IntDataPoint",
	description:    "// IntDataPoint is a single data point in a timeseries that describes the time-varying values of a scalar int metric.",
	originFullName: "otlpmetrics.IntDataPoint",
	fields: []baseField{
		labelsField,
		startTimeField,
		timeField,
		valueInt64Field,
		intExemplarsField,
	},
}

var doubleDataPointSlice = &sliceOfPtrs{
	structName: "DoubleDataPointSlice",
	element:    doubleDataPoint,
}

var doubleDataPoint = &messageValueStruct{
	structName:     "DoubleDataPoint",
	description:    "// DoubleDataPoint is a single data point in a timeseries that describes the time-varying value of a double metric.",
	originFullName: "otlpmetrics.DoubleDataPoint",
	fields: []baseField{
		labelsField,
		startTimeField,
		timeField,
		valueFloat64Field,
		doubleExemplarsField,
	},
}

var intHistogramDataPointSlice = &sliceOfPtrs{
	structName: "IntHistogramDataPointSlice",
	element:    intHistogramDataPoint,
}

var intHistogramDataPoint = &messageValueStruct{
	structName:     "IntHistogramDataPoint",
	description:    "// IntHistogramDataPoint is a single data point in a timeseries that describes the time-varying values of a Histogram of int values.",
	originFullName: "otlpmetrics.IntHistogramDataPoint",
	fields: []baseField{
		labelsField,
		startTimeField,
		timeField,
		countField,
		intSumField,
		bucketCountsField,
		explicitBoundsField,
		intExemplarsField,
	},
}

var doubleHistogramDataPointSlice = &sliceOfPtrs{
	structName: "DoubleHistogramDataPointSlice",
	element:    doubleHistogramDataPoint,
}

var doubleHistogramDataPoint = &messageValueStruct{
	structName:     "DoubleHistogramDataPoint",
	description:    "// DoubleHistogramDataPoint is a single data point in a timeseries that describes the time-varying values of a Histogram of double values.",
	originFullName: "otlpmetrics.DoubleHistogramDataPoint",
	fields: []baseField{
		labelsField,
		startTimeField,
		timeField,
		countField,
		doubleSumField,
		bucketCountsField,
		explicitBoundsField,
		doubleExemplarsField,
	},
}

var doubleSummaryDataPointSlice = &sliceOfPtrs{
	structName: "SummaryDataPointSlice",
	element:    doubleSummaryDataPoint,
}

var doubleSummaryDataPoint = &messageValueStruct{
	structName:     "SummaryDataPoint",
	description:    "// SummaryDataPoint is a single data point in a timeseries that describes the time-varying values of a Summary of double values.",
	originFullName: "otlpmetrics.DoubleSummaryDataPoint",
	fields: []baseField{
		labelsField,
		startTimeField,
		timeField,
		countField,
		doubleSumField,
		&sliceField{
			fieldName:       "QuantileValues",
			originFieldName: "QuantileValues",
			returnSlice:     quantileValuesSlice,
		},
	},
}

var quantileValuesSlice = &sliceOfPtrs{
	structName: "ValueAtQuantileSlice",
	element:    quantileValues,
}

var quantileValues = &messageValueStruct{
	structName:     "ValueAtQuantile",
	description:    "// ValueAtQuantile is a quantile value within a Summary data point",
	originFullName: "otlpmetrics.DoubleSummaryDataPoint_ValueAtQuantile",
	fields: []baseField{
		quantileField,
		valueFloat64Field,
	},
}

var intExemplarSlice = &sliceOfValues{
	structName: "IntExemplarSlice",
	element:    intExemplar,
}

var intExemplar = &messageValueStruct{
	structName: "IntExemplar",
	description: "// IntExemplar is a sample input int measurement.\n//\n" +
		"// Exemplars also hold information about the environment when the measurement was recorded,\n" +
		"// for example the span and trace ID of the active span when the exemplar was recorded.",

	originFullName: "otlpmetrics.IntExemplar",
	fields: []baseField{
		timeField,
		valueInt64Field,
		&sliceField{
			fieldName:       "FilteredLabels",
			originFieldName: "FilteredLabels",
			returnSlice:     stringMap,
		},
	},
}

var doubleExemplarSlice = &sliceOfValues{
	structName: "DoubleExemplarSlice",
	element:    doubleExemplar,
}

var doubleExemplar = &messageValueStruct{
	structName: "DoubleExemplar",
	description: "// DoubleExemplar is a sample input double measurement.\n//\n" +
		"// Exemplars also hold information about the environment when the measurement was recorded,\n" +
		"// for example the span and trace ID of the active span when the exemplar was recorded.",

	originFullName: "otlpmetrics.DoubleExemplar",
	fields: []baseField{
		timeField,
		valueFloat64Field,
		&sliceField{
			fieldName:       "FilteredLabels",
			originFieldName: "FilteredLabels",
			returnSlice:     stringMap,
		},
	},
}

var labelsField = &sliceField{
	fieldName:       "LabelsMap",
	originFieldName: "Labels",
	returnSlice:     stringMap,
}

var intExemplarsField = &sliceField{
	fieldName:       "Exemplars",
	originFieldName: "Exemplars",
	returnSlice:     intExemplarSlice,
}

var doubleExemplarsField = &sliceField{
	fieldName:       "Exemplars",
	originFieldName: "Exemplars",
	returnSlice:     doubleExemplarSlice,
}

var countField = &primitiveField{
	fieldName:       "Count",
	originFieldName: "Count",
	returnType:      "uint64",
	defaultVal:      "uint64(0)",
	testVal:         "uint64(17)",
}

var intSumField = &primitiveField{
	fieldName:       "Sum",
	originFieldName: "Sum",
	returnType:      "int64",
	defaultVal:      "int64(0.0)",
	testVal:         "int64(1713)",
}

var doubleSumField = &primitiveField{
	fieldName:       "Sum",
	originFieldName: "Sum",
	returnType:      "float64",
	defaultVal:      "float64(0.0)",
	testVal:         "float64(17.13)",
}

var valueInt64Field = &primitiveField{
	fieldName:       "Value",
	originFieldName: "Value",
	returnType:      "int64",
	defaultVal:      "int64(0)",
	testVal:         "int64(-17)",
}

var valueFloat64Field = &primitiveField{
	fieldName:       "Value",
	originFieldName: "Value",
	returnType:      "float64",
	defaultVal:      "float64(0.0)",
	testVal:         "float64(17.13)",
}

var bucketCountsField = &primitiveField{
	fieldName:       "BucketCounts",
	originFieldName: "BucketCounts",
	returnType:      "[]uint64",
	defaultVal:      "[]uint64(nil)",
	testVal:         "[]uint64{1, 2, 3}",
}

var explicitBoundsField = &primitiveField{
	fieldName:       "ExplicitBounds",
	originFieldName: "ExplicitBounds",
	returnType:      "[]float64",
	defaultVal:      "[]float64(nil)",
	testVal:         "[]float64{1, 2, 3}",
}

var quantileField = &primitiveField{
	fieldName:       "Quantile",
	originFieldName: "Quantile",
	returnType:      "float64",
	defaultVal:      "float64(0.0)",
	testVal:         "float64(17.13)",
}

var isMonotonicField = &primitiveField{
	fieldName:       "IsMonotonic",
	originFieldName: "IsMonotonic",
	returnType:      "bool",
	defaultVal:      "false",
	testVal:         "true",
}

var aggregationTemporalityField = &primitiveTypedField{
	fieldName:       "AggregationTemporality",
	originFieldName: "AggregationTemporality",
	returnType:      "AggregationTemporality",
	rawType:         "otlpmetrics.AggregationTemporality",
	defaultVal:      "AggregationTemporalityUnspecified",
	testVal:         "AggregationTemporalityCumulative",
}

var oneofDataField = &oneofField{
	copyFuncName:    "copyData",
	originFieldName: "Data",
	testVal:         "&otlpmetrics.Metric_IntGauge{IntGauge: &otlpmetrics.IntGauge{}}",
	fillTestName:    "IntGauge",
}
