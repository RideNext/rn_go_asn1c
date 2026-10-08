
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package m3ap
import (
  log "github.com/sirupsen/logrus"
)
var version = "vf50"

func fmtm3ap() {log.Debug("m3ap")}
func (self *M3APPDU)Unpack(stream *Stream) {
    //coptions := []string{"initiatingMessage","successfulOutcome","unsuccessfulOutcome","Unknown"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in M3APPDU\n", choice, choice_len)
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
func (self * M3APPDU) Pack(stream *Stream) {
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
type M3APPDU struct { //[{'type': 'InitiatingMessage', 'name': 'initiatingMessage'}, {'type': 'SuccessfulOutcome', 'name': 'successfulOutcome'}, {'type': 'UnsuccessfulOutcome', 'name': 'unsuccessfulOutcome'}, None]
    InitiatingMessage *InitiatingMessage
    SuccessfulOutcome *SuccessfulOutcome
    UnsuccessfulOutcome *UnsuccessfulOutcome
} // M3APPDU

type InitiatingMessage struct { // [{'type': 'M3AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M3AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'M3AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['M3AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'M3AP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['M3AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode M3APELEMENTARYPROCEDUREprocedureCode
    Criticality M3APELEMENTARYPROCEDUREcriticality
    Value M3APELEMENTARYPROCEDUREInitiatingMessage
}

func (self * InitiatingMessage) Unpack(stream *Stream) {
    //table {'type': 'M3AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M3AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_M3APELEMENTARYPROCEDURES[key].INITIATINGMESSAGE
    self.Value = out.(M3APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(M3APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'M3AP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['M3AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * InitiatingMessage) Pack(stream *Stream) {
    //table {'type': 'M3AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M3AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_M3APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(M3APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type SuccessfulOutcome struct { // [{'type': 'M3AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M3AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'M3AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['M3AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'M3AP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['M3AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode M3APELEMENTARYPROCEDUREprocedureCode
    Criticality M3APELEMENTARYPROCEDUREcriticality
    Value M3APELEMENTARYPROCEDURESuccessfulOutcome
}

func (self * SuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'M3AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M3AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_M3APELEMENTARYPROCEDURES[key].SUCCESSFULOUTCOME
    self.Value = out.(M3APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(M3APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'M3AP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['M3AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * SuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'M3AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M3AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_M3APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(M3APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type UnsuccessfulOutcome struct { // [{'type': 'M3AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M3AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'M3AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['M3AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'M3AP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['M3AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode M3APELEMENTARYPROCEDUREprocedureCode
    Criticality M3APELEMENTARYPROCEDUREcriticality
    Value M3APELEMENTARYPROCEDUREUnsuccessfulOutcome
}

func (self * UnsuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'M3AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M3AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_M3APELEMENTARYPROCEDURES[key].UNSUCCESSFULOUTCOME
    self.Value = out.(M3APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(M3APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'M3AP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['M3AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * UnsuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'M3AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'M3AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_M3APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(M3APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type MBMSSessionStartRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MBMSSessionStartRequest-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MBMSSessionStartRequestIEs
}

func (self * MBMSSessionStartRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MBMSSessionStartRequestIEs, order_MBMSSessionStartRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSSessionStartRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MBMSSessionStartRequestIEs, order_MBMSSessionStartRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MBMSSessionStartResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MBMSSessionStartResponse-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MBMSSessionStartResponseIEs
}

func (self * MBMSSessionStartResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MBMSSessionStartResponseIEs, order_MBMSSessionStartResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSSessionStartResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MBMSSessionStartResponseIEs, order_MBMSSessionStartResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MBMSSessionStartFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MBMSSessionStartFailure-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MBMSSessionStartFailureIEs
}

func (self * MBMSSessionStartFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MBMSSessionStartFailureIEs, order_MBMSSessionStartFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSSessionStartFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MBMSSessionStartFailureIEs, order_MBMSSessionStartFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MBMSSessionStopRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MBMSSessionStopRequest-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MBMSSessionStopRequestIEs
}

func (self * MBMSSessionStopRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MBMSSessionStopRequestIEs, order_MBMSSessionStopRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSSessionStopRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MBMSSessionStopRequestIEs, order_MBMSSessionStopRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MBMSSessionStopResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MBMSSessionStopResponse-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MBMSSessionStopResponseIEs
}

func (self * MBMSSessionStopResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MBMSSessionStopResponseIEs, order_MBMSSessionStopResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSSessionStopResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MBMSSessionStopResponseIEs, order_MBMSSessionStopResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MBMSSessionUpdateRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MBMSSessionUpdateRequest-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MBMSSessionUpdateRequestIEs
}

func (self * MBMSSessionUpdateRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MBMSSessionUpdateRequestIEs, order_MBMSSessionUpdateRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSSessionUpdateRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MBMSSessionUpdateRequestIEs, order_MBMSSessionUpdateRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MBMSSessionUpdateResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MBMSSessionUpdateResponse-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MBMSSessionUpdateResponseIEs
}

func (self * MBMSSessionUpdateResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MBMSSessionUpdateResponseIEs, order_MBMSSessionUpdateResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSSessionUpdateResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MBMSSessionUpdateResponseIEs, order_MBMSSessionUpdateResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MBMSSessionUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MBMSSessionUpdateFailure-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MBMSSessionUpdateFailureIEs
}

func (self * MBMSSessionUpdateFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MBMSSessionUpdateFailureIEs, order_MBMSSessionUpdateFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSSessionUpdateFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_MBMSSessionUpdateFailureIEs, order_MBMSSessionUpdateFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

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
    //coptions := []string{"m3-Interface","partOfM3-Interface"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in ResetType\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.M3Interface = &ResetAll{}//cho6
        self.M3Interface.Unpack(stream)
    } else if choice == 1 { //ch2
        self.PartOfM3Interface = &MBMSServiceassociatedLogicalM3ConnectionListRes{}//cho6
        self.PartOfM3Interface.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ResetType) Pack(stream *Stream) {
    if self.M3Interface != nil {
        stream.set_choice(0, 1, 1, 2)
        self.M3Interface.Pack(stream)//2
    } else if self.PartOfM3Interface != nil {
        stream.set_choice(1, 1, 1, 2)
        self.PartOfM3Interface.Pack(stream)//2
    }

}
type ResetType struct { //[{'type': 'ResetAll', 'name': 'm3-Interface'}, {'type': 'MBMS-Service-associatedLogicalM3-ConnectionListRes', 'name': 'partOfM3-Interface'}, None]
    M3Interface *ResetAll
    PartOfM3Interface *MBMSServiceassociatedLogicalM3ConnectionListRes
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
func (self *MBMSServiceassociatedLogicalM3ConnectionListRes) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]MBMSServiceassociatedLogicalM3ConnectionItemRes, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *MBMSServiceassociatedLogicalM3ConnectionListRes) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    //for item in table_MBMSServiceassociatedLogicalM3ConnectionItemRes:
    val := ProtocolIESingleContainer{table_MBMSServiceassociatedLogicalM3ConnectionItemRes, order_MBMSServiceassociatedLogicalM3ConnectionItemRes}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type MBMSServiceassociatedLogicalM3ConnectionListRes struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['MBMS-Service-associatedLogicalM3-ConnectionItemRes']}, 'size': [(1, 'maxNrOfIndividualM3ConnectionsToReset')]}
    Items []MBMSServiceassociatedLogicalM3ConnectionItemRes
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

func (self *MBMSServiceassociatedLogicalM3ConnectionListResAck) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]MBMSServiceassociatedLogicalM3ConnectionItemResAck, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *MBMSServiceassociatedLogicalM3ConnectionListResAck) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    //for item in table_MBMSServiceassociatedLogicalM3ConnectionItemResAck:
    val := ProtocolIESingleContainer{table_MBMSServiceassociatedLogicalM3ConnectionItemResAck, order_MBMSServiceassociatedLogicalM3ConnectionItemResAck}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type MBMSServiceassociatedLogicalM3ConnectionListResAck struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Single-Container', 'actual-parameters': ['MBMS-Service-associatedLogicalM3-ConnectionItemResAck']}, 'size': [(1, 'maxNrOfIndividualM3ConnectionsToReset')]}
    Items []MBMSServiceassociatedLogicalM3ConnectionItemResAck
}

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

type M3SetupRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['M3SetupRequestIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs M3SetupRequestIEs
}

func (self * M3SetupRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_M3SetupRequestIEs, order_M3SetupRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * M3SetupRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_M3SetupRequestIEs, order_M3SetupRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *MBMSServiceAreaListItem) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65536)
    _size += 1
    self.Items = make([]MBMSServiceArea1, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MBMSServiceAreaListItem) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65536)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MBMSServiceAreaListItem struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MBMSServiceArea1'}, 'size': [(1, 'maxnoofMBMSServiceAreaIdentitiesPerMCE')]}
    Items []MBMSServiceArea1
}

type M3SetupResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['M3SetupResponseIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs M3SetupResponseIEs
}

func (self * M3SetupResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_M3SetupResponseIEs, order_M3SetupResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * M3SetupResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_M3SetupResponseIEs, order_M3SetupResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type M3SetupFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['M3SetupFailureIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs M3SetupFailureIEs
}

func (self * M3SetupFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_M3SetupFailureIEs, order_M3SetupFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * M3SetupFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_M3SetupFailureIEs, order_M3SetupFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MCEConfigurationUpdate struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MCEConfigurationUpdateIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MCEConfigurationUpdateIEs
}

func (self * MCEConfigurationUpdate) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MCEConfigurationUpdateIEs, order_MCEConfigurationUpdateIEs} // p3
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
    ProtocolIEs := &ProtocolIEContainer {table_MCEConfigurationUpdateIEs, order_MCEConfigurationUpdateIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MCEConfigurationUpdateAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MCEConfigurationUpdateAcknowledgeIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MCEConfigurationUpdateAcknowledgeIEs
}

func (self * MCEConfigurationUpdateAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MCEConfigurationUpdateAcknowledgeIEs, order_MCEConfigurationUpdateAcknowledgeIEs} // p3
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
    ProtocolIEs := &ProtocolIEContainer {table_MCEConfigurationUpdateAcknowledgeIEs, order_MCEConfigurationUpdateAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MCEConfigurationUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['MCEConfigurationUpdateFailureIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs MCEConfigurationUpdateFailureIEs
}

func (self * MCEConfigurationUpdateFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_MCEConfigurationUpdateFailureIEs, order_MCEConfigurationUpdateFailureIEs} // p3
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
    ProtocolIEs := &ProtocolIEContainer {table_MCEConfigurationUpdateFailureIEs, order_MCEConfigurationUpdateFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type AbsoluteTimeofMBMSData struct {
  Len int
  Value HexBytes
}
func (self *AbsoluteTimeofMBMSData) Unpack(st *Stream){
    self.Value = st.parsef_BitString(64, 64)
}
func (self *AbsoluteTimeofMBMSData) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 64)
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
    CauseRadioNetworkunknown_or_already_allocated_MME_MBMS_M3AP_ID = 0
    CauseRadioNetworkunknown_or_already_allocated_MCE_MBMS_M3AP_ID = 1
    CauseRadioNetworkunknown_or_inconsistent_pair_of_MBMS_M3AP_IDs = 2
    CauseRadioNetworkradio_resources_not_available = 3
    CauseRadioNetworkinvalid_QoS_combination = 4
    CauseRadioNetworkinteraction_with_other_procedure = 5
    CauseRadioNetworknot_supported_QCI_value = 6
    CauseRadioNetworkunspecified = 7

    /* Extensions */
    CauseRadioNetworkuninvolved_MCE = 8
)
func (self *CauseRadioNetwork) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 8, 1)
}
func (self *CauseRadioNetwork) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 8, 1)
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
type ExtendedMCEID struct {
  Value HexBytes
}
func (self *ExtendedMCEID) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(1)
}
func (self *ExtendedMCEID) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 1)
}
type GlobalMCEID struct { // [{'type': 'PLMN-Identity', 'name': 'pLMN-Identity'}, {'type': 'MCE-ID', 'name': 'mCE-ID'}, {'type': 'ExtendedMCE-ID', 'name': 'extendedMCE-ID', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GlobalMCE-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNIdentity PLMNIdentity
    MCEID MCEID
    ExtendedMCEID *ExtendedMCEID
    IEExtensions *GlobalMCEIDExtIEs
}

