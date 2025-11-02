package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details []any  `json:"details"`
}

func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	if st, ok := status.FromError(err); ok {
		if st.Code() == codes.Unavailable ||
			strings.Contains(st.Message(), "connection error") ||
			strings.Contains(st.Message(), "transport: Error while dialing") ||
			strings.Contains(st.Message(), "connectex: No connection could be made") {

			customError := CustomErrorResponse{
				Code:    503,
				Message: "Service is currently unavailable. Please try again later.",
				Details: []any{},
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)

			if jsonBytes, marshalErr := marshaler.Marshal(customError); marshalErr == nil {
				w.Write(jsonBytes)
			} else {
				w.Write([]byte(`{"code":503,"message":"Service is currently unavailable. Please try again later.","details":[]}`))
			}
			return
		}
	}

	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}
