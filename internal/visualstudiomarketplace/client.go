package visualstudiomarketplace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type VisualStudioMarketplaceClient struct {
	BaseUrl    string
	HttpClient *http.Client
}

func (e *VisualStudioMarketplaceClient) buildCriteriaForExtensions(extensions []string) string {
	var buffer bytes.Buffer

	for i := range extensions {
		if i > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(
			fmt.Sprintf(`{ "filterType": 7, "value": "%s" }`, extensions[i]),
		)
	}

	return buffer.String()
}

func (e *VisualStudioMarketplaceClient) GetStatistics(extensions []string) []VisualStudioMarketPlaceStatistic {
	var jsonStr = fmt.Sprintf(`{
		"assetTypes": null,
		"filters": [
				{
						"criteria": [%s],
						"direction": 2,
						"pageSize": %d,
						"pageNumber": 1,
						"sortBy": 0,
						"sortOrder": 0,
						"pagingToken": null
				}
		],
		"flags": 870
	}`, e.buildCriteriaForExtensions(extensions), len(extensions))

	req, err := http.NewRequest("POST", e.BaseUrl+"/_apis/public/gallery/extensionquery", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json;api-version=7.2-preview.1;excludeUrls=true")
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var parsedResponse VisualStudioMarketPlaceApiResponse
	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		log.Fatal(err)
	}

	statistics := make([]VisualStudioMarketPlaceStatistic, 0)

	for i := range parsedResponse.Results {
		result := parsedResponse.Results[i]

		for j := range result.Extensions {
			extension := result.Extensions[j]

			for k := range extension.Statistics {
				statistic := extension.Statistics[k]

				statistics = append(statistics, CreateNewVisualStudioMarketPlaceStatistic(extension, statistic))
			}
		}
	}

	return statistics
}

func CreateNewVisualStudioMarketPlaceStatistic(extension VisualStudioMarketPlaceApiResponseExtension, statistic VisualStudioMarketPlaceApiResponseExtensionStatistic) VisualStudioMarketPlaceStatistic {
	return VisualStudioMarketPlaceStatistic{
		ExtensionId:   extension.ExtensionId,
		ExtensionName: extension.ExtensionName,
		Name:          statistic.Name,
		Value:         statistic.Value,
	}
}
