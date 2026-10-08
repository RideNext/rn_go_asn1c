
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package e1ap
import (
  log "github.com/sirupsen/logrus"
)
var version = "vf40"

func fmte1ap() {log.Debug("e1ap")}
func (self *E1APPDU)Unpack(stream *Stream) {
    //coptions := []string{"initiatingMessage","successfulOutcome","unsuccessfulOutcome","Unknown"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in E1APPDU\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.InitiatingMessage = &InitiatingMessage{}//cho6
        self.InitiatingMessage.Unpack(stream)
    } else if choice == 1 { //ch2
        self.SuccessfulOutcome = &SuccessfulOutcome{}//cho6
        self.SuccessfulOutcome.Unpack(stream)
    } else if choice == 2 { //ch2
        self.UnsuccessfulOutcome = &UnsuccessfulOutcome{}//cho6
        self.UnsuccessfulOutcome.Unpack(stream)
    }//end of if else

    if choice >= 3 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * E1APPDU) Pack(stream *Stream) {
    if self.InitiatingMessage != nil {
        stream.set_choice(0, 2, 1, 3)
        self.InitiatingMessage.Pack(stream)//2
    } else if self.SuccessfulOutcome != nil {
        stream.set_choice(1, 2, 1, 3)
        self.SuccessfulOutcome.Pack(stream)//2
    } else if self.UnsuccessfulOutcome != nil {
        stream.set_choice(2, 2, 1, 3)
        self.UnsuccessfulOutcome.Pack(stream)//2
    }

}
type E1APPDU struct { //[{'type': 'InitiatingMessage', 'name': 'initiatingMessage'}, {'type': 'SuccessfulOutcome', 'name': 'successfulOutcome'}, {'type': 'UnsuccessfulOutcome', 'name': 'unsuccessfulOutcome'}, None]
    InitiatingMessage *InitiatingMessage
    SuccessfulOutcome *SuccessfulOutcome
    UnsuccessfulOutcome *UnsuccessfulOutcome
} // E1APPDU

type InitiatingMessage struct { // [{'type': 'E1AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E1AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'E1AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['E1AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'E1AP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['E1AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode E1APELEMENTARYPROCEDUREprocedureCode
    Criticality E1APELEMENTARYPROCEDUREcriticality
    Value E1APELEMENTARYPROCEDUREInitiatingMessage
}

func (self * InitiatingMessage) Unpack(stream *Stream) {
    //table {'type': 'E1AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E1AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_E1APELEMENTARYPROCEDURES[key].INITIATINGMESSAGE
    self.Value = out.(E1APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(E1APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'E1AP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['E1AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * InitiatingMessage) Pack(stream *Stream) {
    //table {'type': 'E1AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E1AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_E1APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(E1APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type SuccessfulOutcome struct { // [{'type': 'E1AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E1AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'E1AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['E1AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'E1AP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['E1AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode E1APELEMENTARYPROCEDUREprocedureCode
    Criticality E1APELEMENTARYPROCEDUREcriticality
    Value E1APELEMENTARYPROCEDURESuccessfulOutcome
}

func (self * SuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'E1AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E1AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_E1APELEMENTARYPROCEDURES[key].SUCCESSFULOUTCOME
    self.Value = out.(E1APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(E1APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'E1AP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['E1AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * SuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'E1AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E1AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_E1APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(E1APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type UnsuccessfulOutcome struct { // [{'type': 'E1AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E1AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'E1AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['E1AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'E1AP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['E1AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode E1APELEMENTARYPROCEDUREprocedureCode
    Criticality E1APELEMENTARYPROCEDUREcriticality
    Value E1APELEMENTARYPROCEDUREUnsuccessfulOutcome
}

func (self * UnsuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'E1AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E1AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_E1APELEMENTARYPROCEDURES[key].UNSUCCESSFULOUTCOME
    self.Value = out.(E1APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(E1APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'E1AP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['E1AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * UnsuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'E1AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E1AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_E1APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(E1APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type Reset struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ResetIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ResetIEs
}

func (self * Reset) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ResetIEs, order_ResetIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * Reset) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ResetIEs, order_ResetIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *ResetType)Unpack(stream *Stream) {
    //coptions := []string{"e1-Interface","partOfE1-Interface","choice-extension","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        self.E1Interface = &ResetAll{}//cho6
        self.E1Interface.Unpack(stream)
    } else if choice == 1 { //ch2
        self.PartOfE1Interface = &UEassociatedLogicalE1ConnectionListRes{}//cho6
        self.PartOfE1Interface.Unpack(stream)
    } else if choice == 2 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &ResetTypeExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * ResetType) Pack(stream *Stream) {
    if self.E1Interface != nil {
        stream.set_choice(0, 2, 0, 3)
        self.E1Interface.Pack(stream)//2
    } else if self.PartOfE1Interface != nil {
        stream.set_choice(1, 2, 0, 3)
        self.PartOfE1Interface.Pack(stream)//2
    } else if self.Choiceextension != nil {
        stream.set_choice(2, 2, 0, 3)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type ResetType struct { //[{'type': 'ResetAll', 'name': 'e1-Interface'}, {'type': 'UE-associatedLogicalE1-ConnectionListRes', 'name': 'partOfE1-Interface'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['ResetType-ExtIEs'], 'name': 'choice-extension'}]
    E1Interface *ResetAll
    PartOfE1Interface *UEassociatedLogicalE1ConnectionListRes
    Choiceextension *ResetTypeExtIEs
} // ResetType

type ResetAll struct {
  Value int
}
const (
    ResetAllreset_all = 0

    /* Extensions */
)
func (self *ResetAll) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *ResetAll) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
func (self *UEassociatedLogicalE1ConnectionListRes) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65536)
    _size += 1
    self.Items = make([]UEassociatedLogicalE1ConnectionItemRes, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *UEassociatedLogicalE1ConnectionListRes) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65536)
    //for item in table_UEassociatedLogicalE1ConnectionItemRes:
    val := ProtocolIESingleContainer{table_UEassociatedLogicalE1ConnectionItemRes, order_UEassociatedLogicalE1ConnectionItemRes}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type UEassociatedLogicalE1ConnectionListRes struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['UE-associatedLogicalE1-ConnectionItemRes']}, 'size': [(1, 'maxnoofIndividualE1ConnectionsToReset')]}
    Items []UEassociatedLogicalE1ConnectionItemRes
}

type ResetAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ResetAcknowledgeIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ResetAcknowledgeIEs
}

func (self * ResetAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ResetAcknowledgeIEs, order_ResetAcknowledgeIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ResetAcknowledge) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ResetAcknowledgeIEs, order_ResetAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *UEassociatedLogicalE1ConnectionListResAck) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65536)
    _size += 1
    self.Items = make([]UEassociatedLogicalE1ConnectionItemResAck, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *UEassociatedLogicalE1ConnectionListResAck) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65536)
    //for item in table_UEassociatedLogicalE1ConnectionItemResAck:
    val := ProtocolIESingleContainer{table_UEassociatedLogicalE1ConnectionItemResAck, order_UEassociatedLogicalE1ConnectionItemResAck}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type UEassociatedLogicalE1ConnectionListResAck struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['UE-associatedLogicalE1-ConnectionItemResAck']}, 'size': [(1, 'maxnoofIndividualE1ConnectionsToReset')]}
    Items []UEassociatedLogicalE1ConnectionItemResAck
}

type ErrorIndication struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ErrorIndication-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ErrorIndicationIEs
}

func (self * ErrorIndication) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ErrorIndicationIEs, order_ErrorIndicationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ErrorIndication) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ErrorIndicationIEs, order_ErrorIndicationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUUPE1SetupRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-UP-E1SetupRequestIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUUPE1SetupRequestIEs
}

func (self * GNBCUUPE1SetupRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUUPE1SetupRequestIEs, order_GNBCUUPE1SetupRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUUPE1SetupRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUUPE1SetupRequestIEs, order_GNBCUUPE1SetupRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *SupportedPLMNsList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(12)
    _size += 1
    self.Items = make([]SupportedPLMNsItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *SupportedPLMNsList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 12)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type SupportedPLMNsList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SupportedPLMNs-Item'}, 'size': [(1, 'maxnoofSPLMNs')]}
    Items []SupportedPLMNsItem
}

type SupportedPLMNsItem struct { // [{'type': 'PLMN-Identity', 'name': 'pLMN-Identity'}, {'type': 'Slice-Support-List', 'name': 'slice-Support-List', 'optional': True}, {'type': 'NR-CGI-Support-List', 'name': 'nR-CGI-Support-List', 'optional': True}, {'type': 'QoS-Parameters-Support-List', 'name': 'qoS-Parameters-Support-List', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SupportedPLMNs-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNIdentity PLMNIdentity
    SliceSupportList *SliceSupportList
    NRCGISupportList *NRCGISupportList
    QoSParametersSupportList *QoSParametersSupportList
    IEExtensions *SupportedPLMNsExtIEs
}

func (self * SupportedPLMNsItem) Unpack(stream *Stream) {
    sliceSupportList_flag := 0x00000002
    nRCGISupportList_flag := 0x00000004
    qoSParametersSupportList_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.PLMNIdentity.Unpack(stream)// p8
    if (sliceSupportList_flag & _flags) == sliceSupportList_flag { //cond2
        self.SliceSupportList = &SliceSupportList{}//7{'type': 'Slice-Support-List', 'name': 'slice-Support-List', 'optional': True}
        self.SliceSupportList.Unpack(stream)// p8
    }
    if (nRCGISupportList_flag & _flags) == nRCGISupportList_flag { //cond2
        self.NRCGISupportList = &NRCGISupportList{}//7{'type': 'NR-CGI-Support-List', 'name': 'nR-CGI-Support-List', 'optional': True}
        self.NRCGISupportList.Unpack(stream)// p8
    }
    if (qoSParametersSupportList_flag & _flags) == qoSParametersSupportList_flag { //cond2
        self.QoSParametersSupportList = &QoSParametersSupportList{}//7{'type': 'QoS-Parameters-Support-List', 'name': 'qoS-Parameters-Support-List', 'optional': True}
        self.QoSParametersSupportList.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SupportedPLMNsExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SupportedPLMNs-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SupportedPLMNsExtIEs, order_SupportedPLMNsExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SupportedPLMNsItem) Pack(stream *Stream) {
    const sliceSupportList_flag uint = 0x00000002
    const nRCGISupportList_flag uint = 0x00000004
    const qoSParametersSupportList_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    if self.SliceSupportList != nil { 
        _flags |= sliceSupportList_flag
        self.SliceSupportList.Pack(stream)
    }//end of optional
    if self.NRCGISupportList != nil { 
        _flags |= nRCGISupportList_flag
        self.NRCGISupportList.Pack(stream)
    }//end of optional
    if self.QoSParametersSupportList != nil { 
        _flags |= qoSParametersSupportList_flag
        self.QoSParametersSupportList.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SupportedPLMNsExtIEs, order_SupportedPLMNsExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

type GNBCUUPE1SetupResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-UP-E1SetupResponseIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUUPE1SetupResponseIEs
}

func (self * GNBCUUPE1SetupResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUUPE1SetupResponseIEs, order_GNBCUUPE1SetupResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUUPE1SetupResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUUPE1SetupResponseIEs, order_GNBCUUPE1SetupResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUUPE1SetupFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-UP-E1SetupFailureIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUUPE1SetupFailureIEs
}

func (self * GNBCUUPE1SetupFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUUPE1SetupFailureIEs, order_GNBCUUPE1SetupFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUUPE1SetupFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUUPE1SetupFailureIEs, order_GNBCUUPE1SetupFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUCPE1SetupRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-CP-E1SetupRequestIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUCPE1SetupRequestIEs
}

func (self * GNBCUCPE1SetupRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUCPE1SetupRequestIEs, order_GNBCUCPE1SetupRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUCPE1SetupRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUCPE1SetupRequestIEs, order_GNBCUCPE1SetupRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUCPE1SetupResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-CP-E1SetupResponseIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUCPE1SetupResponseIEs
}

func (self * GNBCUCPE1SetupResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUCPE1SetupResponseIEs, order_GNBCUCPE1SetupResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUCPE1SetupResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUCPE1SetupResponseIEs, order_GNBCUCPE1SetupResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUCPE1SetupFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-CP-E1SetupFailureIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUCPE1SetupFailureIEs
}

func (self * GNBCUCPE1SetupFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUCPE1SetupFailureIEs, order_GNBCUCPE1SetupFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUCPE1SetupFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUCPE1SetupFailureIEs, order_GNBCUCPE1SetupFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUUPConfigurationUpdate struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-UP-ConfigurationUpdateIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUUPConfigurationUpdateIEs
}

func (self * GNBCUUPConfigurationUpdate) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUUPConfigurationUpdateIEs, order_GNBCUUPConfigurationUpdateIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUUPConfigurationUpdate) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUUPConfigurationUpdateIEs, order_GNBCUUPConfigurationUpdateIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *GNBCUUPTNLAToRemoveList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]GNBCUUPTNLAToRemoveItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *GNBCUUPTNLAToRemoveList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type GNBCUUPTNLAToRemoveList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GNB-CU-UP-TNLA-To-Remove-Item'}, 'size': [(1, 'maxnoofTNLAssociations')]}
    Items []GNBCUUPTNLAToRemoveItem
}

type GNBCUUPConfigurationUpdateAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-UP-ConfigurationUpdateAcknowledgeIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUUPConfigurationUpdateAcknowledgeIEs
}

func (self * GNBCUUPConfigurationUpdateAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUUPConfigurationUpdateAcknowledgeIEs, order_GNBCUUPConfigurationUpdateAcknowledgeIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUUPConfigurationUpdateAcknowledge) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUUPConfigurationUpdateAcknowledgeIEs, order_GNBCUUPConfigurationUpdateAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUUPConfigurationUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-UP-ConfigurationUpdateFailureIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUUPConfigurationUpdateFailureIEs
}

func (self * GNBCUUPConfigurationUpdateFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUUPConfigurationUpdateFailureIEs, order_GNBCUUPConfigurationUpdateFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUUPConfigurationUpdateFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUUPConfigurationUpdateFailureIEs, order_GNBCUUPConfigurationUpdateFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUCPConfigurationUpdate struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-CP-ConfigurationUpdateIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUCPConfigurationUpdateIEs
}

func (self * GNBCUCPConfigurationUpdate) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUCPConfigurationUpdateIEs, order_GNBCUCPConfigurationUpdateIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUCPConfigurationUpdate) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUCPConfigurationUpdateIEs, order_GNBCUCPConfigurationUpdateIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *GNBCUCPTNLAToAddList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]GNBCUCPTNLAToAddItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *GNBCUCPTNLAToAddList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type GNBCUCPTNLAToAddList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GNB-CU-CP-TNLA-To-Add-Item'}, 'size': [(1, 'maxnoofTNLAssociations')]}
    Items []GNBCUCPTNLAToAddItem
}

func (self *GNBCUCPTNLAToRemoveList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]GNBCUCPTNLAToRemoveItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *GNBCUCPTNLAToRemoveList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type GNBCUCPTNLAToRemoveList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GNB-CU-CP-TNLA-To-Remove-Item'}, 'size': [(1, 'maxnoofTNLAssociations')]}
    Items []GNBCUCPTNLAToRemoveItem
}

func (self *GNBCUCPTNLAToUpdateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]GNBCUCPTNLAToUpdateItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *GNBCUCPTNLAToUpdateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type GNBCUCPTNLAToUpdateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GNB-CU-CP-TNLA-To-Update-Item'}, 'size': [(1, 'maxnoofTNLAssociations')]}
    Items []GNBCUCPTNLAToUpdateItem
}

type GNBCUCPConfigurationUpdateAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-CP-ConfigurationUpdateAcknowledgeIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUCPConfigurationUpdateAcknowledgeIEs
}

func (self * GNBCUCPConfigurationUpdateAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUCPConfigurationUpdateAcknowledgeIEs, order_GNBCUCPConfigurationUpdateAcknowledgeIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUCPConfigurationUpdateAcknowledge) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUCPConfigurationUpdateAcknowledgeIEs, order_GNBCUCPConfigurationUpdateAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *GNBCUCPTNLASetupList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]GNBCUCPTNLASetupItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *GNBCUCPTNLASetupList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type GNBCUCPTNLASetupList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GNB-CU-CP-TNLA-Setup-Item'}, 'size': [(1, 'maxnoofTNLAssociations')]}
    Items []GNBCUCPTNLASetupItem
}

func (self *GNBCUCPTNLAFailedToSetupList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]GNBCUCPTNLAFailedToSetupItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *GNBCUCPTNLAFailedToSetupList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type GNBCUCPTNLAFailedToSetupList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GNB-CU-CP-TNLA-Failed-To-Setup-Item'}, 'size': [(1, 'maxnoofTNLAssociations')]}
    Items []GNBCUCPTNLAFailedToSetupItem
}

type GNBCUCPConfigurationUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-CP-ConfigurationUpdateFailureIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUCPConfigurationUpdateFailureIEs
}

func (self * GNBCUCPConfigurationUpdateFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUCPConfigurationUpdateFailureIEs, order_GNBCUCPConfigurationUpdateFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUCPConfigurationUpdateFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUCPConfigurationUpdateFailureIEs, order_GNBCUCPConfigurationUpdateFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E1ReleaseRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E1ReleaseRequestIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E1ReleaseRequestIEs
}

func (self * E1ReleaseRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E1ReleaseRequestIEs, order_E1ReleaseRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E1ReleaseRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E1ReleaseRequestIEs, order_E1ReleaseRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E1ReleaseResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E1ReleaseResponseIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E1ReleaseResponseIEs
}

func (self * E1ReleaseResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E1ReleaseResponseIEs, order_E1ReleaseResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E1ReleaseResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E1ReleaseResponseIEs, order_E1ReleaseResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type BearerContextSetupRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextSetupRequestIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextSetupRequestIEs
}

func (self * BearerContextSetupRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextSetupRequestIEs, order_BearerContextSetupRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextSetupRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextSetupRequestIEs, order_BearerContextSetupRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *SystemBearerContextSetupRequest)Unpack(stream *Stream) {
    //coptions := []string{"e-UTRAN-BearerContextSetupRequest","nG-RAN-BearerContextSetupRequest","choice-extension","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        EUTRANBearerContextSetupRequest := &ProtocolIEContainer{}//cho2
        self.EUTRANBearerContextSetupRequest = &EUTRANBearerContextSetupRequest{}//cho3
        EUTRANBearerContextSetupRequest.Unpack(stream, self.EUTRANBearerContextSetupRequest)
    } else if choice == 1 { //ch2
        NGRANBearerContextSetupRequest := &ProtocolIEContainer{}//cho2
        self.NGRANBearerContextSetupRequest = &NGRANBearerContextSetupRequest{}//cho3
        NGRANBearerContextSetupRequest.Unpack(stream, self.NGRANBearerContextSetupRequest)
    } else if choice == 2 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &SystemBearerContextSetupRequestExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * SystemBearerContextSetupRequest) Pack(stream *Stream) {
    if self.EUTRANBearerContextSetupRequest != nil {
        stream.set_choice(0, 2, 0, 3)
        EUTRANBearerContextSetupRequest := ProtocolIEContainer{}//cho2
        EUTRANBearerContextSetupRequest.Pack(stream, self.EUTRANBearerContextSetupRequest)
    } else if self.NGRANBearerContextSetupRequest != nil {
        stream.set_choice(1, 2, 0, 3)
        NGRANBearerContextSetupRequest := ProtocolIEContainer{}//cho2
        NGRANBearerContextSetupRequest.Pack(stream, self.NGRANBearerContextSetupRequest)
    } else if self.Choiceextension != nil {
        stream.set_choice(2, 2, 0, 3)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type SystemBearerContextSetupRequest struct { //[{'type': 'ProtocolIE-Container', 'actual-parameters': ['EUTRAN-BearerContextSetupRequest'], 'name': 'e-UTRAN-BearerContextSetupRequest'}, {'type': 'ProtocolIE-Container', 'actual-parameters': ['NG-RAN-BearerContextSetupRequest'], 'name': 'nG-RAN-BearerContextSetupRequest'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['System-BearerContextSetupRequest-ExtIEs'], 'name': 'choice-extension'}]
    EUTRANBearerContextSetupRequest *EUTRANBearerContextSetupRequest
    NGRANBearerContextSetupRequest *NGRANBearerContextSetupRequest
    Choiceextension *SystemBearerContextSetupRequestExtIEs
} // SystemBearerContextSetupRequest

type BearerContextSetupResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextSetupResponseIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextSetupResponseIEs
}

func (self * BearerContextSetupResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextSetupResponseIEs, order_BearerContextSetupResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextSetupResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextSetupResponseIEs, order_BearerContextSetupResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *SystemBearerContextSetupResponse)Unpack(stream *Stream) {
    //coptions := []string{"e-UTRAN-BearerContextSetupResponse","nG-RAN-BearerContextSetupResponse","choice-extension","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        EUTRANBearerContextSetupResponse := &ProtocolIEContainer{}//cho2
        self.EUTRANBearerContextSetupResponse = &EUTRANBearerContextSetupResponse{}//cho3
        EUTRANBearerContextSetupResponse.Unpack(stream, self.EUTRANBearerContextSetupResponse)
    } else if choice == 1 { //ch2
        NGRANBearerContextSetupResponse := &ProtocolIEContainer{}//cho2
        self.NGRANBearerContextSetupResponse = &NGRANBearerContextSetupResponse{}//cho3
        NGRANBearerContextSetupResponse.Unpack(stream, self.NGRANBearerContextSetupResponse)
    } else if choice == 2 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &SystemBearerContextSetupResponseExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * SystemBearerContextSetupResponse) Pack(stream *Stream) {
    if self.EUTRANBearerContextSetupResponse != nil {
        stream.set_choice(0, 2, 0, 3)
        EUTRANBearerContextSetupResponse := ProtocolIEContainer{}//cho2
        EUTRANBearerContextSetupResponse.Pack(stream, self.EUTRANBearerContextSetupResponse)
    } else if self.NGRANBearerContextSetupResponse != nil {
        stream.set_choice(1, 2, 0, 3)
        NGRANBearerContextSetupResponse := ProtocolIEContainer{}//cho2
        NGRANBearerContextSetupResponse.Pack(stream, self.NGRANBearerContextSetupResponse)
    } else if self.Choiceextension != nil {
        stream.set_choice(2, 2, 0, 3)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type SystemBearerContextSetupResponse struct { //[{'type': 'ProtocolIE-Container', 'actual-parameters': ['EUTRAN-BearerContextSetupResponse'], 'name': 'e-UTRAN-BearerContextSetupResponse'}, {'type': 'ProtocolIE-Container', 'actual-parameters': ['NG-RAN-BearerContextSetupResponse'], 'name': 'nG-RAN-BearerContextSetupResponse'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['System-BearerContextSetupResponse-ExtIEs'], 'name': 'choice-extension'}]
    EUTRANBearerContextSetupResponse *EUTRANBearerContextSetupResponse
    NGRANBearerContextSetupResponse *NGRANBearerContextSetupResponse
    Choiceextension *SystemBearerContextSetupResponseExtIEs
} // SystemBearerContextSetupResponse

type BearerContextSetupFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextSetupFailureIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextSetupFailureIEs
}

func (self * BearerContextSetupFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextSetupFailureIEs, order_BearerContextSetupFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextSetupFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextSetupFailureIEs, order_BearerContextSetupFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type BearerContextModificationRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextModificationRequestIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextModificationRequestIEs
}

func (self * BearerContextModificationRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextModificationRequestIEs, order_BearerContextModificationRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextModificationRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextModificationRequestIEs, order_BearerContextModificationRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *SystemBearerContextModificationRequest)Unpack(stream *Stream) {
    //coptions := []string{"e-UTRAN-BearerContextModificationRequest","nG-RAN-BearerContextModificationRequest","choice-extension","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        EUTRANBearerContextModificationRequest := &ProtocolIEContainer{}//cho2
        self.EUTRANBearerContextModificationRequest = &EUTRANBearerContextModificationRequest{}//cho3
        EUTRANBearerContextModificationRequest.Unpack(stream, self.EUTRANBearerContextModificationRequest)
    } else if choice == 1 { //ch2
        NGRANBearerContextModificationRequest := &ProtocolIEContainer{}//cho2
        self.NGRANBearerContextModificationRequest = &NGRANBearerContextModificationRequest{}//cho3
        NGRANBearerContextModificationRequest.Unpack(stream, self.NGRANBearerContextModificationRequest)
    } else if choice == 2 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &SystemBearerContextModificationRequestExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * SystemBearerContextModificationRequest) Pack(stream *Stream) {
    if self.EUTRANBearerContextModificationRequest != nil {
        stream.set_choice(0, 2, 0, 3)
        EUTRANBearerContextModificationRequest := ProtocolIEContainer{}//cho2
        EUTRANBearerContextModificationRequest.Pack(stream, self.EUTRANBearerContextModificationRequest)
    } else if self.NGRANBearerContextModificationRequest != nil {
        stream.set_choice(1, 2, 0, 3)
        NGRANBearerContextModificationRequest := ProtocolIEContainer{}//cho2
        NGRANBearerContextModificationRequest.Pack(stream, self.NGRANBearerContextModificationRequest)
    } else if self.Choiceextension != nil {
        stream.set_choice(2, 2, 0, 3)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type SystemBearerContextModificationRequest struct { //[{'type': 'ProtocolIE-Container', 'actual-parameters': ['EUTRAN-BearerContextModificationRequest'], 'name': 'e-UTRAN-BearerContextModificationRequest'}, {'type': 'ProtocolIE-Container', 'actual-parameters': ['NG-RAN-BearerContextModificationRequest'], 'name': 'nG-RAN-BearerContextModificationRequest'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['System-BearerContextModificationRequest-ExtIEs'], 'name': 'choice-extension'}]
    EUTRANBearerContextModificationRequest *EUTRANBearerContextModificationRequest
    NGRANBearerContextModificationRequest *NGRANBearerContextModificationRequest
    Choiceextension *SystemBearerContextModificationRequestExtIEs
} // SystemBearerContextModificationRequest

type BearerContextModificationResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextModificationResponseIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextModificationResponseIEs
}

func (self * BearerContextModificationResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextModificationResponseIEs, order_BearerContextModificationResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextModificationResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextModificationResponseIEs, order_BearerContextModificationResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *SystemBearerContextModificationResponse)Unpack(stream *Stream) {
    //coptions := []string{"e-UTRAN-BearerContextModificationResponse","nG-RAN-BearerContextModificationResponse","choice-extension","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        EUTRANBearerContextModificationResponse := &ProtocolIEContainer{}//cho2
        self.EUTRANBearerContextModificationResponse = &EUTRANBearerContextModificationResponse{}//cho3
        EUTRANBearerContextModificationResponse.Unpack(stream, self.EUTRANBearerContextModificationResponse)
    } else if choice == 1 { //ch2
        NGRANBearerContextModificationResponse := &ProtocolIEContainer{}//cho2
        self.NGRANBearerContextModificationResponse = &NGRANBearerContextModificationResponse{}//cho3
        NGRANBearerContextModificationResponse.Unpack(stream, self.NGRANBearerContextModificationResponse)
    } else if choice == 2 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &SystemBearerContextModificationResponseExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * SystemBearerContextModificationResponse) Pack(stream *Stream) {
    if self.EUTRANBearerContextModificationResponse != nil {
        stream.set_choice(0, 2, 0, 3)
        EUTRANBearerContextModificationResponse := ProtocolIEContainer{}//cho2
        EUTRANBearerContextModificationResponse.Pack(stream, self.EUTRANBearerContextModificationResponse)
    } else if self.NGRANBearerContextModificationResponse != nil {
        stream.set_choice(1, 2, 0, 3)
        NGRANBearerContextModificationResponse := ProtocolIEContainer{}//cho2
        NGRANBearerContextModificationResponse.Pack(stream, self.NGRANBearerContextModificationResponse)
    } else if self.Choiceextension != nil {
        stream.set_choice(2, 2, 0, 3)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type SystemBearerContextModificationResponse struct { //[{'type': 'ProtocolIE-Container', 'actual-parameters': ['EUTRAN-BearerContextModificationResponse'], 'name': 'e-UTRAN-BearerContextModificationResponse'}, {'type': 'ProtocolIE-Container', 'actual-parameters': ['NG-RAN-BearerContextModificationResponse'], 'name': 'nG-RAN-BearerContextModificationResponse'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['System-BearerContextModificationResponse-ExtIEs'], 'name': 'choice-extension'}]
    EUTRANBearerContextModificationResponse *EUTRANBearerContextModificationResponse
    NGRANBearerContextModificationResponse *NGRANBearerContextModificationResponse
    Choiceextension *SystemBearerContextModificationResponseExtIEs
} // SystemBearerContextModificationResponse

type BearerContextModificationFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextModificationFailureIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextModificationFailureIEs
}

func (self * BearerContextModificationFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextModificationFailureIEs, order_BearerContextModificationFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextModificationFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextModificationFailureIEs, order_BearerContextModificationFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type BearerContextModificationRequired struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextModificationRequiredIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextModificationRequiredIEs
}

func (self * BearerContextModificationRequired) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextModificationRequiredIEs, order_BearerContextModificationRequiredIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextModificationRequired) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextModificationRequiredIEs, order_BearerContextModificationRequiredIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *SystemBearerContextModificationRequired)Unpack(stream *Stream) {
    //coptions := []string{"e-UTRAN-BearerContextModificationRequired","nG-RAN-BearerContextModificationRequired","choice-extension","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        EUTRANBearerContextModificationRequired := &ProtocolIEContainer{}//cho2
        self.EUTRANBearerContextModificationRequired = &EUTRANBearerContextModificationRequired{}//cho3
        EUTRANBearerContextModificationRequired.Unpack(stream, self.EUTRANBearerContextModificationRequired)
    } else if choice == 1 { //ch2
        NGRANBearerContextModificationRequired := &ProtocolIEContainer{}//cho2
        self.NGRANBearerContextModificationRequired = &NGRANBearerContextModificationRequired{}//cho3
        NGRANBearerContextModificationRequired.Unpack(stream, self.NGRANBearerContextModificationRequired)
    } else if choice == 2 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &SystemBearerContextModificationRequiredExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * SystemBearerContextModificationRequired) Pack(stream *Stream) {
    if self.EUTRANBearerContextModificationRequired != nil {
        stream.set_choice(0, 2, 0, 3)
        EUTRANBearerContextModificationRequired := ProtocolIEContainer{}//cho2
        EUTRANBearerContextModificationRequired.Pack(stream, self.EUTRANBearerContextModificationRequired)
    } else if self.NGRANBearerContextModificationRequired != nil {
        stream.set_choice(1, 2, 0, 3)
        NGRANBearerContextModificationRequired := ProtocolIEContainer{}//cho2
        NGRANBearerContextModificationRequired.Pack(stream, self.NGRANBearerContextModificationRequired)
    } else if self.Choiceextension != nil {
        stream.set_choice(2, 2, 0, 3)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type SystemBearerContextModificationRequired struct { //[{'type': 'ProtocolIE-Container', 'actual-parameters': ['EUTRAN-BearerContextModificationRequired'], 'name': 'e-UTRAN-BearerContextModificationRequired'}, {'type': 'ProtocolIE-Container', 'actual-parameters': ['NG-RAN-BearerContextModificationRequired'], 'name': 'nG-RAN-BearerContextModificationRequired'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['System-BearerContextModificationRequired-ExtIEs'], 'name': 'choice-extension'}]
    EUTRANBearerContextModificationRequired *EUTRANBearerContextModificationRequired
    NGRANBearerContextModificationRequired *NGRANBearerContextModificationRequired
    Choiceextension *SystemBearerContextModificationRequiredExtIEs
} // SystemBearerContextModificationRequired

type BearerContextModificationConfirm struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextModificationConfirmIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextModificationConfirmIEs
}

func (self * BearerContextModificationConfirm) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextModificationConfirmIEs, order_BearerContextModificationConfirmIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextModificationConfirm) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextModificationConfirmIEs, order_BearerContextModificationConfirmIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *SystemBearerContextModificationConfirm)Unpack(stream *Stream) {
    //coptions := []string{"e-UTRAN-BearerContextModificationConfirm","nG-RAN-BearerContextModificationConfirm","choice-extension","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        EUTRANBearerContextModificationConfirm := &ProtocolIEContainer{}//cho2
        self.EUTRANBearerContextModificationConfirm = &EUTRANBearerContextModificationConfirm{}//cho3
        EUTRANBearerContextModificationConfirm.Unpack(stream, self.EUTRANBearerContextModificationConfirm)
    } else if choice == 1 { //ch2
        NGRANBearerContextModificationConfirm := &ProtocolIEContainer{}//cho2
        self.NGRANBearerContextModificationConfirm = &NGRANBearerContextModificationConfirm{}//cho3
        NGRANBearerContextModificationConfirm.Unpack(stream, self.NGRANBearerContextModificationConfirm)
    } else if choice == 2 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &SystemBearerContextModificationConfirmExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * SystemBearerContextModificationConfirm) Pack(stream *Stream) {
    if self.EUTRANBearerContextModificationConfirm != nil {
        stream.set_choice(0, 2, 0, 3)
        EUTRANBearerContextModificationConfirm := ProtocolIEContainer{}//cho2
        EUTRANBearerContextModificationConfirm.Pack(stream, self.EUTRANBearerContextModificationConfirm)
    } else if self.NGRANBearerContextModificationConfirm != nil {
        stream.set_choice(1, 2, 0, 3)
        NGRANBearerContextModificationConfirm := ProtocolIEContainer{}//cho2
        NGRANBearerContextModificationConfirm.Pack(stream, self.NGRANBearerContextModificationConfirm)
    } else if self.Choiceextension != nil {
        stream.set_choice(2, 2, 0, 3)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type SystemBearerContextModificationConfirm struct { //[{'type': 'ProtocolIE-Container', 'actual-parameters': ['EUTRAN-BearerContextModificationConfirm'], 'name': 'e-UTRAN-BearerContextModificationConfirm'}, {'type': 'ProtocolIE-Container', 'actual-parameters': ['NG-RAN-BearerContextModificationConfirm'], 'name': 'nG-RAN-BearerContextModificationConfirm'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['System-BearerContextModificationConfirm-ExtIEs'], 'name': 'choice-extension'}]
    EUTRANBearerContextModificationConfirm *EUTRANBearerContextModificationConfirm
    NGRANBearerContextModificationConfirm *NGRANBearerContextModificationConfirm
    Choiceextension *SystemBearerContextModificationConfirmExtIEs
} // SystemBearerContextModificationConfirm

type BearerContextReleaseCommand struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextReleaseCommandIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextReleaseCommandIEs
}

func (self * BearerContextReleaseCommand) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextReleaseCommandIEs, order_BearerContextReleaseCommandIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextReleaseCommand) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextReleaseCommandIEs, order_BearerContextReleaseCommandIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type BearerContextReleaseComplete struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextReleaseCompleteIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextReleaseCompleteIEs
}

func (self * BearerContextReleaseComplete) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextReleaseCompleteIEs, order_BearerContextReleaseCompleteIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextReleaseComplete) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextReleaseCompleteIEs, order_BearerContextReleaseCompleteIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type BearerContextReleaseRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextReleaseRequestIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextReleaseRequestIEs
}

func (self * BearerContextReleaseRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextReleaseRequestIEs, order_BearerContextReleaseRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextReleaseRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextReleaseRequestIEs, order_BearerContextReleaseRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *DRBStatusList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBStatusItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBStatusList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBStatusList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Status-Item'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBStatusItem
}

type BearerContextInactivityNotification struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['BearerContextInactivityNotificationIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs BearerContextInactivityNotificationIEs
}

func (self * BearerContextInactivityNotification) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_BearerContextInactivityNotificationIEs, order_BearerContextInactivityNotificationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BearerContextInactivityNotification) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_BearerContextInactivityNotificationIEs, order_BearerContextInactivityNotificationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type DLDataNotification struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['DLDataNotificationIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs DLDataNotificationIEs
}

func (self * DLDataNotification) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_DLDataNotificationIEs, order_DLDataNotificationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DLDataNotification) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_DLDataNotificationIEs, order_DLDataNotificationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type ULDataNotification struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ULDataNotificationIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ULDataNotificationIEs
}

func (self * ULDataNotification) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ULDataNotificationIEs, order_ULDataNotificationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ULDataNotification) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ULDataNotificationIEs, order_ULDataNotificationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type DataUsageReport struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['DataUsageReportIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs DataUsageReportIEs
}

func (self * DataUsageReport) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_DataUsageReportIEs, order_DataUsageReportIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DataUsageReport) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_DataUsageReportIEs, order_DataUsageReportIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUUPCounterCheckRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-UP-CounterCheckRequestIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUUPCounterCheckRequestIEs
}

func (self * GNBCUUPCounterCheckRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUUPCounterCheckRequestIEs, order_GNBCUUPCounterCheckRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUUPCounterCheckRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUUPCounterCheckRequestIEs, order_GNBCUUPCounterCheckRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *SystemGNBCUUPCounterCheckRequest)Unpack(stream *Stream) {
    //coptions := []string{"e-UTRAN-GNB-CU-UP-CounterCheckRequest","nG-RAN-GNB-CU-UP-CounterCheckRequest","choice-extension","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        EUTRANGNBCUUPCounterCheckRequest := &ProtocolIEContainer{}//cho2
        self.EUTRANGNBCUUPCounterCheckRequest = &EUTRANGNBCUUPCounterCheckRequest{}//cho3
        EUTRANGNBCUUPCounterCheckRequest.Unpack(stream, self.EUTRANGNBCUUPCounterCheckRequest)
    } else if choice == 1 { //ch2
        NGRANGNBCUUPCounterCheckRequest := &ProtocolIEContainer{}//cho2
        self.NGRANGNBCUUPCounterCheckRequest = &NGRANGNBCUUPCounterCheckRequest{}//cho3
        NGRANGNBCUUPCounterCheckRequest.Unpack(stream, self.NGRANGNBCUUPCounterCheckRequest)
    } else if choice == 2 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &SystemGNBCUUPCounterCheckRequestExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * SystemGNBCUUPCounterCheckRequest) Pack(stream *Stream) {
    if self.EUTRANGNBCUUPCounterCheckRequest != nil {
        stream.set_choice(0, 2, 0, 3)
        EUTRANGNBCUUPCounterCheckRequest := ProtocolIEContainer{}//cho2
        EUTRANGNBCUUPCounterCheckRequest.Pack(stream, self.EUTRANGNBCUUPCounterCheckRequest)
    } else if self.NGRANGNBCUUPCounterCheckRequest != nil {
        stream.set_choice(1, 2, 0, 3)
        NGRANGNBCUUPCounterCheckRequest := ProtocolIEContainer{}//cho2
        NGRANGNBCUUPCounterCheckRequest.Pack(stream, self.NGRANGNBCUUPCounterCheckRequest)
    } else if self.Choiceextension != nil {
        stream.set_choice(2, 2, 0, 3)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type SystemGNBCUUPCounterCheckRequest struct { //[{'type': 'ProtocolIE-Container', 'actual-parameters': ['EUTRAN-GNB-CU-UP-CounterCheckRequest'], 'name': 'e-UTRAN-GNB-CU-UP-CounterCheckRequest'}, {'type': 'ProtocolIE-Container', 'actual-parameters': ['NG-RAN-GNB-CU-UP-CounterCheckRequest'], 'name': 'nG-RAN-GNB-CU-UP-CounterCheckRequest'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['System-GNB-CU-UP-CounterCheckRequest-ExtIEs'], 'name': 'choice-extension'}]
    EUTRANGNBCUUPCounterCheckRequest *EUTRANGNBCUUPCounterCheckRequest
    NGRANGNBCUUPCounterCheckRequest *NGRANGNBCUUPCounterCheckRequest
    Choiceextension *SystemGNBCUUPCounterCheckRequestExtIEs
} // SystemGNBCUUPCounterCheckRequest

type GNBCUUPStatusIndication struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['GNB-CU-UP-StatusIndicationIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs GNBCUUPStatusIndicationIEs
}

func (self * GNBCUUPStatusIndication) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_GNBCUUPStatusIndicationIEs, order_GNBCUUPStatusIndicationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUUPStatusIndication) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_GNBCUUPStatusIndicationIEs, order_GNBCUUPStatusIndicationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MRDCDataUsageReport struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MRDC-DataUsageReportIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MRDCDataUsageReportIEs
}

func (self * MRDCDataUsageReport) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MRDCDataUsageReportIEs, order_MRDCDataUsageReportIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MRDCDataUsageReport) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MRDCDataUsageReportIEs, order_MRDCDataUsageReportIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type PrivateMessage struct { // [{'type': 'PrivateIE-Container', 'actual-parameters': ['PrivateMessage-IEs'], 'name': 'privateIEs'}, None]
    PrivateIEs PrivateMessageIEs
}

func (self * PrivateMessage) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    PrivateIEs := PrivateIEContainer {table_PrivateMessageIEs, order_PrivateMessageIEs} // p3
    PrivateIEs.Unpack(stream, &self.PrivateIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PrivateMessage) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    PrivateIEs := &PrivateIEContainer {table_PrivateMessageIEs, order_PrivateMessageIEs} // p3
    PrivateIEs.Pack(stream, &self.PrivateIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *ActivityInformation)Unpack(stream *Stream) {
    //coptions := []string{"dRB-Activity-List","pDU-Session-Resource-Activity-List","uE-Activity","choice-extension"}
    choice := stream.get_choice(2, 0, 4)
    if choice == 0 { //ch1
        self.DRBActivityList = &DRBActivityList{}//cho6
        self.DRBActivityList.Unpack(stream)
    } else if choice == 1 { //ch2
        self.PDUSessionResourceActivityList = &PDUSessionResourceActivityList{}//cho6
        self.PDUSessionResourceActivityList.Unpack(stream)
    } else if choice == 2 { //ch2
        self.UEActivity = &UEActivity{}//cho6
        self.UEActivity.Unpack(stream)
    } else if choice == 3 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &ActivityInformationExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * ActivityInformation) Pack(stream *Stream) {
    if self.DRBActivityList != nil {
        stream.set_choice(0, 2, 0, 4)
        self.DRBActivityList.Pack(stream)//2
    } else if self.PDUSessionResourceActivityList != nil {
        stream.set_choice(1, 2, 0, 4)
        self.PDUSessionResourceActivityList.Pack(stream)//2
    } else if self.UEActivity != nil {
        stream.set_choice(2, 2, 0, 4)
        self.UEActivity.Pack(stream)//2
    } else if self.Choiceextension != nil {
        stream.set_choice(3, 2, 0, 4)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type ActivityInformation struct { //[{'type': 'DRB-Activity-List', 'name': 'dRB-Activity-List'}, {'type': 'PDU-Session-Resource-Activity-List', 'name': 'pDU-Session-Resource-Activity-List'}, {'type': 'UE-Activity', 'name': 'uE-Activity'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['ActivityInformation-ExtIEs'], 'name': 'choice-extension'}]
    DRBActivityList *DRBActivityList
    PDUSessionResourceActivityList *PDUSessionResourceActivityList
    UEActivity *UEActivity
    Choiceextension *ActivityInformationExtIEs
} // ActivityInformation

type ActivityNotificationLevel struct {
  Value int
}
const (
    ActivityNotificationLeveldrb = 0
    ActivityNotificationLevelpdu_session = 1
    ActivityNotificationLevelue = 2

    /* Extensions */
)
func (self *ActivityNotificationLevel) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *ActivityNotificationLevel) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type AveragingWindow struct {
  Value uint64
}
func (self *AveragingWindow) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4096, 13, 1, 0)
}
func (self * AveragingWindow) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4096, 13, 1, 0)
}
type BearerContextStatusChange struct {
  Value int
}
const (
    BearerContextStatusChangesuspend = 0
    BearerContextStatusChangeresume = 1

    /* Extensions */
)
func (self *BearerContextStatusChange) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *BearerContextStatusChange) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type BitRate struct {
  Value uint64
}
func (self *BitRate) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4000000000001, 43, 1, 0)
}
func (self * BitRate) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4000000000001, 43, 1, 0)
}
func (self *Cause)Unpack(stream *Stream) {
    //coptions := []string{"radioNetwork","transport","protocol","misc","choice-extension","Unknown","Unknown","Unknown"}
    choice := stream.get_choice(3, 0, 5)
    if choice == 0 { //ch1
        self.RadioNetwork = &CauseRadioNetwork{}//cho6
        self.RadioNetwork.Unpack(stream)
    } else if choice == 1 { //ch2
        self.Transport = &CauseTransport{}//cho6
        self.Transport.Unpack(stream)
    } else if choice == 2 { //ch2
        self.Protocol = &CauseProtocol{}//cho6
        self.Protocol.Unpack(stream)
    } else if choice == 3 { //ch2
        self.Misc = &CauseMisc{}//cho6
        self.Misc.Unpack(stream)
    } else if choice == 4 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &CauseExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * Cause) Pack(stream *Stream) {
    if self.RadioNetwork != nil {
        stream.set_choice(0, 3, 0, 5)
        self.RadioNetwork.Pack(stream)//2
    } else if self.Transport != nil {
        stream.set_choice(1, 3, 0, 5)
        self.Transport.Pack(stream)//2
    } else if self.Protocol != nil {
        stream.set_choice(2, 3, 0, 5)
        self.Protocol.Pack(stream)//2
    } else if self.Misc != nil {
        stream.set_choice(3, 3, 0, 5)
        self.Misc.Pack(stream)//2
    } else if self.Choiceextension != nil {
        stream.set_choice(4, 3, 0, 5)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type Cause struct { //[{'type': 'CauseRadioNetwork', 'name': 'radioNetwork'}, {'type': 'CauseTransport', 'name': 'transport'}, {'type': 'CauseProtocol', 'name': 'protocol'}, {'type': 'CauseMisc', 'name': 'misc'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['Cause-ExtIEs'], 'name': 'choice-extension'}]
    RadioNetwork *CauseRadioNetwork
    Transport *CauseTransport
    Protocol *CauseProtocol
    Misc *CauseMisc
    Choiceextension *CauseExtIEs
} // Cause

type CauseMisc struct {
  Value int
}
const (
    CauseMisccontrol_processing_overload = 0
    CauseMiscnot_enough_user_plane_processing_resources = 1
    CauseMischardware_failure = 2
    CauseMiscom_intervention = 3
    CauseMiscunspecified = 4

    /* Extensions */
)
func (self *CauseMisc) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 5, 1)
}
func (self *CauseMisc) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 5, 1)
}
type CauseProtocol struct {
  Value int
}
const (
    CauseProtocoltransfer_syntax_error = 0
    CauseProtocolabstract_syntax_error_reject = 1
    CauseProtocolabstract_syntax_error_ignore_and_notify = 2
    CauseProtocolmessage_not_compatible_with_receiver_state = 3
    CauseProtocolsemantic_error = 4
    CauseProtocolabstract_syntax_error_fAlsely_constructed_message = 5
    CauseProtocolunspecified = 6

    /* Extensions */
)
func (self *CauseProtocol) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 7, 1)
}
func (self *CauseProtocol) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 7, 1)
}
type CauseRadioNetwork struct {
  Value int
}
const (
    CauseRadioNetworkunspecified = 0
    CauseRadioNetworkunknown_or_already_allocated_gnb_cu_cp_ue_e1ap_id = 1
    CauseRadioNetworkunknown_or_already_allocated_gnb_cu_up_ue_e1ap_id = 2
    CauseRadioNetworkunknown_or_inconsistent_pair_of_ue_e1ap_id = 3
    CauseRadioNetworkinteraction_with_other_procedure = 4
    CauseRadioNetworkpPDCP_Count_wrap_around = 5
    CauseRadioNetworknot_supported_QCI_value = 6
    CauseRadioNetworknot_supported_5QI_value = 7
    CauseRadioNetworkencryption_algorithms_not_supported = 8
    CauseRadioNetworkintegrity_protection_algorithms_not_supported = 9
    CauseRadioNetworkuP_integrity_protection_not_possible = 10
    CauseRadioNetworkuP_confidentiality_protection_not_possible = 11
    CauseRadioNetworkmultiple_PDU_Session_ID_Instances = 12
    CauseRadioNetworkunknown_PDU_Session_ID = 13
    CauseRadioNetworkmultiple_QoS_Flow_ID_Instances = 14
    CauseRadioNetworkunknown_QoS_Flow_ID = 15
    CauseRadioNetworkmultiple_DRB_ID_Instances = 16
    CauseRadioNetworkunknown_DRB_ID = 17
    CauseRadioNetworkinvalid_QoS_combination = 18
    CauseRadioNetworkprocedure_cancelled = 19
    CauseRadioNetworknormal_release = 20
    CauseRadioNetworkno_radio_resources_available = 21
    CauseRadioNetworkaction_desirable_for_radio_reasons = 22
    CauseRadioNetworkresources_not_available_for_the_slice = 23
    CauseRadioNetworkpDCP_configuration_not_supported = 24

    /* Extensions */
    CauseRadioNetworkue_dl_max_IP_data_rate_reason = 25
    CauseRadioNetworkuP_integrity_protection_failure = 26
    CauseRadioNetworkrelease_due_to_pre_emption = 27
)
func (self *CauseRadioNetwork) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(6, 25, 1)
}
func (self *CauseRadioNetwork) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 6, 25, 1)
}
type CauseTransport struct {
  Value int
}
const (
    CauseTransportunspecified = 0
    CauseTransporttransport_resource_unavailable = 1

    /* Extensions */
)
func (self *CauseTransport) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *CauseTransport) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
func (self *CellGroupInformation) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(4)
    _size += 1
    self.Items = make([]CellGroupInformationItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *CellGroupInformation) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 4)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type CellGroupInformation struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'Cell-Group-Information-Item'}, 'size': [(1, 'maxnoofCellGroups')]}
    Items []CellGroupInformationItem
}

type CellGroupInformationItem struct { // [{'type': 'Cell-Group-ID', 'name': 'cell-Group-ID'}, {'type': 'UL-Configuration', 'name': 'uL-Configuration', 'optional': True}, {'type': 'DL-TX-Stop', 'name': 'dL-TX-Stop', 'optional': True}, {'type': 'RAT-Type', 'name': 'rAT-Type', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Cell-Group-Information-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    CellGroupID CellGroupID
    ULConfiguration *ULConfiguration
    DLTXStop *DLTXStop
    RATType *RATType
    IEExtensions *CellGroupInformationItemExtIEs
}

func (self * CellGroupInformationItem) Unpack(stream *Stream) {
    uLConfiguration_flag := 0x00000002
    dLTXStop_flag := 0x00000004
    rATType_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.CellGroupID.Unpack(stream)// p8
    if (uLConfiguration_flag & _flags) == uLConfiguration_flag { //cond2
        self.ULConfiguration = &ULConfiguration{}//7{'type': 'UL-Configuration', 'name': 'uL-Configuration', 'optional': True}
        self.ULConfiguration.Unpack(stream)// p8
    }
    if (dLTXStop_flag & _flags) == dLTXStop_flag { //cond2
        self.DLTXStop = &DLTXStop{}//7{'type': 'DL-TX-Stop', 'name': 'dL-TX-Stop', 'optional': True}
        self.DLTXStop.Unpack(stream)// p8
    }
    if (rATType_flag & _flags) == rATType_flag { //cond2
        self.RATType = &RATType{}//7{'type': 'RAT-Type', 'name': 'rAT-Type', 'optional': True}
        self.RATType.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &CellGroupInformationItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Cell-Group-Information-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_CellGroupInformationItemExtIEs, order_CellGroupInformationItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CellGroupInformationItem) Pack(stream *Stream) {
    const uLConfiguration_flag uint = 0x00000002
    const dLTXStop_flag uint = 0x00000004
    const rATType_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CellGroupID.Pack(stream)
    if self.ULConfiguration != nil { 
        _flags |= uLConfiguration_flag
        self.ULConfiguration.Pack(stream)
    }//end of optional
    if self.DLTXStop != nil { 
        _flags |= dLTXStop_flag
        self.DLTXStop.Pack(stream)
    }//end of optional
    if self.RATType != nil { 
        _flags |= rATType_flag
        self.RATType.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_CellGroupInformationItemExtIEs, order_CellGroupInformationItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

type CellGroupID struct {
  Value uint64
}
func (self *CellGroupID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4, 3, 1, 0)
}
func (self * CellGroupID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4, 3, 1, 0)
}
type CipheringAlgorithm struct {
  Value int
}
const (
    CipheringAlgorithmnEA0 = 0
    CipheringAlgorithmc_128_NEA1 = 1
    CipheringAlgorithmc_128_NEA2 = 2
    CipheringAlgorithmc_128_NEA3 = 3

    /* Extensions */
)
func (self *CipheringAlgorithm) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 4, 1)
}
func (self *CipheringAlgorithm) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 4, 1)
}
type CNSupport struct {
  Value int
}
const (
    CNSupportc_epc = 0
    CNSupportc_5gc = 1
    CNSupportboth = 2

    /* Extensions */
)
func (self *CNSupport) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *CNSupport) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type CommonNetworkInstance struct {
  Value HexBytes
}
func (self *CommonNetworkInstance) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *CommonNetworkInstance) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type ConfidentialityProtectionIndication struct {
  Value int
}
const (
    ConfidentialityProtectionIndicationrequired = 0
    ConfidentialityProtectionIndicationpreferred = 1
    ConfidentialityProtectionIndicationnot_needed = 2

    /* Extensions */
)
func (self *ConfidentialityProtectionIndication) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *ConfidentialityProtectionIndication) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type ConfidentialityProtectionResult struct {
  Value int
}
const (
    ConfidentialityProtectionResultperformed = 0
    ConfidentialityProtectionResultnot_performed = 1

    /* Extensions */
)
func (self *ConfidentialityProtectionResult) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *ConfidentialityProtectionResult) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
func (self *CPTNLInformation)Unpack(stream *Stream) {
    //coptions := []string{"endpoint-IP-Address","choice-extension"}
    choice := stream.get_choice(1, 0, 2)
    if choice == 0 { //ch1
        self.EndpointIPAddress = &TransportLayerAddress{}//cho6
        self.EndpointIPAddress.Unpack(stream)
    } else if choice == 1 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &CPTNLInformationExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * CPTNLInformation) Pack(stream *Stream) {
    if self.EndpointIPAddress != nil {
        stream.set_choice(0, 1, 0, 2)
        self.EndpointIPAddress.Pack(stream)//2
    } else if self.Choiceextension != nil {
        stream.set_choice(1, 1, 0, 2)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type CPTNLInformation struct { //[{'type': 'TransportLayerAddress', 'name': 'endpoint-IP-Address'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['CP-TNL-Information-ExtIEs'], 'name': 'choice-extension'}]
    EndpointIPAddress *TransportLayerAddress
    Choiceextension *CPTNLInformationExtIEs
} // CPTNLInformation

type CriticalityDiagnostics struct { // [{'type': 'ProcedureCode', 'name': 'procedureCode', 'optional': True}, {'type': 'TriggeringMessage', 'name': 'triggeringMessage', 'optional': True}, {'type': 'Criticality', 'name': 'procedureCriticality', 'optional': True}, {'type': 'TransactionID', 'name': 'transactionID', 'optional': True}, {'type': 'CriticalityDiagnostics-IE-List', 'name': 'iEsCriticalityDiagnostics', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    ProcedureCode *ProcedureCode
    TriggeringMessage *TriggeringMessage
    ProcedureCriticality *Criticality
    TransactionID *TransactionID
    IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList
    IEExtensions *CriticalityDiagnosticsExtIEs
}

func (self * CriticalityDiagnostics) Unpack(stream *Stream) {
    procedureCode_flag := 0x00000002
    triggeringMessage_flag := 0x00000004
    procedureCriticality_flag := 0x00000008
    transactionID_flag := 0x00000010
    iEsCriticalityDiagnostics_flag := 0x00000020
    iEExtensions_flag := 0x00000040
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(7)
    if (procedureCode_flag & _flags) == procedureCode_flag { //cond2
        self.ProcedureCode = &ProcedureCode{}//7{'type': 'ProcedureCode', 'name': 'procedureCode', 'optional': True}
        self.ProcedureCode.Unpack(stream)// p8
    }
    if (triggeringMessage_flag & _flags) == triggeringMessage_flag { //cond2
        self.TriggeringMessage = &TriggeringMessage{}//7{'type': 'TriggeringMessage', 'name': 'triggeringMessage', 'optional': True}
        self.TriggeringMessage.Unpack(stream)// p8
    }
    if (procedureCriticality_flag & _flags) == procedureCriticality_flag { //cond2
        self.ProcedureCriticality = &Criticality{}//7{'type': 'Criticality', 'name': 'procedureCriticality', 'optional': True}
        self.ProcedureCriticality.Unpack(stream)// p8
    }
    if (transactionID_flag & _flags) == transactionID_flag { //cond2
        self.TransactionID = &TransactionID{}//7{'type': 'TransactionID', 'name': 'transactionID', 'optional': True}
        self.TransactionID.Unpack(stream)// p8
    }
    if (iEsCriticalityDiagnostics_flag & _flags) == iEsCriticalityDiagnostics_flag { //cond2
        self.IEsCriticalityDiagnostics = &CriticalityDiagnosticsIEList{}//7{'type': 'CriticalityDiagnostics-IE-List', 'name': 'iEsCriticalityDiagnostics', 'optional': True}
        self.IEsCriticalityDiagnostics.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &CriticalityDiagnosticsExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_CriticalityDiagnosticsExtIEs, order_CriticalityDiagnosticsExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CriticalityDiagnostics) Pack(stream *Stream) {
    const procedureCode_flag uint = 0x00000002
    const triggeringMessage_flag uint = 0x00000004
    const procedureCriticality_flag uint = 0x00000008
    const transactionID_flag uint = 0x00000010
    const iEsCriticalityDiagnostics_flag uint = 0x00000020
    const iEExtensions_flag uint = 0x00000040
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(7)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.ProcedureCode != nil { 
        _flags |= procedureCode_flag
        self.ProcedureCode.Pack(stream)
    }//end of optional
    if self.TriggeringMessage != nil { 
        _flags |= triggeringMessage_flag
        self.TriggeringMessage.Pack(stream)
    }//end of optional
    if self.ProcedureCriticality != nil { 
        _flags |= procedureCriticality_flag
        self.ProcedureCriticality.Pack(stream)
    }//end of optional
    if self.TransactionID != nil { 
        _flags |= transactionID_flag
        self.TransactionID.Pack(stream)
    }//end of optional
    if self.IEsCriticalityDiagnostics != nil { 
        _flags |= iEsCriticalityDiagnostics_flag
        self.IEsCriticalityDiagnostics.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_CriticalityDiagnosticsExtIEs, order_CriticalityDiagnosticsExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 7)
}//end

func (self *CriticalityDiagnosticsIEList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]CriticalityDiagnosticsIEList_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *CriticalityDiagnosticsIEList_Item) { //[{'type': 'Criticality', 'name': 'iECriticality'}, {'type': 'ProtocolIE-ID', 'name': 'iE-ID'}, {'type': 'TypeOfError', 'name': 'typeOfError'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-IE-List-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        iEExtensions_flag := 0x00000002
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(2)
        self.IECriticality.Unpack(stream)// p8
        self.IEID.Unpack(stream)// p8
        self.TypeOfError.Unpack(stream)// p8
        if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
            self.IEExtensions = &CriticalityDiagnosticsIEListExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-IE-List-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
            IEExtensions := ProtocolExtensionContainer {table_CriticalityDiagnosticsIEListExtIEs, order_CriticalityDiagnosticsIEListExtIEs} // p3
            IEExtensions.Unpack(stream, &self.IEExtensions) // p3
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    for item := 0; item <_size; item +=1 {
        Unpack_Item(stream, &self.Items[item])
    }
}


func (self *CriticalityDiagnosticsIEList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    var Pack_Item = func(stream *Stream, self CriticalityDiagnosticsIEList_Item) {//seq
        const iEExtensions_flag uint = 0x00000002
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(2)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        self.IECriticality.Pack(stream)
        self.IEID.Pack(stream)
        self.TypeOfError.Pack(stream)
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_CriticalityDiagnosticsIEListExtIEs, order_CriticalityDiagnosticsIEListExtIEs} // p3
            IEExtensions.Pack(stream, &self.IEExtensions)
        }//end of optional
        if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
        stream.set_flags(_flags, _flagReserve, 2)
    }//end
    for _, item := range self.Items { // seqof base type
        Pack_Item(stream, item)
    }
    return

}


type CriticalityDiagnosticsIEList_Item struct { // [{'type': 'Criticality', 'name': 'iECriticality'}, {'type': 'ProtocolIE-ID', 'name': 'iE-ID'}, {'type': 'TypeOfError', 'name': 'typeOfError'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-IE-List-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    IECriticality Criticality
    IEID ProtocolIEID
    TypeOfError TypeOfError
    IEExtensions *CriticalityDiagnosticsIEListExtIEs
}
type CriticalityDiagnosticsIEList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'Criticality', 'name': 'iECriticality'}, {'type': 'ProtocolIE-ID', 'name': 'iE-ID'}, {'type': 'TypeOfError', 'name': 'typeOfError'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-IE-List-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxnoofErrors')]}
    Items []CriticalityDiagnosticsIEList_Item
}

type DataForwardingInformationRequest struct { // [{'type': 'Data-Forwarding-Request', 'name': 'data-Forwarding-Request'}, {'type': 'QoS-Flow-Mapping-List', 'name': 'qoS-Flows-Forwarded-On-Fwd-Tunnels', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Data-Forwarding-Information-Request-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DataForwardingRequest DataForwardingRequest
    QoSFlowsForwardedOnFwdTunnels *QoSFlowMappingList
    IEExtensions *DataForwardingInformationRequestExtIEs
}

func (self * DataForwardingInformationRequest) Unpack(stream *Stream) {
    qoSFlowsForwardedOnFwdTunnels_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.DataForwardingRequest.Unpack(stream)// p8
    if (qoSFlowsForwardedOnFwdTunnels_flag & _flags) == qoSFlowsForwardedOnFwdTunnels_flag { //cond2
        self.QoSFlowsForwardedOnFwdTunnels = &QoSFlowMappingList{}//7{'type': 'QoS-Flow-Mapping-List', 'name': 'qoS-Flows-Forwarded-On-Fwd-Tunnels', 'optional': True}
        self.QoSFlowsForwardedOnFwdTunnels.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DataForwardingInformationRequestExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Data-Forwarding-Information-Request-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DataForwardingInformationRequestExtIEs, order_DataForwardingInformationRequestExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DataForwardingInformationRequest) Pack(stream *Stream) {
    const qoSFlowsForwardedOnFwdTunnels_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DataForwardingRequest.Pack(stream)
    if self.QoSFlowsForwardedOnFwdTunnels != nil { 
        _flags |= qoSFlowsForwardedOnFwdTunnels_flag
        self.QoSFlowsForwardedOnFwdTunnels.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DataForwardingInformationRequestExtIEs, order_DataForwardingInformationRequestExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type DataForwardingInformation struct { // [{'type': 'UP-TNL-Information', 'name': 'uL-Data-Forwarding', 'optional': True}, {'type': 'UP-TNL-Information', 'name': 'dL-Data-Forwarding', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Data-Forwarding-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    ULDataForwarding *UPTNLInformation
    DLDataForwarding *UPTNLInformation
    IEExtensions *DataForwardingInformationExtIEs
}

func (self * DataForwardingInformation) Unpack(stream *Stream) {
    uLDataForwarding_flag := 0x00000002
    dLDataForwarding_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (uLDataForwarding_flag & _flags) == uLDataForwarding_flag { //cond2
        self.ULDataForwarding = &UPTNLInformation{}//7{'type': 'UP-TNL-Information', 'name': 'uL-Data-Forwarding', 'optional': True}
        self.ULDataForwarding.Unpack(stream)// p8
    }
    if (dLDataForwarding_flag & _flags) == dLDataForwarding_flag { //cond2
        self.DLDataForwarding = &UPTNLInformation{}//7{'type': 'UP-TNL-Information', 'name': 'dL-Data-Forwarding', 'optional': True}
        self.DLDataForwarding.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DataForwardingInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Data-Forwarding-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DataForwardingInformationExtIEs, order_DataForwardingInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DataForwardingInformation) Pack(stream *Stream) {
    const uLDataForwarding_flag uint = 0x00000002
    const dLDataForwarding_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.ULDataForwarding != nil { 
        _flags |= uLDataForwarding_flag
        self.ULDataForwarding.Pack(stream)
    }//end of optional
    if self.DLDataForwarding != nil { 
        _flags |= dLDataForwarding_flag
        self.DLDataForwarding.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DataForwardingInformationExtIEs, order_DataForwardingInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type DataForwardingRequest struct {
  Value int
}
const (
    DataForwardingRequestuL = 0
    DataForwardingRequestdL = 1
    DataForwardingRequestboth = 2

    /* Extensions */
)
func (self *DataForwardingRequest) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *DataForwardingRequest) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type DataUsageperPDUSessionReport_PDUsessionTimedReportList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MRDC-Data-Usage-Report-Item'}, 'size': [(1, 'maxnooftimeperiods')], 'name': 'pDU-session-Timed-Report-List'}
    Items []MRDCDataUsageReportItem
}
type DataUsageperPDUSessionReport struct { // [{'type': 'ENUMERATED', 'values': [('nR', 0), ('e-UTRA', 1), None], 'name': 'secondaryRATType'}, {'type': 'SEQUENCE OF', 'element': {'type': 'MRDC-Data-Usage-Report-Item'}, 'size': [(1, 'maxnooftimeperiods')], 'name': 'pDU-session-Timed-Report-List'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Data-Usage-per-PDU-Session-Report-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    SecondaryRATType ENUMERATED
    PDUsessionTimedReportList DataUsageperPDUSessionReport_PDUsessionTimedReportList
    IEExtensions *DataUsageperPDUSessionReportExtIEs
}

func (self * DataUsageperPDUSessionReport) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_secondaryRATType = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_secondaryRATType(stream, &self.SecondaryRATType)// p2
    var Unpack_pDUsessionTimedReportList = func(stream *Stream, self *DataUsageperPDUSessionReport_PDUsessionTimedReportList){// Seq6 DataUsageperPDUSessionReport {'type': 'SEQUENCE OF', 'element': {'type': 'MRDC-Data-Usage-Report-Item'}, 'size': [(1, 'maxnooftimeperiods')], 'name': 'pDU-session-Timed-Report-List'}
        _size := stream.get_listsize(2)
        _size += 1
        self.Items = make([]MRDCDataUsageReportItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_pDUsessionTimedReportList(stream, &self.PDUsessionTimedReportList)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DataUsageperPDUSessionReportExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Data-Usage-per-PDU-Session-Report-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DataUsageperPDUSessionReportExtIEs, order_DataUsageperPDUSessionReportExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DataUsageperPDUSessionReport) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_secondaryRATType = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_secondaryRATType(stream, self.SecondaryRATType) //f2
    var Pack_pDUsessionTimedReportList = func(stream *Stream, self DataUsageperPDUSessionReport_PDUsessionTimedReportList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 2)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_pDUsessionTimedReportList(stream, self.PDUsessionTimedReportList) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DataUsageperPDUSessionReportExtIEs, order_DataUsageperPDUSessionReportExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DataUsageperQoSFlowList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(64)
    _size += 1
    self.Items = make([]DataUsageperQoSFlowItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DataUsageperQoSFlowList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 64)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DataUsageperQoSFlowList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'Data-Usage-per-QoS-Flow-Item'}, 'size': [(1, 'maxnoofQoSFlows')]}
    Items []DataUsageperQoSFlowItem
}

type DataUsageperQoSFlowItem_QoSFlowTimedReportList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MRDC-Data-Usage-Report-Item'}, 'size': [(1, 'maxnooftimeperiods')], 'name': 'qoS-Flow-Timed-Report-List'}
    Items []MRDCDataUsageReportItem
}
type DataUsageperQoSFlowItem struct { // [{'type': 'QoS-Flow-Identifier', 'name': 'qoS-Flow-Identifier'}, {'type': 'ENUMERATED', 'values': [('nR', 0), ('e-UTRA', 1), None], 'name': 'secondaryRATType'}, {'type': 'SEQUENCE OF', 'element': {'type': 'MRDC-Data-Usage-Report-Item'}, 'size': [(1, 'maxnooftimeperiods')], 'name': 'qoS-Flow-Timed-Report-List'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Data-Usage-per-QoS-Flow-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    QoSFlowIdentifier QoSFlowIdentifier
    SecondaryRATType ENUMERATED
    QoSFlowTimedReportList DataUsageperQoSFlowItem_QoSFlowTimedReportList
    IEExtensions *DataUsageperQoSFlowItemExtIEs
}

func (self * DataUsageperQoSFlowItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.QoSFlowIdentifier.Unpack(stream)// p8
    var Unpack_secondaryRATType = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_secondaryRATType(stream, &self.SecondaryRATType)// p2
    var Unpack_qoSFlowTimedReportList = func(stream *Stream, self *DataUsageperQoSFlowItem_QoSFlowTimedReportList){// Seq6 DataUsageperQoSFlowItem {'type': 'SEQUENCE OF', 'element': {'type': 'MRDC-Data-Usage-Report-Item'}, 'size': [(1, 'maxnooftimeperiods')], 'name': 'qoS-Flow-Timed-Report-List'}
        _size := stream.get_listsize(2)
        _size += 1
        self.Items = make([]MRDCDataUsageReportItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_qoSFlowTimedReportList(stream, &self.QoSFlowTimedReportList)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DataUsageperQoSFlowItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Data-Usage-per-QoS-Flow-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DataUsageperQoSFlowItemExtIEs, order_DataUsageperQoSFlowItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DataUsageperQoSFlowItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.QoSFlowIdentifier.Pack(stream)
    var Pack_secondaryRATType = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_secondaryRATType(stream, self.SecondaryRATType) //f2
    var Pack_qoSFlowTimedReportList = func(stream *Stream, self DataUsageperQoSFlowItem_QoSFlowTimedReportList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 2)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_qoSFlowTimedReportList(stream, self.QoSFlowTimedReportList) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DataUsageperQoSFlowItemExtIEs, order_DataUsageperQoSFlowItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DataUsageReportList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DataUsageReportItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DataUsageReportList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DataUsageReportList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'Data-Usage-Report-Item'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DataUsageReportItem
}

type DataUsageReportItem struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'RAT-Type', 'name': 'rAT-Type'}, {'type': 'DRB-Usage-Report-List', 'name': 'dRB-Usage-Report-List'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Data-Usage-Report-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    RATType RATType
    DRBUsageReportList DRBUsageReportList
    IEExtensions *DataUsageReportItemExtIEs
}

func (self * DataUsageReportItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.RATType.Unpack(stream)// p8
    self.DRBUsageReportList.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DataUsageReportItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Data-Usage-Report-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DataUsageReportItemExtIEs, order_DataUsageReportItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DataUsageReportItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.RATType.Pack(stream)
    self.DRBUsageReportList.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DataUsageReportItemExtIEs, order_DataUsageReportItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type DefaultDRB struct {
  Value int
}
const (
    DefaultDRBtRue = 0
    DefaultDRBfAlse = 1

    /* Extensions */
)
func (self *DefaultDRB) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *DefaultDRB) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type DiscardTimer struct {
  Value int
}
const (
    DiscardTimerms10 = 0
    DiscardTimerms20 = 1
    DiscardTimerms30 = 2
    DiscardTimerms40 = 3
    DiscardTimerms50 = 4
    DiscardTimerms60 = 5
    DiscardTimerms75 = 6
    DiscardTimerms100 = 7
    DiscardTimerms150 = 8
    DiscardTimerms200 = 9
    DiscardTimerms250 = 10
    DiscardTimerms300 = 11
    DiscardTimerms500 = 12
    DiscardTimerms750 = 13
    DiscardTimerms1500 = 14
    DiscardTimerinfinity = 15
)
func (self *DiscardTimer) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 16, 0)
}
func (self *DiscardTimer) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 16, 0)
}
type DLTXStop struct {
  Value int
}
const (
    DLTXStopstop = 0
    DLTXStopresume = 1

    /* Extensions */
)
func (self *DLTXStop) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *DLTXStop) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type DRBActivity struct {
  Value int
}
const (
    DRBActivityactive = 0
    DRBActivitynot_active = 1

    /* Extensions */
)
func (self *DRBActivity) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *DRBActivity) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
func (self *DRBActivityList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBActivityItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBActivityList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBActivityList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Activity-Item'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBActivityItem
}

type DRBActivityItem struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'DRB-Activity', 'name': 'dRB-Activity'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Activity-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    DRBActivity DRBActivity
    IEExtensions *DRBActivityItemExtIEs
}

func (self * DRBActivityItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.DRBActivity.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBActivityItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Activity-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBActivityItemExtIEs, order_DRBActivityItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBActivityItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.DRBActivity.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBActivityItemExtIEs, order_DRBActivityItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBConfirmModifiedListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBConfirmModifiedItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBConfirmModifiedListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBConfirmModifiedListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Confirm-Modified-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBConfirmModifiedItemEUTRAN
}

type DRBConfirmModifiedItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-Information', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Confirm-Modified-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    CellGroupInformation *CellGroupInformation
    IEExtensions *DRBConfirmModifiedItemEUTRANExtIEs
}

func (self * DRBConfirmModifiedItemEUTRAN) Unpack(stream *Stream) {
    cellGroupInformation_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.DRBID.Unpack(stream)// p8
    if (cellGroupInformation_flag & _flags) == cellGroupInformation_flag { //cond2
        self.CellGroupInformation = &CellGroupInformation{}//7{'type': 'Cell-Group-Information', 'name': 'cell-Group-Information', 'optional': True}
        self.CellGroupInformation.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBConfirmModifiedItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Confirm-Modified-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBConfirmModifiedItemEUTRANExtIEs, order_DRBConfirmModifiedItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBConfirmModifiedItemEUTRAN) Pack(stream *Stream) {
    const cellGroupInformation_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.CellGroupInformation != nil { 
        _flags |= cellGroupInformation_flag
        self.CellGroupInformation.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBConfirmModifiedItemEUTRANExtIEs, order_DRBConfirmModifiedItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

func (self *DRBConfirmModifiedListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBConfirmModifiedItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBConfirmModifiedListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBConfirmModifiedListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Confirm-Modified-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBConfirmModifiedItemNGRAN
}

type DRBConfirmModifiedItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-Information', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Confirm-Modified-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    CellGroupInformation *CellGroupInformation
    IEExtensions *DRBConfirmModifiedItemNGRANExtIEs
}

func (self * DRBConfirmModifiedItemNGRAN) Unpack(stream *Stream) {
    cellGroupInformation_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.DRBID.Unpack(stream)// p8
    if (cellGroupInformation_flag & _flags) == cellGroupInformation_flag { //cond2
        self.CellGroupInformation = &CellGroupInformation{}//7{'type': 'Cell-Group-Information', 'name': 'cell-Group-Information', 'optional': True}
        self.CellGroupInformation.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBConfirmModifiedItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Confirm-Modified-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBConfirmModifiedItemNGRANExtIEs, order_DRBConfirmModifiedItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBConfirmModifiedItemNGRAN) Pack(stream *Stream) {
    const cellGroupInformation_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.CellGroupInformation != nil { 
        _flags |= cellGroupInformation_flag
        self.CellGroupInformation.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBConfirmModifiedItemNGRANExtIEs, order_DRBConfirmModifiedItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

func (self *DRBFailedListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBFailedItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBFailedListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBFailedListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Failed-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBFailedItemEUTRAN
}

type DRBFailedItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    Cause Cause
    IEExtensions *DRBFailedItemEUTRANExtIEs
}

func (self * DRBFailedItemEUTRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBFailedItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBFailedItemEUTRANExtIEs, order_DRBFailedItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBFailedItemEUTRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBFailedItemEUTRANExtIEs, order_DRBFailedItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBFailedModListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBFailedModItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBFailedModListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBFailedModListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Failed-Mod-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBFailedModItemEUTRAN
}

type DRBFailedModItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-Mod-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    Cause Cause
    IEExtensions *DRBFailedModItemEUTRANExtIEs
}

func (self * DRBFailedModItemEUTRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBFailedModItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-Mod-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBFailedModItemEUTRANExtIEs, order_DRBFailedModItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBFailedModItemEUTRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBFailedModItemEUTRANExtIEs, order_DRBFailedModItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBFailedListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBFailedItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBFailedListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBFailedListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Failed-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBFailedItemNGRAN
}

type DRBFailedItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    Cause Cause
    IEExtensions *DRBFailedItemNGRANExtIEs
}

func (self * DRBFailedItemNGRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBFailedItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBFailedItemNGRANExtIEs, order_DRBFailedItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBFailedItemNGRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBFailedItemNGRANExtIEs, order_DRBFailedItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBFailedModListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBFailedModItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBFailedModListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBFailedModListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Failed-Mod-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBFailedModItemNGRAN
}

type DRBFailedModItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-Mod-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    Cause Cause
    IEExtensions *DRBFailedModItemNGRANExtIEs
}

func (self * DRBFailedModItemNGRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBFailedModItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-Mod-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBFailedModItemNGRANExtIEs, order_DRBFailedModItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBFailedModItemNGRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBFailedModItemNGRANExtIEs, order_DRBFailedModItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBFailedToModifyListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBFailedToModifyItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBFailedToModifyListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBFailedToModifyListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Failed-To-Modify-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBFailedToModifyItemEUTRAN
}

type DRBFailedToModifyItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-To-Modify-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    Cause Cause
    IEExtensions *DRBFailedToModifyItemEUTRANExtIEs
}

func (self * DRBFailedToModifyItemEUTRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBFailedToModifyItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-To-Modify-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBFailedToModifyItemEUTRANExtIEs, order_DRBFailedToModifyItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBFailedToModifyItemEUTRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBFailedToModifyItemEUTRANExtIEs, order_DRBFailedToModifyItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBFailedToModifyListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBFailedToModifyItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBFailedToModifyListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBFailedToModifyListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Failed-To-Modify-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBFailedToModifyItemNGRAN
}

type DRBFailedToModifyItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-To-Modify-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    Cause Cause
    IEExtensions *DRBFailedToModifyItemNGRANExtIEs
}

func (self * DRBFailedToModifyItemNGRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBFailedToModifyItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Failed-To-Modify-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBFailedToModifyItemNGRANExtIEs, order_DRBFailedToModifyItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBFailedToModifyItemNGRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBFailedToModifyItemNGRANExtIEs, order_DRBFailedToModifyItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type DRBID struct {
  Value uint64
}
func (self *DRBID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(32, 6, 1, 1)
}
func (self * DRBID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 32, 6, 1, 1)
}
func (self *DRBModifiedListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBModifiedItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBModifiedListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBModifiedListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Modified-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBModifiedItemEUTRAN
}

type DRBModifiedItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'UP-TNL-Information', 'name': 's1-DL-UP-TNL-Information', 'optional': True}, {'type': 'PDCP-SN-Status-Information', 'name': 'pDCP-SN-Status-Information', 'optional': True}, {'type': 'UP-Parameters', 'name': 'uL-UP-Transport-Parameters', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Modified-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    S1DLUPTNLInformation *UPTNLInformation
    PDCPSNStatusInformation *PDCPSNStatusInformation
    ULUPTransportParameters *UPParameters
    IEExtensions *DRBModifiedItemEUTRANExtIEs
}

func (self * DRBModifiedItemEUTRAN) Unpack(stream *Stream) {
    s1DLUPTNLInformation_flag := 0x00000002
    pDCPSNStatusInformation_flag := 0x00000004
    uLUPTransportParameters_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.DRBID.Unpack(stream)// p8
    if (s1DLUPTNLInformation_flag & _flags) == s1DLUPTNLInformation_flag { //cond2
        self.S1DLUPTNLInformation = &UPTNLInformation{}//7{'type': 'UP-TNL-Information', 'name': 's1-DL-UP-TNL-Information', 'optional': True}
        self.S1DLUPTNLInformation.Unpack(stream)// p8
    }
    if (pDCPSNStatusInformation_flag & _flags) == pDCPSNStatusInformation_flag { //cond2
        self.PDCPSNStatusInformation = &PDCPSNStatusInformation{}//7{'type': 'PDCP-SN-Status-Information', 'name': 'pDCP-SN-Status-Information', 'optional': True}
        self.PDCPSNStatusInformation.Unpack(stream)// p8
    }
    if (uLUPTransportParameters_flag & _flags) == uLUPTransportParameters_flag { //cond2
        self.ULUPTransportParameters = &UPParameters{}//7{'type': 'UP-Parameters', 'name': 'uL-UP-Transport-Parameters', 'optional': True}
        self.ULUPTransportParameters.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBModifiedItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Modified-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBModifiedItemEUTRANExtIEs, order_DRBModifiedItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBModifiedItemEUTRAN) Pack(stream *Stream) {
    const s1DLUPTNLInformation_flag uint = 0x00000002
    const pDCPSNStatusInformation_flag uint = 0x00000004
    const uLUPTransportParameters_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.S1DLUPTNLInformation != nil { 
        _flags |= s1DLUPTNLInformation_flag
        self.S1DLUPTNLInformation.Pack(stream)
    }//end of optional
    if self.PDCPSNStatusInformation != nil { 
        _flags |= pDCPSNStatusInformation_flag
        self.PDCPSNStatusInformation.Pack(stream)
    }//end of optional
    if self.ULUPTransportParameters != nil { 
        _flags |= uLUPTransportParameters_flag
        self.ULUPTransportParameters.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBModifiedItemEUTRANExtIEs, order_DRBModifiedItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *DRBModifiedListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBModifiedItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBModifiedListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBModifiedListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Modified-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBModifiedItemNGRAN
}

type DRBModifiedItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'UP-Parameters', 'name': 'uL-UP-Transport-Parameters', 'optional': True}, {'type': 'PDCP-SN-Status-Information', 'name': 'pDCP-SN-Status-Information', 'optional': True}, {'type': 'QoS-Flow-List', 'name': 'flow-Setup-List', 'optional': True}, {'type': 'QoS-Flow-Failed-List', 'name': 'flow-Failed-List', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Modified-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    ULUPTransportParameters *UPParameters
    PDCPSNStatusInformation *PDCPSNStatusInformation
    FlowSetupList *QoSFlowList
    FlowFailedList *QoSFlowFailedList
    IEExtensions *DRBModifiedItemNGRANExtIEs
}

func (self * DRBModifiedItemNGRAN) Unpack(stream *Stream) {
    uLUPTransportParameters_flag := 0x00000002
    pDCPSNStatusInformation_flag := 0x00000004
    flowSetupList_flag := 0x00000008
    flowFailedList_flag := 0x00000010
    iEExtensions_flag := 0x00000020
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(6)
    self.DRBID.Unpack(stream)// p8
    if (uLUPTransportParameters_flag & _flags) == uLUPTransportParameters_flag { //cond2
        self.ULUPTransportParameters = &UPParameters{}//7{'type': 'UP-Parameters', 'name': 'uL-UP-Transport-Parameters', 'optional': True}
        self.ULUPTransportParameters.Unpack(stream)// p8
    }
    if (pDCPSNStatusInformation_flag & _flags) == pDCPSNStatusInformation_flag { //cond2
        self.PDCPSNStatusInformation = &PDCPSNStatusInformation{}//7{'type': 'PDCP-SN-Status-Information', 'name': 'pDCP-SN-Status-Information', 'optional': True}
        self.PDCPSNStatusInformation.Unpack(stream)// p8
    }
    if (flowSetupList_flag & _flags) == flowSetupList_flag { //cond2
        self.FlowSetupList = &QoSFlowList{}//7{'type': 'QoS-Flow-List', 'name': 'flow-Setup-List', 'optional': True}
        self.FlowSetupList.Unpack(stream)// p8
    }
    if (flowFailedList_flag & _flags) == flowFailedList_flag { //cond2
        self.FlowFailedList = &QoSFlowFailedList{}//7{'type': 'QoS-Flow-Failed-List', 'name': 'flow-Failed-List', 'optional': True}
        self.FlowFailedList.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBModifiedItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Modified-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBModifiedItemNGRANExtIEs, order_DRBModifiedItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBModifiedItemNGRAN) Pack(stream *Stream) {
    const uLUPTransportParameters_flag uint = 0x00000002
    const pDCPSNStatusInformation_flag uint = 0x00000004
    const flowSetupList_flag uint = 0x00000008
    const flowFailedList_flag uint = 0x00000010
    const iEExtensions_flag uint = 0x00000020
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(6)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.ULUPTransportParameters != nil { 
        _flags |= uLUPTransportParameters_flag
        self.ULUPTransportParameters.Pack(stream)
    }//end of optional
    if self.PDCPSNStatusInformation != nil { 
        _flags |= pDCPSNStatusInformation_flag
        self.PDCPSNStatusInformation.Pack(stream)
    }//end of optional
    if self.FlowSetupList != nil { 
        _flags |= flowSetupList_flag
        self.FlowSetupList.Pack(stream)
    }//end of optional
    if self.FlowFailedList != nil { 
        _flags |= flowFailedList_flag
        self.FlowFailedList.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBModifiedItemNGRANExtIEs, order_DRBModifiedItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

func (self *DRBRequiredToModifyListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBRequiredToModifyItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBRequiredToModifyListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBRequiredToModifyListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Required-To-Modify-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBRequiredToModifyItemEUTRAN
}

type DRBRequiredToModifyItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'UP-TNL-Information', 'name': 's1-DL-UP-TNL-Information', 'optional': True}, {'type': 'GNB-CU-UP-CellGroupRelatedConfiguration', 'name': 'gNB-CU-UP-CellGroupRelatedConfiguration', 'optional': True}, {'type': 'Cause', 'name': 'cause', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Required-To-Modify-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    S1DLUPTNLInformation *UPTNLInformation
    GNBCUUPCellGroupRelatedConfiguration *GNBCUUPCellGroupRelatedConfiguration
    Cause *Cause
    IEExtensions *DRBRequiredToModifyItemEUTRANExtIEs
}

func (self * DRBRequiredToModifyItemEUTRAN) Unpack(stream *Stream) {
    s1DLUPTNLInformation_flag := 0x00000002
    gNBCUUPCellGroupRelatedConfiguration_flag := 0x00000004
    cause_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.DRBID.Unpack(stream)// p8
    if (s1DLUPTNLInformation_flag & _flags) == s1DLUPTNLInformation_flag { //cond2
        self.S1DLUPTNLInformation = &UPTNLInformation{}//7{'type': 'UP-TNL-Information', 'name': 's1-DL-UP-TNL-Information', 'optional': True}
        self.S1DLUPTNLInformation.Unpack(stream)// p8
    }
    if (gNBCUUPCellGroupRelatedConfiguration_flag & _flags) == gNBCUUPCellGroupRelatedConfiguration_flag { //cond2
        self.GNBCUUPCellGroupRelatedConfiguration = &GNBCUUPCellGroupRelatedConfiguration{}//7{'type': 'GNB-CU-UP-CellGroupRelatedConfiguration', 'name': 'gNB-CU-UP-CellGroupRelatedConfiguration', 'optional': True}
        self.GNBCUUPCellGroupRelatedConfiguration.Unpack(stream)// p8
    }
    if (cause_flag & _flags) == cause_flag { //cond2
        self.Cause = &Cause{}//7{'type': 'Cause', 'name': 'cause', 'optional': True}
        self.Cause.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBRequiredToModifyItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Required-To-Modify-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBRequiredToModifyItemEUTRANExtIEs, order_DRBRequiredToModifyItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBRequiredToModifyItemEUTRAN) Pack(stream *Stream) {
    const s1DLUPTNLInformation_flag uint = 0x00000002
    const gNBCUUPCellGroupRelatedConfiguration_flag uint = 0x00000004
    const cause_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.S1DLUPTNLInformation != nil { 
        _flags |= s1DLUPTNLInformation_flag
        self.S1DLUPTNLInformation.Pack(stream)
    }//end of optional
    if self.GNBCUUPCellGroupRelatedConfiguration != nil { 
        _flags |= gNBCUUPCellGroupRelatedConfiguration_flag
        self.GNBCUUPCellGroupRelatedConfiguration.Pack(stream)
    }//end of optional
    if self.Cause != nil { 
        _flags |= cause_flag
        self.Cause.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBRequiredToModifyItemEUTRANExtIEs, order_DRBRequiredToModifyItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *DRBRequiredToModifyListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBRequiredToModifyItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBRequiredToModifyListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBRequiredToModifyListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Required-To-Modify-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBRequiredToModifyItemNGRAN
}

type DRBRequiredToModifyItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'GNB-CU-UP-CellGroupRelatedConfiguration', 'name': 'gNB-CU-UP-CellGroupRelatedConfiguration', 'optional': True}, {'type': 'QoS-Flow-List', 'name': 'flow-To-Remove', 'optional': True}, {'type': 'Cause', 'name': 'cause', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Required-To-Modify-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    GNBCUUPCellGroupRelatedConfiguration *GNBCUUPCellGroupRelatedConfiguration
    FlowToRemove *QoSFlowList
    Cause *Cause
    IEExtensions *DRBRequiredToModifyItemNGRANExtIEs
}

func (self * DRBRequiredToModifyItemNGRAN) Unpack(stream *Stream) {
    gNBCUUPCellGroupRelatedConfiguration_flag := 0x00000002
    flowToRemove_flag := 0x00000004
    cause_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.DRBID.Unpack(stream)// p8
    if (gNBCUUPCellGroupRelatedConfiguration_flag & _flags) == gNBCUUPCellGroupRelatedConfiguration_flag { //cond2
        self.GNBCUUPCellGroupRelatedConfiguration = &GNBCUUPCellGroupRelatedConfiguration{}//7{'type': 'GNB-CU-UP-CellGroupRelatedConfiguration', 'name': 'gNB-CU-UP-CellGroupRelatedConfiguration', 'optional': True}
        self.GNBCUUPCellGroupRelatedConfiguration.Unpack(stream)// p8
    }
    if (flowToRemove_flag & _flags) == flowToRemove_flag { //cond2
        self.FlowToRemove = &QoSFlowList{}//7{'type': 'QoS-Flow-List', 'name': 'flow-To-Remove', 'optional': True}
        self.FlowToRemove.Unpack(stream)// p8
    }
    if (cause_flag & _flags) == cause_flag { //cond2
        self.Cause = &Cause{}//7{'type': 'Cause', 'name': 'cause', 'optional': True}
        self.Cause.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBRequiredToModifyItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Required-To-Modify-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBRequiredToModifyItemNGRANExtIEs, order_DRBRequiredToModifyItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBRequiredToModifyItemNGRAN) Pack(stream *Stream) {
    const gNBCUUPCellGroupRelatedConfiguration_flag uint = 0x00000002
    const flowToRemove_flag uint = 0x00000004
    const cause_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.GNBCUUPCellGroupRelatedConfiguration != nil { 
        _flags |= gNBCUUPCellGroupRelatedConfiguration_flag
        self.GNBCUUPCellGroupRelatedConfiguration.Pack(stream)
    }//end of optional
    if self.FlowToRemove != nil { 
        _flags |= flowToRemove_flag
        self.FlowToRemove.Pack(stream)
    }//end of optional
    if self.Cause != nil { 
        _flags |= cause_flag
        self.Cause.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBRequiredToModifyItemNGRANExtIEs, order_DRBRequiredToModifyItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *DRBSetupListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBSetupItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBSetupListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBSetupListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Setup-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBSetupItemEUTRAN
}

type DRBSetupItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'UP-TNL-Information', 'name': 's1-DL-UP-TNL-Information'}, {'type': 'Data-Forwarding-Information', 'name': 'data-Forwarding-Information-Response', 'optional': True}, {'type': 'UP-Parameters', 'name': 'uL-UP-Transport-Parameters'}, {'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 's1-DL-UP-Unchanged', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Setup-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    S1DLUPTNLInformation UPTNLInformation
    DataForwardingInformationResponse *DataForwardingInformation
    ULUPTransportParameters UPParameters
    S1DLUPUnchanged *ENUMERATED
    IEExtensions *DRBSetupItemEUTRANExtIEs
}

func (self * DRBSetupItemEUTRAN) Unpack(stream *Stream) {
    dataForwardingInformationResponse_flag := 0x00000002
    s1DLUPUnchanged_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.DRBID.Unpack(stream)// p8
    self.S1DLUPTNLInformation.Unpack(stream)// p8
    if (dataForwardingInformationResponse_flag & _flags) == dataForwardingInformationResponse_flag { //cond2
        self.DataForwardingInformationResponse = &DataForwardingInformation{}//7{'type': 'Data-Forwarding-Information', 'name': 'data-Forwarding-Information-Response', 'optional': True}
        self.DataForwardingInformationResponse.Unpack(stream)// p8
    }
    self.ULUPTransportParameters.Unpack(stream)// p8
    if (s1DLUPUnchanged_flag & _flags) == s1DLUPUnchanged_flag { //cond1
        var Unpack_s1DLUPUnchanged = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.S1DLUPUnchanged = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 's1-DL-UP-Unchanged', 'optional': True}
        Unpack_s1DLUPUnchanged(stream, self.S1DLUPUnchanged)// p1 {'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 's1-DL-UP-Unchanged', 'optional': True}
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBSetupItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Setup-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBSetupItemEUTRANExtIEs, order_DRBSetupItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBSetupItemEUTRAN) Pack(stream *Stream) {
    const dataForwardingInformationResponse_flag uint = 0x00000002
    const s1DLUPUnchanged_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.S1DLUPTNLInformation.Pack(stream)
    if self.DataForwardingInformationResponse != nil { 
        _flags |= dataForwardingInformationResponse_flag
        self.DataForwardingInformationResponse.Pack(stream)
    }//end of optional
    self.ULUPTransportParameters.Pack(stream)
    if self.S1DLUPUnchanged != nil { //YY
        _flags |= s1DLUPUnchanged_flag
        var Pack_s1DLUPUnchanged = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_s1DLUPUnchanged(stream, *self.S1DLUPUnchanged) //f1
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBSetupItemEUTRANExtIEs, order_DRBSetupItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

func (self *DRBSetupModListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBSetupModItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBSetupModListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBSetupModListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Setup-Mod-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBSetupModItemEUTRAN
}

type DRBSetupModItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'UP-TNL-Information', 'name': 's1-DL-UP-TNL-Information'}, {'type': 'Data-Forwarding-Information', 'name': 'data-Forwarding-Information-Response', 'optional': True}, {'type': 'UP-Parameters', 'name': 'uL-UP-Transport-Parameters'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Setup-Mod-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    S1DLUPTNLInformation UPTNLInformation
    DataForwardingInformationResponse *DataForwardingInformation
    ULUPTransportParameters UPParameters
    IEExtensions *DRBSetupModItemEUTRANExtIEs
}

func (self * DRBSetupModItemEUTRAN) Unpack(stream *Stream) {
    dataForwardingInformationResponse_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.DRBID.Unpack(stream)// p8
    self.S1DLUPTNLInformation.Unpack(stream)// p8
    if (dataForwardingInformationResponse_flag & _flags) == dataForwardingInformationResponse_flag { //cond2
        self.DataForwardingInformationResponse = &DataForwardingInformation{}//7{'type': 'Data-Forwarding-Information', 'name': 'data-Forwarding-Information-Response', 'optional': True}
        self.DataForwardingInformationResponse.Unpack(stream)// p8
    }
    self.ULUPTransportParameters.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBSetupModItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Setup-Mod-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBSetupModItemEUTRANExtIEs, order_DRBSetupModItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBSetupModItemEUTRAN) Pack(stream *Stream) {
    const dataForwardingInformationResponse_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.S1DLUPTNLInformation.Pack(stream)
    if self.DataForwardingInformationResponse != nil { 
        _flags |= dataForwardingInformationResponse_flag
        self.DataForwardingInformationResponse.Pack(stream)
    }//end of optional
    self.ULUPTransportParameters.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBSetupModItemEUTRANExtIEs, order_DRBSetupModItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

func (self *DRBSetupListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBSetupItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBSetupListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBSetupListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Setup-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBSetupItemNGRAN
}

type DRBSetupItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Data-Forwarding-Information', 'name': 'dRB-data-Forwarding-Information-Response', 'optional': True}, {'type': 'UP-Parameters', 'name': 'uL-UP-Transport-Parameters'}, {'type': 'QoS-Flow-List', 'name': 'flow-Setup-List'}, {'type': 'QoS-Flow-Failed-List', 'name': 'flow-Failed-List', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Setup-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    DRBdataForwardingInformationResponse *DataForwardingInformation
    ULUPTransportParameters UPParameters
    FlowSetupList QoSFlowList
    FlowFailedList *QoSFlowFailedList
    IEExtensions *DRBSetupItemNGRANExtIEs
}

func (self * DRBSetupItemNGRAN) Unpack(stream *Stream) {
    dRBdataForwardingInformationResponse_flag := 0x00000002
    flowFailedList_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.DRBID.Unpack(stream)// p8
    if (dRBdataForwardingInformationResponse_flag & _flags) == dRBdataForwardingInformationResponse_flag { //cond2
        self.DRBdataForwardingInformationResponse = &DataForwardingInformation{}//7{'type': 'Data-Forwarding-Information', 'name': 'dRB-data-Forwarding-Information-Response', 'optional': True}
        self.DRBdataForwardingInformationResponse.Unpack(stream)// p8
    }
    self.ULUPTransportParameters.Unpack(stream)// p8
    self.FlowSetupList.Unpack(stream)// p8
    if (flowFailedList_flag & _flags) == flowFailedList_flag { //cond2
        self.FlowFailedList = &QoSFlowFailedList{}//7{'type': 'QoS-Flow-Failed-List', 'name': 'flow-Failed-List', 'optional': True}
        self.FlowFailedList.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBSetupItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Setup-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBSetupItemNGRANExtIEs, order_DRBSetupItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBSetupItemNGRAN) Pack(stream *Stream) {
    const dRBdataForwardingInformationResponse_flag uint = 0x00000002
    const flowFailedList_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.DRBdataForwardingInformationResponse != nil { 
        _flags |= dRBdataForwardingInformationResponse_flag
        self.DRBdataForwardingInformationResponse.Pack(stream)
    }//end of optional
    self.ULUPTransportParameters.Pack(stream)
    self.FlowSetupList.Pack(stream)
    if self.FlowFailedList != nil { 
        _flags |= flowFailedList_flag
        self.FlowFailedList.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBSetupItemNGRANExtIEs, order_DRBSetupItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

func (self *DRBSetupModListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBSetupModItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBSetupModListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBSetupModListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Setup-Mod-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBSetupModItemNGRAN
}

type DRBSetupModItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Data-Forwarding-Information', 'name': 'dRB-data-Forwarding-Information-Response', 'optional': True}, {'type': 'UP-Parameters', 'name': 'uL-UP-Transport-Parameters'}, {'type': 'QoS-Flow-List', 'name': 'flow-Setup-List'}, {'type': 'QoS-Flow-Failed-List', 'name': 'flow-Failed-List', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Setup-Mod-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    DRBdataForwardingInformationResponse *DataForwardingInformation
    ULUPTransportParameters UPParameters
    FlowSetupList QoSFlowList
    FlowFailedList *QoSFlowFailedList
    IEExtensions *DRBSetupModItemNGRANExtIEs
}

func (self * DRBSetupModItemNGRAN) Unpack(stream *Stream) {
    dRBdataForwardingInformationResponse_flag := 0x00000002
    flowFailedList_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.DRBID.Unpack(stream)// p8
    if (dRBdataForwardingInformationResponse_flag & _flags) == dRBdataForwardingInformationResponse_flag { //cond2
        self.DRBdataForwardingInformationResponse = &DataForwardingInformation{}//7{'type': 'Data-Forwarding-Information', 'name': 'dRB-data-Forwarding-Information-Response', 'optional': True}
        self.DRBdataForwardingInformationResponse.Unpack(stream)// p8
    }
    self.ULUPTransportParameters.Unpack(stream)// p8
    self.FlowSetupList.Unpack(stream)// p8
    if (flowFailedList_flag & _flags) == flowFailedList_flag { //cond2
        self.FlowFailedList = &QoSFlowFailedList{}//7{'type': 'QoS-Flow-Failed-List', 'name': 'flow-Failed-List', 'optional': True}
        self.FlowFailedList.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBSetupModItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Setup-Mod-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBSetupModItemNGRANExtIEs, order_DRBSetupModItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBSetupModItemNGRAN) Pack(stream *Stream) {
    const dRBdataForwardingInformationResponse_flag uint = 0x00000002
    const flowFailedList_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.DRBdataForwardingInformationResponse != nil { 
        _flags |= dRBdataForwardingInformationResponse_flag
        self.DRBdataForwardingInformationResponse.Pack(stream)
    }//end of optional
    self.ULUPTransportParameters.Pack(stream)
    self.FlowSetupList.Pack(stream)
    if self.FlowFailedList != nil { 
        _flags |= flowFailedList_flag
        self.FlowFailedList.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBSetupModItemNGRANExtIEs, order_DRBSetupModItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type DRBStatusItem struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'PDCP-Count', 'name': 'pDCP-DL-Count', 'optional': True}, {'type': 'PDCP-Count', 'name': 'pDCP-UL-Count', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Status-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    PDCPDLCount *PDCPCount
    PDCPULCount *PDCPCount
    IEExtensions *DRBStatusItemExtIEs
}

func (self * DRBStatusItem) Unpack(stream *Stream) {
    pDCPDLCount_flag := 0x00000002
    pDCPULCount_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.DRBID.Unpack(stream)// p8
    if (pDCPDLCount_flag & _flags) == pDCPDLCount_flag { //cond2
        self.PDCPDLCount = &PDCPCount{}//7{'type': 'PDCP-Count', 'name': 'pDCP-DL-Count', 'optional': True}
        self.PDCPDLCount.Unpack(stream)// p8
    }
    if (pDCPULCount_flag & _flags) == pDCPULCount_flag { //cond2
        self.PDCPULCount = &PDCPCount{}//7{'type': 'PDCP-Count', 'name': 'pDCP-UL-Count', 'optional': True}
        self.PDCPULCount.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBStatusItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Status-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBStatusItemExtIEs, order_DRBStatusItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBStatusItem) Pack(stream *Stream) {
    const pDCPDLCount_flag uint = 0x00000002
    const pDCPULCount_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.PDCPDLCount != nil { 
        _flags |= pDCPDLCount_flag
        self.PDCPDLCount.Pack(stream)
    }//end of optional
    if self.PDCPULCount != nil { 
        _flags |= pDCPULCount_flag
        self.PDCPULCount.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBStatusItemExtIEs, order_DRBStatusItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

func (self *DRBsSubjectToCounterCheckListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBsSubjectToCounterCheckItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBsSubjectToCounterCheckListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBsSubjectToCounterCheckListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRBs-Subject-To-Counter-Check-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBsSubjectToCounterCheckItemEUTRAN
}

type DRBsSubjectToCounterCheckItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'PDCP-Count', 'name': 'pDCP-UL-Count'}, {'type': 'PDCP-Count', 'name': 'pDCP-DL-Count'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRBs-Subject-To-Counter-Check-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    PDCPULCount PDCPCount
    PDCPDLCount PDCPCount
    IEExtensions *DRBsSubjectToCounterCheckItemEUTRANExtIEs
}

func (self * DRBsSubjectToCounterCheckItemEUTRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.PDCPULCount.Unpack(stream)// p8
    self.PDCPDLCount.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBsSubjectToCounterCheckItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRBs-Subject-To-Counter-Check-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBsSubjectToCounterCheckItemEUTRANExtIEs, order_DRBsSubjectToCounterCheckItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBsSubjectToCounterCheckItemEUTRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.PDCPULCount.Pack(stream)
    self.PDCPDLCount.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBsSubjectToCounterCheckItemEUTRANExtIEs, order_DRBsSubjectToCounterCheckItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBsSubjectToCounterCheckListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBsSubjectToCounterCheckItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBsSubjectToCounterCheckListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBsSubjectToCounterCheckListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRBs-Subject-To-Counter-Check-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBsSubjectToCounterCheckItemNGRAN
}

type DRBsSubjectToCounterCheckItemNGRAN struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'PDCP-Count', 'name': 'pDCP-UL-Count'}, {'type': 'PDCP-Count', 'name': 'pDCP-DL-Count'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRBs-Subject-To-Counter-Check-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    DRBID DRBID
    PDCPULCount PDCPCount
    PDCPDLCount PDCPCount
    IEExtensions *DRBsSubjectToCounterCheckItemNGRANExtIEs
}

func (self * DRBsSubjectToCounterCheckItemNGRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PDUSessionID.Unpack(stream)// p8
    self.DRBID.Unpack(stream)// p8
    self.PDCPULCount.Unpack(stream)// p8
    self.PDCPDLCount.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBsSubjectToCounterCheckItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRBs-Subject-To-Counter-Check-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBsSubjectToCounterCheckItemNGRANExtIEs, order_DRBsSubjectToCounterCheckItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBsSubjectToCounterCheckItemNGRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    self.DRBID.Pack(stream)
    self.PDCPULCount.Pack(stream)
    self.PDCPDLCount.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBsSubjectToCounterCheckItemNGRANExtIEs, order_DRBsSubjectToCounterCheckItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBToModifyListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBToModifyItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBToModifyListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBToModifyListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-To-Modify-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBToModifyItemEUTRAN
}

type DRBToModifyItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'PDCP-Configuration', 'name': 'pDCP-Configuration', 'optional': True}, {'type': 'EUTRAN-QoS', 'name': 'eUTRAN-QoS', 'optional': True}, {'type': 'UP-TNL-Information', 'name': 's1-UL-UP-TNL-Information', 'optional': True}, {'type': 'Data-Forwarding-Information', 'name': 'data-Forwarding-Information', 'optional': True}, {'type': 'PDCP-SN-Status-Request', 'name': 'pDCP-SN-Status-Request', 'optional': True}, {'type': 'PDCP-SN-Status-Information', 'name': 'pDCP-SN-Status-Information', 'optional': True}, {'type': 'UP-Parameters', 'name': 'dL-UP-Parameters', 'optional': True}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Add', 'optional': True}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Modify', 'optional': True}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Remove', 'optional': True}, {'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Modify-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    PDCPConfiguration *PDCPConfiguration
    EUTRANQoS *EUTRANQoS
    S1ULUPTNLInformation *UPTNLInformation
    DataForwardingInformation *DataForwardingInformation
    PDCPSNStatusRequest *PDCPSNStatusRequest
    PDCPSNStatusInformation *PDCPSNStatusInformation
    DLUPParameters *UPParameters
    CellGroupToAdd *CellGroupInformation
    CellGroupToModify *CellGroupInformation
    CellGroupToRemove *CellGroupInformation
    DRBInactivityTimer *InactivityTimer
    IEExtensions *DRBToModifyItemEUTRANExtIEs
}

func (self * DRBToModifyItemEUTRAN) Unpack(stream *Stream) {
    pDCPConfiguration_flag := 0x00000002
    eUTRANQoS_flag := 0x00000004
    s1ULUPTNLInformation_flag := 0x00000008
    dataForwardingInformation_flag := 0x00000010
    pDCPSNStatusRequest_flag := 0x00000020
    pDCPSNStatusInformation_flag := 0x00000040
    dLUPParameters_flag := 0x00000080
    cellGroupToAdd_flag := 0x00000100
    cellGroupToModify_flag := 0x00000200
    cellGroupToRemove_flag := 0x00000400
    dRBInactivityTimer_flag := 0x00000800
    iEExtensions_flag := 0x00001000
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(13)
    self.DRBID.Unpack(stream)// p8
    if (pDCPConfiguration_flag & _flags) == pDCPConfiguration_flag { //cond2
        self.PDCPConfiguration = &PDCPConfiguration{}//7{'type': 'PDCP-Configuration', 'name': 'pDCP-Configuration', 'optional': True}
        self.PDCPConfiguration.Unpack(stream)// p8
    }
    if (eUTRANQoS_flag & _flags) == eUTRANQoS_flag { //cond2
        self.EUTRANQoS = &EUTRANQoS{}//7{'type': 'EUTRAN-QoS', 'name': 'eUTRAN-QoS', 'optional': True}
        self.EUTRANQoS.Unpack(stream)// p8
    }
    if (s1ULUPTNLInformation_flag & _flags) == s1ULUPTNLInformation_flag { //cond2
        self.S1ULUPTNLInformation = &UPTNLInformation{}//7{'type': 'UP-TNL-Information', 'name': 's1-UL-UP-TNL-Information', 'optional': True}
        self.S1ULUPTNLInformation.Unpack(stream)// p8
    }
    if (dataForwardingInformation_flag & _flags) == dataForwardingInformation_flag { //cond2
        self.DataForwardingInformation = &DataForwardingInformation{}//7{'type': 'Data-Forwarding-Information', 'name': 'data-Forwarding-Information', 'optional': True}
        self.DataForwardingInformation.Unpack(stream)// p8
    }
    if (pDCPSNStatusRequest_flag & _flags) == pDCPSNStatusRequest_flag { //cond2
        self.PDCPSNStatusRequest = &PDCPSNStatusRequest{}//7{'type': 'PDCP-SN-Status-Request', 'name': 'pDCP-SN-Status-Request', 'optional': True}
        self.PDCPSNStatusRequest.Unpack(stream)// p8
    }
    if (pDCPSNStatusInformation_flag & _flags) == pDCPSNStatusInformation_flag { //cond2
        self.PDCPSNStatusInformation = &PDCPSNStatusInformation{}//7{'type': 'PDCP-SN-Status-Information', 'name': 'pDCP-SN-Status-Information', 'optional': True}
        self.PDCPSNStatusInformation.Unpack(stream)// p8
    }
    if (dLUPParameters_flag & _flags) == dLUPParameters_flag { //cond2
        self.DLUPParameters = &UPParameters{}//7{'type': 'UP-Parameters', 'name': 'dL-UP-Parameters', 'optional': True}
        self.DLUPParameters.Unpack(stream)// p8
    }
    if (cellGroupToAdd_flag & _flags) == cellGroupToAdd_flag { //cond2
        self.CellGroupToAdd = &CellGroupInformation{}//7{'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Add', 'optional': True}
        self.CellGroupToAdd.Unpack(stream)// p8
    }
    if (cellGroupToModify_flag & _flags) == cellGroupToModify_flag { //cond2
        self.CellGroupToModify = &CellGroupInformation{}//7{'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Modify', 'optional': True}
        self.CellGroupToModify.Unpack(stream)// p8
    }
    if (cellGroupToRemove_flag & _flags) == cellGroupToRemove_flag { //cond2
        self.CellGroupToRemove = &CellGroupInformation{}//7{'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Remove', 'optional': True}
        self.CellGroupToRemove.Unpack(stream)// p8
    }
    if (dRBInactivityTimer_flag & _flags) == dRBInactivityTimer_flag { //cond2
        self.DRBInactivityTimer = &InactivityTimer{}//7{'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}
        self.DRBInactivityTimer.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBToModifyItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Modify-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBToModifyItemEUTRANExtIEs, order_DRBToModifyItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBToModifyItemEUTRAN) Pack(stream *Stream) {
    const pDCPConfiguration_flag uint = 0x00000002
    const eUTRANQoS_flag uint = 0x00000004
    const s1ULUPTNLInformation_flag uint = 0x00000008
    const dataForwardingInformation_flag uint = 0x00000010
    const pDCPSNStatusRequest_flag uint = 0x00000020
    const pDCPSNStatusInformation_flag uint = 0x00000040
    const dLUPParameters_flag uint = 0x00000080
    const cellGroupToAdd_flag uint = 0x00000100
    const cellGroupToModify_flag uint = 0x00000200
    const cellGroupToRemove_flag uint = 0x00000400
    const dRBInactivityTimer_flag uint = 0x00000800
    const iEExtensions_flag uint = 0x00001000
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(13)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.PDCPConfiguration != nil { 
        _flags |= pDCPConfiguration_flag
        self.PDCPConfiguration.Pack(stream)
    }//end of optional
    if self.EUTRANQoS != nil { 
        _flags |= eUTRANQoS_flag
        self.EUTRANQoS.Pack(stream)
    }//end of optional
    if self.S1ULUPTNLInformation != nil { 
        _flags |= s1ULUPTNLInformation_flag
        self.S1ULUPTNLInformation.Pack(stream)
    }//end of optional
    if self.DataForwardingInformation != nil { 
        _flags |= dataForwardingInformation_flag
        self.DataForwardingInformation.Pack(stream)
    }//end of optional
    if self.PDCPSNStatusRequest != nil { 
        _flags |= pDCPSNStatusRequest_flag
        self.PDCPSNStatusRequest.Pack(stream)
    }//end of optional
    if self.PDCPSNStatusInformation != nil { 
        _flags |= pDCPSNStatusInformation_flag
        self.PDCPSNStatusInformation.Pack(stream)
    }//end of optional
    if self.DLUPParameters != nil { 
        _flags |= dLUPParameters_flag
        self.DLUPParameters.Pack(stream)
    }//end of optional
    if self.CellGroupToAdd != nil { 
        _flags |= cellGroupToAdd_flag
        self.CellGroupToAdd.Pack(stream)
    }//end of optional
    if self.CellGroupToModify != nil { 
        _flags |= cellGroupToModify_flag
        self.CellGroupToModify.Pack(stream)
    }//end of optional
    if self.CellGroupToRemove != nil { 
        _flags |= cellGroupToRemove_flag
        self.CellGroupToRemove.Pack(stream)
    }//end of optional
    if self.DRBInactivityTimer != nil { 
        _flags |= dRBInactivityTimer_flag
        self.DRBInactivityTimer.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBToModifyItemEUTRANExtIEs, order_DRBToModifyItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 13)
}//end

func (self *DRBToModifyListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBToModifyItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBToModifyListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBToModifyListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-To-Modify-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBToModifyItemNGRAN
}

type DRBToModifyItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'SDAP-Configuration', 'name': 'sDAP-Configuration', 'optional': True}, {'type': 'PDCP-Configuration', 'name': 'pDCP-Configuration', 'optional': True}, {'type': 'Data-Forwarding-Information', 'name': 'dRB-Data-Forwarding-Information', 'optional': True}, {'type': 'PDCP-SN-Status-Request', 'name': 'pDCP-SN-Status-Request', 'optional': True}, {'type': 'PDCP-SN-Status-Information', 'name': 'pdcp-SN-Status-Information', 'optional': True}, {'type': 'UP-Parameters', 'name': 'dL-UP-Parameters', 'optional': True}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Add', 'optional': True}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Modify', 'optional': True}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Remove', 'optional': True}, {'type': 'QoS-Flow-QoS-Parameter-List', 'name': 'flow-Mapping-Information', 'optional': True}, {'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Modify-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    SDAPConfiguration *SDAPConfiguration
    PDCPConfiguration *PDCPConfiguration
    DRBDataForwardingInformation *DataForwardingInformation
    PDCPSNStatusRequest *PDCPSNStatusRequest
    PdcpSNStatusInformation *PDCPSNStatusInformation
    DLUPParameters *UPParameters
    CellGroupToAdd *CellGroupInformation
    CellGroupToModify *CellGroupInformation
    CellGroupToRemove *CellGroupInformation
    FlowMappingInformation *QoSFlowQoSParameterList
    DRBInactivityTimer *InactivityTimer
    IEExtensions *DRBToModifyItemNGRANExtIEs
}

func (self * DRBToModifyItemNGRAN) Unpack(stream *Stream) {
    sDAPConfiguration_flag := 0x00000002
    pDCPConfiguration_flag := 0x00000004
    dRBDataForwardingInformation_flag := 0x00000008
    pDCPSNStatusRequest_flag := 0x00000010
    pdcpSNStatusInformation_flag := 0x00000020
    dLUPParameters_flag := 0x00000040
    cellGroupToAdd_flag := 0x00000080
    cellGroupToModify_flag := 0x00000100
    cellGroupToRemove_flag := 0x00000200
    flowMappingInformation_flag := 0x00000400
    dRBInactivityTimer_flag := 0x00000800
    iEExtensions_flag := 0x00001000
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(13)
    self.DRBID.Unpack(stream)// p8
    if (sDAPConfiguration_flag & _flags) == sDAPConfiguration_flag { //cond2
        self.SDAPConfiguration = &SDAPConfiguration{}//7{'type': 'SDAP-Configuration', 'name': 'sDAP-Configuration', 'optional': True}
        self.SDAPConfiguration.Unpack(stream)// p8
    }
    if (pDCPConfiguration_flag & _flags) == pDCPConfiguration_flag { //cond2
        self.PDCPConfiguration = &PDCPConfiguration{}//7{'type': 'PDCP-Configuration', 'name': 'pDCP-Configuration', 'optional': True}
        self.PDCPConfiguration.Unpack(stream)// p8
    }
    if (dRBDataForwardingInformation_flag & _flags) == dRBDataForwardingInformation_flag { //cond2
        self.DRBDataForwardingInformation = &DataForwardingInformation{}//7{'type': 'Data-Forwarding-Information', 'name': 'dRB-Data-Forwarding-Information', 'optional': True}
        self.DRBDataForwardingInformation.Unpack(stream)// p8
    }
    if (pDCPSNStatusRequest_flag & _flags) == pDCPSNStatusRequest_flag { //cond2
        self.PDCPSNStatusRequest = &PDCPSNStatusRequest{}//7{'type': 'PDCP-SN-Status-Request', 'name': 'pDCP-SN-Status-Request', 'optional': True}
        self.PDCPSNStatusRequest.Unpack(stream)// p8
    }
    if (pdcpSNStatusInformation_flag & _flags) == pdcpSNStatusInformation_flag { //cond2
        self.PdcpSNStatusInformation = &PDCPSNStatusInformation{}//7{'type': 'PDCP-SN-Status-Information', 'name': 'pdcp-SN-Status-Information', 'optional': True}
        self.PdcpSNStatusInformation.Unpack(stream)// p8
    }
    if (dLUPParameters_flag & _flags) == dLUPParameters_flag { //cond2
        self.DLUPParameters = &UPParameters{}//7{'type': 'UP-Parameters', 'name': 'dL-UP-Parameters', 'optional': True}
        self.DLUPParameters.Unpack(stream)// p8
    }
    if (cellGroupToAdd_flag & _flags) == cellGroupToAdd_flag { //cond2
        self.CellGroupToAdd = &CellGroupInformation{}//7{'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Add', 'optional': True}
        self.CellGroupToAdd.Unpack(stream)// p8
    }
    if (cellGroupToModify_flag & _flags) == cellGroupToModify_flag { //cond2
        self.CellGroupToModify = &CellGroupInformation{}//7{'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Modify', 'optional': True}
        self.CellGroupToModify.Unpack(stream)// p8
    }
    if (cellGroupToRemove_flag & _flags) == cellGroupToRemove_flag { //cond2
        self.CellGroupToRemove = &CellGroupInformation{}//7{'type': 'Cell-Group-Information', 'name': 'cell-Group-To-Remove', 'optional': True}
        self.CellGroupToRemove.Unpack(stream)// p8
    }
    if (flowMappingInformation_flag & _flags) == flowMappingInformation_flag { //cond2
        self.FlowMappingInformation = &QoSFlowQoSParameterList{}//7{'type': 'QoS-Flow-QoS-Parameter-List', 'name': 'flow-Mapping-Information', 'optional': True}
        self.FlowMappingInformation.Unpack(stream)// p8
    }
    if (dRBInactivityTimer_flag & _flags) == dRBInactivityTimer_flag { //cond2
        self.DRBInactivityTimer = &InactivityTimer{}//7{'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}
        self.DRBInactivityTimer.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBToModifyItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Modify-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBToModifyItemNGRANExtIEs, order_DRBToModifyItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBToModifyItemNGRAN) Pack(stream *Stream) {
    const sDAPConfiguration_flag uint = 0x00000002
    const pDCPConfiguration_flag uint = 0x00000004
    const dRBDataForwardingInformation_flag uint = 0x00000008
    const pDCPSNStatusRequest_flag uint = 0x00000010
    const pdcpSNStatusInformation_flag uint = 0x00000020
    const dLUPParameters_flag uint = 0x00000040
    const cellGroupToAdd_flag uint = 0x00000080
    const cellGroupToModify_flag uint = 0x00000100
    const cellGroupToRemove_flag uint = 0x00000200
    const flowMappingInformation_flag uint = 0x00000400
    const dRBInactivityTimer_flag uint = 0x00000800
    const iEExtensions_flag uint = 0x00001000
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(13)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.SDAPConfiguration != nil { 
        _flags |= sDAPConfiguration_flag
        self.SDAPConfiguration.Pack(stream)
    }//end of optional
    if self.PDCPConfiguration != nil { 
        _flags |= pDCPConfiguration_flag
        self.PDCPConfiguration.Pack(stream)
    }//end of optional
    if self.DRBDataForwardingInformation != nil { 
        _flags |= dRBDataForwardingInformation_flag
        self.DRBDataForwardingInformation.Pack(stream)
    }//end of optional
    if self.PDCPSNStatusRequest != nil { 
        _flags |= pDCPSNStatusRequest_flag
        self.PDCPSNStatusRequest.Pack(stream)
    }//end of optional
    if self.PdcpSNStatusInformation != nil { 
        _flags |= pdcpSNStatusInformation_flag
        self.PdcpSNStatusInformation.Pack(stream)
    }//end of optional
    if self.DLUPParameters != nil { 
        _flags |= dLUPParameters_flag
        self.DLUPParameters.Pack(stream)
    }//end of optional
    if self.CellGroupToAdd != nil { 
        _flags |= cellGroupToAdd_flag
        self.CellGroupToAdd.Pack(stream)
    }//end of optional
    if self.CellGroupToModify != nil { 
        _flags |= cellGroupToModify_flag
        self.CellGroupToModify.Pack(stream)
    }//end of optional
    if self.CellGroupToRemove != nil { 
        _flags |= cellGroupToRemove_flag
        self.CellGroupToRemove.Pack(stream)
    }//end of optional
    if self.FlowMappingInformation != nil { 
        _flags |= flowMappingInformation_flag
        self.FlowMappingInformation.Pack(stream)
    }//end of optional
    if self.DRBInactivityTimer != nil { 
        _flags |= dRBInactivityTimer_flag
        self.DRBInactivityTimer.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBToModifyItemNGRANExtIEs, order_DRBToModifyItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 13)
}//end

func (self *DRBToRemoveListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBToRemoveItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBToRemoveListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBToRemoveListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-To-Remove-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBToRemoveItemEUTRAN
}

type DRBToRemoveItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Remove-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    IEExtensions *DRBToRemoveItemEUTRANExtIEs
}

func (self * DRBToRemoveItemEUTRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBToRemoveItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Remove-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBToRemoveItemEUTRANExtIEs, order_DRBToRemoveItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBToRemoveItemEUTRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBToRemoveItemEUTRANExtIEs, order_DRBToRemoveItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBRequiredToRemoveListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBRequiredToRemoveItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBRequiredToRemoveListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBRequiredToRemoveListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Required-To-Remove-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBRequiredToRemoveItemEUTRAN
}

type DRBRequiredToRemoveItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Required-To-Remove-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    Cause Cause
    IEExtensions *DRBRequiredToRemoveItemEUTRANExtIEs
}

func (self * DRBRequiredToRemoveItemEUTRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBRequiredToRemoveItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Required-To-Remove-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBRequiredToRemoveItemEUTRANExtIEs, order_DRBRequiredToRemoveItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBRequiredToRemoveItemEUTRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBRequiredToRemoveItemEUTRANExtIEs, order_DRBRequiredToRemoveItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBToRemoveListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBToRemoveItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBToRemoveListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBToRemoveListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-To-Remove-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBToRemoveItemNGRAN
}

type DRBToRemoveItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Remove-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    IEExtensions *DRBToRemoveItemNGRANExtIEs
}

func (self * DRBToRemoveItemNGRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBToRemoveItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Remove-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBToRemoveItemNGRANExtIEs, order_DRBToRemoveItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBToRemoveItemNGRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBToRemoveItemNGRANExtIEs, order_DRBToRemoveItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBRequiredToRemoveListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBRequiredToRemoveItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBRequiredToRemoveListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBRequiredToRemoveListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Required-To-Remove-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBRequiredToRemoveItemNGRAN
}

type DRBRequiredToRemoveItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Required-To-Remove-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    Cause Cause
    IEExtensions *DRBRequiredToRemoveItemNGRANExtIEs
}

func (self * DRBRequiredToRemoveItemNGRAN) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DRBID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBRequiredToRemoveItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Required-To-Remove-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBRequiredToRemoveItemNGRANExtIEs, order_DRBRequiredToRemoveItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBRequiredToRemoveItemNGRAN) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBRequiredToRemoveItemNGRANExtIEs, order_DRBRequiredToRemoveItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *DRBToSetupListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBToSetupItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBToSetupListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBToSetupListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-To-Setup-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBToSetupItemEUTRAN
}

type DRBToSetupItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'PDCP-Configuration', 'name': 'pDCP-Configuration'}, {'type': 'EUTRAN-QoS', 'name': 'eUTRAN-QoS'}, {'type': 'UP-TNL-Information', 'name': 's1-UL-UP-TNL-Information'}, {'type': 'Data-Forwarding-Information-Request', 'name': 'data-Forwarding-Information-Request', 'optional': True}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-Information'}, {'type': 'UP-Parameters', 'name': 'dL-UP-Parameters', 'optional': True}, {'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}, {'type': 'UP-TNL-Information', 'name': 'existing-Allocated-S1-DL-UP-TNL-Info', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Setup-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    PDCPConfiguration PDCPConfiguration
    EUTRANQoS EUTRANQoS
    S1ULUPTNLInformation UPTNLInformation
    DataForwardingInformationRequest *DataForwardingInformationRequest
    CellGroupInformation CellGroupInformation
    DLUPParameters *UPParameters
    DRBInactivityTimer *InactivityTimer
    ExistingAllocatedS1DLUPTNLInfo *UPTNLInformation
    IEExtensions *DRBToSetupItemEUTRANExtIEs
}

func (self * DRBToSetupItemEUTRAN) Unpack(stream *Stream) {
    dataForwardingInformationRequest_flag := 0x00000002
    dLUPParameters_flag := 0x00000004
    dRBInactivityTimer_flag := 0x00000008
    existingAllocatedS1DLUPTNLInfo_flag := 0x00000010
    iEExtensions_flag := 0x00000020
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(6)
    self.DRBID.Unpack(stream)// p8
    self.PDCPConfiguration.Unpack(stream)// p8
    self.EUTRANQoS.Unpack(stream)// p8
    self.S1ULUPTNLInformation.Unpack(stream)// p8
    if (dataForwardingInformationRequest_flag & _flags) == dataForwardingInformationRequest_flag { //cond2
        self.DataForwardingInformationRequest = &DataForwardingInformationRequest{}//7{'type': 'Data-Forwarding-Information-Request', 'name': 'data-Forwarding-Information-Request', 'optional': True}
        self.DataForwardingInformationRequest.Unpack(stream)// p8
    }
    self.CellGroupInformation.Unpack(stream)// p8
    if (dLUPParameters_flag & _flags) == dLUPParameters_flag { //cond2
        self.DLUPParameters = &UPParameters{}//7{'type': 'UP-Parameters', 'name': 'dL-UP-Parameters', 'optional': True}
        self.DLUPParameters.Unpack(stream)// p8
    }
    if (dRBInactivityTimer_flag & _flags) == dRBInactivityTimer_flag { //cond2
        self.DRBInactivityTimer = &InactivityTimer{}//7{'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}
        self.DRBInactivityTimer.Unpack(stream)// p8
    }
    if (existingAllocatedS1DLUPTNLInfo_flag & _flags) == existingAllocatedS1DLUPTNLInfo_flag { //cond2
        self.ExistingAllocatedS1DLUPTNLInfo = &UPTNLInformation{}//7{'type': 'UP-TNL-Information', 'name': 'existing-Allocated-S1-DL-UP-TNL-Info', 'optional': True}
        self.ExistingAllocatedS1DLUPTNLInfo.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBToSetupItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Setup-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBToSetupItemEUTRANExtIEs, order_DRBToSetupItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBToSetupItemEUTRAN) Pack(stream *Stream) {
    const dataForwardingInformationRequest_flag uint = 0x00000002
    const dLUPParameters_flag uint = 0x00000004
    const dRBInactivityTimer_flag uint = 0x00000008
    const existingAllocatedS1DLUPTNLInfo_flag uint = 0x00000010
    const iEExtensions_flag uint = 0x00000020
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(6)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.PDCPConfiguration.Pack(stream)
    self.EUTRANQoS.Pack(stream)
    self.S1ULUPTNLInformation.Pack(stream)
    if self.DataForwardingInformationRequest != nil { 
        _flags |= dataForwardingInformationRequest_flag
        self.DataForwardingInformationRequest.Pack(stream)
    }//end of optional
    self.CellGroupInformation.Pack(stream)
    if self.DLUPParameters != nil { 
        _flags |= dLUPParameters_flag
        self.DLUPParameters.Pack(stream)
    }//end of optional
    if self.DRBInactivityTimer != nil { 
        _flags |= dRBInactivityTimer_flag
        self.DRBInactivityTimer.Pack(stream)
    }//end of optional
    if self.ExistingAllocatedS1DLUPTNLInfo != nil { 
        _flags |= existingAllocatedS1DLUPTNLInfo_flag
        self.ExistingAllocatedS1DLUPTNLInfo.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBToSetupItemEUTRANExtIEs, order_DRBToSetupItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

func (self *DRBToSetupModListEUTRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBToSetupModItemEUTRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBToSetupModListEUTRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBToSetupModListEUTRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-To-Setup-Mod-Item-EUTRAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBToSetupModItemEUTRAN
}

type DRBToSetupModItemEUTRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'PDCP-Configuration', 'name': 'pDCP-Configuration'}, {'type': 'EUTRAN-QoS', 'name': 'eUTRAN-QoS'}, {'type': 'UP-TNL-Information', 'name': 's1-UL-UP-TNL-Information'}, {'type': 'Data-Forwarding-Information-Request', 'name': 'data-Forwarding-Information-Request', 'optional': True}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-Information'}, {'type': 'UP-Parameters', 'name': 'dL-UP-Parameters', 'optional': True}, {'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Setup-Mod-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    PDCPConfiguration PDCPConfiguration
    EUTRANQoS EUTRANQoS
    S1ULUPTNLInformation UPTNLInformation
    DataForwardingInformationRequest *DataForwardingInformationRequest
    CellGroupInformation CellGroupInformation
    DLUPParameters *UPParameters
    DRBInactivityTimer *InactivityTimer
    IEExtensions *DRBToSetupModItemEUTRANExtIEs
}

func (self * DRBToSetupModItemEUTRAN) Unpack(stream *Stream) {
    dataForwardingInformationRequest_flag := 0x00000002
    dLUPParameters_flag := 0x00000004
    dRBInactivityTimer_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.DRBID.Unpack(stream)// p8
    self.PDCPConfiguration.Unpack(stream)// p8
    self.EUTRANQoS.Unpack(stream)// p8
    self.S1ULUPTNLInformation.Unpack(stream)// p8
    if (dataForwardingInformationRequest_flag & _flags) == dataForwardingInformationRequest_flag { //cond2
        self.DataForwardingInformationRequest = &DataForwardingInformationRequest{}//7{'type': 'Data-Forwarding-Information-Request', 'name': 'data-Forwarding-Information-Request', 'optional': True}
        self.DataForwardingInformationRequest.Unpack(stream)// p8
    }
    self.CellGroupInformation.Unpack(stream)// p8
    if (dLUPParameters_flag & _flags) == dLUPParameters_flag { //cond2
        self.DLUPParameters = &UPParameters{}//7{'type': 'UP-Parameters', 'name': 'dL-UP-Parameters', 'optional': True}
        self.DLUPParameters.Unpack(stream)// p8
    }
    if (dRBInactivityTimer_flag & _flags) == dRBInactivityTimer_flag { //cond2
        self.DRBInactivityTimer = &InactivityTimer{}//7{'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}
        self.DRBInactivityTimer.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBToSetupModItemEUTRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Setup-Mod-Item-EUTRAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBToSetupModItemEUTRANExtIEs, order_DRBToSetupModItemEUTRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBToSetupModItemEUTRAN) Pack(stream *Stream) {
    const dataForwardingInformationRequest_flag uint = 0x00000002
    const dLUPParameters_flag uint = 0x00000004
    const dRBInactivityTimer_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.PDCPConfiguration.Pack(stream)
    self.EUTRANQoS.Pack(stream)
    self.S1ULUPTNLInformation.Pack(stream)
    if self.DataForwardingInformationRequest != nil { 
        _flags |= dataForwardingInformationRequest_flag
        self.DataForwardingInformationRequest.Pack(stream)
    }//end of optional
    self.CellGroupInformation.Pack(stream)
    if self.DLUPParameters != nil { 
        _flags |= dLUPParameters_flag
        self.DLUPParameters.Pack(stream)
    }//end of optional
    if self.DRBInactivityTimer != nil { 
        _flags |= dRBInactivityTimer_flag
        self.DRBInactivityTimer.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBToSetupModItemEUTRANExtIEs, order_DRBToSetupModItemEUTRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *DRBToSetupListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBToSetupItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBToSetupListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBToSetupListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-To-Setup-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBToSetupItemNGRAN
}

type DRBToSetupItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'SDAP-Configuration', 'name': 'sDAP-Configuration'}, {'type': 'PDCP-Configuration', 'name': 'pDCP-Configuration'}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-Information'}, {'type': 'QoS-Flow-QoS-Parameter-List', 'name': 'qos-flow-Information-To-Be-Setup'}, {'type': 'Data-Forwarding-Information-Request', 'name': 'dRB-Data-Forwarding-Information-Request', 'optional': True}, {'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}, {'type': 'PDCP-SN-Status-Information', 'name': 'pDCP-SN-Status-Information', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Setup-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    SDAPConfiguration SDAPConfiguration
    PDCPConfiguration PDCPConfiguration
    CellGroupInformation CellGroupInformation
    QosflowInformationToBeSetup QoSFlowQoSParameterList
    DRBDataForwardingInformationRequest *DataForwardingInformationRequest
    DRBInactivityTimer *InactivityTimer
    PDCPSNStatusInformation *PDCPSNStatusInformation
    IEExtensions *DRBToSetupItemNGRANExtIEs
}

func (self * DRBToSetupItemNGRAN) Unpack(stream *Stream) {
    dRBDataForwardingInformationRequest_flag := 0x00000002
    dRBInactivityTimer_flag := 0x00000004
    pDCPSNStatusInformation_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.DRBID.Unpack(stream)// p8
    self.SDAPConfiguration.Unpack(stream)// p8
    self.PDCPConfiguration.Unpack(stream)// p8
    self.CellGroupInformation.Unpack(stream)// p8
    self.QosflowInformationToBeSetup.Unpack(stream)// p8
    if (dRBDataForwardingInformationRequest_flag & _flags) == dRBDataForwardingInformationRequest_flag { //cond2
        self.DRBDataForwardingInformationRequest = &DataForwardingInformationRequest{}//7{'type': 'Data-Forwarding-Information-Request', 'name': 'dRB-Data-Forwarding-Information-Request', 'optional': True}
        self.DRBDataForwardingInformationRequest.Unpack(stream)// p8
    }
    if (dRBInactivityTimer_flag & _flags) == dRBInactivityTimer_flag { //cond2
        self.DRBInactivityTimer = &InactivityTimer{}//7{'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}
        self.DRBInactivityTimer.Unpack(stream)// p8
    }
    if (pDCPSNStatusInformation_flag & _flags) == pDCPSNStatusInformation_flag { //cond2
        self.PDCPSNStatusInformation = &PDCPSNStatusInformation{}//7{'type': 'PDCP-SN-Status-Information', 'name': 'pDCP-SN-Status-Information', 'optional': True}
        self.PDCPSNStatusInformation.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBToSetupItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Setup-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBToSetupItemNGRANExtIEs, order_DRBToSetupItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBToSetupItemNGRAN) Pack(stream *Stream) {
    const dRBDataForwardingInformationRequest_flag uint = 0x00000002
    const dRBInactivityTimer_flag uint = 0x00000004
    const pDCPSNStatusInformation_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.SDAPConfiguration.Pack(stream)
    self.PDCPConfiguration.Pack(stream)
    self.CellGroupInformation.Pack(stream)
    self.QosflowInformationToBeSetup.Pack(stream)
    if self.DRBDataForwardingInformationRequest != nil { 
        _flags |= dRBDataForwardingInformationRequest_flag
        self.DRBDataForwardingInformationRequest.Pack(stream)
    }//end of optional
    if self.DRBInactivityTimer != nil { 
        _flags |= dRBInactivityTimer_flag
        self.DRBInactivityTimer.Pack(stream)
    }//end of optional
    if self.PDCPSNStatusInformation != nil { 
        _flags |= pDCPSNStatusInformation_flag
        self.PDCPSNStatusInformation.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBToSetupItemNGRANExtIEs, order_DRBToSetupItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *DRBToSetupModListNGRAN) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]DRBToSetupModItemNGRAN, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBToSetupModListNGRAN) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBToSetupModListNGRAN struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-To-Setup-Mod-Item-NG-RAN'}, 'size': [(1, 'maxnoofDRBs')]}
    Items []DRBToSetupModItemNGRAN
}

type DRBToSetupModItemNGRAN struct { // [{'type': 'DRB-ID', 'name': 'dRB-ID'}, {'type': 'SDAP-Configuration', 'name': 'sDAP-Configuration'}, {'type': 'PDCP-Configuration', 'name': 'pDCP-Configuration'}, {'type': 'Cell-Group-Information', 'name': 'cell-Group-Information'}, {'type': 'QoS-Flow-QoS-Parameter-List', 'name': 'flow-Mapping-Information'}, {'type': 'Data-Forwarding-Information-Request', 'name': 'dRB-Data-Forwarding-Information-Request', 'optional': True}, {'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}, {'type': 'PDCP-SN-Status-Information', 'name': 'pDCP-SN-Status-Information', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Setup-Mod-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DRBID DRBID
    SDAPConfiguration SDAPConfiguration
    PDCPConfiguration PDCPConfiguration
    CellGroupInformation CellGroupInformation
    FlowMappingInformation QoSFlowQoSParameterList
    DRBDataForwardingInformationRequest *DataForwardingInformationRequest
    DRBInactivityTimer *InactivityTimer
    PDCPSNStatusInformation *PDCPSNStatusInformation
    IEExtensions *DRBToSetupModItemNGRANExtIEs
}

func (self * DRBToSetupModItemNGRAN) Unpack(stream *Stream) {
    dRBDataForwardingInformationRequest_flag := 0x00000002
    dRBInactivityTimer_flag := 0x00000004
    pDCPSNStatusInformation_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.DRBID.Unpack(stream)// p8
    self.SDAPConfiguration.Unpack(stream)// p8
    self.PDCPConfiguration.Unpack(stream)// p8
    self.CellGroupInformation.Unpack(stream)// p8
    self.FlowMappingInformation.Unpack(stream)// p8
    if (dRBDataForwardingInformationRequest_flag & _flags) == dRBDataForwardingInformationRequest_flag { //cond2
        self.DRBDataForwardingInformationRequest = &DataForwardingInformationRequest{}//7{'type': 'Data-Forwarding-Information-Request', 'name': 'dRB-Data-Forwarding-Information-Request', 'optional': True}
        self.DRBDataForwardingInformationRequest.Unpack(stream)// p8
    }
    if (dRBInactivityTimer_flag & _flags) == dRBInactivityTimer_flag { //cond2
        self.DRBInactivityTimer = &InactivityTimer{}//7{'type': 'Inactivity-Timer', 'name': 'dRB-Inactivity-Timer', 'optional': True}
        self.DRBInactivityTimer.Unpack(stream)// p8
    }
    if (pDCPSNStatusInformation_flag & _flags) == pDCPSNStatusInformation_flag { //cond2
        self.PDCPSNStatusInformation = &PDCPSNStatusInformation{}//7{'type': 'PDCP-SN-Status-Information', 'name': 'pDCP-SN-Status-Information', 'optional': True}
        self.PDCPSNStatusInformation.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBToSetupModItemNGRANExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-To-Setup-Mod-Item-NG-RAN-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBToSetupModItemNGRANExtIEs, order_DRBToSetupModItemNGRANExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBToSetupModItemNGRAN) Pack(stream *Stream) {
    const dRBDataForwardingInformationRequest_flag uint = 0x00000002
    const dRBInactivityTimer_flag uint = 0x00000004
    const pDCPSNStatusInformation_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DRBID.Pack(stream)
    self.SDAPConfiguration.Pack(stream)
    self.PDCPConfiguration.Pack(stream)
    self.CellGroupInformation.Pack(stream)
    self.FlowMappingInformation.Pack(stream)
    if self.DRBDataForwardingInformationRequest != nil { 
        _flags |= dRBDataForwardingInformationRequest_flag
        self.DRBDataForwardingInformationRequest.Pack(stream)
    }//end of optional
    if self.DRBInactivityTimer != nil { 
        _flags |= dRBInactivityTimer_flag
        self.DRBInactivityTimer.Pack(stream)
    }//end of optional
    if self.PDCPSNStatusInformation != nil { 
        _flags |= pDCPSNStatusInformation_flag
        self.PDCPSNStatusInformation.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBToSetupModItemNGRANExtIEs, order_DRBToSetupModItemNGRANExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *DRBUsageReportList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]DRBUsageReportItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *DRBUsageReportList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type DRBUsageReportList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'DRB-Usage-Report-Item'}, 'size': [(1, 'maxnooftimeperiods')]}
    Items []DRBUsageReportItem
}

type DRBUsageReportItem struct { // [{'type': 'OCTET STRING', 'size': [4], 'name': 'startTimeStamp'}, {'type': 'OCTET STRING', 'size': [4], 'name': 'endTimeStamp'}, {'type': 'INTEGER', 'restricted-to': [(0, 18446744073709551615)], 'name': 'usageCountUL'}, {'type': 'INTEGER', 'restricted-to': [(0, 18446744073709551615)], 'name': 'usageCountDL'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Usage-Report-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    StartTimeStamp OCTETSTRING
    EndTimeStamp OCTETSTRING
    UsageCountUL INTEGER
    UsageCountDL INTEGER
    IEExtensions *DRBUsageReportItemExtIEs
}

func (self * DRBUsageReportItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_startTimeStamp = func(st *Stream, self *OCTETSTRING) {
        self.Value = st.parsef_OctString(4)
    }
    Unpack_startTimeStamp(stream, &self.StartTimeStamp)// p2
    var Unpack_endTimeStamp = func(st *Stream, self *OCTETSTRING) {
        self.Value = st.parsef_OctString(4)
    }
    Unpack_endTimeStamp(stream, &self.EndTimeStamp)// p2
    var Unpack_usageCountUL = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(18446744073709551615, 64, 0, 0)
    }
    Unpack_usageCountUL(stream, &self.UsageCountUL)// p2
    var Unpack_usageCountDL = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(18446744073709551615, 64, 0, 0)
    }
    Unpack_usageCountDL(stream, &self.UsageCountDL)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &DRBUsageReportItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRB-Usage-Report-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_DRBUsageReportItemExtIEs, order_DRBUsageReportItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBUsageReportItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_startTimeStamp = func(st *Stream, self OCTETSTRING) {
        st.formatf_OctString(self.Value, 4)
    }
    Pack_startTimeStamp(stream, self.StartTimeStamp) //f2
    var Pack_endTimeStamp = func(st *Stream, self OCTETSTRING) {
        st.formatf_OctString(self.Value, 4)
    }
    Pack_endTimeStamp(stream, self.EndTimeStamp) //f2
    var Pack_usageCountUL = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 18446744073709551615, 64, 0, 0)
    }
    Pack_usageCountUL(stream, self.UsageCountUL) //f2
    var Pack_usageCountDL = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 18446744073709551615, 64, 0, 0)
    }
    Pack_usageCountDL(stream, self.UsageCountDL) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_DRBUsageReportItemExtIEs, order_DRBUsageReportItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type DuplicationActivation struct {
  Value int
}
const (
    DuplicationActivationactive = 0
    DuplicationActivationinactive = 1

    /* Extensions */
)
func (self *DuplicationActivation) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *DuplicationActivation) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type Dynamic5QIDescriptor struct { // [{'type': 'QoSPriorityLevel', 'name': 'qoSPriorityLevel'}, {'type': 'PacketDelayBudget', 'name': 'packetDelayBudget'}, {'type': 'PacketErrorRate', 'name': 'packetErrorRate'}, {'type': 'INTEGER', 'restricted-to': [(0, 255), None], 'name': 'fiveQI', 'optional': True}, {'type': 'ENUMERATED', 'values': [('delay-critical', 0), ('non-delay-critical', 1)], 'name': 'delayCritical', 'optional': True}, {'type': 'AveragingWindow', 'name': 'averagingWindow', 'optional': True}, {'type': 'MaxDataBurstVolume', 'name': 'maxDataBurstVolume', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Dynamic5QIDescriptor-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    QoSPriorityLevel QoSPriorityLevel
    PacketDelayBudget PacketDelayBudget
    PacketErrorRate PacketErrorRate
    FiveQI *INTEGER
    DelayCritical *ENUMERATED
    AveragingWindow *AveragingWindow
    MaxDataBurstVolume *MaxDataBurstVolume
    IEExtensions *Dynamic5QIDescriptorExtIEs
}

func (self * Dynamic5QIDescriptor) Unpack(stream *Stream) {
    fiveQI_flag := 0x00000001
    delayCritical_flag := 0x00000002
    averagingWindow_flag := 0x00000004
    maxDataBurstVolume_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    _flags := 0
    _flags = stream.get_flags(5)
    self.QoSPriorityLevel.Unpack(stream)// p8
    self.PacketDelayBudget.Unpack(stream)// p8
    self.PacketErrorRate.Unpack(stream)// p8
    if (fiveQI_flag & _flags) == fiveQI_flag { //cond1
        var Unpack_fiveQI = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(256, 9, 1, 0)
        }
        self.FiveQI = &INTEGER{}//6{'type': 'INTEGER', 'restricted-to': [(0, 255), None], 'name': 'fiveQI', 'optional': True}
        Unpack_fiveQI(stream, self.FiveQI)// p1 {'type': 'INTEGER', 'restricted-to': [(0, 255), None], 'name': 'fiveQI', 'optional': True}
    }
    if (delayCritical_flag & _flags) == delayCritical_flag { //cond1
        var Unpack_delayCritical = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 2, 0)
        }
        self.DelayCritical = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('delay-critical', 0), ('non-delay-critical', 1)], 'name': 'delayCritical', 'optional': True}
        Unpack_delayCritical(stream, self.DelayCritical)// p1 {'type': 'ENUMERATED', 'values': [('delay-critical', 0), ('non-delay-critical', 1)], 'name': 'delayCritical', 'optional': True}
    }
    if (averagingWindow_flag & _flags) == averagingWindow_flag { //cond2
        self.AveragingWindow = &AveragingWindow{}//7{'type': 'AveragingWindow', 'name': 'averagingWindow', 'optional': True}
        self.AveragingWindow.Unpack(stream)// p8
    }
    if (maxDataBurstVolume_flag & _flags) == maxDataBurstVolume_flag { //cond2
        self.MaxDataBurstVolume = &MaxDataBurstVolume{}//7{'type': 'MaxDataBurstVolume', 'name': 'maxDataBurstVolume', 'optional': True}
        self.MaxDataBurstVolume.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &Dynamic5QIDescriptorExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Dynamic5QIDescriptor-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_Dynamic5QIDescriptorExtIEs, order_Dynamic5QIDescriptorExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * Dynamic5QIDescriptor) Pack(stream *Stream) {
    const fiveQI_flag uint = 0x00000001
    const delayCritical_flag uint = 0x00000002
    const averagingWindow_flag uint = 0x00000004
    const maxDataBurstVolume_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.QoSPriorityLevel.Pack(stream)
    self.PacketDelayBudget.Pack(stream)
    self.PacketErrorRate.Pack(stream)
    if self.FiveQI != nil { //YY
        _flags |= fiveQI_flag
        var Pack_fiveQI = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 256, 9, 1, 0)
        }
        Pack_fiveQI(stream, *self.FiveQI) //f1
    }//end of optional
    if self.DelayCritical != nil { //YY
        _flags |= delayCritical_flag
        var Pack_delayCritical = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 2, 0)
        }
        Pack_delayCritical(stream, *self.DelayCritical) //f1
    }//end of optional
    if self.AveragingWindow != nil { 
        _flags |= averagingWindow_flag
        self.AveragingWindow.Pack(stream)
    }//end of optional
    if self.MaxDataBurstVolume != nil { 
        _flags |= maxDataBurstVolume_flag
        self.MaxDataBurstVolume.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_Dynamic5QIDescriptorExtIEs, order_Dynamic5QIDescriptorExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

type DataDiscardRequired struct {
  Value int
}
const (
    DataDiscardRequiredrequired = 0

    /* Extensions */
)
func (self *DataDiscardRequired) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *DataDiscardRequired) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
type EncryptionKey struct {
  Value HexBytes
}
func (self *EncryptionKey) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *EncryptionKey) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type EndpointIPaddressandport struct { // [{'type': 'TransportLayerAddress', 'name': 'endpoint-IP-Address'}, {'type': 'PortNumber', 'name': 'portNumber'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Endpoint-IP-address-and-port-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    EndpointIPAddress TransportLayerAddress
    PortNumber PortNumber
    IEExtensions *EndpointIPaddressandportExtIEs
}

func (self * EndpointIPaddressandport) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.EndpointIPAddress.Unpack(stream)// p8
    self.PortNumber.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &EndpointIPaddressandportExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Endpoint-IP-address-and-port-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_EndpointIPaddressandportExtIEs, order_EndpointIPaddressandportExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * EndpointIPaddressandport) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.EndpointIPAddress.Pack(stream)
    self.PortNumber.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_EndpointIPaddressandportExtIEs, order_EndpointIPaddressandportExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EUTRANAllocationAndRetentionPriority struct { // [{'type': 'PriorityLevel', 'name': 'priorityLevel'}, {'type': 'Pre-emptionCapability', 'name': 'pre-emptionCapability'}, {'type': 'Pre-emptionVulnerability', 'name': 'pre-emptionVulnerability'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['EUTRANAllocationAndRetentionPriority-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PriorityLevel PriorityLevel
    PreemptionCapability PreemptionCapability
    PreemptionVulnerability PreemptionVulnerability
    IEExtensions *EUTRANAllocationAndRetentionPriorityExtIEs
}

func (self * EUTRANAllocationAndRetentionPriority) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PriorityLevel.Unpack(stream)// p8
    self.PreemptionCapability.Unpack(stream)// p8
    self.PreemptionVulnerability.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &EUTRANAllocationAndRetentionPriorityExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['EUTRANAllocationAndRetentionPriority-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_EUTRANAllocationAndRetentionPriorityExtIEs, order_EUTRANAllocationAndRetentionPriorityExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EUTRANAllocationAndRetentionPriority) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PriorityLevel.Pack(stream)
    self.PreemptionCapability.Pack(stream)
    self.PreemptionVulnerability.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_EUTRANAllocationAndRetentionPriorityExtIEs, order_EUTRANAllocationAndRetentionPriorityExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *EUTRANQoSSupportList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]EUTRANQoSSupportItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *EUTRANQoSSupportList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type EUTRANQoSSupportList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'EUTRAN-QoS-Support-Item'}, 'size': [(1, 'maxnoofEUTRANQOSParameters')]}
    Items []EUTRANQoSSupportItem
}

type EUTRANQoSSupportItem struct { // [{'type': 'EUTRAN-QoS', 'name': 'eUTRAN-QoS'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['EUTRAN-QoS-Support-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    EUTRANQoS EUTRANQoS
    IEExtensions *EUTRANQoSSupportItemExtIEs
}

func (self * EUTRANQoSSupportItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.EUTRANQoS.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &EUTRANQoSSupportItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['EUTRAN-QoS-Support-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_EUTRANQoSSupportItemExtIEs, order_EUTRANQoSSupportItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * EUTRANQoSSupportItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.EUTRANQoS.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_EUTRANQoSSupportItemExtIEs, order_EUTRANQoSSupportItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EUTRANQoS struct { // [{'type': 'QCI', 'name': 'qCI'}, {'type': 'EUTRANAllocationAndRetentionPriority', 'name': 'eUTRANallocationAndRetentionPriority'}, {'type': 'GBR-QosInformation', 'name': 'gbrQosInformation', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['EUTRAN-QoS-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    QCI QCI
    EUTRANallocationAndRetentionPriority EUTRANAllocationAndRetentionPriority
    GbrQosInformation *GBRQosInformation
    IEExtensions *EUTRANQoSExtIEs
}

func (self * EUTRANQoS) Unpack(stream *Stream) {
    gbrQosInformation_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.QCI.Unpack(stream)// p8
    self.EUTRANallocationAndRetentionPriority.Unpack(stream)// p8
    if (gbrQosInformation_flag & _flags) == gbrQosInformation_flag { //cond2
        self.GbrQosInformation = &GBRQosInformation{}//7{'type': 'GBR-QosInformation', 'name': 'gbrQosInformation', 'optional': True}
        self.GbrQosInformation.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &EUTRANQoSExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['EUTRAN-QoS-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_EUTRANQoSExtIEs, order_EUTRANQoSExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EUTRANQoS) Pack(stream *Stream) {
    const gbrQosInformation_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.QCI.Pack(stream)
    self.EUTRANallocationAndRetentionPriority.Pack(stream)
    if self.GbrQosInformation != nil { 
        _flags |= gbrQosInformation_flag
        self.GbrQosInformation.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_EUTRANQoSExtIEs, order_EUTRANQoSExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type GNBCUCPName struct {
  Value string
}
func (self *GNBCUCPName) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in GNB-CU-CP-Name")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *GNBCUCPName) Pack(st *Stream) {
    _eflag := 0
    if len(self.Value) > 150 {
       _eflag = 1
    }
    st.format_ext(_eflag)
    if len(self.Value) < 1 || len(self.Value) > 150 {
        return;
    }
    st.format_olen((len(self.Value))-1, 8)
    st.formatf_PriString(self.Value, len(self.Value))
}
type GNBCUCPUEE1APID struct {
  Value uint64
}
func (self *GNBCUCPUEE1APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967296, 32, 0, 0)
}
func (self * GNBCUCPUEE1APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967296, 32, 0, 0)
}
type GNBCUUPCapacity struct {
  Value uint64
}
func (self *GNBCUUPCapacity) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * GNBCUUPCapacity) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
func (self *GNBCUUPCellGroupRelatedConfiguration) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(8)
    _size += 1
    self.Items = make([]GNBCUUPCellGroupRelatedConfigurationItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *GNBCUUPCellGroupRelatedConfiguration) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 8)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type GNBCUUPCellGroupRelatedConfiguration struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GNB-CU-UP-CellGroupRelatedConfiguration-Item'}, 'size': [(1, 'maxnoofUPParameters')]}
    Items []GNBCUUPCellGroupRelatedConfigurationItem
}

type GNBCUUPCellGroupRelatedConfigurationItem struct { // [{'type': 'Cell-Group-ID', 'name': 'cell-Group-ID'}, {'type': 'UP-TNL-Information', 'name': 'uP-TNL-Information'}, {'type': 'UL-Configuration', 'name': 'uL-Configuration', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-UP-CellGroupRelatedConfiguration-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    CellGroupID CellGroupID
    UPTNLInformation UPTNLInformation
    ULConfiguration *ULConfiguration
    IEExtensions *GNBCUUPCellGroupRelatedConfigurationItemExtIEs
}

func (self * GNBCUUPCellGroupRelatedConfigurationItem) Unpack(stream *Stream) {
    uLConfiguration_flag := 0x00000001
    iEExtensions_flag := 0x00000002
    _flags := 0
    _flags = stream.get_flags(2)
    self.CellGroupID.Unpack(stream)// p8
    self.UPTNLInformation.Unpack(stream)// p8
    if (uLConfiguration_flag & _flags) == uLConfiguration_flag { //cond2
        self.ULConfiguration = &ULConfiguration{}//7{'type': 'UL-Configuration', 'name': 'uL-Configuration', 'optional': True}
        self.ULConfiguration.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GNBCUUPCellGroupRelatedConfigurationItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-UP-CellGroupRelatedConfiguration-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GNBCUUPCellGroupRelatedConfigurationItemExtIEs, order_GNBCUUPCellGroupRelatedConfigurationItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * GNBCUUPCellGroupRelatedConfigurationItem) Pack(stream *Stream) {
    const uLConfiguration_flag uint = 0x00000001
    const iEExtensions_flag uint = 0x00000002
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CellGroupID.Pack(stream)
    self.UPTNLInformation.Pack(stream)
    if self.ULConfiguration != nil { 
        _flags |= uLConfiguration_flag
        self.ULConfiguration.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GNBCUUPCellGroupRelatedConfigurationItemExtIEs, order_GNBCUUPCellGroupRelatedConfigurationItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GNBCUUPID struct {
  Value uint64
}
func (self *GNBCUUPID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(68719476736, 36, 0, 0)
}
func (self * GNBCUUPID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 68719476736, 36, 0, 0)
}
type GNBCUUPName struct {
  Value string
}
func (self *GNBCUUPName) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in GNB-CU-UP-Name")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *GNBCUUPName) Pack(st *Stream) {
    _eflag := 0
    if len(self.Value) > 150 {
       _eflag = 1
    }
    st.format_ext(_eflag)
    if len(self.Value) < 1 || len(self.Value) > 150 {
        return;
    }
    st.format_olen((len(self.Value))-1, 8)
    st.formatf_PriString(self.Value, len(self.Value))
}
type GNBCUUPUEE1APID struct {
  Value uint64
}
func (self *GNBCUUPUEE1APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967296, 32, 0, 0)
}
func (self * GNBCUUPUEE1APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967296, 32, 0, 0)
}
type GNBCUCPTNLASetupItem struct { // [{'type': 'CP-TNL-Information', 'name': 'tNLAssociationTransportLayerAddress'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-CP-TNLA-Setup-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    TNLAssociationTransportLayerAddress CPTNLInformation
    IEExtensions *GNBCUCPTNLASetupItemExtIEs
}

func (self * GNBCUCPTNLASetupItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.TNLAssociationTransportLayerAddress.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GNBCUCPTNLASetupItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-CP-TNLA-Setup-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GNBCUCPTNLASetupItemExtIEs, order_GNBCUCPTNLASetupItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GNBCUCPTNLASetupItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TNLAssociationTransportLayerAddress.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GNBCUCPTNLASetupItemExtIEs, order_GNBCUCPTNLASetupItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GNBCUCPTNLAFailedToSetupItem struct { // [{'type': 'CP-TNL-Information', 'name': 'tNLAssociationTransportLayerAddress'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-CP-TNLA-Failed-To-Setup-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    TNLAssociationTransportLayerAddress CPTNLInformation
    Cause Cause
    IEExtensions *GNBCUCPTNLAFailedToSetupItemExtIEs
}

func (self * GNBCUCPTNLAFailedToSetupItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.TNLAssociationTransportLayerAddress.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GNBCUCPTNLAFailedToSetupItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-CP-TNLA-Failed-To-Setup-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GNBCUCPTNLAFailedToSetupItemExtIEs, order_GNBCUCPTNLAFailedToSetupItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * GNBCUCPTNLAFailedToSetupItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TNLAssociationTransportLayerAddress.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GNBCUCPTNLAFailedToSetupItemExtIEs, order_GNBCUCPTNLAFailedToSetupItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUCPTNLAToAddItem struct { // [{'type': 'CP-TNL-Information', 'name': 'tNLAssociationTransportLayerAddress'}, {'type': 'TNLAssociationUsage', 'name': 'tNLAssociationUsage'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-CP-TNLA-To-Add-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    TNLAssociationTransportLayerAddress CPTNLInformation
    TNLAssociationUsage TNLAssociationUsage
    IEExtensions *GNBCUCPTNLAToAddItemExtIEs
}

func (self * GNBCUCPTNLAToAddItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.TNLAssociationTransportLayerAddress.Unpack(stream)// p8
    self.TNLAssociationUsage.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GNBCUCPTNLAToAddItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-CP-TNLA-To-Add-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GNBCUCPTNLAToAddItemExtIEs, order_GNBCUCPTNLAToAddItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * GNBCUCPTNLAToAddItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TNLAssociationTransportLayerAddress.Pack(stream)
    self.TNLAssociationUsage.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GNBCUCPTNLAToAddItemExtIEs, order_GNBCUCPTNLAToAddItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUCPTNLAToRemoveItem struct { // [{'type': 'CP-TNL-Information', 'name': 'tNLAssociationTransportLayerAddress'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-CP-TNLA-To-Remove-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    TNLAssociationTransportLayerAddress CPTNLInformation
    IEExtensions *GNBCUCPTNLAToRemoveItemExtIEs
}

func (self * GNBCUCPTNLAToRemoveItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.TNLAssociationTransportLayerAddress.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GNBCUCPTNLAToRemoveItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-CP-TNLA-To-Remove-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GNBCUCPTNLAToRemoveItemExtIEs, order_GNBCUCPTNLAToRemoveItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * GNBCUCPTNLAToRemoveItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TNLAssociationTransportLayerAddress.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GNBCUCPTNLAToRemoveItemExtIEs, order_GNBCUCPTNLAToRemoveItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUCPTNLAToUpdateItem struct { // [{'type': 'CP-TNL-Information', 'name': 'tNLAssociationTransportLayerAddress'}, {'type': 'TNLAssociationUsage', 'name': 'tNLAssociationUsage', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-CP-TNLA-To-Update-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    TNLAssociationTransportLayerAddress CPTNLInformation
    TNLAssociationUsage *TNLAssociationUsage
    IEExtensions *GNBCUCPTNLAToUpdateItemExtIEs
}

func (self * GNBCUCPTNLAToUpdateItem) Unpack(stream *Stream) {
    tNLAssociationUsage_flag := 0x00000001
    iEExtensions_flag := 0x00000002
    _flags := 0
    _flags = stream.get_flags(2)
    self.TNLAssociationTransportLayerAddress.Unpack(stream)// p8
    if (tNLAssociationUsage_flag & _flags) == tNLAssociationUsage_flag { //cond2
        self.TNLAssociationUsage = &TNLAssociationUsage{}//7{'type': 'TNLAssociationUsage', 'name': 'tNLAssociationUsage', 'optional': True}
        self.TNLAssociationUsage.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GNBCUCPTNLAToUpdateItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-CP-TNLA-To-Update-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GNBCUCPTNLAToUpdateItemExtIEs, order_GNBCUCPTNLAToUpdateItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * GNBCUCPTNLAToUpdateItem) Pack(stream *Stream) {
    const tNLAssociationUsage_flag uint = 0x00000001
    const iEExtensions_flag uint = 0x00000002
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TNLAssociationTransportLayerAddress.Pack(stream)
    if self.TNLAssociationUsage != nil { 
        _flags |= tNLAssociationUsage_flag
        self.TNLAssociationUsage.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GNBCUCPTNLAToUpdateItemExtIEs, order_GNBCUCPTNLAToUpdateItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GNBCUUPTNLAToRemoveItem struct { // [{'type': 'CP-TNL-Information', 'name': 'tNLAssociationTransportLayerAddress'}, {'type': 'CP-TNL-Information', 'name': 'tNLAssociationTransportLayerAddressgNBCUCP', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-UP-TNLA-To-Remove-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    TNLAssociationTransportLayerAddress CPTNLInformation
    TNLAssociationTransportLayerAddressgNBCUCP *CPTNLInformation
    IEExtensions *GNBCUUPTNLAToRemoveItemExtIEs
}

func (self * GNBCUUPTNLAToRemoveItem) Unpack(stream *Stream) {
    tNLAssociationTransportLayerAddressgNBCUCP_flag := 0x00000001
    iEExtensions_flag := 0x00000002
    _flags := 0
    _flags = stream.get_flags(2)
    self.TNLAssociationTransportLayerAddress.Unpack(stream)// p8
    if (tNLAssociationTransportLayerAddressgNBCUCP_flag & _flags) == tNLAssociationTransportLayerAddressgNBCUCP_flag { //cond2
        self.TNLAssociationTransportLayerAddressgNBCUCP = &CPTNLInformation{}//7{'type': 'CP-TNL-Information', 'name': 'tNLAssociationTransportLayerAddressgNBCUCP', 'optional': True}
        self.TNLAssociationTransportLayerAddressgNBCUCP.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GNBCUUPTNLAToRemoveItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GNB-CU-UP-TNLA-To-Remove-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GNBCUUPTNLAToRemoveItemExtIEs, order_GNBCUUPTNLAToRemoveItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * GNBCUUPTNLAToRemoveItem) Pack(stream *Stream) {
    const tNLAssociationTransportLayerAddressgNBCUCP_flag uint = 0x00000001
    const iEExtensions_flag uint = 0x00000002
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TNLAssociationTransportLayerAddress.Pack(stream)
    if self.TNLAssociationTransportLayerAddressgNBCUCP != nil { 
        _flags |= tNLAssociationTransportLayerAddressgNBCUCP_flag
        self.TNLAssociationTransportLayerAddressgNBCUCP.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GNBCUUPTNLAToRemoveItemExtIEs, order_GNBCUUPTNLAToRemoveItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GBRQosInformation struct { // [{'type': 'BitRate', 'name': 'e-RAB-MaximumBitrateDL'}, {'type': 'BitRate', 'name': 'e-RAB-MaximumBitrateUL'}, {'type': 'BitRate', 'name': 'e-RAB-GuaranteedBitrateDL'}, {'type': 'BitRate', 'name': 'e-RAB-GuaranteedBitrateUL'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GBR-QosInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    ERABMaximumBitrateDL BitRate
    ERABMaximumBitrateUL BitRate
    ERABGuaranteedBitrateDL BitRate
    ERABGuaranteedBitrateUL BitRate
    IEExtensions *GBRQosInformationExtIEs
}

func (self * GBRQosInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.ERABMaximumBitrateDL.Unpack(stream)// p8
    self.ERABMaximumBitrateUL.Unpack(stream)// p8
    self.ERABGuaranteedBitrateDL.Unpack(stream)// p8
    self.ERABGuaranteedBitrateUL.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GBRQosInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GBR-QosInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GBRQosInformationExtIEs, order_GBRQosInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GBRQosInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.ERABMaximumBitrateDL.Pack(stream)
    self.ERABMaximumBitrateUL.Pack(stream)
    self.ERABGuaranteedBitrateDL.Pack(stream)
    self.ERABGuaranteedBitrateUL.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GBRQosInformationExtIEs, order_GBRQosInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GBRQoSFlowInformation struct { // [{'type': 'BitRate', 'name': 'maxFlowBitRateDownlink'}, {'type': 'BitRate', 'name': 'maxFlowBitRateUplink'}, {'type': 'BitRate', 'name': 'guaranteedFlowBitRateDownlink'}, {'type': 'BitRate', 'name': 'guaranteedFlowBitRateUplink'}, {'type': 'MaxPacketLossRate', 'name': 'maxPacketLossRateDownlink', 'optional': True}, {'type': 'MaxPacketLossRate', 'name': 'maxPacketLossRateUplink', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GBR-QosFlowInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    MaxFlowBitRateDownlink BitRate
    MaxFlowBitRateUplink BitRate
    GuaranteedFlowBitRateDownlink BitRate
    GuaranteedFlowBitRateUplink BitRate
    MaxPacketLossRateDownlink *MaxPacketLossRate
    MaxPacketLossRateUplink *MaxPacketLossRate
    IEExtensions *GBRQosFlowInformationExtIEs
}

func (self * GBRQoSFlowInformation) Unpack(stream *Stream) {
    maxPacketLossRateDownlink_flag := 0x00000002
    maxPacketLossRateUplink_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.MaxFlowBitRateDownlink.Unpack(stream)// p8
    self.MaxFlowBitRateUplink.Unpack(stream)// p8
    self.GuaranteedFlowBitRateDownlink.Unpack(stream)// p8
    self.GuaranteedFlowBitRateUplink.Unpack(stream)// p8
    if (maxPacketLossRateDownlink_flag & _flags) == maxPacketLossRateDownlink_flag { //cond2
        self.MaxPacketLossRateDownlink = &MaxPacketLossRate{}//7{'type': 'MaxPacketLossRate', 'name': 'maxPacketLossRateDownlink', 'optional': True}
        self.MaxPacketLossRateDownlink.Unpack(stream)// p8
    }
    if (maxPacketLossRateUplink_flag & _flags) == maxPacketLossRateUplink_flag { //cond2
        self.MaxPacketLossRateUplink = &MaxPacketLossRate{}//7{'type': 'MaxPacketLossRate', 'name': 'maxPacketLossRateUplink', 'optional': True}
        self.MaxPacketLossRateUplink.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GBRQosFlowInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GBR-QosFlowInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GBRQosFlowInformationExtIEs, order_GBRQosFlowInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GBRQoSFlowInformation) Pack(stream *Stream) {
    const maxPacketLossRateDownlink_flag uint = 0x00000002
    const maxPacketLossRateUplink_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MaxFlowBitRateDownlink.Pack(stream)
    self.MaxFlowBitRateUplink.Pack(stream)
    self.GuaranteedFlowBitRateDownlink.Pack(stream)
    self.GuaranteedFlowBitRateUplink.Pack(stream)
    if self.MaxPacketLossRateDownlink != nil { 
        _flags |= maxPacketLossRateDownlink_flag
        self.MaxPacketLossRateDownlink.Pack(stream)
    }//end of optional
    if self.MaxPacketLossRateUplink != nil { 
        _flags |= maxPacketLossRateUplink_flag
        self.MaxPacketLossRateUplink.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GBRQosFlowInformationExtIEs, order_GBRQosFlowInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type GTPTEID struct {
  Value HexBytes
}
func (self *GTPTEID) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(4)
}
func (self *GTPTEID) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 4)
}
type GTPTunnel struct { // [{'type': 'TransportLayerAddress', 'name': 'transportLayerAddress'}, {'type': 'GTP-TEID', 'name': 'gTP-TEID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GTPTunnel-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    TransportLayerAddress TransportLayerAddress
    GTPTEID GTPTEID
    IEExtensions *GTPTunnelExtIEs
}

func (self * GTPTunnel) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.TransportLayerAddress.Unpack(stream)// p8
    self.GTPTEID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GTPTunnelExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GTPTunnel-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GTPTunnelExtIEs, order_GTPTunnelExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GTPTunnel) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TransportLayerAddress.Pack(stream)
    self.GTPTEID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GTPTunnelExtIEs, order_GTPTunnelExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GNBCUUPOverloadInformation struct {
  Value int
}
const (
    GNBCUUPOverloadInformationoverloaded = 0
    GNBCUUPOverloadInformationnot_overloaded = 1
)
func (self *GNBCUUPOverloadInformation) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 2, 0)
}
func (self *GNBCUUPOverloadInformation) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 2, 0)
}
type GNBDUID struct {
  Value uint64
}
func (self *GNBDUID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(68719476736, 36, 0, 0)
}
func (self * GNBDUID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 68719476736, 36, 0, 0)
}
type HFN struct {
  Value uint64
}
func (self *HFN) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967296, 32, 0, 0)
}
func (self * HFN) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967296, 32, 0, 0)
}
type IntegrityProtectionIndication struct {
  Value int
}
const (
    IntegrityProtectionIndicationrequired = 0
    IntegrityProtectionIndicationpreferred = 1
    IntegrityProtectionIndicationnot_needed = 2

    /* Extensions */
)
func (self *IntegrityProtectionIndication) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *IntegrityProtectionIndication) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type IntegrityProtectionAlgorithm struct {
  Value int
}
const (
    IntegrityProtectionAlgorithmnIA0 = 0
    IntegrityProtectionAlgorithmi_128_NIA1 = 1
    IntegrityProtectionAlgorithmi_128_NIA2 = 2
    IntegrityProtectionAlgorithmi_128_NIA3 = 3

    /* Extensions */
)
func (self *IntegrityProtectionAlgorithm) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 4, 1)
}
func (self *IntegrityProtectionAlgorithm) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 4, 1)
}
type IntegrityProtectionKey struct {
  Value HexBytes
}
func (self *IntegrityProtectionKey) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *IntegrityProtectionKey) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type IntegrityProtectionResult struct {
  Value int
}
const (
    IntegrityProtectionResultperformed = 0
    IntegrityProtectionResultnot_performed = 1

    /* Extensions */
)
func (self *IntegrityProtectionResult) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *IntegrityProtectionResult) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type InactivityTimer struct {
  Value uint64
}
func (self *InactivityTimer) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(7200, 14, 1, 1)
}
func (self * InactivityTimer) Pack(st *Stream){
    st.formatf_Integer(self.Value, 7200, 14, 1, 1)
}
type MaxDataBurstVolume struct {
  Value uint64
}
func (self *MaxDataBurstVolume) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4096, 13, 1, 0)
}
func (self * MaxDataBurstVolume) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4096, 13, 1, 0)
}
type MaximumIPdatarate struct { // [{'type': 'MaxIPrate', 'name': 'maxIPrate'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MaximumIPdatarate-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    MaxIPrate MaxIPrate
    IEExtensions *MaximumIPdatarateExtIEs
}

func (self * MaximumIPdatarate) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.MaxIPrate.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MaximumIPdatarateExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MaximumIPdatarate-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MaximumIPdatarateExtIEs, order_MaximumIPdatarateExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MaximumIPdatarate) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MaxIPrate.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MaximumIPdatarateExtIEs, order_MaximumIPdatarateExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type MaxIPrate struct {
  Value int
}
const (
    MaxIPratebitrate64kbs = 0
    MaxIPratemax_UErate = 1

    /* Extensions */
)
func (self *MaxIPrate) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *MaxIPrate) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type MaxPacketLossRate struct {
  Value uint64
}
func (self *MaxPacketLossRate) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(1001, 11, 1, 0)
}
func (self * MaxPacketLossRate) Pack(st *Stream){
    st.formatf_Integer(self.Value, 1001, 11, 1, 0)
}
type MRDCDataUsageReportItem struct { // [{'type': 'OCTET STRING', 'size': [4], 'name': 'startTimeStamp'}, {'type': 'OCTET STRING', 'size': [4], 'name': 'endTimeStamp'}, {'type': 'INTEGER', 'restricted-to': [(0, 18446744073709551615)], 'name': 'usageCountUL'}, {'type': 'INTEGER', 'restricted-to': [(0, 18446744073709551615)], 'name': 'usageCountDL'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MRDC-Data-Usage-Report-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    StartTimeStamp OCTETSTRING
    EndTimeStamp OCTETSTRING
    UsageCountUL INTEGER
    UsageCountDL INTEGER
    IEExtensions *MRDCDataUsageReportItemExtIEs
}

func (self * MRDCDataUsageReportItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_startTimeStamp = func(st *Stream, self *OCTETSTRING) {
        self.Value = st.parsef_OctString(4)
    }
    Unpack_startTimeStamp(stream, &self.StartTimeStamp)// p2
    var Unpack_endTimeStamp = func(st *Stream, self *OCTETSTRING) {
        self.Value = st.parsef_OctString(4)
    }
    Unpack_endTimeStamp(stream, &self.EndTimeStamp)// p2
    var Unpack_usageCountUL = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(18446744073709551615, 64, 0, 0)
    }
    Unpack_usageCountUL(stream, &self.UsageCountUL)// p2
    var Unpack_usageCountDL = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(18446744073709551615, 64, 0, 0)
    }
    Unpack_usageCountDL(stream, &self.UsageCountDL)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MRDCDataUsageReportItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MRDC-Data-Usage-Report-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MRDCDataUsageReportItemExtIEs, order_MRDCDataUsageReportItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MRDCDataUsageReportItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_startTimeStamp = func(st *Stream, self OCTETSTRING) {
        st.formatf_OctString(self.Value, 4)
    }
    Pack_startTimeStamp(stream, self.StartTimeStamp) //f2
    var Pack_endTimeStamp = func(st *Stream, self OCTETSTRING) {
        st.formatf_OctString(self.Value, 4)
    }
    Pack_endTimeStamp(stream, self.EndTimeStamp) //f2
    var Pack_usageCountUL = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 18446744073709551615, 64, 0, 0)
    }
    Pack_usageCountUL(stream, self.UsageCountUL) //f2
    var Pack_usageCountDL = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 18446744073709551615, 64, 0, 0)
    }
    Pack_usageCountDL(stream, self.UsageCountDL) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MRDCDataUsageReportItemExtIEs, order_MRDCDataUsageReportItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type MRDCUsageInformation struct { // [{'type': 'Data-Usage-per-PDU-Session-Report', 'name': 'data-Usage-per-PDU-Session-Report', 'optional': True}, {'type': 'Data-Usage-per-QoS-Flow-List', 'name': 'data-Usage-per-QoS-Flow-List', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MRDC-Usage-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DataUsageperPDUSessionReport *DataUsageperPDUSessionReport
    DataUsageperQoSFlowList *DataUsageperQoSFlowList
    IEExtensions *MRDCUsageInformationExtIEs
}

func (self * MRDCUsageInformation) Unpack(stream *Stream) {
    dataUsageperPDUSessionReport_flag := 0x00000002
    dataUsageperQoSFlowList_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (dataUsageperPDUSessionReport_flag & _flags) == dataUsageperPDUSessionReport_flag { //cond2
        self.DataUsageperPDUSessionReport = &DataUsageperPDUSessionReport{}//7{'type': 'Data-Usage-per-PDU-Session-Report', 'name': 'data-Usage-per-PDU-Session-Report', 'optional': True}
        self.DataUsageperPDUSessionReport.Unpack(stream)// p8
    }
    if (dataUsageperQoSFlowList_flag & _flags) == dataUsageperQoSFlowList_flag { //cond2
        self.DataUsageperQoSFlowList = &DataUsageperQoSFlowList{}//7{'type': 'Data-Usage-per-QoS-Flow-List', 'name': 'data-Usage-per-QoS-Flow-List', 'optional': True}
        self.DataUsageperQoSFlowList.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MRDCUsageInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MRDC-Usage-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MRDCUsageInformationExtIEs, order_MRDCUsageInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MRDCUsageInformation) Pack(stream *Stream) {
    const dataUsageperPDUSessionReport_flag uint = 0x00000002
    const dataUsageperQoSFlowList_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.DataUsageperPDUSessionReport != nil { 
        _flags |= dataUsageperPDUSessionReport_flag
        self.DataUsageperPDUSessionReport.Pack(stream)
    }//end of optional
    if self.DataUsageperQoSFlowList != nil { 
        _flags |= dataUsageperQoSFlowList_flag
        self.DataUsageperQoSFlowList.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MRDCUsageInformationExtIEs, order_MRDCUsageInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type NetworkInstance struct {
  Value uint64
}
func (self *NetworkInstance) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 9, 1, 1)
}
func (self * NetworkInstance) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 9, 1, 1)
}
type NewULTNLInformationRequired struct {
  Value int
}
const (
    NewULTNLInformationRequiredrequired = 0

    /* Extensions */
)
func (self *NewULTNLInformationRequired) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *NewULTNLInformationRequired) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
type NGRANAllocationAndRetentionPriority struct { // [{'type': 'PriorityLevel', 'name': 'priorityLevel'}, {'type': 'Pre-emptionCapability', 'name': 'pre-emptionCapability'}, {'type': 'Pre-emptionVulnerability', 'name': 'pre-emptionVulnerability'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['NGRANAllocationAndRetentionPriority-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PriorityLevel PriorityLevel
    PreemptionCapability PreemptionCapability
    PreemptionVulnerability PreemptionVulnerability
    IEExtensions *NGRANAllocationAndRetentionPriorityExtIEs
}

func (self * NGRANAllocationAndRetentionPriority) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.PriorityLevel.Unpack(stream)// p8
    self.PreemptionCapability.Unpack(stream)// p8
    self.PreemptionVulnerability.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &NGRANAllocationAndRetentionPriorityExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['NGRANAllocationAndRetentionPriority-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_NGRANAllocationAndRetentionPriorityExtIEs, order_NGRANAllocationAndRetentionPriorityExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * NGRANAllocationAndRetentionPriority) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PriorityLevel.Pack(stream)
    self.PreemptionCapability.Pack(stream)
    self.PreemptionVulnerability.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_NGRANAllocationAndRetentionPriorityExtIEs, order_NGRANAllocationAndRetentionPriorityExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *NGRANQoSSupportList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]NGRANQoSSupportItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *NGRANQoSSupportList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type NGRANQoSSupportList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'NG-RAN-QoS-Support-Item'}, 'size': [(1, 'maxnoofNGRANQOSParameters')]}
    Items []NGRANQoSSupportItem
}

type NGRANQoSSupportItem struct { // [{'type': 'Non-Dynamic5QIDescriptor', 'name': 'non-Dynamic5QIDescriptor'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['NG-RAN-QoS-Support-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    NonDynamic5QIDescriptor NonDynamic5QIDescriptor
    IEExtensions *NGRANQoSSupportItemExtIEs
}

func (self * NGRANQoSSupportItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.NonDynamic5QIDescriptor.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &NGRANQoSSupportItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['NG-RAN-QoS-Support-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_NGRANQoSSupportItemExtIEs, order_NGRANQoSSupportItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * NGRANQoSSupportItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.NonDynamic5QIDescriptor.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_NGRANQoSSupportItemExtIEs, order_NGRANQoSSupportItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type NonDynamic5QIDescriptor struct { // [{'type': 'INTEGER', 'restricted-to': [(0, 255), None], 'name': 'fiveQI'}, {'type': 'QoSPriorityLevel', 'name': 'qoSPriorityLevel', 'optional': True}, {'type': 'AveragingWindow', 'name': 'averagingWindow', 'optional': True}, {'type': 'MaxDataBurstVolume', 'name': 'maxDataBurstVolume', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Non-Dynamic5QIDescriptor-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    FiveQI INTEGER
    QoSPriorityLevel *QoSPriorityLevel
    AveragingWindow *AveragingWindow
    MaxDataBurstVolume *MaxDataBurstVolume
    IEExtensions *NonDynamic5QIDescriptorExtIEs
}

func (self * NonDynamic5QIDescriptor) Unpack(stream *Stream) {
    qoSPriorityLevel_flag := 0x00000001
    averagingWindow_flag := 0x00000002
    maxDataBurstVolume_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    _flags := 0
    _flags = stream.get_flags(4)
    var Unpack_fiveQI = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(256, 9, 1, 0)
    }
    Unpack_fiveQI(stream, &self.FiveQI)// p2
    if (qoSPriorityLevel_flag & _flags) == qoSPriorityLevel_flag { //cond2
        self.QoSPriorityLevel = &QoSPriorityLevel{}//7{'type': 'QoSPriorityLevel', 'name': 'qoSPriorityLevel', 'optional': True}
        self.QoSPriorityLevel.Unpack(stream)// p8
    }
    if (averagingWindow_flag & _flags) == averagingWindow_flag { //cond2
        self.AveragingWindow = &AveragingWindow{}//7{'type': 'AveragingWindow', 'name': 'averagingWindow', 'optional': True}
        self.AveragingWindow.Unpack(stream)// p8
    }
    if (maxDataBurstVolume_flag & _flags) == maxDataBurstVolume_flag { //cond2
        self.MaxDataBurstVolume = &MaxDataBurstVolume{}//7{'type': 'MaxDataBurstVolume', 'name': 'maxDataBurstVolume', 'optional': True}
        self.MaxDataBurstVolume.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &NonDynamic5QIDescriptorExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Non-Dynamic5QIDescriptor-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_NonDynamic5QIDescriptorExtIEs, order_NonDynamic5QIDescriptorExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * NonDynamic5QIDescriptor) Pack(stream *Stream) {
    const qoSPriorityLevel_flag uint = 0x00000001
    const averagingWindow_flag uint = 0x00000002
    const maxDataBurstVolume_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_fiveQI = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 256, 9, 1, 0)
    }
    Pack_fiveQI(stream, self.FiveQI) //f2
    if self.QoSPriorityLevel != nil { 
        _flags |= qoSPriorityLevel_flag
        self.QoSPriorityLevel.Pack(stream)
    }//end of optional
    if self.AveragingWindow != nil { 
        _flags |= averagingWindow_flag
        self.AveragingWindow.Pack(stream)
    }//end of optional
    if self.MaxDataBurstVolume != nil { 
        _flags |= maxDataBurstVolume_flag
        self.MaxDataBurstVolume.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_NonDynamic5QIDescriptorExtIEs, order_NonDynamic5QIDescriptorExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type NRCellIdentity struct {
  Len int
  Value HexBytes
}
func (self *NRCellIdentity) Unpack(st *Stream){
    self.Value = st.parsef_BitString(36, 36)
}
func (self *NRCellIdentity) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 36)
}
type NRCGI struct { // [{'type': 'PLMN-Identity', 'name': 'pLMN-Identity'}, {'type': 'NR-Cell-Identity', 'name': 'nR-Cell-Identity'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['NR-CGI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PLMNIdentity PLMNIdentity
    NRCellIdentity NRCellIdentity
    IEExtensions *NRCGIExtIEs
}

func (self * NRCGI) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.PLMNIdentity.Unpack(stream)// p8
    self.NRCellIdentity.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &NRCGIExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['NR-CGI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_NRCGIExtIEs, order_NRCGIExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * NRCGI) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.NRCellIdentity.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_NRCGIExtIEs, order_NRCGIExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *NRCGISupportList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(512)
    _size += 1
    self.Items = make([]NRCGISupportItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *NRCGISupportList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 512)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type NRCGISupportList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'NR-CGI-Support-Item'}, 'size': [(1, 'maxnoofNRCGI')]}
    Items []NRCGISupportItem
}

type NRCGISupportItem struct { // [{'type': 'NR-CGI', 'name': 'nR-CGI'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['NR-CGI-Support-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    NRCGI NRCGI
    IEExtensions *NRCGISupportItemExtIEs
}

func (self * NRCGISupportItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.NRCGI.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &NRCGISupportItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['NR-CGI-Support-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_NRCGISupportItemExtIEs, order_NRCGISupportItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * NRCGISupportItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.NRCGI.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_NRCGISupportItemExtIEs, order_NRCGISupportItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type OutOfOrderDelivery struct {
  Value int
}
const (
    OutOfOrderDeliverytRue = 0

    /* Extensions */
)
func (self *OutOfOrderDelivery) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *OutOfOrderDelivery) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
type PacketDelayBudget struct {
  Value uint64
}
func (self *PacketDelayBudget) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(1024, 11, 1, 0)
}
func (self * PacketDelayBudget) Pack(st *Stream){
    st.formatf_Integer(self.Value, 1024, 11, 1, 0)
}
type PacketErrorRate struct { // [{'type': 'PER-Scalar', 'name': 'pER-Scalar'}, {'type': 'PER-Exponent', 'name': 'pER-Exponent'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PacketErrorRate-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PERScalar PERScalar
    PERExponent PERExponent
    IEExtensions *PacketErrorRateExtIEs
}

func (self * PacketErrorRate) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PERScalar.Unpack(stream)// p8
    self.PERExponent.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PacketErrorRateExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PacketErrorRate-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PacketErrorRateExtIEs, order_PacketErrorRateExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PacketErrorRate) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PERScalar.Pack(stream)
    self.PERExponent.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PacketErrorRateExtIEs, order_PacketErrorRateExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type PERScalar struct {
  Value uint64
}
func (self *PERScalar) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(10, 5, 1, 0)
}
func (self * PERScalar) Pack(st *Stream){
    st.formatf_Integer(self.Value, 10, 5, 1, 0)
}
type PERExponent struct {
  Value uint64
}
func (self *PERExponent) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(10, 5, 1, 0)
}
func (self * PERExponent) Pack(st *Stream){
    st.formatf_Integer(self.Value, 10, 5, 1, 0)
}
type PDCPConfiguration struct { // [{'type': 'PDCP-SN-Size', 'name': 'pDCP-SN-Size-UL'}, {'type': 'PDCP-SN-Size', 'name': 'pDCP-SN-Size-DL'}, {'type': 'RLC-Mode', 'name': 'rLC-Mode'}, {'type': 'ROHC-Parameters', 'name': 'rOHC-Parameters', 'optional': True}, {'type': 'T-ReorderingTimer', 'name': 't-ReorderingTimer', 'optional': True}, {'type': 'DiscardTimer', 'name': 'discardTimer', 'optional': True}, {'type': 'ULDataSplitThreshold', 'name': 'uLDataSplitThreshold', 'optional': True}, {'type': 'PDCP-Duplication', 'name': 'pDCP-Duplication', 'optional': True}, {'type': 'PDCP-Reestablishment', 'name': 'pDCP-Reestablishment', 'optional': True}, {'type': 'PDCP-DataRecovery', 'name': 'pDCP-DataRecovery', 'optional': True}, {'type': 'Duplication-Activation', 'name': 'duplication-Activation', 'optional': True}, {'type': 'OutOfOrderDelivery', 'name': 'outOfOrderDelivery', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDCP-Configuration-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDCPSNSizeUL PDCPSNSize
    PDCPSNSizeDL PDCPSNSize
    RLCMode RLCMode
    ROHCParameters *ROHCParameters
    TReorderingTimer *TReorderingTimer
    DiscardTimer *DiscardTimer
    ULDataSplitThreshold *ULDataSplitThreshold
    PDCPDuplication *PDCPDuplication
    PDCPReestablishment *PDCPReestablishment
    PDCPDataRecovery *PDCPDataRecovery
    DuplicationActivation *DuplicationActivation
    OutOfOrderDelivery *OutOfOrderDelivery
    IEExtensions *PDCPConfigurationExtIEs
}

func (self * PDCPConfiguration) Unpack(stream *Stream) {
    rOHCParameters_flag := 0x00000002
    tReorderingTimer_flag := 0x00000004
    discardTimer_flag := 0x00000008
    uLDataSplitThreshold_flag := 0x00000010
    pDCPDuplication_flag := 0x00000020
    pDCPReestablishment_flag := 0x00000040
    pDCPDataRecovery_flag := 0x00000080
    duplicationActivation_flag := 0x00000100
    outOfOrderDelivery_flag := 0x00000200
    iEExtensions_flag := 0x00000400
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(11)
    self.PDCPSNSizeUL.Unpack(stream)// p8
    self.PDCPSNSizeDL.Unpack(stream)// p8
    self.RLCMode.Unpack(stream)// p8
    if (rOHCParameters_flag & _flags) == rOHCParameters_flag { //cond2
        self.ROHCParameters = &ROHCParameters{}//7{'type': 'ROHC-Parameters', 'name': 'rOHC-Parameters', 'optional': True}
        self.ROHCParameters.Unpack(stream)// p8
    }
    if (tReorderingTimer_flag & _flags) == tReorderingTimer_flag { //cond2
        self.TReorderingTimer = &TReorderingTimer{}//7{'type': 'T-ReorderingTimer', 'name': 't-ReorderingTimer', 'optional': True}
        self.TReorderingTimer.Unpack(stream)// p8
    }
    if (discardTimer_flag & _flags) == discardTimer_flag { //cond2
        self.DiscardTimer = &DiscardTimer{}//7{'type': 'DiscardTimer', 'name': 'discardTimer', 'optional': True}
        self.DiscardTimer.Unpack(stream)// p8
    }
    if (uLDataSplitThreshold_flag & _flags) == uLDataSplitThreshold_flag { //cond2
        self.ULDataSplitThreshold = &ULDataSplitThreshold{}//7{'type': 'ULDataSplitThreshold', 'name': 'uLDataSplitThreshold', 'optional': True}
        self.ULDataSplitThreshold.Unpack(stream)// p8
    }
    if (pDCPDuplication_flag & _flags) == pDCPDuplication_flag { //cond2
        self.PDCPDuplication = &PDCPDuplication{}//7{'type': 'PDCP-Duplication', 'name': 'pDCP-Duplication', 'optional': True}
        self.PDCPDuplication.Unpack(stream)// p8
    }
    if (pDCPReestablishment_flag & _flags) == pDCPReestablishment_flag { //cond2
        self.PDCPReestablishment = &PDCPReestablishment{}//7{'type': 'PDCP-Reestablishment', 'name': 'pDCP-Reestablishment', 'optional': True}
        self.PDCPReestablishment.Unpack(stream)// p8
    }
    if (pDCPDataRecovery_flag & _flags) == pDCPDataRecovery_flag { //cond2
        self.PDCPDataRecovery = &PDCPDataRecovery{}//7{'type': 'PDCP-DataRecovery', 'name': 'pDCP-DataRecovery', 'optional': True}
        self.PDCPDataRecovery.Unpack(stream)// p8
    }
    if (duplicationActivation_flag & _flags) == duplicationActivation_flag { //cond2
        self.DuplicationActivation = &DuplicationActivation{}//7{'type': 'Duplication-Activation', 'name': 'duplication-Activation', 'optional': True}
        self.DuplicationActivation.Unpack(stream)// p8
    }
    if (outOfOrderDelivery_flag & _flags) == outOfOrderDelivery_flag { //cond2
        self.OutOfOrderDelivery = &OutOfOrderDelivery{}//7{'type': 'OutOfOrderDelivery', 'name': 'outOfOrderDelivery', 'optional': True}
        self.OutOfOrderDelivery.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDCPConfigurationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDCP-Configuration-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDCPConfigurationExtIEs, order_PDCPConfigurationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDCPConfiguration) Pack(stream *Stream) {
    const rOHCParameters_flag uint = 0x00000002
    const tReorderingTimer_flag uint = 0x00000004
    const discardTimer_flag uint = 0x00000008
    const uLDataSplitThreshold_flag uint = 0x00000010
    const pDCPDuplication_flag uint = 0x00000020
    const pDCPReestablishment_flag uint = 0x00000040
    const pDCPDataRecovery_flag uint = 0x00000080
    const duplicationActivation_flag uint = 0x00000100
    const outOfOrderDelivery_flag uint = 0x00000200
    const iEExtensions_flag uint = 0x00000400
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(11)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDCPSNSizeUL.Pack(stream)
    self.PDCPSNSizeDL.Pack(stream)
    self.RLCMode.Pack(stream)
    if self.ROHCParameters != nil { 
        _flags |= rOHCParameters_flag
        self.ROHCParameters.Pack(stream)
    }//end of optional
    if self.TReorderingTimer != nil { 
        _flags |= tReorderingTimer_flag
        self.TReorderingTimer.Pack(stream)
    }//end of optional
    if self.DiscardTimer != nil { 
        _flags |= discardTimer_flag
        self.DiscardTimer.Pack(stream)
    }//end of optional
    if self.ULDataSplitThreshold != nil { 
        _flags |= uLDataSplitThreshold_flag
        self.ULDataSplitThreshold.Pack(stream)
    }//end of optional
    if self.PDCPDuplication != nil { 
        _flags |= pDCPDuplication_flag
        self.PDCPDuplication.Pack(stream)
    }//end of optional
    if self.PDCPReestablishment != nil { 
        _flags |= pDCPReestablishment_flag
        self.PDCPReestablishment.Pack(stream)
    }//end of optional
    if self.PDCPDataRecovery != nil { 
        _flags |= pDCPDataRecovery_flag
        self.PDCPDataRecovery.Pack(stream)
    }//end of optional
    if self.DuplicationActivation != nil { 
        _flags |= duplicationActivation_flag
        self.DuplicationActivation.Pack(stream)
    }//end of optional
    if self.OutOfOrderDelivery != nil { 
        _flags |= outOfOrderDelivery_flag
        self.OutOfOrderDelivery.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDCPConfigurationExtIEs, order_PDCPConfigurationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 11)
}//end

type PDCPCount struct { // [{'type': 'PDCP-SN', 'name': 'pDCP-SN'}, {'type': 'HFN', 'name': 'hFN'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDCP-Count-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDCPSN PDCPSN
    HFN HFN
    IEExtensions *PDCPCountExtIEs
}

func (self * PDCPCount) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PDCPSN.Unpack(stream)// p8
    self.HFN.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDCPCountExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDCP-Count-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDCPCountExtIEs, order_PDCPCountExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDCPCount) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDCPSN.Pack(stream)
    self.HFN.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDCPCountExtIEs, order_PDCPCountExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type PDCPSNStatusRequest struct {
  Value int
}
const (
    PDCPSNStatusRequestrequested = 0

    /* Extensions */
)
func (self *PDCPSNStatusRequest) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *PDCPSNStatusRequest) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
type PDCPDataRecovery struct {
  Value int
}
const (
    PDCPDataRecoverytRue = 0

    /* Extensions */
)
func (self *PDCPDataRecovery) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *PDCPDataRecovery) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
type PDCPDuplication struct {
  Value int
}
const (
    PDCPDuplicationtRue = 0

    /* Extensions */
)
func (self *PDCPDuplication) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *PDCPDuplication) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
type PDCPReestablishment struct {
  Value int
}
const (
    PDCPReestablishmenttRue = 0

    /* Extensions */
)
func (self *PDCPReestablishment) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *PDCPReestablishment) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
func (self *PDUSessionResourceDataUsageList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceDataUsageItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceDataUsageList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceDataUsageList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-Data-Usage-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceDataUsageItem
}

type PDUSessionResourceDataUsageItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'MRDC-Usage-Information', 'name': 'mRDC-Usage-Information'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Data-Usage-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    MRDCUsageInformation MRDCUsageInformation
    IEExtensions *PDUSessionResourceDataUsageItemExtIEs
}

func (self * PDUSessionResourceDataUsageItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PDUSessionID.Unpack(stream)// p8
    self.MRDCUsageInformation.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceDataUsageItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Data-Usage-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceDataUsageItemExtIEs, order_PDUSessionResourceDataUsageItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceDataUsageItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    self.MRDCUsageInformation.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceDataUsageItemExtIEs, order_PDUSessionResourceDataUsageItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type PDCPSN struct {
  Value uint64
}
func (self *PDCPSN) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(262144, 18, 0, 0)
}
func (self * PDCPSN) Pack(st *Stream){
    st.formatf_Integer(self.Value, 262144, 18, 0, 0)
}
type PDCPSNSize struct {
  Value int
}
const (
    PDCPSNSizes_12 = 0
    PDCPSNSizes_18 = 1

    /* Extensions */
)
func (self *PDCPSNSize) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *PDCPSNSize) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type PDCPSNStatusInformation struct { // [{'type': 'DRBBStatusTransfer', 'name': 'pdcpStatusTransfer-UL'}, {'type': 'PDCP-Count', 'name': 'pdcpStatusTransfer-DL'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRBsSubjectToStatusTransfer-Item-ExtIEs'], 'name': 'iE-Extension', 'optional': True}, None]
    PdcpStatusTransferUL DRBBStatusTransfer
    PdcpStatusTransferDL PDCPCount
    IEExtension *DRBsSubjectToStatusTransferItemExtIEs
}

func (self * PDCPSNStatusInformation) Unpack(stream *Stream) {
    iEExtension_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PdcpStatusTransferUL.Unpack(stream)// p8
    self.PdcpStatusTransferDL.Unpack(stream)// p8
    if (iEExtension_flag & _flags) == iEExtension_flag { //cond2
        self.IEExtension = &DRBsSubjectToStatusTransferItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRBsSubjectToStatusTransfer-Item-ExtIEs'], 'name': 'iE-Extension', 'optional': True}
        IEExtension := ProtocolExtensionContainer {table_DRBsSubjectToStatusTransferItemExtIEs, order_DRBsSubjectToStatusTransferItemExtIEs} // p3
        IEExtension.Unpack(stream, &self.IEExtension) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDCPSNStatusInformation) Pack(stream *Stream) {
    const iEExtension_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PdcpStatusTransferUL.Pack(stream)
    self.PdcpStatusTransferDL.Pack(stream)
    if self.IEExtension != nil { 
        _flags |= iEExtension_flag
        IEExtension := &ProtocolExtensionContainer {table_DRBsSubjectToStatusTransferItemExtIEs, order_DRBsSubjectToStatusTransferItemExtIEs} // p3
        IEExtension.Pack(stream, &self.IEExtension)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type DRBBStatusTransfer struct { // [{'type': 'BIT STRING', 'size': [(1, 131072)], 'name': 'receiveStatusofPDCPSDU', 'optional': True}, {'type': 'PDCP-Count', 'name': 'countValue'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRBBStatusTransfer-ExtIEs'], 'name': 'iE-Extension', 'optional': True}, None]
    ReceiveStatusofPDCPSDU *BITSTRING
    CountValue PDCPCount
    IEExtension *DRBBStatusTransferExtIEs
}

func (self * DRBBStatusTransfer) Unpack(stream *Stream) {
    receiveStatusofPDCPSDU_flag := 0x00000002
    iEExtension_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    if (receiveStatusofPDCPSDU_flag & _flags) == receiveStatusofPDCPSDU_flag { //cond1
        var Unpack_receiveStatusofPDCPSDU = func(st *Stream, self *BITSTRING){
            self.Len = int(st.parse_blen(18, 0)+1)
            self.Value = st.parsef_BitString(131072, int(self.Len))
        }
        self.ReceiveStatusofPDCPSDU = &BITSTRING{}//6{'type': 'BIT STRING', 'size': [(1, 131072)], 'name': 'receiveStatusofPDCPSDU', 'optional': True}
        Unpack_receiveStatusofPDCPSDU(stream, self.ReceiveStatusofPDCPSDU)// p1 {'type': 'BIT STRING', 'size': [(1, 131072)], 'name': 'receiveStatusofPDCPSDU', 'optional': True}
    }
    self.CountValue.Unpack(stream)// p8
    if (iEExtension_flag & _flags) == iEExtension_flag { //cond2
        self.IEExtension = &DRBBStatusTransferExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DRBBStatusTransfer-ExtIEs'], 'name': 'iE-Extension', 'optional': True}
        IEExtension := ProtocolExtensionContainer {table_DRBBStatusTransferExtIEs, order_DRBBStatusTransferExtIEs} // p3
        IEExtension.Unpack(stream, &self.IEExtension) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBBStatusTransfer) Pack(stream *Stream) {
    const receiveStatusofPDCPSDU_flag uint = 0x00000002
    const iEExtension_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.ReceiveStatusofPDCPSDU != nil { //YY
        _flags |= receiveStatusofPDCPSDU_flag
        var Pack_receiveStatusofPDCPSDU = func(st *Stream, self BITSTRING) {
            st.format_blen(int(self.Len-1), 18, 0)
            st.formatf_BitString(self.Value, int(self.Len))
        }
        Pack_receiveStatusofPDCPSDU(stream, *self.ReceiveStatusofPDCPSDU) //f1
    }//end of optional
    self.CountValue.Pack(stream)
    if self.IEExtension != nil { 
        _flags |= iEExtension_flag
        IEExtension := &ProtocolExtensionContainer {table_DRBBStatusTransferExtIEs, order_DRBBStatusTransferExtIEs} // p3
        IEExtension.Pack(stream, &self.IEExtension)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type PDUSessionID struct {
  Value uint64
}
func (self *PDUSessionID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * PDUSessionID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
type PDUSessionResourceActivity struct {
  Value int
}
const (
    PDUSessionResourceActivityactive = 0
    PDUSessionResourceActivitynot_active = 1

    /* Extensions */
)
func (self *PDUSessionResourceActivity) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *PDUSessionResourceActivity) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
func (self *PDUSessionResourceActivityList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceActivityItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceActivityList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceActivityList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-Activity-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceActivityItem
}

type PDUSessionResourceActivityItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'PDU-Session-Resource-Activity', 'name': 'pDU-Session-Resource-Activity'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Activity-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    PDUSessionResourceActivity PDUSessionResourceActivity
    IEExtensions *PDUSessionResourceActivityItemExtIEs
}

func (self * PDUSessionResourceActivityItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PDUSessionID.Unpack(stream)// p8
    self.PDUSessionResourceActivity.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceActivityItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Activity-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceActivityItemExtIEs, order_PDUSessionResourceActivityItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceActivityItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    self.PDUSessionResourceActivity.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceActivityItemExtIEs, order_PDUSessionResourceActivityItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *PDUSessionResourceConfirmModifiedList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceConfirmModifiedItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceConfirmModifiedList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceConfirmModifiedList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-Confirm-Modified-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceConfirmModifiedItem
}

type PDUSessionResourceConfirmModifiedItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'DRB-Confirm-Modified-List-NG-RAN', 'name': 'dRB-Confirm-Modified-List-NG-RAN', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Confirm-Modified-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    DRBConfirmModifiedListNGRAN *DRBConfirmModifiedListNGRAN
    IEExtensions *PDUSessionResourceConfirmModifiedItemExtIEs
}

func (self * PDUSessionResourceConfirmModifiedItem) Unpack(stream *Stream) {
    dRBConfirmModifiedListNGRAN_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.PDUSessionID.Unpack(stream)// p8
    if (dRBConfirmModifiedListNGRAN_flag & _flags) == dRBConfirmModifiedListNGRAN_flag { //cond2
        self.DRBConfirmModifiedListNGRAN = &DRBConfirmModifiedListNGRAN{}//7{'type': 'DRB-Confirm-Modified-List-NG-RAN', 'name': 'dRB-Confirm-Modified-List-NG-RAN', 'optional': True}
        self.DRBConfirmModifiedListNGRAN.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceConfirmModifiedItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Confirm-Modified-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceConfirmModifiedItemExtIEs, order_PDUSessionResourceConfirmModifiedItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceConfirmModifiedItem) Pack(stream *Stream) {
    const dRBConfirmModifiedListNGRAN_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    if self.DRBConfirmModifiedListNGRAN != nil { 
        _flags |= dRBConfirmModifiedListNGRAN_flag
        self.DRBConfirmModifiedListNGRAN.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceConfirmModifiedItemExtIEs, order_PDUSessionResourceConfirmModifiedItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

func (self *PDUSessionResourceFailedList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceFailedItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceFailedList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceFailedList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-Failed-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceFailedItem
}

type PDUSessionResourceFailedItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Failed-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    Cause Cause
    IEExtensions *PDUSessionResourceFailedItemExtIEs
}

func (self * PDUSessionResourceFailedItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PDUSessionID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceFailedItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Failed-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceFailedItemExtIEs, order_PDUSessionResourceFailedItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceFailedItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceFailedItemExtIEs, order_PDUSessionResourceFailedItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *PDUSessionResourceFailedModList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceFailedModItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceFailedModList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceFailedModList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-Failed-Mod-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceFailedModItem
}

type PDUSessionResourceFailedModItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Failed-Mod-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    Cause Cause
    IEExtensions *PDUSessionResourceFailedModItemExtIEs
}

func (self * PDUSessionResourceFailedModItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PDUSessionID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceFailedModItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Failed-Mod-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceFailedModItemExtIEs, order_PDUSessionResourceFailedModItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceFailedModItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceFailedModItemExtIEs, order_PDUSessionResourceFailedModItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *PDUSessionResourceFailedToModifyList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceFailedToModifyItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceFailedToModifyList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceFailedToModifyList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-Failed-To-Modify-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceFailedToModifyItem
}

type PDUSessionResourceFailedToModifyItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Failed-To-Modify-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    Cause Cause
    IEExtensions *PDUSessionResourceFailedToModifyItemExtIEs
}

func (self * PDUSessionResourceFailedToModifyItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PDUSessionID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceFailedToModifyItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Failed-To-Modify-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceFailedToModifyItemExtIEs, order_PDUSessionResourceFailedToModifyItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceFailedToModifyItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceFailedToModifyItemExtIEs, order_PDUSessionResourceFailedToModifyItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *PDUSessionResourceModifiedList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceModifiedItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceModifiedList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceModifiedList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-Modified-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceModifiedItem
}

type PDUSessionResourceModifiedItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'UP-TNL-Information', 'name': 'nG-DL-UP-TNL-Information', 'optional': True}, {'type': 'SecurityResult', 'name': 'securityResult', 'optional': True}, {'type': 'Data-Forwarding-Information', 'name': 'pDU-Session-Data-Forwarding-Information-Response', 'optional': True}, {'type': 'DRB-Setup-List-NG-RAN', 'name': 'dRB-Setup-List-NG-RAN', 'optional': True}, {'type': 'DRB-Failed-List-NG-RAN', 'name': 'dRB-Failed-List-NG-RAN', 'optional': True}, {'type': 'DRB-Modified-List-NG-RAN', 'name': 'dRB-Modified-List-NG-RAN', 'optional': True}, {'type': 'DRB-Failed-To-Modify-List-NG-RAN', 'name': 'dRB-Failed-To-Modify-List-NG-RAN', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Modified-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    NGDLUPTNLInformation *UPTNLInformation
    SecurityResult *SecurityResult
    PDUSessionDataForwardingInformationResponse *DataForwardingInformation
    DRBSetupListNGRAN *DRBSetupListNGRAN
    DRBFailedListNGRAN *DRBFailedListNGRAN
    DRBModifiedListNGRAN *DRBModifiedListNGRAN
    DRBFailedToModifyListNGRAN *DRBFailedToModifyListNGRAN
    IEExtensions *PDUSessionResourceModifiedItemExtIEs
}

func (self * PDUSessionResourceModifiedItem) Unpack(stream *Stream) {
    nGDLUPTNLInformation_flag := 0x00000002
    securityResult_flag := 0x00000004
    pDUSessionDataForwardingInformationResponse_flag := 0x00000008
    dRBSetupListNGRAN_flag := 0x00000010
    dRBFailedListNGRAN_flag := 0x00000020
    dRBModifiedListNGRAN_flag := 0x00000040
    dRBFailedToModifyListNGRAN_flag := 0x00000080
    iEExtensions_flag := 0x00000100
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(9)
    self.PDUSessionID.Unpack(stream)// p8
    if (nGDLUPTNLInformation_flag & _flags) == nGDLUPTNLInformation_flag { //cond2
        self.NGDLUPTNLInformation = &UPTNLInformation{}//7{'type': 'UP-TNL-Information', 'name': 'nG-DL-UP-TNL-Information', 'optional': True}
        self.NGDLUPTNLInformation.Unpack(stream)// p8
    }
    if (securityResult_flag & _flags) == securityResult_flag { //cond2
        self.SecurityResult = &SecurityResult{}//7{'type': 'SecurityResult', 'name': 'securityResult', 'optional': True}
        self.SecurityResult.Unpack(stream)// p8
    }
    if (pDUSessionDataForwardingInformationResponse_flag & _flags) == pDUSessionDataForwardingInformationResponse_flag { //cond2
        self.PDUSessionDataForwardingInformationResponse = &DataForwardingInformation{}//7{'type': 'Data-Forwarding-Information', 'name': 'pDU-Session-Data-Forwarding-Information-Response', 'optional': True}
        self.PDUSessionDataForwardingInformationResponse.Unpack(stream)// p8
    }
    if (dRBSetupListNGRAN_flag & _flags) == dRBSetupListNGRAN_flag { //cond2
        self.DRBSetupListNGRAN = &DRBSetupListNGRAN{}//7{'type': 'DRB-Setup-List-NG-RAN', 'name': 'dRB-Setup-List-NG-RAN', 'optional': True}
        self.DRBSetupListNGRAN.Unpack(stream)// p8
    }
    if (dRBFailedListNGRAN_flag & _flags) == dRBFailedListNGRAN_flag { //cond2
        self.DRBFailedListNGRAN = &DRBFailedListNGRAN{}//7{'type': 'DRB-Failed-List-NG-RAN', 'name': 'dRB-Failed-List-NG-RAN', 'optional': True}
        self.DRBFailedListNGRAN.Unpack(stream)// p8
    }
    if (dRBModifiedListNGRAN_flag & _flags) == dRBModifiedListNGRAN_flag { //cond2
        self.DRBModifiedListNGRAN = &DRBModifiedListNGRAN{}//7{'type': 'DRB-Modified-List-NG-RAN', 'name': 'dRB-Modified-List-NG-RAN', 'optional': True}
        self.DRBModifiedListNGRAN.Unpack(stream)// p8
    }
    if (dRBFailedToModifyListNGRAN_flag & _flags) == dRBFailedToModifyListNGRAN_flag { //cond2
        self.DRBFailedToModifyListNGRAN = &DRBFailedToModifyListNGRAN{}//7{'type': 'DRB-Failed-To-Modify-List-NG-RAN', 'name': 'dRB-Failed-To-Modify-List-NG-RAN', 'optional': True}
        self.DRBFailedToModifyListNGRAN.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceModifiedItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Modified-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceModifiedItemExtIEs, order_PDUSessionResourceModifiedItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceModifiedItem) Pack(stream *Stream) {
    const nGDLUPTNLInformation_flag uint = 0x00000002
    const securityResult_flag uint = 0x00000004
    const pDUSessionDataForwardingInformationResponse_flag uint = 0x00000008
    const dRBSetupListNGRAN_flag uint = 0x00000010
    const dRBFailedListNGRAN_flag uint = 0x00000020
    const dRBModifiedListNGRAN_flag uint = 0x00000040
    const dRBFailedToModifyListNGRAN_flag uint = 0x00000080
    const iEExtensions_flag uint = 0x00000100
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(9)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    if self.NGDLUPTNLInformation != nil { 
        _flags |= nGDLUPTNLInformation_flag
        self.NGDLUPTNLInformation.Pack(stream)
    }//end of optional
    if self.SecurityResult != nil { 
        _flags |= securityResult_flag
        self.SecurityResult.Pack(stream)
    }//end of optional
    if self.PDUSessionDataForwardingInformationResponse != nil { 
        _flags |= pDUSessionDataForwardingInformationResponse_flag
        self.PDUSessionDataForwardingInformationResponse.Pack(stream)
    }//end of optional
    if self.DRBSetupListNGRAN != nil { 
        _flags |= dRBSetupListNGRAN_flag
        self.DRBSetupListNGRAN.Pack(stream)
    }//end of optional
    if self.DRBFailedListNGRAN != nil { 
        _flags |= dRBFailedListNGRAN_flag
        self.DRBFailedListNGRAN.Pack(stream)
    }//end of optional
    if self.DRBModifiedListNGRAN != nil { 
        _flags |= dRBModifiedListNGRAN_flag
        self.DRBModifiedListNGRAN.Pack(stream)
    }//end of optional
    if self.DRBFailedToModifyListNGRAN != nil { 
        _flags |= dRBFailedToModifyListNGRAN_flag
        self.DRBFailedToModifyListNGRAN.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceModifiedItemExtIEs, order_PDUSessionResourceModifiedItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 9)
}//end

func (self *PDUSessionResourceRequiredToModifyList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceRequiredToModifyItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceRequiredToModifyList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceRequiredToModifyList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-Required-To-Modify-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceRequiredToModifyItem
}

type PDUSessionResourceRequiredToModifyItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'UP-TNL-Information', 'name': 'nG-DL-UP-TNL-Information', 'optional': True}, {'type': 'DRB-Required-To-Modify-List-NG-RAN', 'name': 'dRB-Required-To-Modify-List-NG-RAN', 'optional': True}, {'type': 'DRB-Required-To-Remove-List-NG-RAN', 'name': 'dRB-Required-To-Remove-List-NG-RAN', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Required-To-Modify-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    NGDLUPTNLInformation *UPTNLInformation
    DRBRequiredToModifyListNGRAN *DRBRequiredToModifyListNGRAN
    DRBRequiredToRemoveListNGRAN *DRBRequiredToRemoveListNGRAN
    IEExtensions *PDUSessionResourceRequiredToModifyItemExtIEs
}

func (self * PDUSessionResourceRequiredToModifyItem) Unpack(stream *Stream) {
    nGDLUPTNLInformation_flag := 0x00000002
    dRBRequiredToModifyListNGRAN_flag := 0x00000004
    dRBRequiredToRemoveListNGRAN_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.PDUSessionID.Unpack(stream)// p8
    if (nGDLUPTNLInformation_flag & _flags) == nGDLUPTNLInformation_flag { //cond2
        self.NGDLUPTNLInformation = &UPTNLInformation{}//7{'type': 'UP-TNL-Information', 'name': 'nG-DL-UP-TNL-Information', 'optional': True}
        self.NGDLUPTNLInformation.Unpack(stream)// p8
    }
    if (dRBRequiredToModifyListNGRAN_flag & _flags) == dRBRequiredToModifyListNGRAN_flag { //cond2
        self.DRBRequiredToModifyListNGRAN = &DRBRequiredToModifyListNGRAN{}//7{'type': 'DRB-Required-To-Modify-List-NG-RAN', 'name': 'dRB-Required-To-Modify-List-NG-RAN', 'optional': True}
        self.DRBRequiredToModifyListNGRAN.Unpack(stream)// p8
    }
    if (dRBRequiredToRemoveListNGRAN_flag & _flags) == dRBRequiredToRemoveListNGRAN_flag { //cond2
        self.DRBRequiredToRemoveListNGRAN = &DRBRequiredToRemoveListNGRAN{}//7{'type': 'DRB-Required-To-Remove-List-NG-RAN', 'name': 'dRB-Required-To-Remove-List-NG-RAN', 'optional': True}
        self.DRBRequiredToRemoveListNGRAN.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceRequiredToModifyItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Required-To-Modify-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceRequiredToModifyItemExtIEs, order_PDUSessionResourceRequiredToModifyItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceRequiredToModifyItem) Pack(stream *Stream) {
    const nGDLUPTNLInformation_flag uint = 0x00000002
    const dRBRequiredToModifyListNGRAN_flag uint = 0x00000004
    const dRBRequiredToRemoveListNGRAN_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    if self.NGDLUPTNLInformation != nil { 
        _flags |= nGDLUPTNLInformation_flag
        self.NGDLUPTNLInformation.Pack(stream)
    }//end of optional
    if self.DRBRequiredToModifyListNGRAN != nil { 
        _flags |= dRBRequiredToModifyListNGRAN_flag
        self.DRBRequiredToModifyListNGRAN.Pack(stream)
    }//end of optional
    if self.DRBRequiredToRemoveListNGRAN != nil { 
        _flags |= dRBRequiredToRemoveListNGRAN_flag
        self.DRBRequiredToRemoveListNGRAN.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceRequiredToModifyItemExtIEs, order_PDUSessionResourceRequiredToModifyItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *PDUSessionResourceSetupList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceSetupItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceSetupList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceSetupList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-Setup-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceSetupItem
}

type PDUSessionResourceSetupItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'SecurityResult', 'name': 'securityResult', 'optional': True}, {'type': 'UP-TNL-Information', 'name': 'nG-DL-UP-TNL-Information'}, {'type': 'Data-Forwarding-Information', 'name': 'pDU-Session-Data-Forwarding-Information-Response', 'optional': True}, {'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 'nG-DL-UP-Unchanged', 'optional': True}, {'type': 'DRB-Setup-List-NG-RAN', 'name': 'dRB-Setup-List-NG-RAN'}, {'type': 'DRB-Failed-List-NG-RAN', 'name': 'dRB-Failed-List-NG-RAN', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Setup-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    SecurityResult *SecurityResult
    NGDLUPTNLInformation UPTNLInformation
    PDUSessionDataForwardingInformationResponse *DataForwardingInformation
    NGDLUPUnchanged *ENUMERATED
    DRBSetupListNGRAN DRBSetupListNGRAN
    DRBFailedListNGRAN *DRBFailedListNGRAN
    IEExtensions *PDUSessionResourceSetupItemExtIEs
}

func (self * PDUSessionResourceSetupItem) Unpack(stream *Stream) {
    securityResult_flag := 0x00000002
    pDUSessionDataForwardingInformationResponse_flag := 0x00000004
    nGDLUPUnchanged_flag := 0x00000008
    dRBFailedListNGRAN_flag := 0x00000010
    iEExtensions_flag := 0x00000020
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(6)
    self.PDUSessionID.Unpack(stream)// p8
    if (securityResult_flag & _flags) == securityResult_flag { //cond2
        self.SecurityResult = &SecurityResult{}//7{'type': 'SecurityResult', 'name': 'securityResult', 'optional': True}
        self.SecurityResult.Unpack(stream)// p8
    }
    self.NGDLUPTNLInformation.Unpack(stream)// p8
    if (pDUSessionDataForwardingInformationResponse_flag & _flags) == pDUSessionDataForwardingInformationResponse_flag { //cond2
        self.PDUSessionDataForwardingInformationResponse = &DataForwardingInformation{}//7{'type': 'Data-Forwarding-Information', 'name': 'pDU-Session-Data-Forwarding-Information-Response', 'optional': True}
        self.PDUSessionDataForwardingInformationResponse.Unpack(stream)// p8
    }
    if (nGDLUPUnchanged_flag & _flags) == nGDLUPUnchanged_flag { //cond1
        var Unpack_nGDLUPUnchanged = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.NGDLUPUnchanged = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 'nG-DL-UP-Unchanged', 'optional': True}
        Unpack_nGDLUPUnchanged(stream, self.NGDLUPUnchanged)// p1 {'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 'nG-DL-UP-Unchanged', 'optional': True}
    }
    self.DRBSetupListNGRAN.Unpack(stream)// p8
    if (dRBFailedListNGRAN_flag & _flags) == dRBFailedListNGRAN_flag { //cond2
        self.DRBFailedListNGRAN = &DRBFailedListNGRAN{}//7{'type': 'DRB-Failed-List-NG-RAN', 'name': 'dRB-Failed-List-NG-RAN', 'optional': True}
        self.DRBFailedListNGRAN.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceSetupItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Setup-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceSetupItemExtIEs, order_PDUSessionResourceSetupItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceSetupItem) Pack(stream *Stream) {
    const securityResult_flag uint = 0x00000002
    const pDUSessionDataForwardingInformationResponse_flag uint = 0x00000004
    const nGDLUPUnchanged_flag uint = 0x00000008
    const dRBFailedListNGRAN_flag uint = 0x00000010
    const iEExtensions_flag uint = 0x00000020
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(6)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    if self.SecurityResult != nil { 
        _flags |= securityResult_flag
        self.SecurityResult.Pack(stream)
    }//end of optional
    self.NGDLUPTNLInformation.Pack(stream)
    if self.PDUSessionDataForwardingInformationResponse != nil { 
        _flags |= pDUSessionDataForwardingInformationResponse_flag
        self.PDUSessionDataForwardingInformationResponse.Pack(stream)
    }//end of optional
    if self.NGDLUPUnchanged != nil { //YY
        _flags |= nGDLUPUnchanged_flag
        var Pack_nGDLUPUnchanged = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_nGDLUPUnchanged(stream, *self.NGDLUPUnchanged) //f1
    }//end of optional
    self.DRBSetupListNGRAN.Pack(stream)
    if self.DRBFailedListNGRAN != nil { 
        _flags |= dRBFailedListNGRAN_flag
        self.DRBFailedListNGRAN.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceSetupItemExtIEs, order_PDUSessionResourceSetupItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

func (self *PDUSessionResourceSetupModList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceSetupModItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceSetupModList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceSetupModList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-Setup-Mod-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceSetupModItem
}

type PDUSessionResourceSetupModItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'SecurityResult', 'name': 'securityResult', 'optional': True}, {'type': 'UP-TNL-Information', 'name': 'nG-DL-UP-TNL-Information'}, {'type': 'Data-Forwarding-Information', 'name': 'pDU-Session-Data-Forwarding-Information-Response', 'optional': True}, {'type': 'DRB-Setup-Mod-List-NG-RAN', 'name': 'dRB-Setup-Mod-List-NG-RAN'}, {'type': 'DRB-Failed-Mod-List-NG-RAN', 'name': 'dRB-Failed-Mod-List-NG-RAN', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Setup-Mod-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    SecurityResult *SecurityResult
    NGDLUPTNLInformation UPTNLInformation
    PDUSessionDataForwardingInformationResponse *DataForwardingInformation
    DRBSetupModListNGRAN DRBSetupModListNGRAN
    DRBFailedModListNGRAN *DRBFailedModListNGRAN
    IEExtensions *PDUSessionResourceSetupModItemExtIEs
}

func (self * PDUSessionResourceSetupModItem) Unpack(stream *Stream) {
    securityResult_flag := 0x00000002
    pDUSessionDataForwardingInformationResponse_flag := 0x00000004
    dRBFailedModListNGRAN_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.PDUSessionID.Unpack(stream)// p8
    if (securityResult_flag & _flags) == securityResult_flag { //cond2
        self.SecurityResult = &SecurityResult{}//7{'type': 'SecurityResult', 'name': 'securityResult', 'optional': True}
        self.SecurityResult.Unpack(stream)// p8
    }
    self.NGDLUPTNLInformation.Unpack(stream)// p8
    if (pDUSessionDataForwardingInformationResponse_flag & _flags) == pDUSessionDataForwardingInformationResponse_flag { //cond2
        self.PDUSessionDataForwardingInformationResponse = &DataForwardingInformation{}//7{'type': 'Data-Forwarding-Information', 'name': 'pDU-Session-Data-Forwarding-Information-Response', 'optional': True}
        self.PDUSessionDataForwardingInformationResponse.Unpack(stream)// p8
    }
    self.DRBSetupModListNGRAN.Unpack(stream)// p8
    if (dRBFailedModListNGRAN_flag & _flags) == dRBFailedModListNGRAN_flag { //cond2
        self.DRBFailedModListNGRAN = &DRBFailedModListNGRAN{}//7{'type': 'DRB-Failed-Mod-List-NG-RAN', 'name': 'dRB-Failed-Mod-List-NG-RAN', 'optional': True}
        self.DRBFailedModListNGRAN.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceSetupModItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-Setup-Mod-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceSetupModItemExtIEs, order_PDUSessionResourceSetupModItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceSetupModItem) Pack(stream *Stream) {
    const securityResult_flag uint = 0x00000002
    const pDUSessionDataForwardingInformationResponse_flag uint = 0x00000004
    const dRBFailedModListNGRAN_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    if self.SecurityResult != nil { 
        _flags |= securityResult_flag
        self.SecurityResult.Pack(stream)
    }//end of optional
    self.NGDLUPTNLInformation.Pack(stream)
    if self.PDUSessionDataForwardingInformationResponse != nil { 
        _flags |= pDUSessionDataForwardingInformationResponse_flag
        self.PDUSessionDataForwardingInformationResponse.Pack(stream)
    }//end of optional
    self.DRBSetupModListNGRAN.Pack(stream)
    if self.DRBFailedModListNGRAN != nil { 
        _flags |= dRBFailedModListNGRAN_flag
        self.DRBFailedModListNGRAN.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceSetupModItemExtIEs, order_PDUSessionResourceSetupModItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *PDUSessionResourceToModifyList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceToModifyItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceToModifyList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceToModifyList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-To-Modify-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceToModifyItem
}

type PDUSessionResourceToModifyItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'SecurityIndication', 'name': 'securityIndication', 'optional': True}, {'type': 'BitRate', 'name': 'pDU-Session-Resource-DL-AMBR', 'optional': True}, {'type': 'UP-TNL-Information', 'name': 'nG-UL-UP-TNL-Information', 'optional': True}, {'type': 'Data-Forwarding-Information-Request', 'name': 'pDU-Session-Data-Forwarding-Information-Request', 'optional': True}, {'type': 'Data-Forwarding-Information', 'name': 'pDU-Session-Data-Forwarding-Information', 'optional': True}, {'type': 'Inactivity-Timer', 'name': 'pDU-Session-Inactivity-Timer', 'optional': True}, {'type': 'NetworkInstance', 'name': 'networkInstance', 'optional': True}, {'type': 'DRB-To-Setup-List-NG-RAN', 'name': 'dRB-To-Setup-List-NG-RAN', 'optional': True}, {'type': 'DRB-To-Modify-List-NG-RAN', 'name': 'dRB-To-Modify-List-NG-RAN', 'optional': True}, {'type': 'DRB-To-Remove-List-NG-RAN', 'name': 'dRB-To-Remove-List-NG-RAN', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-To-Modify-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    SecurityIndication *SecurityIndication
    PDUSessionResourceDLAMBR *BitRate
    NGULUPTNLInformation *UPTNLInformation
    PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest
    PDUSessionDataForwardingInformation *DataForwardingInformation
    PDUSessionInactivityTimer *InactivityTimer
    NetworkInstance *NetworkInstance
    DRBToSetupListNGRAN *DRBToSetupListNGRAN
    DRBToModifyListNGRAN *DRBToModifyListNGRAN
    DRBToRemoveListNGRAN *DRBToRemoveListNGRAN
    IEExtensions *PDUSessionResourceToModifyItemExtIEs
}

func (self * PDUSessionResourceToModifyItem) Unpack(stream *Stream) {
    securityIndication_flag := 0x00000002
    pDUSessionResourceDLAMBR_flag := 0x00000004
    nGULUPTNLInformation_flag := 0x00000008
    pDUSessionDataForwardingInformationRequest_flag := 0x00000010
    pDUSessionDataForwardingInformation_flag := 0x00000020
    pDUSessionInactivityTimer_flag := 0x00000040
    networkInstance_flag := 0x00000080
    dRBToSetupListNGRAN_flag := 0x00000100
    dRBToModifyListNGRAN_flag := 0x00000200
    dRBToRemoveListNGRAN_flag := 0x00000400
    iEExtensions_flag := 0x00000800
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(12)
    self.PDUSessionID.Unpack(stream)// p8
    if (securityIndication_flag & _flags) == securityIndication_flag { //cond2
        self.SecurityIndication = &SecurityIndication{}//7{'type': 'SecurityIndication', 'name': 'securityIndication', 'optional': True}
        self.SecurityIndication.Unpack(stream)// p8
    }
    if (pDUSessionResourceDLAMBR_flag & _flags) == pDUSessionResourceDLAMBR_flag { //cond2
        self.PDUSessionResourceDLAMBR = &BitRate{}//7{'type': 'BitRate', 'name': 'pDU-Session-Resource-DL-AMBR', 'optional': True}
        self.PDUSessionResourceDLAMBR.Unpack(stream)// p8
    }
    if (nGULUPTNLInformation_flag & _flags) == nGULUPTNLInformation_flag { //cond2
        self.NGULUPTNLInformation = &UPTNLInformation{}//7{'type': 'UP-TNL-Information', 'name': 'nG-UL-UP-TNL-Information', 'optional': True}
        self.NGULUPTNLInformation.Unpack(stream)// p8
    }
    if (pDUSessionDataForwardingInformationRequest_flag & _flags) == pDUSessionDataForwardingInformationRequest_flag { //cond2
        self.PDUSessionDataForwardingInformationRequest = &DataForwardingInformationRequest{}//7{'type': 'Data-Forwarding-Information-Request', 'name': 'pDU-Session-Data-Forwarding-Information-Request', 'optional': True}
        self.PDUSessionDataForwardingInformationRequest.Unpack(stream)// p8
    }
    if (pDUSessionDataForwardingInformation_flag & _flags) == pDUSessionDataForwardingInformation_flag { //cond2
        self.PDUSessionDataForwardingInformation = &DataForwardingInformation{}//7{'type': 'Data-Forwarding-Information', 'name': 'pDU-Session-Data-Forwarding-Information', 'optional': True}
        self.PDUSessionDataForwardingInformation.Unpack(stream)// p8
    }
    if (pDUSessionInactivityTimer_flag & _flags) == pDUSessionInactivityTimer_flag { //cond2
        self.PDUSessionInactivityTimer = &InactivityTimer{}//7{'type': 'Inactivity-Timer', 'name': 'pDU-Session-Inactivity-Timer', 'optional': True}
        self.PDUSessionInactivityTimer.Unpack(stream)// p8
    }
    if (networkInstance_flag & _flags) == networkInstance_flag { //cond2
        self.NetworkInstance = &NetworkInstance{}//7{'type': 'NetworkInstance', 'name': 'networkInstance', 'optional': True}
        self.NetworkInstance.Unpack(stream)// p8
    }
    if (dRBToSetupListNGRAN_flag & _flags) == dRBToSetupListNGRAN_flag { //cond2
        self.DRBToSetupListNGRAN = &DRBToSetupListNGRAN{}//7{'type': 'DRB-To-Setup-List-NG-RAN', 'name': 'dRB-To-Setup-List-NG-RAN', 'optional': True}
        self.DRBToSetupListNGRAN.Unpack(stream)// p8
    }
    if (dRBToModifyListNGRAN_flag & _flags) == dRBToModifyListNGRAN_flag { //cond2
        self.DRBToModifyListNGRAN = &DRBToModifyListNGRAN{}//7{'type': 'DRB-To-Modify-List-NG-RAN', 'name': 'dRB-To-Modify-List-NG-RAN', 'optional': True}
        self.DRBToModifyListNGRAN.Unpack(stream)// p8
    }
    if (dRBToRemoveListNGRAN_flag & _flags) == dRBToRemoveListNGRAN_flag { //cond2
        self.DRBToRemoveListNGRAN = &DRBToRemoveListNGRAN{}//7{'type': 'DRB-To-Remove-List-NG-RAN', 'name': 'dRB-To-Remove-List-NG-RAN', 'optional': True}
        self.DRBToRemoveListNGRAN.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceToModifyItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-To-Modify-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceToModifyItemExtIEs, order_PDUSessionResourceToModifyItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceToModifyItem) Pack(stream *Stream) {
    const securityIndication_flag uint = 0x00000002
    const pDUSessionResourceDLAMBR_flag uint = 0x00000004
    const nGULUPTNLInformation_flag uint = 0x00000008
    const pDUSessionDataForwardingInformationRequest_flag uint = 0x00000010
    const pDUSessionDataForwardingInformation_flag uint = 0x00000020
    const pDUSessionInactivityTimer_flag uint = 0x00000040
    const networkInstance_flag uint = 0x00000080
    const dRBToSetupListNGRAN_flag uint = 0x00000100
    const dRBToModifyListNGRAN_flag uint = 0x00000200
    const dRBToRemoveListNGRAN_flag uint = 0x00000400
    const iEExtensions_flag uint = 0x00000800
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(12)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    if self.SecurityIndication != nil { 
        _flags |= securityIndication_flag
        self.SecurityIndication.Pack(stream)
    }//end of optional
    if self.PDUSessionResourceDLAMBR != nil { 
        _flags |= pDUSessionResourceDLAMBR_flag
        self.PDUSessionResourceDLAMBR.Pack(stream)
    }//end of optional
    if self.NGULUPTNLInformation != nil { 
        _flags |= nGULUPTNLInformation_flag
        self.NGULUPTNLInformation.Pack(stream)
    }//end of optional
    if self.PDUSessionDataForwardingInformationRequest != nil { 
        _flags |= pDUSessionDataForwardingInformationRequest_flag
        self.PDUSessionDataForwardingInformationRequest.Pack(stream)
    }//end of optional
    if self.PDUSessionDataForwardingInformation != nil { 
        _flags |= pDUSessionDataForwardingInformation_flag
        self.PDUSessionDataForwardingInformation.Pack(stream)
    }//end of optional
    if self.PDUSessionInactivityTimer != nil { 
        _flags |= pDUSessionInactivityTimer_flag
        self.PDUSessionInactivityTimer.Pack(stream)
    }//end of optional
    if self.NetworkInstance != nil { 
        _flags |= networkInstance_flag
        self.NetworkInstance.Pack(stream)
    }//end of optional
    if self.DRBToSetupListNGRAN != nil { 
        _flags |= dRBToSetupListNGRAN_flag
        self.DRBToSetupListNGRAN.Pack(stream)
    }//end of optional
    if self.DRBToModifyListNGRAN != nil { 
        _flags |= dRBToModifyListNGRAN_flag
        self.DRBToModifyListNGRAN.Pack(stream)
    }//end of optional
    if self.DRBToRemoveListNGRAN != nil { 
        _flags |= dRBToRemoveListNGRAN_flag
        self.DRBToRemoveListNGRAN.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceToModifyItemExtIEs, order_PDUSessionResourceToModifyItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 12)
}//end

func (self *PDUSessionResourceToRemoveList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceToRemoveItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceToRemoveList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceToRemoveList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-To-Remove-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceToRemoveItem
}

type PDUSessionResourceToRemoveItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-To-Remove-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    IEExtensions *PDUSessionResourceToRemoveItemExtIEs
}

func (self * PDUSessionResourceToRemoveItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PDUSessionID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceToRemoveItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-To-Remove-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceToRemoveItemExtIEs, order_PDUSessionResourceToRemoveItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceToRemoveItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceToRemoveItemExtIEs, order_PDUSessionResourceToRemoveItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *PDUSessionResourceToSetupList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceToSetupItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceToSetupList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceToSetupList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-To-Setup-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceToSetupItem
}

type PDUSessionResourceToSetupItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'PDU-Session-Type', 'name': 'pDU-Session-Type'}, {'type': 'SNSSAI', 'name': 'sNSSAI'}, {'type': 'SecurityIndication', 'name': 'securityIndication'}, {'type': 'BitRate', 'name': 'pDU-Session-Resource-DL-AMBR', 'optional': True}, {'type': 'UP-TNL-Information', 'name': 'nG-UL-UP-TNL-Information'}, {'type': 'Data-Forwarding-Information-Request', 'name': 'pDU-Session-Data-Forwarding-Information-Request', 'optional': True}, {'type': 'Inactivity-Timer', 'name': 'pDU-Session-Inactivity-Timer', 'optional': True}, {'type': 'UP-TNL-Information', 'name': 'existing-Allocated-NG-DL-UP-TNL-Info', 'optional': True}, {'type': 'NetworkInstance', 'name': 'networkInstance', 'optional': True}, {'type': 'DRB-To-Setup-List-NG-RAN', 'name': 'dRB-To-Setup-List-NG-RAN'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-To-Setup-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    PDUSessionType PDUSessionType
    SNSSAI SNSSAI
    SecurityIndication SecurityIndication
    PDUSessionResourceDLAMBR *BitRate
    NGULUPTNLInformation UPTNLInformation
    PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest
    PDUSessionInactivityTimer *InactivityTimer
    ExistingAllocatedNGDLUPTNLInfo *UPTNLInformation
    NetworkInstance *NetworkInstance
    DRBToSetupListNGRAN DRBToSetupListNGRAN
    IEExtensions *PDUSessionResourceToSetupItemExtIEs
}

func (self * PDUSessionResourceToSetupItem) Unpack(stream *Stream) {
    pDUSessionResourceDLAMBR_flag := 0x00000002
    pDUSessionDataForwardingInformationRequest_flag := 0x00000004
    pDUSessionInactivityTimer_flag := 0x00000008
    existingAllocatedNGDLUPTNLInfo_flag := 0x00000010
    networkInstance_flag := 0x00000020
    iEExtensions_flag := 0x00000040
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(7)
    self.PDUSessionID.Unpack(stream)// p8
    self.PDUSessionType.Unpack(stream)// p8
    self.SNSSAI.Unpack(stream)// p8
    self.SecurityIndication.Unpack(stream)// p8
    if (pDUSessionResourceDLAMBR_flag & _flags) == pDUSessionResourceDLAMBR_flag { //cond2
        self.PDUSessionResourceDLAMBR = &BitRate{}//7{'type': 'BitRate', 'name': 'pDU-Session-Resource-DL-AMBR', 'optional': True}
        self.PDUSessionResourceDLAMBR.Unpack(stream)// p8
    }
    self.NGULUPTNLInformation.Unpack(stream)// p8
    if (pDUSessionDataForwardingInformationRequest_flag & _flags) == pDUSessionDataForwardingInformationRequest_flag { //cond2
        self.PDUSessionDataForwardingInformationRequest = &DataForwardingInformationRequest{}//7{'type': 'Data-Forwarding-Information-Request', 'name': 'pDU-Session-Data-Forwarding-Information-Request', 'optional': True}
        self.PDUSessionDataForwardingInformationRequest.Unpack(stream)// p8
    }
    if (pDUSessionInactivityTimer_flag & _flags) == pDUSessionInactivityTimer_flag { //cond2
        self.PDUSessionInactivityTimer = &InactivityTimer{}//7{'type': 'Inactivity-Timer', 'name': 'pDU-Session-Inactivity-Timer', 'optional': True}
        self.PDUSessionInactivityTimer.Unpack(stream)// p8
    }
    if (existingAllocatedNGDLUPTNLInfo_flag & _flags) == existingAllocatedNGDLUPTNLInfo_flag { //cond2
        self.ExistingAllocatedNGDLUPTNLInfo = &UPTNLInformation{}//7{'type': 'UP-TNL-Information', 'name': 'existing-Allocated-NG-DL-UP-TNL-Info', 'optional': True}
        self.ExistingAllocatedNGDLUPTNLInfo.Unpack(stream)// p8
    }
    if (networkInstance_flag & _flags) == networkInstance_flag { //cond2
        self.NetworkInstance = &NetworkInstance{}//7{'type': 'NetworkInstance', 'name': 'networkInstance', 'optional': True}
        self.NetworkInstance.Unpack(stream)// p8
    }
    self.DRBToSetupListNGRAN.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceToSetupItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-To-Setup-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceToSetupItemExtIEs, order_PDUSessionResourceToSetupItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceToSetupItem) Pack(stream *Stream) {
    const pDUSessionResourceDLAMBR_flag uint = 0x00000002
    const pDUSessionDataForwardingInformationRequest_flag uint = 0x00000004
    const pDUSessionInactivityTimer_flag uint = 0x00000008
    const existingAllocatedNGDLUPTNLInfo_flag uint = 0x00000010
    const networkInstance_flag uint = 0x00000020
    const iEExtensions_flag uint = 0x00000040
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(7)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    self.PDUSessionType.Pack(stream)
    self.SNSSAI.Pack(stream)
    self.SecurityIndication.Pack(stream)
    if self.PDUSessionResourceDLAMBR != nil { 
        _flags |= pDUSessionResourceDLAMBR_flag
        self.PDUSessionResourceDLAMBR.Pack(stream)
    }//end of optional
    self.NGULUPTNLInformation.Pack(stream)
    if self.PDUSessionDataForwardingInformationRequest != nil { 
        _flags |= pDUSessionDataForwardingInformationRequest_flag
        self.PDUSessionDataForwardingInformationRequest.Pack(stream)
    }//end of optional
    if self.PDUSessionInactivityTimer != nil { 
        _flags |= pDUSessionInactivityTimer_flag
        self.PDUSessionInactivityTimer.Pack(stream)
    }//end of optional
    if self.ExistingAllocatedNGDLUPTNLInfo != nil { 
        _flags |= existingAllocatedNGDLUPTNLInfo_flag
        self.ExistingAllocatedNGDLUPTNLInfo.Pack(stream)
    }//end of optional
    if self.NetworkInstance != nil { 
        _flags |= networkInstance_flag
        self.NetworkInstance.Pack(stream)
    }//end of optional
    self.DRBToSetupListNGRAN.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceToSetupItemExtIEs, order_PDUSessionResourceToSetupItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 7)
}//end

func (self *PDUSessionResourceToSetupModList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionResourceToSetupModItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionResourceToSetupModList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionResourceToSetupModList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-Resource-To-Setup-Mod-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionResourceToSetupModItem
}

type PDUSessionResourceToSetupModItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'PDU-Session-Type', 'name': 'pDU-Session-Type'}, {'type': 'SNSSAI', 'name': 'sNSSAI'}, {'type': 'SecurityIndication', 'name': 'securityIndication'}, {'type': 'BitRate', 'name': 'pDU-Session-Resource-AMBR', 'optional': True}, {'type': 'UP-TNL-Information', 'name': 'nG-UL-UP-TNL-Information'}, {'type': 'Data-Forwarding-Information-Request', 'name': 'pDU-Session-Data-Forwarding-Information-Request', 'optional': True}, {'type': 'Inactivity-Timer', 'name': 'pDU-Session-Inactivity-Timer', 'optional': True}, {'type': 'DRB-To-Setup-Mod-List-NG-RAN', 'name': 'dRB-To-Setup-Mod-List-NG-RAN'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-To-Setup-Mod-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    PDUSessionType PDUSessionType
    SNSSAI SNSSAI
    SecurityIndication SecurityIndication
    PDUSessionResourceAMBR *BitRate
    NGULUPTNLInformation UPTNLInformation
    PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest
    PDUSessionInactivityTimer *InactivityTimer
    DRBToSetupModListNGRAN DRBToSetupModListNGRAN
    IEExtensions *PDUSessionResourceToSetupModItemExtIEs
}

func (self * PDUSessionResourceToSetupModItem) Unpack(stream *Stream) {
    pDUSessionResourceAMBR_flag := 0x00000002
    pDUSessionDataForwardingInformationRequest_flag := 0x00000004
    pDUSessionInactivityTimer_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.PDUSessionID.Unpack(stream)// p8
    self.PDUSessionType.Unpack(stream)// p8
    self.SNSSAI.Unpack(stream)// p8
    self.SecurityIndication.Unpack(stream)// p8
    if (pDUSessionResourceAMBR_flag & _flags) == pDUSessionResourceAMBR_flag { //cond2
        self.PDUSessionResourceAMBR = &BitRate{}//7{'type': 'BitRate', 'name': 'pDU-Session-Resource-AMBR', 'optional': True}
        self.PDUSessionResourceAMBR.Unpack(stream)// p8
    }
    self.NGULUPTNLInformation.Unpack(stream)// p8
    if (pDUSessionDataForwardingInformationRequest_flag & _flags) == pDUSessionDataForwardingInformationRequest_flag { //cond2
        self.PDUSessionDataForwardingInformationRequest = &DataForwardingInformationRequest{}//7{'type': 'Data-Forwarding-Information-Request', 'name': 'pDU-Session-Data-Forwarding-Information-Request', 'optional': True}
        self.PDUSessionDataForwardingInformationRequest.Unpack(stream)// p8
    }
    if (pDUSessionInactivityTimer_flag & _flags) == pDUSessionInactivityTimer_flag { //cond2
        self.PDUSessionInactivityTimer = &InactivityTimer{}//7{'type': 'Inactivity-Timer', 'name': 'pDU-Session-Inactivity-Timer', 'optional': True}
        self.PDUSessionInactivityTimer.Unpack(stream)// p8
    }
    self.DRBToSetupModListNGRAN.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionResourceToSetupModItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-Resource-To-Setup-Mod-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionResourceToSetupModItemExtIEs, order_PDUSessionResourceToSetupModItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionResourceToSetupModItem) Pack(stream *Stream) {
    const pDUSessionResourceAMBR_flag uint = 0x00000002
    const pDUSessionDataForwardingInformationRequest_flag uint = 0x00000004
    const pDUSessionInactivityTimer_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    self.PDUSessionType.Pack(stream)
    self.SNSSAI.Pack(stream)
    self.SecurityIndication.Pack(stream)
    if self.PDUSessionResourceAMBR != nil { 
        _flags |= pDUSessionResourceAMBR_flag
        self.PDUSessionResourceAMBR.Pack(stream)
    }//end of optional
    self.NGULUPTNLInformation.Pack(stream)
    if self.PDUSessionDataForwardingInformationRequest != nil { 
        _flags |= pDUSessionDataForwardingInformationRequest_flag
        self.PDUSessionDataForwardingInformationRequest.Pack(stream)
    }//end of optional
    if self.PDUSessionInactivityTimer != nil { 
        _flags |= pDUSessionInactivityTimer_flag
        self.PDUSessionInactivityTimer.Pack(stream)
    }//end of optional
    self.DRBToSetupModListNGRAN.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionResourceToSetupModItemExtIEs, order_PDUSessionResourceToSetupModItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *PDUSessionToNotifyList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]PDUSessionToNotifyItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDUSessionToNotifyList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDUSessionToNotifyList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDU-Session-To-Notify-Item'}, 'size': [(1, 'maxnoofPDUSessionResource')]}
    Items []PDUSessionToNotifyItem
}

type PDUSessionToNotifyItem struct { // [{'type': 'PDU-Session-ID', 'name': 'pDU-Session-ID'}, {'type': 'QoS-Flow-List', 'name': 'qoS-Flow-List'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-To-Notify-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDUSessionID PDUSessionID
    QoSFlowList QoSFlowList
    IEExtensions *PDUSessionToNotifyItemExtIEs
}

func (self * PDUSessionToNotifyItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PDUSessionID.Unpack(stream)// p8
    self.QoSFlowList.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PDUSessionToNotifyItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PDU-Session-To-Notify-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PDUSessionToNotifyItemExtIEs, order_PDUSessionToNotifyItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PDUSessionToNotifyItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PDUSessionID.Pack(stream)
    self.QoSFlowList.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PDUSessionToNotifyItemExtIEs, order_PDUSessionToNotifyItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type PDUSessionType struct {
  Value int
}
const (
    PDUSessionTypeipv4 = 0
    PDUSessionTypeipv6 = 1
    PDUSessionTypeipv4v6 = 2
    PDUSessionTypeethernet = 3
    PDUSessionTypeunstructured = 4

    /* Extensions */
)
func (self *PDUSessionType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 5, 1)
}
func (self *PDUSessionType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 5, 1)
}
type PLMNIdentity struct {
  Value HexBytes
}
func (self *PLMNIdentity) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(3)
}
func (self *PLMNIdentity) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 3)
}
type PortNumber struct {
  Len int
  Value HexBytes
}
func (self *PortNumber) Unpack(st *Stream){
    self.Value = st.parsef_BitString(16, 16)
}
func (self *PortNumber) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 16)
}
type PPI struct {
  Value uint64
}
func (self *PPI) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(8, 4, 1, 0)
}
func (self * PPI) Pack(st *Stream){
    st.formatf_Integer(self.Value, 8, 4, 1, 0)
}
type PriorityLevel struct {
  Value uint64
}
func (self *PriorityLevel) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16, 4, 0, 0)
}
func (self * PriorityLevel) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16, 4, 0, 0)
}
type PreemptionCapability struct {
  Value int
}
const (
    PreemptionCapabilityshall_not_trigger_pre_emption = 0
    PreemptionCapabilitymay_trigger_pre_emption = 1
)
func (self *PreemptionCapability) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 2, 0)
}
func (self *PreemptionCapability) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 2, 0)
}
type PreemptionVulnerability struct {
  Value int
}
const (
    PreemptionVulnerabilitynot_pre_emptable = 0
    PreemptionVulnerabilitypre_emptable = 1
)
func (self *PreemptionVulnerability) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 2, 0)
}
func (self *PreemptionVulnerability) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 2, 0)
}
type QCI struct {
  Value uint64
}
func (self *QCI) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * QCI) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
func (self *QoSCharacteristics)Unpack(stream *Stream) {
    //coptions := []string{"non-Dynamic-5QI","dynamic-5QI","choice-extension","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        self.NonDynamic5QI = &NonDynamic5QIDescriptor{}//cho6
        self.NonDynamic5QI.Unpack(stream)
    } else if choice == 1 { //ch2
        self.Dynamic5QI = &Dynamic5QIDescriptor{}//cho6
        self.Dynamic5QI.Unpack(stream)
    } else if choice == 2 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &QoSCharacteristicsExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * QoSCharacteristics) Pack(stream *Stream) {
    if self.NonDynamic5QI != nil {
        stream.set_choice(0, 2, 0, 3)
        self.NonDynamic5QI.Pack(stream)//2
    } else if self.Dynamic5QI != nil {
        stream.set_choice(1, 2, 0, 3)
        self.Dynamic5QI.Pack(stream)//2
    } else if self.Choiceextension != nil {
        stream.set_choice(2, 2, 0, 3)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type QoSCharacteristics struct { //[{'type': 'Non-Dynamic5QIDescriptor', 'name': 'non-Dynamic-5QI'}, {'type': 'Dynamic5QIDescriptor', 'name': 'dynamic-5QI'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['QoS-Characteristics-ExtIEs'], 'name': 'choice-extension'}]
    NonDynamic5QI *NonDynamic5QIDescriptor
    Dynamic5QI *Dynamic5QIDescriptor
    Choiceextension *QoSCharacteristicsExtIEs
} // QoSCharacteristics

type QoSFlowIdentifier struct {
  Value uint64
}
func (self *QoSFlowIdentifier) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(64, 6, 0, 0)
}
func (self * QoSFlowIdentifier) Pack(st *Stream){
    st.formatf_Integer(self.Value, 64, 6, 0, 0)
}
func (self *QoSFlowList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(64)
    _size += 1
    self.Items = make([]QoSFlowItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *QoSFlowList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 64)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type QoSFlowList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'QoS-Flow-Item'}, 'size': [(1, 'maxnoofQoSFlows')]}
    Items []QoSFlowItem
}

type QoSFlowItem struct { // [{'type': 'QoS-Flow-Identifier', 'name': 'qoS-Flow-Identifier'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoS-Flow-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    QoSFlowIdentifier QoSFlowIdentifier
    IEExtensions *QoSFlowItemExtIEs
}

func (self * QoSFlowItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.QoSFlowIdentifier.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &QoSFlowItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoS-Flow-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_QoSFlowItemExtIEs, order_QoSFlowItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * QoSFlowItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.QoSFlowIdentifier.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_QoSFlowItemExtIEs, order_QoSFlowItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *QoSFlowFailedList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(64)
    _size += 1
    self.Items = make([]QoSFlowFailedItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *QoSFlowFailedList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 64)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type QoSFlowFailedList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'QoS-Flow-Failed-Item'}, 'size': [(1, 'maxnoofQoSFlows')]}
    Items []QoSFlowFailedItem
}

type QoSFlowFailedItem struct { // [{'type': 'QoS-Flow-Identifier', 'name': 'qoS-Flow-Identifier'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoS-Flow-Failed-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    QoSFlowIdentifier QoSFlowIdentifier
    Cause Cause
    IEExtensions *QoSFlowFailedItemExtIEs
}

func (self * QoSFlowFailedItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.QoSFlowIdentifier.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &QoSFlowFailedItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoS-Flow-Failed-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_QoSFlowFailedItemExtIEs, order_QoSFlowFailedItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * QoSFlowFailedItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.QoSFlowIdentifier.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_QoSFlowFailedItemExtIEs, order_QoSFlowFailedItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *QoSFlowMappingList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(64)
    _size += 1
    self.Items = make([]QoSFlowMappingItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *QoSFlowMappingList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 64)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type QoSFlowMappingList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'QoS-Flow-Mapping-Item'}, 'size': [(1, 'maxnoofQoSFlows')]}
    Items []QoSFlowMappingItem
}

type QoSFlowMappingItem struct { // [{'type': 'QoS-Flow-Identifier', 'name': 'qoS-Flow-Identifier'}, {'type': 'QoS-Flow-Mapping-Indication', 'name': 'qoSFlowMappingIndication', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoS-Flow-Mapping-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    QoSFlowIdentifier QoSFlowIdentifier
    QoSFlowMappingIndication *QoSFlowMappingIndication
    IEExtensions *QoSFlowMappingItemExtIEs
}

func (self * QoSFlowMappingItem) Unpack(stream *Stream) {
    qoSFlowMappingIndication_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.QoSFlowIdentifier.Unpack(stream)// p8
    if (qoSFlowMappingIndication_flag & _flags) == qoSFlowMappingIndication_flag { //cond2
        self.QoSFlowMappingIndication = &QoSFlowMappingIndication{}//7{'type': 'QoS-Flow-Mapping-Indication', 'name': 'qoSFlowMappingIndication', 'optional': True}
        self.QoSFlowMappingIndication.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &QoSFlowMappingItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoS-Flow-Mapping-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_QoSFlowMappingItemExtIEs, order_QoSFlowMappingItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * QoSFlowMappingItem) Pack(stream *Stream) {
    const qoSFlowMappingIndication_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.QoSFlowIdentifier.Pack(stream)
    if self.QoSFlowMappingIndication != nil { 
        _flags |= qoSFlowMappingIndication_flag
        self.QoSFlowMappingIndication.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_QoSFlowMappingItemExtIEs, order_QoSFlowMappingItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type QoSFlowMappingIndication struct {
  Value int
}
const (
    QoSFlowMappingIndicationul = 0
    QoSFlowMappingIndicationdl = 1

    /* Extensions */
)
func (self *QoSFlowMappingIndication) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *QoSFlowMappingIndication) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type QoSParametersSupportList struct { // [{'type': 'EUTRAN-QoS-Support-List', 'name': 'eUTRAN-QoS-Support-List', 'optional': True}, {'type': 'NG-RAN-QoS-Support-List', 'name': 'nG-RAN-QoS-Support-List', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoS-Parameters-Support-List-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    EUTRANQoSSupportList *EUTRANQoSSupportList
    NGRANQoSSupportList *NGRANQoSSupportList
    IEExtensions *QoSParametersSupportListItemExtIEs
}

func (self * QoSParametersSupportList) Unpack(stream *Stream) {
    eUTRANQoSSupportList_flag := 0x00000002
    nGRANQoSSupportList_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (eUTRANQoSSupportList_flag & _flags) == eUTRANQoSSupportList_flag { //cond2
        self.EUTRANQoSSupportList = &EUTRANQoSSupportList{}//7{'type': 'EUTRAN-QoS-Support-List', 'name': 'eUTRAN-QoS-Support-List', 'optional': True}
        self.EUTRANQoSSupportList.Unpack(stream)// p8
    }
    if (nGRANQoSSupportList_flag & _flags) == nGRANQoSSupportList_flag { //cond2
        self.NGRANQoSSupportList = &NGRANQoSSupportList{}//7{'type': 'NG-RAN-QoS-Support-List', 'name': 'nG-RAN-QoS-Support-List', 'optional': True}
        self.NGRANQoSSupportList.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &QoSParametersSupportListItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoS-Parameters-Support-List-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_QoSParametersSupportListItemExtIEs, order_QoSParametersSupportListItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * QoSParametersSupportList) Pack(stream *Stream) {
    const eUTRANQoSSupportList_flag uint = 0x00000002
    const nGRANQoSSupportList_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.EUTRANQoSSupportList != nil { 
        _flags |= eUTRANQoSSupportList_flag
        self.EUTRANQoSSupportList.Pack(stream)
    }//end of optional
    if self.NGRANQoSSupportList != nil { 
        _flags |= nGRANQoSSupportList_flag
        self.NGRANQoSSupportList.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_QoSParametersSupportListItemExtIEs, order_QoSParametersSupportListItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type QoSPriorityLevel struct {
  Value uint64
}
func (self *QoSPriorityLevel) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(128, 8, 1, 0)
}
func (self * QoSPriorityLevel) Pack(st *Stream){
    st.formatf_Integer(self.Value, 128, 8, 1, 0)
}
func (self *QoSFlowQoSParameterList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(64)
    _size += 1
    self.Items = make([]QoSFlowQoSParameterItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *QoSFlowQoSParameterList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 64)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type QoSFlowQoSParameterList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'QoS-Flow-QoS-Parameter-Item'}, 'size': [(1, 'maxnoofQoSFlows')]}
    Items []QoSFlowQoSParameterItem
}

type QoSFlowQoSParameterItem struct { // [{'type': 'QoS-Flow-Identifier', 'name': 'qoS-Flow-Identifier'}, {'type': 'QoSFlowLevelQoSParameters', 'name': 'qoSFlowLevelQoSParameters'}, {'type': 'QoS-Flow-Mapping-Indication', 'name': 'qoSFlowMappingIndication', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoS-Flow-QoS-Parameter-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    QoSFlowIdentifier QoSFlowIdentifier
    QoSFlowLevelQoSParameters QoSFlowLevelQoSParameters
    QoSFlowMappingIndication *QoSFlowMappingIndication
    IEExtensions *QoSFlowQoSParameterItemExtIEs
}

func (self * QoSFlowQoSParameterItem) Unpack(stream *Stream) {
    qoSFlowMappingIndication_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.QoSFlowIdentifier.Unpack(stream)// p8
    self.QoSFlowLevelQoSParameters.Unpack(stream)// p8
    if (qoSFlowMappingIndication_flag & _flags) == qoSFlowMappingIndication_flag { //cond2
        self.QoSFlowMappingIndication = &QoSFlowMappingIndication{}//7{'type': 'QoS-Flow-Mapping-Indication', 'name': 'qoSFlowMappingIndication', 'optional': True}
        self.QoSFlowMappingIndication.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &QoSFlowQoSParameterItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoS-Flow-QoS-Parameter-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_QoSFlowQoSParameterItemExtIEs, order_QoSFlowQoSParameterItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * QoSFlowQoSParameterItem) Pack(stream *Stream) {
    const qoSFlowMappingIndication_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.QoSFlowIdentifier.Pack(stream)
    self.QoSFlowLevelQoSParameters.Pack(stream)
    if self.QoSFlowMappingIndication != nil { 
        _flags |= qoSFlowMappingIndication_flag
        self.QoSFlowMappingIndication.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_QoSFlowQoSParameterItemExtIEs, order_QoSFlowQoSParameterItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type QoSFlowLevelQoSParameters struct { // [{'type': 'QoS-Characteristics', 'name': 'qoS-Characteristics'}, {'type': 'NGRANAllocationAndRetentionPriority', 'name': 'nGRANallocationRetentionPriority'}, {'type': 'GBR-QoSFlowInformation', 'name': 'gBR-QoS-Flow-Information', 'optional': True}, {'type': 'ENUMERATED', 'values': [('subject-to', 0), None], 'name': 'reflective-QoS-Attribute', 'optional': True}, {'type': 'ENUMERATED', 'values': [('more-likely', 0), None], 'name': 'additional-QoS-Information', 'optional': True}, {'type': 'INTEGER', 'restricted-to': [(1, 8), None], 'name': 'paging-Policy-Indicator', 'optional': True}, {'type': 'ENUMERATED', 'values': [('enabled', 0), None], 'name': 'reflective-QoS-Indicator', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoSFlowLevelQoSParameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    QoSCharacteristics QoSCharacteristics
    NGRANallocationRetentionPriority NGRANAllocationAndRetentionPriority
    GBRQoSFlowInformation *GBRQoSFlowInformation
    ReflectiveQoSAttribute *ENUMERATED
    AdditionalQoSInformation *ENUMERATED
    PagingPolicyIndicator *INTEGER
    ReflectiveQoSIndicator *ENUMERATED
    IEExtensions *QoSFlowLevelQoSParametersExtIEs
}

func (self * QoSFlowLevelQoSParameters) Unpack(stream *Stream) {
    gBRQoSFlowInformation_flag := 0x00000001
    reflectiveQoSAttribute_flag := 0x00000002
    additionalQoSInformation_flag := 0x00000004
    pagingPolicyIndicator_flag := 0x00000008
    reflectiveQoSIndicator_flag := 0x00000010
    iEExtensions_flag := 0x00000020
    _flags := 0
    _flags = stream.get_flags(6)
    self.QoSCharacteristics.Unpack(stream)// p8
    self.NGRANallocationRetentionPriority.Unpack(stream)// p8
    if (gBRQoSFlowInformation_flag & _flags) == gBRQoSFlowInformation_flag { //cond2
        self.GBRQoSFlowInformation = &GBRQoSFlowInformation{}//7{'type': 'GBR-QoSFlowInformation', 'name': 'gBR-QoS-Flow-Information', 'optional': True}
        self.GBRQoSFlowInformation.Unpack(stream)// p8
    }
    if (reflectiveQoSAttribute_flag & _flags) == reflectiveQoSAttribute_flag { //cond1
        var Unpack_reflectiveQoSAttribute = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.ReflectiveQoSAttribute = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('subject-to', 0), None], 'name': 'reflective-QoS-Attribute', 'optional': True}
        Unpack_reflectiveQoSAttribute(stream, self.ReflectiveQoSAttribute)// p1 {'type': 'ENUMERATED', 'values': [('subject-to', 0), None], 'name': 'reflective-QoS-Attribute', 'optional': True}
    }
    if (additionalQoSInformation_flag & _flags) == additionalQoSInformation_flag { //cond1
        var Unpack_additionalQoSInformation = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.AdditionalQoSInformation = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('more-likely', 0), None], 'name': 'additional-QoS-Information', 'optional': True}
        Unpack_additionalQoSInformation(stream, self.AdditionalQoSInformation)// p1 {'type': 'ENUMERATED', 'values': [('more-likely', 0), None], 'name': 'additional-QoS-Information', 'optional': True}
    }
    if (pagingPolicyIndicator_flag & _flags) == pagingPolicyIndicator_flag { //cond1
        var Unpack_pagingPolicyIndicator = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(8, 4, 1, 1)
        }
        self.PagingPolicyIndicator = &INTEGER{}//6{'type': 'INTEGER', 'restricted-to': [(1, 8), None], 'name': 'paging-Policy-Indicator', 'optional': True}
        Unpack_pagingPolicyIndicator(stream, self.PagingPolicyIndicator)// p1 {'type': 'INTEGER', 'restricted-to': [(1, 8), None], 'name': 'paging-Policy-Indicator', 'optional': True}
    }
    if (reflectiveQoSIndicator_flag & _flags) == reflectiveQoSIndicator_flag { //cond1
        var Unpack_reflectiveQoSIndicator = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.ReflectiveQoSIndicator = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('enabled', 0), None], 'name': 'reflective-QoS-Indicator', 'optional': True}
        Unpack_reflectiveQoSIndicator(stream, self.ReflectiveQoSIndicator)// p1 {'type': 'ENUMERATED', 'values': [('enabled', 0), None], 'name': 'reflective-QoS-Indicator', 'optional': True}
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &QoSFlowLevelQoSParametersExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['QoSFlowLevelQoSParameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_QoSFlowLevelQoSParametersExtIEs, order_QoSFlowLevelQoSParametersExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * QoSFlowLevelQoSParameters) Pack(stream *Stream) {
    const gBRQoSFlowInformation_flag uint = 0x00000001
    const reflectiveQoSAttribute_flag uint = 0x00000002
    const additionalQoSInformation_flag uint = 0x00000004
    const pagingPolicyIndicator_flag uint = 0x00000008
    const reflectiveQoSIndicator_flag uint = 0x00000010
    const iEExtensions_flag uint = 0x00000020
    _flagReserve := stream.reserve_flags(6)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.QoSCharacteristics.Pack(stream)
    self.NGRANallocationRetentionPriority.Pack(stream)
    if self.GBRQoSFlowInformation != nil { 
        _flags |= gBRQoSFlowInformation_flag
        self.GBRQoSFlowInformation.Pack(stream)
    }//end of optional
    if self.ReflectiveQoSAttribute != nil { //YY
        _flags |= reflectiveQoSAttribute_flag
        var Pack_reflectiveQoSAttribute = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_reflectiveQoSAttribute(stream, *self.ReflectiveQoSAttribute) //f1
    }//end of optional
    if self.AdditionalQoSInformation != nil { //YY
        _flags |= additionalQoSInformation_flag
        var Pack_additionalQoSInformation = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_additionalQoSInformation(stream, *self.AdditionalQoSInformation) //f1
    }//end of optional
    if self.PagingPolicyIndicator != nil { //YY
        _flags |= pagingPolicyIndicator_flag
        var Pack_pagingPolicyIndicator = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 8, 4, 1, 1)
        }
        Pack_pagingPolicyIndicator(stream, *self.PagingPolicyIndicator) //f1
    }//end of optional
    if self.ReflectiveQoSIndicator != nil { //YY
        _flags |= reflectiveQoSIndicator_flag
        var Pack_reflectiveQoSIndicator = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_reflectiveQoSIndicator(stream, *self.ReflectiveQoSIndicator) //f1
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_QoSFlowLevelQoSParametersExtIEs, order_QoSFlowLevelQoSParametersExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

type RANUEID struct {
  Value HexBytes
}
func (self *RANUEID) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(8)
}
func (self *RANUEID) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 8)
}
type RATType struct {
  Value int
}
const (
    RATTypee_UTRA = 0
    RATTypenR = 1

    /* Extensions */
)
func (self *RATType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *RATType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type RLCMode struct {
  Value int
}
const (
    RLCModerlc_tm = 0
    RLCModerlc_am = 1
    RLCModerlc_um_bidirectional = 2
    RLCModerlc_um_unidirectional_ul = 3
    RLCModerlc_um_unidirectional_dl = 4

    /* Extensions */
)
func (self *RLCMode) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 5, 1)
}
func (self *RLCMode) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 5, 1)
}
func (self *ROHCParameters)Unpack(stream *Stream) {
    //coptions := []string{"rOHC","uPlinkOnlyROHC","choice-Extension","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        self.ROHC = &ROHC{}//cho6
        self.ROHC.Unpack(stream)
    } else if choice == 1 { //ch2
        self.UPlinkOnlyROHC = &UplinkOnlyROHC{}//cho6
        self.UPlinkOnlyROHC.Unpack(stream)
    } else if choice == 2 { //ch2
        ChoiceExtension := &ProtocolIESingleContainer{}//cho2
        self.ChoiceExtension = &ROHCParametersExtIEs{}//cho3
        ChoiceExtension.Unpack(stream, self.ChoiceExtension)
    }//end of if else

}
func (self * ROHCParameters) Pack(stream *Stream) {
    if self.ROHC != nil {
        stream.set_choice(0, 2, 0, 3)
        self.ROHC.Pack(stream)//2
    } else if self.UPlinkOnlyROHC != nil {
        stream.set_choice(1, 2, 0, 3)
        self.UPlinkOnlyROHC.Pack(stream)//2
    } else if self.ChoiceExtension != nil {
        stream.set_choice(2, 2, 0, 3)
        ChoiceExtension := ProtocolIESingleContainer{}//cho2
        ChoiceExtension.Pack(stream, self.ChoiceExtension)
    }

}
type ROHCParameters struct { //[{'type': 'ROHC', 'name': 'rOHC'}, {'type': 'UplinkOnlyROHC', 'name': 'uPlinkOnlyROHC'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['ROHC-Parameters-ExtIEs'], 'name': 'choice-Extension'}]
    ROHC *ROHC
    UPlinkOnlyROHC *UplinkOnlyROHC
    ChoiceExtension *ROHCParametersExtIEs
} // ROHCParameters

type ROHC struct { // [{'type': 'INTEGER', 'restricted-to': [(0, 16383), None], 'name': 'maxCID'}, {'type': 'INTEGER', 'restricted-to': [(0, 511), None], 'name': 'rOHC-Profiles'}, {'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 'continueROHC', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ROHC-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    MaxCID INTEGER
    ROHCProfiles INTEGER
    ContinueROHC *ENUMERATED
    IEExtensions *ROHCExtIEs
}

func (self * ROHC) Unpack(stream *Stream) {
    continueROHC_flag := 0x00000001
    iEExtensions_flag := 0x00000002
    _flags := 0
    _flags = stream.get_flags(2)
    var Unpack_maxCID = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(16384, 15, 1, 0)
    }
    Unpack_maxCID(stream, &self.MaxCID)// p2
    var Unpack_rOHCProfiles = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(512, 10, 1, 0)
    }
    Unpack_rOHCProfiles(stream, &self.ROHCProfiles)// p2
    if (continueROHC_flag & _flags) == continueROHC_flag { //cond1
        var Unpack_continueROHC = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.ContinueROHC = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 'continueROHC', 'optional': True}
        Unpack_continueROHC(stream, self.ContinueROHC)// p1 {'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 'continueROHC', 'optional': True}
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &ROHCExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ROHC-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_ROHCExtIEs, order_ROHCExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * ROHC) Pack(stream *Stream) {
    const continueROHC_flag uint = 0x00000001
    const iEExtensions_flag uint = 0x00000002
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_maxCID = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 16384, 15, 1, 0)
    }
    Pack_maxCID(stream, self.MaxCID) //f2
    var Pack_rOHCProfiles = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 512, 10, 1, 0)
    }
    Pack_rOHCProfiles(stream, self.ROHCProfiles) //f2
    if self.ContinueROHC != nil { //YY
        _flags |= continueROHC_flag
        var Pack_continueROHC = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_continueROHC(stream, *self.ContinueROHC) //f1
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_ROHCExtIEs, order_ROHCExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SecurityAlgorithm struct { // [{'type': 'CipheringAlgorithm', 'name': 'cipheringAlgorithm'}, {'type': 'IntegrityProtectionAlgorithm', 'name': 'integrityProtectionAlgorithm', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityAlgorithm-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    CipheringAlgorithm CipheringAlgorithm
    IntegrityProtectionAlgorithm *IntegrityProtectionAlgorithm
    IEExtensions *SecurityAlgorithmExtIEs
}

func (self * SecurityAlgorithm) Unpack(stream *Stream) {
    integrityProtectionAlgorithm_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.CipheringAlgorithm.Unpack(stream)// p8
    if (integrityProtectionAlgorithm_flag & _flags) == integrityProtectionAlgorithm_flag { //cond2
        self.IntegrityProtectionAlgorithm = &IntegrityProtectionAlgorithm{}//7{'type': 'IntegrityProtectionAlgorithm', 'name': 'integrityProtectionAlgorithm', 'optional': True}
        self.IntegrityProtectionAlgorithm.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SecurityAlgorithmExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityAlgorithm-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SecurityAlgorithmExtIEs, order_SecurityAlgorithmExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SecurityAlgorithm) Pack(stream *Stream) {
    const integrityProtectionAlgorithm_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CipheringAlgorithm.Pack(stream)
    if self.IntegrityProtectionAlgorithm != nil { 
        _flags |= integrityProtectionAlgorithm_flag
        self.IntegrityProtectionAlgorithm.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SecurityAlgorithmExtIEs, order_SecurityAlgorithmExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type SecurityIndication struct { // [{'type': 'IntegrityProtectionIndication', 'name': 'integrityProtectionIndication'}, {'type': 'ConfidentialityProtectionIndication', 'name': 'confidentialityProtectionIndication'}, {'type': 'MaximumIPdatarate', 'name': 'maximumIPdatarate', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityIndication-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    IntegrityProtectionIndication IntegrityProtectionIndication
    ConfidentialityProtectionIndication ConfidentialityProtectionIndication
    MaximumIPdatarate *MaximumIPdatarate
    IEExtensions *SecurityIndicationExtIEs
}

func (self * SecurityIndication) Unpack(stream *Stream) {
    maximumIPdatarate_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.IntegrityProtectionIndication.Unpack(stream)// p8
    self.ConfidentialityProtectionIndication.Unpack(stream)// p8
    if (maximumIPdatarate_flag & _flags) == maximumIPdatarate_flag { //cond2
        self.MaximumIPdatarate = &MaximumIPdatarate{}//7{'type': 'MaximumIPdatarate', 'name': 'maximumIPdatarate', 'optional': True}
        self.MaximumIPdatarate.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SecurityIndicationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityIndication-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SecurityIndicationExtIEs, order_SecurityIndicationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SecurityIndication) Pack(stream *Stream) {
    const maximumIPdatarate_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IntegrityProtectionIndication.Pack(stream)
    self.ConfidentialityProtectionIndication.Pack(stream)
    if self.MaximumIPdatarate != nil { 
        _flags |= maximumIPdatarate_flag
        self.MaximumIPdatarate.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SecurityIndicationExtIEs, order_SecurityIndicationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type SecurityInformation struct { // [{'type': 'SecurityAlgorithm', 'name': 'securityAlgorithm'}, {'type': 'UPSecuritykey', 'name': 'uPSecuritykey'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    SecurityAlgorithm SecurityAlgorithm
    UPSecuritykey UPSecuritykey
    IEExtensions *SecurityInformationExtIEs
}

func (self * SecurityInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.SecurityAlgorithm.Unpack(stream)// p8
    self.UPSecuritykey.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SecurityInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SecurityInformationExtIEs, order_SecurityInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SecurityInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.SecurityAlgorithm.Pack(stream)
    self.UPSecuritykey.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SecurityInformationExtIEs, order_SecurityInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SecurityResult struct { // [{'type': 'IntegrityProtectionResult', 'name': 'integrityProtectionResult'}, {'type': 'ConfidentialityProtectionResult', 'name': 'confidentialityProtectionResult'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityResult-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    IntegrityProtectionResult IntegrityProtectionResult
    ConfidentialityProtectionResult ConfidentialityProtectionResult
    IEExtensions *SecurityResultExtIEs
}

func (self * SecurityResult) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.IntegrityProtectionResult.Unpack(stream)// p8
    self.ConfidentialityProtectionResult.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SecurityResultExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityResult-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SecurityResultExtIEs, order_SecurityResultExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SecurityResult) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IntegrityProtectionResult.Pack(stream)
    self.ConfidentialityProtectionResult.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SecurityResultExtIEs, order_SecurityResultExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *SliceSupportList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(1024)
    _size += 1
    self.Items = make([]SliceSupportItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *SliceSupportList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 1024)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type SliceSupportList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'Slice-Support-Item'}, 'size': [(1, 'maxnoofSliceItems')]}
    Items []SliceSupportItem
}

type SliceSupportItem struct { // [{'type': 'SNSSAI', 'name': 'sNSSAI'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Slice-Support-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    SNSSAI SNSSAI
    IEExtensions *SliceSupportItemExtIEs
}

func (self * SliceSupportItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.SNSSAI.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SliceSupportItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Slice-Support-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SliceSupportItemExtIEs, order_SliceSupportItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * SliceSupportItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.SNSSAI.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SliceSupportItemExtIEs, order_SliceSupportItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SNSSAI struct { // [{'type': 'OCTET STRING', 'size': [1], 'name': 'sST'}, {'type': 'OCTET STRING', 'size': [3], 'name': 'sD', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SNSSAI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    SST OCTETSTRING
    SD *OCTETSTRING
    IEExtensions *SNSSAIExtIEs
}

func (self * SNSSAI) Unpack(stream *Stream) {
    sD_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    var Unpack_sST = func(st *Stream, self *OCTETSTRING) {
        self.Value = st.parsef_OctString(1)
    }
    Unpack_sST(stream, &self.SST)// p2
    if (sD_flag & _flags) == sD_flag { //cond1
        var Unpack_sD = func(st *Stream, self *OCTETSTRING) {
            self.Value = st.parsef_OctString(3)
        }
        self.SD = &OCTETSTRING{}//6{'type': 'OCTET STRING', 'size': [3], 'name': 'sD', 'optional': True}
        Unpack_sD(stream, self.SD)// p1 {'type': 'OCTET STRING', 'size': [3], 'name': 'sD', 'optional': True}
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SNSSAIExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SNSSAI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SNSSAIExtIEs, order_SNSSAIExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SNSSAI) Pack(stream *Stream) {
    const sD_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_sST = func(st *Stream, self OCTETSTRING) {
        st.formatf_OctString(self.Value, 1)
    }
    Pack_sST(stream, self.SST) //f2
    if self.SD != nil { //YY
        _flags |= sD_flag
        var Pack_sD = func(st *Stream, self OCTETSTRING) {
            st.formatf_OctString(self.Value, 3)
        }
        Pack_sD(stream, *self.SD) //f1
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SNSSAIExtIEs, order_SNSSAIExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type SDAPConfiguration struct { // [{'type': 'DefaultDRB', 'name': 'defaultDRB'}, {'type': 'SDAP-Header-UL', 'name': 'sDAP-Header-UL'}, {'type': 'SDAP-Header-DL', 'name': 'sDAP-Header-DL'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDAP-Configuration-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DefaultDRB DefaultDRB
    SDAPHeaderUL SDAPHeaderUL
    SDAPHeaderDL SDAPHeaderDL
    IEExtensions *SDAPConfigurationExtIEs
}

func (self * SDAPConfiguration) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.DefaultDRB.Unpack(stream)// p8
    self.SDAPHeaderUL.Unpack(stream)// p8
    self.SDAPHeaderDL.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SDAPConfigurationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDAP-Configuration-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SDAPConfigurationExtIEs, order_SDAPConfigurationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SDAPConfiguration) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.DefaultDRB.Pack(stream)
    self.SDAPHeaderUL.Pack(stream)
    self.SDAPHeaderDL.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SDAPConfigurationExtIEs, order_SDAPConfigurationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SDAPHeaderDL struct {
  Value int
}
const (
    SDAPHeaderDLpresent = 0
    SDAPHeaderDLabsent = 1

    /* Extensions */
)
func (self *SDAPHeaderDL) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *SDAPHeaderDL) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type SDAPHeaderUL struct {
  Value int
}
const (
    SDAPHeaderULpresent = 0
    SDAPHeaderULabsent = 1

    /* Extensions */
)
func (self *SDAPHeaderUL) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *SDAPHeaderUL) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type TimeToWait struct {
  Value int
}
const (
    TimeToWaitv1s = 0
    TimeToWaitv2s = 1
    TimeToWaitv5s = 2
    TimeToWaitv10s = 3
    TimeToWaitv20s = 4
    TimeToWaitv60s = 5

    /* Extensions */
)
func (self *TimeToWait) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 6, 1)
}
func (self *TimeToWait) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 6, 1)
}
type TNLAssociationUsage struct {
  Value int
}
const (
    TNLAssociationUsageue = 0
    TNLAssociationUsagenon_ue = 1
    TNLAssociationUsageboth = 2

    /* Extensions */
)
func (self *TNLAssociationUsage) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *TNLAssociationUsage) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type TransportLayerAddress struct {
  Len int
  Value HexBytes
}
func (self *TransportLayerAddress) Unpack(st *Stream){
    self.Len = int(st.parse_blen(9, 1)+1)
    self.Value = st.parsef_BitString(160, int(self.Len))
}
func (self *TransportLayerAddress) Pack(st *Stream) {
    st.format_blen(int(self.Len-1), 9, 1)
    st.formatf_BitString(self.Value, int(self.Len))
}
type TransactionID struct {
  Value uint64
}
func (self *TransactionID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 9, 1, 0)
}
func (self * TransactionID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 9, 1, 0)
}
type TReordering struct {
  Value int
}
const (
    TReorderingms0 = 0
    TReorderingms1 = 1
    TReorderingms2 = 2
    TReorderingms4 = 3
    TReorderingms5 = 4
    TReorderingms8 = 5
    TReorderingms10 = 6
    TReorderingms15 = 7
    TReorderingms20 = 8
    TReorderingms30 = 9
    TReorderingms40 = 10
    TReorderingms50 = 11
    TReorderingms60 = 12
    TReorderingms80 = 13
    TReorderingms100 = 14
    TReorderingms120 = 15
    TReorderingms140 = 16
    TReorderingms160 = 17
    TReorderingms180 = 18
    TReorderingms200 = 19
    TReorderingms220 = 20
    TReorderingms240 = 21
    TReorderingms260 = 22
    TReorderingms280 = 23
    TReorderingms300 = 24
    TReorderingms500 = 25
    TReorderingms750 = 26
    TReorderingms1000 = 27
    TReorderingms1250 = 28
    TReorderingms1500 = 29
    TReorderingms1750 = 30
    TReorderingms2000 = 31
    TReorderingms2250 = 32
    TReorderingms2500 = 33
    TReorderingms2750 = 34
    TReorderingms3000 = 35

    /* Extensions */
)
func (self *TReordering) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(7, 36, 1)
}
func (self *TReordering) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 7, 36, 1)
}
type TReorderingTimer struct { // [{'type': 'T-Reordering', 'name': 't-Reordering'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['T-ReorderingTimer-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    TReordering TReordering
    IEExtensions *TReorderingTimerExtIEs
}

func (self * TReorderingTimer) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.TReordering.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &TReorderingTimerExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['T-ReorderingTimer-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_TReorderingTimerExtIEs, order_TReorderingTimerExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TReorderingTimer) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TReordering.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_TReorderingTimerExtIEs, order_TReorderingTimerExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type TypeOfError struct {
  Value int
}
const (
    TypeOfErrornot_understood = 0
    TypeOfErrormissing = 1

    /* Extensions */
)
func (self *TypeOfError) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *TypeOfError) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type UEActivity struct {
  Value int
}
const (
    UEActivityactive = 0
    UEActivitynot_active = 1

    /* Extensions */
)
func (self *UEActivity) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *UEActivity) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type UEassociatedLogicalE1ConnectionItem struct { // [{'type': 'GNB-CU-CP-UE-E1AP-ID', 'name': 'gNB-CU-CP-UE-E1AP-ID', 'optional': True}, {'type': 'GNB-CU-UP-UE-E1AP-ID', 'name': 'gNB-CU-UP-UE-E1AP-ID', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UE-associatedLogicalE1-ConnectionItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    GNBCUCPUEE1APID *GNBCUCPUEE1APID
    GNBCUUPUEE1APID *GNBCUUPUEE1APID
    IEExtensions *UEassociatedLogicalE1ConnectionItemExtIEs
}

func (self * UEassociatedLogicalE1ConnectionItem) Unpack(stream *Stream) {
    gNBCUCPUEE1APID_flag := 0x00000002
    gNBCUUPUEE1APID_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (gNBCUCPUEE1APID_flag & _flags) == gNBCUCPUEE1APID_flag { //cond2
        self.GNBCUCPUEE1APID = &GNBCUCPUEE1APID{}//7{'type': 'GNB-CU-CP-UE-E1AP-ID', 'name': 'gNB-CU-CP-UE-E1AP-ID', 'optional': True}
        self.GNBCUCPUEE1APID.Unpack(stream)// p8
    }
    if (gNBCUUPUEE1APID_flag & _flags) == gNBCUUPUEE1APID_flag { //cond2
        self.GNBCUUPUEE1APID = &GNBCUUPUEE1APID{}//7{'type': 'GNB-CU-UP-UE-E1AP-ID', 'name': 'gNB-CU-UP-UE-E1AP-ID', 'optional': True}
        self.GNBCUUPUEE1APID.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &UEassociatedLogicalE1ConnectionItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UE-associatedLogicalE1-ConnectionItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_UEassociatedLogicalE1ConnectionItemExtIEs, order_UEassociatedLogicalE1ConnectionItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEassociatedLogicalE1ConnectionItem) Pack(stream *Stream) {
    const gNBCUCPUEE1APID_flag uint = 0x00000002
    const gNBCUUPUEE1APID_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.GNBCUCPUEE1APID != nil { 
        _flags |= gNBCUCPUEE1APID_flag
        self.GNBCUCPUEE1APID.Pack(stream)
    }//end of optional
    if self.GNBCUUPUEE1APID != nil { 
        _flags |= gNBCUUPUEE1APID_flag
        self.GNBCUUPUEE1APID.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_UEassociatedLogicalE1ConnectionItemExtIEs, order_UEassociatedLogicalE1ConnectionItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type ULConfiguration struct {
  Value int
}
const (
    ULConfigurationno_data = 0
    ULConfigurationshared = 1
    ULConfigurationonly = 2

    /* Extensions */
)
func (self *ULConfiguration) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *ULConfiguration) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type ULDataSplitThreshold struct {
  Value int
}
const (
    ULDataSplitThresholdb0 = 0
    ULDataSplitThresholdb100 = 1
    ULDataSplitThresholdb200 = 2
    ULDataSplitThresholdb400 = 3
    ULDataSplitThresholdb800 = 4
    ULDataSplitThresholdb1600 = 5
    ULDataSplitThresholdb3200 = 6
    ULDataSplitThresholdb6400 = 7
    ULDataSplitThresholdb12800 = 8
    ULDataSplitThresholdb25600 = 9
    ULDataSplitThresholdb51200 = 10
    ULDataSplitThresholdb102400 = 11
    ULDataSplitThresholdb204800 = 12
    ULDataSplitThresholdb409600 = 13
    ULDataSplitThresholdb819200 = 14
    ULDataSplitThresholdb1228800 = 15
    ULDataSplitThresholdb1638400 = 16
    ULDataSplitThresholdb2457600 = 17
    ULDataSplitThresholdb3276800 = 18
    ULDataSplitThresholdb4096000 = 19
    ULDataSplitThresholdb4915200 = 20
    ULDataSplitThresholdb5734400 = 21
    ULDataSplitThresholdb6553600 = 22
    ULDataSplitThresholdinfinity = 23

    /* Extensions */
)
func (self *ULDataSplitThreshold) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(6, 24, 1)
}
func (self *ULDataSplitThreshold) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 6, 24, 1)
}
func (self *UPParameters) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(8)
    _size += 1
    self.Items = make([]UPParametersItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *UPParameters) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 8)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type UPParameters struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'UP-Parameters-Item'}, 'size': [(1, 'maxnoofUPParameters')]}
    Items []UPParametersItem
}

type UPParametersItem struct { // [{'type': 'UP-TNL-Information', 'name': 'uP-TNL-Information'}, {'type': 'Cell-Group-ID', 'name': 'cell-Group-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UP-Parameters-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    UPTNLInformation UPTNLInformation
    CellGroupID CellGroupID
    IEExtensions *UPParametersItemExtIEs
}

func (self * UPParametersItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.UPTNLInformation.Unpack(stream)// p8
    self.CellGroupID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &UPParametersItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UP-Parameters-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_UPParametersItemExtIEs, order_UPParametersItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UPParametersItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UPTNLInformation.Pack(stream)
    self.CellGroupID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_UPParametersItemExtIEs, order_UPParametersItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UPSecuritykey struct { // [{'type': 'EncryptionKey', 'name': 'encryptionKey'}, {'type': 'IntegrityProtectionKey', 'name': 'integrityProtectionKey', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UPSecuritykey-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    EncryptionKey EncryptionKey
    IntegrityProtectionKey *IntegrityProtectionKey
    IEExtensions *UPSecuritykeyExtIEs
}

func (self * UPSecuritykey) Unpack(stream *Stream) {
    integrityProtectionKey_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.EncryptionKey.Unpack(stream)// p8
    if (integrityProtectionKey_flag & _flags) == integrityProtectionKey_flag { //cond2
        self.IntegrityProtectionKey = &IntegrityProtectionKey{}//7{'type': 'IntegrityProtectionKey', 'name': 'integrityProtectionKey', 'optional': True}
        self.IntegrityProtectionKey.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &UPSecuritykeyExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UPSecuritykey-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_UPSecuritykeyExtIEs, order_UPSecuritykeyExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UPSecuritykey) Pack(stream *Stream) {
    const integrityProtectionKey_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.EncryptionKey.Pack(stream)
    if self.IntegrityProtectionKey != nil { 
        _flags |= integrityProtectionKey_flag
        self.IntegrityProtectionKey.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_UPSecuritykeyExtIEs, order_UPSecuritykeyExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

func (self *UPTNLInformation)Unpack(stream *Stream) {
    //coptions := []string{"gTPTunnel","choice-extension"}
    choice := stream.get_choice(1, 0, 2)
    if choice == 0 { //ch1
        self.GTPTunnel = &GTPTunnel{}//cho6
        self.GTPTunnel.Unpack(stream)
    } else if choice == 1 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &UPTNLInformationExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

}
func (self * UPTNLInformation) Pack(stream *Stream) {
    if self.GTPTunnel != nil {
        stream.set_choice(0, 1, 0, 2)
        self.GTPTunnel.Pack(stream)//2
    } else if self.Choiceextension != nil {
        stream.set_choice(1, 1, 0, 2)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type UPTNLInformation struct { //[{'type': 'GTPTunnel', 'name': 'gTPTunnel'}, {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['UP-TNL-Information-ExtIEs'], 'name': 'choice-extension'}]
    GTPTunnel *GTPTunnel
    Choiceextension *UPTNLInformationExtIEs
} // UPTNLInformation

type UplinkOnlyROHC struct { // [{'type': 'INTEGER', 'restricted-to': [(0, 16383), None], 'name': 'maxCID'}, {'type': 'INTEGER', 'restricted-to': [(0, 511), None], 'name': 'rOHC-Profiles'}, {'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 'continueROHC', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UplinkOnlyROHC-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    MaxCID INTEGER
    ROHCProfiles INTEGER
    ContinueROHC *ENUMERATED
    IEExtensions *UplinkOnlyROHCExtIEs
}

func (self * UplinkOnlyROHC) Unpack(stream *Stream) {
    continueROHC_flag := 0x00000001
    iEExtensions_flag := 0x00000002
    _flags := 0
    _flags = stream.get_flags(2)
    var Unpack_maxCID = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(16384, 15, 1, 0)
    }
    Unpack_maxCID(stream, &self.MaxCID)// p2
    var Unpack_rOHCProfiles = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(512, 10, 1, 0)
    }
    Unpack_rOHCProfiles(stream, &self.ROHCProfiles)// p2
    if (continueROHC_flag & _flags) == continueROHC_flag { //cond1
        var Unpack_continueROHC = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.ContinueROHC = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 'continueROHC', 'optional': True}
        Unpack_continueROHC(stream, self.ContinueROHC)// p1 {'type': 'ENUMERATED', 'values': [('tRue', 0), None], 'name': 'continueROHC', 'optional': True}
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &UplinkOnlyROHCExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UplinkOnlyROHC-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_UplinkOnlyROHCExtIEs, order_UplinkOnlyROHCExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * UplinkOnlyROHC) Pack(stream *Stream) {
    const continueROHC_flag uint = 0x00000001
    const iEExtensions_flag uint = 0x00000002
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_maxCID = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 16384, 15, 1, 0)
    }
    Pack_maxCID(stream, self.MaxCID) //f2
    var Pack_rOHCProfiles = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 512, 10, 1, 0)
    }
    Pack_rOHCProfiles(stream, self.ROHCProfiles) //f2
    if self.ContinueROHC != nil { //YY
        _flags |= continueROHC_flag
        var Pack_continueROHC = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_continueROHC(stream, *self.ContinueROHC) //f1
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_UplinkOnlyROHCExtIEs, order_UplinkOnlyROHCExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type Criticality struct {
  Value int
}
const (
    Criticalityreject = 0
    Criticalityignore = 1
    Criticalitynotify = 2
)
func (self *Criticality) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 3, 0)
}
func (self *Criticality) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 3, 0)
}
type Presence struct {
  Value int
}
const (
    Presenceoptional = 0
    Presenceconditional = 1
    Presencemandatory = 2
)
func (self *Presence) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 3, 0)
}
func (self *Presence) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 3, 0)
}
func (self *PrivateIEID)Unpack(stream *Stream) {
    //coptions := []string{"local","global"}
    choice := stream.get_choice(1, 0, 2)
    if choice == 0 { //ch1
        var Unpack_local = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(65536, 16, 0, 0)
        }
        self.Local = &INTEGER{}//cho5
        Unpack_local(stream, self.Local);
    } else if choice == 1 { //ch2
        var Unpack_global = func(st *Stream, self *OBJECTIDENTIFIER ) {
          self.Value = st.parsef_ObjectID()
        }
        self.Global = &OBJECTIDENTIFIER{}//cho5
        Unpack_global(stream, self.Global);
    }//end of if else

}
func (self * PrivateIEID) Pack(stream *Stream) {
    if self.Local != nil {
        stream.set_choice(0, 1, 0, 2)
        var Pack_local = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 65536, 16, 0, 0)
        }
        Pack_local(stream, *self.Local)//3
    } else if self.Global != nil {
        stream.set_choice(1, 1, 0, 2)
        var Pack_global = func(stream *Stream, self OBJECTIDENTIFIER) {
          //st.formatf_ObjectID(self.Value)
        }
        Pack_global(stream, *self.Global)//3
    }

}
type PrivateIEID struct { //[{'type': 'INTEGER', 'restricted-to': [(0, 'maxPrivateIEs')], 'name': 'local'}, {'type': 'OBJECT IDENTIFIER', 'name': 'global'}]
    Local *INTEGER
    Global *OBJECTIDENTIFIER
} // PrivateIEID

type ProcedureCode struct {
  Value uint64
}
func (self *ProcedureCode) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * ProcedureCode) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
type ProtocolExtensionID struct {
  Value uint64
}
func (self *ProtocolExtensionID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * ProtocolExtensionID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type ProtocolIEID struct {
  Value uint64
}
func (self *ProtocolIEID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * ProtocolIEID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type TriggeringMessage struct {
  Value int
}
const (
    TriggeringMessageinitiating_message = 0
    TriggeringMessagesuccessful_outcome = 1
    TriggeringMessageunsuccessful_outcome = 2
)
func (self *TriggeringMessage) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 3, 0)
}
func (self *TriggeringMessage) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 3, 0)
}
func (self *ProtocolIEContainer) Unpack(stream *Stream, out interface{}) { //SeqOF1
    _size := stream.get_listsize(65536) + 0
    for i := 0; i <_size; i +=1 {
        item := ProtocolIEField{}
        item.Unpack(stream, out)
    }
    return 
}


func (self * ProtocolIEContainer) Pack(stream *Stream, data interface{}) { //sqof 1
    _size := data.(E1APPROTOCOLIES_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['E1AP-PROTOCOL-IES']}
    Items map[int]*E1APPROTOCOLIES
    order []int
}

type ProtocolIESingleContainer struct{ //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['E1AP-PROTOCOL-IES']}
   Item map[int]*E1APPROTOCOLIES //UserType
   order []int
}
func (self *ProtocolIESingleContainer) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['E1AP-PROTOCOL-IES']}
    val := ProtocolIEField{} //ut2
    val.Unpack(stream, out) //ut2
}

func (self *ProtocolIESingleContainer) Pack(stream *Stream, data interface{}){
    // UserType
    if len(self.order) > 0 { //ut2
       k := self.order[0] 
       val := self.Item[k] //ut2
       val.Pack(stream, data)
    }
}

type ProtocolIEField struct { // [{'type': 'E1AP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'E1AP-PROTOCOL-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'E1AP-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id E1APPROTOCOLIESid
    Criticality E1APPROTOCOLIEScriticality
    Value E1APPROTOCOLIESValue
}

func (self * ProtocolIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'E1AP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := E1APPROTOCOLIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(E1APPROTOCOLIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'E1AP-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * E1APPROTOCOLIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'E1AP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (E1APPROTOCOLIESid)(self.ID)
    if out.(E1APPROTOCOLIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(E1APPROTOCOLIES_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

func (self *ProtocolIEContainerList)Unpack(stream *Stream, arg1 uint64, arg2 uint64,  data interface{}) { //Seq3
    //cloc := stream.get_current_location();
    _size := stream.get_listsize(arg2-arg1+1) + int(arg1)
    val := ProtocolIEContainer{}
    for item := 0; item <_size; item +=1 {
        val.Unpack(stream, data);
        //values = append(values, val.(interface{}))
    }
    return
}


func (self *ProtocolIEContainerList) Pack (stream *Stream, arg1 int, arg2 int, data interface{}) { //seqof 4
    _size := len(self.Items)
    stream.set_listsize(_size-arg1, uint64(arg2-arg1+1))
    val := ProtocolIEContainer{}
    for _, item := range self.Items {
        val.Pack(stream, item);
    }
}


type ProtocolIEContainerList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['IEsSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IEsSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'E1AP-PROTOCOL-IES']}
    Items map[int]*E1APPROTOCOLIES
    order []int
}

func (self *ProtocolExtensionContainer) Unpack(stream *Stream, out interface{}) { //SeqOF1
    _size := stream.get_listsize(65535) + 1
    for i := 0; i <_size; i +=1 {
        item := ProtocolExtensionField{}
        item.Unpack(stream, out)
    }
    return 
}


func (self * ProtocolExtensionContainer) Pack(stream *Stream, data interface{}) { //sqof 1
    _size := data.(E1APPROTOCOLEXTENSION_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolExtensionContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolExtensionField', 'actual-parameters': ['ExtensionSetParam']}, 'size': [(1, 'maxProtocolExtensions')], 'parameters': ['ExtensionSetParam'], 'param-types': ['E1AP-PROTOCOL-EXTENSION']}
    Items map[int]*E1APPROTOCOLEXTENSION
    order []int
}

type ProtocolExtensionField struct { // [{'type': 'E1AP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}, {'type': 'E1AP-PROTOCOL-EXTENSION.&criticality', 'table': ['ExtensionSetParam', ['id']], 'name': 'criticality'}, {'type': 'E1AP-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}]
    Id E1APPROTOCOLEXTENSIONid
    Criticality E1APPROTOCOLEXTENSIONcriticality
    ExtensionValue E1APPROTOCOLEXTENSIONExtension
}

func (self * ProtocolExtensionField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'E1AP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := E1APPROTOCOLEXTENSION{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(E1APPROTOCOLEXTENSION_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'E1AP-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}
    stream.set_location(location, _len)
    return
}

func (self * E1APPROTOCOLEXTENSION) Pack(stream *Stream, out interface{}) {
    //table {'type': 'E1AP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    key := (E1APPROTOCOLEXTENSIONid)(self.ID)
    if out.(E1APPROTOCOLEXTENSION_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(E1APPROTOCOLEXTENSION_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

func (self *PrivateIEContainer) Unpack(stream *Stream, out interface{}) { //SeqOF1
    _size := stream.get_listsize(65535) + 1
    for i := 0; i <_size; i +=1 {
        item := PrivateIEField{}
        item.Unpack(stream, out)
    }
    return 
}


func (self * PrivateIEContainer) Pack(stream *Stream, data interface{}) { //sqof 1
    _size := data.(E1APPRIVATEIES_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type PrivateIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PrivateIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(1, 'maxPrivateIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['E1AP-PRIVATE-IES']}
    Items map[int]*E1APPRIVATEIES
    order []int
}

type PrivateIEField struct { // [{'type': 'E1AP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'E1AP-PRIVATE-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'E1AP-PRIVATE-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id E1APPRIVATEIESid
    Criticality E1APPRIVATEIEScriticality
    Value E1APPRIVATEIESValue
}

func (self * PrivateIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'E1AP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := E1APPRIVATEIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(E1APPRIVATEIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'E1AP-PRIVATE-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * E1APPRIVATEIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'E1AP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (E1APPRIVATEIESid)(self.ID)
    if out.(E1APPRIVATEIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(E1APPRIVATEIES_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type E1APELEMENTARYPROCEDURE struct {//CLASS {'members': [{'type': 'OpenType', 'name': '&InitiatingMessage'}, {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, {'type': 'ProcedureCode', 'name': '&procedureCode'}, {'type': 'Criticality', 'name': '&criticality'}], 'with-members': [{'with-order': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME', 'PROCEDURE CODE', 'CRITICALITY'], 'with-type': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME']}], 'alt-type': {'INITIATING MESSAGE': {'type': 'OpenType', 'name': '&InitiatingMessage'}, 'SUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, 'UNSUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, 'PROCEDURE CODE': {'type': 'ProcedureCode', 'name': '&procedureCode'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}}, 'id-type': {'type': 'ProcedureCode', 'name': '&procedureCode'}}
    INITIATINGMESSAGE interface{}
    SUCCESSFULOUTCOME interface{}
    UNSUCCESSFULOUTCOME interface{}
    PROCEDURECODE ProcedureCode
    CRITICALITY Criticality
}
type E1APELEMENTARYPROCEDUREInitiatingMessage interface{}
type E1APELEMENTARYPROCEDURESuccessfulOutcome interface{}
type E1APELEMENTARYPROCEDUREUnsuccessfulOutcome interface{}
type E1APELEMENTARYPROCEDUREprocedureCode ProcedureCode
func (self *E1APELEMENTARYPROCEDUREprocedureCode) Pack(st *Stream) {
    ieID := ProcedureCode(*self)
    ieID.Pack(st)
}
func (self *E1APELEMENTARYPROCEDUREprocedureCode) Unpack(st *Stream) {
    val := ProcedureCode{}
    val.Unpack(st)
    *self = E1APELEMENTARYPROCEDUREprocedureCode(val)
}
type E1APELEMENTARYPROCEDUREcriticality Criticality
func (self *E1APELEMENTARYPROCEDUREcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *E1APELEMENTARYPROCEDUREcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = E1APELEMENTARYPROCEDUREcriticality(val)
}

type E1APELEMENTARYPROCEDURE_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type E1APPROTOCOLIES struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type E1APPROTOCOLIESid ProtocolIEID
func (self *E1APPROTOCOLIESid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *E1APPROTOCOLIESid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = E1APPROTOCOLIESid(val)
}
type E1APPROTOCOLIEScriticality Criticality
func (self *E1APPROTOCOLIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *E1APPROTOCOLIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = E1APPROTOCOLIEScriticality(val)
}
type E1APPROTOCOLIESValue interface{}
type E1APPROTOCOLIESpresence Presence
func (self *E1APPROTOCOLIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *E1APPROTOCOLIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = E1APPROTOCOLIESpresence(val)
}

type E1APPROTOCOLIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type E1APPROTOCOLEXTENSION struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Extension'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'EXTENSION', 'PRESENCE'], 'with-type': ['EXTENSION']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    EXTENSION interface{}
    PRESENCE Presence
}
type E1APPROTOCOLEXTENSIONid ProtocolIEID
func (self *E1APPROTOCOLEXTENSIONid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *E1APPROTOCOLEXTENSIONid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = E1APPROTOCOLEXTENSIONid(val)
}
type E1APPROTOCOLEXTENSIONcriticality Criticality
func (self *E1APPROTOCOLEXTENSIONcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *E1APPROTOCOLEXTENSIONcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = E1APPROTOCOLEXTENSIONcriticality(val)
}
type E1APPROTOCOLEXTENSIONExtension interface{}
type E1APPROTOCOLEXTENSIONpresence Presence
func (self *E1APPROTOCOLEXTENSIONpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *E1APPROTOCOLEXTENSIONpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = E1APPROTOCOLEXTENSIONpresence(val)
}

type E1APPROTOCOLEXTENSION_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type E1APPRIVATEIES struct {//CLASS {'members': [{'type': 'PrivateIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'PrivateIE-ID', 'name': '&id'}}
    ID PrivateIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type E1APPRIVATEIESid PrivateIEID
func (self *E1APPRIVATEIESid) Pack(st *Stream) {
    ieID := PrivateIEID(*self)
    ieID.Pack(st)
}
func (self *E1APPRIVATEIESid) Unpack(st *Stream) {
    val := PrivateIEID{}
    val.Unpack(st)
    *self = E1APPRIVATEIESid(val)
}
type E1APPRIVATEIEScriticality Criticality
func (self *E1APPRIVATEIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *E1APPRIVATEIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = E1APPRIVATEIEScriticality(val)
}
type E1APPRIVATEIESValue interface{}
type E1APPRIVATEIESpresence Presence
func (self *E1APPRIVATEIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *E1APPRIVATEIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = E1APPRIVATEIESpresence(val)
}

type E1APPRIVATEIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
//class E1APELEMENTARYPROCEDURES: #OBJSET1 {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'members': [{}, {}, None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_E1APELEMENTARYPROCEDURES = make(map[E1APELEMENTARYPROCEDUREprocedureCode]*E1APELEMENTARYPROCEDURE)


//class E1APELEMENTARYPROCEDURESCLASS1: #OBJSET1 {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_E1APELEMENTARYPROCEDURESCLASS1 = make(map[E1APELEMENTARYPROCEDUREprocedureCode]*E1APELEMENTARYPROCEDURE)


//class E1APELEMENTARYPROCEDURESCLASS2: #OBJSET1 {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_E1APELEMENTARYPROCEDURESCLASS2 = make(map[E1APELEMENTARYPROCEDUREprocedureCode]*E1APELEMENTARYPROCEDURE)


type ResetIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-ResetType', 'CRITICALITY': 'reject', 'TYPE': 'ResetType', 'PRESENCE': 'mandatory'}, None]}
   TransactionID  TransactionID
   Cause  Cause
   ResetType  ResetType
   list []interface{}
}
func (self *ResetIEs)createOT() interface{}{
    return nil
}
var table_ResetIEs = make(map[int]*E1APPROTOCOLIES)

var order_ResetIEs = make([]int, 3)

func (self *ResetIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.Cause
   count +=1 //self.ResetType
   return count//ObjSet
}
func (self *ResetIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 0: //Cause
        return true //self.Cause
      case 4: //ResetType
        return true //self.ResetType
   }
   return false//ObjSet
}
func (self *ResetIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 0: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 4: //ResetType
        self.ResetType.Unpack(st)
        self.list = append(self.list, &self.ResetType)
   }
}
func (self *ResetIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 0: //Cause
        self.Cause.Pack(st)
      case 4: //ResetType
        self.ResetType.Pack(st)
      default:
      break
   }
}
func init() {
table_ResetIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetIEs[0] = 57
table_ResetIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetIEs[1] = 0
table_ResetIEs[4] = &E1APPROTOCOLIES{ID:ProtocolIEID{idResetType}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ResetType{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetIEs[2] = 4
   }

type ResetTypeExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *ResetTypeExtIEs)createOT() interface{}{
    return nil
}
var table_ResetTypeExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_ResetTypeExtIEs = make([]int, 0)

type UEassociatedLogicalE1ConnectionItemRes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-UE-associatedLogicalE1-ConnectionItem', 'CRITICALITY': 'reject', 'TYPE': 'UE-associatedLogicalE1-ConnectionItem', 'PRESENCE': 'mandatory'}, None]}
   UEassociatedLogicalE1ConnectionItem  UEassociatedLogicalE1ConnectionItem
   list []interface{}
}
func (self *UEassociatedLogicalE1ConnectionItemRes)createOT() interface{}{
    return nil
}
var table_UEassociatedLogicalE1ConnectionItemRes = make(map[int]*E1APPROTOCOLIES)

var order_UEassociatedLogicalE1ConnectionItemRes = make([]int, 1)

func (self *UEassociatedLogicalE1ConnectionItemRes) GetIECount() int{
   count := 0
   count +=1 //self.UEassociatedLogicalE1ConnectionItem
   return count//ObjSet
}
func (self *UEassociatedLogicalE1ConnectionItemRes) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 5: //UEassociatedLogicalE1ConnectionItem
        return true //self.UEassociatedLogicalE1ConnectionItem
   }
   return false//ObjSet
}
func (self *UEassociatedLogicalE1ConnectionItemRes)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 5: //UEassociatedLogicalE1ConnectionItem
        self.UEassociatedLogicalE1ConnectionItem.Unpack(st)
        self.list = append(self.list, &self.UEassociatedLogicalE1ConnectionItem)
   }
}
func (self *UEassociatedLogicalE1ConnectionItemRes)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 5: //UEassociatedLogicalE1ConnectionItem
        self.UEassociatedLogicalE1ConnectionItem.Pack(st)
      default:
      break
   }
}
func init() {
table_UEassociatedLogicalE1ConnectionItemRes[5] = &E1APPROTOCOLIES{ID:ProtocolIEID{idUEassociatedLogicalE1ConnectionItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&UEassociatedLogicalE1ConnectionItem{}, PRESENCE:Presence{Presencemandatory}, }
order_UEassociatedLogicalE1ConnectionItemRes[0] = 5
   }

type ResetAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-UE-associatedLogicalE1-ConnectionListResAck', 'CRITICALITY': 'ignore', 'TYPE': 'UE-associatedLogicalE1-ConnectionListResAck', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   UEassociatedLogicalE1ConnectionListResAck  *UEassociatedLogicalE1ConnectionListResAck
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ResetAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_ResetAcknowledgeIEs = make(map[int]*E1APPROTOCOLIES)

var order_ResetAcknowledgeIEs = make([]int, 3)

func (self *ResetAcknowledgeIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.UEassociatedLogicalE1ConnectionListResAck != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 6: //UEassociatedLogicalE1ConnectionListResAck
        if self.UEassociatedLogicalE1ConnectionListResAck != nil { return true }
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 6: //UEassociatedLogicalE1ConnectionListResAck
        self.UEassociatedLogicalE1ConnectionListResAck = &UEassociatedLogicalE1ConnectionListResAck{}
        self.UEassociatedLogicalE1ConnectionListResAck.Unpack(st)
        self.list = append(self.list, self.UEassociatedLogicalE1ConnectionListResAck)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ResetAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 6: //UEassociatedLogicalE1ConnectionListResAck
        if self.UEassociatedLogicalE1ConnectionListResAck != nil {self.UEassociatedLogicalE1ConnectionListResAck.Pack(st)}
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetAcknowledgeIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetAcknowledgeIEs[0] = 57
table_ResetAcknowledgeIEs[6] = &E1APPROTOCOLIES{ID:ProtocolIEID{idUEassociatedLogicalE1ConnectionListResAck}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&UEassociatedLogicalE1ConnectionListResAck{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetAcknowledgeIEs[1] = 6
table_ResetAcknowledgeIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetAcknowledgeIEs[2] = 1
   }

type UEassociatedLogicalE1ConnectionItemResAck struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-UE-associatedLogicalE1-ConnectionItem', 'CRITICALITY': 'ignore', 'TYPE': 'UE-associatedLogicalE1-ConnectionItem', 'PRESENCE': 'mandatory'}, None]}
   UEassociatedLogicalE1ConnectionItem  UEassociatedLogicalE1ConnectionItem
   list []interface{}
}
func (self *UEassociatedLogicalE1ConnectionItemResAck)createOT() interface{}{
    return nil
}
var table_UEassociatedLogicalE1ConnectionItemResAck = make(map[int]*E1APPROTOCOLIES)

var order_UEassociatedLogicalE1ConnectionItemResAck = make([]int, 1)

func (self *UEassociatedLogicalE1ConnectionItemResAck) GetIECount() int{
   count := 0
   count +=1 //self.UEassociatedLogicalE1ConnectionItem
   return count//ObjSet
}
func (self *UEassociatedLogicalE1ConnectionItemResAck) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 5: //UEassociatedLogicalE1ConnectionItem
        return true //self.UEassociatedLogicalE1ConnectionItem
   }
   return false//ObjSet
}
func (self *UEassociatedLogicalE1ConnectionItemResAck)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 5: //UEassociatedLogicalE1ConnectionItem
        self.UEassociatedLogicalE1ConnectionItem.Unpack(st)
        self.list = append(self.list, &self.UEassociatedLogicalE1ConnectionItem)
   }
}
func (self *UEassociatedLogicalE1ConnectionItemResAck)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 5: //UEassociatedLogicalE1ConnectionItem
        self.UEassociatedLogicalE1ConnectionItem.Pack(st)
      default:
      break
   }
}
func init() {
table_UEassociatedLogicalE1ConnectionItemResAck[5] = &E1APPROTOCOLIES{ID:ProtocolIEID{idUEassociatedLogicalE1ConnectionItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&UEassociatedLogicalE1ConnectionItem{}, PRESENCE:Presence{Presencemandatory}, }
order_UEassociatedLogicalE1ConnectionItemResAck[0] = 5
   }

type ErrorIndicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'optional'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'optional'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   GNBCUCPUEE1APID  *GNBCUCPUEE1APID
   GNBCUUPUEE1APID  *GNBCUUPUEE1APID
   Cause  *Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ErrorIndicationIEs)createOT() interface{}{
    return nil
}
var table_ErrorIndicationIEs = make(map[int]*E1APPROTOCOLIES)

var order_ErrorIndicationIEs = make([]int, 5)

func (self *ErrorIndicationIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.GNBCUCPUEE1APID != nil { count += 1 }
   if self.GNBCUUPUEE1APID != nil { count += 1 }
   if self.Cause != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ErrorIndicationIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 2: //GNBCUCPUEE1APID
        if self.GNBCUCPUEE1APID != nil { return true }
      case 3: //GNBCUUPUEE1APID
        if self.GNBCUUPUEE1APID != nil { return true }
      case 0: //Cause
        if self.Cause != nil { return true }
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ErrorIndicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID = &GNBCUCPUEE1APID{}
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID = &GNBCUUPUEE1APID{}
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, self.GNBCUUPUEE1APID)
      case 0: //Cause
        self.Cause = &Cause{}
        self.Cause.Unpack(st)
        self.list = append(self.list, self.Cause)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ErrorIndicationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 2: //GNBCUCPUEE1APID
        if self.GNBCUCPUEE1APID != nil {self.GNBCUCPUEE1APID.Pack(st)}
      case 3: //GNBCUUPUEE1APID
        if self.GNBCUUPUEE1APID != nil {self.GNBCUUPUEE1APID.Pack(st)}
      case 0: //Cause
        if self.Cause != nil {self.Cause.Pack(st)}
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ErrorIndicationIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_ErrorIndicationIEs[0] = 57
table_ErrorIndicationIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[1] = 2
table_ErrorIndicationIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[2] = 3
table_ErrorIndicationIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[3] = 0
table_ErrorIndicationIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[4] = 1
   }

type GNBCUUPE1SetupRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-Name', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-UP-Name', 'PRESENCE': 'optional'}, {'ID': 'id-CNSupport', 'CRITICALITY': 'reject', 'TYPE': 'CNSupport', 'PRESENCE': 'mandatory'}, {'ID': 'id-SupportedPLMNs', 'CRITICALITY': 'reject', 'TYPE': 'SupportedPLMNs-List', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-Capacity', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-UP-Capacity', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   GNBCUUPID  GNBCUUPID
   GNBCUUPName  *GNBCUUPName
   CNSupport  CNSupport
   SupportedPLMNs  SupportedPLMNsList
   GNBCUUPCapacity  *GNBCUUPCapacity
   list []interface{}
}
func (self *GNBCUUPE1SetupRequestIEs)createOT() interface{}{
    return nil
}
var table_GNBCUUPE1SetupRequestIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUUPE1SetupRequestIEs = make([]int, 6)

func (self *GNBCUUPE1SetupRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.GNBCUUPID
   if self.GNBCUUPName != nil { count += 1 }
   count +=1 //self.CNSupport
   count +=1 //self.SupportedPLMNs
   if self.GNBCUUPCapacity != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUUPE1SetupRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 7: //GNBCUUPID
        return true //self.GNBCUUPID
      case 8: //GNBCUUPName
        if self.GNBCUUPName != nil { return true }
      case 10: //CNSupport
        return true //self.CNSupport
      case 11: //SupportedPLMNs
        return true //self.SupportedPLMNs
      case 64: //GNBCUUPCapacity
        if self.GNBCUUPCapacity != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUUPE1SetupRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 7: //GNBCUUPID
        self.GNBCUUPID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPID)
      case 8: //GNBCUUPName
        self.GNBCUUPName = &GNBCUUPName{}
        self.GNBCUUPName.Unpack(st)
        self.list = append(self.list, self.GNBCUUPName)
      case 10: //CNSupport
        self.CNSupport.Unpack(st)
        self.list = append(self.list, &self.CNSupport)
      case 11: //SupportedPLMNs
        self.SupportedPLMNs.Unpack(st)
        self.list = append(self.list, &self.SupportedPLMNs)
      case 64: //GNBCUUPCapacity
        self.GNBCUUPCapacity = &GNBCUUPCapacity{}
        self.GNBCUUPCapacity.Unpack(st)
        self.list = append(self.list, self.GNBCUUPCapacity)
   }
}
func (self *GNBCUUPE1SetupRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 7: //GNBCUUPID
        self.GNBCUUPID.Pack(st)
      case 8: //GNBCUUPName
        if self.GNBCUUPName != nil {self.GNBCUUPName.Pack(st)}
      case 10: //CNSupport
        self.CNSupport.Pack(st)
      case 11: //SupportedPLMNs
        self.SupportedPLMNs.Pack(st)
      case 64: //GNBCUUPCapacity
        if self.GNBCUUPCapacity != nil {self.GNBCUUPCapacity.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUUPE1SetupRequestIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPE1SetupRequestIEs[0] = 57
table_GNBCUUPE1SetupRequestIEs[7] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPE1SetupRequestIEs[1] = 7
table_GNBCUUPE1SetupRequestIEs[8] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPName}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUUPName{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPE1SetupRequestIEs[2] = 8
table_GNBCUUPE1SetupRequestIEs[10] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCNSupport}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNSupport{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPE1SetupRequestIEs[3] = 10
table_GNBCUUPE1SetupRequestIEs[11] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSupportedPLMNs}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SupportedPLMNsList{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPE1SetupRequestIEs[4] = 11
table_GNBCUUPE1SetupRequestIEs[64] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPCapacity}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUUPCapacity{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPE1SetupRequestIEs[5] = 64
   }

type SupportedPLMNsExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SupportedPLMNsExtIEs)createOT() interface{}{
    return nil
}
var table_SupportedPLMNsExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_SupportedPLMNsExtIEs = make([]int, 0)

type GNBCUUPE1SetupResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-CP-Name', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-CP-Name', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   GNBCUCPName  *GNBCUCPName
   list []interface{}
}
func (self *GNBCUUPE1SetupResponseIEs)createOT() interface{}{
    return nil
}
var table_GNBCUUPE1SetupResponseIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUUPE1SetupResponseIEs = make([]int, 2)

func (self *GNBCUUPE1SetupResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.GNBCUCPName != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUUPE1SetupResponseIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 9: //GNBCUCPName
        if self.GNBCUCPName != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUUPE1SetupResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 9: //GNBCUCPName
        self.GNBCUCPName = &GNBCUCPName{}
        self.GNBCUCPName.Unpack(st)
        self.list = append(self.list, self.GNBCUCPName)
   }
}
func (self *GNBCUUPE1SetupResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 9: //GNBCUCPName
        if self.GNBCUCPName != nil {self.GNBCUCPName.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUUPE1SetupResponseIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPE1SetupResponseIEs[0] = 57
table_GNBCUUPE1SetupResponseIEs[9] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPName}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUCPName{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPE1SetupResponseIEs[1] = 9
   }

type GNBCUUPE1SetupFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *GNBCUUPE1SetupFailureIEs)createOT() interface{}{
    return nil
}
var table_GNBCUUPE1SetupFailureIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUUPE1SetupFailureIEs = make([]int, 4)

func (self *GNBCUUPE1SetupFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUUPE1SetupFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 0: //Cause
        return true //self.Cause
      case 12: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUUPE1SetupFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 0: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 12: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *GNBCUUPE1SetupFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 0: //Cause
        self.Cause.Pack(st)
      case 12: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUUPE1SetupFailureIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPE1SetupFailureIEs[0] = 57
table_GNBCUUPE1SetupFailureIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPE1SetupFailureIEs[1] = 0
table_GNBCUUPE1SetupFailureIEs[12] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPE1SetupFailureIEs[2] = 12
table_GNBCUUPE1SetupFailureIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPE1SetupFailureIEs[3] = 1
   }

type GNBCUCPE1SetupRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-CP-Name', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-CP-Name', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   GNBCUCPName  *GNBCUCPName
   list []interface{}
}
func (self *GNBCUCPE1SetupRequestIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPE1SetupRequestIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUCPE1SetupRequestIEs = make([]int, 2)

func (self *GNBCUCPE1SetupRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.GNBCUCPName != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUCPE1SetupRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 9: //GNBCUCPName
        if self.GNBCUCPName != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUCPE1SetupRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 9: //GNBCUCPName
        self.GNBCUCPName = &GNBCUCPName{}
        self.GNBCUCPName.Unpack(st)
        self.list = append(self.list, self.GNBCUCPName)
   }
}
func (self *GNBCUCPE1SetupRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 9: //GNBCUCPName
        if self.GNBCUCPName != nil {self.GNBCUCPName.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUCPE1SetupRequestIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPE1SetupRequestIEs[0] = 57
table_GNBCUCPE1SetupRequestIEs[9] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPName}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUCPName{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPE1SetupRequestIEs[1] = 9
   }

type GNBCUCPE1SetupResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-Name', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-UP-Name', 'PRESENCE': 'optional'}, {'ID': 'id-CNSupport', 'CRITICALITY': 'reject', 'TYPE': 'CNSupport', 'PRESENCE': 'mandatory'}, {'ID': 'id-SupportedPLMNs', 'CRITICALITY': 'reject', 'TYPE': 'SupportedPLMNs-List', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-Capacity', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-UP-Capacity', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   GNBCUUPID  GNBCUUPID
   GNBCUUPName  *GNBCUUPName
   CNSupport  CNSupport
   SupportedPLMNs  SupportedPLMNsList
   GNBCUUPCapacity  *GNBCUUPCapacity
   list []interface{}
}
func (self *GNBCUCPE1SetupResponseIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPE1SetupResponseIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUCPE1SetupResponseIEs = make([]int, 6)

func (self *GNBCUCPE1SetupResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.GNBCUUPID
   if self.GNBCUUPName != nil { count += 1 }
   count +=1 //self.CNSupport
   count +=1 //self.SupportedPLMNs
   if self.GNBCUUPCapacity != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUCPE1SetupResponseIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 7: //GNBCUUPID
        return true //self.GNBCUUPID
      case 8: //GNBCUUPName
        if self.GNBCUUPName != nil { return true }
      case 10: //CNSupport
        return true //self.CNSupport
      case 11: //SupportedPLMNs
        return true //self.SupportedPLMNs
      case 64: //GNBCUUPCapacity
        if self.GNBCUUPCapacity != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUCPE1SetupResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 7: //GNBCUUPID
        self.GNBCUUPID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPID)
      case 8: //GNBCUUPName
        self.GNBCUUPName = &GNBCUUPName{}
        self.GNBCUUPName.Unpack(st)
        self.list = append(self.list, self.GNBCUUPName)
      case 10: //CNSupport
        self.CNSupport.Unpack(st)
        self.list = append(self.list, &self.CNSupport)
      case 11: //SupportedPLMNs
        self.SupportedPLMNs.Unpack(st)
        self.list = append(self.list, &self.SupportedPLMNs)
      case 64: //GNBCUUPCapacity
        self.GNBCUUPCapacity = &GNBCUUPCapacity{}
        self.GNBCUUPCapacity.Unpack(st)
        self.list = append(self.list, self.GNBCUUPCapacity)
   }
}
func (self *GNBCUCPE1SetupResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 7: //GNBCUUPID
        self.GNBCUUPID.Pack(st)
      case 8: //GNBCUUPName
        if self.GNBCUUPName != nil {self.GNBCUUPName.Pack(st)}
      case 10: //CNSupport
        self.CNSupport.Pack(st)
      case 11: //SupportedPLMNs
        self.SupportedPLMNs.Pack(st)
      case 64: //GNBCUUPCapacity
        if self.GNBCUUPCapacity != nil {self.GNBCUUPCapacity.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUCPE1SetupResponseIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPE1SetupResponseIEs[0] = 57
table_GNBCUCPE1SetupResponseIEs[7] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPE1SetupResponseIEs[1] = 7
table_GNBCUCPE1SetupResponseIEs[8] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPName}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUUPName{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPE1SetupResponseIEs[2] = 8
table_GNBCUCPE1SetupResponseIEs[10] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCNSupport}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNSupport{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPE1SetupResponseIEs[3] = 10
table_GNBCUCPE1SetupResponseIEs[11] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSupportedPLMNs}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SupportedPLMNsList{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPE1SetupResponseIEs[4] = 11
table_GNBCUCPE1SetupResponseIEs[64] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPCapacity}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUUPCapacity{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPE1SetupResponseIEs[5] = 64
   }

type GNBCUCPE1SetupFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *GNBCUCPE1SetupFailureIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPE1SetupFailureIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUCPE1SetupFailureIEs = make([]int, 4)

func (self *GNBCUCPE1SetupFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUCPE1SetupFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 0: //Cause
        return true //self.Cause
      case 12: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUCPE1SetupFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 0: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 12: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *GNBCUCPE1SetupFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 0: //Cause
        self.Cause.Pack(st)
      case 12: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUCPE1SetupFailureIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPE1SetupFailureIEs[0] = 57
table_GNBCUCPE1SetupFailureIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPE1SetupFailureIEs[1] = 0
table_GNBCUCPE1SetupFailureIEs[12] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPE1SetupFailureIEs[2] = 12
table_GNBCUCPE1SetupFailureIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPE1SetupFailureIEs[3] = 1
   }

type GNBCUUPConfigurationUpdateIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-Name', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-UP-Name', 'PRESENCE': 'optional'}, {'ID': 'id-SupportedPLMNs', 'CRITICALITY': 'reject', 'TYPE': 'SupportedPLMNs-List', 'PRESENCE': 'optional'}, {'ID': 'id-gNB-CU-UP-Capacity', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-UP-Capacity', 'PRESENCE': 'optional'}, {'ID': 'id-GNB-CU-UP-TNLA-To-Remove-List', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-TNLA-To-Remove-List', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   GNBCUUPID  GNBCUUPID
   GNBCUUPName  *GNBCUUPName
   SupportedPLMNs  *SupportedPLMNsList
   GNBCUUPCapacity  *GNBCUUPCapacity
   GNBCUUPTNLAToRemoveList  *GNBCUUPTNLAToRemoveList
   list []interface{}
}
func (self *GNBCUUPConfigurationUpdateIEs)createOT() interface{}{
    return nil
}
var table_GNBCUUPConfigurationUpdateIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUUPConfigurationUpdateIEs = make([]int, 6)

func (self *GNBCUUPConfigurationUpdateIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.GNBCUUPID
   if self.GNBCUUPName != nil { count += 1 }
   if self.SupportedPLMNs != nil { count += 1 }
   if self.GNBCUUPCapacity != nil { count += 1 }
   if self.GNBCUUPTNLAToRemoveList != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUUPConfigurationUpdateIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 7: //GNBCUUPID
        return true //self.GNBCUUPID
      case 8: //GNBCUUPName
        if self.GNBCUUPName != nil { return true }
      case 11: //SupportedPLMNs
        if self.SupportedPLMNs != nil { return true }
      case 64: //GNBCUUPCapacity
        if self.GNBCUUPCapacity != nil { return true }
      case 73: //GNBCUUPTNLAToRemoveList
        if self.GNBCUUPTNLAToRemoveList != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUUPConfigurationUpdateIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 7: //GNBCUUPID
        self.GNBCUUPID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPID)
      case 8: //GNBCUUPName
        self.GNBCUUPName = &GNBCUUPName{}
        self.GNBCUUPName.Unpack(st)
        self.list = append(self.list, self.GNBCUUPName)
      case 11: //SupportedPLMNs
        self.SupportedPLMNs = &SupportedPLMNsList{}
        self.SupportedPLMNs.Unpack(st)
        self.list = append(self.list, self.SupportedPLMNs)
      case 64: //GNBCUUPCapacity
        self.GNBCUUPCapacity = &GNBCUUPCapacity{}
        self.GNBCUUPCapacity.Unpack(st)
        self.list = append(self.list, self.GNBCUUPCapacity)
      case 73: //GNBCUUPTNLAToRemoveList
        self.GNBCUUPTNLAToRemoveList = &GNBCUUPTNLAToRemoveList{}
        self.GNBCUUPTNLAToRemoveList.Unpack(st)
        self.list = append(self.list, self.GNBCUUPTNLAToRemoveList)
   }
}
func (self *GNBCUUPConfigurationUpdateIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 7: //GNBCUUPID
        self.GNBCUUPID.Pack(st)
      case 8: //GNBCUUPName
        if self.GNBCUUPName != nil {self.GNBCUUPName.Pack(st)}
      case 11: //SupportedPLMNs
        if self.SupportedPLMNs != nil {self.SupportedPLMNs.Pack(st)}
      case 64: //GNBCUUPCapacity
        if self.GNBCUUPCapacity != nil {self.GNBCUUPCapacity.Pack(st)}
      case 73: //GNBCUUPTNLAToRemoveList
        if self.GNBCUUPTNLAToRemoveList != nil {self.GNBCUUPTNLAToRemoveList.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUUPConfigurationUpdateIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPConfigurationUpdateIEs[0] = 57
table_GNBCUUPConfigurationUpdateIEs[7] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPConfigurationUpdateIEs[1] = 7
table_GNBCUUPConfigurationUpdateIEs[8] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPName}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUUPName{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPConfigurationUpdateIEs[2] = 8
table_GNBCUUPConfigurationUpdateIEs[11] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSupportedPLMNs}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SupportedPLMNsList{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPConfigurationUpdateIEs[3] = 11
table_GNBCUUPConfigurationUpdateIEs[64] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPCapacity}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUUPCapacity{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPConfigurationUpdateIEs[4] = 64
table_GNBCUUPConfigurationUpdateIEs[73] = &E1APPROTOCOLIES{ID:ProtocolIEID{idGNBCUUPTNLAToRemoveList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPTNLAToRemoveList{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPConfigurationUpdateIEs[5] = 73
   }

type GNBCUUPConfigurationUpdateAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *GNBCUUPConfigurationUpdateAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_GNBCUUPConfigurationUpdateAcknowledgeIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUUPConfigurationUpdateAcknowledgeIEs = make([]int, 2)

func (self *GNBCUUPConfigurationUpdateAcknowledgeIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUUPConfigurationUpdateAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUUPConfigurationUpdateAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *GNBCUUPConfigurationUpdateAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUUPConfigurationUpdateAcknowledgeIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPConfigurationUpdateAcknowledgeIEs[0] = 57
table_GNBCUUPConfigurationUpdateAcknowledgeIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPConfigurationUpdateAcknowledgeIEs[1] = 1
   }

type GNBCUUPConfigurationUpdateFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *GNBCUUPConfigurationUpdateFailureIEs)createOT() interface{}{
    return nil
}
var table_GNBCUUPConfigurationUpdateFailureIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUUPConfigurationUpdateFailureIEs = make([]int, 4)

func (self *GNBCUUPConfigurationUpdateFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUUPConfigurationUpdateFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 0: //Cause
        return true //self.Cause
      case 12: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUUPConfigurationUpdateFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 0: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 12: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *GNBCUUPConfigurationUpdateFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 0: //Cause
        self.Cause.Pack(st)
      case 12: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUUPConfigurationUpdateFailureIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPConfigurationUpdateFailureIEs[0] = 57
table_GNBCUUPConfigurationUpdateFailureIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPConfigurationUpdateFailureIEs[1] = 0
table_GNBCUUPConfigurationUpdateFailureIEs[12] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPConfigurationUpdateFailureIEs[2] = 12
table_GNBCUUPConfigurationUpdateFailureIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUUPConfigurationUpdateFailureIEs[3] = 1
   }

type GNBCUCPConfigurationUpdateIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-CP-Name', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-UP-Name', 'PRESENCE': 'optional'}, {'ID': 'id-GNB-CU-CP-TNLA-To-Add-List', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-CP-TNLA-To-Add-List', 'PRESENCE': 'optional'}, {'ID': 'id-GNB-CU-CP-TNLA-To-Remove-List', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-CP-TNLA-To-Remove-List', 'PRESENCE': 'optional'}, {'ID': 'id-GNB-CU-CP-TNLA-To-Update-List', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-CP-TNLA-To-Update-List', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   GNBCUCPName  *GNBCUUPName
   GNBCUCPTNLAToAddList  *GNBCUCPTNLAToAddList
   GNBCUCPTNLAToRemoveList  *GNBCUCPTNLAToRemoveList
   GNBCUCPTNLAToUpdateList  *GNBCUCPTNLAToUpdateList
   list []interface{}
}
func (self *GNBCUCPConfigurationUpdateIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPConfigurationUpdateIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUCPConfigurationUpdateIEs = make([]int, 5)

func (self *GNBCUCPConfigurationUpdateIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.GNBCUCPName != nil { count += 1 }
   if self.GNBCUCPTNLAToAddList != nil { count += 1 }
   if self.GNBCUCPTNLAToRemoveList != nil { count += 1 }
   if self.GNBCUCPTNLAToUpdateList != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUCPConfigurationUpdateIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 9: //GNBCUCPName
        if self.GNBCUCPName != nil { return true }
      case 27: //GNBCUCPTNLAToAddList
        if self.GNBCUCPTNLAToAddList != nil { return true }
      case 28: //GNBCUCPTNLAToRemoveList
        if self.GNBCUCPTNLAToRemoveList != nil { return true }
      case 29: //GNBCUCPTNLAToUpdateList
        if self.GNBCUCPTNLAToUpdateList != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUCPConfigurationUpdateIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 9: //GNBCUCPName
        self.GNBCUCPName = &GNBCUUPName{}
        self.GNBCUCPName.Unpack(st)
        self.list = append(self.list, self.GNBCUCPName)
      case 27: //GNBCUCPTNLAToAddList
        self.GNBCUCPTNLAToAddList = &GNBCUCPTNLAToAddList{}
        self.GNBCUCPTNLAToAddList.Unpack(st)
        self.list = append(self.list, self.GNBCUCPTNLAToAddList)
      case 28: //GNBCUCPTNLAToRemoveList
        self.GNBCUCPTNLAToRemoveList = &GNBCUCPTNLAToRemoveList{}
        self.GNBCUCPTNLAToRemoveList.Unpack(st)
        self.list = append(self.list, self.GNBCUCPTNLAToRemoveList)
      case 29: //GNBCUCPTNLAToUpdateList
        self.GNBCUCPTNLAToUpdateList = &GNBCUCPTNLAToUpdateList{}
        self.GNBCUCPTNLAToUpdateList.Unpack(st)
        self.list = append(self.list, self.GNBCUCPTNLAToUpdateList)
   }
}
func (self *GNBCUCPConfigurationUpdateIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 9: //GNBCUCPName
        if self.GNBCUCPName != nil {self.GNBCUCPName.Pack(st)}
      case 27: //GNBCUCPTNLAToAddList
        if self.GNBCUCPTNLAToAddList != nil {self.GNBCUCPTNLAToAddList.Pack(st)}
      case 28: //GNBCUCPTNLAToRemoveList
        if self.GNBCUCPTNLAToRemoveList != nil {self.GNBCUCPTNLAToRemoveList.Pack(st)}
      case 29: //GNBCUCPTNLAToUpdateList
        if self.GNBCUCPTNLAToUpdateList != nil {self.GNBCUCPTNLAToUpdateList.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUCPConfigurationUpdateIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPConfigurationUpdateIEs[0] = 57
table_GNBCUCPConfigurationUpdateIEs[9] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPName}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUUPName{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPConfigurationUpdateIEs[1] = 9
table_GNBCUCPConfigurationUpdateIEs[27] = &E1APPROTOCOLIES{ID:ProtocolIEID{idGNBCUCPTNLAToAddList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUCPTNLAToAddList{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPConfigurationUpdateIEs[2] = 27
table_GNBCUCPConfigurationUpdateIEs[28] = &E1APPROTOCOLIES{ID:ProtocolIEID{idGNBCUCPTNLAToRemoveList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUCPTNLAToRemoveList{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPConfigurationUpdateIEs[3] = 28
table_GNBCUCPConfigurationUpdateIEs[29] = &E1APPROTOCOLIES{ID:ProtocolIEID{idGNBCUCPTNLAToUpdateList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUCPTNLAToUpdateList{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPConfigurationUpdateIEs[4] = 29
   }

type GNBCUCPConfigurationUpdateAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, {'ID': 'id-GNB-CU-CP-TNLA-Setup-List', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-CP-TNLA-Setup-List', 'PRESENCE': 'optional'}, {'ID': 'id-GNB-CU-CP-TNLA-Failed-To-Setup-List', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-CP-TNLA-Failed-To-Setup-List', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   CriticalityDiagnostics  *CriticalityDiagnostics
   GNBCUCPTNLASetupList  *GNBCUCPTNLASetupList
   GNBCUCPTNLAFailedToSetupList  *GNBCUCPTNLAFailedToSetupList
   list []interface{}
}
func (self *GNBCUCPConfigurationUpdateAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPConfigurationUpdateAcknowledgeIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUCPConfigurationUpdateAcknowledgeIEs = make([]int, 4)

func (self *GNBCUCPConfigurationUpdateAcknowledgeIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.CriticalityDiagnostics != nil { count += 1 }
   if self.GNBCUCPTNLASetupList != nil { count += 1 }
   if self.GNBCUCPTNLAFailedToSetupList != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUCPConfigurationUpdateAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
      case 30: //GNBCUCPTNLASetupList
        if self.GNBCUCPTNLASetupList != nil { return true }
      case 31: //GNBCUCPTNLAFailedToSetupList
        if self.GNBCUCPTNLAFailedToSetupList != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUCPConfigurationUpdateAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
      case 30: //GNBCUCPTNLASetupList
        self.GNBCUCPTNLASetupList = &GNBCUCPTNLASetupList{}
        self.GNBCUCPTNLASetupList.Unpack(st)
        self.list = append(self.list, self.GNBCUCPTNLASetupList)
      case 31: //GNBCUCPTNLAFailedToSetupList
        self.GNBCUCPTNLAFailedToSetupList = &GNBCUCPTNLAFailedToSetupList{}
        self.GNBCUCPTNLAFailedToSetupList.Unpack(st)
        self.list = append(self.list, self.GNBCUCPTNLAFailedToSetupList)
   }
}
func (self *GNBCUCPConfigurationUpdateAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      case 30: //GNBCUCPTNLASetupList
        if self.GNBCUCPTNLASetupList != nil {self.GNBCUCPTNLASetupList.Pack(st)}
      case 31: //GNBCUCPTNLAFailedToSetupList
        if self.GNBCUCPTNLAFailedToSetupList != nil {self.GNBCUCPTNLAFailedToSetupList.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUCPConfigurationUpdateAcknowledgeIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPConfigurationUpdateAcknowledgeIEs[0] = 57
table_GNBCUCPConfigurationUpdateAcknowledgeIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPConfigurationUpdateAcknowledgeIEs[1] = 1
table_GNBCUCPConfigurationUpdateAcknowledgeIEs[30] = &E1APPROTOCOLIES{ID:ProtocolIEID{idGNBCUCPTNLASetupList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUCPTNLASetupList{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPConfigurationUpdateAcknowledgeIEs[2] = 30
table_GNBCUCPConfigurationUpdateAcknowledgeIEs[31] = &E1APPROTOCOLIES{ID:ProtocolIEID{idGNBCUCPTNLAFailedToSetupList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUCPTNLAFailedToSetupList{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPConfigurationUpdateAcknowledgeIEs[3] = 31
   }

type GNBCUCPConfigurationUpdateFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *GNBCUCPConfigurationUpdateFailureIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPConfigurationUpdateFailureIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUCPConfigurationUpdateFailureIEs = make([]int, 4)

func (self *GNBCUCPConfigurationUpdateFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUCPConfigurationUpdateFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 0: //Cause
        return true //self.Cause
      case 12: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUCPConfigurationUpdateFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 0: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 12: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *GNBCUCPConfigurationUpdateFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 0: //Cause
        self.Cause.Pack(st)
      case 12: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUCPConfigurationUpdateFailureIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPConfigurationUpdateFailureIEs[0] = 57
table_GNBCUCPConfigurationUpdateFailureIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUCPConfigurationUpdateFailureIEs[1] = 0
table_GNBCUCPConfigurationUpdateFailureIEs[12] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPConfigurationUpdateFailureIEs[2] = 12
table_GNBCUCPConfigurationUpdateFailureIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPConfigurationUpdateFailureIEs[3] = 1
   }

type E1ReleaseRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, None]}
   TransactionID  TransactionID
   Cause  Cause
   list []interface{}
}
func (self *E1ReleaseRequestIEs)createOT() interface{}{
    return nil
}
var table_E1ReleaseRequestIEs = make(map[int]*E1APPROTOCOLIES)

var order_E1ReleaseRequestIEs = make([]int, 2)

func (self *E1ReleaseRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.Cause
   return count//ObjSet
}
func (self *E1ReleaseRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 0: //Cause
        return true //self.Cause
   }
   return false//ObjSet
}
func (self *E1ReleaseRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 0: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
   }
}
func (self *E1ReleaseRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 0: //Cause
        self.Cause.Pack(st)
      default:
      break
   }
}
func init() {
table_E1ReleaseRequestIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E1ReleaseRequestIEs[0] = 57
table_E1ReleaseRequestIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_E1ReleaseRequestIEs[1] = 0
   }

type E1ReleaseResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, None]}
   TransactionID  TransactionID
   list []interface{}
}
func (self *E1ReleaseResponseIEs)createOT() interface{}{
    return nil
}
var table_E1ReleaseResponseIEs = make(map[int]*E1APPROTOCOLIES)

var order_E1ReleaseResponseIEs = make([]int, 1)

func (self *E1ReleaseResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   return count//ObjSet
}
func (self *E1ReleaseResponseIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
   }
   return false//ObjSet
}
func (self *E1ReleaseResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
   }
}
func (self *E1ReleaseResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      default:
      break
   }
}
func init() {
table_E1ReleaseResponseIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E1ReleaseResponseIEs[0] = 57
   }

type BearerContextSetupRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-SecurityInformation', 'CRITICALITY': 'reject', 'TYPE': 'SecurityInformation', 'PRESENCE': 'mandatory'}, {'ID': 'id-UEDLAggregateMaximumBitRate', 'CRITICALITY': 'reject', 'TYPE': 'BitRate', 'PRESENCE': 'mandatory'}, {'ID': 'id-UEDLMaximumIntegrityProtectedDataRate', 'CRITICALITY': 'reject', 'TYPE': 'BitRate', 'PRESENCE': 'optional'}, {'ID': 'id-Serving-PLMN', 'CRITICALITY': 'ignore', 'TYPE': 'PLMN-Identity', 'PRESENCE': 'mandatory'}, {'ID': 'id-ActivityNotificationLevel', 'CRITICALITY': 'reject', 'TYPE': 'ActivityNotificationLevel', 'PRESENCE': 'mandatory'}, {'ID': 'id-UE-Inactivity-Timer', 'CRITICALITY': 'reject', 'TYPE': 'Inactivity-Timer', 'PRESENCE': 'optional'}, {'ID': 'id-BearerContextStatusChange', 'CRITICALITY': 'reject', 'TYPE': 'BearerContextStatusChange', 'PRESENCE': 'optional'}, {'ID': 'id-System-BearerContextSetupRequest', 'CRITICALITY': 'reject', 'TYPE': 'System-BearerContextSetupRequest', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANUEID', 'CRITICALITY': 'ignore', 'TYPE': 'RANUEID', 'PRESENCE': 'optional'}, {'ID': 'id-GNB-DU-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-DU-ID', 'PRESENCE': 'optional'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   SecurityInformation  SecurityInformation
   UEDLAggregateMaximumBitRate  BitRate
   UEDLMaximumIntegrityProtectedDataRate  *BitRate
   ServingPLMN  PLMNIdentity
   ActivityNotificationLevel  ActivityNotificationLevel
   UEInactivityTimer  *InactivityTimer
   BearerContextStatusChange  *BearerContextStatusChange
   SystemBearerContextSetupRequest  SystemBearerContextSetupRequest
   RANUEID  *RANUEID
   GNBDUID  *GNBDUID
   list []interface{}
}
func (self *BearerContextSetupRequestIEs)createOT() interface{}{
    return nil
}
var table_BearerContextSetupRequestIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextSetupRequestIEs = make([]int, 11)

func (self *BearerContextSetupRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.SecurityInformation
   count +=1 //self.UEDLAggregateMaximumBitRate
   if self.UEDLMaximumIntegrityProtectedDataRate != nil { count += 1 }
   count +=1 //self.ServingPLMN
   count +=1 //self.ActivityNotificationLevel
   if self.UEInactivityTimer != nil { count += 1 }
   if self.BearerContextStatusChange != nil { count += 1 }
   count +=1 //self.SystemBearerContextSetupRequest
   if self.RANUEID != nil { count += 1 }
   if self.GNBDUID != nil { count += 1 }
   return count//ObjSet
}
func (self *BearerContextSetupRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 13: //SecurityInformation
        return true //self.SecurityInformation
      case 14: //UEDLAggregateMaximumBitRate
        return true //self.UEDLAggregateMaximumBitRate
      case 66: //UEDLMaximumIntegrityProtectedDataRate
        if self.UEDLMaximumIntegrityProtectedDataRate != nil { return true }
      case 58: //ServingPLMN
        return true //self.ServingPLMN
      case 23: //ActivityNotificationLevel
        return true //self.ActivityNotificationLevel
      case 59: //UEInactivityTimer
        if self.UEInactivityTimer != nil { return true }
      case 17: //BearerContextStatusChange
        if self.BearerContextStatusChange != nil { return true }
      case 15: //SystemBearerContextSetupRequest
        return true //self.SystemBearerContextSetupRequest
      case 76: //RANUEID
        if self.RANUEID != nil { return true }
      case 77: //GNBDUID
        if self.GNBDUID != nil { return true }
   }
   return false//ObjSet
}
func (self *BearerContextSetupRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 13: //SecurityInformation
        self.SecurityInformation.Unpack(st)
        self.list = append(self.list, &self.SecurityInformation)
      case 14: //UEDLAggregateMaximumBitRate
        self.UEDLAggregateMaximumBitRate.Unpack(st)
        self.list = append(self.list, &self.UEDLAggregateMaximumBitRate)
      case 66: //UEDLMaximumIntegrityProtectedDataRate
        self.UEDLMaximumIntegrityProtectedDataRate = &BitRate{}
        self.UEDLMaximumIntegrityProtectedDataRate.Unpack(st)
        self.list = append(self.list, self.UEDLMaximumIntegrityProtectedDataRate)
      case 58: //ServingPLMN
        self.ServingPLMN.Unpack(st)
        self.list = append(self.list, &self.ServingPLMN)
      case 23: //ActivityNotificationLevel
        self.ActivityNotificationLevel.Unpack(st)
        self.list = append(self.list, &self.ActivityNotificationLevel)
      case 59: //UEInactivityTimer
        self.UEInactivityTimer = &InactivityTimer{}
        self.UEInactivityTimer.Unpack(st)
        self.list = append(self.list, self.UEInactivityTimer)
      case 17: //BearerContextStatusChange
        self.BearerContextStatusChange = &BearerContextStatusChange{}
        self.BearerContextStatusChange.Unpack(st)
        self.list = append(self.list, self.BearerContextStatusChange)
      case 15: //SystemBearerContextSetupRequest
        self.SystemBearerContextSetupRequest.Unpack(st)
        self.list = append(self.list, &self.SystemBearerContextSetupRequest)
      case 76: //RANUEID
        self.RANUEID = &RANUEID{}
        self.RANUEID.Unpack(st)
        self.list = append(self.list, self.RANUEID)
      case 77: //GNBDUID
        self.GNBDUID = &GNBDUID{}
        self.GNBDUID.Unpack(st)
        self.list = append(self.list, self.GNBDUID)
   }
}
func (self *BearerContextSetupRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 13: //SecurityInformation
        self.SecurityInformation.Pack(st)
      case 14: //UEDLAggregateMaximumBitRate
        self.UEDLAggregateMaximumBitRate.Pack(st)
      case 66: //UEDLMaximumIntegrityProtectedDataRate
        if self.UEDLMaximumIntegrityProtectedDataRate != nil {self.UEDLMaximumIntegrityProtectedDataRate.Pack(st)}
      case 58: //ServingPLMN
        self.ServingPLMN.Pack(st)
      case 23: //ActivityNotificationLevel
        self.ActivityNotificationLevel.Pack(st)
      case 59: //UEInactivityTimer
        if self.UEInactivityTimer != nil {self.UEInactivityTimer.Pack(st)}
      case 17: //BearerContextStatusChange
        if self.BearerContextStatusChange != nil {self.BearerContextStatusChange.Pack(st)}
      case 15: //SystemBearerContextSetupRequest
        self.SystemBearerContextSetupRequest.Pack(st)
      case 76: //RANUEID
        if self.RANUEID != nil {self.RANUEID.Pack(st)}
      case 77: //GNBDUID
        if self.GNBDUID != nil {self.GNBDUID.Pack(st)}
      default:
      break
   }
}
func init() {
table_BearerContextSetupRequestIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupRequestIEs[0] = 2
table_BearerContextSetupRequestIEs[13] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSecurityInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SecurityInformation{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupRequestIEs[1] = 13
table_BearerContextSetupRequestIEs[14] = &E1APPROTOCOLIES{ID:ProtocolIEID{idUEDLAggregateMaximumBitRate}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&BitRate{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupRequestIEs[2] = 14
table_BearerContextSetupRequestIEs[66] = &E1APPROTOCOLIES{ID:ProtocolIEID{idUEDLMaximumIntegrityProtectedDataRate}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&BitRate{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextSetupRequestIEs[3] = 66
table_BearerContextSetupRequestIEs[58] = &E1APPROTOCOLIES{ID:ProtocolIEID{idServingPLMN}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PLMNIdentity{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupRequestIEs[4] = 58
table_BearerContextSetupRequestIEs[23] = &E1APPROTOCOLIES{ID:ProtocolIEID{idActivityNotificationLevel}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ActivityNotificationLevel{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupRequestIEs[5] = 23
table_BearerContextSetupRequestIEs[59] = &E1APPROTOCOLIES{ID:ProtocolIEID{idUEInactivityTimer}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&InactivityTimer{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextSetupRequestIEs[6] = 59
table_BearerContextSetupRequestIEs[17] = &E1APPROTOCOLIES{ID:ProtocolIEID{idBearerContextStatusChange}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&BearerContextStatusChange{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextSetupRequestIEs[7] = 17
table_BearerContextSetupRequestIEs[15] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSystemBearerContextSetupRequest}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SystemBearerContextSetupRequest{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupRequestIEs[8] = 15
table_BearerContextSetupRequestIEs[76] = &E1APPROTOCOLIES{ID:ProtocolIEID{idRANUEID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RANUEID{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextSetupRequestIEs[9] = 76
table_BearerContextSetupRequestIEs[77] = &E1APPROTOCOLIES{ID:ProtocolIEID{idGNBDUID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBDUID{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextSetupRequestIEs[10] = 77
   }

type SystemBearerContextSetupRequestExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *SystemBearerContextSetupRequestExtIEs)createOT() interface{}{
    return nil
}
var table_SystemBearerContextSetupRequestExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_SystemBearerContextSetupRequestExtIEs = make([]int, 0)

type EUTRANBearerContextSetupRequest struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-DRB-To-Setup-List-EUTRAN', 'CRITICALITY': 'reject', 'TYPE': 'DRB-To-Setup-List-EUTRAN', 'PRESENCE': 'mandatory'}, None]}
   DRBToSetupListEUTRAN  DRBToSetupListEUTRAN
   list []interface{}
}
func (self *EUTRANBearerContextSetupRequest)createOT() interface{}{
    return nil
}
var table_EUTRANBearerContextSetupRequest = make(map[int]*E1APPROTOCOLIES)

var order_EUTRANBearerContextSetupRequest = make([]int, 1)

func (self *EUTRANBearerContextSetupRequest) GetIECount() int{
   count := 0
   count +=1 //self.DRBToSetupListEUTRAN
   return count//ObjSet
}
func (self *EUTRANBearerContextSetupRequest) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 32: //DRBToSetupListEUTRAN
        return true //self.DRBToSetupListEUTRAN
   }
   return false//ObjSet
}
func (self *EUTRANBearerContextSetupRequest)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 32: //DRBToSetupListEUTRAN
        self.DRBToSetupListEUTRAN.Unpack(st)
        self.list = append(self.list, &self.DRBToSetupListEUTRAN)
   }
}
func (self *EUTRANBearerContextSetupRequest)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 32: //DRBToSetupListEUTRAN
        self.DRBToSetupListEUTRAN.Pack(st)
      default:
      break
   }
}
func init() {
table_EUTRANBearerContextSetupRequest[32] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBToSetupListEUTRAN}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&DRBToSetupListEUTRAN{}, PRESENCE:Presence{Presencemandatory}, }
order_EUTRANBearerContextSetupRequest[0] = 32
   }

type NGRANBearerContextSetupRequest struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-PDU-Session-Resource-To-Setup-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-Resource-To-Setup-List', 'PRESENCE': 'mandatory'}, None]}
   PDUSessionResourceToSetupList  PDUSessionResourceToSetupList
   list []interface{}
}
func (self *NGRANBearerContextSetupRequest)createOT() interface{}{
    return nil
}
var table_NGRANBearerContextSetupRequest = make(map[int]*E1APPROTOCOLIES)

var order_NGRANBearerContextSetupRequest = make([]int, 1)

func (self *NGRANBearerContextSetupRequest) GetIECount() int{
   count := 0
   count +=1 //self.PDUSessionResourceToSetupList
   return count//ObjSet
}
func (self *NGRANBearerContextSetupRequest) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 42: //PDUSessionResourceToSetupList
        return true //self.PDUSessionResourceToSetupList
   }
   return false//ObjSet
}
func (self *NGRANBearerContextSetupRequest)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 42: //PDUSessionResourceToSetupList
        self.PDUSessionResourceToSetupList.Unpack(st)
        self.list = append(self.list, &self.PDUSessionResourceToSetupList)
   }
}
func (self *NGRANBearerContextSetupRequest)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 42: //PDUSessionResourceToSetupList
        self.PDUSessionResourceToSetupList.Pack(st)
      default:
      break
   }
}
func init() {
table_NGRANBearerContextSetupRequest[42] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceToSetupList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionResourceToSetupList{}, PRESENCE:Presence{Presencemandatory}, }
order_NGRANBearerContextSetupRequest[0] = 42
   }

type BearerContextSetupResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-System-BearerContextSetupResponse', 'CRITICALITY': 'ignore', 'TYPE': 'System-BearerContextSetupResponse', 'PRESENCE': 'mandatory'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   SystemBearerContextSetupResponse  SystemBearerContextSetupResponse
   list []interface{}
}
func (self *BearerContextSetupResponseIEs)createOT() interface{}{
    return nil
}
var table_BearerContextSetupResponseIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextSetupResponseIEs = make([]int, 3)

func (self *BearerContextSetupResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   count +=1 //self.SystemBearerContextSetupResponse
   return count//ObjSet
}
func (self *BearerContextSetupResponseIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 16: //SystemBearerContextSetupResponse
        return true //self.SystemBearerContextSetupResponse
   }
   return false//ObjSet
}
func (self *BearerContextSetupResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 16: //SystemBearerContextSetupResponse
        self.SystemBearerContextSetupResponse.Unpack(st)
        self.list = append(self.list, &self.SystemBearerContextSetupResponse)
   }
}
func (self *BearerContextSetupResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 16: //SystemBearerContextSetupResponse
        self.SystemBearerContextSetupResponse.Pack(st)
      default:
      break
   }
}
func init() {
table_BearerContextSetupResponseIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupResponseIEs[0] = 2
table_BearerContextSetupResponseIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupResponseIEs[1] = 3
table_BearerContextSetupResponseIEs[16] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSystemBearerContextSetupResponse}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&SystemBearerContextSetupResponse{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupResponseIEs[2] = 16
   }

type SystemBearerContextSetupResponseExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *SystemBearerContextSetupResponseExtIEs)createOT() interface{}{
    return nil
}
var table_SystemBearerContextSetupResponseExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_SystemBearerContextSetupResponseExtIEs = make([]int, 0)

type EUTRANBearerContextSetupResponse struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-DRB-Setup-List-EUTRAN', 'CRITICALITY': 'ignore', 'TYPE': 'DRB-Setup-List-EUTRAN', 'PRESENCE': 'mandatory'}, {'ID': 'id-DRB-Failed-List-EUTRAN', 'CRITICALITY': 'ignore', 'TYPE': 'DRB-Failed-List-EUTRAN', 'PRESENCE': 'optional'}, None]}
   DRBSetupListEUTRAN  DRBSetupListEUTRAN
   DRBFailedListEUTRAN  *DRBFailedListEUTRAN
   list []interface{}
}
func (self *EUTRANBearerContextSetupResponse)createOT() interface{}{
    return nil
}
var table_EUTRANBearerContextSetupResponse = make(map[int]*E1APPROTOCOLIES)

var order_EUTRANBearerContextSetupResponse = make([]int, 2)

func (self *EUTRANBearerContextSetupResponse) GetIECount() int{
   count := 0
   count +=1 //self.DRBSetupListEUTRAN
   if self.DRBFailedListEUTRAN != nil { count += 1 }
   return count//ObjSet
}
func (self *EUTRANBearerContextSetupResponse) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 37: //DRBSetupListEUTRAN
        return true //self.DRBSetupListEUTRAN
      case 38: //DRBFailedListEUTRAN
        if self.DRBFailedListEUTRAN != nil { return true }
   }
   return false//ObjSet
}
func (self *EUTRANBearerContextSetupResponse)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 37: //DRBSetupListEUTRAN
        self.DRBSetupListEUTRAN.Unpack(st)
        self.list = append(self.list, &self.DRBSetupListEUTRAN)
      case 38: //DRBFailedListEUTRAN
        self.DRBFailedListEUTRAN = &DRBFailedListEUTRAN{}
        self.DRBFailedListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBFailedListEUTRAN)
   }
}
func (self *EUTRANBearerContextSetupResponse)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 37: //DRBSetupListEUTRAN
        self.DRBSetupListEUTRAN.Pack(st)
      case 38: //DRBFailedListEUTRAN
        if self.DRBFailedListEUTRAN != nil {self.DRBFailedListEUTRAN.Pack(st)}
      default:
      break
   }
}
func init() {
table_EUTRANBearerContextSetupResponse[37] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBSetupListEUTRAN}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRBSetupListEUTRAN{}, PRESENCE:Presence{Presencemandatory}, }
order_EUTRANBearerContextSetupResponse[0] = 37
table_EUTRANBearerContextSetupResponse[38] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBFailedListEUTRAN}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRBFailedListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextSetupResponse[1] = 38
   }

type NGRANBearerContextSetupResponse struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-PDU-Session-Resource-Setup-List', 'CRITICALITY': 'ignore', 'TYPE': 'PDU-Session-Resource-Setup-List', 'PRESENCE': 'mandatory'}, {'ID': 'id-PDU-Session-Resource-Failed-List', 'CRITICALITY': 'ignore', 'TYPE': 'PDU-Session-Resource-Failed-List', 'PRESENCE': 'optional'}, None]}
   PDUSessionResourceSetupList  PDUSessionResourceSetupList
   PDUSessionResourceFailedList  *PDUSessionResourceFailedList
   list []interface{}
}
func (self *NGRANBearerContextSetupResponse)createOT() interface{}{
    return nil
}
var table_NGRANBearerContextSetupResponse = make(map[int]*E1APPROTOCOLIES)

var order_NGRANBearerContextSetupResponse = make([]int, 2)

func (self *NGRANBearerContextSetupResponse) GetIECount() int{
   count := 0
   count +=1 //self.PDUSessionResourceSetupList
   if self.PDUSessionResourceFailedList != nil { count += 1 }
   return count//ObjSet
}
func (self *NGRANBearerContextSetupResponse) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 46: //PDUSessionResourceSetupList
        return true //self.PDUSessionResourceSetupList
      case 47: //PDUSessionResourceFailedList
        if self.PDUSessionResourceFailedList != nil { return true }
   }
   return false//ObjSet
}
func (self *NGRANBearerContextSetupResponse)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 46: //PDUSessionResourceSetupList
        self.PDUSessionResourceSetupList.Unpack(st)
        self.list = append(self.list, &self.PDUSessionResourceSetupList)
      case 47: //PDUSessionResourceFailedList
        self.PDUSessionResourceFailedList = &PDUSessionResourceFailedList{}
        self.PDUSessionResourceFailedList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceFailedList)
   }
}
func (self *NGRANBearerContextSetupResponse)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 46: //PDUSessionResourceSetupList
        self.PDUSessionResourceSetupList.Pack(st)
      case 47: //PDUSessionResourceFailedList
        if self.PDUSessionResourceFailedList != nil {self.PDUSessionResourceFailedList.Pack(st)}
      default:
      break
   }
}
func init() {
table_NGRANBearerContextSetupResponse[46] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceSetupList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PDUSessionResourceSetupList{}, PRESENCE:Presence{Presencemandatory}, }
order_NGRANBearerContextSetupResponse[0] = 46
table_NGRANBearerContextSetupResponse[47] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceFailedList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PDUSessionResourceFailedList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextSetupResponse[1] = 47
   }

type BearerContextSetupFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'optional'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  *GNBCUUPUEE1APID
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *BearerContextSetupFailureIEs)createOT() interface{}{
    return nil
}
var table_BearerContextSetupFailureIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextSetupFailureIEs = make([]int, 4)

func (self *BearerContextSetupFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   if self.GNBCUUPUEE1APID != nil { count += 1 }
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *BearerContextSetupFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        if self.GNBCUUPUEE1APID != nil { return true }
      case 0: //Cause
        return true //self.Cause
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *BearerContextSetupFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID = &GNBCUUPUEE1APID{}
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, self.GNBCUUPUEE1APID)
      case 0: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *BearerContextSetupFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        if self.GNBCUUPUEE1APID != nil {self.GNBCUUPUEE1APID.Pack(st)}
      case 0: //Cause
        self.Cause.Pack(st)
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_BearerContextSetupFailureIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupFailureIEs[0] = 2
table_BearerContextSetupFailureIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextSetupFailureIEs[1] = 3
table_BearerContextSetupFailureIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextSetupFailureIEs[2] = 0
table_BearerContextSetupFailureIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextSetupFailureIEs[3] = 1
   }

type BearerContextModificationRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-SecurityInformation', 'CRITICALITY': 'reject', 'TYPE': 'SecurityInformation', 'PRESENCE': 'optional'}, {'ID': 'id-UEDLAggregateMaximumBitRate', 'CRITICALITY': 'reject', 'TYPE': 'BitRate', 'PRESENCE': 'optional'}, {'ID': 'id-UEDLMaximumIntegrityProtectedDataRate', 'CRITICALITY': 'reject', 'TYPE': 'BitRate', 'PRESENCE': 'optional'}, {'ID': 'id-BearerContextStatusChange', 'CRITICALITY': 'reject', 'TYPE': 'BearerContextStatusChange', 'PRESENCE': 'optional'}, {'ID': 'id-New-UL-TNL-Information-Required', 'CRITICALITY': 'reject', 'TYPE': 'New-UL-TNL-Information-Required', 'PRESENCE': 'optional'}, {'ID': 'id-UE-Inactivity-Timer', 'CRITICALITY': 'reject', 'TYPE': 'Inactivity-Timer', 'PRESENCE': 'optional'}, {'ID': 'id-DataDiscardRequired', 'CRITICALITY': 'ignore', 'TYPE': 'DataDiscardRequired', 'PRESENCE': 'optional'}, {'ID': 'id-System-BearerContextModificationRequest', 'CRITICALITY': 'reject', 'TYPE': 'System-BearerContextModificationRequest', 'PRESENCE': 'optional'}, {'ID': 'id-RANUEID', 'CRITICALITY': 'ignore', 'TYPE': 'RANUEID', 'PRESENCE': 'optional'}, {'ID': 'id-GNB-DU-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GNB-DU-ID', 'PRESENCE': 'optional'}, {'ID': 'id-ActivityNotificationLevel', 'CRITICALITY': 'ignore', 'TYPE': 'ActivityNotificationLevel', 'PRESENCE': 'optional'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   SecurityInformation  *SecurityInformation
   UEDLAggregateMaximumBitRate  *BitRate
   UEDLMaximumIntegrityProtectedDataRate  *BitRate
   BearerContextStatusChange  *BearerContextStatusChange
   NewULTNLInformationRequired  *NewULTNLInformationRequired
   UEInactivityTimer  *InactivityTimer
   DataDiscardRequired  *DataDiscardRequired
   SystemBearerContextModificationRequest  *SystemBearerContextModificationRequest
   RANUEID  *RANUEID
   GNBDUID  *GNBDUID
   ActivityNotificationLevel  *ActivityNotificationLevel
   list []interface{}
}
func (self *BearerContextModificationRequestIEs)createOT() interface{}{
    return nil
}
var table_BearerContextModificationRequestIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextModificationRequestIEs = make([]int, 13)

func (self *BearerContextModificationRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   if self.SecurityInformation != nil { count += 1 }
   if self.UEDLAggregateMaximumBitRate != nil { count += 1 }
   if self.UEDLMaximumIntegrityProtectedDataRate != nil { count += 1 }
   if self.BearerContextStatusChange != nil { count += 1 }
   if self.NewULTNLInformationRequired != nil { count += 1 }
   if self.UEInactivityTimer != nil { count += 1 }
   if self.DataDiscardRequired != nil { count += 1 }
   if self.SystemBearerContextModificationRequest != nil { count += 1 }
   if self.RANUEID != nil { count += 1 }
   if self.GNBDUID != nil { count += 1 }
   if self.ActivityNotificationLevel != nil { count += 1 }
   return count//ObjSet
}
func (self *BearerContextModificationRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 13: //SecurityInformation
        if self.SecurityInformation != nil { return true }
      case 14: //UEDLAggregateMaximumBitRate
        if self.UEDLAggregateMaximumBitRate != nil { return true }
      case 66: //UEDLMaximumIntegrityProtectedDataRate
        if self.UEDLMaximumIntegrityProtectedDataRate != nil { return true }
      case 17: //BearerContextStatusChange
        if self.BearerContextStatusChange != nil { return true }
      case 26: //NewULTNLInformationRequired
        if self.NewULTNLInformationRequired != nil { return true }
      case 59: //UEInactivityTimer
        if self.UEInactivityTimer != nil { return true }
      case 70: //DataDiscardRequired
        if self.DataDiscardRequired != nil { return true }
      case 18: //SystemBearerContextModificationRequest
        if self.SystemBearerContextModificationRequest != nil { return true }
      case 76: //RANUEID
        if self.RANUEID != nil { return true }
      case 77: //GNBDUID
        if self.GNBDUID != nil { return true }
      case 23: //ActivityNotificationLevel
        if self.ActivityNotificationLevel != nil { return true }
   }
   return false//ObjSet
}
func (self *BearerContextModificationRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 13: //SecurityInformation
        self.SecurityInformation = &SecurityInformation{}
        self.SecurityInformation.Unpack(st)
        self.list = append(self.list, self.SecurityInformation)
      case 14: //UEDLAggregateMaximumBitRate
        self.UEDLAggregateMaximumBitRate = &BitRate{}
        self.UEDLAggregateMaximumBitRate.Unpack(st)
        self.list = append(self.list, self.UEDLAggregateMaximumBitRate)
      case 66: //UEDLMaximumIntegrityProtectedDataRate
        self.UEDLMaximumIntegrityProtectedDataRate = &BitRate{}
        self.UEDLMaximumIntegrityProtectedDataRate.Unpack(st)
        self.list = append(self.list, self.UEDLMaximumIntegrityProtectedDataRate)
      case 17: //BearerContextStatusChange
        self.BearerContextStatusChange = &BearerContextStatusChange{}
        self.BearerContextStatusChange.Unpack(st)
        self.list = append(self.list, self.BearerContextStatusChange)
      case 26: //NewULTNLInformationRequired
        self.NewULTNLInformationRequired = &NewULTNLInformationRequired{}
        self.NewULTNLInformationRequired.Unpack(st)
        self.list = append(self.list, self.NewULTNLInformationRequired)
      case 59: //UEInactivityTimer
        self.UEInactivityTimer = &InactivityTimer{}
        self.UEInactivityTimer.Unpack(st)
        self.list = append(self.list, self.UEInactivityTimer)
      case 70: //DataDiscardRequired
        self.DataDiscardRequired = &DataDiscardRequired{}
        self.DataDiscardRequired.Unpack(st)
        self.list = append(self.list, self.DataDiscardRequired)
      case 18: //SystemBearerContextModificationRequest
        self.SystemBearerContextModificationRequest = &SystemBearerContextModificationRequest{}
        self.SystemBearerContextModificationRequest.Unpack(st)
        self.list = append(self.list, self.SystemBearerContextModificationRequest)
      case 76: //RANUEID
        self.RANUEID = &RANUEID{}
        self.RANUEID.Unpack(st)
        self.list = append(self.list, self.RANUEID)
      case 77: //GNBDUID
        self.GNBDUID = &GNBDUID{}
        self.GNBDUID.Unpack(st)
        self.list = append(self.list, self.GNBDUID)
      case 23: //ActivityNotificationLevel
        self.ActivityNotificationLevel = &ActivityNotificationLevel{}
        self.ActivityNotificationLevel.Unpack(st)
        self.list = append(self.list, self.ActivityNotificationLevel)
   }
}
func (self *BearerContextModificationRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 13: //SecurityInformation
        if self.SecurityInformation != nil {self.SecurityInformation.Pack(st)}
      case 14: //UEDLAggregateMaximumBitRate
        if self.UEDLAggregateMaximumBitRate != nil {self.UEDLAggregateMaximumBitRate.Pack(st)}
      case 66: //UEDLMaximumIntegrityProtectedDataRate
        if self.UEDLMaximumIntegrityProtectedDataRate != nil {self.UEDLMaximumIntegrityProtectedDataRate.Pack(st)}
      case 17: //BearerContextStatusChange
        if self.BearerContextStatusChange != nil {self.BearerContextStatusChange.Pack(st)}
      case 26: //NewULTNLInformationRequired
        if self.NewULTNLInformationRequired != nil {self.NewULTNLInformationRequired.Pack(st)}
      case 59: //UEInactivityTimer
        if self.UEInactivityTimer != nil {self.UEInactivityTimer.Pack(st)}
      case 70: //DataDiscardRequired
        if self.DataDiscardRequired != nil {self.DataDiscardRequired.Pack(st)}
      case 18: //SystemBearerContextModificationRequest
        if self.SystemBearerContextModificationRequest != nil {self.SystemBearerContextModificationRequest.Pack(st)}
      case 76: //RANUEID
        if self.RANUEID != nil {self.RANUEID.Pack(st)}
      case 77: //GNBDUID
        if self.GNBDUID != nil {self.GNBDUID.Pack(st)}
      case 23: //ActivityNotificationLevel
        if self.ActivityNotificationLevel != nil {self.ActivityNotificationLevel.Pack(st)}
      default:
      break
   }
}
func init() {
table_BearerContextModificationRequestIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationRequestIEs[0] = 2
table_BearerContextModificationRequestIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationRequestIEs[1] = 3
table_BearerContextModificationRequestIEs[13] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSecurityInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SecurityInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[2] = 13
table_BearerContextModificationRequestIEs[14] = &E1APPROTOCOLIES{ID:ProtocolIEID{idUEDLAggregateMaximumBitRate}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&BitRate{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[3] = 14
table_BearerContextModificationRequestIEs[66] = &E1APPROTOCOLIES{ID:ProtocolIEID{idUEDLMaximumIntegrityProtectedDataRate}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&BitRate{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[4] = 66
table_BearerContextModificationRequestIEs[17] = &E1APPROTOCOLIES{ID:ProtocolIEID{idBearerContextStatusChange}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&BearerContextStatusChange{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[5] = 17
table_BearerContextModificationRequestIEs[26] = &E1APPROTOCOLIES{ID:ProtocolIEID{idNewULTNLInformationRequired}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&NewULTNLInformationRequired{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[6] = 26
table_BearerContextModificationRequestIEs[59] = &E1APPROTOCOLIES{ID:ProtocolIEID{idUEInactivityTimer}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&InactivityTimer{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[7] = 59
table_BearerContextModificationRequestIEs[70] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDataDiscardRequired}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DataDiscardRequired{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[8] = 70
table_BearerContextModificationRequestIEs[18] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSystemBearerContextModificationRequest}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SystemBearerContextModificationRequest{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[9] = 18
table_BearerContextModificationRequestIEs[76] = &E1APPROTOCOLIES{ID:ProtocolIEID{idRANUEID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RANUEID{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[10] = 76
table_BearerContextModificationRequestIEs[77] = &E1APPROTOCOLIES{ID:ProtocolIEID{idGNBDUID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GNBDUID{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[11] = 77
table_BearerContextModificationRequestIEs[23] = &E1APPROTOCOLIES{ID:ProtocolIEID{idActivityNotificationLevel}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ActivityNotificationLevel{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationRequestIEs[12] = 23
   }

type SystemBearerContextModificationRequestExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *SystemBearerContextModificationRequestExtIEs)createOT() interface{}{
    return nil
}
var table_SystemBearerContextModificationRequestExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_SystemBearerContextModificationRequestExtIEs = make([]int, 0)

type EUTRANBearerContextModificationRequest struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-DRB-To-Setup-Mod-List-EUTRAN', 'CRITICALITY': 'reject', 'TYPE': 'DRB-To-Setup-Mod-List-EUTRAN', 'PRESENCE': 'optional'}, {'ID': 'id-DRB-To-Modify-List-EUTRAN', 'CRITICALITY': 'reject', 'TYPE': 'DRB-To-Modify-List-EUTRAN', 'PRESENCE': 'optional'}, {'ID': 'id-DRB-To-Remove-List-EUTRAN', 'CRITICALITY': 'reject', 'TYPE': 'DRB-To-Remove-List-EUTRAN', 'PRESENCE': 'optional'}, None]}
   DRBToSetupModListEUTRAN  *DRBToSetupModListEUTRAN
   DRBToModifyListEUTRAN  *DRBToModifyListEUTRAN
   DRBToRemoveListEUTRAN  *DRBToRemoveListEUTRAN
   list []interface{}
}
func (self *EUTRANBearerContextModificationRequest)createOT() interface{}{
    return nil
}
var table_EUTRANBearerContextModificationRequest = make(map[int]*E1APPROTOCOLIES)

var order_EUTRANBearerContextModificationRequest = make([]int, 3)

func (self *EUTRANBearerContextModificationRequest) GetIECount() int{
   count := 0
   if self.DRBToSetupModListEUTRAN != nil { count += 1 }
   if self.DRBToModifyListEUTRAN != nil { count += 1 }
   if self.DRBToRemoveListEUTRAN != nil { count += 1 }
   return count//ObjSet
}
func (self *EUTRANBearerContextModificationRequest) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 51: //DRBToSetupModListEUTRAN
        if self.DRBToSetupModListEUTRAN != nil { return true }
      case 33: //DRBToModifyListEUTRAN
        if self.DRBToModifyListEUTRAN != nil { return true }
      case 34: //DRBToRemoveListEUTRAN
        if self.DRBToRemoveListEUTRAN != nil { return true }
   }
   return false//ObjSet
}
func (self *EUTRANBearerContextModificationRequest)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 51: //DRBToSetupModListEUTRAN
        self.DRBToSetupModListEUTRAN = &DRBToSetupModListEUTRAN{}
        self.DRBToSetupModListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBToSetupModListEUTRAN)
      case 33: //DRBToModifyListEUTRAN
        self.DRBToModifyListEUTRAN = &DRBToModifyListEUTRAN{}
        self.DRBToModifyListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBToModifyListEUTRAN)
      case 34: //DRBToRemoveListEUTRAN
        self.DRBToRemoveListEUTRAN = &DRBToRemoveListEUTRAN{}
        self.DRBToRemoveListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBToRemoveListEUTRAN)
   }
}
func (self *EUTRANBearerContextModificationRequest)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 51: //DRBToSetupModListEUTRAN
        if self.DRBToSetupModListEUTRAN != nil {self.DRBToSetupModListEUTRAN.Pack(st)}
      case 33: //DRBToModifyListEUTRAN
        if self.DRBToModifyListEUTRAN != nil {self.DRBToModifyListEUTRAN.Pack(st)}
      case 34: //DRBToRemoveListEUTRAN
        if self.DRBToRemoveListEUTRAN != nil {self.DRBToRemoveListEUTRAN.Pack(st)}
      default:
      break
   }
}
func init() {
table_EUTRANBearerContextModificationRequest[51] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBToSetupModListEUTRAN}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&DRBToSetupModListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextModificationRequest[0] = 51
table_EUTRANBearerContextModificationRequest[33] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBToModifyListEUTRAN}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&DRBToModifyListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextModificationRequest[1] = 33
table_EUTRANBearerContextModificationRequest[34] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBToRemoveListEUTRAN}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&DRBToRemoveListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextModificationRequest[2] = 34
   }

type NGRANBearerContextModificationRequest struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-PDU-Session-Resource-To-Setup-Mod-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-Resource-To-Setup-Mod-List', 'PRESENCE': 'optional'}, {'ID': 'id-PDU-Session-Resource-To-Modify-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-Resource-To-Modify-List', 'PRESENCE': 'optional'}, {'ID': 'id-PDU-Session-Resource-To-Remove-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-Resource-To-Remove-List', 'PRESENCE': 'optional'}, None]}
   PDUSessionResourceToSetupModList  *PDUSessionResourceToSetupModList
   PDUSessionResourceToModifyList  *PDUSessionResourceToModifyList
   PDUSessionResourceToRemoveList  *PDUSessionResourceToRemoveList
   list []interface{}
}
func (self *NGRANBearerContextModificationRequest)createOT() interface{}{
    return nil
}
var table_NGRANBearerContextModificationRequest = make(map[int]*E1APPROTOCOLIES)

var order_NGRANBearerContextModificationRequest = make([]int, 3)

func (self *NGRANBearerContextModificationRequest) GetIECount() int{
   count := 0
   if self.PDUSessionResourceToSetupModList != nil { count += 1 }
   if self.PDUSessionResourceToModifyList != nil { count += 1 }
   if self.PDUSessionResourceToRemoveList != nil { count += 1 }
   return count//ObjSet
}
func (self *NGRANBearerContextModificationRequest) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 56: //PDUSessionResourceToSetupModList
        if self.PDUSessionResourceToSetupModList != nil { return true }
      case 43: //PDUSessionResourceToModifyList
        if self.PDUSessionResourceToModifyList != nil { return true }
      case 44: //PDUSessionResourceToRemoveList
        if self.PDUSessionResourceToRemoveList != nil { return true }
   }
   return false//ObjSet
}
func (self *NGRANBearerContextModificationRequest)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 56: //PDUSessionResourceToSetupModList
        self.PDUSessionResourceToSetupModList = &PDUSessionResourceToSetupModList{}
        self.PDUSessionResourceToSetupModList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceToSetupModList)
      case 43: //PDUSessionResourceToModifyList
        self.PDUSessionResourceToModifyList = &PDUSessionResourceToModifyList{}
        self.PDUSessionResourceToModifyList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceToModifyList)
      case 44: //PDUSessionResourceToRemoveList
        self.PDUSessionResourceToRemoveList = &PDUSessionResourceToRemoveList{}
        self.PDUSessionResourceToRemoveList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceToRemoveList)
   }
}
func (self *NGRANBearerContextModificationRequest)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 56: //PDUSessionResourceToSetupModList
        if self.PDUSessionResourceToSetupModList != nil {self.PDUSessionResourceToSetupModList.Pack(st)}
      case 43: //PDUSessionResourceToModifyList
        if self.PDUSessionResourceToModifyList != nil {self.PDUSessionResourceToModifyList.Pack(st)}
      case 44: //PDUSessionResourceToRemoveList
        if self.PDUSessionResourceToRemoveList != nil {self.PDUSessionResourceToRemoveList.Pack(st)}
      default:
      break
   }
}
func init() {
table_NGRANBearerContextModificationRequest[56] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceToSetupModList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionResourceToSetupModList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextModificationRequest[0] = 56
table_NGRANBearerContextModificationRequest[43] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceToModifyList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionResourceToModifyList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextModificationRequest[1] = 43
table_NGRANBearerContextModificationRequest[44] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceToRemoveList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionResourceToRemoveList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextModificationRequest[2] = 44
   }

type BearerContextModificationResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-System-BearerContextModificationResponse', 'CRITICALITY': 'ignore', 'TYPE': 'System-BearerContextModificationResponse', 'PRESENCE': 'optional'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   SystemBearerContextModificationResponse  *SystemBearerContextModificationResponse
   list []interface{}
}
func (self *BearerContextModificationResponseIEs)createOT() interface{}{
    return nil
}
var table_BearerContextModificationResponseIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextModificationResponseIEs = make([]int, 3)

func (self *BearerContextModificationResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   if self.SystemBearerContextModificationResponse != nil { count += 1 }
   return count//ObjSet
}
func (self *BearerContextModificationResponseIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 19: //SystemBearerContextModificationResponse
        if self.SystemBearerContextModificationResponse != nil { return true }
   }
   return false//ObjSet
}
func (self *BearerContextModificationResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 19: //SystemBearerContextModificationResponse
        self.SystemBearerContextModificationResponse = &SystemBearerContextModificationResponse{}
        self.SystemBearerContextModificationResponse.Unpack(st)
        self.list = append(self.list, self.SystemBearerContextModificationResponse)
   }
}
func (self *BearerContextModificationResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 19: //SystemBearerContextModificationResponse
        if self.SystemBearerContextModificationResponse != nil {self.SystemBearerContextModificationResponse.Pack(st)}
      default:
      break
   }
}
func init() {
table_BearerContextModificationResponseIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationResponseIEs[0] = 2
table_BearerContextModificationResponseIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationResponseIEs[1] = 3
table_BearerContextModificationResponseIEs[19] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSystemBearerContextModificationResponse}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&SystemBearerContextModificationResponse{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationResponseIEs[2] = 19
   }

type SystemBearerContextModificationResponseExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *SystemBearerContextModificationResponseExtIEs)createOT() interface{}{
    return nil
}
var table_SystemBearerContextModificationResponseExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_SystemBearerContextModificationResponseExtIEs = make([]int, 0)

type EUTRANBearerContextModificationResponse struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-DRB-Setup-Mod-List-EUTRAN', 'CRITICALITY': 'ignore', 'TYPE': 'DRB-Setup-Mod-List-EUTRAN', 'PRESENCE': 'optional'}, {'ID': 'id-DRB-Failed-Mod-List-EUTRAN', 'CRITICALITY': 'ignore', 'TYPE': 'DRB-Failed-Mod-List-EUTRAN', 'PRESENCE': 'optional'}, {'ID': 'id-DRB-Modified-List-EUTRAN', 'CRITICALITY': 'ignore', 'TYPE': 'DRB-Modified-List-EUTRAN', 'PRESENCE': 'optional'}, {'ID': 'id-DRB-Failed-To-Modify-List-EUTRAN', 'CRITICALITY': 'ignore', 'TYPE': 'DRB-Failed-To-Modify-List-EUTRAN', 'PRESENCE': 'optional'}, None]}
   DRBSetupModListEUTRAN  *DRBSetupModListEUTRAN
   DRBFailedModListEUTRAN  *DRBFailedModListEUTRAN
   DRBModifiedListEUTRAN  *DRBModifiedListEUTRAN
   DRBFailedToModifyListEUTRAN  *DRBFailedToModifyListEUTRAN
   list []interface{}
}
func (self *EUTRANBearerContextModificationResponse)createOT() interface{}{
    return nil
}
var table_EUTRANBearerContextModificationResponse = make(map[int]*E1APPROTOCOLIES)

var order_EUTRANBearerContextModificationResponse = make([]int, 4)

func (self *EUTRANBearerContextModificationResponse) GetIECount() int{
   count := 0
   if self.DRBSetupModListEUTRAN != nil { count += 1 }
   if self.DRBFailedModListEUTRAN != nil { count += 1 }
   if self.DRBModifiedListEUTRAN != nil { count += 1 }
   if self.DRBFailedToModifyListEUTRAN != nil { count += 1 }
   return count//ObjSet
}
func (self *EUTRANBearerContextModificationResponse) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 52: //DRBSetupModListEUTRAN
        if self.DRBSetupModListEUTRAN != nil { return true }
      case 53: //DRBFailedModListEUTRAN
        if self.DRBFailedModListEUTRAN != nil { return true }
      case 39: //DRBModifiedListEUTRAN
        if self.DRBModifiedListEUTRAN != nil { return true }
      case 40: //DRBFailedToModifyListEUTRAN
        if self.DRBFailedToModifyListEUTRAN != nil { return true }
   }
   return false//ObjSet
}
func (self *EUTRANBearerContextModificationResponse)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 52: //DRBSetupModListEUTRAN
        self.DRBSetupModListEUTRAN = &DRBSetupModListEUTRAN{}
        self.DRBSetupModListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBSetupModListEUTRAN)
      case 53: //DRBFailedModListEUTRAN
        self.DRBFailedModListEUTRAN = &DRBFailedModListEUTRAN{}
        self.DRBFailedModListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBFailedModListEUTRAN)
      case 39: //DRBModifiedListEUTRAN
        self.DRBModifiedListEUTRAN = &DRBModifiedListEUTRAN{}
        self.DRBModifiedListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBModifiedListEUTRAN)
      case 40: //DRBFailedToModifyListEUTRAN
        self.DRBFailedToModifyListEUTRAN = &DRBFailedToModifyListEUTRAN{}
        self.DRBFailedToModifyListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBFailedToModifyListEUTRAN)
   }
}
func (self *EUTRANBearerContextModificationResponse)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 52: //DRBSetupModListEUTRAN
        if self.DRBSetupModListEUTRAN != nil {self.DRBSetupModListEUTRAN.Pack(st)}
      case 53: //DRBFailedModListEUTRAN
        if self.DRBFailedModListEUTRAN != nil {self.DRBFailedModListEUTRAN.Pack(st)}
      case 39: //DRBModifiedListEUTRAN
        if self.DRBModifiedListEUTRAN != nil {self.DRBModifiedListEUTRAN.Pack(st)}
      case 40: //DRBFailedToModifyListEUTRAN
        if self.DRBFailedToModifyListEUTRAN != nil {self.DRBFailedToModifyListEUTRAN.Pack(st)}
      default:
      break
   }
}
func init() {
table_EUTRANBearerContextModificationResponse[52] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBSetupModListEUTRAN}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRBSetupModListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextModificationResponse[0] = 52
table_EUTRANBearerContextModificationResponse[53] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBFailedModListEUTRAN}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRBFailedModListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextModificationResponse[1] = 53
table_EUTRANBearerContextModificationResponse[39] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBModifiedListEUTRAN}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRBModifiedListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextModificationResponse[2] = 39
table_EUTRANBearerContextModificationResponse[40] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBFailedToModifyListEUTRAN}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRBFailedToModifyListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextModificationResponse[3] = 40
   }

type NGRANBearerContextModificationResponse struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-PDU-Session-Resource-Setup-Mod-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-Resource-Setup-Mod-List', 'PRESENCE': 'optional'}, {'ID': 'id-PDU-Session-Resource-Failed-Mod-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-Resource-Failed-Mod-List', 'PRESENCE': 'optional'}, {'ID': 'id-PDU-Session-Resource-Modified-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-Resource-Modified-List', 'PRESENCE': 'optional'}, {'ID': 'id-PDU-Session-Resource-Failed-To-Modify-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-Resource-Failed-To-Modify-List', 'PRESENCE': 'optional'}, None]}
   PDUSessionResourceSetupModList  *PDUSessionResourceSetupModList
   PDUSessionResourceFailedModList  *PDUSessionResourceFailedModList
   PDUSessionResourceModifiedList  *PDUSessionResourceModifiedList
   PDUSessionResourceFailedToModifyList  *PDUSessionResourceFailedToModifyList
   list []interface{}
}
func (self *NGRANBearerContextModificationResponse)createOT() interface{}{
    return nil
}
var table_NGRANBearerContextModificationResponse = make(map[int]*E1APPROTOCOLIES)

var order_NGRANBearerContextModificationResponse = make([]int, 4)

func (self *NGRANBearerContextModificationResponse) GetIECount() int{
   count := 0
   if self.PDUSessionResourceSetupModList != nil { count += 1 }
   if self.PDUSessionResourceFailedModList != nil { count += 1 }
   if self.PDUSessionResourceModifiedList != nil { count += 1 }
   if self.PDUSessionResourceFailedToModifyList != nil { count += 1 }
   return count//ObjSet
}
func (self *NGRANBearerContextModificationResponse) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 54: //PDUSessionResourceSetupModList
        if self.PDUSessionResourceSetupModList != nil { return true }
      case 55: //PDUSessionResourceFailedModList
        if self.PDUSessionResourceFailedModList != nil { return true }
      case 48: //PDUSessionResourceModifiedList
        if self.PDUSessionResourceModifiedList != nil { return true }
      case 49: //PDUSessionResourceFailedToModifyList
        if self.PDUSessionResourceFailedToModifyList != nil { return true }
   }
   return false//ObjSet
}
func (self *NGRANBearerContextModificationResponse)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 54: //PDUSessionResourceSetupModList
        self.PDUSessionResourceSetupModList = &PDUSessionResourceSetupModList{}
        self.PDUSessionResourceSetupModList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceSetupModList)
      case 55: //PDUSessionResourceFailedModList
        self.PDUSessionResourceFailedModList = &PDUSessionResourceFailedModList{}
        self.PDUSessionResourceFailedModList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceFailedModList)
      case 48: //PDUSessionResourceModifiedList
        self.PDUSessionResourceModifiedList = &PDUSessionResourceModifiedList{}
        self.PDUSessionResourceModifiedList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceModifiedList)
      case 49: //PDUSessionResourceFailedToModifyList
        self.PDUSessionResourceFailedToModifyList = &PDUSessionResourceFailedToModifyList{}
        self.PDUSessionResourceFailedToModifyList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceFailedToModifyList)
   }
}
func (self *NGRANBearerContextModificationResponse)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 54: //PDUSessionResourceSetupModList
        if self.PDUSessionResourceSetupModList != nil {self.PDUSessionResourceSetupModList.Pack(st)}
      case 55: //PDUSessionResourceFailedModList
        if self.PDUSessionResourceFailedModList != nil {self.PDUSessionResourceFailedModList.Pack(st)}
      case 48: //PDUSessionResourceModifiedList
        if self.PDUSessionResourceModifiedList != nil {self.PDUSessionResourceModifiedList.Pack(st)}
      case 49: //PDUSessionResourceFailedToModifyList
        if self.PDUSessionResourceFailedToModifyList != nil {self.PDUSessionResourceFailedToModifyList.Pack(st)}
      default:
      break
   }
}
func init() {
table_NGRANBearerContextModificationResponse[54] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceSetupModList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionResourceSetupModList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextModificationResponse[0] = 54
table_NGRANBearerContextModificationResponse[55] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceFailedModList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionResourceFailedModList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextModificationResponse[1] = 55
table_NGRANBearerContextModificationResponse[48] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceModifiedList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionResourceModifiedList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextModificationResponse[2] = 48
table_NGRANBearerContextModificationResponse[49] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceFailedToModifyList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionResourceFailedToModifyList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextModificationResponse[3] = 49
   }

type BearerContextModificationFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *BearerContextModificationFailureIEs)createOT() interface{}{
    return nil
}
var table_BearerContextModificationFailureIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextModificationFailureIEs = make([]int, 4)

func (self *BearerContextModificationFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *BearerContextModificationFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 0: //Cause
        return true //self.Cause
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *BearerContextModificationFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 0: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *BearerContextModificationFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 0: //Cause
        self.Cause.Pack(st)
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_BearerContextModificationFailureIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationFailureIEs[0] = 2
table_BearerContextModificationFailureIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationFailureIEs[1] = 3
table_BearerContextModificationFailureIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationFailureIEs[2] = 0
table_BearerContextModificationFailureIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationFailureIEs[3] = 1
   }

type BearerContextModificationRequiredIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-System-BearerContextModificationRequired', 'CRITICALITY': 'reject', 'TYPE': 'System-BearerContextModificationRequired', 'PRESENCE': 'mandatory'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   SystemBearerContextModificationRequired  SystemBearerContextModificationRequired
   list []interface{}
}
func (self *BearerContextModificationRequiredIEs)createOT() interface{}{
    return nil
}
var table_BearerContextModificationRequiredIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextModificationRequiredIEs = make([]int, 3)

func (self *BearerContextModificationRequiredIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   count +=1 //self.SystemBearerContextModificationRequired
   return count//ObjSet
}
func (self *BearerContextModificationRequiredIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 21: //SystemBearerContextModificationRequired
        return true //self.SystemBearerContextModificationRequired
   }
   return false//ObjSet
}
func (self *BearerContextModificationRequiredIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 21: //SystemBearerContextModificationRequired
        self.SystemBearerContextModificationRequired.Unpack(st)
        self.list = append(self.list, &self.SystemBearerContextModificationRequired)
   }
}
func (self *BearerContextModificationRequiredIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 21: //SystemBearerContextModificationRequired
        self.SystemBearerContextModificationRequired.Pack(st)
      default:
      break
   }
}
func init() {
table_BearerContextModificationRequiredIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationRequiredIEs[0] = 2
table_BearerContextModificationRequiredIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationRequiredIEs[1] = 3
table_BearerContextModificationRequiredIEs[21] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSystemBearerContextModificationRequired}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SystemBearerContextModificationRequired{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationRequiredIEs[2] = 21
   }

type SystemBearerContextModificationRequiredExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *SystemBearerContextModificationRequiredExtIEs)createOT() interface{}{
    return nil
}
var table_SystemBearerContextModificationRequiredExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_SystemBearerContextModificationRequiredExtIEs = make([]int, 0)

type EUTRANBearerContextModificationRequired struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-DRB-Required-To-Modify-List-EUTRAN', 'CRITICALITY': 'reject', 'TYPE': 'DRB-Required-To-Modify-List-EUTRAN', 'PRESENCE': 'optional'}, {'ID': 'id-DRB-Required-To-Remove-List-EUTRAN', 'CRITICALITY': 'reject', 'TYPE': 'DRB-Required-To-Remove-List-EUTRAN', 'PRESENCE': 'optional'}, None]}
   DRBRequiredToModifyListEUTRAN  *DRBRequiredToModifyListEUTRAN
   DRBRequiredToRemoveListEUTRAN  *DRBRequiredToRemoveListEUTRAN
   list []interface{}
}
func (self *EUTRANBearerContextModificationRequired)createOT() interface{}{
    return nil
}
var table_EUTRANBearerContextModificationRequired = make(map[int]*E1APPROTOCOLIES)

var order_EUTRANBearerContextModificationRequired = make([]int, 2)

func (self *EUTRANBearerContextModificationRequired) GetIECount() int{
   count := 0
   if self.DRBRequiredToModifyListEUTRAN != nil { count += 1 }
   if self.DRBRequiredToRemoveListEUTRAN != nil { count += 1 }
   return count//ObjSet
}
func (self *EUTRANBearerContextModificationRequired) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 35: //DRBRequiredToModifyListEUTRAN
        if self.DRBRequiredToModifyListEUTRAN != nil { return true }
      case 36: //DRBRequiredToRemoveListEUTRAN
        if self.DRBRequiredToRemoveListEUTRAN != nil { return true }
   }
   return false//ObjSet
}
func (self *EUTRANBearerContextModificationRequired)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 35: //DRBRequiredToModifyListEUTRAN
        self.DRBRequiredToModifyListEUTRAN = &DRBRequiredToModifyListEUTRAN{}
        self.DRBRequiredToModifyListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBRequiredToModifyListEUTRAN)
      case 36: //DRBRequiredToRemoveListEUTRAN
        self.DRBRequiredToRemoveListEUTRAN = &DRBRequiredToRemoveListEUTRAN{}
        self.DRBRequiredToRemoveListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBRequiredToRemoveListEUTRAN)
   }
}
func (self *EUTRANBearerContextModificationRequired)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 35: //DRBRequiredToModifyListEUTRAN
        if self.DRBRequiredToModifyListEUTRAN != nil {self.DRBRequiredToModifyListEUTRAN.Pack(st)}
      case 36: //DRBRequiredToRemoveListEUTRAN
        if self.DRBRequiredToRemoveListEUTRAN != nil {self.DRBRequiredToRemoveListEUTRAN.Pack(st)}
      default:
      break
   }
}
func init() {
table_EUTRANBearerContextModificationRequired[35] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBRequiredToModifyListEUTRAN}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&DRBRequiredToModifyListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextModificationRequired[0] = 35
table_EUTRANBearerContextModificationRequired[36] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBRequiredToRemoveListEUTRAN}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&DRBRequiredToRemoveListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextModificationRequired[1] = 36
   }

type NGRANBearerContextModificationRequired struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-PDU-Session-Resource-Required-To-Modify-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-Resource-Required-To-Modify-List', 'PRESENCE': 'optional'}, {'ID': 'id-PDU-Session-Resource-To-Remove-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-Resource-To-Remove-List', 'PRESENCE': 'optional'}, None]}
   PDUSessionResourceRequiredToModifyList  *PDUSessionResourceRequiredToModifyList
   PDUSessionResourceToRemoveList  *PDUSessionResourceToRemoveList
   list []interface{}
}
func (self *NGRANBearerContextModificationRequired)createOT() interface{}{
    return nil
}
var table_NGRANBearerContextModificationRequired = make(map[int]*E1APPROTOCOLIES)

var order_NGRANBearerContextModificationRequired = make([]int, 2)

func (self *NGRANBearerContextModificationRequired) GetIECount() int{
   count := 0
   if self.PDUSessionResourceRequiredToModifyList != nil { count += 1 }
   if self.PDUSessionResourceToRemoveList != nil { count += 1 }
   return count//ObjSet
}
func (self *NGRANBearerContextModificationRequired) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 45: //PDUSessionResourceRequiredToModifyList
        if self.PDUSessionResourceRequiredToModifyList != nil { return true }
      case 44: //PDUSessionResourceToRemoveList
        if self.PDUSessionResourceToRemoveList != nil { return true }
   }
   return false//ObjSet
}
func (self *NGRANBearerContextModificationRequired)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 45: //PDUSessionResourceRequiredToModifyList
        self.PDUSessionResourceRequiredToModifyList = &PDUSessionResourceRequiredToModifyList{}
        self.PDUSessionResourceRequiredToModifyList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceRequiredToModifyList)
      case 44: //PDUSessionResourceToRemoveList
        self.PDUSessionResourceToRemoveList = &PDUSessionResourceToRemoveList{}
        self.PDUSessionResourceToRemoveList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceToRemoveList)
   }
}
func (self *NGRANBearerContextModificationRequired)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 45: //PDUSessionResourceRequiredToModifyList
        if self.PDUSessionResourceRequiredToModifyList != nil {self.PDUSessionResourceRequiredToModifyList.Pack(st)}
      case 44: //PDUSessionResourceToRemoveList
        if self.PDUSessionResourceToRemoveList != nil {self.PDUSessionResourceToRemoveList.Pack(st)}
      default:
      break
   }
}
func init() {
table_NGRANBearerContextModificationRequired[45] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceRequiredToModifyList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionResourceRequiredToModifyList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextModificationRequired[0] = 45
table_NGRANBearerContextModificationRequired[44] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceToRemoveList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionResourceToRemoveList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextModificationRequired[1] = 44
   }

type BearerContextModificationConfirmIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-System-BearerContextModificationConfirm', 'CRITICALITY': 'ignore', 'TYPE': 'System-BearerContextModificationConfirm', 'PRESENCE': 'optional'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   SystemBearerContextModificationConfirm  *SystemBearerContextModificationConfirm
   list []interface{}
}
func (self *BearerContextModificationConfirmIEs)createOT() interface{}{
    return nil
}
var table_BearerContextModificationConfirmIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextModificationConfirmIEs = make([]int, 3)

func (self *BearerContextModificationConfirmIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   if self.SystemBearerContextModificationConfirm != nil { count += 1 }
   return count//ObjSet
}
func (self *BearerContextModificationConfirmIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 20: //SystemBearerContextModificationConfirm
        if self.SystemBearerContextModificationConfirm != nil { return true }
   }
   return false//ObjSet
}
func (self *BearerContextModificationConfirmIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 20: //SystemBearerContextModificationConfirm
        self.SystemBearerContextModificationConfirm = &SystemBearerContextModificationConfirm{}
        self.SystemBearerContextModificationConfirm.Unpack(st)
        self.list = append(self.list, self.SystemBearerContextModificationConfirm)
   }
}
func (self *BearerContextModificationConfirmIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 20: //SystemBearerContextModificationConfirm
        if self.SystemBearerContextModificationConfirm != nil {self.SystemBearerContextModificationConfirm.Pack(st)}
      default:
      break
   }
}
func init() {
table_BearerContextModificationConfirmIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationConfirmIEs[0] = 2
table_BearerContextModificationConfirmIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextModificationConfirmIEs[1] = 3
table_BearerContextModificationConfirmIEs[20] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSystemBearerContextModificationConfirm}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&SystemBearerContextModificationConfirm{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextModificationConfirmIEs[2] = 20
   }

type SystemBearerContextModificationConfirmExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *SystemBearerContextModificationConfirmExtIEs)createOT() interface{}{
    return nil
}
var table_SystemBearerContextModificationConfirmExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_SystemBearerContextModificationConfirmExtIEs = make([]int, 0)

type EUTRANBearerContextModificationConfirm struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-DRB-Confirm-Modified-List-EUTRAN', 'CRITICALITY': 'ignore', 'TYPE': 'DRB-Confirm-Modified-List-EUTRAN', 'PRESENCE': 'optional'}, None]}
   DRBConfirmModifiedListEUTRAN  *DRBConfirmModifiedListEUTRAN
   list []interface{}
}
func (self *EUTRANBearerContextModificationConfirm)createOT() interface{}{
    return nil
}
var table_EUTRANBearerContextModificationConfirm = make(map[int]*E1APPROTOCOLIES)

var order_EUTRANBearerContextModificationConfirm = make([]int, 1)

func (self *EUTRANBearerContextModificationConfirm) GetIECount() int{
   count := 0
   if self.DRBConfirmModifiedListEUTRAN != nil { count += 1 }
   return count//ObjSet
}
func (self *EUTRANBearerContextModificationConfirm) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 41: //DRBConfirmModifiedListEUTRAN
        if self.DRBConfirmModifiedListEUTRAN != nil { return true }
   }
   return false//ObjSet
}
func (self *EUTRANBearerContextModificationConfirm)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 41: //DRBConfirmModifiedListEUTRAN
        self.DRBConfirmModifiedListEUTRAN = &DRBConfirmModifiedListEUTRAN{}
        self.DRBConfirmModifiedListEUTRAN.Unpack(st)
        self.list = append(self.list, self.DRBConfirmModifiedListEUTRAN)
   }
}
func (self *EUTRANBearerContextModificationConfirm)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 41: //DRBConfirmModifiedListEUTRAN
        if self.DRBConfirmModifiedListEUTRAN != nil {self.DRBConfirmModifiedListEUTRAN.Pack(st)}
      default:
      break
   }
}
func init() {
table_EUTRANBearerContextModificationConfirm[41] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBConfirmModifiedListEUTRAN}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRBConfirmModifiedListEUTRAN{}, PRESENCE:Presence{Presenceoptional}, }
order_EUTRANBearerContextModificationConfirm[0] = 41
   }

type NGRANBearerContextModificationConfirm struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-PDU-Session-Resource-Confirm-Modified-List', 'CRITICALITY': 'ignore', 'TYPE': 'PDU-Session-Resource-Confirm-Modified-List', 'PRESENCE': 'optional'}, None]}
   PDUSessionResourceConfirmModifiedList  *PDUSessionResourceConfirmModifiedList
   list []interface{}
}
func (self *NGRANBearerContextModificationConfirm)createOT() interface{}{
    return nil
}
var table_NGRANBearerContextModificationConfirm = make(map[int]*E1APPROTOCOLIES)

var order_NGRANBearerContextModificationConfirm = make([]int, 1)

func (self *NGRANBearerContextModificationConfirm) GetIECount() int{
   count := 0
   if self.PDUSessionResourceConfirmModifiedList != nil { count += 1 }
   return count//ObjSet
}
func (self *NGRANBearerContextModificationConfirm) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 50: //PDUSessionResourceConfirmModifiedList
        if self.PDUSessionResourceConfirmModifiedList != nil { return true }
   }
   return false//ObjSet
}
func (self *NGRANBearerContextModificationConfirm)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 50: //PDUSessionResourceConfirmModifiedList
        self.PDUSessionResourceConfirmModifiedList = &PDUSessionResourceConfirmModifiedList{}
        self.PDUSessionResourceConfirmModifiedList.Unpack(st)
        self.list = append(self.list, self.PDUSessionResourceConfirmModifiedList)
   }
}
func (self *NGRANBearerContextModificationConfirm)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 50: //PDUSessionResourceConfirmModifiedList
        if self.PDUSessionResourceConfirmModifiedList != nil {self.PDUSessionResourceConfirmModifiedList.Pack(st)}
      default:
      break
   }
}
func init() {
table_NGRANBearerContextModificationConfirm[50] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceConfirmModifiedList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PDUSessionResourceConfirmModifiedList{}, PRESENCE:Presence{Presenceoptional}, }
order_NGRANBearerContextModificationConfirm[0] = 50
   }

type BearerContextReleaseCommandIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   Cause  Cause
   list []interface{}
}
func (self *BearerContextReleaseCommandIEs)createOT() interface{}{
    return nil
}
var table_BearerContextReleaseCommandIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextReleaseCommandIEs = make([]int, 3)

func (self *BearerContextReleaseCommandIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   count +=1 //self.Cause
   return count//ObjSet
}
func (self *BearerContextReleaseCommandIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 0: //Cause
        return true //self.Cause
   }
   return false//ObjSet
}
func (self *BearerContextReleaseCommandIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 0: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
   }
}
func (self *BearerContextReleaseCommandIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 0: //Cause
        self.Cause.Pack(st)
      default:
      break
   }
}
func init() {
table_BearerContextReleaseCommandIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextReleaseCommandIEs[0] = 2
table_BearerContextReleaseCommandIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextReleaseCommandIEs[1] = 3
table_BearerContextReleaseCommandIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextReleaseCommandIEs[2] = 0
   }

type BearerContextReleaseCompleteIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *BearerContextReleaseCompleteIEs)createOT() interface{}{
    return nil
}
var table_BearerContextReleaseCompleteIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextReleaseCompleteIEs = make([]int, 3)

func (self *BearerContextReleaseCompleteIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *BearerContextReleaseCompleteIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *BearerContextReleaseCompleteIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 1: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *BearerContextReleaseCompleteIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 1: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_BearerContextReleaseCompleteIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextReleaseCompleteIEs[0] = 2
table_BearerContextReleaseCompleteIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextReleaseCompleteIEs[1] = 3
table_BearerContextReleaseCompleteIEs[1] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextReleaseCompleteIEs[2] = 1
   }

type BearerContextReleaseRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-DRB-Status-List', 'CRITICALITY': 'ignore', 'TYPE': 'DRB-Status-List', 'PRESENCE': 'optional'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   DRBStatusList  *DRBStatusList
   Cause  Cause
   list []interface{}
}
func (self *BearerContextReleaseRequestIEs)createOT() interface{}{
    return nil
}
var table_BearerContextReleaseRequestIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextReleaseRequestIEs = make([]int, 4)

func (self *BearerContextReleaseRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   if self.DRBStatusList != nil { count += 1 }
   count +=1 //self.Cause
   return count//ObjSet
}
func (self *BearerContextReleaseRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 22: //DRBStatusList
        if self.DRBStatusList != nil { return true }
      case 0: //Cause
        return true //self.Cause
   }
   return false//ObjSet
}
func (self *BearerContextReleaseRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 22: //DRBStatusList
        self.DRBStatusList = &DRBStatusList{}
        self.DRBStatusList.Unpack(st)
        self.list = append(self.list, self.DRBStatusList)
      case 0: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
   }
}
func (self *BearerContextReleaseRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 22: //DRBStatusList
        if self.DRBStatusList != nil {self.DRBStatusList.Pack(st)}
      case 0: //Cause
        self.Cause.Pack(st)
      default:
      break
   }
}
func init() {
table_BearerContextReleaseRequestIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextReleaseRequestIEs[0] = 2
table_BearerContextReleaseRequestIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextReleaseRequestIEs[1] = 3
table_BearerContextReleaseRequestIEs[22] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBStatusList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRBStatusList{}, PRESENCE:Presence{Presenceoptional}, }
order_BearerContextReleaseRequestIEs[2] = 22
table_BearerContextReleaseRequestIEs[0] = &E1APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextReleaseRequestIEs[3] = 0
   }

type BearerContextInactivityNotificationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-ActivityInformation', 'CRITICALITY': 'reject', 'TYPE': 'ActivityInformation', 'PRESENCE': 'mandatory'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   ActivityInformation  ActivityInformation
   list []interface{}
}
func (self *BearerContextInactivityNotificationIEs)createOT() interface{}{
    return nil
}
var table_BearerContextInactivityNotificationIEs = make(map[int]*E1APPROTOCOLIES)

var order_BearerContextInactivityNotificationIEs = make([]int, 3)

func (self *BearerContextInactivityNotificationIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   count +=1 //self.ActivityInformation
   return count//ObjSet
}
func (self *BearerContextInactivityNotificationIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 24: //ActivityInformation
        return true //self.ActivityInformation
   }
   return false//ObjSet
}
func (self *BearerContextInactivityNotificationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 24: //ActivityInformation
        self.ActivityInformation.Unpack(st)
        self.list = append(self.list, &self.ActivityInformation)
   }
}
func (self *BearerContextInactivityNotificationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 24: //ActivityInformation
        self.ActivityInformation.Pack(st)
      default:
      break
   }
}
func init() {
table_BearerContextInactivityNotificationIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextInactivityNotificationIEs[0] = 2
table_BearerContextInactivityNotificationIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextInactivityNotificationIEs[1] = 3
table_BearerContextInactivityNotificationIEs[24] = &E1APPROTOCOLIES{ID:ProtocolIEID{idActivityInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ActivityInformation{}, PRESENCE:Presence{Presencemandatory}, }
order_BearerContextInactivityNotificationIEs[2] = 24
   }

type DLDataNotificationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-PPI', 'CRITICALITY': 'ignore', 'TYPE': 'PPI', 'PRESENCE': 'optional'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   PPI  *PPI
   list []interface{}
}
func (self *DLDataNotificationIEs)createOT() interface{}{
    return nil
}
var table_DLDataNotificationIEs = make(map[int]*E1APPROTOCOLIES)

var order_DLDataNotificationIEs = make([]int, 3)

func (self *DLDataNotificationIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   if self.PPI != nil { count += 1 }
   return count//ObjSet
}
func (self *DLDataNotificationIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 63: //PPI
        if self.PPI != nil { return true }
   }
   return false//ObjSet
}
func (self *DLDataNotificationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 63: //PPI
        self.PPI = &PPI{}
        self.PPI.Unpack(st)
        self.list = append(self.list, self.PPI)
   }
}
func (self *DLDataNotificationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 63: //PPI
        if self.PPI != nil {self.PPI.Pack(st)}
      default:
      break
   }
}
func init() {
table_DLDataNotificationIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_DLDataNotificationIEs[0] = 2
table_DLDataNotificationIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_DLDataNotificationIEs[1] = 3
table_DLDataNotificationIEs[63] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPPI}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PPI{}, PRESENCE:Presence{Presenceoptional}, }
order_DLDataNotificationIEs[2] = 63
   }

type ULDataNotificationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-PDU-Session-To-Notify-List', 'CRITICALITY': 'reject', 'TYPE': 'PDU-Session-To-Notify-List', 'PRESENCE': 'mandatory'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   PDUSessionToNotifyList  PDUSessionToNotifyList
   list []interface{}
}
func (self *ULDataNotificationIEs)createOT() interface{}{
    return nil
}
var table_ULDataNotificationIEs = make(map[int]*E1APPROTOCOLIES)

var order_ULDataNotificationIEs = make([]int, 3)

func (self *ULDataNotificationIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   count +=1 //self.PDUSessionToNotifyList
   return count//ObjSet
}
func (self *ULDataNotificationIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 67: //PDUSessionToNotifyList
        return true //self.PDUSessionToNotifyList
   }
   return false//ObjSet
}
func (self *ULDataNotificationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 67: //PDUSessionToNotifyList
        self.PDUSessionToNotifyList.Unpack(st)
        self.list = append(self.list, &self.PDUSessionToNotifyList)
   }
}
func (self *ULDataNotificationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 67: //PDUSessionToNotifyList
        self.PDUSessionToNotifyList.Pack(st)
      default:
      break
   }
}
func init() {
table_ULDataNotificationIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_ULDataNotificationIEs[0] = 2
table_ULDataNotificationIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_ULDataNotificationIEs[1] = 3
table_ULDataNotificationIEs[67] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionToNotifyList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PDUSessionToNotifyList{}, PRESENCE:Presence{Presencemandatory}, }
order_ULDataNotificationIEs[2] = 67
   }

type DataUsageReportIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Data-Usage-Report-List', 'CRITICALITY': 'ignore', 'TYPE': 'Data-Usage-Report-List', 'PRESENCE': 'mandatory'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   DataUsageReportList  DataUsageReportList
   list []interface{}
}
func (self *DataUsageReportIEs)createOT() interface{}{
    return nil
}
var table_DataUsageReportIEs = make(map[int]*E1APPROTOCOLIES)

var order_DataUsageReportIEs = make([]int, 3)

func (self *DataUsageReportIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   count +=1 //self.DataUsageReportList
   return count//ObjSet
}
func (self *DataUsageReportIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 25: //DataUsageReportList
        return true //self.DataUsageReportList
   }
   return false//ObjSet
}
func (self *DataUsageReportIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 25: //DataUsageReportList
        self.DataUsageReportList.Unpack(st)
        self.list = append(self.list, &self.DataUsageReportList)
   }
}
func (self *DataUsageReportIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 25: //DataUsageReportList
        self.DataUsageReportList.Pack(st)
      default:
      break
   }
}
func init() {
table_DataUsageReportIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_DataUsageReportIEs[0] = 2
table_DataUsageReportIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_DataUsageReportIEs[1] = 3
table_DataUsageReportIEs[25] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDataUsageReportList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DataUsageReportList{}, PRESENCE:Presence{Presencemandatory}, }
order_DataUsageReportIEs[2] = 25
   }

type GNBCUUPCounterCheckRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-System-GNB-CU-UP-CounterCheckRequest', 'CRITICALITY': 'reject', 'TYPE': 'System-GNB-CU-UP-CounterCheckRequest', 'PRESENCE': 'mandatory'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   SystemGNBCUUPCounterCheckRequest  SystemGNBCUUPCounterCheckRequest
   list []interface{}
}
func (self *GNBCUUPCounterCheckRequestIEs)createOT() interface{}{
    return nil
}
var table_GNBCUUPCounterCheckRequestIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUUPCounterCheckRequestIEs = make([]int, 3)

func (self *GNBCUUPCounterCheckRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   count +=1 //self.SystemGNBCUUPCounterCheckRequest
   return count//ObjSet
}
func (self *GNBCUUPCounterCheckRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 60: //SystemGNBCUUPCounterCheckRequest
        return true //self.SystemGNBCUUPCounterCheckRequest
   }
   return false//ObjSet
}
func (self *GNBCUUPCounterCheckRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 60: //SystemGNBCUUPCounterCheckRequest
        self.SystemGNBCUUPCounterCheckRequest.Unpack(st)
        self.list = append(self.list, &self.SystemGNBCUUPCounterCheckRequest)
   }
}
func (self *GNBCUUPCounterCheckRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 60: //SystemGNBCUUPCounterCheckRequest
        self.SystemGNBCUUPCounterCheckRequest.Pack(st)
      default:
      break
   }
}
func init() {
table_GNBCUUPCounterCheckRequestIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPCounterCheckRequestIEs[0] = 2
table_GNBCUUPCounterCheckRequestIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPCounterCheckRequestIEs[1] = 3
table_GNBCUUPCounterCheckRequestIEs[60] = &E1APPROTOCOLIES{ID:ProtocolIEID{idSystemGNBCUUPCounterCheckRequest}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SystemGNBCUUPCounterCheckRequest{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPCounterCheckRequestIEs[2] = 60
   }

type SystemGNBCUUPCounterCheckRequestExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *SystemGNBCUUPCounterCheckRequestExtIEs)createOT() interface{}{
    return nil
}
var table_SystemGNBCUUPCounterCheckRequestExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_SystemGNBCUUPCounterCheckRequestExtIEs = make([]int, 0)

type EUTRANGNBCUUPCounterCheckRequest struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-DRBs-Subject-To-Counter-Check-List-EUTRAN', 'CRITICALITY': 'ignore', 'TYPE': 'DRBs-Subject-To-Counter-Check-List-EUTRAN', 'PRESENCE': 'mandatory'}, None]}
   DRBsSubjectToCounterCheckListEUTRAN  DRBsSubjectToCounterCheckListEUTRAN
   list []interface{}
}
func (self *EUTRANGNBCUUPCounterCheckRequest)createOT() interface{}{
    return nil
}
var table_EUTRANGNBCUUPCounterCheckRequest = make(map[int]*E1APPROTOCOLIES)

var order_EUTRANGNBCUUPCounterCheckRequest = make([]int, 1)

func (self *EUTRANGNBCUUPCounterCheckRequest) GetIECount() int{
   count := 0
   count +=1 //self.DRBsSubjectToCounterCheckListEUTRAN
   return count//ObjSet
}
func (self *EUTRANGNBCUUPCounterCheckRequest) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 61: //DRBsSubjectToCounterCheckListEUTRAN
        return true //self.DRBsSubjectToCounterCheckListEUTRAN
   }
   return false//ObjSet
}
func (self *EUTRANGNBCUUPCounterCheckRequest)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 61: //DRBsSubjectToCounterCheckListEUTRAN
        self.DRBsSubjectToCounterCheckListEUTRAN.Unpack(st)
        self.list = append(self.list, &self.DRBsSubjectToCounterCheckListEUTRAN)
   }
}
func (self *EUTRANGNBCUUPCounterCheckRequest)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 61: //DRBsSubjectToCounterCheckListEUTRAN
        self.DRBsSubjectToCounterCheckListEUTRAN.Pack(st)
      default:
      break
   }
}
func init() {
table_EUTRANGNBCUUPCounterCheckRequest[61] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBsSubjectToCounterCheckListEUTRAN}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRBsSubjectToCounterCheckListEUTRAN{}, PRESENCE:Presence{Presencemandatory}, }
order_EUTRANGNBCUUPCounterCheckRequest[0] = 61
   }

type NGRANGNBCUUPCounterCheckRequest struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-DRBs-Subject-To-Counter-Check-List-NG-RAN', 'CRITICALITY': 'ignore', 'TYPE': 'DRBs-Subject-To-Counter-Check-List-NG-RAN', 'PRESENCE': 'mandatory'}, None]}
   DRBsSubjectToCounterCheckListNGRAN  DRBsSubjectToCounterCheckListNGRAN
   list []interface{}
}
func (self *NGRANGNBCUUPCounterCheckRequest)createOT() interface{}{
    return nil
}
var table_NGRANGNBCUUPCounterCheckRequest = make(map[int]*E1APPROTOCOLIES)

var order_NGRANGNBCUUPCounterCheckRequest = make([]int, 1)

func (self *NGRANGNBCUUPCounterCheckRequest) GetIECount() int{
   count := 0
   count +=1 //self.DRBsSubjectToCounterCheckListNGRAN
   return count//ObjSet
}
func (self *NGRANGNBCUUPCounterCheckRequest) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 62: //DRBsSubjectToCounterCheckListNGRAN
        return true //self.DRBsSubjectToCounterCheckListNGRAN
   }
   return false//ObjSet
}
func (self *NGRANGNBCUUPCounterCheckRequest)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 62: //DRBsSubjectToCounterCheckListNGRAN
        self.DRBsSubjectToCounterCheckListNGRAN.Unpack(st)
        self.list = append(self.list, &self.DRBsSubjectToCounterCheckListNGRAN)
   }
}
func (self *NGRANGNBCUUPCounterCheckRequest)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 62: //DRBsSubjectToCounterCheckListNGRAN
        self.DRBsSubjectToCounterCheckListNGRAN.Pack(st)
      default:
      break
   }
}
func init() {
table_NGRANGNBCUUPCounterCheckRequest[62] = &E1APPROTOCOLIES{ID:ProtocolIEID{idDRBsSubjectToCounterCheckListNGRAN}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRBsSubjectToCounterCheckListNGRAN{}, PRESENCE:Presence{Presencemandatory}, }
order_NGRANGNBCUUPCounterCheckRequest[0] = 62
   }

type GNBCUUPStatusIndicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-GNB-CU-UP-OverloadInformation', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-OverloadInformation', 'PRESENCE': 'mandatory'}, None]}
   TransactionID  TransactionID
   GNBCUUPOverloadInformation  GNBCUUPOverloadInformation
   list []interface{}
}
func (self *GNBCUUPStatusIndicationIEs)createOT() interface{}{
    return nil
}
var table_GNBCUUPStatusIndicationIEs = make(map[int]*E1APPROTOCOLIES)

var order_GNBCUUPStatusIndicationIEs = make([]int, 2)

func (self *GNBCUUPStatusIndicationIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.GNBCUUPOverloadInformation
   return count//ObjSet
}
func (self *GNBCUUPStatusIndicationIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        return true //self.TransactionID
      case 65: //GNBCUUPOverloadInformation
        return true //self.GNBCUUPOverloadInformation
   }
   return false//ObjSet
}
func (self *GNBCUUPStatusIndicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 65: //GNBCUUPOverloadInformation
        self.GNBCUUPOverloadInformation.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPOverloadInformation)
   }
}
func (self *GNBCUUPStatusIndicationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 57: //TransactionID
        self.TransactionID.Pack(st)
      case 65: //GNBCUUPOverloadInformation
        self.GNBCUUPOverloadInformation.Pack(st)
      default:
      break
   }
}
func init() {
table_GNBCUUPStatusIndicationIEs[57] = &E1APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPStatusIndicationIEs[0] = 57
table_GNBCUUPStatusIndicationIEs[65] = &E1APPROTOCOLIES{ID:ProtocolIEID{idGNBCUUPOverloadInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPOverloadInformation{}, PRESENCE:Presence{Presencemandatory}, }
order_GNBCUUPStatusIndicationIEs[1] = 65
   }

type MRDCDataUsageReportIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-gNB-CU-CP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-CP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-gNB-CU-UP-UE-E1AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'GNB-CU-UP-UE-E1AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-PDU-Session-Resource-Data-Usage-List', 'CRITICALITY': 'ignore', 'TYPE': 'PDU-Session-Resource-Data-Usage-List', 'PRESENCE': 'mandatory'}, None]}
   GNBCUCPUEE1APID  GNBCUCPUEE1APID
   GNBCUUPUEE1APID  GNBCUUPUEE1APID
   PDUSessionResourceDataUsageList  PDUSessionResourceDataUsageList
   list []interface{}
}
func (self *MRDCDataUsageReportIEs)createOT() interface{}{
    return nil
}
var table_MRDCDataUsageReportIEs = make(map[int]*E1APPROTOCOLIES)

var order_MRDCDataUsageReportIEs = make([]int, 3)

func (self *MRDCDataUsageReportIEs) GetIECount() int{
   count := 0
   count +=1 //self.GNBCUCPUEE1APID
   count +=1 //self.GNBCUUPUEE1APID
   count +=1 //self.PDUSessionResourceDataUsageList
   return count//ObjSet
}
func (self *MRDCDataUsageReportIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        return true //self.GNBCUCPUEE1APID
      case 3: //GNBCUUPUEE1APID
        return true //self.GNBCUUPUEE1APID
      case 68: //PDUSessionResourceDataUsageList
        return true //self.PDUSessionResourceDataUsageList
   }
   return false//ObjSet
}
func (self *MRDCDataUsageReportIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUCPUEE1APID)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Unpack(st)
        self.list = append(self.list, &self.GNBCUUPUEE1APID)
      case 68: //PDUSessionResourceDataUsageList
        self.PDUSessionResourceDataUsageList.Unpack(st)
        self.list = append(self.list, &self.PDUSessionResourceDataUsageList)
   }
}
func (self *MRDCDataUsageReportIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 2: //GNBCUCPUEE1APID
        self.GNBCUCPUEE1APID.Pack(st)
      case 3: //GNBCUUPUEE1APID
        self.GNBCUUPUEE1APID.Pack(st)
      case 68: //PDUSessionResourceDataUsageList
        self.PDUSessionResourceDataUsageList.Pack(st)
      default:
      break
   }
}
func init() {
table_MRDCDataUsageReportIEs[2] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUCPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUCPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MRDCDataUsageReportIEs[0] = 2
table_MRDCDataUsageReportIEs[3] = &E1APPROTOCOLIES{ID:ProtocolIEID{idgNBCUUPUEE1APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GNBCUUPUEE1APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MRDCDataUsageReportIEs[1] = 3
table_MRDCDataUsageReportIEs[68] = &E1APPROTOCOLIES{ID:ProtocolIEID{idPDUSessionResourceDataUsageList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PDUSessionResourceDataUsageList{}, PRESENCE:Presence{Presencemandatory}, }
order_MRDCDataUsageReportIEs[2] = 68
   }

type PrivateMessageIEs struct { //ObjSet 1 {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PRIVATE-IES', 'members': [None]}
   list []interface{}
}
func (self *PrivateMessageIEs)createOT() interface{}{
    return nil
}
var table_PrivateMessageIEs = make(map[int]*E1APPRIVATEIES)

var order_PrivateMessageIEs = make([]int, 0)

type ActivityInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *ActivityInformationExtIEs)createOT() interface{}{
    return nil
}
var table_ActivityInformationExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_ActivityInformationExtIEs = make([]int, 0)

type CauseExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *CauseExtIEs)createOT() interface{}{
    return nil
}
var table_CauseExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_CauseExtIEs = make([]int, 0)

type CellGroupInformationItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CellGroupInformationItemExtIEs)createOT() interface{}{
    return nil
}
var table_CellGroupInformationItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_CellGroupInformationItemExtIEs = make([]int, 0)

type CPTNLInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [{'ID': 'id-endpoint-IP-Address-and-Port', 'CRITICALITY': 'reject', 'TYPE': 'Endpoint-IP-address-and-port', 'PRESENCE': 'mandatory'}, None]}
   EndpointIPAddressandPort  EndpointIPaddressandport
   list []interface{}
}
func (self *CPTNLInformationExtIEs)createOT() interface{}{
    return nil
}
var table_CPTNLInformationExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_CPTNLInformationExtIEs = make([]int, 1)

func (self *CPTNLInformationExtIEs) GetIECount() int{
   count := 0
   count +=1 //self.EndpointIPAddressandPort
   return count//ObjSet
}
func (self *CPTNLInformationExtIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 74: //EndpointIPAddressandPort
        return true //self.EndpointIPAddressandPort
   }
   return false//ObjSet
}
func (self *CPTNLInformationExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 74: //EndpointIPAddressandPort
        self.EndpointIPAddressandPort.Unpack(st)
        self.list = append(self.list, &self.EndpointIPAddressandPort)
   }
}
func (self *CPTNLInformationExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLIESid).Value
   switch cat {
      case 74: //EndpointIPAddressandPort
        self.EndpointIPAddressandPort.Pack(st)
      default:
      break
   }
}
func init() {
table_CPTNLInformationExtIEs[74] = &E1APPROTOCOLIES{ID:ProtocolIEID{idendpointIPAddressandPort}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&EndpointIPaddressandport{}, PRESENCE:Presence{Presencemandatory}, }
order_CPTNLInformationExtIEs[0] = 74
   }

type CriticalityDiagnosticsExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsExtIEs = make([]int, 0)

type CriticalityDiagnosticsIEListExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsIEListExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsIEListExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsIEListExtIEs = make([]int, 0)

type DataForwardingInformationRequestExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DataForwardingInformationRequestExtIEs)createOT() interface{}{
    return nil
}
var table_DataForwardingInformationRequestExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DataForwardingInformationRequestExtIEs = make([]int, 0)

type DataForwardingInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DataForwardingInformationExtIEs)createOT() interface{}{
    return nil
}
var table_DataForwardingInformationExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DataForwardingInformationExtIEs = make([]int, 0)

type DataUsageperPDUSessionReportExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DataUsageperPDUSessionReportExtIEs)createOT() interface{}{
    return nil
}
var table_DataUsageperPDUSessionReportExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DataUsageperPDUSessionReportExtIEs = make([]int, 0)

type DataUsageperQoSFlowItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DataUsageperQoSFlowItemExtIEs)createOT() interface{}{
    return nil
}
var table_DataUsageperQoSFlowItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DataUsageperQoSFlowItemExtIEs = make([]int, 0)

type DataUsageReportItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DataUsageReportItemExtIEs)createOT() interface{}{
    return nil
}
var table_DataUsageReportItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DataUsageReportItemExtIEs = make([]int, 0)

type DRBActivityItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBActivityItemExtIEs)createOT() interface{}{
    return nil
}
var table_DRBActivityItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBActivityItemExtIEs = make([]int, 0)

type DRBConfirmModifiedItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBConfirmModifiedItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBConfirmModifiedItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBConfirmModifiedItemEUTRANExtIEs = make([]int, 0)

type DRBConfirmModifiedItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBConfirmModifiedItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBConfirmModifiedItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBConfirmModifiedItemNGRANExtIEs = make([]int, 0)

type DRBFailedItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBFailedItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBFailedItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBFailedItemEUTRANExtIEs = make([]int, 0)

type DRBFailedModItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBFailedModItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBFailedModItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBFailedModItemEUTRANExtIEs = make([]int, 0)

type DRBFailedItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBFailedItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBFailedItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBFailedItemNGRANExtIEs = make([]int, 0)

type DRBFailedModItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBFailedModItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBFailedModItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBFailedModItemNGRANExtIEs = make([]int, 0)

type DRBFailedToModifyItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBFailedToModifyItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBFailedToModifyItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBFailedToModifyItemEUTRANExtIEs = make([]int, 0)

type DRBFailedToModifyItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBFailedToModifyItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBFailedToModifyItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBFailedToModifyItemNGRANExtIEs = make([]int, 0)

type DRBModifiedItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBModifiedItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBModifiedItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBModifiedItemEUTRANExtIEs = make([]int, 0)

type DRBModifiedItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBModifiedItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBModifiedItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBModifiedItemNGRANExtIEs = make([]int, 0)

type DRBRequiredToModifyItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBRequiredToModifyItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBRequiredToModifyItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBRequiredToModifyItemEUTRANExtIEs = make([]int, 0)

type DRBRequiredToModifyItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBRequiredToModifyItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBRequiredToModifyItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBRequiredToModifyItemNGRANExtIEs = make([]int, 0)

type DRBSetupItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBSetupItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBSetupItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBSetupItemEUTRANExtIEs = make([]int, 0)

type DRBSetupModItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBSetupModItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBSetupModItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBSetupModItemEUTRANExtIEs = make([]int, 0)

type DRBSetupItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBSetupItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBSetupItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBSetupItemNGRANExtIEs = make([]int, 0)

type DRBSetupModItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBSetupModItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBSetupModItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBSetupModItemNGRANExtIEs = make([]int, 0)

type DRBStatusItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBStatusItemExtIEs)createOT() interface{}{
    return nil
}
var table_DRBStatusItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBStatusItemExtIEs = make([]int, 0)

type DRBsSubjectToCounterCheckItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBsSubjectToCounterCheckItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBsSubjectToCounterCheckItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBsSubjectToCounterCheckItemEUTRANExtIEs = make([]int, 0)

type DRBsSubjectToCounterCheckItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBsSubjectToCounterCheckItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBsSubjectToCounterCheckItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBsSubjectToCounterCheckItemNGRANExtIEs = make([]int, 0)

type DRBToModifyItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBToModifyItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBToModifyItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBToModifyItemEUTRANExtIEs = make([]int, 0)

type DRBToModifyItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-OldQoSFlowMap-ULendmarkerexpected', 'CRITICALITY': 'reject', 'EXTENSION': 'QoS-Flow-List', 'PRESENCE': 'optional'}, {'ID': 'id-DRB-QoS', 'CRITICALITY': 'ignore', 'EXTENSION': 'QoSFlowLevelQoSParameters', 'PRESENCE': 'optional'}, None]}
   OldQoSFlowMapULendmarkerexpected  *QoSFlowList
   DRBQoS  *QoSFlowLevelQoSParameters
   list []interface{}
}
func (self *DRBToModifyItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBToModifyItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBToModifyItemNGRANExtIEs = make([]int, 2)

func (self *DRBToModifyItemNGRANExtIEs) GetIECount() int{
   count := 0
   if self.OldQoSFlowMapULendmarkerexpected != nil { count += 1 }
   if self.DRBQoS != nil { count += 1 }
   return count//ObjSet
}
func (self *DRBToModifyItemNGRANExtIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 71: //OldQoSFlowMapULendmarkerexpected
        if self.OldQoSFlowMapULendmarkerexpected != nil { return true }
      case 72: //DRBQoS
        if self.DRBQoS != nil { return true }
   }
   return false//ObjSet
}
func (self *DRBToModifyItemNGRANExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 71: //OldQoSFlowMapULendmarkerexpected
        self.OldQoSFlowMapULendmarkerexpected = &QoSFlowList{}
        self.OldQoSFlowMapULendmarkerexpected.Unpack(st)
        self.list = append(self.list, self.OldQoSFlowMapULendmarkerexpected)
      case 72: //DRBQoS
        self.DRBQoS = &QoSFlowLevelQoSParameters{}
        self.DRBQoS.Unpack(st)
        self.list = append(self.list, self.DRBQoS)
   }
}
func (self *DRBToModifyItemNGRANExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 71: //OldQoSFlowMapULendmarkerexpected
        if self.OldQoSFlowMapULendmarkerexpected != nil {self.OldQoSFlowMapULendmarkerexpected.Pack(st)}
      case 72: //DRBQoS
        if self.DRBQoS != nil {self.DRBQoS.Pack(st)}
      default:
      break
   }
}
func init() {
table_DRBToModifyItemNGRANExtIEs[71] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idOldQoSFlowMapULendmarkerexpected}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&QoSFlowList{}, PRESENCE:Presence{Presenceoptional}, }
order_DRBToModifyItemNGRANExtIEs[0] = 71
table_DRBToModifyItemNGRANExtIEs[72] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idDRBQoS}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&QoSFlowLevelQoSParameters{}, PRESENCE:Presence{Presenceoptional}, }
order_DRBToModifyItemNGRANExtIEs[1] = 72
   }

type DRBToRemoveItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBToRemoveItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBToRemoveItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBToRemoveItemEUTRANExtIEs = make([]int, 0)

type DRBRequiredToRemoveItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBRequiredToRemoveItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBRequiredToRemoveItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBRequiredToRemoveItemEUTRANExtIEs = make([]int, 0)

type DRBToRemoveItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBToRemoveItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBToRemoveItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBToRemoveItemNGRANExtIEs = make([]int, 0)

type DRBRequiredToRemoveItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBRequiredToRemoveItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBRequiredToRemoveItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBRequiredToRemoveItemNGRANExtIEs = make([]int, 0)

type DRBToSetupItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBToSetupItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBToSetupItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBToSetupItemEUTRANExtIEs = make([]int, 0)

type DRBToSetupModItemEUTRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBToSetupModItemEUTRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBToSetupModItemEUTRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBToSetupModItemEUTRANExtIEs = make([]int, 0)

type DRBToSetupItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-DRB-QoS', 'CRITICALITY': 'ignore', 'EXTENSION': 'QoSFlowLevelQoSParameters', 'PRESENCE': 'optional'}, None]}
   DRBQoS  *QoSFlowLevelQoSParameters
   list []interface{}
}
func (self *DRBToSetupItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBToSetupItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBToSetupItemNGRANExtIEs = make([]int, 1)

func (self *DRBToSetupItemNGRANExtIEs) GetIECount() int{
   count := 0
   if self.DRBQoS != nil { count += 1 }
   return count//ObjSet
}
func (self *DRBToSetupItemNGRANExtIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 72: //DRBQoS
        if self.DRBQoS != nil { return true }
   }
   return false//ObjSet
}
func (self *DRBToSetupItemNGRANExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 72: //DRBQoS
        self.DRBQoS = &QoSFlowLevelQoSParameters{}
        self.DRBQoS.Unpack(st)
        self.list = append(self.list, self.DRBQoS)
   }
}
func (self *DRBToSetupItemNGRANExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 72: //DRBQoS
        if self.DRBQoS != nil {self.DRBQoS.Pack(st)}
      default:
      break
   }
}
func init() {
table_DRBToSetupItemNGRANExtIEs[72] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idDRBQoS}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&QoSFlowLevelQoSParameters{}, PRESENCE:Presence{Presenceoptional}, }
order_DRBToSetupItemNGRANExtIEs[0] = 72
   }

type DRBToSetupModItemNGRANExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-DRB-QoS', 'CRITICALITY': 'ignore', 'EXTENSION': 'QoSFlowLevelQoSParameters', 'PRESENCE': 'optional'}, None]}
   DRBQoS  *QoSFlowLevelQoSParameters
   list []interface{}
}
func (self *DRBToSetupModItemNGRANExtIEs)createOT() interface{}{
    return nil
}
var table_DRBToSetupModItemNGRANExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBToSetupModItemNGRANExtIEs = make([]int, 1)

func (self *DRBToSetupModItemNGRANExtIEs) GetIECount() int{
   count := 0
   if self.DRBQoS != nil { count += 1 }
   return count//ObjSet
}
func (self *DRBToSetupModItemNGRANExtIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 72: //DRBQoS
        if self.DRBQoS != nil { return true }
   }
   return false//ObjSet
}
func (self *DRBToSetupModItemNGRANExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 72: //DRBQoS
        self.DRBQoS = &QoSFlowLevelQoSParameters{}
        self.DRBQoS.Unpack(st)
        self.list = append(self.list, self.DRBQoS)
   }
}
func (self *DRBToSetupModItemNGRANExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 72: //DRBQoS
        if self.DRBQoS != nil {self.DRBQoS.Pack(st)}
      default:
      break
   }
}
func init() {
table_DRBToSetupModItemNGRANExtIEs[72] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idDRBQoS}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&QoSFlowLevelQoSParameters{}, PRESENCE:Presence{Presenceoptional}, }
order_DRBToSetupModItemNGRANExtIEs[0] = 72
   }

type DRBUsageReportItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBUsageReportItemExtIEs)createOT() interface{}{
    return nil
}
var table_DRBUsageReportItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBUsageReportItemExtIEs = make([]int, 0)

type Dynamic5QIDescriptorExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *Dynamic5QIDescriptorExtIEs)createOT() interface{}{
    return nil
}
var table_Dynamic5QIDescriptorExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_Dynamic5QIDescriptorExtIEs = make([]int, 0)

type EndpointIPaddressandportExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *EndpointIPaddressandportExtIEs)createOT() interface{}{
    return nil
}
var table_EndpointIPaddressandportExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_EndpointIPaddressandportExtIEs = make([]int, 0)

type EUTRANAllocationAndRetentionPriorityExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *EUTRANAllocationAndRetentionPriorityExtIEs)createOT() interface{}{
    return nil
}
var table_EUTRANAllocationAndRetentionPriorityExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_EUTRANAllocationAndRetentionPriorityExtIEs = make([]int, 0)

type EUTRANQoSSupportItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *EUTRANQoSSupportItemExtIEs)createOT() interface{}{
    return nil
}
var table_EUTRANQoSSupportItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_EUTRANQoSSupportItemExtIEs = make([]int, 0)

type EUTRANQoSExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *EUTRANQoSExtIEs)createOT() interface{}{
    return nil
}
var table_EUTRANQoSExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_EUTRANQoSExtIEs = make([]int, 0)

type GNBCUUPCellGroupRelatedConfigurationItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GNBCUUPCellGroupRelatedConfigurationItemExtIEs)createOT() interface{}{
    return nil
}
var table_GNBCUUPCellGroupRelatedConfigurationItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_GNBCUUPCellGroupRelatedConfigurationItemExtIEs = make([]int, 0)

type GNBCUCPTNLASetupItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GNBCUCPTNLASetupItemExtIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPTNLASetupItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_GNBCUCPTNLASetupItemExtIEs = make([]int, 0)

type GNBCUCPTNLAFailedToSetupItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GNBCUCPTNLAFailedToSetupItemExtIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPTNLAFailedToSetupItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_GNBCUCPTNLAFailedToSetupItemExtIEs = make([]int, 0)

type GNBCUCPTNLAToAddItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GNBCUCPTNLAToAddItemExtIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPTNLAToAddItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_GNBCUCPTNLAToAddItemExtIEs = make([]int, 0)

type GNBCUCPTNLAToRemoveItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-TNLAssociationTransportLayerAddressgNBCUUP', 'CRITICALITY': 'reject', 'EXTENSION': 'CP-TNL-Information', 'PRESENCE': 'optional'}, None]}
   TNLAssociationTransportLayerAddressgNBCUUP  *CPTNLInformation
   list []interface{}
}
func (self *GNBCUCPTNLAToRemoveItemExtIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPTNLAToRemoveItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_GNBCUCPTNLAToRemoveItemExtIEs = make([]int, 1)

func (self *GNBCUCPTNLAToRemoveItemExtIEs) GetIECount() int{
   count := 0
   if self.TNLAssociationTransportLayerAddressgNBCUUP != nil { count += 1 }
   return count//ObjSet
}
func (self *GNBCUCPTNLAToRemoveItemExtIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 75: //TNLAssociationTransportLayerAddressgNBCUUP
        if self.TNLAssociationTransportLayerAddressgNBCUUP != nil { return true }
   }
   return false//ObjSet
}
func (self *GNBCUCPTNLAToRemoveItemExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 75: //TNLAssociationTransportLayerAddressgNBCUUP
        self.TNLAssociationTransportLayerAddressgNBCUUP = &CPTNLInformation{}
        self.TNLAssociationTransportLayerAddressgNBCUUP.Unpack(st)
        self.list = append(self.list, self.TNLAssociationTransportLayerAddressgNBCUUP)
   }
}
func (self *GNBCUCPTNLAToRemoveItemExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 75: //TNLAssociationTransportLayerAddressgNBCUUP
        if self.TNLAssociationTransportLayerAddressgNBCUUP != nil {self.TNLAssociationTransportLayerAddressgNBCUUP.Pack(st)}
      default:
      break
   }
}
func init() {
table_GNBCUCPTNLAToRemoveItemExtIEs[75] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idTNLAssociationTransportLayerAddressgNBCUUP}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&CPTNLInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_GNBCUCPTNLAToRemoveItemExtIEs[0] = 75
   }

type GNBCUCPTNLAToUpdateItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GNBCUCPTNLAToUpdateItemExtIEs)createOT() interface{}{
    return nil
}
var table_GNBCUCPTNLAToUpdateItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_GNBCUCPTNLAToUpdateItemExtIEs = make([]int, 0)

type GNBCUUPTNLAToRemoveItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GNBCUUPTNLAToRemoveItemExtIEs)createOT() interface{}{
    return nil
}
var table_GNBCUUPTNLAToRemoveItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_GNBCUUPTNLAToRemoveItemExtIEs = make([]int, 0)

type GBRQosInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GBRQosInformationExtIEs)createOT() interface{}{
    return nil
}
var table_GBRQosInformationExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_GBRQosInformationExtIEs = make([]int, 0)

type GBRQosFlowInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GBRQosFlowInformationExtIEs)createOT() interface{}{
    return nil
}
var table_GBRQosFlowInformationExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_GBRQosFlowInformationExtIEs = make([]int, 0)

type GTPTunnelExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GTPTunnelExtIEs)createOT() interface{}{
    return nil
}
var table_GTPTunnelExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_GTPTunnelExtIEs = make([]int, 0)

type MaximumIPdatarateExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MaximumIPdatarateExtIEs)createOT() interface{}{
    return nil
}
var table_MaximumIPdatarateExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_MaximumIPdatarateExtIEs = make([]int, 0)

type MRDCDataUsageReportItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MRDCDataUsageReportItemExtIEs)createOT() interface{}{
    return nil
}
var table_MRDCDataUsageReportItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_MRDCDataUsageReportItemExtIEs = make([]int, 0)

type MRDCUsageInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MRDCUsageInformationExtIEs)createOT() interface{}{
    return nil
}
var table_MRDCUsageInformationExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_MRDCUsageInformationExtIEs = make([]int, 0)

type NGRANAllocationAndRetentionPriorityExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *NGRANAllocationAndRetentionPriorityExtIEs)createOT() interface{}{
    return nil
}
var table_NGRANAllocationAndRetentionPriorityExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_NGRANAllocationAndRetentionPriorityExtIEs = make([]int, 0)

type NGRANQoSSupportItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *NGRANQoSSupportItemExtIEs)createOT() interface{}{
    return nil
}
var table_NGRANQoSSupportItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_NGRANQoSSupportItemExtIEs = make([]int, 0)

type NonDynamic5QIDescriptorExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *NonDynamic5QIDescriptorExtIEs)createOT() interface{}{
    return nil
}
var table_NonDynamic5QIDescriptorExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_NonDynamic5QIDescriptorExtIEs = make([]int, 0)

type NRCGIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *NRCGIExtIEs)createOT() interface{}{
    return nil
}
var table_NRCGIExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_NRCGIExtIEs = make([]int, 0)

type NRCGISupportItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *NRCGISupportItemExtIEs)createOT() interface{}{
    return nil
}
var table_NRCGISupportItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_NRCGISupportItemExtIEs = make([]int, 0)

type PacketErrorRateExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PacketErrorRateExtIEs)createOT() interface{}{
    return nil
}
var table_PacketErrorRateExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PacketErrorRateExtIEs = make([]int, 0)

type PDCPConfigurationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDCPConfigurationExtIEs)createOT() interface{}{
    return nil
}
var table_PDCPConfigurationExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDCPConfigurationExtIEs = make([]int, 0)

type PDCPCountExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDCPCountExtIEs)createOT() interface{}{
    return nil
}
var table_PDCPCountExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDCPCountExtIEs = make([]int, 0)

type PDUSessionResourceDataUsageItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionResourceDataUsageItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceDataUsageItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceDataUsageItemExtIEs = make([]int, 0)

type DRBsSubjectToStatusTransferItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBsSubjectToStatusTransferItemExtIEs)createOT() interface{}{
    return nil
}
var table_DRBsSubjectToStatusTransferItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBsSubjectToStatusTransferItemExtIEs = make([]int, 0)

type DRBBStatusTransferExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DRBBStatusTransferExtIEs)createOT() interface{}{
    return nil
}
var table_DRBBStatusTransferExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_DRBBStatusTransferExtIEs = make([]int, 0)

type PDUSessionResourceActivityItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionResourceActivityItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceActivityItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceActivityItemExtIEs = make([]int, 0)

type PDUSessionResourceConfirmModifiedItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionResourceConfirmModifiedItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceConfirmModifiedItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceConfirmModifiedItemExtIEs = make([]int, 0)

type PDUSessionResourceFailedItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionResourceFailedItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceFailedItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceFailedItemExtIEs = make([]int, 0)

type PDUSessionResourceFailedModItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionResourceFailedModItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceFailedModItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceFailedModItemExtIEs = make([]int, 0)

type PDUSessionResourceFailedToModifyItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionResourceFailedToModifyItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceFailedToModifyItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceFailedToModifyItemExtIEs = make([]int, 0)

type PDUSessionResourceModifiedItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionResourceModifiedItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceModifiedItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceModifiedItemExtIEs = make([]int, 0)

type PDUSessionResourceRequiredToModifyItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionResourceRequiredToModifyItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceRequiredToModifyItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceRequiredToModifyItemExtIEs = make([]int, 0)

type PDUSessionResourceSetupItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionResourceSetupItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceSetupItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceSetupItemExtIEs = make([]int, 0)

type PDUSessionResourceSetupModItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionResourceSetupModItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceSetupModItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceSetupModItemExtIEs = make([]int, 0)

type PDUSessionResourceToModifyItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-SNSSAI', 'CRITICALITY': 'reject', 'EXTENSION': 'SNSSAI', 'PRESENCE': 'optional'}, {'ID': 'id-CommonNetworkInstance', 'CRITICALITY': 'ignore', 'EXTENSION': 'CommonNetworkInstance', 'PRESENCE': 'optional'}, None]}
   SNSSAI  *SNSSAI
   CommonNetworkInstance  *CommonNetworkInstance
   list []interface{}
}
func (self *PDUSessionResourceToModifyItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceToModifyItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceToModifyItemExtIEs = make([]int, 2)

func (self *PDUSessionResourceToModifyItemExtIEs) GetIECount() int{
   count := 0
   if self.SNSSAI != nil { count += 1 }
   if self.CommonNetworkInstance != nil { count += 1 }
   return count//ObjSet
}
func (self *PDUSessionResourceToModifyItemExtIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 69: //SNSSAI
        if self.SNSSAI != nil { return true }
      case 78: //CommonNetworkInstance
        if self.CommonNetworkInstance != nil { return true }
   }
   return false//ObjSet
}
func (self *PDUSessionResourceToModifyItemExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 69: //SNSSAI
        self.SNSSAI = &SNSSAI{}
        self.SNSSAI.Unpack(st)
        self.list = append(self.list, self.SNSSAI)
      case 78: //CommonNetworkInstance
        self.CommonNetworkInstance = &CommonNetworkInstance{}
        self.CommonNetworkInstance.Unpack(st)
        self.list = append(self.list, self.CommonNetworkInstance)
   }
}
func (self *PDUSessionResourceToModifyItemExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 69: //SNSSAI
        if self.SNSSAI != nil {self.SNSSAI.Pack(st)}
      case 78: //CommonNetworkInstance
        if self.CommonNetworkInstance != nil {self.CommonNetworkInstance.Pack(st)}
      default:
      break
   }
}
func init() {
table_PDUSessionResourceToModifyItemExtIEs[69] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idSNSSAI}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&SNSSAI{}, PRESENCE:Presence{Presenceoptional}, }
order_PDUSessionResourceToModifyItemExtIEs[0] = 69
table_PDUSessionResourceToModifyItemExtIEs[78] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idCommonNetworkInstance}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&CommonNetworkInstance{}, PRESENCE:Presence{Presenceoptional}, }
order_PDUSessionResourceToModifyItemExtIEs[1] = 78
   }

type PDUSessionResourceToRemoveItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'EXTENSION': 'Cause', 'PRESENCE': 'optional'}, None]}
   Cause  *Cause
   list []interface{}
}
func (self *PDUSessionResourceToRemoveItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceToRemoveItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceToRemoveItemExtIEs = make([]int, 1)

func (self *PDUSessionResourceToRemoveItemExtIEs) GetIECount() int{
   count := 0
   if self.Cause != nil { count += 1 }
   return count//ObjSet
}
func (self *PDUSessionResourceToRemoveItemExtIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 0: //Cause
        if self.Cause != nil { return true }
   }
   return false//ObjSet
}
func (self *PDUSessionResourceToRemoveItemExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 0: //Cause
        self.Cause = &Cause{}
        self.Cause.Unpack(st)
        self.list = append(self.list, self.Cause)
   }
}
func (self *PDUSessionResourceToRemoveItemExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 0: //Cause
        if self.Cause != nil {self.Cause.Pack(st)}
      default:
      break
   }
}
func init() {
table_PDUSessionResourceToRemoveItemExtIEs[0] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&Cause{}, PRESENCE:Presence{Presenceoptional}, }
order_PDUSessionResourceToRemoveItemExtIEs[0] = 0
   }

type PDUSessionResourceToSetupItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-CommonNetworkInstance', 'CRITICALITY': 'ignore', 'EXTENSION': 'CommonNetworkInstance', 'PRESENCE': 'optional'}, None]}
   CommonNetworkInstance  *CommonNetworkInstance
   list []interface{}
}
func (self *PDUSessionResourceToSetupItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceToSetupItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceToSetupItemExtIEs = make([]int, 1)

func (self *PDUSessionResourceToSetupItemExtIEs) GetIECount() int{
   count := 0
   if self.CommonNetworkInstance != nil { count += 1 }
   return count//ObjSet
}
func (self *PDUSessionResourceToSetupItemExtIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 78: //CommonNetworkInstance
        if self.CommonNetworkInstance != nil { return true }
   }
   return false//ObjSet
}
func (self *PDUSessionResourceToSetupItemExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 78: //CommonNetworkInstance
        self.CommonNetworkInstance = &CommonNetworkInstance{}
        self.CommonNetworkInstance.Unpack(st)
        self.list = append(self.list, self.CommonNetworkInstance)
   }
}
func (self *PDUSessionResourceToSetupItemExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 78: //CommonNetworkInstance
        if self.CommonNetworkInstance != nil {self.CommonNetworkInstance.Pack(st)}
      default:
      break
   }
}
func init() {
table_PDUSessionResourceToSetupItemExtIEs[78] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idCommonNetworkInstance}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&CommonNetworkInstance{}, PRESENCE:Presence{Presenceoptional}, }
order_PDUSessionResourceToSetupItemExtIEs[0] = 78
   }

type PDUSessionResourceToSetupModItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-NetworkInstance', 'CRITICALITY': 'ignore', 'EXTENSION': 'NetworkInstance', 'PRESENCE': 'optional'}, {'ID': 'id-CommonNetworkInstance', 'CRITICALITY': 'ignore', 'EXTENSION': 'CommonNetworkInstance', 'PRESENCE': 'optional'}, None]}
   NetworkInstance  *NetworkInstance
   CommonNetworkInstance  *CommonNetworkInstance
   list []interface{}
}
func (self *PDUSessionResourceToSetupModItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionResourceToSetupModItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionResourceToSetupModItemExtIEs = make([]int, 2)

func (self *PDUSessionResourceToSetupModItemExtIEs) GetIECount() int{
   count := 0
   if self.NetworkInstance != nil { count += 1 }
   if self.CommonNetworkInstance != nil { count += 1 }
   return count//ObjSet
}
func (self *PDUSessionResourceToSetupModItemExtIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 79: //NetworkInstance
        if self.NetworkInstance != nil { return true }
      case 78: //CommonNetworkInstance
        if self.CommonNetworkInstance != nil { return true }
   }
   return false//ObjSet
}
func (self *PDUSessionResourceToSetupModItemExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 79: //NetworkInstance
        self.NetworkInstance = &NetworkInstance{}
        self.NetworkInstance.Unpack(st)
        self.list = append(self.list, self.NetworkInstance)
      case 78: //CommonNetworkInstance
        self.CommonNetworkInstance = &CommonNetworkInstance{}
        self.CommonNetworkInstance.Unpack(st)
        self.list = append(self.list, self.CommonNetworkInstance)
   }
}
func (self *PDUSessionResourceToSetupModItemExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 79: //NetworkInstance
        if self.NetworkInstance != nil {self.NetworkInstance.Pack(st)}
      case 78: //CommonNetworkInstance
        if self.CommonNetworkInstance != nil {self.CommonNetworkInstance.Pack(st)}
      default:
      break
   }
}
func init() {
table_PDUSessionResourceToSetupModItemExtIEs[79] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idNetworkInstance}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&NetworkInstance{}, PRESENCE:Presence{Presenceoptional}, }
order_PDUSessionResourceToSetupModItemExtIEs[0] = 79
table_PDUSessionResourceToSetupModItemExtIEs[78] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idCommonNetworkInstance}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&CommonNetworkInstance{}, PRESENCE:Presence{Presenceoptional}, }
order_PDUSessionResourceToSetupModItemExtIEs[1] = 78
   }

type PDUSessionToNotifyItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PDUSessionToNotifyItemExtIEs)createOT() interface{}{
    return nil
}
var table_PDUSessionToNotifyItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_PDUSessionToNotifyItemExtIEs = make([]int, 0)

type QoSCharacteristicsExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *QoSCharacteristicsExtIEs)createOT() interface{}{
    return nil
}
var table_QoSCharacteristicsExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_QoSCharacteristicsExtIEs = make([]int, 0)

type QoSFlowItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-QoSFlowMappingIndication', 'CRITICALITY': 'ignore', 'EXTENSION': 'QoS-Flow-Mapping-Indication', 'PRESENCE': 'optional'}, None]}
   QoSFlowMappingIndication  *QoSFlowMappingIndication
   list []interface{}
}
func (self *QoSFlowItemExtIEs)createOT() interface{}{
    return nil
}
var table_QoSFlowItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_QoSFlowItemExtIEs = make([]int, 1)

func (self *QoSFlowItemExtIEs) GetIECount() int{
   count := 0
   if self.QoSFlowMappingIndication != nil { count += 1 }
   return count//ObjSet
}
func (self *QoSFlowItemExtIEs) GetOT(id interface{}) bool{
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 80: //QoSFlowMappingIndication
        if self.QoSFlowMappingIndication != nil { return true }
   }
   return false//ObjSet
}
func (self *QoSFlowItemExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 80: //QoSFlowMappingIndication
        self.QoSFlowMappingIndication = &QoSFlowMappingIndication{}
        self.QoSFlowMappingIndication.Unpack(st)
        self.list = append(self.list, self.QoSFlowMappingIndication)
   }
}
func (self *QoSFlowItemExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E1APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 80: //QoSFlowMappingIndication
        if self.QoSFlowMappingIndication != nil {self.QoSFlowMappingIndication.Pack(st)}
      default:
      break
   }
}
func init() {
table_QoSFlowItemExtIEs[80] = &E1APPROTOCOLEXTENSION{ID:ProtocolIEID{idQoSFlowMappingIndication}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&QoSFlowMappingIndication{}, PRESENCE:Presence{Presenceoptional}, }
order_QoSFlowItemExtIEs[0] = 80
   }

type QoSFlowFailedItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *QoSFlowFailedItemExtIEs)createOT() interface{}{
    return nil
}
var table_QoSFlowFailedItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_QoSFlowFailedItemExtIEs = make([]int, 0)

type QoSFlowMappingItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *QoSFlowMappingItemExtIEs)createOT() interface{}{
    return nil
}
var table_QoSFlowMappingItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_QoSFlowMappingItemExtIEs = make([]int, 0)

type QoSParametersSupportListItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *QoSParametersSupportListItemExtIEs)createOT() interface{}{
    return nil
}
var table_QoSParametersSupportListItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_QoSParametersSupportListItemExtIEs = make([]int, 0)

type QoSFlowQoSParameterItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *QoSFlowQoSParameterItemExtIEs)createOT() interface{}{
    return nil
}
var table_QoSFlowQoSParameterItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_QoSFlowQoSParameterItemExtIEs = make([]int, 0)

type QoSFlowLevelQoSParametersExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *QoSFlowLevelQoSParametersExtIEs)createOT() interface{}{
    return nil
}
var table_QoSFlowLevelQoSParametersExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_QoSFlowLevelQoSParametersExtIEs = make([]int, 0)

type ROHCParametersExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *ROHCParametersExtIEs)createOT() interface{}{
    return nil
}
var table_ROHCParametersExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_ROHCParametersExtIEs = make([]int, 0)

type ROHCExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *ROHCExtIEs)createOT() interface{}{
    return nil
}
var table_ROHCExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_ROHCExtIEs = make([]int, 0)

type SecurityAlgorithmExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SecurityAlgorithmExtIEs)createOT() interface{}{
    return nil
}
var table_SecurityAlgorithmExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_SecurityAlgorithmExtIEs = make([]int, 0)

type SecurityIndicationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SecurityIndicationExtIEs)createOT() interface{}{
    return nil
}
var table_SecurityIndicationExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_SecurityIndicationExtIEs = make([]int, 0)

type SecurityInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SecurityInformationExtIEs)createOT() interface{}{
    return nil
}
var table_SecurityInformationExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_SecurityInformationExtIEs = make([]int, 0)

type SecurityResultExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SecurityResultExtIEs)createOT() interface{}{
    return nil
}
var table_SecurityResultExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_SecurityResultExtIEs = make([]int, 0)

type SliceSupportItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SliceSupportItemExtIEs)createOT() interface{}{
    return nil
}
var table_SliceSupportItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_SliceSupportItemExtIEs = make([]int, 0)

type SNSSAIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SNSSAIExtIEs)createOT() interface{}{
    return nil
}
var table_SNSSAIExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_SNSSAIExtIEs = make([]int, 0)

type SDAPConfigurationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SDAPConfigurationExtIEs)createOT() interface{}{
    return nil
}
var table_SDAPConfigurationExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_SDAPConfigurationExtIEs = make([]int, 0)

type TReorderingTimerExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TReorderingTimerExtIEs)createOT() interface{}{
    return nil
}
var table_TReorderingTimerExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_TReorderingTimerExtIEs = make([]int, 0)

type UEassociatedLogicalE1ConnectionItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UEassociatedLogicalE1ConnectionItemExtIEs)createOT() interface{}{
    return nil
}
var table_UEassociatedLogicalE1ConnectionItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_UEassociatedLogicalE1ConnectionItemExtIEs = make([]int, 0)

type UPParametersItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UPParametersItemExtIEs)createOT() interface{}{
    return nil
}
var table_UPParametersItemExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_UPParametersItemExtIEs = make([]int, 0)

type UPSecuritykeyExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UPSecuritykeyExtIEs)createOT() interface{}{
    return nil
}
var table_UPSecuritykeyExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_UPSecuritykeyExtIEs = make([]int, 0)

type UPTNLInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *UPTNLInformationExtIEs)createOT() interface{}{
    return nil
}
var table_UPTNLInformationExtIEs = make(map[int]*E1APPROTOCOLIES)

var order_UPTNLInformationExtIEs = make([]int, 0)

type UplinkOnlyROHCExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E1AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UplinkOnlyROHCExtIEs)createOT() interface{}{
    return nil
}
var table_UplinkOnlyROHCExtIEs = make(map[int]*E1APPROTOCOLEXTENSION)

var order_UplinkOnlyROHCExtIEs = make([]int, 0)

func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'Reset', 'SUCCESSFUL OUTCOME': 'ResetAcknowledge', 'PROCEDURE CODE': 'id-reset', 'CRITICALITY': 'reject'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idreset}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&Reset{}, SUCCESSFULOUTCOME:&ResetAcknowledge{}, PROCEDURECODE:ProcedureCode{idreset}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetResetINITIATINGMESSAGE() (*Reset, uint64, int) {/*TYPE, ID, Cricality*/
 return &Reset{}, uint64(idreset), int(Criticalityreject)
}
func GetResetSUCCESSFULOUTCOME() (*ResetAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &ResetAcknowledge{}, uint64(idreset), int(Criticalityreject)
}
func (self *Reset) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *Reset) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *Reset) createOT() interface{} {
   return &Reset{}
}
func (self *Reset) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *Reset) GetIECount() int{
    return 0
}
func (self *ResetAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ResetAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ResetAcknowledge) createOT() interface{} {
   return &ResetAcknowledge{}
}
func (self *ResetAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ResetAcknowledge) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ErrorIndication', 'PROCEDURE CODE': 'id-errorIndication', 'CRITICALITY': 'ignore'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{iderrorIndication}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ErrorIndication{}, PROCEDURECODE:ProcedureCode{iderrorIndication}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetErrorIndicationINITIATINGMESSAGE() (*ErrorIndication, uint64, int) {/*TYPE, ID, Cricality*/
 return &ErrorIndication{}, uint64(iderrorIndication), int(Criticalityignore)
}
func (self *ErrorIndication) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ErrorIndication) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ErrorIndication) createOT() interface{} {
   return &ErrorIndication{}
}
func (self *ErrorIndication) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ErrorIndication) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'GNB-CU-UP-E1SetupRequest', 'SUCCESSFUL OUTCOME': 'GNB-CU-UP-E1SetupResponse', 'UNSUCCESSFUL OUTCOME': 'GNB-CU-UP-E1SetupFailure', 'PROCEDURE CODE': 'id-gNB-CU-UP-E1Setup', 'CRITICALITY': 'reject'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idgNBCUUPE1Setup}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&GNBCUUPE1SetupRequest{}, SUCCESSFULOUTCOME:&GNBCUUPE1SetupResponse{}, UNSUCCESSFULOUTCOME:&GNBCUUPE1SetupFailure{}, PROCEDURECODE:ProcedureCode{idgNBCUUPE1Setup}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetGNBCUUPE1SetupINITIATINGMESSAGE() (*GNBCUUPE1SetupRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUUPE1SetupRequest{}, uint64(idgNBCUUPE1Setup), int(Criticalityreject)
}
func GetGNBCUUPE1SetupSUCCESSFULOUTCOME() (*GNBCUUPE1SetupResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUUPE1SetupResponse{}, uint64(idgNBCUUPE1Setup), int(Criticalityreject)
}
func GetGNBCUUPE1SetupUNSUCCESSFULOUTCOME() (*GNBCUUPE1SetupFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUUPE1SetupFailure{}, uint64(idgNBCUUPE1Setup), int(Criticalityreject)
}
func (self *GNBCUUPE1SetupRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUUPE1SetupRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUUPE1SetupRequest) createOT() interface{} {
   return &GNBCUUPE1SetupRequest{}
}
func (self *GNBCUUPE1SetupRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUUPE1SetupRequest) GetIECount() int{
    return 0
}
func (self *GNBCUUPE1SetupResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUUPE1SetupResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUUPE1SetupResponse) createOT() interface{} {
   return &GNBCUUPE1SetupResponse{}
}
func (self *GNBCUUPE1SetupResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUUPE1SetupResponse) GetIECount() int{
    return 0
}
func (self *GNBCUUPE1SetupFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUUPE1SetupFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUUPE1SetupFailure) createOT() interface{} {
   return &GNBCUUPE1SetupFailure{}
}
func (self *GNBCUUPE1SetupFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUUPE1SetupFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'GNB-CU-CP-E1SetupRequest', 'SUCCESSFUL OUTCOME': 'GNB-CU-CP-E1SetupResponse', 'UNSUCCESSFUL OUTCOME': 'GNB-CU-CP-E1SetupFailure', 'PROCEDURE CODE': 'id-gNB-CU-CP-E1Setup', 'CRITICALITY': 'reject'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idgNBCUCPE1Setup}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&GNBCUCPE1SetupRequest{}, SUCCESSFULOUTCOME:&GNBCUCPE1SetupResponse{}, UNSUCCESSFULOUTCOME:&GNBCUCPE1SetupFailure{}, PROCEDURECODE:ProcedureCode{idgNBCUCPE1Setup}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetGNBCUCPE1SetupINITIATINGMESSAGE() (*GNBCUCPE1SetupRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUCPE1SetupRequest{}, uint64(idgNBCUCPE1Setup), int(Criticalityreject)
}
func GetGNBCUCPE1SetupSUCCESSFULOUTCOME() (*GNBCUCPE1SetupResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUCPE1SetupResponse{}, uint64(idgNBCUCPE1Setup), int(Criticalityreject)
}
func GetGNBCUCPE1SetupUNSUCCESSFULOUTCOME() (*GNBCUCPE1SetupFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUCPE1SetupFailure{}, uint64(idgNBCUCPE1Setup), int(Criticalityreject)
}
func (self *GNBCUCPE1SetupRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUCPE1SetupRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUCPE1SetupRequest) createOT() interface{} {
   return &GNBCUCPE1SetupRequest{}
}
func (self *GNBCUCPE1SetupRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUCPE1SetupRequest) GetIECount() int{
    return 0
}
func (self *GNBCUCPE1SetupResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUCPE1SetupResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUCPE1SetupResponse) createOT() interface{} {
   return &GNBCUCPE1SetupResponse{}
}
func (self *GNBCUCPE1SetupResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUCPE1SetupResponse) GetIECount() int{
    return 0
}
func (self *GNBCUCPE1SetupFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUCPE1SetupFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUCPE1SetupFailure) createOT() interface{} {
   return &GNBCUCPE1SetupFailure{}
}
func (self *GNBCUCPE1SetupFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUCPE1SetupFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'GNB-CU-UP-ConfigurationUpdate', 'SUCCESSFUL OUTCOME': 'GNB-CU-UP-ConfigurationUpdateAcknowledge', 'UNSUCCESSFUL OUTCOME': 'GNB-CU-UP-ConfigurationUpdateFailure', 'PROCEDURE CODE': 'id-gNB-CU-UP-ConfigurationUpdate', 'CRITICALITY': 'reject'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idgNBCUUPConfigurationUpdate}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&GNBCUUPConfigurationUpdate{}, SUCCESSFULOUTCOME:&GNBCUUPConfigurationUpdateAcknowledge{}, UNSUCCESSFULOUTCOME:&GNBCUUPConfigurationUpdateFailure{}, PROCEDURECODE:ProcedureCode{idgNBCUUPConfigurationUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetGNBCUUPConfigurationUpdateINITIATINGMESSAGE() (*GNBCUUPConfigurationUpdate, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUUPConfigurationUpdate{}, uint64(idgNBCUUPConfigurationUpdate), int(Criticalityreject)
}
func GetGNBCUUPConfigurationUpdateSUCCESSFULOUTCOME() (*GNBCUUPConfigurationUpdateAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUUPConfigurationUpdateAcknowledge{}, uint64(idgNBCUUPConfigurationUpdate), int(Criticalityreject)
}
func GetGNBCUUPConfigurationUpdateUNSUCCESSFULOUTCOME() (*GNBCUUPConfigurationUpdateFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUUPConfigurationUpdateFailure{}, uint64(idgNBCUUPConfigurationUpdate), int(Criticalityreject)
}
func (self *GNBCUUPConfigurationUpdate) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUUPConfigurationUpdate) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUUPConfigurationUpdate) createOT() interface{} {
   return &GNBCUUPConfigurationUpdate{}
}
func (self *GNBCUUPConfigurationUpdate) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUUPConfigurationUpdate) GetIECount() int{
    return 0
}
func (self *GNBCUUPConfigurationUpdateAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUUPConfigurationUpdateAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUUPConfigurationUpdateAcknowledge) createOT() interface{} {
   return &GNBCUUPConfigurationUpdateAcknowledge{}
}
func (self *GNBCUUPConfigurationUpdateAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUUPConfigurationUpdateAcknowledge) GetIECount() int{
    return 0
}
func (self *GNBCUUPConfigurationUpdateFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUUPConfigurationUpdateFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUUPConfigurationUpdateFailure) createOT() interface{} {
   return &GNBCUUPConfigurationUpdateFailure{}
}
func (self *GNBCUUPConfigurationUpdateFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUUPConfigurationUpdateFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'GNB-CU-CP-ConfigurationUpdate', 'SUCCESSFUL OUTCOME': 'GNB-CU-CP-ConfigurationUpdateAcknowledge', 'UNSUCCESSFUL OUTCOME': 'GNB-CU-CP-ConfigurationUpdateFailure', 'PROCEDURE CODE': 'id-gNB-CU-CP-ConfigurationUpdate', 'CRITICALITY': 'reject'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idgNBCUCPConfigurationUpdate}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&GNBCUCPConfigurationUpdate{}, SUCCESSFULOUTCOME:&GNBCUCPConfigurationUpdateAcknowledge{}, UNSUCCESSFULOUTCOME:&GNBCUCPConfigurationUpdateFailure{}, PROCEDURECODE:ProcedureCode{idgNBCUCPConfigurationUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetGNBCUCPConfigurationUpdateINITIATINGMESSAGE() (*GNBCUCPConfigurationUpdate, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUCPConfigurationUpdate{}, uint64(idgNBCUCPConfigurationUpdate), int(Criticalityreject)
}
func GetGNBCUCPConfigurationUpdateSUCCESSFULOUTCOME() (*GNBCUCPConfigurationUpdateAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUCPConfigurationUpdateAcknowledge{}, uint64(idgNBCUCPConfigurationUpdate), int(Criticalityreject)
}
func GetGNBCUCPConfigurationUpdateUNSUCCESSFULOUTCOME() (*GNBCUCPConfigurationUpdateFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUCPConfigurationUpdateFailure{}, uint64(idgNBCUCPConfigurationUpdate), int(Criticalityreject)
}
func (self *GNBCUCPConfigurationUpdate) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUCPConfigurationUpdate) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUCPConfigurationUpdate) createOT() interface{} {
   return &GNBCUCPConfigurationUpdate{}
}
func (self *GNBCUCPConfigurationUpdate) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUCPConfigurationUpdate) GetIECount() int{
    return 0
}
func (self *GNBCUCPConfigurationUpdateAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUCPConfigurationUpdateAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUCPConfigurationUpdateAcknowledge) createOT() interface{} {
   return &GNBCUCPConfigurationUpdateAcknowledge{}
}
func (self *GNBCUCPConfigurationUpdateAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUCPConfigurationUpdateAcknowledge) GetIECount() int{
    return 0
}
func (self *GNBCUCPConfigurationUpdateFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUCPConfigurationUpdateFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUCPConfigurationUpdateFailure) createOT() interface{} {
   return &GNBCUCPConfigurationUpdateFailure{}
}
func (self *GNBCUCPConfigurationUpdateFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUCPConfigurationUpdateFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'E1ReleaseRequest', 'SUCCESSFUL OUTCOME': 'E1ReleaseResponse', 'PROCEDURE CODE': 'id-e1Release', 'CRITICALITY': 'reject'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{ide1Release}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&E1ReleaseRequest{}, SUCCESSFULOUTCOME:&E1ReleaseResponse{}, PROCEDURECODE:ProcedureCode{ide1Release}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetE1ReleaseINITIATINGMESSAGE() (*E1ReleaseRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &E1ReleaseRequest{}, uint64(ide1Release), int(Criticalityreject)
}
func GetE1ReleaseSUCCESSFULOUTCOME() (*E1ReleaseResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &E1ReleaseResponse{}, uint64(ide1Release), int(Criticalityreject)
}
func (self *E1ReleaseRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E1ReleaseRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E1ReleaseRequest) createOT() interface{} {
   return &E1ReleaseRequest{}
}
func (self *E1ReleaseRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E1ReleaseRequest) GetIECount() int{
    return 0
}
func (self *E1ReleaseResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E1ReleaseResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E1ReleaseResponse) createOT() interface{} {
   return &E1ReleaseResponse{}
}
func (self *E1ReleaseResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E1ReleaseResponse) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'BearerContextSetupRequest', 'SUCCESSFUL OUTCOME': 'BearerContextSetupResponse', 'UNSUCCESSFUL OUTCOME': 'BearerContextSetupFailure', 'PROCEDURE CODE': 'id-bearerContextSetup', 'CRITICALITY': 'reject'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idbearerContextSetup}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&BearerContextSetupRequest{}, SUCCESSFULOUTCOME:&BearerContextSetupResponse{}, UNSUCCESSFULOUTCOME:&BearerContextSetupFailure{}, PROCEDURECODE:ProcedureCode{idbearerContextSetup}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetBearerContextSetupINITIATINGMESSAGE() (*BearerContextSetupRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextSetupRequest{}, uint64(idbearerContextSetup), int(Criticalityreject)
}
func GetBearerContextSetupSUCCESSFULOUTCOME() (*BearerContextSetupResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextSetupResponse{}, uint64(idbearerContextSetup), int(Criticalityreject)
}
func GetBearerContextSetupUNSUCCESSFULOUTCOME() (*BearerContextSetupFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextSetupFailure{}, uint64(idbearerContextSetup), int(Criticalityreject)
}
func (self *BearerContextSetupRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextSetupRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextSetupRequest) createOT() interface{} {
   return &BearerContextSetupRequest{}
}
func (self *BearerContextSetupRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextSetupRequest) GetIECount() int{
    return 0
}
func (self *BearerContextSetupResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextSetupResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextSetupResponse) createOT() interface{} {
   return &BearerContextSetupResponse{}
}
func (self *BearerContextSetupResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextSetupResponse) GetIECount() int{
    return 0
}
func (self *BearerContextSetupFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextSetupFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextSetupFailure) createOT() interface{} {
   return &BearerContextSetupFailure{}
}
func (self *BearerContextSetupFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextSetupFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'BearerContextModificationRequest', 'SUCCESSFUL OUTCOME': 'BearerContextModificationResponse', 'UNSUCCESSFUL OUTCOME': 'BearerContextModificationFailure', 'PROCEDURE CODE': 'id-bearerContextModification', 'CRITICALITY': 'reject'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idbearerContextModification}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&BearerContextModificationRequest{}, SUCCESSFULOUTCOME:&BearerContextModificationResponse{}, UNSUCCESSFULOUTCOME:&BearerContextModificationFailure{}, PROCEDURECODE:ProcedureCode{idbearerContextModification}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetBearerContextModificationINITIATINGMESSAGE() (*BearerContextModificationRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextModificationRequest{}, uint64(idbearerContextModification), int(Criticalityreject)
}
func GetBearerContextModificationSUCCESSFULOUTCOME() (*BearerContextModificationResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextModificationResponse{}, uint64(idbearerContextModification), int(Criticalityreject)
}
func GetBearerContextModificationUNSUCCESSFULOUTCOME() (*BearerContextModificationFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextModificationFailure{}, uint64(idbearerContextModification), int(Criticalityreject)
}
func (self *BearerContextModificationRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextModificationRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextModificationRequest) createOT() interface{} {
   return &BearerContextModificationRequest{}
}
func (self *BearerContextModificationRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextModificationRequest) GetIECount() int{
    return 0
}
func (self *BearerContextModificationResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextModificationResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextModificationResponse) createOT() interface{} {
   return &BearerContextModificationResponse{}
}
func (self *BearerContextModificationResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextModificationResponse) GetIECount() int{
    return 0
}
func (self *BearerContextModificationFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextModificationFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextModificationFailure) createOT() interface{} {
   return &BearerContextModificationFailure{}
}
func (self *BearerContextModificationFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextModificationFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'BearerContextModificationRequired', 'SUCCESSFUL OUTCOME': 'BearerContextModificationConfirm', 'PROCEDURE CODE': 'id-bearerContextModificationRequired', 'CRITICALITY': 'reject'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idbearerContextModificationRequired}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&BearerContextModificationRequired{}, SUCCESSFULOUTCOME:&BearerContextModificationConfirm{}, PROCEDURECODE:ProcedureCode{idbearerContextModificationRequired}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetBearerContextModificationRequiredINITIATINGMESSAGE() (*BearerContextModificationRequired, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextModificationRequired{}, uint64(idbearerContextModificationRequired), int(Criticalityreject)
}
func GetBearerContextModificationRequiredSUCCESSFULOUTCOME() (*BearerContextModificationConfirm, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextModificationConfirm{}, uint64(idbearerContextModificationRequired), int(Criticalityreject)
}
func (self *BearerContextModificationRequired) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextModificationRequired) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextModificationRequired) createOT() interface{} {
   return &BearerContextModificationRequired{}
}
func (self *BearerContextModificationRequired) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextModificationRequired) GetIECount() int{
    return 0
}
func (self *BearerContextModificationConfirm) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextModificationConfirm) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextModificationConfirm) createOT() interface{} {
   return &BearerContextModificationConfirm{}
}
func (self *BearerContextModificationConfirm) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextModificationConfirm) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'BearerContextReleaseCommand', 'SUCCESSFUL OUTCOME': 'BearerContextReleaseComplete', 'PROCEDURE CODE': 'id-bearerContextRelease', 'CRITICALITY': 'reject'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idbearerContextRelease}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&BearerContextReleaseCommand{}, SUCCESSFULOUTCOME:&BearerContextReleaseComplete{}, PROCEDURECODE:ProcedureCode{idbearerContextRelease}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetBearerContextReleaseINITIATINGMESSAGE() (*BearerContextReleaseCommand, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextReleaseCommand{}, uint64(idbearerContextRelease), int(Criticalityreject)
}
func GetBearerContextReleaseSUCCESSFULOUTCOME() (*BearerContextReleaseComplete, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextReleaseComplete{}, uint64(idbearerContextRelease), int(Criticalityreject)
}
func (self *BearerContextReleaseCommand) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextReleaseCommand) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextReleaseCommand) createOT() interface{} {
   return &BearerContextReleaseCommand{}
}
func (self *BearerContextReleaseCommand) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextReleaseCommand) GetIECount() int{
    return 0
}
func (self *BearerContextReleaseComplete) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextReleaseComplete) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextReleaseComplete) createOT() interface{} {
   return &BearerContextReleaseComplete{}
}
func (self *BearerContextReleaseComplete) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextReleaseComplete) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'BearerContextReleaseRequest', 'PROCEDURE CODE': 'id-bearerContextReleaseRequest', 'CRITICALITY': 'ignore'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idbearerContextReleaseRequest}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&BearerContextReleaseRequest{}, PROCEDURECODE:ProcedureCode{idbearerContextReleaseRequest}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetBearerContextReleaseRequestINITIATINGMESSAGE() (*BearerContextReleaseRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextReleaseRequest{}, uint64(idbearerContextReleaseRequest), int(Criticalityignore)
}
func (self *BearerContextReleaseRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextReleaseRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextReleaseRequest) createOT() interface{} {
   return &BearerContextReleaseRequest{}
}
func (self *BearerContextReleaseRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextReleaseRequest) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'BearerContextInactivityNotification', 'PROCEDURE CODE': 'id-bearerContextInactivityNotification', 'CRITICALITY': 'ignore'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idbearerContextInactivityNotification}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&BearerContextInactivityNotification{}, PROCEDURECODE:ProcedureCode{idbearerContextInactivityNotification}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetBearerContextInactivityNotificationINITIATINGMESSAGE() (*BearerContextInactivityNotification, uint64, int) {/*TYPE, ID, Cricality*/
 return &BearerContextInactivityNotification{}, uint64(idbearerContextInactivityNotification), int(Criticalityignore)
}
func (self *BearerContextInactivityNotification) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *BearerContextInactivityNotification) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *BearerContextInactivityNotification) createOT() interface{} {
   return &BearerContextInactivityNotification{}
}
func (self *BearerContextInactivityNotification) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *BearerContextInactivityNotification) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'DLDataNotification', 'PROCEDURE CODE': 'id-dLDataNotification', 'CRITICALITY': 'ignore'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{iddLDataNotification}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&DLDataNotification{}, PROCEDURECODE:ProcedureCode{iddLDataNotification}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetDLDataNotificationINITIATINGMESSAGE() (*DLDataNotification, uint64, int) {/*TYPE, ID, Cricality*/
 return &DLDataNotification{}, uint64(iddLDataNotification), int(Criticalityignore)
}
func (self *DLDataNotification) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *DLDataNotification) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *DLDataNotification) createOT() interface{} {
   return &DLDataNotification{}
}
func (self *DLDataNotification) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *DLDataNotification) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ULDataNotification', 'PROCEDURE CODE': 'id-uLDataNotification', 'CRITICALITY': 'ignore'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{iduLDataNotification}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ULDataNotification{}, PROCEDURECODE:ProcedureCode{iduLDataNotification}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetULDataNotificationINITIATINGMESSAGE() (*ULDataNotification, uint64, int) {/*TYPE, ID, Cricality*/
 return &ULDataNotification{}, uint64(iduLDataNotification), int(Criticalityignore)
}
func (self *ULDataNotification) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ULDataNotification) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ULDataNotification) createOT() interface{} {
   return &ULDataNotification{}
}
func (self *ULDataNotification) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ULDataNotification) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'DataUsageReport', 'PROCEDURE CODE': 'id-dataUsageReport', 'CRITICALITY': 'ignore'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{iddataUsageReport}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&DataUsageReport{}, PROCEDURECODE:ProcedureCode{iddataUsageReport}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetDataUsageReportINITIATINGMESSAGE() (*DataUsageReport, uint64, int) {/*TYPE, ID, Cricality*/
 return &DataUsageReport{}, uint64(iddataUsageReport), int(Criticalityignore)
}
func (self *DataUsageReport) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *DataUsageReport) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *DataUsageReport) createOT() interface{} {
   return &DataUsageReport{}
}
func (self *DataUsageReport) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *DataUsageReport) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'GNB-CU-UP-CounterCheckRequest', 'PROCEDURE CODE': 'id-gNB-CU-UP-CounterCheck', 'CRITICALITY': 'ignore'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idgNBCUUPCounterCheck}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&GNBCUUPCounterCheckRequest{}, PROCEDURECODE:ProcedureCode{idgNBCUUPCounterCheck}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetGNBCUUPCounterCheckINITIATINGMESSAGE() (*GNBCUUPCounterCheckRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUUPCounterCheckRequest{}, uint64(idgNBCUUPCounterCheck), int(Criticalityignore)
}
func (self *GNBCUUPCounterCheckRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUUPCounterCheckRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUUPCounterCheckRequest) createOT() interface{} {
   return &GNBCUUPCounterCheckRequest{}
}
func (self *GNBCUUPCounterCheckRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUUPCounterCheckRequest) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'GNB-CU-UP-StatusIndication', 'PROCEDURE CODE': 'id-gNB-CU-UP-StatusIndication', 'CRITICALITY': 'ignore'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idgNBCUUPStatusIndication}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&GNBCUUPStatusIndication{}, PROCEDURECODE:ProcedureCode{idgNBCUUPStatusIndication}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetGNBCUUPStatusIndicationINITIATINGMESSAGE() (*GNBCUUPStatusIndication, uint64, int) {/*TYPE, ID, Cricality*/
 return &GNBCUUPStatusIndication{}, uint64(idgNBCUUPStatusIndication), int(Criticalityignore)
}
func (self *GNBCUUPStatusIndication) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *GNBCUUPStatusIndication) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *GNBCUUPStatusIndication) createOT() interface{} {
   return &GNBCUUPStatusIndication{}
}
func (self *GNBCUUPStatusIndication) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *GNBCUUPStatusIndication) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'PrivateMessage', 'PROCEDURE CODE': 'id-privateMessage', 'CRITICALITY': 'ignore'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idprivateMessage}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&PrivateMessage{}, PROCEDURECODE:ProcedureCode{idprivateMessage}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetPrivateMessageINITIATINGMESSAGE() (*PrivateMessage, uint64, int) {/*TYPE, ID, Cricality*/
 return &PrivateMessage{}, uint64(idprivateMessage), int(Criticalityignore)
}
func (self *PrivateMessage) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *PrivateMessage) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *PrivateMessage) createOT() interface{} {
   return &PrivateMessage{}
}
func (self *PrivateMessage) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *PrivateMessage) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E1AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'MRDC-DataUsageReport', 'PROCEDURE CODE': 'id-mRDC-DataUsageReport', 'CRITICALITY': 'ignore'}]}
table_E1APELEMENTARYPROCEDURES[E1APELEMENTARYPROCEDUREprocedureCode{idmRDCDataUsageReport}] = &E1APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&MRDCDataUsageReport{}, PROCEDURECODE:ProcedureCode{idmRDCDataUsageReport}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetMRDCDataUsageReportINITIATINGMESSAGE() (*MRDCDataUsageReport, uint64, int) {/*TYPE, ID, Cricality*/
 return &MRDCDataUsageReport{}, uint64(idmRDCDataUsageReport), int(Criticalityignore)
}
func (self *MRDCDataUsageReport) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MRDCDataUsageReport) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MRDCDataUsageReport) createOT() interface{} {
   return &MRDCDataUsageReport{}
}
func (self *MRDCDataUsageReport) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MRDCDataUsageReport) GetIECount() int{
    return 0
}
var maxPrivateIEs uint64 = 65535
var maxProtocolExtensions uint64 = 65535
var maxProtocolIEs uint64 = 65535
var idreset uint64 = 0
const ProcedureCodereset = 0
var iderrorIndication uint64 = 1
const ProcedureCodeerrorIndication = 1
var idprivateMessage uint64 = 2
const ProcedureCodeprivateMessage = 2
var idgNBCUUPE1Setup uint64 = 3
const ProcedureCodegNBCUUPE1Setup = 3
var idgNBCUCPE1Setup uint64 = 4
const ProcedureCodegNBCUCPE1Setup = 4
var idgNBCUUPConfigurationUpdate uint64 = 5
const ProcedureCodegNBCUUPConfigurationUpdate = 5
var idgNBCUCPConfigurationUpdate uint64 = 6
const ProcedureCodegNBCUCPConfigurationUpdate = 6
var ide1Release uint64 = 7
const ProcedureCodee1Release = 7
var idbearerContextSetup uint64 = 8
const ProcedureCodebearerContextSetup = 8
var idbearerContextModification uint64 = 9
const ProcedureCodebearerContextModification = 9
var idbearerContextModificationRequired uint64 = 10
const ProcedureCodebearerContextModificationRequired = 10
var idbearerContextRelease uint64 = 11
const ProcedureCodebearerContextRelease = 11
var idbearerContextReleaseRequest uint64 = 12
const ProcedureCodebearerContextReleaseRequest = 12
var idbearerContextInactivityNotification uint64 = 13
const ProcedureCodebearerContextInactivityNotification = 13
var iddLDataNotification uint64 = 14
const ProcedureCodedLDataNotification = 14
var iddataUsageReport uint64 = 15
const ProcedureCodedataUsageReport = 15
var idgNBCUUPCounterCheck uint64 = 16
const ProcedureCodegNBCUUPCounterCheck = 16
var idgNBCUUPStatusIndication uint64 = 17
const ProcedureCodegNBCUUPStatusIndication = 17
var iduLDataNotification uint64 = 18
const ProcedureCodeuLDataNotification = 18
var idmRDCDataUsageReport uint64 = 19
const ProcedureCodemRDCDataUsageReport = 19
var maxnoofErrors uint64 = 256
var maxnoofSPLMNs uint64 = 12
var maxnoofSliceItems uint64 = 1024
var maxnoofIndividualE1ConnectionsToReset uint64 = 65536
var maxnoofEUTRANQOSParameters uint64 = 256
var maxnoofNGRANQOSParameters uint64 = 256
var maxnoofDRBs uint64 = 32
var maxnoofNRCGI uint64 = 512
var maxnoofPDUSessionResource uint64 = 256
var maxnoofQoSFlows uint64 = 64
var maxnoofUPParameters uint64 = 8
var maxnoofCellGroups uint64 = 4
var maxnooftimeperiods uint64 = 2
var maxnoofTNLAssociations uint64 = 32
var idCause uint64 = 0
const ProtocolIEIDCause = 0
var idCriticalityDiagnostics uint64 = 1
const ProtocolIEIDCriticalityDiagnostics = 1
var idgNBCUCPUEE1APID uint64 = 2
const ProtocolIEIDgNBCUCPUEE1APID = 2
var idgNBCUUPUEE1APID uint64 = 3
const ProtocolIEIDgNBCUUPUEE1APID = 3
var idResetType uint64 = 4
const ProtocolIEIDResetType = 4
var idUEassociatedLogicalE1ConnectionItem uint64 = 5
const ProtocolIEIDUEassociatedLogicalE1ConnectionItem = 5
var idUEassociatedLogicalE1ConnectionListResAck uint64 = 6
const ProtocolIEIDUEassociatedLogicalE1ConnectionListResAck = 6
var idgNBCUUPID uint64 = 7
const ProtocolIEIDgNBCUUPID = 7
var idgNBCUUPName uint64 = 8
const ProtocolIEIDgNBCUUPName = 8
var idgNBCUCPName uint64 = 9
const ProtocolIEIDgNBCUCPName = 9
var idCNSupport uint64 = 10
const ProtocolIEIDCNSupport = 10
var idSupportedPLMNs uint64 = 11
const ProtocolIEIDSupportedPLMNs = 11
var idTimeToWait uint64 = 12
const ProtocolIEIDTimeToWait = 12
var idSecurityInformation uint64 = 13
const ProtocolIEIDSecurityInformation = 13
var idUEDLAggregateMaximumBitRate uint64 = 14
const ProtocolIEIDUEDLAggregateMaximumBitRate = 14
var idSystemBearerContextSetupRequest uint64 = 15
const ProtocolIEIDSystemBearerContextSetupRequest = 15
var idSystemBearerContextSetupResponse uint64 = 16
const ProtocolIEIDSystemBearerContextSetupResponse = 16
var idBearerContextStatusChange uint64 = 17
const ProtocolIEIDBearerContextStatusChange = 17
var idSystemBearerContextModificationRequest uint64 = 18
const ProtocolIEIDSystemBearerContextModificationRequest = 18
var idSystemBearerContextModificationResponse uint64 = 19
const ProtocolIEIDSystemBearerContextModificationResponse = 19
var idSystemBearerContextModificationConfirm uint64 = 20
const ProtocolIEIDSystemBearerContextModificationConfirm = 20
var idSystemBearerContextModificationRequired uint64 = 21
const ProtocolIEIDSystemBearerContextModificationRequired = 21
var idDRBStatusList uint64 = 22
const ProtocolIEIDDRBStatusList = 22
var idActivityNotificationLevel uint64 = 23
const ProtocolIEIDActivityNotificationLevel = 23
var idActivityInformation uint64 = 24
const ProtocolIEIDActivityInformation = 24
var idDataUsageReportList uint64 = 25
const ProtocolIEIDDataUsageReportList = 25
var idNewULTNLInformationRequired uint64 = 26
const ProtocolIEIDNewULTNLInformationRequired = 26
var idGNBCUCPTNLAToAddList uint64 = 27
const ProtocolIEIDGNBCUCPTNLAToAddList = 27
var idGNBCUCPTNLAToRemoveList uint64 = 28
const ProtocolIEIDGNBCUCPTNLAToRemoveList = 28
var idGNBCUCPTNLAToUpdateList uint64 = 29
const ProtocolIEIDGNBCUCPTNLAToUpdateList = 29
var idGNBCUCPTNLASetupList uint64 = 30
const ProtocolIEIDGNBCUCPTNLASetupList = 30
var idGNBCUCPTNLAFailedToSetupList uint64 = 31
const ProtocolIEIDGNBCUCPTNLAFailedToSetupList = 31
var idDRBToSetupListEUTRAN uint64 = 32
const ProtocolIEIDDRBToSetupListEUTRAN = 32
var idDRBToModifyListEUTRAN uint64 = 33
const ProtocolIEIDDRBToModifyListEUTRAN = 33
var idDRBToRemoveListEUTRAN uint64 = 34
const ProtocolIEIDDRBToRemoveListEUTRAN = 34
var idDRBRequiredToModifyListEUTRAN uint64 = 35
const ProtocolIEIDDRBRequiredToModifyListEUTRAN = 35
var idDRBRequiredToRemoveListEUTRAN uint64 = 36
const ProtocolIEIDDRBRequiredToRemoveListEUTRAN = 36
var idDRBSetupListEUTRAN uint64 = 37
const ProtocolIEIDDRBSetupListEUTRAN = 37
var idDRBFailedListEUTRAN uint64 = 38
const ProtocolIEIDDRBFailedListEUTRAN = 38
var idDRBModifiedListEUTRAN uint64 = 39
const ProtocolIEIDDRBModifiedListEUTRAN = 39
var idDRBFailedToModifyListEUTRAN uint64 = 40
const ProtocolIEIDDRBFailedToModifyListEUTRAN = 40
var idDRBConfirmModifiedListEUTRAN uint64 = 41
const ProtocolIEIDDRBConfirmModifiedListEUTRAN = 41
var idPDUSessionResourceToSetupList uint64 = 42
const ProtocolIEIDPDUSessionResourceToSetupList = 42
var idPDUSessionResourceToModifyList uint64 = 43
const ProtocolIEIDPDUSessionResourceToModifyList = 43
var idPDUSessionResourceToRemoveList uint64 = 44
const ProtocolIEIDPDUSessionResourceToRemoveList = 44
var idPDUSessionResourceRequiredToModifyList uint64 = 45
const ProtocolIEIDPDUSessionResourceRequiredToModifyList = 45
var idPDUSessionResourceSetupList uint64 = 46
const ProtocolIEIDPDUSessionResourceSetupList = 46
var idPDUSessionResourceFailedList uint64 = 47
const ProtocolIEIDPDUSessionResourceFailedList = 47
var idPDUSessionResourceModifiedList uint64 = 48
const ProtocolIEIDPDUSessionResourceModifiedList = 48
var idPDUSessionResourceFailedToModifyList uint64 = 49
const ProtocolIEIDPDUSessionResourceFailedToModifyList = 49
var idPDUSessionResourceConfirmModifiedList uint64 = 50
const ProtocolIEIDPDUSessionResourceConfirmModifiedList = 50
var idDRBToSetupModListEUTRAN uint64 = 51
const ProtocolIEIDDRBToSetupModListEUTRAN = 51
var idDRBSetupModListEUTRAN uint64 = 52
const ProtocolIEIDDRBSetupModListEUTRAN = 52
var idDRBFailedModListEUTRAN uint64 = 53
const ProtocolIEIDDRBFailedModListEUTRAN = 53
var idPDUSessionResourceSetupModList uint64 = 54
const ProtocolIEIDPDUSessionResourceSetupModList = 54
var idPDUSessionResourceFailedModList uint64 = 55
const ProtocolIEIDPDUSessionResourceFailedModList = 55
var idPDUSessionResourceToSetupModList uint64 = 56
const ProtocolIEIDPDUSessionResourceToSetupModList = 56
var idTransactionID uint64 = 57
const ProtocolIEIDTransactionID = 57
var idServingPLMN uint64 = 58
const ProtocolIEIDServingPLMN = 58
var idUEInactivityTimer uint64 = 59
const ProtocolIEIDUEInactivityTimer = 59
var idSystemGNBCUUPCounterCheckRequest uint64 = 60
const ProtocolIEIDSystemGNBCUUPCounterCheckRequest = 60
var idDRBsSubjectToCounterCheckListEUTRAN uint64 = 61
const ProtocolIEIDDRBsSubjectToCounterCheckListEUTRAN = 61
var idDRBsSubjectToCounterCheckListNGRAN uint64 = 62
const ProtocolIEIDDRBsSubjectToCounterCheckListNGRAN = 62
var idPPI uint64 = 63
const ProtocolIEIDPPI = 63
var idgNBCUUPCapacity uint64 = 64
const ProtocolIEIDgNBCUUPCapacity = 64
var idGNBCUUPOverloadInformation uint64 = 65
const ProtocolIEIDGNBCUUPOverloadInformation = 65
var idUEDLMaximumIntegrityProtectedDataRate uint64 = 66
const ProtocolIEIDUEDLMaximumIntegrityProtectedDataRate = 66
var idPDUSessionToNotifyList uint64 = 67
const ProtocolIEIDPDUSessionToNotifyList = 67
var idPDUSessionResourceDataUsageList uint64 = 68
const ProtocolIEIDPDUSessionResourceDataUsageList = 68
var idSNSSAI uint64 = 69
const ProtocolIEIDSNSSAI = 69
var idDataDiscardRequired uint64 = 70
const ProtocolIEIDDataDiscardRequired = 70
var idOldQoSFlowMapULendmarkerexpected uint64 = 71
const ProtocolIEIDOldQoSFlowMapULendmarkerexpected = 71
var idDRBQoS uint64 = 72
const ProtocolIEIDDRBQoS = 72
var idGNBCUUPTNLAToRemoveList uint64 = 73
const ProtocolIEIDGNBCUUPTNLAToRemoveList = 73
var idendpointIPAddressandPort uint64 = 74
const ProtocolIEIDendpointIPAddressandPort = 74
var idTNLAssociationTransportLayerAddressgNBCUUP uint64 = 75
const ProtocolIEIDTNLAssociationTransportLayerAddressgNBCUUP = 75
var idRANUEID uint64 = 76
const ProtocolIEIDRANUEID = 76
var idGNBDUID uint64 = 77
const ProtocolIEIDGNBDUID = 77
var idCommonNetworkInstance uint64 = 78
const ProtocolIEIDCommonNetworkInstance = 78
var idNetworkInstance uint64 = 79
const ProtocolIEIDNetworkInstance = 79
var idQoSFlowMappingIndication uint64 = 80
const ProtocolIEIDQoSFlowMappingIndication = 80
