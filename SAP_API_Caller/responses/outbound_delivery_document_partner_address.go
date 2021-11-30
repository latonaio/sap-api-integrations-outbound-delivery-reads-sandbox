package responses

type OutboundDeliveryDocumentPartnerAddress struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			DeliveryDocument       string `json:"DeliveryDocument"`
			PartnerFunction        string `json:"PartnerFunction"`
			AddressID              string `json:"AddressID"`
			Country                string `json:"Country"`
			Region                 string `json:"Region"`
			StreetName             string `json:"StreetName"`
			CityName               string `json:"CityName"`
			PostalCode             string `json:"PostalCode"`
			CorrespondenceLanguage string `json:"CorrespondenceLanguage"`
			EmailAddress           string `json:"EmailAddress"`
			FaxNumber              string `json:"FaxNumber"`
			MobilePhoneNumber      string `json:"MobilePhoneNumber"`
			PhoneNumber            string `json:"PhoneNumber"`
		} `json:"results"`
	} `json:"d"`
}
