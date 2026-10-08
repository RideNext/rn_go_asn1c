
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package hnbap
import (
  log "github.com/sirupsen/logrus"
)
var version = "vf50"

func fmthnbap() {log.Debug("hnbap")}
func (self *HNBAPPDU)Unpack(stream *Stream) {
    //coptions := []string{"initiatingMessage","successfulOutcome","unsuccessfulOutcome","Unknown"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in HNBAPPDU\n", choice, choice_len)
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
func (self * HNBAPPDU) Pack(stream *Stream) {
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
type HNBAPPDU struct { //[{'type': 'InitiatingMessage', 'name': 'initiatingMessage'}, {'type': 'SuccessfulOutcome', 'name': 'successfulOutcome'}, {'type': 'UnsuccessfulOutcome', 'name': 'unsuccessfulOutcome'}, None]
    InitiatingMessage *InitiatingMessage
    SuccessfulOutcome *SuccessfulOutcome
    UnsuccessfulOutcome *UnsuccessfulOutcome
} // HNBAPPDU

type InitiatingMessage struct { // [{'type': 'HNBAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'HNBAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['HNBAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['HNBAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode HNBAPELEMENTARYPROCEDUREprocedureCode
    Criticality HNBAPELEMENTARYPROCEDUREcriticality
    Value HNBAPELEMENTARYPROCEDUREInitiatingMessage
}

func (self * InitiatingMessage) Unpack(stream *Stream) {
    //table {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'HNBAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_HNBAPELEMENTARYPROCEDURES[key].INITIATINGMESSAGE
    self.Value = out.(HNBAPELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(HNBAPELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'HNBAP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['HNBAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * InitiatingMessage) Pack(stream *Stream) {
    //table {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'HNBAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_HNBAPELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(HNBAPELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type SuccessfulOutcome struct { // [{'type': 'HNBAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'HNBAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['HNBAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['HNBAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode HNBAPELEMENTARYPROCEDUREprocedureCode
    Criticality HNBAPELEMENTARYPROCEDUREcriticality
    Value HNBAPELEMENTARYPROCEDURESuccessfulOutcome
}

func (self * SuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'HNBAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_HNBAPELEMENTARYPROCEDURES[key].SUCCESSFULOUTCOME
    self.Value = out.(HNBAPELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(HNBAPELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'HNBAP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['HNBAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * SuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'HNBAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_HNBAPELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(HNBAPELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type UnsuccessfulOutcome struct { // [{'type': 'HNBAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'HNBAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['HNBAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['HNBAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode HNBAPELEMENTARYPROCEDUREprocedureCode
    Criticality HNBAPELEMENTARYPROCEDUREcriticality
    Value HNBAPELEMENTARYPROCEDUREUnsuccessfulOutcome
}

func (self * UnsuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'HNBAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_HNBAPELEMENTARYPROCEDURES[key].UNSUCCESSFULOUTCOME
    self.Value = out.(HNBAPELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(HNBAPELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'HNBAP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['HNBAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * UnsuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'HNBAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'HNBAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_HNBAPELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(HNBAPELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type HNBRegisterRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['HNBRegisterRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBRegisterRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs HNBRegisterRequestIEs
    ProtocolExtensions *HNBRegisterRequestExtensions
}

func (self * HNBRegisterRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_HNBRegisterRequestIEs, order_HNBRegisterRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &HNBRegisterRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBRegisterRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_HNBRegisterRequestExtensions, order_HNBRegisterRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBRegisterRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_HNBRegisterRequestIEs, order_HNBRegisterRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_HNBRegisterRequestExtensions, order_HNBRegisterRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type HNBRegisterAccept struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['HNBRegisterResponseIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBRegisterResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs HNBRegisterResponseIEs
    ProtocolExtensions *HNBRegisterResponseExtensions
}

func (self * HNBRegisterAccept) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_HNBRegisterResponseIEs, order_HNBRegisterResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &HNBRegisterResponseExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBRegisterResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_HNBRegisterResponseExtensions, order_HNBRegisterResponseExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBRegisterAccept) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_HNBRegisterResponseIEs, order_HNBRegisterResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_HNBRegisterResponseExtensions, order_HNBRegisterResponseExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type HNBRegisterReject struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['HNBRegisterRejectIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBRegisterRejectExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs HNBRegisterRejectIEs
    ProtocolExtensions *HNBRegisterRejectExtensions
}

func (self * HNBRegisterReject) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_HNBRegisterRejectIEs, order_HNBRegisterRejectIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &HNBRegisterRejectExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBRegisterRejectExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_HNBRegisterRejectExtensions, order_HNBRegisterRejectExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBRegisterReject) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_HNBRegisterRejectIEs, order_HNBRegisterRejectIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_HNBRegisterRejectExtensions, order_HNBRegisterRejectExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type HNBDeRegister struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['HNBDe-RegisterIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBDe-RegisterExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs HNBDeRegisterIEs
    ProtocolExtensions *HNBDeRegisterExtensions
}

func (self * HNBDeRegister) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_HNBDeRegisterIEs, order_HNBDeRegisterIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &HNBDeRegisterExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBDe-RegisterExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_HNBDeRegisterExtensions, order_HNBDeRegisterExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBDeRegister) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_HNBDeRegisterIEs, order_HNBDeRegisterIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_HNBDeRegisterExtensions, order_HNBDeRegisterExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UERegisterRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['UERegisterRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UERegisterRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs UERegisterRequestIEs
    ProtocolExtensions *UERegisterRequestExtensions
}

func (self * UERegisterRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_UERegisterRequestIEs, order_UERegisterRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &UERegisterRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UERegisterRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_UERegisterRequestExtensions, order_UERegisterRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UERegisterRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_UERegisterRequestIEs, order_UERegisterRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_UERegisterRequestExtensions, order_UERegisterRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UERegisterAccept struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['UERegisterAcceptIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UERegisterAcceptExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs UERegisterAcceptIEs
    ProtocolExtensions *UERegisterAcceptExtensions
}

func (self * UERegisterAccept) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_UERegisterAcceptIEs, order_UERegisterAcceptIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &UERegisterAcceptExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UERegisterAcceptExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_UERegisterAcceptExtensions, order_UERegisterAcceptExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UERegisterAccept) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_UERegisterAcceptIEs, order_UERegisterAcceptIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_UERegisterAcceptExtensions, order_UERegisterAcceptExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UERegisterReject struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['UERegisterRejectIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UERegisterRejectExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs UERegisterRejectIEs
    ProtocolExtensions *UERegisterRejectExtensions
}

func (self * UERegisterReject) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_UERegisterRejectIEs, order_UERegisterRejectIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &UERegisterRejectExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UERegisterRejectExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_UERegisterRejectExtensions, order_UERegisterRejectExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UERegisterReject) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_UERegisterRejectIEs, order_UERegisterRejectIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_UERegisterRejectExtensions, order_UERegisterRejectExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UEDeRegister struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['UEDe-RegisterIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UEDe-RegisterExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs UEDeRegisterIEs
    ProtocolExtensions *UEDeRegisterExtensions
}

func (self * UEDeRegister) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_UEDeRegisterIEs, order_UEDeRegisterIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &UEDeRegisterExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UEDe-RegisterExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_UEDeRegisterExtensions, order_UEDeRegisterExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEDeRegister) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_UEDeRegisterIEs, order_UEDeRegisterIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_UEDeRegisterExtensions, order_UEDeRegisterExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type CSGMembershipUpdate struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['CSGMembershipUpdateIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CSGMembershipUpdateExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs CSGMembershipUpdateIEs
    ProtocolExtensions *CSGMembershipUpdateExtensions
}

func (self * CSGMembershipUpdate) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_CSGMembershipUpdateIEs, order_CSGMembershipUpdateIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &CSGMembershipUpdateExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CSGMembershipUpdateExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_CSGMembershipUpdateExtensions, order_CSGMembershipUpdateExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CSGMembershipUpdate) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_CSGMembershipUpdateIEs, order_CSGMembershipUpdateIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_CSGMembershipUpdateExtensions, order_CSGMembershipUpdateExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type AccessControlQuery struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['AccessControlQueryIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AccessControlQueryExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs AccessControlQueryIEs
    ProtocolExtensions *AccessControlQueryExtensions
}

func (self * AccessControlQuery) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_AccessControlQueryIEs, order_AccessControlQueryIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &AccessControlQueryExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AccessControlQueryExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_AccessControlQueryExtensions, order_AccessControlQueryExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AccessControlQuery) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_AccessControlQueryIEs, order_AccessControlQueryIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_AccessControlQueryExtensions, order_AccessControlQueryExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type AccessControlResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['AccessControlResponseIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AccessControlResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs AccessControlResponseIEs
    ProtocolExtensions *AccessControlResponseExtensions
}

func (self * AccessControlResponse) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_AccessControlResponseIEs, order_AccessControlResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &AccessControlResponseExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AccessControlResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_AccessControlResponseExtensions, order_AccessControlResponseExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AccessControlResponse) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_AccessControlResponseIEs, order_AccessControlResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_AccessControlResponseExtensions, order_AccessControlResponseExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type TNLUpdateRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['TNLUpdateRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TNLUpdateExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs TNLUpdateRequestIEs
    ProtocolExtensions *TNLUpdateExtensions
}

func (self * TNLUpdateRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_TNLUpdateRequestIEs, order_TNLUpdateRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &TNLUpdateExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TNLUpdateExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_TNLUpdateExtensions, order_TNLUpdateExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TNLUpdateRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_TNLUpdateRequestIEs, order_TNLUpdateRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_TNLUpdateExtensions, order_TNLUpdateExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type TNLUpdateResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['TNLUpdateResponseIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TNLUpdateResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs TNLUpdateResponseIEs
    ProtocolExtensions *TNLUpdateResponseExtensions
}

func (self * TNLUpdateResponse) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_TNLUpdateResponseIEs, order_TNLUpdateResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &TNLUpdateResponseExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TNLUpdateResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_TNLUpdateResponseExtensions, order_TNLUpdateResponseExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TNLUpdateResponse) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_TNLUpdateResponseIEs, order_TNLUpdateResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_TNLUpdateResponseExtensions, order_TNLUpdateResponseExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type TNLUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['TNLUpdateFailureIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TNLUpdateFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs TNLUpdateFailureIEs
    ProtocolExtensions *TNLUpdateFailureExtensions
}

func (self * TNLUpdateFailure) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_TNLUpdateFailureIEs, order_TNLUpdateFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &TNLUpdateFailureExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TNLUpdateFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_TNLUpdateFailureExtensions, order_TNLUpdateFailureExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TNLUpdateFailure) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_TNLUpdateFailureIEs, order_TNLUpdateFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_TNLUpdateFailureExtensions, order_TNLUpdateFailureExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type HNBConfigTransferRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['HNBConfigTransferRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBConfigTransferRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs HNBConfigTransferRequestIEs
    ProtocolExtensions *HNBConfigTransferRequestExtensions
}

func (self * HNBConfigTransferRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_HNBConfigTransferRequestIEs, order_HNBConfigTransferRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &HNBConfigTransferRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBConfigTransferRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_HNBConfigTransferRequestExtensions, order_HNBConfigTransferRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBConfigTransferRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_HNBConfigTransferRequestIEs, order_HNBConfigTransferRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_HNBConfigTransferRequestExtensions, order_HNBConfigTransferRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type HNBConfigTransferResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['HNBConfigTransferResponseIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBConfigTransferResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs HNBConfigTransferResponseIEs
    ProtocolExtensions *HNBConfigTransferResponseExtensions
}

func (self * HNBConfigTransferResponse) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_HNBConfigTransferResponseIEs, order_HNBConfigTransferResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &HNBConfigTransferResponseExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBConfigTransferResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_HNBConfigTransferResponseExtensions, order_HNBConfigTransferResponseExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBConfigTransferResponse) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_HNBConfigTransferResponseIEs, order_HNBConfigTransferResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_HNBConfigTransferResponseExtensions, order_HNBConfigTransferResponseExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type HNBConfigTransferFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['HNBConfigTransferFailureIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBConfigTransferFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs HNBConfigTransferFailureIEs
    ProtocolExtensions *HNBConfigTransferFailureExtensions
}

func (self * HNBConfigTransferFailure) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_HNBConfigTransferFailureIEs, order_HNBConfigTransferFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &HNBConfigTransferFailureExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBConfigTransferFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_HNBConfigTransferFailureExtensions, order_HNBConfigTransferFailureExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBConfigTransferFailure) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_HNBConfigTransferFailureIEs, order_HNBConfigTransferFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_HNBConfigTransferFailureExtensions, order_HNBConfigTransferFailureExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RelocationComplete struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RelocationCompleteIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationCompleteExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RelocationCompleteIEs
    ProtocolExtensions *RelocationCompleteExtensions
}

func (self * RelocationComplete) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RelocationCompleteIEs, order_RelocationCompleteIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RelocationCompleteExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationCompleteExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RelocationCompleteExtensions, order_RelocationCompleteExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RelocationComplete) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RelocationCompleteIEs, order_RelocationCompleteIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RelocationCompleteExtensions, order_RelocationCompleteExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ErrorIndication struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ErrorIndicationIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ErrorIndicationExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs ErrorIndicationIEs
    ProtocolExtensions *ErrorIndicationExtensions
}

func (self * ErrorIndication) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_ErrorIndicationIEs, order_ErrorIndicationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &ErrorIndicationExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ErrorIndicationExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_ErrorIndicationExtensions, order_ErrorIndicationExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ErrorIndication) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ErrorIndicationIEs, order_ErrorIndicationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_ErrorIndicationExtensions, order_ErrorIndicationExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
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

type HNBHeartBeatRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['HNBHeartBeatRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBHeartBeatRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs HNBHeartBeatRequestIEs
    ProtocolExtensions *HNBHeartBeatRequestExtensions
}

func (self * HNBHeartBeatRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_HNBHeartBeatRequestIEs, order_HNBHeartBeatRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &HNBHeartBeatRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBHeartBeatRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_HNBHeartBeatRequestExtensions, order_HNBHeartBeatRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBHeartBeatRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_HNBHeartBeatRequestIEs, order_HNBHeartBeatRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_HNBHeartBeatRequestExtensions, order_HNBHeartBeatRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type HNBHeartBeatResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['HNBHeartBeatResponseIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBHeartBeatResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs HNBHeartBeatResponseIEs
    ProtocolExtensions *HNBHeartBeatResponseExtensions
}

func (self * HNBHeartBeatResponse) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_HNBHeartBeatResponseIEs, order_HNBHeartBeatResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &HNBHeartBeatResponseExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBHeartBeatResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_HNBHeartBeatResponseExtensions, order_HNBHeartBeatResponseExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBHeartBeatResponse) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_HNBHeartBeatResponseIEs, order_HNBHeartBeatResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_HNBHeartBeatResponseExtensions, order_HNBHeartBeatResponseExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type Accessstratumreleaseindicator struct {
  Value int
}
const (
    Accessstratumreleaseindicatorr99 = 0
    Accessstratumreleaseindicatorrel_4 = 1
    Accessstratumreleaseindicatorrel_5 = 2
    Accessstratumreleaseindicatorrel_6 = 3
    Accessstratumreleaseindicatorrel_7 = 4
    Accessstratumreleaseindicatorrel_8_and_beyond = 5

    /* Extensions */
)
func (self *Accessstratumreleaseindicator) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 6, 1)
}
func (self *Accessstratumreleaseindicator) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 6, 1)
}
type AccessResult struct {
  Value int
}
const (
    AccessResultallowed = 0
    AccessResultnotAllowed = 1

    /* Extensions */
)
func (self *AccessResult) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *AccessResult) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type AltitudeAndDirection struct { // [{'type': 'ENUMERATED', 'values': [('height', 0), ('depth', 1)], 'name': 'directionOfAltitude'}, {'type': 'INTEGER', 'restricted-to': [(0, 32767)], 'name': 'altitude'}, None]
    DirectionOfAltitude ENUMERATED
    Altitude INTEGER
}

