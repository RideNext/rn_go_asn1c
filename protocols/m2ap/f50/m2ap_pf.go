
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package m2ap
import (
  log "github.com/sirupsen/logrus"
)
var version = "vf50"

func fmtm2ap() {log.Debug("m2ap")}
func (self *M2APPDU)Unpack(stream *Stream) {
    //coptions := []string{"initiatingMessage","successfulOutcome","unsuccessfulOutcome","Unknown"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in M2APPDU\n", choice, choice_len)
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
func (self * M2APPDU) Pack(stream *Stream) {
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
type M2APPDU struct { //[{'type': 'InitiatingMessage', 'name': 'initiatingMessage'}, {'type': 'SuccessfulOutcome', 'name': 'successfulOutcome'}, {'type': 'UnsuccessfulOutcome', 'name': 'unsuccessfulOutcome'}, None]
    InitiatingMessage *InitiatingMessage
    SuccessfulOutcome *SuccessfulOutcome
    UnsuccessfulOutcome *UnsuccessfulOutcome
} // M2APPDU

type InitiatingMessage struct { // [{'type': 'M2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'M2AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['M2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'M2AP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['M2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode M2APELEMENTARYPROCEDUREprocedureCode
    Criticality M2APELEMENTARYPROCEDUREcriticality
    Value M2APELEMENTARYPROCEDUREInitiatingMessage
}

func (self * InitiatingMessage) Unpack(stream *Stream) {
    //table {'type': 'M2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_M2APELEMENTARYPROCEDURES[key].INITIATINGMESSAGE
    self.Value = out.(M2APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(M2APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'M2AP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['M2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * InitiatingMessage) Pack(stream *Stream) {
    //table {'type': 'M2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_M2APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(M2APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type SuccessfulOutcome struct { // [{'type': 'M2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'M2AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['M2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'M2AP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['M2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode M2APELEMENTARYPROCEDUREprocedureCode
    Criticality M2APELEMENTARYPROCEDUREcriticality
    Value M2APELEMENTARYPROCEDURESuccessfulOutcome
}

func (self * SuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'M2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_M2APELEMENTARYPROCEDURES[key].SUCCESSFULOUTCOME
    self.Value = out.(M2APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(M2APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'M2AP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['M2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * SuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'M2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_M2APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(M2APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type UnsuccessfulOutcome struct { // [{'type': 'M2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'M2AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['M2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'M2AP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['M2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode M2APELEMENTARYPROCEDUREprocedureCode
    Criticality M2APELEMENTARYPROCEDUREcriticality
    Value M2APELEMENTARYPROCEDUREUnsuccessfulOutcome
}

func (self * UnsuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'M2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_M2APELEMENTARYPROCEDURES[key].UNSUCCESSFULOUTCOME
    self.Value = out.(M2APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(M2APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'M2AP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['M2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * UnsuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'M2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_M2APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(M2APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type SessionStartRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SessionStartRequest-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs SessionStartRequestIes
}

func (self * SessionStartRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_SessionStartRequestIes, order_SessionStartRequestIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SessionStartRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SessionStartRequestIes, order_SessionStartRequestIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SessionStartResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SessionStartResponse-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs SessionStartResponseIes
}

func (self * SessionStartResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_SessionStartResponseIes, order_SessionStartResponseIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SessionStartResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SessionStartResponseIes, order_SessionStartResponseIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SessionStartFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SessionStartFailure-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs SessionStartFailureIes
}

func (self * SessionStartFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_SessionStartFailureIes, order_SessionStartFailureIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SessionStartFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SessionStartFailureIes, order_SessionStartFailureIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SessionStopRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SessionStopRequest-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs SessionStopRequestIes
}

func (self * SessionStopRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_SessionStopRequestIes, order_SessionStopRequestIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SessionStopRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SessionStopRequestIes, order_SessionStopRequestIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SessionStopResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SessionStopResponse-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs SessionStopResponseIes
}

func (self * SessionStopResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_SessionStopResponseIes, order_SessionStopResponseIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SessionStopResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SessionStopResponseIes, order_SessionStopResponseIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SessionUpdateRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SessionUpdateRequest-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs SessionUpdateRequestIes
}

func (self * SessionUpdateRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_SessionUpdateRequestIes, order_SessionUpdateRequestIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SessionUpdateRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SessionUpdateRequestIes, order_SessionUpdateRequestIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SessionUpdateResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SessionUpdateResponse-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs SessionUpdateResponseIes
}

func (self * SessionUpdateResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_SessionUpdateResponseIes, order_SessionUpdateResponseIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SessionUpdateResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SessionUpdateResponseIes, order_SessionUpdateResponseIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SessionUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SessionUpdateFailure-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs SessionUpdateFailureIes
}

func (self * SessionUpdateFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_SessionUpdateFailureIes, order_SessionUpdateFailureIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SessionUpdateFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SessionUpdateFailureIes, order_SessionUpdateFailureIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MbmsSchedulingInformation struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MbmsSchedulingInformation-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MbmsSchedulingInformationIes
}

func (self * MbmsSchedulingInformation) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MbmsSchedulingInformationIes, order_MbmsSchedulingInformationIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MbmsSchedulingInformation) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MbmsSchedulingInformationIes, order_MbmsSchedulingInformationIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *MBSFNAreaConfigurationList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]MBSFNAreaConfigurationItem, _size)
    val := ProtocolIEContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *MBSFNAreaConfigurationList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    //for item in table_MBSFNAreaConfigurationItem:
    val := ProtocolIEContainer{table_MBSFNAreaConfigurationItem, order_MBSFNAreaConfigurationItem}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type MBSFNAreaConfigurationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['MBSFN-Area-Configuration-Item']}, 'size': [(1, 'maxnoofMBSFNareas')]}
    Items []MBSFNAreaConfigurationItem
}

func (self *PMCHConfigurationList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(16)
    _size += 0
    self.Items = make([]PMCHConfigurationItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *PMCHConfigurationList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-0, 16)
    //for item in table_PMCHConfigurationItemIEs:
    val := ProtocolIESingleContainer{table_PMCHConfigurationItemIEs, order_PMCHConfigurationItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type PMCHConfigurationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['PMCH-Configuration-ItemIEs']}, 'size': [(0, 'maxnoofPMCHsperMBSFNarea')]}
    Items []PMCHConfigurationItemIEs
}

type PMCHConfigurationItem struct { // [{'type': 'PMCH-Configuration', 'name': 'pmch-Configuration'}, {'type': 'MBMSsessionListPerPMCH-Item', 'name': 'mbms-Session-List'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PMCH-Configuration-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PmchConfiguration PMCHConfiguration
    MbmsSessionList MBMSsessionListPerPMCHItem
    IEExtensions *PMCHConfigurationItemExtIEs
}

func (self * PMCHConfigurationItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PmchConfiguration.Unpack(stream)// p8
    self.MbmsSessionList.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PMCHConfigurationItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PMCH-Configuration-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PMCHConfigurationItemExtIEs, order_PMCHConfigurationItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PMCHConfigurationItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PmchConfiguration.Pack(stream)
    self.MbmsSessionList.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PMCHConfigurationItemExtIEs, order_PMCHConfigurationItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *MBSFNSubframeConfigurationList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(8)
    _size += 1
    self.Items = make([]MBSFNSubframeConfigurationItem, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *MBSFNSubframeConfigurationList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 8)
    //for item in table_MBSFNSubframeConfigurationItem:
    val := ProtocolIESingleContainer{table_MBSFNSubframeConfigurationItem, order_MBSFNSubframeConfigurationItem}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type MBSFNSubframeConfigurationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['MBSFN-Subframe-ConfigurationItem']}, 'size': [(1, 'maxnoofMBSFN-Allocations')]}
    Items []MBSFNSubframeConfigurationItem
}

func (self *MBMSSuspensionNotificationList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(15)
    _size += 1
    self.Items = make([]MBMSSuspensionNotificationItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *MBMSSuspensionNotificationList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 15)
    //for item in table_MBMSSuspensionNotificationItemIEs:
    val := ProtocolIESingleContainer{table_MBMSSuspensionNotificationItemIEs, order_MBMSSuspensionNotificationItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type MBMSSuspensionNotificationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['MBMS-Suspension-Notification-ItemIEs']}, 'size': [(1, 'maxnoofPMCHsperMBSFNarea')]}
    Items []MBMSSuspensionNotificationItemIEs
}

type MBMSSuspensionNotificationItem struct { // [{'type': 'SFN', 'name': 'sfn'}, {'type': 'MBMSsessionsToBeSuspendedListPerPMCH-Item', 'name': 'mbms-Sessions-To-Be-Suspended-List'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-Suspension-Notification-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    Sfn SFN
    MbmsSessionsToBeSuspendedList MBMSsessionsToBeSuspendedListPerPMCHItem
    IEExtensions *MBMSSuspensionNotificationItemExtIEs
}

func (self * MBMSSuspensionNotificationItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.Sfn.Unpack(stream)// p8
    self.MbmsSessionsToBeSuspendedList.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MBMSSuspensionNotificationItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-Suspension-Notification-ItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MBMSSuspensionNotificationItemExtIEs, order_MBMSSuspensionNotificationItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSSuspensionNotificationItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.Sfn.Pack(stream)
    self.MbmsSessionsToBeSuspendedList.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MBMSSuspensionNotificationItemExtIEs, order_MBMSSuspensionNotificationItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type MbmsSchedulingInformationResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MbmsSchedulingInformationResponse-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MbmsSchedulingInformationResponseIes
}

func (self * MbmsSchedulingInformationResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MbmsSchedulingInformationResponseIes, order_MbmsSchedulingInformationResponseIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MbmsSchedulingInformationResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MbmsSchedulingInformationResponseIes, order_MbmsSchedulingInformationResponseIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type M2SetupRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['M2SetupRequest-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs M2SetupRequestIes
}

func (self * M2SetupRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_M2SetupRequestIes, order_M2SetupRequestIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * M2SetupRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_M2SetupRequestIes, order_M2SetupRequestIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *ENBMBMSConfigurationdataList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]ENBMBMSConfigurationdataItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *ENBMBMSConfigurationdataList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    //for item in table_ENBMBMSConfigurationdataItemIEs:
    val := ProtocolIESingleContainer{table_ENBMBMSConfigurationdataItemIEs, order_ENBMBMSConfigurationdataItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type ENBMBMSConfigurationdataList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['ENB-MBMS-Configuration-data-ItemIEs']}, 'size': [(1, 'maxnoofCells')]}
    Items []ENBMBMSConfigurationdataItemIEs
}

type M2SetupResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['M2SetupResponse-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs M2SetupResponseIes
}

func (self * M2SetupResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_M2SetupResponseIes, order_M2SetupResponseIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * M2SetupResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_M2SetupResponseIes, order_M2SetupResponseIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *MCCHrelatedBCCHConfigPerMBSFNArea) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *MCCHrelatedBCCHConfigPerMBSFNArea) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    //for item in table_MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs:
    val := ProtocolIESingleContainer{table_MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs, order_MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type MCCHrelatedBCCHConfigPerMBSFNArea struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['MCCHrelatedBCCH-ConfigPerMBSFNArea-ItemIEs']}, 'size': [(1, 'maxnoofMBSFNareas')]}
    Items []MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs
}

type M2SetupFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['M2SetupFailure-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs M2SetupFailureIes
}

func (self * M2SetupFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_M2SetupFailureIes, order_M2SetupFailureIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * M2SetupFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_M2SetupFailureIes, order_M2SetupFailureIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type ENBConfigurationUpdate struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ENBConfigurationUpdate-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ENBConfigurationUpdateIes
}

func (self * ENBConfigurationUpdate) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ENBConfigurationUpdateIes, order_ENBConfigurationUpdateIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ENBConfigurationUpdate) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ENBConfigurationUpdateIes, order_ENBConfigurationUpdateIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *ENBMBMSConfigurationdataListConfigUpdate) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]ENBMBMSConfigurationdataConfigUpdateItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *ENBMBMSConfigurationdataListConfigUpdate) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    //for item in table_ENBMBMSConfigurationdataConfigUpdateItemIEs:
    val := ProtocolIESingleContainer{table_ENBMBMSConfigurationdataConfigUpdateItemIEs, order_ENBMBMSConfigurationdataConfigUpdateItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type ENBMBMSConfigurationdataListConfigUpdate struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['ENB-MBMS-Configuration-data-ConfigUpdate-ItemIEs']}, 'size': [(1, 'maxnoofCells')]}
    Items []ENBMBMSConfigurationdataConfigUpdateItemIEs
}

type ENBConfigurationUpdateAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ENBConfigurationUpdateAcknowledge-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ENBConfigurationUpdateAcknowledgeIes
}

func (self * ENBConfigurationUpdateAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ENBConfigurationUpdateAcknowledgeIes, order_ENBConfigurationUpdateAcknowledgeIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ENBConfigurationUpdateAcknowledge) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ENBConfigurationUpdateAcknowledgeIes, order_ENBConfigurationUpdateAcknowledgeIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type ENBConfigurationUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ENBConfigurationUpdateFailure-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ENBConfigurationUpdateFailureIes
}

func (self * ENBConfigurationUpdateFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ENBConfigurationUpdateFailureIes, order_ENBConfigurationUpdateFailureIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ENBConfigurationUpdateFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ENBConfigurationUpdateFailureIes, order_ENBConfigurationUpdateFailureIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MCEConfigurationUpdate struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MCEConfigurationUpdate-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MCEConfigurationUpdateIes
}

func (self * MCEConfigurationUpdate) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MCEConfigurationUpdateIes, order_MCEConfigurationUpdateIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MCEConfigurationUpdate) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MCEConfigurationUpdateIes, order_MCEConfigurationUpdateIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MCEConfigurationUpdateAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MCEConfigurationUpdateAcknowledge-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MCEConfigurationUpdateAcknowledgeIes
}

func (self * MCEConfigurationUpdateAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MCEConfigurationUpdateAcknowledgeIes, order_MCEConfigurationUpdateAcknowledgeIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MCEConfigurationUpdateAcknowledge) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MCEConfigurationUpdateAcknowledgeIes, order_MCEConfigurationUpdateAcknowledgeIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MCEConfigurationUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MCEConfigurationUpdateFailure-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MCEConfigurationUpdateFailureIes
}

func (self * MCEConfigurationUpdateFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MCEConfigurationUpdateFailureIes, order_MCEConfigurationUpdateFailureIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MCEConfigurationUpdateFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MCEConfigurationUpdateFailureIes, order_MCEConfigurationUpdateFailureIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type ErrorIndication struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ErrorIndication-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ErrorIndicationIes
}

func (self * ErrorIndication) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ErrorIndicationIes, order_ErrorIndicationIes} // p3
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
    ProtocolIEs := &ProtocolIEContainer {table_ErrorIndicationIes, order_ErrorIndicationIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type Reset struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['Reset-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ResetIes
}

func (self * Reset) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ResetIes, order_ResetIes} // p3
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
    ProtocolIEs := &ProtocolIEContainer {table_ResetIes, order_ResetIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *ResetType)Unpack(stream *Stream) {
    //coptions := []string{"m2-Interface","partOfM2-Interface"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in ResetType\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.M2Interface = &ResetAll{}//cho6
        self.M2Interface.Unpack(stream)
    } else if choice == 1 { //ch2
        self.PartOfM2Interface = &MBMSServiceassociatedLogicalM2ConnectionListRes{}//cho6
        self.PartOfM2Interface.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ResetType) Pack(stream *Stream) {
    if self.M2Interface != nil {
        stream.set_choice(0, 1, 1, 2)
        self.M2Interface.Pack(stream)//2
    } else if self.PartOfM2Interface != nil {
        stream.set_choice(1, 1, 1, 2)
        self.PartOfM2Interface.Pack(stream)//2
    }

}
type ResetType struct { //[{'type': 'ResetAll', 'name': 'm2-Interface'}, {'type': 'MBMS-Service-associatedLogicalM2-ConnectionListRes', 'name': 'partOfM2-Interface'}, None]
    M2Interface *ResetAll
    PartOfM2Interface *MBMSServiceassociatedLogicalM2ConnectionListRes
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
func (self *MBMSServiceassociatedLogicalM2ConnectionListRes) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]MBMSServiceassociatedLogicalM2ConnectionItemRes, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *MBMSServiceassociatedLogicalM2ConnectionListRes) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    //for item in table_MBMSServiceassociatedLogicalM2ConnectionItemRes:
    val := ProtocolIESingleContainer{table_MBMSServiceassociatedLogicalM2ConnectionItemRes, order_MBMSServiceassociatedLogicalM2ConnectionItemRes}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type MBMSServiceassociatedLogicalM2ConnectionListRes struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['MBMS-Service-associatedLogicalM2-ConnectionItemRes']}, 'size': [(1, 'maxNrOfIndividualM2ConnectionsToReset')]}
    Items []MBMSServiceassociatedLogicalM2ConnectionItemRes
}

type ResetAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ResetAcknowledge-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ResetAcknowledgeIes
}

func (self * ResetAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ResetAcknowledgeIes, order_ResetAcknowledgeIes} // p3
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
    ProtocolIEs := &ProtocolIEContainer {table_ResetAcknowledgeIes, order_ResetAcknowledgeIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *MBMSServiceassociatedLogicalM2ConnectionListResAck) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]MBMSServiceassociatedLogicalM2ConnectionItemResAck, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *MBMSServiceassociatedLogicalM2ConnectionListResAck) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    //for item in table_MBMSServiceassociatedLogicalM2ConnectionItemResAck:
    val := ProtocolIESingleContainer{table_MBMSServiceassociatedLogicalM2ConnectionItemResAck, order_MBMSServiceassociatedLogicalM2ConnectionItemResAck}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type MBMSServiceassociatedLogicalM2ConnectionListResAck struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['MBMS-Service-associatedLogicalM2-ConnectionItemResAck']}, 'size': [(1, 'maxNrOfIndividualM2ConnectionsToReset')]}
    Items []MBMSServiceassociatedLogicalM2ConnectionItemResAck
}

type PrivateMessage struct { // [{'type': 'PrivateIE-Container', 'actual-parameters': ['PrivateMessage-Ies'], 'name': 'privateIEs'}, None]
    PrivateIEs PrivateMessageIes
}

func (self * PrivateMessage) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    PrivateIEs := PrivateIEContainer {table_PrivateMessageIes, order_PrivateMessageIes} // p3
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
    PrivateIEs := &PrivateIEContainer {table_PrivateMessageIes, order_PrivateMessageIes} // p3
    PrivateIEs.Pack(stream, &self.PrivateIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MbmsServiceCountingRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MbmsServiceCountingRequest-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MbmsServiceCountingRequestIes
}

func (self * MbmsServiceCountingRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MbmsServiceCountingRequestIes, order_MbmsServiceCountingRequestIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MbmsServiceCountingRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MbmsServiceCountingRequestIes, order_MbmsServiceCountingRequestIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *MBMSCountingRequestSession) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(16)
    _size += 1
    self.Items = make([]MBMSCountingRequestSessionItem, _size)
    val := ProtocolIEContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *MBMSCountingRequestSession) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 16)
    //for item in table_MBMSCountingRequestSessionItem:
    val := ProtocolIEContainer{table_MBMSCountingRequestSessionItem, order_MBMSCountingRequestSessionItem}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type MBMSCountingRequestSession struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['MBMS-Counting-Request-Session-Item']}, 'size': [(1, 'maxnoofCountingService')]}
    Items []MBMSCountingRequestSessionItem
}

type MBMSCountingRequestSessionIE struct { // [{'type': 'TMGI', 'name': 'tmgi'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-Counting-Request-SessionIE-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    Tmgi TMGI
    IEExtensions *MBMSCountingRequestSessionIEExtIEs
}

func (self * MBMSCountingRequestSessionIE) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.Tmgi.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MBMSCountingRequestSessionIEExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-Counting-Request-SessionIE-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MBMSCountingRequestSessionIEExtIEs, order_MBMSCountingRequestSessionIEExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSCountingRequestSessionIE) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.Tmgi.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MBMSCountingRequestSessionIEExtIEs, order_MBMSCountingRequestSessionIEExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type MbmsServiceCountingResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MbmsServiceCountingResponse-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MbmsServiceCountingResponseIes
}

func (self * MbmsServiceCountingResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MbmsServiceCountingResponseIes, order_MbmsServiceCountingResponseIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MbmsServiceCountingResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MbmsServiceCountingResponseIes, order_MbmsServiceCountingResponseIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MbmsServiceCountingFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MbmsServiceCountingFailure-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MbmsServiceCountingFailureIes
}

func (self * MbmsServiceCountingFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MbmsServiceCountingFailureIes, order_MbmsServiceCountingFailureIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MbmsServiceCountingFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MbmsServiceCountingFailureIes, order_MbmsServiceCountingFailureIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MbmsServiceCountingResultsReport struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MbmsServiceCountingResultsReport-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MbmsServiceCountingResultsReportIes
}

func (self * MbmsServiceCountingResultsReport) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MbmsServiceCountingResultsReportIes, order_MbmsServiceCountingResultsReportIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MbmsServiceCountingResultsReport) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MbmsServiceCountingResultsReportIes, order_MbmsServiceCountingResultsReportIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *MBMSCountingResultList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(16)
    _size += 1
    self.Items = make([]MBMSCountingResultItem, _size)
    val := ProtocolIEContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *MBMSCountingResultList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 16)
    //for item in table_MBMSCountingResultItem:
    val := ProtocolIEContainer{table_MBMSCountingResultItem, order_MBMSCountingResultItem}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type MBMSCountingResultList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['MBMS-Counting-Result-Item']}, 'size': [(1, 'maxnoofCountingService')]}
    Items []MBMSCountingResultItem
}

type MBMSCountingResult struct { // [{'type': 'TMGI', 'name': 'tmgi'}, {'type': 'CountingResult', 'name': 'countingResult'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-Counting-Result-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    Tmgi TMGI
    CountingResult CountingResult
    IEExtensions *MBMSCountingResultExtIEs
}

func (self * MBMSCountingResult) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.Tmgi.Unpack(stream)// p8
    self.CountingResult.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MBMSCountingResultExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-Counting-Result-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MBMSCountingResultExtIEs, order_MBMSCountingResultExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSCountingResult) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.Tmgi.Pack(stream)
    self.CountingResult.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MBMSCountingResultExtIEs, order_MBMSCountingResultExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type CountingResult struct {
  Value uint64
}
func (self *CountingResult) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(1024, 10, 0, 0)
}
func (self * CountingResult) Pack(st *Stream){
    st.formatf_Integer(self.Value, 1024, 10, 0, 0)
}
type MbmsOverloadNotification struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MbmsOverloadNotification-Ies'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MbmsOverloadNotificationIes
}

func (self * MbmsOverloadNotification) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MbmsOverloadNotificationIes, order_MbmsOverloadNotificationIes} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MbmsOverloadNotification) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MbmsOverloadNotificationIes, order_MbmsOverloadNotificationIes} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *OverloadStatusPerPMCHList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(15)
    _size += 1
    self.Items = make([]OverloadStatusPerPMCHItem, _size)
    val := ProtocolIEContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *OverloadStatusPerPMCHList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 15)
    //for item in table_OverloadStatusPerPMCHItem:
    val := ProtocolIEContainer{table_OverloadStatusPerPMCHItem, order_OverloadStatusPerPMCHItem}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type OverloadStatusPerPMCHList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['Overload-Status-Per-PMCH-Item']}, 'size': [(1, 'maxnoofPMCHsperMBSFNarea')]}
    Items []OverloadStatusPerPMCHItem
}

type PMCHOverloadStatus struct {
  Value int
}
const (
    PMCHOverloadStatusnormal = 0
    PMCHOverloadStatusoverload = 1

    /* Extensions */
)
func (self *PMCHOverloadStatus) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *PMCHOverloadStatus) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
func (self *ActiveMBMSSessionList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(29)
    _size += 1
    self.Items = make([]ActiveMBMSSessionItem, _size)
    val := ProtocolIEContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *ActiveMBMSSessionList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 29)
    //for item in table_ActiveMBMSSessionItem:
    val := ProtocolIEContainer{table_ActiveMBMSSessionItem, order_ActiveMBMSSessionItem}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type ActiveMBMSSessionList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['Active-MBMS-Session-Item']}, 'size': [(1, 'maxnoofSessionsPerPMCH')]}
    Items []ActiveMBMSSessionItem
}

type AllocatedSubframesEnd struct {
  Value uint64
}
func (self *AllocatedSubframesEnd) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(1536, 11, 0, 0)
}
func (self * AllocatedSubframesEnd) Pack(st *Stream){
    st.formatf_Integer(self.Value, 1536, 11, 0, 0)
}
type AllocationAndRetentionPriority struct { // [{'type': 'PriorityLevel', 'name': 'priorityLevel'}, {'type': 'Pre-emptionCapability', 'name': 'pre-emptionCapability'}, {'type': 'Pre-emptionVulnerability', 'name': 'pre-emptionVulnerability'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AllocationAndRetentionPriority-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PriorityLevel PriorityLevel
    PreemptionCapability PreemptionCapability
    PreemptionVulnerability PreemptionVulnerability
    IEExtensions *AllocationAndRetentionPriorityExtIEs
}

