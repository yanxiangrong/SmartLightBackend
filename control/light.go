package control

import (
	"SmartLightBackend/models"
	"SmartLightBackend/network"
	"SmartLightBackend/pkg/logging"
	"fmt"
)

func init() {
	network.SetCallBack(receiveCallBack)
}

func OpenLight(devicesId int) {
	SetStatus(devicesId, true)
}

func CloseLight(devicesId int) {
	SetStatus(devicesId, false)
}

func OpenAllLight() {
	devices := models.GetAllDevices()
	for i := 0; i < len(devices); i++ {
		OpenLight(devices[i].Addr)
	}
}

func CloseAllLight() {
	devices := models.GetAllDevices()
	for i := 0; i < len(devices); i++ {
		CloseLight(devices[i].Addr)
	}
}

// SetStatus 设置电灯状态，开（true）或关
func SetStatus(devicesId int, status bool) {
	var s uint8
	if status {
		s = 0x01
	} else {
		s = 0x00
	}
	command := []uint8{0xFE, 0x06, 0x91, 0x90, uint8(devicesId >> 8), uint8(devicesId & 0xFF), 0x00, s, 0xFF}
	network.SendBroadcast(command)
}

// 解析从设备接收到的数据
func parsingReportData(bytes []byte) models.Device {
	id := (uint16(bytes[4]) << 8) | uint16(bytes[5])
	str := string(bytes[6:13])

	var current float32
	var switchStatus, fault int

	_, err := fmt.Sscanf(str, "%d,%d,%f", &fault, &switchStatus, &current)
	if err != nil {
		logging.Debug("Parsing Error ", err.Error())
		return models.Device{}
	}

	var device = models.Device{
		Addr:    int(id),
		Current: new(int),
		Switch:  new(bool),
		Fault:   new(bool),
	}
	*device.Current = int(current * 10)
	*device.Switch = switchStatus == 1
	*device.Fault = fault == 1

	return device
}

// 接收到数据的回调函数
func receiveCallBack(bytes []byte) {
	device := parsingReportData(bytes)
	if !models.IsDeviceExist(device) {
		logging.Info("Add devices to database")
		models.AddDevice(device)
	} else {
		models.UpdateDevice(device)
	}
}
