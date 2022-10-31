package mqttadapter

import (
	"encoding/json"
	"strings"
	"time"

	"k8s.io/klog/v2"

	"github.com/thb-cmyk/kubeedge-sd-demo/internal/clients/mqttclient"
	"github.com/thb-cmyk/kubeedge-sd-demo/internal/common"
	"github.com/thb-cmyk/kubeedge-sd-demo/internal/configmap"
	"github.com/thb-cmyk/kubeedge-sd-demo/internal/controller"
	"github.com/thb-cmyk/kubeedge-sd-demo/pkg/di"
	"github.com/thb-cmyk/kubeedge-sd-demo/pkg/models"
)

// TwinData the structure of device twin
type TwinData struct {
	Name       string
	Type       string
	Topic      string
	Value      string
	MqttClient mqttclient.MqttClient
	driverUnit DriverUnit
}

// DriverUnit the structure necessary to send a message
type DriverUnit struct {
	instanceID string
	twin       configmap.Twin
	drivers    models.ProtocolDriver
	mutex      *common.Lock
	dic        *di.Container
}

// Run start timer function to get device's twin or data, and send it to mqtt broker
func (td *TwinData) Run() {
	var err error
	sData, err := controller.GetDeviceData(td.driverUnit.instanceID, td.driverUnit.twin, td.driverUnit.drivers, td.driverUnit.mutex, td.driverUnit.dic)
	if err != nil {
		klog.Errorf("Get %s data error:", td.driverUnit.instanceID, err.Error())
		return
	}
	td.Value = sData
	var payload []byte
	if strings.Contains(td.Topic, "$hw") {
		if payload, err = CreateMessageTwinUpdate(td.Name, td.Type, td.Value); err != nil {
			klog.Errorf("Create %s message twin update failed: %v", td.driverUnit.instanceID, err)
			return
		}
	} else {
		if payload, err = CreateMessageData(td.Name, td.Type, td.Value); err != nil {
			klog.Errorf("Create %s message data failed: %v", td.driverUnit.instanceID, err)
			return
		}
	}
	if err := td.MqttClient.Publish(td.Topic, payload); err != nil {
		klog.Errorf("Publish topic %v failed, err: %v", td.Topic, err)
	}
}

// CreateMessageTwinUpdate create twin update message.
func CreateMessageTwinUpdate(name string, valueType string, value string) (msg []byte, err error) {
	var updateMsg DeviceTwinUpdate

	updateMsg.BaseMessage.Timestamp = time.Now().UnixNano() / 1e6
	updateMsg.Twin = map[string]*MsgTwin{}
	updateMsg.Twin[name] = &MsgTwin{}
	updateMsg.Twin[name].Actual = &TwinValue{Value: &value}
	updateMsg.Twin[name].Metadata = &TypeMetadata{Type: valueType}

	msg, err = json.Marshal(updateMsg)
	return
}
