import { readdir } from "node:fs/promises";

// this is location of local checkout of https://github.com/marlersoft/win32json
let win32jsonDir = "..\\win32json\\api";

const files = [
  // "AI.MachineLearning.DirectML.json"
  // "AI.MachineLearning.WinML.json"
  // "Data.HtmlHelp.json"
  // "Data.RightsManagement.json"
  // "Data.Xml.MsXml.json"
  // "Data.Xml.XmlLite.json"
  // "Devices.AllJoyn.json"
  // "Devices.BiometricFramework.json"
  // "Devices.Bluetooth.json"
  // "Devices.Communication.json"
  // "Devices.DeviceAccess.json"
  // "Devices.DeviceAndDriverInstallation.json"
  // "Devices.DeviceQuery.json"
  // "Devices.Display.json"
  // "Devices.Enumeration.Pnp.json"
  // "Devices.Fax.json"
  // "Devices.FunctionDiscovery.json"
  // "Devices.Geolocation.json"
  // "Devices.HumanInterfaceDevice.json"
  // "Devices.ImageAcquisition.json"
  // "Devices.PortableDevices.json"
  // "Devices.Properties.json"
  // "Devices.Pwm.json"
  // "Devices.Sensors.json"
  // "Devices.SerialCommunication.json"
  // "Devices.Tapi.json"
  // "Devices.Usb.json"
  // "Devices.WebServicesOnDevices.json"
  "Foundation.json",
  // "Gaming.json"
  // "Globalization.json"
  // "Graphics.CompositionSwapchain.json"
  // "Graphics.Direct2D.Common.json"
  // "Graphics.Direct2D.json"
  // "Graphics.Direct3D.Dxc.json"
  // "Graphics.Direct3D.Fxc.json"
  // "Graphics.Direct3D.json"
  // "Graphics.Direct3D10.json"
  // "Graphics.Direct3D11.json"
  // "Graphics.Direct3D11on12.json"
  // "Graphics.Direct3D12.json"
  // "Graphics.Direct3D9.json"
  // "Graphics.Direct3D9on12.json"
  // "Graphics.DirectComposition.json"
  // "Graphics.DirectDraw.json"
  // "Graphics.DirectManipulation.json"
  // "Graphics.DirectWrite.json"
  // "Graphics.Dwm.json"
  // "Graphics.DXCore.json"
  // "Graphics.Dxgi.Common.json"
  // "Graphics.Dxgi.json"
  // "Graphics.Gdi.json"
  // "Graphics.Hlsl.json"
  // "Graphics.Imaging.D2D.json"
  // "Graphics.Imaging.json"
  // "Graphics.OpenGL.json"
  // "Graphics.Printing.json"
  // "Graphics.Printing.PrintTicket.json"
  // "Management.MobileDeviceManagementRegistration.json"
  // "Media.Audio.Apo.json"
  // "Media.Audio.DirectMusic.json"
  // "Media.Audio.DirectSound.json"
  // "Media.Audio.Endpoints.json"
  // "Media.Audio.json"
  // "Media.Audio.XAudio2.json"
  // "Media.DeviceManager.json"
  // "Media.DirectShow.json"
  // "Media.DirectShow.Xml.json"
  // "Media.DxMediaObjects.json"
  // "Media.json"
  // "Media.KernelStreaming.json"
  // "Media.LibrarySharingServices.json"
  // "Media.MediaFoundation.json"
  // "Media.MediaPlayer.json"
  // "Media.Multimedia.json"
  // "Media.PictureAcquisition.json"
  // "Media.Speech.json"
  // "Media.Streaming.json"
  // "Media.WindowsMediaFormat.json"
  // "Networking.ActiveDirectory.json"
  // "Networking.BackgroundIntelligentTransferService.json"
  // "Networking.Clustering.json"
  // "Networking.HttpServer.json"
  // "Networking.Ldap.json"
  // "Networking.NetworkListManager.json"
  // "Networking.RemoteDifferentialCompression.json"
  // "Networking.WebSocket.json"
  // "Networking.WindowsWebServices.json"
  // "Networking.WinHttp.json"
  // "Networking.WinInet.json"
  // "Networking.WinSock.json"
  // "NetworkManagement.Dhcp.json"
  // "NetworkManagement.Dns.json"
  // "NetworkManagement.InternetConnectionWizard.json"
  // "NetworkManagement.IpHelper.json"
  // "NetworkManagement.MobileBroadband.json"
  // "NetworkManagement.Multicast.json"
  // "NetworkManagement.Ndis.json"
  // "NetworkManagement.NetBios.json"
  // "NetworkManagement.NetManagement.json"
  // "NetworkManagement.NetShell.json"
  // "NetworkManagement.NetworkDiagnosticsFramework.json"
  // "NetworkManagement.NetworkPolicyServer.json"
  // "NetworkManagement.P2P.json"
  // "NetworkManagement.QoS.json"
  // "NetworkManagement.Rras.json"
  // "NetworkManagement.Snmp.json"
  // "NetworkManagement.WebDav.json"
  // "NetworkManagement.WiFi.json"
  // "NetworkManagement.WindowsConnectionManager.json"
  // "NetworkManagement.WindowsConnectNow.json"
  // "NetworkManagement.WindowsFilteringPlatform.json"
  // "NetworkManagement.WindowsFirewall.json"
  // "NetworkManagement.WindowsNetworkVirtualization.json"
  // "NetworkManagement.WNet.json"
  // "Security.AppLocker.json"
  // "Security.Authentication.Identity.json"
  // "Security.Authentication.Identity.Provider.json"
  // "Security.Authorization.json"
  // "Security.Authorization.UI.json"
  // "Security.ConfigurationSnapin.json"
  // "Security.Credentials.json"
  // "Security.Cryptography.Catalog.json"
  // "Security.Cryptography.Certificates.json"
  // "Security.Cryptography.json"
  // "Security.Cryptography.Sip.json"
  // "Security.Cryptography.UI.json"
  // "Security.DiagnosticDataQuery.json"
  // "Security.DirectoryServices.json"
  // "Security.EnterpriseData.json"
  // "Security.ExtensibleAuthenticationProtocol.json"
  // "Security.Isolation.json"
  // "Security.json"
  // "Security.LicenseProtection.json"
  // "Security.NetworkAccessProtection.json"
  // "Security.Tpm.json"
  // "Security.WinTrust.json"
  // "Security.WinWlx.json"
  // "Storage.Cabinets.json"
  // "Storage.CloudFilters.json"
  // "Storage.Compression.json"
  // "Storage.DataDeduplication.json"
  // "Storage.DistributedFileSystem.json"
  // "Storage.EnhancedStorage.json"
  // "Storage.FileHistory.json"
  // "Storage.FileServerResourceManager.json"
  // "Storage.FileSystem.json"
  // "Storage.Imapi.json"
  // "Storage.IndexServer.json"
  // "Storage.InstallableFileSystems.json"
  // "Storage.IscsiDisc.json"
  // "Storage.Jet.json"
  // "Storage.OfflineFiles.json"
  // "Storage.OperationRecorder.json"
  // "Storage.Packaging.Appx.json"
  // "Storage.Packaging.Opc.json"
  // "Storage.ProjectedFileSystem.json"
  // "Storage.StructuredStorage.json"
  // "Storage.Vhd.json"
  // "Storage.VirtualDiskService.json"
  // "Storage.Vss.json"
  // "Storage.Xps.json"
  // "Storage.Xps.Printing.json"
  // "System.AddressBook.json"
  // "System.Antimalware.json"
  // "System.ApplicationInstallationAndServicing.json"
  // "System.ApplicationVerifier.json"
  // "System.AssessmentTool.json"
  // "System.Com.CallObj.json"
  // "System.Com.ChannelCredentials.json"
  // "System.Com.Events.json"
  // "System.Com.json"
  // "System.Com.Marshal.json"
  // "System.Com.StructuredStorage.json"
  // "System.Com.UI.json"
  // "System.Com.Urlmon.json"
  // "System.ComponentServices.json"
  // "System.Console.json"
  // "System.Contacts.json"
  // "System.CorrelationVector.json"
  // "System.DataExchange.json"
  // "System.DeploymentServices.json"
  // "System.DesktopSharing.json"
  // "System.DeveloperLicensing.json"
  // "System.Diagnostics.Ceip.json"
  // "System.Diagnostics.Debug.json"
  // "System.Diagnostics.Debug.WebApp.json"
  // "System.Diagnostics.Etw.json"
  // "System.Diagnostics.ProcessSnapshotting.json"
  // "System.Diagnostics.ToolHelp.json"
  // "System.DistributedTransactionCoordinator.json"
  // "System.Environment.json"
  // "System.ErrorReporting.json"
  // "System.EventCollector.json"
  // "System.EventLog.json"
  // "System.EventNotificationService.json"
  // "System.GroupPolicy.json"
  // "System.HostCompute.json"
  // "System.HostComputeNetwork.json"
  // "System.HostComputeSystem.json"
  // "System.Hypervisor.json"
  // "System.Iis.json"
  // "System.IO.json"
  // "System.Ioctl.json"
  // "System.JobObjects.json"
  // "System.Js.json"
  // "System.Kernel.json"
  // "System.LibraryLoader.json"
  // "System.Mailslots.json"
  // "System.Mapi.json"
  "System.Memory.json",
  // "System.Memory.NonVolatile.json"
  // "System.MessageQueuing.json"
  // "System.MixedReality.json"
  // "System.Mmc.json"
  // "System.Ole.json"
  // "System.ParentalControls.json"
  // "System.PasswordManagement.json"
  // "System.Performance.HardwareCounterProfiling.json"
  // "System.Performance.json"
  // "System.Pipes.json"
  // "System.Power.json"
  // "System.ProcessStatus.json"
  // "System.RealTimeCommunications.json"
  // "System.Recovery.json"
  // "System.Registry.json"
  // "System.RemoteAssistance.json"
  // "System.RemoteDesktop.json"
  // "System.RemoteManagement.json"
  // "System.RestartManager.json"
  // "System.Restore.json"
  // "System.Rpc.json"
  // "System.Search.Common.json"
  // "System.Search.json"
  // "System.SecurityCenter.json"
  // "System.ServerBackup.json"
  // "System.Services.json"
  // "System.SettingsManagementInfrastructure.json"
  // "System.SetupAndMigration.json"
  // "System.Shutdown.json"
  // "System.SideShow.json"
  // "System.StationsAndDesktops.json"
  // "System.SubsystemForLinux.json"
  // "System.SystemInformation.json"
  // "System.SystemServices.json"
  // "System.TaskScheduler.json"
  // "System.Threading.json"
  // "System.Time.json"
  // "System.TpmBaseServices.json"
  // "System.TransactionServer.json"
  // "System.UpdateAgent.json"
  // "System.UpdateAssessment.json"
  // "System.UserAccessLogging.json"
  // "System.VirtualDosMachines.json"
  "System.WindowsProgramming.json",
  // "System.WindowsSync.json"
  // "System.WinRT.AllJoyn.json"
  // "System.WinRT.Composition.json"
  // "System.WinRT.CoreInputView.json"
  // "System.WinRT.Direct3D11.json"
  // "System.WinRT.Display.json"
  // "System.WinRT.Graphics.Capture.json"
  // "System.WinRT.Graphics.Direct2D.json"
  // "System.WinRT.Graphics.Imaging.json"
  // "System.WinRT.Holographic.json"
  // "System.WinRT.Isolation.json"
  // "System.WinRT.json"
  // "System.WinRT.Media.json"
  // "System.WinRT.ML.json"
  // "System.WinRT.Pdf.json"
  // "System.WinRT.Printing.json"
  // "System.WinRT.Shell.json"
  // "System.WinRT.Storage.json"
  // "System.WinRT.Xaml.json"
  // "System.Wmi.json"
  // "UI.Accessibility.json"
  // "UI.Animation.json"
  // "UI.ColorSystem.json"
  // "UI.Controls.Dialogs.json"
  // "UI.Controls.json"
  // "UI.Controls.RichEdit.json"
  // "UI.HiDpi.json"
  // "UI.Input.Ime.json"
  // "UI.Input.Ink.json"
  // "UI.Input.json"
  // "UI.Input.KeyboardAndMouse.json"
  // "UI.Input.Pointer.json"
  // "UI.Input.Radial.json"
  // "UI.Input.Touch.json"
  // "UI.Input.XboxController.json"
  // "UI.InteractionContext.json"
  // "UI.LegacyWindowsEnvironmentFeatures.json"
  // "UI.Magnification.json"
  // "UI.Notifications.json"
  // "UI.Ribbon.json"
  // "UI.Shell.Common.json"
  // "UI.Shell.json"
  // "UI.Shell.PropertiesSystem.json"
  // "UI.TabletPC.json"
  // "UI.TextServices.json"
  // "UI.WindowsAndMessaging.json"
  // "UI.Wpf.json"
  // "UI.Xaml.Diagnostics.json"
  // "Web.MsHtml.json"
];

