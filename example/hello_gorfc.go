package main

import (
	"fmt"
	"time"

	"github.com/lulouis/gorfc/gorfc"
)

func abapSystem() gorfc.ConnectionParameter {
	return gorfc.ConnectionParameter{
		Dest:   "I64",
		Client: "310",
		User:   "1008",
		Passwd: "m4.com",
		Lang:   "EN",
		Ashost: "192.168.2.26",
		Sysnr:  "00",
	}
}

func main() {
	c, err := gorfc.ConnectionFromParams(abapSystem())
	defer c.Close()
	// var t *testing.T

	if err != nil {
		println(err.Error())
		return
	}

	params := map[string]interface{}{
		"IMPORTSTRUCT": map[string]interface{}{
			"RFCFLOAT": 1.23456789,
			"RFCCHAR1": "A",
			"RFCCHAR2": "BC",
			"RFCCHAR4": "ÄBC",
			"RFCINT1":  0xfe,
			"RFCINT2":  0x7ffe,
			"RFCINT4":  999999999,
			"RFCHEX3":  []byte{255, 254, 253},
			"RFCTIME":  time.Now(),
			"RFCDATE":  time.Now(),
			"RFCDATA1": "测试结果：通过",
			"RFCDATA2": "DATA222",
		},
	}
	r, _ := c.Call("STFC_STRUCTURE", params)
	echoStruct := r["ECHOSTRUCT"].(map[string]interface{})
	fmt.Println(echoStruct["RFCDATA1"])
	// assert.NotNil(t, r["ECHOSTRUCT"])
	// importStruct := params["IMPORTSTRUCT"].(map[string]interface{})

	// assert.Equal(t, importStruct["RFCFLOAT"], echoStruct["RFCFLOAT"])
	// assert.Equal(t, importStruct["RFCCHAR1"], echoStruct["RFCCHAR1"])
	// assert.Equal(t, importStruct["RFCCHAR2"], echoStruct["RFCCHAR2"])
	// assert.Equal(t, importStruct["RFCCHAR4"], echoStruct["RFCCHAR4"])
	// assert.Equal(t, importStruct["RFCINT1"], echoStruct["RFCINT1"])
	// assert.Equal(t, importStruct["RFCINT2"], echoStruct["RFCINT2"])
	// assert.Equal(t, importStruct["RFCINT4"], echoStruct["RFCINT4"])
	// //	assert.Equal(t, importStruct["RFCHEX3"], echoStruct["RFCHEX3"])
	// assert.Equal(t, importStruct["RFCTIME"].(time.Time).Format("150405"), echoStruct["RFCTIME"].(time.Time).Format("150405"))
	// assert.Equal(t, importStruct["RFCDATE"].(time.Time).Format("20060102"), echoStruct["RFCDATE"].(time.Time).Format("20060102"))
	// assert.Equal(t, importStruct["RFCDATA1"], echoStruct["RFCDATA1"])
	// assert.Equal(t, importStruct["RFCDATA2"], echoStruct["RFCDATA2"])

	// fmt.Println(reflect.TypeOf(importStruct["RFCDATE"]))
	// fmt.Println(reflect.TypeOf(importStruct["RFCTIME"]))

}