func (self * AllocationAndRetentionPriority) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.PriorityLevel.Unpack(stream)// p8
    self.PreemptionCapability.Unpack(stream)// p8
    self.PreemptionVulnerability.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &AllocationAndRetentionPriorityExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AllocationAndRetentionPriority-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_AllocationAndRetentionPriorityExtIEs, order_AllocationAndRetentionPriorityExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * AllocationAndRetentionPriority) Pack(stream *Stream) {
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
        IEExtensions := &ProtocolExtensionContainer {table_AllocationAndRetentionPriorityExtIEs, order_AllocationAndRetentionPriorityExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type BitRate struct {
  Value uint64
}
func (self *BitRate) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(10000000001, 34, 0, 0)
}
func (self * BitRate) Pack(st *Stream){
    st.formatf_Integer(self.Value, 10000000001, 34, 0, 0)
}
func (self *Cause)Unpack(stream *Stream) {
    //coptions := []string{"radioNetwork","transport","nAS","protocol","misc","Unknown","Unknown","Unknown"}
    choice := stream.get_choice(3, 1, 5)
    choice_len := 0
    choice_loc := 0
    if choice >= 5 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in Cause\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.RadioNetwork = &CauseRadioNetwork{}//cho6
        self.RadioNetwork.Unpack(stream)
    } else if choice == 1 { //ch2
        self.Transport = &CauseTransport{}//cho6
        self.Transport.Unpack(stream)
    } else if choice == 2 { //ch2
        self.NAS = &CauseNAS{}//cho6
        self.NAS.Unpack(stream)
    } else if choice == 3 { //ch2
        self.Protocol = &CauseProtocol{}//cho6
        self.Protocol.Unpack(stream)
    } else if choice == 4 { //ch2
        self.Misc = &CauseMisc{}//cho6
        self.Misc.Unpack(stream)
    }//end of if else

    if choice >= 5 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * Cause) Pack(stream *Stream) {
    if self.RadioNetwork != nil {
        stream.set_choice(0, 3, 1, 5)
        self.RadioNetwork.Pack(stream)//2
    } else if self.Transport != nil {
        stream.set_choice(1, 3, 1, 5)
        self.Transport.Pack(stream)//2
    } else if self.NAS != nil {
        stream.set_choice(2, 3, 1, 5)
        self.NAS.Pack(stream)//2
    } else if self.Protocol != nil {
        stream.set_choice(3, 3, 1, 5)
        self.Protocol.Pack(stream)//2
    } else if self.Misc != nil {
        stream.set_choice(4, 3, 1, 5)
        self.Misc.Pack(stream)//2
    }

}
type Cause struct { //[{'type': 'CauseRadioNetwork', 'name': 'radioNetwork'}, {'type': 'CauseTransport', 'name': 'transport'}, {'type': 'CauseNAS', 'name': 'nAS'}, {'type': 'CauseProtocol', 'name': 'protocol'}, {'type': 'CauseMisc', 'name': 'misc'}, None]
    RadioNetwork *CauseRadioNetwork
    Transport *CauseTransport
    NAS *CauseNAS
    Protocol *CauseProtocol
    Misc *CauseMisc
} // Cause

type CauseMisc struct {
  Value int
}
const (
    CauseMisccontrol_processing_overload = 0
    CauseMischardware_failure = 1
    CauseMiscom_intervention = 2
    CauseMiscunspecified = 3

    /* Extensions */
)
func (self *CauseMisc) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 4, 1)
}
func (self *CauseMisc) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 4, 1)
}
type CauseNAS struct {
  Value int
}
const (
    CauseNASunspecified = 0

    /* Extensions */
)
func (self *CauseNAS) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *CauseNAS) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
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
    CauseProtocolabstract_syntax_error_falsely_constructed_message = 5
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
    CauseRadioNetworkunknown_or_already_allocated_MCE_MBMS_M2AP_ID = 0
    CauseRadioNetworkunknown_or_already_allocated_eNB_MBMS_M2AP_ID = 1
    CauseRadioNetworkunknown_or_inconsistent_pair_of_MBMS_M2AP_IDs = 2
    CauseRadioNetworkradio_resources_not_available = 3
    CauseRadioNetworkinteraction_with_other_procedure = 4
    CauseRadioNetworkunspecified = 5

    /* Extensions */
    CauseRadioNetworkinvalid_QoS_combination = 6
    CauseRadioNetworknot_supported_QCI_value = 7
)
func (self *CauseRadioNetwork) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 6, 1)
}
func (self *CauseRadioNetwork) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 6, 1)
}
type CauseTransport struct {
  Value int
}
const (
    CauseTransporttransport_resource_unavailable = 0
    CauseTransportunspecified = 1

    /* Extensions */
)
func (self *CauseTransport) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *CauseTransport) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type CellInformation struct { // [{'type': 'ECGI', 'name': 'eCGI'}, {'type': 'ENUMERATED', 'values': [('reservedCell', 0), ('nonReservedCell', 1), None], 'name': 'cellReservationInfo'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Cell-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    ECGI ECGI
    CellReservationInfo ENUMERATED
    IEExtensions *CellInformationExtIEs
}

func (self * CellInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.ECGI.Unpack(stream)// p8
    var Unpack_cellReservationInfo = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_cellReservationInfo(stream, &self.CellReservationInfo)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &CellInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Cell-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_CellInformationExtIEs, order_CellInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CellInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.ECGI.Pack(stream)
    var Pack_cellReservationInfo = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_cellReservationInfo(stream, self.CellReservationInfo) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_CellInformationExtIEs, order_CellInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *CellInformationList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]CellInformation, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *CellInformationList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type CellInformationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'Cell-Information'}, 'size': [(1, 'maxnoofCells')]}
    Items []CellInformation
}

type CriticalityDiagnostics struct { // [{'type': 'ProcedureCode', 'name': 'procedureCode', 'optional': True}, {'type': 'TriggeringMessage', 'name': 'triggeringMessage', 'optional': True}, {'type': 'Criticality', 'name': 'procedureCriticality', 'optional': True}, {'type': 'CriticalityDiagnostics-IE-List', 'name': 'iEsCriticalityDiagnostics', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    ProcedureCode *ProcedureCode
    TriggeringMessage *TriggeringMessage
    ProcedureCriticality *Criticality
    IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList
    IEExtensions *CriticalityDiagnosticsExtIEs
}

func (self * CriticalityDiagnostics) Unpack(stream *Stream) {
    procedureCode_flag := 0x00000002
    triggeringMessage_flag := 0x00000004
    procedureCriticality_flag := 0x00000008
    iEsCriticalityDiagnostics_flag := 0x00000010
    iEExtensions_flag := 0x00000020
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(6)
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
    const iEsCriticalityDiagnostics_flag uint = 0x00000010
    const iEExtensions_flag uint = 0x00000020
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(6)
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
    stream.set_flags(_flags, _flagReserve, 6)
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
type CriticalityDiagnosticsIEList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'Criticality', 'name': 'iECriticality'}, {'type': 'ProtocolIE-ID', 'name': 'iE-ID'}, {'type': 'TypeOfError', 'name': 'typeOfError'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-IE-List-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxnooferrors')]}
    Items []CriticalityDiagnosticsIEList_Item
}

type ECGI struct { // [{'type': 'PLMN-Identity', 'name': 'pLMN-Identity'}, {'type': 'EUTRANCellIdentifier', 'name': 'eUTRANcellIdentifier'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ECGI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNIdentity PLMNIdentity
    EUTRANcellIdentifier EUTRANCellIdentifier
    IEExtensions *ECGIExtIEs
}

func (self * ECGI) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PLMNIdentity.Unpack(stream)// p8
    self.EUTRANcellIdentifier.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &ECGIExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ECGI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_ECGIExtIEs, order_ECGIExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ECGI) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.EUTRANcellIdentifier.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_ECGIExtIEs, order_ECGIExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *ENBID)Unpack(stream *Stream) {
    //coptions := []string{"macro-eNB-ID","short-Macro-eNB-ID","long-Macro-eNB-ID"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in ENBID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_macroeNBID = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(20, 20)
        }
        self.MacroeNBID = &BITSTRING{}//cho5
        Unpack_macroeNBID(stream, self.MacroeNBID);
    } else if choice == 1 { //ch2
        var Unpack_shortMacroeNBID = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(18, 18)
        }
        self.ShortMacroeNBID = &BITSTRING{}//cho5
        Unpack_shortMacroeNBID(stream, self.ShortMacroeNBID);
    } else if choice == 2 { //ch2
        var Unpack_longMacroeNBID = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(21, 21)
        }
        self.LongMacroeNBID = &BITSTRING{}//cho5
        Unpack_longMacroeNBID(stream, self.LongMacroeNBID);
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ENBID) Pack(stream *Stream) {
    if self.MacroeNBID != nil {
        stream.set_choice(0, 0, 1, 1)
        var Pack_macroeNBID = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 20)
        }
        Pack_macroeNBID(stream, *self.MacroeNBID)//3
    } else if self.ShortMacroeNBID != nil {
        stream.set_choice(1, 0, 1, 1)
        lenLoc := stream.reserve_len()
        var Pack_shortMacroeNBID = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 18)
        }
        Pack_shortMacroeNBID(stream, *self.ShortMacroeNBID)//3
        stream.set_len(lenLoc)
    } else if self.LongMacroeNBID != nil {
        stream.set_choice(2, 0, 1, 1)
        lenLoc := stream.reserve_len()
        var Pack_longMacroeNBID = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 21)
        }
        Pack_longMacroeNBID(stream, *self.LongMacroeNBID)//3
        stream.set_len(lenLoc)
    }

}
type ENBID struct { //[{'type': 'BIT STRING', 'size': [20], 'name': 'macro-eNB-ID'}, None, {'type': 'BIT STRING', 'size': [18], 'name': 'short-Macro-eNB-ID'}, {'type': 'BIT STRING', 'size': [21], 'name': 'long-Macro-eNB-ID'}]
    MacroeNBID *BITSTRING
    ShortMacroeNBID *BITSTRING
    LongMacroeNBID *BITSTRING
} // ENBID

type ENBMBMSConfigurationdataItem struct { // [{'type': 'ECGI', 'name': 'eCGI'}, {'type': 'MBSFN-SynchronisationArea-ID', 'name': 'mbsfnSynchronisationArea'}, {'type': 'MBMS-Service-Area-ID-List', 'name': 'mbmsServiceAreaList'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ENB-MBMS-Configuration-data-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    ECGI ECGI
    MbsfnSynchronisationArea MBSFNSynchronisationAreaID
    MbmsServiceAreaList MBMSServiceAreaIDList
    IEExtensions *ENBMBMSConfigurationdataItemExtIEs
}

func (self * ENBMBMSConfigurationdataItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.ECGI.Unpack(stream)// p8
    self.MbsfnSynchronisationArea.Unpack(stream)// p8
    self.MbmsServiceAreaList.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &ENBMBMSConfigurationdataItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ENB-MBMS-Configuration-data-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_ENBMBMSConfigurationdataItemExtIEs, order_ENBMBMSConfigurationdataItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ENBMBMSConfigurationdataItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.ECGI.Pack(stream)
    self.MbsfnSynchronisationArea.Pack(stream)
    self.MbmsServiceAreaList.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_ENBMBMSConfigurationdataItemExtIEs, order_ENBMBMSConfigurationdataItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *ENBMBMSConfigurationdataConfigUpdateItem)Unpack(stream *Stream) {
    //coptions := []string{"mBMSConfigData","eCGI"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in ENBMBMSConfigurationdataConfigUpdateItem\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.MBMSConfigData = &ENBMBMSConfigurationdataItem{}//cho6
        self.MBMSConfigData.Unpack(stream)
    } else if choice == 1 { //ch2
        self.ECGI = &ECGI{}//cho6
        self.ECGI.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ENBMBMSConfigurationdataConfigUpdateItem) Pack(stream *Stream) {
    if self.MBMSConfigData != nil {
        stream.set_choice(0, 1, 1, 2)
        self.MBMSConfigData.Pack(stream)//2
    } else if self.ECGI != nil {
        stream.set_choice(1, 1, 1, 2)
        self.ECGI.Pack(stream)//2
    }

}
type ENBMBMSConfigurationdataConfigUpdateItem struct { //[{'type': 'ENB-MBMS-Configuration-data-Item', 'name': 'mBMSConfigData'}, {'type': 'ECGI', 'name': 'eCGI'}, None]
    MBMSConfigData *ENBMBMSConfigurationdataItem
    ECGI *ECGI
} // ENBMBMSConfigurationdataConfigUpdateItem

type ENBMBMSM2APID struct {
  Value uint64
}
func (self *ENBMBMSM2APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * ENBMBMSM2APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type ENBname struct {
  Value string
}
func (self *ENBname) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in ENBname")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *ENBname) Pack(st *Stream) {
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
type EUTRANCellIdentifier struct {
  Len int
  Value HexBytes
}
func (self *EUTRANCellIdentifier) Unpack(st *Stream){
    self.Value = st.parsef_BitString(28, 28)
}
func (self *EUTRANCellIdentifier) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 28)
}
type GBRQosInformation struct { // [{'type': 'BitRate', 'name': 'mBMS-E-RAB-MaximumBitrateDL'}, {'type': 'BitRate', 'name': 'mBMS-E-RAB-GuaranteedBitrateDL'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GBR-QosInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    MBMSERABMaximumBitrateDL BitRate
    MBMSERABGuaranteedBitrateDL BitRate
    IEExtensions *GBRQosInformationExtIEs
}

func (self * GBRQosInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.MBMSERABMaximumBitrateDL.Unpack(stream)// p8
    self.MBMSERABGuaranteedBitrateDL.Unpack(stream)// p8
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
    self.MBMSERABMaximumBitrateDL.Pack(stream)
    self.MBMSERABGuaranteedBitrateDL.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GBRQosInformationExtIEs, order_GBRQosInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GlobalENBID struct { // [{'type': 'PLMN-Identity', 'name': 'pLMN-Identity'}, {'type': 'ENB-ID', 'name': 'eNB-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GlobalENB-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNIdentity PLMNIdentity
    ENBID ENBID
    IEExtensions *GlobalENBIDExtIEs
}

func (self * GlobalENBID) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PLMNIdentity.Unpack(stream)// p8
    self.ENBID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GlobalENBIDExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GlobalENB-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GlobalENBIDExtIEs, order_GlobalENBIDExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalENBID) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.ENBID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GlobalENBIDExtIEs, order_GlobalENBIDExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GlobalMCEID struct { // [{'type': 'PLMN-Identity', 'name': 'pLMN-Identity'}, {'type': 'MCE-ID', 'name': 'mCE-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GlobalMCE-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNIdentity PLMNIdentity
    MCEID MCEID
    IEExtensions *GlobalMCEIDExtIEs
}

func (self * GlobalMCEID) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PLMNIdentity.Unpack(stream)// p8
    self.MCEID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GlobalMCEIDExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GlobalMCE-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GlobalMCEIDExtIEs, order_GlobalMCEIDExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalMCEID) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.MCEID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GlobalMCEIDExtIEs, order_GlobalMCEIDExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
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
type IPAddress struct {
  Value HexBytes
}
func (self *IPAddress) Unpack(st *Stream) {
    _len := st.parse_olen(4)+4
    if _len < 4 || _len > 16 {
        //fmt.Println ("Invalid len in IPAddress")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *IPAddress) Pack(st *Stream) {
    if len(self.Value) < 4 || len(self.Value) > 16 {
        log.Error ("Invalid len in IPAddress")
        return
}
    st.format_olen((len(self.Value))-4, 4)
    st.formatf_OctString(self.Value, 0)
}
type LCID struct {
  Value uint64
}
func (self *LCID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(29, 5, 0, 0)
}
func (self * LCID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 29, 5, 0, 0)
}
func (self *MBMSCellList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(4096)
    _size += 1
    self.Items = make([]ECGI, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MBMSCellList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 4096)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MBMSCellList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ECGI'}, 'size': [(1, 'maxnoofCellsforMBMS')]}
    Items []ECGI
}

type MBMSERABQoSParameters struct { // [{'type': 'QCI', 'name': 'qCI'}, {'type': 'GBR-QosInformation', 'name': 'gbrQosInformation', 'optional': True}, {'type': 'AllocationAndRetentionPriority', 'name': 'allocationAndRetentionPriority'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-E-RAB-QoS-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    QCI QCI
    GbrQosInformation *GBRQosInformation
    AllocationAndRetentionPriority AllocationAndRetentionPriority
    IEExtensions *MBMSERABQoSParametersExtIEs
}

func (self * MBMSERABQoSParameters) Unpack(stream *Stream) {
    gbrQosInformation_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.QCI.Unpack(stream)// p8
    if (gbrQosInformation_flag & _flags) == gbrQosInformation_flag { //cond2
        self.GbrQosInformation = &GBRQosInformation{}//7{'type': 'GBR-QosInformation', 'name': 'gbrQosInformation', 'optional': True}
        self.GbrQosInformation.Unpack(stream)// p8
    }
    self.AllocationAndRetentionPriority.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MBMSERABQoSParametersExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-E-RAB-QoS-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MBMSERABQoSParametersExtIEs, order_MBMSERABQoSParametersExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSERABQoSParameters) Pack(stream *Stream) {
    const gbrQosInformation_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.QCI.Pack(stream)
    if self.GbrQosInformation != nil { 
        _flags |= gbrQosInformation_flag
        self.GbrQosInformation.Pack(stream)
    }//end of optional
    self.AllocationAndRetentionPriority.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MBMSERABQoSParametersExtIEs, order_MBMSERABQoSParametersExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type MBMSServiceassociatedLogicalM2ConnectionItem struct { // [{'type': 'ENB-MBMS-M2AP-ID', 'name': 'eNB-MBMS-M2AP-ID', 'optional': True}, {'type': 'MCE-MBMS-M2AP-ID', 'name': 'mCE-MBMS-M2AP-ID', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-Service-associatedLogicalM2-ConnectionItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    ENBMBMSM2APID *ENBMBMSM2APID
    MCEMBMSM2APID *MCEMBMSM2APID
    IEExtensions *MBMSServiceassociatedLogicalM2ConnectionItemExtIEs
}

func (self * MBMSServiceassociatedLogicalM2ConnectionItem) Unpack(stream *Stream) {
    eNBMBMSM2APID_flag := 0x00000002
    mCEMBMSM2APID_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (eNBMBMSM2APID_flag & _flags) == eNBMBMSM2APID_flag { //cond2
        self.ENBMBMSM2APID = &ENBMBMSM2APID{}//7{'type': 'ENB-MBMS-M2AP-ID', 'name': 'eNB-MBMS-M2AP-ID', 'optional': True}
        self.ENBMBMSM2APID.Unpack(stream)// p8
    }
    if (mCEMBMSM2APID_flag & _flags) == mCEMBMSM2APID_flag { //cond2
        self.MCEMBMSM2APID = &MCEMBMSM2APID{}//7{'type': 'MCE-MBMS-M2AP-ID', 'name': 'mCE-MBMS-M2AP-ID', 'optional': True}
        self.MCEMBMSM2APID.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MBMSServiceassociatedLogicalM2ConnectionItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-Service-associatedLogicalM2-ConnectionItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MBMSServiceassociatedLogicalM2ConnectionItemExtIEs, order_MBMSServiceassociatedLogicalM2ConnectionItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSServiceassociatedLogicalM2ConnectionItem) Pack(stream *Stream) {
    const eNBMBMSM2APID_flag uint = 0x00000002
    const mCEMBMSM2APID_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.ENBMBMSM2APID != nil { 
        _flags |= eNBMBMSM2APID_flag
        self.ENBMBMSM2APID.Pack(stream)
    }//end of optional
    if self.MCEMBMSM2APID != nil { 
        _flags |= mCEMBMSM2APID_flag
        self.MCEMBMSM2APID.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MBMSServiceassociatedLogicalM2ConnectionItemExtIEs, order_MBMSServiceassociatedLogicalM2ConnectionItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type MBMSServiceArea struct {
  Value HexBytes
}
func (self *MBMSServiceArea) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *MBMSServiceArea) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
func (self *MBMSServiceAreaIDList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]MBMSServiceArea, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MBMSServiceAreaIDList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MBMSServiceAreaIDList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MBMS-Service-Area'}, 'size': [(1, 'maxnoofMBMSServiceAreasPerCell')]}
    Items []MBMSServiceArea
}

type MBMSSessionID struct {
  Value HexBytes
}
func (self *MBMSSessionID) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(1)
}
func (self *MBMSSessionID) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 1)
}
func (self *MBMSsessionListPerPMCHItem) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(29)
    _size += 1
    self.Items = make([]MBMSsessionListPerPMCHItem_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *MBMSsessionListPerPMCHItem_Item) { //[{'type': 'TMGI', 'name': 'tmgi'}, {'type': 'LCID', 'name': 'lcid'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMSsessionListPerPMCH-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        iEExtensions_flag := 0x00000002
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(2)
        self.Tmgi.Unpack(stream)// p8
        self.Lcid.Unpack(stream)// p8
        if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
            self.IEExtensions = &MBMSsessionListPerPMCHItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMSsessionListPerPMCH-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
            IEExtensions := ProtocolExtensionContainer {table_MBMSsessionListPerPMCHItemExtIEs, order_MBMSsessionListPerPMCHItemExtIEs} // p3
            IEExtensions.Unpack(stream, &self.IEExtensions) // p3
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    for item := 0; item <_size; item +=1 {
        Unpack_Item(stream, &self.Items[item])
    }
}


func (self *MBMSsessionListPerPMCHItem) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 29)
    var Pack_Item = func(stream *Stream, self MBMSsessionListPerPMCHItem_Item) {//seq
        const iEExtensions_flag uint = 0x00000002
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(2)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        self.Tmgi.Pack(stream)
        self.Lcid.Pack(stream)
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_MBMSsessionListPerPMCHItemExtIEs, order_MBMSsessionListPerPMCHItemExtIEs} // p3
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


type MBMSsessionListPerPMCHItem_Item struct { // [{'type': 'TMGI', 'name': 'tmgi'}, {'type': 'LCID', 'name': 'lcid'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMSsessionListPerPMCH-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    Tmgi TMGI
    Lcid LCID
    IEExtensions *MBMSsessionListPerPMCHItemExtIEs
}
type MBMSsessionListPerPMCHItem struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'TMGI', 'name': 'tmgi'}, {'type': 'LCID', 'name': 'lcid'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMSsessionListPerPMCH-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxnoofSessionsPerPMCH')]}
    Items []MBMSsessionListPerPMCHItem_Item
}

func (self *MBMSsessionsToBeSuspendedListPerPMCHItem) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(29)
    _size += 1
    self.Items = make([]MBMSsessionsToBeSuspendedListPerPMCHItem_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *MBMSsessionsToBeSuspendedListPerPMCHItem_Item) { //[{'type': 'TMGI', 'name': 'tmgi'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMSsessionsToBeSuspendedListPerPMCH-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        iEExtensions_flag := 0x00000002
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(2)
        self.Tmgi.Unpack(stream)// p8
        if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
            self.IEExtensions = &MBMSsessionsToBeSuspendedListPerPMCHItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMSsessionsToBeSuspendedListPerPMCH-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
            IEExtensions := ProtocolExtensionContainer {table_MBMSsessionsToBeSuspendedListPerPMCHItemExtIEs, order_MBMSsessionsToBeSuspendedListPerPMCHItemExtIEs} // p3
            IEExtensions.Unpack(stream, &self.IEExtensions) // p3
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    for item := 0; item <_size; item +=1 {
        Unpack_Item(stream, &self.Items[item])
    }
}


func (self *MBMSsessionsToBeSuspendedListPerPMCHItem) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 29)
    var Pack_Item = func(stream *Stream, self MBMSsessionsToBeSuspendedListPerPMCHItem_Item) {//seq
        const iEExtensions_flag uint = 0x00000002
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(2)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        self.Tmgi.Pack(stream)
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_MBMSsessionsToBeSuspendedListPerPMCHItemExtIEs, order_MBMSsessionsToBeSuspendedListPerPMCHItemExtIEs} // p3
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


type MBMSsessionsToBeSuspendedListPerPMCHItem_Item struct { // [{'type': 'TMGI', 'name': 'tmgi'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMSsessionsToBeSuspendedListPerPMCH-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    Tmgi TMGI
    IEExtensions *MBMSsessionsToBeSuspendedListPerPMCHItemExtIEs
}
type MBMSsessionsToBeSuspendedListPerPMCHItem struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'TMGI', 'name': 'tmgi'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMSsessionsToBeSuspendedListPerPMCH-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxnoofSessionsPerPMCH')]}
    Items []MBMSsessionsToBeSuspendedListPerPMCHItem_Item
}