function len(o) {
  if (o) {
    return o.length;
  }
  return 0;
}

// content of current file
let curr;
let skipped;

// capitalize the first letter of each word and lower the rest
function capitalize(s) {
  const s2 = s.toLowerCase();
  return s2.charAt(0).toUpperCase() + s2.slice(1);
}

// convert "SECURITY_ATTRIBUTES" => "SecurityAttributes"
// this is naming in Go syscall package
function toCamelCase(str) {
  const words = str.split("_");
  let n = len(words);
  for (let i = 0; i < n; i++) {
    words[i] = capitalize(words[i]);
  }
  return words.join("");
}

function assertEq(v1, v2) {
  if (v1 === v2) {
    return;
  }
  throw new Error(`'${v1}' !== '${v2}'`);
}

function tests() {
  assertEq("SecurityAttributes", toCamelCase("SECURITY_ATTRIBUTES"));
}

/*
Platform frequency (if not null)
windows6.0.6000 : 1023
windows6.1 : 660
windows5.1.2600 : 651
windows5.0 : 589
windows8.0 : 370
windows8.1 : 115
windows10.0.10240 : 107
windowsServer2008 : 107
windowsServer2012 : 55
windows10.0.15063 : 49
windowsServer2003 : 40
windows10.0.19041 : 15
windows10.0.14393 : 14
windows10.0.17134 : 10
windows10.0.17763 : 8
windows10.0.16299 : 3
windows10.0.18362 : 1
windowsServer2016 : 1
/*

types is array of objects:
{
  "Name":"LUID"
  ,"Architectures":[] // array of X64, Arm64, X86
  ,"Platform":null // or windows5.0, etc. see above
  ,"Kind":"Struct" // "Enum", "Struct", "Union", "Com", "FunctionPointer", "ComClassID", "NativeTypedef"
  ,"Size":0 // always zero, so ignore
  ,"PackingSize":0
  ,"Fields":[
    {"Name":"LowPart","Type":{"Kind":"Native","Name":"UInt32"},"Attrs":[]}
    ,{"Name":"HighPart","Type":{"Kind":"Native","Name":"Int32"},"Attrs":[]}
  ]
  ,"NestedTypes":[
  ]
}

{
  "Name":"DEVPROPKEY"
  ,"Architectures":[]
  ,"Platform":null
  ,"Kind":"Struct"
  ,"Size":0
  ,"PackingSize":0
  ,"Fields":[
    {"Name":"fmtid","Type":{"Kind":"Native","Name":"Guid"},"Attrs":[]}
    ,{"Name":"pid","Type":{"Kind":"Native","Name":"UInt32"},"Attrs":[]}
  ]
  ,"NestedTypes":[
  ]
}
*/
function gen_types(types) {
  curr += "\n";
  curr += `// Section: types (${len(types)})\n`;
  for (let t of types) {
    // TODO: implement me
  }
}

