package jsonanalyzer

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/AEKDA/aereq/internal/pkg/logger"
	"github.com/itchyny/gojq"
	"github.com/pkg/errors"
)

type analyzer struct{}

func (s *analyzer) Do(ctx context.Context, input string, query string) (string, error) {
	ctx = logger.WithName(ctx, "jsonanalyzer")

	queryJq, err := gojq.Parse(query)
	if err != nil {
		return "", errors.Wrap(err, "gojq.Parse")
	}

	var jsonInput interface{}
	err = json.Unmarshal([]byte(input), &jsonInput)
	if err != nil {
		return "", errors.Wrap(err, "json.Unmarshal")
	}

	iter := queryJq.RunWithContext(ctx, jsonInput)

	out := strings.Builder{}
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil {
				break
			}
			logger.Errorf(ctx, "Ошибка", err.Error())
			return "", err
		}
		err := json.NewEncoder(&out).Encode(v)
		if err != nil {
			return "", errors.Wrap(err, "json.Encode")
		}
	}

	res := out.String()

	return res, nil
}

func New() *analyzer {
	return &analyzer{}
}
