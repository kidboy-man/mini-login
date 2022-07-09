package constants

// general
const (
	InternalServerErrCode = 1001 + iota
	BadRequesErrCode
	NotAuthorizedErrCode
	OrmHookDataErrCode // meaning this error occurs in gorm model hooks
)

// repositories
const (
	QueryInternalServerErrCode = 2001 + iota
	QueryNotFoundErrCode
)

// user
const (
	RegisterUsernameNotAvailableErrCode = 3001 + iota
	HashPasswordInternalErrCode
	LoginUsernameNotFoundErrCode
	LoginInvalidPasswordErrCode
)
