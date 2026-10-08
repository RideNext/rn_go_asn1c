
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package e2ap
import (
  log "github.com/sirupsen/logrus"
)
var version = "vv02"

func fmte2ap() {log.Debug("e2ap")}
func (self *E2APPDU)Unpack(stream *Stream) {
    //coptions := []string{"initiatingMessage","successfulOutcome","unsuccessfulOutcome","Unknown"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in E2APPDU\n", choice, choice_len)
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
func (self * E2APPDU) Pack(stream *Stream) {
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
type E2APPDU struct { //[{'type': 'InitiatingMessage', 'name': 'initiatingMessage'}, {'type': 'SuccessfulOutcome', 'name': 'successfulOutcome'}, {'type': 'UnsuccessfulOutcome', 'name': 'unsuccessfulOutcome'}, None]
    InitiatingMessage *InitiatingMessage
    SuccessfulOutcome *SuccessfulOutcome
    UnsuccessfulOutcome *UnsuccessfulOutcome
} // E2APPDU

type InitiatingMessage struct { // [{'type': 'E2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'E2AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['E2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'E2AP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['E2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode E2APELEMENTARYPROCEDUREprocedureCode
    Criticality E2APELEMENTARYPROCEDUREcriticality
    Value E2APELEMENTARYPROCEDUREInitiatingMessage
}

func (self * InitiatingMessage) Unpack(stream *Stream) {
    //table {'type': 'E2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_E2APELEMENTARYPROCEDURES[key].INITIATINGMESSAGE
    self.Value = out.(E2APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(E2APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'E2AP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['E2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * InitiatingMessage) Pack(stream *Stream) {
    //table {'type': 'E2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_E2APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(E2APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type SuccessfulOutcome struct { // [{'type': 'E2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'E2AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['E2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'E2AP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['E2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode E2APELEMENTARYPROCEDUREprocedureCode
    Criticality E2APELEMENTARYPROCEDUREcriticality
    Value E2APELEMENTARYPROCEDURESuccessfulOutcome
}

func (self * SuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'E2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_E2APELEMENTARYPROCEDURES[key].SUCCESSFULOUTCOME
    self.Value = out.(E2APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(E2APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'E2AP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['E2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * SuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'E2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_E2APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(E2APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type UnsuccessfulOutcome struct { // [{'type': 'E2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'E2AP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['E2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'E2AP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['E2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode E2APELEMENTARYPROCEDUREprocedureCode
    Criticality E2APELEMENTARYPROCEDUREcriticality
    Value E2APELEMENTARYPROCEDUREUnsuccessfulOutcome
}

func (self * UnsuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'E2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_E2APELEMENTARYPROCEDURES[key].UNSUCCESSFULOUTCOME
    self.Value = out.(E2APELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(E2APELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'E2AP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['E2AP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * UnsuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'E2AP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'E2AP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_E2APELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(E2APELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type RICsubscriptionRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICsubscriptionRequest-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICsubscriptionRequestIEs
}

func (self * RICsubscriptionRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICsubscriptionRequestIEs, order_RICsubscriptionRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICsubscriptionRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICsubscriptionRequestIEs, order_RICsubscriptionRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICsubscriptionDetails struct { // [{'type': 'RICeventTriggerDefinition', 'name': 'ricEventTriggerDefinition'}, {'type': 'RICactions-ToBeSetup-List', 'name': 'ricAction-ToBeSetup-List'}, None]
    RicEventTriggerDefinition RICeventTriggerDefinition
    RicActionToBeSetupList RICactionsToBeSetupList
}

func (self * RICsubscriptionDetails) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicEventTriggerDefinition.Unpack(stream)// p8
    self.RicActionToBeSetupList.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICsubscriptionDetails) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicEventTriggerDefinition.Pack(stream)
    self.RicActionToBeSetupList.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *RICactionsToBeSetupList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(16)
    _size += 1
    self.Items = make([]RICactionToBeSetupItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *RICactionsToBeSetupList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 16)
    //for item in table_RICactionToBeSetupItemIEs:
    val := ProtocolIESingleContainer{table_RICactionToBeSetupItemIEs, order_RICactionToBeSetupItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type RICactionsToBeSetupList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['RICaction-ToBeSetup-ItemIEs']}, 'size': [(1, 'maxofRICactionID')]}
    Items []RICactionToBeSetupItemIEs
}

type RICactionToBeSetupItem struct { // [{'type': 'RICactionID', 'name': 'ricActionID'}, {'type': 'RICactionType', 'name': 'ricActionType'}, {'type': 'RICactionDefinition', 'name': 'ricActionDefinition', 'optional': True}, {'type': 'RICsubsequentAction', 'name': 'ricSubsequentAction', 'optional': True}, None]
    RicActionID RICactionID
    RicActionType RICactionType
    RicActionDefinition *RICactionDefinition
    RicSubsequentAction *RICsubsequentAction
}

func (self * RICactionToBeSetupItem) Unpack(stream *Stream) {
    ricActionDefinition_flag := 0x00000002
    ricSubsequentAction_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RicActionID.Unpack(stream)// p8
    self.RicActionType.Unpack(stream)// p8
    if (ricActionDefinition_flag & _flags) == ricActionDefinition_flag { //cond2
        self.RicActionDefinition = &RICactionDefinition{}//7{'type': 'RICactionDefinition', 'name': 'ricActionDefinition', 'optional': True}
        self.RicActionDefinition.Unpack(stream)// p8
    }
    if (ricSubsequentAction_flag & _flags) == ricSubsequentAction_flag { //cond2
        self.RicSubsequentAction = &RICsubsequentAction{}//7{'type': 'RICsubsequentAction', 'name': 'ricSubsequentAction', 'optional': True}
        self.RicSubsequentAction.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICactionToBeSetupItem) Pack(stream *Stream) {
    const ricActionDefinition_flag uint = 0x00000002
    const ricSubsequentAction_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicActionID.Pack(stream)
    self.RicActionType.Pack(stream)
    if self.RicActionDefinition != nil { 
        _flags |= ricActionDefinition_flag
        self.RicActionDefinition.Pack(stream)
    }//end of optional
    if self.RicSubsequentAction != nil { 
        _flags |= ricSubsequentAction_flag
        self.RicSubsequentAction.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type RICsubscriptionResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICsubscriptionResponse-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICsubscriptionResponseIEs
}

func (self * RICsubscriptionResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICsubscriptionResponseIEs, order_RICsubscriptionResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICsubscriptionResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICsubscriptionResponseIEs, order_RICsubscriptionResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *RICactionAdmittedList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(16)
    _size += 1
    self.Items = make([]RICactionAdmittedItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *RICactionAdmittedList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 16)
    //for item in table_RICactionAdmittedItemIEs:
    val := ProtocolIESingleContainer{table_RICactionAdmittedItemIEs, order_RICactionAdmittedItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type RICactionAdmittedList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['RICaction-Admitted-ItemIEs']}, 'size': [(1, 'maxofRICactionID')]}
    Items []RICactionAdmittedItemIEs
}

type RICactionAdmittedItem struct { // [{'type': 'RICactionID', 'name': 'ricActionID'}, None]
    RicActionID RICactionID
}

func (self * RICactionAdmittedItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicActionID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICactionAdmittedItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicActionID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *RICactionNotAdmittedList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(17)
    _size += 0
    self.Items = make([]RICactionNotAdmittedItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *RICactionNotAdmittedList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-0, 17)
    //for item in table_RICactionNotAdmittedItemIEs:
    val := ProtocolIESingleContainer{table_RICactionNotAdmittedItemIEs, order_RICactionNotAdmittedItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type RICactionNotAdmittedList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['RICaction-NotAdmitted-ItemIEs']}, 'size': [(0, 'maxofRICactionID')]}
    Items []RICactionNotAdmittedItemIEs
}

type RICactionNotAdmittedItem struct { // [{'type': 'RICactionID', 'name': 'ricActionID'}, {'type': 'Cause', 'name': 'cause'}, None]
    RicActionID RICactionID
    Cause Cause
}

func (self * RICactionNotAdmittedItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicActionID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICactionNotAdmittedItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicActionID.Pack(stream)
    self.Cause.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICsubscriptionFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICsubscriptionFailure-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICsubscriptionFailureIEs
}

func (self * RICsubscriptionFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICsubscriptionFailureIEs, order_RICsubscriptionFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICsubscriptionFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICsubscriptionFailureIEs, order_RICsubscriptionFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICsubscriptionDeleteRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICsubscriptionDeleteRequest-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICsubscriptionDeleteRequestIEs
}

func (self * RICsubscriptionDeleteRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICsubscriptionDeleteRequestIEs, order_RICsubscriptionDeleteRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICsubscriptionDeleteRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICsubscriptionDeleteRequestIEs, order_RICsubscriptionDeleteRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICsubscriptionDeleteResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICsubscriptionDeleteResponse-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICsubscriptionDeleteResponseIEs
}

func (self * RICsubscriptionDeleteResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICsubscriptionDeleteResponseIEs, order_RICsubscriptionDeleteResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICsubscriptionDeleteResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICsubscriptionDeleteResponseIEs, order_RICsubscriptionDeleteResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICsubscriptionDeleteFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICsubscriptionDeleteFailure-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICsubscriptionDeleteFailureIEs
}

func (self * RICsubscriptionDeleteFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICsubscriptionDeleteFailureIEs, order_RICsubscriptionDeleteFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICsubscriptionDeleteFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICsubscriptionDeleteFailureIEs, order_RICsubscriptionDeleteFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICsubscriptionDeleteRequired struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICsubscriptionDeleteRequired-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICsubscriptionDeleteRequiredIEs
}

func (self * RICsubscriptionDeleteRequired) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICsubscriptionDeleteRequiredIEs, order_RICsubscriptionDeleteRequiredIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICsubscriptionDeleteRequired) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICsubscriptionDeleteRequiredIEs, order_RICsubscriptionDeleteRequiredIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *RICsubscriptionListwithCause) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(4294967295)
    self.Items = make([]RICsubscriptionwithCauseItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *RICsubscriptionListwithCause) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size, 4294967295)
    //for item in table_RICsubscriptionwithCauseItemIEs:
    val := ProtocolIESingleContainer{table_RICsubscriptionwithCauseItemIEs, order_RICsubscriptionwithCauseItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type RICsubscriptionListwithCause struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['RICsubscription-withCause-ItemIEs']}, 'size': [(1, 'maxofRICrequestID')]}
    Items []RICsubscriptionwithCauseItemIEs
}

type RICsubscriptionwithCauseItem struct { // [{'type': 'RICrequestID', 'name': 'ricRequestID'}, {'type': 'RANfunctionID', 'name': 'ranFunctionID'}, {'type': 'Cause', 'name': 'cause'}, None]
    RicRequestID RICrequestID
    RanFunctionID RANfunctionID
    Cause Cause
}

func (self * RICsubscriptionwithCauseItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicRequestID.Unpack(stream)// p8
    self.RanFunctionID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICsubscriptionwithCauseItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicRequestID.Pack(stream)
    self.RanFunctionID.Pack(stream)
    self.Cause.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICindication struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICindication-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICindicationIEs
}

func (self * RICindication) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICindicationIEs, order_RICindicationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICindication) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICindicationIEs, order_RICindicationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICcontrolRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICcontrolRequest-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICcontrolRequestIEs
}

func (self * RICcontrolRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICcontrolRequestIEs, order_RICcontrolRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICcontrolRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICcontrolRequestIEs, order_RICcontrolRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICcontrolAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICcontrolAcknowledge-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICcontrolAcknowledgeIEs
}

func (self * RICcontrolAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICcontrolAcknowledgeIEs, order_RICcontrolAcknowledgeIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICcontrolAcknowledge) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICcontrolAcknowledgeIEs, order_RICcontrolAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICcontrolFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICcontrolFailure-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICcontrolFailureIEs
}

func (self * RICcontrolFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICcontrolFailureIEs, order_RICcontrolFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICcontrolFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICcontrolFailureIEs, order_RICcontrolFailureIEs} // p3
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

type E2setupRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E2setupRequestIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E2setupRequestIEs
}

func (self * E2setupRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E2setupRequestIEs, order_E2setupRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2setupRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E2setupRequestIEs, order_E2setupRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2setupResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E2setupResponseIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E2setupResponseIEs
}

func (self * E2setupResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E2setupResponseIEs, order_E2setupResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2setupResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E2setupResponseIEs, order_E2setupResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2setupFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E2setupFailureIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E2setupFailureIEs
}

func (self * E2setupFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E2setupFailureIEs, order_E2setupFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2setupFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E2setupFailureIEs, order_E2setupFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2connectionUpdate struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E2connectionUpdate-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E2connectionUpdateIEs
}

func (self * E2connectionUpdate) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E2connectionUpdateIEs, order_E2connectionUpdateIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2connectionUpdate) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E2connectionUpdateIEs, order_E2connectionUpdateIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *E2connectionUpdateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]E2connectionUpdateItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *E2connectionUpdateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    //for item in table_E2connectionUpdateItemIEs:
    val := ProtocolIESingleContainer{table_E2connectionUpdateItemIEs, order_E2connectionUpdateItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type E2connectionUpdateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['E2connectionUpdate-ItemIEs']}, 'size': [(1, 'maxofTNLA')]}
    Items []E2connectionUpdateItemIEs
}

type E2connectionUpdateItem struct { // [{'type': 'TNLinformation', 'name': 'tnlInformation'}, {'type': 'TNLusage', 'name': 'tnlUsage'}, None]
    TnlInformation TNLinformation
    TnlUsage TNLusage
}

func (self * E2connectionUpdateItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.TnlInformation.Unpack(stream)// p8
    self.TnlUsage.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2connectionUpdateItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TnlInformation.Pack(stream)
    self.TnlUsage.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *E2connectionUpdateRemoveList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]E2connectionUpdateRemoveItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *E2connectionUpdateRemoveList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    //for item in table_E2connectionUpdateRemoveItemIEs:
    val := ProtocolIESingleContainer{table_E2connectionUpdateRemoveItemIEs, order_E2connectionUpdateRemoveItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type E2connectionUpdateRemoveList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['E2connectionUpdateRemove-ItemIEs']}, 'size': [(1, 'maxofTNLA')]}
    Items []E2connectionUpdateRemoveItemIEs
}

type E2connectionUpdateRemoveItem struct { // [{'type': 'TNLinformation', 'name': 'tnlInformation'}, None]
    TnlInformation TNLinformation
}

func (self * E2connectionUpdateRemoveItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.TnlInformation.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2connectionUpdateRemoveItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TnlInformation.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2connectionUpdateAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E2connectionUpdateAck-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E2connectionUpdateAckIEs
}

func (self * E2connectionUpdateAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E2connectionUpdateAckIEs, order_E2connectionUpdateAckIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2connectionUpdateAcknowledge) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E2connectionUpdateAckIEs, order_E2connectionUpdateAckIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *E2connectionSetupFailedList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]E2connectionSetupFailedItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *E2connectionSetupFailedList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    //for item in table_E2connectionSetupFailedItemIEs:
    val := ProtocolIESingleContainer{table_E2connectionSetupFailedItemIEs, order_E2connectionSetupFailedItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type E2connectionSetupFailedList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['E2connectionSetupFailed-ItemIEs']}, 'size': [(1, 'maxofTNLA')]}
    Items []E2connectionSetupFailedItemIEs
}

type E2connectionSetupFailedItem struct { // [{'type': 'TNLinformation', 'name': 'tnlInformation'}, {'type': 'Cause', 'name': 'cause'}, None]
    TnlInformation TNLinformation
    Cause Cause
}

func (self * E2connectionSetupFailedItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.TnlInformation.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2connectionSetupFailedItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TnlInformation.Pack(stream)
    self.Cause.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2connectionUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E2connectionUpdateFailure-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E2connectionUpdateFailureIEs
}

func (self * E2connectionUpdateFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E2connectionUpdateFailureIEs, order_E2connectionUpdateFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2connectionUpdateFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E2connectionUpdateFailureIEs, order_E2connectionUpdateFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2nodeConfigurationUpdate struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E2nodeConfigurationUpdate-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E2nodeConfigurationUpdateIEs
}

func (self * E2nodeConfigurationUpdate) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E2nodeConfigurationUpdateIEs, order_E2nodeConfigurationUpdateIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeConfigurationUpdate) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E2nodeConfigurationUpdateIEs, order_E2nodeConfigurationUpdateIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *E2nodeComponentConfigAdditionList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(1024)
    _size += 1
    self.Items = make([]E2nodeComponentConfigAdditionItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *E2nodeComponentConfigAdditionList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 1024)
    //for item in table_E2nodeComponentConfigAdditionItemIEs:
    val := ProtocolIESingleContainer{table_E2nodeComponentConfigAdditionItemIEs, order_E2nodeComponentConfigAdditionItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type E2nodeComponentConfigAdditionList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['E2nodeComponentConfigAddition-ItemIEs']}, 'size': [(1, 'maxofE2nodeComponents')]}
    Items []E2nodeComponentConfigAdditionItemIEs
}

type E2nodeComponentConfigAdditionItem struct { // [{'type': 'E2nodeComponentInterfaceType', 'name': 'e2nodeComponentInterfaceType'}, {'type': 'E2nodeComponentID', 'name': 'e2nodeComponentID'}, {'type': 'E2nodeComponentConfiguration', 'name': 'e2nodeComponentConfiguration'}, None]
    E2nodeComponentInterfaceType E2nodeComponentInterfaceType
    E2nodeComponentID E2nodeComponentID
    E2nodeComponentConfiguration E2nodeComponentConfiguration
}

func (self * E2nodeComponentConfigAdditionItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.E2nodeComponentInterfaceType.Unpack(stream)// p8
    self.E2nodeComponentID.Unpack(stream)// p8
    self.E2nodeComponentConfiguration.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentConfigAdditionItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.E2nodeComponentInterfaceType.Pack(stream)
    self.E2nodeComponentID.Pack(stream)
    self.E2nodeComponentConfiguration.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *E2nodeComponentConfigUpdateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(1024)
    _size += 1
    self.Items = make([]E2nodeComponentConfigUpdateItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *E2nodeComponentConfigUpdateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 1024)
    //for item in table_E2nodeComponentConfigUpdateItemIEs:
    val := ProtocolIESingleContainer{table_E2nodeComponentConfigUpdateItemIEs, order_E2nodeComponentConfigUpdateItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type E2nodeComponentConfigUpdateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['E2nodeComponentConfigUpdate-ItemIEs']}, 'size': [(1, 'maxofE2nodeComponents')]}
    Items []E2nodeComponentConfigUpdateItemIEs
}

type E2nodeComponentConfigUpdateItem struct { // [{'type': 'E2nodeComponentInterfaceType', 'name': 'e2nodeComponentInterfaceType'}, {'type': 'E2nodeComponentID', 'name': 'e2nodeComponentID'}, {'type': 'E2nodeComponentConfiguration', 'name': 'e2nodeComponentConfiguration'}, None]
    E2nodeComponentInterfaceType E2nodeComponentInterfaceType
    E2nodeComponentID E2nodeComponentID
    E2nodeComponentConfiguration E2nodeComponentConfiguration
}

func (self * E2nodeComponentConfigUpdateItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.E2nodeComponentInterfaceType.Unpack(stream)// p8
    self.E2nodeComponentID.Unpack(stream)// p8
    self.E2nodeComponentConfiguration.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentConfigUpdateItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.E2nodeComponentInterfaceType.Pack(stream)
    self.E2nodeComponentID.Pack(stream)
    self.E2nodeComponentConfiguration.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *E2nodeComponentConfigRemovalList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(1024)
    _size += 1
    self.Items = make([]E2nodeComponentConfigRemovalItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *E2nodeComponentConfigRemovalList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 1024)
    //for item in table_E2nodeComponentConfigRemovalItemIEs:
    val := ProtocolIESingleContainer{table_E2nodeComponentConfigRemovalItemIEs, order_E2nodeComponentConfigRemovalItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type E2nodeComponentConfigRemovalList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['E2nodeComponentConfigRemoval-ItemIEs']}, 'size': [(1, 'maxofE2nodeComponents')]}
    Items []E2nodeComponentConfigRemovalItemIEs
}

type E2nodeComponentConfigRemovalItem struct { // [{'type': 'E2nodeComponentInterfaceType', 'name': 'e2nodeComponentInterfaceType'}, {'type': 'E2nodeComponentID', 'name': 'e2nodeComponentID'}, None]
    E2nodeComponentInterfaceType E2nodeComponentInterfaceType
    E2nodeComponentID E2nodeComponentID
}

func (self * E2nodeComponentConfigRemovalItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.E2nodeComponentInterfaceType.Unpack(stream)// p8
    self.E2nodeComponentID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentConfigRemovalItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.E2nodeComponentInterfaceType.Pack(stream)
    self.E2nodeComponentID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *E2nodeTNLassociationRemovalList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]E2nodeTNLassociationRemovalItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *E2nodeTNLassociationRemovalList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    //for item in table_E2nodeTNLassociationRemovalItemIEs:
    val := ProtocolIESingleContainer{table_E2nodeTNLassociationRemovalItemIEs, order_E2nodeTNLassociationRemovalItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type E2nodeTNLassociationRemovalList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['E2nodeTNLassociationRemoval-ItemIEs']}, 'size': [(1, 'maxofTNLA')]}
    Items []E2nodeTNLassociationRemovalItemIEs
}

type E2nodeTNLassociationRemovalItem struct { // [{'type': 'TNLinformation', 'name': 'tnlInformation'}, {'type': 'TNLinformation', 'name': 'tnlInformationRIC'}, None]
    TnlInformation TNLinformation
    TnlInformationRIC TNLinformation
}

func (self * E2nodeTNLassociationRemovalItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.TnlInformation.Unpack(stream)// p8
    self.TnlInformationRIC.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeTNLassociationRemovalItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TnlInformation.Pack(stream)
    self.TnlInformationRIC.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2nodeConfigurationUpdateAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E2nodeConfigurationUpdateAcknowledge-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E2nodeConfigurationUpdateAcknowledgeIEs
}

func (self * E2nodeConfigurationUpdateAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E2nodeConfigurationUpdateAcknowledgeIEs, order_E2nodeConfigurationUpdateAcknowledgeIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeConfigurationUpdateAcknowledge) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E2nodeConfigurationUpdateAcknowledgeIEs, order_E2nodeConfigurationUpdateAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *E2nodeComponentConfigAdditionAckList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(1024)
    _size += 1
    self.Items = make([]E2nodeComponentConfigAdditionAckItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *E2nodeComponentConfigAdditionAckList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 1024)
    //for item in table_E2nodeComponentConfigAdditionAckItemIEs:
    val := ProtocolIESingleContainer{table_E2nodeComponentConfigAdditionAckItemIEs, order_E2nodeComponentConfigAdditionAckItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type E2nodeComponentConfigAdditionAckList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['E2nodeComponentConfigAdditionAck-ItemIEs']}, 'size': [(1, 'maxofE2nodeComponents')]}
    Items []E2nodeComponentConfigAdditionAckItemIEs
}

type E2nodeComponentConfigAdditionAckItem struct { // [{'type': 'E2nodeComponentInterfaceType', 'name': 'e2nodeComponentInterfaceType'}, {'type': 'E2nodeComponentID', 'name': 'e2nodeComponentID'}, {'type': 'E2nodeComponentConfigurationAck', 'name': 'e2nodeComponentConfigurationAck'}, None]
    E2nodeComponentInterfaceType E2nodeComponentInterfaceType
    E2nodeComponentID E2nodeComponentID
    E2nodeComponentConfigurationAck E2nodeComponentConfigurationAck
}

func (self * E2nodeComponentConfigAdditionAckItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.E2nodeComponentInterfaceType.Unpack(stream)// p8
    self.E2nodeComponentID.Unpack(stream)// p8
    self.E2nodeComponentConfigurationAck.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentConfigAdditionAckItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.E2nodeComponentInterfaceType.Pack(stream)
    self.E2nodeComponentID.Pack(stream)
    self.E2nodeComponentConfigurationAck.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *E2nodeComponentConfigUpdateAckList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(1024)
    _size += 1
    self.Items = make([]E2nodeComponentConfigUpdateAckItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *E2nodeComponentConfigUpdateAckList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 1024)
    //for item in table_E2nodeComponentConfigUpdateAckItemIEs:
    val := ProtocolIESingleContainer{table_E2nodeComponentConfigUpdateAckItemIEs, order_E2nodeComponentConfigUpdateAckItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type E2nodeComponentConfigUpdateAckList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['E2nodeComponentConfigUpdateAck-ItemIEs']}, 'size': [(1, 'maxofE2nodeComponents')]}
    Items []E2nodeComponentConfigUpdateAckItemIEs
}

type E2nodeComponentConfigUpdateAckItem struct { // [{'type': 'E2nodeComponentInterfaceType', 'name': 'e2nodeComponentInterfaceType'}, {'type': 'E2nodeComponentID', 'name': 'e2nodeComponentID'}, {'type': 'E2nodeComponentConfigurationAck', 'name': 'e2nodeComponentConfigurationAck'}, None]
    E2nodeComponentInterfaceType E2nodeComponentInterfaceType
    E2nodeComponentID E2nodeComponentID
    E2nodeComponentConfigurationAck E2nodeComponentConfigurationAck
}