type MBSFNAreaID struct {
  Value uint64
}
func (self *MBSFNAreaID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * MBSFNAreaID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
type MBSFNSynchronisationAreaID struct {
  Value uint64
}
func (self *MBSFNSynchronisationAreaID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * MBSFNSynchronisationAreaID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type MBSFNSubframeConfiguration_SubframeAllocation struct { //[{'type': 'BIT STRING', 'size': [6], 'name': 'oneFrame'}, {'type': 'BIT STRING', 'size': [24], 'name': 'fourFrames'}]
    OneFrame *BITSTRING
    FourFrames *BITSTRING
} // MBSFNSubframeConfiguration_SubframeAllocation

type MBSFNSubframeConfiguration struct { // [{'type': 'ENUMERATED', 'values': [('n1', 0), ('n2', 1), ('n4', 2), ('n8', 3), ('n16', 4), ('n32', 5)], 'name': 'radioframeAllocationPeriod'}, {'type': 'INTEGER', 'restricted-to': [(0, 7)], 'name': 'radioframeAllocationOffset'}, {'type': 'CHOICE', 'members': [{'type': 'BIT STRING', 'size': [6], 'name': 'oneFrame'}, {'type': 'BIT STRING', 'size': [24], 'name': 'fourFrames'}], 'name': 'subframeAllocation'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBSFN-Subframe-Configuration-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RadioframeAllocationPeriod ENUMERATED
    RadioframeAllocationOffset INTEGER
    SubframeAllocation MBSFNSubframeConfiguration_SubframeAllocation
    IEExtensions *MBSFNSubframeConfigurationExtIEs
}

func (self * MBSFNSubframeConfiguration) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_radioframeAllocationPeriod = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(3, 6, 0)
    }
    Unpack_radioframeAllocationPeriod(stream, &self.RadioframeAllocationPeriod)// p2
    var Unpack_radioframeAllocationOffset = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(8, 3, 0, 0)
    }
    Unpack_radioframeAllocationOffset(stream, &self.RadioframeAllocationOffset)// p2
    var Unpack_subframeAllocation = func(stream *Stream, self *MBSFNSubframeConfiguration_SubframeAllocation) {
        //coptions := []string{"oneFrame","fourFrames"}
        choice := stream.get_choice(1, 0, 2)
        if choice == 0 { //ch1
            var Unpack_oneFrame = func(st *Stream, self *BITSTRING){
                self.Value = st.parsef_BitString(6, 6)
            }
            self.OneFrame = &BITSTRING{}//cho5
            Unpack_oneFrame(stream, self.OneFrame);
        } else if choice == 1 { //ch2
            var Unpack_fourFrames = func(st *Stream, self *BITSTRING){
                self.Value = st.parsef_BitString(24, 24)
            }
            self.FourFrames = &BITSTRING{}//cho5
            Unpack_fourFrames(stream, self.FourFrames);
        }//end of if else

    }
    Unpack_subframeAllocation(stream, &self.SubframeAllocation)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MBSFNSubframeConfigurationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBSFN-Subframe-Configuration-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MBSFNSubframeConfigurationExtIEs, order_MBSFNSubframeConfigurationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBSFNSubframeConfiguration) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_radioframeAllocationPeriod = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 3, 6, 0)
    }
    Pack_radioframeAllocationPeriod(stream, self.RadioframeAllocationPeriod) //f2
    var Pack_radioframeAllocationOffset = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 8, 3, 0, 0)
    }
    Pack_radioframeAllocationOffset(stream, self.RadioframeAllocationOffset) //f2
    var Pack_subframeAllocation = func(stream *Stream, self MBSFNSubframeConfiguration_SubframeAllocation) {
        if self.OneFrame != nil {
            stream.set_choice(0, 1, 0, 2)
            var Pack_oneFrame = func(st *Stream, self BITSTRING) {
                st.formatf_BitString(self.Value, 6)
            }
            Pack_oneFrame(stream, *self.OneFrame)//3
        } else if self.FourFrames != nil {
            stream.set_choice(1, 1, 0, 2)
            var Pack_fourFrames = func(st *Stream, self BITSTRING) {
                st.formatf_BitString(self.Value, 24)
            }
            Pack_fourFrames(stream, *self.FourFrames)//3
        }

    }
    Pack_subframeAllocation(stream, self.SubframeAllocation) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MBSFNSubframeConfigurationExtIEs, order_MBSFNSubframeConfigurationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type MCCHUpdateTime struct {
  Value uint64
}
func (self *MCCHUpdateTime) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * MCCHUpdateTime) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
type MCCHrelatedBCCHConfigPerMBSFNAreaItem struct { // [{'type': 'MBSFN-Area-ID', 'name': 'mbsfnArea'}, {'type': 'ENUMERATED', 'values': [('s1', 0), ('s2', 1), None], 'name': 'pdcchLength'}, {'type': 'ENUMERATED', 'values': [('rf32', 0), ('rf64', 1), ('rf128', 2), ('rf256', 3)], 'name': 'repetitionPeriod'}, {'type': 'INTEGER', 'restricted-to': [(0, 10)], 'name': 'offset'}, {'type': 'ENUMERATED', 'values': [('rf512', 0), ('rf1024', 1)], 'name': 'modificationPeriod'}, {'type': 'BIT STRING', 'size': [6], 'name': 'subframeAllocationInfo'}, {'type': 'ENUMERATED', 'values': [('n2', 0), ('n7', 1), ('n13', 2), ('n19', 3)], 'name': 'modulationAndCodingScheme'}, {'type': 'Cell-Information-List', 'name': 'cellInformationList', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MCCHrelatedBCCH-ConfigPerMBSFNArea-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    MbsfnArea MBSFNAreaID
    PdcchLength ENUMERATED
    RepetitionPeriod ENUMERATED
    Offset INTEGER
    ModificationPeriod ENUMERATED
    SubframeAllocationInfo BITSTRING
    ModulationAndCodingScheme ENUMERATED
    CellInformationList *CellInformationList
    IEExtensions *MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs
}

func (self * MCCHrelatedBCCHConfigPerMBSFNAreaItem) Unpack(stream *Stream) {
    cellInformationList_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.MbsfnArea.Unpack(stream)// p8
    var Unpack_pdcchLength = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_pdcchLength(stream, &self.PdcchLength)// p2
    var Unpack_repetitionPeriod = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 4, 0)
    }
    Unpack_repetitionPeriod(stream, &self.RepetitionPeriod)// p2
    var Unpack_offset = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(11, 4, 0, 0)
    }
    Unpack_offset(stream, &self.Offset)// p2
    var Unpack_modificationPeriod = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(1, 2, 0)
    }
    Unpack_modificationPeriod(stream, &self.ModificationPeriod)// p2
    var Unpack_subframeAllocationInfo = func(st *Stream, self *BITSTRING){
        self.Value = st.parsef_BitString(6, 6)
    }
    Unpack_subframeAllocationInfo(stream, &self.SubframeAllocationInfo)// p2
    var Unpack_modulationAndCodingScheme = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 4, 0)
    }
    Unpack_modulationAndCodingScheme(stream, &self.ModulationAndCodingScheme)// p2
    if (cellInformationList_flag & _flags) == cellInformationList_flag { //cond2
        self.CellInformationList = &CellInformationList{}//7{'type': 'Cell-Information-List', 'name': 'cellInformationList', 'optional': True}
        self.CellInformationList.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MCCHrelatedBCCH-ConfigPerMBSFNArea-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs, order_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MCCHrelatedBCCHConfigPerMBSFNAreaItem) Pack(stream *Stream) {
    const cellInformationList_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MbsfnArea.Pack(stream)
    var Pack_pdcchLength = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_pdcchLength(stream, self.PdcchLength) //f2
    var Pack_repetitionPeriod = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 4, 0)
    }
    Pack_repetitionPeriod(stream, self.RepetitionPeriod) //f2
    var Pack_offset = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 11, 4, 0, 0)
    }
    Pack_offset(stream, self.Offset) //f2
    var Pack_modificationPeriod = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 1, 2, 0)
    }
    Pack_modificationPeriod(stream, self.ModificationPeriod) //f2
    var Pack_subframeAllocationInfo = func(st *Stream, self BITSTRING) {
        st.formatf_BitString(self.Value, 6)
    }
    Pack_subframeAllocationInfo(stream, self.SubframeAllocationInfo) //f2
    var Pack_modulationAndCodingScheme = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 4, 0)
    }
    Pack_modulationAndCodingScheme(stream, self.ModulationAndCodingScheme) //f2
    if self.CellInformationList != nil { 
        _flags |= cellInformationList_flag
        self.CellInformationList.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs, order_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type MCEID struct {
  Value HexBytes
}
func (self *MCEID) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(2)
}
func (self *MCEID) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 2)
}
type MCEMBMSM2APID struct {
  Value uint64
}
func (self *MCEMBMSM2APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16777216, 24, 0, 0)
}
func (self * MCEMBMSM2APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16777216, 24, 0, 0)
}
type MCEname struct {
  Value string
}
func (self *MCEname) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in MCEname")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *MCEname) Pack(st *Stream) {
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
type MCHSchedulingPeriod struct {
  Value int
}
const (
    MCHSchedulingPeriodrf8 = 0
    MCHSchedulingPeriodrf16 = 1
    MCHSchedulingPeriodrf32 = 2
    MCHSchedulingPeriodrf64 = 3
    MCHSchedulingPeriodrf128 = 4
    MCHSchedulingPeriodrf256 = 5
    MCHSchedulingPeriodrf512 = 6
    MCHSchedulingPeriodrf1024 = 7
)
func (self *MCHSchedulingPeriod) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 8, 0)
}
func (self *MCHSchedulingPeriod) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 8, 0)
}
type MCHSchedulingPeriodExtended struct {
  Value int
}
const (
    MCHSchedulingPeriodExtendedrf4 = 0

    /* Extensions */
)
func (self *MCHSchedulingPeriodExtended) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *MCHSchedulingPeriodExtended) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
type MCHSchedulingPeriodExtended2 struct {
  Value int
}
const (
    MCHSchedulingPeriodExtended2rf1 = 0
    MCHSchedulingPeriodExtended2rf2 = 1

    /* Extensions */
)
func (self *MCHSchedulingPeriodExtended2) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *MCHSchedulingPeriodExtended2) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type ModulationCodingScheme2 struct {
  Value uint64
}
func (self *ModulationCodingScheme2) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(28, 5, 0, 0)
}
func (self * ModulationCodingScheme2) Pack(st *Stream){
    st.formatf_Integer(self.Value, 28, 5, 0, 0)
}
type ModificationPeriodExtended struct {
  Value int
}
const (
    ModificationPeriodExtendedrf1 = 0
    ModificationPeriodExtendedrf2 = 1
    ModificationPeriodExtendedrf4 = 2
    ModificationPeriodExtendedrf8 = 3
    ModificationPeriodExtendedrf16 = 4
    ModificationPeriodExtendedrf32 = 5
    ModificationPeriodExtendedrf64 = 6
    ModificationPeriodExtendedrf128 = 7
    ModificationPeriodExtendedrf256 = 8

    /* Extensions */
)
func (self *ModificationPeriodExtended) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(5, 9, 1)
}
func (self *ModificationPeriodExtended) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 5, 9, 1)
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
type PMCHConfiguration struct { // [{'type': 'AllocatedSubframesEnd', 'name': 'allocatedSubframesEnd'}, {'type': 'INTEGER', 'restricted-to': [(0, 28)], 'name': 'dataMCS'}, {'type': 'MCH-Scheduling-Period', 'name': 'mchSchedulingPeriod'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PMCH-Configuration-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    AllocatedSubframesEnd AllocatedSubframesEnd
    DataMCS INTEGER
    MchSchedulingPeriod MCHSchedulingPeriod
    IEExtensions *PMCHConfigurationExtIEs
}

func (self * PMCHConfiguration) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.AllocatedSubframesEnd.Unpack(stream)// p8
    var Unpack_dataMCS = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(29, 5, 0, 0)
    }
    Unpack_dataMCS(stream, &self.DataMCS)// p2
    self.MchSchedulingPeriod.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PMCHConfigurationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PMCH-Configuration-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PMCHConfigurationExtIEs, order_PMCHConfigurationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PMCHConfiguration) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.AllocatedSubframesEnd.Pack(stream)
    var Pack_dataMCS = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 29, 5, 0, 0)
    }
    Pack_dataMCS(stream, self.DataMCS) //f2
    self.MchSchedulingPeriod.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PMCHConfigurationExtIEs, order_PMCHConfigurationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type CommonSubframeAllocationPeriod struct {
  Value int
}
const (
    CommonSubframeAllocationPeriodrf4 = 0
    CommonSubframeAllocationPeriodrf8 = 1
    CommonSubframeAllocationPeriodrf16 = 2
    CommonSubframeAllocationPeriodrf32 = 3
    CommonSubframeAllocationPeriodrf64 = 4
    CommonSubframeAllocationPeriodrf128 = 5
    CommonSubframeAllocationPeriodrf256 = 6
)
func (self *CommonSubframeAllocationPeriod) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 7, 0)
}
func (self *CommonSubframeAllocationPeriod) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 7, 0)
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
type PriorityLevel struct {
  Value uint64
}
func (self *PriorityLevel) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16, 4, 0, 0)
}
func (self * PriorityLevel) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16, 4, 0, 0)
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
type RepetitionPeriodExtended struct {
  Value int
}
const (
    RepetitionPeriodExtendedrf1 = 0
    RepetitionPeriodExtendedrf2 = 1
    RepetitionPeriodExtendedrf4 = 2
    RepetitionPeriodExtendedrf8 = 3
    RepetitionPeriodExtendedrf16 = 4

    /* Extensions */
)
func (self *RepetitionPeriodExtended) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 5, 1)
}
func (self *RepetitionPeriodExtended) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 5, 1)
}
type SCPTMInformation struct { // [{'type': 'MBMS-Cell-List', 'name': 'mbmsCellList'}, {'type': 'MBMS-E-RAB-QoS-Parameters', 'name': 'mbms-E-RAB-QoS-Parameters'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SC-PTM-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    MbmsCellList MBMSCellList
    MbmsERABQoSParameters MBMSERABQoSParameters
    IEExtensions *SCPTMInformationExtIEs
}

func (self * SCPTMInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.MbmsCellList.Unpack(stream)// p8
    self.MbmsERABQoSParameters.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SCPTMInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SC-PTM-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SCPTMInformationExtIEs, order_SCPTMInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SCPTMInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MbmsCellList.Pack(stream)
    self.MbmsERABQoSParameters.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SCPTMInformationExtIEs, order_SCPTMInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SFN struct {
  Value uint64
}
func (self *SFN) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(1024, 10, 0, 0)
}
func (self * SFN) Pack(st *Stream){
    st.formatf_Integer(self.Value, 1024, 10, 0, 0)
}
type SubcarrierSpacingMBMS struct {
  Value int
}
const (
    SubcarrierSpacingMBMSkhz_7dot5 = 0
    SubcarrierSpacingMBMSkhz_1dot25 = 1

    /* Extensions */
)
func (self *SubcarrierSpacingMBMS) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *SubcarrierSpacingMBMS) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
func (self *SubframeAllocationExtended)Unpack(stream *Stream) {
    //coptions := []string{"oneFrameExtension","fourFrameExtension","choice-extension","Unknown"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in SubframeAllocationExtended\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_oneFrameExtension = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(2, 2)
        }
        self.OneFrameExtension = &BITSTRING{}//cho5
        Unpack_oneFrameExtension(stream, self.OneFrameExtension);
    } else if choice == 1 { //ch2
        var Unpack_fourFrameExtension = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(8, 8)
        }
        self.FourFrameExtension = &BITSTRING{}//cho5
        Unpack_fourFrameExtension(stream, self.FourFrameExtension);
    } else if choice == 2 { //ch2
        Choiceextension := &ProtocolIESingleContainer{}//cho2
        self.Choiceextension = &SubframeAllocationExtendedExtIEs{}//cho3
        Choiceextension.Unpack(stream, self.Choiceextension)
    }//end of if else

    if choice >= 3 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * SubframeAllocationExtended) Pack(stream *Stream) {
    if self.OneFrameExtension != nil {
        stream.set_choice(0, 2, 1, 3)
        var Pack_oneFrameExtension = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 2)
        }
        Pack_oneFrameExtension(stream, *self.OneFrameExtension)//3
    } else if self.FourFrameExtension != nil {
        stream.set_choice(1, 2, 1, 3)
        var Pack_fourFrameExtension = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 8)
        }
        Pack_fourFrameExtension(stream, *self.FourFrameExtension)//3
    } else if self.Choiceextension != nil {
        stream.set_choice(2, 2, 1, 3)
        Choiceextension := ProtocolIESingleContainer{}//cho2
        Choiceextension.Pack(stream, self.Choiceextension)
    }

}
type SubframeAllocationExtended struct { //[{'type': 'BIT STRING', 'size': [2], 'name': 'oneFrameExtension'}, {'type': 'BIT STRING', 'size': [8], 'name': 'fourFrameExtension'}, {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['SubframeAllocationExtended-ExtIEs'], 'name': 'choice-extension'}, None]
    OneFrameExtension *BITSTRING
    FourFrameExtension *BITSTRING
    Choiceextension *SubframeAllocationExtendedExtIEs
} // SubframeAllocationExtended

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
type TMGI struct { // [{'type': 'PLMN-Identity', 'name': 'pLMNidentity'}, {'type': 'OCTET STRING', 'size': [3], 'name': 'serviceID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TMGI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNidentity PLMNIdentity
    ServiceID OCTETSTRING
    IEExtensions *TMGIExtIEs
}

func (self * TMGI) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PLMNidentity.Unpack(stream)// p8
    var Unpack_serviceID = func(st *Stream, self *OCTETSTRING) {
        self.Value = st.parsef_OctString(3)
    }
    Unpack_serviceID(stream, &self.ServiceID)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &TMGIExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TMGI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_TMGIExtIEs, order_TMGIExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TMGI) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNidentity.Pack(stream)
    var Pack_serviceID = func(st *Stream, self OCTETSTRING) {
        st.formatf_OctString(self.Value, 3)
    }
    Pack_serviceID(stream, self.ServiceID) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_TMGIExtIEs, order_TMGIExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type TNLInformation struct { // [{'type': 'IPAddress', 'name': 'iPMCAddress'}, {'type': 'IPAddress', 'name': 'iPSourceAddress'}, {'type': 'GTP-TEID', 'name': 'gTP-TEID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TNL-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    IPMCAddress IPAddress
    IPSourceAddress IPAddress
    GTPTEID GTPTEID
    IEExtensions *TNLInformationExtIEs
}

func (self * TNLInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.IPMCAddress.Unpack(stream)// p8
    self.IPSourceAddress.Unpack(stream)// p8
    self.GTPTEID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &TNLInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TNL-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_TNLInformationExtIEs, order_TNLInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TNLInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IPMCAddress.Pack(stream)
    self.IPSourceAddress.Pack(stream)
    self.GTPTEID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_TNLInformationExtIEs, order_TNLInformationExtIEs} // p3
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
    _size := data.(M2APPROTOCOLIES_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Field', 'actual-parameters': ['IesSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IesSetParam'], 'param-types': ['M2AP-PROTOCOL-IES']}
    Items map[int]*M2APPROTOCOLIES
    order []int
}

type ProtocolIESingleContainer struct{ //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IesSetParam'], 'parameters': ['IesSetParam'], 'param-types': ['M2AP-PROTOCOL-IES']}
   Item map[int]*M2APPROTOCOLIES //UserType
   order []int
}
func (self *ProtocolIESingleContainer) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IesSetParam'], 'parameters': ['IesSetParam'], 'param-types': ['M2AP-PROTOCOL-IES']}
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

type ProtocolIEField struct { // [{'type': 'M2AP-PROTOCOL-IES.&id', 'table': {'type': 'IesSetParam'}, 'name': 'id'}, {'type': 'M2AP-PROTOCOL-IES.&criticality', 'table': ['IesSetParam', ['id']], 'name': 'criticality'}, {'type': 'M2AP-PROTOCOL-IES.&Value', 'table': ['IesSetParam', ['id']], 'name': 'value'}]
    Id M2APPROTOCOLIESid
    Criticality M2APPROTOCOLIEScriticality
    Value M2APPROTOCOLIESValue
}

