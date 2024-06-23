package tracing

import (
	"context"

	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
)

type NoTracingAdapter struct{}
type SpanMock struct{}

func (SpanMock) End() {}

func NewNotracingAdapter() applicationrepository.TracingRepository {
	return &NoTracingAdapter{}
}

func (NoTracingAdapter) CaptureError(ctx context.Context, err error) {}
func (t NoTracingAdapter) StartSpan(ctx context.Context, spanName string, spanType string) (applicationrepository.Span, context.Context) {
	return SpanMock{}, ctx
}