func (self * E2nodeComponentConfigUpdateAckItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.E2nodeComponentInterfaceType.Unpack(stream)// p8
    self.E2nodeComponentID.Unpack(stream)// p8
    self.E2nodeComponentConfigurationAck.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentConfigUpdateAckItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.E2nodeComponentInterfaceType.Pack(stream)
    self.E2nodeComponentID.Pack(stream)
    self.E2nodeComponentConfigurationAck.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *E2nodeComponentConfigRemovalAckList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(1024)
    _size += 1
    self.Items = make([]E2nodeComponentConfigRemovalAckItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *E2nodeComponentConfigRemovalAckList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 1024)
    //for item in table_E2nodeComponentConfigRemovalAckItemIEs:
    val := ProtocolIESingleContainer{table_E2nodeComponentConfigRemovalAckItemIEs, order_E2nodeComponentConfigRemovalAckItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type E2nodeComponentConfigRemovalAckList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['E2nodeComponentConfigRemovalAck-ItemIEs']}, 'size': [(1, 'maxofE2nodeComponents')]}
    Items []E2nodeComponentConfigRemovalAckItemIEs
}

type E2nodeComponentConfigRemovalAckItem struct { // [{'type': 'E2nodeComponentInterfaceType', 'name': 'e2nodeComponentInterfaceType'}, {'type': 'E2nodeComponentID', 'name': 'e2nodeComponentID'}, {'type': 'E2nodeComponentConfigurationAck', 'name': 'e2nodeComponentConfigurationAck'}, None]
    E2nodeComponentInterfaceType E2nodeComponentInterfaceType
    E2nodeComponentID E2nodeComponentID
    E2nodeComponentConfigurationAck E2nodeComponentConfigurationAck
}

func (self * E2nodeComponentConfigRemovalAckItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.E2nodeComponentInterfaceType.Unpack(stream)// p8
    self.E2nodeComponentID.Unpack(stream)// p8
    self.E2nodeComponentConfigurationAck.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentConfigRemovalAckItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.E2nodeComponentInterfaceType.Pack(stream)
    self.E2nodeComponentID.Pack(stream)
    self.E2nodeComponentConfigurationAck.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2nodeConfigurationUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['E2nodeConfigurationUpdateFailure-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs E2nodeConfigurationUpdateFailureIEs
}

func (self * E2nodeConfigurationUpdateFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_E2nodeConfigurationUpdateFailureIEs, order_E2nodeConfigurationUpdateFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeConfigurationUpdateFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_E2nodeConfigurationUpdateFailureIEs, order_E2nodeConfigurationUpdateFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type ResetRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ResetRequestIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ResetRequestIEs
}

func (self * ResetRequest) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ResetRequestIEs, order_ResetRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ResetRequest) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ResetRequestIEs, order_ResetRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type ResetResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ResetResponseIEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs ResetResponseIEs
}

func (self * ResetResponse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_ResetResponseIEs, order_ResetResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ResetResponse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ResetResponseIEs, order_ResetResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICserviceUpdate struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICserviceUpdate-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICserviceUpdateIEs
}

func (self * RICserviceUpdate) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICserviceUpdateIEs, order_RICserviceUpdateIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICserviceUpdate) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICserviceUpdateIEs, order_RICserviceUpdateIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *RANfunctionsList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(257)
    _size += 0
    self.Items = make([]RANfunctionItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *RANfunctionsList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-0, 257)
    //for item in table_RANfunctionItemIEs:
    val := ProtocolIESingleContainer{table_RANfunctionItemIEs, order_RANfunctionItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type RANfunctionsList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['RANfunction-ItemIEs']}, 'size': [(0, 'maxofRANfunctionID')]}
    Items []RANfunctionItemIEs
}

type RANfunctionItem struct { // [{'type': 'RANfunctionID', 'name': 'ranFunctionID'}, {'type': 'RANfunctionDefinition', 'name': 'ranFunctionDefinition'}, {'type': 'RANfunctionRevision', 'name': 'ranFunctionRevision'}, {'type': 'RANfunctionOID', 'name': 'ranFunctionOID'}, None]
    RanFunctionID RANfunctionID
    RanFunctionDefinition RANfunctionDefinition
    RanFunctionRevision RANfunctionRevision
    RanFunctionOID RANfunctionOID
}

func (self * RANfunctionItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanFunctionID.Unpack(stream)// p8
    self.RanFunctionDefinition.Unpack(stream)// p8
    self.RanFunctionRevision.Unpack(stream)// p8
    self.RanFunctionOID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANfunctionItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanFunctionID.Pack(stream)
    self.RanFunctionDefinition.Pack(stream)
    self.RanFunctionRevision.Pack(stream)
    self.RanFunctionOID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *RANfunctionsIDList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]RANfunctionIDItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *RANfunctionsIDList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    //for item in table_RANfunctionIDItemIEs:
    val := ProtocolIESingleContainer{table_RANfunctionIDItemIEs, order_RANfunctionIDItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type RANfunctionsIDList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['RANfunctionID-ItemIEs']}, 'size': [(1, 'maxofRANfunctionID')]}
    Items []RANfunctionIDItemIEs
}

type RANfunctionIDItem struct { // [{'type': 'RANfunctionID', 'name': 'ranFunctionID'}, {'type': 'RANfunctionRevision', 'name': 'ranFunctionRevision'}, None]
    RanFunctionID RANfunctionID
    RanFunctionRevision RANfunctionRevision
}

func (self * RANfunctionIDItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanFunctionID.Unpack(stream)// p8
    self.RanFunctionRevision.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANfunctionIDItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanFunctionID.Pack(stream)
    self.RanFunctionRevision.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICserviceUpdateAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICserviceUpdateAcknowledge-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICserviceUpdateAcknowledgeIEs
}

func (self * RICserviceUpdateAcknowledge) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICserviceUpdateAcknowledgeIEs, order_RICserviceUpdateAcknowledgeIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICserviceUpdateAcknowledge) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICserviceUpdateAcknowledgeIEs, order_RICserviceUpdateAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *RANfunctionsIDcauseList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]RANfunctionIDcauseItemIEs, _size)
    val := ProtocolIESingleContainer{}
    for i :=0 ; i < _size; i++ { //aparams
        val.Unpack(stream, &self.Items[i])
    }
}


func (self *RANfunctionsIDcauseList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    //for item in table_RANfunctionIDcauseItemIEs:
    val := ProtocolIESingleContainer{table_RANfunctionIDcauseItemIEs, order_RANfunctionIDcauseItemIEs}
    for _, item := range self.Items {
        val.Pack(stream, &item)
    }
    return

}


type RANfunctionsIDcauseList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['RANfunctionIDcause-ItemIEs']}, 'size': [(1, 'maxofRANfunctionID')]}
    Items []RANfunctionIDcauseItemIEs
}

type RANfunctionIDcauseItem struct { // [{'type': 'RANfunctionID', 'name': 'ranFunctionID'}, {'type': 'Cause', 'name': 'cause'}, None]
    RanFunctionID RANfunctionID
    Cause Cause
}

func (self * RANfunctionIDcauseItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanFunctionID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANfunctionIDcauseItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanFunctionID.Pack(stream)
    self.Cause.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICserviceUpdateFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICserviceUpdateFailure-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICserviceUpdateFailureIEs
}

func (self * RICserviceUpdateFailure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICserviceUpdateFailureIEs, order_RICserviceUpdateFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICserviceUpdateFailure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICserviceUpdateFailureIEs, order_RICserviceUpdateFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICserviceQuery struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RICserviceQuery-IEs'], 'name': 'protocolIEs'}, None]
    ProtocolIEs RICserviceQueryIEs
}

func (self * RICserviceQuery) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    ProtocolIEs := ProtocolIEContainer {table_RICserviceQueryIEs, order_RICserviceQueryIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICserviceQuery) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RICserviceQueryIEs, order_RICserviceQueryIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type AMFName struct {
  Value string
}
func (self *AMFName) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in AMFName")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *AMFName) Pack(st *Stream) {
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
func (self *Cause)Unpack(stream *Stream) {
    //coptions := []string{"ricRequest","ricService","e2Node","transport","protocol","misc","Unknown","Unknown"}
    choice := stream.get_choice(3, 1, 6)
    choice_len := 0
    choice_loc := 0
    if choice >= 6 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in Cause\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.RicRequest = &CauseRICrequest{}//cho6
        self.RicRequest.Unpack(stream)
    } else if choice == 1 { //ch2
        self.RicService = &CauseRICservice{}//cho6
        self.RicService.Unpack(stream)
    } else if choice == 2 { //ch2
        self.E2Node = &CauseE2node{}//cho6
        self.E2Node.Unpack(stream)
    } else if choice == 3 { //ch2
        self.Transport = &CauseTransport{}//cho6
        self.Transport.Unpack(stream)
    } else if choice == 4 { //ch2
        self.Protocol = &CauseProtocol{}//cho6
        self.Protocol.Unpack(stream)
    } else if choice == 5 { //ch2
        self.Misc = &CauseMisc{}//cho6
        self.Misc.Unpack(stream)
    }//end of if else

    if choice >= 6 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * Cause) Pack(stream *Stream) {
    if self.RicRequest != nil {
        stream.set_choice(0, 3, 1, 6)
        self.RicRequest.Pack(stream)//2
    } else if self.RicService != nil {
        stream.set_choice(1, 3, 1, 6)
        self.RicService.Pack(stream)//2
    } else if self.E2Node != nil {
        stream.set_choice(2, 3, 1, 6)
        self.E2Node.Pack(stream)//2
    } else if self.Transport != nil {
        stream.set_choice(3, 3, 1, 6)
        self.Transport.Pack(stream)//2
    } else if self.Protocol != nil {
        stream.set_choice(4, 3, 1, 6)
        self.Protocol.Pack(stream)//2
    } else if self.Misc != nil {
        stream.set_choice(5, 3, 1, 6)
        self.Misc.Pack(stream)//2
    }

}
type Cause struct { //[{'type': 'CauseRICrequest', 'name': 'ricRequest'}, {'type': 'CauseRICservice', 'name': 'ricService'}, {'type': 'CauseE2node', 'name': 'e2Node'}, {'type': 'CauseTransport', 'name': 'transport'}, {'type': 'CauseProtocol', 'name': 'protocol'}, {'type': 'CauseMisc', 'name': 'misc'}, None]
    RicRequest *CauseRICrequest
    RicService *CauseRICservice
    E2Node *CauseE2node
    Transport *CauseTransport
    Protocol *CauseProtocol
    Misc *CauseMisc
} // Cause

type CauseE2node struct {
  Value int
}
const (
    CauseE2nodee2node_component_unknown = 0

    /* Extensions */
)
func (self *CauseE2node) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *CauseE2node) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
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
type CauseRICrequest struct {
  Value int
}
const (
    CauseRICrequestran_function_id_invalid = 0
    CauseRICrequestaction_not_supported = 1
    CauseRICrequestexcessive_actions = 2
    CauseRICrequestduplicate_action = 3
    CauseRICrequestduplicate_event_trigger = 4
    CauseRICrequestfunction_resource_limit = 5
    CauseRICrequestrequest_id_unknown = 6
    CauseRICrequestinconsistent_action_subsequent_action_sequence = 7
    CauseRICrequestcontrol_message_invalid = 8
    CauseRICrequestric_call_process_id_invalid = 9
    CauseRICrequestcontrol_timer_expired = 10
    CauseRICrequestcontrol_failed_to_execute = 11
    CauseRICrequestsystem_not_ready = 12
    CauseRICrequestunspecified = 13

    /* Extensions */
)
func (self *CauseRICrequest) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(5, 14, 1)
}
func (self *CauseRICrequest) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 5, 14, 1)
}
type CauseRICservice struct {
  Value int
}
const (
    CauseRICserviceran_function_not_supported = 0
    CauseRICserviceexcessive_functions = 1
    CauseRICserviceric_resource_limit = 2

    /* Extensions */
)
func (self *CauseRICservice) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *CauseRICservice) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
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
type CriticalityDiagnostics struct { // [{'type': 'ProcedureCode', 'name': 'procedureCode', 'optional': True}, {'type': 'TriggeringMessage', 'name': 'triggeringMessage', 'optional': True}, {'type': 'Criticality', 'name': 'procedureCriticality', 'optional': True}, {'type': 'RICrequestID', 'name': 'ricRequestorID', 'optional': True}, {'type': 'CriticalityDiagnostics-IE-List', 'name': 'iEsCriticalityDiagnostics', 'optional': True}, None]
    ProcedureCode *ProcedureCode
    TriggeringMessage *TriggeringMessage
    ProcedureCriticality *Criticality
    RicRequestorID *RICrequestID
    IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList
}

func (self * CriticalityDiagnostics) Unpack(stream *Stream) {
    procedureCode_flag := 0x00000002
    triggeringMessage_flag := 0x00000004
    procedureCriticality_flag := 0x00000008
    ricRequestorID_flag := 0x00000010
    iEsCriticalityDiagnostics_flag := 0x00000020
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
    if (ricRequestorID_flag & _flags) == ricRequestorID_flag { //cond2
        self.RicRequestorID = &RICrequestID{}//7{'type': 'RICrequestID', 'name': 'ricRequestorID', 'optional': True}
        self.RicRequestorID.Unpack(stream)// p8
    }
    if (iEsCriticalityDiagnostics_flag & _flags) == iEsCriticalityDiagnostics_flag { //cond2
        self.IEsCriticalityDiagnostics = &CriticalityDiagnosticsIEList{}//7{'type': 'CriticalityDiagnostics-IE-List', 'name': 'iEsCriticalityDiagnostics', 'optional': True}
        self.IEsCriticalityDiagnostics.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CriticalityDiagnostics) Pack(stream *Stream) {
    const procedureCode_flag uint = 0x00000002
    const triggeringMessage_flag uint = 0x00000004
    const procedureCriticality_flag uint = 0x00000008
    const ricRequestorID_flag uint = 0x00000010
    const iEsCriticalityDiagnostics_flag uint = 0x00000020
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
    if self.RicRequestorID != nil { 
        _flags |= ricRequestorID_flag
        self.RicRequestorID.Pack(stream)
    }//end of optional
    if self.IEsCriticalityDiagnostics != nil { 
        _flags |= iEsCriticalityDiagnostics_flag
        self.IEsCriticalityDiagnostics.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

func (self *CriticalityDiagnosticsIEList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]CriticalityDiagnosticsIEItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *CriticalityDiagnosticsIEList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type CriticalityDiagnosticsIEList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'CriticalityDiagnostics-IE-Item'}, 'size': [(1, 'maxnoofErrors')]}
    Items []CriticalityDiagnosticsIEItem
}

type CriticalityDiagnosticsIEItem struct { // [{'type': 'Criticality', 'name': 'iECriticality'}, {'type': 'ProtocolIE-ID', 'name': 'iE-ID'}, {'type': 'TypeOfError', 'name': 'typeOfError'}, None]
    IECriticality Criticality
    IEID ProtocolIEID
    TypeOfError TypeOfError
}

func (self * CriticalityDiagnosticsIEItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.IECriticality.Unpack(stream)// p8
    self.IEID.Unpack(stream)// p8
    self.TypeOfError.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CriticalityDiagnosticsIEItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IECriticality.Pack(stream)
    self.IEID.Pack(stream)
    self.TypeOfError.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2nodeComponentConfiguration struct { // [{'type': 'OCTET STRING', 'name': 'e2nodeComponentRequestPart'}, {'type': 'OCTET STRING', 'name': 'e2nodeComponentResponsePart'}, None]
    E2nodeComponentRequestPart OCTETSTRING
    E2nodeComponentResponsePart OCTETSTRING
}

func (self * E2nodeComponentConfiguration) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_e2nodeComponentRequestPart = func(st *Stream, self *OCTETSTRING) {
        _len := st.parse_len(0)
        self.Value = st.parsef_OctString(_len)
    }
    Unpack_e2nodeComponentRequestPart(stream, &self.E2nodeComponentRequestPart)// p2
    var Unpack_e2nodeComponentResponsePart = func(st *Stream, self *OCTETSTRING) {
        _len := st.parse_len(0)
        self.Value = st.parsef_OctString(_len)
    }
    Unpack_e2nodeComponentResponsePart(stream, &self.E2nodeComponentResponsePart)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentConfiguration) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_e2nodeComponentRequestPart = func(st *Stream, self OCTETSTRING) {
        st.format_len(len(self.Value), 0)
        st.formatf_OctString(self.Value, 0)
    }
    Pack_e2nodeComponentRequestPart(stream, self.E2nodeComponentRequestPart) //f2
    var Pack_e2nodeComponentResponsePart = func(st *Stream, self OCTETSTRING) {
        st.format_len(len(self.Value), 0)
        st.formatf_OctString(self.Value, 0)
    }
    Pack_e2nodeComponentResponsePart(stream, self.E2nodeComponentResponsePart) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2nodeComponentConfigurationAck struct { // [{'type': 'ENUMERATED', 'values': [('success', 0), ('failure', 1), None], 'name': 'updateOutcome'}, {'type': 'Cause', 'name': 'failureCause', 'optional': True}, None]
    UpdateOutcome ENUMERATED
    FailureCause *Cause
}

func (self * E2nodeComponentConfigurationAck) Unpack(stream *Stream) {
    failureCause_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_updateOutcome = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_updateOutcome(stream, &self.UpdateOutcome)// p2
    if (failureCause_flag & _flags) == failureCause_flag { //cond2
        self.FailureCause = &Cause{}//7{'type': 'Cause', 'name': 'failureCause', 'optional': True}
        self.FailureCause.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentConfigurationAck) Pack(stream *Stream) {
    const failureCause_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_updateOutcome = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_updateOutcome(stream, self.UpdateOutcome) //f2
    if self.FailureCause != nil { 
        _flags |= failureCause_flag
        self.FailureCause.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2nodeComponentInterfaceType struct {
  Value int
}
const (
    E2nodeComponentInterfaceTypeng = 0
    E2nodeComponentInterfaceTypexn = 1
    E2nodeComponentInterfaceTypee1 = 2
    E2nodeComponentInterfaceTypef1 = 3
    E2nodeComponentInterfaceTypew1 = 4
    E2nodeComponentInterfaceTypes1 = 5
    E2nodeComponentInterfaceTypex2 = 6

    /* Extensions */
)
func (self *E2nodeComponentInterfaceType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 7, 1)
}
func (self *E2nodeComponentInterfaceType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 7, 1)
}
func (self *E2nodeComponentID)Unpack(stream *Stream) {
    //coptions := []string{"e2nodeComponentInterfaceTypeNG","e2nodeComponentInterfaceTypeXn","e2nodeComponentInterfaceTypeE1","e2nodeComponentInterfaceTypeF1","e2nodeComponentInterfaceTypeW1","e2nodeComponentInterfaceTypeS1","e2nodeComponentInterfaceTypeX2","Unknown"}
    choice := stream.get_choice(3, 1, 7)
    choice_len := 0
    choice_loc := 0
    if choice >= 7 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in E2nodeComponentID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.E2nodeComponentInterfaceTypeNG = &E2nodeComponentInterfaceNG{}//cho6
        self.E2nodeComponentInterfaceTypeNG.Unpack(stream)
    } else if choice == 1 { //ch2
        self.E2nodeComponentInterfaceTypeXn = &E2nodeComponentInterfaceXn{}//cho6
        self.E2nodeComponentInterfaceTypeXn.Unpack(stream)
    } else if choice == 2 { //ch2
        self.E2nodeComponentInterfaceTypeE1 = &E2nodeComponentInterfaceE1{}//cho6
        self.E2nodeComponentInterfaceTypeE1.Unpack(stream)
    } else if choice == 3 { //ch2
        self.E2nodeComponentInterfaceTypeF1 = &E2nodeComponentInterfaceF1{}//cho6
        self.E2nodeComponentInterfaceTypeF1.Unpack(stream)
    } else if choice == 4 { //ch2
        self.E2nodeComponentInterfaceTypeW1 = &E2nodeComponentInterfaceW1{}//cho6
        self.E2nodeComponentInterfaceTypeW1.Unpack(stream)
    } else if choice == 5 { //ch2
        self.E2nodeComponentInterfaceTypeS1 = &E2nodeComponentInterfaceS1{}//cho6
        self.E2nodeComponentInterfaceTypeS1.Unpack(stream)
    } else if choice == 6 { //ch2
        self.E2nodeComponentInterfaceTypeX2 = &E2nodeComponentInterfaceX2{}//cho6
        self.E2nodeComponentInterfaceTypeX2.Unpack(stream)
    }//end of if else

    if choice >= 7 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * E2nodeComponentID) Pack(stream *Stream) {
    if self.E2nodeComponentInterfaceTypeNG != nil {
        stream.set_choice(0, 3, 1, 7)
        self.E2nodeComponentInterfaceTypeNG.Pack(stream)//2
    } else if self.E2nodeComponentInterfaceTypeXn != nil {
        stream.set_choice(1, 3, 1, 7)
        self.E2nodeComponentInterfaceTypeXn.Pack(stream)//2
    } else if self.E2nodeComponentInterfaceTypeE1 != nil {
        stream.set_choice(2, 3, 1, 7)
        self.E2nodeComponentInterfaceTypeE1.Pack(stream)//2
    } else if self.E2nodeComponentInterfaceTypeF1 != nil {
        stream.set_choice(3, 3, 1, 7)
        self.E2nodeComponentInterfaceTypeF1.Pack(stream)//2
    } else if self.E2nodeComponentInterfaceTypeW1 != nil {
        stream.set_choice(4, 3, 1, 7)
        self.E2nodeComponentInterfaceTypeW1.Pack(stream)//2
    } else if self.E2nodeComponentInterfaceTypeS1 != nil {
        stream.set_choice(5, 3, 1, 7)
        self.E2nodeComponentInterfaceTypeS1.Pack(stream)//2
    } else if self.E2nodeComponentInterfaceTypeX2 != nil {
        stream.set_choice(6, 3, 1, 7)
        self.E2nodeComponentInterfaceTypeX2.Pack(stream)//2
    }

}
type E2nodeComponentID struct { //[{'type': 'E2nodeComponentInterfaceNG', 'name': 'e2nodeComponentInterfaceTypeNG'}, {'type': 'E2nodeComponentInterfaceXn', 'name': 'e2nodeComponentInterfaceTypeXn'}, {'type': 'E2nodeComponentInterfaceE1', 'name': 'e2nodeComponentInterfaceTypeE1'}, {'type': 'E2nodeComponentInterfaceF1', 'name': 'e2nodeComponentInterfaceTypeF1'}, {'type': 'E2nodeComponentInterfaceW1', 'name': 'e2nodeComponentInterfaceTypeW1'}, {'type': 'E2nodeComponentInterfaceS1', 'name': 'e2nodeComponentInterfaceTypeS1'}, {'type': 'E2nodeComponentInterfaceX2', 'name': 'e2nodeComponentInterfaceTypeX2'}, None]
    E2nodeComponentInterfaceTypeNG *E2nodeComponentInterfaceNG
    E2nodeComponentInterfaceTypeXn *E2nodeComponentInterfaceXn
    E2nodeComponentInterfaceTypeE1 *E2nodeComponentInterfaceE1
    E2nodeComponentInterfaceTypeF1 *E2nodeComponentInterfaceF1
    E2nodeComponentInterfaceTypeW1 *E2nodeComponentInterfaceW1
    E2nodeComponentInterfaceTypeS1 *E2nodeComponentInterfaceS1
    E2nodeComponentInterfaceTypeX2 *E2nodeComponentInterfaceX2
} // E2nodeComponentID

type E2nodeComponentInterfaceE1 struct { // [{'type': 'GNB-CU-UP-ID', 'name': 'gNB-CU-CP-ID'}, None]
    GNBCUCPID GNBCUUPID
}

