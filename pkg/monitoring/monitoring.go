package monitoring

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"

	"github.com/abdivasiyev/microservice_template/pkg/monitoring/open_telemetry"
	"github.com/abdivasiyev/microservice_template/pkg/monitoring/sentry"
)

var FxOption = fx.Provide(
	fx.Annotate(sentry.New, fx.As(new(Sentry))),
	fx.Annotate(open_telemetry.New, fx.As(new(Tracer))),
)

type Sentry interface {
	Error(err error)
	Message(format string, args ...any)
}

type Tracer interface {
	Start(ctx context.Context, spanName string) (trace.Span, context.Context)
}
