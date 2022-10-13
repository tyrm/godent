package http

import "net/http"

type ErrCode string

const (
	ErrCodeSuccess ErrCode = "M_SUCCESS"

	ErrCodeBadJSON             ErrCode = "M_BAD_JSON"
	ErrCodeEmailSendError      ErrCode = "M_EMAIL_SEND_ERROR"
	ErrCodeInvalidAddress      ErrCode = "M_INVALID_ADDRESS"
	ErrCodeInvalidEmail        ErrCode = "M_INVALID_EMAIL"
	ErrCodeInvalidParam        ErrCode = "M_INVALID_PARAM"
	ErrCodeMethodNotAllowed    ErrCode = "M_METHOD_NOT_ALLOWED" // not part of spec
	ErrCodeMissingParams       ErrCode = "M_MISSING_PARAMS"
	ErrCodeNoValidSession      ErrCode = "M_NO_VALID_SESSION"
	ErrCodeNotFound            ErrCode = "M_NOT_FOUND"
	ErrCodeSendError           ErrCode = "M_SEND_ERROR"
	ErrCodeSessionExpired      ErrCode = "M_SESSION_EXPIRED"
	ErrCodeSessionNotValidated ErrCode = "M_SESSION_NOT_VALIDATED"
	ErrCodeTermsNotSigned      ErrCode = "M_TERMS_NOT_SIGNED"
	ErrCodeThreePIDInUse       ErrCode = "M_THREEPID_IN_USE"
	ErrCodeUnauthorized        ErrCode = "M_UNAUTHORIZED"
	ErrCodeUnknown             ErrCode = "M_UNKNOWN"
	ErrCodeUnrecognized        ErrCode = "M_UNRECOGNIZED"
)

var errCodeHTTPStatus = map[ErrCode]int{
	ErrCodeBadJSON:             http.StatusBadRequest,
	ErrCodeEmailSendError:      http.StatusInternalServerError,
	ErrCodeInvalidAddress:      http.StatusBadRequest,
	ErrCodeInvalidEmail:        http.StatusBadRequest,
	ErrCodeInvalidParam:        http.StatusBadRequest,
	ErrCodeMissingParams:       http.StatusBadRequest,
	ErrCodeNoValidSession:      http.StatusForbidden,
	ErrCodeMethodNotAllowed:    http.StatusMethodNotAllowed,
	ErrCodeNotFound:            http.StatusNotFound,
	ErrCodeSendError:           http.StatusInternalServerError,
	ErrCodeSessionExpired:      http.StatusForbidden,
	ErrCodeSessionNotValidated: http.StatusForbidden,
	ErrCodeTermsNotSigned:      http.StatusForbidden,
	ErrCodeThreePIDInUse:       http.StatusConflict,
	ErrCodeUnauthorized:        http.StatusUnauthorized,
	ErrCodeUnknown:             http.StatusInternalServerError,
	ErrCodeUnrecognized:        http.StatusBadRequest,
}
