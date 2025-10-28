package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CustomErrorResponse định nghĩa cấu trúc response lỗi tùy chỉnh
type CustomErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details []any  `json:"details"`
}

// CustomErrorHandler xử lý lỗi gRPC và trả về thông báo tùy chỉnh
func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	// Kiểm tra nếu là lỗi kết nối gRPC
	if st, ok := status.FromError(err); ok {
		if st.Code() == codes.Unavailable ||
			strings.Contains(st.Message(), "connection error") ||
			strings.Contains(st.Message(), "transport: Error while dialing") ||
			strings.Contains(st.Message(), "connectex: No connection could be made") {

			// Trả về lỗi tùy chỉnh
			customError := CustomErrorResponse{
				Code:    503,
				Message: "Service is currently unavailable. Please try again later.",
				Details: []any{},
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)

			// Marshal và gửi response
			if jsonBytes, marshalErr := marshaler.Marshal(customError); marshalErr == nil {
				w.Write(jsonBytes)
			} else {
				// Fallback nếu không thể marshal
				w.Write([]byte(`{"code":503,"message":"Service is currently unavailable. Please try again later.","details":[]}`))
			}
			return
		}
	}

	// Nếu không phải lỗi kết nối, sử dụng default handler
	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}
