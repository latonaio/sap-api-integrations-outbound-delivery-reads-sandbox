package main

import (
	sap_api_caller "/SAP_API_Caller"
	"sap-api-integrations-outbound-delivery-reads/file_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := file_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Outbound_Delivery_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	caller.AsyncGetOutboundDelivery(
		inoutSDC.DeliveryDocument,
		inoutSDC.DeliveryDocument.PartnerFunction,
		inoutSDC.DeliveryDocument.DeliveryDocumentItem,
	)
}