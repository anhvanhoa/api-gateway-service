package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcErrorResponse struct {
	Code    codes.Code `json:"code"`
	Message string     `json:"message"`
	Details any        `json:"details,omitempty"`
}

func HandleGrpcError(err error) (int, GrpcErrorResponse) {
	if err == nil {
		return http.StatusOK, GrpcErrorResponse{}
	}

	if st, ok := status.FromError(err); ok {
		switch st.Code() {
		case codes.InvalidArgument:
			return http.StatusBadRequest, GrpcErrorResponse{
				Code:    codes.InvalidArgument,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.Unauthenticated:
			return http.StatusUnauthorized, GrpcErrorResponse{
				Code:    codes.Unauthenticated,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.PermissionDenied:
			return http.StatusForbidden, GrpcErrorResponse{
				Code:    codes.PermissionDenied,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.NotFound:
			return http.StatusNotFound, GrpcErrorResponse{
				Code:    codes.NotFound,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.AlreadyExists:
			return http.StatusConflict, GrpcErrorResponse{
				Code:    codes.AlreadyExists,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.FailedPrecondition:
			return http.StatusPreconditionFailed, GrpcErrorResponse{
				Code:    codes.FailedPrecondition,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.ResourceExhausted:
			return http.StatusTooManyRequests, GrpcErrorResponse{
				Code:    codes.ResourceExhausted,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.Unimplemented:
			return http.StatusNotImplemented, GrpcErrorResponse{
				Code:    codes.Unimplemented,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.Internal:
			return http.StatusInternalServerError, GrpcErrorResponse{
				Code:    codes.Internal,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.Unavailable:
			return http.StatusServiceUnavailable, GrpcErrorResponse{
				Code:    codes.Unavailable,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.DeadlineExceeded:
			return http.StatusRequestTimeout, GrpcErrorResponse{
				Code:    codes.DeadlineExceeded,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.Canceled:
			return http.StatusRequestTimeout, GrpcErrorResponse{
				Code:    codes.Canceled,
				Message: st.Message(),
				Details: st.Details(),
			}
		case codes.DataLoss:
			return http.StatusInternalServerError, GrpcErrorResponse{
				Code:    codes.DataLoss,
				Message: st.Message(),
				Details: st.Details(),
			}
		default:
			return http.StatusInternalServerError, GrpcErrorResponse{
				Code:    codes.Unknown,
				Message: st.Message(),
				Details: st.Details(),
			}
		}
	}

	return http.StatusInternalServerError, GrpcErrorResponse{
		Code:    codes.Internal,
		Message: err.Error(),
	}
}

func RespondWithGrpcError(c *gin.Context, err error) {
	statusCode, errorResponse := HandleGrpcError(err)
	c.JSON(statusCode, errorResponse)
}