func (self * E2nodeComponentInterfaceE1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GNBCUCPID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentInterfaceE1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GNBCUCPID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2nodeComponentInterfaceF1 struct { // [{'type': 'GNB-DU-ID', 'name': 'gNB-DU-ID'}, None]
    GNBDUID GNBDUID
}

func (self * E2nodeComponentInterfaceF1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GNBDUID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentInterfaceF1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GNBDUID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2nodeComponentInterfaceNG struct { // [{'type': 'AMFName', 'name': 'amf-name'}, None]
    Amfname AMFName
}

func (self * E2nodeComponentInterfaceNG) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.Amfname.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentInterfaceNG) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.Amfname.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2nodeComponentInterfaceS1 struct { // [{'type': 'MMEname', 'name': 'mme-name'}, None]
    Mmename MMEname
}

func (self * E2nodeComponentInterfaceS1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.Mmename.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentInterfaceS1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.Mmename.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2nodeComponentInterfaceX2 struct { // [{'type': 'GlobalENB-ID', 'name': 'global-eNB-ID', 'optional': True}, {'type': 'GlobalenGNB-ID', 'name': 'global-en-gNB-ID', 'optional': True}, None]
    GlobaleNBID *GlobalENBID
    GlobalengNBID *GlobalenGNBID
}

func (self * E2nodeComponentInterfaceX2) Unpack(stream *Stream) {
    globaleNBID_flag := 0x00000002
    globalengNBID_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    if (globaleNBID_flag & _flags) == globaleNBID_flag { //cond2
        self.GlobaleNBID = &GlobalENBID{}//7{'type': 'GlobalENB-ID', 'name': 'global-eNB-ID', 'optional': True}
        self.GlobaleNBID.Unpack(stream)// p8
    }
    if (globalengNBID_flag & _flags) == globalengNBID_flag { //cond2
        self.GlobalengNBID = &GlobalenGNBID{}//7{'type': 'GlobalenGNB-ID', 'name': 'global-en-gNB-ID', 'optional': True}
        self.GlobalengNBID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentInterfaceX2) Pack(stream *Stream) {
    const globaleNBID_flag uint = 0x00000002
    const globalengNBID_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.GlobaleNBID != nil { 
        _flags |= globaleNBID_flag
        self.GlobaleNBID.Pack(stream)
    }//end of optional
    if self.GlobalengNBID != nil { 
        _flags |= globalengNBID_flag
        self.GlobalengNBID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type E2nodeComponentInterfaceXn struct { // [{'type': 'GlobalNG-RANNode-ID', 'name': 'global-NG-RAN-Node-ID'}, None]
    GlobalNGRANNodeID GlobalNGRANNodeID
}

func (self * E2nodeComponentInterfaceXn) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GlobalNGRANNodeID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentInterfaceXn) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GlobalNGRANNodeID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2nodeComponentInterfaceW1 struct { // [{'type': 'NGENB-DU-ID', 'name': 'ng-eNB-DU-ID'}, None]
    NgeNBDUID NGENBDUID
}

func (self * E2nodeComponentInterfaceW1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.NgeNBDUID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2nodeComponentInterfaceW1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.NgeNBDUID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *ENBID)Unpack(stream *Stream) {
    //coptions := []string{"macro-eNB-ID","home-eNB-ID","short-Macro-eNB-ID","long-Macro-eNB-ID"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
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
        var Unpack_homeeNBID = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(28, 28)
        }
        self.HomeeNBID = &BITSTRING{}//cho5
        Unpack_homeeNBID(stream, self.HomeeNBID);
    } else if choice == 2 { //ch2
        var Unpack_shortMacroeNBID = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(18, 18)
        }
        self.ShortMacroeNBID = &BITSTRING{}//cho5
        Unpack_shortMacroeNBID(stream, self.ShortMacroeNBID);
    } else if choice == 3 { //ch2
        var Unpack_longMacroeNBID = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(21, 21)
        }
        self.LongMacroeNBID = &BITSTRING{}//cho5
        Unpack_longMacroeNBID(stream, self.LongMacroeNBID);
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ENBID) Pack(stream *Stream) {
    if self.MacroeNBID != nil {
        stream.set_choice(0, 1, 1, 2)
        var Pack_macroeNBID = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 20)
        }
        Pack_macroeNBID(stream, *self.MacroeNBID)//3
    } else if self.HomeeNBID != nil {
        stream.set_choice(1, 1, 1, 2)
        var Pack_homeeNBID = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 28)
        }
        Pack_homeeNBID(stream, *self.HomeeNBID)//3
    } else if self.ShortMacroeNBID != nil {
        stream.set_choice(2, 1, 1, 2)
        lenLoc := stream.reserve_len()
        var Pack_shortMacroeNBID = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 18)
        }
        Pack_shortMacroeNBID(stream, *self.ShortMacroeNBID)//3
        stream.set_len(lenLoc)
    } else if self.LongMacroeNBID != nil {
        stream.set_choice(3, 1, 1, 2)
        lenLoc := stream.reserve_len()
        var Pack_longMacroeNBID = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 21)
        }
        Pack_longMacroeNBID(stream, *self.LongMacroeNBID)//3
        stream.set_len(lenLoc)
    }

}
type ENBID struct { //[{'type': 'BIT STRING', 'size': [20], 'name': 'macro-eNB-ID'}, {'type': 'BIT STRING', 'size': [28], 'name': 'home-eNB-ID'}, None, {'type': 'BIT STRING', 'size': [18], 'name': 'short-Macro-eNB-ID'}, {'type': 'BIT STRING', 'size': [21], 'name': 'long-Macro-eNB-ID'}]
    MacroeNBID *BITSTRING
    HomeeNBID *BITSTRING
    ShortMacroeNBID *BITSTRING
    LongMacroeNBID *BITSTRING
} // ENBID

func (self *ENBIDChoice)Unpack(stream *Stream) {
    //coptions := []string{"enb-ID-macro","enb-ID-shortmacro","enb-ID-longmacro","Unknown"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in ENBIDChoice\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_enbIDmacro = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(20, 20)
        }
        self.EnbIDmacro = &BITSTRING{}//cho5
        Unpack_enbIDmacro(stream, self.EnbIDmacro);
    } else if choice == 1 { //ch2
        var Unpack_enbIDshortmacro = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(18, 18)
        }
        self.EnbIDshortmacro = &BITSTRING{}//cho5
        Unpack_enbIDshortmacro(stream, self.EnbIDshortmacro);
    } else if choice == 2 { //ch2
        var Unpack_enbIDlongmacro = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(21, 21)
        }
        self.EnbIDlongmacro = &BITSTRING{}//cho5
        Unpack_enbIDlongmacro(stream, self.EnbIDlongmacro);
    }//end of if else

    if choice >= 3 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ENBIDChoice) Pack(stream *Stream) {
    if self.EnbIDmacro != nil {
        stream.set_choice(0, 2, 1, 3)
        var Pack_enbIDmacro = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 20)
        }
        Pack_enbIDmacro(stream, *self.EnbIDmacro)//3
    } else if self.EnbIDshortmacro != nil {
        stream.set_choice(1, 2, 1, 3)
        var Pack_enbIDshortmacro = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 18)
        }
        Pack_enbIDshortmacro(stream, *self.EnbIDshortmacro)//3
    } else if self.EnbIDlongmacro != nil {
        stream.set_choice(2, 2, 1, 3)
        var Pack_enbIDlongmacro = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 21)
        }
        Pack_enbIDlongmacro(stream, *self.EnbIDlongmacro)//3
    }

}
type ENBIDChoice struct { //[{'type': 'BIT STRING', 'size': [20], 'name': 'enb-ID-macro'}, {'type': 'BIT STRING', 'size': [18], 'name': 'enb-ID-shortmacro'}, {'type': 'BIT STRING', 'size': [21], 'name': 'enb-ID-longmacro'}, None]
    EnbIDmacro *BITSTRING
    EnbIDshortmacro *BITSTRING
    EnbIDlongmacro *BITSTRING
} // ENBIDChoice

func (self *ENGNBID)Unpack(stream *Stream) {
    //coptions := []string{"gNB-ID"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in ENGNBID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_gNBID = func(st *Stream, self *BITSTRING){
            self.Len = int(st.parse_blen(4, 0)+22)
            self.Value = st.parsef_BitString(32, int(self.Len))
        }
        self.GNBID = &BITSTRING{}//cho5
        Unpack_gNBID(stream, self.GNBID);
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ENGNBID) Pack(stream *Stream) {
    if self.GNBID != nil {
        stream.set_choice(0, 0, 1, 1)
        var Pack_gNBID = func(st *Stream, self BITSTRING) {
            st.format_blen(int(self.Len-22), 4, 0)
            st.formatf_BitString(self.Value, int(self.Len))
        }
        Pack_gNBID(stream, *self.GNBID)//3
    }

}
type ENGNBID struct { //[{'type': 'BIT STRING', 'size': [(22, 32)], 'name': 'gNB-ID'}, None]
    GNBID *BITSTRING
} // ENGNBID

func (self *GlobalE2nodeID)Unpack(stream *Stream) {
    //coptions := []string{"gNB","en-gNB","ng-eNB","eNB"}
    choice := stream.get_choice(2, 1, 4)
    choice_len := 0
    choice_loc := 0
    if choice >= 4 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in GlobalE2nodeID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.GNB = &GlobalE2nodegNBID{}//cho6
        self.GNB.Unpack(stream)
    } else if choice == 1 { //ch2
        self.EngNB = &GlobalE2nodeengNBID{}//cho6
        self.EngNB.Unpack(stream)
    } else if choice == 2 { //ch2
        self.NgeNB = &GlobalE2nodengeNBID{}//cho6
        self.NgeNB.Unpack(stream)
    } else if choice == 3 { //ch2
        self.ENB = &GlobalE2nodeeNBID{}//cho6
        self.ENB.Unpack(stream)
    }//end of if else

    if choice >= 4 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * GlobalE2nodeID) Pack(stream *Stream) {
    if self.GNB != nil {
        stream.set_choice(0, 2, 1, 4)
        self.GNB.Pack(stream)//2
    } else if self.EngNB != nil {
        stream.set_choice(1, 2, 1, 4)
        self.EngNB.Pack(stream)//2
    } else if self.NgeNB != nil {
        stream.set_choice(2, 2, 1, 4)
        self.NgeNB.Pack(stream)//2
    } else if self.ENB != nil {
        stream.set_choice(3, 2, 1, 4)
        self.ENB.Pack(stream)//2
    }

}
type GlobalE2nodeID struct { //[{'type': 'GlobalE2node-gNB-ID', 'name': 'gNB'}, {'type': 'GlobalE2node-en-gNB-ID', 'name': 'en-gNB'}, {'type': 'GlobalE2node-ng-eNB-ID', 'name': 'ng-eNB'}, {'type': 'GlobalE2node-eNB-ID', 'name': 'eNB'}, None]
    GNB *GlobalE2nodegNBID
    EngNB *GlobalE2nodeengNBID
    NgeNB *GlobalE2nodengeNBID
    ENB *GlobalE2nodeeNBID
} // GlobalE2nodeID

type GlobalE2nodeengNBID struct { // [{'type': 'GlobalenGNB-ID', 'name': 'global-en-gNB-ID'}, {'type': 'GNB-CU-UP-ID', 'name': 'en-gNB-CU-UP-ID', 'optional': True}, {'type': 'GNB-DU-ID', 'name': 'en-gNB-DU-ID', 'optional': True}, None]
    GlobalengNBID GlobalenGNBID
    EngNBCUUPID *GNBCUUPID
    EngNBDUID *GNBDUID
}

func (self * GlobalE2nodeengNBID) Unpack(stream *Stream) {
    engNBCUUPID_flag := 0x00000002
    engNBDUID_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.GlobalengNBID.Unpack(stream)// p8
    if (engNBCUUPID_flag & _flags) == engNBCUUPID_flag { //cond2
        self.EngNBCUUPID = &GNBCUUPID{}//7{'type': 'GNB-CU-UP-ID', 'name': 'en-gNB-CU-UP-ID', 'optional': True}
        self.EngNBCUUPID.Unpack(stream)// p8
    }
    if (engNBDUID_flag & _flags) == engNBDUID_flag { //cond2
        self.EngNBDUID = &GNBDUID{}//7{'type': 'GNB-DU-ID', 'name': 'en-gNB-DU-ID', 'optional': True}
        self.EngNBDUID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalE2nodeengNBID) Pack(stream *Stream) {
    const engNBCUUPID_flag uint = 0x00000002
    const engNBDUID_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GlobalengNBID.Pack(stream)
    if self.EngNBCUUPID != nil { 
        _flags |= engNBCUUPID_flag
        self.EngNBCUUPID.Pack(stream)
    }//end of optional
    if self.EngNBDUID != nil { 
        _flags |= engNBDUID_flag
        self.EngNBDUID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type GlobalE2nodeeNBID struct { // [{'type': 'GlobalENB-ID', 'name': 'global-eNB-ID'}, None]
    GlobaleNBID GlobalENBID
}

func (self * GlobalE2nodeeNBID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GlobaleNBID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalE2nodeeNBID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GlobaleNBID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GlobalE2nodegNBID struct { // [{'type': 'GlobalgNB-ID', 'name': 'global-gNB-ID'}, {'type': 'GlobalenGNB-ID', 'name': 'global-en-gNB-ID', 'optional': True}, {'type': 'GNB-CU-UP-ID', 'name': 'gNB-CU-UP-ID', 'optional': True}, {'type': 'GNB-DU-ID', 'name': 'gNB-DU-ID', 'optional': True}, None]
    GlobalgNBID GlobalgNBID
    GlobalengNBID *GlobalenGNBID
    GNBCUUPID *GNBCUUPID
    GNBDUID *GNBDUID
}

func (self * GlobalE2nodegNBID) Unpack(stream *Stream) {
    globalengNBID_flag := 0x00000002
    gNBCUUPID_flag := 0x00000004
    gNBDUID_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.GlobalgNBID.Unpack(stream)// p8
    if (globalengNBID_flag & _flags) == globalengNBID_flag { //cond2
        self.GlobalengNBID = &GlobalenGNBID{}//7{'type': 'GlobalenGNB-ID', 'name': 'global-en-gNB-ID', 'optional': True}
        self.GlobalengNBID.Unpack(stream)// p8
    }
    if (gNBCUUPID_flag & _flags) == gNBCUUPID_flag { //cond2
        self.GNBCUUPID = &GNBCUUPID{}//7{'type': 'GNB-CU-UP-ID', 'name': 'gNB-CU-UP-ID', 'optional': True}
        self.GNBCUUPID.Unpack(stream)// p8
    }
    if (gNBDUID_flag & _flags) == gNBDUID_flag { //cond2
        self.GNBDUID = &GNBDUID{}//7{'type': 'GNB-DU-ID', 'name': 'gNB-DU-ID', 'optional': True}
        self.GNBDUID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalE2nodegNBID) Pack(stream *Stream) {
    const globalengNBID_flag uint = 0x00000002
    const gNBCUUPID_flag uint = 0x00000004
    const gNBDUID_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GlobalgNBID.Pack(stream)
    if self.GlobalengNBID != nil { 
        _flags |= globalengNBID_flag
        self.GlobalengNBID.Pack(stream)
    }//end of optional
    if self.GNBCUUPID != nil { 
        _flags |= gNBCUUPID_flag
        self.GNBCUUPID.Pack(stream)
    }//end of optional
    if self.GNBDUID != nil { 
        _flags |= gNBDUID_flag
        self.GNBDUID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type GlobalE2nodengeNBID struct { // [{'type': 'GlobalngeNB-ID', 'name': 'global-ng-eNB-ID'}, {'type': 'GlobalENB-ID', 'name': 'global-eNB-ID', 'optional': True}, {'type': 'NGENB-DU-ID', 'name': 'ngENB-DU-ID', 'optional': True}, None]
    GlobalngeNBID GlobalngeNBID
    GlobaleNBID *GlobalENBID
    NgENBDUID *NGENBDUID
}

func (self * GlobalE2nodengeNBID) Unpack(stream *Stream) {
    globaleNBID_flag := 0x00000002
    ngENBDUID_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.GlobalngeNBID.Unpack(stream)// p8
    if (globaleNBID_flag & _flags) == globaleNBID_flag { //cond2
        self.GlobaleNBID = &GlobalENBID{}//7{'type': 'GlobalENB-ID', 'name': 'global-eNB-ID', 'optional': True}
        self.GlobaleNBID.Unpack(stream)// p8
    }
    if (ngENBDUID_flag & _flags) == ngENBDUID_flag { //cond2
        self.NgENBDUID = &NGENBDUID{}//7{'type': 'NGENB-DU-ID', 'name': 'ngENB-DU-ID', 'optional': True}
        self.NgENBDUID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalE2nodengeNBID) Pack(stream *Stream) {
    const globaleNBID_flag uint = 0x00000002
    const ngENBDUID_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GlobalngeNBID.Pack(stream)
    if self.GlobaleNBID != nil { 
        _flags |= globaleNBID_flag
        self.GlobaleNBID.Pack(stream)
    }//end of optional
    if self.NgENBDUID != nil { 
        _flags |= ngENBDUID_flag
        self.NgENBDUID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type GlobalENBID struct { // [{'type': 'PLMN-Identity', 'name': 'pLMN-Identity'}, {'type': 'ENB-ID', 'name': 'eNB-ID'}, None]
    PLMNIdentity PLMNIdentity
    ENBID ENBID
}

func (self * GlobalENBID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PLMNIdentity.Unpack(stream)// p8
    self.ENBID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalENBID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.ENBID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GlobalenGNBID struct { // [{'type': 'PLMN-Identity', 'name': 'pLMN-Identity'}, {'type': 'ENGNB-ID', 'name': 'gNB-ID'}, None]
    PLMNIdentity PLMNIdentity
    GNBID ENGNBID
}

func (self * GlobalenGNBID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PLMNIdentity.Unpack(stream)// p8
    self.GNBID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalenGNBID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.GNBID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GlobalgNBID struct { // [{'type': 'PLMN-Identity', 'name': 'plmn-id'}, {'type': 'GNB-ID-Choice', 'name': 'gnb-id'}, None]
    Plmnid PLMNIdentity
    Gnbid GNBIDChoice
}

func (self * GlobalgNBID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.Plmnid.Unpack(stream)// p8
    self.Gnbid.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalgNBID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.Plmnid.Pack(stream)
    self.Gnbid.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GlobalngeNBID struct { // [{'type': 'PLMN-Identity', 'name': 'plmn-id'}, {'type': 'ENB-ID-Choice', 'name': 'enb-id'}, None]
    Plmnid PLMNIdentity
    Enbid ENBIDChoice
}

func (self * GlobalngeNBID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.Plmnid.Unpack(stream)// p8
    self.Enbid.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalngeNBID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.Plmnid.Pack(stream)
    self.Enbid.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *GlobalNGRANNodeID)Unpack(stream *Stream) {
    //coptions := []string{"gNB","ng-eNB"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in GlobalNGRANNodeID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.GNB = &GlobalgNBID{}//cho6
        self.GNB.Unpack(stream)
    } else if choice == 1 { //ch2
        self.NgeNB = &GlobalngeNBID{}//cho6
        self.NgeNB.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * GlobalNGRANNodeID) Pack(stream *Stream) {
    if self.GNB != nil {
        stream.set_choice(0, 1, 1, 2)
        self.GNB.Pack(stream)//2
    } else if self.NgeNB != nil {
        stream.set_choice(1, 1, 1, 2)
        self.NgeNB.Pack(stream)//2
    }

}
type GlobalNGRANNodeID struct { //[{'type': 'GlobalgNB-ID', 'name': 'gNB'}, {'type': 'GlobalngeNB-ID', 'name': 'ng-eNB'}, None]
    GNB *GlobalgNBID
    NgeNB *GlobalngeNBID
} // GlobalNGRANNodeID

type GlobalRICID struct { // [{'type': 'PLMN-Identity', 'name': 'pLMN-Identity'}, {'type': 'BIT STRING', 'size': [20], 'name': 'ric-ID'}, None]
    PLMNIdentity PLMNIdentity
    RicID BITSTRING
}

func (self * GlobalRICID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PLMNIdentity.Unpack(stream)// p8
    var Unpack_ricID = func(st *Stream, self *BITSTRING){
        self.Value = st.parsef_BitString(20, 20)
    }
    Unpack_ricID(stream, &self.RicID)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalRICID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    var Pack_ricID = func(st *Stream, self BITSTRING) {
        st.formatf_BitString(self.Value, 20)
    }
    Pack_ricID(stream, self.RicID) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
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
type GNBDUID struct {
  Value uint64
}
func (self *GNBDUID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(68719476736, 36, 0, 0)
}
func (self * GNBDUID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 68719476736, 36, 0, 0)
}
func (self *GNBIDChoice)Unpack(stream *Stream) {
    //coptions := []string{"gnb-ID"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in GNBIDChoice\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_gnbID = func(st *Stream, self *BITSTRING){
            self.Len = int(st.parse_blen(4, 0)+22)
            self.Value = st.parsef_BitString(32, int(self.Len))
        }
        self.GnbID = &BITSTRING{}//cho5
        Unpack_gnbID(stream, self.GnbID);
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * GNBIDChoice) Pack(stream *Stream) {
    if self.GnbID != nil {
        stream.set_choice(0, 0, 1, 1)
        var Pack_gnbID = func(st *Stream, self BITSTRING) {
            st.format_blen(int(self.Len-22), 4, 0)
            st.formatf_BitString(self.Value, int(self.Len))
        }
        Pack_gnbID(stream, *self.GnbID)//3
    }

}
type GNBIDChoice struct { //[{'type': 'BIT STRING', 'size': [(22, 32)], 'name': 'gnb-ID'}, None]
    GnbID *BITSTRING
} // GNBIDChoice

type MMEname struct {
  Value string
}
func (self *MMEname) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in MMEname")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *MMEname) Pack(st *Stream) {
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
type NGENBDUID struct {
  Value uint64
}
func (self *NGENBDUID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(68719476736, 36, 0, 0)
}
func (self * NGENBDUID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 68719476736, 36, 0, 0)
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
type RANfunctionDefinition struct {
  Value HexBytes
}
func (self *RANfunctionDefinition) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RANfunctionDefinition) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RANfunctionID struct {
  Value uint64
}
func (self *RANfunctionID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4096, 12, 0, 0)
}
func (self * RANfunctionID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4096, 12, 0, 0)
}
type RANfunctionOID struct {
  Value string
}
func (self *RANfunctionOID) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(10)+1
    if _len < 1 || _len > 1000 {
        print ("Invalid len in RANfunctionOID")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *RANfunctionOID) Pack(st *Stream) {
    _eflag := 0
    if len(self.Value) > 1000 {
       _eflag = 1
    }
    st.format_ext(_eflag)
    if len(self.Value) < 1 || len(self.Value) > 1000 {
        return;
    }
    st.format_olen((len(self.Value))-1, 10)
    st.formatf_PriString(self.Value, len(self.Value))
}
type RANfunctionRevision struct {
  Value uint64
}
func (self *RANfunctionRevision) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4096, 12, 0, 0)
}
func (self * RANfunctionRevision) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4096, 12, 0, 0)
}
type RICactionDefinition struct {
  Value HexBytes
}
func (self *RICactionDefinition) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RICactionDefinition) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RICactionID struct {
  Value uint64
}
func (self *RICactionID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * RICactionID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
type RICactionType struct {
  Value int
}
const (
    RICactionTypereport = 0
    RICactionTypeinsert = 1
    RICactionTypepolicy = 2

    /* Extensions */
)
func (self *RICactionType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *RICactionType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type RICcallProcessID struct {
  Value HexBytes
}
func (self *RICcallProcessID) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RICcallProcessID) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RICcontrolAckRequest struct {
  Value int
}
const (
    RICcontrolAckRequestnoAck = 0
    RICcontrolAckRequestack = 1

    /* Extensions */
)
func (self *RICcontrolAckRequest) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *RICcontrolAckRequest) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type RICcontrolHeader struct {
  Value HexBytes
}
func (self *RICcontrolHeader) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RICcontrolHeader) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RICcontrolMessage struct {
  Value HexBytes
}
func (self *RICcontrolMessage) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RICcontrolMessage) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RICcontrolOutcome struct {
  Value HexBytes
}
func (self *RICcontrolOutcome) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RICcontrolOutcome) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RICeventTriggerDefinition struct {
  Value HexBytes
}
func (self *RICeventTriggerDefinition) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RICeventTriggerDefinition) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RICindicationHeader struct {
  Value HexBytes
}
func (self *RICindicationHeader) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RICindicationHeader) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RICindicationMessage struct {
  Value HexBytes
}
func (self *RICindicationMessage) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RICindicationMessage) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RICindicationSN struct {
  Value uint64
}
func (self *RICindicationSN) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * RICindicationSN) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type RICindicationType struct {
  Value int
}
const (
    RICindicationTypereport = 0
    RICindicationTypeinsert = 1

    /* Extensions */
)
func (self *RICindicationType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *RICindicationType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type RICrequestID struct { // [{'type': 'INTEGER', 'restricted-to': [(0, 65535)], 'name': 'ricRequestorID'}, {'type': 'INTEGER', 'restricted-to': [(0, 65535)], 'name': 'ricInstanceID'}, None]
    RicRequestorID INTEGER
    RicInstanceID INTEGER
}

func (self * RICrequestID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricRequestorID = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(65536, 16, 0, 0)
    }
    Unpack_ricRequestorID(stream, &self.RicRequestorID)// p2
    var Unpack_ricInstanceID = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(65536, 16, 0, 0)
    }
    Unpack_ricInstanceID(stream, &self.RicInstanceID)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICrequestID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricRequestorID = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 65536, 16, 0, 0)
    }
    Pack_ricRequestorID(stream, self.RicRequestorID) //f2
    var Pack_ricInstanceID = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 65536, 16, 0, 0)
    }
    Pack_ricInstanceID(stream, self.RicInstanceID) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICsubsequentAction struct { // [{'type': 'RICsubsequentActionType', 'name': 'ricSubsequentActionType'}, {'type': 'RICtimeToWait', 'name': 'ricTimeToWait'}, None]
    RicSubsequentActionType RICsubsequentActionType
    RicTimeToWait RICtimeToWait
}

