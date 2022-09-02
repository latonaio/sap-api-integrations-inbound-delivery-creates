package responses

type Item struct {
	D struct {
 			ActualDeliveredQtyInBaseUnit   string `json:"ActualDeliveredQtyInBaseUnit"`
			ActualDeliveryQuantity         string `json:"ActualDeliveryQuantity"`
			AdditionalCustomerGroup1       string `json:"AdditionalCustomerGroup1"`
			AdditionalCustomerGroup2       string `json:"AdditionalCustomerGroup2"`
			AdditionalCustomerGroup3       string `json:"AdditionalCustomerGroup3"`
			AdditionalCustomerGroup4       string `json:"AdditionalCustomerGroup4"`
			AdditionalCustomerGroup5       string `json:"AdditionalCustomerGroup5"`
			BaseUnit                       string `json:"BaseUnit"`
			Batch                          string `json:"Batch"`
			BatchBySupplier                string `json:"BatchBySupplier"`
			BOMExplosion                   string `json:"BOMExplosion"`
			BusinessArea                   string `json:"BusinessArea"`
			ControllingArea                string `json:"ControllingArea"`
			CostCenter                     string `json:"CostCenter"`
			CreationDate                   string `json:"CreationDate"`
			CreationTime                   string `json:"CreationTime"`
			DeliveryDocument               string `json:"DeliveryDocument"`
			DeliveryDocumentItem           string `json:"DeliveryDocumentItem"`
			DeliveryDocumentItemCategory   string `json:"DeliveryDocumentItemCategory"`
			DeliveryDocumentItemText       string `json:"DeliveryDocumentItemText"`
			DeliveryGroup                  string `json:"DeliveryGroup"`
			DeliveryQuantityUnit           string `json:"DeliveryQuantityUnit"`
			DeliveryRelatedBillingStatus   string `json:"DeliveryRelatedBillingStatus"`
			DistributionChannel            string `json:"DistributionChannel"`
			Division                       string `json:"Division"`
			GLAccount                      string `json:"GLAccount"`
			GoodsMovementReasonCode        string `json:"GoodsMovementReasonCode"`
			GoodsMovementStatus            string `json:"GoodsMovementStatus"`
			GoodsMovementType              string `json:"GoodsMovementType"`
			InternationalArticleNumber     string `json:"InternationalArticleNumber"`
			InventorySpecialStockType      string `json:"InventorySpecialStockType"`
			IsCompletelyDelivered          bool   `json:"IsCompletelyDelivered"`
			IsNotGoodsMovementsRelevant    string `json:"IsNotGoodsMovementsRelevant"`
			IssuingOrReceivingPlant        string `json:"IssuingOrReceivingPlant"`
			IssuingOrReceivingStorageLoc   string `json:"IssuingOrReceivingStorageLoc"`
			ItemBillingBlockReason         string `json:"ItemBillingBlockReason"`
			ItemBillingIncompletionStatus  string `json:"ItemBillingIncompletionStatus"`
			ItemDeliveryIncompletionStatus string `json:"ItemDeliveryIncompletionStatus"`
			ItemGdsMvtIncompletionSts      string `json:"ItemGdsMvtIncompletionSts"`
			ItemGeneralIncompletionStatus  string `json:"ItemGeneralIncompletionStatus"`
			ItemGrossWeight                string `json:"ItemGrossWeight"`
			ItemIsBillingRelevant          string `json:"ItemIsBillingRelevant"`
			ItemNetWeight                  string `json:"ItemNetWeight"`
			ItemPackingIncompletionStatus  string `json:"ItemPackingIncompletionStatus"`
			ItemPickingIncompletionStatus  string `json:"ItemPickingIncompletionStatus"`
			ItemVolume                     string `json:"ItemVolume"`
			ItemVolumeUnit                 string `json:"ItemVolumeUnit"`
			ItemWeightUnit                 string `json:"ItemWeightUnit"`
			LastChangeDate                 string `json:"LastChangeDate"`
			LoadingGroup                   string `json:"LoadingGroup"`
			Material                       string `json:"Material"`
			MaterialByCustomer             string `json:"MaterialByCustomer"`
			MaterialFreightGroup           string `json:"MaterialFreightGroup"`
			MaterialGroup                  string `json:"MaterialGroup"`
			MaterialIsBatchManaged         bool   `json:"MaterialIsBatchManaged"`
			OrderID                        string `json:"OrderID"`
			OrderItem                      string `json:"OrderItem"`
			OriginalDeliveryQuantity       string `json:"OriginalDeliveryQuantity"`
			PackingStatus                  string `json:"PackingStatus"`
			PartialDeliveryIsAllowed       string `json:"PartialDeliveryIsAllowed"`
			PickingConfirmationStatus      string `json:"PickingConfirmationStatus"`
			PickingStatus                  string `json:"PickingStatus"`
			Plant                          string `json:"Plant"`
			ProductAvailabilityDate        string `json:"ProductAvailabilityDate"`
			ProductAvailabilityTime        string `json:"ProductAvailabilityTime"`
			ProfitabilitySegment           string `json:"ProfitabilitySegment"`
			ProfitCenter                   string `json:"ProfitCenter"`
			QuantityIsFixed                bool   `json:"QuantityIsFixed"`
			ReceivingPoint                 string `json:"ReceivingPoint"`
			ReferenceSDDocument            string `json:"ReferenceSDDocument"`
			ReferenceSDDocumentItem        string `json:"ReferenceSDDocumentItem"`
			SalesDocumentItemType          string `json:"SalesDocumentItemType"`
			SalesGroup                     string `json:"SalesGroup"`
			SalesOffice                    string `json:"SalesOffice"`
			SDDocumentCategory             string `json:"SDDocumentCategory"`
			SDProcessStatus                string `json:"SDProcessStatus"`
			ShelfLifeExpirationDate        string `json:"ShelfLifeExpirationDate"`
			StockType                      string `json:"StockType"`
			StorageLocation                string `json:"StorageLocation"`
			TransportationGroup            string `json:"TransportationGroup"`
			UnlimitedOverdeliveryIsAllowed bool   `json:"UnlimitedOverdeliveryIsAllowed"`
	} `json:"d"`
}
