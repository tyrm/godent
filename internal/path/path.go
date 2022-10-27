package path

const (
	BasePath = "/_matrix/identity"
	V2       = BasePath + "/v2"

	// account.

	V2Account         = V2 + "/account"
	V2AccountLogout   = V2Account + "/logout"
	V2AccountRegister = V2Account + "/register"

	// pubkey.

	V2Pubkey                 = V2 + "/pubkey"
	V2PubkeyEphemeralIsvalid = V2Pubkey + "/ephemeral/isvalid"
	V2PubkeyIsvalid          = V2Pubkey + "/isvalid"
	V2PubkeyKey              = V2Pubkey + "/" + VarKey

	// terms.

	V2Terms = V2 + "/terms"

	// versions.

	Versions = BasePath + "/versions"
)
