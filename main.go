package main

import (
	sap_api_caller "sap-api-integrations-outbound-delivery-reads/SAP_API_Caller"
	"sap-api-integrations-outbound-delivery-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Outbound_Delivery_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetOutboundDelivery(
		inoutSDC.OutboundDelivery.DeliveryDocument,
		inoutSDC.SDDocument,
		inoutSDC.OutboundDelivery.PartnerFunction.PartnerFunction,
		inoutSDC.OutboundDelivery.DeliveryDocumentItem.DeliveryDocumentItem,
	)
}
