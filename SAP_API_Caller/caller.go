package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-outbound-delivery-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

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

func (c *SAPAPICaller) Header(deliveryDocument string) {
	data, err := c.callOutboundDeliverySrvAPIRequirementHeader("A_OutbDeliveryHeader", deliveryDocument)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementHeader(api, deliveryDocument string) (*sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV;v=0002", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, deliveryDocument)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PartnerFunction(sDDocument, partnerFunction string) {
	data, err := c.callOutboundDeliverySrvAPIRequirementPartnerFunction(fmt.Sprintf("A_OutbDeliveryHeader('%s')/to_DeliveryDocumentPartner", sDDocument), sDDocument, partnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementPartnerFunction(api, sDDocument, partnerFunction string) (*sap_api_output_formatter.PartnerFunction, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV;v=0002", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPartnerFunction(req, sDDocument, partnerFunction)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerFunction(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PartnerAddress(partnerFunction, sDDocument string) {
	data, err := c.callOutboundDeliverySrvAPIRequirementPartnerAddress(fmt.Sprintf("A_OutbDeliveryPartner(PartnerFunction='%s',SDDocument='%s')/to_Address2", partnerFunction, sDDocument), partnerFunction, sDDocument)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementPartnerAddress(api, partnerFunction, sDDocument string) (*sap_api_output_formatter.PartnerAddress, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV;v=0002", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	//	c.getQueryWithPartnerAddress(req, partnerFunction, sDDocument)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(deliveryDocument, deliveryDocumentItem string) {
	data, err := c.callOutboundDeliverySrvAPIRequirementItem("A_OutbDeliveryItem", deliveryDocument, deliveryDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementItem(api, deliveryDocument, deliveryDocumentItem string) (*sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV;v=0002", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, deliveryDocument, deliveryDocumentItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, deliveryDocument string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("DeliveryDocument eq '%s'", deliveryDocument))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPartnerFunction(req *http.Request, sDDocument, partnerFunction string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("SDDocument eq '%s' and PartnerFunction eq '%s'", sDDocument, partnerFunction))
	req.URL.RawQuery = params.Encode()
}

//func (c *SAPAPICaller) getQueryWithPartnerAddress(req *http.Request, partnerFunction, sDDocument string) {
//	params := req.URL.Query()
//	params.Add("$filter", fmt.Sprintf("PartnerFunction eq '%s' and SDDocument eq '%s'", partnerFunction, sDDocument))
//	req.URL.RawQuery = params.Encode()
//}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, deliveryDocument, deliveryDocumentItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("DeliveryDocument eq '%s' and DeliveryDocumentItem eq '%s'", deliveryDocument, deliveryDocumentItem))
	req.URL.RawQuery = params.Encode()
}
