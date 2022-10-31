package instancepool

import (
	"github.com/thb-cmyk/kubeedge-sd-demo/pkg/di"
)

// ConfigMapName contains the name of device service's ConfigurationStruct implementation in the DIC.
var ConfigMapName = di.TypeInstanceToName(string("configMap"))

// ConfigMapNameFrom helper function queries the DIC and returns device service's ConfigurationStruct implementation.
func ConfigMapNameFrom(get di.Get) string {
	return get(ConfigMapName).(string)
}
