package errorsRes

import "errors"

// Services
// Token Service

var TokenGenerateError = errors.New("Token generate error")
var InvalidTokenError = errors.New("Invalid token")
