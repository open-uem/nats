package nats

import "time"

type OSUpdateSettingsType struct {
	AutoCheckEnabled                bool      `plist:"AutoCheckEnabled"`
	AutomaticAppInstallationEnabled bool      `plist:"AutomaticAppInstallationEnabled"`
	AutomaticOSInstallationEnabled  bool      `plist:"AutomaticOSInstallationEnabled"`
	AutomaticSecurityUpdatesEnabled bool      `plist:"AutomaticSecurityUpdatesEnabled"`
	BackgroundDownloadEnabled       bool      `plist:"BackgroundDownloadEnabled"`
	CatalogURL                      string    `plist:"CatalogURL"`
	IsDefaultCatalog                bool      `plist:"IsDefaultCatalog"`
	PreviousScanDate                time.Time `plist:"PreviousScanDate"`
	PreviousScanResult              int       `plist:"PreviousScanResult"`
}

type NanoMDMDeviceInfoResponse struct {
	ActiveManagedUsers               []string             `plist:"ActiveManagedUsers"`
	AutoSetupAdminAccounts           []string             `plist:"AutoSetupAdminAccounts"`
	AvailableDeviceCapacity          float64              `plist:"AvailableDeviceCapacity"`
	AwaitingConfiguration            bool                 `plist:"AwaitingConfiguration"`
	BatteryLevel                     float64              `plist:"BatteryLevel"`
	BluetoothMAC                     string               `plist:"BluetoothMAC"`
	BuildVersion                     string               `plist:"BuildVersion"`
	CurrentConsoleManagedUser        string               `plist:"CurrentConsoleManagedUser"`
	DeviceCapacity                   float64              `plist:"DeviceCapacity"`
	DeviceName                       string               `plist:"DeviceName"`
	EACSPreflight                    string               `plist:"EACSPreflight"`
	EthernetMAC                      string               `plist:"EthernetMAC"`
	HasBattery                       bool                 `plist:"HasBattery"`
	HostName                         string               `plist:"HostName"`
	IsActivationLockEnabled          bool                 `plist:"IsActivationLockEnabled"`
	IsActivationLockSupported        bool                 `plist:"IsActivationLockSupported"`
	IsAppleSilicon                   bool                 `plist:"IsAppleSilicon"`
	IsSupervised                     bool                 `plist:"IsSupervised"`
	LocalHostName                    string               `plist:"LocalHostName"`
	Model                            string               `plist:"Model"`
	ModelName                        string               `plist:"ModelName"`
	OSUpdateSettings                 OSUpdateSettingsType `plist:"OSUpdateSettings"`
	OSVersion                        string               `plist:"OSVersion"`
	OSXSoftwareUpdateStatus          OSUpdateSettingsType `plist:"OSXSoftwareUpdateStatus"`
	PINRequiredForDeviceLock         bool                 `plist:"PINRequiredForDeviceLock"`
	PINRequiredForEraseDevice        bool                 `plist:"PINRequiredForEraseDevice"`
	ProductName                      string               `plist:"ProductName"`
	ProvisioningUDID                 string               `plist:"ProvisioningUDID"`
	SerialNumber                     string               `plist:"SerialNumber"`
	SoftwareUpdateDeviceID           string               `plist:"SoftwareUpdateDeviceID"`
	SupplementalBuildVersion         string               `plist:"SupplementalBuildVersion"`
	SupportsLOMDevice                bool                 `plist:"SupportsLOMDevice"`
	SupportsiOSAppInstalls           bool                 `plist:"SupportsiOSAppInstalls"`
	SystemIntegrityProtectionEnabled bool                 `plist:"SystemIntegrityProtectionEnabled"`
	Status                           string               `plist:"Status"`
	UDID                             string               `plist:"UDID"`
	WifiMAC                          string               `plist:"WiFiMAC"`
}

type DeviceInformationResponse struct {
	CommandUUID    string                    `plist:"CommandUUID"`
	QueryResponses NanoMDMDeviceInfoResponse `plist:"QueryResponses"`
}

type NanoMDMUser struct {
	DataQuota      int    `plist:"DataQuota"`
	DataUsed       int    `plist:"DataUsed"`
	HasDataToSync  bool   `plist:"HasDataToSync"`
	HasSecureToken bool   `plist:"HasSecureToken"`
	IsLoggedIn     bool   `plist:"IsLoggedIn"`
	UserName       string `plist:"UserName"`
	FullName       string `plist:"FullName"`
	MobileAccount  bool   `plist:"MobileAccount"`
	UID            int    `plist:"UID"`
	UserGUID       int    `plist:"UserGUID"`
}

type UserListResponse struct {
	CommandUUID string        `plist:"CommandUUID"`
	Status      string        `plist:"Status"`
	UDID        string        `plist:"UDID"`
	Users       []NanoMDMUser `plist:"Users"`
}
