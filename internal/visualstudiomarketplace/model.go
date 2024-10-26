package visualstudiomarketplace

type VisualStudioMarketPlaceApiResponse struct {
	Results []VisualStudioMarketPlaceApiResponseResult `json:"results"`
}

type VisualStudioMarketPlaceApiResponseResult struct {
	Extensions []VisualStudioMarketPlaceApiResponseExtension `json:"extensions"`
}

type VisualStudioMarketPlaceApiResponseExtension struct {
	ExtensionId   string                                                 `json:"extensionId"`
	ExtensionName string                                                 `json:"extensionName"`
	DisplayName   string                                                 `json:"displayName"`
	Statistics    []VisualStudioMarketPlaceApiResponseExtensionStatistic `json:"statistics"`
}

type VisualStudioMarketPlaceApiResponseExtensionStatistic struct {
	Name  string  `json:"statisticName"`
	Value float64 `json:"value"`
}

type VisualStudioMarketPlaceStatistic struct {
	ExtensionId   string
	ExtensionName string
	Name          string
	Value         float64
}
