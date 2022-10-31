// Package requests used to call define add device request's struct
package requests

import (
	"github.com/thb-cmyk/kubeedge-sd-demo/internal/configmap"
)

// AddDeviceRequest the struct of device request
type AddDeviceRequest struct {
	DeviceInstance *configmap.DeviceInstance `json:"deviceInstances"`
	DeviceModels   []*configmap.DeviceModel  `json:"deviceModels"`
	Protocol       []*configmap.Protocol     `json:"protocols"`
}
