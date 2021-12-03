package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-outbound-delivery-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Header{
		DeliveryDocument:              data.DeliveryDocument,
		DeliveryDocumentType:          data.DeliveryDocumentType,
		DocumentDate:                  data.DocumentDate,
		ActualGoodsMovementDate:       data.ActualGoodsMovementDate,
		ActualDeliveryRoute:           data.ActualDeliveryRoute,
        Shippinglocationtimezone:      data.Shippinglocationtimezone,
        Receivinglocationtimezone:     data.Receivinglocationtimezone,
		ActualGoodsMovementTime:       data.ActualGoodsMovementTime,
		BillingDocumentDate:           data.BillingDocumentDate,
		CompleteDeliveryIsDefined:     data.CompleteDeliveryIsDefined,
		ConfirmationTime:              data.ConfirmationTime,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		CustomerGroup:                 data.CustomerGroup,
		DeliveryBlockReason:           data.DeliveryBlockReason,
		DeliveryDate:                  data.DeliveryDate,
		DeliveryDocumentBySupplier:    data.DeliveryDocumentBySupplier,
		DeliveryIsInPlant:             data.DeliveryIsInPlant,
		DeliveryPriority:              data.DeliveryPriority,
		DeliveryTime:                  data.DeliveryTime,
		GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
		GoodsIssueTime:                data.GoodsIssueTime,
		HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
		HeaderGrossWeight:             data.HeaderGrossWeight,
		HeaderNetWeight:               data.HeaderNetWeight,
		HeaderVolume:                  data.HeaderVolume,
		HeaderVolumeUnit:              data.HeaderVolumeUnit,
		HeaderWeightUnit:              data.HeaderWeightUnit,
		IncotermsClassification:       data.IncotermsClassification,
		IsExportDelivery:              data.IsExportDelivery,
		LastChangeDate:                data.LastChangeDate,
		LoadingDate:                   data.LoadingDate,
		LoadingPoint:                  data.LoadingPoint,
		LoadingTime:                   data.LoadingTime,
		MeansOfTransport:              data.MeansOfTransport,
		OrderCombinationIsAllowed:     data.OrderCombinationIsAllowed,
		OrderID:                       data.OrderID,
		OverallDelivConfStatus:        data.OverallDelivConfStatus,
		OverallDelivReltdBillgStatus:  data.OverallDelivReltdBillgStatus,
		OverallGoodsMovementStatus:    data.OverallGoodsMovementStatus,
		OverallPackingStatus:          data.OverallPackingStatus,
		OverallPickingConfStatus:      data.OverallPickingConfStatus,
		OverallPickingStatus:          data.OverallPickingStatus,
		PickingDate:                   data.PickingDate,
		PickingTime:                   data.PickingTime,
		PlannedGoodsIssueDate:         data.PlannedGoodsIssueDate,
		ReceivingPlant:                data.ReceivingPlant,
		ShippingCondition:             data.ShippingCondition,
		ShippingPoint:                 data.ShippingPoint,
		ShippingType:                  data.ShippingType,
		ShipToParty:                   data.ShipToParty,
		SoldToParty:                   data.SoldToParty,
		Supplier:                      data.Supplier,
		TransportationGroup:           data.TransportationGroup,
		TransportationPlanningDate:    data.TransportationPlanningDate,
		TransportationPlanningTime:    data.TransportationPlanningTime,
	}, nil
}

func ConvertToPartnerFunction(raw []byte, l *logger.Logger) (*PartnerFunction, error) {
	pm := &responses.PartnerFunction{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerFunction. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PartnerFunction{
		SDDocument:                    data.SDDocument,
        PartnerFunction:               data.PartnerFunction,
		Customer:                      data.Customer,
		Supplier:                      data.Supplier,
	}, nil
}

func ConvertToPartnerAddress(raw []byte, l *logger.Logger) (*PartnerAddress, error) {
	pm := &responses.PartnerAddress{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerAddress. unmarshal error: %w", err)
	}
	if pm.D == nil {
		return nil, xerrors.New("Metadata data is not exist")
	}
	data := pm.D

	return &PartnerAddress{
        AddressID:                     data.AddressID, 
        BusinessPartnerName1:          data.BusinessPartnerName1,
        Country:                       data.Country,
		Region:                        data.Region,
		StreetName:                    data.StreetName,
		CityName:                      data.CityName,
		PostalCode:                    data.PostalCode,
		CorrespondenceLanguage:        data.CorrespondenceLanguage,
		FaxNumber:                     data.FaxNumber,
		PhoneNumber:                   data.PhoneNumber,
	}, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Item{
		DeliveryDocument:              data.DeliveryDocument,
        DeliveryDocumentItem:          data.DeliveryDocumentItem,
		BaseUnit:                      data.BaseUnit,
		ActualDeliveryQuantity:        data.ActualDeliveryQuantity,
		Batch:                         data.Batch,
		BatchBySupplier:               data.BatchBySupplier,
		CostCenter:                    data.CostCenter,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		DeliveryDocumentItemCategory:  data.DeliveryDocumentItemCategory,
		DeliveryDocumentItemText:      data.DeliveryDocumentItemText,
		DeliveryQuantityUnit:          data.DeliveryQuantityUnit,
		DistributionChannel:           data.DistributionChannel,
		Division:                      data.Division,
		GLAccount:                     data.GLAccount,
		GoodsMovementStatus:           data.GoodsMovementStatus,
		GoodsMovementType:             data.GoodsMovementType,
		InternationalArticleNumber:    data.InternationalArticleNumber,
		InventorySpecialStockType:     data.InventorySpecialStockType,
		IsCompletelyDelivered:         data.IsCompletelyDelivered,
		ItemBillingBlockReason:        data.ItemBillingBlockReason,
		ItemDeliveryIncompletionStatus: data.ItemDeliveryIncompletionStatus,
		ItemGdsMvtIncompletionSts:     data.ItemGdsMvtIncompletionSts,
		ItemGrossWeight:               data.ItemGrossWeight,
		ItemNetWeight:                 data.ItemNetWeight,
		ItemWeightUnit:                data.ItemWeightUnit,
		ItemIsBillingRelevant:         data.ItemIsBillingRelevant,
		ItemPackingIncompletionStatus: data.ItemPackingIncompletionStatus,
		ItemPickingIncompletionStatus: data.ItemPickingIncompletionStatus,
		ItemVolume:                    data.ItemVolume,
		LastChangeDate:                data.LastChangeDate,
		LoadingGroup:                  data.LoadingGroup,
		Material:                      data.Material,
		MaterialByCustomer:            data.MaterialByCustomer,
		MaterialFreightGroup:          data.MaterialFreightGroup,
		NumberOfSerialNumbers:         data.NumberOfSerialNumbers,
		OrderID:                       data.OrderID,
		OrderItem:                     data.OrderItem,
		OriginalDeliveryQuantity:      data.OriginalDeliveryQuantity,
		PackingStatus:                 data.PackingStatus,
		PartialDeliveryIsAllowed:      data.PartialDeliveryIsAllowed,
		PickingConfirmationStatus:     data.PickingConfirmationStatus,
		PickingStatus:                 data.PickingStatus,
		Plant:                         data.Plant,
		ProductAvailabilityDate:       data.ProductAvailabilityDate,
		ProductAvailabilityTime:       data.ProductAvailabilityTime,
		ProfitCenter:                  data.ProfitCenter,
		StorageLocation:               data.StorageLocation,
		TransportationGroup:           data.TransportationGroup,
	}, nil
}