func (self * GlobalMCEID) Unpack(stream *Stream) {
    extendedMCEID_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.PLMNIdentity.Unpack(stream)// p8
    self.MCEID.Unpack(stream)// p8
    if (extendedMCEID_flag & _flags) == extendedMCEID_flag { //cond2
        self.ExtendedMCEID = &ExtendedMCEID{}//7{'type': 'ExtendedMCE-ID', 'name': 'extendedMCE-ID', 'optional': True}
        self.ExtendedMCEID.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GlobalMCEIDExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GlobalMCE-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GlobalMCEIDExtIEs, order_GlobalMCEIDExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalMCEID) Pack(stream *Stream) {
    const extendedMCEID_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.MCEID.Pack(stream)
    if self.ExtendedMCEID != nil { 
        _flags |= extendedMCEID_flag
        self.ExtendedMCEID.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GlobalMCEIDExtIEs, order_GlobalMCEIDExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

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
    st.parse_ext()
    _len := st.parse_olen(5)+4
    if _len < 4 || _len > 16 {
        //fmt.Println ("Invalid len in IPAddress")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *IPAddress) Pack(st *Stream) {
    _eflag := 0
    if len(self.Value) > 16:
       var _eflag int = 1
    st.format_ext(_eflag)
    if len(self.Value) < 4 || len(self.Value) > 16 {
        log.Error ("Invalid len in IPAddress")
        return
}
    st.format_olen((len(self.Value))-4, 5)
    st.formatf_OctString(self.Value, 0)
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

type MBMSERABQoSParameters struct { // [{'type': 'QCI', 'name': 'qCI'}, {'type': 'GBR-QosInformation', 'name': 'gbrQosInformation', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-E-RAB-QoS-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    QCI QCI
    GbrQosInformation *GBRQosInformation
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
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MBMSERABQoSParametersExtIEs, order_MBMSERABQoSParametersExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type MBMSServiceassociatedLogicalM3ConnectionItem struct { // [{'type': 'MME-MBMS-M3AP-ID', 'name': 'mME-MBMS-M3AP-ID', 'optional': True}, {'type': 'MCE-MBMS-M3AP-ID', 'name': 'mCE-MBMS-M3AP-ID', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-Service-associatedLogicalM3-ConnectionItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    MMEMBMSM3APID *MMEMBMSM3APID
    MCEMBMSM3APID *MCEMBMSM3APID
    IEExtensions *MBMSServiceassociatedLogicalM3ConnectionItemExtIEs
}

func (self * MBMSServiceassociatedLogicalM3ConnectionItem) Unpack(stream *Stream) {
    mMEMBMSM3APID_flag := 0x00000002
    mCEMBMSM3APID_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (mMEMBMSM3APID_flag & _flags) == mMEMBMSM3APID_flag { //cond2
        self.MMEMBMSM3APID = &MMEMBMSM3APID{}//7{'type': 'MME-MBMS-M3AP-ID', 'name': 'mME-MBMS-M3AP-ID', 'optional': True}
        self.MMEMBMSM3APID.Unpack(stream)// p8
    }
    if (mCEMBMSM3APID_flag & _flags) == mCEMBMSM3APID_flag { //cond2
        self.MCEMBMSM3APID = &MCEMBMSM3APID{}//7{'type': 'MCE-MBMS-M3AP-ID', 'name': 'mCE-MBMS-M3AP-ID', 'optional': True}
        self.MCEMBMSM3APID.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MBMSServiceassociatedLogicalM3ConnectionItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MBMS-Service-associatedLogicalM3-ConnectionItemExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MBMSServiceassociatedLogicalM3ConnectionItemExtIEs, order_MBMSServiceassociatedLogicalM3ConnectionItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MBMSServiceassociatedLogicalM3ConnectionItem) Pack(stream *Stream) {
    const mMEMBMSM3APID_flag uint = 0x00000002
    const mCEMBMSM3APID_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.MMEMBMSM3APID != nil { 
        _flags |= mMEMBMSM3APID_flag
        self.MMEMBMSM3APID.Pack(stream)
    }//end of optional
    if self.MCEMBMSM3APID != nil { 
        _flags |= mCEMBMSM3APID_flag
        self.MCEMBMSM3APID.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MBMSServiceassociatedLogicalM3ConnectionItemExtIEs, order_MBMSServiceassociatedLogicalM3ConnectionItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type MBMSServiceArea1 struct {
  Value HexBytes
}
func (self *MBMSServiceArea1) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(2)
}
func (self *MBMSServiceArea1) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 2)
}
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
type MBMSSessionDuration struct {
  Value HexBytes
}
func (self *MBMSSessionDuration) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(3)
}
func (self *MBMSSessionDuration) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 3)
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
type MCEMBMSM3APID struct {
  Value uint64
}
func (self *MCEMBMSM3APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * MCEMBMSM3APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type MCEID struct {
  Value HexBytes
}
func (self *MCEID) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(2)
}
func (self *MCEID) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 2)
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
type MinimumTimeToMBMSDataTransfer struct {
  Value HexBytes
}
func (self *MinimumTimeToMBMSDataTransfer) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(1)
}
func (self *MinimumTimeToMBMSDataTransfer) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 1)
}
type MMEMBMSM3APID struct {
  Value uint64
}
func (self *MMEMBMSM3APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * MMEMBMSM3APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
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
type PLMNIdentity struct {
  Value HexBytes
}
func (self *PLMNIdentity) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(3)
}
func (self *PLMNIdentity) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 3)
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
type Reestablishment struct {
  Value int
}
const (
    ReestablishmenttRue = 0

    /* Extensions */
)
func (self *Reestablishment) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *Reestablishment) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
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
type TMGI struct { // [{'type': 'PLMN-Identity', 'name': 'pLMNidentity'}, {'type': 'OCTET STRING', 'size': [3], 'name': 'serviceID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TMGI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PLMNidentity PLMNIdentity
    ServiceID OCTETSTRING
    IEExtensions *TMGIExtIEs
}

func (self * TMGI) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
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
    return
}

func (self * TMGI) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
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
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type TNLInformation struct { // [{'type': 'IPAddress', 'name': 'iPMCAddress'}, {'type': 'IPAddress', 'name': 'iPSourceAddress'}, {'type': 'GTP-TEID', 'name': 'gTP-DLTEID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TNL-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    IPMCAddress IPAddress
    IPSourceAddress IPAddress
    GTPDLTEID GTPTEID
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
    self.GTPDLTEID.Unpack(stream)// p8
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
    self.GTPDLTEID.Pack(stream)
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
    _size := data.(M3APPROTOCOLIES_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['M3AP-PROTOCOL-IES']}
    Items map[int]*M3APPROTOCOLIES
    order []int
}

type ProtocolIESingleContainer struct{ //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['M3AP-PROTOCOL-IES']}
   Item map[int]*M3APPROTOCOLIES //UserType
   order []int
}
func (self *ProtocolIESingleContainer) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['M3AP-PROTOCOL-IES']}
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

type ProtocolIEField struct { // [{'type': 'M3AP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'M3AP-PROTOCOL-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'M3AP-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id M3APPROTOCOLIESid
    Criticality M3APPROTOCOLIEScriticality
    Value M3APPROTOCOLIESValue
}

func (self * ProtocolIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'M3AP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := M3APPROTOCOLIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(M3APPROTOCOLIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'M3AP-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * M3APPROTOCOLIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'M3AP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (M3APPROTOCOLIESid)(self.ID)
    if out.(M3APPROTOCOLIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(M3APPROTOCOLIES_IF).PackOT(stream, key)
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
    _size := data.(M3APPROTOCOLIESPAIR_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainerPair struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-FieldPair', 'actual-parameters': ['IEsSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['M3AP-PROTOCOL-IES-PAIR']}
    Items map[int]*M3APPROTOCOLIESPAIR
    order []int
}

type ProtocolIEFieldPair struct { // [{'type': 'M3AP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'M3AP-PROTOCOL-IES-PAIR.&firstCriticality', 'table': ['IEsSetParam', ['id']], 'name': 'firstCriticality'}, {'type': 'M3AP-PROTOCOL-IES-PAIR.&FirstValue', 'table': ['IEsSetParam', ['id']], 'name': 'firstValue'}, {'type': 'M3AP-PROTOCOL-IES-PAIR.&secondCriticality', 'table': ['IEsSetParam', ['id']], 'name': 'secondCriticality'}, {'type': 'M3AP-PROTOCOL-IES-PAIR.&SecondValue', 'table': ['IEsSetParam', ['id']], 'name': 'secondValue'}]
    Id M3APPROTOCOLIESPAIRid
    FirstCriticality M3APPROTOCOLIESPAIRfirstCriticality
    FirstValue M3APPROTOCOLIESPAIRFirstValue
    SecondCriticality M3APPROTOCOLIESPAIRsecondCriticality
    SecondValue M3APPROTOCOLIESPAIRSecondValue
}

func (self * ProtocolIEFieldPair) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'M3AP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := M3APPROTOCOLIESPAIR{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.FirstCriticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(M3APPROTOCOLIESPAIR_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'M3AP-PROTOCOL-IES-PAIR.&FirstValue', 'table': ['IEsSetParam', ['id']], 'name': 'firstValue'}
    self.SecondCriticality.Unpack(stream)//p9
    out.(M3APPROTOCOLIESPAIR_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'M3AP-PROTOCOL-IES-PAIR.&SecondValue', 'table': ['IEsSetParam', ['id']], 'name': 'secondValue'}
    stream.set_location(location, _len)
    return
}

func (self * M3APPROTOCOLIESPAIR) Pack(stream *Stream, out interface{}) {
    //table {'type': 'M3AP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (M3APPROTOCOLIESPAIRid)(self.ID)
    if out.(M3APPROTOCOLIESPAIR_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.FIRSTCRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(M3APPROTOCOLIESPAIR_IF).PackOT(stream, key)
    self.SECONDCRITICALITY.Pack(stream)
    out.(M3APPROTOCOLIESPAIR_IF).PackOT(stream, key)
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


type ProtocolIEContainerList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['IEsSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IEsSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'M3AP-PROTOCOL-IES']}
    Items map[int]*M3APPROTOCOLIES
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


type ProtocolIEContainerPairList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-ContainerPair', 'actual-parameters': ['IEsSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IEsSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'M3AP-PROTOCOL-IES-PAIR']}
    Items map[int]*M3APPROTOCOLIESPAIR
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
    _size := data.(M3APPROTOCOLEXTENSION_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolExtensionContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolExtensionField', 'actual-parameters': ['ExtensionSetParam']}, 'size': [(1, 'maxProtocolExtensions')], 'parameters': ['ExtensionSetParam'], 'param-types': ['M3AP-PROTOCOL-EXTENSION']}
    Items map[int]*M3APPROTOCOLEXTENSION
    order []int
}

type ProtocolExtensionField struct { // [{'type': 'M3AP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}, {'type': 'M3AP-PROTOCOL-EXTENSION.&criticality', 'table': ['ExtensionSetParam', ['id']], 'name': 'criticality'}, {'type': 'M3AP-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}]
    Id M3APPROTOCOLEXTENSIONid
    Criticality M3APPROTOCOLEXTENSIONcriticality
    ExtensionValue M3APPROTOCOLEXTENSIONExtension
}

func (self * ProtocolExtensionField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'M3AP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := M3APPROTOCOLEXTENSION{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(M3APPROTOCOLEXTENSION_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'M3AP-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}
    stream.set_location(location, _len)
    return
}

