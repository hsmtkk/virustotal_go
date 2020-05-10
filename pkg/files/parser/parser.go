package parser

import (
	"encoding/json"
)

type ResponseParser interface {
	ParseResponse(js string) (string, error)
}

func New() ResponseParser {
	return &responseParserImpl{}
}

type responseSchema struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

type responseParserImpl struct{}

func (imp *responseParserImpl) ParseResponse(js string) (string, error) {
	rs := responseSchema{}
	err := json.Unmarshal([]byte(js), &rs)
	if err != nil {
		return "", err
	}
	return rs.Data.ID, nil
}
