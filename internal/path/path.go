package path

const (
	BasePath = "/_matrix/identity"
	V2       = BasePath + "/v2"

	// account.

	V2Account         = V2 + "/account"
	V2AccountRegister = V2Account + "/register"

	// terms.

	V2Terms = V2 + "/terms"

	// versions.

	Versions = BasePath + "/versions"
)
