package constant

import "errors"

var (
	UnprocessableEntity = errors.New("unprocessable entity")

	AccessDenied = errors.New("access denied")

	InvalidJsonPayload     = errors.New("invalid request payload")
	InvalidQueryParameters = errors.New("invalid query parameters")
	FailedToLoadTimeZone   = errors.New("failed to load timezone")
	UnexpectedError        = errors.New("unexpected error")
	JsonUnmarshalError     = errors.New("json unmarshal error")
	JsonMarshalError       = errors.New("json marshal error")

	InvalidHeaderFormat             = errors.New("invalid header format")
	MalformedToken                  = errors.New("malformed token")
	TokenHasExpired                 = errors.New("token has expired")
	TokenIsNotValid                 = errors.New("token is not valid")
	TokenInvalid                    = errors.New("token is invalid")
	InvalidJwtExpiredDurationFormat = errors.New("invalid jwt expired duration")

	RoleNotFound = errors.New("role not found")

	EmailAlreadyExists        = errors.New("email already exists")
	UsernameOrPasswordInvalid = errors.New("username or password invalid")
	UserNotFound              = errors.New("user not found")

	ConsumerNotFound         = errors.New("consumer not found")
	NIKConsumerAlreadyExists = errors.New("nik consumer already exists")
)
