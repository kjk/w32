package main

import "slices"

var namespaces = []string{
	// "Windows.Win32.System.ApplicationInstallationAndServicing",
	// "Windows.Win32.Media.Audio",
	// "Windows.Win32.System.Com",
	// "Windows.Win32.UI.Controls",
	// "Windows.Win32.Security.Cryptography.Certificates",
	// "Windows.Win32.System.Diagnostics.Debug",
	// "Windows.Win32.Devices.Properties",
	// "Windows.Win32.Storage.IscsiDisc",
	// "Windows.Win32.Graphics.Direct3D11",
	// "Windows.Win32.Graphics.Direct3D12",
	// "Windows.Win32.Graphics.Direct3D9",
	// "Windows.Win32.Graphics.Hlsl",
	// "Windows.Win32.Graphics.DirectWrite",
	// "Windows.Win32.System.Diagnostics.Etw",
	// "Windows.Win32.Storage.FileSystem",
	"Windows.Win32.Foundation",
	// "Windows.Win32.Graphics.Gdi",
	// "Windows.Win32.UI.HiDpi",
	// "Windows.Win32.Security.Authentication.Identity",
	// "Windows.Win32.Web.MsHtml",
	// "Windows.Win32.System.Ioctl",
	// "Windows.Win32.System.Js",
	// "Windows.Win32.Media.MediaFoundation",
	// "Windows.Win32.System.Wmi",
	// "Windows.Win32.Media.Multimedia",
	// "Windows.Win32.NetworkManagement.WiFi",
	// "Windows.Win32.System.Ole",
	// "Windows.Win32.System.Registry",
	// "Windows.Win32.System.RemoteDesktop",
	// "Windows.Win32.System.RestartManager",
	// "Windows.Win32.Storage.DistributedFileSystem",
	// "Windows.Win32.Security.Cryptography",
	// "Windows.Win32.Security",
	// "Windows.Win32.Security.Credentials",
	// "Windows.Win32.System.Services",
	// "Windows.Win32.UI.Shell",
	// "Windows.Win32.System.SystemInformation",
	// "Windows.Win32.System.SystemServices",
	// "Windows.Win32.UI.TabletPC",
	// "Windows.Win32.UI.TextServices",
	// "Windows.Win32.System.Threading",
	// "Windows.Win32.AI.MachineLearning.DirectML",
	// "Windows.Win32.AI.MachineLearning.WinML",
	// "Windows.Win32.Data.HtmlHelp",
	// "Windows.Win32.Data.RightsManagement",
	// "Windows.Win32.Data.Xml.MsXml",
	// "Windows.Win32.Devices.AllJoyn",
	// "Windows.Win32.Devices.BiometricFramework",
	// "Windows.Win32.Devices.Bluetooth",
	// "Windows.Win32.Devices.Communication",
	// "Windows.Win32.Devices.DeviceAccess",
	// "Windows.Win32.Devices.DeviceAndDriverInstallation",
	// "Windows.Win32.Devices.Display",
	// "Windows.Win32.Devices.Enumeration.Pnp",
	// "Windows.Win32.Devices.Fax",
	// "Windows.Win32.Devices.FunctionDiscovery",
	// "Windows.Win32.Devices.Geolocation",
	// "Windows.Win32.Devices.HumanInterfaceDevice",
	// "Windows.Win32.Devices.ImageAcquisition",
	// "Windows.Win32.Devices.PortableDevices",
	// "Windows.Win32.Devices.Pwm",
	// "Windows.Win32.Devices.Sensors",
	// "Windows.Win32.Devices.SerialCommunication",
	// "Windows.Win32.Devices.Tapi",
	// "Windows.Win32.Devices.Usb",
	// "Windows.Win32.Devices.WebServicesOnDevices",
	// "Windows.Win32.Gaming",
	// "Windows.Win32.Globalization",
	// "Windows.Win32.Graphics.Direct2D",
	// "Windows.Win32.Graphics.Direct3D",
	// "Windows.Win32.Graphics.Direct3D.Dxc",
	// "Windows.Win32.Graphics.Direct3D.Fxc",
	// "Windows.Win32.Graphics.Direct3D10",
	// "Windows.Win32.Graphics.Direct3D9on12",
	// "Windows.Win32.Graphics.DirectComposition",
	// "Windows.Win32.Graphics.DirectDraw",
	// "Windows.Win32.Graphics.DirectManipulation",
	// "Windows.Win32.Graphics.Dwm",
	// "Windows.Win32.Graphics.DXCore",
	// "Windows.Win32.Graphics.Dxgi.Common",
	// "Windows.Win32.Graphics.Dxgi",
	// "Windows.Win32.Graphics.GdiPlus",
	// "Windows.Win32.Graphics.Imaging",
	// "Windows.Win32.Graphics.OpenGL",
	// "Windows.Win32.Graphics.Printing",
	// "Windows.Win32.Graphics.Printing.PrintTicket",
	// "Windows.Win32.Management.MobileDeviceManagementRegistration",
	// "Windows.Win32.Media.Audio.Apo",
	// "Windows.Win32.Media.Audio.DirectMusic",
	// "Windows.Win32.Media.Audio.DirectSound",
	// "Windows.Win32.Media.Audio.Endpoints",
	// "Windows.Win32.Media.Audio.XAudio2",
	// "Windows.Win32.Media",
	// "Windows.Win32.Media.DeviceManager",
	// "Windows.Win32.Media.DirectShow",
	// "Windows.Win32.Media.DirectShow.Tv",
	// "Windows.Win32.Media.DirectShow.Xml",
	// "Windows.Win32.Media.DxMediaObjects",
	// "Windows.Win32.Media.KernelStreaming",
	// "Windows.Win32.Media.MediaPlayer",
	// "Windows.Win32.Media.PictureAcquisition",
	// "Windows.Win32.Media.Speech",
	// "Windows.Win32.Media.Streaming",
	// "Windows.Win32.Media.WindowsMediaFormat",
	// "Windows.Win32.Networking.ActiveDirectory",
	// "Windows.Win32.Networking.BackgroundIntelligentTransferService",
	// "Windows.Win32.Networking.Clustering",
	// "Windows.Win32.Networking.HttpServer",
	// "Windows.Win32.Networking.Ldap",
	// "Windows.Win32.Networking.NetworkListManager",
	// "Windows.Win32.Networking.RemoteDifferentialCompression",
	// "Windows.Win32.Networking.WebSocket",
	// "Windows.Win32.Networking.WindowsWebServices",
	// "Windows.Win32.Networking.WinHttp",
	// "Windows.Win32.Networking.WinInet",
	// "Windows.Win32.Networking.WinSock",
	// "Windows.Win32.NetworkManagement.Dhcp",
	// "Windows.Win32.NetworkManagement.Dns",
	// "Windows.Win32.NetworkManagement.InternetConnectionWizard",
	// "Windows.Win32.NetworkManagement.IpHelper",
	// "Windows.Win32.NetworkManagement.Multicast",
	// "Windows.Win32.NetworkManagement.Ndis",
	// "Windows.Win32.NetworkManagement.NetBios",
	// "Windows.Win32.NetworkManagement.NetManagement",
	// "Windows.Win32.NetworkManagement.NetShell",
	// "Windows.Win32.NetworkManagement.NetworkDiagnosticsFramework",
	// "Windows.Win32.NetworkManagement.NetworkPolicyServer",
	// "Windows.Win32.NetworkManagement.P2P",
	// "Windows.Win32.NetworkManagement.QoS",
	// "Windows.Win32.NetworkManagement.Rras",
	// "Windows.Win32.NetworkManagement.Snmp",
	// "Windows.Win32.NetworkManagement.WebDav",
	// "Windows.Win32.NetworkManagement.WindowsConnectionManager",
	// "Windows.Win32.NetworkManagement.WindowsConnectNow",
	// "Windows.Win32.NetworkManagement.WindowsFilteringPlatform",
	// "Windows.Win32.NetworkManagement.WindowsFirewall",
	// "Windows.Win32.NetworkManagement.WindowsNetworkVirtualization",
	// "Windows.Win32.NetworkManagement.WNet",
	// "Windows.Win32.Security.AppLocker",
	// "Windows.Win32.Security.Authentication.Identity.Provider",
	// "Windows.Win32.Security.Authorization",
	// "Windows.Win32.Security.Authorization.UI",
	// "Windows.Win32.Security.ConfigurationSnapin",
	// "Windows.Win32.Security.Cryptography.Catalog",
	// "Windows.Win32.Security.Cryptography.Sip",
	// "Windows.Win32.Security.Cryptography.UI",
	// "Windows.Win32.Security.DirectoryServices",
	// "Windows.Win32.Security.ExtensibleAuthenticationProtocol",
	// "Windows.Win32.Security.Isolation",
	// "Windows.Win32.Security.NetworkAccessProtection",
	// "Windows.Win32.Security.Tpm",
	// "Windows.Win32.Security.WinTrust",
	// "Windows.Win32.Security.WinWlx",
	// "Windows.Win32.Storage.Cabinets",
	// "Windows.Win32.Storage.CloudFilters",
	// "Windows.Win32.Storage.Compression",
	// "Windows.Win32.Storage.DataDeduplication",
	// "Windows.Win32.Storage.EnhancedStorage",
	// "Windows.Win32.Storage.FileHistory",
	// "Windows.Win32.Storage.FileServerResourceManager",
	// "Windows.Win32.Storage.Imapi",
	// "Windows.Win32.Storage.IndexServer",
	// "Windows.Win32.Storage.InstallableFileSystems",
	// "Windows.Win32.Storage.Jet",
	// "Windows.Win32.Storage.Nvme",
	// "Windows.Win32.Storage.OfflineFiles",
	// "Windows.Win32.Storage.Packaging.Appx",
	// "Windows.Win32.Storage.Packaging.Opc",
	// "Windows.Win32.Storage.Vhd",
	// "Windows.Win32.Storage.VirtualDiskService",
	// "Windows.Win32.Storage.Vss",
	// "Windows.Win32.Storage.Xps",
	// "Windows.Win32.Storage.Xps.Printing",
	// "Windows.Win32.System.AddressBook",
	// "Windows.Win32.System.ApplicationVerifier",
	// "Windows.Win32.System.ClrHosting",
	// "Windows.Win32.System.Com.StructuredStorage",
	// "Windows.Win32.System.Com.Urlmon",
	// "Windows.Win32.System.ComponentServices",
	// "Windows.Win32.System.Console",
	// "Windows.Win32.System.Contacts",
	// "Windows.Win32.System.CorrelationVector",
	// "Windows.Win32.System.DataExchange",
	// "Windows.Win32.System.DeploymentServices",
	// "Windows.Win32.System.DesktopSharing",
	// "Windows.Win32.System.Diagnostics.Debug.ActiveScript",
	// "Windows.Win32.System.Diagnostics.Debug.Extensions",
	// "Windows.Win32.System.Diagnostics.ProcessSnapshotting",
	// "Windows.Win32.System.Diagnostics.ToolHelp",
	// "Windows.Win32.System.DistributedTransactionCoordinator",
	// "Windows.Win32.System.Environment",
	// "Windows.Win32.System.ErrorReporting",
	// "Windows.Win32.System.EventCollector",
	// "Windows.Win32.System.EventLog",
	// "Windows.Win32.System.EventNotificationService",
	// "Windows.Win32.System.GroupPolicy",
	// "Windows.Win32.System.Hypervisor",
	// "Windows.Win32.System.Iis",
	// "Windows.Win32.System.Kernel",
	// "Windows.Win32.System.LibraryLoader",
	// "Windows.Win32.System.Mapi",
	// "Windows.Win32.System.Memory",
	// "Windows.Win32.System.MessageQueuing",
	// "Windows.Win32.System.MixedReality",
	// "Windows.Win32.System.Mmc",
	// "Windows.Win32.System.ParentalControls",
	// "Windows.Win32.System.Performance",
	// "Windows.Win32.System.Pipes",
	// "Windows.Win32.System.Power",
	// "Windows.Win32.System.ProcessStatus",
	// "Windows.Win32.System.RealTimeCommunications",
	// "Windows.Win32.System.RemoteAssistance",
	// "Windows.Win32.System.RemoteManagement",
	// "Windows.Win32.System.Restore",
	// "Windows.Win32.System.Rpc",
	// "Windows.Win32.System.Search",
	// "Windows.Win32.System.ServerBackup",
	// "Windows.Win32.System.SettingsManagementInfrastructure",
	// "Windows.Win32.System.Shutdown",
	// "Windows.Win32.System.SideShow",
	// "Windows.Win32.System.TaskScheduler",
	// "Windows.Win32.System.Time",
	// "Windows.Win32.System.TpmBaseServices",
	// "Windows.Win32.System.UpdateAgent",
	// "Windows.Win32.System.VirtualDosMachines",
	// "Windows.Win32.System.WindowsProgramming",
	// "Windows.Win32.System.WindowsSync",
	// "Windows.Win32.System.WinRT",
	// "Windows.Win32.System.WinRT.Metadata",
	// "Windows.Win32.System.WinRT.Xaml",
	// "Windows.Win32.UI.Accessibility",
	// "Windows.Win32.UI.Animation",
	// "Windows.Win32.UI.ColorSystem",
	// "Windows.Win32.UI.Controls.Dialogs",
	// "Windows.Win32.UI.Controls.RichEdit",
	// "Windows.Win32.UI.Input.Ime",
	// "Windows.Win32.UI.Input.KeyboardAndMouse",
	// "Windows.Win32.UI.Input.XboxController",
	// "Windows.Win32.UI.LegacyWindowsEnvironmentFeatures",
	// "Windows.Win32.UI.Magnification",
	// "Windows.Win32.UI.Ribbon",
	// "Windows.Win32.UI.Shell.Common",
	// "Windows.Win32.UI.Shell.PropertiesSystem",
	// "Windows.Win32.UI.WindowsAndMessaging",
	// "Windows.Win32.UI.Wpf",
	// "Windows.Win32.UI.Xaml.Diagnostics",
	// "Windows.Win32.Web.InternetExplorer",
	// "Windows.Win32.System.Antimalware",
	// "Windows.Win32.System.Com.Marshal",
	// "Windows.Win32.System.Variant",
	// "Windows.Win32.System.Diagnostics.Ceip",
	// "Windows.Win32.System.Diagnostics.ClrProfiling",
	// "Windows.Win32.System.Com.CallObj",
	// "Windows.Win32.System.Com.ChannelCredentials",
	// "Windows.Win32.System.Com.Events",
	// "Windows.Win32.Graphics.CompositionSwapchain",
	// "Windows.Win32.System.Diagnostics.Debug.WebApp",
	// "Windows.Win32.Devices.DeviceQuery",
	// "Windows.Win32.System.DeveloperLicensing",
	// "Windows.Win32.Graphics.Direct3D11on12",
	// "Windows.Win32.Security.EnterpriseData",
	// "Windows.Win32.System.HostComputeNetwork",
	// "Windows.Win32.System.HostComputeSystem",
	// "Windows.Win32.UI.Input.Radial",
	// "Windows.Win32.UI.Input.Ink",
	// "Windows.Win32.UI.InteractionContext",
	// "Windows.Win32.System.IO",
	// "Windows.Win32.System.JobObjects",
	// "Windows.Win32.NetworkManagement.MobileBroadband",
	// "Windows.Win32.System.PasswordManagement",
	// "Windows.Win32.Storage.ProjectedFileSystem",
	// "Windows.Win32.Security.DiagnosticDataQuery",
	// "Windows.Win32.Security.LicenseProtection",
	// "Windows.Win32.System.SecurityCenter",
	// "Windows.Win32.System.TransactionServer",
	// "Windows.Win32.System.UserAccessLogging",
	// "Windows.Win32.System.UpdateAssessment",
	// "Windows.Win32.UI.Notifications",
	// "Windows.Win32.System.SetupAndMigration",
	// "Windows.Win32.System.WinRT.AllJoyn",
	// "Windows.Win32.System.WinRT.Composition",
	// "Windows.Win32.System.WinRT.CoreInputView",
	// "Windows.Win32.System.WinRT.Graphics.Direct2D",
	// "Windows.Win32.System.WinRT.Direct3D11",
	// "Windows.Win32.System.WinRT.Display",
	// "Windows.Win32.System.WinRT.Graphics.Capture",
	// "Windows.Win32.System.WinRT.Graphics.Imaging",
	// "Windows.Win32.System.WinRT.Holographic",
	// "Windows.Win32.System.WinRT.Isolation",
	// "Windows.Win32.System.WinRT.Media",
	// "Windows.Win32.System.WinRT.ML",
	// "Windows.Win32.System.WinRT.Pdf",
	// "Windows.Win32.System.WinRT.Printing",
	// "Windows.Win32.System.WinRT.Storage",
	// "Windows.Win32.System.AssessmentTool",
	// "Windows.Win32.UI.Input.Touch",
	// "Windows.Win32.Media.LibrarySharingServices",
	// "Windows.Win32.System.SubsystemForLinux",
	// "Windows.Win32.Data.Xml.XmlLite",
	// "Windows.Win32.System.Memory.NonVolatile",
	// "Windows.Win32.System.StationsAndDesktops",
	// "Windows.Win32.UI.Input.Pointer",
	// "Windows.Win32.UI.Input",
	// "Windows.Win32.Storage.OperationRecorder",
	// "Windows.Win32.System.Mailslots",
	// "Windows.Win32.System.Recovery",
	// "Windows.Win32.System.Performance.HardwareCounterProfiling",
}

func skipNamespace(namespace string) bool {
	return !slices.Contains(namespaces, namespace)
}