func (self * ProtocolIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'M2AP-PROTOCOL-IES.&id', 'table': {'type': 'IesSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := M2APPROTOCOLIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(M2APPROTOCOLIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'M2AP-PROTOCOL-IES.&Value', 'table': ['IesSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * M2APPROTOCOLIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'M2AP-PROTOCOL-IES.&id', 'table': {'type': 'IesSetParam'}, 'name': 'id'}
    key := (M2APPROTOCOLIESid)(self.ID)
    if out.(M2APPROTOCOLIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(M2APPROTOCOLIES_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

func (self *ProtocolIEContainerPair) Unpack(stream *Stream, out interface{}) { //SeqOF1
    _size := stream.get_listsize(65536) + 0
    for i := 0; i <_size; i +=1 {
        item := ProtocolIEFieldPair{}
        item.Unpack(stream, out)
    }
    return 
}


func (self * ProtocolIEContainerPair) Pack(stream *Stream, data interface{}) { //sqof 1
    _size := data.(M2APPROTOCOLIESPAIR_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainerPair struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-FieldPair', 'actual-parameters': ['IesSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IesSetParam'], 'param-types': ['M2AP-PROTOCOL-IES-PAIR']}
    Items map[int]*M2APPROTOCOLIESPAIR
    order []int
}

type ProtocolIEFieldPair struct { // [{'type': 'M2AP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IesSetParam'}, 'name': 'id'}, {'type': 'M2AP-PROTOCOL-IES-PAIR.&firstCriticality', 'table': ['IesSetParam', ['id']], 'name': 'firstCriticality'}, {'type': 'M2AP-PROTOCOL-IES-PAIR.&FirstValue', 'table': ['IesSetParam', ['id']], 'name': 'firstValue'}, {'type': 'M2AP-PROTOCOL-IES-PAIR.&secondCriticality', 'table': ['IesSetParam', ['id']], 'name': 'secondCriticality'}, {'type': 'M2AP-PROTOCOL-IES-PAIR.&SecondValue', 'table': ['IesSetParam', ['id']], 'name': 'secondValue'}]
    Id M2APPROTOCOLIESPAIRid
    FirstCriticality M2APPROTOCOLIESPAIRfirstCriticality
    FirstValue M2APPROTOCOLIESPAIRFirstValue
    SecondCriticality M2APPROTOCOLIESPAIRsecondCriticality
    SecondValue M2APPROTOCOLIESPAIRSecondValue
}

func (self * ProtocolIEFieldPair) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'M2AP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IesSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := M2APPROTOCOLIESPAIR{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.FirstCriticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(M2APPROTOCOLIESPAIR_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'M2AP-PROTOCOL-IES-PAIR.&FirstValue', 'table': ['IesSetParam', ['id']], 'name': 'firstValue'}
    self.SecondCriticality.Unpack(stream)//p9
    out.(M2APPROTOCOLIESPAIR_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'M2AP-PROTOCOL-IES-PAIR.&SecondValue', 'table': ['IesSetParam', ['id']], 'name': 'secondValue'}
    stream.set_location(location, _len)
    return
}

func (self * M2APPROTOCOLIESPAIR) Pack(stream *Stream, out interface{}) {
    //table {'type': 'M2AP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IesSetParam'}, 'name': 'id'}
    key := (M2APPROTOCOLIESPAIRid)(self.ID)
    if out.(M2APPROTOCOLIESPAIR_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.FIRSTCRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(M2APPROTOCOLIESPAIR_IF).PackOT(stream, key)
    self.SECONDCRITICALITY.Pack(stream)
    out.(M2APPROTOCOLIESPAIR_IF).PackOT(stream, key)
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


type ProtocolIEContainerList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['IesSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IesSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'M2AP-PROTOCOL-IES']}
    Items map[int]*M2APPROTOCOLIES
    order []int
}

func (self *ProtocolIEContainerPairList)Unpack(stream *Stream, arg1 uint64, arg2 uint64,  data interface{}) { //Seq3
    //cloc := stream.get_current_location();
    _size := stream.get_listsize(arg2-arg1+1) + int(arg1)
    val := ProtocolIEContainerPair{}
    for item := 0; item <_size; item +=1 {
        val.Unpack(stream, data);
        //values = append(values, val.(interface{}))
    }
    return
}


func (self *ProtocolIEContainerPairList) Pack (stream *Stream, arg1 int, arg2 int, data interface{}) { //seqof 4
    _size := len(self.Items)
    stream.set_listsize(_size-arg1, uint64(arg2-arg1+1))
    val := ProtocolIEContainerPair{}
    for _, item := range self.Items {
        val.Pack(stream, item);
    }
}


type ProtocolIEContainerPairList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-ContainerPair', 'actual-parameters': ['IesSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IesSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'M2AP-PROTOCOL-IES-PAIR']}
    Items map[int]*M2APPROTOCOLIESPAIR
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
    _size := data.(M2APPROTOCOLEXTENSION_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolExtensionContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolExtensionField', 'actual-parameters': ['ExtensionSetParam']}, 'size': [(1, 'maxProtocolExtensions')], 'parameters': ['ExtensionSetParam'], 'param-types': ['M2AP-PROTOCOL-EXTENSION']}
    Items map[int]*M2APPROTOCOLEXTENSION
    order []int
}

type ProtocolExtensionField struct { // [{'type': 'M2AP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}, {'type': 'M2AP-PROTOCOL-EXTENSION.&criticality', 'table': ['ExtensionSetParam', ['id']], 'name': 'criticality'}, {'type': 'M2AP-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}]
    Id M2APPROTOCOLEXTENSIONid
    Criticality M2APPROTOCOLEXTENSIONcriticality
    ExtensionValue M2APPROTOCOLEXTENSIONExtension
}

func (self * ProtocolExtensionField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'M2AP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := M2APPROTOCOLEXTENSION{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(M2APPROTOCOLEXTENSION_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'M2AP-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}
    stream.set_location(location, _len)
    return
}

func (self * M2APPROTOCOLEXTENSION) Pack(stream *Stream, out interface{}) {
    //table {'type': 'M2AP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    key := (M2APPROTOCOLEXTENSIONid)(self.ID)
    if out.(M2APPROTOCOLEXTENSION_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(M2APPROTOCOLEXTENSION_IF).PackOT(stream, key)
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
    _size := data.(M2APPRIVATEIES_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type PrivateIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PrivateIE-Field', 'actual-parameters': ['IesSetParam']}, 'size': [(1, 'maxPrivateIEs')], 'parameters': ['IesSetParam'], 'param-types': ['M2AP-PRIVATE-IES']}
    Items map[int]*M2APPRIVATEIES
    order []int
}

type PrivateIEField struct { // [{'type': 'M2AP-PRIVATE-IES.&id', 'table': {'type': 'IesSetParam'}, 'name': 'id'}, {'type': 'M2AP-PRIVATE-IES.&criticality', 'table': ['IesSetParam', ['id']], 'name': 'criticality'}, {'type': 'M2AP-PRIVATE-IES.&Value', 'table': ['IesSetParam', ['id']], 'name': 'value'}]
    Id M2APPRIVATEIESid
    Criticality M2APPRIVATEIEScriticality
    Value M2APPRIVATEIESValue
}

func (self * PrivateIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'M2AP-PRIVATE-IES.&id', 'table': {'type': 'IesSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := M2APPRIVATEIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(M2APPRIVATEIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'M2AP-PRIVATE-IES.&Value', 'table': ['IesSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * M2APPRIVATEIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'M2AP-PRIVATE-IES.&id', 'table': {'type': 'IesSetParam'}, 'name': 'id'}
    key := (M2APPRIVATEIESid)(self.ID)
    if out.(M2APPRIVATEIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(M2APPRIVATEIES_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type M2APELEMENTARYPROCEDURE struct {//CLASS {'members': [{'type': 'OpenType', 'name': '&InitiatingMessage'}, {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, {'type': 'ProcedureCode', 'name': '&procedureCode'}, {'type': 'Criticality', 'name': '&criticality'}], 'with-members': [{'with-order': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME', 'PROCEDURE CODE', 'CRITICALITY'], 'with-type': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME']}], 'alt-type': {'INITIATING MESSAGE': {'type': 'OpenType', 'name': '&InitiatingMessage'}, 'SUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, 'UNSUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, 'PROCEDURE CODE': {'type': 'ProcedureCode', 'name': '&procedureCode'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}}, 'id-type': {'type': 'ProcedureCode', 'name': '&procedureCode'}}
    INITIATINGMESSAGE interface{}
    SUCCESSFULOUTCOME interface{}
    UNSUCCESSFULOUTCOME interface{}
    PROCEDURECODE ProcedureCode
    CRITICALITY Criticality
}
type M2APELEMENTARYPROCEDUREInitiatingMessage interface{}
type M2APELEMENTARYPROCEDURESuccessfulOutcome interface{}
type M2APELEMENTARYPROCEDUREUnsuccessfulOutcome interface{}
type M2APELEMENTARYPROCEDUREprocedureCode ProcedureCode
func (self *M2APELEMENTARYPROCEDUREprocedureCode) Pack(st *Stream) {
    ieID := ProcedureCode(*self)
    ieID.Pack(st)
}
func (self *M2APELEMENTARYPROCEDUREprocedureCode) Unpack(st *Stream) {
    val := ProcedureCode{}
    val.Unpack(st)
    *self = M2APELEMENTARYPROCEDUREprocedureCode(val)
}
type M2APELEMENTARYPROCEDUREcriticality Criticality
func (self *M2APELEMENTARYPROCEDUREcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M2APELEMENTARYPROCEDUREcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M2APELEMENTARYPROCEDUREcriticality(val)
}

type M2APELEMENTARYPROCEDURE_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type M2APPROTOCOLIES struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type M2APPROTOCOLIESid ProtocolIEID
func (self *M2APPROTOCOLIESid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *M2APPROTOCOLIESid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = M2APPROTOCOLIESid(val)
}
type M2APPROTOCOLIEScriticality Criticality
func (self *M2APPROTOCOLIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M2APPROTOCOLIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M2APPROTOCOLIEScriticality(val)
}
type M2APPROTOCOLIESValue interface{}
type M2APPROTOCOLIESpresence Presence
func (self *M2APPROTOCOLIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *M2APPROTOCOLIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = M2APPROTOCOLIESpresence(val)
}

type M2APPROTOCOLIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type M2APPROTOCOLIESPAIR struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&firstCriticality'}, {'type': 'OpenType', 'name': '&FirstValue'}, {'type': 'Criticality', 'name': '&secondCriticality'}, {'type': 'OpenType', 'name': '&SecondValue'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'FIRST CRITICALITY', 'FIRST TYPE', 'SECOND CRITICALITY', 'SECOND TYPE', 'PRESENCE'], 'with-type': ['FIRST TYPE', 'SECOND TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'FIRST CRITICALITY': {'type': 'Criticality', 'name': '&firstCriticality'}, 'FIRST TYPE': {'type': 'OpenType', 'name': '&FirstValue'}, 'SECOND CRITICALITY': {'type': 'Criticality', 'name': '&secondCriticality'}, 'SECOND TYPE': {'type': 'OpenType', 'name': '&SecondValue'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    FIRSTCRITICALITY Criticality
    FIRSTTYPE interface{}
    SECONDCRITICALITY Criticality
    SECONDTYPE interface{}
    PRESENCE Presence
}
type M2APPROTOCOLIESPAIRid ProtocolIEID
func (self *M2APPROTOCOLIESPAIRid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *M2APPROTOCOLIESPAIRid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = M2APPROTOCOLIESPAIRid(val)
}
type M2APPROTOCOLIESPAIRfirstCriticality Criticality
func (self *M2APPROTOCOLIESPAIRfirstCriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M2APPROTOCOLIESPAIRfirstCriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M2APPROTOCOLIESPAIRfirstCriticality(val)
}
type M2APPROTOCOLIESPAIRFirstValue interface{}
type M2APPROTOCOLIESPAIRsecondCriticality Criticality
func (self *M2APPROTOCOLIESPAIRsecondCriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M2APPROTOCOLIESPAIRsecondCriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M2APPROTOCOLIESPAIRsecondCriticality(val)
}
type M2APPROTOCOLIESPAIRSecondValue interface{}
type M2APPROTOCOLIESPAIRpresence Presence
func (self *M2APPROTOCOLIESPAIRpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *M2APPROTOCOLIESPAIRpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = M2APPROTOCOLIESPAIRpresence(val)
}

type M2APPROTOCOLIESPAIR_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type M2APPROTOCOLEXTENSION struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Extension'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'EXTENSION', 'PRESENCE'], 'with-type': ['EXTENSION']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    EXTENSION interface{}
    PRESENCE Presence
}
type M2APPROTOCOLEXTENSIONid ProtocolIEID
func (self *M2APPROTOCOLEXTENSIONid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *M2APPROTOCOLEXTENSIONid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = M2APPROTOCOLEXTENSIONid(val)
}
type M2APPROTOCOLEXTENSIONcriticality Criticality
func (self *M2APPROTOCOLEXTENSIONcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M2APPROTOCOLEXTENSIONcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M2APPROTOCOLEXTENSIONcriticality(val)
}
type M2APPROTOCOLEXTENSIONExtension interface{}
type M2APPROTOCOLEXTENSIONpresence Presence
func (self *M2APPROTOCOLEXTENSIONpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *M2APPROTOCOLEXTENSIONpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = M2APPROTOCOLEXTENSIONpresence(val)
}

type M2APPROTOCOLEXTENSION_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type M2APPRIVATEIES struct {//CLASS {'members': [{'type': 'PrivateIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'PrivateIE-ID', 'name': '&id'}}
    ID PrivateIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type M2APPRIVATEIESid PrivateIEID
func (self *M2APPRIVATEIESid) Pack(st *Stream) {
    ieID := PrivateIEID(*self)
    ieID.Pack(st)
}
func (self *M2APPRIVATEIESid) Unpack(st *Stream) {
    val := PrivateIEID{}
    val.Unpack(st)
    *self = M2APPRIVATEIESid(val)
}
type M2APPRIVATEIEScriticality Criticality
func (self *M2APPRIVATEIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M2APPRIVATEIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M2APPRIVATEIEScriticality(val)
}
type M2APPRIVATEIESValue interface{}
type M2APPRIVATEIESpresence Presence
func (self *M2APPRIVATEIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *M2APPRIVATEIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = M2APPRIVATEIESpresence(val)
}

type M2APPRIVATEIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
//class M2APELEMENTARYPROCEDURES: #OBJSET1 {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'members': [{}, {}, None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_M2APELEMENTARYPROCEDURES = make(map[M2APELEMENTARYPROCEDUREprocedureCode]*M2APELEMENTARYPROCEDURE)


//class M2APELEMENTARYPROCEDURESCLASS1: #OBJSET1 {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_M2APELEMENTARYPROCEDURESCLASS1 = make(map[M2APELEMENTARYPROCEDUREprocedureCode]*M2APELEMENTARYPROCEDURE)


//class M2APELEMENTARYPROCEDURESCLASS2: #OBJSET1 {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_M2APELEMENTARYPROCEDURESCLASS2 = make(map[M2APELEMENTARYPROCEDUREprocedureCode]*M2APELEMENTARYPROCEDURE)


type SessionStartRequestIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCE-MBMS-M2AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'MCE-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-TMGI', 'CRITICALITY': 'reject', 'TYPE': 'TMGI', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Session-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Session-ID', 'PRESENCE': 'optional'}, {'ID': 'id-MBMS-Service-Area', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Service-Area', 'PRESENCE': 'mandatory'}, {'ID': 'id-TNL-Information', 'CRITICALITY': 'reject', 'TYPE': 'TNL-Information', 'PRESENCE': 'mandatory'}, {'ID': 'id-Alternative-TNL-Information', 'CRITICALITY': 'ignore', 'TYPE': 'TNL-Information', 'PRESENCE': 'optional'}, {'ID': 'id-SC-PTM-Information', 'CRITICALITY': 'reject', 'TYPE': 'SC-PTM-Information', 'PRESENCE': 'optional'}, None]}
   MCEMBMSM2APID  MCEMBMSM2APID
   TMGI  TMGI
   MBMSSessionID  *MBMSSessionID
   MBMSServiceArea  MBMSServiceArea
   TNLInformation  TNLInformation
   AlternativeTNLInformation  *TNLInformation
   SCPTMInformation  *SCPTMInformation
   list []interface{}
}
func (self *SessionStartRequestIes)createOT() interface{}{
    return nil
}
var table_SessionStartRequestIes = make(map[int]*M2APPROTOCOLIES)

var order_SessionStartRequestIes = make([]int, 7)

func (self *SessionStartRequestIes) GetIECount() int{
   count := 0
   count +=1 //self.MCEMBMSM2APID
   count +=1 //self.TMGI
   if self.MBMSSessionID != nil { count += 1 }
   count +=1 //self.MBMSServiceArea
   count +=1 //self.TNLInformation
   if self.AlternativeTNLInformation != nil { count += 1 }
   if self.SCPTMInformation != nil { count += 1 }
   return count//ObjSet
}
func (self *SessionStartRequestIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        return true //self.MCEMBMSM2APID
      case 2: //TMGI
        return true //self.TMGI
      case 3: //MBMSSessionID
        if self.MBMSSessionID != nil { return true }
      case 6: //MBMSServiceArea
        return true //self.MBMSServiceArea
      case 7: //TNLInformation
        return true //self.TNLInformation
      case 38: //AlternativeTNLInformation
        if self.AlternativeTNLInformation != nil { return true }
      case 45: //SCPTMInformation
        if self.SCPTMInformation != nil { return true }
   }
   return false//ObjSet
}
func (self *SessionStartRequestIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM2APID)
      case 2: //TMGI
        self.TMGI.Unpack(st)
        self.list = append(self.list, &self.TMGI)
      case 3: //MBMSSessionID
        self.MBMSSessionID = &MBMSSessionID{}
        self.MBMSSessionID.Unpack(st)
        self.list = append(self.list, self.MBMSSessionID)
      case 6: //MBMSServiceArea
        self.MBMSServiceArea.Unpack(st)
        self.list = append(self.list, &self.MBMSServiceArea)
      case 7: //TNLInformation
        self.TNLInformation.Unpack(st)
        self.list = append(self.list, &self.TNLInformation)
      case 38: //AlternativeTNLInformation
        self.AlternativeTNLInformation = &TNLInformation{}
        self.AlternativeTNLInformation.Unpack(st)
        self.list = append(self.list, self.AlternativeTNLInformation)
      case 45: //SCPTMInformation
        self.SCPTMInformation = &SCPTMInformation{}
        self.SCPTMInformation.Unpack(st)
        self.list = append(self.list, self.SCPTMInformation)
   }
}
func (self *SessionStartRequestIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Pack(st)
      case 2: //TMGI
        self.TMGI.Pack(st)
      case 3: //MBMSSessionID
        if self.MBMSSessionID != nil {self.MBMSSessionID.Pack(st)}
      case 6: //MBMSServiceArea
        self.MBMSServiceArea.Pack(st)
      case 7: //TNLInformation
        self.TNLInformation.Pack(st)
      case 38: //AlternativeTNLInformation
        if self.AlternativeTNLInformation != nil {self.AlternativeTNLInformation.Pack(st)}
      case 45: //SCPTMInformation
        if self.SCPTMInformation != nil {self.SCPTMInformation.Pack(st)}
      default:
      break
   }
}
func init() {
table_SessionStartRequestIes[0] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM2APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCEMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStartRequestIes[0] = 0
table_SessionStartRequestIes[2] = &M2APPROTOCOLIES{ID:ProtocolIEID{idTMGI}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TMGI{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStartRequestIes[1] = 2
table_SessionStartRequestIes[3] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSSessionID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSSessionID{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionStartRequestIes[2] = 3
table_SessionStartRequestIes[6] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceArea}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSServiceArea{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStartRequestIes[3] = 6
table_SessionStartRequestIes[7] = &M2APPROTOCOLIES{ID:ProtocolIEID{idTNLInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TNLInformation{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStartRequestIes[4] = 7
table_SessionStartRequestIes[38] = &M2APPROTOCOLIES{ID:ProtocolIEID{idAlternativeTNLInformation}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TNLInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionStartRequestIes[5] = 38
table_SessionStartRequestIes[45] = &M2APPROTOCOLIES{ID:ProtocolIEID{idSCPTMInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SCPTMInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionStartRequestIes[6] = 45
   }

type SessionStartResponseIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCE-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-ENB-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'ENB-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MCEMBMSM2APID  MCEMBMSM2APID
   ENBMBMSM2APID  ENBMBMSM2APID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *SessionStartResponseIes)createOT() interface{}{
    return nil
}
var table_SessionStartResponseIes = make(map[int]*M2APPROTOCOLIES)

var order_SessionStartResponseIes = make([]int, 3)

func (self *SessionStartResponseIes) GetIECount() int{
   count := 0
   count +=1 //self.MCEMBMSM2APID
   count +=1 //self.ENBMBMSM2APID
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *SessionStartResponseIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        return true //self.MCEMBMSM2APID
      case 1: //ENBMBMSM2APID
        return true //self.ENBMBMSM2APID
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *SessionStartResponseIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM2APID)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.ENBMBMSM2APID)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *SessionStartResponseIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Pack(st)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_SessionStartResponseIes[0] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStartResponseIes[0] = 0
table_SessionStartResponseIes[1] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ENBMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStartResponseIes[1] = 1
table_SessionStartResponseIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionStartResponseIes[2] = 8
   }

type SessionStartFailureIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCE-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MCEMBMSM2APID  MCEMBMSM2APID
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *SessionStartFailureIes)createOT() interface{}{
    return nil
}
var table_SessionStartFailureIes = make(map[int]*M2APPROTOCOLIES)

var order_SessionStartFailureIes = make([]int, 3)

func (self *SessionStartFailureIes) GetIECount() int{
   count := 0
   count +=1 //self.MCEMBMSM2APID
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *SessionStartFailureIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        return true //self.MCEMBMSM2APID
      case 9: //Cause
        return true //self.Cause
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *SessionStartFailureIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM2APID)
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *SessionStartFailureIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Pack(st)
      case 9: //Cause
        self.Cause.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_SessionStartFailureIes[0] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStartFailureIes[0] = 0
table_SessionStartFailureIes[9] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStartFailureIes[1] = 9
table_SessionStartFailureIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionStartFailureIes[2] = 8
   }

type SessionStopRequestIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCE-MBMS-M2AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'MCE-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-ENB-MBMS-M2AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'ENB-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, None]}
   MCEMBMSM2APID  MCEMBMSM2APID
   ENBMBMSM2APID  ENBMBMSM2APID
   list []interface{}
}
func (self *SessionStopRequestIes)createOT() interface{}{
    return nil
}
var table_SessionStopRequestIes = make(map[int]*M2APPROTOCOLIES)

var order_SessionStopRequestIes = make([]int, 2)

func (self *SessionStopRequestIes) GetIECount() int{
   count := 0
   count +=1 //self.MCEMBMSM2APID
   count +=1 //self.ENBMBMSM2APID
   return count//ObjSet
}
func (self *SessionStopRequestIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        return true //self.MCEMBMSM2APID
      case 1: //ENBMBMSM2APID
        return true //self.ENBMBMSM2APID
   }
   return false//ObjSet
}
func (self *SessionStopRequestIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM2APID)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.ENBMBMSM2APID)
   }
}
func (self *SessionStopRequestIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Pack(st)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Pack(st)
      default:
      break
   }
}
func init() {
table_SessionStopRequestIes[0] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM2APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCEMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStopRequestIes[0] = 0
table_SessionStopRequestIes[1] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSM2APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ENBMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStopRequestIes[1] = 1
   }

type SessionStopResponseIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCE-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-ENB-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'ENB-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MCEMBMSM2APID  MCEMBMSM2APID
   ENBMBMSM2APID  ENBMBMSM2APID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *SessionStopResponseIes)createOT() interface{}{
    return nil
}
var table_SessionStopResponseIes = make(map[int]*M2APPROTOCOLIES)

var order_SessionStopResponseIes = make([]int, 3)

func (self *SessionStopResponseIes) GetIECount() int{
   count := 0
   count +=1 //self.MCEMBMSM2APID
   count +=1 //self.ENBMBMSM2APID
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *SessionStopResponseIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        return true //self.MCEMBMSM2APID
      case 1: //ENBMBMSM2APID
        return true //self.ENBMBMSM2APID
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *SessionStopResponseIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM2APID)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.ENBMBMSM2APID)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *SessionStopResponseIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Pack(st)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_SessionStopResponseIes[0] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStopResponseIes[0] = 0
table_SessionStopResponseIes[1] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ENBMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionStopResponseIes[1] = 1
table_SessionStopResponseIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionStopResponseIes[2] = 8
   }

type SessionUpdateRequestIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCE-MBMS-M2AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'MCE-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-ENB-MBMS-M2AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'ENB-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-TMGI', 'CRITICALITY': 'reject', 'TYPE': 'TMGI', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Session-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Session-ID', 'PRESENCE': 'optional'}, {'ID': 'id-MBMS-Service-Area', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Service-Area', 'PRESENCE': 'optional'}, {'ID': 'id-TNL-Information', 'CRITICALITY': 'reject', 'TYPE': 'TNL-Information', 'PRESENCE': 'optional'}, {'ID': 'id-SC-PTM-Information', 'CRITICALITY': 'reject', 'TYPE': 'SC-PTM-Information', 'PRESENCE': 'optional'}, None]}
   MCEMBMSM2APID  MCEMBMSM2APID
   ENBMBMSM2APID  ENBMBMSM2APID
   TMGI  TMGI
   MBMSSessionID  *MBMSSessionID
   MBMSServiceArea  *MBMSServiceArea
   TNLInformation  *TNLInformation
   SCPTMInformation  *SCPTMInformation
   list []interface{}
}
func (self *SessionUpdateRequestIes)createOT() interface{}{
    return nil
}
var table_SessionUpdateRequestIes = make(map[int]*M2APPROTOCOLIES)

var order_SessionUpdateRequestIes = make([]int, 7)

func (self *SessionUpdateRequestIes) GetIECount() int{
   count := 0
   count +=1 //self.MCEMBMSM2APID
   count +=1 //self.ENBMBMSM2APID
   count +=1 //self.TMGI
   if self.MBMSSessionID != nil { count += 1 }
   if self.MBMSServiceArea != nil { count += 1 }
   if self.TNLInformation != nil { count += 1 }
   if self.SCPTMInformation != nil { count += 1 }
   return count//ObjSet
}
func (self *SessionUpdateRequestIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        return true //self.MCEMBMSM2APID
      case 1: //ENBMBMSM2APID
        return true //self.ENBMBMSM2APID
      case 2: //TMGI
        return true //self.TMGI
      case 3: //MBMSSessionID
        if self.MBMSSessionID != nil { return true }
      case 6: //MBMSServiceArea
        if self.MBMSServiceArea != nil { return true }
      case 7: //TNLInformation
        if self.TNLInformation != nil { return true }
      case 45: //SCPTMInformation
        if self.SCPTMInformation != nil { return true }
   }
   return false//ObjSet
}
func (self *SessionUpdateRequestIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM2APID)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.ENBMBMSM2APID)
      case 2: //TMGI
        self.TMGI.Unpack(st)
        self.list = append(self.list, &self.TMGI)
      case 3: //MBMSSessionID
        self.MBMSSessionID = &MBMSSessionID{}
        self.MBMSSessionID.Unpack(st)
        self.list = append(self.list, self.MBMSSessionID)
      case 6: //MBMSServiceArea
        self.MBMSServiceArea = &MBMSServiceArea{}
        self.MBMSServiceArea.Unpack(st)
        self.list = append(self.list, self.MBMSServiceArea)
      case 7: //TNLInformation
        self.TNLInformation = &TNLInformation{}
        self.TNLInformation.Unpack(st)
        self.list = append(self.list, self.TNLInformation)
      case 45: //SCPTMInformation
        self.SCPTMInformation = &SCPTMInformation{}
        self.SCPTMInformation.Unpack(st)
        self.list = append(self.list, self.SCPTMInformation)
   }
}
func (self *SessionUpdateRequestIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Pack(st)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Pack(st)
      case 2: //TMGI
        self.TMGI.Pack(st)
      case 3: //MBMSSessionID
        if self.MBMSSessionID != nil {self.MBMSSessionID.Pack(st)}
      case 6: //MBMSServiceArea
        if self.MBMSServiceArea != nil {self.MBMSServiceArea.Pack(st)}
      case 7: //TNLInformation
        if self.TNLInformation != nil {self.TNLInformation.Pack(st)}
      case 45: //SCPTMInformation
        if self.SCPTMInformation != nil {self.SCPTMInformation.Pack(st)}
      default:
      break
   }
}
func init() {
table_SessionUpdateRequestIes[0] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM2APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCEMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionUpdateRequestIes[0] = 0
table_SessionUpdateRequestIes[1] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSM2APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ENBMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionUpdateRequestIes[1] = 1
table_SessionUpdateRequestIes[2] = &M2APPROTOCOLIES{ID:ProtocolIEID{idTMGI}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TMGI{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionUpdateRequestIes[2] = 2
table_SessionUpdateRequestIes[3] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSSessionID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSSessionID{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionUpdateRequestIes[3] = 3
table_SessionUpdateRequestIes[6] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceArea}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSServiceArea{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionUpdateRequestIes[4] = 6
table_SessionUpdateRequestIes[7] = &M2APPROTOCOLIES{ID:ProtocolIEID{idTNLInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TNLInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionUpdateRequestIes[5] = 7
table_SessionUpdateRequestIes[45] = &M2APPROTOCOLIES{ID:ProtocolIEID{idSCPTMInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SCPTMInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionUpdateRequestIes[6] = 45
   }

type SessionUpdateResponseIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCE-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-ENB-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'ENB-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MCEMBMSM2APID  MCEMBMSM2APID
   ENBMBMSM2APID  ENBMBMSM2APID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *SessionUpdateResponseIes)createOT() interface{}{
    return nil
}
var table_SessionUpdateResponseIes = make(map[int]*M2APPROTOCOLIES)

var order_SessionUpdateResponseIes = make([]int, 3)

func (self *SessionUpdateResponseIes) GetIECount() int{
   count := 0
   count +=1 //self.MCEMBMSM2APID
   count +=1 //self.ENBMBMSM2APID
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *SessionUpdateResponseIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        return true //self.MCEMBMSM2APID
      case 1: //ENBMBMSM2APID
        return true //self.ENBMBMSM2APID
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *SessionUpdateResponseIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM2APID)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.ENBMBMSM2APID)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *SessionUpdateResponseIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Pack(st)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_SessionUpdateResponseIes[0] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionUpdateResponseIes[0] = 0
table_SessionUpdateResponseIes[1] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ENBMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionUpdateResponseIes[1] = 1
table_SessionUpdateResponseIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionUpdateResponseIes[2] = 8
   }

type SessionUpdateFailureIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCE-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-ENB-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'ENB-MBMS-M2AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MCEMBMSM2APID  MCEMBMSM2APID
   ENBMBMSM2APID  ENBMBMSM2APID
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *SessionUpdateFailureIes)createOT() interface{}{
    return nil
}
var table_SessionUpdateFailureIes = make(map[int]*M2APPROTOCOLIES)

var order_SessionUpdateFailureIes = make([]int, 4)

func (self *SessionUpdateFailureIes) GetIECount() int{
   count := 0
   count +=1 //self.MCEMBMSM2APID
   count +=1 //self.ENBMBMSM2APID
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *SessionUpdateFailureIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        return true //self.MCEMBMSM2APID
      case 1: //ENBMBMSM2APID
        return true //self.ENBMBMSM2APID
      case 9: //Cause
        return true //self.Cause
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *SessionUpdateFailureIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM2APID)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Unpack(st)
        self.list = append(self.list, &self.ENBMBMSM2APID)
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *SessionUpdateFailureIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID.Pack(st)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID.Pack(st)
      case 9: //Cause
        self.Cause.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_SessionUpdateFailureIes[0] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionUpdateFailureIes[0] = 0
table_SessionUpdateFailureIes[1] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ENBMBMSM2APID{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionUpdateFailureIes[1] = 1
table_SessionUpdateFailureIes[9] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_SessionUpdateFailureIes[2] = 9
table_SessionUpdateFailureIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_SessionUpdateFailureIes[3] = 8
   }

type MbmsSchedulingInformationIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCCH-Update-Time', 'CRITICALITY': 'reject', 'TYPE': 'MCCH-Update-Time', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBSFN-Area-Configuration-List', 'CRITICALITY': 'reject', 'TYPE': 'MBSFN-Area-Configuration-List', 'PRESENCE': 'mandatory'}, None]}
   MCCHUpdateTime  MCCHUpdateTime
   MBSFNAreaConfigurationList  MBSFNAreaConfigurationList
   list []interface{}
}
func (self *MbmsSchedulingInformationIes)createOT() interface{}{
    return nil
}
var table_MbmsSchedulingInformationIes = make(map[int]*M2APPROTOCOLIES)