func (self * M3APPROTOCOLEXTENSION) Pack(stream *Stream, out interface{}) {
    //table {'type': 'M3AP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    key := (M3APPROTOCOLEXTENSIONid)(self.ID)
    if out.(M3APPROTOCOLEXTENSION_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(M3APPROTOCOLEXTENSION_IF).PackOT(stream, key)
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
    _size := data.(M3APPRIVATEIES_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type PrivateIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PrivateIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(1, 'maxPrivateIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['M3AP-PRIVATE-IES']}
    Items map[int]*M3APPRIVATEIES
    order []int
}

type PrivateIEField struct { // [{'type': 'M3AP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'M3AP-PRIVATE-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'M3AP-PRIVATE-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id M3APPRIVATEIESid
    Criticality M3APPRIVATEIEScriticality
    Value M3APPRIVATEIESValue
}

func (self * PrivateIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'M3AP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := M3APPRIVATEIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(M3APPRIVATEIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'M3AP-PRIVATE-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * M3APPRIVATEIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'M3AP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (M3APPRIVATEIESid)(self.ID)
    if out.(M3APPRIVATEIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(M3APPRIVATEIES_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type M3APELEMENTARYPROCEDURE struct {//CLASS {'members': [{'type': 'OpenType', 'name': '&InitiatingMessage'}, {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, {'type': 'ProcedureCode', 'name': '&procedureCode'}, {'type': 'Criticality', 'name': '&criticality'}], 'with-members': [{'with-order': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME', 'PROCEDURE CODE', 'CRITICALITY'], 'with-type': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME']}], 'alt-type': {'INITIATING MESSAGE': {'type': 'OpenType', 'name': '&InitiatingMessage'}, 'SUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, 'UNSUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, 'PROCEDURE CODE': {'type': 'ProcedureCode', 'name': '&procedureCode'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}}, 'id-type': {'type': 'ProcedureCode', 'name': '&procedureCode'}}
    INITIATINGMESSAGE interface{}
    SUCCESSFULOUTCOME interface{}
    UNSUCCESSFULOUTCOME interface{}
    PROCEDURECODE ProcedureCode
    CRITICALITY Criticality
}
type M3APELEMENTARYPROCEDUREInitiatingMessage interface{}
type M3APELEMENTARYPROCEDURESuccessfulOutcome interface{}
type M3APELEMENTARYPROCEDUREUnsuccessfulOutcome interface{}
type M3APELEMENTARYPROCEDUREprocedureCode ProcedureCode
func (self *M3APELEMENTARYPROCEDUREprocedureCode) Pack(st *Stream) {
    ieID := ProcedureCode(*self)
    ieID.Pack(st)
}
func (self *M3APELEMENTARYPROCEDUREprocedureCode) Unpack(st *Stream) {
    val := ProcedureCode{}
    val.Unpack(st)
    *self = M3APELEMENTARYPROCEDUREprocedureCode(val)
}
type M3APELEMENTARYPROCEDUREcriticality Criticality
func (self *M3APELEMENTARYPROCEDUREcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M3APELEMENTARYPROCEDUREcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M3APELEMENTARYPROCEDUREcriticality(val)
}

type M3APELEMENTARYPROCEDURE_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type M3APPROTOCOLIES struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type M3APPROTOCOLIESid ProtocolIEID
func (self *M3APPROTOCOLIESid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *M3APPROTOCOLIESid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = M3APPROTOCOLIESid(val)
}
type M3APPROTOCOLIEScriticality Criticality
func (self *M3APPROTOCOLIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M3APPROTOCOLIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M3APPROTOCOLIEScriticality(val)
}
type M3APPROTOCOLIESValue interface{}
type M3APPROTOCOLIESpresence Presence
func (self *M3APPROTOCOLIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *M3APPROTOCOLIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = M3APPROTOCOLIESpresence(val)
}

type M3APPROTOCOLIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type M3APPROTOCOLIESPAIR struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&firstCriticality'}, {'type': 'OpenType', 'name': '&FirstValue'}, {'type': 'Criticality', 'name': '&secondCriticality'}, {'type': 'OpenType', 'name': '&SecondValue'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'FIRST CRITICALITY', 'FIRST TYPE', 'SECOND CRITICALITY', 'SECOND TYPE', 'PRESENCE'], 'with-type': ['FIRST TYPE', 'SECOND TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'FIRST CRITICALITY': {'type': 'Criticality', 'name': '&firstCriticality'}, 'FIRST TYPE': {'type': 'OpenType', 'name': '&FirstValue'}, 'SECOND CRITICALITY': {'type': 'Criticality', 'name': '&secondCriticality'}, 'SECOND TYPE': {'type': 'OpenType', 'name': '&SecondValue'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    FIRSTCRITICALITY Criticality
    FIRSTTYPE interface{}
    SECONDCRITICALITY Criticality
    SECONDTYPE interface{}
    PRESENCE Presence
}
type M3APPROTOCOLIESPAIRid ProtocolIEID
func (self *M3APPROTOCOLIESPAIRid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *M3APPROTOCOLIESPAIRid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = M3APPROTOCOLIESPAIRid(val)
}
type M3APPROTOCOLIESPAIRfirstCriticality Criticality
func (self *M3APPROTOCOLIESPAIRfirstCriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M3APPROTOCOLIESPAIRfirstCriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M3APPROTOCOLIESPAIRfirstCriticality(val)
}
type M3APPROTOCOLIESPAIRFirstValue interface{}
type M3APPROTOCOLIESPAIRsecondCriticality Criticality
func (self *M3APPROTOCOLIESPAIRsecondCriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M3APPROTOCOLIESPAIRsecondCriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M3APPROTOCOLIESPAIRsecondCriticality(val)
}
type M3APPROTOCOLIESPAIRSecondValue interface{}
type M3APPROTOCOLIESPAIRpresence Presence
func (self *M3APPROTOCOLIESPAIRpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *M3APPROTOCOLIESPAIRpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = M3APPROTOCOLIESPAIRpresence(val)
}

type M3APPROTOCOLIESPAIR_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type M3APPROTOCOLEXTENSION struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Extension'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'EXTENSION', 'PRESENCE'], 'with-type': ['EXTENSION']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    EXTENSION interface{}
    PRESENCE Presence
}
type M3APPROTOCOLEXTENSIONid ProtocolIEID
func (self *M3APPROTOCOLEXTENSIONid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *M3APPROTOCOLEXTENSIONid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = M3APPROTOCOLEXTENSIONid(val)
}
type M3APPROTOCOLEXTENSIONcriticality Criticality
func (self *M3APPROTOCOLEXTENSIONcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M3APPROTOCOLEXTENSIONcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M3APPROTOCOLEXTENSIONcriticality(val)
}
type M3APPROTOCOLEXTENSIONExtension interface{}
type M3APPROTOCOLEXTENSIONpresence Presence
func (self *M3APPROTOCOLEXTENSIONpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *M3APPROTOCOLEXTENSIONpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = M3APPROTOCOLEXTENSIONpresence(val)
}

type M3APPROTOCOLEXTENSION_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type M3APPRIVATEIES struct {//CLASS {'members': [{'type': 'PrivateIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'PrivateIE-ID', 'name': '&id'}}
    ID PrivateIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type M3APPRIVATEIESid PrivateIEID
func (self *M3APPRIVATEIESid) Pack(st *Stream) {
    ieID := PrivateIEID(*self)
    ieID.Pack(st)
}
func (self *M3APPRIVATEIESid) Unpack(st *Stream) {
    val := PrivateIEID{}
    val.Unpack(st)
    *self = M3APPRIVATEIESid(val)
}
type M3APPRIVATEIEScriticality Criticality
func (self *M3APPRIVATEIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *M3APPRIVATEIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = M3APPRIVATEIEScriticality(val)
}
type M3APPRIVATEIESValue interface{}
type M3APPRIVATEIESpresence Presence
func (self *M3APPRIVATEIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *M3APPRIVATEIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = M3APPRIVATEIESpresence(val)
}

type M3APPRIVATEIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
//class M3APELEMENTARYPROCEDURES: #OBJSET1 {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'members': [{}, {}, None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_M3APELEMENTARYPROCEDURES = make(map[M3APELEMENTARYPROCEDUREprocedureCode]*M3APELEMENTARYPROCEDURE)


//class M3APELEMENTARYPROCEDURESCLASS1: #OBJSET1 {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_M3APELEMENTARYPROCEDURESCLASS1 = make(map[M3APELEMENTARYPROCEDUREprocedureCode]*M3APELEMENTARYPROCEDURE)


//class M3APELEMENTARYPROCEDURESCLASS2: #OBJSET1 {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_M3APELEMENTARYPROCEDURESCLASS2 = make(map[M3APELEMENTARYPROCEDUREprocedureCode]*M3APELEMENTARYPROCEDURE)


type MBMSSessionStartRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MME-MBMS-M3AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'MME-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-TMGI', 'CRITICALITY': 'reject', 'TYPE': 'TMGI', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Session-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Session-ID', 'PRESENCE': 'optional'}, {'ID': 'id-MBMS-E-RAB-QoS-Parameters', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-E-RAB-QoS-Parameters', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Session-Duration', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Session-Duration', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Service-Area', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Service-Area', 'PRESENCE': 'mandatory'}, {'ID': 'id-MinimumTimeToMBMSDataTransfer', 'CRITICALITY': 'reject', 'TYPE': 'MinimumTimeToMBMSDataTransfer', 'PRESENCE': 'mandatory'}, {'ID': 'id-TNL-Information', 'CRITICALITY': 'reject', 'TYPE': 'TNL-Information', 'PRESENCE': 'mandatory'}, {'ID': 'id-Time-ofMBMS-DataTransfer', 'CRITICALITY': 'ignore', 'TYPE': 'Absolute-Time-ofMBMS-Data', 'PRESENCE': 'optional'}, {'ID': 'id-Reestablishment', 'CRITICALITY': 'ignore', 'TYPE': 'Reestablishment', 'PRESENCE': 'optional'}, {'ID': 'id-Alternative-TNL-Information', 'CRITICALITY': 'ignore', 'TYPE': 'TNL-Information', 'PRESENCE': 'optional'}, {'ID': 'id-MBMS-Cell-List', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Cell-List', 'PRESENCE': 'optional'}, None]}
   MMEMBMSM3APID  MMEMBMSM3APID
   TMGI  TMGI
   MBMSSessionID  *MBMSSessionID
   MBMSERABQoSParameters  MBMSERABQoSParameters
   MBMSSessionDuration  MBMSSessionDuration
   MBMSServiceArea  MBMSServiceArea
   MinimumTimeToMBMSDataTransfer  MinimumTimeToMBMSDataTransfer
   TNLInformation  TNLInformation
   TimeofMBMSDataTransfer  *AbsoluteTimeofMBMSData
   Reestablishment  *Reestablishment
   AlternativeTNLInformation  *TNLInformation
   MBMSCellList  *MBMSCellList
   list []interface{}
}
func (self *MBMSSessionStartRequestIEs)createOT() interface{}{
    return nil
}
var table_MBMSSessionStartRequestIEs = make(map[int]*M3APPROTOCOLIES)

var order_MBMSSessionStartRequestIEs = make([]int, 12)

func (self *MBMSSessionStartRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.MMEMBMSM3APID
   count +=1 //self.TMGI
   if self.MBMSSessionID != nil { count += 1 }
   count +=1 //self.MBMSERABQoSParameters
   count +=1 //self.MBMSSessionDuration
   count +=1 //self.MBMSServiceArea
   count +=1 //self.MinimumTimeToMBMSDataTransfer
   count +=1 //self.TNLInformation
   if self.TimeofMBMSDataTransfer != nil { count += 1 }
   if self.Reestablishment != nil { count += 1 }
   if self.AlternativeTNLInformation != nil { count += 1 }
   if self.MBMSCellList != nil { count += 1 }
   return count//ObjSet
}
func (self *MBMSSessionStartRequestIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        return true //self.MMEMBMSM3APID
      case 2: //TMGI
        return true //self.TMGI
      case 3: //MBMSSessionID
        if self.MBMSSessionID != nil { return true }
      case 4: //MBMSERABQoSParameters
        return true //self.MBMSERABQoSParameters
      case 5: //MBMSSessionDuration
        return true //self.MBMSSessionDuration
      case 6: //MBMSServiceArea
        return true //self.MBMSServiceArea
      case 16: //MinimumTimeToMBMSDataTransfer
        return true //self.MinimumTimeToMBMSDataTransfer
      case 7: //TNLInformation
        return true //self.TNLInformation
      case 21: //TimeofMBMSDataTransfer
        if self.TimeofMBMSDataTransfer != nil { return true }
      case 23: //Reestablishment
        if self.Reestablishment != nil { return true }
      case 24: //AlternativeTNLInformation
        if self.AlternativeTNLInformation != nil { return true }
      case 25: //MBMSCellList
        if self.MBMSCellList != nil { return true }
   }
   return false//ObjSet
}
func (self *MBMSSessionStartRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MMEMBMSM3APID)
      case 2: //TMGI
        self.TMGI.Unpack(st)
        self.list = append(self.list, &self.TMGI)
      case 3: //MBMSSessionID
        self.MBMSSessionID = &MBMSSessionID{}
        self.MBMSSessionID.Unpack(st)
        self.list = append(self.list, self.MBMSSessionID)
      case 4: //MBMSERABQoSParameters
        self.MBMSERABQoSParameters.Unpack(st)
        self.list = append(self.list, &self.MBMSERABQoSParameters)
      case 5: //MBMSSessionDuration
        self.MBMSSessionDuration.Unpack(st)
        self.list = append(self.list, &self.MBMSSessionDuration)
      case 6: //MBMSServiceArea
        self.MBMSServiceArea.Unpack(st)
        self.list = append(self.list, &self.MBMSServiceArea)
      case 16: //MinimumTimeToMBMSDataTransfer
        self.MinimumTimeToMBMSDataTransfer.Unpack(st)
        self.list = append(self.list, &self.MinimumTimeToMBMSDataTransfer)
      case 7: //TNLInformation
        self.TNLInformation.Unpack(st)
        self.list = append(self.list, &self.TNLInformation)
      case 21: //TimeofMBMSDataTransfer
        self.TimeofMBMSDataTransfer = &AbsoluteTimeofMBMSData{}
        self.TimeofMBMSDataTransfer.Unpack(st)
        self.list = append(self.list, self.TimeofMBMSDataTransfer)
      case 23: //Reestablishment
        self.Reestablishment = &Reestablishment{}
        self.Reestablishment.Unpack(st)
        self.list = append(self.list, self.Reestablishment)
      case 24: //AlternativeTNLInformation
        self.AlternativeTNLInformation = &TNLInformation{}
        self.AlternativeTNLInformation.Unpack(st)
        self.list = append(self.list, self.AlternativeTNLInformation)
      case 25: //MBMSCellList
        self.MBMSCellList = &MBMSCellList{}
        self.MBMSCellList.Unpack(st)
        self.list = append(self.list, self.MBMSCellList)
   }
}
func (self *MBMSSessionStartRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Pack(st)
      case 2: //TMGI
        self.TMGI.Pack(st)
      case 3: //MBMSSessionID
        if self.MBMSSessionID != nil {self.MBMSSessionID.Pack(st)}
      case 4: //MBMSERABQoSParameters
        self.MBMSERABQoSParameters.Pack(st)
      case 5: //MBMSSessionDuration
        self.MBMSSessionDuration.Pack(st)
      case 6: //MBMSServiceArea
        self.MBMSServiceArea.Pack(st)
      case 16: //MinimumTimeToMBMSDataTransfer
        self.MinimumTimeToMBMSDataTransfer.Pack(st)
      case 7: //TNLInformation
        self.TNLInformation.Pack(st)
      case 21: //TimeofMBMSDataTransfer
        if self.TimeofMBMSDataTransfer != nil {self.TimeofMBMSDataTransfer.Pack(st)}
      case 23: //Reestablishment
        if self.Reestablishment != nil {self.Reestablishment.Pack(st)}
      case 24: //AlternativeTNLInformation
        if self.AlternativeTNLInformation != nil {self.AlternativeTNLInformation.Pack(st)}
      case 25: //MBMSCellList
        if self.MBMSCellList != nil {self.MBMSCellList.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBMSSessionStartRequestIEs[0] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMMEMBMSM3APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MMEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartRequestIEs[0] = 0
table_MBMSSessionStartRequestIEs[2] = &M3APPROTOCOLIES{ID:ProtocolIEID{idTMGI}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TMGI{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartRequestIEs[1] = 2
table_MBMSSessionStartRequestIEs[3] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSSessionID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSSessionID{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionStartRequestIEs[2] = 3
table_MBMSSessionStartRequestIEs[4] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSERABQoSParameters}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSERABQoSParameters{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartRequestIEs[3] = 4
table_MBMSSessionStartRequestIEs[5] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSSessionDuration}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSSessionDuration{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartRequestIEs[4] = 5
table_MBMSSessionStartRequestIEs[6] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceArea}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSServiceArea{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartRequestIEs[5] = 6
table_MBMSSessionStartRequestIEs[16] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMinimumTimeToMBMSDataTransfer}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MinimumTimeToMBMSDataTransfer{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartRequestIEs[6] = 16
table_MBMSSessionStartRequestIEs[7] = &M3APPROTOCOLIES{ID:ProtocolIEID{idTNLInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TNLInformation{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartRequestIEs[7] = 7
table_MBMSSessionStartRequestIEs[21] = &M3APPROTOCOLIES{ID:ProtocolIEID{idTimeofMBMSDataTransfer}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&AbsoluteTimeofMBMSData{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionStartRequestIEs[8] = 21
table_MBMSSessionStartRequestIEs[23] = &M3APPROTOCOLIES{ID:ProtocolIEID{idReestablishment}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Reestablishment{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionStartRequestIEs[9] = 23
table_MBMSSessionStartRequestIEs[24] = &M3APPROTOCOLIES{ID:ProtocolIEID{idAlternativeTNLInformation}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TNLInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionStartRequestIEs[10] = 24
table_MBMSSessionStartRequestIEs[25] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSCellList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSCellList{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionStartRequestIEs[11] = 25
   }

type MBMSSessionStartResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MME-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MME-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MCE-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MMEMBMSM3APID  MMEMBMSM3APID
   MCEMBMSM3APID  MCEMBMSM3APID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MBMSSessionStartResponseIEs)createOT() interface{}{
    return nil
}
var table_MBMSSessionStartResponseIEs = make(map[int]*M3APPROTOCOLIES)

var order_MBMSSessionStartResponseIEs = make([]int, 3)

func (self *MBMSSessionStartResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.MMEMBMSM3APID
   count +=1 //self.MCEMBMSM3APID
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MBMSSessionStartResponseIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        return true //self.MMEMBMSM3APID
      case 1: //MCEMBMSM3APID
        return true //self.MCEMBMSM3APID
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MBMSSessionStartResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MMEMBMSM3APID)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM3APID)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MBMSSessionStartResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Pack(st)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBMSSessionStartResponseIEs[0] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMMEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MMEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartResponseIEs[0] = 0
table_MBMSSessionStartResponseIEs[1] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartResponseIEs[1] = 1
table_MBMSSessionStartResponseIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionStartResponseIEs[2] = 8
   }

type MBMSSessionStartFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MME-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MME-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MMEMBMSM3APID  MMEMBMSM3APID
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MBMSSessionStartFailureIEs)createOT() interface{}{
    return nil
}
var table_MBMSSessionStartFailureIEs = make(map[int]*M3APPROTOCOLIES)

var order_MBMSSessionStartFailureIEs = make([]int, 3)

func (self *MBMSSessionStartFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.MMEMBMSM3APID
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MBMSSessionStartFailureIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        return true //self.MMEMBMSM3APID
      case 9: //Cause
        return true //self.Cause
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MBMSSessionStartFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MMEMBMSM3APID)
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MBMSSessionStartFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Pack(st)
      case 9: //Cause
        self.Cause.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBMSSessionStartFailureIEs[0] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMMEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MMEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartFailureIEs[0] = 0
table_MBMSSessionStartFailureIEs[9] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStartFailureIEs[1] = 9
table_MBMSSessionStartFailureIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionStartFailureIEs[2] = 8
   }

type MBMSSessionStopRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MME-MBMS-M3AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'MME-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MCE-MBMS-M3AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'MCE-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Time-ofMBMS-DataStop', 'CRITICALITY': 'ignore', 'TYPE': 'Absolute-Time-ofMBMS-Data', 'PRESENCE': 'optional'}, None]}
   MMEMBMSM3APID  MMEMBMSM3APID
   MCEMBMSM3APID  MCEMBMSM3APID
   TimeofMBMSDataStop  *AbsoluteTimeofMBMSData
   list []interface{}
}
func (self *MBMSSessionStopRequestIEs)createOT() interface{}{
    return nil
}
var table_MBMSSessionStopRequestIEs = make(map[int]*M3APPROTOCOLIES)

var order_MBMSSessionStopRequestIEs = make([]int, 3)

func (self *MBMSSessionStopRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.MMEMBMSM3APID
   count +=1 //self.MCEMBMSM3APID
   if self.TimeofMBMSDataStop != nil { count += 1 }
   return count//ObjSet
}
func (self *MBMSSessionStopRequestIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        return true //self.MMEMBMSM3APID
      case 1: //MCEMBMSM3APID
        return true //self.MCEMBMSM3APID
      case 22: //TimeofMBMSDataStop
        if self.TimeofMBMSDataStop != nil { return true }
   }
   return false//ObjSet
}
func (self *MBMSSessionStopRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MMEMBMSM3APID)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM3APID)
      case 22: //TimeofMBMSDataStop
        self.TimeofMBMSDataStop = &AbsoluteTimeofMBMSData{}
        self.TimeofMBMSDataStop.Unpack(st)
        self.list = append(self.list, self.TimeofMBMSDataStop)
   }
}
func (self *MBMSSessionStopRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Pack(st)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Pack(st)
      case 22: //TimeofMBMSDataStop
        if self.TimeofMBMSDataStop != nil {self.TimeofMBMSDataStop.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBMSSessionStopRequestIEs[0] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMMEMBMSM3APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MMEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStopRequestIEs[0] = 0
table_MBMSSessionStopRequestIEs[1] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM3APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStopRequestIEs[1] = 1
table_MBMSSessionStopRequestIEs[22] = &M3APPROTOCOLIES{ID:ProtocolIEID{idTimeofMBMSDataStop}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&AbsoluteTimeofMBMSData{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionStopRequestIEs[2] = 22
   }

type MBMSSessionStopResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MME-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MME-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MCE-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MMEMBMSM3APID  MMEMBMSM3APID
   MCEMBMSM3APID  MCEMBMSM3APID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MBMSSessionStopResponseIEs)createOT() interface{}{
    return nil
}
var table_MBMSSessionStopResponseIEs = make(map[int]*M3APPROTOCOLIES)

var order_MBMSSessionStopResponseIEs = make([]int, 3)

func (self *MBMSSessionStopResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.MMEMBMSM3APID
   count +=1 //self.MCEMBMSM3APID
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MBMSSessionStopResponseIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        return true //self.MMEMBMSM3APID
      case 1: //MCEMBMSM3APID
        return true //self.MCEMBMSM3APID
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MBMSSessionStopResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MMEMBMSM3APID)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM3APID)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MBMSSessionStopResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Pack(st)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBMSSessionStopResponseIEs[0] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMMEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MMEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStopResponseIEs[0] = 0
table_MBMSSessionStopResponseIEs[1] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionStopResponseIEs[1] = 1
table_MBMSSessionStopResponseIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionStopResponseIEs[2] = 8
   }

type MBMSSessionUpdateRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MME-MBMS-M3AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'MME-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MCE-MBMS-M3AP-ID', 'CRITICALITY': 'reject', 'TYPE': 'MCE-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-TMGI', 'CRITICALITY': 'reject', 'TYPE': 'TMGI', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Session-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Session-ID', 'PRESENCE': 'optional'}, {'ID': 'id-MBMS-E-RAB-QoS-Parameters', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-E-RAB-QoS-Parameters', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Session-Duration', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Session-Duration', 'PRESENCE': 'mandatory'}, {'ID': 'id-MBMS-Service-Area', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Service-Area', 'PRESENCE': 'optional'}, {'ID': 'id-MinimumTimeToMBMSDataTransfer', 'CRITICALITY': 'reject', 'TYPE': 'MinimumTimeToMBMSDataTransfer', 'PRESENCE': 'mandatory'}, {'ID': 'id-TNL-Information', 'CRITICALITY': 'ignore', 'TYPE': 'TNL-Information', 'PRESENCE': 'optional'}, {'ID': 'id-Time-ofMBMS-DataTransfer', 'CRITICALITY': 'ignore', 'TYPE': 'Absolute-Time-ofMBMS-Data', 'PRESENCE': 'optional'}, {'ID': 'id-MBMS-Cell-List', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Cell-List', 'PRESENCE': 'optional'}, None]}
   MMEMBMSM3APID  MMEMBMSM3APID
   MCEMBMSM3APID  MCEMBMSM3APID
   TMGI  TMGI
   MBMSSessionID  *MBMSSessionID
   MBMSERABQoSParameters  MBMSERABQoSParameters
   MBMSSessionDuration  MBMSSessionDuration
   MBMSServiceArea  *MBMSServiceArea
   MinimumTimeToMBMSDataTransfer  MinimumTimeToMBMSDataTransfer
   TNLInformation  *TNLInformation
   TimeofMBMSDataTransfer  *AbsoluteTimeofMBMSData
   MBMSCellList  *MBMSCellList
   list []interface{}
}
func (self *MBMSSessionUpdateRequestIEs)createOT() interface{}{
    return nil
}
var table_MBMSSessionUpdateRequestIEs = make(map[int]*M3APPROTOCOLIES)

var order_MBMSSessionUpdateRequestIEs = make([]int, 11)

func (self *MBMSSessionUpdateRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.MMEMBMSM3APID
   count +=1 //self.MCEMBMSM3APID
   count +=1 //self.TMGI
   if self.MBMSSessionID != nil { count += 1 }
   count +=1 //self.MBMSERABQoSParameters
   count +=1 //self.MBMSSessionDuration
   if self.MBMSServiceArea != nil { count += 1 }
   count +=1 //self.MinimumTimeToMBMSDataTransfer
   if self.TNLInformation != nil { count += 1 }
   if self.TimeofMBMSDataTransfer != nil { count += 1 }
   if self.MBMSCellList != nil { count += 1 }
   return count//ObjSet
}
func (self *MBMSSessionUpdateRequestIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        return true //self.MMEMBMSM3APID
      case 1: //MCEMBMSM3APID
        return true //self.MCEMBMSM3APID
      case 2: //TMGI
        return true //self.TMGI
      case 3: //MBMSSessionID
        if self.MBMSSessionID != nil { return true }
      case 4: //MBMSERABQoSParameters
        return true //self.MBMSERABQoSParameters
      case 5: //MBMSSessionDuration
        return true //self.MBMSSessionDuration
      case 6: //MBMSServiceArea
        if self.MBMSServiceArea != nil { return true }
      case 16: //MinimumTimeToMBMSDataTransfer
        return true //self.MinimumTimeToMBMSDataTransfer
      case 7: //TNLInformation
        if self.TNLInformation != nil { return true }
      case 21: //TimeofMBMSDataTransfer
        if self.TimeofMBMSDataTransfer != nil { return true }
      case 25: //MBMSCellList
        if self.MBMSCellList != nil { return true }
   }
   return false//ObjSet
}
func (self *MBMSSessionUpdateRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MMEMBMSM3APID)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM3APID)
      case 2: //TMGI
        self.TMGI.Unpack(st)
        self.list = append(self.list, &self.TMGI)
      case 3: //MBMSSessionID
        self.MBMSSessionID = &MBMSSessionID{}
        self.MBMSSessionID.Unpack(st)
        self.list = append(self.list, self.MBMSSessionID)
      case 4: //MBMSERABQoSParameters
        self.MBMSERABQoSParameters.Unpack(st)
        self.list = append(self.list, &self.MBMSERABQoSParameters)
      case 5: //MBMSSessionDuration
        self.MBMSSessionDuration.Unpack(st)
        self.list = append(self.list, &self.MBMSSessionDuration)
      case 6: //MBMSServiceArea
        self.MBMSServiceArea = &MBMSServiceArea{}
        self.MBMSServiceArea.Unpack(st)
        self.list = append(self.list, self.MBMSServiceArea)
      case 16: //MinimumTimeToMBMSDataTransfer
        self.MinimumTimeToMBMSDataTransfer.Unpack(st)
        self.list = append(self.list, &self.MinimumTimeToMBMSDataTransfer)
      case 7: //TNLInformation
        self.TNLInformation = &TNLInformation{}
        self.TNLInformation.Unpack(st)
        self.list = append(self.list, self.TNLInformation)
      case 21: //TimeofMBMSDataTransfer
        self.TimeofMBMSDataTransfer = &AbsoluteTimeofMBMSData{}
        self.TimeofMBMSDataTransfer.Unpack(st)
        self.list = append(self.list, self.TimeofMBMSDataTransfer)
      case 25: //MBMSCellList
        self.MBMSCellList = &MBMSCellList{}
        self.MBMSCellList.Unpack(st)
        self.list = append(self.list, self.MBMSCellList)
   }
}
func (self *MBMSSessionUpdateRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Pack(st)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Pack(st)
      case 2: //TMGI
        self.TMGI.Pack(st)
      case 3: //MBMSSessionID
        if self.MBMSSessionID != nil {self.MBMSSessionID.Pack(st)}
      case 4: //MBMSERABQoSParameters
        self.MBMSERABQoSParameters.Pack(st)
      case 5: //MBMSSessionDuration
        self.MBMSSessionDuration.Pack(st)
      case 6: //MBMSServiceArea
        if self.MBMSServiceArea != nil {self.MBMSServiceArea.Pack(st)}
      case 16: //MinimumTimeToMBMSDataTransfer
        self.MinimumTimeToMBMSDataTransfer.Pack(st)
      case 7: //TNLInformation
        if self.TNLInformation != nil {self.TNLInformation.Pack(st)}
      case 21: //TimeofMBMSDataTransfer
        if self.TimeofMBMSDataTransfer != nil {self.TimeofMBMSDataTransfer.Pack(st)}
      case 25: //MBMSCellList
        if self.MBMSCellList != nil {self.MBMSCellList.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBMSSessionUpdateRequestIEs[0] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMMEMBMSM3APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MMEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateRequestIEs[0] = 0
table_MBMSSessionUpdateRequestIEs[1] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM3APID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MCEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateRequestIEs[1] = 1
table_MBMSSessionUpdateRequestIEs[2] = &M3APPROTOCOLIES{ID:ProtocolIEID{idTMGI}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TMGI{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateRequestIEs[2] = 2
table_MBMSSessionUpdateRequestIEs[3] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSSessionID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSSessionID{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionUpdateRequestIEs[3] = 3
table_MBMSSessionUpdateRequestIEs[4] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSERABQoSParameters}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSERABQoSParameters{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateRequestIEs[4] = 4
table_MBMSSessionUpdateRequestIEs[5] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSSessionDuration}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSSessionDuration{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateRequestIEs[5] = 5
table_MBMSSessionUpdateRequestIEs[6] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceArea}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSServiceArea{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionUpdateRequestIEs[6] = 6
table_MBMSSessionUpdateRequestIEs[16] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMinimumTimeToMBMSDataTransfer}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MinimumTimeToMBMSDataTransfer{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateRequestIEs[7] = 16
table_MBMSSessionUpdateRequestIEs[7] = &M3APPROTOCOLIES{ID:ProtocolIEID{idTNLInformation}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TNLInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionUpdateRequestIEs[8] = 7
table_MBMSSessionUpdateRequestIEs[21] = &M3APPROTOCOLIES{ID:ProtocolIEID{idTimeofMBMSDataTransfer}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&AbsoluteTimeofMBMSData{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionUpdateRequestIEs[9] = 21
table_MBMSSessionUpdateRequestIEs[25] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSCellList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSCellList{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionUpdateRequestIEs[10] = 25
   }

type MBMSSessionUpdateResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MME-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MME-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MCE-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MMEMBMSM3APID  MMEMBMSM3APID
   MCEMBMSM3APID  MCEMBMSM3APID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MBMSSessionUpdateResponseIEs)createOT() interface{}{
    return nil
}
var table_MBMSSessionUpdateResponseIEs = make(map[int]*M3APPROTOCOLIES)

var order_MBMSSessionUpdateResponseIEs = make([]int, 3)

func (self *MBMSSessionUpdateResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.MMEMBMSM3APID
   count +=1 //self.MCEMBMSM3APID
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MBMSSessionUpdateResponseIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        return true //self.MMEMBMSM3APID
      case 1: //MCEMBMSM3APID
        return true //self.MCEMBMSM3APID
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MBMSSessionUpdateResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MMEMBMSM3APID)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM3APID)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MBMSSessionUpdateResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Pack(st)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBMSSessionUpdateResponseIEs[0] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMMEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MMEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateResponseIEs[0] = 0
table_MBMSSessionUpdateResponseIEs[1] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateResponseIEs[1] = 1
table_MBMSSessionUpdateResponseIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionUpdateResponseIEs[2] = 8
   }

type MBMSSessionUpdateFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MME-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MME-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MCE-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M3AP-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MMEMBMSM3APID  MMEMBMSM3APID
   MCEMBMSM3APID  MCEMBMSM3APID
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MBMSSessionUpdateFailureIEs)createOT() interface{}{
    return nil
}
var table_MBMSSessionUpdateFailureIEs = make(map[int]*M3APPROTOCOLIES)

var order_MBMSSessionUpdateFailureIEs = make([]int, 4)

func (self *MBMSSessionUpdateFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.MMEMBMSM3APID
   count +=1 //self.MCEMBMSM3APID
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MBMSSessionUpdateFailureIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        return true //self.MMEMBMSM3APID
      case 1: //MCEMBMSM3APID
        return true //self.MCEMBMSM3APID
      case 9: //Cause
        return true //self.Cause
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MBMSSessionUpdateFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MMEMBMSM3APID)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Unpack(st)
        self.list = append(self.list, &self.MCEMBMSM3APID)
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MBMSSessionUpdateFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID.Pack(st)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID.Pack(st)
      case 9: //Cause
        self.Cause.Pack(st)
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MBMSSessionUpdateFailureIEs[0] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMMEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MMEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateFailureIEs[0] = 0
table_MBMSSessionUpdateFailureIEs[1] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM3APID{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateFailureIEs[1] = 1
table_MBMSSessionUpdateFailureIEs[9] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSSessionUpdateFailureIEs[2] = 9
table_MBMSSessionUpdateFailureIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MBMSSessionUpdateFailureIEs[3] = 8
   }

type ErrorIndicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MME-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MME-MBMS-M3AP-ID', 'PRESENCE': 'optional'}, {'ID': 'id-MCE-MBMS-M3AP-ID', 'CRITICALITY': 'ignore', 'TYPE': 'MCE-MBMS-M3AP-ID', 'PRESENCE': 'optional'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MMEMBMSM3APID  *MMEMBMSM3APID
   MCEMBMSM3APID  *MCEMBMSM3APID
   Cause  *Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ErrorIndicationIEs)createOT() interface{}{
    return nil
}
var table_ErrorIndicationIEs = make(map[int]*M3APPROTOCOLIES)

var order_ErrorIndicationIEs = make([]int, 4)

func (self *ErrorIndicationIEs) GetIECount() int{
   count := 0
   if self.MMEMBMSM3APID != nil { count += 1 }
   if self.MCEMBMSM3APID != nil { count += 1 }
   if self.Cause != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ErrorIndicationIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        if self.MMEMBMSM3APID != nil { return true }
      case 1: //MCEMBMSM3APID
        if self.MCEMBMSM3APID != nil { return true }
      case 9: //Cause
        if self.Cause != nil { return true }
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ErrorIndicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        self.MMEMBMSM3APID = &MMEMBMSM3APID{}
        self.MMEMBMSM3APID.Unpack(st)
        self.list = append(self.list, self.MMEMBMSM3APID)
      case 1: //MCEMBMSM3APID
        self.MCEMBMSM3APID = &MCEMBMSM3APID{}
        self.MCEMBMSM3APID.Unpack(st)
        self.list = append(self.list, self.MCEMBMSM3APID)
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
func (self *ErrorIndicationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 0: //MMEMBMSM3APID
        if self.MMEMBMSM3APID != nil {self.MMEMBMSM3APID.Pack(st)}
      case 1: //MCEMBMSM3APID
        if self.MCEMBMSM3APID != nil {self.MCEMBMSM3APID.Pack(st)}
      case 9: //Cause
        if self.Cause != nil {self.Cause.Pack(st)}
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ErrorIndicationIEs[0] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMMEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MMEMBMSM3APID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[0] = 0
table_ErrorIndicationIEs[1] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMCEMBMSM3APID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEMBMSM3APID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[1] = 1
table_ErrorIndicationIEs[9] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[2] = 9
table_ErrorIndicationIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[3] = 8
   }

type ResetIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-ResetType', 'CRITICALITY': 'reject', 'TYPE': 'ResetType', 'PRESENCE': 'mandatory'}, None]}
   Cause  Cause
   ResetType  ResetType
   list []interface{}
}
func (self *ResetIEs)createOT() interface{}{
    return nil
}
var table_ResetIEs = make(map[int]*M3APPROTOCOLIES)

var order_ResetIEs = make([]int, 2)

func (self *ResetIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   count +=1 //self.ResetType
   return count//ObjSet
}
func (self *ResetIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        return true //self.Cause
      case 13: //ResetType
        return true //self.ResetType
   }
   return false//ObjSet
}
func (self *ResetIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 13: //ResetType
        self.ResetType.Unpack(st)
        self.list = append(self.list, &self.ResetType)
   }
}
func (self *ResetIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Pack(st)
      case 13: //ResetType
        self.ResetType.Pack(st)
      default:
      break
   }
}
func init() {
table_ResetIEs[9] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetIEs[0] = 9
table_ResetIEs[13] = &M3APPROTOCOLIES{ID:ProtocolIEID{idResetType}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ResetType{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetIEs[1] = 13
   }

type MBMSServiceassociatedLogicalM3ConnectionItemRes struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBMS-Service-associatedLogicalM3-ConnectionItem', 'CRITICALITY': 'reject', 'TYPE': 'MBMS-Service-associatedLogicalM3-ConnectionItem', 'PRESENCE': 'mandatory'}, None]}
   MBMSServiceassociatedLogicalM3ConnectionItem  MBMSServiceassociatedLogicalM3ConnectionItem
   list []interface{}
}
func (self *MBMSServiceassociatedLogicalM3ConnectionItemRes)createOT() interface{}{
    return nil
}
var table_MBMSServiceassociatedLogicalM3ConnectionItemRes = make(map[int]*M3APPROTOCOLIES)

var order_MBMSServiceassociatedLogicalM3ConnectionItemRes = make([]int, 1)

func (self *MBMSServiceassociatedLogicalM3ConnectionItemRes) GetIECount() int{
   count := 0
   count +=1 //self.MBMSServiceassociatedLogicalM3ConnectionItem
   return count//ObjSet
}
func (self *MBMSServiceassociatedLogicalM3ConnectionItemRes) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 14: //MBMSServiceassociatedLogicalM3ConnectionItem
        return true //self.MBMSServiceassociatedLogicalM3ConnectionItem
   }
   return false//ObjSet
}
func (self *MBMSServiceassociatedLogicalM3ConnectionItemRes)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 14: //MBMSServiceassociatedLogicalM3ConnectionItem
        self.MBMSServiceassociatedLogicalM3ConnectionItem.Unpack(st)
        self.list = append(self.list, &self.MBMSServiceassociatedLogicalM3ConnectionItem)
   }
}
func (self *MBMSServiceassociatedLogicalM3ConnectionItemRes)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 14: //MBMSServiceassociatedLogicalM3ConnectionItem
        self.MBMSServiceassociatedLogicalM3ConnectionItem.Pack(st)
      default:
      break
   }
}
func init() {
table_MBMSServiceassociatedLogicalM3ConnectionItemRes[14] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceassociatedLogicalM3ConnectionItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSServiceassociatedLogicalM3ConnectionItem{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSServiceassociatedLogicalM3ConnectionItemRes[0] = 14
   }

type ResetAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBMS-Service-associatedLogicalM3-ConnectionListResAck', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Service-associatedLogicalM3-ConnectionListResAck', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   MBMSServiceassociatedLogicalM3ConnectionListResAck  *MBMSServiceassociatedLogicalM3ConnectionListResAck
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ResetAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_ResetAcknowledgeIEs = make(map[int]*M3APPROTOCOLIES)

var order_ResetAcknowledgeIEs = make([]int, 2)

func (self *ResetAcknowledgeIEs) GetIECount() int{
   count := 0
   if self.MBMSServiceassociatedLogicalM3ConnectionListResAck != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 15: //MBMSServiceassociatedLogicalM3ConnectionListResAck
        if self.MBMSServiceassociatedLogicalM3ConnectionListResAck != nil { return true }
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 15: //MBMSServiceassociatedLogicalM3ConnectionListResAck
        self.MBMSServiceassociatedLogicalM3ConnectionListResAck = &MBMSServiceassociatedLogicalM3ConnectionListResAck{}
        self.MBMSServiceassociatedLogicalM3ConnectionListResAck.Unpack(st)
        self.list = append(self.list, self.MBMSServiceassociatedLogicalM3ConnectionListResAck)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ResetAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 15: //MBMSServiceassociatedLogicalM3ConnectionListResAck
        if self.MBMSServiceassociatedLogicalM3ConnectionListResAck != nil {self.MBMSServiceassociatedLogicalM3ConnectionListResAck.Pack(st)}
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetAcknowledgeIEs[15] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceassociatedLogicalM3ConnectionListResAck}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSServiceassociatedLogicalM3ConnectionListResAck{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetAcknowledgeIEs[0] = 15
table_ResetAcknowledgeIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetAcknowledgeIEs[1] = 8
   }

type MBMSServiceassociatedLogicalM3ConnectionItemResAck struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-MBMS-Service-associatedLogicalM3-ConnectionItem', 'CRITICALITY': 'ignore', 'TYPE': 'MBMS-Service-associatedLogicalM3-ConnectionItem', 'PRESENCE': 'mandatory'}, None]}
   MBMSServiceassociatedLogicalM3ConnectionItem  MBMSServiceassociatedLogicalM3ConnectionItem
   list []interface{}
}
func (self *MBMSServiceassociatedLogicalM3ConnectionItemResAck)createOT() interface{}{
    return nil
}
var table_MBMSServiceassociatedLogicalM3ConnectionItemResAck = make(map[int]*M3APPROTOCOLIES)

