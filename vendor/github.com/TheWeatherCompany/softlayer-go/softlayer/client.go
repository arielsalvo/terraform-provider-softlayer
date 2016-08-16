package softlayer

import (
	"bytes"
)

type Client interface {
	GetService(name string) (Service, error)

	GetSoftLayer_Account_Service() (SoftLayer_Account_Service, error)
	GetSoftLayer_User_Customer_Service() (SoftLayer_User_Customer_Service, error)
	GetSoftLayer_Virtual_Guest_Service() (SoftLayer_Virtual_Guest_Service, error)
	GetSoftLayer_Virtual_Disk_Image_Service() (SoftLayer_Virtual_Disk_Image_Service, error)
	GetSoftLayer_Security_Ssh_Key_Service() (SoftLayer_Security_Ssh_Key_Service, error)
	GetSoftLayer_Product_Order_Service() (SoftLayer_Product_Order_Service, error)
	GetSoftLayer_Product_Package_Service() (SoftLayer_Product_Package_Service, error)
	GetSoftLayer_Network_Storage_Service() (SoftLayer_Network_Storage_Service, error)
	GetSoftLayer_Network_Storage_Allowed_Host_Service() (SoftLayer_Network_Storage_Allowed_Host_Service, error)
	GetSoftLayer_Billing_Item_Cancellation_Request_Service() (SoftLayer_Billing_Item_Cancellation_Request_Service, error)
	GetSoftLayer_Billing_Item_Service() (SoftLayer_Billing_Item_Service, error)
	GetSoftLayer_Virtual_Guest_Block_Device_Template_Group_Service() (SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service, error)
	GetSoftLayer_Hardware_Service() (SoftLayer_Hardware_Service, error)
	GetSoftLayer_Dns_Domain_Service() (SoftLayer_Dns_Domain_Service, error)
	GetSoftLayer_Network_Application_Delivery_Controller_Service() (SoftLayer_Network_Application_Delivery_Controller_Service, error)
	GetSoftLayer_Dns_Domain_ResourceRecord_Service() (SoftLayer_Dns_Domain_ResourceRecord_Service, error)
	GetSoftLayer_Security_Certificate_Service() (SoftLayer_Security_Certificate_Service, error)
	GetSoftLayer_Scale_Group_Service() (SoftLayer_Scale_Group_Service, error)
	GetSoftLayer_Scale_Network_Vlan_Service() (SoftLayer_Scale_Network_Vlan_Service, error)
	GetSoftLayer_Scale_Policy_Service() (SoftLayer_Scale_Policy_Service, error)
	GetSoftLayer_Scale_Policy_Trigger_Service() (SoftLayer_Scale_Policy_Trigger_Service, error)
	GetSoftLayer_Load_Balancer_Service() (SoftLayer_Load_Balancer_Service, error)
	GetSoftLayer_Load_Balancer_Service_Group_Service() (SoftLayer_Load_Balancer_Service_Group_Service, error)
	GetSoftLayer_Provisioning_Hook_Service() (SoftLayer_Provisioning_Hook_Service, error)

	GetHttpClient() HttpClient
}

type HttpClient interface {
	DoRawHttpRequest(path string, requestType string, requestBody *bytes.Buffer) ([]byte, int, error)
	DoRawHttpRequestWithObjectMask(path string, masks []string, requestType string, requestBody *bytes.Buffer) ([]byte, int, error)
	DoRawHttpRequestWithObjectFilter(path string, filters string, requestType string, requestBody *bytes.Buffer) ([]byte, int, error)
	DoRawHttpRequestWithObjectFilterAndObjectMask(path string, masks []string, filters string, requestType string, requestBody *bytes.Buffer) ([]byte, int, error)
	GenerateRequestBody(templateData interface{}) (*bytes.Buffer, error)
	HasErrors(body map[string]interface{}) error

	CheckForHttpResponseErrors(data []byte) error
}