package file_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	DeliveryDocument   struct {
		DeliveryDocument               string      `json:"document_no"`
		ShipToParty                    string      `json:"deliver_to"`
		OriginalDeliveryQuantity       float64     `json:"quantity"`
		ActualDeliveryQuantity         float64     `json:"picked_quantity"`
		Price                          float64     `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             float64     `json:"quantity"`
		CompletedQuantity    float64     `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 float64     `json:"quantity"`
			CompletedQuantity        float64     `json:"completed_quantity"`
			ErroredQuantity          float64     `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity float64     `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema               string      `json:"api_schema"`
	Material                string      `json:"material_code"`
	ShippingPoint           string      `json:"plant/supplier"`
	Stock                   float64     `json:"stock"`
	DeliveryDocumentType    string      `json:"document_type"`
	DeliveryDocument        string      `json:"document_no"`
	PlannedGoodsIssueDate   string      `json:"planned_date"`
	ActualGoodsMovementDate string      `json:"validated_date"`
	Deleted                 string      `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string `json:"connection_key"`
	Result           bool   `json:"result"`
	RedisKey         string `json:"redis_key"`
	Filepath         string `json:"filepath"`
	DeliveryDocument struct {
		DeliveryDocument              string `json:"DeliveryDocument"`
		DeliveryDocumentType          string `json:"DeliveryDocumentType"`
		DocumentDate                  string `json:"DocumentDate"`
		ActualGoodsMovementDate       string `json:"ActualGoodsMovementDate"`
		ActualDeliveryRoute           string `json:"ActualDeliveryRoute"`
		Shippinglocationtimezone      string `json:"Shippinglocationtimezone"`
		Receivinglocationtimezone     string `json:"Receivinglocationtimezone"`
		ActualGoodsMovementTime       string `json:"ActualGoodsMovementTime"`
		BillingDocumentDate           string `json:"BillingDocumentDate"`
		CompleteDeliveryIsDefined     string `json:"CompleteDeliveryIsDefined"`
		ConfirmationTime              string `json:"ConfirmationTime"`
		CreationDate                  string `json:"CreationDate"`
		CreationTime                  string `json:"CreationTime"`
		CustomerGroup                 string `json:"CustomerGroup"`
		DeliveryBlockReason           string `json:"DeliveryBlockReason"`
		DeliveryDate                  string `json:"DeliveryDate"`
		DeliveryDocumentBySupplier    string `json:"DeliveryDocumentBySupplier"`
		DeliveryIsInPlant             string `json:"DeliveryIsInPlant"`
		DeliveryPriority              string `json:"DeliveryPriority"`
		DeliveryTime                  string `json:"DeliveryTime"`
		GoodsIssueOrReceiptSlipNumber string `json:"GoodsIssueOrReceiptSlipNumber"`
		GoodsIssueTime                string `json:"GoodsIssueTime"`
		HeaderBillingBlockReason      string `json:"HeaderBillingBlockReason"`
		HeaderGrossWeight             float64 `json:"HeaderGrossWeight"`
		HeaderNetWeight               float64 `json:"HeaderNetWeight"`
		HeaderVolume                  float64 `json:"HeaderVolume"`
		HeaderVolumeUnit              string `json:"HeaderVolumeUnit"`
		HeaderWeightUnit              string `json:"HeaderWeightUnit"`
		IncotermsClassification       string `json:"IncotermsClassification"`
		IsExportDelivery              string `json:"IsExportDelivery"`
		LastChangeDate                string `json:"LastChangeDate"`
		LoadingDate                   string `json:"LoadingDate"`
		LoadingPoint                  string `json:"LoadingPoint"`
		LoadingTime                   string `json:"LoadingTime"`
		MeansOfTransport              string `json:"MeansOfTransport"`
		OrderCombinationIsAllowed     string `json:"OrderCombinationIsAllowed"`
		OrderID                       string `json:"OrderID"`
		OverallDelivConfStatus        string `json:"OverallDelivConfStatus"`
		OverallDelivReltdBillgStatus  string `json:"OverallDelivReltdBillgStatus"`
		OverallGoodsMovementStatus    string `json:"OverallGoodsMovementStatus"`
		OverallPackingStatus          string `json:"OverallPackingStatus"`
		OverallPickingConfStatus      string `json:"OverallPickingConfStatus"`
		OverallPickingStatus          string `json:"OverallPickingStatus"`
		PickingDate                   string `json:"PickingDate"`
		PickingTime                   string `json:"PickingTime"`
		PlannedGoodsIssueDate         string `json:"PlannedGoodsIssueDate"`
		ReceivingPlant                string `json:"ReceivingPlant"`
		ShippingCondition             string `json:"ShippingCondition"`
		ShippingPoint                 string `json:"ShippingPoint"`
		ShippingType                  string `json:"ShippingType"`
		ShipToParty                   string `json:"ShipToParty"`
		SoldToParty                   string `json:"SoldToParty"`
		Supplier                      string `json:"Supplier"`
		TransportationGroup           string `json:"TransportationGroup"`
		TransportationPlanningDate    string `json:"TransportationPlanningDate"`
		TransportationPlanningTime    string `json:"TransportationPlanningTime"`
		DeliveryDocumentItem          struct {
			DeliveryDocumentItem           int         `json:"DeliveryDocumentItem"`
			BaseUnit                       string      `json:"BaseUnit"`
			ActualDeliveryQuantity         float64     `json:"ActualDeliveryQuantity"`
			Batch                          string      `json:"Batch"`
			BatchBySupplier                string      `json:"BatchBySupplier"`
			CostCenter                     string      `json:"CostCenter"`
			CreationDate                   string      `json:"CreationDate"`
			CreationTime                   string      `json:"CreationTime"`
			DeliveryDocumentItemCategory   string      `json:"DeliveryDocumentItemCategory"`
			DeliveryDocumentItemText       string      `json:"DeliveryDocumentItemText"`
			DeliveryQuantityUnit           string      `json:"DeliveryQuantityUnit"`
			DistributionChannel            string      `json:"DistributionChannel"`
			Division                       string      `json:"Division"`
			GLAccount                      string      `json:"GLAccount"`
			GoodsMovementStatus            string      `json:"GoodsMovementStatus"`
			GoodsMovementType              string      `json:"GoodsMovementType"`
			InternationalArticleNumber     string      `json:"InternationalArticleNumber"`
			InventorySpecialStockType      string      `json:"InventorySpecialStockType"`
			IsCompletelyDelivered          string      `json:"IsCompletelyDelivered"`
			ItemBillingBlockReason         string      `json:"ItemBillingBlockReason"`
			ItemDeliveryIncompletionStatus string      `json:"ItemDeliveryIncompletionStatus"`
			ItemGdsMvtIncompletionSts      string      `json:"ItemGdsMvtIncompletionSts"`
			ItemGrossWeight                float64     `json:"ItemGrossWeight"`
			ItemNetWeight                  float64     `json:"ItemNetWeight"`
			ItemWeightUnit                 string      `json:"ItemWeightUnit"`
			ItemIsBillingRelevant          string      `json:"ItemIsBillingRelevant"`
			ItemPackingIncompletionStatus  string      `json:"ItemPackingIncompletionStatus"`
			ItemPickingIncompletionStatus  string      `json:"ItemPickingIncompletionStatus"`
			ItemVolume                     string      `json:"ItemVolume"`
			LastChangeDate                 string      `json:"LastChangeDate"`
			LoadingGroup                   string      `json:"LoadingGroup"`
			Material                       string      `json:"Material"`
			MaterialByCustomer             string      `json:"MaterialByCustomer"`
			MaterialFreightGroup           string      `json:"MaterialFreightGroup"`
			NumberOfSerialNumbers          string      `json:"NumberOfSerialNumbers"`
			OrderID                        string      `json:"OrderID"`
			OrderItem                      int         `json:"OrderItem"`
			OriginalDeliveryQuantity       float64     `json:"OriginalDeliveryQuantity"`
			PackingStatus                  string      `json:"PackingStatus"`
			PartialDeliveryIsAllowed       string      `json:"PartialDeliveryIsAllowed"`
			PickingConfirmationStatus      string      `json:"PickingConfirmationStatus"`
			PickingStatus                  string      `json:"PickingStatus"`
			Plant                          string      `json:"Plant"`
			ProductAvailabilityDate        string      `json:"ProductAvailabilityDate"`
			ProductAvailabilityTime        string      `json:"ProductAvailabilityTime"`
			ProfitCenter                   string      `json:"ProfitCenter"`
			StorageLocation                string      `json:"StorageLocation"`
			TransportationGroup            string      `json:"TransportationGroup"`
		} `json:"DeliveryDocumentItem"`
		PartnerFunction struct {
			PartnerFunction        string `json:"PartnerFunction"`
			AddressID              string `json:"AddressID"`
			Customer               string `json:"Customer"`
			Supplier               string `json:"Supplier"`
			BusinessPartnerName1   string `json:"BusinessPartnerName1"`
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
		} `json:"PartnerFunction"`
	} `json:"DeliveryDocument"`
	APISchema        string `json:"api_schema"`
	DeliveryDocument string `json:"delivery_document"`
	Deleted          string `json:"deleted"`
}