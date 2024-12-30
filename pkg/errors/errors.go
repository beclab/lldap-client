package errors

type Status struct {
	// Status of the operation
	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`

	Reason StatusReason `json:"reason,omitempty"`

	Code int32 `json:"code,omitempty"`
}

func ReasonForError(err error) StatusReason {
	if err == nil {
		return StatusReasonUnknown
	}
	status := ErrorToAPIStatus(err)
	return status.Reason
}

func IsNotFound(err error) bool {
	return ReasonForError(err) == StatusReasonNotFound
}

func IsAlreadyExists(err error) bool {
	return ReasonForError(err) == StatusReasonAlreadyExists
}

func IsForbidden(err error) bool {
	return ReasonForError(err) == StatusReasonForbidden
}

type StatusError struct {
	ErrStatus Status
}

type APIStatus interface {
	Status() Status
}

var _ error = &StatusError{}

func (e *StatusError) Error() string {
	return e.ErrStatus.Message
}

func (e *StatusError) Status() Status {
	return e.ErrStatus
}