/**
 * @param {number} number
 * @param {number} padding
 * @returns string
 */
function numToHex(number, padding = 0) {
  if (number < 0) {
    number = 0xffffffff + number + 1;
  }
  let hexString = number.toString(16);
  return "0x" + hexString.padStart(padding, "0");
}

function startsWithPrefix(s, prefixes) {
  for (let p of prefixes) {
    if (s.startsWith(p)) {
      return true;
    }
  }
  return false;
}

function constAsHex(name, file, number) {
  if (file === "Foundation.json") {
    let prefixes = ["EXCEPTION_", "E_", "DISP_E_", "STATUS_"];
    return startsWithPrefix(name, prefixes);
  }
  return false;
}

/*
ValueType: "UInt32", "Int32", "String", "UInt64", "Byte", "UInt16", "PropertyKey", "Single", "Double", "Int64"
*/
function getGoValueType(vt) {
  switch (vt) {
    case "Byte":
      return "uint8";
    case "UInt32":
      return "uint32";
    case "UInt64":
      return "uint64";
    case "Int32":
      return "int32";
    case "Single":
      return "float32";
    case "Double":
      return "float64";
    case "String":
      return "string";
    default:
      throw new Error(`ValueType '${vt}' not yet supported`);
    //case "PropertyKey":
    // not supported yet
  }
}
/*
{
  "Name":"INVALID_HANDLE_VALUE"
  ,"Type":{"Kind":"ApiRef","Name":"HANDLE","TargetKind":"Default","Api":"Foundation","Parents":[]}
  ,"ValueType":"Int32"
  ,"Value":-1
  ,"Attrs":[]
}

{
		"Name":"DEVPKEY_IndirectDisplay"
		,"Type":{"Kind":"ApiRef","Name":"DEVPROPKEY","TargetKind":"Default","Api":"Devices.Properties","Parents":[]}
		,"ValueType":"PropertyKey"
		,"Value":{"Fmtid":"c50a3f10-aa5c-4247-b830-d6a6f8eaa310","Pid":1}
		,"Attrs":[]
	}
*/
/*
ValueType: PropertyKey
pub const DEVPKEY_Device_TerminalLuid = DEVPROPKEY { .fmtid = Guid.initString("c50a3f10-aa5c-4247-b830-d6a6f8eaa310"), .pid = 2 };
*/

