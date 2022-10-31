package main

import (
	"github.com/thb-cmyk/kubeedge-sd-demo/driver"
	"github.com/thb-cmyk/kubeedge-sd-demo/pkg/service"
)

// main Template device program entry
func main() {
	d := &driver.Template{}
	// TODO: Modify your protocol name to be consistent with the CRDs definition
	service.Bootstrap("Template", d)
}