func (self * RICsubsequentAction) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicSubsequentActionType.Unpack(stream)// p8
    self.RicTimeToWait.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICsubsequentAction) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicSubsequentActionType.Pack(stream)
    self.RicTimeToWait.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RICsubsequentActionType struct {
  Value int
}
const (
    RICsubsequentActionTypecontinue = 0
    RICsubsequentActionTypewait = 1

    /* Extensions */
)
func (self *RICsubsequentActionType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *RICsubsequentActionType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type RICtimeToWait struct {
  Value int
}
const (
    RICtimeToWaitw1ms = 0
    RICtimeToWaitw2ms = 1
    RICtimeToWaitw5ms = 2
    RICtimeToWaitw10ms = 3
    RICtimeToWaitw20ms = 4
    RICtimeToWaitw30ms = 5
    RICtimeToWaitw40ms = 6
    RICtimeToWaitw50ms = 7
    RICtimeToWaitw100ms = 8
    RICtimeToWaitw200ms = 9
    RICtimeToWaitw500ms = 10
    RICtimeToWaitw1s = 11
    RICtimeToWaitw2s = 12
    RICtimeToWaitw5s = 13
    RICtimeToWaitw10s = 14
    RICtimeToWaitw20s = 15
    RICtimeToWaitw60s = 16

    /* Extensions */
)
func (self *RICtimeToWait) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(6, 17, 1)
}
func (self *RICtimeToWait) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 6, 17, 1)
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
type TNLinformation struct { // [{'type': 'BIT STRING', 'size': [(1, 160), None], 'name': 'tnlAddress'}, {'type': 'BIT STRING', 'size': [16], 'name': 'tnlPort', 'optional': True}, None]
    TnlAddress BITSTRING
    TnlPort *BITSTRING
}

func (self * TNLinformation) Unpack(stream *Stream) {
    tnlPort_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_tnlAddress = func(st *Stream, self *BITSTRING){
        self.Len = int(st.parse_blen(9, 1)+1)
        self.Value = st.parsef_BitString(160, int(self.Len))
    }
    Unpack_tnlAddress(stream, &self.TnlAddress)// p2
    if (tnlPort_flag & _flags) == tnlPort_flag { //cond1
        var Unpack_tnlPort = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(16, 16)
        }
        self.TnlPort = &BITSTRING{}//6{'type': 'BIT STRING', 'size': [16], 'name': 'tnlPort', 'optional': True}
        Unpack_tnlPort(stream, self.TnlPort)// p1 {'type': 'BIT STRING', 'size': [16], 'name': 'tnlPort', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TNLinformation) Pack(stream *Stream) {
    const tnlPort_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_tnlAddress = func(st *Stream, self BITSTRING) {
        st.format_blen(int(self.Len-1), 9, 1)
        st.formatf_BitString(self.Value, int(self.Len))
    }
    Pack_tnlAddress(stream, self.TnlAddress) //f2
    if self.TnlPort != nil { //YY
        _flags |= tnlPort_flag
        var Pack_tnlPort = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 16)
        }
        Pack_tnlPort(stream, *self.TnlPort) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type TNLusage struct {
  Value int
}
const (
    TNLusageric_service = 0
    TNLusagesupport_function = 1
    TNLusageboth = 2

    /* Extensions */
)
func (self *TNLusage) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *TNLusage) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
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
    TriggeringMessageunsuccessfull_outcome = 2
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
    _size := data.(E2APPROTOCOLIES_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['E2AP-PROTOCOL-IES']}
    Items map[int]*E2APPROTOCOLIES
    order []int
}

type ProtocolIESingleContainer struct{ //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['E2AP-PROTOCOL-IES']}
   Item map[int]*E2APPROTOCOLIES //UserType
   order []int
}
func (self *ProtocolIESingleContainer) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['E2AP-PROTOCOL-IES']}
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

type ProtocolIEField struct { // [{'type': 'E2AP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'E2AP-PROTOCOL-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'E2AP-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id E2APPROTOCOLIESid
    Criticality E2APPROTOCOLIEScriticality
    Value E2APPROTOCOLIESValue
}

func (self * ProtocolIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'E2AP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := E2APPROTOCOLIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(E2APPROTOCOLIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'E2AP-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * E2APPROTOCOLIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'E2AP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (E2APPROTOCOLIESid)(self.ID)
    if out.(E2APPROTOCOLIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(E2APPROTOCOLIES_IF).PackOT(stream, key)
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
    _size := data.(E2APPROTOCOLIESPAIR_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainerPair struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-FieldPair', 'actual-parameters': ['IEsSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['E2AP-PROTOCOL-IES-PAIR']}
    Items map[int]*E2APPROTOCOLIESPAIR
    order []int
}

type ProtocolIEFieldPair struct { // [{'type': 'E2AP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'E2AP-PROTOCOL-IES-PAIR.&firstCriticality', 'table': ['IEsSetParam', ['id']], 'name': 'firstCriticality'}, {'type': 'E2AP-PROTOCOL-IES-PAIR.&FirstValue', 'table': ['IEsSetParam', ['id']], 'name': 'firstValue'}, {'type': 'E2AP-PROTOCOL-IES-PAIR.&secondCriticality', 'table': ['IEsSetParam', ['id']], 'name': 'secondCriticality'}, {'type': 'E2AP-PROTOCOL-IES-PAIR.&SecondValue', 'table': ['IEsSetParam', ['id']], 'name': 'secondValue'}]
    Id E2APPROTOCOLIESPAIRid
    FirstCriticality E2APPROTOCOLIESPAIRfirstCriticality
    FirstValue E2APPROTOCOLIESPAIRFirstValue
    SecondCriticality E2APPROTOCOLIESPAIRsecondCriticality
    SecondValue E2APPROTOCOLIESPAIRSecondValue
}

func (self * ProtocolIEFieldPair) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'E2AP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := E2APPROTOCOLIESPAIR{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.FirstCriticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(E2APPROTOCOLIESPAIR_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'E2AP-PROTOCOL-IES-PAIR.&FirstValue', 'table': ['IEsSetParam', ['id']], 'name': 'firstValue'}
    self.SecondCriticality.Unpack(stream)//p9
    out.(E2APPROTOCOLIESPAIR_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'E2AP-PROTOCOL-IES-PAIR.&SecondValue', 'table': ['IEsSetParam', ['id']], 'name': 'secondValue'}
    stream.set_location(location, _len)
    return
}

func (self * E2APPROTOCOLIESPAIR) Pack(stream *Stream, out interface{}) {
    //table {'type': 'E2AP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (E2APPROTOCOLIESPAIRid)(self.ID)
    if out.(E2APPROTOCOLIESPAIR_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.FIRSTCRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(E2APPROTOCOLIESPAIR_IF).PackOT(stream, key)
    self.SECONDCRITICALITY.Pack(stream)
    out.(E2APPROTOCOLIESPAIR_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

func (self *ProtocolIEContainerList)Unpack(stream *Stream, arg1 uint64, arg2 uint64,  data interface{}) { //Seq3
    //cloc := stream.get_current_location();
    _size := stream.get_listsize(arg2-arg1+1) + int(arg1)
    val := ProtocolIESingleContainer{}
    for item := 0; item <_size; item +=1 {
        val.Unpack(stream, data);
        //values = append(values, val.(interface{}))
    }
    return
}


func (self *ProtocolIEContainerList) Pack (stream *Stream, arg1 int, arg2 int, data interface{}) { //seqof 4
    _size := len(self.Items)
    stream.set_listsize(_size-arg1, uint64(arg2-arg1+1))
    val := ProtocolIESingleContainer{}
    for _, item := range self.Items {
        val.Pack(stream, item);
    }
}


type ProtocolIEContainerList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-SingleContainer', 'actual-parameters': ['IEsSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IEsSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'E2AP-PROTOCOL-IES']}
    Items map[int]*E2APPROTOCOLIES
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


type ProtocolIEContainerPairList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-ContainerPair', 'actual-parameters': ['IEsSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IEsSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'E2AP-PROTOCOL-IES-PAIR']}
    Items map[int]*E2APPROTOCOLIESPAIR
    order []int
}

type E2APELEMENTARYPROCEDURE struct {//CLASS {'members': [{'type': 'OpenType', 'name': '&InitiatingMessage'}, {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, {'type': 'ProcedureCode', 'name': '&procedureCode'}, {'type': 'Criticality', 'name': '&criticality'}], 'with-members': [{'with-order': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME', 'PROCEDURE CODE', 'CRITICALITY'], 'with-type': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME']}], 'alt-type': {'INITIATING MESSAGE': {'type': 'OpenType', 'name': '&InitiatingMessage'}, 'SUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, 'UNSUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, 'PROCEDURE CODE': {'type': 'ProcedureCode', 'name': '&procedureCode'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}}, 'id-type': {'type': 'ProcedureCode', 'name': '&procedureCode'}}
    INITIATINGMESSAGE interface{}
    SUCCESSFULOUTCOME interface{}
    UNSUCCESSFULOUTCOME interface{}
    PROCEDURECODE ProcedureCode
    CRITICALITY Criticality
}
type E2APELEMENTARYPROCEDUREInitiatingMessage interface{}
type E2APELEMENTARYPROCEDURESuccessfulOutcome interface{}
type E2APELEMENTARYPROCEDUREUnsuccessfulOutcome interface{}
type E2APELEMENTARYPROCEDUREprocedureCode ProcedureCode
func (self *E2APELEMENTARYPROCEDUREprocedureCode) Pack(st *Stream) {
    ieID := ProcedureCode(*self)
    ieID.Pack(st)
}
func (self *E2APELEMENTARYPROCEDUREprocedureCode) Unpack(st *Stream) {
    val := ProcedureCode{}
    val.Unpack(st)
    *self = E2APELEMENTARYPROCEDUREprocedureCode(val)
}
type E2APELEMENTARYPROCEDUREcriticality Criticality
func (self *E2APELEMENTARYPROCEDUREcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *E2APELEMENTARYPROCEDUREcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = E2APELEMENTARYPROCEDUREcriticality(val)
}

type E2APELEMENTARYPROCEDURE_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type E2APPROTOCOLIES struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type E2APPROTOCOLIESid ProtocolIEID
func (self *E2APPROTOCOLIESid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *E2APPROTOCOLIESid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = E2APPROTOCOLIESid(val)
}
type E2APPROTOCOLIEScriticality Criticality
func (self *E2APPROTOCOLIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *E2APPROTOCOLIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = E2APPROTOCOLIEScriticality(val)
}
type E2APPROTOCOLIESValue interface{}
type E2APPROTOCOLIESpresence Presence
func (self *E2APPROTOCOLIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *E2APPROTOCOLIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = E2APPROTOCOLIESpresence(val)
}

type E2APPROTOCOLIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type E2APPROTOCOLIESPAIR struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&firstCriticality'}, {'type': 'OpenType', 'name': '&FirstValue'}, {'type': 'Criticality', 'name': '&secondCriticality'}, {'type': 'OpenType', 'name': '&SecondValue'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'FIRST CRITICALITY', 'FIRST TYPE', 'SECOND CRITICALITY', 'SECOND TYPE', 'PRESENCE'], 'with-type': ['FIRST TYPE', 'SECOND TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'FIRST CRITICALITY': {'type': 'Criticality', 'name': '&firstCriticality'}, 'FIRST TYPE': {'type': 'OpenType', 'name': '&FirstValue'}, 'SECOND CRITICALITY': {'type': 'Criticality', 'name': '&secondCriticality'}, 'SECOND TYPE': {'type': 'OpenType', 'name': '&SecondValue'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    FIRSTCRITICALITY Criticality
    FIRSTTYPE interface{}
    SECONDCRITICALITY Criticality
    SECONDTYPE interface{}
    PRESENCE Presence
}
type E2APPROTOCOLIESPAIRid ProtocolIEID
func (self *E2APPROTOCOLIESPAIRid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *E2APPROTOCOLIESPAIRid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = E2APPROTOCOLIESPAIRid(val)
}
type E2APPROTOCOLIESPAIRfirstCriticality Criticality
func (self *E2APPROTOCOLIESPAIRfirstCriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *E2APPROTOCOLIESPAIRfirstCriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = E2APPROTOCOLIESPAIRfirstCriticality(val)
}
type E2APPROTOCOLIESPAIRFirstValue interface{}
type E2APPROTOCOLIESPAIRsecondCriticality Criticality
func (self *E2APPROTOCOLIESPAIRsecondCriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *E2APPROTOCOLIESPAIRsecondCriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = E2APPROTOCOLIESPAIRsecondCriticality(val)
}
type E2APPROTOCOLIESPAIRSecondValue interface{}
type E2APPROTOCOLIESPAIRpresence Presence
func (self *E2APPROTOCOLIESPAIRpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *E2APPROTOCOLIESPAIRpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = E2APPROTOCOLIESPAIRpresence(val)
}

type E2APPROTOCOLIESPAIR_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
//class E2APELEMENTARYPROCEDURES: #OBJSET1 {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'members': [{}, {}, None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_E2APELEMENTARYPROCEDURES = make(map[E2APELEMENTARYPROCEDUREprocedureCode]*E2APELEMENTARYPROCEDURE)


//class E2APELEMENTARYPROCEDURESCLASS1: #OBJSET1 {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_E2APELEMENTARYPROCEDURESCLASS1 = make(map[E2APELEMENTARYPROCEDUREprocedureCode]*E2APELEMENTARYPROCEDURE)


//class E2APELEMENTARYPROCEDURESCLASS2: #OBJSET1 {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_E2APELEMENTARYPROCEDURESCLASS2 = make(map[E2APELEMENTARYPROCEDUREprocedureCode]*E2APELEMENTARYPROCEDURE)


type RICsubscriptionRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICsubscriptionDetails', 'CRITICALITY': 'reject', 'TYPE': 'RICsubscriptionDetails', 'PRESENCE': 'mandatory'}, None]}
   RICrequestID  RICrequestID
   RANfunctionID  RANfunctionID
   RICsubscriptionDetails  RICsubscriptionDetails
   list []interface{}
}
func (self *RICsubscriptionRequestIEs)createOT() interface{}{
    return nil
}
var table_RICsubscriptionRequestIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICsubscriptionRequestIEs = make([]int, 3)

func (self *RICsubscriptionRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICrequestID
   count +=1 //self.RANfunctionID
   count +=1 //self.RICsubscriptionDetails
   return count//ObjSet
}
func (self *RICsubscriptionRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        return true //self.RICrequestID
      case 5: //RANfunctionID
        return true //self.RANfunctionID
      case 30: //RICsubscriptionDetails
        return true //self.RICsubscriptionDetails
   }
   return false//ObjSet
}
func (self *RICsubscriptionRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, &self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, &self.RANfunctionID)
      case 30: //RICsubscriptionDetails
        self.RICsubscriptionDetails.Unpack(st)
        self.list = append(self.list, &self.RICsubscriptionDetails)
   }
}
func (self *RICsubscriptionRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Pack(st)
      case 5: //RANfunctionID
        self.RANfunctionID.Pack(st)
      case 30: //RICsubscriptionDetails
        self.RICsubscriptionDetails.Pack(st)
      default:
      break
   }
}
func init() {
table_RICsubscriptionRequestIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionRequestIEs[0] = 29
table_RICsubscriptionRequestIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionRequestIEs[1] = 5
table_RICsubscriptionRequestIEs[30] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICsubscriptionDetails}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICsubscriptionDetails{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionRequestIEs[2] = 30
   }

type RICactionToBeSetupItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICaction-ToBeSetup-Item', 'CRITICALITY': 'ignore', 'TYPE': 'RICaction-ToBeSetup-Item', 'PRESENCE': 'mandatory'}, None]}
   RICactionToBeSetupItem  RICactionToBeSetupItem
   list []interface{}
}
func (self *RICactionToBeSetupItemIEs)createOT() interface{}{
    return nil
}
var table_RICactionToBeSetupItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICactionToBeSetupItemIEs = make([]int, 1)

func (self *RICactionToBeSetupItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICactionToBeSetupItem
   return count//ObjSet
}
func (self *RICactionToBeSetupItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 19: //RICactionToBeSetupItem
        return true //self.RICactionToBeSetupItem
   }
   return false//ObjSet
}
func (self *RICactionToBeSetupItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 19: //RICactionToBeSetupItem
        self.RICactionToBeSetupItem.Unpack(st)
        self.list = append(self.list, &self.RICactionToBeSetupItem)
   }
}
func (self *RICactionToBeSetupItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 19: //RICactionToBeSetupItem
        self.RICactionToBeSetupItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RICactionToBeSetupItemIEs[19] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICactionToBeSetupItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RICactionToBeSetupItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RICactionToBeSetupItemIEs[0] = 19
   }

type RICsubscriptionResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICactions-Admitted', 'CRITICALITY': 'reject', 'TYPE': 'RICaction-Admitted-List', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICactions-NotAdmitted', 'CRITICALITY': 'reject', 'TYPE': 'RICaction-NotAdmitted-List', 'PRESENCE': 'optional'}, None]}
   RICrequestID  RICrequestID
   RANfunctionID  RANfunctionID
   RICactionsAdmitted  RICactionAdmittedList
   RICactionsNotAdmitted  *RICactionNotAdmittedList
   list []interface{}
}
func (self *RICsubscriptionResponseIEs)createOT() interface{}{
    return nil
}
var table_RICsubscriptionResponseIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICsubscriptionResponseIEs = make([]int, 4)

func (self *RICsubscriptionResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICrequestID
   count +=1 //self.RANfunctionID
   count +=1 //self.RICactionsAdmitted
   if self.RICactionsNotAdmitted != nil { count += 1 }
   return count//ObjSet
}
func (self *RICsubscriptionResponseIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        return true //self.RICrequestID
      case 5: //RANfunctionID
        return true //self.RANfunctionID
      case 17: //RICactionsAdmitted
        return true //self.RICactionsAdmitted
      case 18: //RICactionsNotAdmitted
        if self.RICactionsNotAdmitted != nil { return true }
   }
   return false//ObjSet
}
func (self *RICsubscriptionResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, &self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, &self.RANfunctionID)
      case 17: //RICactionsAdmitted
        self.RICactionsAdmitted.Unpack(st)
        self.list = append(self.list, &self.RICactionsAdmitted)
      case 18: //RICactionsNotAdmitted
        self.RICactionsNotAdmitted = &RICactionNotAdmittedList{}
        self.RICactionsNotAdmitted.Unpack(st)
        self.list = append(self.list, self.RICactionsNotAdmitted)
   }
}
func (self *RICsubscriptionResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Pack(st)
      case 5: //RANfunctionID
        self.RANfunctionID.Pack(st)
      case 17: //RICactionsAdmitted
        self.RICactionsAdmitted.Pack(st)
      case 18: //RICactionsNotAdmitted
        if self.RICactionsNotAdmitted != nil {self.RICactionsNotAdmitted.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICsubscriptionResponseIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionResponseIEs[0] = 29
table_RICsubscriptionResponseIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionResponseIEs[1] = 5
table_RICsubscriptionResponseIEs[17] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICactionsAdmitted}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICactionAdmittedList{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionResponseIEs[2] = 17
table_RICsubscriptionResponseIEs[18] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICactionsNotAdmitted}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICactionNotAdmittedList{}, PRESENCE:Presence{Presenceoptional}, }
order_RICsubscriptionResponseIEs[3] = 18
   }

type RICactionAdmittedItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICaction-Admitted-Item', 'CRITICALITY': 'ignore', 'TYPE': 'RICaction-Admitted-Item', 'PRESENCE': 'mandatory'}, None]}
   RICactionAdmittedItem  RICactionAdmittedItem
   list []interface{}
}
func (self *RICactionAdmittedItemIEs)createOT() interface{}{
    return nil
}
var table_RICactionAdmittedItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICactionAdmittedItemIEs = make([]int, 1)

func (self *RICactionAdmittedItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICactionAdmittedItem
   return count//ObjSet
}
func (self *RICactionAdmittedItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 14: //RICactionAdmittedItem
        return true //self.RICactionAdmittedItem
   }
   return false//ObjSet
}
func (self *RICactionAdmittedItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 14: //RICactionAdmittedItem
        self.RICactionAdmittedItem.Unpack(st)
        self.list = append(self.list, &self.RICactionAdmittedItem)
   }
}
func (self *RICactionAdmittedItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 14: //RICactionAdmittedItem
        self.RICactionAdmittedItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RICactionAdmittedItemIEs[14] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICactionAdmittedItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RICactionAdmittedItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RICactionAdmittedItemIEs[0] = 14
   }

type RICactionNotAdmittedItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICaction-NotAdmitted-Item', 'CRITICALITY': 'ignore', 'TYPE': 'RICaction-NotAdmitted-Item', 'PRESENCE': 'mandatory'}, None]}
   RICactionNotAdmittedItem  RICactionNotAdmittedItem
   list []interface{}
}
func (self *RICactionNotAdmittedItemIEs)createOT() interface{}{
    return nil
}
var table_RICactionNotAdmittedItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICactionNotAdmittedItemIEs = make([]int, 1)

func (self *RICactionNotAdmittedItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICactionNotAdmittedItem
   return count//ObjSet
}
func (self *RICactionNotAdmittedItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 16: //RICactionNotAdmittedItem
        return true //self.RICactionNotAdmittedItem
   }
   return false//ObjSet
}
func (self *RICactionNotAdmittedItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 16: //RICactionNotAdmittedItem
        self.RICactionNotAdmittedItem.Unpack(st)
        self.list = append(self.list, &self.RICactionNotAdmittedItem)
   }
}
func (self *RICactionNotAdmittedItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 16: //RICactionNotAdmittedItem
        self.RICactionNotAdmittedItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RICactionNotAdmittedItemIEs[16] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICactionNotAdmittedItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RICactionNotAdmittedItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RICactionNotAdmittedItemIEs[0] = 16
   }

type RICsubscriptionFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'reject', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   RICrequestID  RICrequestID
   RANfunctionID  RANfunctionID
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *RICsubscriptionFailureIEs)createOT() interface{}{
    return nil
}
var table_RICsubscriptionFailureIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICsubscriptionFailureIEs = make([]int, 4)

func (self *RICsubscriptionFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICrequestID
   count +=1 //self.RANfunctionID
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *RICsubscriptionFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        return true //self.RICrequestID
      case 5: //RANfunctionID
        return true //self.RANfunctionID
      case 1: //Cause
        return true //self.Cause
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *RICsubscriptionFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, &self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, &self.RANfunctionID)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *RICsubscriptionFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Pack(st)
      case 5: //RANfunctionID
        self.RANfunctionID.Pack(st)
      case 1: //Cause
        self.Cause.Pack(st)
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICsubscriptionFailureIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionFailureIEs[0] = 29
table_RICsubscriptionFailureIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionFailureIEs[1] = 5
table_RICsubscriptionFailureIEs[1] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionFailureIEs[2] = 1
table_RICsubscriptionFailureIEs[2] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_RICsubscriptionFailureIEs[3] = 2
   }

type RICsubscriptionDeleteRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'mandatory'}, None]}
   RICrequestID  RICrequestID
   RANfunctionID  RANfunctionID
   list []interface{}
}
func (self *RICsubscriptionDeleteRequestIEs)createOT() interface{}{
    return nil
}
var table_RICsubscriptionDeleteRequestIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICsubscriptionDeleteRequestIEs = make([]int, 2)

func (self *RICsubscriptionDeleteRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICrequestID
   count +=1 //self.RANfunctionID
   return count//ObjSet
}
func (self *RICsubscriptionDeleteRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        return true //self.RICrequestID
      case 5: //RANfunctionID
        return true //self.RANfunctionID
   }
   return false//ObjSet
}
func (self *RICsubscriptionDeleteRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, &self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, &self.RANfunctionID)
   }
}
func (self *RICsubscriptionDeleteRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Pack(st)
      case 5: //RANfunctionID
        self.RANfunctionID.Pack(st)
      default:
      break
   }
}
func init() {
table_RICsubscriptionDeleteRequestIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionDeleteRequestIEs[0] = 29
table_RICsubscriptionDeleteRequestIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionDeleteRequestIEs[1] = 5
   }

type RICsubscriptionDeleteResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'mandatory'}, None]}
   RICrequestID  RICrequestID
   RANfunctionID  RANfunctionID
   list []interface{}
}
func (self *RICsubscriptionDeleteResponseIEs)createOT() interface{}{
    return nil
}
var table_RICsubscriptionDeleteResponseIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICsubscriptionDeleteResponseIEs = make([]int, 2)

func (self *RICsubscriptionDeleteResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICrequestID
   count +=1 //self.RANfunctionID
   return count//ObjSet
}
func (self *RICsubscriptionDeleteResponseIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        return true //self.RICrequestID
      case 5: //RANfunctionID
        return true //self.RANfunctionID
   }
   return false//ObjSet
}
func (self *RICsubscriptionDeleteResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, &self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, &self.RANfunctionID)
   }
}
func (self *RICsubscriptionDeleteResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Pack(st)
      case 5: //RANfunctionID
        self.RANfunctionID.Pack(st)
      default:
      break
   }
}
func init() {
table_RICsubscriptionDeleteResponseIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionDeleteResponseIEs[0] = 29
table_RICsubscriptionDeleteResponseIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionDeleteResponseIEs[1] = 5
   }

type RICsubscriptionDeleteFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   RICrequestID  RICrequestID
   RANfunctionID  RANfunctionID
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *RICsubscriptionDeleteFailureIEs)createOT() interface{}{
    return nil
}
var table_RICsubscriptionDeleteFailureIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICsubscriptionDeleteFailureIEs = make([]int, 4)

func (self *RICsubscriptionDeleteFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICrequestID
   count +=1 //self.RANfunctionID
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *RICsubscriptionDeleteFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        return true //self.RICrequestID
      case 5: //RANfunctionID
        return true //self.RANfunctionID
      case 1: //Cause
        return true //self.Cause
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *RICsubscriptionDeleteFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, &self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, &self.RANfunctionID)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *RICsubscriptionDeleteFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Pack(st)
      case 5: //RANfunctionID
        self.RANfunctionID.Pack(st)
      case 1: //Cause
        self.Cause.Pack(st)
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICsubscriptionDeleteFailureIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionDeleteFailureIEs[0] = 29
table_RICsubscriptionDeleteFailureIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionDeleteFailureIEs[1] = 5
table_RICsubscriptionDeleteFailureIEs[1] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionDeleteFailureIEs[2] = 1
table_RICsubscriptionDeleteFailureIEs[2] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_RICsubscriptionDeleteFailureIEs[3] = 2
   }

type RICsubscriptionDeleteRequiredIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICsubscriptionToBeRemoved', 'CRITICALITY': 'ignore', 'TYPE': 'RICsubscription-List-withCause', 'PRESENCE': 'mandatory'}, None]}
   RICsubscriptionToBeRemoved  RICsubscriptionListwithCause
   list []interface{}
}
func (self *RICsubscriptionDeleteRequiredIEs)createOT() interface{}{
    return nil
}
var table_RICsubscriptionDeleteRequiredIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICsubscriptionDeleteRequiredIEs = make([]int, 1)

func (self *RICsubscriptionDeleteRequiredIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICsubscriptionToBeRemoved
   return count//ObjSet
}
func (self *RICsubscriptionDeleteRequiredIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 50: //RICsubscriptionToBeRemoved
        return true //self.RICsubscriptionToBeRemoved
   }
   return false//ObjSet
}
func (self *RICsubscriptionDeleteRequiredIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 50: //RICsubscriptionToBeRemoved
        self.RICsubscriptionToBeRemoved.Unpack(st)
        self.list = append(self.list, &self.RICsubscriptionToBeRemoved)
   }
}
func (self *RICsubscriptionDeleteRequiredIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 50: //RICsubscriptionToBeRemoved
        self.RICsubscriptionToBeRemoved.Pack(st)
      default:
      break
   }
}
func init() {
table_RICsubscriptionDeleteRequiredIEs[50] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICsubscriptionToBeRemoved}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RICsubscriptionListwithCause{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionDeleteRequiredIEs[0] = 50
   }

type RICsubscriptionwithCauseItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICsubscription-withCause-Item', 'CRITICALITY': 'ignore', 'TYPE': 'RICsubscription-withCause-Item', 'PRESENCE': 'mandatory'}, None]}
   RICsubscriptionwithCauseItem  RICsubscriptionwithCauseItem
   list []interface{}
}
func (self *RICsubscriptionwithCauseItemIEs)createOT() interface{}{
    return nil
}
var table_RICsubscriptionwithCauseItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICsubscriptionwithCauseItemIEs = make([]int, 1)

func (self *RICsubscriptionwithCauseItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICsubscriptionwithCauseItem
   return count//ObjSet
}
func (self *RICsubscriptionwithCauseItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 51: //RICsubscriptionwithCauseItem
        return true //self.RICsubscriptionwithCauseItem
   }
   return false//ObjSet
}
func (self *RICsubscriptionwithCauseItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 51: //RICsubscriptionwithCauseItem
        self.RICsubscriptionwithCauseItem.Unpack(st)
        self.list = append(self.list, &self.RICsubscriptionwithCauseItem)
   }
}
func (self *RICsubscriptionwithCauseItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 51: //RICsubscriptionwithCauseItem
        self.RICsubscriptionwithCauseItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RICsubscriptionwithCauseItemIEs[51] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICsubscriptionwithCauseItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RICsubscriptionwithCauseItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RICsubscriptionwithCauseItemIEs[0] = 51
   }

type RICindicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICactionID', 'CRITICALITY': 'reject', 'TYPE': 'RICactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICindicationSN', 'CRITICALITY': 'reject', 'TYPE': 'RICindicationSN', 'PRESENCE': 'optional'}, {'ID': 'id-RICindicationType', 'CRITICALITY': 'reject', 'TYPE': 'RICindicationType', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICindicationHeader', 'CRITICALITY': 'reject', 'TYPE': 'RICindicationHeader', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICindicationMessage', 'CRITICALITY': 'reject', 'TYPE': 'RICindicationMessage', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICcallProcessID', 'CRITICALITY': 'reject', 'TYPE': 'RICcallProcessID', 'PRESENCE': 'optional'}, None]}
   RICrequestID  RICrequestID
   RANfunctionID  RANfunctionID
   RICactionID  RICactionID
   RICindicationSN  *RICindicationSN
   RICindicationType  RICindicationType
   RICindicationHeader  RICindicationHeader
   RICindicationMessage  RICindicationMessage
   RICcallProcessID  *RICcallProcessID
   list []interface{}
}
func (self *RICindicationIEs)createOT() interface{}{
    return nil
}
var table_RICindicationIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICindicationIEs = make([]int, 8)

func (self *RICindicationIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICrequestID
   count +=1 //self.RANfunctionID
   count +=1 //self.RICactionID
   if self.RICindicationSN != nil { count += 1 }
   count +=1 //self.RICindicationType
   count +=1 //self.RICindicationHeader
   count +=1 //self.RICindicationMessage
   if self.RICcallProcessID != nil { count += 1 }
   return count//ObjSet
}
func (self *RICindicationIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        return true //self.RICrequestID
      case 5: //RANfunctionID
        return true //self.RANfunctionID
      case 15: //RICactionID
        return true //self.RICactionID
      case 27: //RICindicationSN
        if self.RICindicationSN != nil { return true }
      case 28: //RICindicationType
        return true //self.RICindicationType
      case 25: //RICindicationHeader
        return true //self.RICindicationHeader
      case 26: //RICindicationMessage
        return true //self.RICindicationMessage
      case 20: //RICcallProcessID
        if self.RICcallProcessID != nil { return true }
   }
   return false//ObjSet
}
func (self *RICindicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, &self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, &self.RANfunctionID)
      case 15: //RICactionID
        self.RICactionID.Unpack(st)
        self.list = append(self.list, &self.RICactionID)
      case 27: //RICindicationSN
        self.RICindicationSN = &RICindicationSN{}
        self.RICindicationSN.Unpack(st)
        self.list = append(self.list, self.RICindicationSN)
      case 28: //RICindicationType
        self.RICindicationType.Unpack(st)
        self.list = append(self.list, &self.RICindicationType)
      case 25: //RICindicationHeader
        self.RICindicationHeader.Unpack(st)
        self.list = append(self.list, &self.RICindicationHeader)
      case 26: //RICindicationMessage
        self.RICindicationMessage.Unpack(st)
        self.list = append(self.list, &self.RICindicationMessage)
      case 20: //RICcallProcessID
        self.RICcallProcessID = &RICcallProcessID{}
        self.RICcallProcessID.Unpack(st)
        self.list = append(self.list, self.RICcallProcessID)
   }
}
func (self *RICindicationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Pack(st)
      case 5: //RANfunctionID
        self.RANfunctionID.Pack(st)
      case 15: //RICactionID
        self.RICactionID.Pack(st)
      case 27: //RICindicationSN
        if self.RICindicationSN != nil {self.RICindicationSN.Pack(st)}
      case 28: //RICindicationType
        self.RICindicationType.Pack(st)
      case 25: //RICindicationHeader
        self.RICindicationHeader.Pack(st)
      case 26: //RICindicationMessage
        self.RICindicationMessage.Pack(st)
      case 20: //RICcallProcessID
        if self.RICcallProcessID != nil {self.RICcallProcessID.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICindicationIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICindicationIEs[0] = 29
table_RICindicationIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICindicationIEs[1] = 5
table_RICindicationIEs[15] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICindicationIEs[2] = 15
table_RICindicationIEs[27] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICindicationSN}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICindicationSN{}, PRESENCE:Presence{Presenceoptional}, }
order_RICindicationIEs[3] = 27
table_RICindicationIEs[28] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICindicationType}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICindicationType{}, PRESENCE:Presence{Presencemandatory}, }
order_RICindicationIEs[4] = 28
table_RICindicationIEs[25] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICindicationHeader}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICindicationHeader{}, PRESENCE:Presence{Presencemandatory}, }
order_RICindicationIEs[5] = 25
table_RICindicationIEs[26] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICindicationMessage}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICindicationMessage{}, PRESENCE:Presence{Presencemandatory}, }
order_RICindicationIEs[6] = 26
table_RICindicationIEs[20] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICcallProcessID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICcallProcessID{}, PRESENCE:Presence{Presenceoptional}, }
order_RICindicationIEs[7] = 20
   }

type RICcontrolRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICcallProcessID', 'CRITICALITY': 'reject', 'TYPE': 'RICcallProcessID', 'PRESENCE': 'optional'}, {'ID': 'id-RICcontrolHeader', 'CRITICALITY': 'reject', 'TYPE': 'RICcontrolHeader', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICcontrolMessage', 'CRITICALITY': 'reject', 'TYPE': 'RICcontrolMessage', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICcontrolAckRequest', 'CRITICALITY': 'reject', 'TYPE': 'RICcontrolAckRequest', 'PRESENCE': 'optional'}, None]}
   RICrequestID  RICrequestID
   RANfunctionID  RANfunctionID
   RICcallProcessID  *RICcallProcessID
   RICcontrolHeader  RICcontrolHeader
   RICcontrolMessage  RICcontrolMessage
   RICcontrolAckRequest  *RICcontrolAckRequest
   list []interface{}
}
func (self *RICcontrolRequestIEs)createOT() interface{}{
    return nil
}
var table_RICcontrolRequestIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICcontrolRequestIEs = make([]int, 6)

func (self *RICcontrolRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICrequestID
   count +=1 //self.RANfunctionID
   if self.RICcallProcessID != nil { count += 1 }
   count +=1 //self.RICcontrolHeader
   count +=1 //self.RICcontrolMessage
   if self.RICcontrolAckRequest != nil { count += 1 }
   return count//ObjSet
}
func (self *RICcontrolRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        return true //self.RICrequestID
      case 5: //RANfunctionID
        return true //self.RANfunctionID
      case 20: //RICcallProcessID
        if self.RICcallProcessID != nil { return true }
      case 22: //RICcontrolHeader
        return true //self.RICcontrolHeader
      case 23: //RICcontrolMessage
        return true //self.RICcontrolMessage
      case 21: //RICcontrolAckRequest
        if self.RICcontrolAckRequest != nil { return true }
   }
   return false//ObjSet
}
func (self *RICcontrolRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, &self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, &self.RANfunctionID)
      case 20: //RICcallProcessID
        self.RICcallProcessID = &RICcallProcessID{}
        self.RICcallProcessID.Unpack(st)
        self.list = append(self.list, self.RICcallProcessID)
      case 22: //RICcontrolHeader
        self.RICcontrolHeader.Unpack(st)
        self.list = append(self.list, &self.RICcontrolHeader)
      case 23: //RICcontrolMessage
        self.RICcontrolMessage.Unpack(st)
        self.list = append(self.list, &self.RICcontrolMessage)
      case 21: //RICcontrolAckRequest
        self.RICcontrolAckRequest = &RICcontrolAckRequest{}
        self.RICcontrolAckRequest.Unpack(st)
        self.list = append(self.list, self.RICcontrolAckRequest)
   }
}
func (self *RICcontrolRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Pack(st)
      case 5: //RANfunctionID
        self.RANfunctionID.Pack(st)
      case 20: //RICcallProcessID
        if self.RICcallProcessID != nil {self.RICcallProcessID.Pack(st)}
      case 22: //RICcontrolHeader
        self.RICcontrolHeader.Pack(st)
      case 23: //RICcontrolMessage
        self.RICcontrolMessage.Pack(st)
      case 21: //RICcontrolAckRequest
        if self.RICcontrolAckRequest != nil {self.RICcontrolAckRequest.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICcontrolRequestIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICcontrolRequestIEs[0] = 29
table_RICcontrolRequestIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICcontrolRequestIEs[1] = 5
table_RICcontrolRequestIEs[20] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICcallProcessID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICcallProcessID{}, PRESENCE:Presence{Presenceoptional}, }
order_RICcontrolRequestIEs[2] = 20
table_RICcontrolRequestIEs[22] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICcontrolHeader}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICcontrolHeader{}, PRESENCE:Presence{Presencemandatory}, }
order_RICcontrolRequestIEs[3] = 22
table_RICcontrolRequestIEs[23] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICcontrolMessage}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICcontrolMessage{}, PRESENCE:Presence{Presencemandatory}, }
order_RICcontrolRequestIEs[4] = 23
table_RICcontrolRequestIEs[21] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICcontrolAckRequest}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICcontrolAckRequest{}, PRESENCE:Presence{Presenceoptional}, }
order_RICcontrolRequestIEs[5] = 21
   }

type RICcontrolAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICcallProcessID', 'CRITICALITY': 'reject', 'TYPE': 'RICcallProcessID', 'PRESENCE': 'optional'}, {'ID': 'id-RICcontrolOutcome', 'CRITICALITY': 'reject', 'TYPE': 'RICcontrolOutcome', 'PRESENCE': 'optional'}, None]}
   RICrequestID  RICrequestID
   RANfunctionID  RANfunctionID
   RICcallProcessID  *RICcallProcessID
   RICcontrolOutcome  *RICcontrolOutcome
   list []interface{}
}
func (self *RICcontrolAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_RICcontrolAcknowledgeIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICcontrolAcknowledgeIEs = make([]int, 4)

func (self *RICcontrolAcknowledgeIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICrequestID
   count +=1 //self.RANfunctionID
   if self.RICcallProcessID != nil { count += 1 }
   if self.RICcontrolOutcome != nil { count += 1 }
   return count//ObjSet
}
func (self *RICcontrolAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        return true //self.RICrequestID
      case 5: //RANfunctionID
        return true //self.RANfunctionID
      case 20: //RICcallProcessID
        if self.RICcallProcessID != nil { return true }
      case 32: //RICcontrolOutcome
        if self.RICcontrolOutcome != nil { return true }
   }
   return false//ObjSet
}
func (self *RICcontrolAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, &self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, &self.RANfunctionID)
      case 20: //RICcallProcessID
        self.RICcallProcessID = &RICcallProcessID{}
        self.RICcallProcessID.Unpack(st)
        self.list = append(self.list, self.RICcallProcessID)
      case 32: //RICcontrolOutcome
        self.RICcontrolOutcome = &RICcontrolOutcome{}
        self.RICcontrolOutcome.Unpack(st)
        self.list = append(self.list, self.RICcontrolOutcome)
   }
}
func (self *RICcontrolAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Pack(st)
      case 5: //RANfunctionID
        self.RANfunctionID.Pack(st)
      case 20: //RICcallProcessID
        if self.RICcallProcessID != nil {self.RICcallProcessID.Pack(st)}
      case 32: //RICcontrolOutcome
        if self.RICcontrolOutcome != nil {self.RICcontrolOutcome.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICcontrolAcknowledgeIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICcontrolAcknowledgeIEs[0] = 29
table_RICcontrolAcknowledgeIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICcontrolAcknowledgeIEs[1] = 5
table_RICcontrolAcknowledgeIEs[20] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICcallProcessID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICcallProcessID{}, PRESENCE:Presence{Presenceoptional}, }
order_RICcontrolAcknowledgeIEs[2] = 20
table_RICcontrolAcknowledgeIEs[32] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICcontrolOutcome}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICcontrolOutcome{}, PRESENCE:Presence{Presenceoptional}, }
order_RICcontrolAcknowledgeIEs[3] = 32
   }

type RICcontrolFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICcallProcessID', 'CRITICALITY': 'reject', 'TYPE': 'RICcallProcessID', 'PRESENCE': 'optional'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-RICcontrolOutcome', 'CRITICALITY': 'reject', 'TYPE': 'RICcontrolOutcome', 'PRESENCE': 'optional'}, None]}
   RICrequestID  RICrequestID
   RANfunctionID  RANfunctionID
   RICcallProcessID  *RICcallProcessID
   Cause  Cause
   RICcontrolOutcome  *RICcontrolOutcome
   list []interface{}
}
func (self *RICcontrolFailureIEs)createOT() interface{}{
    return nil
}
var table_RICcontrolFailureIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICcontrolFailureIEs = make([]int, 5)

func (self *RICcontrolFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.RICrequestID
   count +=1 //self.RANfunctionID
   if self.RICcallProcessID != nil { count += 1 }
   count +=1 //self.Cause
   if self.RICcontrolOutcome != nil { count += 1 }
   return count//ObjSet
}
func (self *RICcontrolFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        return true //self.RICrequestID
      case 5: //RANfunctionID
        return true //self.RANfunctionID
      case 20: //RICcallProcessID
        if self.RICcallProcessID != nil { return true }
      case 1: //Cause
        return true //self.Cause
      case 32: //RICcontrolOutcome
        if self.RICcontrolOutcome != nil { return true }
   }
   return false//ObjSet
}
func (self *RICcontrolFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, &self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, &self.RANfunctionID)
      case 20: //RICcallProcessID
        self.RICcallProcessID = &RICcallProcessID{}
        self.RICcallProcessID.Unpack(st)
        self.list = append(self.list, self.RICcallProcessID)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 32: //RICcontrolOutcome
        self.RICcontrolOutcome = &RICcontrolOutcome{}
        self.RICcontrolOutcome.Unpack(st)
        self.list = append(self.list, self.RICcontrolOutcome)
   }
}
func (self *RICcontrolFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 29: //RICrequestID
        self.RICrequestID.Pack(st)
      case 5: //RANfunctionID
        self.RANfunctionID.Pack(st)
      case 20: //RICcallProcessID
        if self.RICcallProcessID != nil {self.RICcallProcessID.Pack(st)}
      case 1: //Cause
        self.Cause.Pack(st)
      case 32: //RICcontrolOutcome
        if self.RICcontrolOutcome != nil {self.RICcontrolOutcome.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICcontrolFailureIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICcontrolFailureIEs[0] = 29
table_RICcontrolFailureIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICcontrolFailureIEs[1] = 5
table_RICcontrolFailureIEs[20] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICcallProcessID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICcallProcessID{}, PRESENCE:Presence{Presenceoptional}, }
order_RICcontrolFailureIEs[2] = 20
table_RICcontrolFailureIEs[1] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_RICcontrolFailureIEs[3] = 1
table_RICcontrolFailureIEs[32] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICcontrolOutcome}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICcontrolOutcome{}, PRESENCE:Presence{Presenceoptional}, }
order_RICcontrolFailureIEs[4] = 32
   }

type ErrorIndicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'optional'}, {'ID': 'id-RICrequestID', 'CRITICALITY': 'reject', 'TYPE': 'RICrequestID', 'PRESENCE': 'optional'}, {'ID': 'id-RANfunctionID', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionID', 'PRESENCE': 'optional'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  *TransactionID
   RICrequestID  *RICrequestID
   RANfunctionID  *RANfunctionID
   Cause  *Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ErrorIndicationIEs)createOT() interface{}{
    return nil
}
var table_ErrorIndicationIEs = make(map[int]*E2APPROTOCOLIES)

var order_ErrorIndicationIEs = make([]int, 5)

func (self *ErrorIndicationIEs) GetIECount() int{
   count := 0
   if self.TransactionID != nil { count += 1 }
   if self.RICrequestID != nil { count += 1 }
   if self.RANfunctionID != nil { count += 1 }
   if self.Cause != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ErrorIndicationIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        if self.TransactionID != nil { return true }
      case 29: //RICrequestID
        if self.RICrequestID != nil { return true }
      case 5: //RANfunctionID
        if self.RANfunctionID != nil { return true }
      case 1: //Cause
        if self.Cause != nil { return true }
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ErrorIndicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID = &TransactionID{}
        self.TransactionID.Unpack(st)
        self.list = append(self.list, self.TransactionID)
      case 29: //RICrequestID
        self.RICrequestID = &RICrequestID{}
        self.RICrequestID.Unpack(st)
        self.list = append(self.list, self.RICrequestID)
      case 5: //RANfunctionID
        self.RANfunctionID = &RANfunctionID{}
        self.RANfunctionID.Unpack(st)
        self.list = append(self.list, self.RANfunctionID)
      case 1: //Cause
        self.Cause = &Cause{}
        self.Cause.Unpack(st)
        self.list = append(self.list, self.Cause)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ErrorIndicationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        if self.TransactionID != nil {self.TransactionID.Pack(st)}
      case 29: //RICrequestID
        if self.RICrequestID != nil {self.RICrequestID.Pack(st)}
      case 5: //RANfunctionID
        if self.RANfunctionID != nil {self.RANfunctionID.Pack(st)}
      case 1: //Cause
        if self.Cause != nil {self.Cause.Pack(st)}
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ErrorIndicationIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[0] = 49
table_ErrorIndicationIEs[29] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRICrequestID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RICrequestID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[1] = 29
table_ErrorIndicationIEs[5] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[2] = 5
table_ErrorIndicationIEs[1] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[3] = 1
table_ErrorIndicationIEs[2] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[4] = 2
   }

type E2setupRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalE2node-ID', 'CRITICALITY': 'reject', 'TYPE': 'GlobalE2node-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionsAdded', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctions-List', 'PRESENCE': 'mandatory'}, {'ID': 'id-E2nodeComponentConfigAddition', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigAddition-List', 'PRESENCE': 'mandatory'}, None]}
   TransactionID  TransactionID
   GlobalE2nodeID  GlobalE2nodeID
   RANfunctionsAdded  RANfunctionsList
   E2nodeComponentConfigAddition  E2nodeComponentConfigAdditionList
   list []interface{}
}
func (self *E2setupRequestIEs)createOT() interface{}{
    return nil
}
var table_E2setupRequestIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2setupRequestIEs = make([]int, 4)

func (self *E2setupRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.GlobalE2nodeID
   count +=1 //self.RANfunctionsAdded
   count +=1 //self.E2nodeComponentConfigAddition
   return count//ObjSet
}
func (self *E2setupRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 3: //GlobalE2nodeID
        return true //self.GlobalE2nodeID
      case 10: //RANfunctionsAdded
        return true //self.RANfunctionsAdded
      case 50: //E2nodeComponentConfigAddition
        return true //self.E2nodeComponentConfigAddition
   }
   return false//ObjSet
}
func (self *E2setupRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 3: //GlobalE2nodeID
        self.GlobalE2nodeID.Unpack(st)
        self.list = append(self.list, &self.GlobalE2nodeID)
      case 10: //RANfunctionsAdded
        self.RANfunctionsAdded.Unpack(st)
        self.list = append(self.list, &self.RANfunctionsAdded)
      case 50: //E2nodeComponentConfigAddition
        self.E2nodeComponentConfigAddition.Unpack(st)
        self.list = append(self.list, &self.E2nodeComponentConfigAddition)
   }
}
func (self *E2setupRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 3: //GlobalE2nodeID
        self.GlobalE2nodeID.Pack(st)
      case 10: //RANfunctionsAdded
        self.RANfunctionsAdded.Pack(st)
      case 50: //E2nodeComponentConfigAddition
        self.E2nodeComponentConfigAddition.Pack(st)
      default:
      break
   }
}
func init() {
table_E2setupRequestIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2setupRequestIEs[0] = 49
table_E2setupRequestIEs[3] = &E2APPROTOCOLIES{ID:ProtocolIEID{idGlobalE2nodeID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GlobalE2nodeID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2setupRequestIEs[1] = 3
table_E2setupRequestIEs[10] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionsAdded}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionsList{}, PRESENCE:Presence{Presencemandatory}, }
order_E2setupRequestIEs[2] = 10
table_E2setupRequestIEs[50] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigAddition}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigAdditionList{}, PRESENCE:Presence{Presencemandatory}, }
order_E2setupRequestIEs[3] = 50
   }

type E2setupResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalRIC-ID', 'CRITICALITY': 'reject', 'TYPE': 'GlobalRIC-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionsAccepted', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionsID-List', 'PRESENCE': 'optional'}, {'ID': 'id-RANfunctionsRejected', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionsIDcause-List', 'PRESENCE': 'optional'}, {'ID': 'id-E2nodeComponentConfigAdditionAck', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigAdditionAck-List', 'PRESENCE': 'mandatory'}, None]}
   TransactionID  TransactionID
   GlobalRICID  GlobalRICID
   RANfunctionsAccepted  *RANfunctionsIDList
   RANfunctionsRejected  *RANfunctionsIDcauseList
   E2nodeComponentConfigAdditionAck  E2nodeComponentConfigAdditionAckList
   list []interface{}
}
func (self *E2setupResponseIEs)createOT() interface{}{
    return nil
}
var table_E2setupResponseIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2setupResponseIEs = make([]int, 5)

