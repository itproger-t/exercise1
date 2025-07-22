package main

import "fmt"

type HDMI interface {
	ConnectHDMI()
}

type USBDevice struct{}

func (u *USBDevice) ConnectUSB() {
	fmt.Println("✅ USB device connected via USB port")
}

type USBToHDMIAdapter struct {
	device *USBDevice
}

func (a *USBToHDMIAdapter) ConnectHDMI() {
	fmt.Println("🔌 Adapter converts HDMI signal to USB")
	a.device.ConnectUSB()
}

func main() {
	usb := &USBDevice{}

	var hdmiDevice HDMI = &USBToHDMIAdapter{device: usb}

	hdmiDevice.ConnectHDMI()
}
