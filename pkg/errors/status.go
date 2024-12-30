package errors

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
)

type statusError interface {
	Status() Status
}

func ErrorToAPIStatus(err error) *Status {
	switch t := err.(type) {
	case statusError:
		status := t.Status()
		if len(status.Status) == 0 {
			status.Status = StatusFailure
		}
		switch status.Status {
		case StatusSuccess:
			if status.Code == 0 {
				status.Code = http.StatusOK
			}
		case StatusFailure:
			if status.Code == 0 {
				status.Code = http.StatusInternalServerError
			}
		default:
			if status.Code == 0 {
				status.Code = http.StatusInternalServerError
			}
		}
		return &status
	case gqlerror.List:
		status := statusFromError(err)
		return &status
	default:
		return &Status{
			Status:  StatusFailure,
			Code:    http.StatusInternalServerError,
			Reason:  StatusReasonInternalError,
			Message: err.Error(),
		}
	}
}

func statusFromError(err error) Status {
	status := Status{Status: StatusFailure}
	switch t := err.(type) {
	case gqlerror.List:
		if len(t) == 0 {
			return status
		}
		err := t[0]
		if err.Extensions != nil {
			if code, ok := err.Extensions["code"].(int32); ok {
				status.Code = code
			}
			if message, ok := err.Extensions["message"].(string); ok {
				status.Message = message
			}
			if reason, ok := err.Extensions["reason"].(string); ok {
				status.Reason = StatusReason(reason)
			}
		}
	default:
		status.Message = err.Error()
		status.Reason = StatusReasonUnknown
	}

	return status
}
