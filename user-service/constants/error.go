package constants

// general
const (
	InternalServerErrCode = 1001 + iota
	BadRequestErrCode
	NotAuthorizedErrCode
)

// repositories
const (
	QueryInternalServerErrCode = 2001 + iota
	QueryNotFoundErrCode
)

// user
const (
	UpdateUserEmailExistErrCode = 3001 + iota
	InvalidEmailFormatErrCode
	HashPasswordInternalErrCode
	LoginUsernameNotFoundErrCode
	LoginInvalidPasswordErrCode
)