var order_MBMSServiceassociatedLogicalM3ConnectionItemResAck = make([]int, 1)

func (self *MBMSServiceassociatedLogicalM3ConnectionItemResAck) GetIECount() int{
   count := 0
   count +=1 //self.MBMSServiceassociatedLogicalM3ConnectionItem
   return count//ObjSet
}
func (self *MBMSServiceassociatedLogicalM3ConnectionItemResAck) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 14: //MBMSServiceassociatedLogicalM3ConnectionItem
        return true //self.MBMSServiceassociatedLogicalM3ConnectionItem
   }
   return false//ObjSet
}
func (self *MBMSServiceassociatedLogicalM3ConnectionItemResAck)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 14: //MBMSServiceassociatedLogicalM3ConnectionItem
        self.MBMSServiceassociatedLogicalM3ConnectionItem.Unpack(st)
        self.list = append(self.list, &self.MBMSServiceassociatedLogicalM3ConnectionItem)
   }
}
func (self *MBMSServiceassociatedLogicalM3ConnectionItemResAck)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 14: //MBMSServiceassociatedLogicalM3ConnectionItem
        self.MBMSServiceassociatedLogicalM3ConnectionItem.Pack(st)
      default:
      break
   }
}
func init() {
table_MBMSServiceassociatedLogicalM3ConnectionItemResAck[14] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceassociatedLogicalM3ConnectionItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MBMSServiceassociatedLogicalM3ConnectionItem{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSServiceassociatedLogicalM3ConnectionItemResAck[0] = 14
   }

type PrivateMessageIEs struct { //ObjSet 1 {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PRIVATE-IES', 'members': [None]}
   list []interface{}
}
func (self *PrivateMessageIEs)createOT() interface{}{
    return nil
}
var table_PrivateMessageIEs = make(map[int]*M3APPRIVATEIES)

var order_PrivateMessageIEs = make([]int, 0)

type M3SetupRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-Global-MCE-ID', 'CRITICALITY': 'reject', 'TYPE': 'Global-MCE-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-MCEname', 'CRITICALITY': 'ignore', 'TYPE': 'MCEname', 'PRESENCE': 'optional'}, {'ID': 'id-MBMSServiceAreaList', 'CRITICALITY': 'reject', 'TYPE': 'MBMSServiceAreaListItem', 'PRESENCE': 'mandatory'}, None]}
   GlobalMCEID  GlobalMCEID
   MCEname  *MCEname
   MBMSServiceAreaList  MBMSServiceAreaListItem
   list []interface{}
}
func (self *M3SetupRequestIEs)createOT() interface{}{
    return nil
}
var table_M3SetupRequestIEs = make(map[int]*M3APPROTOCOLIES)

