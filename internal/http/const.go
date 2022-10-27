package http

const (
	HeaderCacheControl = "Cache-Control"
	HeaderExpires      = "Expires"

	QueryAccessToken      = "access_token"
	QueryExpiresIn        = "expires_in"
	QueryMatrixServerName = "matrix_server_name"
	QueryPublicKey        = "public_key"
	QueryTokenType        = "token_type"

	ResponseDatabaseError  = "Database Error"
	ResponseMalformedJSON  = "Malformed JSON"
	ResponseTermsNotSigned = "Terms not signed"
	ResponseUnauthorized   = "Unauthorized"
)