func (self * AltitudeAndDirection) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_directionOfAltitude = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(1, 2, 0)
    }
    Unpack_directionOfAltitude(stream, &self.DirectionOfAltitude)// p2
    var Unpack_altitude = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(32768, 15, 0, 0)
    }
    Unpack_altitude(stream, &self.Altitude)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AltitudeAndDirection) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_directionOfAltitude = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 1, 2, 0)
    }
    Pack_directionOfAltitude(stream, self.DirectionOfAltitude) //f2
    var Pack_altitude = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 32768, 15, 0, 0)
    }
    Pack_altitude(stream, self.Altitude) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type BackoffTimer struct {
  Value uint64
}
func (self *BackoffTimer) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(3601, 12, 0, 0)
}
func (self * BackoffTimer) Pack(st *Stream){
    st.formatf_Integer(self.Value, 3601, 12, 0, 0)
}
type BindingID struct {
  Value HexBytes
}
func (self *BindingID) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(3)+1
    if _len < 1 || _len > 4 {
        //fmt.Println ("Invalid len in BindingID")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *BindingID) Pack(st *Stream) {
    _eflag := 0
    if len(self.Value) > 4:
       var _eflag int = 1
    st.format_ext(_eflag)
    if len(self.Value) < 1 || len(self.Value) > 4 {
        log.Error ("Invalid len in BindingID")
        return
}
    st.format_olen((len(self.Value))-1, 3)
    st.formatf_OctString(self.Value, 0)
}
func (self *Cause)Unpack(stream *Stream) {
    //coptions := []string{"radioNetwork","transport","protocol","misc"}
    choice := stream.get_choice(2, 1, 4)
    choice_len := 0
    choice_loc := 0
    if choice >= 4 {
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
        self.Protocol = &CauseProtocol{}//cho6
        self.Protocol.Unpack(stream)
    } else if choice == 3 { //ch2
        self.Misc = &CauseMisc{}//cho6
        self.Misc.Unpack(stream)
    }//end of if else

    if choice >= 4 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * Cause) Pack(stream *Stream) {
    if self.RadioNetwork != nil {
        stream.set_choice(0, 2, 1, 4)
        self.RadioNetwork.Pack(stream)//2
    } else if self.Transport != nil {
        stream.set_choice(1, 2, 1, 4)
        self.Transport.Pack(stream)//2
    } else if self.Protocol != nil {
        stream.set_choice(2, 2, 1, 4)
        self.Protocol.Pack(stream)//2
    } else if self.Misc != nil {
        stream.set_choice(3, 2, 1, 4)
        self.Misc.Pack(stream)//2
    }

}
type Cause struct { //[{'type': 'CauseRadioNetwork', 'name': 'radioNetwork'}, {'type': 'CauseTransport', 'name': 'transport'}, {'type': 'CauseProtocol', 'name': 'protocol'}, {'type': 'CauseMisc', 'name': 'misc'}, None]
    RadioNetwork *CauseRadioNetwork
    Transport *CauseTransport
    Protocol *CauseProtocol
    Misc *CauseMisc
} // Cause

type CauseRadioNetwork struct {
  Value int
}
const (
    CauseRadioNetworkoverload = 0
    CauseRadioNetworkunauthorised_Location = 1
    CauseRadioNetworkunauthorised_HNB = 2
    CauseRadioNetworkhNB_parameter_mismatch = 3
    CauseRadioNetworkinvalid_UE_identity = 4
    CauseRadioNetworkuE_not_allowed_on_this_HNB = 5
    CauseRadioNetworkuE_unauthorised = 6
    CauseRadioNetworkconnection_with_UE_lost = 7
    CauseRadioNetworkue_RRC_telease = 8
    CauseRadioNetworkhNB_not_registered = 9
    CauseRadioNetworkunspecified = 10
    CauseRadioNetworknormal = 11
    CauseRadioNetworkuE_relocated = 12
    CauseRadioNetworkue_registered_in_another_HNB = 13

    /* Extensions */
)
func (self *CauseRadioNetwork) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(5, 14, 1)
}
func (self *CauseRadioNetwork) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 5, 14, 1)
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
type CauseProtocol struct {
  Value int
}
const (
    CauseProtocoltransfer_syntax_error = 0
    CauseProtocolabstract_syntax_error_reject = 1
    CauseProtocolabstract_syntax_error_ignore_and_notify = 2
    CauseProtocolmessage_not_compatible_with_receiver_state = 3
    CauseProtocolsemantic_error = 4
    CauseProtocolunspecified = 5
    CauseProtocolabstract_syntax_error_falsely_constructed_message = 6

    /* Extensions */
)
func (self *CauseProtocol) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 7, 1)
}
func (self *CauseProtocol) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 7, 1)
}
type CauseMisc struct {
  Value int
}
const (
    CauseMiscprocessing_overload = 0
    CauseMischardware_failure = 1
    CauseMisco_and_m_intervention = 2
    CauseMiscunspecified = 3

    /* Extensions */
)
func (self *CauseMisc) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 4, 1)
}
func (self *CauseMisc) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 4, 1)
}
type CellIdentity struct {
  Len int
  Value HexBytes
}
func (self *CellIdentity) Unpack(st *Stream){
    self.Value = st.parsef_BitString(28, 28)
}
func (self *CellIdentity) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 28)
}
type ContextID struct {
  Len int
  Value HexBytes
}
func (self *ContextID) Unpack(st *Stream){
    self.Value = st.parsef_BitString(24, 24)
}
func (self *ContextID) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 24)
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
type CriticalityDiagnosticsIEList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'Criticality', 'name': 'iECriticality'}, {'type': 'ProtocolIE-ID', 'name': 'iE-ID'}, {'type': 'TypeOfError', 'name': 'typeOfError'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-IE-List-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxNrOfErrors')]}
    Items []CriticalityDiagnosticsIEList_Item
}

type CSGID struct {
  Len int
  Value HexBytes
}
func (self *CSGID) Unpack(st *Stream){
    self.Value = st.parsef_BitString(27, 27)
}
func (self *CSGID) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 27)
}
type CSGIndicator struct {
  Value int
}
const (
    CSGIndicatorcsg_capable = 0
    CSGIndicatornot_csg_capable = 1

    /* Extensions */
)
func (self *CSGIndicator) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *CSGIndicator) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type CSGMembershipStatus struct {
  Value int
}
const (
    CSGMembershipStatusmember = 0
    CSGMembershipStatusnon_member = 1

    /* Extensions */
)
func (self *CSGMembershipStatus) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *CSGMembershipStatus) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type CGI struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'LAC', 'name': 'lAC'}, {'type': 'CI', 'name': 'cI'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CGI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PLMNidentity PLMNidentity
    LAC LAC
    CI CI
    IEExtensions *CGIExtIEs
}

func (self * CGI) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.PLMNidentity.Unpack(stream)// p8
    self.LAC.Unpack(stream)// p8
    self.CI.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &CGIExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CGI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_CGIExtIEs, order_CGIExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * CGI) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNidentity.Pack(stream)
    self.LAC.Pack(stream)
    self.CI.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_CGIExtIEs, order_CGIExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type CI struct {
  Value HexBytes
}
func (self *CI) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(2)
}
func (self *CI) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 2)
}
type CNDomainIndicator struct {
  Value int
}
const (
    CNDomainIndicatorcs_domain = 0
    CNDomainIndicatorps_domain = 1
)
func (self *CNDomainIndicator) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 2, 0)
}
func (self *CNDomainIndicator) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 2, 0)
}
type ESN struct {
  Len int
  Value HexBytes
}
func (self *ESN) Unpack(st *Stream){
    self.Value = st.parsef_BitString(32, 32)
}
func (self *ESN) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 32)
}
type GeographicalLocation struct { // [{'type': 'GeographicalCoordinates', 'name': 'geographicalCoordinates'}, {'type': 'AltitudeAndDirection', 'name': 'altitudeAndDirection'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GeographicLocation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    GeographicalCoordinates GeographicalCoordinates
    AltitudeAndDirection AltitudeAndDirection
    IEExtensions *GeographicLocationExtIEs
}

func (self * GeographicalLocation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.GeographicalCoordinates.Unpack(stream)// p8
    self.AltitudeAndDirection.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GeographicLocationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GeographicLocation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GeographicLocationExtIEs, order_GeographicLocationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GeographicalLocation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GeographicalCoordinates.Pack(stream)
    self.AltitudeAndDirection.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GeographicLocationExtIEs, order_GeographicLocationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GeographicalCoordinates struct { // [{'type': 'ENUMERATED', 'values': [('north', 0), ('south', 1)], 'name': 'latitudeSign'}, {'type': 'INTEGER', 'restricted-to': [(0, 8388607)], 'name': 'latitude'}, {'type': 'INTEGER', 'restricted-to': [(-8388608, 8388607)], 'name': 'longitude'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GeographicalCoordinates-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    LatitudeSign ENUMERATED
    Latitude INTEGER
    Longitude INTEGER
    IEExtensions *GeographicalCoordinatesExtIEs
}

func (self * GeographicalCoordinates) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_latitudeSign = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(1, 2, 0)
    }
    Unpack_latitudeSign(stream, &self.LatitudeSign)// p2
    var Unpack_latitude = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(8388608, 23, 0, 0)
    }
    Unpack_latitude(stream, &self.Latitude)// p2
    var Unpack_longitude = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(16777216, 24, 0, -8388608)
    }
    Unpack_longitude(stream, &self.Longitude)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GeographicalCoordinatesExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GeographicalCoordinates-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GeographicalCoordinatesExtIEs, order_GeographicalCoordinatesExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GeographicalCoordinates) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_latitudeSign = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 1, 2, 0)
    }
    Pack_latitudeSign(stream, self.LatitudeSign) //f2
    var Pack_latitude = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 8388608, 23, 0, 0)
    }
    Pack_latitude(stream, self.Latitude) //f2
    var Pack_longitude = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 16777216, 24, 0, -8388608)
    }
    Pack_longitude(stream, self.Longitude) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GeographicalCoordinatesExtIEs, order_GeographicalCoordinatesExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GlobalRNCID struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'RNC-ID', 'name': 'rNCid'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Global-RNC-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNidentity PLMNidentity
    RNCid RNCID
    IEExtensions *GlobalRNCIDExtIEs
}

func (self * GlobalRNCID) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PLMNidentity.Unpack(stream)// p8
    self.RNCid.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GlobalRNCIDExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Global-RNC-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GlobalRNCIDExtIEs, order_GlobalRNCIDExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalRNCID) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNidentity.Pack(stream)
    self.RNCid.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GlobalRNCIDExtIEs, order_GlobalRNCIDExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GTPTEI struct {
  Value HexBytes
}
func (self *GTPTEI) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(4)
}
func (self *GTPTEI) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 4)
}
type HNBCellAccessMode struct {
  Value int
}
const (
    HNBCellAccessModeclosed = 0
    HNBCellAccessModehybrid = 1
    HNBCellAccessModeopen = 2

    /* Extensions */
)
func (self *HNBCellAccessMode) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *HNBCellAccessMode) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type HNBCellIdentifier struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'CellIdentity', 'name': 'cellIdentity'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNB-Cell-Identifier-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNidentity PLMNidentity
    CellIdentity CellIdentity
    IEExtensions *HNBCellIdentifierExtIEs
}

func (self * HNBCellIdentifier) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PLMNidentity.Unpack(stream)// p8
    self.CellIdentity.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &HNBCellIdentifierExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNB-Cell-Identifier-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_HNBCellIdentifierExtIEs, order_HNBCellIdentifierExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBCellIdentifier) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNidentity.Pack(stream)
    self.CellIdentity.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_HNBCellIdentifierExtIEs, order_HNBCellIdentifierExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *HNBRNLIdentity)Unpack(stream *Stream) {
    //coptions := []string{"hNB-Identity-as-Cell-Identifier"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in HNBRNLIdentity\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.HNBIdentityasCellIdentifier = &HNBCellIdentifier{}//cho6
        self.HNBIdentityasCellIdentifier.Unpack(stream)
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * HNBRNLIdentity) Pack(stream *Stream) {
    if self.HNBIdentityasCellIdentifier != nil {
        stream.set_choice(0, 0, 1, 1)
        self.HNBIdentityasCellIdentifier.Pack(stream)//2
    }

}
type HNBRNLIdentity struct { //[{'type': 'HNB-Cell-Identifier', 'name': 'hNB-Identity-as-Cell-Identifier'}, None]
    HNBIdentityasCellIdentifier *HNBCellIdentifier
} // HNBRNLIdentity

type HNBConfigInfo struct { // [{'type': 'HNB-RNL-Identity', 'name': 'hnb-RNL-Identity'}, {'type': 'PSC', 'name': 'psc', 'optional': True}, {'type': 'CSG-ID', 'name': 'cSG-ID', 'optional': True}, {'type': 'HNB-Cell-Access-Mode', 'name': 'hNB-Cell-Access-Mode'}, {'type': 'IP-Address', 'name': 'local-Iurh-IP-Address', 'optional': True}, {'type': 'IP-Address', 'name': 'remote-Iurh-IP-Address', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBConfigInfo-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    HnbRNLIdentity HNBRNLIdentity
    Psc *PSC
    CSGID *CSGID
    HNBCellAccessMode HNBCellAccessMode
    LocalIurhIPAddress *IPAddress
    RemoteIurhIPAddress *IPAddress
    IEExtensions *HNBConfigInfoExtIEs
}

