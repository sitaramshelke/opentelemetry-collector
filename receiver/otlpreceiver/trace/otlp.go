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

package trace

import (
	"context"

	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/internal"
	collectortrace "go.opentelemetry.io/collector/internal/data/protogen/collector/trace/v1"
	otlptrace "go.opentelemetry.io/collector/internal/data/protogen/trace/v1"
	"go.opentelemetry.io/collector/obsreport"
)

const (
	dataFormatProtobuf = "protobuf"
)

// Receiver is the type used to handle spans from OpenTelemetry exporters.
type Receiver struct {
	instanceName string
	nextConsumer consumer.Traces
}

// New creates a new Receiver reference.
func New(instanceName string, nextConsumer consumer.Traces) *Receiver {
	r := &Receiver{
		instanceName: instanceName,
		nextConsumer: nextConsumer,
	}

	return r
}

const (
	receiverTagValue  = "otlp_trace"
	receiverTransport = "grpc"
)

func (r *Receiver) Export(ctx context.Context, req *collectortrace.ExportTraceServiceRequest) (*collectortrace.ExportTraceServiceResponse, error) {
	// We need to ensure that it propagates the receiver name as a tag
	ctxWithReceiverName := obsreport.ReceiverContext(ctx, r.instanceName, receiverTransport)

	// Perform backward compatibility conversion of Span Status code according to
	// OTLP specification as we are a new receiver and sender (we are pushing data to the pipelines):
	// See https://github.com/open-telemetry/opentelemetry-proto/blob/59c488bfb8fb6d0458ad6425758b70259ff4a2bd/opentelemetry/proto/trace/v1/trace.proto#L239
	// See https://github.com/open-telemetry/opentelemetry-proto/blob/59c488bfb8fb6d0458ad6425758b70259ff4a2bd/opentelemetry/proto/trace/v1/trace.proto#L253
	for _, rss := range req.ResourceSpans {
		for _, ils := range rss.InstrumentationLibrarySpans {
			for _, span := range ils.Spans {
				switch span.Status.Code {
				case otlptrace.Status_STATUS_CODE_UNSET:
					if span.Status.DeprecatedCode != otlptrace.Status_DEPRECATED_STATUS_CODE_OK {
						span.Status.Code = otlptrace.Status_STATUS_CODE_ERROR
					}
				case otlptrace.Status_STATUS_CODE_OK:
					// If status code is set then overwrites deprecated.
					span.Status.DeprecatedCode = otlptrace.Status_DEPRECATED_STATUS_CODE_OK
				case otlptrace.Status_STATUS_CODE_ERROR:
					span.Status.DeprecatedCode = otlptrace.Status_DEPRECATED_STATUS_CODE_UNKNOWN_ERROR
				}
			}
		}
	}

	td := pdata.TracesFromInternalRep(internal.TracesFromOtlp(req))
	err := r.sendToNextConsumer(ctxWithReceiverName, td)
	if err != nil {
		return nil, err
	}

	return &collectortrace.ExportTraceServiceResponse{}, nil
}

func (r *Receiver) sendToNextConsumer(ctx context.Context, td pdata.Traces) error {
	numSpans := td.SpanCount()
	if numSpans == 0 {
		return nil
	}

	if c, ok := client.FromGRPC(ctx); ok {
		ctx = client.NewContext(ctx, c)
	}

	ctx = obsreport.StartTraceDataReceiveOp(ctx, r.instanceName, receiverTransport)
	err := r.nextConsumer.ConsumeTraces(ctx, td)
	obsreport.EndTraceDataReceiveOp(ctx, dataFormatProtobuf, numSpans, err)

	return err
}