func (self *E2setupResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.GlobalRICID
   if self.RANfunctionsAccepted != nil { count += 1 }
   if self.RANfunctionsRejected != nil { count += 1 }
   count +=1 //self.E2nodeComponentConfigAdditionAck
   return count//ObjSet
}
func (self *E2setupResponseIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 4: //GlobalRICID
        return true //self.GlobalRICID
      case 9: //RANfunctionsAccepted
        if self.RANfunctionsAccepted != nil { return true }
      case 13: //RANfunctionsRejected
        if self.RANfunctionsRejected != nil { return true }
      case 52: //E2nodeComponentConfigAdditionAck
        return true //self.E2nodeComponentConfigAdditionAck
   }
   return false//ObjSet
}
func (self *E2setupResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 4: //GlobalRICID
        self.GlobalRICID.Unpack(st)
        self.list = append(self.list, &self.GlobalRICID)
      case 9: //RANfunctionsAccepted
        self.RANfunctionsAccepted = &RANfunctionsIDList{}
        self.RANfunctionsAccepted.Unpack(st)
        self.list = append(self.list, self.RANfunctionsAccepted)
      case 13: //RANfunctionsRejected
        self.RANfunctionsRejected = &RANfunctionsIDcauseList{}
        self.RANfunctionsRejected.Unpack(st)
        self.list = append(self.list, self.RANfunctionsRejected)
      case 52: //E2nodeComponentConfigAdditionAck
        self.E2nodeComponentConfigAdditionAck.Unpack(st)
        self.list = append(self.list, &self.E2nodeComponentConfigAdditionAck)
   }
}
func (self *E2setupResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 4: //GlobalRICID
        self.GlobalRICID.Pack(st)
      case 9: //RANfunctionsAccepted
        if self.RANfunctionsAccepted != nil {self.RANfunctionsAccepted.Pack(st)}
      case 13: //RANfunctionsRejected
        if self.RANfunctionsRejected != nil {self.RANfunctionsRejected.Pack(st)}
      case 52: //E2nodeComponentConfigAdditionAck
        self.E2nodeComponentConfigAdditionAck.Pack(st)
      default:
      break
   }
}
func init() {
table_E2setupResponseIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2setupResponseIEs[0] = 49
table_E2setupResponseIEs[4] = &E2APPROTOCOLIES{ID:ProtocolIEID{idGlobalRICID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GlobalRICID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2setupResponseIEs[1] = 4
table_E2setupResponseIEs[9] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionsAccepted}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionsIDList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2setupResponseIEs[2] = 9
table_E2setupResponseIEs[13] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionsRejected}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionsIDcauseList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2setupResponseIEs[3] = 13
table_E2setupResponseIEs[52] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigAdditionAck}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigAdditionAckList{}, PRESENCE:Presence{Presencemandatory}, }
order_E2setupResponseIEs[4] = 52
   }

type E2setupFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, {'ID': 'id-TNLinformation', 'CRITICALITY': 'ignore', 'TYPE': 'TNLinformation', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   TNLinformation  *TNLinformation
   list []interface{}
}
func (self *E2setupFailureIEs)createOT() interface{}{
    return nil
}
var table_E2setupFailureIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2setupFailureIEs = make([]int, 5)

func (self *E2setupFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   if self.TNLinformation != nil { count += 1 }
   return count//ObjSet
}
func (self *E2setupFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 1: //Cause
        return true //self.Cause
      case 31: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
      case 48: //TNLinformation
        if self.TNLinformation != nil { return true }
   }
   return false//ObjSet
}
func (self *E2setupFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 31: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
      case 48: //TNLinformation
        self.TNLinformation = &TNLinformation{}
        self.TNLinformation.Unpack(st)
        self.list = append(self.list, self.TNLinformation)
   }
}
func (self *E2setupFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 1: //Cause
        self.Cause.Pack(st)
      case 31: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      case 48: //TNLinformation
        if self.TNLinformation != nil {self.TNLinformation.Pack(st)}
      default:
      break
   }
}
func init() {
table_E2setupFailureIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2setupFailureIEs[0] = 49
table_E2setupFailureIEs[1] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_E2setupFailureIEs[1] = 1
table_E2setupFailureIEs[31] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_E2setupFailureIEs[2] = 31
table_E2setupFailureIEs[2] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_E2setupFailureIEs[3] = 2
table_E2setupFailureIEs[48] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTNLinformation}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TNLinformation{}, PRESENCE:Presence{Presenceoptional}, }
order_E2setupFailureIEs[4] = 48
   }

type E2connectionUpdateIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-E2connectionUpdateAdd', 'CRITICALITY': 'reject', 'TYPE': 'E2connectionUpdate-List', 'PRESENCE': 'optional'}, {'ID': 'id-E2connectionUpdateRemove', 'CRITICALITY': 'reject', 'TYPE': 'E2connectionUpdateRemove-List', 'PRESENCE': 'optional'}, {'ID': 'id-E2connectionUpdateModify', 'CRITICALITY': 'reject', 'TYPE': 'E2connectionUpdate-List', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   E2connectionUpdateAdd  *E2connectionUpdateList
   E2connectionUpdateRemove  *E2connectionUpdateRemoveList
   E2connectionUpdateModify  *E2connectionUpdateList
   list []interface{}
}
func (self *E2connectionUpdateIEs)createOT() interface{}{
    return nil
}
var table_E2connectionUpdateIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2connectionUpdateIEs = make([]int, 4)

func (self *E2connectionUpdateIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.E2connectionUpdateAdd != nil { count += 1 }
   if self.E2connectionUpdateRemove != nil { count += 1 }
   if self.E2connectionUpdateModify != nil { count += 1 }
   return count//ObjSet
}
func (self *E2connectionUpdateIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 44: //E2connectionUpdateAdd
        if self.E2connectionUpdateAdd != nil { return true }
      case 46: //E2connectionUpdateRemove
        if self.E2connectionUpdateRemove != nil { return true }
      case 45: //E2connectionUpdateModify
        if self.E2connectionUpdateModify != nil { return true }
   }
   return false//ObjSet
}
func (self *E2connectionUpdateIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 44: //E2connectionUpdateAdd
        self.E2connectionUpdateAdd = &E2connectionUpdateList{}
        self.E2connectionUpdateAdd.Unpack(st)
        self.list = append(self.list, self.E2connectionUpdateAdd)
      case 46: //E2connectionUpdateRemove
        self.E2connectionUpdateRemove = &E2connectionUpdateRemoveList{}
        self.E2connectionUpdateRemove.Unpack(st)
        self.list = append(self.list, self.E2connectionUpdateRemove)
      case 45: //E2connectionUpdateModify
        self.E2connectionUpdateModify = &E2connectionUpdateList{}
        self.E2connectionUpdateModify.Unpack(st)
        self.list = append(self.list, self.E2connectionUpdateModify)
   }
}
func (self *E2connectionUpdateIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 44: //E2connectionUpdateAdd
        if self.E2connectionUpdateAdd != nil {self.E2connectionUpdateAdd.Pack(st)}
      case 46: //E2connectionUpdateRemove
        if self.E2connectionUpdateRemove != nil {self.E2connectionUpdateRemove.Pack(st)}
      case 45: //E2connectionUpdateModify
        if self.E2connectionUpdateModify != nil {self.E2connectionUpdateModify.Pack(st)}
      default:
      break
   }
}
func init() {
table_E2connectionUpdateIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2connectionUpdateIEs[0] = 49
table_E2connectionUpdateIEs[44] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2connectionUpdateAdd}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2connectionUpdateList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2connectionUpdateIEs[1] = 44
table_E2connectionUpdateIEs[46] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2connectionUpdateRemove}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2connectionUpdateRemoveList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2connectionUpdateIEs[2] = 46
table_E2connectionUpdateIEs[45] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2connectionUpdateModify}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2connectionUpdateList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2connectionUpdateIEs[3] = 45
   }

type E2connectionUpdateItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-E2connectionUpdate-Item', 'CRITICALITY': 'ignore', 'TYPE': 'E2connectionUpdate-Item', 'PRESENCE': 'mandatory'}, None]}
   E2connectionUpdateItem  E2connectionUpdateItem
   list []interface{}
}
func (self *E2connectionUpdateItemIEs)createOT() interface{}{
    return nil
}
var table_E2connectionUpdateItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2connectionUpdateItemIEs = make([]int, 1)

func (self *E2connectionUpdateItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.E2connectionUpdateItem
   return count//ObjSet
}
func (self *E2connectionUpdateItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 43: //E2connectionUpdateItem
        return true //self.E2connectionUpdateItem
   }
   return false//ObjSet
}
func (self *E2connectionUpdateItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 43: //E2connectionUpdateItem
        self.E2connectionUpdateItem.Unpack(st)
        self.list = append(self.list, &self.E2connectionUpdateItem)
   }
}
func (self *E2connectionUpdateItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 43: //E2connectionUpdateItem
        self.E2connectionUpdateItem.Pack(st)
      default:
      break
   }
}
func init() {
table_E2connectionUpdateItemIEs[43] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2connectionUpdateItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&E2connectionUpdateItem{}, PRESENCE:Presence{Presencemandatory}, }
order_E2connectionUpdateItemIEs[0] = 43
   }

type E2connectionUpdateRemoveItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-E2connectionUpdateRemove-Item', 'CRITICALITY': 'ignore', 'TYPE': 'E2connectionUpdateRemove-Item', 'PRESENCE': 'mandatory'}, None]}
   E2connectionUpdateRemoveItem  E2connectionUpdateRemoveItem
   list []interface{}
}
func (self *E2connectionUpdateRemoveItemIEs)createOT() interface{}{
    return nil
}
var table_E2connectionUpdateRemoveItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2connectionUpdateRemoveItemIEs = make([]int, 1)

func (self *E2connectionUpdateRemoveItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.E2connectionUpdateRemoveItem
   return count//ObjSet
}
func (self *E2connectionUpdateRemoveItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 47: //E2connectionUpdateRemoveItem
        return true //self.E2connectionUpdateRemoveItem
   }
   return false//ObjSet
}
func (self *E2connectionUpdateRemoveItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 47: //E2connectionUpdateRemoveItem
        self.E2connectionUpdateRemoveItem.Unpack(st)
        self.list = append(self.list, &self.E2connectionUpdateRemoveItem)
   }
}
func (self *E2connectionUpdateRemoveItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 47: //E2connectionUpdateRemoveItem
        self.E2connectionUpdateRemoveItem.Pack(st)
      default:
      break
   }
}
func init() {
table_E2connectionUpdateRemoveItemIEs[47] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2connectionUpdateRemoveItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&E2connectionUpdateRemoveItem{}, PRESENCE:Presence{Presencemandatory}, }
order_E2connectionUpdateRemoveItemIEs[0] = 47
   }

type E2connectionUpdateAckIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-E2connectionSetup', 'CRITICALITY': 'reject', 'TYPE': 'E2connectionUpdate-List', 'PRESENCE': 'optional'}, {'ID': 'id-E2connectionSetupFailed', 'CRITICALITY': 'reject', 'TYPE': 'E2connectionSetupFailed-List', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   E2connectionSetup  *E2connectionUpdateList
   E2connectionSetupFailed  *E2connectionSetupFailedList
   list []interface{}
}
func (self *E2connectionUpdateAckIEs)createOT() interface{}{
    return nil
}
var table_E2connectionUpdateAckIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2connectionUpdateAckIEs = make([]int, 3)

func (self *E2connectionUpdateAckIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.E2connectionSetup != nil { count += 1 }
   if self.E2connectionSetupFailed != nil { count += 1 }
   return count//ObjSet
}
func (self *E2connectionUpdateAckIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 39: //E2connectionSetup
        if self.E2connectionSetup != nil { return true }
      case 40: //E2connectionSetupFailed
        if self.E2connectionSetupFailed != nil { return true }
   }
   return false//ObjSet
}
func (self *E2connectionUpdateAckIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 39: //E2connectionSetup
        self.E2connectionSetup = &E2connectionUpdateList{}
        self.E2connectionSetup.Unpack(st)
        self.list = append(self.list, self.E2connectionSetup)
      case 40: //E2connectionSetupFailed
        self.E2connectionSetupFailed = &E2connectionSetupFailedList{}
        self.E2connectionSetupFailed.Unpack(st)
        self.list = append(self.list, self.E2connectionSetupFailed)
   }
}
func (self *E2connectionUpdateAckIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 39: //E2connectionSetup
        if self.E2connectionSetup != nil {self.E2connectionSetup.Pack(st)}
      case 40: //E2connectionSetupFailed
        if self.E2connectionSetupFailed != nil {self.E2connectionSetupFailed.Pack(st)}
      default:
      break
   }
}
func init() {
table_E2connectionUpdateAckIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2connectionUpdateAckIEs[0] = 49
table_E2connectionUpdateAckIEs[39] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2connectionSetup}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2connectionUpdateList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2connectionUpdateAckIEs[1] = 39
table_E2connectionUpdateAckIEs[40] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2connectionSetupFailed}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2connectionSetupFailedList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2connectionUpdateAckIEs[2] = 40
   }

type E2connectionSetupFailedItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-E2connectionSetupFailed-Item', 'CRITICALITY': 'ignore', 'TYPE': 'E2connectionSetupFailed-Item', 'PRESENCE': 'mandatory'}, None]}
   E2connectionSetupFailedItem  E2connectionSetupFailedItem
   list []interface{}
}
func (self *E2connectionSetupFailedItemIEs)createOT() interface{}{
    return nil
}
var table_E2connectionSetupFailedItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2connectionSetupFailedItemIEs = make([]int, 1)

func (self *E2connectionSetupFailedItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.E2connectionSetupFailedItem
   return count//ObjSet
}
func (self *E2connectionSetupFailedItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 41: //E2connectionSetupFailedItem
        return true //self.E2connectionSetupFailedItem
   }
   return false//ObjSet
}
func (self *E2connectionSetupFailedItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 41: //E2connectionSetupFailedItem
        self.E2connectionSetupFailedItem.Unpack(st)
        self.list = append(self.list, &self.E2connectionSetupFailedItem)
   }
}
func (self *E2connectionSetupFailedItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 41: //E2connectionSetupFailedItem
        self.E2connectionSetupFailedItem.Pack(st)
      default:
      break
   }
}
func init() {
table_E2connectionSetupFailedItemIEs[41] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2connectionSetupFailedItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&E2connectionSetupFailedItem{}, PRESENCE:Presence{Presencemandatory}, }
order_E2connectionSetupFailedItemIEs[0] = 41
   }

type E2connectionUpdateFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'reject', 'TYPE': 'Cause', 'PRESENCE': 'optional'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   Cause  *Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *E2connectionUpdateFailureIEs)createOT() interface{}{
    return nil
}
var table_E2connectionUpdateFailureIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2connectionUpdateFailureIEs = make([]int, 4)

func (self *E2connectionUpdateFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.Cause != nil { count += 1 }
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *E2connectionUpdateFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 1: //Cause
        if self.Cause != nil { return true }
      case 31: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *E2connectionUpdateFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 1: //Cause
        self.Cause = &Cause{}
        self.Cause.Unpack(st)
        self.list = append(self.list, self.Cause)
      case 31: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *E2connectionUpdateFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 1: //Cause
        if self.Cause != nil {self.Cause.Pack(st)}
      case 31: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_E2connectionUpdateFailureIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2connectionUpdateFailureIEs[0] = 49
table_E2connectionUpdateFailureIEs[1] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&Cause{}, PRESENCE:Presence{Presenceoptional}, }
order_E2connectionUpdateFailureIEs[1] = 1
table_E2connectionUpdateFailureIEs[31] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_E2connectionUpdateFailureIEs[2] = 31
table_E2connectionUpdateFailureIEs[2] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_E2connectionUpdateFailureIEs[3] = 2
   }

type E2nodeConfigurationUpdateIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalE2node-ID', 'CRITICALITY': 'reject', 'TYPE': 'GlobalE2node-ID', 'PRESENCE': 'optional'}, {'ID': 'id-E2nodeComponentConfigAddition', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigAddition-List', 'PRESENCE': 'optional'}, {'ID': 'id-E2nodeComponentConfigUpdate', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigUpdate-List', 'PRESENCE': 'optional'}, {'ID': 'id-E2nodeComponentConfigRemoval', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigRemoval-List', 'PRESENCE': 'optional'}, {'ID': 'id-E2nodeTNLassociationRemoval', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeTNLassociationRemoval-List', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   GlobalE2nodeID  *GlobalE2nodeID
   E2nodeComponentConfigAddition  *E2nodeComponentConfigAdditionList
   E2nodeComponentConfigUpdate  *E2nodeComponentConfigUpdateList
   E2nodeComponentConfigRemoval  *E2nodeComponentConfigRemovalList
   E2nodeTNLassociationRemoval  *E2nodeTNLassociationRemovalList
   list []interface{}
}
func (self *E2nodeConfigurationUpdateIEs)createOT() interface{}{
    return nil
}
var table_E2nodeConfigurationUpdateIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2nodeConfigurationUpdateIEs = make([]int, 6)

func (self *E2nodeConfigurationUpdateIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.GlobalE2nodeID != nil { count += 1 }
   if self.E2nodeComponentConfigAddition != nil { count += 1 }
   if self.E2nodeComponentConfigUpdate != nil { count += 1 }
   if self.E2nodeComponentConfigRemoval != nil { count += 1 }
   if self.E2nodeTNLassociationRemoval != nil { count += 1 }
   return count//ObjSet
}
func (self *E2nodeConfigurationUpdateIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 3: //GlobalE2nodeID
        if self.GlobalE2nodeID != nil { return true }
      case 50: //E2nodeComponentConfigAddition
        if self.E2nodeComponentConfigAddition != nil { return true }
      case 33: //E2nodeComponentConfigUpdate
        if self.E2nodeComponentConfigUpdate != nil { return true }
      case 54: //E2nodeComponentConfigRemoval
        if self.E2nodeComponentConfigRemoval != nil { return true }
      case 58: //E2nodeTNLassociationRemoval
        if self.E2nodeTNLassociationRemoval != nil { return true }
   }
   return false//ObjSet
}
func (self *E2nodeConfigurationUpdateIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 3: //GlobalE2nodeID
        self.GlobalE2nodeID = &GlobalE2nodeID{}
        self.GlobalE2nodeID.Unpack(st)
        self.list = append(self.list, self.GlobalE2nodeID)
      case 50: //E2nodeComponentConfigAddition
        self.E2nodeComponentConfigAddition = &E2nodeComponentConfigAdditionList{}
        self.E2nodeComponentConfigAddition.Unpack(st)
        self.list = append(self.list, self.E2nodeComponentConfigAddition)
      case 33: //E2nodeComponentConfigUpdate
        self.E2nodeComponentConfigUpdate = &E2nodeComponentConfigUpdateList{}
        self.E2nodeComponentConfigUpdate.Unpack(st)
        self.list = append(self.list, self.E2nodeComponentConfigUpdate)
      case 54: //E2nodeComponentConfigRemoval
        self.E2nodeComponentConfigRemoval = &E2nodeComponentConfigRemovalList{}
        self.E2nodeComponentConfigRemoval.Unpack(st)
        self.list = append(self.list, self.E2nodeComponentConfigRemoval)
      case 58: //E2nodeTNLassociationRemoval
        self.E2nodeTNLassociationRemoval = &E2nodeTNLassociationRemovalList{}
        self.E2nodeTNLassociationRemoval.Unpack(st)
        self.list = append(self.list, self.E2nodeTNLassociationRemoval)
   }
}
func (self *E2nodeConfigurationUpdateIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 3: //GlobalE2nodeID
        if self.GlobalE2nodeID != nil {self.GlobalE2nodeID.Pack(st)}
      case 50: //E2nodeComponentConfigAddition
        if self.E2nodeComponentConfigAddition != nil {self.E2nodeComponentConfigAddition.Pack(st)}
      case 33: //E2nodeComponentConfigUpdate
        if self.E2nodeComponentConfigUpdate != nil {self.E2nodeComponentConfigUpdate.Pack(st)}
      case 54: //E2nodeComponentConfigRemoval
        if self.E2nodeComponentConfigRemoval != nil {self.E2nodeComponentConfigRemoval.Pack(st)}
      case 58: //E2nodeTNLassociationRemoval
        if self.E2nodeTNLassociationRemoval != nil {self.E2nodeTNLassociationRemoval.Pack(st)}
      default:
      break
   }
}
func init() {
table_E2nodeConfigurationUpdateIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeConfigurationUpdateIEs[0] = 49
table_E2nodeConfigurationUpdateIEs[3] = &E2APPROTOCOLIES{ID:ProtocolIEID{idGlobalE2nodeID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GlobalE2nodeID{}, PRESENCE:Presence{Presenceoptional}, }
order_E2nodeConfigurationUpdateIEs[1] = 3
table_E2nodeConfigurationUpdateIEs[50] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigAddition}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigAdditionList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2nodeConfigurationUpdateIEs[2] = 50
table_E2nodeConfigurationUpdateIEs[33] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigUpdate}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigUpdateList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2nodeConfigurationUpdateIEs[3] = 33
table_E2nodeConfigurationUpdateIEs[54] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigRemoval}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigRemovalList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2nodeConfigurationUpdateIEs[4] = 54
table_E2nodeConfigurationUpdateIEs[58] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeTNLassociationRemoval}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeTNLassociationRemovalList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2nodeConfigurationUpdateIEs[5] = 58
   }

type E2nodeComponentConfigAdditionItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-E2nodeComponentConfigAddition-Item', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigAddition-Item', 'PRESENCE': 'mandatory'}, None]}
   E2nodeComponentConfigAdditionItem  E2nodeComponentConfigAdditionItem
   list []interface{}
}
func (self *E2nodeComponentConfigAdditionItemIEs)createOT() interface{}{
    return nil
}
var table_E2nodeComponentConfigAdditionItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2nodeComponentConfigAdditionItemIEs = make([]int, 1)

func (self *E2nodeComponentConfigAdditionItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.E2nodeComponentConfigAdditionItem
   return count//ObjSet
}
func (self *E2nodeComponentConfigAdditionItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 51: //E2nodeComponentConfigAdditionItem
        return true //self.E2nodeComponentConfigAdditionItem
   }
   return false//ObjSet
}
func (self *E2nodeComponentConfigAdditionItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 51: //E2nodeComponentConfigAdditionItem
        self.E2nodeComponentConfigAdditionItem.Unpack(st)
        self.list = append(self.list, &self.E2nodeComponentConfigAdditionItem)
   }
}
func (self *E2nodeComponentConfigAdditionItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 51: //E2nodeComponentConfigAdditionItem
        self.E2nodeComponentConfigAdditionItem.Pack(st)
      default:
      break
   }
}
func init() {
table_E2nodeComponentConfigAdditionItemIEs[51] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigAdditionItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigAdditionItem{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeComponentConfigAdditionItemIEs[0] = 51
   }

type E2nodeComponentConfigUpdateItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-E2nodeComponentConfigUpdate-Item', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigUpdate-Item', 'PRESENCE': 'mandatory'}, None]}
   E2nodeComponentConfigUpdateItem  E2nodeComponentConfigUpdateItem
   list []interface{}
}
func (self *E2nodeComponentConfigUpdateItemIEs)createOT() interface{}{
    return nil
}
var table_E2nodeComponentConfigUpdateItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2nodeComponentConfigUpdateItemIEs = make([]int, 1)

func (self *E2nodeComponentConfigUpdateItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.E2nodeComponentConfigUpdateItem
   return count//ObjSet
}
func (self *E2nodeComponentConfigUpdateItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 34: //E2nodeComponentConfigUpdateItem
        return true //self.E2nodeComponentConfigUpdateItem
   }
   return false//ObjSet
}
func (self *E2nodeComponentConfigUpdateItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 34: //E2nodeComponentConfigUpdateItem
        self.E2nodeComponentConfigUpdateItem.Unpack(st)
        self.list = append(self.list, &self.E2nodeComponentConfigUpdateItem)
   }
}
func (self *E2nodeComponentConfigUpdateItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 34: //E2nodeComponentConfigUpdateItem
        self.E2nodeComponentConfigUpdateItem.Pack(st)
      default:
      break
   }
}
func init() {
table_E2nodeComponentConfigUpdateItemIEs[34] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigUpdateItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigUpdateItem{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeComponentConfigUpdateItemIEs[0] = 34
   }

type E2nodeComponentConfigRemovalItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-E2nodeComponentConfigRemoval-Item', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigRemoval-Item', 'PRESENCE': 'mandatory'}, None]}
   E2nodeComponentConfigRemovalItem  E2nodeComponentConfigRemovalItem
   list []interface{}
}
func (self *E2nodeComponentConfigRemovalItemIEs)createOT() interface{}{
    return nil
}
var table_E2nodeComponentConfigRemovalItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2nodeComponentConfigRemovalItemIEs = make([]int, 1)

func (self *E2nodeComponentConfigRemovalItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.E2nodeComponentConfigRemovalItem
   return count//ObjSet
}
func (self *E2nodeComponentConfigRemovalItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 55: //E2nodeComponentConfigRemovalItem
        return true //self.E2nodeComponentConfigRemovalItem
   }
   return false//ObjSet
}
func (self *E2nodeComponentConfigRemovalItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 55: //E2nodeComponentConfigRemovalItem
        self.E2nodeComponentConfigRemovalItem.Unpack(st)
        self.list = append(self.list, &self.E2nodeComponentConfigRemovalItem)
   }
}
func (self *E2nodeComponentConfigRemovalItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 55: //E2nodeComponentConfigRemovalItem
        self.E2nodeComponentConfigRemovalItem.Pack(st)
      default:
      break
   }
}
func init() {
table_E2nodeComponentConfigRemovalItemIEs[55] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigRemovalItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigRemovalItem{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeComponentConfigRemovalItemIEs[0] = 55
   }

type E2nodeTNLassociationRemovalItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-E2nodeTNLassociationRemoval-Item', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeTNLassociationRemoval-Item', 'PRESENCE': 'mandatory'}, None]}
   E2nodeTNLassociationRemovalItem  E2nodeTNLassociationRemovalItem
   list []interface{}
}
func (self *E2nodeTNLassociationRemovalItemIEs)createOT() interface{}{
    return nil
}
var table_E2nodeTNLassociationRemovalItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2nodeTNLassociationRemovalItemIEs = make([]int, 1)

func (self *E2nodeTNLassociationRemovalItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.E2nodeTNLassociationRemovalItem
   return count//ObjSet
}
func (self *E2nodeTNLassociationRemovalItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 59: //E2nodeTNLassociationRemovalItem
        return true //self.E2nodeTNLassociationRemovalItem
   }
   return false//ObjSet
}
func (self *E2nodeTNLassociationRemovalItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 59: //E2nodeTNLassociationRemovalItem
        self.E2nodeTNLassociationRemovalItem.Unpack(st)
        self.list = append(self.list, &self.E2nodeTNLassociationRemovalItem)
   }
}
func (self *E2nodeTNLassociationRemovalItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 59: //E2nodeTNLassociationRemovalItem
        self.E2nodeTNLassociationRemovalItem.Pack(st)
      default:
      break
   }
}
func init() {
table_E2nodeTNLassociationRemovalItemIEs[59] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeTNLassociationRemovalItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeTNLassociationRemovalItem{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeTNLassociationRemovalItemIEs[0] = 59
   }

type E2nodeConfigurationUpdateAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-E2nodeComponentConfigAdditionAck', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigAdditionAck-List', 'PRESENCE': 'optional'}, {'ID': 'id-E2nodeComponentConfigUpdateAck', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigUpdateAck-List', 'PRESENCE': 'optional'}, {'ID': 'id-E2nodeComponentConfigRemovalAck', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigRemovalAck-List', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   E2nodeComponentConfigAdditionAck  *E2nodeComponentConfigAdditionAckList
   E2nodeComponentConfigUpdateAck  *E2nodeComponentConfigUpdateAckList
   E2nodeComponentConfigRemovalAck  *E2nodeComponentConfigRemovalAckList
   list []interface{}
}
func (self *E2nodeConfigurationUpdateAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_E2nodeConfigurationUpdateAcknowledgeIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2nodeConfigurationUpdateAcknowledgeIEs = make([]int, 4)

func (self *E2nodeConfigurationUpdateAcknowledgeIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.E2nodeComponentConfigAdditionAck != nil { count += 1 }
   if self.E2nodeComponentConfigUpdateAck != nil { count += 1 }
   if self.E2nodeComponentConfigRemovalAck != nil { count += 1 }
   return count//ObjSet
}
func (self *E2nodeConfigurationUpdateAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 52: //E2nodeComponentConfigAdditionAck
        if self.E2nodeComponentConfigAdditionAck != nil { return true }
      case 35: //E2nodeComponentConfigUpdateAck
        if self.E2nodeComponentConfigUpdateAck != nil { return true }
      case 56: //E2nodeComponentConfigRemovalAck
        if self.E2nodeComponentConfigRemovalAck != nil { return true }
   }
   return false//ObjSet
}
func (self *E2nodeConfigurationUpdateAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 52: //E2nodeComponentConfigAdditionAck
        self.E2nodeComponentConfigAdditionAck = &E2nodeComponentConfigAdditionAckList{}
        self.E2nodeComponentConfigAdditionAck.Unpack(st)
        self.list = append(self.list, self.E2nodeComponentConfigAdditionAck)
      case 35: //E2nodeComponentConfigUpdateAck
        self.E2nodeComponentConfigUpdateAck = &E2nodeComponentConfigUpdateAckList{}
        self.E2nodeComponentConfigUpdateAck.Unpack(st)
        self.list = append(self.list, self.E2nodeComponentConfigUpdateAck)
      case 56: //E2nodeComponentConfigRemovalAck
        self.E2nodeComponentConfigRemovalAck = &E2nodeComponentConfigRemovalAckList{}
        self.E2nodeComponentConfigRemovalAck.Unpack(st)
        self.list = append(self.list, self.E2nodeComponentConfigRemovalAck)
   }
}
func (self *E2nodeConfigurationUpdateAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 52: //E2nodeComponentConfigAdditionAck
        if self.E2nodeComponentConfigAdditionAck != nil {self.E2nodeComponentConfigAdditionAck.Pack(st)}
      case 35: //E2nodeComponentConfigUpdateAck
        if self.E2nodeComponentConfigUpdateAck != nil {self.E2nodeComponentConfigUpdateAck.Pack(st)}
      case 56: //E2nodeComponentConfigRemovalAck
        if self.E2nodeComponentConfigRemovalAck != nil {self.E2nodeComponentConfigRemovalAck.Pack(st)}
      default:
      break
   }
}
func init() {
table_E2nodeConfigurationUpdateAcknowledgeIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeConfigurationUpdateAcknowledgeIEs[0] = 49
table_E2nodeConfigurationUpdateAcknowledgeIEs[52] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigAdditionAck}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigAdditionAckList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2nodeConfigurationUpdateAcknowledgeIEs[1] = 52
table_E2nodeConfigurationUpdateAcknowledgeIEs[35] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigUpdateAck}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigUpdateAckList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2nodeConfigurationUpdateAcknowledgeIEs[2] = 35
table_E2nodeConfigurationUpdateAcknowledgeIEs[56] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigRemovalAck}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigRemovalAckList{}, PRESENCE:Presence{Presenceoptional}, }
order_E2nodeConfigurationUpdateAcknowledgeIEs[3] = 56
   }

type E2nodeComponentConfigAdditionAckItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-E2nodeComponentConfigAdditionAck-Item', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigAdditionAck-Item', 'PRESENCE': 'mandatory'}, None]}
   E2nodeComponentConfigAdditionAckItem  E2nodeComponentConfigAdditionAckItem
   list []interface{}
}
func (self *E2nodeComponentConfigAdditionAckItemIEs)createOT() interface{}{
    return nil
}
var table_E2nodeComponentConfigAdditionAckItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2nodeComponentConfigAdditionAckItemIEs = make([]int, 1)

func (self *E2nodeComponentConfigAdditionAckItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.E2nodeComponentConfigAdditionAckItem
   return count//ObjSet
}
func (self *E2nodeComponentConfigAdditionAckItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 53: //E2nodeComponentConfigAdditionAckItem
        return true //self.E2nodeComponentConfigAdditionAckItem
   }
   return false//ObjSet
}
func (self *E2nodeComponentConfigAdditionAckItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 53: //E2nodeComponentConfigAdditionAckItem
        self.E2nodeComponentConfigAdditionAckItem.Unpack(st)
        self.list = append(self.list, &self.E2nodeComponentConfigAdditionAckItem)
   }
}
func (self *E2nodeComponentConfigAdditionAckItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 53: //E2nodeComponentConfigAdditionAckItem
        self.E2nodeComponentConfigAdditionAckItem.Pack(st)
      default:
      break
   }
}
func init() {
table_E2nodeComponentConfigAdditionAckItemIEs[53] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigAdditionAckItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigAdditionAckItem{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeComponentConfigAdditionAckItemIEs[0] = 53
   }

type E2nodeComponentConfigUpdateAckItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-E2nodeComponentConfigUpdateAck-Item', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigUpdateAck-Item', 'PRESENCE': 'mandatory'}, None]}
   E2nodeComponentConfigUpdateAckItem  E2nodeComponentConfigUpdateAckItem
   list []interface{}
}
func (self *E2nodeComponentConfigUpdateAckItemIEs)createOT() interface{}{
    return nil
}
var table_E2nodeComponentConfigUpdateAckItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2nodeComponentConfigUpdateAckItemIEs = make([]int, 1)

func (self *E2nodeComponentConfigUpdateAckItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.E2nodeComponentConfigUpdateAckItem
   return count//ObjSet
}
func (self *E2nodeComponentConfigUpdateAckItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 36: //E2nodeComponentConfigUpdateAckItem
        return true //self.E2nodeComponentConfigUpdateAckItem
   }
   return false//ObjSet
}
func (self *E2nodeComponentConfigUpdateAckItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 36: //E2nodeComponentConfigUpdateAckItem
        self.E2nodeComponentConfigUpdateAckItem.Unpack(st)
        self.list = append(self.list, &self.E2nodeComponentConfigUpdateAckItem)
   }
}
func (self *E2nodeComponentConfigUpdateAckItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 36: //E2nodeComponentConfigUpdateAckItem
        self.E2nodeComponentConfigUpdateAckItem.Pack(st)
      default:
      break
   }
}
func init() {
table_E2nodeComponentConfigUpdateAckItemIEs[36] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigUpdateAckItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigUpdateAckItem{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeComponentConfigUpdateAckItemIEs[0] = 36
   }

type E2nodeComponentConfigRemovalAckItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-E2nodeComponentConfigRemovalAck-Item', 'CRITICALITY': 'reject', 'TYPE': 'E2nodeComponentConfigRemovalAck-Item', 'PRESENCE': 'mandatory'}, None]}
   E2nodeComponentConfigRemovalAckItem  E2nodeComponentConfigRemovalAckItem
   list []interface{}
}
func (self *E2nodeComponentConfigRemovalAckItemIEs)createOT() interface{}{
    return nil
}
var table_E2nodeComponentConfigRemovalAckItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2nodeComponentConfigRemovalAckItemIEs = make([]int, 1)

func (self *E2nodeComponentConfigRemovalAckItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.E2nodeComponentConfigRemovalAckItem
   return count//ObjSet
}
func (self *E2nodeComponentConfigRemovalAckItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 57: //E2nodeComponentConfigRemovalAckItem
        return true //self.E2nodeComponentConfigRemovalAckItem
   }
   return false//ObjSet
}
func (self *E2nodeComponentConfigRemovalAckItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 57: //E2nodeComponentConfigRemovalAckItem
        self.E2nodeComponentConfigRemovalAckItem.Unpack(st)
        self.list = append(self.list, &self.E2nodeComponentConfigRemovalAckItem)
   }
}
func (self *E2nodeComponentConfigRemovalAckItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 57: //E2nodeComponentConfigRemovalAckItem
        self.E2nodeComponentConfigRemovalAckItem.Pack(st)
      default:
      break
   }
}
func init() {
table_E2nodeComponentConfigRemovalAckItemIEs[57] = &E2APPROTOCOLIES{ID:ProtocolIEID{idE2nodeComponentConfigRemovalAckItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&E2nodeComponentConfigRemovalAckItem{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeComponentConfigRemovalAckItemIEs[0] = 57
   }

type E2nodeConfigurationUpdateFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *E2nodeConfigurationUpdateFailureIEs)createOT() interface{}{
    return nil
}
var table_E2nodeConfigurationUpdateFailureIEs = make(map[int]*E2APPROTOCOLIES)

var order_E2nodeConfigurationUpdateFailureIEs = make([]int, 4)

func (self *E2nodeConfigurationUpdateFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *E2nodeConfigurationUpdateFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 1: //Cause
        return true //self.Cause
      case 31: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *E2nodeConfigurationUpdateFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 31: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *E2nodeConfigurationUpdateFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 1: //Cause
        self.Cause.Pack(st)
      case 31: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_E2nodeConfigurationUpdateFailureIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeConfigurationUpdateFailureIEs[0] = 49
table_E2nodeConfigurationUpdateFailureIEs[1] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_E2nodeConfigurationUpdateFailureIEs[1] = 1
table_E2nodeConfigurationUpdateFailureIEs[31] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_E2nodeConfigurationUpdateFailureIEs[2] = 31
table_E2nodeConfigurationUpdateFailureIEs[2] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_E2nodeConfigurationUpdateFailureIEs[3] = 2
   }

type ResetRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, None]}
   TransactionID  TransactionID
   Cause  Cause
   list []interface{}
}
func (self *ResetRequestIEs)createOT() interface{}{
    return nil
}
var table_ResetRequestIEs = make(map[int]*E2APPROTOCOLIES)

var order_ResetRequestIEs = make([]int, 2)

func (self *ResetRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.Cause
   return count//ObjSet
}
func (self *ResetRequestIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 1: //Cause
        return true //self.Cause
   }
   return false//ObjSet
}
func (self *ResetRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
   }
}
func (self *ResetRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 1: //Cause
        self.Cause.Pack(st)
      default:
      break
   }
}
func init() {
table_ResetRequestIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetRequestIEs[0] = 49
table_ResetRequestIEs[1] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetRequestIEs[1] = 1
   }

type ResetResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ResetResponseIEs)createOT() interface{}{
    return nil
}
var table_ResetResponseIEs = make(map[int]*E2APPROTOCOLIES)

var order_ResetResponseIEs = make([]int, 2)

func (self *ResetResponseIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetResponseIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ResetResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetResponseIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetResponseIEs[0] = 49
table_ResetResponseIEs[2] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetResponseIEs[1] = 2
   }

type RICserviceUpdateIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionsAdded', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctions-List', 'PRESENCE': 'optional'}, {'ID': 'id-RANfunctionsModified', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctions-List', 'PRESENCE': 'optional'}, {'ID': 'id-RANfunctionsDeleted', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionsID-List', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   RANfunctionsAdded  *RANfunctionsList
   RANfunctionsModified  *RANfunctionsList
   RANfunctionsDeleted  *RANfunctionsIDList
   list []interface{}
}
func (self *RICserviceUpdateIEs)createOT() interface{}{
    return nil
}
var table_RICserviceUpdateIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICserviceUpdateIEs = make([]int, 4)

func (self *RICserviceUpdateIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.RANfunctionsAdded != nil { count += 1 }
   if self.RANfunctionsModified != nil { count += 1 }
   if self.RANfunctionsDeleted != nil { count += 1 }
   return count//ObjSet
}
func (self *RICserviceUpdateIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 10: //RANfunctionsAdded
        if self.RANfunctionsAdded != nil { return true }
      case 12: //RANfunctionsModified
        if self.RANfunctionsModified != nil { return true }
      case 11: //RANfunctionsDeleted
        if self.RANfunctionsDeleted != nil { return true }
   }
   return false//ObjSet
}
func (self *RICserviceUpdateIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 10: //RANfunctionsAdded
        self.RANfunctionsAdded = &RANfunctionsList{}
        self.RANfunctionsAdded.Unpack(st)
        self.list = append(self.list, self.RANfunctionsAdded)
      case 12: //RANfunctionsModified
        self.RANfunctionsModified = &RANfunctionsList{}
        self.RANfunctionsModified.Unpack(st)
        self.list = append(self.list, self.RANfunctionsModified)
      case 11: //RANfunctionsDeleted
        self.RANfunctionsDeleted = &RANfunctionsIDList{}
        self.RANfunctionsDeleted.Unpack(st)
        self.list = append(self.list, self.RANfunctionsDeleted)
   }
}
func (self *RICserviceUpdateIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 10: //RANfunctionsAdded
        if self.RANfunctionsAdded != nil {self.RANfunctionsAdded.Pack(st)}
      case 12: //RANfunctionsModified
        if self.RANfunctionsModified != nil {self.RANfunctionsModified.Pack(st)}
      case 11: //RANfunctionsDeleted
        if self.RANfunctionsDeleted != nil {self.RANfunctionsDeleted.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICserviceUpdateIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICserviceUpdateIEs[0] = 49
table_RICserviceUpdateIEs[10] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionsAdded}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionsList{}, PRESENCE:Presence{Presenceoptional}, }
order_RICserviceUpdateIEs[1] = 10
table_RICserviceUpdateIEs[12] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionsModified}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionsList{}, PRESENCE:Presence{Presenceoptional}, }
order_RICserviceUpdateIEs[2] = 12
table_RICserviceUpdateIEs[11] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionsDeleted}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionsIDList{}, PRESENCE:Presence{Presenceoptional}, }
order_RICserviceUpdateIEs[3] = 11
   }

type RANfunctionItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RANfunction-Item', 'CRITICALITY': 'ignore', 'TYPE': 'RANfunction-Item', 'PRESENCE': 'mandatory'}, None]}
   RANfunctionItem  RANfunctionItem
   list []interface{}
}
func (self *RANfunctionItemIEs)createOT() interface{}{
    return nil
}
var table_RANfunctionItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_RANfunctionItemIEs = make([]int, 1)

func (self *RANfunctionItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RANfunctionItem
   return count//ObjSet
}
func (self *RANfunctionItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 8: //RANfunctionItem
        return true //self.RANfunctionItem
   }
   return false//ObjSet
}
func (self *RANfunctionItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 8: //RANfunctionItem
        self.RANfunctionItem.Unpack(st)
        self.list = append(self.list, &self.RANfunctionItem)
   }
}
func (self *RANfunctionItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 8: //RANfunctionItem
        self.RANfunctionItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RANfunctionItemIEs[8] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RANfunctionItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RANfunctionItemIEs[0] = 8
   }

type RANfunctionIDItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RANfunctionID-Item', 'CRITICALITY': 'ignore', 'TYPE': 'RANfunctionID-Item', 'PRESENCE': 'mandatory'}, None]}
   RANfunctionIDItem  RANfunctionIDItem
   list []interface{}
}
func (self *RANfunctionIDItemIEs)createOT() interface{}{
    return nil
}
var table_RANfunctionIDItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_RANfunctionIDItemIEs = make([]int, 1)

func (self *RANfunctionIDItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RANfunctionIDItem
   return count//ObjSet
}
func (self *RANfunctionIDItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 6: //RANfunctionIDItem
        return true //self.RANfunctionIDItem
   }
   return false//ObjSet
}
func (self *RANfunctionIDItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 6: //RANfunctionIDItem
        self.RANfunctionIDItem.Unpack(st)
        self.list = append(self.list, &self.RANfunctionIDItem)
   }
}
func (self *RANfunctionIDItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 6: //RANfunctionIDItem
        self.RANfunctionIDItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RANfunctionIDItemIEs[6] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionIDItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RANfunctionIDItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RANfunctionIDItemIEs[0] = 6
   }

type RICserviceUpdateAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionsAccepted', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionsID-List', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionsRejected', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionsIDcause-List', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   RANfunctionsAccepted  RANfunctionsIDList
   RANfunctionsRejected  *RANfunctionsIDcauseList
   list []interface{}
}
func (self *RICserviceUpdateAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_RICserviceUpdateAcknowledgeIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICserviceUpdateAcknowledgeIEs = make([]int, 3)

func (self *RICserviceUpdateAcknowledgeIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.RANfunctionsAccepted
   if self.RANfunctionsRejected != nil { count += 1 }
   return count//ObjSet
}
func (self *RICserviceUpdateAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 9: //RANfunctionsAccepted
        return true //self.RANfunctionsAccepted
      case 13: //RANfunctionsRejected
        if self.RANfunctionsRejected != nil { return true }
   }
   return false//ObjSet
}
func (self *RICserviceUpdateAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 9: //RANfunctionsAccepted
        self.RANfunctionsAccepted.Unpack(st)
        self.list = append(self.list, &self.RANfunctionsAccepted)
      case 13: //RANfunctionsRejected
        self.RANfunctionsRejected = &RANfunctionsIDcauseList{}
        self.RANfunctionsRejected.Unpack(st)
        self.list = append(self.list, self.RANfunctionsRejected)
   }
}
func (self *RICserviceUpdateAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 9: //RANfunctionsAccepted
        self.RANfunctionsAccepted.Pack(st)
      case 13: //RANfunctionsRejected
        if self.RANfunctionsRejected != nil {self.RANfunctionsRejected.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICserviceUpdateAcknowledgeIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICserviceUpdateAcknowledgeIEs[0] = 49
table_RICserviceUpdateAcknowledgeIEs[9] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionsAccepted}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionsIDList{}, PRESENCE:Presence{Presencemandatory}, }
order_RICserviceUpdateAcknowledgeIEs[1] = 9
table_RICserviceUpdateAcknowledgeIEs[13] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionsRejected}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionsIDcauseList{}, PRESENCE:Presence{Presenceoptional}, }
order_RICserviceUpdateAcknowledgeIEs[2] = 13
   }

type RANfunctionIDcauseItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-RANfunctionIEcause-Item', 'CRITICALITY': 'ignore', 'TYPE': 'RANfunctionIDcause-Item', 'PRESENCE': 'mandatory'}, None]}
   RANfunctionIEcauseItem  RANfunctionIDcauseItem
   list []interface{}
}
func (self *RANfunctionIDcauseItemIEs)createOT() interface{}{
    return nil
}
var table_RANfunctionIDcauseItemIEs = make(map[int]*E2APPROTOCOLIES)

var order_RANfunctionIDcauseItemIEs = make([]int, 1)

func (self *RANfunctionIDcauseItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RANfunctionIEcauseItem
   return count//ObjSet
}
func (self *RANfunctionIDcauseItemIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 7: //RANfunctionIEcauseItem
        return true //self.RANfunctionIEcauseItem
   }
   return false//ObjSet
}
func (self *RANfunctionIDcauseItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 7: //RANfunctionIEcauseItem
        self.RANfunctionIEcauseItem.Unpack(st)
        self.list = append(self.list, &self.RANfunctionIEcauseItem)
   }
}
func (self *RANfunctionIDcauseItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 7: //RANfunctionIEcauseItem
        self.RANfunctionIEcauseItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RANfunctionIDcauseItemIEs[7] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionIEcauseItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RANfunctionIDcauseItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RANfunctionIDcauseItemIEs[0] = 7
   }

type RICserviceUpdateFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'reject', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-TimeToWait', 'CRITICALITY': 'ignore', 'TYPE': 'TimeToWait', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   Cause  Cause
   TimeToWait  *TimeToWait
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *RICserviceUpdateFailureIEs)createOT() interface{}{
    return nil
}
var table_RICserviceUpdateFailureIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICserviceUpdateFailureIEs = make([]int, 4)

func (self *RICserviceUpdateFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   count +=1 //self.Cause
   if self.TimeToWait != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *RICserviceUpdateFailureIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 1: //Cause
        return true //self.Cause
      case 31: //TimeToWait
        if self.TimeToWait != nil { return true }
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *RICserviceUpdateFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 1: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 31: //TimeToWait
        self.TimeToWait = &TimeToWait{}
        self.TimeToWait.Unpack(st)
        self.list = append(self.list, self.TimeToWait)
      case 2: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *RICserviceUpdateFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 1: //Cause
        self.Cause.Pack(st)
      case 31: //TimeToWait
        if self.TimeToWait != nil {self.TimeToWait.Pack(st)}
      case 2: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICserviceUpdateFailureIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICserviceUpdateFailureIEs[0] = 49
table_RICserviceUpdateFailureIEs[1] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_RICserviceUpdateFailureIEs[1] = 1
table_RICserviceUpdateFailureIEs[31] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTimeToWait}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TimeToWait{}, PRESENCE:Presence{Presenceoptional}, }
order_RICserviceUpdateFailureIEs[2] = 31
table_RICserviceUpdateFailureIEs[2] = &E2APPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_RICserviceUpdateFailureIEs[3] = 2
   }

type RICserviceQueryIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'E2AP-PROTOCOL-IES', 'members': [{'ID': 'id-TransactionID', 'CRITICALITY': 'reject', 'TYPE': 'TransactionID', 'PRESENCE': 'mandatory'}, {'ID': 'id-RANfunctionsAccepted', 'CRITICALITY': 'reject', 'TYPE': 'RANfunctionsID-List', 'PRESENCE': 'optional'}, None]}
   TransactionID  TransactionID
   RANfunctionsAccepted  *RANfunctionsIDList
   list []interface{}
}
func (self *RICserviceQueryIEs)createOT() interface{}{
    return nil
}
var table_RICserviceQueryIEs = make(map[int]*E2APPROTOCOLIES)

var order_RICserviceQueryIEs = make([]int, 2)