func (self * HNBConfigInfo) Unpack(stream *Stream) {
    psc_flag := 0x00000002
    cSGID_flag := 0x00000004
    localIurhIPAddress_flag := 0x00000008
    remoteIurhIPAddress_flag := 0x00000010
    iEExtensions_flag := 0x00000020
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(6)
    self.HnbRNLIdentity.Unpack(stream)// p8
    if (psc_flag & _flags) == psc_flag { //cond2
        self.Psc = &PSC{}//7{'type': 'PSC', 'name': 'psc', 'optional': True}
        self.Psc.Unpack(stream)// p8
    }
    if (cSGID_flag & _flags) == cSGID_flag { //cond2
        self.CSGID = &CSGID{}//7{'type': 'CSG-ID', 'name': 'cSG-ID', 'optional': True}
        self.CSGID.Unpack(stream)// p8
    }
    self.HNBCellAccessMode.Unpack(stream)// p8
    if (localIurhIPAddress_flag & _flags) == localIurhIPAddress_flag { //cond2
        self.LocalIurhIPAddress = &IPAddress{}//7{'type': 'IP-Address', 'name': 'local-Iurh-IP-Address', 'optional': True}
        self.LocalIurhIPAddress.Unpack(stream)// p8
    }
    if (remoteIurhIPAddress_flag & _flags) == remoteIurhIPAddress_flag { //cond2
        self.RemoteIurhIPAddress = &IPAddress{}//7{'type': 'IP-Address', 'name': 'remote-Iurh-IP-Address', 'optional': True}
        self.RemoteIurhIPAddress.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &HNBConfigInfoExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNBConfigInfo-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_HNBConfigInfoExtIEs, order_HNBConfigInfoExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBConfigInfo) Pack(stream *Stream) {
    const psc_flag uint = 0x00000002
    const cSGID_flag uint = 0x00000004
    const localIurhIPAddress_flag uint = 0x00000008
    const remoteIurhIPAddress_flag uint = 0x00000010
    const iEExtensions_flag uint = 0x00000020
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(6)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.HnbRNLIdentity.Pack(stream)
    if self.Psc != nil { 
        _flags |= psc_flag
        self.Psc.Pack(stream)
    }//end of optional
    if self.CSGID != nil { 
        _flags |= cSGID_flag
        self.CSGID.Pack(stream)
    }//end of optional
    self.HNBCellAccessMode.Pack(stream)
    if self.LocalIurhIPAddress != nil { 
        _flags |= localIurhIPAddress_flag
        self.LocalIurhIPAddress.Pack(stream)
    }//end of optional
    if self.RemoteIurhIPAddress != nil { 
        _flags |= remoteIurhIPAddress_flag
        self.RemoteIurhIPAddress.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_HNBConfigInfoExtIEs, order_HNBConfigInfoExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

type HNBLocationInformation struct { // [{'type': 'MacroCoverageInformation', 'name': 'macroCoverageInfo', 'optional': True}, {'type': 'GeographicalLocation', 'name': 'geographicalCoordinates', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNB-Location-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    MacroCoverageInfo *MacroCoverageInformation
    GeographicalCoordinates *GeographicalLocation
    IEExtensions *HNBLocationInformationExtIEs
}

func (self * HNBLocationInformation) Unpack(stream *Stream) {
    macroCoverageInfo_flag := 0x00000002
    geographicalCoordinates_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (macroCoverageInfo_flag & _flags) == macroCoverageInfo_flag { //cond2
        self.MacroCoverageInfo = &MacroCoverageInformation{}//7{'type': 'MacroCoverageInformation', 'name': 'macroCoverageInfo', 'optional': True}
        self.MacroCoverageInfo.Unpack(stream)// p8
    }
    if (geographicalCoordinates_flag & _flags) == geographicalCoordinates_flag { //cond2
        self.GeographicalCoordinates = &GeographicalLocation{}//7{'type': 'GeographicalLocation', 'name': 'geographicalCoordinates', 'optional': True}
        self.GeographicalCoordinates.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &HNBLocationInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNB-Location-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_HNBLocationInformationExtIEs, order_HNBLocationInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBLocationInformation) Pack(stream *Stream) {
    const macroCoverageInfo_flag uint = 0x00000002
    const geographicalCoordinates_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.MacroCoverageInfo != nil { 
        _flags |= macroCoverageInfo_flag
        self.MacroCoverageInfo.Pack(stream)
    }//end of optional
    if self.GeographicalCoordinates != nil { 
        _flags |= geographicalCoordinates_flag
        self.GeographicalCoordinates.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_HNBLocationInformationExtIEs, order_HNBLocationInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type HNBIdentity struct { // [{'type': 'HNB-Identity-Info', 'name': 'hNB-Identity-Info'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNB-Identity-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    HNBIdentityInfo HNBIdentityInfo
    IEExtensions *HNBIdentityExtIEs
}

func (self * HNBIdentity) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.HNBIdentityInfo.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &HNBIdentityExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['HNB-Identity-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_HNBIdentityExtIEs, order_HNBIdentityExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * HNBIdentity) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.HNBIdentityInfo.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_HNBIdentityExtIEs, order_HNBIdentityExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type HNBIdentityInfo struct {
  Value HexBytes
}
func (self *HNBIdentityInfo) Unpack(st *Stream) {
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 255 {
        //fmt.Println ("Invalid len in HNB-Identity-Info")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *HNBIdentityInfo) Pack(st *Stream) {
    if len(self.Value) < 1 || len(self.Value) > 255 {
        log.Error ("Invalid len in HNB-Identity-Info")
        return
}
    st.format_olen((len(self.Value))-1, 8)
    st.formatf_OctString(self.Value, 0)
}
type IMEI struct {
  Len int
  Value HexBytes
}
func (self *IMEI) Unpack(st *Stream){
    self.Value = st.parsef_BitString(60, 60)
}
func (self *IMEI) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 60)
}
type IMSI struct {
  Value HexBytes
}
func (self *IMSI) Unpack(st *Stream) {
    _len := st.parse_olen(3)+3
    if _len < 3 || _len > 8 {
        //fmt.Println ("Invalid len in IMSI")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *IMSI) Pack(st *Stream) {
    if len(self.Value) < 3 || len(self.Value) > 8 {
        log.Error ("Invalid len in IMSI")
        return
}
    st.format_olen((len(self.Value))-3, 3)
    st.formatf_OctString(self.Value, 0)
}
type IMSIDS41 struct {
  Value HexBytes
}
func (self *IMSIDS41) Unpack(st *Stream) {
    _len := st.parse_olen(2)+5
    if _len < 5 || _len > 7 {
        //fmt.Println ("Invalid len in IMSIDS41")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *IMSIDS41) Pack(st *Stream) {
    if len(self.Value) < 5 || len(self.Value) > 7 {
        log.Error ("Invalid len in IMSIDS41")
        return
}
    st.format_olen((len(self.Value))-5, 2)
    st.formatf_OctString(self.Value, 0)
}
type IMSIESN struct { // [{'type': 'IMSIDS41', 'name': 'iMSIDS41'}, {'type': 'ESN', 'name': 'eSN'}]
    IMSIDS41 IMSIDS41
    ESN ESN
}

func (self * IMSIESN) Unpack(stream *Stream) {
    self.IMSIDS41.Unpack(stream)// p8
    self.ESN.Unpack(stream)// p8
    return
}

func (self * IMSIESN) Pack(stream *Stream) {
    self.IMSIDS41.Pack(stream)
    self.ESN.Pack(stream)
}//end

type IPAddress_Ipaddress struct { //[{'type': 'Ipv4Address', 'name': 'ipv4info'}, {'type': 'Ipv6Address', 'name': 'ipv6info'}, None]
    Ipv4info *Ipv4Address
    Ipv6info *Ipv6Address
} // IPAddress_Ipaddress

type IPAddress struct { // [{'type': 'CHOICE', 'members': [{'type': 'Ipv4Address', 'name': 'ipv4info'}, {'type': 'Ipv6Address', 'name': 'ipv6info'}, None], 'name': 'ipaddress'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['IP-Address-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    Ipaddress IPAddress_Ipaddress
    IEExtensions *IPAddressExtIEs
}

func (self * IPAddress) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_ipaddress = func(stream *Stream, self *IPAddress_Ipaddress) {
        //coptions := []string{"ipv4info","ipv6info"}
        choice := stream.get_choice(1, 1, 2)
        choice_len := 0
        choice_loc := 0
        if choice >= 2 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in IPAddress_Ipaddress\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.Ipv4info = &Ipv4Address{}//cho6
            self.Ipv4info.Unpack(stream)
        } else if choice == 1 { //ch2
            self.Ipv6info = &Ipv6Address{}//cho6
            self.Ipv6info.Unpack(stream)
        }//end of if else

        if choice >= 2 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_ipaddress(stream, &self.Ipaddress)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &IPAddressExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['IP-Address-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_IPAddressExtIEs, order_IPAddressExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * IPAddress) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ipaddress = func(stream *Stream, self IPAddress_Ipaddress) {
        if self.Ipv4info != nil {
            stream.set_choice(0, 1, 1, 2)
            self.Ipv4info.Pack(stream)//2
        } else if self.Ipv6info != nil {
            stream.set_choice(1, 1, 1, 2)
            self.Ipv6info.Pack(stream)//2
        }

    }
    Pack_ipaddress(stream, self.Ipaddress) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_IPAddressExtIEs, order_IPAddressExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type Ipv4Address struct {
  Value HexBytes
}
func (self *Ipv4Address) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(4)
}
func (self *Ipv4Address) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 4)
}
type Ipv6Address struct {
  Value HexBytes
}
func (self *Ipv6Address) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(16)
}
func (self *Ipv6Address) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 16)
}
type LAC struct {
  Value HexBytes
}
func (self *LAC) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(2)
}
func (self *LAC) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 2)
}
type HNBHBInfo struct {
  Value HexBytes
}
func (self *HNBHBInfo) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(16)
}
func (self *HNBHBInfo) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 16)
}
type LAI struct { // [{'type': 'PLMNidentity', 'name': 'pLMNID'}, {'type': 'LAC', 'name': 'lAC'}, None]
    PLMNID PLMNidentity
    LAC LAC
}

func (self * LAI) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PLMNID.Unpack(stream)// p8
    self.LAC.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * LAI) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNID.Pack(stream)
    self.LAC.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MacroCoverageInformation struct { // [{'type': 'MacroCellID', 'name': 'cellIdentity'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MacroCoverageInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    CellIdentity MacroCellID
    IEExtensions *MacroCoverageInformationExtIEs
}

func (self * MacroCoverageInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.CellIdentity.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MacroCoverageInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MacroCoverageInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MacroCoverageInformationExtIEs, order_MacroCoverageInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MacroCoverageInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CellIdentity.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MacroCoverageInformationExtIEs, order_MacroCoverageInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *MacroCellID)Unpack(stream *Stream) {
    //coptions := []string{"uTRANCellID","gERANCellID"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in MacroCellID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.UTRANCellID = &UTRANCellID{}//cho6
        self.UTRANCellID.Unpack(stream)
    } else if choice == 1 { //ch2
        self.GERANCellID = &CGI{}//cho6
        self.GERANCellID.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * MacroCellID) Pack(stream *Stream) {
    if self.UTRANCellID != nil {
        stream.set_choice(0, 1, 1, 2)
        self.UTRANCellID.Pack(stream)//2
    } else if self.GERANCellID != nil {
        stream.set_choice(1, 1, 1, 2)
        self.GERANCellID.Pack(stream)//2
    }

}
type MacroCellID struct { //[{'type': 'UTRANCellID', 'name': 'uTRANCellID'}, {'type': 'CGI', 'name': 'gERANCellID'}, None]
    UTRANCellID *UTRANCellID
    GERANCellID *CGI
} // MacroCellID

type MuxPortNumber struct {
  Value uint64
}
func (self *MuxPortNumber) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(64512, 16, 0, 1024)
}
func (self * MuxPortNumber) Pack(st *Stream){
    st.formatf_Integer(self.Value, 64512, 16, 0, 1024)
}
func (self *NeighbourInfoList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]HNBConfigInfo, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *NeighbourInfoList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type NeighbourInfoList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'HNBConfigInfo'}, 'size': [(1, 'maxnoofNeighbours')]}
    Items []HNBConfigInfo
}

func (self *NeighbourInfoRequestList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]HNBRNLIdentity, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *NeighbourInfoRequestList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type NeighbourInfoRequestList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'HNB-RNL-Identity'}, 'size': [(1, 'maxnoofNeighbours')]}
    Items []HNBRNLIdentity
}

type PLMNidentity struct {
  Value HexBytes
}
func (self *PLMNidentity) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(3)
}
func (self *PLMNidentity) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 3)
}
type PSC struct {
  Len int
  Value HexBytes
}
func (self *PSC) Unpack(st *Stream){
    self.Value = st.parsef_BitString(9, 9)
}
func (self *PSC) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 9)
}
type PTMSI struct {
  Len int
  Value HexBytes
}
func (self *PTMSI) Unpack(st *Stream){
    self.Value = st.parsef_BitString(32, 32)
}
func (self *PTMSI) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 32)
}
type PTMSIRAI struct { // [{'type': 'PTMSI', 'name': 'pTMSI'}, {'type': 'RAI', 'name': 'rAI'}, None]
    PTMSI PTMSI
    RAI RAI
}

func (self * PTMSIRAI) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PTMSI.Unpack(stream)// p8
    self.RAI.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PTMSIRAI) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PTMSI.Pack(stream)
    self.RAI.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *RABList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]RABListItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *RABList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type RABList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RABListItem'}, 'size': [(1, 'maxnoofRABs')]}
    Items []RABListItem
}

type RABListItem struct { // [{'type': 'TransportInfo', 'name': 'old-transport-Info'}, {'type': 'TransportInfo', 'name': 'new-transport-Info'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RABListItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    OldtransportInfo TransportInfo
    NewtransportInfo TransportInfo
    IEExtensions *RABListItemExtIEs
}

func (self * RABListItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.OldtransportInfo.Unpack(stream)// p8
    self.NewtransportInfo.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABListItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RABListItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABListItemExtIEs, order_RABListItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABListItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.OldtransportInfo.Pack(stream)
    self.NewtransportInfo.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABListItemExtIEs, order_RABListItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RAC struct {
  Value HexBytes
}
func (self *RAC) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(1)
}
func (self *RAC) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 1)
}
type RAI struct { // [{'type': 'LAI', 'name': 'lAI'}, {'type': 'RAC', 'name': 'rAC'}, None]
    LAI LAI
    RAC RAC
}

func (self * RAI) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.LAI.Unpack(stream)// p8
    self.RAC.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RAI) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.LAI.Pack(stream)
    self.RAC.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RegistrationCause struct {
  Value int
}
const (
    RegistrationCauseemergency_call = 0
    RegistrationCausenormal = 1

    /* Extensions */
    RegistrationCauseue_relocation = 2
)
func (self *RegistrationCause) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *RegistrationCause) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type RNCID struct {
  Value uint64
}
func (self *RNCID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * RNCID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type SAC struct {
  Value HexBytes
}
func (self *SAC) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(2)
}
func (self *SAC) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 2)
}
type TMSILAI struct { // [{'type': 'BIT STRING', 'size': [32], 'name': 'tMSI'}, {'type': 'LAI', 'name': 'lAI'}]
    TMSI BITSTRING
    LAI LAI
}

func (self * TMSILAI) Unpack(stream *Stream) {
    var Unpack_tMSI = func(st *Stream, self *BITSTRING){
        self.Value = st.parsef_BitString(32, 32)
    }
    Unpack_tMSI(stream, &self.TMSI)// p2
    self.LAI.Unpack(stream)// p8
    return
}

func (self * TMSILAI) Pack(stream *Stream) {
    var Pack_tMSI = func(st *Stream, self BITSTRING) {
        st.formatf_BitString(self.Value, 32)
    }
    Pack_tMSI(stream, self.TMSI) //f2
    self.LAI.Pack(stream)
}//end

type TMSIDS41 struct {
  Value HexBytes
}
func (self *TMSIDS41) Unpack(st *Stream) {
    _len := st.parse_olen(4)+2
    if _len < 2 || _len > 17 {
        //fmt.Println ("Invalid len in TMSIDS41")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *TMSIDS41) Pack(st *Stream) {
    if len(self.Value) < 2 || len(self.Value) > 17 {
        log.Error ("Invalid len in TMSIDS41")
        return
}
    st.format_olen((len(self.Value))-2, 4)
    st.formatf_OctString(self.Value, 0)
}
type TransportInfo_TransportAssociation struct { //[{'type': 'GTP-TEI', 'name': 'gtp-TEI'}, {'type': 'BindingID', 'name': 'bindingID'}, None]
    GtpTEI *GTPTEI
    BindingID *BindingID
} // TransportInfo_TransportAssociation

type TransportInfo struct { // [{'type': 'TransportLayerAddress', 'name': 'transportLayerAddress'}, {'type': 'CHOICE', 'members': [{'type': 'GTP-TEI', 'name': 'gtp-TEI'}, {'type': 'BindingID', 'name': 'bindingID'}, None], 'name': 'transportAssociation'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TransportInfo-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    TransportLayerAddress TransportLayerAddress
    TransportAssociation TransportInfo_TransportAssociation
    IEExtensions *TransportInfoExtIEs
}

func (self * TransportInfo) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.TransportLayerAddress.Unpack(stream)// p8
    var Unpack_transportAssociation = func(stream *Stream, self *TransportInfo_TransportAssociation) {
        //coptions := []string{"gtp-TEI","bindingID"}
        choice := stream.get_choice(1, 1, 2)
        choice_len := 0
        choice_loc := 0
        if choice >= 2 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in TransportInfo_TransportAssociation\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.GtpTEI = &GTPTEI{}//cho6
            self.GtpTEI.Unpack(stream)
        } else if choice == 1 { //ch2
            self.BindingID = &BindingID{}//cho6
            self.BindingID.Unpack(stream)
        }//end of if else

        if choice >= 2 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_transportAssociation(stream, &self.TransportAssociation)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &TransportInfoExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TransportInfo-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_TransportInfoExtIEs, order_TransportInfoExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TransportInfo) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TransportLayerAddress.Pack(stream)
    var Pack_transportAssociation = func(stream *Stream, self TransportInfo_TransportAssociation) {
        if self.GtpTEI != nil {
            stream.set_choice(0, 1, 1, 2)
            self.GtpTEI.Pack(stream)//2
        } else if self.BindingID != nil {
            stream.set_choice(1, 1, 1, 2)
            self.BindingID.Pack(stream)//2
        }

    }
    Pack_transportAssociation(stream, self.TransportAssociation) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_TransportInfoExtIEs, order_TransportInfoExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

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
type UECapabilities struct { // [{'type': 'Access-stratum-release-indicator', 'name': 'access-stratum-release-indicator'}, {'type': 'CSG-Indicator', 'name': 'csg-indicator'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UE-Capabilities-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    Accessstratumreleaseindicator Accessstratumreleaseindicator
    Csgindicator CSGIndicator
    IEExtensions *UECapabilitiesExtIEs
}

func (self * UECapabilities) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.Accessstratumreleaseindicator.Unpack(stream)// p8
    self.Csgindicator.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &UECapabilitiesExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UE-Capabilities-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_UECapabilitiesExtIEs, order_UECapabilitiesExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UECapabilities) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.Accessstratumreleaseindicator.Pack(stream)
    self.Csgindicator.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_UECapabilitiesExtIEs, order_UECapabilitiesExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UTRANCellID struct { // [{'type': 'LAC', 'name': 'lAC'}, {'type': 'RAC', 'name': 'rAC'}, {'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'CellIdentity', 'name': 'uTRANcellID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UTRANCellID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    LAC LAC
    RAC RAC
    PLMNidentity PLMNidentity
    UTRANcellID CellIdentity
    IEExtensions *UTRANCellIDExtIEs
}

