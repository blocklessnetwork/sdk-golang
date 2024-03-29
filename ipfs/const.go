package ipfs

const (
	SUCCESS Error = iota
	INVALID_HANDLE
	UTF8_ERROR
	INVALID_METHOD
	INVALID_PARAMETER
	INVALID_ENCODING
	REQUEST_ERROR
	RUNTIME_ERROR
	TOO_MANY_SESSIONS
	INVALID_DRIVER
	PERMISSION_DENY

	BUFFER_TOO_SMALL
)
