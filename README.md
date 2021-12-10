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

## 本レポジトリ が 対応する API サービス
sap-api-integrations-outbound-delivery-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_OUTBOUND_DELIVERY_SRV_0002/overview  
* APIサービス名(=baseURL): API_OUTBOUND_DELIVERY_SRV;v=0002

## 本レポジトリ に 含まれる API名
sap-api-integrations-outbound-delivery-reads には、次の API をコールするためのリソースが含まれています。  

* A_OutbDeliveryHeader（出荷伝票 - ヘッダ）
* A_OutbDeliveryHeader('{DeliveryDocument}')/to_DeliveryDocumentPartner（出荷伝票 - 取引先機能）
* A_OutbDeliveryPartner(PartnerFunction='{PartnerFunction}',SDDocument='{SDDocument}')/to_Address2（出荷伝票 - 取引先アドレス）
* A_OutbDeliveryItem（出荷伝票 - 明細）

## API への 値入力条件 の 初期値
sap-api-integrations-outbound-delivery-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.OutboundDelivery.DeliveryDocument（出荷伝票）
* inoutSDC.SDDocument（販売伝票 ※出荷伝票の取引先機能関連のAPIをコールするときに出荷伝票ではなく販売伝票としての項目値が必要です。通常は、出荷伝票の値＝販売伝票の値、となります）
* inoutSDC.OutboundDelivery.PartnerFunction.PartnerFunction（取引先機能）
* inoutSDC.OutboundDelivery.DeliveryDocumentItem.DeliveryDocumentItem（出荷伝票明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "sap.s4.beh.outbounddelivery.v1.OutboundDelivery.Created.v1",
	"accepter": ["Header"],
	"delivery_document": "80000000",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.outbounddelivery.v1.OutboundDelivery.Created.v1",
	"accepter": ["All"],
	"delivery_document": "80000000",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetOutboundDelivery(deliveryDocument, sDDocument, partnerFunction, deliveryDocumentItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(deliveryDocument)
				wg.Done()
			}()
		case "PartnerFunction":
			func() {
				c.PartnerFunction(sDDocument, partnerFunction)
				wg.Done()
			}()
		case "PartnerAddress":
			func() {
				c.PartnerAddress(partnerFunction, sDDocument)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(deliveryDocument, deliveryDocumentItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## SAP API Business Hub における API サービス の バージョン と バージョン におけるデータレイアウトの相違

SAP API Business Hub における API サービス のうちの 殆どの API サービス の バージョンは、v1.X.X であり、そのため殆どの API サービス 間 の データレイアウトは統一されています。
従って、Latona および AION における リソースにおいても、これらの 揃った v1.X.X のバージョンに従い、データレイアウトが統一されています。
一方、本レポジトリ に関わる API である Outbound Delivery のサービスは、v2.X.X までバージョンが進んでおり、その結果、本レポジトリ内の一部のAPIのデータレイアウトが、他のAPIサービスのものと異なっています。

* v1.X.X のバージョンである API サービス の データレイアウト（=responses）  
v1.X.X のバージョンであるAPIサービスのデータレイアウト（=responses）は、例えば、次の通りです。

```
type PartnerFunction struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SDDocument           string `json:"SDDocument"`
			PartnerFunction      string `json:"PartnerFunction"`
			Customer             string `json:"Customer"`
			Supplier             string `json:"Supplier"`
		} `json:"results"`
	} `json:"d"`
}


```

* 本レポジトリ内の一部のAPIのデータレイアウトを含む、v2.X.X のバージョンである API サービス の データレイアウト（=responses）  
本レポジトリ内の一部のAPIのデータレイアウトを含む、v2.X.X のバージョンである API サービスのデータレイアウト（=responses）は、次の通りです。  
ここで言う本レポジトリ内の一部のAPIのデータレイアウトには、PartnerAddressが含まれ、  
本レポジトリ内の他のデータレイアウトは、v1.X.X と同様のデータレイアウトとなっています。
（つまり、本レポジトリ内においても、現時点でv2.X.X のデータレイアウトを踏襲しているのは、PartnerAddress のみで、Header、PartnerFunction、Item、などのAPIでは、v1.X.X と同様のデータレイアウトとなっています）  

```
type PartnerAddress struct {
	D *struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			    Etag string `json:"etag"`
			} `json:"__metadata"`
			AddressID              string `json:"AddressID"`
			BusinessPartnerName1   string `json:"BusinessPartnerName1"`
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

このように、v1.X.X のバージョンである API サービス の データレイアウトと、v2.X.X のバージョンである API サービス の データレイアウトは、Results の配列構造を持っているか持っていないかという点が異なります。  