func (self * UTRANCellID) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.LAC.Unpack(stream)// p8
    self.RAC.Unpack(stream)// p8
    self.PLMNidentity.Unpack(stream)// p8
    self.UTRANcellID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &UTRANCellIDExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UTRANCellID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_UTRANCellIDExtIEs, order_UTRANCellIDExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * UTRANCellID) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.LAC.Pack(stream)
    self.RAC.Pack(stream)
    self.PLMNidentity.Pack(stream)
    self.UTRANcellID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_UTRANCellIDExtIEs, order_UTRANCellIDExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *UEIdentity)Unpack(stream *Stream) {
    //coptions := []string{"iMSI","tMSILAI","pTMSIRAI","iMEI","eSN","iMSIDS41","iMSIESN","tMSIDS41"}
    choice := stream.get_choice(3, 1, 8)
    choice_len := 0
    choice_loc := 0
    if choice >= 8 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in UEIdentity\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.IMSI = &IMSI{}//cho6
        self.IMSI.Unpack(stream)
    } else if choice == 1 { //ch2
        self.TMSILAI = &TMSILAI{}//cho6
        self.TMSILAI.Unpack(stream)
    } else if choice == 2 { //ch2
        self.PTMSIRAI = &PTMSIRAI{}//cho6
        self.PTMSIRAI.Unpack(stream)
    } else if choice == 3 { //ch2
        self.IMEI = &IMEI{}//cho6
        self.IMEI.Unpack(stream)
    } else if choice == 4 { //ch2
        self.ESN = &ESN{}//cho6
        self.ESN.Unpack(stream)
    } else if choice == 5 { //ch2
        self.IMSIDS41 = &IMSIDS41{}//cho6
        self.IMSIDS41.Unpack(stream)
    } else if choice == 6 { //ch2
        self.IMSIESN = &IMSIESN{}//cho6
        self.IMSIESN.Unpack(stream)
    } else if choice == 7 { //ch2
        self.TMSIDS41 = &TMSIDS41{}//cho6
        self.TMSIDS41.Unpack(stream)
    }//end of if else

    if choice >= 8 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * UEIdentity) Pack(stream *Stream) {
    if self.IMSI != nil {
        stream.set_choice(0, 3, 1, 8)
        self.IMSI.Pack(stream)//2
    } else if self.TMSILAI != nil {
        stream.set_choice(1, 3, 1, 8)
        self.TMSILAI.Pack(stream)//2
    } else if self.PTMSIRAI != nil {
        stream.set_choice(2, 3, 1, 8)
        self.PTMSIRAI.Pack(stream)//2
    } else if self.IMEI != nil {
        stream.set_choice(3, 3, 1, 8)
        self.IMEI.Pack(stream)//2
    } else if self.ESN != nil {
        stream.set_choice(4, 3, 1, 8)
        self.ESN.Pack(stream)//2
    } else if self.IMSIDS41 != nil {
        stream.set_choice(5, 3, 1, 8)
        self.IMSIDS41.Pack(stream)//2
    } else if self.IMSIESN != nil {
        stream.set_choice(6, 3, 1, 8)
        self.IMSIESN.Pack(stream)//2
    } else if self.TMSIDS41 != nil {
        stream.set_choice(7, 3, 1, 8)
        self.TMSIDS41.Pack(stream)//2
    }

}
type UEIdentity struct { //[{'type': 'IMSI', 'name': 'iMSI'}, {'type': 'TMSILAI', 'name': 'tMSILAI'}, {'type': 'PTMSIRAI', 'name': 'pTMSIRAI'}, {'type': 'IMEI', 'name': 'iMEI'}, {'type': 'ESN', 'name': 'eSN'}, {'type': 'IMSIDS41', 'name': 'iMSIDS41'}, {'type': 'IMSIESN', 'name': 'iMSIESN'}, {'type': 'TMSIDS41', 'name': 'tMSIDS41'}, None]
    IMSI *IMSI
    TMSILAI *TMSILAI
    PTMSIRAI *PTMSIRAI
    IMEI *IMEI
    ESN *ESN
    IMSIDS41 *IMSIDS41
    IMSIESN *IMSIESN
    TMSIDS41 *TMSIDS41
} // UEIdentity

type Updatecause struct {
  Value int
}
const (
    Updatecauserelocation_preparation = 0

    /* Extensions */
)
func (self *Updatecause) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *Updatecause) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
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
type PrivateIEID struct { //[{'type': 'INTEGER', 'restricted-to': [(0, 65535)], 'name': 'local'}, {'type': 'OBJECT IDENTIFIER', 'name': 'global'}]
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
    _size := data.(HNBAPPROTOCOLIES_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['HNBAP-PROTOCOL-IES']}
    Items map[int]*HNBAPPROTOCOLIES
    order []int
}

type ProtocolIESingleContainer struct{ //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['HNBAP-PROTOCOL-IES']}
   Item map[int]*HNBAPPROTOCOLIES //UserType
   order []int
}
func (self *ProtocolIESingleContainer) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['HNBAP-PROTOCOL-IES']}
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

type ProtocolIEField struct { // [{'type': 'HNBAP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'HNBAP-PROTOCOL-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'HNBAP-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id HNBAPPROTOCOLIESid
    Criticality HNBAPPROTOCOLIEScriticality
    Value HNBAPPROTOCOLIESValue
}

func (self * ProtocolIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'HNBAP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := HNBAPPROTOCOLIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(HNBAPPROTOCOLIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'HNBAP-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * HNBAPPROTOCOLIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'HNBAP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (HNBAPPROTOCOLIESid)(self.ID)
    if out.(HNBAPPROTOCOLIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(HNBAPPROTOCOLIES_IF).PackOT(stream, key)
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


type ProtocolIEContainerList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['IEsSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IEsSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'HNBAP-PROTOCOL-IES']}
    Items map[int]*HNBAPPROTOCOLIES
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
    _size := data.(HNBAPPROTOCOLEXTENSION_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolExtensionContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolExtensionField', 'actual-parameters': ['ExtensionSetParam']}, 'size': [(1, 'maxProtocolExtensions')], 'parameters': ['ExtensionSetParam'], 'param-types': ['HNBAP-PROTOCOL-EXTENSION']}
    Items map[int]*HNBAPPROTOCOLEXTENSION
    order []int
}

type ProtocolExtensionField struct { // [{'type': 'HNBAP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}, {'type': 'HNBAP-PROTOCOL-EXTENSION.&criticality', 'table': ['ExtensionSetParam', ['id']], 'name': 'criticality'}, {'type': 'HNBAP-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}]
    Id HNBAPPROTOCOLEXTENSIONid
    Criticality HNBAPPROTOCOLEXTENSIONcriticality
    ExtensionValue HNBAPPROTOCOLEXTENSIONExtension
}

func (self * ProtocolExtensionField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'HNBAP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := HNBAPPROTOCOLEXTENSION{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(HNBAPPROTOCOLEXTENSION_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'HNBAP-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}
    stream.set_location(location, _len)
    return
}

func (self * HNBAPPROTOCOLEXTENSION) Pack(stream *Stream, out interface{}) {
    //table {'type': 'HNBAP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    key := (HNBAPPROTOCOLEXTENSIONid)(self.ID)
    if out.(HNBAPPROTOCOLEXTENSION_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(HNBAPPROTOCOLEXTENSION_IF).PackOT(stream, key)
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
    _size := data.(HNBAPPRIVATEIES_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type PrivateIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PrivateIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(1, 'maxPrivateIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['HNBAP-PRIVATE-IES']}
    Items map[int]*HNBAPPRIVATEIES
    order []int
}

type PrivateIEField struct { // [{'type': 'HNBAP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'HNBAP-PRIVATE-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'HNBAP-PRIVATE-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id HNBAPPRIVATEIESid
    Criticality HNBAPPRIVATEIEScriticality
    Value HNBAPPRIVATEIESValue
}

func (self * PrivateIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'HNBAP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := HNBAPPRIVATEIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(HNBAPPRIVATEIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'HNBAP-PRIVATE-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * HNBAPPRIVATEIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'HNBAP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (HNBAPPRIVATEIESid)(self.ID)
    if out.(HNBAPPRIVATEIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(HNBAPPRIVATEIES_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type HNBAPELEMENTARYPROCEDURE struct {//CLASS {'members': [{'type': 'OpenType', 'name': '&InitiatingMessage'}, {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, {'type': 'ProcedureCode', 'name': '&procedureCode'}, {'type': 'Criticality', 'name': '&criticality'}], 'with-members': [{'with-order': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME', 'PROCEDURE CODE', 'CRITICALITY'], 'with-type': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME']}], 'alt-type': {'INITIATING MESSAGE': {'type': 'OpenType', 'name': '&InitiatingMessage'}, 'SUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, 'UNSUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, 'PROCEDURE CODE': {'type': 'ProcedureCode', 'name': '&procedureCode'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}}, 'id-type': {'type': 'ProcedureCode', 'name': '&procedureCode'}}
    INITIATINGMESSAGE interface{}
    SUCCESSFULOUTCOME interface{}
    UNSUCCESSFULOUTCOME interface{}
    PROCEDURECODE ProcedureCode
    CRITICALITY Criticality
}
type HNBAPELEMENTARYPROCEDUREInitiatingMessage interface{}
type HNBAPELEMENTARYPROCEDURESuccessfulOutcome interface{}
type HNBAPELEMENTARYPROCEDUREUnsuccessfulOutcome interface{}
type HNBAPELEMENTARYPROCEDUREprocedureCode ProcedureCode
func (self *HNBAPELEMENTARYPROCEDUREprocedureCode) Pack(st *Stream) {
    ieID := ProcedureCode(*self)
    ieID.Pack(st)
}
func (self *HNBAPELEMENTARYPROCEDUREprocedureCode) Unpack(st *Stream) {
    val := ProcedureCode{}
    val.Unpack(st)
    *self = HNBAPELEMENTARYPROCEDUREprocedureCode(val)
}
type HNBAPELEMENTARYPROCEDUREcriticality Criticality
func (self *HNBAPELEMENTARYPROCEDUREcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *HNBAPELEMENTARYPROCEDUREcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = HNBAPELEMENTARYPROCEDUREcriticality(val)
}

type HNBAPELEMENTARYPROCEDURE_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type HNBAPPROTOCOLIES struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type HNBAPPROTOCOLIESid ProtocolIEID
func (self *HNBAPPROTOCOLIESid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *HNBAPPROTOCOLIESid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = HNBAPPROTOCOLIESid(val)
}
type HNBAPPROTOCOLIEScriticality Criticality
func (self *HNBAPPROTOCOLIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *HNBAPPROTOCOLIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = HNBAPPROTOCOLIEScriticality(val)
}
type HNBAPPROTOCOLIESValue interface{}
type HNBAPPROTOCOLIESpresence Presence
func (self *HNBAPPROTOCOLIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *HNBAPPROTOCOLIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = HNBAPPROTOCOLIESpresence(val)
}

type HNBAPPROTOCOLIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type HNBAPPROTOCOLEXTENSION struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Extension'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'EXTENSION', 'PRESENCE'], 'with-type': ['EXTENSION']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    EXTENSION interface{}
    PRESENCE Presence
}
type HNBAPPROTOCOLEXTENSIONid ProtocolIEID
func (self *HNBAPPROTOCOLEXTENSIONid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *HNBAPPROTOCOLEXTENSIONid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = HNBAPPROTOCOLEXTENSIONid(val)
}
type HNBAPPROTOCOLEXTENSIONcriticality Criticality
func (self *HNBAPPROTOCOLEXTENSIONcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *HNBAPPROTOCOLEXTENSIONcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = HNBAPPROTOCOLEXTENSIONcriticality(val)
}
type HNBAPPROTOCOLEXTENSIONExtension interface{}
type HNBAPPROTOCOLEXTENSIONpresence Presence
func (self *HNBAPPROTOCOLEXTENSIONpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *HNBAPPROTOCOLEXTENSIONpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = HNBAPPROTOCOLEXTENSIONpresence(val)
}

type HNBAPPROTOCOLEXTENSION_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type HNBAPPRIVATEIES struct {//CLASS {'members': [{'type': 'PrivateIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'PrivateIE-ID', 'name': '&id'}}
    ID PrivateIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type HNBAPPRIVATEIESid PrivateIEID
func (self *HNBAPPRIVATEIESid) Pack(st *Stream) {
    ieID := PrivateIEID(*self)
    ieID.Pack(st)
}
func (self *HNBAPPRIVATEIESid) Unpack(st *Stream) {
    val := PrivateIEID{}
    val.Unpack(st)
    *self = HNBAPPRIVATEIESid(val)
}
type HNBAPPRIVATEIEScriticality Criticality
func (self *HNBAPPRIVATEIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *HNBAPPRIVATEIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = HNBAPPRIVATEIEScriticality(val)
}
type HNBAPPRIVATEIESValue interface{}
type HNBAPPRIVATEIESpresence Presence
func (self *HNBAPPRIVATEIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *HNBAPPRIVATEIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = HNBAPPRIVATEIESpresence(val)
}

type HNBAPPRIVATEIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
//class HNBAPELEMENTARYPROCEDURES: #OBJSET1 {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'members': [{}, {}, None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_HNBAPELEMENTARYPROCEDURES = make(map[HNBAPELEMENTARYPROCEDUREprocedureCode]*HNBAPELEMENTARYPROCEDURE)


//class HNBAPELEMENTARYPROCEDURESCLASS1: #OBJSET1 {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_HNBAPELEMENTARYPROCEDURESCLASS1 = make(map[HNBAPELEMENTARYPROCEDUREprocedureCode]*HNBAPELEMENTARYPROCEDURE)


//class HNBAPELEMENTARYPROCEDURESCLASS2: #OBJSET1 {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_HNBAPELEMENTARYPROCEDURESCLASS2 = make(map[HNBAPELEMENTARYPROCEDUREprocedureCode]*HNBAPELEMENTARYPROCEDURE)


type HNBRegisterRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-HNB-Identity', 'CRITICALITY': 'reject', 'TYPE': 'HNB-Identity', 'PRESENCE': 'mandatory'}, {'ID': 'id-HNB-Location-Information', 'CRITICALITY': 'reject', 'TYPE': 'HNB-Location-Information', 'PRESENCE': 'mandatory'}, {'ID': 'id-PLMNidentity', 'CRITICALITY': 'reject', 'TYPE': 'PLMNidentity', 'PRESENCE': 'mandatory'}, {'ID': 'id-CellIdentity', 'CRITICALITY': 'reject', 'TYPE': 'CellIdentity', 'PRESENCE': 'mandatory'}, {'ID': 'id-LAC', 'CRITICALITY': 'reject', 'TYPE': 'LAC', 'PRESENCE': 'mandatory'}, {'ID': 'id-RAC', 'CRITICALITY': 'reject', 'TYPE': 'RAC', 'PRESENCE': 'mandatory'}, {'ID': 'id-SAC', 'CRITICALITY': 'reject', 'TYPE': 'SAC', 'PRESENCE': 'mandatory'}, {'ID': 'id-CSG-ID', 'CRITICALITY': 'reject', 'TYPE': 'CSG-ID', 'PRESENCE': 'optional'}, None]}
   HNBIdentity  HNBIdentity
   HNBLocationInformation  HNBLocationInformation
   PLMNidentity  PLMNidentity
   CellIdentity  CellIdentity
   LAC  LAC
   RAC  RAC
   SAC  SAC
   CSGID  *CSGID
   list []interface{}
}
func (self *HNBRegisterRequestIEs)createOT() interface{}{
    return nil
}
var table_HNBRegisterRequestIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_HNBRegisterRequestIEs = make([]int, 8)

func (self *HNBRegisterRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.HNBIdentity
   count +=1 //self.HNBLocationInformation
   count +=1 //self.PLMNidentity
   count +=1 //self.CellIdentity
   count +=1 //self.LAC
   count +=1 //self.RAC
   count +=1 //self.SAC
   if self.CSGID != nil { count += 1 }
   return count//ObjSet
}
func (self *HNBRegisterRequestIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 3: //HNBIdentity
        return true //self.HNBIdentity
      case 8: //HNBLocationInformation
        return true //self.HNBLocationInformation
      case 9: //PLMNidentity
        return true //self.PLMNidentity
      case 11: //CellIdentity
        return true //self.CellIdentity
      case 6: //LAC
        return true //self.LAC
      case 7: //RAC
        return true //self.RAC
      case 10: //SAC
        return true //self.SAC
      case 15: //CSGID
        if self.CSGID != nil { return true }
   }
   return false//ObjSet
}
func (self *HNBRegisterRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 3: //HNBIdentity
        self.HNBIdentity.Unpack(st)
        self.list = append(self.list, &self.HNBIdentity)
      case 8: //HNBLocationInformation
        self.HNBLocationInformation.Unpack(st)
        self.list = append(self.list, &self.HNBLocationInformation)
      case 9: //PLMNidentity
        self.PLMNidentity.Unpack(st)
        self.list = append(self.list, &self.PLMNidentity)
      case 11: //CellIdentity
        self.CellIdentity.Unpack(st)
        self.list = append(self.list, &self.CellIdentity)
      case 6: //LAC
        self.LAC.Unpack(st)
        self.list = append(self.list, &self.LAC)
      case 7: //RAC
        self.RAC.Unpack(st)
        self.list = append(self.list, &self.RAC)
      case 10: //SAC
        self.SAC.Unpack(st)
        self.list = append(self.list, &self.SAC)
      case 15: //CSGID
        self.CSGID = &CSGID{}
        self.CSGID.Unpack(st)
        self.list = append(self.list, self.CSGID)
   }
}
func (self *HNBRegisterRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 3: //HNBIdentity
        self.HNBIdentity.Pack(st)
      case 8: //HNBLocationInformation
        self.HNBLocationInformation.Pack(st)
      case 9: //PLMNidentity
        self.PLMNidentity.Pack(st)
      case 11: //CellIdentity
        self.CellIdentity.Pack(st)
      case 6: //LAC
        self.LAC.Pack(st)
      case 7: //RAC
        self.RAC.Pack(st)
      case 10: //SAC
        self.SAC.Pack(st)
      case 15: //CSGID
        if self.CSGID != nil {self.CSGID.Pack(st)}
      default:
      break
   }
}
func init() {
table_HNBRegisterRequestIEs[3] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idHNBIdentity}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&HNBIdentity{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBRegisterRequestIEs[0] = 3
table_HNBRegisterRequestIEs[8] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idHNBLocationInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&HNBLocationInformation{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBRegisterRequestIEs[1] = 8
table_HNBRegisterRequestIEs[9] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idPLMNidentity}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&PLMNidentity{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBRegisterRequestIEs[2] = 9
table_HNBRegisterRequestIEs[11] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCellIdentity}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CellIdentity{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBRegisterRequestIEs[3] = 11
table_HNBRegisterRequestIEs[6] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idLAC}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&LAC{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBRegisterRequestIEs[4] = 6
table_HNBRegisterRequestIEs[7] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idRAC}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RAC{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBRegisterRequestIEs[5] = 7
table_HNBRegisterRequestIEs[10] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idSAC}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SAC{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBRegisterRequestIEs[6] = 10
table_HNBRegisterRequestIEs[15] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCSGID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CSGID{}, PRESENCE:Presence{Presenceoptional}, }
order_HNBRegisterRequestIEs[7] = 15
   }

type HNBRegisterRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-Service-Area-For-Broadcast', 'CRITICALITY': 'ignore', 'EXTENSION': 'SAC', 'PRESENCE': 'optional'}, {'ID': 'id-HNB-Cell-Access-Mode', 'CRITICALITY': 'reject', 'EXTENSION': 'HNB-Cell-Access-Mode', 'PRESENCE': 'optional'}, {'ID': 'id-PSC', 'CRITICALITY': 'ignore', 'EXTENSION': 'PSC', 'PRESENCE': 'optional'}, {'ID': 'id-Local-Iurh-IP-Address', 'CRITICALITY': 'ignore', 'EXTENSION': 'IP-Address', 'PRESENCE': 'optional'}, None]}
   ServiceAreaForBroadcast  *SAC
   HNBCellAccessMode  *HNBCellAccessMode
   PSC  *PSC
   LocalIurhIPAddress  *IPAddress
   list []interface{}
}
func (self *HNBRegisterRequestExtensions)createOT() interface{}{
    return nil
}
var table_HNBRegisterRequestExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBRegisterRequestExtensions = make([]int, 4)

func (self *HNBRegisterRequestExtensions) GetIECount() int{
   count := 0
   if self.ServiceAreaForBroadcast != nil { count += 1 }
   if self.HNBCellAccessMode != nil { count += 1 }
   if self.PSC != nil { count += 1 }
   if self.LocalIurhIPAddress != nil { count += 1 }
   return count//ObjSet
}
func (self *HNBRegisterRequestExtensions) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 20: //ServiceAreaForBroadcast
        if self.ServiceAreaForBroadcast != nil { return true }
      case 18: //HNBCellAccessMode
        if self.HNBCellAccessMode != nil { return true }
      case 30: //PSC
        if self.PSC != nil { return true }
      case 24: //LocalIurhIPAddress
        if self.LocalIurhIPAddress != nil { return true }
   }
   return false//ObjSet
}
func (self *HNBRegisterRequestExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 20: //ServiceAreaForBroadcast
        self.ServiceAreaForBroadcast = &SAC{}
        self.ServiceAreaForBroadcast.Unpack(st)
        self.list = append(self.list, self.ServiceAreaForBroadcast)
      case 18: //HNBCellAccessMode
        self.HNBCellAccessMode = &HNBCellAccessMode{}
        self.HNBCellAccessMode.Unpack(st)
        self.list = append(self.list, self.HNBCellAccessMode)
      case 30: //PSC
        self.PSC = &PSC{}
        self.PSC.Unpack(st)
        self.list = append(self.list, self.PSC)
      case 24: //LocalIurhIPAddress
        self.LocalIurhIPAddress = &IPAddress{}
        self.LocalIurhIPAddress.Unpack(st)
        self.list = append(self.list, self.LocalIurhIPAddress)
   }
}
func (self *HNBRegisterRequestExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 20: //ServiceAreaForBroadcast
        if self.ServiceAreaForBroadcast != nil {self.ServiceAreaForBroadcast.Pack(st)}
      case 18: //HNBCellAccessMode
        if self.HNBCellAccessMode != nil {self.HNBCellAccessMode.Pack(st)}
      case 30: //PSC
        if self.PSC != nil {self.PSC.Pack(st)}
      case 24: //LocalIurhIPAddress
        if self.LocalIurhIPAddress != nil {self.LocalIurhIPAddress.Pack(st)}
      default:
      break
   }
}
func init() {
table_HNBRegisterRequestExtensions[20] = &HNBAPPROTOCOLEXTENSION{ID:ProtocolIEID{idServiceAreaForBroadcast}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&SAC{}, PRESENCE:Presence{Presenceoptional}, }
order_HNBRegisterRequestExtensions[0] = 20
table_HNBRegisterRequestExtensions[18] = &HNBAPPROTOCOLEXTENSION{ID:ProtocolIEID{idHNBCellAccessMode}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&HNBCellAccessMode{}, PRESENCE:Presence{Presenceoptional}, }
order_HNBRegisterRequestExtensions[1] = 18
table_HNBRegisterRequestExtensions[30] = &HNBAPPROTOCOLEXTENSION{ID:ProtocolIEID{idPSC}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&PSC{}, PRESENCE:Presence{Presenceoptional}, }
order_HNBRegisterRequestExtensions[2] = 30
table_HNBRegisterRequestExtensions[24] = &HNBAPPROTOCOLEXTENSION{ID:ProtocolIEID{idLocalIurhIPAddress}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&IPAddress{}, PRESENCE:Presence{Presenceoptional}, }
order_HNBRegisterRequestExtensions[3] = 24
   }

type HNBRegisterResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-RNC-ID', 'CRITICALITY': 'reject', 'TYPE': 'RNC-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Remote-Iurh-IP-Address', 'CRITICALITY': 'ignore', 'TYPE': 'IP-Address', 'PRESENCE': 'optional'}, None]}
   RNCID  RNCID
   RemoteIurhIPAddress  *IPAddress
   list []interface{}
}
func (self *HNBRegisterResponseIEs)createOT() interface{}{
    return nil
}
var table_HNBRegisterResponseIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_HNBRegisterResponseIEs = make([]int, 2)

func (self *HNBRegisterResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.RNCID
   if self.RemoteIurhIPAddress != nil { count += 1 }
   return count//ObjSet
}
func (self *HNBRegisterResponseIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 14: //RNCID
        return true //self.RNCID
      case 29: //RemoteIurhIPAddress
        if self.RemoteIurhIPAddress != nil { return true }
   }
   return false//ObjSet
}
func (self *HNBRegisterResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 14: //RNCID
        self.RNCID.Unpack(st)
        self.list = append(self.list, &self.RNCID)
      case 29: //RemoteIurhIPAddress
        self.RemoteIurhIPAddress = &IPAddress{}
        self.RemoteIurhIPAddress.Unpack(st)
        self.list = append(self.list, self.RemoteIurhIPAddress)
   }
}
func (self *HNBRegisterResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 14: //RNCID
        self.RNCID.Pack(st)
      case 29: //RemoteIurhIPAddress
        if self.RemoteIurhIPAddress != nil {self.RemoteIurhIPAddress.Pack(st)}
      default:
      break
   }
}
func init() {
table_HNBRegisterResponseIEs[14] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idRNCID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RNCID{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBRegisterResponseIEs[0] = 14
table_HNBRegisterResponseIEs[29] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idRemoteIurhIPAddress}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&IPAddress{}, PRESENCE:Presence{Presenceoptional}, }
order_HNBRegisterResponseIEs[1] = 29
   }

type HNBRegisterResponseExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-MuxPortNumber', 'CRITICALITY': 'ignore', 'EXTENSION': 'MuxPortNumber', 'PRESENCE': 'optional'}, None]}
   MuxPortNumber  *MuxPortNumber
   list []interface{}
}
func (self *HNBRegisterResponseExtensions)createOT() interface{}{
    return nil
}
var table_HNBRegisterResponseExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBRegisterResponseExtensions = make([]int, 1)

func (self *HNBRegisterResponseExtensions) GetIECount() int{
   count := 0
   if self.MuxPortNumber != nil { count += 1 }
   return count//ObjSet
}
func (self *HNBRegisterResponseExtensions) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 19: //MuxPortNumber
        if self.MuxPortNumber != nil { return true }
   }
   return false//ObjSet
}
func (self *HNBRegisterResponseExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 19: //MuxPortNumber
        self.MuxPortNumber = &MuxPortNumber{}
        self.MuxPortNumber.Unpack(st)
        self.list = append(self.list, self.MuxPortNumber)
   }
}
func (self *HNBRegisterResponseExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 19: //MuxPortNumber
        if self.MuxPortNumber != nil {self.MuxPortNumber.Pack(st)}
      default:
      break
   }
}
func init() {
table_HNBRegisterResponseExtensions[19] = &HNBAPPROTOCOLEXTENSION{ID:ProtocolIEID{idMuxPortNumber}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&MuxPortNumber{}, PRESENCE:Presence{Presenceoptional}, }
order_HNBRegisterResponseExtensions[0] = 19
   }

type HNBRegisterRejectIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, {'ID': 'id-BackoffTimer', 'CRITICALITY': 'reject', 'TYPE': 'BackoffTimer', 'PRESENCE': 'conditional'}, None]}
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   BackoffTimer  BackoffTimer
   list []interface{}
}
func (self *HNBRegisterRejectIEs)createOT() interface{}{
    return nil
}
var table_HNBRegisterRejectIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_HNBRegisterRejectIEs = make([]int, 3)

func (self *HNBRegisterRejectIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   count +=1 //self.BackoffTimer
   return count//ObjSet
}
func (self *HNBRegisterRejectIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        return true //self.Cause
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
      case 16: //BackoffTimer
        return true //self.BackoffTimer
   }
   return false//ObjSet
}
func (self *HNBRegisterRejectIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
      case 16: //BackoffTimer
        self.BackoffTimer.Unpack(st)
        self.list = append(self.list, &self.BackoffTimer)
   }
}
func (self *HNBRegisterRejectIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        self.Cause.Pack(st)
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      case 16: //BackoffTimer
        self.BackoffTimer.Pack(st)
      default:
      break
   }
}
func init() {
table_HNBRegisterRejectIEs[1] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBRegisterRejectIEs[0] = 1
table_HNBRegisterRejectIEs[2] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_HNBRegisterRejectIEs[1] = 2
table_HNBRegisterRejectIEs[16] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idBackoffTimer}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&BackoffTimer{}, PRESENCE:Presence{Presenceconditional}, }
order_HNBRegisterRejectIEs[2] = 16
   }

type HNBRegisterRejectExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *HNBRegisterRejectExtensions)createOT() interface{}{
    return nil
}
var table_HNBRegisterRejectExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBRegisterRejectExtensions = make([]int, 0)

type HNBDeRegisterIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-BackoffTimer', 'CRITICALITY': 'reject', 'TYPE': 'BackoffTimer', 'PRESENCE': 'conditional'}, None]}
   Cause  Cause
   BackoffTimer  BackoffTimer
   list []interface{}
}
func (self *HNBDeRegisterIEs)createOT() interface{}{
    return nil
}
var table_HNBDeRegisterIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_HNBDeRegisterIEs = make([]int, 2)

func (self *HNBDeRegisterIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   count +=1 //self.BackoffTimer
   return count//ObjSet
}
func (self *HNBDeRegisterIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        return true //self.Cause
      case 16: //BackoffTimer
        return true //self.BackoffTimer
   }
   return false//ObjSet
}
func (self *HNBDeRegisterIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 16: //BackoffTimer
        self.BackoffTimer.Unpack(st)
        self.list = append(self.list, &self.BackoffTimer)
   }
}
func (self *HNBDeRegisterIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        self.Cause.Pack(st)
      case 16: //BackoffTimer
        self.BackoffTimer.Pack(st)
      default:
      break
   }
}
func init() {
table_HNBDeRegisterIEs[1] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBDeRegisterIEs[0] = 1
table_HNBDeRegisterIEs[16] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idBackoffTimer}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&BackoffTimer{}, PRESENCE:Presence{Presenceconditional}, }
order_HNBDeRegisterIEs[1] = 16
   }

type HNBDeRegisterExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *HNBDeRegisterExtensions)createOT() interface{}{
    return nil
}
var table_HNBDeRegisterExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBDeRegisterExtensions = make([]int, 0)

type UERegisterRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-UE-Identity', 'CRITICALITY': 'reject', 'TYPE': 'UE-Identity', 'PRESENCE': 'mandatory'}, {'ID': 'id-Registration-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Registration-Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-UE-Capabilities', 'CRITICALITY': 'reject', 'TYPE': 'UE-Capabilities', 'PRESENCE': 'mandatory'}, None]}
   UEIdentity  UEIdentity
   RegistrationCause  RegistrationCause
   UECapabilities  UECapabilities
   list []interface{}
}
func (self *UERegisterRequestIEs)createOT() interface{}{
    return nil
}
var table_UERegisterRequestIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_UERegisterRequestIEs = make([]int, 3)

func (self *UERegisterRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.UEIdentity
   count +=1 //self.RegistrationCause
   count +=1 //self.UECapabilities
   return count//ObjSet
}
func (self *UERegisterRequestIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        return true //self.UEIdentity
      case 12: //RegistrationCause
        return true //self.RegistrationCause
      case 13: //UECapabilities
        return true //self.UECapabilities
   }
   return false//ObjSet
}
func (self *UERegisterRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        self.UEIdentity.Unpack(st)
        self.list = append(self.list, &self.UEIdentity)
      case 12: //RegistrationCause
        self.RegistrationCause.Unpack(st)
        self.list = append(self.list, &self.RegistrationCause)
      case 13: //UECapabilities
        self.UECapabilities.Unpack(st)
        self.list = append(self.list, &self.UECapabilities)
   }
}
func (self *UERegisterRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        self.UEIdentity.Pack(st)
      case 12: //RegistrationCause
        self.RegistrationCause.Pack(st)
      case 13: //UECapabilities
        self.UECapabilities.Pack(st)
      default:
      break
   }
}
func init() {
table_UERegisterRequestIEs[5] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idUEIdentity}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&UEIdentity{}, PRESENCE:Presence{Presencemandatory}, }
order_UERegisterRequestIEs[0] = 5
table_UERegisterRequestIEs[12] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idRegistrationCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RegistrationCause{}, PRESENCE:Presence{Presencemandatory}, }
order_UERegisterRequestIEs[1] = 12
table_UERegisterRequestIEs[13] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idUECapabilities}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&UECapabilities{}, PRESENCE:Presence{Presencemandatory}, }
order_UERegisterRequestIEs[2] = 13
   }

type UERegisterRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UERegisterRequestExtensions)createOT() interface{}{
    return nil
}
var table_UERegisterRequestExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_UERegisterRequestExtensions = make([]int, 0)

type UERegisterAcceptIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-UE-Identity', 'CRITICALITY': 'reject', 'TYPE': 'UE-Identity', 'PRESENCE': 'mandatory'}, {'ID': 'id-Context-ID', 'CRITICALITY': 'reject', 'TYPE': 'Context-ID', 'PRESENCE': 'mandatory'}, None]}
   UEIdentity  UEIdentity
   ContextID  ContextID
   list []interface{}
}
func (self *UERegisterAcceptIEs)createOT() interface{}{
    return nil
}
var table_UERegisterAcceptIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_UERegisterAcceptIEs = make([]int, 2)

func (self *UERegisterAcceptIEs) GetIECount() int{
   count := 0
   count +=1 //self.UEIdentity
   count +=1 //self.ContextID
   return count//ObjSet
}
func (self *UERegisterAcceptIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        return true //self.UEIdentity
      case 4: //ContextID
        return true //self.ContextID
   }
   return false//ObjSet
}
func (self *UERegisterAcceptIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        self.UEIdentity.Unpack(st)
        self.list = append(self.list, &self.UEIdentity)
      case 4: //ContextID
        self.ContextID.Unpack(st)
        self.list = append(self.list, &self.ContextID)
   }
}
func (self *UERegisterAcceptIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        self.UEIdentity.Pack(st)
      case 4: //ContextID
        self.ContextID.Pack(st)
      default:
      break
   }
}
func init() {
table_UERegisterAcceptIEs[5] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idUEIdentity}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&UEIdentity{}, PRESENCE:Presence{Presencemandatory}, }
order_UERegisterAcceptIEs[0] = 5
table_UERegisterAcceptIEs[4] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idContextID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ContextID{}, PRESENCE:Presence{Presencemandatory}, }
order_UERegisterAcceptIEs[1] = 4
   }

type UERegisterAcceptExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-CSGMembershipStatus', 'CRITICALITY': 'reject', 'EXTENSION': 'CSGMembershipStatus', 'PRESENCE': 'optional'}, None]}
   CSGMembershipStatus  *CSGMembershipStatus
   list []interface{}
}
func (self *UERegisterAcceptExtensions)createOT() interface{}{
    return nil
}
var table_UERegisterAcceptExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_UERegisterAcceptExtensions = make([]int, 1)

func (self *UERegisterAcceptExtensions) GetIECount() int{
   count := 0
   if self.CSGMembershipStatus != nil { count += 1 }
   return count//ObjSet
}
func (self *UERegisterAcceptExtensions) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 21: //CSGMembershipStatus
        if self.CSGMembershipStatus != nil { return true }
   }
   return false//ObjSet
}
func (self *UERegisterAcceptExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 21: //CSGMembershipStatus
        self.CSGMembershipStatus = &CSGMembershipStatus{}
        self.CSGMembershipStatus.Unpack(st)
        self.list = append(self.list, self.CSGMembershipStatus)
   }
}
func (self *UERegisterAcceptExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 21: //CSGMembershipStatus
        if self.CSGMembershipStatus != nil {self.CSGMembershipStatus.Pack(st)}
      default:
      break
   }
}
func init() {
table_UERegisterAcceptExtensions[21] = &HNBAPPROTOCOLEXTENSION{ID:ProtocolIEID{idCSGMembershipStatus}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&CSGMembershipStatus{}, PRESENCE:Presence{Presenceoptional}, }
order_UERegisterAcceptExtensions[0] = 21
   }

type UERegisterRejectIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-UE-Identity', 'CRITICALITY': 'reject', 'TYPE': 'UE-Identity', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   UEIdentity  UEIdentity
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *UERegisterRejectIEs)createOT() interface{}{
    return nil
}
var table_UERegisterRejectIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_UERegisterRejectIEs = make([]int, 3)

func (self *UERegisterRejectIEs) GetIECount() int{
   count := 0
   count +=1 //self.UEIdentity
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *UERegisterRejectIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        return true //self.UEIdentity
      case 1: //Cause
        return true //self.Cause
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *UERegisterRejectIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        self.UEIdentity.Unpack(st)
        self.list = append(self.list, &self.UEIdentity)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *UERegisterRejectIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        self.UEIdentity.Pack(st)
      case 1: //Cause
        self.Cause.Pack(st)
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_UERegisterRejectIEs[5] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idUEIdentity}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&UEIdentity{}, PRESENCE:Presence{Presencemandatory}, }
order_UERegisterRejectIEs[0] = 5
table_UERegisterRejectIEs[1] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_UERegisterRejectIEs[1] = 1
table_UERegisterRejectIEs[2] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_UERegisterRejectIEs[2] = 2
   }

type UERegisterRejectExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UERegisterRejectExtensions)createOT() interface{}{
    return nil
}
var table_UERegisterRejectExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_UERegisterRejectExtensions = make([]int, 0)

type UEDeRegisterIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-Context-ID', 'CRITICALITY': 'reject', 'TYPE': 'Context-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, None]}
   ContextID  ContextID
   Cause  Cause
   list []interface{}
}
func (self *UEDeRegisterIEs)createOT() interface{}{
    return nil
}
var table_UEDeRegisterIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_UEDeRegisterIEs = make([]int, 2)

func (self *UEDeRegisterIEs) GetIECount() int{
   count := 0
   count +=1 //self.ContextID
   count +=1 //self.Cause
   return count//ObjSet
}
func (self *UEDeRegisterIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        return true //self.ContextID
      case 1: //Cause
        return true //self.Cause
   }
   return false//ObjSet
}
func (self *UEDeRegisterIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Unpack(st)
        self.list = append(self.list, &self.ContextID)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
   }
}
func (self *UEDeRegisterIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Pack(st)
      case 1: //Cause
        self.Cause.Pack(st)
      default:
      break
   }
}
func init() {
table_UEDeRegisterIEs[4] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idContextID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ContextID{}, PRESENCE:Presence{Presencemandatory}, }
order_UEDeRegisterIEs[0] = 4
table_UEDeRegisterIEs[1] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_UEDeRegisterIEs[1] = 1
   }

type UEDeRegisterExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UEDeRegisterExtensions)createOT() interface{}{
    return nil
}
var table_UEDeRegisterExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_UEDeRegisterExtensions = make([]int, 0)

type CSGMembershipUpdateIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-Context-ID', 'CRITICALITY': 'reject', 'TYPE': 'Context-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CSGMembershipStatus', 'CRITICALITY': 'reject', 'TYPE': 'CSGMembershipStatus', 'PRESENCE': 'mandatory'}, None]}
   ContextID  ContextID
   CSGMembershipStatus  CSGMembershipStatus
   list []interface{}
}
func (self *CSGMembershipUpdateIEs)createOT() interface{}{
    return nil
}
var table_CSGMembershipUpdateIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_CSGMembershipUpdateIEs = make([]int, 2)

func (self *CSGMembershipUpdateIEs) GetIECount() int{
   count := 0
   count +=1 //self.ContextID
   count +=1 //self.CSGMembershipStatus
   return count//ObjSet
}
func (self *CSGMembershipUpdateIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        return true //self.ContextID
      case 21: //CSGMembershipStatus
        return true //self.CSGMembershipStatus
   }
   return false//ObjSet
}
func (self *CSGMembershipUpdateIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Unpack(st)
        self.list = append(self.list, &self.ContextID)
      case 21: //CSGMembershipStatus
        self.CSGMembershipStatus.Unpack(st)
        self.list = append(self.list, &self.CSGMembershipStatus)
   }
}
func (self *CSGMembershipUpdateIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Pack(st)
      case 21: //CSGMembershipStatus
        self.CSGMembershipStatus.Pack(st)
      default:
      break
   }
}
func init() {
table_CSGMembershipUpdateIEs[4] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idContextID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ContextID{}, PRESENCE:Presence{Presencemandatory}, }
order_CSGMembershipUpdateIEs[0] = 4
table_CSGMembershipUpdateIEs[21] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCSGMembershipStatus}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CSGMembershipStatus{}, PRESENCE:Presence{Presencemandatory}, }
order_CSGMembershipUpdateIEs[1] = 21
   }

type CSGMembershipUpdateExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CSGMembershipUpdateExtensions)createOT() interface{}{
    return nil
}
var table_CSGMembershipUpdateExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_CSGMembershipUpdateExtensions = make([]int, 0)

type AccessControlQueryIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-UE-Identity', 'CRITICALITY': 'reject', 'TYPE': 'UE-Identity', 'PRESENCE': 'mandatory'}, None]}
   UEIdentity  UEIdentity
   list []interface{}
}
func (self *AccessControlQueryIEs)createOT() interface{}{
    return nil
}
var table_AccessControlQueryIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_AccessControlQueryIEs = make([]int, 1)

func (self *AccessControlQueryIEs) GetIECount() int{
   count := 0
   count +=1 //self.UEIdentity
   return count//ObjSet
}
func (self *AccessControlQueryIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        return true //self.UEIdentity
   }
   return false//ObjSet
}
func (self *AccessControlQueryIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        self.UEIdentity.Unpack(st)
        self.list = append(self.list, &self.UEIdentity)
   }
}
func (self *AccessControlQueryIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        self.UEIdentity.Pack(st)
      default:
      break
   }
}
func init() {
table_AccessControlQueryIEs[5] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idUEIdentity}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&UEIdentity{}, PRESENCE:Presence{Presencemandatory}, }
order_AccessControlQueryIEs[0] = 5
   }

type AccessControlQueryExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *AccessControlQueryExtensions)createOT() interface{}{
    return nil
}
var table_AccessControlQueryExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_AccessControlQueryExtensions = make([]int, 0)

type AccessControlResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-UE-Identity', 'CRITICALITY': 'reject', 'TYPE': 'UE-Identity', 'PRESENCE': 'mandatory'}, {'ID': 'id-AccessResult', 'CRITICALITY': 'reject', 'TYPE': 'AccessResult', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'optional'}, None]}
   UEIdentity  UEIdentity
   AccessResult  AccessResult
   Cause  *Cause
   list []interface{}
}
func (self *AccessControlResponseIEs)createOT() interface{}{
    return nil
}
var table_AccessControlResponseIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_AccessControlResponseIEs = make([]int, 3)

func (self *AccessControlResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.UEIdentity
   count +=1 //self.AccessResult
   if self.Cause != nil { count += 1 }
   return count//ObjSet
}
func (self *AccessControlResponseIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        return true //self.UEIdentity
      case 25: //AccessResult
        return true //self.AccessResult
      case 1: //Cause
        if self.Cause != nil { return true }
   }
   return false//ObjSet
}
func (self *AccessControlResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        self.UEIdentity.Unpack(st)
        self.list = append(self.list, &self.UEIdentity)
      case 25: //AccessResult
        self.AccessResult.Unpack(st)
        self.list = append(self.list, &self.AccessResult)
      case 1: //Cause
        self.Cause = &Cause{}
        self.Cause.Unpack(st)
        self.list = append(self.list, self.Cause)
   }
}
func (self *AccessControlResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 5: //UEIdentity
        self.UEIdentity.Pack(st)
      case 25: //AccessResult
        self.AccessResult.Pack(st)
      case 1: //Cause
        if self.Cause != nil {self.Cause.Pack(st)}
      default:
      break
   }
}
func init() {
table_AccessControlResponseIEs[5] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idUEIdentity}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&UEIdentity{}, PRESENCE:Presence{Presencemandatory}, }
order_AccessControlResponseIEs[0] = 5
table_AccessControlResponseIEs[25] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idAccessResult}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&AccessResult{}, PRESENCE:Presence{Presencemandatory}, }
order_AccessControlResponseIEs[1] = 25
table_AccessControlResponseIEs[1] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presenceoptional}, }
order_AccessControlResponseIEs[2] = 1
   }

type AccessControlResponseExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *AccessControlResponseExtensions)createOT() interface{}{
    return nil
}
var table_AccessControlResponseExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_AccessControlResponseExtensions = make([]int, 0)

type TNLUpdateRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-Context-ID', 'CRITICALITY': 'reject', 'TYPE': 'Context-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RABList', 'CRITICALITY': 'reject', 'TYPE': 'RABList', 'PRESENCE': 'mandatory'}, {'ID': 'id-Update-cause', 'CRITICALITY': 'reject', 'TYPE': 'Update-cause', 'PRESENCE': 'mandatory'}, None]}
   ContextID  ContextID
   RABList  RABList
   Updatecause  Updatecause
   list []interface{}
}
func (self *TNLUpdateRequestIEs)createOT() interface{}{
    return nil
}
var table_TNLUpdateRequestIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_TNLUpdateRequestIEs = make([]int, 3)

func (self *TNLUpdateRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.ContextID
   count +=1 //self.RABList
   count +=1 //self.Updatecause
   return count//ObjSet
}
func (self *TNLUpdateRequestIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        return true //self.ContextID
      case 22: //RABList
        return true //self.RABList
      case 26: //Updatecause
        return true //self.Updatecause
   }
   return false//ObjSet
}
func (self *TNLUpdateRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Unpack(st)
        self.list = append(self.list, &self.ContextID)
      case 22: //RABList
        self.RABList.Unpack(st)
        self.list = append(self.list, &self.RABList)
      case 26: //Updatecause
        self.Updatecause.Unpack(st)
        self.list = append(self.list, &self.Updatecause)
   }
}
func (self *TNLUpdateRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Pack(st)
      case 22: //RABList
        self.RABList.Pack(st)
      case 26: //Updatecause
        self.Updatecause.Pack(st)
      default:
      break
   }
}
func init() {
table_TNLUpdateRequestIEs[4] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idContextID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ContextID{}, PRESENCE:Presence{Presencemandatory}, }
order_TNLUpdateRequestIEs[0] = 4
table_TNLUpdateRequestIEs[22] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idRABList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RABList{}, PRESENCE:Presence{Presencemandatory}, }
order_TNLUpdateRequestIEs[1] = 22
table_TNLUpdateRequestIEs[26] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idUpdatecause}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&Updatecause{}, PRESENCE:Presence{Presencemandatory}, }
order_TNLUpdateRequestIEs[2] = 26
   }

type TNLUpdateExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TNLUpdateExtensions)createOT() interface{}{
    return nil
}
var table_TNLUpdateExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_TNLUpdateExtensions = make([]int, 0)

type TNLUpdateResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-Context-ID', 'CRITICALITY': 'reject', 'TYPE': 'Context-ID', 'PRESENCE': 'mandatory'}, None]}
   ContextID  ContextID
   list []interface{}
}
func (self *TNLUpdateResponseIEs)createOT() interface{}{
    return nil
}
var table_TNLUpdateResponseIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_TNLUpdateResponseIEs = make([]int, 1)

func (self *TNLUpdateResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.ContextID
   return count//ObjSet
}
func (self *TNLUpdateResponseIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        return true //self.ContextID
   }
   return false//ObjSet
}
func (self *TNLUpdateResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Unpack(st)
        self.list = append(self.list, &self.ContextID)
   }
}
func (self *TNLUpdateResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Pack(st)
      default:
      break
   }
}
func init() {
table_TNLUpdateResponseIEs[4] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idContextID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ContextID{}, PRESENCE:Presence{Presencemandatory}, }
order_TNLUpdateResponseIEs[0] = 4
   }

type TNLUpdateResponseExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TNLUpdateResponseExtensions)createOT() interface{}{
    return nil
}
var table_TNLUpdateResponseExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_TNLUpdateResponseExtensions = make([]int, 0)

type TNLUpdateFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-Context-ID', 'CRITICALITY': 'reject', 'TYPE': 'Context-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   ContextID  ContextID
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *TNLUpdateFailureIEs)createOT() interface{}{
    return nil
}
var table_TNLUpdateFailureIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_TNLUpdateFailureIEs = make([]int, 3)

