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
	ActiveManagedUsers               []string             `plist:"ActiveManagedUsers" json:"activeManagedUsers"`
	AutoSetupAdminAccounts           []string             `plist:"AutoSetupAdminAccounts" json:"autoSetupAdminAccounts"`
	AvailableDeviceCapacity          float64              `plist:"AvailableDeviceCapacity" json:"availableDeviceCapacity"`
	AwaitingConfiguration            bool                 `plist:"AwaitingConfiguration" json:"awaitingConfiguration"`
	BatteryLevel                     float64              `plist:"BatteryLevel" json:"batteryLevel"`
	BluetoothMAC                     string               `plist:"BluetoothMAC" json:"bluetoothMAC"`
	BuildVersion                     string               `plist:"BuildVersion" json:"buildVersion"`
	CurrentConsoleManagedUser        string               `plist:"CurrentConsoleManagedUser" json:"currentConsoleManagedUser"`
	DeviceCapacity                   float64              `plist:"DeviceCapacity" json:"deviceCapacity"`
	DeviceName                       string               `plist:"DeviceName" json:"deviceName"`
	EACSPreflight                    string               `plist:"EACSPreflight" json:"eacsPreflight"`
	EthernetMAC                      string               `plist:"EthernetMAC" json:"ethernetMAC"`
	HasBattery                       bool                 `plist:"HasBattery" json:"hasBattery"`
	HostName                         string               `plist:"HostName" json:"hostname"`
	IsActivationLockEnabled          bool                 `plist:"IsActivationLockEnabled" json:"isActivationLockEnabled"`
	IsActivationLockSupported        bool                 `plist:"IsActivationLockSupported" json:"isActivationLockSupported"`
	IsAppleSilicon                   bool                 `plist:"IsAppleSilicon" json:"isAppleSilicon"`
	IsSupervised                     bool                 `plist:"IsSupervised" json:"isSupervised"`
	LocalHostName                    string               `plist:"LocalHostName" json:"localhostname"`
	Model                            string               `plist:"Model" json:"model"`
	ModelName                        string               `plist:"ModelName" json:"modelName"`
	OSUpdateSettings                 OSUpdateSettingsType `plist:"OSUpdateSettings" json:"osUpdateSettings"`
	OSVersion                        string               `plist:"OSVersion" json:"osVersion"`
	OSXSoftwareUpdateStatus          OSUpdateSettingsType `plist:"OSXSoftwareUpdateStatus"`
	PINRequiredForDeviceLock         bool                 `plist:"PINRequiredForDeviceLock" json:"pinRequiredForDeviceLock"`
	PINRequiredForEraseDevice        bool                 `plist:"PINRequiredForEraseDevice" json:"pinRequiredForEraseDevice"`
	ProductName                      string               `plist:"ProductName" json:"productName"`
	ProvisioningUDID                 string               `plist:"ProvisioningUDID" json:"provisioningUDID"`
	SerialNumber                     string               `plist:"SerialNumber" json:"serialNumber"`
	SoftwareUpdateDeviceID           string               `plist:"SoftwareUpdateDeviceID" json:"softwareUpdateDeviceID"`
	SupplementalBuildVersion         string               `plist:"SupplementalBuildVersion" json:"supplementalBuildVersion"`
	SupportsLOMDevice                bool                 `plist:"SupportsLOMDevice" json:"supportsLOMDevice"`
	SupportsiOSAppInstalls           bool                 `plist:"SupportsiOSAppInstalls" json:"supportsiOSAppInstalls"`
	SystemIntegrityProtectionEnabled bool                 `plist:"SystemIntegrityProtectionEnabled" json:"systemIntegrityProtectionEnabled"`
	Status                           string               `plist:"Status" json:"status"`
	UDID                             string               `plist:"UDID" json:"udid"`
}
