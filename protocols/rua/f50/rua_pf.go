
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package rua
import (
  log "github.com/sirupsen/logrus"
)
var version = "vf50"

func fmtrua() {log.Debug("rua")}
func (self *RUAPDU)Unpack(stream *Stream) {
    //coptions := []string{"initiatingMessage","successfulOutcome","unsuccessfulOutcome","Unknown"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in RUAPDU\n", choice, choice_len)
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
func (self * RUAPDU) Pack(stream *Stream) {
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
type RUAPDU struct { //[{'type': 'InitiatingMessage', 'name': 'initiatingMessage'}, {'type': 'SuccessfulOutcome', 'name': 'successfulOutcome'}, {'type': 'UnsuccessfulOutcome', 'name': 'unsuccessfulOutcome'}, None]
    InitiatingMessage *InitiatingMessage
    SuccessfulOutcome *SuccessfulOutcome
    UnsuccessfulOutcome *UnsuccessfulOutcome
} // RUAPDU

type InitiatingMessage struct { // [{'type': 'RUA-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RUA-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'RUA-ELEMENTARY-PROCEDURE.&criticality', 'table': ['RUA-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'RUA-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['RUA-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode RUAELEMENTARYPROCEDUREprocedureCode
    Criticality RUAELEMENTARYPROCEDUREcriticality
    Value RUAELEMENTARYPROCEDUREInitiatingMessage
}

func (self * InitiatingMessage) Unpack(stream *Stream) {
    //table {'type': 'RUA-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RUA-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_RUAELEMENTARYPROCEDURES[key].INITIATINGMESSAGE
    self.Value = out.(RUAELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(RUAELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'RUA-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['RUA-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * InitiatingMessage) Pack(stream *Stream) {
    //table {'type': 'RUA-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RUA-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_RUAELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(RUAELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type SuccessfulOutcome struct { // [{'type': 'RUA-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RUA-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'RUA-ELEMENTARY-PROCEDURE.&criticality', 'table': ['RUA-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'RUA-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['RUA-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode RUAELEMENTARYPROCEDUREprocedureCode
    Criticality RUAELEMENTARYPROCEDUREcriticality
    Value RUAELEMENTARYPROCEDURESuccessfulOutcome
}

func (self * SuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'RUA-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RUA-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_RUAELEMENTARYPROCEDURES[key].SUCCESSFULOUTCOME
    self.Value = out.(RUAELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(RUAELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'RUA-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['RUA-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * SuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'RUA-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RUA-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_RUAELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(RUAELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type UnsuccessfulOutcome struct { // [{'type': 'RUA-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RUA-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'RUA-ELEMENTARY-PROCEDURE.&criticality', 'table': ['RUA-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'RUA-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['RUA-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode RUAELEMENTARYPROCEDUREprocedureCode
    Criticality RUAELEMENTARYPROCEDUREcriticality
    Value RUAELEMENTARYPROCEDUREUnsuccessfulOutcome
}

func (self * UnsuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'RUA-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RUA-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_RUAELEMENTARYPROCEDURES[key].UNSUCCESSFULOUTCOME
    self.Value = out.(RUAELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(RUAELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'RUA-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['RUA-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * UnsuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'RUA-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RUA-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_RUAELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(RUAELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type Connect struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ConnectIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ConnectExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs ConnectIEs
    ProtocolExtensions *ConnectExtensions
}

func (self * Connect) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_ConnectIEs, order_ConnectIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &ConnectExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ConnectExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_ConnectExtensions, order_ConnectExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * Connect) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ConnectIEs, order_ConnectIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_ConnectExtensions, order_ConnectExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type DirectTransfer struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['DirectTransferIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DirectTransferExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs DirectTransferIEs
    ProtocolExtensions *DirectTransferExtensions
}

func (self * DirectTransfer) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_DirectTransferIEs, order_DirectTransferIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &DirectTransferExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DirectTransferExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_DirectTransferExtensions, order_DirectTransferExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DirectTransfer) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_DirectTransferIEs, order_DirectTransferIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_DirectTransferExtensions, order_DirectTransferExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type Disconnect struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['DisconnectIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DisconnectExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs DisconnectIEs
    ProtocolExtensions *DisconnectExtensions
}

func (self * Disconnect) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_DisconnectIEs, order_DisconnectIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &DisconnectExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DisconnectExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_DisconnectExtensions, order_DisconnectExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * Disconnect) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_DisconnectIEs, order_DisconnectIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_DisconnectExtensions, order_DisconnectExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ConnectionlessTransfer struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ConnectionlessTransferIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ConnectionlessTransferExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs ConnectionlessTransferIEs
    ProtocolExtensions *ConnectionlessTransferExtensions
}

func (self * ConnectionlessTransfer) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_ConnectionlessTransferIEs, order_ConnectionlessTransferIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &ConnectionlessTransferExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ConnectionlessTransferExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_ConnectionlessTransferExtensions, order_ConnectionlessTransferExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ConnectionlessTransfer) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ConnectionlessTransferIEs, order_ConnectionlessTransferIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_ConnectionlessTransferExtensions, order_ConnectionlessTransferExtensions} // p3
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
type EstablishmentCause struct {
  Value int
}
const (
    EstablishmentCauseemergency_call = 0
    EstablishmentCausenormal_call = 1

    /* Extensions */
)
func (self *EstablishmentCause) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *EstablishmentCause) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
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
type IntraDomainNasNodeSelector_Version_Later struct { // [{'type': 'BIT STRING', 'size': [15], 'name': 'futurecoding'}]
    Futurecoding BITSTRING
}
type IntraDomainNasNodeSelector_Version_Release99_CnType struct { //[{'type': 'Gsm-map-IDNNS', 'name': 'gsm-Map-IDNNS'}, {'type': 'Ansi-41-IDNNS', 'name': 'ansi-41-IDNNS'}]
    GsmMapIDNNS *GsmmapIDNNS
    Ansi41IDNNS *Ansi41IDNNS
} // IntraDomainNasNodeSelector_Version_Release99_CnType

type IntraDomainNasNodeSelector_Version_Release99 struct { // [{'type': 'CHOICE', 'members': [{'type': 'Gsm-map-IDNNS', 'name': 'gsm-Map-IDNNS'}, {'type': 'Ansi-41-IDNNS', 'name': 'ansi-41-IDNNS'}], 'name': 'cn-Type'}]
    CnType IntraDomainNasNodeSelector_Version_Release99_CnType
}
type IntraDomainNasNodeSelector_Version struct { //[{'type': 'SEQUENCE', 'members': [{'type': 'CHOICE', 'members': [{'type': 'Gsm-map-IDNNS', 'name': 'gsm-Map-IDNNS'}, {'type': 'Ansi-41-IDNNS', 'name': 'ansi-41-IDNNS'}], 'name': 'cn-Type'}], 'name': 'release99'}, {'type': 'SEQUENCE', 'members': [{'type': 'BIT STRING', 'size': [15], 'name': 'futurecoding'}], 'name': 'later'}]
    Release99 *IntraDomainNasNodeSelector_Version_Release99
    Later *IntraDomainNasNodeSelector_Version_Later
} // IntraDomainNasNodeSelector_Version

type IntraDomainNasNodeSelector struct { // [{'type': 'CHOICE', 'members': [{'type': 'SEQUENCE', 'members': [{'type': 'CHOICE', 'members': [{'type': 'Gsm-map-IDNNS', 'name': 'gsm-Map-IDNNS'}, {'type': 'Ansi-41-IDNNS', 'name': 'ansi-41-IDNNS'}], 'name': 'cn-Type'}], 'name': 'release99'}, {'type': 'SEQUENCE', 'members': [{'type': 'BIT STRING', 'size': [15], 'name': 'futurecoding'}], 'name': 'later'}], 'name': 'version'}]
    Version IntraDomainNasNodeSelector_Version
}

func (self * IntraDomainNasNodeSelector) Unpack(stream *Stream) {
    var Unpack_version = func(stream *Stream, self *IntraDomainNasNodeSelector_Version) {
        //coptions := []string{"release99","later"}
        choice := stream.get_choice(1, 0, 2)
        if choice == 0 { //ch1
            var Unpack_release99 = func(stream *Stream, self *IntraDomainNasNodeSelector_Version_Release99) { //[{'type': 'CHOICE', 'members': [{'type': 'Gsm-map-IDNNS', 'name': 'gsm-Map-IDNNS'}, {'type': 'Ansi-41-IDNNS', 'name': 'ansi-41-IDNNS'}], 'name': 'cn-Type'}]
                var Unpack_cnType = func(stream *Stream, self *IntraDomainNasNodeSelector_Version_Release99_CnType) {
                    //coptions := []string{"gsm-Map-IDNNS","ansi-41-IDNNS"}
                    choice := stream.get_choice(1, 0, 2)
                    if choice == 0 { //ch1
                        self.GsmMapIDNNS = &GsmmapIDNNS{}//cho6
                        self.GsmMapIDNNS.Unpack(stream)
                    } else if choice == 1 { //ch2
                        self.Ansi41IDNNS = &Ansi41IDNNS{}//cho6
                        self.Ansi41IDNNS.Unpack(stream)
                    }//end of if else

                }
                Unpack_cnType(stream, &self.CnType)// p2
                return
            }
            self.Release99 = &IntraDomainNasNodeSelector_Version_Release99{}//cho4
            Unpack_release99(stream, self.Release99);
        } else if choice == 1 { //ch2
            var Unpack_later = func(stream *Stream, self *IntraDomainNasNodeSelector_Version_Later) { //[{'type': 'BIT STRING', 'size': [15], 'name': 'futurecoding'}]
                var Unpack_futurecoding = func(st *Stream, self *BITSTRING){
                    self.Value = st.parsef_BitString(15, 15)
                }
                Unpack_futurecoding(stream, &self.Futurecoding)// p2
                return
            }
            self.Later = &IntraDomainNasNodeSelector_Version_Later{}//cho4
            Unpack_later(stream, self.Later);
        }//end of if else

    }
    Unpack_version(stream, &self.Version)// p2
    return
}

func (self * IntraDomainNasNodeSelector) Pack(stream *Stream) {
    var Pack_version = func(stream *Stream, self IntraDomainNasNodeSelector_Version) {
        if self.Release99 != nil {
            stream.set_choice(0, 1, 0, 2)
            var Pack_release99 = func(stream *Stream, self IntraDomainNasNodeSelector_Version_Release99) {//seq
                var Pack_cnType = func(stream *Stream, self IntraDomainNasNodeSelector_Version_Release99_CnType) {
                    if self.GsmMapIDNNS != nil {
                        stream.set_choice(0, 1, 0, 2)
                        self.GsmMapIDNNS.Pack(stream)//2
                    } else if self.Ansi41IDNNS != nil {
                        stream.set_choice(1, 1, 0, 2)
                        self.Ansi41IDNNS.Pack(stream)//2
                    }

                }
                Pack_cnType(stream, self.CnType) //f2
            }//end
            Pack_release99(stream, *self.Release99)//1
        } else if self.Later != nil {
            stream.set_choice(1, 1, 0, 2)
            var Pack_later = func(stream *Stream, self IntraDomainNasNodeSelector_Version_Later) {//seq
                var Pack_futurecoding = func(st *Stream, self BITSTRING) {
                    st.formatf_BitString(self.Value, 15)
                }
                Pack_futurecoding(stream, self.Futurecoding) //f2
            }//end
            Pack_later(stream, *self.Later)//1
        }

    }
    Pack_version(stream, self.Version) //f2
}//end

type GsmmapIDNNS_Routingbasis_Spare1 struct { // [{'type': 'RoutingParameter', 'name': 'routingparameter'}]
    Routingparameter RoutingParameter
}
type GsmmapIDNNS_Routingbasis_Spare2 struct { // [{'type': 'RoutingParameter', 'name': 'routingparameter'}]
    Routingparameter RoutingParameter
}
type GsmmapIDNNS_Routingbasis_IMEI struct { // [{'type': 'RoutingParameter', 'name': 'routingparameter'}]
    Routingparameter RoutingParameter
}
type GsmmapIDNNS_Routingbasis_IMSIcauseUEinitiatedEvent struct { // [{'type': 'RoutingParameter', 'name': 'routingparameter'}]
    Routingparameter RoutingParameter
}
type GsmmapIDNNS_Routingbasis_IMSIresponsetopaging struct { // [{'type': 'RoutingParameter', 'name': 'routingparameter'}]
    Routingparameter RoutingParameter
}
type GsmmapIDNNS_Routingbasis_TMSIofdifferentPLMN struct { // [{'type': 'RoutingParameter', 'name': 'routingparameter'}]
    Routingparameter RoutingParameter
}
type GsmmapIDNNS_Routingbasis_TMSIofsamePLMN struct { // [{'type': 'RoutingParameter', 'name': 'routingparameter'}]
    Routingparameter RoutingParameter
}
type GsmmapIDNNS_Routingbasis_LocalPTMSI struct { // [{'type': 'RoutingParameter', 'name': 'routingparameter'}]
    Routingparameter RoutingParameter
}
type GsmmapIDNNS_Routingbasis struct { //[{'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'localPTMSI'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'tMSIofsamePLMN'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'tMSIofdifferentPLMN'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'iMSIresponsetopaging'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'iMSIcauseUEinitiatedEvent'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'iMEI'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'spare2'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'spare1'}]
    LocalPTMSI *GsmmapIDNNS_Routingbasis_LocalPTMSI
    TMSIofsamePLMN *GsmmapIDNNS_Routingbasis_TMSIofsamePLMN
    TMSIofdifferentPLMN *GsmmapIDNNS_Routingbasis_TMSIofdifferentPLMN
    IMSIresponsetopaging *GsmmapIDNNS_Routingbasis_IMSIresponsetopaging
    IMSIcauseUEinitiatedEvent *GsmmapIDNNS_Routingbasis_IMSIcauseUEinitiatedEvent
    IMEI *GsmmapIDNNS_Routingbasis_IMEI
    Spare2 *GsmmapIDNNS_Routingbasis_Spare2
    Spare1 *GsmmapIDNNS_Routingbasis_Spare1
} // GsmmapIDNNS_Routingbasis

type GsmmapIDNNS struct { // [{'type': 'CHOICE', 'members': [{'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'localPTMSI'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'tMSIofsamePLMN'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'tMSIofdifferentPLMN'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'iMSIresponsetopaging'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'iMSIcauseUEinitiatedEvent'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'iMEI'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'spare2'}, {'type': 'SEQUENCE', 'members': [{'type': 'RoutingParameter', 'name': 'routingparameter'}], 'name': 'spare1'}], 'name': 'routingbasis'}, {'type': 'BOOLEAN', 'name': 'dummy'}]
    Routingbasis GsmmapIDNNS_Routingbasis
    Dummy BOOLEAN
}

func (self * GsmmapIDNNS) Unpack(stream *Stream) {
    var Unpack_routingbasis = func(stream *Stream, self *GsmmapIDNNS_Routingbasis) {
        //coptions := []string{"localPTMSI","tMSIofsamePLMN","tMSIofdifferentPLMN","iMSIresponsetopaging","iMSIcauseUEinitiatedEvent","iMEI","spare2","spare1"}
        choice := stream.get_choice(3, 0, 8)
        if choice == 0 { //ch1
            var Unpack_localPTMSI = func(stream *Stream, self *GsmmapIDNNS_Routingbasis_LocalPTMSI) { //[{'type': 'RoutingParameter', 'name': 'routingparameter'}]
                self.Routingparameter.Unpack(stream)// p8
                return
            }
            self.LocalPTMSI = &GsmmapIDNNS_Routingbasis_LocalPTMSI{}//cho4
            Unpack_localPTMSI(stream, self.LocalPTMSI);
        } else if choice == 1 { //ch2
            var Unpack_tMSIofsamePLMN = func(stream *Stream, self *GsmmapIDNNS_Routingbasis_TMSIofsamePLMN) { //[{'type': 'RoutingParameter', 'name': 'routingparameter'}]
                self.Routingparameter.Unpack(stream)// p8
                return
            }
            self.TMSIofsamePLMN = &GsmmapIDNNS_Routingbasis_TMSIofsamePLMN{}//cho4
            Unpack_tMSIofsamePLMN(stream, self.TMSIofsamePLMN);
        } else if choice == 2 { //ch2
            var Unpack_tMSIofdifferentPLMN = func(stream *Stream, self *GsmmapIDNNS_Routingbasis_TMSIofdifferentPLMN) { //[{'type': 'RoutingParameter', 'name': 'routingparameter'}]
                self.Routingparameter.Unpack(stream)// p8
                return
            }
            self.TMSIofdifferentPLMN = &GsmmapIDNNS_Routingbasis_TMSIofdifferentPLMN{}//cho4
            Unpack_tMSIofdifferentPLMN(stream, self.TMSIofdifferentPLMN);
        } else if choice == 3 { //ch2
            var Unpack_iMSIresponsetopaging = func(stream *Stream, self *GsmmapIDNNS_Routingbasis_IMSIresponsetopaging) { //[{'type': 'RoutingParameter', 'name': 'routingparameter'}]
                self.Routingparameter.Unpack(stream)// p8
                return
            }
            self.IMSIresponsetopaging = &GsmmapIDNNS_Routingbasis_IMSIresponsetopaging{}//cho4
            Unpack_iMSIresponsetopaging(stream, self.IMSIresponsetopaging);
        } else if choice == 4 { //ch2
            var Unpack_iMSIcauseUEinitiatedEvent = func(stream *Stream, self *GsmmapIDNNS_Routingbasis_IMSIcauseUEinitiatedEvent) { //[{'type': 'RoutingParameter', 'name': 'routingparameter'}]
                self.Routingparameter.Unpack(stream)// p8
                return
            }
            self.IMSIcauseUEinitiatedEvent = &GsmmapIDNNS_Routingbasis_IMSIcauseUEinitiatedEvent{}//cho4
            Unpack_iMSIcauseUEinitiatedEvent(stream, self.IMSIcauseUEinitiatedEvent);
        } else if choice == 5 { //ch2
            var Unpack_iMEI = func(stream *Stream, self *GsmmapIDNNS_Routingbasis_IMEI) { //[{'type': 'RoutingParameter', 'name': 'routingparameter'}]
                self.Routingparameter.Unpack(stream)// p8
                return
            }
            self.IMEI = &GsmmapIDNNS_Routingbasis_IMEI{}//cho4
            Unpack_iMEI(stream, self.IMEI);
        } else if choice == 6 { //ch2
            var Unpack_spare2 = func(stream *Stream, self *GsmmapIDNNS_Routingbasis_Spare2) { //[{'type': 'RoutingParameter', 'name': 'routingparameter'}]
                self.Routingparameter.Unpack(stream)// p8
                return
            }
            self.Spare2 = &GsmmapIDNNS_Routingbasis_Spare2{}//cho4
            Unpack_spare2(stream, self.Spare2);
        } else if choice == 7 { //ch2
            var Unpack_spare1 = func(stream *Stream, self *GsmmapIDNNS_Routingbasis_Spare1) { //[{'type': 'RoutingParameter', 'name': 'routingparameter'}]
                self.Routingparameter.Unpack(stream)// p8
                return
            }
            self.Spare1 = &GsmmapIDNNS_Routingbasis_Spare1{}//cho4
            Unpack_spare1(stream, self.Spare1);
        }//end of if else

    }
    Unpack_routingbasis(stream, &self.Routingbasis)// p2
    var  Unpack_dummy = func (st *Stream, self *BOOLEAN) {
        self.Value = st.parsef_bool()
        }
    Unpack_dummy(stream, &self.Dummy)// p2
    return
}

func (self * GsmmapIDNNS) Pack(stream *Stream) {
    var Pack_routingbasis = func(stream *Stream, self GsmmapIDNNS_Routingbasis) {
        if self.LocalPTMSI != nil {
            stream.set_choice(0, 3, 0, 8)
            var Pack_localPTMSI = func(stream *Stream, self GsmmapIDNNS_Routingbasis_LocalPTMSI) {//seq
                self.Routingparameter.Pack(stream)
            }//end
            Pack_localPTMSI(stream, *self.LocalPTMSI)//1
        } else if self.TMSIofsamePLMN != nil {
            stream.set_choice(1, 3, 0, 8)
            var Pack_tMSIofsamePLMN = func(stream *Stream, self GsmmapIDNNS_Routingbasis_TMSIofsamePLMN) {//seq
                self.Routingparameter.Pack(stream)
            }//end
            Pack_tMSIofsamePLMN(stream, *self.TMSIofsamePLMN)//1
        } else if self.TMSIofdifferentPLMN != nil {
            stream.set_choice(2, 3, 0, 8)
            var Pack_tMSIofdifferentPLMN = func(stream *Stream, self GsmmapIDNNS_Routingbasis_TMSIofdifferentPLMN) {//seq
                self.Routingparameter.Pack(stream)
            }//end
            Pack_tMSIofdifferentPLMN(stream, *self.TMSIofdifferentPLMN)//1
        } else if self.IMSIresponsetopaging != nil {
            stream.set_choice(3, 3, 0, 8)
            var Pack_iMSIresponsetopaging = func(stream *Stream, self GsmmapIDNNS_Routingbasis_IMSIresponsetopaging) {//seq
                self.Routingparameter.Pack(stream)
            }//end
            Pack_iMSIresponsetopaging(stream, *self.IMSIresponsetopaging)//1
        } else if self.IMSIcauseUEinitiatedEvent != nil {
            stream.set_choice(4, 3, 0, 8)
            var Pack_iMSIcauseUEinitiatedEvent = func(stream *Stream, self GsmmapIDNNS_Routingbasis_IMSIcauseUEinitiatedEvent) {//seq
                self.Routingparameter.Pack(stream)
            }//end
            Pack_iMSIcauseUEinitiatedEvent(stream, *self.IMSIcauseUEinitiatedEvent)//1
        } else if self.IMEI != nil {
            stream.set_choice(5, 3, 0, 8)
            var Pack_iMEI = func(stream *Stream, self GsmmapIDNNS_Routingbasis_IMEI) {//seq
                self.Routingparameter.Pack(stream)
            }//end
            Pack_iMEI(stream, *self.IMEI)//1
        } else if self.Spare2 != nil {
            stream.set_choice(6, 3, 0, 8)
            var Pack_spare2 = func(stream *Stream, self GsmmapIDNNS_Routingbasis_Spare2) {//seq
                self.Routingparameter.Pack(stream)
            }//end
            Pack_spare2(stream, *self.Spare2)//1
        } else if self.Spare1 != nil {
            stream.set_choice(7, 3, 0, 8)
            var Pack_spare1 = func(stream *Stream, self GsmmapIDNNS_Routingbasis_Spare1) {//seq
                self.Routingparameter.Pack(stream)
            }//end
            Pack_spare1(stream, *self.Spare1)//1
        }

    }
    Pack_routingbasis(stream, self.Routingbasis) //f2
    var Pack_dummy = func(st *Stream, self BOOLEAN) {
        st.formatf_bool(self.Value);
        }
    Pack_dummy(stream, self.Dummy) //f2
}//end

type Ansi41IDNNS struct {
  Len int
  Value HexBytes
}
func (self *Ansi41IDNNS) Unpack(st *Stream){
    self.Value = st.parsef_BitString(14, 14)
}
func (self *Ansi41IDNNS) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 14)
}
type RANAPMessage struct {
  Value HexBytes
}
func (self *RANAPMessage) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RANAPMessage) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RoutingParameter struct {
  Len int
  Value HexBytes
}
func (self *RoutingParameter) Unpack(st *Stream){
    self.Value = st.parsef_BitString(10, 10)
}
func (self *RoutingParameter) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 10)
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
    CauseRadioNetworknormal = 0
    CauseRadioNetworkconnect_failed = 1
    CauseRadioNetworknetwork_release = 2
    CauseRadioNetworkunspecified = 3

    /* Extensions */
)
func (self *CauseRadioNetwork) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 4, 1)
}
func (self *CauseRadioNetwork) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 4, 1)
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
type ProcedureCode struct {
  Value uint64
}
func (self *ProcedureCode) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * ProcedureCode) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
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
    _size := data.(RUAPROTOCOLIES_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['RUA-PROTOCOL-IES']}
    Items map[int]*RUAPROTOCOLIES
    order []int
}

