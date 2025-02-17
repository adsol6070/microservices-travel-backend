package errors

const (
	// General Errors
	Success        = 200
	BadRequest     = 400
	Unauthorized   = 401
	Forbidden      = 403
	NotFound       = 404
	Conflict       = 409
	InternalError  = 500
	ServiceUnavailable = 503

	// Database Errors
	DBConnectionFailed = 600
	DBQueryFailed      = 601
	DBTransactionFailed = 602

	// Authentication Errors
	TokenExpired    = 700
	TokenInvalid    = 701
	InvalidCredentials = 702

	// Business Logic Errors
	UserAlreadyExists = 800
	OrderNotFound     = 801
	PaymentFailed     = 802
)