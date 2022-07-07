package constants

// general
const (
	InternalServerErrCode = 1001 + iota
	BadRequesErrCode
	NotAuthorizedErrCode
)

// repositories
const (
	QueryInternalServerErrCode = 2001 + iota
	QueryNotFoundErrCode
)

// user
const (
	RegisterEmailNotAvailableErrCode = 3001 + iota
	RegisterUsernameNotAvailableErrCode
	HashPasswordInternalErrCode
	LoginUsernameNotFoundErrCode
	LoginInvalidPasswordErrCode
)