type ProtocolIESingleContainer struct{ //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RUA-PROTOCOL-IES']}
   Item map[int]*RUAPROTOCOLIES //UserType
   order []int
}
func (self *ProtocolIESingleContainer) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RUA-PROTOCOL-IES']}
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

type ProtocolIEField struct { // [{'type': 'RUA-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'RUA-PROTOCOL-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'RUA-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id RUAPROTOCOLIESid
    Criticality RUAPROTOCOLIEScriticality
    Value RUAPROTOCOLIESValue
}

func (self * ProtocolIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'RUA-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := RUAPROTOCOLIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(RUAPROTOCOLIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'RUA-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * RUAPROTOCOLIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'RUA-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (RUAPROTOCOLIESid)(self.ID)
    if out.(RUAPROTOCOLIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(RUAPROTOCOLIES_IF).PackOT(stream, key)
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


type ProtocolIEContainerList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['IEsSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IEsSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'RUA-PROTOCOL-IES']}
    Items map[int]*RUAPROTOCOLIES
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
    _size := data.(RUAPROTOCOLEXTENSION_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolExtensionContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolExtensionField', 'actual-parameters': ['ExtensionSetParam']}, 'size': [(1, 'maxProtocolExtensions')], 'parameters': ['ExtensionSetParam'], 'param-types': ['RUA-PROTOCOL-EXTENSION']}
    Items map[int]*RUAPROTOCOLEXTENSION
    order []int
}

type ProtocolExtensionField struct { // [{'type': 'RUA-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}, {'type': 'RUA-PROTOCOL-EXTENSION.&criticality', 'table': ['ExtensionSetParam', ['id']], 'name': 'criticality'}, {'type': 'RUA-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}]
    Id RUAPROTOCOLEXTENSIONid
    Criticality RUAPROTOCOLEXTENSIONcriticality
    ExtensionValue RUAPROTOCOLEXTENSIONExtension
}

func (self * ProtocolExtensionField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'RUA-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := RUAPROTOCOLEXTENSION{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(RUAPROTOCOLEXTENSION_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'RUA-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}
    stream.set_location(location, _len)
    return
}

func (self * RUAPROTOCOLEXTENSION) Pack(stream *Stream, out interface{}) {
    //table {'type': 'RUA-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    key := (RUAPROTOCOLEXTENSIONid)(self.ID)
    if out.(RUAPROTOCOLEXTENSION_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(RUAPROTOCOLEXTENSION_IF).PackOT(stream, key)
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
    _size := data.(RUAPRIVATEIES_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type PrivateIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PrivateIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(1, 'maxPrivateIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['RUA-PRIVATE-IES']}
    Items map[int]*RUAPRIVATEIES
    order []int
}

type PrivateIEField struct { // [{'type': 'RUA-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'RUA-PRIVATE-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'RUA-PRIVATE-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id RUAPRIVATEIESid
    Criticality RUAPRIVATEIEScriticality
    Value RUAPRIVATEIESValue
}

func (self * PrivateIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'RUA-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := RUAPRIVATEIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(RUAPRIVATEIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'RUA-PRIVATE-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * RUAPRIVATEIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'RUA-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (RUAPRIVATEIESid)(self.ID)
    if out.(RUAPRIVATEIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(RUAPRIVATEIES_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type RUAELEMENTARYPROCEDURE struct {//CLASS {'members': [{'type': 'OpenType', 'name': '&InitiatingMessage'}, {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, {'type': 'ProcedureCode', 'name': '&procedureCode'}, {'type': 'Criticality', 'name': '&criticality'}], 'with-members': [{'with-order': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME', 'PROCEDURE CODE', 'CRITICALITY'], 'with-type': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME']}], 'alt-type': {'INITIATING MESSAGE': {'type': 'OpenType', 'name': '&InitiatingMessage'}, 'SUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, 'UNSUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, 'PROCEDURE CODE': {'type': 'ProcedureCode', 'name': '&procedureCode'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}}, 'id-type': {'type': 'ProcedureCode', 'name': '&procedureCode'}}
    INITIATINGMESSAGE interface{}
    SUCCESSFULOUTCOME interface{}
    UNSUCCESSFULOUTCOME interface{}
    PROCEDURECODE ProcedureCode
    CRITICALITY Criticality
}
type RUAELEMENTARYPROCEDUREInitiatingMessage interface{}
type RUAELEMENTARYPROCEDURESuccessfulOutcome interface{}
type RUAELEMENTARYPROCEDUREUnsuccessfulOutcome interface{}
type RUAELEMENTARYPROCEDUREprocedureCode ProcedureCode
func (self *RUAELEMENTARYPROCEDUREprocedureCode) Pack(st *Stream) {
    ieID := ProcedureCode(*self)
    ieID.Pack(st)
}
func (self *RUAELEMENTARYPROCEDUREprocedureCode) Unpack(st *Stream) {
    val := ProcedureCode{}
    val.Unpack(st)
    *self = RUAELEMENTARYPROCEDUREprocedureCode(val)
}
type RUAELEMENTARYPROCEDUREcriticality Criticality
func (self *RUAELEMENTARYPROCEDUREcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *RUAELEMENTARYPROCEDUREcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = RUAELEMENTARYPROCEDUREcriticality(val)
}

type RUAELEMENTARYPROCEDURE_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type RUAPRIVATEIES struct {//CLASS {'members': [{'type': 'PrivateIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'PrivateIE-ID', 'name': '&id'}}
    ID PrivateIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type RUAPRIVATEIESid PrivateIEID
func (self *RUAPRIVATEIESid) Pack(st *Stream) {
    ieID := PrivateIEID(*self)
    ieID.Pack(st)
}
func (self *RUAPRIVATEIESid) Unpack(st *Stream) {
    val := PrivateIEID{}
    val.Unpack(st)
    *self = RUAPRIVATEIESid(val)
}
type RUAPRIVATEIEScriticality Criticality
func (self *RUAPRIVATEIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *RUAPRIVATEIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = RUAPRIVATEIEScriticality(val)
}
type RUAPRIVATEIESValue interface{}
type RUAPRIVATEIESpresence Presence
func (self *RUAPRIVATEIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *RUAPRIVATEIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = RUAPRIVATEIESpresence(val)
}

type RUAPRIVATEIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type RUAPROTOCOLIES struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type RUAPROTOCOLIESid ProtocolIEID
func (self *RUAPROTOCOLIESid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *RUAPROTOCOLIESid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = RUAPROTOCOLIESid(val)
}
type RUAPROTOCOLIEScriticality Criticality
func (self *RUAPROTOCOLIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *RUAPROTOCOLIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = RUAPROTOCOLIEScriticality(val)
}
type RUAPROTOCOLIESValue interface{}
type RUAPROTOCOLIESpresence Presence
func (self *RUAPROTOCOLIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *RUAPROTOCOLIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = RUAPROTOCOLIESpresence(val)
}

type RUAPROTOCOLIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type RUAPROTOCOLEXTENSION struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Extension'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'EXTENSION', 'PRESENCE'], 'with-type': ['EXTENSION']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    EXTENSION interface{}
    PRESENCE Presence
}
type RUAPROTOCOLEXTENSIONid ProtocolIEID
func (self *RUAPROTOCOLEXTENSIONid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *RUAPROTOCOLEXTENSIONid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = RUAPROTOCOLEXTENSIONid(val)
}
type RUAPROTOCOLEXTENSIONcriticality Criticality
func (self *RUAPROTOCOLEXTENSIONcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *RUAPROTOCOLEXTENSIONcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = RUAPROTOCOLEXTENSIONcriticality(val)
}
type RUAPROTOCOLEXTENSIONExtension interface{}
type RUAPROTOCOLEXTENSIONpresence Presence
func (self *RUAPROTOCOLEXTENSIONpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *RUAPROTOCOLEXTENSIONpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = RUAPROTOCOLEXTENSIONpresence(val)
}

type RUAPROTOCOLEXTENSION_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
//class RUAELEMENTARYPROCEDURES: #OBJSET1 {'class': 'RUA-ELEMENTARY-PROCEDURE', 'members': [{}, {}, None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_RUAELEMENTARYPROCEDURES = make(map[RUAELEMENTARYPROCEDUREprocedureCode]*RUAELEMENTARYPROCEDURE)


//class RUAELEMENTARYPROCEDURESCLASS1: #OBJSET1 {'class': 'RUA-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_RUAELEMENTARYPROCEDURESCLASS1 = make(map[RUAELEMENTARYPROCEDUREprocedureCode]*RUAELEMENTARYPROCEDURE)


//class RUAELEMENTARYPROCEDURESCLASS2: #OBJSET1 {'class': 'RUA-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_RUAELEMENTARYPROCEDURESCLASS2 = make(map[RUAELEMENTARYPROCEDUREprocedureCode]*RUAELEMENTARYPROCEDURE)


type ConnectIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-IES', 'members': [{'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'reject', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-Context-ID', 'CRITICALITY': 'reject', 'TYPE': 'Context-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-IntraDomainNasNodeSelector', 'CRITICALITY': 'ignore', 'TYPE': 'IntraDomainNasNodeSelector', 'PRESENCE': 'optional'}, {'ID': 'id-Establishment-Cause', 'CRITICALITY': 'reject', 'TYPE': 'Establishment-Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANAP-Message', 'CRITICALITY': 'reject', 'TYPE': 'RANAP-Message', 'PRESENCE': 'mandatory'}, None]}
   CNDomainIndicator  CNDomainIndicator
   ContextID  ContextID
   IntraDomainNasNodeSelector  *IntraDomainNasNodeSelector
   EstablishmentCause  EstablishmentCause
   RANAPMessage  RANAPMessage
   list []interface{}
}
func (self *ConnectIEs)createOT() interface{}{
    return nil
}
var table_ConnectIEs = make(map[int]*RUAPROTOCOLIES)

var order_ConnectIEs = make([]int, 5)

func (self *ConnectIEs) GetIECount() int{
   count := 0
   count +=1 //self.CNDomainIndicator
   count +=1 //self.ContextID
   if self.IntraDomainNasNodeSelector != nil { count += 1 }
   count +=1 //self.EstablishmentCause
   count +=1 //self.RANAPMessage
   return count//ObjSet
}
func (self *ConnectIEs) GetOT(id interface{}) bool{
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 7: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 3: //ContextID
        return true //self.ContextID
      case 5: //IntraDomainNasNodeSelector
        if self.IntraDomainNasNodeSelector != nil { return true }
      case 6: //EstablishmentCause
        return true //self.EstablishmentCause
      case 4: //RANAPMessage
        return true //self.RANAPMessage
   }
   return false//ObjSet
}
func (self *ConnectIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 7: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 3: //ContextID
        self.ContextID.Unpack(st)
        self.list = append(self.list, &self.ContextID)
      case 5: //IntraDomainNasNodeSelector
        self.IntraDomainNasNodeSelector = &IntraDomainNasNodeSelector{}
        self.IntraDomainNasNodeSelector.Unpack(st)
        self.list = append(self.list, self.IntraDomainNasNodeSelector)
      case 6: //EstablishmentCause
        self.EstablishmentCause.Unpack(st)
        self.list = append(self.list, &self.EstablishmentCause)
      case 4: //RANAPMessage
        self.RANAPMessage.Unpack(st)
        self.list = append(self.list, &self.RANAPMessage)
   }
}
func (self *ConnectIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 7: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 3: //ContextID
        self.ContextID.Pack(st)
      case 5: //IntraDomainNasNodeSelector
        if self.IntraDomainNasNodeSelector != nil {self.IntraDomainNasNodeSelector.Pack(st)}
      case 6: //EstablishmentCause
        self.EstablishmentCause.Pack(st)
      case 4: //RANAPMessage
        self.RANAPMessage.Pack(st)
      default:
      break
   }
}
func init() {
table_ConnectIEs[7] = &RUAPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_ConnectIEs[0] = 7
table_ConnectIEs[3] = &RUAPROTOCOLIES{ID:ProtocolIEID{idContextID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ContextID{}, PRESENCE:Presence{Presencemandatory}, }
order_ConnectIEs[1] = 3
table_ConnectIEs[5] = &RUAPROTOCOLIES{ID:ProtocolIEID{idIntraDomainNasNodeSelector}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&IntraDomainNasNodeSelector{}, PRESENCE:Presence{Presenceoptional}, }
order_ConnectIEs[2] = 5
table_ConnectIEs[6] = &RUAPROTOCOLIES{ID:ProtocolIEID{idEstablishmentCause}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&EstablishmentCause{}, PRESENCE:Presence{Presencemandatory}, }
order_ConnectIEs[3] = 6
table_ConnectIEs[4] = &RUAPROTOCOLIES{ID:ProtocolIEID{idRANAPMessage}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANAPMessage{}, PRESENCE:Presence{Presencemandatory}, }
order_ConnectIEs[4] = 4
   }

type ConnectExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-CSGMembershipStatus', 'CRITICALITY': 'ignore', 'EXTENSION': 'CSGMembershipStatus', 'PRESENCE': 'optional'}, None]}
   CSGMembershipStatus  *CSGMembershipStatus
   list []interface{}
}
func (self *ConnectExtensions)createOT() interface{}{
    return nil
}
var table_ConnectExtensions = make(map[int]*RUAPROTOCOLEXTENSION)

var order_ConnectExtensions = make([]int, 1)

func (self *ConnectExtensions) GetIECount() int{
   count := 0
   if self.CSGMembershipStatus != nil { count += 1 }
   return count//ObjSet
}
func (self *ConnectExtensions) GetOT(id interface{}) bool{
   cat := id.(RUAPROTOCOLEXTENSIONid).Value
   switch cat {
      case 9: //CSGMembershipStatus
        if self.CSGMembershipStatus != nil { return true }
   }
   return false//ObjSet
}
func (self *ConnectExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLEXTENSIONid).Value
   switch cat {
      case 9: //CSGMembershipStatus
        self.CSGMembershipStatus = &CSGMembershipStatus{}
        self.CSGMembershipStatus.Unpack(st)
        self.list = append(self.list, self.CSGMembershipStatus)
   }
}
func (self *ConnectExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLEXTENSIONid).Value
   switch cat {
      case 9: //CSGMembershipStatus
        if self.CSGMembershipStatus != nil {self.CSGMembershipStatus.Pack(st)}
      default:
      break
   }
}
func init() {
table_ConnectExtensions[9] = &RUAPROTOCOLEXTENSION{ID:ProtocolIEID{idCSGMembershipStatus}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&CSGMembershipStatus{}, PRESENCE:Presence{Presenceoptional}, }
order_ConnectExtensions[0] = 9
   }

type DirectTransferIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-IES', 'members': [{'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'reject', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-Context-ID', 'CRITICALITY': 'reject', 'TYPE': 'Context-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANAP-Message', 'CRITICALITY': 'reject', 'TYPE': 'RANAP-Message', 'PRESENCE': 'mandatory'}, None]}
   CNDomainIndicator  CNDomainIndicator
   ContextID  ContextID
   RANAPMessage  RANAPMessage
   list []interface{}
}
func (self *DirectTransferIEs)createOT() interface{}{
    return nil
}
var table_DirectTransferIEs = make(map[int]*RUAPROTOCOLIES)

var order_DirectTransferIEs = make([]int, 3)

func (self *DirectTransferIEs) GetIECount() int{
   count := 0
   count +=1 //self.CNDomainIndicator
   count +=1 //self.ContextID
   count +=1 //self.RANAPMessage
   return count//ObjSet
}
func (self *DirectTransferIEs) GetOT(id interface{}) bool{
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 7: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 3: //ContextID
        return true //self.ContextID
      case 4: //RANAPMessage
        return true //self.RANAPMessage
   }
   return false//ObjSet
}
func (self *DirectTransferIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 7: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 3: //ContextID
        self.ContextID.Unpack(st)
        self.list = append(self.list, &self.ContextID)
      case 4: //RANAPMessage
        self.RANAPMessage.Unpack(st)
        self.list = append(self.list, &self.RANAPMessage)
   }
}
func (self *DirectTransferIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 7: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 3: //ContextID
        self.ContextID.Pack(st)
      case 4: //RANAPMessage
        self.RANAPMessage.Pack(st)
      default:
      break
   }
}
func init() {
table_DirectTransferIEs[7] = &RUAPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_DirectTransferIEs[0] = 7
table_DirectTransferIEs[3] = &RUAPROTOCOLIES{ID:ProtocolIEID{idContextID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ContextID{}, PRESENCE:Presence{Presencemandatory}, }
order_DirectTransferIEs[1] = 3
table_DirectTransferIEs[4] = &RUAPROTOCOLIES{ID:ProtocolIEID{idRANAPMessage}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANAPMessage{}, PRESENCE:Presence{Presencemandatory}, }
order_DirectTransferIEs[2] = 4
   }

type DirectTransferExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DirectTransferExtensions)createOT() interface{}{
    return nil
}
var table_DirectTransferExtensions = make(map[int]*RUAPROTOCOLEXTENSION)

