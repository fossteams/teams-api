package errors

import "fmt"

type GenericHTTPError struct {
	statusCode int
	expectedStatusCode int
	body []byte
}

func NewHTTPError(expectedStatusCode int, statusCode int, body []byte) GenericHTTPError {
	return GenericHTTPError{
		statusCode:         statusCode,
		expectedStatusCode: expectedStatusCode,
		body:               body,
	}
}

func (g GenericHTTPError) Error() string {
	if g.body == nil {
		return fmt.Sprintf("invalid status code returned: %d expected but %d returned",
			g.expectedStatusCode,
			g.statusCode)
	}

	return fmt.Sprintf("invalid status code returned: %d expected but %d returned: %v",
		g.expectedStatusCode,
		g.statusCode,
		string(g.body),
	)
}

var _ error = GenericHTTPError{}