func (self *RICserviceQueryIEs) GetIECount() int{
   count := 0
   count +=1 //self.TransactionID
   if self.RANfunctionsAccepted != nil { count += 1 }
   return count//ObjSet
}
func (self *RICserviceQueryIEs) GetOT(id interface{}) bool{
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        return true //self.TransactionID
      case 9: //RANfunctionsAccepted
        if self.RANfunctionsAccepted != nil { return true }
   }
   return false//ObjSet
}
func (self *RICserviceQueryIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Unpack(st)
        self.list = append(self.list, &self.TransactionID)
      case 9: //RANfunctionsAccepted
        self.RANfunctionsAccepted = &RANfunctionsIDList{}
        self.RANfunctionsAccepted.Unpack(st)
        self.list = append(self.list, self.RANfunctionsAccepted)
   }
}
func (self *RICserviceQueryIEs)PackOT(st *Stream, id interface{}){
   cat := id.(E2APPROTOCOLIESid).Value
   switch cat {
      case 49: //TransactionID
        self.TransactionID.Pack(st)
      case 9: //RANfunctionsAccepted
        if self.RANfunctionsAccepted != nil {self.RANfunctionsAccepted.Pack(st)}
      default:
      break
   }
}
func init() {
table_RICserviceQueryIEs[49] = &E2APPROTOCOLIES{ID:ProtocolIEID{idTransactionID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TransactionID{}, PRESENCE:Presence{Presencemandatory}, }
order_RICserviceQueryIEs[0] = 49
table_RICserviceQueryIEs[9] = &E2APPROTOCOLIES{ID:ProtocolIEID{idRANfunctionsAccepted}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RANfunctionsIDList{}, PRESENCE:Presence{Presenceoptional}, }
order_RICserviceQueryIEs[1] = 9
   }

func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'E2connectionUpdate', 'SUCCESSFUL OUTCOME': 'E2connectionUpdateAcknowledge', 'UNSUCCESSFUL OUTCOME': 'E2connectionUpdateFailure', 'PROCEDURE CODE': 'id-E2connectionUpdate', 'CRITICALITY': 'reject'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idE2connectionUpdate}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&E2connectionUpdate{}, SUCCESSFULOUTCOME:&E2connectionUpdateAcknowledge{}, UNSUCCESSFULOUTCOME:&E2connectionUpdateFailure{}, PROCEDURECODE:ProcedureCode{idE2connectionUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetE2connectionUpdateINITIATINGMESSAGE() (*E2connectionUpdate, uint64, int) {/*TYPE, ID, Cricality*/
 return &E2connectionUpdate{}, uint64(idE2connectionUpdate), int(Criticalityreject)
}
func GetE2connectionUpdateSUCCESSFULOUTCOME() (*E2connectionUpdateAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &E2connectionUpdateAcknowledge{}, uint64(idE2connectionUpdate), int(Criticalityreject)
}
func GetE2connectionUpdateUNSUCCESSFULOUTCOME() (*E2connectionUpdateFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &E2connectionUpdateFailure{}, uint64(idE2connectionUpdate), int(Criticalityreject)
}
func (self *E2connectionUpdate) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E2connectionUpdate) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E2connectionUpdate) createOT() interface{} {
   return &E2connectionUpdate{}
}
func (self *E2connectionUpdate) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E2connectionUpdate) GetIECount() int{
    return 0
}
func (self *E2connectionUpdateAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E2connectionUpdateAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E2connectionUpdateAcknowledge) createOT() interface{} {
   return &E2connectionUpdateAcknowledge{}
}
func (self *E2connectionUpdateAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E2connectionUpdateAcknowledge) GetIECount() int{
    return 0
}
func (self *E2connectionUpdateFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E2connectionUpdateFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E2connectionUpdateFailure) createOT() interface{} {
   return &E2connectionUpdateFailure{}
}
func (self *E2connectionUpdateFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E2connectionUpdateFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'E2nodeConfigurationUpdate', 'SUCCESSFUL OUTCOME': 'E2nodeConfigurationUpdateAcknowledge', 'UNSUCCESSFUL OUTCOME': 'E2nodeConfigurationUpdateFailure', 'PROCEDURE CODE': 'id-E2nodeConfigurationUpdate', 'CRITICALITY': 'reject'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idE2nodeConfigurationUpdate}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&E2nodeConfigurationUpdate{}, SUCCESSFULOUTCOME:&E2nodeConfigurationUpdateAcknowledge{}, UNSUCCESSFULOUTCOME:&E2nodeConfigurationUpdateFailure{}, PROCEDURECODE:ProcedureCode{idE2nodeConfigurationUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetE2nodeConfigurationUpdateINITIATINGMESSAGE() (*E2nodeConfigurationUpdate, uint64, int) {/*TYPE, ID, Cricality*/
 return &E2nodeConfigurationUpdate{}, uint64(idE2nodeConfigurationUpdate), int(Criticalityreject)
}
func GetE2nodeConfigurationUpdateSUCCESSFULOUTCOME() (*E2nodeConfigurationUpdateAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &E2nodeConfigurationUpdateAcknowledge{}, uint64(idE2nodeConfigurationUpdate), int(Criticalityreject)
}
func GetE2nodeConfigurationUpdateUNSUCCESSFULOUTCOME() (*E2nodeConfigurationUpdateFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &E2nodeConfigurationUpdateFailure{}, uint64(idE2nodeConfigurationUpdate), int(Criticalityreject)
}
func (self *E2nodeConfigurationUpdate) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E2nodeConfigurationUpdate) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E2nodeConfigurationUpdate) createOT() interface{} {
   return &E2nodeConfigurationUpdate{}
}
func (self *E2nodeConfigurationUpdate) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E2nodeConfigurationUpdate) GetIECount() int{
    return 0
}
func (self *E2nodeConfigurationUpdateAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E2nodeConfigurationUpdateAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E2nodeConfigurationUpdateAcknowledge) createOT() interface{} {
   return &E2nodeConfigurationUpdateAcknowledge{}
}
func (self *E2nodeConfigurationUpdateAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E2nodeConfigurationUpdateAcknowledge) GetIECount() int{
    return 0
}
func (self *E2nodeConfigurationUpdateFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E2nodeConfigurationUpdateFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E2nodeConfigurationUpdateFailure) createOT() interface{} {
   return &E2nodeConfigurationUpdateFailure{}
}
func (self *E2nodeConfigurationUpdateFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E2nodeConfigurationUpdateFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'E2setupRequest', 'SUCCESSFUL OUTCOME': 'E2setupResponse', 'UNSUCCESSFUL OUTCOME': 'E2setupFailure', 'PROCEDURE CODE': 'id-E2setup', 'CRITICALITY': 'reject'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idE2setup}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&E2setupRequest{}, SUCCESSFULOUTCOME:&E2setupResponse{}, UNSUCCESSFULOUTCOME:&E2setupFailure{}, PROCEDURECODE:ProcedureCode{idE2setup}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetE2setupINITIATINGMESSAGE() (*E2setupRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &E2setupRequest{}, uint64(idE2setup), int(Criticalityreject)
}
func GetE2setupSUCCESSFULOUTCOME() (*E2setupResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &E2setupResponse{}, uint64(idE2setup), int(Criticalityreject)
}
func GetE2setupUNSUCCESSFULOUTCOME() (*E2setupFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &E2setupFailure{}, uint64(idE2setup), int(Criticalityreject)
}
func (self *E2setupRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E2setupRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E2setupRequest) createOT() interface{} {
   return &E2setupRequest{}
}
func (self *E2setupRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E2setupRequest) GetIECount() int{
    return 0
}
func (self *E2setupResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E2setupResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E2setupResponse) createOT() interface{} {
   return &E2setupResponse{}
}
func (self *E2setupResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E2setupResponse) GetIECount() int{
    return 0
}
func (self *E2setupFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *E2setupFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *E2setupFailure) createOT() interface{} {
   return &E2setupFailure{}
}
func (self *E2setupFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *E2setupFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ErrorIndication', 'PROCEDURE CODE': 'id-ErrorIndication', 'CRITICALITY': 'ignore'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idErrorIndication}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ErrorIndication{}, PROCEDURECODE:ProcedureCode{idErrorIndication}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ResetRequest', 'SUCCESSFUL OUTCOME': 'ResetResponse', 'PROCEDURE CODE': 'id-Reset', 'CRITICALITY': 'reject'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idReset}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ResetRequest{}, SUCCESSFULOUTCOME:&ResetResponse{}, PROCEDURECODE:ProcedureCode{idReset}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetResetINITIATINGMESSAGE() (*ResetRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &ResetRequest{}, uint64(idReset), int(Criticalityreject)
}
func GetResetSUCCESSFULOUTCOME() (*ResetResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &ResetResponse{}, uint64(idReset), int(Criticalityreject)
}
func (self *ResetRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ResetRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ResetRequest) createOT() interface{} {
   return &ResetRequest{}
}
func (self *ResetRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ResetRequest) GetIECount() int{
    return 0
}
func (self *ResetResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ResetResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ResetResponse) createOT() interface{} {
   return &ResetResponse{}
}
func (self *ResetResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ResetResponse) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RICcontrolRequest', 'SUCCESSFUL OUTCOME': 'RICcontrolAcknowledge', 'UNSUCCESSFUL OUTCOME': 'RICcontrolFailure', 'PROCEDURE CODE': 'id-RICcontrol', 'CRITICALITY': 'reject'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idRICcontrol}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RICcontrolRequest{}, SUCCESSFULOUTCOME:&RICcontrolAcknowledge{}, UNSUCCESSFULOUTCOME:&RICcontrolFailure{}, PROCEDURECODE:ProcedureCode{idRICcontrol}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetRicControlINITIATINGMESSAGE() (*RICcontrolRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICcontrolRequest{}, uint64(idRICcontrol), int(Criticalityreject)
}
func GetRicControlSUCCESSFULOUTCOME() (*RICcontrolAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICcontrolAcknowledge{}, uint64(idRICcontrol), int(Criticalityreject)
}
func GetRicControlUNSUCCESSFULOUTCOME() (*RICcontrolFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICcontrolFailure{}, uint64(idRICcontrol), int(Criticalityreject)
}
func (self *RICcontrolRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICcontrolRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICcontrolRequest) createOT() interface{} {
   return &RICcontrolRequest{}
}
func (self *RICcontrolRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICcontrolRequest) GetIECount() int{
    return 0
}
func (self *RICcontrolAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICcontrolAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICcontrolAcknowledge) createOT() interface{} {
   return &RICcontrolAcknowledge{}
}
func (self *RICcontrolAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICcontrolAcknowledge) GetIECount() int{
    return 0
}
func (self *RICcontrolFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICcontrolFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICcontrolFailure) createOT() interface{} {
   return &RICcontrolFailure{}
}
func (self *RICcontrolFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICcontrolFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RICindication', 'PROCEDURE CODE': 'id-RICindication', 'CRITICALITY': 'ignore'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idRICindication}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RICindication{}, PROCEDURECODE:ProcedureCode{idRICindication}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetRicIndicationINITIATINGMESSAGE() (*RICindication, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICindication{}, uint64(idRICindication), int(Criticalityignore)
}
func (self *RICindication) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICindication) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICindication) createOT() interface{} {
   return &RICindication{}
}
func (self *RICindication) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICindication) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RICserviceQuery', 'PROCEDURE CODE': 'id-RICserviceQuery', 'CRITICALITY': 'ignore'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idRICserviceQuery}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RICserviceQuery{}, PROCEDURECODE:ProcedureCode{idRICserviceQuery}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetRicServiceQueryINITIATINGMESSAGE() (*RICserviceQuery, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICserviceQuery{}, uint64(idRICserviceQuery), int(Criticalityignore)
}
func (self *RICserviceQuery) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICserviceQuery) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICserviceQuery) createOT() interface{} {
   return &RICserviceQuery{}
}
func (self *RICserviceQuery) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICserviceQuery) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RICserviceUpdate', 'SUCCESSFUL OUTCOME': 'RICserviceUpdateAcknowledge', 'UNSUCCESSFUL OUTCOME': 'RICserviceUpdateFailure', 'PROCEDURE CODE': 'id-RICserviceUpdate', 'CRITICALITY': 'reject'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idRICserviceUpdate}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RICserviceUpdate{}, SUCCESSFULOUTCOME:&RICserviceUpdateAcknowledge{}, UNSUCCESSFULOUTCOME:&RICserviceUpdateFailure{}, PROCEDURECODE:ProcedureCode{idRICserviceUpdate}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetRicServiceUpdateINITIATINGMESSAGE() (*RICserviceUpdate, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICserviceUpdate{}, uint64(idRICserviceUpdate), int(Criticalityreject)
}
func GetRicServiceUpdateSUCCESSFULOUTCOME() (*RICserviceUpdateAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICserviceUpdateAcknowledge{}, uint64(idRICserviceUpdate), int(Criticalityreject)
}
func GetRicServiceUpdateUNSUCCESSFULOUTCOME() (*RICserviceUpdateFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICserviceUpdateFailure{}, uint64(idRICserviceUpdate), int(Criticalityreject)
}
func (self *RICserviceUpdate) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICserviceUpdate) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICserviceUpdate) createOT() interface{} {
   return &RICserviceUpdate{}
}
func (self *RICserviceUpdate) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICserviceUpdate) GetIECount() int{
    return 0
}
func (self *RICserviceUpdateAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICserviceUpdateAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICserviceUpdateAcknowledge) createOT() interface{} {
   return &RICserviceUpdateAcknowledge{}
}
func (self *RICserviceUpdateAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICserviceUpdateAcknowledge) GetIECount() int{
    return 0
}
func (self *RICserviceUpdateFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICserviceUpdateFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICserviceUpdateFailure) createOT() interface{} {
   return &RICserviceUpdateFailure{}
}
func (self *RICserviceUpdateFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICserviceUpdateFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RICsubscriptionRequest', 'SUCCESSFUL OUTCOME': 'RICsubscriptionResponse', 'UNSUCCESSFUL OUTCOME': 'RICsubscriptionFailure', 'PROCEDURE CODE': 'id-RICsubscription', 'CRITICALITY': 'reject'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idRICsubscription}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RICsubscriptionRequest{}, SUCCESSFULOUTCOME:&RICsubscriptionResponse{}, UNSUCCESSFULOUTCOME:&RICsubscriptionFailure{}, PROCEDURECODE:ProcedureCode{idRICsubscription}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetRicSubscriptionINITIATINGMESSAGE() (*RICsubscriptionRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICsubscriptionRequest{}, uint64(idRICsubscription), int(Criticalityreject)
}
func GetRicSubscriptionSUCCESSFULOUTCOME() (*RICsubscriptionResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICsubscriptionResponse{}, uint64(idRICsubscription), int(Criticalityreject)
}
func GetRicSubscriptionUNSUCCESSFULOUTCOME() (*RICsubscriptionFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICsubscriptionFailure{}, uint64(idRICsubscription), int(Criticalityreject)
}
func (self *RICsubscriptionRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICsubscriptionRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICsubscriptionRequest) createOT() interface{} {
   return &RICsubscriptionRequest{}
}
func (self *RICsubscriptionRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICsubscriptionRequest) GetIECount() int{
    return 0
}
func (self *RICsubscriptionResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICsubscriptionResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICsubscriptionResponse) createOT() interface{} {
   return &RICsubscriptionResponse{}
}
func (self *RICsubscriptionResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICsubscriptionResponse) GetIECount() int{
    return 0
}
func (self *RICsubscriptionFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICsubscriptionFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICsubscriptionFailure) createOT() interface{} {
   return &RICsubscriptionFailure{}
}
func (self *RICsubscriptionFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICsubscriptionFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RICsubscriptionDeleteRequest', 'SUCCESSFUL OUTCOME': 'RICsubscriptionDeleteResponse', 'UNSUCCESSFUL OUTCOME': 'RICsubscriptionDeleteFailure', 'PROCEDURE CODE': 'id-RICsubscriptionDelete', 'CRITICALITY': 'reject'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idRICsubscriptionDelete}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RICsubscriptionDeleteRequest{}, SUCCESSFULOUTCOME:&RICsubscriptionDeleteResponse{}, UNSUCCESSFULOUTCOME:&RICsubscriptionDeleteFailure{}, PROCEDURECODE:ProcedureCode{idRICsubscriptionDelete}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetRicSubscriptionDeleteINITIATINGMESSAGE() (*RICsubscriptionDeleteRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICsubscriptionDeleteRequest{}, uint64(idRICsubscriptionDelete), int(Criticalityreject)
}
func GetRicSubscriptionDeleteSUCCESSFULOUTCOME() (*RICsubscriptionDeleteResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICsubscriptionDeleteResponse{}, uint64(idRICsubscriptionDelete), int(Criticalityreject)
}
func GetRicSubscriptionDeleteUNSUCCESSFULOUTCOME() (*RICsubscriptionDeleteFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICsubscriptionDeleteFailure{}, uint64(idRICsubscriptionDelete), int(Criticalityreject)
}
func (self *RICsubscriptionDeleteRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICsubscriptionDeleteRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICsubscriptionDeleteRequest) createOT() interface{} {
   return &RICsubscriptionDeleteRequest{}
}
func (self *RICsubscriptionDeleteRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICsubscriptionDeleteRequest) GetIECount() int{
    return 0
}
func (self *RICsubscriptionDeleteResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICsubscriptionDeleteResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICsubscriptionDeleteResponse) createOT() interface{} {
   return &RICsubscriptionDeleteResponse{}
}
func (self *RICsubscriptionDeleteResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICsubscriptionDeleteResponse) GetIECount() int{
    return 0
}
func (self *RICsubscriptionDeleteFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICsubscriptionDeleteFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICsubscriptionDeleteFailure) createOT() interface{} {
   return &RICsubscriptionDeleteFailure{}
}
func (self *RICsubscriptionDeleteFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICsubscriptionDeleteFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'E2AP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RICsubscriptionDeleteRequired', 'PROCEDURE CODE': 'id-RICsubscriptionDeleteRequired', 'CRITICALITY': 'ignore'}]}
table_E2APELEMENTARYPROCEDURES[E2APELEMENTARYPROCEDUREprocedureCode{idRICsubscriptionDeleteRequired}] = &E2APELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RICsubscriptionDeleteRequired{}, PROCEDURECODE:ProcedureCode{idRICsubscriptionDeleteRequired}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetRicSubscriptionDeleteRequiredINITIATINGMESSAGE() (*RICsubscriptionDeleteRequired, uint64, int) {/*TYPE, ID, Cricality*/
 return &RICsubscriptionDeleteRequired{}, uint64(idRICsubscriptionDeleteRequired), int(Criticalityignore)
}
func (self *RICsubscriptionDeleteRequired) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RICsubscriptionDeleteRequired) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RICsubscriptionDeleteRequired) createOT() interface{} {
   return &RICsubscriptionDeleteRequired{}
}
func (self *RICsubscriptionDeleteRequired) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RICsubscriptionDeleteRequired) GetIECount() int{
    return 0
}
var idE2setup uint64 = 1
const ProcedureCodeE2setup = 1
var idErrorIndication uint64 = 2
const ProcedureCodeErrorIndication = 2
var idReset uint64 = 3
const ProcedureCodeReset = 3
var idRICcontrol uint64 = 4
const ProcedureCodeRICcontrol = 4
var idRICindication uint64 = 5
const ProcedureCodeRICindication = 5
var idRICserviceQuery uint64 = 6
const ProcedureCodeRICserviceQuery = 6
var idRICserviceUpdate uint64 = 7
const ProcedureCodeRICserviceUpdate = 7
var idRICsubscription uint64 = 8
const ProcedureCodeRICsubscription = 8
var idRICsubscriptionDelete uint64 = 9
const ProcedureCodeRICsubscriptionDelete = 9
var idE2nodeConfigurationUpdate uint64 = 10
const ProcedureCodeE2nodeConfigurationUpdate = 10
var idE2connectionUpdate uint64 = 11
const ProcedureCodeE2connectionUpdate = 11
var idRICsubscriptionDeleteRequired uint64 = 12
const ProcedureCodeRICsubscriptionDeleteRequired = 12
var maxProtocolIEs uint64 = 65535
var maxnoofErrors uint64 = 256
var maxofE2nodeComponents uint64 = 1024
var maxofRANfunctionID uint64 = 256
var maxofRICactionID uint64 = 16
var maxofTNLA uint64 = 32
var maxofRICrequestID uint64 = 4294967295
var idCause uint64 = 1
const ProtocolIEIDCause = 1
var idCriticalityDiagnostics uint64 = 2
const ProtocolIEIDCriticalityDiagnostics = 2
var idGlobalE2nodeID uint64 = 3
const ProtocolIEIDGlobalE2nodeID = 3
var idGlobalRICID uint64 = 4
const ProtocolIEIDGlobalRICID = 4
var idRANfunctionID uint64 = 5
const ProtocolIEIDRANfunctionID = 5
var idRANfunctionIDItem uint64 = 6
const ProtocolIEIDRANfunctionIDItem = 6
var idRANfunctionIEcauseItem uint64 = 7
const ProtocolIEIDRANfunctionIEcauseItem = 7
var idRANfunctionItem uint64 = 8
const ProtocolIEIDRANfunctionItem = 8
var idRANfunctionsAccepted uint64 = 9
const ProtocolIEIDRANfunctionsAccepted = 9
var idRANfunctionsAdded uint64 = 10
const ProtocolIEIDRANfunctionsAdded = 10
var idRANfunctionsDeleted uint64 = 11
const ProtocolIEIDRANfunctionsDeleted = 11
var idRANfunctionsModified uint64 = 12
const ProtocolIEIDRANfunctionsModified = 12
var idRANfunctionsRejected uint64 = 13
const ProtocolIEIDRANfunctionsRejected = 13
var idRICactionAdmittedItem uint64 = 14
const ProtocolIEIDRICactionAdmittedItem = 14
var idRICactionID uint64 = 15
const ProtocolIEIDRICactionID = 15
var idRICactionNotAdmittedItem uint64 = 16
const ProtocolIEIDRICactionNotAdmittedItem = 16
var idRICactionsAdmitted uint64 = 17
const ProtocolIEIDRICactionsAdmitted = 17
var idRICactionsNotAdmitted uint64 = 18
const ProtocolIEIDRICactionsNotAdmitted = 18
var idRICactionToBeSetupItem uint64 = 19
const ProtocolIEIDRICactionToBeSetupItem = 19
var idRICcallProcessID uint64 = 20
const ProtocolIEIDRICcallProcessID = 20
var idRICcontrolAckRequest uint64 = 21
const ProtocolIEIDRICcontrolAckRequest = 21
var idRICcontrolHeader uint64 = 22
const ProtocolIEIDRICcontrolHeader = 22
var idRICcontrolMessage uint64 = 23
const ProtocolIEIDRICcontrolMessage = 23
var idRICcontrolStatus uint64 = 24
const ProtocolIEIDRICcontrolStatus = 24
var idRICindicationHeader uint64 = 25
const ProtocolIEIDRICindicationHeader = 25
var idRICindicationMessage uint64 = 26
const ProtocolIEIDRICindicationMessage = 26
var idRICindicationSN uint64 = 27
const ProtocolIEIDRICindicationSN = 27
var idRICindicationType uint64 = 28
const ProtocolIEIDRICindicationType = 28
var idRICrequestID uint64 = 29
const ProtocolIEIDRICrequestID = 29
var idRICsubscriptionDetails uint64 = 30
const ProtocolIEIDRICsubscriptionDetails = 30
var idTimeToWait uint64 = 31
const ProtocolIEIDTimeToWait = 31
var idRICcontrolOutcome uint64 = 32
const ProtocolIEIDRICcontrolOutcome = 32
var idE2nodeComponentConfigUpdate uint64 = 33
const ProtocolIEIDE2nodeComponentConfigUpdate = 33
var idE2nodeComponentConfigUpdateItem uint64 = 34
const ProtocolIEIDE2nodeComponentConfigUpdateItem = 34
var idE2nodeComponentConfigUpdateAck uint64 = 35
const ProtocolIEIDE2nodeComponentConfigUpdateAck = 35
var idE2nodeComponentConfigUpdateAckItem uint64 = 36
const ProtocolIEIDE2nodeComponentConfigUpdateAckItem = 36
var idE2connectionSetup uint64 = 39
const ProtocolIEIDE2connectionSetup = 39
var idE2connectionSetupFailed uint64 = 40
const ProtocolIEIDE2connectionSetupFailed = 40
var idE2connectionSetupFailedItem uint64 = 41
const ProtocolIEIDE2connectionSetupFailedItem = 41
var idE2connectionFailedItem uint64 = 42
const ProtocolIEIDE2connectionFailedItem = 42
var idE2connectionUpdateItem uint64 = 43
const ProtocolIEIDE2connectionUpdateItem = 43
var idE2connectionUpdateAdd uint64 = 44
const ProtocolIEIDE2connectionUpdateAdd = 44
var idE2connectionUpdateModify uint64 = 45
const ProtocolIEIDE2connectionUpdateModify = 45
var idE2connectionUpdateRemove uint64 = 46
const ProtocolIEIDE2connectionUpdateRemove = 46
var idE2connectionUpdateRemoveItem uint64 = 47
const ProtocolIEIDE2connectionUpdateRemoveItem = 47
var idTNLinformation uint64 = 48
const ProtocolIEIDTNLinformation = 48
var idTransactionID uint64 = 49
const ProtocolIEIDTransactionID = 49
var idE2nodeComponentConfigAddition uint64 = 50
const ProtocolIEIDE2nodeComponentConfigAddition = 50
var idE2nodeComponentConfigAdditionItem uint64 = 51
const ProtocolIEIDE2nodeComponentConfigAdditionItem = 51
var idE2nodeComponentConfigAdditionAck uint64 = 52
const ProtocolIEIDE2nodeComponentConfigAdditionAck = 52
var idE2nodeComponentConfigAdditionAckItem uint64 = 53
const ProtocolIEIDE2nodeComponentConfigAdditionAckItem = 53
var idE2nodeComponentConfigRemoval uint64 = 54
const ProtocolIEIDE2nodeComponentConfigRemoval = 54
var idE2nodeComponentConfigRemovalItem uint64 = 55
const ProtocolIEIDE2nodeComponentConfigRemovalItem = 55
var idE2nodeComponentConfigRemovalAck uint64 = 56
const ProtocolIEIDE2nodeComponentConfigRemovalAck = 56
var idE2nodeComponentConfigRemovalAckItem uint64 = 57
const ProtocolIEIDE2nodeComponentConfigRemovalAckItem = 57
var idE2nodeTNLassociationRemoval uint64 = 58
const ProtocolIEIDE2nodeTNLassociationRemoval = 58
var idE2nodeTNLassociationRemovalItem uint64 = 59
const ProtocolIEIDE2nodeTNLassociationRemovalItem = 59
var idRICsubscriptionToBeRemoved uint64 = 50
const ProtocolIEIDRICsubscriptionToBeRemoved = 50
var idRICsubscriptionwithCauseItem uint64 = 51
const ProtocolIEIDRICsubscriptionwithCauseItem = 51