var order_DirectTransferExtensions = make([]int, 0)

type DisconnectIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-IES', 'members': [{'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'reject', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-Context-ID', 'CRITICALITY': 'reject', 'TYPE': 'Context-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'reject', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANAP-Message', 'CRITICALITY': 'reject', 'TYPE': 'RANAP-Message', 'PRESENCE': 'conditional'}, None]}
   CNDomainIndicator  CNDomainIndicator
   ContextID  ContextID
   Cause  Cause
   RANAPMessage  RANAPMessage
   list []interface{}
}
func (self *DisconnectIEs)createOT() interface{}{
    return nil
}
var table_DisconnectIEs = make(map[int]*RUAPROTOCOLIES)

var order_DisconnectIEs = make([]int, 4)

func (self *DisconnectIEs) GetIECount() int{
   count := 0
   count +=1 //self.CNDomainIndicator
   count +=1 //self.ContextID
   count +=1 //self.Cause
   count +=1 //self.RANAPMessage
   return count//ObjSet
}
func (self *DisconnectIEs) GetOT(id interface{}) bool{
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 7: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 3: //ContextID
        return true //self.ContextID
      case 1: //Cause
        return true //self.Cause
      case 4: //RANAPMessage
        return true //self.RANAPMessage
   }
   return false//ObjSet
}
func (self *DisconnectIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 7: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 3: //ContextID
        self.ContextID.Unpack(st)
        self.list = append(self.list, &self.ContextID)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 4: //RANAPMessage
        self.RANAPMessage.Unpack(st)
        self.list = append(self.list, &self.RANAPMessage)
   }
}
func (self *DisconnectIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 7: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 3: //ContextID
        self.ContextID.Pack(st)
      case 1: //Cause
        self.Cause.Pack(st)
      case 4: //RANAPMessage
        self.RANAPMessage.Pack(st)
      default:
      break
   }
}
func init() {
table_DisconnectIEs[7] = &RUAPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_DisconnectIEs[0] = 7
table_DisconnectIEs[3] = &RUAPROTOCOLIES{ID:ProtocolIEID{idContextID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ContextID{}, PRESENCE:Presence{Presencemandatory}, }
order_DisconnectIEs[1] = 3
table_DisconnectIEs[1] = &RUAPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_DisconnectIEs[2] = 1
table_DisconnectIEs[4] = &RUAPROTOCOLIES{ID:ProtocolIEID{idRANAPMessage}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANAPMessage{}, PRESENCE:Presence{Presenceconditional}, }
order_DisconnectIEs[3] = 4
   }

type DisconnectExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DisconnectExtensions)createOT() interface{}{
    return nil
}
var table_DisconnectExtensions = make(map[int]*RUAPROTOCOLEXTENSION)

