import { readdir } from "node:fs/promises";

// this is location of local checkout of https://github.com/marlersoft/win32json
let win32jsonDir = "..\\win32json\\api";

function len(o) {
  if (o) {
    return o.length;
  }
  return 0;
}

// content of current file
let curr;
let skipped;

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
    default:
      throw new Error(`ValueType '${vt}' not yet supported`);
    //case "PropertyKey":
    //case "String":
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
  name = name.replace(".json", ".go");
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

if (true) {
  await process_file("Foundation.json");
}
if (false) {
  await do_stats();
  await process_all_files();
}