var order_M3SetupRequestIEs = make([]int, 3)

func (self *M3SetupRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.GlobalMCEID
   if self.MCEname != nil { count += 1 }
   count +=1 //self.MBMSServiceAreaList
   return count//ObjSet
}
func (self *M3SetupRequestIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 18: //GlobalMCEID
        return true //self.GlobalMCEID
      case 19: //MCEname
        if self.MCEname != nil { return true }
      case 20: //MBMSServiceAreaList
        return true //self.MBMSServiceAreaList
   }
   return false//ObjSet
}
func (self *M3SetupRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 18: //GlobalMCEID
        self.GlobalMCEID.Unpack(st)
        self.list = append(self.list, &self.GlobalMCEID)
      case 19: //MCEname
        self.MCEname = &MCEname{}
        self.MCEname.Unpack(st)
        self.list = append(self.list, self.MCEname)
      case 20: //MBMSServiceAreaList
        self.MBMSServiceAreaList.Unpack(st)
        self.list = append(self.list, &self.MBMSServiceAreaList)
   }
}
func (self *M3SetupRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 18: //GlobalMCEID
        self.GlobalMCEID.Pack(st)
      case 19: //MCEname
        if self.MCEname != nil {self.MCEname.Pack(st)}
      case 20: //MBMSServiceAreaList
        self.MBMSServiceAreaList.Pack(st)
      default:
      break
   }
}
func init() {
table_M3SetupRequestIEs[18] = &M3APPROTOCOLIES{ID:ProtocolIEID{idGlobalMCEID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GlobalMCEID{}, PRESENCE:Presence{Presencemandatory}, }
order_M3SetupRequestIEs[0] = 18
table_M3SetupRequestIEs[19] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMCEname}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEname{}, PRESENCE:Presence{Presenceoptional}, }
order_M3SetupRequestIEs[1] = 19
table_M3SetupRequestIEs[20] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceAreaList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSServiceAreaListItem{}, PRESENCE:Presence{Presencemandatory}, }
order_M3SetupRequestIEs[2] = 20
   }

