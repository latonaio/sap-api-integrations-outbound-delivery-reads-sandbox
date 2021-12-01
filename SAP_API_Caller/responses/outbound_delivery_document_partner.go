package responses

type PartnerFunction struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			DeliveryDocument     string `json:"DeliveryDocument"`
			PartnerFunction      string `json:"PartnerFunction"`
			Customer             string `json:"Customer"`
			Supplier             string `json:"Supplier"`
			BusinessPartnerName1 string `json:"BusinessPartnerName1"`
		} `json:"results"`
	} `json:"d"`
}
