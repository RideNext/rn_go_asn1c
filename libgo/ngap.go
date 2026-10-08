package main

import (
	ngap "ridenext.co.in/go-asn1/ngap"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {
   //msg, _ := hex.DecodeString("00150025000003001b000800310410100000120066000d000000000000310410000000000015400100")
   //msg, _ := hex.DecodeString("000f403c00000400550002006900260018177e004179000d0164f6290000000010000011102e02e0e00079000e0031041000000000000000000000005a400100")
   //msg, _ := hex.DecodeString("000e0066000008000a00032003e8005500020069006e0006080186a00000001c00070041031000000000000002000000770009000000000000000000005e0020000000000000000000000000000000000000000000000000000000000000000000264006057e00420101")
   //msg, _ := hex.DecodeString("002e403d000004000a00032003e8005500034003e800260016157e00572d106a179eb117b1a36d9bb3893a204801c80079400e0000000000000000000000000000")
   //msg, _ := hex.DecodeString("200e0010000002000a40032003e8005540020000")
   msg, _ := hex.DecodeString("001d005e000004000a00032003e8005500020000002600403f2e0506c2110014050011223d01013d09110101010101010101220906040001040001290501c00000022514134f70656e52616469737973496e7465726e6574004a0006000001000000")
   st := ngap.Stream{}
   st1 := ngap.Stream{}
   fmt.Println(msg)
   st.Init(msg)
   psr := &ngap.PDUSessionResourceSetupRequestTransfer{}
   d5qi := &ngap.Dynamic5QIDescriptor{}
   qfsri := make([]ngap.QosFlowSetupRequestItem, 1)
   qfsri[0].QosFlowLevelQosParameters.QosCharacteristics.Dynamic5QI = d5qi
   psr.ProtocolIEs.QosFlowSetupRequestList.Items =  qfsri
   pdu := ngap.NGAPPDU{}
   rpdu := ngap.NGAPPDU{}
   req, pc, c := ngap.GetInitialUEMessageINITIATINGMESSAGE()
   rpdu.InitiatingMessage = &ngap.InitiatingMessage{}
   //ngap.NGAPELEMENTARYPROCEDUREprocedureCode{pc}, ngap.NGAPELEMENTARYPROCEDUREcriticality{c}, req}
   rpdu.InitiatingMessage.ProcedureCode = ngap.NGAPELEMENTARYPROCEDUREprocedureCode{pc}
   rpdu.InitiatingMessage.Criticality = ngap.NGAPELEMENTARYPROCEDUREcriticality{c}
   rpdu.InitiatingMessage.Value = req
   req.ProtocolIEs.RANUENGAPID.Value = 1000
   req.ProtocolIEs.NASPDU.Value = []byte {0x7e,0x00,0x00}
   req.ProtocolIEs.UserLocationInformation.UserLocationInformationNR = &ngap.UserLocationInformationNR{}
   req.ProtocolIEs.UserLocationInformation.UserLocationInformationNR.NRCGI.PLMNIdentity.Value=[]byte {0x00, 0xf1, 0x10}
   req.ProtocolIEs.UserLocationInformation.UserLocationInformationNR.NRCGI.NRCellIdentity.Len=22
   req.ProtocolIEs.UserLocationInformation.UserLocationInformationNR.NRCGI.NRCellIdentity.Value= []byte {0x11, 0x22, 0x33}
   req.ProtocolIEs.AMFSetID = &ngap.AMFSetID{}
   req.ProtocolIEs.UEContextRequest = &ngap.UEContextRequest{}
   pdu.Unpack(&st)
   switch pdutype := pdu.InitiatingMessage.Value.(type) {
       case *ngap.PDUSessionResourceSetupRequest:
          pdutype.ProtocolIEs.PDUSessionResourceSetupListSUReq.Items[0].PDUSessionResourceSetupRequestTransfer.Item = psr
	  fmt.Printf("ngap.PDUSessionResourceSetupRequest %+v\n", psr)
          pdu.InitiatingMessage.Value =  pdutype
       default:
	  fmt.Printf("\n%+V\n", pdutype)
	  break;
   }
   jsonData, _ := json.Marshal(rpdu)
   fmt.Printf("PDU = %s\n", jsonData)
   pdu.Pack(&st1)
   fmt.Println(st.Get_buff())
   fmt.Println(st1.Get_buff())
   st1.Reset()
   pdu.Unpack(&st1)
   jsonData, _ = json.Marshal(pdu)
   fmt.Printf("PDU = %s\n", jsonData)
   st3 := ngap.Stream{}
   pdu.Pack(&st3)
   fmt.Println(st3.Get_buff())
   buff := buildAndSendInitialUEMessage()
   st.Init(buff)
   fmt.Printf("buildAndSendInitialUEMessage %v\n", buff)
   pdu.Unpack(&st)
   jsonData, _ = json.Marshal(pdu)
   fmt.Printf("PDU = %s\n", jsonData)

}

func InitiaUEMessagePDU( ngapId uint64, nasBuff [] byte) ngap.NGAPPDU {   
   pdu := ngap.NGAPPDU{} // it creates a generic ngap PDU.
   msg, procedure, criticality := ngap.GetInitialUEMessageINITIATINGMESSAGE()
   pdu.InitiatingMessage = &ngap.InitiatingMessage{}
   pdu.InitiatingMessage.ProcedureCode = ngap.NGAPELEMENTARYPROCEDUREprocedureCode{procedure}
   pdu.InitiatingMessage.Criticality = ngap.NGAPELEMENTARYPROCEDUREcriticality{criticality}
   msg.ProtocolIEs.RANUENGAPID.Value = ngapId
   msg.ProtocolIEs.NASPDU.Value = nasBuff
   msg.ProtocolIEs.UserLocationInformation.UserLocationInformationNR = &ngap.UserLocationInformationNR{}
   msg.ProtocolIEs.UserLocationInformation.UserLocationInformationNR.NRCGI.PLMNIdentity.Value=[]byte {0x00, 0xf1, 0x10}
   msg.ProtocolIEs.UserLocationInformation.UserLocationInformationNR.NRCGI.NRCellIdentity.Len=22
   msg.ProtocolIEs.UserLocationInformation.UserLocationInformationNR.NRCGI.NRCellIdentity.Value= []byte {0x11, 0x22, 0x33, 0x44, 0x55}
   msg.ProtocolIEs.UserLocationInformation.UserLocationInformationNR.TAI.PLMNIdentity.Value=[]byte {0x00, 0xf1, 0x10}
   msg.ProtocolIEs.UserLocationInformation.UserLocationInformationNR.TAI.TAC.Value=[]byte {0xab, 0xcd, 0xef}
   msg.ProtocolIEs.AMFSetID = &ngap.AMFSetID{}
   msg.ProtocolIEs.AMFSetID.Value = []byte{0x10, 0xc0}
   msg.ProtocolIEs.UEContextRequest = &ngap.UEContextRequest{}
   msg.ProtocolIEs.UEContextRequest.Value = 5
   pdu.InitiatingMessage.Value = msg
   return pdu
}
func buildAndSendInitialUEMessage() []byte {
    st := ngap.Stream{}
    pdu := InitiaUEMessagePDU(1000, []byte{0x7e,0x05, 0x05})
    pdu.Pack(&st) 
    return st.Get_buff()
    //send buff out over transport
}

