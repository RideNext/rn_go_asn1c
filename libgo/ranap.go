package main

import (
	e2ap "ridenext.co.in/goasn1/ranap"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {
   msg, _ := hex.DecodeString("0009004b000003001d00050000030d4000050002000d001e0035001b18000300000020000000000120000100000220000200000320000300001340134001000f00010400000210012b10012c10018f00")
   st := e2ap.Stream{}
   st1 := e2ap.Stream{}
   fmt.Println(msg)
   st.Init(msg)
   pdu := e2ap.RANAPPDU{}
   pdu.Unpack(&st)
   jsonData, _ := json.Marshal(pdu)
   fmt.Printf("PDU = %s\n", jsonData)
   pdu.Pack(&st1)
   fmt.Println(st1.Get_buff())
   fmt.Println(st.Get_buff())
}
