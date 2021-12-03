# sap-api-integrations-outbound-delivery-reads 
sap-api-integrations-outbound-delivery-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で出荷データ を取得するマイクロサービスです。    
sap-api-integrations-outbound-delivery-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-outbound-delivery-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_OUTBOUND_DELIVERY_SRV_0002/overview  

## 動作環境  

sap-api-integrations-outbound-delivery-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用

sap-api-integrations-outbound-delivery-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。 

## SAP_API_Caller/responses/outbound_delivery_document_partner_address.goの形式

SAP_API_Caller/responses/outbound_delivery_document_partner_address.goは、SAP　APIのバージョンが２のため形式が異なります。
```
type PartnerAddress struct {
	D struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			    Etag string `json:"etag"`
			} `json:"__metadata"`
			AddressID              string `json:"AddressID"`
			Country                string `json:"Country"`
			Region                 string `json:"Region"`
			StreetName             string `json:"StreetName"`
			CityName               string `json:"CityName"`
			PostalCode             string `json:"PostalCode"`
			CorrespondenceLanguage string `json:"CorrespondenceLanguage"`
			FaxNumber              string `json:"FaxNumber"`
			PhoneNumber            string `json:"PhoneNumber"`
	} `json:"d"`
}
```
SAP_API_Caller/caller.goにおいての該当箇所は下記のとおりです。
```
func (c *SAPAPICaller) PartnerAddress(partnerFunction, sDDocument string) {
	data, err := c.callOutboundDeliverySrvAPIRequirementPartnerAddress(fmt.Sprintf("A_OutbDeliveryPartner(PartnerFunction='%s',SDDocument='%s')/to_Address2", partnerFunction, sDDocument), partnerFunction, sDDocument)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}
```