package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
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

func (c *SAPAPICaller) AsyncGetOutboundDelivery(DeliveryDocument, PartnerFunction, DeliveryDocumentItem string) {
	wg := &sync.WaitGroup{}

	wg.Add(4)
	go func() {
		c.Header(DeliveryDocument)
		wg.Done()
	}()

	go func() {
		c.PartnerFunction(DeliveryDocument, PartnerFunction)
		wg.Done()
	}()
	
	go func() {
		c.PartnerAddress(DeliveryDocument, PartnerFunction)
		wg.Done()
	}()
	
	go func() {
		c.Item(DeliveryDocument, DeliveryDocumentItem)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Header(DeliveryDocument string) {
	res, err := c.callOutboundDeliverySrvAPIRequirementHeader("A_OutbDeliveryHeader", DeliveryDocument)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) PartnerFunction(DeliveryDocument, PartnerFunction string) {
	res, err := c.callOutboundDeliverySrvAPIRequirementPartnerFunction("A_OutbDeliveryHeader('{DeliveryDocument}')/to_DeliveryDocumentPartner", DeliveryDocument, PartnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}
func (c *SAPAPICaller) PartnerAddress(PartnerFunction, DeliveryDocument, AddressID string) {
	res, err := c.callOutboundDeliverySrvAPIRequirementPartnerAddress("/A_OutbDeliveryPartner(PartnerFunction='{PartnerFunction}',SDDocument='{SDDocument}')/to_Address", PartnerFunction, DeliveryDocument, AddressID)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Item(DeliveryDocument, DeliveryDocumentItem string) {
	res, err := c.callOutboundDeliverySrvAPIRequirementItem("A_OutbDeliveryItem", DeliveryDocument, DeliveryDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)


func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementHeader(api, DeliveryDocument string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "DeliveryDocument")
	params.Add("$filter", fmt.Sprintf("DeliveryDocument eq '%s'", DeliveryDocument))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementPartnerFunction(api, DeliveryDocument, PartnerFunction string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "DeliveryDocument, PartnerFunction")
	params.Add("$filter", fmt.Sprintf("DeliveryDocument eq '%s' and PartnerFunction eq 'SH'", DeliveryDocument, PartnerFunction))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementPartnerAddress(api, PartnerFunction, DeliveryDocument, AddressID string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "PartnerFunction, DeliveryDocument, AddressID")
	params.Add("$filter", fmt.Sprintf("PartnerFunction eq 'SH' and DeliveryDocument eq '%s' and AddressID eq '%s'", PartnerFunction, DeliveryDocument, AddressID))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementItem(api, DeliveryDocument, DeliveryDocumentItem string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "DeliveryDocument, DeliveryDocumentItem")
	params.Add("$filter", fmt.Sprintf("DeliveryDocument eq '%s' and DeliveryDocumentItem eq '%s'", DeliveryDocument, DeliveryDocumentItem))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}
