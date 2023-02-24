// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-15
// Based on yundong by liasica, magicrolan@qq.com.

package parser

import (
    "github.com/auroraride/adapter"
    "net"
    "strconv"
)

var loginDataConfig = []Field{
    {Length: 4, Name: "Version"},
    {Length: 1, Name: "Type"},
    {Length: 16, Name: "Serial"},
    {Length: 15, Name: "IMEI"},
    {Length: 15, Name: "IMSI"},
    {Length: 6, Name: "WifiMac"},
    {Length: 6, Name: "EthMac"},
    {Length: 16, Name: "Key"},
    {Length: 4, Name: "IP"},
}

type Version [4]byte

func (v Version) String() string {
    // major.minor.patch
    return strconv.Itoa(int(v[1])) + "." + strconv.Itoa(int(v[2])) + "." + strconv.Itoa(int(v[3]))
}

type LoginData struct {
    Version Version `json:"version"` // 固件版本号
    Type    uint8   `json:"type"`    // 电池柜类型，用于区分不同的电池柜类型
    Serial  string  `json:"serial"`  // 电柜编码
    IMEI    string  `json:"imei"`    // 电柜IMEI编码
    IMSI    string  `json:"imsi"`    // 电柜IMSI编码
    WifiMac string  `json:"wifiMac"` // WIFI Mac地址
    EthMac  string  `json:"ethMac"`  // 网口mac地址
    Key     string  `json:"key"`     // AES-128-ECB密钥
    IP      string  `json:"ip"`      // 本次登录使用的接入服务器IP地址
}

func (d *LoginData) SetField(b []byte, name string) {
    switch name {
    case "Version":
        copy(d.Version[:], b)
    case "Type":
        d.Type = b[0]
    case "Serial":
        d.Serial = adapter.ConvertBytes2String(b)
    case "IMEI":
        d.IMSI = adapter.ConvertBytes2String(b)
    case "IMSI":
        d.IMSI = adapter.ConvertBytes2String(b)
    case "WifiMac":
        d.WifiMac = net.HardwareAddr(b).String()
    case "EthMac":
        d.EthMac = net.HardwareAddr(b).String()
    case "Key":
        d.Key = adapter.ConvertBytes2String(b)
    case "IP":
        d.IP = net.IPv4(b[0], b[1], b[2], b[3]).String()
    }
}

func (d *LoginData) String() string {
    buf := adapter.NewBuffer()
    defer adapter.ReleaseBuffer(buf)

    buf.WriteString("固件版本号=")
    buf.WriteString(d.Version.String())

    buf.WriteString(", 电池柜类型=")
    buf.WriteString(strconv.Itoa(int(d.Type)))

    buf.WriteString(", 电柜编码=")
    buf.WriteString(d.Serial)

    buf.WriteString(", IMEI=")
    buf.WriteString(d.IMEI)

    buf.WriteString(", IMSI=")
    buf.WriteString(d.IMSI)

    buf.WriteString(", WifiMAC=")
    buf.WriteString(d.WifiMac)

    buf.WriteString(", EthMAC=")
    buf.WriteString(d.EthMac)

    buf.WriteString(", Key=")
    buf.WriteString(d.Key)

    buf.WriteString(", IP=")
    buf.WriteString(d.IP)

    return buf.String()
}

func (p *Parser) Login(b []byte) (serial string, logdata string) {
    // TODO Save
    data := new(LoginData)
    Parse(data, b, loginDataConfig)
    // c.serial = data.Serial

    // 保存客户端
    // clients.Store(data.Serial, c)

    return data.Serial, data.String()
}