type M3SetupResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *M3SetupResponseIEs)createOT() interface{}{
    return nil
}
var table_M3SetupResponseIEs = make(map[int]*M3APPROTOCOLIES)

var order_M3SetupResponseIEs = make([]int, 1)

func (self *M3SetupResponseIEs) GetIECount() int{
   count := 0
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *M3SetupResponseIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *M3SetupResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *M3SetupResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_M3SetupResponseIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_M3SetupResponseIEs[0] = 8
   }

type M3SetupFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *M3SetupFailureIEs)createOT() interface{}{
    return nil
}
var table_M3SetupFailureIEs = make(map[int]*M3APPROTOCOLIES)

var order_M3SetupFailureIEs = make([]int, 3)

func (self *M3SetupFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *M3SetupFailureIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        return true //self.Cause
      case 12: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *M3SetupFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 12: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *M3SetupFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Pack(st)
      case 12: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_M3SetupFailureIEs[9] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_M3SetupFailureIEs[0] = 9
table_M3SetupFailureIEs[12] = &M3APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_M3SetupFailureIEs[1] = 12
table_M3SetupFailureIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_M3SetupFailureIEs[2] = 8
   }

type MCEConfigurationUpdateIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-Global-MCE-ID', 'CRITICALITY': 'reject', 'TYPE': 'Global-MCE-ID', 'PRESENCE': 'optional'}, {'ID': 'id-MCEname', 'CRITICALITY': 'ignore', 'TYPE': 'MCEname', 'PRESENCE': 'optional'}, {'ID': 'id-MBMSServiceAreaList', 'CRITICALITY': 'reject', 'TYPE': 'MBMSServiceAreaListItem', 'PRESENCE': 'optional'}, None]}
   GlobalMCEID  *GlobalMCEID
   MCEname  *MCEname
   MBMSServiceAreaList  *MBMSServiceAreaListItem
   list []interface{}
}
func (self *MCEConfigurationUpdateIEs)createOT() interface{}{
    return nil
}
var table_MCEConfigurationUpdateIEs = make(map[int]*M3APPROTOCOLIES)