var order_MbmsSchedulingInformationIes = make([]int, 2)

func (self *MbmsSchedulingInformationIes) GetIECount() int{
   count := 0
   count +=1 //self.MCCHUpdateTime
   count +=1 //self.MBSFNAreaConfigurationList
   return count//ObjSet
}
func (self *MbmsSchedulingInformationIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 25: //MCCHUpdateTime
        return true //self.MCCHUpdateTime
      case 10: //MBSFNAreaConfigurationList
        return true //self.MBSFNAreaConfigurationList
   }
   return false//ObjSet
}
func (self *MbmsSchedulingInformationIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 25: //MCCHUpdateTime
        self.MCCHUpdateTime.Unpack(st)
        self.list = append(self.list, &self.MCCHUpdateTime)
      case 10: //MBSFNAreaConfigurationList
        self.MBSFNAreaConfigurationList.Unpack(st)
        self.list = append(self.list, &self.MBSFNAreaConfigurationList)
   }
}
func (self *MbmsSchedulingInformationIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 25: //MCCHUpdateTime
        self.MCCHUpdateTime.Pack(st)
      case 10: //MBSFNAreaConfigurationList
        self.MBSFNAreaConfigurationList.Pack(st)
      default:
      break
   }
}
func init() {
table_MbmsSchedulingInformationIes[25] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCCHUpdateTime}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCCHUpdateTime{}, PRESENCE:Presence{Presencemandatory}, }
order_MbmsSchedulingInformationIes[0] = 25
table_MbmsSchedulingInformationIes[10] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBSFNAreaConfigurationList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBSFNAreaConfigurationList{}, PRESENCE:Presence{Presencemandatory}, }
order_MbmsSchedulingInformationIes[1] = 10
   }

type MBSFNAreaConfigurationItem struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-PMCH-Configuration-List', 'CRITICALITY': 'reject', 'TYPE': 'PMCH-Configuration-List', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBSFN-Subframe-Configuration-List', 'CRITICALITY': 'reject', 'TYPE': 'MBSFN-Subframe-ConfigurationList', 'PRESENCE': 'mandatory'}, {'ID': 'id-Common-Subframe-Allocation-Period', 'CRITICALITY': 'reject', 'TYPE': 'Common-Subframe-Allocation-Period', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBSFN-Area-ID', 'CRITICALITY': 'reject', 'TYPE': 'MBSFN-Area-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Suspension-Notification-List', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Suspension-Notification-List', 'PRESENCE': 'optional'}, None]}
   PMCHConfigurationList  PMCHConfigurationList
   MBSFNSubframeConfigurationList  MBSFNSubframeConfigurationList
   CommonSubframeAllocationPeriod  CommonSubframeAllocationPeriod
   MBSFNAreaID  MBSFNAreaID
   MBMSSuspensionNotificationList  *MBMSSuspensionNotificationList
   list []interface{}
}
func (self *MBSFNAreaConfigurationItem)createOT() interface{}{
    return nil
}
var table_MBSFNAreaConfigurationItem = make(map[int]*M2APPROTOCOLIES)

var order_MBSFNAreaConfigurationItem = make([]int, 5)

func (self *MBSFNAreaConfigurationItem) GetIECount() int{
   count := 0
   count +=1 //self.PMCHConfigurationList
   count +=1 //self.MBSFNSubframeConfigurationList
   count +=1 //self.CommonSubframeAllocationPeriod
   count +=1 //self.MBSFNAreaID
   if self.MBMSSuspensionNotificationList != nil { count += 1 }
   return count//ObjSet
}
func (self *MBSFNAreaConfigurationItem) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 11: //PMCHConfigurationList
        return true //self.PMCHConfigurationList
      case 22: //MBSFNSubframeConfigurationList
        return true //self.MBSFNSubframeConfigurationList
      case 24: //CommonSubframeAllocationPeriod
        return true //self.CommonSubframeAllocationPeriod
      case 29: //MBSFNAreaID
        return true //self.MBSFNAreaID
      case 43: //MBMSSuspensionNotificationList
        if self.MBMSSuspensionNotificationList != nil { return true }
   }
   return false//ObjSet
}
func (self *MBSFNAreaConfigurationItem)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 11: //PMCHConfigurationList
        self.PMCHConfigurationList.Unpack(st)
        self.list = append(self.list, &self.PMCHConfigurationList)
      case 22: //MBSFNSubframeConfigurationList
        self.MBSFNSubframeConfigurationList.Unpack(st)
        self.list = append(self.list, &self.MBSFNSubframeConfigurationList)
      case 24: //CommonSubframeAllocationPeriod
        self.CommonSubframeAllocationPeriod.Unpack(st)
        self.list = append(self.list, &self.CommonSubframeAllocationPeriod)
      case 29: //MBSFNAreaID
        self.MBSFNAreaID.Unpack(st)
        self.list = append(self.list, &self.MBSFNAreaID)
      case 43: //MBMSSuspensionNotificationList
        self.MBMSSuspensionNotificationList = &MBMSSuspensionNotificationList{}
        self.MBMSSuspensionNotificationList.Unpack(st)
        self.list = append(self.list, self.MBMSSuspensionNotificationList)
   }
}
func (self *MBSFNAreaConfigurationItem)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 11: //PMCHConfigurationList
        self.PMCHConfigurationList.Pack(st)
      case 22: //MBSFNSubframeConfigurationList
        self.MBSFNSubframeConfigurationList.Pack(st)
      case 24: //CommonSubframeAllocationPeriod
        self.CommonSubframeAllocationPeriod.Pack(st)
      case 29: //MBSFNAreaID
        self.MBSFNAreaID.Pack(st)
      case 43: //MBMSSuspensionNotificationList
        if self.MBMSSuspensionNotificationList != nil {self.MBMSSuspensionNotificationList.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBSFNAreaConfigurationItem[11] = &M2APPROTOCOLIES{ID:ProtocolIEID{idPMCHConfigurationList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PMCHConfigurationList{}, PRESENCE:Presence{Presencemandatory}, }
order_MBSFNAreaConfigurationItem[0] = 11
table_MBSFNAreaConfigurationItem[22] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBSFNSubframeConfigurationList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBSFNSubframeConfigurationList{}, PRESENCE:Presence{Presencemandatory}, }
order_MBSFNAreaConfigurationItem[1] = 22
table_MBSFNAreaConfigurationItem[24] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCommonSubframeAllocationPeriod}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CommonSubframeAllocationPeriod{}, PRESENCE:Presence{Presencemandatory}, }
order_MBSFNAreaConfigurationItem[2] = 24
table_MBSFNAreaConfigurationItem[29] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBSFNAreaID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBSFNAreaID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBSFNAreaConfigurationItem[3] = 29
table_MBSFNAreaConfigurationItem[43] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSSuspensionNotificationList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSSuspensionNotificationList{}, PRESENCE:Presence{Presenceoptional}, }
order_MBSFNAreaConfigurationItem[4] = 43
   }

type PMCHConfigurationItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-PMCH-Configuration-Item', 'CRITICALITY': 'reject', 'TYPE': 'PMCH-Configuration-Item', 'PRESENCE': 'mandatory'}, None]}
   PMCHConfigurationItem  PMCHConfigurationItem
   list []interface{}
}
func (self *PMCHConfigurationItemIEs)createOT() interface{}{
    return nil
}
var table_PMCHConfigurationItemIEs = make(map[int]*M2APPROTOCOLIES)

var order_PMCHConfigurationItemIEs = make([]int, 1)

func (self *PMCHConfigurationItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.PMCHConfigurationItem
   return count//ObjSet
}
func (self *PMCHConfigurationItemIEs) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 12: //PMCHConfigurationItem
        return true //self.PMCHConfigurationItem
   }
   return false//ObjSet
}
func (self *PMCHConfigurationItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 12: //PMCHConfigurationItem
        self.PMCHConfigurationItem.Unpack(st)
        self.list = append(self.list, &self.PMCHConfigurationItem)
   }
}
func (self *PMCHConfigurationItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 12: //PMCHConfigurationItem
        self.PMCHConfigurationItem.Pack(st)
      default:
      break
   }
}
func init() {
table_PMCHConfigurationItemIEs[12] = &M2APPROTOCOLIES{ID:ProtocolIEID{idPMCHConfigurationItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PMCHConfigurationItem{}, PRESENCE:Presence{Presencemandatory}, }
order_PMCHConfigurationItemIEs[0] = 12
   }

type PMCHConfigurationItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PMCHConfigurationItemExtIEs)createOT() interface{}{
    return nil
}
var table_PMCHConfigurationItemExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_PMCHConfigurationItemExtIEs = make([]int, 0)

type MBSFNSubframeConfigurationItem struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBSFN-Subframe-Configuration-Item', 'CRITICALITY': 'reject', 'TYPE': 'MBSFN-Subframe-Configuration', 'PRESENCE': 'mandatory'}, None]}
   MBSFNSubframeConfigurationItem  MBSFNSubframeConfiguration
   list []interface{}
}
func (self *MBSFNSubframeConfigurationItem)createOT() interface{}{
    return nil
}
var table_MBSFNSubframeConfigurationItem = make(map[int]*M2APPROTOCOLIES)

var order_MBSFNSubframeConfigurationItem = make([]int, 1)

func (self *MBSFNSubframeConfigurationItem) GetIECount() int{
   count := 0
   count +=1 //self.MBSFNSubframeConfigurationItem
   return count//ObjSet
}
func (self *MBSFNSubframeConfigurationItem) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 23: //MBSFNSubframeConfigurationItem
        return true //self.MBSFNSubframeConfigurationItem
   }
   return false//ObjSet
}
func (self *MBSFNSubframeConfigurationItem)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 23: //MBSFNSubframeConfigurationItem
        self.MBSFNSubframeConfigurationItem.Unpack(st)
        self.list = append(self.list, &self.MBSFNSubframeConfigurationItem)
   }
}
func (self *MBSFNSubframeConfigurationItem)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 23: //MBSFNSubframeConfigurationItem
        self.MBSFNSubframeConfigurationItem.Pack(st)
      default:
      break
   }
}
func init() {
table_MBSFNSubframeConfigurationItem[23] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBSFNSubframeConfigurationItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBSFNSubframeConfiguration{}, PRESENCE:Presence{Presencemandatory}, }
order_MBSFNSubframeConfigurationItem[0] = 23
   }

type MBMSSuspensionNotificationItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBMS-Suspension-Notification-Item', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Suspension-Notification-Item', 'PRESENCE': 'optional'}, None]}
   MBMSSuspensionNotificationItem  *MBMSSuspensionNotificationItem
   list []interface{}
}
func (self *MBMSSuspensionNotificationItemIEs)createOT() interface{}{
    return nil
}
var table_MBMSSuspensionNotificationItemIEs = make(map[int]*M2APPROTOCOLIES)

var order_MBMSSuspensionNotificationItemIEs = make([]int, 1)

func (self *MBMSSuspensionNotificationItemIEs) GetIECount() int{
   count := 0
   if self.MBMSSuspensionNotificationItem != nil { count += 1 }
   return count//ObjSet
}
func (self *MBMSSuspensionNotificationItemIEs) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 44: //MBMSSuspensionNotificationItem
        if self.MBMSSuspensionNotificationItem != nil { return true }
   }
   return false//ObjSet
}
func (self *MBMSSuspensionNotificationItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 44: //MBMSSuspensionNotificationItem
        self.MBMSSuspensionNotificationItem = &MBMSSuspensionNotificationItem{}
        self.MBMSSuspensionNotificationItem.Unpack(st)
        self.list = append(self.list, self.MBMSSuspensionNotificationItem)
   }
}
func (self *MBMSSuspensionNotificationItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 44: //MBMSSuspensionNotificationItem
        if self.MBMSSuspensionNotificationItem != nil {self.MBMSSuspensionNotificationItem.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBMSSuspensionNotificationItemIEs[44] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSSuspensionNotificationItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSSuspensionNotificationItem{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSuspensionNotificationItemIEs[0] = 44
   }

type MBMSSuspensionNotificationItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MBMSSuspensionNotificationItemExtIEs)createOT() interface{}{
    return nil
}
var table_MBMSSuspensionNotificationItemExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_MBMSSuspensionNotificationItemExtIEs = make([]int, 0)

type MbmsSchedulingInformationResponseIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MbmsSchedulingInformationResponseIes)createOT() interface{}{
    return nil
}
var table_MbmsSchedulingInformationResponseIes = make(map[int]*M2APPROTOCOLIES)

var order_MbmsSchedulingInformationResponseIes = make([]int, 1)

func (self *MbmsSchedulingInformationResponseIes) GetIECount() int{
   count := 0
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MbmsSchedulingInformationResponseIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MbmsSchedulingInformationResponseIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MbmsSchedulingInformationResponseIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MbmsSchedulingInformationResponseIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MbmsSchedulingInformationResponseIes[0] = 8
   }

type M2SetupRequestIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-GlobalENB-ID', 'CRITICALITY': 'reject', 'TYPE': 'GlobalENB-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-ENBname', 'CRITICALITY': 'ignore', 'TYPE': 'ENBname', 'PRESENCE': 'optional'}, {'ID': 'id-ENB-MBMS-Configuration-data-List', 'CRITICALITY': 'reject', 'TYPE': 'ENB-MBMS-Configuration-data-List', 'PRESENCE': 'mandatory'}, None]}
   GlobalENBID  GlobalENBID
   ENBname  *ENBname
   ENBMBMSConfigurationdataList  ENBMBMSConfigurationdataList
   list []interface{}
}
func (self *M2SetupRequestIes)createOT() interface{}{
    return nil
}
var table_M2SetupRequestIes = make(map[int]*M2APPROTOCOLIES)

var order_M2SetupRequestIes = make([]int, 3)

func (self *M2SetupRequestIes) GetIECount() int{
   count := 0
   count +=1 //self.GlobalENBID
   if self.ENBname != nil { count += 1 }
   count +=1 //self.ENBMBMSConfigurationdataList
   return count//ObjSet
}
func (self *M2SetupRequestIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 13: //GlobalENBID
        return true //self.GlobalENBID
      case 14: //ENBname
        if self.ENBname != nil { return true }
      case 15: //ENBMBMSConfigurationdataList
        return true //self.ENBMBMSConfigurationdataList
   }
   return false//ObjSet
}
func (self *M2SetupRequestIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 13: //GlobalENBID
        self.GlobalENBID.Unpack(st)
        self.list = append(self.list, &self.GlobalENBID)
      case 14: //ENBname
        self.ENBname = &ENBname{}
        self.ENBname.Unpack(st)
        self.list = append(self.list, self.ENBname)
      case 15: //ENBMBMSConfigurationdataList
        self.ENBMBMSConfigurationdataList.Unpack(st)
        self.list = append(self.list, &self.ENBMBMSConfigurationdataList)
   }
}
func (self *M2SetupRequestIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 13: //GlobalENBID
        self.GlobalENBID.Pack(st)
      case 14: //ENBname
        if self.ENBname != nil {self.ENBname.Pack(st)}
      case 15: //ENBMBMSConfigurationdataList
        self.ENBMBMSConfigurationdataList.Pack(st)
      default:
      break
   }
}
func init() {
table_M2SetupRequestIes[13] = &M2APPROTOCOLIES{ID:ProtocolIEID{idGlobalENBID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GlobalENBID{}, PRESENCE:Presence{Presencemandatory}, }
order_M2SetupRequestIes[0] = 13
table_M2SetupRequestIes[14] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBname}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ENBname{}, PRESENCE:Presence{Presenceoptional}, }
order_M2SetupRequestIes[1] = 14
table_M2SetupRequestIes[15] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSConfigurationdataList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ENBMBMSConfigurationdataList{}, PRESENCE:Presence{Presencemandatory}, }
order_M2SetupRequestIes[2] = 15
   }

type ENBMBMSConfigurationdataItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-ENB-MBMS-Configuration-data-Item', 'CRITICALITY': 'reject', 'TYPE': 'ENB-MBMS-Configuration-data-Item', 'PRESENCE': 'mandatory'}, None]}
   ENBMBMSConfigurationdataItem  ENBMBMSConfigurationdataItem
   list []interface{}
}
func (self *ENBMBMSConfigurationdataItemIEs)createOT() interface{}{
    return nil
}
var table_ENBMBMSConfigurationdataItemIEs = make(map[int]*M2APPROTOCOLIES)

var order_ENBMBMSConfigurationdataItemIEs = make([]int, 1)

func (self *ENBMBMSConfigurationdataItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.ENBMBMSConfigurationdataItem
   return count//ObjSet
}
func (self *ENBMBMSConfigurationdataItemIEs) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 16: //ENBMBMSConfigurationdataItem
        return true //self.ENBMBMSConfigurationdataItem
   }
   return false//ObjSet
}
func (self *ENBMBMSConfigurationdataItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 16: //ENBMBMSConfigurationdataItem
        self.ENBMBMSConfigurationdataItem.Unpack(st)
        self.list = append(self.list, &self.ENBMBMSConfigurationdataItem)
   }
}
func (self *ENBMBMSConfigurationdataItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 16: //ENBMBMSConfigurationdataItem
        self.ENBMBMSConfigurationdataItem.Pack(st)
      default:
      break
   }
}
func init() {
table_ENBMBMSConfigurationdataItemIEs[16] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSConfigurationdataItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ENBMBMSConfigurationdataItem{}, PRESENCE:Presence{Presencemandatory}, }
order_ENBMBMSConfigurationdataItemIEs[0] = 16
   }

type M2SetupResponseIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-GlobalMCE-ID', 'CRITICALITY': 'reject', 'TYPE': 'GlobalMCE-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MCEname', 'CRITICALITY': 'ignore', 'TYPE': 'MCEname', 'PRESENCE': 'optional'}, {'ID': 'id-MCCHrelatedBCCH-ConfigPerMBSFNArea', 'CRITICALITY': 'reject', 'TYPE': 'MCCHrelatedBCCH-ConfigPerMBSFNArea', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   GlobalMCEID  GlobalMCEID
   MCEname  *MCEname
   MCCHrelatedBCCHConfigPerMBSFNArea  MCCHrelatedBCCHConfigPerMBSFNArea
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *M2SetupResponseIes)createOT() interface{}{
    return nil
}
var table_M2SetupResponseIes = make(map[int]*M2APPROTOCOLIES)

var order_M2SetupResponseIes = make([]int, 4)

func (self *M2SetupResponseIes) GetIECount() int{
   count := 0
   count +=1 //self.GlobalMCEID
   if self.MCEname != nil { count += 1 }
   count +=1 //self.MCCHrelatedBCCHConfigPerMBSFNArea
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *M2SetupResponseIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 17: //GlobalMCEID
        return true //self.GlobalMCEID
      case 18: //MCEname
        if self.MCEname != nil { return true }
      case 19: //MCCHrelatedBCCHConfigPerMBSFNArea
        return true //self.MCCHrelatedBCCHConfigPerMBSFNArea
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *M2SetupResponseIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 17: //GlobalMCEID
        self.GlobalMCEID.Unpack(st)
        self.list = append(self.list, &self.GlobalMCEID)
      case 18: //MCEname
        self.MCEname = &MCEname{}
        self.MCEname.Unpack(st)
        self.list = append(self.list, self.MCEname)
      case 19: //MCCHrelatedBCCHConfigPerMBSFNArea
        self.MCCHrelatedBCCHConfigPerMBSFNArea.Unpack(st)
        self.list = append(self.list, &self.MCCHrelatedBCCHConfigPerMBSFNArea)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *M2SetupResponseIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 17: //GlobalMCEID
        self.GlobalMCEID.Pack(st)
      case 18: //MCEname
        if self.MCEname != nil {self.MCEname.Pack(st)}
      case 19: //MCCHrelatedBCCHConfigPerMBSFNArea
        self.MCCHrelatedBCCHConfigPerMBSFNArea.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_M2SetupResponseIes[17] = &M2APPROTOCOLIES{ID:ProtocolIEID{idGlobalMCEID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GlobalMCEID{}, PRESENCE:Presence{Presencemandatory}, }
order_M2SetupResponseIes[0] = 17
table_M2SetupResponseIes[18] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEname}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEname{}, PRESENCE:Presence{Presenceoptional}, }
order_M2SetupResponseIes[1] = 18
table_M2SetupResponseIes[19] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCCHrelatedBCCHConfigPerMBSFNArea}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCCHrelatedBCCHConfigPerMBSFNArea{}, PRESENCE:Presence{Presencemandatory}, }
order_M2SetupResponseIes[2] = 19
table_M2SetupResponseIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_M2SetupResponseIes[3] = 8
   }

type MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCCHrelatedBCCH-ConfigPerMBSFNArea-Item', 'CRITICALITY': 'reject', 'TYPE': 'MCCHrelatedBCCH-ConfigPerMBSFNArea-Item', 'PRESENCE': 'mandatory'}, None]}
   MCCHrelatedBCCHConfigPerMBSFNAreaItem  MCCHrelatedBCCHConfigPerMBSFNAreaItem
   list []interface{}
}
func (self *MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs)createOT() interface{}{
    return nil
}
var table_MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs = make(map[int]*M2APPROTOCOLIES)

var order_MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs = make([]int, 1)

func (self *MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.MCCHrelatedBCCHConfigPerMBSFNAreaItem
   return count//ObjSet
}
func (self *MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 20: //MCCHrelatedBCCHConfigPerMBSFNAreaItem
        return true //self.MCCHrelatedBCCHConfigPerMBSFNAreaItem
   }
   return false//ObjSet
}
func (self *MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 20: //MCCHrelatedBCCHConfigPerMBSFNAreaItem
        self.MCCHrelatedBCCHConfigPerMBSFNAreaItem.Unpack(st)
        self.list = append(self.list, &self.MCCHrelatedBCCHConfigPerMBSFNAreaItem)
   }
}
func (self *MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 20: //MCCHrelatedBCCHConfigPerMBSFNAreaItem
        self.MCCHrelatedBCCHConfigPerMBSFNAreaItem.Pack(st)
      default:
      break
   }
}
func init() {
table_MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs[20] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCCHrelatedBCCHConfigPerMBSFNAreaItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCCHrelatedBCCHConfigPerMBSFNAreaItem{}, PRESENCE:Presence{Presencemandatory}, }
order_MCCHrelatedBCCHConfigPerMBSFNAreaItemIEs[0] = 20
   }

type M2SetupFailureIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *M2SetupFailureIes)createOT() interface{}{
    return nil
}
var table_M2SetupFailureIes = make(map[int]*M2APPROTOCOLIES)

var order_M2SetupFailureIes = make([]int, 3)

func (self *M2SetupFailureIes) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *M2SetupFailureIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        return true //self.Cause
      case 21: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *M2SetupFailureIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 21: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *M2SetupFailureIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Pack(st)
      case 21: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_M2SetupFailureIes[9] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_M2SetupFailureIes[0] = 9
table_M2SetupFailureIes[21] = &M2APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_M2SetupFailureIes[1] = 21
table_M2SetupFailureIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_M2SetupFailureIes[2] = 8
   }

type ENBConfigurationUpdateIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-GlobalENB-ID', 'CRITICALITY': 'reject', 'TYPE': 'GlobalENB-ID', 'PRESENCE': 'optional'}, {'ID': 'id-ENBname', 'CRITICALITY': 'ignore', 'TYPE': 'ENBname', 'PRESENCE': 'optional'}, {'ID': 'id-ENB-MBMS-Configuration-data-List-ConfigUpdate', 'CRITICALITY': 'reject', 'TYPE': 'ENB-MBMS-Configuration-data-List-ConfigUpdate', 'PRESENCE': 'optional'}, None]}
   GlobalENBID  *GlobalENBID
   ENBname  *ENBname
   ENBMBMSConfigurationdataListConfigUpdate  *ENBMBMSConfigurationdataListConfigUpdate
   list []interface{}
}
func (self *ENBConfigurationUpdateIes)createOT() interface{}{
    return nil
}
var table_ENBConfigurationUpdateIes = make(map[int]*M2APPROTOCOLIES)

var order_ENBConfigurationUpdateIes = make([]int, 3)

