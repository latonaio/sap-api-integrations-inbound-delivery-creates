# sap-api-integrations-creates  
sap-api-integrations-inbound-delivery-creates は、他のすべての sap-api-integrations-inbound-delivery-creates 作成更新の際の 参照元となる マスタレポジトリです。  
sap-api-integrations-inbound-delivery-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 入荷データ を登録するマイクロサービスです。  
sap-api-integrations-inbound-delivery-creates には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-inbound-delivery-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_INBOUND_DELIVERY/overview  

## 動作環境  
sap-api-integrations-inbound-delivery-creates は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-inbound-delivery-creates は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。


## 本レポジトリ が 対応する API サービス
sap-api-integrations-inbound-delivery-creates が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_INBOUND_DELIVERY/overview  
* APIサービス名(=baseURL): API_INBOUND_DELIVERY_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-inbound-delivery-creates には、次の API をコールするためのリソースが含まれています。  

* A_InboundDelivery（入荷 - ヘッダ）※入荷の詳細データを取得するために、ToAdress、ToItem、ToPartnerと合わせて利用されます。
* ToXXXXX（入荷 - 明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
"api_schema": "SAPInboundDeliveryCreates",
"accepter": ["Header"],
"inbound-delivery": "",
"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
"api_schema": "SAPInboundDeliveryCreates",
"accepter": ["All"],
"inbound-delivery": "",
"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostInboundDelivery(
	header            *requests.Header,
	item              *requests.Item,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(header)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(item)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 入荷 の ヘッダデータ が登録された結果の JSON の例です。  
以下の項目のうち、"ReceivingLocationTimeZone" ～ "UnloadingPointName" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/home/ampamman/go/src/sap-api-integrations-inbound-delivery-creates/SAP_API_Caller/caller.go#L66",
    "function": "sap-api-integrations-inbound-delivery-creates/SAP_API_Caller.(*SAPAPICaller).Header",
    "level": "INFO",
    "message": {
        "ReceivingLocationTimeZone": "",
		"ActualDeliveryRoute": "",
		"ActualGoodsMovementDate": "",
		"ActualGoodsMovementTime": "",
		"BillingDocumentDate": "",
		"CompleteDeliveryIsDefined": null,
		"ConfirmationTime": "",
		"CreationDate": "",
		"CreationTime": "",
		"CustomerGroup": "",
		"DeliveryBlockReason": "",
		"DeliveryDate": "",
		"DeliveryDocument": "",
		"DeliveryDocumentBySupplier": "",
		"DeliveryDocumentType": "",
		"DeliveryIsInPlant": null,
		"DeliveryPriority": "",
		"DeliveryTime": "",
		"DocumentDate": "",
		"GoodsIssueOrReceiptSlipNumber": "",
		"GoodsIssueTime": "",
		"HeaderBillgIncompletionStatus": "",
		"HeaderBillingBlockReason": "",
		"HeaderDelivIncompletionStatus": "",
		"HeaderGrossWeight": "",
		"HeaderNetWeight": "",
		"HeaderPackingIncompletionSts": "",
		"HeaderPickgIncompletionStatus": "",
		"HeaderVolume": "",
		"HeaderVolumeUnit": "",
		"HeaderWeightUnit": "",
		"IncotermsClassification": "",
		"IsExportDelivery": "",
		"LastChangeDate": "",
		"LoadingDate": "",
		"LoadingPoint": "",
		"LoadingTime": "",
		"MeansOfTransport": "",
		"OrderCombinationIsAllowed": null,
		"OrderID": "",
		"PickedItemsLocation": "",
		"PickingDate": "",
		"PickingTime": "",
		"PlannedGoodsIssueDate": "",
		"ProposedDeliveryRoute": "",
		"ReceivingPlant": "",
		"RouteSchedule": "",
		"SalesDistrict": "",
		"SalesOffice": "",
		"SalesOrganization": "",
		"SDDocumentCategory": "",
		"ShipmentBlockReason": "",
		"ShippingCondition": "",
		"ShippingPoint": "",
		"ShippingType": "",
		"ShipToParty": "",
		"SoldToParty": "",
		"Supplier": "",
		"TotalBlockStatus": "",
		"TotalCreditCheckStatus": "",
		"TotalNumberOfPackage": "",
		"TransactionCurrency": "",
		"TransportationGroup": "",
		"TransportationPlanningDate": "",
		"TransportationPlanningStatus": "",
		"TransportationPlanningTime": "",
		"UnloadingPointName": "",
    },
    "time": "2022-09-02T12:06:08.872009881+09:00"
}
```