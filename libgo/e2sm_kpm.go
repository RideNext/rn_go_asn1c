package main

import (
	e2sm "ridenext.co.in/goasn1/e2sm_kpm"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {
   msgAD, _ := hex.DecodeString("20000102200000010037340300000000000001e04176657261676520444c205545207468726f75676870757420696e20674e42012000004001f3")
   msgIH, _ := hex.DecodeString("101ee9df44a5fae147ae485645525f325f335f3100000a4b504d5f5354594c45313870697474657374064e4543")
   st := e2sm.Stream{}
   fmt.Println(msgAD)
   st.Init(msgIH)
   pdu := e2sm.E2SMPDU{}
   pdu.Unpack(&st)
   jsonData, _ := json.Marshal(pdu)
   fmt.Printf("PDU = %s\n", jsonData)
}