func (self *ENBConfigurationUpdateIes) GetIECount() int{
   count := 0
   if self.GlobalENBID != nil { count += 1 }
   if self.ENBname != nil { count += 1 }
   if self.ENBMBMSConfigurationdataListConfigUpdate != nil { count += 1 }
   return count//ObjSet
}
func (self *ENBConfigurationUpdateIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 13: //GlobalENBID
        if self.GlobalENBID != nil { return true }
      case 14: //ENBname
        if self.ENBname != nil { return true }
      case 26: //ENBMBMSConfigurationdataListConfigUpdate
        if self.ENBMBMSConfigurationdataListConfigUpdate != nil { return true }
   }
   return false//ObjSet
}
func (self *ENBConfigurationUpdateIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 13: //GlobalENBID
        self.GlobalENBID = &GlobalENBID{}
        self.GlobalENBID.Unpack(st)
        self.list = append(self.list, self.GlobalENBID)
      case 14: //ENBname
        self.ENBname = &ENBname{}
        self.ENBname.Unpack(st)
        self.list = append(self.list, self.ENBname)
      case 26: //ENBMBMSConfigurationdataListConfigUpdate
        self.ENBMBMSConfigurationdataListConfigUpdate = &ENBMBMSConfigurationdataListConfigUpdate{}
        self.ENBMBMSConfigurationdataListConfigUpdate.Unpack(st)
        self.list = append(self.list, self.ENBMBMSConfigurationdataListConfigUpdate)
   }
}
func (self *ENBConfigurationUpdateIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 13: //GlobalENBID
        if self.GlobalENBID != nil {self.GlobalENBID.Pack(st)}
      case 14: //ENBname
        if self.ENBname != nil {self.ENBname.Pack(st)}
      case 26: //ENBMBMSConfigurationdataListConfigUpdate
        if self.ENBMBMSConfigurationdataListConfigUpdate != nil {self.ENBMBMSConfigurationdataListConfigUpdate.Pack(st)}
      default:
      break
   }
}
func init() {
table_ENBConfigurationUpdateIes[13] = &M2APPROTOCOLIES{ID:ProtocolIEID{idGlobalENBID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GlobalENBID{}, PRESENCE:Presence{Presenceoptional}, }
order_ENBConfigurationUpdateIes[0] = 13
table_ENBConfigurationUpdateIes[14] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBname}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ENBname{}, PRESENCE:Presence{Presenceoptional}, }
order_ENBConfigurationUpdateIes[1] = 14
table_ENBConfigurationUpdateIes[26] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSConfigurationdataListConfigUpdate}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ENBMBMSConfigurationdataListConfigUpdate{}, PRESENCE:Presence{Presenceoptional}, }
order_ENBConfigurationUpdateIes[2] = 26
   }

type ENBMBMSConfigurationdataConfigUpdateItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-ENB-MBMS-Configuration-data-ConfigUpdate-Item', 'CRITICALITY': 'reject', 'TYPE': 'ENB-MBMS-Configuration-data-ConfigUpdate-Item', 'PRESENCE': 'mandatory'}, None]}
   ENBMBMSConfigurationdataConfigUpdateItem  ENBMBMSConfigurationdataConfigUpdateItem
   list []interface{}
}
func (self *ENBMBMSConfigurationdataConfigUpdateItemIEs)createOT() interface{}{
    return nil
}
var table_ENBMBMSConfigurationdataConfigUpdateItemIEs = make(map[int]*M2APPROTOCOLIES)

var order_ENBMBMSConfigurationdataConfigUpdateItemIEs = make([]int, 1)

func (self *ENBMBMSConfigurationdataConfigUpdateItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.ENBMBMSConfigurationdataConfigUpdateItem
   return count//ObjSet
}
func (self *ENBMBMSConfigurationdataConfigUpdateItemIEs) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 27: //ENBMBMSConfigurationdataConfigUpdateItem
        return true //self.ENBMBMSConfigurationdataConfigUpdateItem
   }
   return false//ObjSet
}
func (self *ENBMBMSConfigurationdataConfigUpdateItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 27: //ENBMBMSConfigurationdataConfigUpdateItem
        self.ENBMBMSConfigurationdataConfigUpdateItem.Unpack(st)
        self.list = append(self.list, &self.ENBMBMSConfigurationdataConfigUpdateItem)
   }
}
func (self *ENBMBMSConfigurationdataConfigUpdateItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 27: //ENBMBMSConfigurationdataConfigUpdateItem
        self.ENBMBMSConfigurationdataConfigUpdateItem.Pack(st)
      default:
      break
   }
}
func init() {
table_ENBMBMSConfigurationdataConfigUpdateItemIEs[27] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSConfigurationdataConfigUpdateItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ENBMBMSConfigurationdataConfigUpdateItem{}, PRESENCE:Presence{Presencemandatory}, }
order_ENBMBMSConfigurationdataConfigUpdateItemIEs[0] = 27
   }

type ENBConfigurationUpdateAcknowledgeIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCCHrelatedBCCH-ConfigPerMBSFNArea', 'CRITICALITY': 'reject', 'TYPE': 'MCCHrelatedBCCH-ConfigPerMBSFNArea', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MCCHrelatedBCCHConfigPerMBSFNArea  *MCCHrelatedBCCHConfigPerMBSFNArea
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ENBConfigurationUpdateAcknowledgeIes)createOT() interface{}{
    return nil
}
var table_ENBConfigurationUpdateAcknowledgeIes = make(map[int]*M2APPROTOCOLIES)

var order_ENBConfigurationUpdateAcknowledgeIes = make([]int, 2)

func (self *ENBConfigurationUpdateAcknowledgeIes) GetIECount() int{
   count := 0
   if self.MCCHrelatedBCCHConfigPerMBSFNArea != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ENBConfigurationUpdateAcknowledgeIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 19: //MCCHrelatedBCCHConfigPerMBSFNArea
        if self.MCCHrelatedBCCHConfigPerMBSFNArea != nil { return true }
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ENBConfigurationUpdateAcknowledgeIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 19: //MCCHrelatedBCCHConfigPerMBSFNArea
        self.MCCHrelatedBCCHConfigPerMBSFNArea = &MCCHrelatedBCCHConfigPerMBSFNArea{}
        self.MCCHrelatedBCCHConfigPerMBSFNArea.Unpack(st)
        self.list = append(self.list, self.MCCHrelatedBCCHConfigPerMBSFNArea)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ENBConfigurationUpdateAcknowledgeIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 19: //MCCHrelatedBCCHConfigPerMBSFNArea
        if self.MCCHrelatedBCCHConfigPerMBSFNArea != nil {self.MCCHrelatedBCCHConfigPerMBSFNArea.Pack(st)}
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ENBConfigurationUpdateAcknowledgeIes[19] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCCHrelatedBCCHConfigPerMBSFNArea}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCCHrelatedBCCHConfigPerMBSFNArea{}, PRESENCE:Presence{Presenceoptional}, }
order_ENBConfigurationUpdateAcknowledgeIes[0] = 19
table_ENBConfigurationUpdateAcknowledgeIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ENBConfigurationUpdateAcknowledgeIes[1] = 8
   }

type ENBConfigurationUpdateFailureIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ENBConfigurationUpdateFailureIes)createOT() interface{}{
    return nil
}
var table_ENBConfigurationUpdateFailureIes = make(map[int]*M2APPROTOCOLIES)

var order_ENBConfigurationUpdateFailureIes = make([]int, 3)

func (self *ENBConfigurationUpdateFailureIes) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ENBConfigurationUpdateFailureIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        return true //self.Cause
      case 21: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ENBConfigurationUpdateFailureIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 21: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ENBConfigurationUpdateFailureIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Pack(st)
      case 21: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ENBConfigurationUpdateFailureIes[9] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_ENBConfigurationUpdateFailureIes[0] = 9
table_ENBConfigurationUpdateFailureIes[21] = &M2APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_ENBConfigurationUpdateFailureIes[1] = 21
table_ENBConfigurationUpdateFailureIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ENBConfigurationUpdateFailureIes[2] = 8
   }

type MCEConfigurationUpdateIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-GlobalMCE-ID', 'CRITICALITY': 'reject', 'TYPE': 'GlobalMCE-ID', 'PRESENCE': 'optional'}, {'ID': 'id-MCEname', 'CRITICALITY': 'ignore', 'TYPE': 'MCEname', 'PRESENCE': 'optional'}, {'ID': 'id-MCCHrelatedBCCH-ConfigPerMBSFNArea', 'CRITICALITY': 'reject', 'TYPE': 'MCCHrelatedBCCH-ConfigPerMBSFNArea', 'PRESENCE': 'optional'}, None]}
   GlobalMCEID  *GlobalMCEID
   MCEname  *MCEname
   MCCHrelatedBCCHConfigPerMBSFNArea  *MCCHrelatedBCCHConfigPerMBSFNArea
   list []interface{}
}
func (self *MCEConfigurationUpdateIes)createOT() interface{}{
    return nil
}
var table_MCEConfigurationUpdateIes = make(map[int]*M2APPROTOCOLIES)

var order_MCEConfigurationUpdateIes = make([]int, 3)

func (self *MCEConfigurationUpdateIes) GetIECount() int{
   count := 0
   if self.GlobalMCEID != nil { count += 1 }
   if self.MCEname != nil { count += 1 }
   if self.MCCHrelatedBCCHConfigPerMBSFNArea != nil { count += 1 }
   return count//ObjSet
}
func (self *MCEConfigurationUpdateIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 17: //GlobalMCEID
        if self.GlobalMCEID != nil { return true }
      case 18: //MCEname
        if self.MCEname != nil { return true }
      case 19: //MCCHrelatedBCCHConfigPerMBSFNArea
        if self.MCCHrelatedBCCHConfigPerMBSFNArea != nil { return true }
   }
   return false//ObjSet
}
func (self *MCEConfigurationUpdateIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 17: //GlobalMCEID
        self.GlobalMCEID = &GlobalMCEID{}
        self.GlobalMCEID.Unpack(st)
        self.list = append(self.list, self.GlobalMCEID)
      case 18: //MCEname
        self.MCEname = &MCEname{}
        self.MCEname.Unpack(st)
        self.list = append(self.list, self.MCEname)
      case 19: //MCCHrelatedBCCHConfigPerMBSFNArea
        self.MCCHrelatedBCCHConfigPerMBSFNArea = &MCCHrelatedBCCHConfigPerMBSFNArea{}
        self.MCCHrelatedBCCHConfigPerMBSFNArea.Unpack(st)
        self.list = append(self.list, self.MCCHrelatedBCCHConfigPerMBSFNArea)
   }
}
func (self *MCEConfigurationUpdateIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 17: //GlobalMCEID
        if self.GlobalMCEID != nil {self.GlobalMCEID.Pack(st)}
      case 18: //MCEname
        if self.MCEname != nil {self.MCEname.Pack(st)}
      case 19: //MCCHrelatedBCCHConfigPerMBSFNArea
        if self.MCCHrelatedBCCHConfigPerMBSFNArea != nil {self.MCCHrelatedBCCHConfigPerMBSFNArea.Pack(st)}
      default:
      break
   }
}
func init() {
table_MCEConfigurationUpdateIes[17] = &M2APPROTOCOLIES{ID:ProtocolIEID{idGlobalMCEID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GlobalMCEID{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateIes[0] = 17
table_MCEConfigurationUpdateIes[18] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEname}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEname{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateIes[1] = 18
table_MCEConfigurationUpdateIes[19] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCCHrelatedBCCHConfigPerMBSFNArea}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCCHrelatedBCCHConfigPerMBSFNArea{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateIes[2] = 19
   }

type MCEConfigurationUpdateAcknowledgeIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MCEConfigurationUpdateAcknowledgeIes)createOT() interface{}{
    return nil
}
var table_MCEConfigurationUpdateAcknowledgeIes = make(map[int]*M2APPROTOCOLIES)

var order_MCEConfigurationUpdateAcknowledgeIes = make([]int, 1)

func (self *MCEConfigurationUpdateAcknowledgeIes) GetIECount() int{
   count := 0
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MCEConfigurationUpdateAcknowledgeIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MCEConfigurationUpdateAcknowledgeIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MCEConfigurationUpdateAcknowledgeIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MCEConfigurationUpdateAcknowledgeIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateAcknowledgeIes[0] = 8
   }

type MCEConfigurationUpdateFailureIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MCEConfigurationUpdateFailureIes)createOT() interface{}{
    return nil
}
var table_MCEConfigurationUpdateFailureIes = make(map[int]*M2APPROTOCOLIES)

var order_MCEConfigurationUpdateFailureIes = make([]int, 3)

func (self *MCEConfigurationUpdateFailureIes) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MCEConfigurationUpdateFailureIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        return true //self.Cause
      case 21: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MCEConfigurationUpdateFailureIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 21: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MCEConfigurationUpdateFailureIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Pack(st)
      case 21: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MCEConfigurationUpdateFailureIes[9] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_MCEConfigurationUpdateFailureIes[0] = 9
table_MCEConfigurationUpdateFailureIes[21] = &M2APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateFailureIes[1] = 21
table_MCEConfigurationUpdateFailureIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateFailureIes[2] = 8
   }

type ErrorIndicationIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCE-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M2AP-ID', 'PRESENCE': 'optional'}, {'ID': 'id-ENB-MBMS-M2AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'ENB-MBMS-M2AP-ID', 'PRESENCE': 'optional'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MCEMBMSM2APID  *MCEMBMSM2APID
   ENBMBMSM2APID  *ENBMBMSM2APID
   Cause  *Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ErrorIndicationIes)createOT() interface{}{
    return nil
}
var table_ErrorIndicationIes = make(map[int]*M2APPROTOCOLIES)

var order_ErrorIndicationIes = make([]int, 4)

func (self *ErrorIndicationIes) GetIECount() int{
   count := 0
   if self.MCEMBMSM2APID != nil { count += 1 }
   if self.ENBMBMSM2APID != nil { count += 1 }
   if self.Cause != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ErrorIndicationIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        if self.MCEMBMSM2APID != nil { return true }
      case 1: //ENBMBMSM2APID
        if self.ENBMBMSM2APID != nil { return true }
      case 9: //Cause
        if self.Cause != nil { return true }
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ErrorIndicationIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        self.MCEMBMSM2APID = &MCEMBMSM2APID{}
        self.MCEMBMSM2APID.Unpack(st)
        self.list = append(self.list, self.MCEMBMSM2APID)
      case 1: //ENBMBMSM2APID
        self.ENBMBMSM2APID = &ENBMBMSM2APID{}
        self.ENBMBMSM2APID.Unpack(st)
        self.list = append(self.list, self.ENBMBMSM2APID)
      case 9: //Cause
        self.Cause = &Cause{}
        self.Cause.Unpack(st)
        self.list = append(self.list, self.Cause)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ErrorIndicationIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 0: //MCEMBMSM2APID
        if self.MCEMBMSM2APID != nil {self.MCEMBMSM2APID.Pack(st)}
      case 1: //ENBMBMSM2APID
        if self.ENBMBMSM2APID != nil {self.ENBMBMSM2APID.Pack(st)}
      case 9: //Cause
        if self.Cause != nil {self.Cause.Pack(st)}
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ErrorIndicationIes[0] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM2APID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIes[0] = 0
table_ErrorIndicationIes[1] = &M2APPROTOCOLIES{ID:ProtocolIEID{idENBMBMSM2APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ENBMBMSM2APID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIes[1] = 1
table_ErrorIndicationIes[9] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIes[2] = 9
table_ErrorIndicationIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIes[3] = 8
   }

type ResetIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-ResetType', 'CRITICALITY': 'reject', 'TYPE': 'ResetType', 'PRESENCE': 'mandatory'}, None]}
   Cause  Cause
   ResetType  ResetType
   list []interface{}
}
func (self *ResetIes)createOT() interface{}{
    return nil
}
var table_ResetIes = make(map[int]*M2APPROTOCOLIES)

var order_ResetIes = make([]int, 2)

func (self *ResetIes) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   count +=1 //self.ResetType
   return count//ObjSet
}
func (self *ResetIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        return true //self.Cause
      case 30: //ResetType
        return true //self.ResetType
   }
   return false//ObjSet
}
func (self *ResetIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 30: //ResetType
        self.ResetType.Unpack(st)
        self.list = append(self.list, &self.ResetType)
   }
}
func (self *ResetIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Pack(st)
      case 30: //ResetType
        self.ResetType.Pack(st)
      default:
      break
   }
}
func init() {
table_ResetIes[9] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetIes[0] = 9
table_ResetIes[30] = &M2APPROTOCOLIES{ID:ProtocolIEID{idResetType}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ResetType{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetIes[1] = 30
   }

type MBMSServiceassociatedLogicalM2ConnectionItemRes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBMS-Service-associatedLogicalM2-ConnectionItem', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Service-associatedLogicalM2-ConnectionItem', 'PRESENCE': 'mandatory'}, None]}
   MBMSServiceassociatedLogicalM2ConnectionItem  MBMSServiceassociatedLogicalM2ConnectionItem
   list []interface{}
}
func (self *MBMSServiceassociatedLogicalM2ConnectionItemRes)createOT() interface{}{
    return nil
}
var table_MBMSServiceassociatedLogicalM2ConnectionItemRes = make(map[int]*M2APPROTOCOLIES)

var order_MBMSServiceassociatedLogicalM2ConnectionItemRes = make([]int, 1)

func (self *MBMSServiceassociatedLogicalM2ConnectionItemRes) GetIECount() int{
   count := 0
   count +=1 //self.MBMSServiceassociatedLogicalM2ConnectionItem
   return count//ObjSet
}
func (self *MBMSServiceassociatedLogicalM2ConnectionItemRes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 28: //MBMSServiceassociatedLogicalM2ConnectionItem
        return true //self.MBMSServiceassociatedLogicalM2ConnectionItem
   }
   return false//ObjSet
}
func (self *MBMSServiceassociatedLogicalM2ConnectionItemRes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 28: //MBMSServiceassociatedLogicalM2ConnectionItem
        self.MBMSServiceassociatedLogicalM2ConnectionItem.Unpack(st)
        self.list = append(self.list, &self.MBMSServiceassociatedLogicalM2ConnectionItem)
   }
}
func (self *MBMSServiceassociatedLogicalM2ConnectionItemRes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 28: //MBMSServiceassociatedLogicalM2ConnectionItem
        self.MBMSServiceassociatedLogicalM2ConnectionItem.Pack(st)
      default:
      break
   }
}
func init() {
table_MBMSServiceassociatedLogicalM2ConnectionItemRes[28] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceassociatedLogicalM2ConnectionItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSServiceassociatedLogicalM2ConnectionItem{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSServiceassociatedLogicalM2ConnectionItemRes[0] = 28
   }

type ResetAcknowledgeIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBMS-Service-associatedLogicalM2-ConnectionListResAck', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Service-associatedLogicalM2-ConnectionListResAck', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MBMSServiceassociatedLogicalM2ConnectionListResAck  *MBMSServiceassociatedLogicalM2ConnectionListResAck
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ResetAcknowledgeIes)createOT() interface{}{
    return nil
}
var table_ResetAcknowledgeIes = make(map[int]*M2APPROTOCOLIES)

var order_ResetAcknowledgeIes = make([]int, 2)

func (self *ResetAcknowledgeIes) GetIECount() int{
   count := 0
   if self.MBMSServiceassociatedLogicalM2ConnectionListResAck != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetAcknowledgeIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 31: //MBMSServiceassociatedLogicalM2ConnectionListResAck
        if self.MBMSServiceassociatedLogicalM2ConnectionListResAck != nil { return true }
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetAcknowledgeIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 31: //MBMSServiceassociatedLogicalM2ConnectionListResAck
        self.MBMSServiceassociatedLogicalM2ConnectionListResAck = &MBMSServiceassociatedLogicalM2ConnectionListResAck{}
        self.MBMSServiceassociatedLogicalM2ConnectionListResAck.Unpack(st)
        self.list = append(self.list, self.MBMSServiceassociatedLogicalM2ConnectionListResAck)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ResetAcknowledgeIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 31: //MBMSServiceassociatedLogicalM2ConnectionListResAck
        if self.MBMSServiceassociatedLogicalM2ConnectionListResAck != nil {self.MBMSServiceassociatedLogicalM2ConnectionListResAck.Pack(st)}
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetAcknowledgeIes[31] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceassociatedLogicalM2ConnectionListResAck}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSServiceassociatedLogicalM2ConnectionListResAck{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetAcknowledgeIes[0] = 31
table_ResetAcknowledgeIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetAcknowledgeIes[1] = 8
   }

type MBMSServiceassociatedLogicalM2ConnectionItemResAck struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBMS-Service-associatedLogicalM2-ConnectionItem', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Service-associatedLogicalM2-ConnectionItem', 'PRESENCE': 'mandatory'}, None]}
   MBMSServiceassociatedLogicalM2ConnectionItem  MBMSServiceassociatedLogicalM2ConnectionItem
   list []interface{}
}
func (self *MBMSServiceassociatedLogicalM2ConnectionItemResAck)createOT() interface{}{
    return nil
}
var table_MBMSServiceassociatedLogicalM2ConnectionItemResAck = make(map[int]*M2APPROTOCOLIES)

var order_MBMSServiceassociatedLogicalM2ConnectionItemResAck = make([]int, 1)

func (self *MBMSServiceassociatedLogicalM2ConnectionItemResAck) GetIECount() int{
   count := 0
   count +=1 //self.MBMSServiceassociatedLogicalM2ConnectionItem
   return count//ObjSet
}
func (self *MBMSServiceassociatedLogicalM2ConnectionItemResAck) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 28: //MBMSServiceassociatedLogicalM2ConnectionItem
        return true //self.MBMSServiceassociatedLogicalM2ConnectionItem
   }
   return false//ObjSet
}
func (self *MBMSServiceassociatedLogicalM2ConnectionItemResAck)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 28: //MBMSServiceassociatedLogicalM2ConnectionItem
        self.MBMSServiceassociatedLogicalM2ConnectionItem.Unpack(st)
        self.list = append(self.list, &self.MBMSServiceassociatedLogicalM2ConnectionItem)
   }
}
func (self *MBMSServiceassociatedLogicalM2ConnectionItemResAck)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 28: //MBMSServiceassociatedLogicalM2ConnectionItem
        self.MBMSServiceassociatedLogicalM2ConnectionItem.Pack(st)
      default:
      break
   }
}
func init() {
table_MBMSServiceassociatedLogicalM2ConnectionItemResAck[28] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceassociatedLogicalM2ConnectionItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSServiceassociatedLogicalM2ConnectionItem{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSServiceassociatedLogicalM2ConnectionItemResAck[0] = 28
   }

type PrivateMessageIes struct { //ObjSet 1 {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PRIVATE-IES', 'members': [None]}
   list []interface{}
}
func (self *PrivateMessageIes)createOT() interface{}{
    return nil
}
var table_PrivateMessageIes = make(map[int]*M2APPRIVATEIES)

var order_PrivateMessageIes = make([]int, 0)

type MbmsServiceCountingRequestIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MCCH-Update-Time', 'CRITICALITY': 'reject', 'TYPE': 'MCCH-Update-Time', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBSFN-Area-ID', 'CRITICALITY': 'reject', 'TYPE': 'MBSFN-Area-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Counting-Request-Session', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Counting-Request-Session', 'PRESENCE': 'mandatory'}, None]}
   MCCHUpdateTime  MCCHUpdateTime
   MBSFNAreaID  MBSFNAreaID
   MBMSCountingRequestSession  MBMSCountingRequestSession
   list []interface{}
}
func (self *MbmsServiceCountingRequestIes)createOT() interface{}{
    return nil
}
var table_MbmsServiceCountingRequestIes = make(map[int]*M2APPROTOCOLIES)

var order_MbmsServiceCountingRequestIes = make([]int, 3)

func (self *MbmsServiceCountingRequestIes) GetIECount() int{
   count := 0
   count +=1 //self.MCCHUpdateTime
   count +=1 //self.MBSFNAreaID
   count +=1 //self.MBMSCountingRequestSession
   return count//ObjSet
}
func (self *MbmsServiceCountingRequestIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 25: //MCCHUpdateTime
        return true //self.MCCHUpdateTime
      case 29: //MBSFNAreaID
        return true //self.MBSFNAreaID
      case 32: //MBMSCountingRequestSession
        return true //self.MBMSCountingRequestSession
   }
   return false//ObjSet
}
func (self *MbmsServiceCountingRequestIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 25: //MCCHUpdateTime
        self.MCCHUpdateTime.Unpack(st)
        self.list = append(self.list, &self.MCCHUpdateTime)
      case 29: //MBSFNAreaID
        self.MBSFNAreaID.Unpack(st)
        self.list = append(self.list, &self.MBSFNAreaID)
      case 32: //MBMSCountingRequestSession
        self.MBMSCountingRequestSession.Unpack(st)
        self.list = append(self.list, &self.MBMSCountingRequestSession)
   }
}
func (self *MbmsServiceCountingRequestIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 25: //MCCHUpdateTime
        self.MCCHUpdateTime.Pack(st)
      case 29: //MBSFNAreaID
        self.MBSFNAreaID.Pack(st)
      case 32: //MBMSCountingRequestSession
        self.MBMSCountingRequestSession.Pack(st)
      default:
      break
   }
}
func init() {
table_MbmsServiceCountingRequestIes[25] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMCCHUpdateTime}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCCHUpdateTime{}, PRESENCE:Presence{Presencemandatory}, }
order_MbmsServiceCountingRequestIes[0] = 25
table_MbmsServiceCountingRequestIes[29] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBSFNAreaID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBSFNAreaID{}, PRESENCE:Presence{Presencemandatory}, }
order_MbmsServiceCountingRequestIes[1] = 29
table_MbmsServiceCountingRequestIes[32] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSCountingRequestSession}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSCountingRequestSession{}, PRESENCE:Presence{Presencemandatory}, }
order_MbmsServiceCountingRequestIes[2] = 32
   }

