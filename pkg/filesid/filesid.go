package filesid

import (
	"io/ioutil"
	"net/http"

	"github.com/hsmtkk/virustotal_go/pkg"
	"github.com/hsmtkk/virustotal_go/pkg/filesid/parser"
)

type InformationRetriever interface {
	RetrieveInformation(hash string) (pkg.LastAnalysisStats, error)
}

func New(apiKey string) InformationRetriever {
	endPoint := "https://www.virustotal.com/api/v3/files"
	return &informationRetrieverImpl{client: http.DefaultClient, endPoint: endPoint, apiKey: apiKey}
}

func NewForTest(client *http.Client, endPoint string) InformationRetriever {
	return &informationRetrieverImpl{client: client, endPoint: endPoint, apiKey: "dummy"}
}

type informationRetrieverImpl struct {
	client   *http.Client
	endPoint string
	apiKey   string
}

func (imp *informationRetrieverImpl) RetrieveInformation(hash string) (pkg.LastAnalysisStats, error) {
	url := imp.endPoint + "/" + hash
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return pkg.LastAnalysisStats{}, err
	}
	req.Header.Set(pkg.XAPIKey, imp.apiKey)
	resp, err := imp.client.Do(req)
	if err != nil {
		return pkg.LastAnalysisStats{}, err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pkg.LastAnalysisStats{}, err
	}
	ps := parser.New()
	stats, err := ps.ParseResponse(string(bs))
	if err != nil {
		return pkg.LastAnalysisStats{}, err
	}
	return stats, nil
}