/**
 *
 * @param {string} name
 * @param {string} file
 * @returns
 */
function skip_constant(name, file) {
  if (file === "Foundation.json") {
    if (name.startsWith("SQLITE_E_")) {
      return true;
    }
  }
  return false;
}

function throwKUnkownKeyValueInObject(k, o) {
  let v = o[k];
  let os = JSON.stringify(o, null, 2);
  throw new Error(`Uknown value '${v}' of key '${k}' in ${os}`);
}

/*
		,"Type":{"Kind":"ApiRef","Name":"HANDLE","TargetKind":"Default","Api":"Foundation","Parents":[]}
*/
function get_const_type(t) {
  let kind = t["Kind"];
  if (kind === "ApiRef") {
    let tk = t["TargetKind"];
    if (tk != "Default") {
      throwKUnkownKeyValueInObject("TargetKind", t);
    }
    let parents = t["Parents"];
    if (len(parents) != 0) {
      throwKUnkownKeyValueInObject("Parents", t);
    }
    let name = t["Name"];
    let validNames = ["HANDLE", "NTSTATUS", "HRESULT"];
    if (!validNames.includes(name)) {
      throwKUnkownKeyValueInObject("Name", t);
    }
    let api = t["Api"];
    let validApis = ["Foundation"];
    if (!validApis.includes(api)) {
      throwKUnkownKeyValueInObject("Api", t);
    }
    return name;
  } else if (kind === "Native") {
    let name = t["Name"];
    return getGoValueType(name);
  } else {
    throwKUnkownKeyValueInObject("Kind", t);
  }
}

