package parser

import (
	"encoding/json"

	"github.com/hsmtkk/virustotal_go/pkg"
)

type ResponseParser interface {
	ParseResponse(js string) (pkg.LastAnalysisStats, error)
}

func New() ResponseParser {
	return &responseParserImpl{}
}

type responseSchema struct {
	Data struct {
		Attributes struct {
			LastAnalysisStats pkg.LastAnalysisStats `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}

type responseParserImpl struct{}

func (imp *responseParserImpl) ParseResponse(js string) (pkg.LastAnalysisStats, error) {
	rs := responseSchema{}
	err := json.Unmarshal([]byte(js), &rs)
	if err != nil {
		return pkg.LastAnalysisStats{}, err
	}
	return rs.Data.Attributes.LastAnalysisStats, nil
}
