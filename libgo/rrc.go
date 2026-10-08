package main

import (
	e2ap "ridenext.co.in/goasn1/rrc"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {
   msg, _ := hex.DecodeString("0008004b000003001d00050000030d4000050002000d001e0035001b18000300000020000000000120000100000220000200000320000300001340134001000f00010400000210012b10012c10018f00")
   msg,_ =  hex.DecodeString("13102000005df80105e400340593d8a4000000004000004440b80b838000")
   st := e2ap.Stream{}
   fmt.Println(msg)
   st.Init(msg)
   pdu := e2ap.RRCPDU{}
   pdu.Unpack(&st)
   jsonData, _ := json.Marshal(pdu)
   fmt.Printf("PDU = %s\n", jsonData)
}
