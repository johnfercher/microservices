package apilog

import (
	"github.com/google/uuid"
	"net/http"
)

func LifecycleCtxSetup() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(HeaderRequestId)

			if requestId == "" {
				id, _ := uuid.NewRandom()
				requestId = id.String()
			}

			ctx := r.Context()

			logger := New()

			logger = logger.WithKeyValue(CtxRequestId, requestId)
			logger = logger.WithKeyValue(CtxRequestMethod, r.Method)
			logger = logger.WithKeyValue(CtxRequestHost, r.Host)
			logger = logger.WithKeyValue(CtxContentLength, r.ContentLength)

			if r.TLS != nil {
				logger = logger.WithKeyValue(CtxTlsVersion, r.ContentLength)
			}

			if r.URL != nil {
				logger = logger.WithKeyValue(CtxRequestPath, r.URL.Path)
				logger = logger.WithKeyValue(CtxRequestQuery, r.URL.RawQuery)
			}

			ctx = AddContextLogger(ctx, logger)
			ctx = AddContextField(ctx, CtxRequestId, requestId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