var order_DisconnectExtensions = make([]int, 0)

type ConnectionlessTransferIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-IES', 'members': [{'ID': 'id-RANAP-Message', 'CRITICALITY': 'reject', 'TYPE': 'RANAP-Message', 'PRESENCE': 'mandatory'}, None]}
   RANAPMessage  RANAPMessage
   list []interface{}
}
func (self *ConnectionlessTransferIEs)createOT() interface{}{
    return nil
}
var table_ConnectionlessTransferIEs = make(map[int]*RUAPROTOCOLIES)

var order_ConnectionlessTransferIEs = make([]int, 1)

func (self *ConnectionlessTransferIEs) GetIECount() int{
   count := 0
   count +=1 //self.RANAPMessage
   return count//ObjSet
}
func (self *ConnectionlessTransferIEs) GetOT(id interface{}) bool{
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 4: //RANAPMessage
        return true //self.RANAPMessage
   }
   return false//ObjSet
}
func (self *ConnectionlessTransferIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 4: //RANAPMessage
        self.RANAPMessage.Unpack(st)
        self.list = append(self.list, &self.RANAPMessage)
   }
}
func (self *ConnectionlessTransferIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 4: //RANAPMessage
        self.RANAPMessage.Pack(st)
      default:
      break
   }
}
func init() {
table_ConnectionlessTransferIEs[4] = &RUAPROTOCOLIES{ID:ProtocolIEID{idRANAPMessage}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANAPMessage{}, PRESENCE:Presence{Presencemandatory}, }
order_ConnectionlessTransferIEs[0] = 4
   }

type ConnectionlessTransferExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *ConnectionlessTransferExtensions)createOT() interface{}{
    return nil
}
var table_ConnectionlessTransferExtensions = make(map[int]*RUAPROTOCOLEXTENSION)

var order_ConnectionlessTransferExtensions = make([]int, 0)

type ErrorIndicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ErrorIndicationIEs)createOT() interface{}{
    return nil
}
var table_ErrorIndicationIEs = make(map[int]*RUAPROTOCOLIES)

var order_ErrorIndicationIEs = make([]int, 2)

func (self *ErrorIndicationIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ErrorIndicationIEs) GetOT(id interface{}) bool{
   cat := id.(RUAPROTOCOLIESid).Value
   switch cat {
      case 1: //Cause
        return true //self.Cause
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ErrorIndicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RUAPROTOCOLIESid).Value
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
   cat := id.(RUAPROTOCOLIESid).Value
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
table_ErrorIndicationIEs[1] = &RUAPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_ErrorIndicationIEs[0] = 1
table_ErrorIndicationIEs[2] = &RUAPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[1] = 2
   }

type ErrorIndicationExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *ErrorIndicationExtensions)createOT() interface{}{
    return nil
}
var table_ErrorIndicationExtensions = make(map[int]*RUAPROTOCOLEXTENSION)