/*
"d20beec4-5ca8-4905-ae3b-bf251ea09b53"
=>
GUID{0xd20beec4, 0x5ca8, 0x4905, [8]byte{0xae, 0x3b, 0xbf, 0x25, 0x1e, 0xa0, 0x9b, 0x53}}

TODO: potentially change GUID to sth. else e.g. KNOWNFOLDERID
*/
function genGUID(s) {
  return "";
}

function gen_constants(consts, file) {
  curr += `// Section: constants (${len(consts)})\n`;
  curr += `const (\n`;
  for (let c of consts) {
    let name = c["Name"];
    if (skip_constant(name, file)) {
      skipped += `const ${name}\n`;
      continue;
    }
    get_const_type(c["Type"]); // validate type
    let v = c["Value"];
    let goValueType = getGoValueType(c["ValueType"]);
    switch (goValueType) {
      case "uint8":
      case "uint32":
      case "uint64":
      case "int32":
        if (constAsHex(name, file, v)) {
          v = numToHex(v);
        }
        curr += `  ${name} = ${v}\n`;
        break;
      case "string":
        curr += `  ${name} = "${v}"\n`;
        break;
      case "float32":
      case "float64":
        curr += `  ${name} = ${goValueType}(${v})\n`;
        console.log(curr);
        break;
      default:
        console.log(goValueType);
        throw new Error(`'${goValueType}' not supprted`);
    }
  }
  curr += `)\n`;
}

