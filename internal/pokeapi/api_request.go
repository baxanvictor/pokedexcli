package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendGetRequest[T any](url string, resultContainer *T) ([]byte, *ApiError) {
	res, err := http.Get(url)
	if err != nil {
		return nil, ApiErrorFromError(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, &ApiError{
			message: fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body),
		}
	}
	if err != nil {
		return nil, ApiErrorFromError(err)
	}
	err = UnmarshalResponseBody(&body, resultContainer)
	if err != nil {
		return nil, ApiErrorFromError(err)
	}

	return body, nil
}

func UnmarshalResponseBody[T any](body *[]byte, resultContainer *T) error {
	return json.Unmarshal(*body, &resultContainer)
}

type ApiError struct {
	message string
}

func (ae *ApiError) Error() string {
	return ae.message
}

func ApiErrorFromError(err error) *ApiError {
	apiError := ApiError{
		message: err.Error(),
	}
	return &apiError
}