var order_ErrorIndicationExtensions = make([]int, 0)

type PrivateMessageIEs struct { //ObjSet 1 {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PRIVATE-IES', 'members': [None]}
   list []interface{}
}
func (self *PrivateMessageIEs)createOT() interface{}{
    return nil
}
var table_PrivateMessageIEs = make(map[int]*RUAPRIVATEIES)

var order_PrivateMessageIEs = make([]int, 0)

type CriticalityDiagnosticsIEListExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsIEListExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsIEListExtIEs = make(map[int]*RUAPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsIEListExtIEs = make([]int, 0)

type CriticalityDiagnosticsExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RUA-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsExtIEs = make(map[int]*RUAPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsExtIEs = make([]int, 0)

func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RUA-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'Connect', 'PROCEDURE CODE': 'id-Connect', 'CRITICALITY': 'ignore'}]}
table_RUAELEMENTARYPROCEDURES[RUAELEMENTARYPROCEDUREprocedureCode{idConnect}] = &RUAELEMENTARYPROCEDURE{INITIATINGMESSAGE:&Connect{}, PROCEDURECODE:ProcedureCode{idConnect}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetConnectionRequestINITIATINGMESSAGE() (*Connect, uint64, int) {/*TYPE, ID, Cricality*/
 return &Connect{}, uint64(idConnect), int(Criticalityignore)
}
func (self *Connect) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *Connect) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *Connect) createOT() interface{} {
   return &Connect{}
}
func (self *Connect) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *Connect) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RUA-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'DirectTransfer', 'PROCEDURE CODE': 'id-DirectTransfer', 'CRITICALITY': 'ignore'}]}
table_RUAELEMENTARYPROCEDURES[RUAELEMENTARYPROCEDUREprocedureCode{idDirectTransfer}] = &RUAELEMENTARYPROCEDURE{INITIATINGMESSAGE:&DirectTransfer{}, PROCEDURECODE:ProcedureCode{idDirectTransfer}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetDirectTransferINITIATINGMESSAGE() (*DirectTransfer, uint64, int) {/*TYPE, ID, Cricality*/
 return &DirectTransfer{}, uint64(idDirectTransfer), int(Criticalityignore)
}
func (self *DirectTransfer) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *DirectTransfer) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *DirectTransfer) createOT() interface{} {
   return &DirectTransfer{}
}
func (self *DirectTransfer) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *DirectTransfer) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RUA-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'Disconnect', 'PROCEDURE CODE': 'id-Disconnect', 'CRITICALITY': 'ignore'}]}
table_RUAELEMENTARYPROCEDURES[RUAELEMENTARYPROCEDUREprocedureCode{idDisconnect}] = &RUAELEMENTARYPROCEDURE{INITIATINGMESSAGE:&Disconnect{}, PROCEDURECODE:ProcedureCode{idDisconnect}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetDisconnectRequestINITIATINGMESSAGE() (*Disconnect, uint64, int) {/*TYPE, ID, Cricality*/
 return &Disconnect{}, uint64(idDisconnect), int(Criticalityignore)
}
func (self *Disconnect) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *Disconnect) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *Disconnect) createOT() interface{} {
   return &Disconnect{}
}
func (self *Disconnect) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *Disconnect) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RUA-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ConnectionlessTransfer', 'PROCEDURE CODE': 'id-ConnectionlessTransfer', 'CRITICALITY': 'ignore'}]}
table_RUAELEMENTARYPROCEDURES[RUAELEMENTARYPROCEDUREprocedureCode{idConnectionlessTransfer}] = &RUAELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ConnectionlessTransfer{}, PROCEDURECODE:ProcedureCode{idConnectionlessTransfer}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetConnectionlessTransferINITIATINGMESSAGE() (*ConnectionlessTransfer, uint64, int) {/*TYPE, ID, Cricality*/
 return &ConnectionlessTransfer{}, uint64(idConnectionlessTransfer), int(Criticalityignore)
}
func (self *ConnectionlessTransfer) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ConnectionlessTransfer) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ConnectionlessTransfer) createOT() interface{} {
   return &ConnectionlessTransfer{}
}
func (self *ConnectionlessTransfer) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ConnectionlessTransfer) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RUA-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ErrorIndication', 'PROCEDURE CODE': 'id-ErrorIndication', 'CRITICALITY': 'ignore'}]}
table_RUAELEMENTARYPROCEDURES[RUAELEMENTARYPROCEDUREprocedureCode{idErrorIndication}] = &RUAELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ErrorIndication{}, PROCEDURECODE:ProcedureCode{idErrorIndication}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RUA-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'PrivateMessage', 'PROCEDURE CODE': 'id-privateMessage', 'CRITICALITY': 'ignore'}]}
table_RUAELEMENTARYPROCEDURES[RUAELEMENTARYPROCEDUREprocedureCode{idprivateMessage}] = &RUAELEMENTARYPROCEDURE{INITIATINGMESSAGE:&PrivateMessage{}, PROCEDURECODE:ProcedureCode{idprivateMessage}, CRITICALITY:Criticality{Criticalityignore}, }
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
var maxPrivateIEs uint64 = 65535
var maxProtocolExtensions uint64 = 65535
var maxProtocolIEs uint64 = 65535
var idConnect uint64 = 1
const ProcedureCodeConnect = 1
var idDirectTransfer uint64 = 2
const ProcedureCodeDirectTransfer = 2
var idDisconnect uint64 = 3
const ProcedureCodeDisconnect = 3
var idConnectionlessTransfer uint64 = 4
const ProcedureCodeConnectionlessTransfer = 4
var idErrorIndication uint64 = 5
const ProcedureCodeErrorIndication = 5
var idprivateMessage uint64 = 6
const ProcedureCodeprivateMessage = 6
var maxNrOfErrors uint64 = 256
var maxUEs uint64 = 64
var idCause uint64 = 1
const ProtocolIEIDCause = 1
var idCriticalityDiagnostics uint64 = 2
const ProtocolIEIDCriticalityDiagnostics = 2
var idContextID uint64 = 3
const ProtocolIEIDContextID = 3
var idRANAPMessage uint64 = 4
const ProtocolIEIDRANAPMessage = 4
var idIntraDomainNasNodeSelector uint64 = 5
const ProtocolIEIDIntraDomainNasNodeSelector = 5
var idEstablishmentCause uint64 = 6
const ProtocolIEIDEstablishmentCause = 6
var idCNDomainIndicator uint64 = 7
const ProtocolIEIDCNDomainIndicator = 7
var idCSGMembershipStatus uint64 = 9
const ProtocolIEIDCSGMembershipStatus = 9
