/*

№ 21 Реализовать паттерн «адаптер» на любом примере.

*/

package main

import "fmt"

type OldDeviceImpl struct{}

func (o *OldDeviceImpl) UseOldMethod() {
	fmt.Println("Старый метод работы с устройством")
}

type NewDeviceImpl struct{}

func (n *NewDeviceImpl) UseNewMethod() {
	fmt.Println("Новый метод работы с устройством")
}

type Adapter struct {
	newDeviceImpl NewDeviceImpl
}

func (a *Adapter) UseOldMethod() {
	a.newDeviceImpl.UseNewMethod()
}

func main() {
	oldDevice := &OldDeviceImpl{}
	oldDevice.UseOldMethod()

	newDevice := &NewDeviceImpl{}
	adapter := &Adapter{*newDevice}
	adapter.UseOldMethod()
}
