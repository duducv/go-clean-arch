package applicationrepository

import "context"

type TracingRepository interface {
	StartSpan(ctx context.Context, spanName string, spanType string) (Span, context.Context)
	CaptureError(ctx context.Context, err error)
}

type Span interface {
	End()
}
