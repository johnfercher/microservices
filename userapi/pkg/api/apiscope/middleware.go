package apiscope

import (
	"github.com/google/uuid"
	"github.com/johnfercher/microservices/userapi/pkg/api"
	"github.com/johnfercher/microservices/userapi/pkg/api/apilog"
	"go.uber.org/zap"
	"net/http"
)

func LifecycleCtxSetup() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(api.HeaderRequestId)

			if requestId == "" {
				id, _ := uuid.NewRandom()
				requestId = id.String()
			}

			ctx := r.Context()

			logger := apilog.New()
			logger = logger.With(zap.String(api.CtxRequestId, requestId))
			logger = logger.With(zap.String(api.CtxRequestMethod, r.Method))
			logger = logger.With(zap.String(api.CtxRequestHost, r.Host))
			logger = logger.With(zap.Int64(api.CtxContentLength, r.ContentLength))

			if r.TLS != nil {
				logger = logger.With(zap.Int64(api.CtxTlsVersion, r.ContentLength))
			}

			if r.URL != nil {
				logger = logger.With(zap.String(api.CtxRequestPath, r.URL.Path))
				logger = logger.With(zap.String(api.CtxRequestQuery, r.URL.RawQuery))
			}

			ctx = api.AddContextLogger(ctx, logger)
			ctx = api.AddContextField(ctx, api.CtxRequestId, requestId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
