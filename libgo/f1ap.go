package main

import (
	f1ap "ridenext.co.in/goasn1/f1ap"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {
   //msg, _ := hex.DecodeString("000b402f000005002900020006006f0009003204100000000000005f0003000000003200080710000000006000004e40020000")
   //msg, _ := hex.DecodeString("000d4029000004002800034003ea0029000200060040000120003200100f000008600000000000000000000000")
   msg, _ := hex.DecodeString("00070015000003002800034003ea0029000200060057400100")
   st := f1ap.Stream{}
   st1 := f1ap.Stream{}
   fmt.Println(msg)
   st.Init(msg)
   pdu := f1ap.F1APPDU{}
   pdu.Unpack(&st)
   jsonData, _ := json.Marshal(pdu)
   fmt.Printf("PDU = %s\n", jsonData)
   pdu.Pack(&st1)
   fmt.Println(st1.Get_buff())
   fmt.Println(st.Get_buff())
}