func (self *TNLUpdateFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.ContextID
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *TNLUpdateFailureIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        return true //self.ContextID
      case 1: //Cause
        return true //self.Cause
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *TNLUpdateFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Unpack(st)
        self.list = append(self.list, &self.ContextID)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *TNLUpdateFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Pack(st)
      case 1: //Cause
        self.Cause.Pack(st)
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_TNLUpdateFailureIEs[4] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idContextID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ContextID{}, PRESENCE:Presence{Presencemandatory}, }
order_TNLUpdateFailureIEs[0] = 4
table_TNLUpdateFailureIEs[1] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_TNLUpdateFailureIEs[1] = 1
table_TNLUpdateFailureIEs[2] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_TNLUpdateFailureIEs[2] = 2
   }

type TNLUpdateFailureExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TNLUpdateFailureExtensions)createOT() interface{}{
    return nil
}
var table_TNLUpdateFailureExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_TNLUpdateFailureExtensions = make([]int, 0)

type HNBConfigTransferRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-NeighbourInfoRequestList', 'CRITICALITY': 'reject', 'TYPE': 'NeighbourInfoRequestList', 'PRESENCE': 'mandatory'}, None]}
   NeighbourInfoRequestList  NeighbourInfoRequestList
   list []interface{}
}
func (self *HNBConfigTransferRequestIEs)createOT() interface{}{
    return nil
}
var table_HNBConfigTransferRequestIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_HNBConfigTransferRequestIEs = make([]int, 1)

func (self *HNBConfigTransferRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.NeighbourInfoRequestList
   return count//ObjSet
}
func (self *HNBConfigTransferRequestIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 28: //NeighbourInfoRequestList
        return true //self.NeighbourInfoRequestList
   }
   return false//ObjSet
}
func (self *HNBConfigTransferRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 28: //NeighbourInfoRequestList
        self.NeighbourInfoRequestList.Unpack(st)
        self.list = append(self.list, &self.NeighbourInfoRequestList)
   }
}
func (self *HNBConfigTransferRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 28: //NeighbourInfoRequestList
        self.NeighbourInfoRequestList.Pack(st)
      default:
      break
   }
}
func init() {
table_HNBConfigTransferRequestIEs[28] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idNeighbourInfoRequestList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&NeighbourInfoRequestList{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBConfigTransferRequestIEs[0] = 28
   }

type HNBConfigTransferRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *HNBConfigTransferRequestExtensions)createOT() interface{}{
    return nil
}
var table_HNBConfigTransferRequestExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBConfigTransferRequestExtensions = make([]int, 0)

type HNBConfigTransferResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-NeighbourInfoList', 'CRITICALITY': 'reject', 'TYPE': 'NeighbourInfoList', 'PRESENCE': 'mandatory'}, None]}
   NeighbourInfoList  NeighbourInfoList
   list []interface{}
}
func (self *HNBConfigTransferResponseIEs)createOT() interface{}{
    return nil
}
var table_HNBConfigTransferResponseIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_HNBConfigTransferResponseIEs = make([]int, 1)

func (self *HNBConfigTransferResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.NeighbourInfoList
   return count//ObjSet
}
func (self *HNBConfigTransferResponseIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 27: //NeighbourInfoList
        return true //self.NeighbourInfoList
   }
   return false//ObjSet
}
func (self *HNBConfigTransferResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 27: //NeighbourInfoList
        self.NeighbourInfoList.Unpack(st)
        self.list = append(self.list, &self.NeighbourInfoList)
   }
}
func (self *HNBConfigTransferResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 27: //NeighbourInfoList
        self.NeighbourInfoList.Pack(st)
      default:
      break
   }
}
func init() {
table_HNBConfigTransferResponseIEs[27] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idNeighbourInfoList}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&NeighbourInfoList{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBConfigTransferResponseIEs[0] = 27
   }

type HNBConfigTransferResponseExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *HNBConfigTransferResponseExtensions)createOT() interface{}{
    return nil
}
var table_HNBConfigTransferResponseExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBConfigTransferResponseExtensions = make([]int, 0)

type HNBConfigTransferFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *HNBConfigTransferFailureIEs)createOT() interface{}{
    return nil
}
var table_HNBConfigTransferFailureIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_HNBConfigTransferFailureIEs = make([]int, 2)

func (self *HNBConfigTransferFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *HNBConfigTransferFailureIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        return true //self.Cause
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *HNBConfigTransferFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *HNBConfigTransferFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        self.Cause.Pack(st)
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_HNBConfigTransferFailureIEs[1] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBConfigTransferFailureIEs[0] = 1
table_HNBConfigTransferFailureIEs[2] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_HNBConfigTransferFailureIEs[1] = 2
   }

type HNBConfigTransferFailureExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *HNBConfigTransferFailureExtensions)createOT() interface{}{
    return nil
}
var table_HNBConfigTransferFailureExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBConfigTransferFailureExtensions = make([]int, 0)

type RelocationCompleteIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-Context-ID', 'CRITICALITY': 'reject', 'TYPE': 'Context-ID', 'PRESENCE': 'mandatory'}, None]}
   ContextID  ContextID
   list []interface{}
}
func (self *RelocationCompleteIEs)createOT() interface{}{
    return nil
}
var table_RelocationCompleteIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_RelocationCompleteIEs = make([]int, 1)

func (self *RelocationCompleteIEs) GetIECount() int{
   count := 0
   count +=1 //self.ContextID
   return count//ObjSet
}
func (self *RelocationCompleteIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        return true //self.ContextID
   }
   return false//ObjSet
}
func (self *RelocationCompleteIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Unpack(st)
        self.list = append(self.list, &self.ContextID)
   }
}
func (self *RelocationCompleteIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 4: //ContextID
        self.ContextID.Pack(st)
      default:
      break
   }
}
func init() {
table_RelocationCompleteIEs[4] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idContextID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ContextID{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationCompleteIEs[0] = 4
   }

type RelocationCompleteExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RelocationCompleteExtensions)createOT() interface{}{
    return nil
}
var table_RelocationCompleteExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_RelocationCompleteExtensions = make([]int, 0)

type ErrorIndicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ErrorIndicationIEs)createOT() interface{}{
    return nil
}
var table_ErrorIndicationIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_ErrorIndicationIEs = make([]int, 2)

func (self *ErrorIndicationIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ErrorIndicationIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        return true //self.Cause
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ErrorIndicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ErrorIndicationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        self.Cause.Pack(st)
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ErrorIndicationIEs[1] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_ErrorIndicationIEs[0] = 1
table_ErrorIndicationIEs[2] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[1] = 2
   }

type ErrorIndicationExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *ErrorIndicationExtensions)createOT() interface{}{
    return nil
}
var table_ErrorIndicationExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_ErrorIndicationExtensions = make([]int, 0)

type PrivateMessageIEs struct { //ObjSet 1 {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PRIVATE-IES', 'members': [None]}
   list []interface{}
}
func (self *PrivateMessageIEs)createOT() interface{}{
    return nil
}
var table_PrivateMessageIEs = make(map[int]*HNBAPPRIVATEIES)

var order_PrivateMessageIEs = make([]int, 0)

type HNBHeartBeatRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-HNB-HBInfo', 'CRITICALITY': 'ignore', 'TYPE': 'HNB-HBInfo', 'PRESENCE': 'mandatory'}]}
   HNBHBInfo  HNBHBInfo
   list []interface{}
}
func (self *HNBHeartBeatRequestIEs)createOT() interface{}{
    return nil
}
var table_HNBHeartBeatRequestIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_HNBHeartBeatRequestIEs = make([]int, 1)

func (self *HNBHeartBeatRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.HNBHBInfo
   return count//ObjSet
}
func (self *HNBHeartBeatRequestIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 22: //HNBHBInfo
        return true //self.HNBHBInfo
   }
   return false//ObjSet
}
func (self *HNBHeartBeatRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 22: //HNBHBInfo
        self.HNBHBInfo.Unpack(st)
        self.list = append(self.list, &self.HNBHBInfo)
   }
}
func (self *HNBHeartBeatRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 22: //HNBHBInfo
        self.HNBHBInfo.Pack(st)
      default:
      break
   }
}
func init() {
table_HNBHeartBeatRequestIEs[22] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idHNBHBInfo}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&HNBHBInfo{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBHeartBeatRequestIEs[0] = 22
   }

type HNBHeartBeatRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *HNBHeartBeatRequestExtensions)createOT() interface{}{
    return nil
}
var table_HNBHeartBeatRequestExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBHeartBeatRequestExtensions = make([]int, 0)

type HNBHeartBeatResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-IES', 'members': [{'ID': 'id-HNB-HBInfo', 'CRITICALITY': 'ignore', 'TYPE': 'HNB-HBInfo', 'PRESENCE': 'mandatory'}]}
   HNBHBInfo  HNBHBInfo
   list []interface{}
}
func (self *HNBHeartBeatResponseIEs)createOT() interface{}{
    return nil
}
var table_HNBHeartBeatResponseIEs = make(map[int]*HNBAPPROTOCOLIES)

var order_HNBHeartBeatResponseIEs = make([]int, 1)

func (self *HNBHeartBeatResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.HNBHBInfo
   return count//ObjSet
}
func (self *HNBHeartBeatResponseIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 22: //HNBHBInfo
        return true //self.HNBHBInfo
   }
   return false//ObjSet
}
func (self *HNBHeartBeatResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 22: //HNBHBInfo
        self.HNBHBInfo.Unpack(st)
        self.list = append(self.list, &self.HNBHBInfo)
   }
}
func (self *HNBHeartBeatResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLIESid).Value
   switch cat {
      case 22: //HNBHBInfo
        self.HNBHBInfo.Pack(st)
      default:
      break
   }
}
func init() {
table_HNBHeartBeatResponseIEs[22] = &HNBAPPROTOCOLIES{ID:ProtocolIEID{idHNBHBInfo}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&HNBHBInfo{}, PRESENCE:Presence{Presencemandatory}, }
order_HNBHeartBeatResponseIEs[0] = 22
   }

type HNBHeartBeatResponseExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *HNBHeartBeatResponseExtensions)createOT() interface{}{
    return nil
}
var table_HNBHeartBeatResponseExtensions = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBHeartBeatResponseExtensions = make([]int, 0)

type CriticalityDiagnosticsIEListExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsIEListExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsIEListExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsIEListExtIEs = make([]int, 0)

type CriticalityDiagnosticsExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsExtIEs = make([]int, 0)

type CGIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CGIExtIEs)createOT() interface{}{
    return nil
}
var table_CGIExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_CGIExtIEs = make([]int, 0)

type GeographicLocationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GeographicLocationExtIEs)createOT() interface{}{
    return nil
}
var table_GeographicLocationExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_GeographicLocationExtIEs = make([]int, 0)

type GeographicalCoordinatesExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GeographicalCoordinatesExtIEs)createOT() interface{}{
    return nil
}
var table_GeographicalCoordinatesExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_GeographicalCoordinatesExtIEs = make([]int, 0)

type GlobalRNCIDExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GlobalRNCIDExtIEs)createOT() interface{}{
    return nil
}
var table_GlobalRNCIDExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_GlobalRNCIDExtIEs = make([]int, 0)

type HNBCellIdentifierExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *HNBCellIdentifierExtIEs)createOT() interface{}{
    return nil
}
var table_HNBCellIdentifierExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBCellIdentifierExtIEs = make([]int, 0)

type HNBConfigInfoExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *HNBConfigInfoExtIEs)createOT() interface{}{
    return nil
}
var table_HNBConfigInfoExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBConfigInfoExtIEs = make([]int, 0)

type HNBLocationInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-HNB-Internet-Information', 'CRITICALITY': 'reject', 'EXTENSION': 'IP-Address', 'PRESENCE': 'optional'}, None]}
   HNBInternetInformation  *IPAddress
   list []interface{}
}
func (self *HNBLocationInformationExtIEs)createOT() interface{}{
    return nil
}
var table_HNBLocationInformationExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBLocationInformationExtIEs = make([]int, 1)

func (self *HNBLocationInformationExtIEs) GetIECount() int{
   count := 0
   if self.HNBInternetInformation != nil { count += 1 }
   return count//ObjSet
}
func (self *HNBLocationInformationExtIEs) GetOT(id interface{}) bool{
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 17: //HNBInternetInformation
        if self.HNBInternetInformation != nil { return true }
   }
   return false//ObjSet
}
func (self *HNBLocationInformationExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 17: //HNBInternetInformation
        self.HNBInternetInformation = &IPAddress{}
        self.HNBInternetInformation.Unpack(st)
        self.list = append(self.list, self.HNBInternetInformation)
   }
}
func (self *HNBLocationInformationExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(HNBAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 17: //HNBInternetInformation
        if self.HNBInternetInformation != nil {self.HNBInternetInformation.Pack(st)}
      default:
      break
   }
}
func init() {
table_HNBLocationInformationExtIEs[17] = &HNBAPPROTOCOLEXTENSION{ID:ProtocolIEID{idHNBInternetInformation}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&IPAddress{}, PRESENCE:Presence{Presenceoptional}, }
order_HNBLocationInformationExtIEs[0] = 17
   }

type HNBIdentityExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *HNBIdentityExtIEs)createOT() interface{}{
    return nil
}
var table_HNBIdentityExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_HNBIdentityExtIEs = make([]int, 0)

type IPAddressExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *IPAddressExtIEs)createOT() interface{}{
    return nil
}
var table_IPAddressExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_IPAddressExtIEs = make([]int, 0)

type MacroCoverageInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MacroCoverageInformationExtIEs)createOT() interface{}{
    return nil
}
var table_MacroCoverageInformationExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_MacroCoverageInformationExtIEs = make([]int, 0)

type RABListItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABListItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABListItemExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_RABListItemExtIEs = make([]int, 0)

type TransportInfoExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TransportInfoExtIEs)createOT() interface{}{
    return nil
}
var table_TransportInfoExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_TransportInfoExtIEs = make([]int, 0)

type UECapabilitiesExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UECapabilitiesExtIEs)createOT() interface{}{
    return nil
}
var table_UECapabilitiesExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_UECapabilitiesExtIEs = make([]int, 0)

type UTRANCellIDExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'HNBAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UTRANCellIDExtIEs)createOT() interface{}{
    return nil
}
var table_UTRANCellIDExtIEs = make(map[int]*HNBAPPROTOCOLEXTENSION)

var order_UTRANCellIDExtIEs = make([]int, 0)