type MBMSCountingRequestSessionItem struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBMS-Counting-Request-Session-Item', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Counting-Request-SessionIE', 'PRESENCE': 'mandatory'}, None]}
   MBMSCountingRequestSessionItem  MBMSCountingRequestSessionIE
   list []interface{}
}
func (self *MBMSCountingRequestSessionItem)createOT() interface{}{
    return nil
}
var table_MBMSCountingRequestSessionItem = make(map[int]*M2APPROTOCOLIES)

var order_MBMSCountingRequestSessionItem = make([]int, 1)

func (self *MBMSCountingRequestSessionItem) GetIECount() int{
   count := 0
   count +=1 //self.MBMSCountingRequestSessionItem
   return count//ObjSet
}
func (self *MBMSCountingRequestSessionItem) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 33: //MBMSCountingRequestSessionItem
        return true //self.MBMSCountingRequestSessionItem
   }
   return false//ObjSet
}
func (self *MBMSCountingRequestSessionItem)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 33: //MBMSCountingRequestSessionItem
        self.MBMSCountingRequestSessionItem.Unpack(st)
        self.list = append(self.list, &self.MBMSCountingRequestSessionItem)
   }
}
func (self *MBMSCountingRequestSessionItem)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 33: //MBMSCountingRequestSessionItem
        self.MBMSCountingRequestSessionItem.Pack(st)
      default:
      break
   }
}
func init() {
table_MBMSCountingRequestSessionItem[33] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSCountingRequestSessionItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSCountingRequestSessionIE{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSCountingRequestSessionItem[0] = 33
   }

type MBMSCountingRequestSessionIEExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MBMSCountingRequestSessionIEExtIEs)createOT() interface{}{
    return nil
}
var table_MBMSCountingRequestSessionIEExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_MBMSCountingRequestSessionIEExtIEs = make([]int, 0)

type MbmsServiceCountingResponseIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MbmsServiceCountingResponseIes)createOT() interface{}{
    return nil
}
var table_MbmsServiceCountingResponseIes = make(map[int]*M2APPROTOCOLIES)

var order_MbmsServiceCountingResponseIes = make([]int, 1)

func (self *MbmsServiceCountingResponseIes) GetIECount() int{
   count := 0
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MbmsServiceCountingResponseIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MbmsServiceCountingResponseIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MbmsServiceCountingResponseIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MbmsServiceCountingResponseIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MbmsServiceCountingResponseIes[0] = 8
   }

type MbmsServiceCountingFailureIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MbmsServiceCountingFailureIes)createOT() interface{}{
    return nil
}
var table_MbmsServiceCountingFailureIes = make(map[int]*M2APPROTOCOLIES)

var order_MbmsServiceCountingFailureIes = make([]int, 2)

func (self *MbmsServiceCountingFailureIes) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MbmsServiceCountingFailureIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        return true //self.Cause
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MbmsServiceCountingFailureIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MbmsServiceCountingFailureIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MbmsServiceCountingFailureIes[9] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_MbmsServiceCountingFailureIes[0] = 9
table_MbmsServiceCountingFailureIes[8] = &M2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MbmsServiceCountingFailureIes[1] = 8
   }

type MbmsServiceCountingResultsReportIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBSFN-Area-ID', 'CRITICALITY': 'reject', 'TYPE': 'MBSFN-Area-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Counting-Result-List', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Counting-Result-List', 'PRESENCE': 'mandatory'}, None]}
   MBSFNAreaID  MBSFNAreaID
   MBMSCountingResultList  MBMSCountingResultList
   list []interface{}
}
func (self *MbmsServiceCountingResultsReportIes)createOT() interface{}{
    return nil
}
var table_MbmsServiceCountingResultsReportIes = make(map[int]*M2APPROTOCOLIES)

var order_MbmsServiceCountingResultsReportIes = make([]int, 2)

func (self *MbmsServiceCountingResultsReportIes) GetIECount() int{
   count := 0
   count +=1 //self.MBSFNAreaID
   count +=1 //self.MBMSCountingResultList
   return count//ObjSet
}
func (self *MbmsServiceCountingResultsReportIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 29: //MBSFNAreaID
        return true //self.MBSFNAreaID
      case 34: //MBMSCountingResultList
        return true //self.MBMSCountingResultList
   }
   return false//ObjSet
}
func (self *MbmsServiceCountingResultsReportIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 29: //MBSFNAreaID
        self.MBSFNAreaID.Unpack(st)
        self.list = append(self.list, &self.MBSFNAreaID)
      case 34: //MBMSCountingResultList
        self.MBMSCountingResultList.Unpack(st)
        self.list = append(self.list, &self.MBMSCountingResultList)
   }
}
func (self *MbmsServiceCountingResultsReportIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 29: //MBSFNAreaID
        self.MBSFNAreaID.Pack(st)
      case 34: //MBMSCountingResultList
        self.MBMSCountingResultList.Pack(st)
      default:
      break
   }
}
func init() {
table_MbmsServiceCountingResultsReportIes[29] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBSFNAreaID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBSFNAreaID{}, PRESENCE:Presence{Presencemandatory}, }
order_MbmsServiceCountingResultsReportIes[0] = 29
table_MbmsServiceCountingResultsReportIes[34] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSCountingResultList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSCountingResultList{}, PRESENCE:Presence{Presencemandatory}, }
order_MbmsServiceCountingResultsReportIes[1] = 34
   }

type MBMSCountingResultItem struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBMS-Counting-Result-Item', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Counting-Result', 'PRESENCE': 'mandatory'}, None]}
   MBMSCountingResultItem  MBMSCountingResult
   list []interface{}
}
func (self *MBMSCountingResultItem)createOT() interface{}{
    return nil
}
var table_MBMSCountingResultItem = make(map[int]*M2APPROTOCOLIES)

var order_MBMSCountingResultItem = make([]int, 1)

func (self *MBMSCountingResultItem) GetIECount() int{
   count := 0
   count +=1 //self.MBMSCountingResultItem
   return count//ObjSet
}
func (self *MBMSCountingResultItem) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 35: //MBMSCountingResultItem
        return true //self.MBMSCountingResultItem
   }
   return false//ObjSet
}
func (self *MBMSCountingResultItem)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 35: //MBMSCountingResultItem
        self.MBMSCountingResultItem.Unpack(st)
        self.list = append(self.list, &self.MBMSCountingResultItem)
   }
}
func (self *MBMSCountingResultItem)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 35: //MBMSCountingResultItem
        self.MBMSCountingResultItem.Pack(st)
      default:
      break
   }
}
func init() {
table_MBMSCountingResultItem[35] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBMSCountingResultItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSCountingResult{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSCountingResultItem[0] = 35
   }

type MBMSCountingResultExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MBMSCountingResultExtIEs)createOT() interface{}{
    return nil
}
var table_MBMSCountingResultExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_MBMSCountingResultExtIEs = make([]int, 0)

type MbmsOverloadNotificationIes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBSFN-Area-ID', 'CRITICALITY': 'reject', 'TYPE': 'MBSFN-Area-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Overload-Status-Per-PMCH-List', 'CRITICALITY': 'reject', 'TYPE': 'Overload-Status-Per-PMCH-List', 'PRESENCE': 'mandatory'}, None]}
   MBSFNAreaID  MBSFNAreaID
   OverloadStatusPerPMCHList  OverloadStatusPerPMCHList
   list []interface{}
}
func (self *MbmsOverloadNotificationIes)createOT() interface{}{
    return nil
}
var table_MbmsOverloadNotificationIes = make(map[int]*M2APPROTOCOLIES)

var order_MbmsOverloadNotificationIes = make([]int, 2)

func (self *MbmsOverloadNotificationIes) GetIECount() int{
   count := 0
   count +=1 //self.MBSFNAreaID
   count +=1 //self.OverloadStatusPerPMCHList
   return count//ObjSet
}
func (self *MbmsOverloadNotificationIes) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 29: //MBSFNAreaID
        return true //self.MBSFNAreaID
      case 39: //OverloadStatusPerPMCHList
        return true //self.OverloadStatusPerPMCHList
   }
   return false//ObjSet
}
func (self *MbmsOverloadNotificationIes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 29: //MBSFNAreaID
        self.MBSFNAreaID.Unpack(st)
        self.list = append(self.list, &self.MBSFNAreaID)
      case 39: //OverloadStatusPerPMCHList
        self.OverloadStatusPerPMCHList.Unpack(st)
        self.list = append(self.list, &self.OverloadStatusPerPMCHList)
   }
}
func (self *MbmsOverloadNotificationIes)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 29: //MBSFNAreaID
        self.MBSFNAreaID.Pack(st)
      case 39: //OverloadStatusPerPMCHList
        self.OverloadStatusPerPMCHList.Pack(st)
      default:
      break
   }
}
func init() {
table_MbmsOverloadNotificationIes[29] = &M2APPROTOCOLIES{ID:ProtocolIEID{idMBSFNAreaID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBSFNAreaID{}, PRESENCE:Presence{Presencemandatory}, }
order_MbmsOverloadNotificationIes[0] = 29
table_MbmsOverloadNotificationIes[39] = &M2APPROTOCOLIES{ID:ProtocolIEID{idOverloadStatusPerPMCHList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&OverloadStatusPerPMCHList{}, PRESENCE:Presence{Presencemandatory}, }
order_MbmsOverloadNotificationIes[1] = 39
   }

type OverloadStatusPerPMCHItem struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-PMCH-Overload-Status', 'CRITICALITY': 'reject', 'TYPE': 'PMCH-Overload-Status', 'PRESENCE': 'mandatory'}, {'ID': 'id-Active-MBMS-Session-List', 'CRITICALITY': 'reject', 'TYPE': 'Active-MBMS-Session-List', 'PRESENCE': 'optional'}, None]}
   PMCHOverloadStatus  PMCHOverloadStatus
   ActiveMBMSSessionList  *ActiveMBMSSessionList
   list []interface{}
}
func (self *OverloadStatusPerPMCHItem)createOT() interface{}{
    return nil
}
var table_OverloadStatusPerPMCHItem = make(map[int]*M2APPROTOCOLIES)

var order_OverloadStatusPerPMCHItem = make([]int, 2)

func (self *OverloadStatusPerPMCHItem) GetIECount() int{
   count := 0
   count +=1 //self.PMCHOverloadStatus
   if self.ActiveMBMSSessionList != nil { count += 1 }
   return count//ObjSet
}
func (self *OverloadStatusPerPMCHItem) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 41: //PMCHOverloadStatus
        return true //self.PMCHOverloadStatus
      case 42: //ActiveMBMSSessionList
        if self.ActiveMBMSSessionList != nil { return true }
   }
   return false//ObjSet
}
func (self *OverloadStatusPerPMCHItem)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 41: //PMCHOverloadStatus
        self.PMCHOverloadStatus.Unpack(st)
        self.list = append(self.list, &self.PMCHOverloadStatus)
      case 42: //ActiveMBMSSessionList
        self.ActiveMBMSSessionList = &ActiveMBMSSessionList{}
        self.ActiveMBMSSessionList.Unpack(st)
        self.list = append(self.list, self.ActiveMBMSSessionList)
   }
}
func (self *OverloadStatusPerPMCHItem)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 41: //PMCHOverloadStatus
        self.PMCHOverloadStatus.Pack(st)
      case 42: //ActiveMBMSSessionList
        if self.ActiveMBMSSessionList != nil {self.ActiveMBMSSessionList.Pack(st)}
      default:
      break
   }
}
func init() {
table_OverloadStatusPerPMCHItem[41] = &M2APPROTOCOLIES{ID:ProtocolIEID{idPMCHOverloadStatus}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PMCHOverloadStatus{}, PRESENCE:Presence{Presencemandatory}, }
order_OverloadStatusPerPMCHItem[0] = 41
table_OverloadStatusPerPMCHItem[42] = &M2APPROTOCOLIES{ID:ProtocolIEID{idActiveMBMSSessionList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ActiveMBMSSessionList{}, PRESENCE:Presence{Presenceoptional}, }
order_OverloadStatusPerPMCHItem[1] = 42
   }

type ActiveMBMSSessionItem struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TMGI', 'CRITICALITY': 'reject', 'TYPE': 'TMGI', 'PRESENCE': 'mandatory'}, None]}
   TMGI  TMGI
   list []interface{}
}
func (self *ActiveMBMSSessionItem)createOT() interface{}{
    return nil
}
var table_ActiveMBMSSessionItem = make(map[int]*M2APPROTOCOLIES)

var order_ActiveMBMSSessionItem = make([]int, 1)

func (self *ActiveMBMSSessionItem) GetIECount() int{
   count := 0
   count +=1 //self.TMGI
   return count//ObjSet
}
func (self *ActiveMBMSSessionItem) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 2: //TMGI
        return true //self.TMGI
   }
   return false//ObjSet
}
func (self *ActiveMBMSSessionItem)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 2: //TMGI
        self.TMGI.Unpack(st)
        self.list = append(self.list, &self.TMGI)
   }
}
func (self *ActiveMBMSSessionItem)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLIESid).Value
   switch cat {
      case 2: //TMGI
        self.TMGI.Pack(st)
      default:
      break
   }
}
func init() {
table_ActiveMBMSSessionItem[2] = &M2APPROTOCOLIES{ID:ProtocolIEID{idTMGI}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TMGI{}, PRESENCE:Presence{Presencemandatory}, }
order_ActiveMBMSSessionItem[0] = 2
   }

type AllocationAndRetentionPriorityExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *AllocationAndRetentionPriorityExtIEs)createOT() interface{}{
    return nil
}
var table_AllocationAndRetentionPriorityExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_AllocationAndRetentionPriorityExtIEs = make([]int, 0)

type CellInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CellInformationExtIEs)createOT() interface{}{
    return nil
}
var table_CellInformationExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_CellInformationExtIEs = make([]int, 0)

type CriticalityDiagnosticsExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsExtIEs = make([]int, 0)

type CriticalityDiagnosticsIEListExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsIEListExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsIEListExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsIEListExtIEs = make([]int, 0)

type ECGIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *ECGIExtIEs)createOT() interface{}{
    return nil
}
var table_ECGIExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_ECGIExtIEs = make([]int, 0)

type ENBMBMSConfigurationdataItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *ENBMBMSConfigurationdataItemExtIEs)createOT() interface{}{
    return nil
}
var table_ENBMBMSConfigurationdataItemExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_ENBMBMSConfigurationdataItemExtIEs = make([]int, 0)

type GBRQosInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GBRQosInformationExtIEs)createOT() interface{}{
    return nil
}
var table_GBRQosInformationExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_GBRQosInformationExtIEs = make([]int, 0)

type GlobalENBIDExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GlobalENBIDExtIEs)createOT() interface{}{
    return nil
}
var table_GlobalENBIDExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_GlobalENBIDExtIEs = make([]int, 0)

type GlobalMCEIDExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GlobalMCEIDExtIEs)createOT() interface{}{
    return nil
}
var table_GlobalMCEIDExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_GlobalMCEIDExtIEs = make([]int, 0)

type MBMSERABQoSParametersExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MBMSERABQoSParametersExtIEs)createOT() interface{}{
    return nil
}
var table_MBMSERABQoSParametersExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_MBMSERABQoSParametersExtIEs = make([]int, 0)

type MBMSServiceassociatedLogicalM2ConnectionItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MBMSServiceassociatedLogicalM2ConnectionItemExtIEs)createOT() interface{}{
    return nil
}
var table_MBMSServiceassociatedLogicalM2ConnectionItemExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_MBMSServiceassociatedLogicalM2ConnectionItemExtIEs = make([]int, 0)

type MBMSsessionListPerPMCHItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MBMSsessionListPerPMCHItemExtIEs)createOT() interface{}{
    return nil
}
var table_MBMSsessionListPerPMCHItemExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_MBMSsessionListPerPMCHItemExtIEs = make([]int, 0)

type MBMSsessionsToBeSuspendedListPerPMCHItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MBMSsessionsToBeSuspendedListPerPMCHItemExtIEs)createOT() interface{}{
    return nil
}
var table_MBMSsessionsToBeSuspendedListPerPMCHItemExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_MBMSsessionsToBeSuspendedListPerPMCHItemExtIEs = make([]int, 0)

type MBSFNSubframeConfigurationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-SubframeAllocationExtended', 'CRITICALITY': 'reject', 'EXTENSION': 'SubframeAllocationExtended', 'PRESENCE': 'optional'}, None]}
   SubframeAllocationExtended  *SubframeAllocationExtended
   list []interface{}
}
func (self *MBSFNSubframeConfigurationExtIEs)createOT() interface{}{
    return nil
}
var table_MBSFNSubframeConfigurationExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_MBSFNSubframeConfigurationExtIEs = make([]int, 1)

func (self *MBSFNSubframeConfigurationExtIEs) GetIECount() int{
   count := 0
   if self.SubframeAllocationExtended != nil { count += 1 }
   return count//ObjSet
}
func (self *MBSFNSubframeConfigurationExtIEs) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 50: //SubframeAllocationExtended
        if self.SubframeAllocationExtended != nil { return true }
   }
   return false//ObjSet
}
func (self *MBSFNSubframeConfigurationExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 50: //SubframeAllocationExtended
        self.SubframeAllocationExtended = &SubframeAllocationExtended{}
        self.SubframeAllocationExtended.Unpack(st)
        self.list = append(self.list, self.SubframeAllocationExtended)
   }
}
func (self *MBSFNSubframeConfigurationExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 50: //SubframeAllocationExtended
        if self.SubframeAllocationExtended != nil {self.SubframeAllocationExtended.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBSFNSubframeConfigurationExtIEs[50] = &M2APPROTOCOLEXTENSION{ID:ProtocolIEID{idSubframeAllocationExtended}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&SubframeAllocationExtended{}, PRESENCE:Presence{Presenceoptional}, }
order_MBSFNSubframeConfigurationExtIEs[0] = 50
   }

type MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-Repetition-PeriodExtended', 'CRITICALITY': 'reject', 'EXTENSION': 'Repetition-PeriodExtended', 'PRESENCE': 'optional'}, {'ID': 'id-Modification-PeriodExtended', 'CRITICALITY': 'reject', 'EXTENSION': 'Modification-PeriodExtended', 'PRESENCE': 'optional'}, {'ID': 'id-Subcarrier-SpacingMBMS', 'CRITICALITY': 'reject', 'EXTENSION': 'Subcarrier-SpacingMBMS', 'PRESENCE': 'optional'}, None]}
   RepetitionPeriodExtended  *RepetitionPeriodExtended
   ModificationPeriodExtended  *ModificationPeriodExtended
   SubcarrierSpacingMBMS  *SubcarrierSpacingMBMS
   list []interface{}
}
func (self *MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs)createOT() interface{}{
    return nil
}
var table_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs = make([]int, 3)

func (self *MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs) GetIECount() int{
   count := 0
   if self.RepetitionPeriodExtended != nil { count += 1 }
   if self.ModificationPeriodExtended != nil { count += 1 }
   if self.SubcarrierSpacingMBMS != nil { count += 1 }
   return count//ObjSet
}
func (self *MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 47: //RepetitionPeriodExtended
        if self.RepetitionPeriodExtended != nil { return true }
      case 46: //ModificationPeriodExtended
        if self.ModificationPeriodExtended != nil { return true }
      case 49: //SubcarrierSpacingMBMS
        if self.SubcarrierSpacingMBMS != nil { return true }
   }
   return false//ObjSet
}
func (self *MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 47: //RepetitionPeriodExtended
        self.RepetitionPeriodExtended = &RepetitionPeriodExtended{}
        self.RepetitionPeriodExtended.Unpack(st)
        self.list = append(self.list, self.RepetitionPeriodExtended)
      case 46: //ModificationPeriodExtended
        self.ModificationPeriodExtended = &ModificationPeriodExtended{}
        self.ModificationPeriodExtended.Unpack(st)
        self.list = append(self.list, self.ModificationPeriodExtended)
      case 49: //SubcarrierSpacingMBMS
        self.SubcarrierSpacingMBMS = &SubcarrierSpacingMBMS{}
        self.SubcarrierSpacingMBMS.Unpack(st)
        self.list = append(self.list, self.SubcarrierSpacingMBMS)
   }
}
func (self *MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 47: //RepetitionPeriodExtended
        if self.RepetitionPeriodExtended != nil {self.RepetitionPeriodExtended.Pack(st)}
      case 46: //ModificationPeriodExtended
        if self.ModificationPeriodExtended != nil {self.ModificationPeriodExtended.Pack(st)}
      case 49: //SubcarrierSpacingMBMS
        if self.SubcarrierSpacingMBMS != nil {self.SubcarrierSpacingMBMS.Pack(st)}
      default:
      break
   }
}
func init() {
table_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs[47] = &M2APPROTOCOLEXTENSION{ID:ProtocolIEID{idRepetitionPeriodExtended}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&RepetitionPeriodExtended{}, PRESENCE:Presence{Presenceoptional}, }
order_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs[0] = 47
table_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs[46] = &M2APPROTOCOLEXTENSION{ID:ProtocolIEID{idModificationPeriodExtended}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&ModificationPeriodExtended{}, PRESENCE:Presence{Presenceoptional}, }
order_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs[1] = 46
table_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs[49] = &M2APPROTOCOLEXTENSION{ID:ProtocolIEID{idSubcarrierSpacingMBMS}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&SubcarrierSpacingMBMS{}, PRESENCE:Presence{Presenceoptional}, }
order_MCCHrelatedBCCHConfigPerMBSFNAreaItemExtIEs[2] = 49
   }

type PMCHConfigurationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-Modulation-Coding-Scheme2', 'CRITICALITY': 'reject', 'EXTENSION': 'Modulation-Coding-Scheme2', 'PRESENCE': 'optional'}, {'ID': 'id-MCH-Scheduling-PeriodExtended', 'CRITICALITY': 'reject', 'EXTENSION': 'MCH-Scheduling-PeriodExtended', 'PRESENCE': 'optional'}, {'ID': 'id-MCH-Scheduling-PeriodExtended2', 'CRITICALITY': 'reject', 'EXTENSION': 'MCH-Scheduling-PeriodExtended2', 'PRESENCE': 'optional'}, None]}
   ModulationCodingScheme2  *ModulationCodingScheme2
   MCHSchedulingPeriodExtended  *MCHSchedulingPeriodExtended
   MCHSchedulingPeriodExtended2  *MCHSchedulingPeriodExtended2
   list []interface{}
}
func (self *PMCHConfigurationExtIEs)createOT() interface{}{
    return nil
}
var table_PMCHConfigurationExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_PMCHConfigurationExtIEs = make([]int, 3)

func (self *PMCHConfigurationExtIEs) GetIECount() int{
   count := 0
   if self.ModulationCodingScheme2 != nil { count += 1 }
   if self.MCHSchedulingPeriodExtended != nil { count += 1 }
   if self.MCHSchedulingPeriodExtended2 != nil { count += 1 }
   return count//ObjSet
}
func (self *PMCHConfigurationExtIEs) GetOT(id interface{}) bool{
   cat := id.(M2APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 36: //ModulationCodingScheme2
        if self.ModulationCodingScheme2 != nil { return true }
      case 37: //MCHSchedulingPeriodExtended
        if self.MCHSchedulingPeriodExtended != nil { return true }
      case 48: //MCHSchedulingPeriodExtended2
        if self.MCHSchedulingPeriodExtended2 != nil { return true }
   }
   return false//ObjSet
}
func (self *PMCHConfigurationExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 36: //ModulationCodingScheme2
        self.ModulationCodingScheme2 = &ModulationCodingScheme2{}
        self.ModulationCodingScheme2.Unpack(st)
        self.list = append(self.list, self.ModulationCodingScheme2)
      case 37: //MCHSchedulingPeriodExtended
        self.MCHSchedulingPeriodExtended = &MCHSchedulingPeriodExtended{}
        self.MCHSchedulingPeriodExtended.Unpack(st)
        self.list = append(self.list, self.MCHSchedulingPeriodExtended)
      case 48: //MCHSchedulingPeriodExtended2
        self.MCHSchedulingPeriodExtended2 = &MCHSchedulingPeriodExtended2{}
        self.MCHSchedulingPeriodExtended2.Unpack(st)
        self.list = append(self.list, self.MCHSchedulingPeriodExtended2)
   }
}
func (self *PMCHConfigurationExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M2APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 36: //ModulationCodingScheme2
        if self.ModulationCodingScheme2 != nil {self.ModulationCodingScheme2.Pack(st)}
      case 37: //MCHSchedulingPeriodExtended
        if self.MCHSchedulingPeriodExtended != nil {self.MCHSchedulingPeriodExtended.Pack(st)}
      case 48: //MCHSchedulingPeriodExtended2
        if self.MCHSchedulingPeriodExtended2 != nil {self.MCHSchedulingPeriodExtended2.Pack(st)}
      default:
      break
   }
}
func init() {
table_PMCHConfigurationExtIEs[36] = &M2APPROTOCOLEXTENSION{ID:ProtocolIEID{idModulationCodingScheme2}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&ModulationCodingScheme2{}, PRESENCE:Presence{Presenceoptional}, }
order_PMCHConfigurationExtIEs[0] = 36
table_PMCHConfigurationExtIEs[37] = &M2APPROTOCOLEXTENSION{ID:ProtocolIEID{idMCHSchedulingPeriodExtended}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&MCHSchedulingPeriodExtended{}, PRESENCE:Presence{Presenceoptional}, }
order_PMCHConfigurationExtIEs[1] = 37
table_PMCHConfigurationExtIEs[48] = &M2APPROTOCOLEXTENSION{ID:ProtocolIEID{idMCHSchedulingPeriodExtended2}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&MCHSchedulingPeriodExtended2{}, PRESENCE:Presence{Presenceoptional}, }
order_PMCHConfigurationExtIEs[2] = 48
   }

type SCPTMInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SCPTMInformationExtIEs)createOT() interface{}{
    return nil
}
var table_SCPTMInformationExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_SCPTMInformationExtIEs = make([]int, 0)

type SubframeAllocationExtendedExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *SubframeAllocationExtendedExtIEs)createOT() interface{}{
    return nil
}
var table_SubframeAllocationExtendedExtIEs = make(map[int]*M2APPROTOCOLIES)

var order_SubframeAllocationExtendedExtIEs = make([]int, 0)

type TMGIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TMGIExtIEs)createOT() interface{}{
    return nil
}
var table_TMGIExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_TMGIExtIEs = make([]int, 0)

type TNLInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M2AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TNLInformationExtIEs)createOT() interface{}{
    return nil
}
var table_TNLInformationExtIEs = make(map[int]*M2APPROTOCOLEXTENSION)

var order_TNLInformationExtIEs = make([]int, 0)

func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'SessionStartRequest', 'SUCCESSFUL OUTCOME': 'SessionStartResponse', 'UNSUCCESSFUL OUTCOME': 'SessionStartFailure', 'PROCEDURE CODE': 'id-sessionStart', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idsessionStart}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&SessionStartRequest{}, SUCCESSFULOUTCOME:&SessionStartResponse{}, UNSUCCESSFULOUTCOME:&SessionStartFailure{}, PROCEDURECODE:ProcedureCode{idsessionStart}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetSessionStartINITIATINGMESSAGE() (*SessionStartRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &SessionStartRequest{}, uint64(idsessionStart), int(Criticalityreject)
}
func GetSessionStartSUCCESSFULOUTCOME() (*SessionStartResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &SessionStartResponse{}, uint64(idsessionStart), int(Criticalityreject)
}
func GetSessionStartUNSUCCESSFULOUTCOME() (*SessionStartFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &SessionStartFailure{}, uint64(idsessionStart), int(Criticalityreject)
}
func (self *SessionStartRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SessionStartRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SessionStartRequest) createOT() interface{} {
   return &SessionStartRequest{}
}
func (self *SessionStartRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SessionStartRequest) GetIECount() int{
    return 0
}
func (self *SessionStartResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SessionStartResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SessionStartResponse) createOT() interface{} {
   return &SessionStartResponse{}
}
func (self *SessionStartResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SessionStartResponse) GetIECount() int{
    return 0
}
func (self *SessionStartFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SessionStartFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SessionStartFailure) createOT() interface{} {
   return &SessionStartFailure{}
}
func (self *SessionStartFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SessionStartFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'SessionStopRequest', 'SUCCESSFUL OUTCOME': 'SessionStopResponse', 'PROCEDURE CODE': 'id-sessionStop', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idsessionStop}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&SessionStopRequest{}, SUCCESSFULOUTCOME:&SessionStopResponse{}, PROCEDURECODE:ProcedureCode{idsessionStop}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetSessionStopINITIATINGMESSAGE() (*SessionStopRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &SessionStopRequest{}, uint64(idsessionStop), int(Criticalityreject)
}
func GetSessionStopSUCCESSFULOUTCOME() (*SessionStopResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &SessionStopResponse{}, uint64(idsessionStop), int(Criticalityreject)
}
func (self *SessionStopRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SessionStopRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SessionStopRequest) createOT() interface{} {
   return &SessionStopRequest{}
}
func (self *SessionStopRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SessionStopRequest) GetIECount() int{
    return 0
}
func (self *SessionStopResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SessionStopResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SessionStopResponse) createOT() interface{} {
   return &SessionStopResponse{}
}
func (self *SessionStopResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SessionStopResponse) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'SessionUpdateRequest', 'SUCCESSFUL OUTCOME': 'SessionUpdateResponse', 'UNSUCCESSFUL OUTCOME': 'SessionUpdateFailure', 'PROCEDURE CODE': 'id-sessionUpdate', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idsessionUpdate}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&SessionUpdateRequest{}, SUCCESSFULOUTCOME:&SessionUpdateResponse{}, UNSUCCESSFULOUTCOME:&SessionUpdateFailure{}, PROCEDURECODE:ProcedureCode{idsessionUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetSessionUpdateINITIATINGMESSAGE() (*SessionUpdateRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &SessionUpdateRequest{}, uint64(idsessionUpdate), int(Criticalityreject)
}
func GetSessionUpdateSUCCESSFULOUTCOME() (*SessionUpdateResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &SessionUpdateResponse{}, uint64(idsessionUpdate), int(Criticalityreject)
}
func GetSessionUpdateUNSUCCESSFULOUTCOME() (*SessionUpdateFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &SessionUpdateFailure{}, uint64(idsessionUpdate), int(Criticalityreject)
}
func (self *SessionUpdateRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SessionUpdateRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SessionUpdateRequest) createOT() interface{} {
   return &SessionUpdateRequest{}
}
func (self *SessionUpdateRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SessionUpdateRequest) GetIECount() int{
    return 0
}
func (self *SessionUpdateResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SessionUpdateResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SessionUpdateResponse) createOT() interface{} {
   return &SessionUpdateResponse{}
}
func (self *SessionUpdateResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SessionUpdateResponse) GetIECount() int{
    return 0
}
func (self *SessionUpdateFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SessionUpdateFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SessionUpdateFailure) createOT() interface{} {
   return &SessionUpdateFailure{}
}
func (self *SessionUpdateFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SessionUpdateFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'MbmsSchedulingInformation', 'SUCCESSFUL OUTCOME': 'MbmsSchedulingInformationResponse', 'PROCEDURE CODE': 'id-mbmsSchedulingInformation', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idmbmsSchedulingInformation}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&MbmsSchedulingInformation{}, SUCCESSFULOUTCOME:&MbmsSchedulingInformationResponse{}, PROCEDURECODE:ProcedureCode{idmbmsSchedulingInformation}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetMbmsSchedulingInformationINITIATINGMESSAGE() (*MbmsSchedulingInformation, uint64, int) {/*TYPE, ID, Cricality*/
 return &MbmsSchedulingInformation{}, uint64(idmbmsSchedulingInformation), int(Criticalityreject)
}
func GetMbmsSchedulingInformationSUCCESSFULOUTCOME() (*MbmsSchedulingInformationResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &MbmsSchedulingInformationResponse{}, uint64(idmbmsSchedulingInformation), int(Criticalityreject)
}
func (self *MbmsSchedulingInformation) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MbmsSchedulingInformation) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MbmsSchedulingInformation) createOT() interface{} {
   return &MbmsSchedulingInformation{}
}
func (self *MbmsSchedulingInformation) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MbmsSchedulingInformation) GetIECount() int{
    return 0
}
func (self *MbmsSchedulingInformationResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MbmsSchedulingInformationResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MbmsSchedulingInformationResponse) createOT() interface{} {
   return &MbmsSchedulingInformationResponse{}
}
func (self *MbmsSchedulingInformationResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MbmsSchedulingInformationResponse) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ErrorIndication', 'PROCEDURE CODE': 'id-errorIndication', 'CRITICALITY': 'ignore'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{iderrorIndication}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ErrorIndication{}, PROCEDURECODE:ProcedureCode{iderrorIndication}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'Reset', 'SUCCESSFUL OUTCOME': 'ResetAcknowledge', 'PROCEDURE CODE': 'id-reset', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idreset}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&Reset{}, SUCCESSFULOUTCOME:&ResetAcknowledge{}, PROCEDURECODE:ProcedureCode{idreset}, CRITICALITY:Criticality{Criticalityreject}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'M2SetupRequest', 'SUCCESSFUL OUTCOME': 'M2SetupResponse', 'UNSUCCESSFUL OUTCOME': 'M2SetupFailure', 'PROCEDURE CODE': 'id-m2Setup', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idm2Setup}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&M2SetupRequest{}, SUCCESSFULOUTCOME:&M2SetupResponse{}, UNSUCCESSFULOUTCOME:&M2SetupFailure{}, PROCEDURECODE:ProcedureCode{idm2Setup}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetM2SetupINITIATINGMESSAGE() (*M2SetupRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &M2SetupRequest{}, uint64(idm2Setup), int(Criticalityreject)
}
func GetM2SetupSUCCESSFULOUTCOME() (*M2SetupResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &M2SetupResponse{}, uint64(idm2Setup), int(Criticalityreject)
}
func GetM2SetupUNSUCCESSFULOUTCOME() (*M2SetupFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &M2SetupFailure{}, uint64(idm2Setup), int(Criticalityreject)
}
func (self *M2SetupRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *M2SetupRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *M2SetupRequest) createOT() interface{} {
   return &M2SetupRequest{}
}
func (self *M2SetupRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *M2SetupRequest) GetIECount() int{
    return 0
}
func (self *M2SetupResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *M2SetupResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *M2SetupResponse) createOT() interface{} {
   return &M2SetupResponse{}
}
func (self *M2SetupResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *M2SetupResponse) GetIECount() int{
    return 0
}
func (self *M2SetupFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *M2SetupFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *M2SetupFailure) createOT() interface{} {
   return &M2SetupFailure{}
}
func (self *M2SetupFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *M2SetupFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ENBConfigurationUpdate', 'SUCCESSFUL OUTCOME': 'ENBConfigurationUpdateAcknowledge', 'UNSUCCESSFUL OUTCOME': 'ENBConfigurationUpdateFailure', 'PROCEDURE CODE': 'id-eNBConfigurationUpdate', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{ideNBConfigurationUpdate}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ENBConfigurationUpdate{}, SUCCESSFULOUTCOME:&ENBConfigurationUpdateAcknowledge{}, UNSUCCESSFULOUTCOME:&ENBConfigurationUpdateFailure{}, PROCEDURECODE:ProcedureCode{ideNBConfigurationUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetENBConfigurationUpdateINITIATINGMESSAGE() (*ENBConfigurationUpdate, uint64, int) {/*TYPE, ID, Cricality*/
 return &ENBConfigurationUpdate{}, uint64(ideNBConfigurationUpdate), int(Criticalityreject)
}
func GetENBConfigurationUpdateSUCCESSFULOUTCOME() (*ENBConfigurationUpdateAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &ENBConfigurationUpdateAcknowledge{}, uint64(ideNBConfigurationUpdate), int(Criticalityreject)
}
func GetENBConfigurationUpdateUNSUCCESSFULOUTCOME() (*ENBConfigurationUpdateFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &ENBConfigurationUpdateFailure{}, uint64(ideNBConfigurationUpdate), int(Criticalityreject)
}
func (self *ENBConfigurationUpdate) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ENBConfigurationUpdate) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ENBConfigurationUpdate) createOT() interface{} {
   return &ENBConfigurationUpdate{}
}
func (self *ENBConfigurationUpdate) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ENBConfigurationUpdate) GetIECount() int{
    return 0
}
func (self *ENBConfigurationUpdateAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ENBConfigurationUpdateAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ENBConfigurationUpdateAcknowledge) createOT() interface{} {
   return &ENBConfigurationUpdateAcknowledge{}
}
func (self *ENBConfigurationUpdateAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ENBConfigurationUpdateAcknowledge) GetIECount() int{
    return 0
}
func (self *ENBConfigurationUpdateFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ENBConfigurationUpdateFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ENBConfigurationUpdateFailure) createOT() interface{} {
   return &ENBConfigurationUpdateFailure{}
}
func (self *ENBConfigurationUpdateFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ENBConfigurationUpdateFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'MCEConfigurationUpdate', 'SUCCESSFUL OUTCOME': 'MCEConfigurationUpdateAcknowledge', 'UNSUCCESSFUL OUTCOME': 'MCEConfigurationUpdateFailure', 'PROCEDURE CODE': 'id-mCEConfigurationUpdate', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idmCEConfigurationUpdate}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&MCEConfigurationUpdate{}, SUCCESSFULOUTCOME:&MCEConfigurationUpdateAcknowledge{}, UNSUCCESSFULOUTCOME:&MCEConfigurationUpdateFailure{}, PROCEDURECODE:ProcedureCode{idmCEConfigurationUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetMCEConfigurationUpdateINITIATINGMESSAGE() (*MCEConfigurationUpdate, uint64, int) {/*TYPE, ID, Cricality*/
 return &MCEConfigurationUpdate{}, uint64(idmCEConfigurationUpdate), int(Criticalityreject)
}
func GetMCEConfigurationUpdateSUCCESSFULOUTCOME() (*MCEConfigurationUpdateAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &MCEConfigurationUpdateAcknowledge{}, uint64(idmCEConfigurationUpdate), int(Criticalityreject)
}
func GetMCEConfigurationUpdateUNSUCCESSFULOUTCOME() (*MCEConfigurationUpdateFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &MCEConfigurationUpdateFailure{}, uint64(idmCEConfigurationUpdate), int(Criticalityreject)
}
func (self *MCEConfigurationUpdate) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MCEConfigurationUpdate) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MCEConfigurationUpdate) createOT() interface{} {
   return &MCEConfigurationUpdate{}
}
func (self *MCEConfigurationUpdate) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MCEConfigurationUpdate) GetIECount() int{
    return 0
}
func (self *MCEConfigurationUpdateAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MCEConfigurationUpdateAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MCEConfigurationUpdateAcknowledge) createOT() interface{} {
   return &MCEConfigurationUpdateAcknowledge{}
}
func (self *MCEConfigurationUpdateAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MCEConfigurationUpdateAcknowledge) GetIECount() int{
    return 0
}
func (self *MCEConfigurationUpdateFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MCEConfigurationUpdateFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MCEConfigurationUpdateFailure) createOT() interface{} {
   return &MCEConfigurationUpdateFailure{}
}
func (self *MCEConfigurationUpdateFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MCEConfigurationUpdateFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'MbmsServiceCountingRequest', 'SUCCESSFUL OUTCOME': 'MbmsServiceCountingResponse', 'UNSUCCESSFUL OUTCOME': 'MbmsServiceCountingFailure', 'PROCEDURE CODE': 'id-mbmsServiceCounting', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idmbmsServiceCounting}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&MbmsServiceCountingRequest{}, SUCCESSFULOUTCOME:&MbmsServiceCountingResponse{}, UNSUCCESSFULOUTCOME:&MbmsServiceCountingFailure{}, PROCEDURECODE:ProcedureCode{idmbmsServiceCounting}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetMbmsServiceCountingINITIATINGMESSAGE() (*MbmsServiceCountingRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &MbmsServiceCountingRequest{}, uint64(idmbmsServiceCounting), int(Criticalityreject)
}
func GetMbmsServiceCountingSUCCESSFULOUTCOME() (*MbmsServiceCountingResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &MbmsServiceCountingResponse{}, uint64(idmbmsServiceCounting), int(Criticalityreject)
}
func GetMbmsServiceCountingUNSUCCESSFULOUTCOME() (*MbmsServiceCountingFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &MbmsServiceCountingFailure{}, uint64(idmbmsServiceCounting), int(Criticalityreject)
}
func (self *MbmsServiceCountingRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MbmsServiceCountingRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MbmsServiceCountingRequest) createOT() interface{} {
   return &MbmsServiceCountingRequest{}
}
func (self *MbmsServiceCountingRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MbmsServiceCountingRequest) GetIECount() int{
    return 0
}
func (self *MbmsServiceCountingResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MbmsServiceCountingResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MbmsServiceCountingResponse) createOT() interface{} {
   return &MbmsServiceCountingResponse{}
}
func (self *MbmsServiceCountingResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MbmsServiceCountingResponse) GetIECount() int{
    return 0
}
func (self *MbmsServiceCountingFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MbmsServiceCountingFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MbmsServiceCountingFailure) createOT() interface{} {
   return &MbmsServiceCountingFailure{}
}
func (self *MbmsServiceCountingFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MbmsServiceCountingFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'MbmsServiceCountingResultsReport', 'PROCEDURE CODE': 'id-mbmsServiceCountingResultsReport', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idmbmsServiceCountingResultsReport}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&MbmsServiceCountingResultsReport{}, PROCEDURECODE:ProcedureCode{idmbmsServiceCountingResultsReport}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetMbmsServiceCountingResultsReportINITIATINGMESSAGE() (*MbmsServiceCountingResultsReport, uint64, int) {/*TYPE, ID, Cricality*/
 return &MbmsServiceCountingResultsReport{}, uint64(idmbmsServiceCountingResultsReport), int(Criticalityreject)
}
func (self *MbmsServiceCountingResultsReport) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MbmsServiceCountingResultsReport) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MbmsServiceCountingResultsReport) createOT() interface{} {
   return &MbmsServiceCountingResultsReport{}
}
func (self *MbmsServiceCountingResultsReport) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MbmsServiceCountingResultsReport) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'PrivateMessage', 'PROCEDURE CODE': 'id-privateMessage', 'CRITICALITY': 'ignore'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idprivateMessage}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&PrivateMessage{}, PROCEDURECODE:ProcedureCode{idprivateMessage}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'MbmsOverloadNotification', 'PROCEDURE CODE': 'id-mbmsOverloadNotification', 'CRITICALITY': 'reject'}]}
table_M2APELEMENTARYPROCEDURES[M2APELEMENTARYPROCEDUREprocedureCode{idmbmsOverloadNotification}] = &M2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&MbmsOverloadNotification{}, PROCEDURECODE:ProcedureCode{idmbmsOverloadNotification}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetMbmsOverloadNotificationINITIATINGMESSAGE() (*MbmsOverloadNotification, uint64, int) {/*TYPE, ID, Cricality*/
 return &MbmsOverloadNotification{}, uint64(idmbmsOverloadNotification), int(Criticalityreject)
}
func (self *MbmsOverloadNotification) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MbmsOverloadNotification) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MbmsOverloadNotification) createOT() interface{} {
   return &MbmsOverloadNotification{}
}
func (self *MbmsOverloadNotification) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MbmsOverloadNotification) GetIECount() int{
    return 0
}
var maxPrivateIEs uint64 = 65535
var maxProtocolExtensions uint64 = 65535
var maxProtocolIEs uint64 = 65535
var idsessionStart uint64 = 0
const ProcedureCodesessionStart = 0
var idsessionStop uint64 = 1
const ProcedureCodesessionStop = 1
var idmbmsSchedulingInformation uint64 = 2
const ProcedureCodembmsSchedulingInformation = 2
var iderrorIndication uint64 = 3
const ProcedureCodeerrorIndication = 3
var idreset uint64 = 4
const ProcedureCodereset = 4
var idm2Setup uint64 = 5
const ProcedureCodem2Setup = 5
var ideNBConfigurationUpdate uint64 = 6
const ProcedureCodeeNBConfigurationUpdate = 6
var idmCEConfigurationUpdate uint64 = 7
const ProcedureCodemCEConfigurationUpdate = 7
var idprivateMessage uint64 = 8
const ProcedureCodeprivateMessage = 8
var idsessionUpdate uint64 = 9
const ProcedureCodesessionUpdate = 9
var idmbmsServiceCounting uint64 = 10
const ProcedureCodembmsServiceCounting = 10
var idmbmsServiceCountingResultsReport uint64 = 11
const ProcedureCodembmsServiceCountingResultsReport = 11
var idmbmsOverloadNotification uint64 = 12
const ProcedureCodembmsOverloadNotification = 12
var maxnoofMBSFNareas uint64 = 256
var maxnoofMBSFNAllocations uint64 = 8
var maxnoofPMCHsperMBSFNarea uint64 = 15
var maxnoofCells uint64 = 256
var maxnoofMBMSServiceAreasPerCell uint64 = 256
var maxnoofSessionsPerPMCH uint64 = 29
var maxnooferrors uint64 = 256
var maxNrOfIndividualM2ConnectionsToReset uint64 = 256
var maxnoofCountingService uint64 = 16
var maxnoofCellsforMBMS uint64 = 4096
var idMCEMBMSM2APID uint64 = 0
const ProtocolIEIDMCEMBMSM2APID = 0
var idENBMBMSM2APID uint64 = 1
const ProtocolIEIDENBMBMSM2APID = 1
var idTMGI uint64 = 2
const ProtocolIEIDTMGI = 2
var idMBMSSessionID uint64 = 3
const ProtocolIEIDMBMSSessionID = 3
var idMBMSServiceArea uint64 = 6
const ProtocolIEIDMBMSServiceArea = 6
var idTNLInformation uint64 = 7
const ProtocolIEIDTNLInformation = 7
var idCriticalityDiagnostics uint64 = 8
const ProtocolIEIDCriticalityDiagnostics = 8
var idCause uint64 = 9
const ProtocolIEIDCause = 9
var idMBSFNAreaConfigurationList uint64 = 10
const ProtocolIEIDMBSFNAreaConfigurationList = 10
var idPMCHConfigurationList uint64 = 11
const ProtocolIEIDPMCHConfigurationList = 11
var idPMCHConfigurationItem uint64 = 12
const ProtocolIEIDPMCHConfigurationItem = 12
var idGlobalENBID uint64 = 13
const ProtocolIEIDGlobalENBID = 13
var idENBname uint64 = 14
const ProtocolIEIDENBname = 14
var idENBMBMSConfigurationdataList uint64 = 15
const ProtocolIEIDENBMBMSConfigurationdataList = 15
var idENBMBMSConfigurationdataItem uint64 = 16
const ProtocolIEIDENBMBMSConfigurationdataItem = 16
var idGlobalMCEID uint64 = 17
const ProtocolIEIDGlobalMCEID = 17
var idMCEname uint64 = 18
const ProtocolIEIDMCEname = 18
var idMCCHrelatedBCCHConfigPerMBSFNArea uint64 = 19
const ProtocolIEIDMCCHrelatedBCCHConfigPerMBSFNArea = 19
var idMCCHrelatedBCCHConfigPerMBSFNAreaItem uint64 = 20
const ProtocolIEIDMCCHrelatedBCCHConfigPerMBSFNAreaItem = 20
var idTimeToWait uint64 = 21
const ProtocolIEIDTimeToWait = 21
var idMBSFNSubframeConfigurationList uint64 = 22
const ProtocolIEIDMBSFNSubframeConfigurationList = 22
var idMBSFNSubframeConfigurationItem uint64 = 23
const ProtocolIEIDMBSFNSubframeConfigurationItem = 23
var idCommonSubframeAllocationPeriod uint64 = 24
const ProtocolIEIDCommonSubframeAllocationPeriod = 24
var idMCCHUpdateTime uint64 = 25
const ProtocolIEIDMCCHUpdateTime = 25
var idENBMBMSConfigurationdataListConfigUpdate uint64 = 26
const ProtocolIEIDENBMBMSConfigurationdataListConfigUpdate = 26
var idENBMBMSConfigurationdataConfigUpdateItem uint64 = 27
const ProtocolIEIDENBMBMSConfigurationdataConfigUpdateItem = 27
var idMBMSServiceassociatedLogicalM2ConnectionItem uint64 = 28
const ProtocolIEIDMBMSServiceassociatedLogicalM2ConnectionItem = 28
var idMBSFNAreaID uint64 = 29
const ProtocolIEIDMBSFNAreaID = 29
var idResetType uint64 = 30
const ProtocolIEIDResetType = 30
var idMBMSServiceassociatedLogicalM2ConnectionListResAck uint64 = 31
const ProtocolIEIDMBMSServiceassociatedLogicalM2ConnectionListResAck = 31
var idMBMSCountingRequestSession uint64 = 32
const ProtocolIEIDMBMSCountingRequestSession = 32
var idMBMSCountingRequestSessionItem uint64 = 33
const ProtocolIEIDMBMSCountingRequestSessionItem = 33
var idMBMSCountingResultList uint64 = 34
const ProtocolIEIDMBMSCountingResultList = 34
var idMBMSCountingResultItem uint64 = 35
const ProtocolIEIDMBMSCountingResultItem = 35
var idModulationCodingScheme2 uint64 = 36
const ProtocolIEIDModulationCodingScheme2 = 36
var idMCHSchedulingPeriodExtended uint64 = 37
const ProtocolIEIDMCHSchedulingPeriodExtended = 37
var idAlternativeTNLInformation uint64 = 38
const ProtocolIEIDAlternativeTNLInformation = 38
var idOverloadStatusPerPMCHList uint64 = 39
const ProtocolIEIDOverloadStatusPerPMCHList = 39
var idPMCHOverloadStatus uint64 = 41
const ProtocolIEIDPMCHOverloadStatus = 41
var idActiveMBMSSessionList uint64 = 42
const ProtocolIEIDActiveMBMSSessionList = 42
var idMBMSSuspensionNotificationList uint64 = 43
const ProtocolIEIDMBMSSuspensionNotificationList = 43
var idMBMSSuspensionNotificationItem uint64 = 44
const ProtocolIEIDMBMSSuspensionNotificationItem = 44
var idSCPTMInformation uint64 = 45
const ProtocolIEIDSCPTMInformation = 45
var idModificationPeriodExtended uint64 = 46
const ProtocolIEIDModificationPeriodExtended = 46
var idRepetitionPeriodExtended uint64 = 47
const ProtocolIEIDRepetitionPeriodExtended = 47
var idMCHSchedulingPeriodExtended2 uint64 = 48
const ProtocolIEIDMCHSchedulingPeriodExtended2 = 48
var idSubcarrierSpacingMBMS uint64 = 49
const ProtocolIEIDSubcarrierSpacingMBMS = 49
var idSubframeAllocationExtended uint64 = 50
const ProtocolIEIDSubframeAllocationExtended = 50