var order_MCEConfigurationUpdateIEs = make([]int, 3)

func (self *MCEConfigurationUpdateIEs) GetIECount() int{
   count := 0
   if self.GlobalMCEID != nil { count += 1 }
   if self.MCEname != nil { count += 1 }
   if self.MBMSServiceAreaList != nil { count += 1 }
   return count//ObjSet
}
func (self *MCEConfigurationUpdateIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 18: //GlobalMCEID
        if self.GlobalMCEID != nil { return true }
      case 19: //MCEname
        if self.MCEname != nil { return true }
      case 20: //MBMSServiceAreaList
        if self.MBMSServiceAreaList != nil { return true }
   }
   return false//ObjSet
}
func (self *MCEConfigurationUpdateIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 18: //GlobalMCEID
        self.GlobalMCEID = &GlobalMCEID{}
        self.GlobalMCEID.Unpack(st)
        self.list = append(self.list, self.GlobalMCEID)
      case 19: //MCEname
        self.MCEname = &MCEname{}
        self.MCEname.Unpack(st)
        self.list = append(self.list, self.MCEname)
      case 20: //MBMSServiceAreaList
        self.MBMSServiceAreaList = &MBMSServiceAreaListItem{}
        self.MBMSServiceAreaList.Unpack(st)
        self.list = append(self.list, self.MBMSServiceAreaList)
   }
}
func (self *MCEConfigurationUpdateIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 18: //GlobalMCEID
        if self.GlobalMCEID != nil {self.GlobalMCEID.Pack(st)}
      case 19: //MCEname
        if self.MCEname != nil {self.MCEname.Pack(st)}
      case 20: //MBMSServiceAreaList
        if self.MBMSServiceAreaList != nil {self.MBMSServiceAreaList.Pack(st)}
      default:
      break
   }
}
func init() {
table_MCEConfigurationUpdateIEs[18] = &M3APPROTOCOLIES{ID:ProtocolIEID{idGlobalMCEID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GlobalMCEID{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateIEs[0] = 18
table_MCEConfigurationUpdateIEs[19] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMCEname}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&MCEname{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateIEs[1] = 19
table_MCEConfigurationUpdateIEs[20] = &M3APPROTOCOLIES{ID:ProtocolIEID{idMBMSServiceAreaList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&MBMSServiceAreaListItem{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateIEs[2] = 20
   }

type MCEConfigurationUpdateAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MCEConfigurationUpdateAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_MCEConfigurationUpdateAcknowledgeIEs = make(map[int]*M3APPROTOCOLIES)

var order_MCEConfigurationUpdateAcknowledgeIEs = make([]int, 1)

func (self *MCEConfigurationUpdateAcknowledgeIEs) GetIECount() int{
   count := 0
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MCEConfigurationUpdateAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MCEConfigurationUpdateAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MCEConfigurationUpdateAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MCEConfigurationUpdateAcknowledgeIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateAcknowledgeIEs[0] = 8
   }

type MCEConfigurationUpdateFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *MCEConfigurationUpdateFailureIEs)createOT() interface{}{
    return nil
}
var table_MCEConfigurationUpdateFailureIEs = make(map[int]*M3APPROTOCOLIES)

var order_MCEConfigurationUpdateFailureIEs = make([]int, 3)

func (self *MCEConfigurationUpdateFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *MCEConfigurationUpdateFailureIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        return true //self.Cause
      case 12: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *MCEConfigurationUpdateFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 12: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 8: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *MCEConfigurationUpdateFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLIESid).Value
   switch cat {
      case 9: //Cause
        self.Cause.Pack(st)
      case 12: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 8: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_MCEConfigurationUpdateFailureIEs[9] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_MCEConfigurationUpdateFailureIEs[0] = 9
table_MCEConfigurationUpdateFailureIEs[12] = &M3APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateFailureIEs[1] = 12
table_MCEConfigurationUpdateFailureIEs[8] = &M3APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_MCEConfigurationUpdateFailureIEs[2] = 8
   }

type AllocationAndRetentionPriorityExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *AllocationAndRetentionPriorityExtIEs)createOT() interface{}{
    return nil
}
var table_AllocationAndRetentionPriorityExtIEs = make(map[int]*M3APPROTOCOLEXTENSION)

var order_AllocationAndRetentionPriorityExtIEs = make([]int, 0)

type CriticalityDiagnosticsExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsExtIEs = make(map[int]*M3APPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsExtIEs = make([]int, 0)

type CriticalityDiagnosticsIEListExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsIEListExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsIEListExtIEs = make(map[int]*M3APPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsIEListExtIEs = make([]int, 0)

type ECGIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *ECGIExtIEs)createOT() interface{}{
    return nil
}
var table_ECGIExtIEs = make(map[int]*M3APPROTOCOLEXTENSION)

var order_ECGIExtIEs = make([]int, 0)

type GlobalMCEIDExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GlobalMCEIDExtIEs)createOT() interface{}{
    return nil
}
var table_GlobalMCEIDExtIEs = make(map[int]*M3APPROTOCOLEXTENSION)

var order_GlobalMCEIDExtIEs = make([]int, 0)

type GBRQosInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GBRQosInformationExtIEs)createOT() interface{}{
    return nil
}
var table_GBRQosInformationExtIEs = make(map[int]*M3APPROTOCOLEXTENSION)

var order_GBRQosInformationExtIEs = make([]int, 0)

type MBMSERABQoSParametersExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-AllocationAndRetentionPriority', 'CRITICALITY': 'ignore', 'EXTENSION': 'AllocationAndRetentionPriority', 'PRESENCE': 'mandatory'}, None]}
   AllocationAndRetentionPriority  AllocationAndRetentionPriority
   list []interface{}
}
func (self *MBMSERABQoSParametersExtIEs)createOT() interface{}{
    return nil
}
var table_MBMSERABQoSParametersExtIEs = make(map[int]*M3APPROTOCOLEXTENSION)

var order_MBMSERABQoSParametersExtIEs = make([]int, 1)

func (self *MBMSERABQoSParametersExtIEs) GetIECount() int{
   count := 0
   count +=1 //self.AllocationAndRetentionPriority
   return count//ObjSet
}
func (self *MBMSERABQoSParametersExtIEs) GetOT(id interface{}) bool{
   cat := id.(M3APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 17: //AllocationAndRetentionPriority
        return true //self.AllocationAndRetentionPriority
   }
   return false//ObjSet
}
func (self *MBMSERABQoSParametersExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 17: //AllocationAndRetentionPriority
        self.AllocationAndRetentionPriority.Unpack(st)
        self.list = append(self.list, &self.AllocationAndRetentionPriority)
   }
}
func (self *MBMSERABQoSParametersExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(M3APPROTOCOLEXTENSIONid).Value
   switch cat {
      case 17: //AllocationAndRetentionPriority
        self.AllocationAndRetentionPriority.Pack(st)
      default:
      break
   }
}
func init() {
table_MBMSERABQoSParametersExtIEs[17] = &M3APPROTOCOLEXTENSION{ID:ProtocolIEID{idAllocationAndRetentionPriority}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&AllocationAndRetentionPriority{}, PRESENCE:Presence{Presencemandatory}, }
order_MBMSERABQoSParametersExtIEs[0] = 17
   }

type MBMSServiceassociatedLogicalM3ConnectionItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MBMSServiceassociatedLogicalM3ConnectionItemExtIEs)createOT() interface{}{
    return nil
}
var table_MBMSServiceassociatedLogicalM3ConnectionItemExtIEs = make(map[int]*M3APPROTOCOLEXTENSION)

var order_MBMSServiceassociatedLogicalM3ConnectionItemExtIEs = make([]int, 0)

type TMGIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TMGIExtIEs)createOT() interface{}{
    return nil
}
var table_TMGIExtIEs = make(map[int]*M3APPROTOCOLEXTENSION)

var order_TMGIExtIEs = make([]int, 0)

type TNLInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'M3AP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TNLInformationExtIEs)createOT() interface{}{
    return nil
}
var table_TNLInformationExtIEs = make(map[int]*M3APPROTOCOLEXTENSION)

var order_TNLInformationExtIEs = make([]int, 0)