func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'HNBRegisterRequest', 'SUCCESSFUL OUTCOME': 'HNBRegisterAccept', 'UNSUCCESSFUL OUTCOME': 'HNBRegisterReject', 'PROCEDURE CODE': 'id-HNBRegister', 'CRITICALITY': 'reject'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idHNBRegister}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&HNBRegisterRequest{}, SUCCESSFULOUTCOME:&HNBRegisterAccept{}, UNSUCCESSFULOUTCOME:&HNBRegisterReject{}, PROCEDURECODE:ProcedureCode{idHNBRegister}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetHNBRegisterINITIATINGMESSAGE() (*HNBRegisterRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &HNBRegisterRequest{}, uint64(idHNBRegister), int(Criticalityreject)
}
func GetHNBRegisterSUCCESSFULOUTCOME() (*HNBRegisterAccept, uint64, int) {/*TYPE, ID, Cricality*/
 return &HNBRegisterAccept{}, uint64(idHNBRegister), int(Criticalityreject)
}
func GetHNBRegisterUNSUCCESSFULOUTCOME() (*HNBRegisterReject, uint64, int) {/*TYPE, ID, Cricality*/
 return &HNBRegisterReject{}, uint64(idHNBRegister), int(Criticalityreject)
}
func (self *HNBRegisterRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *HNBRegisterRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *HNBRegisterRequest) createOT() interface{} {
   return &HNBRegisterRequest{}
}
func (self *HNBRegisterRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *HNBRegisterRequest) GetIECount() int{
    return 0
}
func (self *HNBRegisterAccept) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *HNBRegisterAccept) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *HNBRegisterAccept) createOT() interface{} {
   return &HNBRegisterAccept{}
}
func (self *HNBRegisterAccept) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *HNBRegisterAccept) GetIECount() int{
    return 0
}
func (self *HNBRegisterReject) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *HNBRegisterReject) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *HNBRegisterReject) createOT() interface{} {
   return &HNBRegisterReject{}
}
func (self *HNBRegisterReject) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *HNBRegisterReject) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'UERegisterRequest', 'SUCCESSFUL OUTCOME': 'UERegisterAccept', 'UNSUCCESSFUL OUTCOME': 'UERegisterReject', 'PROCEDURE CODE': 'id-UERegister', 'CRITICALITY': 'reject'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idUERegister}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&UERegisterRequest{}, SUCCESSFULOUTCOME:&UERegisterAccept{}, UNSUCCESSFULOUTCOME:&UERegisterReject{}, PROCEDURECODE:ProcedureCode{idUERegister}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetUERegisterINITIATINGMESSAGE() (*UERegisterRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &UERegisterRequest{}, uint64(idUERegister), int(Criticalityreject)
}
func GetUERegisterSUCCESSFULOUTCOME() (*UERegisterAccept, uint64, int) {/*TYPE, ID, Cricality*/
 return &UERegisterAccept{}, uint64(idUERegister), int(Criticalityreject)
}
func GetUERegisterUNSUCCESSFULOUTCOME() (*UERegisterReject, uint64, int) {/*TYPE, ID, Cricality*/
 return &UERegisterReject{}, uint64(idUERegister), int(Criticalityreject)
}
func (self *UERegisterRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *UERegisterRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *UERegisterRequest) createOT() interface{} {
   return &UERegisterRequest{}
}
func (self *UERegisterRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *UERegisterRequest) GetIECount() int{
    return 0
}
func (self *UERegisterAccept) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *UERegisterAccept) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *UERegisterAccept) createOT() interface{} {
   return &UERegisterAccept{}
}
func (self *UERegisterAccept) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *UERegisterAccept) GetIECount() int{
    return 0
}
func (self *UERegisterReject) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *UERegisterReject) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *UERegisterReject) createOT() interface{} {
   return &UERegisterReject{}
}
func (self *UERegisterReject) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *UERegisterReject) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'UEDe-Register', 'PROCEDURE CODE': 'id-UEDe-Register', 'CRITICALITY': 'ignore'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idUEDeRegister}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&UEDeRegister{}, PROCEDURECODE:ProcedureCode{idUEDeRegister}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetUEDeRegisterINITIATINGMESSAGE() (*UEDeRegister, uint64, int) {/*TYPE, ID, Cricality*/
 return &UEDeRegister{}, uint64(idUEDeRegister), int(Criticalityignore)
}
func (self *UEDeRegister) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *UEDeRegister) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *UEDeRegister) createOT() interface{} {
   return &UEDeRegister{}
}
func (self *UEDeRegister) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *UEDeRegister) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'HNBDe-Register', 'PROCEDURE CODE': 'id-HNBDe-Register', 'CRITICALITY': 'ignore'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idHNBDeRegister}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&HNBDeRegister{}, PROCEDURECODE:ProcedureCode{idHNBDeRegister}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetHNBDeRegisterINITIATINGMESSAGE() (*HNBDeRegister, uint64, int) {/*TYPE, ID, Cricality*/
 return &HNBDeRegister{}, uint64(idHNBDeRegister), int(Criticalityignore)
}
func (self *HNBDeRegister) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *HNBDeRegister) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *HNBDeRegister) createOT() interface{} {
   return &HNBDeRegister{}
}
func (self *HNBDeRegister) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *HNBDeRegister) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ErrorIndication', 'PROCEDURE CODE': 'id-ErrorIndication', 'CRITICALITY': 'ignore'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idErrorIndication}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ErrorIndication{}, PROCEDURECODE:ProcedureCode{idErrorIndication}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetErrorIndicationINITIATINGMESSAGE() (*ErrorIndication, uint64, int) {/*TYPE, ID, Cricality*/
 return &ErrorIndication{}, uint64(idErrorIndication), int(Criticalityignore)
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'CSGMembershipUpdate', 'PROCEDURE CODE': 'id-CSGMembershipUpdate', 'CRITICALITY': 'ignore'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idCSGMembershipUpdate}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&CSGMembershipUpdate{}, PROCEDURECODE:ProcedureCode{idCSGMembershipUpdate}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetCsgmembershipupdateINITIATINGMESSAGE() (*CSGMembershipUpdate, uint64, int) {/*TYPE, ID, Cricality*/
 return &CSGMembershipUpdate{}, uint64(idCSGMembershipUpdate), int(Criticalityignore)
}
func (self *CSGMembershipUpdate) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *CSGMembershipUpdate) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *CSGMembershipUpdate) createOT() interface{} {
   return &CSGMembershipUpdate{}
}
func (self *CSGMembershipUpdate) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *CSGMembershipUpdate) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'AccessControlQuery', 'SUCCESSFUL OUTCOME': 'AccessControlResponse', 'PROCEDURE CODE': 'id-AccessControlQuery', 'CRITICALITY': 'reject'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idAccessControlQuery}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&AccessControlQuery{}, SUCCESSFULOUTCOME:&AccessControlResponse{}, PROCEDURECODE:ProcedureCode{idAccessControlQuery}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetAccessControlQueryINITIATINGMESSAGE() (*AccessControlQuery, uint64, int) {/*TYPE, ID, Cricality*/
 return &AccessControlQuery{}, uint64(idAccessControlQuery), int(Criticalityreject)
}
func GetAccessControlQuerySUCCESSFULOUTCOME() (*AccessControlResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &AccessControlResponse{}, uint64(idAccessControlQuery), int(Criticalityreject)
}
func (self *AccessControlQuery) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *AccessControlQuery) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *AccessControlQuery) createOT() interface{} {
   return &AccessControlQuery{}
}
func (self *AccessControlQuery) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *AccessControlQuery) GetIECount() int{
    return 0
}
func (self *AccessControlResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *AccessControlResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *AccessControlResponse) createOT() interface{} {
   return &AccessControlResponse{}
}
func (self *AccessControlResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *AccessControlResponse) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'TNLUpdateRequest', 'SUCCESSFUL OUTCOME': 'TNLUpdateResponse', 'UNSUCCESSFUL OUTCOME': 'TNLUpdateFailure', 'PROCEDURE CODE': 'id-TNLUpdate', 'CRITICALITY': 'reject'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idTNLUpdate}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&TNLUpdateRequest{}, SUCCESSFULOUTCOME:&TNLUpdateResponse{}, UNSUCCESSFULOUTCOME:&TNLUpdateFailure{}, PROCEDURECODE:ProcedureCode{idTNLUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetTnlUpdateINITIATINGMESSAGE() (*TNLUpdateRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &TNLUpdateRequest{}, uint64(idTNLUpdate), int(Criticalityreject)
}
func GetTnlUpdateSUCCESSFULOUTCOME() (*TNLUpdateResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &TNLUpdateResponse{}, uint64(idTNLUpdate), int(Criticalityreject)
}
func GetTnlUpdateUNSUCCESSFULOUTCOME() (*TNLUpdateFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &TNLUpdateFailure{}, uint64(idTNLUpdate), int(Criticalityreject)
}
func (self *TNLUpdateRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *TNLUpdateRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *TNLUpdateRequest) createOT() interface{} {
   return &TNLUpdateRequest{}
}
func (self *TNLUpdateRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *TNLUpdateRequest) GetIECount() int{
    return 0
}
func (self *TNLUpdateResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *TNLUpdateResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *TNLUpdateResponse) createOT() interface{} {
   return &TNLUpdateResponse{}
}
func (self *TNLUpdateResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *TNLUpdateResponse) GetIECount() int{
    return 0
}
func (self *TNLUpdateFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *TNLUpdateFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *TNLUpdateFailure) createOT() interface{} {
   return &TNLUpdateFailure{}
}
func (self *TNLUpdateFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *TNLUpdateFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'HNBConfigTransferRequest', 'SUCCESSFUL OUTCOME': 'HNBConfigTransferResponse', 'UNSUCCESSFUL OUTCOME': 'HNBConfigTransferFailure', 'PROCEDURE CODE': 'id-HNBConfigTransfer', 'CRITICALITY': 'reject'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idHNBConfigTransfer}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&HNBConfigTransferRequest{}, SUCCESSFULOUTCOME:&HNBConfigTransferResponse{}, UNSUCCESSFULOUTCOME:&HNBConfigTransferFailure{}, PROCEDURECODE:ProcedureCode{idHNBConfigTransfer}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetHNBConfigTransferINITIATINGMESSAGE() (*HNBConfigTransferRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &HNBConfigTransferRequest{}, uint64(idHNBConfigTransfer), int(Criticalityreject)
}
func GetHNBConfigTransferSUCCESSFULOUTCOME() (*HNBConfigTransferResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &HNBConfigTransferResponse{}, uint64(idHNBConfigTransfer), int(Criticalityreject)
}
func GetHNBConfigTransferUNSUCCESSFULOUTCOME() (*HNBConfigTransferFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &HNBConfigTransferFailure{}, uint64(idHNBConfigTransfer), int(Criticalityreject)
}
func (self *HNBConfigTransferRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *HNBConfigTransferRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *HNBConfigTransferRequest) createOT() interface{} {
   return &HNBConfigTransferRequest{}
}
func (self *HNBConfigTransferRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *HNBConfigTransferRequest) GetIECount() int{
    return 0
}
func (self *HNBConfigTransferResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *HNBConfigTransferResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *HNBConfigTransferResponse) createOT() interface{} {
   return &HNBConfigTransferResponse{}
}
func (self *HNBConfigTransferResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *HNBConfigTransferResponse) GetIECount() int{
    return 0
}
func (self *HNBConfigTransferFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *HNBConfigTransferFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *HNBConfigTransferFailure) createOT() interface{} {
   return &HNBConfigTransferFailure{}
}
func (self *HNBConfigTransferFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *HNBConfigTransferFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RelocationComplete', 'PROCEDURE CODE': 'id-RelocationComplete', 'CRITICALITY': 'ignore'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idRelocationComplete}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RelocationComplete{}, PROCEDURECODE:ProcedureCode{idRelocationComplete}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetRelocationCompleteINITIATINGMESSAGE() (*RelocationComplete, uint64, int) {/*TYPE, ID, Cricality*/
 return &RelocationComplete{}, uint64(idRelocationComplete), int(Criticalityignore)
}
func (self *RelocationComplete) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RelocationComplete) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RelocationComplete) createOT() interface{} {
   return &RelocationComplete{}
}
func (self *RelocationComplete) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RelocationComplete) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'PrivateMessage', 'PROCEDURE CODE': 'id-privateMessage', 'CRITICALITY': 'ignore'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idprivateMessage}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&PrivateMessage{}, PROCEDURECODE:ProcedureCode{idprivateMessage}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'HNBAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'HNBHeartBeatRequest', 'SUCCESSFUL OUTCOME': 'HNBHeartBeatResponse', 'PROCEDURE CODE': 'id-HNBHeartBeat', 'CRITICALITY': 'ignore'}]}
table_HNBAPELEMENTARYPROCEDURES[HNBAPELEMENTARYPROCEDUREprocedureCode{idHNBHeartBeat}] = &HNBAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&HNBHeartBeatRequest{}, SUCCESSFULOUTCOME:&HNBHeartBeatResponse{}, PROCEDURECODE:ProcedureCode{idHNBHeartBeat}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetHNBHeartBeatINITIATINGMESSAGE() (*HNBHeartBeatRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &HNBHeartBeatRequest{}, uint64(idHNBHeartBeat), int(Criticalityignore)
}
func GetHNBHeartBeatSUCCESSFULOUTCOME() (*HNBHeartBeatResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &HNBHeartBeatResponse{}, uint64(idHNBHeartBeat), int(Criticalityignore)
}
func (self *HNBHeartBeatRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *HNBHeartBeatRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *HNBHeartBeatRequest) createOT() interface{} {
   return &HNBHeartBeatRequest{}
}
func (self *HNBHeartBeatRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *HNBHeartBeatRequest) GetIECount() int{
    return 0
}
func (self *HNBHeartBeatResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *HNBHeartBeatResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *HNBHeartBeatResponse) createOT() interface{} {
   return &HNBHeartBeatResponse{}
}
func (self *HNBHeartBeatResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *HNBHeartBeatResponse) GetIECount() int{
    return 0
}
var maxPrivateIEs uint64 = 65535
var maxProtocolExtensions uint64 = 65535
var maxProtocolIEs uint64 = 65535
var idHNBRegister uint64 = 1
const ProcedureCodeHNBRegister = 1
var idHNBDeRegister uint64 = 2
const ProcedureCodeHNBDeRegister = 2
var idUERegister uint64 = 3
const ProcedureCodeUERegister = 3
var idUEDeRegister uint64 = 4
const ProcedureCodeUEDeRegister = 4
var idErrorIndication uint64 = 5
const ProcedureCodeErrorIndication = 5
var idprivateMessage uint64 = 6
const ProcedureCodeprivateMessage = 6
var idCSGMembershipUpdate uint64 = 7
const ProcedureCodeCSGMembershipUpdate = 7
var idAccessControlQuery uint64 = 8
const ProcedureCodeAccessControlQuery = 8
var idTNLUpdate uint64 = 9
const ProcedureCodeTNLUpdate = 9
var idHNBConfigTransfer uint64 = 10
const ProcedureCodeHNBConfigTransfer = 10
var idRelocationComplete uint64 = 11
const ProcedureCodeRelocationComplete = 11
var idHNBHeartBeat uint64 = 12
const ProcedureCodeHNBHeartBeat = 12
var maxNrOfErrors uint64 = 256
var maxnoofRABs uint64 = 256
var maxnoofNeighbours uint64 = 32
var idCause uint64 = 1
const ProtocolIEIDCause = 1
var idCriticalityDiagnostics uint64 = 2
const ProtocolIEIDCriticalityDiagnostics = 2
var idHNBIdentity uint64 = 3
const ProtocolIEIDHNBIdentity = 3
var idContextID uint64 = 4
const ProtocolIEIDContextID = 4
var idUEIdentity uint64 = 5
const ProtocolIEIDUEIdentity = 5
var idLAC uint64 = 6
const ProtocolIEIDLAC = 6
var idRAC uint64 = 7
const ProtocolIEIDRAC = 7
var idHNBLocationInformation uint64 = 8
const ProtocolIEIDHNBLocationInformation = 8
var idPLMNidentity uint64 = 9
const ProtocolIEIDPLMNidentity = 9
var idSAC uint64 = 10
const ProtocolIEIDSAC = 10
var idCellIdentity uint64 = 11
const ProtocolIEIDCellIdentity = 11
var idRegistrationCause uint64 = 12
const ProtocolIEIDRegistrationCause = 12
var idUECapabilities uint64 = 13
const ProtocolIEIDUECapabilities = 13
var idRNCID uint64 = 14
const ProtocolIEIDRNCID = 14
var idCSGID uint64 = 15
const ProtocolIEIDCSGID = 15
var idBackoffTimer uint64 = 16
const ProtocolIEIDBackoffTimer = 16
var idHNBInternetInformation uint64 = 17
const ProtocolIEIDHNBInternetInformation = 17
var idHNBCellAccessMode uint64 = 18
const ProtocolIEIDHNBCellAccessMode = 18
var idMuxPortNumber uint64 = 19
const ProtocolIEIDMuxPortNumber = 19
var idServiceAreaForBroadcast uint64 = 20
const ProtocolIEIDServiceAreaForBroadcast = 20
var idCSGMembershipStatus uint64 = 21
const ProtocolIEIDCSGMembershipStatus = 21
var idRABList uint64 = 22
const ProtocolIEIDRABList = 22
var idHNBConfigInfo uint64 = 23
const ProtocolIEIDHNBConfigInfo = 23
var idLocalIurhIPAddress uint64 = 24
const ProtocolIEIDLocalIurhIPAddress = 24
var idAccessResult uint64 = 25
const ProtocolIEIDAccessResult = 25
var idUpdatecause uint64 = 26
const ProtocolIEIDUpdatecause = 26
var idNeighbourInfoList uint64 = 27
const ProtocolIEIDNeighbourInfoList = 27
var idNeighbourInfoRequestList uint64 = 28
const ProtocolIEIDNeighbourInfoRequestList = 28
var idRemoteIurhIPAddress uint64 = 29
const ProtocolIEIDRemoteIurhIPAddress = 29
var idPSC uint64 = 30
const ProtocolIEIDPSC = 30
var idHNBCellIdentifier uint64 = 31
const ProtocolIEIDHNBCellIdentifier = 31
var idHNBHBInfo uint64 = 22
const ProtocolIEIDHNBHBInfo = 22
