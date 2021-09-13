package utils

import (
	"encoding/json"
	"io"

	"github.com/kimnai/online_shopping_utils/error"
)

func RequestBodyParser(reader io.Reader) (map[string]string, error.ErrorInterface) {
	result := make(map[string]string)

	bytes, err := io.ReadAll(reader)
	if err != nil {
		e := error.NewBadRequestError("error when parsing body")
		return nil, e
	}

	if err := json.Unmarshal(bytes, &result); err != nil {
		e := error.NewBadRequestError("error when parsing body")
		return nil, e
	}

	return result, nil
}
