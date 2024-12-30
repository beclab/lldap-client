package errors

const (
	StatusSuccess = "Success"
	StatusFailure = "Failure"
)

type StatusReason string

const (
	StatusReasonUnknown StatusReason = ""

	StatusReasonUnauthorized StatusReason = "Unauthorized"

	StatusReasonForbidden StatusReason = "Forbidden"
	StatusReasonNotFound  StatusReason = "NotFound"

	StatusReasonAlreadyExists StatusReason = "AlreadyExists"

	StatusReasonBadRequest         StatusReason = "BadRequest"
	StatusReasonInternalError      StatusReason = "InternalError"
	StatusReasonServiceUnavailable StatusReason = "ServiceUnavailable"
)