func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'MBMSSessionStartRequest', 'SUCCESSFUL OUTCOME': 'MBMSSessionStartResponse', 'UNSUCCESSFUL OUTCOME': 'MBMSSessionStartFailure', 'PROCEDURE CODE': 'id-mBMSsessionStart', 'CRITICALITY': 'reject'}]}
table_M3APELEMENTARYPROCEDURES[M3APELEMENTARYPROCEDUREprocedureCode{idmBMSsessionStart}] = &M3APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&MBMSSessionStartRequest{}, SUCCESSFULOUTCOME:&MBMSSessionStartResponse{}, UNSUCCESSFULOUTCOME:&MBMSSessionStartFailure{}, PROCEDURECODE:ProcedureCode{idmBMSsessionStart}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetMBMSsessionStartINITIATINGMESSAGE() (*MBMSSessionStartRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &MBMSSessionStartRequest{}, uint64(idmBMSsessionStart), int(Criticalityreject)
}
func GetMBMSsessionStartSUCCESSFULOUTCOME() (*MBMSSessionStartResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &MBMSSessionStartResponse{}, uint64(idmBMSsessionStart), int(Criticalityreject)
}
func GetMBMSsessionStartUNSUCCESSFULOUTCOME() (*MBMSSessionStartFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &MBMSSessionStartFailure{}, uint64(idmBMSsessionStart), int(Criticalityreject)
}
func (self *MBMSSessionStartRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MBMSSessionStartRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MBMSSessionStartRequest) createOT() interface{} {
   return &MBMSSessionStartRequest{}
}
func (self *MBMSSessionStartRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MBMSSessionStartRequest) GetIECount() int{
    return 0
}
func (self *MBMSSessionStartResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MBMSSessionStartResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MBMSSessionStartResponse) createOT() interface{} {
   return &MBMSSessionStartResponse{}
}
func (self *MBMSSessionStartResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MBMSSessionStartResponse) GetIECount() int{
    return 0
}
func (self *MBMSSessionStartFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MBMSSessionStartFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MBMSSessionStartFailure) createOT() interface{} {
   return &MBMSSessionStartFailure{}
}
func (self *MBMSSessionStartFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MBMSSessionStartFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'MBMSSessionStopRequest', 'SUCCESSFUL OUTCOME': 'MBMSSessionStopResponse', 'PROCEDURE CODE': 'id-mBMSsessionStop', 'CRITICALITY': 'reject'}]}
table_M3APELEMENTARYPROCEDURES[M3APELEMENTARYPROCEDUREprocedureCode{idmBMSsessionStop}] = &M3APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&MBMSSessionStopRequest{}, SUCCESSFULOUTCOME:&MBMSSessionStopResponse{}, PROCEDURECODE:ProcedureCode{idmBMSsessionStop}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetMBMSsessionStopINITIATINGMESSAGE() (*MBMSSessionStopRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &MBMSSessionStopRequest{}, uint64(idmBMSsessionStop), int(Criticalityreject)
}
func GetMBMSsessionStopSUCCESSFULOUTCOME() (*MBMSSessionStopResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &MBMSSessionStopResponse{}, uint64(idmBMSsessionStop), int(Criticalityreject)
}
func (self *MBMSSessionStopRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MBMSSessionStopRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MBMSSessionStopRequest) createOT() interface{} {
   return &MBMSSessionStopRequest{}
}
func (self *MBMSSessionStopRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MBMSSessionStopRequest) GetIECount() int{
    return 0
}
func (self *MBMSSessionStopResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MBMSSessionStopResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MBMSSessionStopResponse) createOT() interface{} {
   return &MBMSSessionStopResponse{}
}
func (self *MBMSSessionStopResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MBMSSessionStopResponse) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'MBMSSessionUpdateRequest', 'SUCCESSFUL OUTCOME': 'MBMSSessionUpdateResponse', 'UNSUCCESSFUL OUTCOME': 'MBMSSessionUpdateFailure', 'PROCEDURE CODE': 'id-mBMSsessionUpdate', 'CRITICALITY': 'reject'}]}
table_M3APELEMENTARYPROCEDURES[M3APELEMENTARYPROCEDUREprocedureCode{idmBMSsessionUpdate}] = &M3APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&MBMSSessionUpdateRequest{}, SUCCESSFULOUTCOME:&MBMSSessionUpdateResponse{}, UNSUCCESSFULOUTCOME:&MBMSSessionUpdateFailure{}, PROCEDURECODE:ProcedureCode{idmBMSsessionUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetMBMSsessionUpdateINITIATINGMESSAGE() (*MBMSSessionUpdateRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &MBMSSessionUpdateRequest{}, uint64(idmBMSsessionUpdate), int(Criticalityreject)
}
func GetMBMSsessionUpdateSUCCESSFULOUTCOME() (*MBMSSessionUpdateResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &MBMSSessionUpdateResponse{}, uint64(idmBMSsessionUpdate), int(Criticalityreject)
}
func GetMBMSsessionUpdateUNSUCCESSFULOUTCOME() (*MBMSSessionUpdateFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &MBMSSessionUpdateFailure{}, uint64(idmBMSsessionUpdate), int(Criticalityreject)
}
func (self *MBMSSessionUpdateRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MBMSSessionUpdateRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MBMSSessionUpdateRequest) createOT() interface{} {
   return &MBMSSessionUpdateRequest{}
}
func (self *MBMSSessionUpdateRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MBMSSessionUpdateRequest) GetIECount() int{
    return 0
}
func (self *MBMSSessionUpdateResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MBMSSessionUpdateResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MBMSSessionUpdateResponse) createOT() interface{} {
   return &MBMSSessionUpdateResponse{}
}
func (self *MBMSSessionUpdateResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MBMSSessionUpdateResponse) GetIECount() int{
    return 0
}
func (self *MBMSSessionUpdateFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *MBMSSessionUpdateFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *MBMSSessionUpdateFailure) createOT() interface{} {
   return &MBMSSessionUpdateFailure{}
}
func (self *MBMSSessionUpdateFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *MBMSSessionUpdateFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ErrorIndication', 'PROCEDURE CODE': 'id-errorIndication', 'CRITICALITY': 'ignore'}]}
table_M3APELEMENTARYPROCEDURES[M3APELEMENTARYPROCEDUREprocedureCode{iderrorIndication}] = &M3APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ErrorIndication{}, PROCEDURECODE:ProcedureCode{iderrorIndication}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'Reset', 'SUCCESSFUL OUTCOME': 'ResetAcknowledge', 'PROCEDURE CODE': 'id-Reset', 'CRITICALITY': 'reject'}]}
table_M3APELEMENTARYPROCEDURES[M3APELEMENTARYPROCEDUREprocedureCode{idReset}] = &M3APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&Reset{}, SUCCESSFULOUTCOME:&ResetAcknowledge{}, PROCEDURECODE:ProcedureCode{idReset}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetResetINITIATINGMESSAGE() (*Reset, uint64, int) {/*TYPE, ID, Cricality*/
 return &Reset{}, uint64(idReset), int(Criticalityreject)
}
func GetResetSUCCESSFULOUTCOME() (*ResetAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &ResetAcknowledge{}, uint64(idReset), int(Criticalityreject)
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'PrivateMessage', 'PROCEDURE CODE': 'id-privateMessage', 'CRITICALITY': 'ignore'}]}
table_M3APELEMENTARYPROCEDURES[M3APELEMENTARYPROCEDUREprocedureCode{idprivateMessage}] = &M3APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&PrivateMessage{}, PROCEDURECODE:ProcedureCode{idprivateMessage}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'MCEConfigurationUpdate', 'SUCCESSFUL OUTCOME': 'MCEConfigurationUpdateAcknowledge', 'UNSUCCESSFUL OUTCOME': 'MCEConfigurationUpdateFailure', 'PROCEDURE CODE': 'id-mCEConfigurationUpdate', 'CRITICALITY': 'reject'}]}
table_M3APELEMENTARYPROCEDURES[M3APELEMENTARYPROCEDUREprocedureCode{idmCEConfigurationUpdate}] = &M3APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&MCEConfigurationUpdate{}, SUCCESSFULOUTCOME:&MCEConfigurationUpdateAcknowledge{}, UNSUCCESSFULOUTCOME:&MCEConfigurationUpdateFailure{}, PROCEDURECODE:ProcedureCode{idmCEConfigurationUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'M3AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'M3SetupRequest', 'SUCCESSFUL OUTCOME': 'M3SetupResponse', 'UNSUCCESSFUL OUTCOME': 'M3SetupFailure', 'PROCEDURE CODE': 'id-m3Setup', 'CRITICALITY': 'reject'}]}
table_M3APELEMENTARYPROCEDURES[M3APELEMENTARYPROCEDUREprocedureCode{idm3Setup}] = &M3APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&M3SetupRequest{}, SUCCESSFULOUTCOME:&M3SetupResponse{}, UNSUCCESSFULOUTCOME:&M3SetupFailure{}, PROCEDURECODE:ProcedureCode{idm3Setup}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetM3SetupINITIATINGMESSAGE() (*M3SetupRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &M3SetupRequest{}, uint64(idm3Setup), int(Criticalityreject)
}
func GetM3SetupSUCCESSFULOUTCOME() (*M3SetupResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &M3SetupResponse{}, uint64(idm3Setup), int(Criticalityreject)
}
func GetM3SetupUNSUCCESSFULOUTCOME() (*M3SetupFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &M3SetupFailure{}, uint64(idm3Setup), int(Criticalityreject)
}
func (self *M3SetupRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *M3SetupRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *M3SetupRequest) createOT() interface{} {
   return &M3SetupRequest{}
}
func (self *M3SetupRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *M3SetupRequest) GetIECount() int{
    return 0
}
func (self *M3SetupResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *M3SetupResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *M3SetupResponse) createOT() interface{} {
   return &M3SetupResponse{}
}
func (self *M3SetupResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *M3SetupResponse) GetIECount() int{
    return 0
}
func (self *M3SetupFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *M3SetupFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *M3SetupFailure) createOT() interface{} {
   return &M3SetupFailure{}
}
func (self *M3SetupFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *M3SetupFailure) GetIECount() int{
    return 0
}
var maxPrivateIEs uint64 = 65535
var maxProtocolExtensions uint64 = 65535
var maxProtocolIEs uint64 = 65535
var idmBMSsessionStart uint64 = 0
const ProcedureCodemBMSsessionStart = 0
var idmBMSsessionStop uint64 = 1
const ProcedureCodemBMSsessionStop = 1
var iderrorIndication uint64 = 2
const ProcedureCodeerrorIndication = 2
var idprivateMessage uint64 = 3
const ProcedureCodeprivateMessage = 3
var idReset uint64 = 4
const ProcedureCodeReset = 4
var idmBMSsessionUpdate uint64 = 5
const ProcedureCodemBMSsessionUpdate = 5
var idmCEConfigurationUpdate uint64 = 6
const ProcedureCodemCEConfigurationUpdate = 6
var idm3Setup uint64 = 7
const ProcedureCodem3Setup = 7
var maxnoofMBMSServiceAreaIdentitiesPerMCE uint64 = 65536
var maxnooferrors uint64 = 256
var maxNrOfIndividualM3ConnectionsToReset uint64 = 256
var maxnoofCellsforMBMS uint64 = 4096
var idMMEMBMSM3APID uint64 = 0
const ProtocolIEIDMMEMBMSM3APID = 0
var idMCEMBMSM3APID uint64 = 1
const ProtocolIEIDMCEMBMSM3APID = 1
var idTMGI uint64 = 2
const ProtocolIEIDTMGI = 2
var idMBMSSessionID uint64 = 3
const ProtocolIEIDMBMSSessionID = 3
var idMBMSERABQoSParameters uint64 = 4
const ProtocolIEIDMBMSERABQoSParameters = 4
var idMBMSSessionDuration uint64 = 5
const ProtocolIEIDMBMSSessionDuration = 5
var idMBMSServiceArea uint64 = 6
const ProtocolIEIDMBMSServiceArea = 6
var idTNLInformation uint64 = 7
const ProtocolIEIDTNLInformation = 7
var idCriticalityDiagnostics uint64 = 8
const ProtocolIEIDCriticalityDiagnostics = 8
var idCause uint64 = 9
const ProtocolIEIDCause = 9
var idMBMSServiceAreaList uint64 = 10
const ProtocolIEIDMBMSServiceAreaList = 10
var idMBMSServiceAreaListItem uint64 = 11
const ProtocolIEIDMBMSServiceAreaListItem = 11
var idTimeToWait uint64 = 12
const ProtocolIEIDTimeToWait = 12
var idResetType uint64 = 13
const ProtocolIEIDResetType = 13
var idMBMSServiceassociatedLogicalM3ConnectionItem uint64 = 14
const ProtocolIEIDMBMSServiceassociatedLogicalM3ConnectionItem = 14
var idMBMSServiceassociatedLogicalM3ConnectionListResAck uint64 = 15
const ProtocolIEIDMBMSServiceassociatedLogicalM3ConnectionListResAck = 15
var idMinimumTimeToMBMSDataTransfer uint64 = 16
const ProtocolIEIDMinimumTimeToMBMSDataTransfer = 16
var idAllocationAndRetentionPriority uint64 = 17
const ProtocolIEIDAllocationAndRetentionPriority = 17
var idGlobalMCEID uint64 = 18
const ProtocolIEIDGlobalMCEID = 18
var idMCEname uint64 = 19
const ProtocolIEIDMCEname = 19
var idMBMSServiceAreaList uint64 = 20
const ProtocolIEIDMBMSServiceAreaList = 20
var idTimeofMBMSDataTransfer uint64 = 21
const ProtocolIEIDTimeofMBMSDataTransfer = 21
var idTimeofMBMSDataStop uint64 = 22
const ProtocolIEIDTimeofMBMSDataStop = 22
var idReestablishment uint64 = 23
const ProtocolIEIDReestablishment = 23
var idAlternativeTNLInformation uint64 = 24
const ProtocolIEIDAlternativeTNLInformation = 24
var idMBMSCellList uint64 = 25
const ProtocolIEIDMBMSCellList = 25