async function readAsJSON(path) {
  let f = Bun.file(path);
  // console.log("readJSON:", path, "size:", f.size);
  let s = await f.text();
  // console.log("len(s):", len(s));
  let json = JSON.parse(s);
  return json;
}

/**
 * @param {string} name_json
 * @param {string} s
 */
async function save_go(name_json, s) {
  let name = name_json.toLowerCase();
  name = name.replace(/\./g, "_");
  name = name.replace("_json", ".go");
  await Bun.write(name, s);
}

/**
 * @param {string} name_json
 * @param {string} s
 */
async function save_skipped(name_json, s) {
  if (len(s) === 0) {
    return;
  }
  let name = name_json.toLowerCase();
  name = name.replace(".json", ".skipped.txt");
  await Bun.write(name, s);
}

async function process_file(name) {
  let path = win32jsonDir + "\\" + name;
  console.log("process_file:", name, "path:", path);
  let json = await readAsJSON(path);
  // console.log(Object.keys(json));

  curr = "package w32\n\n";
  skipped = "";
  curr += "//! NOTE: AUTO-GENERATED!\n";
  // [ "Constants", "Types", "Functions", "UnicodeAliases" ]
  let consts = json["Constants"];
  gen_constants(consts, name);
  let types = json["Types"];
  gen_types(types);
  await save_go(name, curr);
  await save_skipped(name, skipped);
}

// m is {} that maps string => number (frequency)
// print key : frequency sorted by frequency
function print_map_by_freq(m) {
  let entries = Object.entries(m);
  entries.sort((a, b) => b[1] - a[1]);
  for (let e of entries) {
    console.log(`${e[0]} : ${e[1]}`);
  }
  console.log(Object.keys(m));
}

function incKey(o, key) {
  let n = o[key] || 0;
  o[key] = n + 1;
}

function collect_by_key(a, m, key) {
  for (let o of a) {
    let p = o[key];
    if (p !== null) {
      incKey(m, p);
    }
  }
}

function collect_array_by_key(a, archs, key) {
  for (let o of a) {
    let a = o[key];
    if (len(a) === 0) {
      continue;
    }
    for (let arch of a) {
      incKey(archs, arch);
    }
  }
}

// this is to extract various stats / info about values for a given json type
async function stats_process_all_files() {
  let dir = "..\\..\\win32json\\api";
  let files = await readdir(dir, {});
  let platforms = {};
  let archs = {};
  let kinds = {};
  let valueTypes = {};
  for (let name of files) {
    if (!name.endsWith(".json")) {
      continue;
    }
    console.log(name);
    let path = dir + "\\" + name;
    let json = await readAsJSON(path);
    let types = json["Types"];
    collect_by_key(types, platforms, "Platform");
    collect_by_key(types, kinds, "Kind");
    collect_array_by_key(types, archs, "Architectures");
    let consts = json["Constants"];
    collect_by_key(consts, valueTypes, "ValueType");
  }
  //print_map_by_freq(platforms);
  //print_map_by_freq(archs);
  // print_map_by_freq(kinds);
  print_map_by_freq(valueTypes);
}

async function do_stats() {
  await stats_process_all_files();
}

async function process_all_files() {
  let files = await readdir(win32jsonDir, {});
  for (let name of files) {
    if (!name.endsWith(".json")) {
      continue;
    }
    console.log(name);
    process_file(name);
  }
}

tests();

async function gen_files_list() {
  let s = "const files = [\n";
  let files = await readdir(win32jsonDir, {});
  for (let name of files) {
    if (!name.endsWith(".json")) {
      continue;
    }
    s += '  // "' + name + '"\n';
  }
  s += "];\n";
  console.log(s);
}

if (true) {
  for (let f of files) {
    await process_file(f);
  }
}
if (false) {
  await gen_files_list();
  await do_stats();
  await process_all_files();
}
