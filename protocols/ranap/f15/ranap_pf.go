
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package ranap
import (
  log "github.com/sirupsen/logrus"
)
var version = "vf15"

func fmtranap() {log.Debug("ranap")}
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
    TriggeringMessageunsuccessfull_outcome = 2
    TriggeringMessageoutcome = 3
)
func (self *TriggeringMessage) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 4, 0)
}
func (self *TriggeringMessage) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 4, 0)
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
    _size := data.(RANAPPROTOCOLIES_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES']}
    Items map[int]*RANAPPROTOCOLIES
    order []int
}

type ProtocolIEField struct { // [{'type': 'RANAP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'RANAP-PROTOCOL-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'RANAP-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id RANAPPROTOCOLIESid
    Criticality RANAPPROTOCOLIEScriticality
    Value RANAPPROTOCOLIESValue
}

func (self * ProtocolIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'RANAP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := RANAPPROTOCOLIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(RANAPPROTOCOLIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'RANAP-PROTOCOL-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * RANAPPROTOCOLIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'RANAP-PROTOCOL-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (RANAPPROTOCOLIESid)(self.ID)
    if out.(RANAPPROTOCOLIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(RANAPPROTOCOLIES_IF).PackOT(stream, key)
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
    _size := data.(RANAPPROTOCOLIESPAIR_IF).GetIECount()
    stream.set_listsize(_size-0, 65536)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolIEContainerPair struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-FieldPair', 'actual-parameters': ['IEsSetParam']}, 'size': [(0, 'maxProtocolIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES-PAIR']}
    Items map[int]*RANAPPROTOCOLIESPAIR
    order []int
}

type ProtocolIEFieldPair struct { // [{'type': 'RANAP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'RANAP-PROTOCOL-IES-PAIR.&firstCriticality', 'table': ['IEsSetParam', ['id']], 'name': 'firstCriticality'}, {'type': 'RANAP-PROTOCOL-IES-PAIR.&FirstValue', 'table': ['IEsSetParam', ['id']], 'name': 'firstValue'}, {'type': 'RANAP-PROTOCOL-IES-PAIR.&secondCriticality', 'table': ['IEsSetParam', ['id']], 'name': 'secondCriticality'}, {'type': 'RANAP-PROTOCOL-IES-PAIR.&SecondValue', 'table': ['IEsSetParam', ['id']], 'name': 'secondValue'}]
    Id RANAPPROTOCOLIESPAIRid
    FirstCriticality RANAPPROTOCOLIESPAIRfirstCriticality
    FirstValue RANAPPROTOCOLIESPAIRFirstValue
    SecondCriticality RANAPPROTOCOLIESPAIRsecondCriticality
    SecondValue RANAPPROTOCOLIESPAIRSecondValue
}

func (self * ProtocolIEFieldPair) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'RANAP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := RANAPPROTOCOLIESPAIR{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.FirstCriticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(RANAPPROTOCOLIESPAIR_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'RANAP-PROTOCOL-IES-PAIR.&FirstValue', 'table': ['IEsSetParam', ['id']], 'name': 'firstValue'}
    self.SecondCriticality.Unpack(stream)//p9
    out.(RANAPPROTOCOLIESPAIR_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'RANAP-PROTOCOL-IES-PAIR.&SecondValue', 'table': ['IEsSetParam', ['id']], 'name': 'secondValue'}
    stream.set_location(location, _len)
    return
}

func (self * RANAPPROTOCOLIESPAIR) Pack(stream *Stream, out interface{}) {
    //table {'type': 'RANAP-PROTOCOL-IES-PAIR.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (RANAPPROTOCOLIESPAIRid)(self.ID)
    if out.(RANAPPROTOCOLIESPAIR_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.FIRSTCRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(RANAPPROTOCOLIESPAIR_IF).PackOT(stream, key)
    self.SECONDCRITICALITY.Pack(stream)
    out.(RANAPPROTOCOLIESPAIR_IF).PackOT(stream, key)
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


type ProtocolIEContainerList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-Container', 'actual-parameters': ['IEsSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IEsSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'RANAP-PROTOCOL-IES']}
    Items map[int]*RANAPPROTOCOLIES
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


type ProtocolIEContainerPairList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolIE-ContainerPair', 'actual-parameters': ['IEsSetParam']}, 'size': [('lowerBound', 'upperBound')], 'parameters': ['lowerBound', 'upperBound', 'IEsSetParam'], 'param-types': ['INTEGER', 'INTEGER', 'RANAP-PROTOCOL-IES-PAIR']}
    Items map[int]*RANAPPROTOCOLIESPAIR
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
    _size := data.(RANAPPROTOCOLEXTENSION_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type ProtocolExtensionContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ProtocolExtensionField', 'actual-parameters': ['ExtensionSetParam']}, 'size': [(1, 'maxProtocolExtensions')], 'parameters': ['ExtensionSetParam'], 'param-types': ['RANAP-PROTOCOL-EXTENSION']}
    Items map[int]*RANAPPROTOCOLEXTENSION
    order []int
}

type ProtocolExtensionField struct { // [{'type': 'RANAP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}, {'type': 'RANAP-PROTOCOL-EXTENSION.&criticality', 'table': ['ExtensionSetParam', ['id']], 'name': 'criticality'}, {'type': 'RANAP-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}]
    Id RANAPPROTOCOLEXTENSIONid
    Criticality RANAPPROTOCOLEXTENSIONcriticality
    ExtensionValue RANAPPROTOCOLEXTENSIONExtension
}

func (self * ProtocolExtensionField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'RANAP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := RANAPPROTOCOLEXTENSION{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(RANAPPROTOCOLEXTENSION_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'RANAP-PROTOCOL-EXTENSION.&Extension', 'table': ['ExtensionSetParam', ['id']], 'name': 'extensionValue'}
    stream.set_location(location, _len)
    return
}

func (self * RANAPPROTOCOLEXTENSION) Pack(stream *Stream, out interface{}) {
    //table {'type': 'RANAP-PROTOCOL-EXTENSION.&id', 'table': {'type': 'ExtensionSetParam'}, 'name': 'id'}
    key := (RANAPPROTOCOLEXTENSIONid)(self.ID)
    if out.(RANAPPROTOCOLEXTENSION_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(RANAPPROTOCOLEXTENSION_IF).PackOT(stream, key)
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
    _size := data.(RANAPPRIVATEIES_IF).GetIECount()
    stream.set_listsize(_size-1, 65535)
    for _, ieID := range self.order {
        item := self.Items[ieID]
        //log.Debug("Packing item %+v\n", item)
        item.Pack(stream, data);
    }
}


type PrivateIEContainer struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PrivateIE-Field', 'actual-parameters': ['IEsSetParam']}, 'size': [(1, 'maxPrivateIEs')], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PRIVATE-IES']}
    Items map[int]*RANAPPRIVATEIES
    order []int
}

type PrivateIEField struct { // [{'type': 'RANAP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}, {'type': 'RANAP-PRIVATE-IES.&criticality', 'table': ['IEsSetParam', ['id']], 'name': 'criticality'}, {'type': 'RANAP-PRIVATE-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}]
    Id RANAPPRIVATEIESid
    Criticality RANAPPRIVATEIEScriticality
    Value RANAPPRIVATEIESValue
}

func (self * PrivateIEField) Unpack(stream *Stream, out interface{}) {
    //table {'type': 'RANAP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    var _len int = 0
    //cobj := RANAPPRIVATEIES{}
    self.Id.Unpack(stream)// p6
    key := self.Id
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out.(RANAPPRIVATEIES_IF).UnpackOT(stream, key)
    //log.Debug("%+v\n",out)
    //{'type': 'RANAP-PRIVATE-IES.&Value', 'table': ['IEsSetParam', ['id']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * RANAPPRIVATEIES) Pack(stream *Stream, out interface{}) {
    //table {'type': 'RANAP-PRIVATE-IES.&id', 'table': {'type': 'IEsSetParam'}, 'name': 'id'}
    key := (RANAPPRIVATEIESid)(self.ID)
    if out.(RANAPPRIVATEIES_IF).GetOT(key) == false {return}
    //fmt.Printf("formating item value %+v\n", self)
    self.ID.Pack(stream)
    self.CRITICALITY.Pack(stream)
    location := stream.reserve_len()
    out.(RANAPPRIVATEIES_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type RABIEContainerList struct{ //{'type': 'ProtocolIE-ContainerList', 'actual-parameters': [1, 'maxNrOfRABs', 'IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES']}
   Item map[int]*RANAPPROTOCOLIES //UserType
   order []int
}
func (self *RABIEContainerList) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-ContainerList', 'actual-parameters': [1, 'maxNrOfRABs', 'IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES']}
    val := ProtocolIEContainerList{self.Item, self.order} //ut2
    val.Unpack(stream, 1, maxNrOfRABs, out) //ut1 args
}

func (self *RABIEContainerList) Pack(stream *Stream, data interface{}){
    // UserType
    val := ProtocolIEContainerList{self.Item, self.order} //ut2
    val.Pack(stream, 1, 256, data)
}

type RABIEContainerPairList struct{ //{'type': 'ProtocolIE-ContainerPairList', 'actual-parameters': [1, 'maxNrOfRABs', 'IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES-PAIR']}
   Item map[int]*RANAPPROTOCOLIESPAIR //UserType
   order []int
}
func (self *RABIEContainerPairList) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-ContainerPairList', 'actual-parameters': [1, 'maxNrOfRABs', 'IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES-PAIR']}
    val := ProtocolIEContainerPairList{self.Item, self.order} //ut2
    val.Unpack(stream, 1, maxNrOfRABs, out) //ut1 args
}

func (self *RABIEContainerPairList) Pack(stream *Stream, data interface{}){
    // UserType
    val := ProtocolIEContainerPairList{self.Item, self.order} //ut2
    val.Pack(stream, 1, 256, data)
}

type ProtocolErrorIEContainerList struct{ //{'type': 'ProtocolIE-ContainerList', 'actual-parameters': [1, 'maxNrOfRABs', 'IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES']}
   Item map[int]*RANAPPROTOCOLIES //UserType
   order []int
}
func (self *ProtocolErrorIEContainerList) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-ContainerList', 'actual-parameters': [1, 'maxNrOfRABs', 'IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES']}
    val := ProtocolIEContainerList{self.Item, self.order} //ut2
    val.Unpack(stream, 1, maxNrOfRABs, out) //ut1 args
}

func (self *ProtocolErrorIEContainerList) Pack(stream *Stream, data interface{}){
    // UserType
    val := ProtocolIEContainerList{self.Item, self.order} //ut2
    val.Pack(stream, 1, 256, data)
}

type IuSigConIdIEContainerList struct{ //{'type': 'ProtocolIE-ContainerList', 'actual-parameters': [1, 'maxNrOfIuSigConIds', 'IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES']}
   Item map[int]*RANAPPROTOCOLIES //UserType
   order []int
}
func (self *IuSigConIdIEContainerList) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-ContainerList', 'actual-parameters': [1, 'maxNrOfIuSigConIds', 'IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES']}
    val := ProtocolIEContainerList{self.Item, self.order} //ut2
    val.Unpack(stream, 1, maxNrOfIuSigConIds, out) //ut1 args
}

func (self *IuSigConIdIEContainerList) Pack(stream *Stream, data interface{}){
    // UserType
    val := ProtocolIEContainerList{self.Item, self.order} //ut2
    val.Pack(stream, 1, 250, data)
}

type DirectTransferIEContainerList struct{ //{'type': 'ProtocolIE-ContainerList', 'actual-parameters': [1, 'maxNrOfDTs', 'IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES']}
   Item map[int]*RANAPPROTOCOLIES //UserType
   order []int
}
func (self *DirectTransferIEContainerList) Unpack(stream *Stream, out interface{}) { //{'type': 'ProtocolIE-ContainerList', 'actual-parameters': [1, 'maxNrOfDTs', 'IEsSetParam'], 'parameters': ['IEsSetParam'], 'param-types': ['RANAP-PROTOCOL-IES']}
    val := ProtocolIEContainerList{self.Item, self.order} //ut2
    val.Unpack(stream, 1, maxNrOfDTs, out) //ut1 args
}

func (self *DirectTransferIEContainerList) Pack(stream *Stream, data interface{}){
    // UserType
    val := ProtocolIEContainerList{self.Item, self.order} //ut2
    val.Pack(stream, 1, 15, data)
}

type IuReleaseCommand struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['Iu-ReleaseCommandIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Iu-ReleaseCommandExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs IuReleaseCommandIEs
    ProtocolExtensions *IuReleaseCommandExtensions
}

func (self * IuReleaseCommand) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_IuReleaseCommandIEs, order_IuReleaseCommandIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &IuReleaseCommandExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Iu-ReleaseCommandExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_IuReleaseCommandExtensions, order_IuReleaseCommandExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * IuReleaseCommand) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_IuReleaseCommandIEs, order_IuReleaseCommandIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_IuReleaseCommandExtensions, order_IuReleaseCommandExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type IuReleaseComplete struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['Iu-ReleaseCompleteIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Iu-ReleaseCompleteExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs IuReleaseCompleteIEs
    ProtocolExtensions *IuReleaseCompleteExtensions
}

func (self * IuReleaseComplete) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_IuReleaseCompleteIEs, order_IuReleaseCompleteIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &IuReleaseCompleteExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Iu-ReleaseCompleteExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_IuReleaseCompleteExtensions, order_IuReleaseCompleteExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * IuReleaseComplete) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_IuReleaseCompleteIEs, order_IuReleaseCompleteIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_IuReleaseCompleteExtensions, order_IuReleaseCompleteExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABDataVolumeReportList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-DataVolumeReportItemIEs']}
    Item RABDataVolumeReportItemIEs //UserType
}
func (self *RABDataVolumeReportList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABDataVolumeReportItemIEs,
       order_RABDataVolumeReportItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABDataVolumeReportList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABDataVolumeReportItemIEs,
        order_RABDataVolumeReportItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABDataVolumeReportItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'DataVolumeList', 'name': 'dl-UnsuccessfullyTransmittedDataVolume', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-DataVolumeReportItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    DlUnsuccessfullyTransmittedDataVolume *DataVolumeList
    IEExtensions *RABDataVolumeReportItemExtIEs
}

func (self * RABDataVolumeReportItem) Unpack(stream *Stream) {
    dlUnsuccessfullyTransmittedDataVolume_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RABID.Unpack(stream)// p8
    if (dlUnsuccessfullyTransmittedDataVolume_flag & _flags) == dlUnsuccessfullyTransmittedDataVolume_flag { //cond2
        self.DlUnsuccessfullyTransmittedDataVolume = &DataVolumeList{}//7{'type': 'DataVolumeList', 'name': 'dl-UnsuccessfullyTransmittedDataVolume', 'optional': True}
        self.DlUnsuccessfullyTransmittedDataVolume.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABDataVolumeReportItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-DataVolumeReportItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABDataVolumeReportItemExtIEs, order_RABDataVolumeReportItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABDataVolumeReportItem) Pack(stream *Stream) {
    const dlUnsuccessfullyTransmittedDataVolume_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.DlUnsuccessfullyTransmittedDataVolume != nil { 
        _flags |= dlUnsuccessfullyTransmittedDataVolume_flag
        self.DlUnsuccessfullyTransmittedDataVolume.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABDataVolumeReportItemExtIEs, order_RABDataVolumeReportItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type RABReleasedListIuRelComp struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-ReleasedItem-IuRelComp-IEs']}
    Item RABReleasedItemIuRelCompIEs //UserType
}
func (self *RABReleasedListIuRelComp) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABReleasedItemIuRelCompIEs,
       order_RABReleasedItemIuRelCompIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABReleasedListIuRelComp) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABReleasedItemIuRelCompIEs,
        order_RABReleasedItemIuRelCompIEs,
    }
        val.Pack(st, self.Item)
}

type RABReleasedItemIuRelComp struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'DL-GTP-PDU-SequenceNumber', 'name': 'dL-GTP-PDU-SequenceNumber', 'optional': True}, {'type': 'UL-GTP-PDU-SequenceNumber', 'name': 'uL-GTP-PDU-SequenceNumber', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ReleasedItem-IuRelComp-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    DLGTPPDUSequenceNumber *DLGTPPDUSequenceNumber
    ULGTPPDUSequenceNumber *ULGTPPDUSequenceNumber
    IEExtensions *RABReleasedItemIuRelCompExtIEs
}

func (self * RABReleasedItemIuRelComp) Unpack(stream *Stream) {
    dLGTPPDUSequenceNumber_flag := 0x00000002
    uLGTPPDUSequenceNumber_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.RABID.Unpack(stream)// p8
    if (dLGTPPDUSequenceNumber_flag & _flags) == dLGTPPDUSequenceNumber_flag { //cond2
        self.DLGTPPDUSequenceNumber = &DLGTPPDUSequenceNumber{}//7{'type': 'DL-GTP-PDU-SequenceNumber', 'name': 'dL-GTP-PDU-SequenceNumber', 'optional': True}
        self.DLGTPPDUSequenceNumber.Unpack(stream)// p8
    }
    if (uLGTPPDUSequenceNumber_flag & _flags) == uLGTPPDUSequenceNumber_flag { //cond2
        self.ULGTPPDUSequenceNumber = &ULGTPPDUSequenceNumber{}//7{'type': 'UL-GTP-PDU-SequenceNumber', 'name': 'uL-GTP-PDU-SequenceNumber', 'optional': True}
        self.ULGTPPDUSequenceNumber.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABReleasedItemIuRelCompExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ReleasedItem-IuRelComp-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABReleasedItemIuRelCompExtIEs, order_RABReleasedItemIuRelCompExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABReleasedItemIuRelComp) Pack(stream *Stream) {
    const dLGTPPDUSequenceNumber_flag uint = 0x00000002
    const uLGTPPDUSequenceNumber_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.DLGTPPDUSequenceNumber != nil { 
        _flags |= dLGTPPDUSequenceNumber_flag
        self.DLGTPPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.ULGTPPDUSequenceNumber != nil { 
        _flags |= uLGTPPDUSequenceNumber_flag
        self.ULGTPPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABReleasedItemIuRelCompExtIEs, order_RABReleasedItemIuRelCompExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type RelocationRequired struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RelocationRequiredIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationRequiredExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RelocationRequiredIEs
    ProtocolExtensions *RelocationRequiredExtensions
}

func (self * RelocationRequired) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RelocationRequiredIEs, order_RelocationRequiredIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RelocationRequiredExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationRequiredExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RelocationRequiredExtensions, order_RelocationRequiredExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RelocationRequired) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RelocationRequiredIEs, order_RelocationRequiredIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RelocationRequiredExtensions, order_RelocationRequiredExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RelocationCommand struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RelocationCommandIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationCommandExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RelocationCommandIEs
    ProtocolExtensions *RelocationCommandExtensions
}

func (self * RelocationCommand) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RelocationCommandIEs, order_RelocationCommandIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RelocationCommandExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationCommandExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RelocationCommandExtensions, order_RelocationCommandExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RelocationCommand) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RelocationCommandIEs, order_RelocationCommandIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RelocationCommandExtensions, order_RelocationCommandExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABRelocationReleaseList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-RelocationReleaseItemIEs']}
    Item RABRelocationReleaseItemIEs //UserType
}
func (self *RABRelocationReleaseList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABRelocationReleaseItemIEs,
       order_RABRelocationReleaseItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABRelocationReleaseList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABRelocationReleaseItemIEs,
        order_RABRelocationReleaseItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABRelocationReleaseItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-RelocationReleaseItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    IEExtensions *RABRelocationReleaseItemExtIEs
}

func (self * RABRelocationReleaseItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABRelocationReleaseItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-RelocationReleaseItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABRelocationReleaseItemExtIEs, order_RABRelocationReleaseItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABRelocationReleaseItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABRelocationReleaseItemExtIEs, order_RABRelocationReleaseItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABDataForwardingList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-DataForwardingItemIEs']}
    Item RABDataForwardingItemIEs //UserType
}
func (self *RABDataForwardingList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABDataForwardingItemIEs,
       order_RABDataForwardingItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABDataForwardingList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABDataForwardingItemIEs,
        order_RABDataForwardingItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABDataForwardingItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'TransportLayerAddress', 'name': 'transportLayerAddress'}, {'type': 'IuTransportAssociation', 'name': 'iuTransportAssociation'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-DataForwardingItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    TransportLayerAddress TransportLayerAddress
    IuTransportAssociation IuTransportAssociation
    IEExtensions *RABDataForwardingItemExtIEs
}

func (self * RABDataForwardingItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    self.TransportLayerAddress.Unpack(stream)// p8
    self.IuTransportAssociation.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABDataForwardingItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-DataForwardingItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABDataForwardingItemExtIEs, order_RABDataForwardingItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABDataForwardingItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    self.TransportLayerAddress.Pack(stream)
    self.IuTransportAssociation.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABDataForwardingItemExtIEs, order_RABDataForwardingItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RelocationPreparationFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RelocationPreparationFailureIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationPreparationFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RelocationPreparationFailureIEs
    ProtocolExtensions *RelocationPreparationFailureExtensions
}

func (self * RelocationPreparationFailure) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RelocationPreparationFailureIEs, order_RelocationPreparationFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RelocationPreparationFailureExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationPreparationFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RelocationPreparationFailureExtensions, order_RelocationPreparationFailureExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RelocationPreparationFailure) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RelocationPreparationFailureIEs, order_RelocationPreparationFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RelocationPreparationFailureExtensions, order_RelocationPreparationFailureExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RelocationRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RelocationRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RelocationRequestIEs
    ProtocolExtensions *RelocationRequestExtensions
}

func (self * RelocationRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RelocationRequestIEs, order_RelocationRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RelocationRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RelocationRequestExtensions, order_RelocationRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RelocationRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RelocationRequestIEs, order_RelocationRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RelocationRequestExtensions, order_RelocationRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABSetupListRelocReq struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-SetupItem-RelocReq-IEs']}
    Item RABSetupItemRelocReqIEs //UserType
}
func (self *RABSetupListRelocReq) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABSetupItemRelocReqIEs,
       order_RABSetupItemRelocReqIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABSetupListRelocReq) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABSetupItemRelocReqIEs,
        order_RABSetupItemRelocReqIEs,
    }
        val.Pack(st, self.Item)
}

type RABSetupItemRelocReq struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'NAS-SynchronisationIndicator', 'name': 'nAS-SynchronisationIndicator', 'optional': True}, {'type': 'RAB-Parameters', 'name': 'rAB-Parameters'}, {'type': 'DataVolumeReportingIndication', 'name': 'dataVolumeReportingIndication', 'optional': True}, {'type': 'PDP-TypeInformation', 'name': 'pDP-TypeInformation', 'optional': True}, {'type': 'UserPlaneInformation', 'name': 'userPlaneInformation'}, {'type': 'TransportLayerAddress', 'name': 'transportLayerAddress'}, {'type': 'IuTransportAssociation', 'name': 'iuTransportAssociation'}, {'type': 'Service-Handover', 'name': 'service-Handover', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-SetupItem-RelocReq-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    NASSynchronisationIndicator *NASSynchronisationIndicator
    RABParameters RABParameters
    DataVolumeReportingIndication *DataVolumeReportingIndication
    PDPTypeInformation *PDPTypeInformation
    UserPlaneInformation UserPlaneInformation
    TransportLayerAddress TransportLayerAddress
    IuTransportAssociation IuTransportAssociation
    ServiceHandover *ServiceHandover
    IEExtensions *RABSetupItemRelocReqExtIEs
}

func (self * RABSetupItemRelocReq) Unpack(stream *Stream) {
    nASSynchronisationIndicator_flag := 0x00000002
    dataVolumeReportingIndication_flag := 0x00000004
    pDPTypeInformation_flag := 0x00000008
    serviceHandover_flag := 0x00000010
    iEExtensions_flag := 0x00000020
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(6)
    self.RABID.Unpack(stream)// p8
    if (nASSynchronisationIndicator_flag & _flags) == nASSynchronisationIndicator_flag { //cond2
        self.NASSynchronisationIndicator = &NASSynchronisationIndicator{}//7{'type': 'NAS-SynchronisationIndicator', 'name': 'nAS-SynchronisationIndicator', 'optional': True}
        self.NASSynchronisationIndicator.Unpack(stream)// p8
    }
    self.RABParameters.Unpack(stream)// p8
    if (dataVolumeReportingIndication_flag & _flags) == dataVolumeReportingIndication_flag { //cond2
        self.DataVolumeReportingIndication = &DataVolumeReportingIndication{}//7{'type': 'DataVolumeReportingIndication', 'name': 'dataVolumeReportingIndication', 'optional': True}
        self.DataVolumeReportingIndication.Unpack(stream)// p8
    }
    if (pDPTypeInformation_flag & _flags) == pDPTypeInformation_flag { //cond2
        self.PDPTypeInformation = &PDPTypeInformation{}//7{'type': 'PDP-TypeInformation', 'name': 'pDP-TypeInformation', 'optional': True}
        self.PDPTypeInformation.Unpack(stream)// p8
    }
    self.UserPlaneInformation.Unpack(stream)// p8
    self.TransportLayerAddress.Unpack(stream)// p8
    self.IuTransportAssociation.Unpack(stream)// p8
    if (serviceHandover_flag & _flags) == serviceHandover_flag { //cond2
        self.ServiceHandover = &ServiceHandover{}//7{'type': 'Service-Handover', 'name': 'service-Handover', 'optional': True}
        self.ServiceHandover.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABSetupItemRelocReqExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-SetupItem-RelocReq-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABSetupItemRelocReqExtIEs, order_RABSetupItemRelocReqExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABSetupItemRelocReq) Pack(stream *Stream) {
    const nASSynchronisationIndicator_flag uint = 0x00000002
    const dataVolumeReportingIndication_flag uint = 0x00000004
    const pDPTypeInformation_flag uint = 0x00000008
    const serviceHandover_flag uint = 0x00000010
    const iEExtensions_flag uint = 0x00000020
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(6)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.NASSynchronisationIndicator != nil { 
        _flags |= nASSynchronisationIndicator_flag
        self.NASSynchronisationIndicator.Pack(stream)
    }//end of optional
    self.RABParameters.Pack(stream)
    if self.DataVolumeReportingIndication != nil { 
        _flags |= dataVolumeReportingIndication_flag
        self.DataVolumeReportingIndication.Pack(stream)
    }//end of optional
    if self.PDPTypeInformation != nil { 
        _flags |= pDPTypeInformation_flag
        self.PDPTypeInformation.Pack(stream)
    }//end of optional
    self.UserPlaneInformation.Pack(stream)
    self.TransportLayerAddress.Pack(stream)
    self.IuTransportAssociation.Pack(stream)
    if self.ServiceHandover != nil { 
        _flags |= serviceHandover_flag
        self.ServiceHandover.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABSetupItemRelocReqExtIEs, order_RABSetupItemRelocReqExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

type UserPlaneInformation struct { // [{'type': 'UserPlaneMode', 'name': 'userPlaneMode'}, {'type': 'UP-ModeVersions', 'name': 'uP-ModeVersions'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UserPlaneInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    UserPlaneMode UserPlaneMode
    UPModeVersions UPModeVersions
    IEExtensions *UserPlaneInformationExtIEs
}

func (self * UserPlaneInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.UserPlaneMode.Unpack(stream)// p8
    self.UPModeVersions.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &UserPlaneInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UserPlaneInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_UserPlaneInformationExtIEs, order_UserPlaneInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UserPlaneInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UserPlaneMode.Pack(stream)
    self.UPModeVersions.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_UserPlaneInformationExtIEs, order_UserPlaneInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RelocationRequestAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RelocationRequestAcknowledgeIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationRequestAcknowledgeExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RelocationRequestAcknowledgeIEs
    ProtocolExtensions *RelocationRequestAcknowledgeExtensions
}

func (self * RelocationRequestAcknowledge) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RelocationRequestAcknowledgeIEs, order_RelocationRequestAcknowledgeIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RelocationRequestAcknowledgeExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationRequestAcknowledgeExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RelocationRequestAcknowledgeExtensions, order_RelocationRequestAcknowledgeExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RelocationRequestAcknowledge) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RelocationRequestAcknowledgeIEs, order_RelocationRequestAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RelocationRequestAcknowledgeExtensions, order_RelocationRequestAcknowledgeExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABSetupListRelocReqAck struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-SetupItem-RelocReqAck-IEs']}
    Item RABSetupItemRelocReqAckIEs //UserType
}
func (self *RABSetupListRelocReqAck) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABSetupItemRelocReqAckIEs,
       order_RABSetupItemRelocReqAckIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABSetupListRelocReqAck) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABSetupItemRelocReqAckIEs,
        order_RABSetupItemRelocReqAckIEs,
    }
        val.Pack(st, self.Item)
}

type RABSetupItemRelocReqAck struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'TransportLayerAddress', 'name': 'transportLayerAddress', 'optional': True}, {'type': 'IuTransportAssociation', 'name': 'iuTransportAssociation', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-SetupItem-RelocReqAck-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    TransportLayerAddress *TransportLayerAddress
    IuTransportAssociation *IuTransportAssociation
    IEExtensions *RABSetupItemRelocReqAckExtIEs
}

func (self * RABSetupItemRelocReqAck) Unpack(stream *Stream) {
    transportLayerAddress_flag := 0x00000002
    iuTransportAssociation_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.RABID.Unpack(stream)// p8
    if (transportLayerAddress_flag & _flags) == transportLayerAddress_flag { //cond2
        self.TransportLayerAddress = &TransportLayerAddress{}//7{'type': 'TransportLayerAddress', 'name': 'transportLayerAddress', 'optional': True}
        self.TransportLayerAddress.Unpack(stream)// p8
    }
    if (iuTransportAssociation_flag & _flags) == iuTransportAssociation_flag { //cond2
        self.IuTransportAssociation = &IuTransportAssociation{}//7{'type': 'IuTransportAssociation', 'name': 'iuTransportAssociation', 'optional': True}
        self.IuTransportAssociation.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABSetupItemRelocReqAckExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-SetupItem-RelocReqAck-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABSetupItemRelocReqAckExtIEs, order_RABSetupItemRelocReqAckExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABSetupItemRelocReqAck) Pack(stream *Stream) {
    const transportLayerAddress_flag uint = 0x00000002
    const iuTransportAssociation_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.TransportLayerAddress != nil { 
        _flags |= transportLayerAddress_flag
        self.TransportLayerAddress.Pack(stream)
    }//end of optional
    if self.IuTransportAssociation != nil { 
        _flags |= iuTransportAssociation_flag
        self.IuTransportAssociation.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABSetupItemRelocReqAckExtIEs, order_RABSetupItemRelocReqAckExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type RABFailedList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-FailedItemIEs']}
    Item RABFailedItemIEs //UserType
}
func (self *RABFailedList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABFailedItemIEs,
       order_RABFailedItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABFailedList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABFailedItemIEs,
        order_RABFailedItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABFailedItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-FailedItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    Cause Cause
    IEExtensions *RABFailedItemExtIEs
}

func (self * RABFailedItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABFailedItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-FailedItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABFailedItemExtIEs, order_RABFailedItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABFailedItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABFailedItemExtIEs, order_RABFailedItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RelocationFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RelocationFailureIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RelocationFailureIEs
    ProtocolExtensions *RelocationFailureExtensions
}

func (self * RelocationFailure) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RelocationFailureIEs, order_RelocationFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RelocationFailureExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RelocationFailureExtensions, order_RelocationFailureExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RelocationFailure) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RelocationFailureIEs, order_RelocationFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RelocationFailureExtensions, order_RelocationFailureExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RelocationCancel struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RelocationCancelIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationCancelExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RelocationCancelIEs
    ProtocolExtensions *RelocationCancelExtensions
}

func (self * RelocationCancel) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RelocationCancelIEs, order_RelocationCancelIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RelocationCancelExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationCancelExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RelocationCancelExtensions, order_RelocationCancelExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RelocationCancel) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RelocationCancelIEs, order_RelocationCancelIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RelocationCancelExtensions, order_RelocationCancelExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RelocationCancelAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RelocationCancelAcknowledgeIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationCancelAcknowledgeExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RelocationCancelAcknowledgeIEs
    ProtocolExtensions *RelocationCancelAcknowledgeExtensions
}

func (self * RelocationCancelAcknowledge) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RelocationCancelAcknowledgeIEs, order_RelocationCancelAcknowledgeIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RelocationCancelAcknowledgeExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationCancelAcknowledgeExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RelocationCancelAcknowledgeExtensions, order_RelocationCancelAcknowledgeExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RelocationCancelAcknowledge) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RelocationCancelAcknowledgeIEs, order_RelocationCancelAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RelocationCancelAcknowledgeExtensions, order_RelocationCancelAcknowledgeExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SRNSContextRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SRNS-ContextRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SRNS-ContextRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs SRNSContextRequestIEs
    ProtocolExtensions *SRNSContextRequestExtensions
}

func (self * SRNSContextRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_SRNSContextRequestIEs, order_SRNSContextRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &SRNSContextRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SRNS-ContextRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_SRNSContextRequestExtensions, order_SRNSContextRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SRNSContextRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SRNSContextRequestIEs, order_SRNSContextRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_SRNSContextRequestExtensions, order_SRNSContextRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABDataForwardingListSRNSCtxReq struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-DataForwardingItem-SRNS-CtxReq-IEs']}
    Item RABDataForwardingItemSRNSCtxReqIEs //UserType
}
func (self *RABDataForwardingListSRNSCtxReq) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABDataForwardingItemSRNSCtxReqIEs,
       order_RABDataForwardingItemSRNSCtxReqIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABDataForwardingListSRNSCtxReq) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABDataForwardingItemSRNSCtxReqIEs,
        order_RABDataForwardingItemSRNSCtxReqIEs,
    }
        val.Pack(st, self.Item)
}

type RABDataForwardingItemSRNSCtxReq struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-DataForwardingItem-SRNS-CtxReq-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    IEExtensions *RABDataForwardingItemSRNSCtxReqExtIEs
}

func (self * RABDataForwardingItemSRNSCtxReq) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABDataForwardingItemSRNSCtxReqExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-DataForwardingItem-SRNS-CtxReq-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABDataForwardingItemSRNSCtxReqExtIEs, order_RABDataForwardingItemSRNSCtxReqExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABDataForwardingItemSRNSCtxReq) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABDataForwardingItemSRNSCtxReqExtIEs, order_RABDataForwardingItemSRNSCtxReqExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SRNSContextResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SRNS-ContextResponseIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SRNS-ContextResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs SRNSContextResponseIEs
    ProtocolExtensions *SRNSContextResponseExtensions
}

func (self * SRNSContextResponse) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_SRNSContextResponseIEs, order_SRNSContextResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &SRNSContextResponseExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SRNS-ContextResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_SRNSContextResponseExtensions, order_SRNSContextResponseExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SRNSContextResponse) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SRNSContextResponseIEs, order_SRNSContextResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_SRNSContextResponseExtensions, order_SRNSContextResponseExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABContextList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-ContextItemIEs']}
    Item RABContextItemIEs //UserType
}
func (self *RABContextList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABContextItemIEs,
       order_RABContextItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABContextList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABContextItemIEs,
        order_RABContextItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABContextItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'DL-GTP-PDU-SequenceNumber', 'name': 'dl-GTP-PDU-SequenceNumber', 'optional': True}, {'type': 'UL-GTP-PDU-SequenceNumber', 'name': 'ul-GTP-PDU-SequenceNumber', 'optional': True}, {'type': 'DL-N-PDU-SequenceNumber', 'name': 'dl-N-PDU-SequenceNumber', 'optional': True}, {'type': 'UL-N-PDU-SequenceNumber', 'name': 'ul-N-PDU-SequenceNumber', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ContextItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    DlGTPPDUSequenceNumber *DLGTPPDUSequenceNumber
    UlGTPPDUSequenceNumber *ULGTPPDUSequenceNumber
    DlNPDUSequenceNumber *DLNPDUSequenceNumber
    UlNPDUSequenceNumber *ULNPDUSequenceNumber
    IEExtensions *RABContextItemExtIEs
}

func (self * RABContextItem) Unpack(stream *Stream) {
    dlGTPPDUSequenceNumber_flag := 0x00000002
    ulGTPPDUSequenceNumber_flag := 0x00000004
    dlNPDUSequenceNumber_flag := 0x00000008
    ulNPDUSequenceNumber_flag := 0x00000010
    iEExtensions_flag := 0x00000020
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(6)
    self.RABID.Unpack(stream)// p8
    if (dlGTPPDUSequenceNumber_flag & _flags) == dlGTPPDUSequenceNumber_flag { //cond2
        self.DlGTPPDUSequenceNumber = &DLGTPPDUSequenceNumber{}//7{'type': 'DL-GTP-PDU-SequenceNumber', 'name': 'dl-GTP-PDU-SequenceNumber', 'optional': True}
        self.DlGTPPDUSequenceNumber.Unpack(stream)// p8
    }
    if (ulGTPPDUSequenceNumber_flag & _flags) == ulGTPPDUSequenceNumber_flag { //cond2
        self.UlGTPPDUSequenceNumber = &ULGTPPDUSequenceNumber{}//7{'type': 'UL-GTP-PDU-SequenceNumber', 'name': 'ul-GTP-PDU-SequenceNumber', 'optional': True}
        self.UlGTPPDUSequenceNumber.Unpack(stream)// p8
    }
    if (dlNPDUSequenceNumber_flag & _flags) == dlNPDUSequenceNumber_flag { //cond2
        self.DlNPDUSequenceNumber = &DLNPDUSequenceNumber{}//7{'type': 'DL-N-PDU-SequenceNumber', 'name': 'dl-N-PDU-SequenceNumber', 'optional': True}
        self.DlNPDUSequenceNumber.Unpack(stream)// p8
    }
    if (ulNPDUSequenceNumber_flag & _flags) == ulNPDUSequenceNumber_flag { //cond2
        self.UlNPDUSequenceNumber = &ULNPDUSequenceNumber{}//7{'type': 'UL-N-PDU-SequenceNumber', 'name': 'ul-N-PDU-SequenceNumber', 'optional': True}
        self.UlNPDUSequenceNumber.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABContextItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ContextItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABContextItemExtIEs, order_RABContextItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABContextItem) Pack(stream *Stream) {
    const dlGTPPDUSequenceNumber_flag uint = 0x00000002
    const ulGTPPDUSequenceNumber_flag uint = 0x00000004
    const dlNPDUSequenceNumber_flag uint = 0x00000008
    const ulNPDUSequenceNumber_flag uint = 0x00000010
    const iEExtensions_flag uint = 0x00000020
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(6)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.DlGTPPDUSequenceNumber != nil { 
        _flags |= dlGTPPDUSequenceNumber_flag
        self.DlGTPPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.UlGTPPDUSequenceNumber != nil { 
        _flags |= ulGTPPDUSequenceNumber_flag
        self.UlGTPPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.DlNPDUSequenceNumber != nil { 
        _flags |= dlNPDUSequenceNumber_flag
        self.DlNPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.UlNPDUSequenceNumber != nil { 
        _flags |= ulNPDUSequenceNumber_flag
        self.UlNPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABContextItemExtIEs, order_RABContextItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

type RABContextFailedtoTransferList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RABs-ContextFailedtoTransferItemIEs']}
    Item RABsContextFailedtoTransferItemIEs //UserType
}
func (self *RABContextFailedtoTransferList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABsContextFailedtoTransferItemIEs,
       order_RABsContextFailedtoTransferItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABContextFailedtoTransferList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABsContextFailedtoTransferItemIEs,
        order_RABsContextFailedtoTransferItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABsContextFailedtoTransferItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RABs-ContextFailedtoTransferItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    Cause Cause
    IEExtensions *RABsContextFailedtoTransferItemExtIEs
}

func (self * RABsContextFailedtoTransferItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABsContextFailedtoTransferItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RABs-ContextFailedtoTransferItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABsContextFailedtoTransferItemExtIEs, order_RABsContextFailedtoTransferItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABsContextFailedtoTransferItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABsContextFailedtoTransferItemExtIEs, order_RABsContextFailedtoTransferItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SecurityModeCommand struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SecurityModeCommandIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityModeCommandExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs SecurityModeCommandIEs
    ProtocolExtensions *SecurityModeCommandExtensions
}

func (self * SecurityModeCommand) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_SecurityModeCommandIEs, order_SecurityModeCommandIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &SecurityModeCommandExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityModeCommandExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_SecurityModeCommandExtensions, order_SecurityModeCommandExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SecurityModeCommand) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SecurityModeCommandIEs, order_SecurityModeCommandIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_SecurityModeCommandExtensions, order_SecurityModeCommandExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SecurityModeComplete struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SecurityModeCompleteIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityModeCompleteExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs SecurityModeCompleteIEs
    ProtocolExtensions *SecurityModeCompleteExtensions
}

func (self * SecurityModeComplete) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_SecurityModeCompleteIEs, order_SecurityModeCompleteIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &SecurityModeCompleteExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityModeCompleteExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_SecurityModeCompleteExtensions, order_SecurityModeCompleteExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SecurityModeComplete) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SecurityModeCompleteIEs, order_SecurityModeCompleteIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_SecurityModeCompleteExtensions, order_SecurityModeCompleteExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SecurityModeReject struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SecurityModeRejectIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityModeRejectExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs SecurityModeRejectIEs
    ProtocolExtensions *SecurityModeRejectExtensions
}

func (self * SecurityModeReject) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_SecurityModeRejectIEs, order_SecurityModeRejectIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &SecurityModeRejectExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SecurityModeRejectExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_SecurityModeRejectExtensions, order_SecurityModeRejectExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SecurityModeReject) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SecurityModeRejectIEs, order_SecurityModeRejectIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_SecurityModeRejectExtensions, order_SecurityModeRejectExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type DataVolumeReportRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['DataVolumeReportRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DataVolumeReportRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs DataVolumeReportRequestIEs
    ProtocolExtensions *DataVolumeReportRequestExtensions
}

func (self * DataVolumeReportRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_DataVolumeReportRequestIEs, order_DataVolumeReportRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &DataVolumeReportRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DataVolumeReportRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_DataVolumeReportRequestExtensions, order_DataVolumeReportRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DataVolumeReportRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_DataVolumeReportRequestIEs, order_DataVolumeReportRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_DataVolumeReportRequestExtensions, order_DataVolumeReportRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABDataVolumeReportRequestList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-DataVolumeReportRequestItemIEs']}
    Item RABDataVolumeReportRequestItemIEs //UserType
}
func (self *RABDataVolumeReportRequestList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABDataVolumeReportRequestItemIEs,
       order_RABDataVolumeReportRequestItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABDataVolumeReportRequestList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABDataVolumeReportRequestItemIEs,
        order_RABDataVolumeReportRequestItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABDataVolumeReportRequestItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-DataVolumeReportRequestItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    IEExtensions *RABDataVolumeReportRequestItemExtIEs
}

func (self * RABDataVolumeReportRequestItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABDataVolumeReportRequestItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-DataVolumeReportRequestItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABDataVolumeReportRequestItemExtIEs, order_RABDataVolumeReportRequestItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABDataVolumeReportRequestItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABDataVolumeReportRequestItemExtIEs, order_RABDataVolumeReportRequestItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type DataVolumeReport struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['DataVolumeReportIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DataVolumeReportExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs DataVolumeReportIEs
    ProtocolExtensions *DataVolumeReportExtensions
}

func (self * DataVolumeReport) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_DataVolumeReportIEs, order_DataVolumeReportIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &DataVolumeReportExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DataVolumeReportExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_DataVolumeReportExtensions, order_DataVolumeReportExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DataVolumeReport) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_DataVolumeReportIEs, order_DataVolumeReportIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_DataVolumeReportExtensions, order_DataVolumeReportExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABFailedtoReportList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RABs-failed-to-reportItemIEs']}
    Item RABsfailedtoreportItemIEs //UserType
}
func (self *RABFailedtoReportList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABsfailedtoreportItemIEs,
       order_RABsfailedtoreportItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABFailedtoReportList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABsfailedtoreportItemIEs,
        order_RABsfailedtoreportItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABsfailedtoreportItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RABs-failed-to-reportItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    Cause Cause
    IEExtensions *RABsfailedtoreportItemExtIEs
}

func (self * RABsfailedtoreportItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABsfailedtoreportItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RABs-failed-to-reportItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABsfailedtoreportItemExtIEs, order_RABsfailedtoreportItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABsfailedtoreportItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABsfailedtoreportItemExtIEs, order_RABsfailedtoreportItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type Reset struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ResetIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs ResetIEs
    ProtocolExtensions *ResetExtensions
}

func (self * Reset) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_ResetIEs, order_ResetIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &ResetExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_ResetExtensions, order_ResetExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * Reset) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ResetIEs, order_ResetIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_ResetExtensions, order_ResetExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ResetAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ResetAcknowledgeIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetAcknowledgeExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs ResetAcknowledgeIEs
    ProtocolExtensions *ResetAcknowledgeExtensions
}

func (self * ResetAcknowledge) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_ResetAcknowledgeIEs, order_ResetAcknowledgeIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &ResetAcknowledgeExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetAcknowledgeExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_ResetAcknowledgeExtensions, order_ResetAcknowledgeExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ResetAcknowledge) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ResetAcknowledgeIEs, order_ResetAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_ResetAcknowledgeExtensions, order_ResetAcknowledgeExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ResetResource struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ResetResourceIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetResourceExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs ResetResourceIEs
    ProtocolExtensions *ResetResourceExtensions
}

func (self * ResetResource) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_ResetResourceIEs, order_ResetResourceIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &ResetResourceExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetResourceExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_ResetResourceExtensions, order_ResetResourceExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ResetResource) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ResetResourceIEs, order_ResetResourceIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_ResetResourceExtensions, order_ResetResourceExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ResetResourceList struct{ //{'type': 'IuSigConId-IE-ContainerList', 'actual-parameters': ['ResetResourceItemIEs']}
    Item ResetResourceItemIEs //UserType
}
func (self *ResetResourceList) Unpack(st *Stream){ // ut3 'IuSigConIdIEContainerList'
    val := IuSigConIdIEContainerList{// ut31
       table_ResetResourceItemIEs,
       order_ResetResourceItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *ResetResourceList) Pack(st *Stream){
    val := IuSigConIdIEContainerList{// ut31
        table_ResetResourceItemIEs,
        order_ResetResourceItemIEs,
    }
        val.Pack(st, self.Item)
}

type ResetResourceItem struct { // [{'type': 'IuSignallingConnectionIdentifier', 'name': 'iuSigConId'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetResourceItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    IuSigConId IuSignallingConnectionIdentifier
    IEExtensions *ResetResourceItemExtIEs
}

func (self * ResetResourceItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.IuSigConId.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &ResetResourceItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetResourceItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_ResetResourceItemExtIEs, order_ResetResourceItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ResetResourceItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IuSigConId.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_ResetResourceItemExtIEs, order_ResetResourceItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ResetResourceAcknowledge struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ResetResourceAcknowledgeIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetResourceAcknowledgeExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs ResetResourceAcknowledgeIEs
    ProtocolExtensions *ResetResourceAcknowledgeExtensions
}

func (self * ResetResourceAcknowledge) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_ResetResourceAcknowledgeIEs, order_ResetResourceAcknowledgeIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &ResetResourceAcknowledgeExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetResourceAcknowledgeExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_ResetResourceAcknowledgeExtensions, order_ResetResourceAcknowledgeExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ResetResourceAcknowledge) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ResetResourceAcknowledgeIEs, order_ResetResourceAcknowledgeIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_ResetResourceAcknowledgeExtensions, order_ResetResourceAcknowledgeExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ResetResourceAckList struct{ //{'type': 'IuSigConId-IE-ContainerList', 'actual-parameters': ['ResetResourceAckItemIEs']}
    Item ResetResourceAckItemIEs //UserType
}
func (self *ResetResourceAckList) Unpack(st *Stream){ // ut3 'IuSigConIdIEContainerList'
    val := IuSigConIdIEContainerList{// ut31
       table_ResetResourceAckItemIEs,
       order_ResetResourceAckItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *ResetResourceAckList) Pack(st *Stream){
    val := IuSigConIdIEContainerList{// ut31
        table_ResetResourceAckItemIEs,
        order_ResetResourceAckItemIEs,
    }
        val.Pack(st, self.Item)
}

type ResetResourceAckItem struct { // [{'type': 'IuSignallingConnectionIdentifier', 'name': 'iuSigConId'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetResourceAckItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    IuSigConId IuSignallingConnectionIdentifier
    IEExtensions *ResetResourceAckItemExtIEs
}

func (self * ResetResourceAckItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.IuSigConId.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &ResetResourceAckItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResetResourceAckItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_ResetResourceAckItemExtIEs, order_ResetResourceAckItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ResetResourceAckItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IuSigConId.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_ResetResourceAckItemExtIEs, order_ResetResourceAckItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABReleaseRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RAB-ReleaseRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ReleaseRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RABReleaseRequestIEs
    ProtocolExtensions *RABReleaseRequestExtensions
}

func (self * RABReleaseRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RABReleaseRequestIEs, order_RABReleaseRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RABReleaseRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ReleaseRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RABReleaseRequestExtensions, order_RABReleaseRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABReleaseRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RABReleaseRequestIEs, order_RABReleaseRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RABReleaseRequestExtensions, order_RABReleaseRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABReleaseList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-ReleaseItemIEs']}
    Item RABReleaseItemIEs //UserType
}
func (self *RABReleaseList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABReleaseItemIEs,
       order_RABReleaseItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABReleaseList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABReleaseItemIEs,
        order_RABReleaseItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABReleaseItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ReleaseItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    Cause Cause
    IEExtensions *RABReleaseItemExtIEs
}

func (self * RABReleaseItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABReleaseItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ReleaseItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABReleaseItemExtIEs, order_RABReleaseItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABReleaseItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    self.Cause.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABReleaseItemExtIEs, order_RABReleaseItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type IuReleaseRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['Iu-ReleaseRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Iu-ReleaseRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs IuReleaseRequestIEs
    ProtocolExtensions *IuReleaseRequestExtensions
}

func (self * IuReleaseRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_IuReleaseRequestIEs, order_IuReleaseRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &IuReleaseRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Iu-ReleaseRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_IuReleaseRequestExtensions, order_IuReleaseRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * IuReleaseRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_IuReleaseRequestIEs, order_IuReleaseRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_IuReleaseRequestExtensions, order_IuReleaseRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RelocationDetect struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RelocationDetectIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationDetectExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RelocationDetectIEs
    ProtocolExtensions *RelocationDetectExtensions
}

func (self * RelocationDetect) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RelocationDetectIEs, order_RelocationDetectIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RelocationDetectExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RelocationDetectExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RelocationDetectExtensions, order_RelocationDetectExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RelocationDetect) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RelocationDetectIEs, order_RelocationDetectIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RelocationDetectExtensions, order_RelocationDetectExtensions} // p3
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

type Paging struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['PagingIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PagingExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs PagingIEs
    ProtocolExtensions *PagingExtensions
}

func (self * Paging) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_PagingIEs, order_PagingIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &PagingExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PagingExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_PagingExtensions, order_PagingExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * Paging) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_PagingIEs, order_PagingIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_PagingExtensions, order_PagingExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type CommonID struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['CommonID-IEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CommonIDExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs CommonIDIEs
    ProtocolExtensions *CommonIDExtensions
}

func (self * CommonID) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_CommonIDIEs, order_CommonIDIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &CommonIDExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CommonIDExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_CommonIDExtensions, order_CommonIDExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CommonID) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_CommonIDIEs, order_CommonIDIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_CommonIDExtensions, order_CommonIDExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type CNInvokeTrace struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['CN-InvokeTraceIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CN-InvokeTraceExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs CNInvokeTraceIEs
    ProtocolExtensions *CNInvokeTraceExtensions
}

func (self * CNInvokeTrace) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_CNInvokeTraceIEs, order_CNInvokeTraceIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &CNInvokeTraceExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CN-InvokeTraceExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_CNInvokeTraceExtensions, order_CNInvokeTraceExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CNInvokeTrace) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_CNInvokeTraceIEs, order_CNInvokeTraceIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_CNInvokeTraceExtensions, order_CNInvokeTraceExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type CNDeactivateTrace struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['CN-DeactivateTraceIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CN-DeactivateTraceExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs CNDeactivateTraceIEs
    ProtocolExtensions *CNDeactivateTraceExtensions
}

func (self * CNDeactivateTrace) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_CNDeactivateTraceIEs, order_CNDeactivateTraceIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &CNDeactivateTraceExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CN-DeactivateTraceExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_CNDeactivateTraceExtensions, order_CNDeactivateTraceExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CNDeactivateTrace) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_CNDeactivateTraceIEs, order_CNDeactivateTraceIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_CNDeactivateTraceExtensions, order_CNDeactivateTraceExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type LocationReportingControl struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['LocationReportingControlIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LocationReportingControlExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs LocationReportingControlIEs
    ProtocolExtensions *LocationReportingControlExtensions
}

func (self * LocationReportingControl) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_LocationReportingControlIEs, order_LocationReportingControlIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &LocationReportingControlExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LocationReportingControlExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_LocationReportingControlExtensions, order_LocationReportingControlExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * LocationReportingControl) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_LocationReportingControlIEs, order_LocationReportingControlIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_LocationReportingControlExtensions, order_LocationReportingControlExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type LocationReport struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['LocationReportIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LocationReportExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs LocationReportIEs
    ProtocolExtensions *LocationReportExtensions
}

func (self * LocationReport) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_LocationReportIEs, order_LocationReportIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &LocationReportExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LocationReportExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_LocationReportExtensions, order_LocationReportExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * LocationReport) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_LocationReportIEs, order_LocationReportIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_LocationReportExtensions, order_LocationReportExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type InitialUEMessage struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['InitialUE-MessageIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InitialUE-MessageExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs InitialUEMessageIEs
    ProtocolExtensions *InitialUEMessageExtensions
}

func (self * InitialUEMessage) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_InitialUEMessageIEs, order_InitialUEMessageIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &InitialUEMessageExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InitialUE-MessageExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_InitialUEMessageExtensions, order_InitialUEMessageExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InitialUEMessage) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_InitialUEMessageIEs, order_InitialUEMessageIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_InitialUEMessageExtensions, order_InitialUEMessageExtensions} // p3
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

type Overload struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['OverloadIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['OverloadExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs OverloadIEs
    ProtocolExtensions *OverloadExtensions
}

func (self * Overload) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_OverloadIEs, order_OverloadIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &OverloadExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['OverloadExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_OverloadExtensions, order_OverloadExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * Overload) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_OverloadIEs, order_OverloadIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_OverloadExtensions, order_OverloadExtensions} // p3
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

type SRNSDataForwardCommand struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['SRNS-DataForwardCommandIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SRNS-DataForwardCommandExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs SRNSDataForwardCommandIEs
    ProtocolExtensions *SRNSDataForwardCommandExtensions
}

func (self * SRNSDataForwardCommand) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_SRNSDataForwardCommandIEs, order_SRNSDataForwardCommandIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &SRNSDataForwardCommandExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SRNS-DataForwardCommandExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_SRNSDataForwardCommandExtensions, order_SRNSDataForwardCommandExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SRNSDataForwardCommand) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_SRNSDataForwardCommandIEs, order_SRNSDataForwardCommandIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_SRNSDataForwardCommandExtensions, order_SRNSDataForwardCommandExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ForwardSRNSContext struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['ForwardSRNS-ContextIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ForwardSRNS-ContextExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs ForwardSRNSContextIEs
    ProtocolExtensions *ForwardSRNSContextExtensions
}

func (self * ForwardSRNSContext) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_ForwardSRNSContextIEs, order_ForwardSRNSContextIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &ForwardSRNSContextExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ForwardSRNS-ContextExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_ForwardSRNSContextExtensions, order_ForwardSRNSContextExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ForwardSRNSContext) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_ForwardSRNSContextIEs, order_ForwardSRNSContextIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_ForwardSRNSContextExtensions, order_ForwardSRNSContextExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABAssignmentRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RAB-AssignmentRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-AssignmentRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RABAssignmentRequestIEs
    ProtocolExtensions *RABAssignmentRequestExtensions
}

func (self * RABAssignmentRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RABAssignmentRequestIEs, order_RABAssignmentRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RABAssignmentRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-AssignmentRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RABAssignmentRequestExtensions, order_RABAssignmentRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABAssignmentRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RABAssignmentRequestIEs, order_RABAssignmentRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RABAssignmentRequestExtensions, order_RABAssignmentRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABSetupOrModifyList struct{ //{'type': 'RAB-IE-ContainerPairList', 'actual-parameters': ['RAB-SetupOrModifyItem-IEs']}
    Item RABSetupOrModifyItemIEs //UserType
}
func (self *RABSetupOrModifyList) Unpack(st *Stream){ // ut3 'RABIEContainerPairList'
    val := RABIEContainerPairList{// ut31
       table_RABSetupOrModifyItemIEs,
       order_RABSetupOrModifyItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABSetupOrModifyList) Pack(st *Stream){
    val := RABIEContainerPairList{// ut31
        table_RABSetupOrModifyItemIEs,
        order_RABSetupOrModifyItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABSetupOrModifyItemFirst struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'NAS-SynchronisationIndicator', 'name': 'nAS-SynchronisationIndicator', 'optional': True}, {'type': 'RAB-Parameters', 'name': 'rAB-Parameters', 'optional': True}, {'type': 'UserPlaneInformation', 'name': 'userPlaneInformation', 'optional': True}, {'type': 'TransportLayerInformation', 'name': 'transportLayerInformation', 'optional': True}, {'type': 'Service-Handover', 'name': 'service-Handover', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-SetupOrModifyItemFirst-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    NASSynchronisationIndicator *NASSynchronisationIndicator
    RABParameters *RABParameters
    UserPlaneInformation *UserPlaneInformation
    TransportLayerInformation *TransportLayerInformation
    ServiceHandover *ServiceHandover
    IEExtensions *RABSetupOrModifyItemFirstExtIEs
}

func (self * RABSetupOrModifyItemFirst) Unpack(stream *Stream) {
    nASSynchronisationIndicator_flag := 0x00000002
    rABParameters_flag := 0x00000004
    userPlaneInformation_flag := 0x00000008
    transportLayerInformation_flag := 0x00000010
    serviceHandover_flag := 0x00000020
    iEExtensions_flag := 0x00000040
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(7)
    self.RABID.Unpack(stream)// p8
    if (nASSynchronisationIndicator_flag & _flags) == nASSynchronisationIndicator_flag { //cond2
        self.NASSynchronisationIndicator = &NASSynchronisationIndicator{}//7{'type': 'NAS-SynchronisationIndicator', 'name': 'nAS-SynchronisationIndicator', 'optional': True}
        self.NASSynchronisationIndicator.Unpack(stream)// p8
    }
    if (rABParameters_flag & _flags) == rABParameters_flag { //cond2
        self.RABParameters = &RABParameters{}//7{'type': 'RAB-Parameters', 'name': 'rAB-Parameters', 'optional': True}
        self.RABParameters.Unpack(stream)// p8
    }
    if (userPlaneInformation_flag & _flags) == userPlaneInformation_flag { //cond2
        self.UserPlaneInformation = &UserPlaneInformation{}//7{'type': 'UserPlaneInformation', 'name': 'userPlaneInformation', 'optional': True}
        self.UserPlaneInformation.Unpack(stream)// p8
    }
    if (transportLayerInformation_flag & _flags) == transportLayerInformation_flag { //cond2
        self.TransportLayerInformation = &TransportLayerInformation{}//7{'type': 'TransportLayerInformation', 'name': 'transportLayerInformation', 'optional': True}
        self.TransportLayerInformation.Unpack(stream)// p8
    }
    if (serviceHandover_flag & _flags) == serviceHandover_flag { //cond2
        self.ServiceHandover = &ServiceHandover{}//7{'type': 'Service-Handover', 'name': 'service-Handover', 'optional': True}
        self.ServiceHandover.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABSetupOrModifyItemFirstExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-SetupOrModifyItemFirst-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABSetupOrModifyItemFirstExtIEs, order_RABSetupOrModifyItemFirstExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABSetupOrModifyItemFirst) Pack(stream *Stream) {
    const nASSynchronisationIndicator_flag uint = 0x00000002
    const rABParameters_flag uint = 0x00000004
    const userPlaneInformation_flag uint = 0x00000008
    const transportLayerInformation_flag uint = 0x00000010
    const serviceHandover_flag uint = 0x00000020
    const iEExtensions_flag uint = 0x00000040
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(7)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.NASSynchronisationIndicator != nil { 
        _flags |= nASSynchronisationIndicator_flag
        self.NASSynchronisationIndicator.Pack(stream)
    }//end of optional
    if self.RABParameters != nil { 
        _flags |= rABParameters_flag
        self.RABParameters.Pack(stream)
    }//end of optional
    if self.UserPlaneInformation != nil { 
        _flags |= userPlaneInformation_flag
        self.UserPlaneInformation.Pack(stream)
    }//end of optional
    if self.TransportLayerInformation != nil { 
        _flags |= transportLayerInformation_flag
        self.TransportLayerInformation.Pack(stream)
    }//end of optional
    if self.ServiceHandover != nil { 
        _flags |= serviceHandover_flag
        self.ServiceHandover.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABSetupOrModifyItemFirstExtIEs, order_RABSetupOrModifyItemFirstExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 7)
}//end

type TransportLayerInformation struct { // [{'type': 'TransportLayerAddress', 'name': 'transportLayerAddress'}, {'type': 'IuTransportAssociation', 'name': 'iuTransportAssociation'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TransportLayerInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    TransportLayerAddress TransportLayerAddress
    IuTransportAssociation IuTransportAssociation
    IEExtensions *TransportLayerInformationExtIEs
}

func (self * TransportLayerInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.TransportLayerAddress.Unpack(stream)// p8
    self.IuTransportAssociation.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &TransportLayerInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TransportLayerInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_TransportLayerInformationExtIEs, order_TransportLayerInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TransportLayerInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TransportLayerAddress.Pack(stream)
    self.IuTransportAssociation.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_TransportLayerInformationExtIEs, order_TransportLayerInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABSetupOrModifyItemSecond struct { // [{'type': 'PDP-TypeInformation', 'name': 'pDP-TypeInformation', 'optional': True}, {'type': 'DataVolumeReportingIndication', 'name': 'dataVolumeReportingIndication', 'optional': True}, {'type': 'DL-GTP-PDU-SequenceNumber', 'name': 'dl-GTP-PDU-SequenceNumber', 'optional': True}, {'type': 'UL-GTP-PDU-SequenceNumber', 'name': 'ul-GTP-PDU-SequenceNumber', 'optional': True}, {'type': 'DL-N-PDU-SequenceNumber', 'name': 'dl-N-PDU-SequenceNumber', 'optional': True}, {'type': 'UL-N-PDU-SequenceNumber', 'name': 'ul-N-PDU-SequenceNumber', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-SetupOrModifyItemSecond-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PDPTypeInformation *PDPTypeInformation
    DataVolumeReportingIndication *DataVolumeReportingIndication
    DlGTPPDUSequenceNumber *DLGTPPDUSequenceNumber
    UlGTPPDUSequenceNumber *ULGTPPDUSequenceNumber
    DlNPDUSequenceNumber *DLNPDUSequenceNumber
    UlNPDUSequenceNumber *ULNPDUSequenceNumber
    IEExtensions *RABSetupOrModifyItemSecondExtIEs
}

func (self * RABSetupOrModifyItemSecond) Unpack(stream *Stream) {
    pDPTypeInformation_flag := 0x00000002
    dataVolumeReportingIndication_flag := 0x00000004
    dlGTPPDUSequenceNumber_flag := 0x00000008
    ulGTPPDUSequenceNumber_flag := 0x00000010
    dlNPDUSequenceNumber_flag := 0x00000020
    ulNPDUSequenceNumber_flag := 0x00000040
    iEExtensions_flag := 0x00000080
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(8)
    if (pDPTypeInformation_flag & _flags) == pDPTypeInformation_flag { //cond2
        self.PDPTypeInformation = &PDPTypeInformation{}//7{'type': 'PDP-TypeInformation', 'name': 'pDP-TypeInformation', 'optional': True}
        self.PDPTypeInformation.Unpack(stream)// p8
    }
    if (dataVolumeReportingIndication_flag & _flags) == dataVolumeReportingIndication_flag { //cond2
        self.DataVolumeReportingIndication = &DataVolumeReportingIndication{}//7{'type': 'DataVolumeReportingIndication', 'name': 'dataVolumeReportingIndication', 'optional': True}
        self.DataVolumeReportingIndication.Unpack(stream)// p8
    }
    if (dlGTPPDUSequenceNumber_flag & _flags) == dlGTPPDUSequenceNumber_flag { //cond2
        self.DlGTPPDUSequenceNumber = &DLGTPPDUSequenceNumber{}//7{'type': 'DL-GTP-PDU-SequenceNumber', 'name': 'dl-GTP-PDU-SequenceNumber', 'optional': True}
        self.DlGTPPDUSequenceNumber.Unpack(stream)// p8
    }
    if (ulGTPPDUSequenceNumber_flag & _flags) == ulGTPPDUSequenceNumber_flag { //cond2
        self.UlGTPPDUSequenceNumber = &ULGTPPDUSequenceNumber{}//7{'type': 'UL-GTP-PDU-SequenceNumber', 'name': 'ul-GTP-PDU-SequenceNumber', 'optional': True}
        self.UlGTPPDUSequenceNumber.Unpack(stream)// p8
    }
    if (dlNPDUSequenceNumber_flag & _flags) == dlNPDUSequenceNumber_flag { //cond2
        self.DlNPDUSequenceNumber = &DLNPDUSequenceNumber{}//7{'type': 'DL-N-PDU-SequenceNumber', 'name': 'dl-N-PDU-SequenceNumber', 'optional': True}
        self.DlNPDUSequenceNumber.Unpack(stream)// p8
    }
    if (ulNPDUSequenceNumber_flag & _flags) == ulNPDUSequenceNumber_flag { //cond2
        self.UlNPDUSequenceNumber = &ULNPDUSequenceNumber{}//7{'type': 'UL-N-PDU-SequenceNumber', 'name': 'ul-N-PDU-SequenceNumber', 'optional': True}
        self.UlNPDUSequenceNumber.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABSetupOrModifyItemSecondExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-SetupOrModifyItemSecond-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABSetupOrModifyItemSecondExtIEs, order_RABSetupOrModifyItemSecondExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABSetupOrModifyItemSecond) Pack(stream *Stream) {
    const pDPTypeInformation_flag uint = 0x00000002
    const dataVolumeReportingIndication_flag uint = 0x00000004
    const dlGTPPDUSequenceNumber_flag uint = 0x00000008
    const ulGTPPDUSequenceNumber_flag uint = 0x00000010
    const dlNPDUSequenceNumber_flag uint = 0x00000020
    const ulNPDUSequenceNumber_flag uint = 0x00000040
    const iEExtensions_flag uint = 0x00000080
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(8)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.PDPTypeInformation != nil { 
        _flags |= pDPTypeInformation_flag
        self.PDPTypeInformation.Pack(stream)
    }//end of optional
    if self.DataVolumeReportingIndication != nil { 
        _flags |= dataVolumeReportingIndication_flag
        self.DataVolumeReportingIndication.Pack(stream)
    }//end of optional
    if self.DlGTPPDUSequenceNumber != nil { 
        _flags |= dlGTPPDUSequenceNumber_flag
        self.DlGTPPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.UlGTPPDUSequenceNumber != nil { 
        _flags |= ulGTPPDUSequenceNumber_flag
        self.UlGTPPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.DlNPDUSequenceNumber != nil { 
        _flags |= dlNPDUSequenceNumber_flag
        self.DlNPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.UlNPDUSequenceNumber != nil { 
        _flags |= ulNPDUSequenceNumber_flag
        self.UlNPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABSetupOrModifyItemSecondExtIEs, order_RABSetupOrModifyItemSecondExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 8)
}//end

type RABAssignmentResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RAB-AssignmentResponseIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-AssignmentResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RABAssignmentResponseIEs
    ProtocolExtensions *RABAssignmentResponseExtensions
}

func (self * RABAssignmentResponse) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RABAssignmentResponseIEs, order_RABAssignmentResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RABAssignmentResponseExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-AssignmentResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RABAssignmentResponseExtensions, order_RABAssignmentResponseExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABAssignmentResponse) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RABAssignmentResponseIEs, order_RABAssignmentResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RABAssignmentResponseExtensions, order_RABAssignmentResponseExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABSetupOrModifiedList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-SetupOrModifiedItemIEs']}
    Item RABSetupOrModifiedItemIEs //UserType
}
func (self *RABSetupOrModifiedList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABSetupOrModifiedItemIEs,
       order_RABSetupOrModifiedItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABSetupOrModifiedList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABSetupOrModifiedItemIEs,
        order_RABSetupOrModifiedItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABSetupOrModifiedItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'TransportLayerAddress', 'name': 'transportLayerAddress', 'optional': True}, {'type': 'IuTransportAssociation', 'name': 'iuTransportAssociation', 'optional': True}, {'type': 'DataVolumeList', 'name': 'dl-dataVolumes', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-SetupOrModifiedItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    TransportLayerAddress *TransportLayerAddress
    IuTransportAssociation *IuTransportAssociation
    DldataVolumes *DataVolumeList
    IEExtensions *RABSetupOrModifiedItemExtIEs
}

func (self * RABSetupOrModifiedItem) Unpack(stream *Stream) {
    transportLayerAddress_flag := 0x00000002
    iuTransportAssociation_flag := 0x00000004
    dldataVolumes_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.RABID.Unpack(stream)// p8
    if (transportLayerAddress_flag & _flags) == transportLayerAddress_flag { //cond2
        self.TransportLayerAddress = &TransportLayerAddress{}//7{'type': 'TransportLayerAddress', 'name': 'transportLayerAddress', 'optional': True}
        self.TransportLayerAddress.Unpack(stream)// p8
    }
    if (iuTransportAssociation_flag & _flags) == iuTransportAssociation_flag { //cond2
        self.IuTransportAssociation = &IuTransportAssociation{}//7{'type': 'IuTransportAssociation', 'name': 'iuTransportAssociation', 'optional': True}
        self.IuTransportAssociation.Unpack(stream)// p8
    }
    if (dldataVolumes_flag & _flags) == dldataVolumes_flag { //cond2
        self.DldataVolumes = &DataVolumeList{}//7{'type': 'DataVolumeList', 'name': 'dl-dataVolumes', 'optional': True}
        self.DldataVolumes.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABSetupOrModifiedItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-SetupOrModifiedItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABSetupOrModifiedItemExtIEs, order_RABSetupOrModifiedItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABSetupOrModifiedItem) Pack(stream *Stream) {
    const transportLayerAddress_flag uint = 0x00000002
    const iuTransportAssociation_flag uint = 0x00000004
    const dldataVolumes_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.TransportLayerAddress != nil { 
        _flags |= transportLayerAddress_flag
        self.TransportLayerAddress.Pack(stream)
    }//end of optional
    if self.IuTransportAssociation != nil { 
        _flags |= iuTransportAssociation_flag
        self.IuTransportAssociation.Pack(stream)
    }//end of optional
    if self.DldataVolumes != nil { 
        _flags |= dldataVolumes_flag
        self.DldataVolumes.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABSetupOrModifiedItemExtIEs, order_RABSetupOrModifiedItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

type RABReleasedList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-ReleasedItemIEs']}
    Item RABReleasedItemIEs //UserType
}
func (self *RABReleasedList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABReleasedItemIEs,
       order_RABReleasedItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABReleasedList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABReleasedItemIEs,
        order_RABReleasedItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABReleasedItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'DataVolumeList', 'name': 'dl-dataVolumes', 'optional': True}, {'type': 'DL-GTP-PDU-SequenceNumber', 'name': 'dL-GTP-PDU-SequenceNumber', 'optional': True}, {'type': 'UL-GTP-PDU-SequenceNumber', 'name': 'uL-GTP-PDU-SequenceNumber', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ReleasedItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    DldataVolumes *DataVolumeList
    DLGTPPDUSequenceNumber *DLGTPPDUSequenceNumber
    ULGTPPDUSequenceNumber *ULGTPPDUSequenceNumber
    IEExtensions *RABReleasedItemExtIEs
}

func (self * RABReleasedItem) Unpack(stream *Stream) {
    dldataVolumes_flag := 0x00000002
    dLGTPPDUSequenceNumber_flag := 0x00000004
    uLGTPPDUSequenceNumber_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.RABID.Unpack(stream)// p8
    if (dldataVolumes_flag & _flags) == dldataVolumes_flag { //cond2
        self.DldataVolumes = &DataVolumeList{}//7{'type': 'DataVolumeList', 'name': 'dl-dataVolumes', 'optional': True}
        self.DldataVolumes.Unpack(stream)// p8
    }
    if (dLGTPPDUSequenceNumber_flag & _flags) == dLGTPPDUSequenceNumber_flag { //cond2
        self.DLGTPPDUSequenceNumber = &DLGTPPDUSequenceNumber{}//7{'type': 'DL-GTP-PDU-SequenceNumber', 'name': 'dL-GTP-PDU-SequenceNumber', 'optional': True}
        self.DLGTPPDUSequenceNumber.Unpack(stream)// p8
    }
    if (uLGTPPDUSequenceNumber_flag & _flags) == uLGTPPDUSequenceNumber_flag { //cond2
        self.ULGTPPDUSequenceNumber = &ULGTPPDUSequenceNumber{}//7{'type': 'UL-GTP-PDU-SequenceNumber', 'name': 'uL-GTP-PDU-SequenceNumber', 'optional': True}
        self.ULGTPPDUSequenceNumber.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABReleasedItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ReleasedItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABReleasedItemExtIEs, order_RABReleasedItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABReleasedItem) Pack(stream *Stream) {
    const dldataVolumes_flag uint = 0x00000002
    const dLGTPPDUSequenceNumber_flag uint = 0x00000004
    const uLGTPPDUSequenceNumber_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.DldataVolumes != nil { 
        _flags |= dldataVolumes_flag
        self.DldataVolumes.Pack(stream)
    }//end of optional
    if self.DLGTPPDUSequenceNumber != nil { 
        _flags |= dLGTPPDUSequenceNumber_flag
        self.DLGTPPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.ULGTPPDUSequenceNumber != nil { 
        _flags |= uLGTPPDUSequenceNumber_flag
        self.ULGTPPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABReleasedItemExtIEs, order_RABReleasedItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *DataVolumeList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]DataVolumeList_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *DataVolumeList_Item) { //[{'type': 'UnsuccessfullyTransmittedDataVolume', 'name': 'dl-UnsuccessfullyTransmittedDataVolume'}, {'type': 'DataVolumeReference', 'name': 'dataVolumeReference', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DataVolumeList-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        dataVolumeReference_flag := 0x00000002
        iEExtensions_flag := 0x00000004
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(3)
        self.DlUnsuccessfullyTransmittedDataVolume.Unpack(stream)// p8
        if (dataVolumeReference_flag & _flags) == dataVolumeReference_flag { //cond2
            self.DataVolumeReference = &DataVolumeReference{}//7{'type': 'DataVolumeReference', 'name': 'dataVolumeReference', 'optional': True}
            self.DataVolumeReference.Unpack(stream)// p8
        }
        if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
            self.IEExtensions = &DataVolumeListExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DataVolumeList-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
            IEExtensions := ProtocolExtensionContainer {table_DataVolumeListExtIEs, order_DataVolumeListExtIEs} // p3
            IEExtensions.Unpack(stream, &self.IEExtensions) // p3
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    for item := 0; item <_size; item +=1 {
        Unpack_Item(stream, &self.Items[item])
    }
}


func (self *DataVolumeList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    var Pack_Item = func(stream *Stream, self DataVolumeList_Item) {//seq
        const dataVolumeReference_flag uint = 0x00000002
        const iEExtensions_flag uint = 0x00000004
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(3)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        self.DlUnsuccessfullyTransmittedDataVolume.Pack(stream)
        if self.DataVolumeReference != nil { 
            _flags |= dataVolumeReference_flag
            self.DataVolumeReference.Pack(stream)
        }//end of optional
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_DataVolumeListExtIEs, order_DataVolumeListExtIEs} // p3
            IEExtensions.Pack(stream, &self.IEExtensions)
        }//end of optional
        if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
        stream.set_flags(_flags, _flagReserve, 3)
    }//end
    for _, item := range self.Items { // seqof base type
        Pack_Item(stream, item)
    }
    return

}


type DataVolumeList_Item struct { // [{'type': 'UnsuccessfullyTransmittedDataVolume', 'name': 'dl-UnsuccessfullyTransmittedDataVolume'}, {'type': 'DataVolumeReference', 'name': 'dataVolumeReference', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DataVolumeList-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DlUnsuccessfullyTransmittedDataVolume UnsuccessfullyTransmittedDataVolume
    DataVolumeReference *DataVolumeReference
    IEExtensions *DataVolumeListExtIEs
}
type DataVolumeList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'UnsuccessfullyTransmittedDataVolume', 'name': 'dl-UnsuccessfullyTransmittedDataVolume'}, {'type': 'DataVolumeReference', 'name': 'dataVolumeReference', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DataVolumeList-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxNrOfVol')]}
    Items []DataVolumeList_Item
}

type RABQueuedList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-QueuedItemIEs']}
    Item RABQueuedItemIEs //UserType
}
func (self *RABQueuedList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABQueuedItemIEs,
       order_RABQueuedItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABQueuedList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABQueuedItemIEs,
        order_RABQueuedItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABQueuedItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-QueuedItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    IEExtensions *RABQueuedItemExtIEs
}

func (self * RABQueuedItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABQueuedItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-QueuedItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABQueuedItemExtIEs, order_RABQueuedItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABQueuedItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABQueuedItemExtIEs, order_RABQueuedItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABReleaseFailedList struct{ //{'type': 'RAB-FailedList'}
    RABFailedList //UserTypeTODO: Unimplimented 
}
func (self *RABReleaseFailedList) Unpack(st *Stream){ // ut3 'RABFailedList'
    self.RABFailedList.Unpack(st) // ut32
}

func (self *RABReleaseFailedList) Pack(st *Stream){
    self.RABFailedList.Pack(st) // ut32
}

type GERANIumodeRABFailedListRABAssgntResponse struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['GERAN-Iumode-RAB-Failed-RABAssgntResponse-ItemIEs']}
    Item GERANIumodeRABFailedRABAssgntResponseItemIEs //UserType
}
func (self *GERANIumodeRABFailedListRABAssgntResponse) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_GERANIumodeRABFailedRABAssgntResponseItemIEs,
       order_GERANIumodeRABFailedRABAssgntResponseItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *GERANIumodeRABFailedListRABAssgntResponse) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_GERANIumodeRABFailedRABAssgntResponseItemIEs,
        order_GERANIumodeRABFailedRABAssgntResponseItemIEs,
    }
        val.Pack(st, self.Item)
}

type GERANIumodeRABFailedRABAssgntResponseItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'Cause', 'name': 'cause'}, {'type': 'GERAN-Classmark', 'name': 'gERAN-Classmark', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GERAN-Iumode-RAB-Failed-RABAssgntResponse-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    Cause Cause
    GERANClassmark *GERANClassmark
    IEExtensions *GERANIumodeRABFailedRABAssgntResponseItemExtIEs
}

func (self * GERANIumodeRABFailedRABAssgntResponseItem) Unpack(stream *Stream) {
    gERANClassmark_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RABID.Unpack(stream)// p8
    self.Cause.Unpack(stream)// p8
    if (gERANClassmark_flag & _flags) == gERANClassmark_flag { //cond2
        self.GERANClassmark = &GERANClassmark{}//7{'type': 'GERAN-Classmark', 'name': 'gERAN-Classmark', 'optional': True}
        self.GERANClassmark.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GERANIumodeRABFailedRABAssgntResponseItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GERAN-Iumode-RAB-Failed-RABAssgntResponse-Item-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GERANIumodeRABFailedRABAssgntResponseItemExtIEs, order_GERANIumodeRABFailedRABAssgntResponseItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GERANIumodeRABFailedRABAssgntResponseItem) Pack(stream *Stream) {
    const gERANClassmark_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    self.Cause.Pack(stream)
    if self.GERANClassmark != nil { 
        _flags |= gERANClassmark_flag
        self.GERANClassmark.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GERANIumodeRABFailedRABAssgntResponseItemExtIEs, order_GERANIumodeRABFailedRABAssgntResponseItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
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

type RANAPRelocationInformation struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RANAP-RelocationInformationIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RANAP-RelocationInformationExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RANAPRelocationInformationIEs
    ProtocolExtensions *RANAPRelocationInformationExtensions
}

func (self * RANAPRelocationInformation) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RANAPRelocationInformationIEs, order_RANAPRelocationInformationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RANAPRelocationInformationExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RANAP-RelocationInformationExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RANAPRelocationInformationExtensions, order_RANAPRelocationInformationExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANAPRelocationInformation) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RANAPRelocationInformationIEs, order_RANAPRelocationInformationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RANAPRelocationInformationExtensions, order_RANAPRelocationInformationExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type DirectTransferInformationListRANAPRelocInf struct{ //{'type': 'DirectTransfer-IE-ContainerList', 'actual-parameters': ['DirectTransferInformationItemIEs-RANAP-RelocInf']}
    Item DirectTransferInformationItemIEsRANAPRelocInf //UserType
}
func (self *DirectTransferInformationListRANAPRelocInf) Unpack(st *Stream){ // ut3 'DirectTransferIEContainerList'
    val := DirectTransferIEContainerList{// ut31
       table_DirectTransferInformationItemIEsRANAPRelocInf,
       order_DirectTransferInformationItemIEsRANAPRelocInf,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *DirectTransferInformationListRANAPRelocInf) Pack(st *Stream){
    val := DirectTransferIEContainerList{// ut31
        table_DirectTransferInformationItemIEsRANAPRelocInf,
        order_DirectTransferInformationItemIEsRANAPRelocInf,
    }
        val.Pack(st, self.Item)
}

type DirectTransferInformationItemRANAPRelocInf struct { // [{'type': 'NAS-PDU', 'name': 'nAS-PDU'}, {'type': 'SAPI', 'name': 'sAPI'}, {'type': 'CN-DomainIndicator', 'name': 'cN-DomainIndicator'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RANAP-DirectTransferInformationItem-ExtIEs-RANAP-RelocInf'], 'name': 'iE-Extensions', 'optional': True}, None]
    NASPDU NASPDU
    SAPI SAPI
    CNDomainIndicator CNDomainIndicator
    IEExtensions *RANAPDirectTransferInformationItemExtIEsRANAPRelocInf
}

func (self * DirectTransferInformationItemRANAPRelocInf) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.NASPDU.Unpack(stream)// p8
    self.SAPI.Unpack(stream)// p8
    self.CNDomainIndicator.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RANAPDirectTransferInformationItemExtIEsRANAPRelocInf{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RANAP-DirectTransferInformationItem-ExtIEs-RANAP-RelocInf'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RANAPDirectTransferInformationItemExtIEsRANAPRelocInf, order_RANAPDirectTransferInformationItemExtIEsRANAPRelocInf} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DirectTransferInformationItemRANAPRelocInf) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.NASPDU.Pack(stream)
    self.SAPI.Pack(stream)
    self.CNDomainIndicator.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RANAPDirectTransferInformationItemExtIEsRANAPRelocInf, order_RANAPDirectTransferInformationItemExtIEsRANAPRelocInf} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABContextListRANAPRelocInf struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-ContextItemIEs-RANAP-RelocInf']}
    Item RABContextItemIEsRANAPRelocInf //UserType
}
func (self *RABContextListRANAPRelocInf) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABContextItemIEsRANAPRelocInf,
       order_RABContextItemIEsRANAPRelocInf,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABContextListRANAPRelocInf) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABContextItemIEsRANAPRelocInf,
        order_RABContextItemIEsRANAPRelocInf,
    }
        val.Pack(st, self.Item)
}

type RABContextItemRANAPRelocInf struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'DL-GTP-PDU-SequenceNumber', 'name': 'dl-GTP-PDU-SequenceNumber', 'optional': True}, {'type': 'UL-GTP-PDU-SequenceNumber', 'name': 'ul-GTP-PDU-SequenceNumber', 'optional': True}, {'type': 'DL-N-PDU-SequenceNumber', 'name': 'dl-N-PDU-SequenceNumber', 'optional': True}, {'type': 'UL-N-PDU-SequenceNumber', 'name': 'ul-N-PDU-SequenceNumber', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ContextItem-ExtIEs-RANAP-RelocInf'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    DlGTPPDUSequenceNumber *DLGTPPDUSequenceNumber
    UlGTPPDUSequenceNumber *ULGTPPDUSequenceNumber
    DlNPDUSequenceNumber *DLNPDUSequenceNumber
    UlNPDUSequenceNumber *ULNPDUSequenceNumber
    IEExtensions *RABContextItemExtIEsRANAPRelocInf
}

func (self * RABContextItemRANAPRelocInf) Unpack(stream *Stream) {
    dlGTPPDUSequenceNumber_flag := 0x00000002
    ulGTPPDUSequenceNumber_flag := 0x00000004
    dlNPDUSequenceNumber_flag := 0x00000008
    ulNPDUSequenceNumber_flag := 0x00000010
    iEExtensions_flag := 0x00000020
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(6)
    self.RABID.Unpack(stream)// p8
    if (dlGTPPDUSequenceNumber_flag & _flags) == dlGTPPDUSequenceNumber_flag { //cond2
        self.DlGTPPDUSequenceNumber = &DLGTPPDUSequenceNumber{}//7{'type': 'DL-GTP-PDU-SequenceNumber', 'name': 'dl-GTP-PDU-SequenceNumber', 'optional': True}
        self.DlGTPPDUSequenceNumber.Unpack(stream)// p8
    }
    if (ulGTPPDUSequenceNumber_flag & _flags) == ulGTPPDUSequenceNumber_flag { //cond2
        self.UlGTPPDUSequenceNumber = &ULGTPPDUSequenceNumber{}//7{'type': 'UL-GTP-PDU-SequenceNumber', 'name': 'ul-GTP-PDU-SequenceNumber', 'optional': True}
        self.UlGTPPDUSequenceNumber.Unpack(stream)// p8
    }
    if (dlNPDUSequenceNumber_flag & _flags) == dlNPDUSequenceNumber_flag { //cond2
        self.DlNPDUSequenceNumber = &DLNPDUSequenceNumber{}//7{'type': 'DL-N-PDU-SequenceNumber', 'name': 'dl-N-PDU-SequenceNumber', 'optional': True}
        self.DlNPDUSequenceNumber.Unpack(stream)// p8
    }
    if (ulNPDUSequenceNumber_flag & _flags) == ulNPDUSequenceNumber_flag { //cond2
        self.UlNPDUSequenceNumber = &ULNPDUSequenceNumber{}//7{'type': 'UL-N-PDU-SequenceNumber', 'name': 'ul-N-PDU-SequenceNumber', 'optional': True}
        self.UlNPDUSequenceNumber.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABContextItemExtIEsRANAPRelocInf{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ContextItem-ExtIEs-RANAP-RelocInf'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABContextItemExtIEsRANAPRelocInf, order_RABContextItemExtIEsRANAPRelocInf} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABContextItemRANAPRelocInf) Pack(stream *Stream) {
    const dlGTPPDUSequenceNumber_flag uint = 0x00000002
    const ulGTPPDUSequenceNumber_flag uint = 0x00000004
    const dlNPDUSequenceNumber_flag uint = 0x00000008
    const ulNPDUSequenceNumber_flag uint = 0x00000010
    const iEExtensions_flag uint = 0x00000020
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(6)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    if self.DlGTPPDUSequenceNumber != nil { 
        _flags |= dlGTPPDUSequenceNumber_flag
        self.DlGTPPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.UlGTPPDUSequenceNumber != nil { 
        _flags |= ulGTPPDUSequenceNumber_flag
        self.UlGTPPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.DlNPDUSequenceNumber != nil { 
        _flags |= dlNPDUSequenceNumber_flag
        self.DlNPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.UlNPDUSequenceNumber != nil { 
        _flags |= ulNPDUSequenceNumber_flag
        self.UlNPDUSequenceNumber.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABContextItemExtIEsRANAPRelocInf, order_RABContextItemExtIEsRANAPRelocInf} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

type RABModifyRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['RAB-ModifyRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ModifyRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs RABModifyRequestIEs
    ProtocolExtensions *RABModifyRequestExtensions
}

func (self * RABModifyRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_RABModifyRequestIEs, order_RABModifyRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &RABModifyRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ModifyRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_RABModifyRequestExtensions, order_RABModifyRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABModifyRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_RABModifyRequestIEs, order_RABModifyRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_RABModifyRequestExtensions, order_RABModifyRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RABModifyList struct{ //{'type': 'RAB-IE-ContainerList', 'actual-parameters': ['RAB-ModifyItemIEs']}
    Item RABModifyItemIEs //UserType
}
func (self *RABModifyList) Unpack(st *Stream){ // ut3 'RABIEContainerList'
    val := RABIEContainerList{// ut31
       table_RABModifyItemIEs,
       order_RABModifyItemIEs,
    }
    val.Unpack(st, &self.Item) // ut31
}

func (self *RABModifyList) Pack(st *Stream){
    val := RABIEContainerList{// ut31
        table_RABModifyItemIEs,
        order_RABModifyItemIEs,
    }
        val.Pack(st, self.Item)
}

type RABModifyItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'Requested-RAB-Parameter-Values', 'name': 'requested-RAB-Parameter-Values'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ModifyItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    RequestedRABParameterValues RequestedRABParameterValues
    IEExtensions *RABModifyItemExtIEs
}

func (self * RABModifyItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    self.RequestedRABParameterValues.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABModifyItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-ModifyItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABModifyItemExtIEs, order_RABModifyItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABModifyItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    self.RequestedRABParameterValues.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABModifyItemExtIEs, order_RABModifyItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type LocationRelatedDataRequest struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['LocationRelatedDataRequestIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LocationRelatedDataRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs LocationRelatedDataRequestIEs
    ProtocolExtensions *LocationRelatedDataRequestExtensions
}

func (self * LocationRelatedDataRequest) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_LocationRelatedDataRequestIEs, order_LocationRelatedDataRequestIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &LocationRelatedDataRequestExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LocationRelatedDataRequestExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_LocationRelatedDataRequestExtensions, order_LocationRelatedDataRequestExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * LocationRelatedDataRequest) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_LocationRelatedDataRequestIEs, order_LocationRelatedDataRequestIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_LocationRelatedDataRequestExtensions, order_LocationRelatedDataRequestExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type LocationRelatedDataResponse struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['LocationRelatedDataResponseIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LocationRelatedDataResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs LocationRelatedDataResponseIEs
    ProtocolExtensions *LocationRelatedDataResponseExtensions
}

func (self * LocationRelatedDataResponse) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_LocationRelatedDataResponseIEs, order_LocationRelatedDataResponseIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &LocationRelatedDataResponseExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LocationRelatedDataResponseExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_LocationRelatedDataResponseExtensions, order_LocationRelatedDataResponseExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * LocationRelatedDataResponse) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_LocationRelatedDataResponseIEs, order_LocationRelatedDataResponseIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_LocationRelatedDataResponseExtensions, order_LocationRelatedDataResponseExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type LocationRelatedDataFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['LocationRelatedDataFailureIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LocationRelatedDataFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs LocationRelatedDataFailureIEs
    ProtocolExtensions *LocationRelatedDataFailureExtensions
}

func (self * LocationRelatedDataFailure) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_LocationRelatedDataFailureIEs, order_LocationRelatedDataFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &LocationRelatedDataFailureExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LocationRelatedDataFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_LocationRelatedDataFailureExtensions, order_LocationRelatedDataFailureExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * LocationRelatedDataFailure) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_LocationRelatedDataFailureIEs, order_LocationRelatedDataFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_LocationRelatedDataFailureExtensions, order_LocationRelatedDataFailureExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type InformationTransferIndication struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['InformationTransferIndicationIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InformationTransferIndicationExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs InformationTransferIndicationIEs
    ProtocolExtensions *InformationTransferIndicationExtensions
}

func (self * InformationTransferIndication) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_InformationTransferIndicationIEs, order_InformationTransferIndicationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &InformationTransferIndicationExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InformationTransferIndicationExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_InformationTransferIndicationExtensions, order_InformationTransferIndicationExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InformationTransferIndication) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_InformationTransferIndicationIEs, order_InformationTransferIndicationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_InformationTransferIndicationExtensions, order_InformationTransferIndicationExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type InformationTransferConfirmation struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['InformationTransferConfirmationIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InformationTransferConfirmationExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs InformationTransferConfirmationIEs
    ProtocolExtensions *InformationTransferConfirmationExtensions
}

func (self * InformationTransferConfirmation) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_InformationTransferConfirmationIEs, order_InformationTransferConfirmationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &InformationTransferConfirmationExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InformationTransferConfirmationExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_InformationTransferConfirmationExtensions, order_InformationTransferConfirmationExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InformationTransferConfirmation) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_InformationTransferConfirmationIEs, order_InformationTransferConfirmationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_InformationTransferConfirmationExtensions, order_InformationTransferConfirmationExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type InformationTransferFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['InformationTransferFailureIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InformationTransferFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs InformationTransferFailureIEs
    ProtocolExtensions *InformationTransferFailureExtensions
}

func (self * InformationTransferFailure) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_InformationTransferFailureIEs, order_InformationTransferFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &InformationTransferFailureExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InformationTransferFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_InformationTransferFailureExtensions, order_InformationTransferFailureExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InformationTransferFailure) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_InformationTransferFailureIEs, order_InformationTransferFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_InformationTransferFailureExtensions, order_InformationTransferFailureExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UESpecificInformationIndication struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['UESpecificInformationIndicationIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UESpecificInformationIndicationExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs UESpecificInformationIndicationIEs
    ProtocolExtensions *UESpecificInformationIndicationExtensions
}

func (self * UESpecificInformationIndication) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_UESpecificInformationIndicationIEs, order_UESpecificInformationIndicationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &UESpecificInformationIndicationExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UESpecificInformationIndicationExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_UESpecificInformationIndicationExtensions, order_UESpecificInformationIndicationExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UESpecificInformationIndication) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_UESpecificInformationIndicationIEs, order_UESpecificInformationIndicationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_UESpecificInformationIndicationExtensions, order_UESpecificInformationIndicationExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type DirectInformationTransfer struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['DirectInformationTransferIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DirectInformationTransferExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs DirectInformationTransferIEs
    ProtocolExtensions *DirectInformationTransferExtensions
}

func (self * DirectInformationTransfer) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_DirectInformationTransferIEs, order_DirectInformationTransferIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &DirectInformationTransferExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['DirectInformationTransferExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_DirectInformationTransferExtensions, order_DirectInformationTransferExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DirectInformationTransfer) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_DirectInformationTransferIEs, order_DirectInformationTransferIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_DirectInformationTransferExtensions, order_DirectInformationTransferExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UplinkInformationTransferIndication struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['UplinkInformationTransferIndicationIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UplinkInformationTransferIndicationExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs UplinkInformationTransferIndicationIEs
    ProtocolExtensions *UplinkInformationTransferIndicationExtensions
}

func (self * UplinkInformationTransferIndication) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_UplinkInformationTransferIndicationIEs, order_UplinkInformationTransferIndicationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &UplinkInformationTransferIndicationExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UplinkInformationTransferIndicationExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_UplinkInformationTransferIndicationExtensions, order_UplinkInformationTransferIndicationExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UplinkInformationTransferIndication) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_UplinkInformationTransferIndicationIEs, order_UplinkInformationTransferIndicationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_UplinkInformationTransferIndicationExtensions, order_UplinkInformationTransferIndicationExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UplinkInformationTransferConfirmation struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['UplinkInformationTransferConfirmationIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UplinkInformationTransferConfirmationExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs UplinkInformationTransferConfirmationIEs
    ProtocolExtensions *UplinkInformationTransferConfirmationExtensions
}

func (self * UplinkInformationTransferConfirmation) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_UplinkInformationTransferConfirmationIEs, order_UplinkInformationTransferConfirmationIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &UplinkInformationTransferConfirmationExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UplinkInformationTransferConfirmationExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_UplinkInformationTransferConfirmationExtensions, order_UplinkInformationTransferConfirmationExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UplinkInformationTransferConfirmation) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_UplinkInformationTransferConfirmationIEs, order_UplinkInformationTransferConfirmationIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_UplinkInformationTransferConfirmationExtensions, order_UplinkInformationTransferConfirmationExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UplinkInformationTransferFailure struct { // [{'type': 'ProtocolIE-Container', 'actual-parameters': ['UplinkInformationTransferFailureIEs'], 'name': 'protocolIEs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UplinkInformationTransferFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}, None]
    ProtocolIEs UplinkInformationTransferFailureIEs
    ProtocolExtensions *UplinkInformationTransferFailureExtensions
}

func (self * UplinkInformationTransferFailure) Unpack(stream *Stream) {
    protocolExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    ProtocolIEs := ProtocolIEContainer {table_UplinkInformationTransferFailureIEs, order_UplinkInformationTransferFailureIEs} // p3
    ProtocolIEs.Unpack(stream, &self.ProtocolIEs) // p3
    if (protocolExtensions_flag & _flags) == protocolExtensions_flag { //cond2
        self.ProtocolExtensions = &UplinkInformationTransferFailureExtensions{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UplinkInformationTransferFailureExtensions'], 'name': 'protocolExtensions', 'optional': True}
        ProtocolExtensions := ProtocolExtensionContainer {table_UplinkInformationTransferFailureExtensions, order_UplinkInformationTransferFailureExtensions} // p3
        ProtocolExtensions.Unpack(stream, &self.ProtocolExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UplinkInformationTransferFailure) Pack(stream *Stream) {
    const protocolExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    ProtocolIEs := &ProtocolIEContainer {table_UplinkInformationTransferFailureIEs, order_UplinkInformationTransferFailureIEs} // p3
    ProtocolIEs.Pack(stream, &self.ProtocolIEs)
    if self.ProtocolExtensions != nil { 
        _flags |= protocolExtensions_flag
        ProtocolExtensions := &ProtocolExtensionContainer {table_UplinkInformationTransferFailureExtensions, order_UplinkInformationTransferFailureExtensions} // p3
        ProtocolExtensions.Pack(stream, &self.ProtocolExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type AccuracyFulfilmentIndicator struct {
  Value int
}
const (
    AccuracyFulfilmentIndicatorrequested_Accuracy_Fulfilled = 0
    AccuracyFulfilmentIndicatorrequested_Accuracy_Not_Fulfilled = 1

    /* Extensions */
)
func (self *AccuracyFulfilmentIndicator) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *AccuracyFulfilmentIndicator) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type AllocationOrRetentionPriority struct { // [{'type': 'PriorityLevel', 'name': 'priorityLevel'}, {'type': 'Pre-emptionCapability', 'name': 'pre-emptionCapability'}, {'type': 'Pre-emptionVulnerability', 'name': 'pre-emptionVulnerability'}, {'type': 'QueuingAllowed', 'name': 'queuingAllowed'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AllocationOrRetentionPriority-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PriorityLevel PriorityLevel
    PreemptionCapability PreemptionCapability
    PreemptionVulnerability PreemptionVulnerability
    QueuingAllowed QueuingAllowed
    IEExtensions *AllocationOrRetentionPriorityExtIEs
}

func (self * AllocationOrRetentionPriority) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PriorityLevel.Unpack(stream)// p8
    self.PreemptionCapability.Unpack(stream)// p8
    self.PreemptionVulnerability.Unpack(stream)// p8
    self.QueuingAllowed.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &AllocationOrRetentionPriorityExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AllocationOrRetentionPriority-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_AllocationOrRetentionPriorityExtIEs, order_AllocationOrRetentionPriorityExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AllocationOrRetentionPriority) Pack(stream *Stream) {
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
    self.QueuingAllowed.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_AllocationOrRetentionPriorityExtIEs, order_AllocationOrRetentionPriorityExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type AltRABParameters struct { // [{'type': 'Alt-RAB-Parameter-MaxBitrateInf', 'name': 'altMaxBitrateInf', 'optional': True}, {'type': 'Alt-RAB-Parameter-GuaranteedBitrateInf', 'name': 'altGuaranteedBitRateInf', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Alt-RAB-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    AltMaxBitrateInf *AltRABParameterMaxBitrateInf
    AltGuaranteedBitRateInf *AltRABParameterGuaranteedBitrateInf
    IEExtensions *AltRABParametersExtIEs
}

func (self * AltRABParameters) Unpack(stream *Stream) {
    altMaxBitrateInf_flag := 0x00000002
    altGuaranteedBitRateInf_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (altMaxBitrateInf_flag & _flags) == altMaxBitrateInf_flag { //cond2
        self.AltMaxBitrateInf = &AltRABParameterMaxBitrateInf{}//7{'type': 'Alt-RAB-Parameter-MaxBitrateInf', 'name': 'altMaxBitrateInf', 'optional': True}
        self.AltMaxBitrateInf.Unpack(stream)// p8
    }
    if (altGuaranteedBitRateInf_flag & _flags) == altGuaranteedBitRateInf_flag { //cond2
        self.AltGuaranteedBitRateInf = &AltRABParameterGuaranteedBitrateInf{}//7{'type': 'Alt-RAB-Parameter-GuaranteedBitrateInf', 'name': 'altGuaranteedBitRateInf', 'optional': True}
        self.AltGuaranteedBitRateInf.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &AltRABParametersExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Alt-RAB-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_AltRABParametersExtIEs, order_AltRABParametersExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AltRABParameters) Pack(stream *Stream) {
    const altMaxBitrateInf_flag uint = 0x00000002
    const altGuaranteedBitRateInf_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.AltMaxBitrateInf != nil { 
        _flags |= altMaxBitrateInf_flag
        self.AltMaxBitrateInf.Pack(stream)
    }//end of optional
    if self.AltGuaranteedBitRateInf != nil { 
        _flags |= altGuaranteedBitRateInf_flag
        self.AltGuaranteedBitRateInf.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_AltRABParametersExtIEs, order_AltRABParametersExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type AltRABParameterGuaranteedBitrateInf struct { // [{'type': 'Alt-RAB-Parameter-GuaranteedBitrateType', 'name': 'altGuaranteedBitrateType'}, {'type': 'Alt-RAB-Parameter-GuaranteedBitrates', 'name': 'altGuaranteedBitrates', 'optional': True}, None]
    AltGuaranteedBitrateType AltRABParameterGuaranteedBitrateType
    AltGuaranteedBitrates *AltRABParameterGuaranteedBitrates
}

func (self * AltRABParameterGuaranteedBitrateInf) Unpack(stream *Stream) {
    altGuaranteedBitrates_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.AltGuaranteedBitrateType.Unpack(stream)// p8
    if (altGuaranteedBitrates_flag & _flags) == altGuaranteedBitrates_flag { //cond2
        self.AltGuaranteedBitrates = &AltRABParameterGuaranteedBitrates{}//7{'type': 'Alt-RAB-Parameter-GuaranteedBitrates', 'name': 'altGuaranteedBitrates', 'optional': True}
        self.AltGuaranteedBitrates.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AltRABParameterGuaranteedBitrateInf) Pack(stream *Stream) {
    const altGuaranteedBitrates_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.AltGuaranteedBitrateType.Pack(stream)
    if self.AltGuaranteedBitrates != nil { 
        _flags |= altGuaranteedBitrates_flag
        self.AltGuaranteedBitrates.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type AltRABParameterGuaranteedBitrateType struct {
  Value int
}
const (
    AltRABParameterGuaranteedBitrateTypeunspecified = 0
    AltRABParameterGuaranteedBitrateTypevalue_range = 1
    AltRABParameterGuaranteedBitrateTypediscrete_values = 2

    /* Extensions */
)
func (self *AltRABParameterGuaranteedBitrateType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *AltRABParameterGuaranteedBitrateType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
func (self *AltRABParameterGuaranteedBitrates) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(16)
    _size += 1
    self.Items = make([]AltRABParameterGuaranteedBitrateList, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *AltRABParameterGuaranteedBitrates) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 16)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type AltRABParameterGuaranteedBitrates struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'Alt-RAB-Parameter-GuaranteedBitrateList'}, 'size': [(1, 'maxNrOfAltValues')]}
    Items []AltRABParameterGuaranteedBitrateList
}

func (self *AltRABParameterGuaranteedBitrateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]GuaranteedBitrate, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *AltRABParameterGuaranteedBitrateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type AltRABParameterGuaranteedBitrateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GuaranteedBitrate'}, 'size': [(1, 'maxNrOfSeparateTrafficDirections')]}
    Items []GuaranteedBitrate
}

type AltRABParameterMaxBitrateInf struct { // [{'type': 'Alt-RAB-Parameter-MaxBitrateType', 'name': 'altMaxBitrateType'}, {'type': 'Alt-RAB-Parameter-MaxBitrates', 'name': 'altMaxBitrates', 'optional': True}, None]
    AltMaxBitrateType AltRABParameterMaxBitrateType
    AltMaxBitrates *AltRABParameterMaxBitrates
}

func (self * AltRABParameterMaxBitrateInf) Unpack(stream *Stream) {
    altMaxBitrates_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.AltMaxBitrateType.Unpack(stream)// p8
    if (altMaxBitrates_flag & _flags) == altMaxBitrates_flag { //cond2
        self.AltMaxBitrates = &AltRABParameterMaxBitrates{}//7{'type': 'Alt-RAB-Parameter-MaxBitrates', 'name': 'altMaxBitrates', 'optional': True}
        self.AltMaxBitrates.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AltRABParameterMaxBitrateInf) Pack(stream *Stream) {
    const altMaxBitrates_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.AltMaxBitrateType.Pack(stream)
    if self.AltMaxBitrates != nil { 
        _flags |= altMaxBitrates_flag
        self.AltMaxBitrates.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type AltRABParameterMaxBitrateType struct {
  Value int
}
const (
    AltRABParameterMaxBitrateTypeunspecified = 0
    AltRABParameterMaxBitrateTypevalue_range = 1
    AltRABParameterMaxBitrateTypediscrete_values = 2

    /* Extensions */
)
func (self *AltRABParameterMaxBitrateType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *AltRABParameterMaxBitrateType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
func (self *AltRABParameterMaxBitrates) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(16)
    _size += 1
    self.Items = make([]AltRABParameterMaxBitrateList, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *AltRABParameterMaxBitrates) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 16)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type AltRABParameterMaxBitrates struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'Alt-RAB-Parameter-MaxBitrateList'}, 'size': [(1, 'maxNrOfAltValues')]}
    Items []AltRABParameterMaxBitrateList
}

func (self *AltRABParameterMaxBitrateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]MaxBitrate, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *AltRABParameterMaxBitrateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type AltRABParameterMaxBitrateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MaxBitrate'}, 'size': [(1, 'maxNrOfSeparateTrafficDirections')]}
    Items []MaxBitrate
}

func (self *AreaIdentity)Unpack(stream *Stream) {
    //coptions := []string{"sAI","geographicalArea"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in AreaIdentity\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.SAI = &SAI{}//cho6
        self.SAI.Unpack(stream)
    } else if choice == 1 { //ch2
        self.GeographicalArea = &GeographicalArea{}//cho6
        self.GeographicalArea.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * AreaIdentity) Pack(stream *Stream) {
    if self.SAI != nil {
        stream.set_choice(0, 1, 1, 2)
        self.SAI.Pack(stream)//2
    } else if self.GeographicalArea != nil {
        stream.set_choice(1, 1, 1, 2)
        self.GeographicalArea.Pack(stream)//2
    }

}
type AreaIdentity struct { //[{'type': 'SAI', 'name': 'sAI'}, {'type': 'GeographicalArea', 'name': 'geographicalArea'}, None]
    SAI *SAI
    GeographicalArea *GeographicalArea
} // AreaIdentity

type AssRABParameters struct { // [{'type': 'Ass-RAB-Parameter-MaxBitrateList', 'name': 'assMaxBitrateInf', 'optional': True}, {'type': 'Ass-RAB-Parameter-GuaranteedBitrateList', 'name': 'assGuaranteedBitRateInf', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Ass-RAB-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    AssMaxBitrateInf *AssRABParameterMaxBitrateList
    AssGuaranteedBitRateInf *AssRABParameterGuaranteedBitrateList
    IEExtensions *AssRABParametersExtIEs
}

func (self * AssRABParameters) Unpack(stream *Stream) {
    assMaxBitrateInf_flag := 0x00000002
    assGuaranteedBitRateInf_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (assMaxBitrateInf_flag & _flags) == assMaxBitrateInf_flag { //cond2
        self.AssMaxBitrateInf = &AssRABParameterMaxBitrateList{}//7{'type': 'Ass-RAB-Parameter-MaxBitrateList', 'name': 'assMaxBitrateInf', 'optional': True}
        self.AssMaxBitrateInf.Unpack(stream)// p8
    }
    if (assGuaranteedBitRateInf_flag & _flags) == assGuaranteedBitRateInf_flag { //cond2
        self.AssGuaranteedBitRateInf = &AssRABParameterGuaranteedBitrateList{}//7{'type': 'Ass-RAB-Parameter-GuaranteedBitrateList', 'name': 'assGuaranteedBitRateInf', 'optional': True}
        self.AssGuaranteedBitRateInf.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &AssRABParametersExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Ass-RAB-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_AssRABParametersExtIEs, order_AssRABParametersExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AssRABParameters) Pack(stream *Stream) {
    const assMaxBitrateInf_flag uint = 0x00000002
    const assGuaranteedBitRateInf_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.AssMaxBitrateInf != nil { 
        _flags |= assMaxBitrateInf_flag
        self.AssMaxBitrateInf.Pack(stream)
    }//end of optional
    if self.AssGuaranteedBitRateInf != nil { 
        _flags |= assGuaranteedBitRateInf_flag
        self.AssGuaranteedBitRateInf.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_AssRABParametersExtIEs, order_AssRABParametersExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

func (self *AssRABParameterGuaranteedBitrateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]GuaranteedBitrate, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *AssRABParameterGuaranteedBitrateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type AssRABParameterGuaranteedBitrateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GuaranteedBitrate'}, 'size': [(1, 'maxNrOfSeparateTrafficDirections')]}
    Items []GuaranteedBitrate
}

func (self *AssRABParameterMaxBitrateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]MaxBitrate, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *AssRABParameterMaxBitrateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type AssRABParameterMaxBitrateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MaxBitrate'}, 'size': [(1, 'maxNrOfSeparateTrafficDirections')]}
    Items []MaxBitrate
}

func (self *AuthorisedPLMNs) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]AuthorisedPLMNs_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *AuthorisedPLMNs_Item) { //[{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'AuthorisedSNAs', 'name': 'authorisedSNAsList', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AuthorisedPLMNs-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        authorisedSNAsList_flag := 0x00000002
        iEExtensions_flag := 0x00000004
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(3)
        self.PLMNidentity.Unpack(stream)// p8
        if (authorisedSNAsList_flag & _flags) == authorisedSNAsList_flag { //cond2
            self.AuthorisedSNAsList = &AuthorisedSNAs{}//7{'type': 'AuthorisedSNAs', 'name': 'authorisedSNAsList', 'optional': True}
            self.AuthorisedSNAsList.Unpack(stream)// p8
        }
        if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
            self.IEExtensions = &AuthorisedPLMNsExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AuthorisedPLMNs-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
            IEExtensions := ProtocolExtensionContainer {table_AuthorisedPLMNsExtIEs, order_AuthorisedPLMNsExtIEs} // p3
            IEExtensions.Unpack(stream, &self.IEExtensions) // p3
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    for item := 0; item <_size; item +=1 {
        Unpack_Item(stream, &self.Items[item])
    }
}


func (self *AuthorisedPLMNs) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    var Pack_Item = func(stream *Stream, self AuthorisedPLMNs_Item) {//seq
        const authorisedSNAsList_flag uint = 0x00000002
        const iEExtensions_flag uint = 0x00000004
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(3)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        self.PLMNidentity.Pack(stream)
        if self.AuthorisedSNAsList != nil { 
            _flags |= authorisedSNAsList_flag
            self.AuthorisedSNAsList.Pack(stream)
        }//end of optional
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_AuthorisedPLMNsExtIEs, order_AuthorisedPLMNsExtIEs} // p3
            IEExtensions.Pack(stream, &self.IEExtensions)
        }//end of optional
        if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
        stream.set_flags(_flags, _flagReserve, 3)
    }//end
    for _, item := range self.Items { // seqof base type
        Pack_Item(stream, item)
    }
    return

}


type AuthorisedPLMNs_Item struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'AuthorisedSNAs', 'name': 'authorisedSNAsList', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AuthorisedPLMNs-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNidentity PLMNidentity
    AuthorisedSNAsList *AuthorisedSNAs
    IEExtensions *AuthorisedPLMNsExtIEs
}
type AuthorisedPLMNs struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'AuthorisedSNAs', 'name': 'authorisedSNAsList', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['AuthorisedPLMNs-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxNrOfPLMNsSN')]}
    Items []AuthorisedPLMNs_Item
}

func (self *AuthorisedSNAs) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65536)
    _size += 1
    self.Items = make([]SNAC, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *AuthorisedSNAs) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65536)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type AuthorisedSNAs struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SNAC'}, 'size': [(1, 'maxNrOfSNAs')]}
    Items []SNAC
}

type BindingID struct {
  Value HexBytes
}
func (self *BindingID) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(4)
}
func (self *BindingID) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 4)
}
type BroadcastAssistanceDataDecipheringKeys struct { // [{'type': 'BIT STRING', 'size': [1], 'name': 'cipheringKeyFlag'}, {'type': 'BIT STRING', 'size': [56], 'name': 'currentDecipheringKey'}, {'type': 'BIT STRING', 'size': [56], 'name': 'nextDecipheringKey'}, None]
    CipheringKeyFlag BITSTRING
    CurrentDecipheringKey BITSTRING
    NextDecipheringKey BITSTRING
}

func (self * BroadcastAssistanceDataDecipheringKeys) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_cipheringKeyFlag = func(st *Stream, self *BITSTRING){
        self.Value = st.parsef_BitString(1, 1)
    }
    Unpack_cipheringKeyFlag(stream, &self.CipheringKeyFlag)// p2
    var Unpack_currentDecipheringKey = func(st *Stream, self *BITSTRING){
        self.Value = st.parsef_BitString(56, 56)
    }
    Unpack_currentDecipheringKey(stream, &self.CurrentDecipheringKey)// p2
    var Unpack_nextDecipheringKey = func(st *Stream, self *BITSTRING){
        self.Value = st.parsef_BitString(56, 56)
    }
    Unpack_nextDecipheringKey(stream, &self.NextDecipheringKey)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * BroadcastAssistanceDataDecipheringKeys) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_cipheringKeyFlag = func(st *Stream, self BITSTRING) {
        st.formatf_BitString(self.Value, 1)
    }
    Pack_cipheringKeyFlag(stream, self.CipheringKeyFlag) //f2
    var Pack_currentDecipheringKey = func(st *Stream, self BITSTRING) {
        st.formatf_BitString(self.Value, 56)
    }
    Pack_currentDecipheringKey(stream, self.CurrentDecipheringKey) //f2
    var Pack_nextDecipheringKey = func(st *Stream, self BITSTRING) {
        st.formatf_BitString(self.Value, 56)
    }
    Pack_nextDecipheringKey(stream, self.NextDecipheringKey) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *Cause)Unpack(stream *Stream) {
    //coptions := []string{"radioNetwork","transmissionNetwork","nAS","protocol","misc","non-Standard","Unknown","Unknown"}
    choice := stream.get_choice(3, 1, 6)
    choice_len := 0
    choice_loc := 0
    if choice >= 6 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in Cause\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.RadioNetwork = &CauseRadioNetwork{}//cho6
        self.RadioNetwork.Unpack(stream)
    } else if choice == 1 { //ch2
        self.TransmissionNetwork = &CauseTransmissionNetwork{}//cho6
        self.TransmissionNetwork.Unpack(stream)
    } else if choice == 2 { //ch2
        self.NAS = &CauseNAS{}//cho6
        self.NAS.Unpack(stream)
    } else if choice == 3 { //ch2
        self.Protocol = &CauseProtocol{}//cho6
        self.Protocol.Unpack(stream)
    } else if choice == 4 { //ch2
        self.Misc = &CauseMisc{}//cho6
        self.Misc.Unpack(stream)
    } else if choice == 5 { //ch2
        self.NonStandard = &CauseNonStandard{}//cho6
        self.NonStandard.Unpack(stream)
    }//end of if else

    if choice >= 6 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * Cause) Pack(stream *Stream) {
    if self.RadioNetwork != nil {
        stream.set_choice(0, 3, 1, 6)
        self.RadioNetwork.Pack(stream)//2
    } else if self.TransmissionNetwork != nil {
        stream.set_choice(1, 3, 1, 6)
        self.TransmissionNetwork.Pack(stream)//2
    } else if self.NAS != nil {
        stream.set_choice(2, 3, 1, 6)
        self.NAS.Pack(stream)//2
    } else if self.Protocol != nil {
        stream.set_choice(3, 3, 1, 6)
        self.Protocol.Pack(stream)//2
    } else if self.Misc != nil {
        stream.set_choice(4, 3, 1, 6)
        self.Misc.Pack(stream)//2
    } else if self.NonStandard != nil {
        stream.set_choice(5, 3, 1, 6)
        self.NonStandard.Pack(stream)//2
    }

}
type Cause struct { //[{'type': 'CauseRadioNetwork', 'name': 'radioNetwork'}, {'type': 'CauseTransmissionNetwork', 'name': 'transmissionNetwork'}, {'type': 'CauseNAS', 'name': 'nAS'}, {'type': 'CauseProtocol', 'name': 'protocol'}, {'type': 'CauseMisc', 'name': 'misc'}, {'type': 'CauseNon-Standard', 'name': 'non-Standard'}, None]
    RadioNetwork *CauseRadioNetwork
    TransmissionNetwork *CauseTransmissionNetwork
    NAS *CauseNAS
    Protocol *CauseProtocol
    Misc *CauseMisc
    NonStandard *CauseNonStandard
} // Cause

type CauseMisc struct {
  Value uint64
}
func (self *CauseMisc) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16, 4, 0, 113)
}
func (self * CauseMisc) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16, 4, 0, 113)
}
type CauseNAS struct {
  Value uint64
}
func (self *CauseNAS) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16, 4, 0, 81)
}
func (self * CauseNAS) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16, 4, 0, 81)
}
type CauseProtocol struct {
  Value uint64
}
func (self *CauseProtocol) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16, 4, 0, 97)
}
func (self * CauseProtocol) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16, 4, 0, 97)
}
type CauseRadioNetwork struct {
  Value uint64
}
func (self *CauseRadioNetwork) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(64, 6, 0, 1)
}
func (self * CauseRadioNetwork) Pack(st *Stream){
    st.formatf_Integer(self.Value, 64, 6, 0, 1)
}
type CauseNonStandard struct {
  Value uint64
}
func (self *CauseNonStandard) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(128, 7, 0, 129)
}
func (self * CauseNonStandard) Pack(st *Stream){
    st.formatf_Integer(self.Value, 128, 7, 0, 129)
}
type CauseTransmissionNetwork struct {
  Value uint64
}
func (self *CauseTransmissionNetwork) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16, 4, 0, 65)
}
func (self * CauseTransmissionNetwork) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16, 4, 0, 65)
}
type CellCapacityClassValue struct {
  Value uint64
}
func (self *CellCapacityClassValue) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(100, 8, 1, 1)
}
func (self * CellCapacityClassValue) Pack(st *Stream){
    st.formatf_Integer(self.Value, 100, 8, 1, 1)
}
type CellLoadInformation struct { // [{'type': 'Cell-Capacity-Class-Value', 'name': 'cell-Capacity-Class-Value'}, {'type': 'LoadValue', 'name': 'loadValue'}, {'type': 'RTLoadValue', 'name': 'rTLoadValue', 'optional': True}, {'type': 'NRTLoadInformationValue', 'name': 'nRTLoadInformationValue', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CellLoadInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    CellCapacityClassValue CellCapacityClassValue
    LoadValue LoadValue
    RTLoadValue *RTLoadValue
    NRTLoadInformationValue *NRTLoadInformationValue
    IEExtensions *CellLoadInformationExtIEs
}

func (self * CellLoadInformation) Unpack(stream *Stream) {
    rTLoadValue_flag := 0x00000002
    nRTLoadInformationValue_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.CellCapacityClassValue.Unpack(stream)// p8
    self.LoadValue.Unpack(stream)// p8
    if (rTLoadValue_flag & _flags) == rTLoadValue_flag { //cond2
        self.RTLoadValue = &RTLoadValue{}//7{'type': 'RTLoadValue', 'name': 'rTLoadValue', 'optional': True}
        self.RTLoadValue.Unpack(stream)// p8
    }
    if (nRTLoadInformationValue_flag & _flags) == nRTLoadInformationValue_flag { //cond2
        self.NRTLoadInformationValue = &NRTLoadInformationValue{}//7{'type': 'NRTLoadInformationValue', 'name': 'nRTLoadInformationValue', 'optional': True}
        self.NRTLoadInformationValue.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &CellLoadInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CellLoadInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_CellLoadInformationExtIEs, order_CellLoadInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CellLoadInformation) Pack(stream *Stream) {
    const rTLoadValue_flag uint = 0x00000002
    const nRTLoadInformationValue_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CellCapacityClassValue.Pack(stream)
    self.LoadValue.Pack(stream)
    if self.RTLoadValue != nil { 
        _flags |= rTLoadValue_flag
        self.RTLoadValue.Pack(stream)
    }//end of optional
    if self.NRTLoadInformationValue != nil { 
        _flags |= nRTLoadInformationValue_flag
        self.NRTLoadInformationValue.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_CellLoadInformationExtIEs, order_CellLoadInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type CellLoadInformationGroup struct { // [{'type': 'SourceCellID', 'name': 'sourceCellID'}, {'type': 'CellLoadInformation', 'name': 'uplinkCellLoadInformation', 'optional': True}, {'type': 'CellLoadInformation', 'name': 'downlinkCellLoadInformation', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CellLoadInformationGroup-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    SourceCellID SourceCellID
    UplinkCellLoadInformation *CellLoadInformation
    DownlinkCellLoadInformation *CellLoadInformation
    IEExtensions *CellLoadInformationGroupExtIEs
}

func (self * CellLoadInformationGroup) Unpack(stream *Stream) {
    uplinkCellLoadInformation_flag := 0x00000002
    downlinkCellLoadInformation_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.SourceCellID.Unpack(stream)// p8
    if (uplinkCellLoadInformation_flag & _flags) == uplinkCellLoadInformation_flag { //cond2
        self.UplinkCellLoadInformation = &CellLoadInformation{}//7{'type': 'CellLoadInformation', 'name': 'uplinkCellLoadInformation', 'optional': True}
        self.UplinkCellLoadInformation.Unpack(stream)// p8
    }
    if (downlinkCellLoadInformation_flag & _flags) == downlinkCellLoadInformation_flag { //cond2
        self.DownlinkCellLoadInformation = &CellLoadInformation{}//7{'type': 'CellLoadInformation', 'name': 'downlinkCellLoadInformation', 'optional': True}
        self.DownlinkCellLoadInformation.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &CellLoadInformationGroupExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CellLoadInformationGroup-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_CellLoadInformationGroupExtIEs, order_CellLoadInformationGroupExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CellLoadInformationGroup) Pack(stream *Stream) {
    const uplinkCellLoadInformation_flag uint = 0x00000002
    const downlinkCellLoadInformation_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.SourceCellID.Pack(stream)
    if self.UplinkCellLoadInformation != nil { 
        _flags |= uplinkCellLoadInformation_flag
        self.UplinkCellLoadInformation.Pack(stream)
    }//end of optional
    if self.DownlinkCellLoadInformation != nil { 
        _flags |= downlinkCellLoadInformation_flag
        self.DownlinkCellLoadInformation.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_CellLoadInformationGroupExtIEs, order_CellLoadInformationGroupExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type ClientType struct {
  Value int
}
const (
    ClientTypeemergency_Services = 0
    ClientTypevalue_Added_Services = 1
    ClientTypepLMN_Operator_Services = 2
    ClientTypelawful_Intercept_Services = 3
    ClientTypepLMN_Operator_Broadcast_Services = 4
    ClientTypepLMN_Operator_O_et_M = 5
    ClientTypepLMN_Operator_Anonymous_Statistics = 6
    ClientTypepLMN_Operator_Target_MS_Service_Support = 7

    /* Extensions */
)
func (self *ClientType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 8, 1)
}
func (self *ClientType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 8, 1)
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

type RNMessageStructure struct { // [{'type': 'ProtocolIE-ID', 'name': 'iE-ID'}, {'type': 'RepetitionNumber1', 'name': 'repetitionNumber', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MessageStructure-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    IEID ProtocolIEID
    RepetitionNumber *RepetitionNumber1
    IEExtensions *MessageStructureExtIEs
}

func (self * RNMessageStructure) Unpack(stream *Stream) {
    repetitionNumber_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.IEID.Unpack(stream)// p8
    if (repetitionNumber_flag & _flags) == repetitionNumber_flag { //cond2
        self.RepetitionNumber = &RepetitionNumber1{}//7{'type': 'RepetitionNumber1', 'name': 'repetitionNumber', 'optional': True}
        self.RepetitionNumber.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &MessageStructureExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['MessageStructure-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_MessageStructureExtIEs, order_MessageStructureExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RNMessageStructure) Pack(stream *Stream) {
    const repetitionNumber_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IEID.Pack(stream)
    if self.RepetitionNumber != nil { 
        _flags |= repetitionNumber_flag
        self.RepetitionNumber.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_MessageStructureExtIEs, order_MessageStructureExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

func (self *MessageStructure) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]RNMessageStructure, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MessageStructure) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MessageStructure struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RNMessageStructure'}, 'size': [(1, 'maxNrOfLevels')]}
    Items []RNMessageStructure
}

func (self *CriticalityDiagnosticsIEList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]CriticalityDiagnosticsIEList_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *CriticalityDiagnosticsIEList_Item) { //[{'type': 'Criticality', 'name': 'iECriticality'}, {'type': 'ProtocolIE-ID', 'name': 'iE-ID'}, {'type': 'RepetitionNumber0', 'name': 'repetitionNumber', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-IE-List-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        repetitionNumber_flag := 0x00000002
        iEExtensions_flag := 0x00000004
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(3)
        self.IECriticality.Unpack(stream)// p8
        self.IEID.Unpack(stream)// p8
        if (repetitionNumber_flag & _flags) == repetitionNumber_flag { //cond2
            self.RepetitionNumber = &RepetitionNumber0{}//7{'type': 'RepetitionNumber0', 'name': 'repetitionNumber', 'optional': True}
            self.RepetitionNumber.Unpack(stream)// p8
        }
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
        const repetitionNumber_flag uint = 0x00000002
        const iEExtensions_flag uint = 0x00000004
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(3)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        self.IECriticality.Pack(stream)
        self.IEID.Pack(stream)
        if self.RepetitionNumber != nil { 
            _flags |= repetitionNumber_flag
            self.RepetitionNumber.Pack(stream)
        }//end of optional
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_CriticalityDiagnosticsIEListExtIEs, order_CriticalityDiagnosticsIEListExtIEs} // p3
            IEExtensions.Pack(stream, &self.IEExtensions)
        }//end of optional
        if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
        stream.set_flags(_flags, _flagReserve, 3)
    }//end
    for _, item := range self.Items { // seqof base type
        Pack_Item(stream, item)
    }
    return

}


type CriticalityDiagnosticsIEList_Item struct { // [{'type': 'Criticality', 'name': 'iECriticality'}, {'type': 'ProtocolIE-ID', 'name': 'iE-ID'}, {'type': 'RepetitionNumber0', 'name': 'repetitionNumber', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-IE-List-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    IECriticality Criticality
    IEID ProtocolIEID
    RepetitionNumber *RepetitionNumber0
    IEExtensions *CriticalityDiagnosticsIEListExtIEs
}
type CriticalityDiagnosticsIEList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'Criticality', 'name': 'iECriticality'}, {'type': 'ProtocolIE-ID', 'name': 'iE-ID'}, {'type': 'RepetitionNumber0', 'name': 'repetitionNumber', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['CriticalityDiagnostics-IE-List-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxNrOfErrors')]}
    Items []CriticalityDiagnosticsIEList_Item
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

type ChosenEncryptionAlgorithm struct{ //{'type': 'EncryptionAlgorithm'}
    EncryptionAlgorithm //UserTypeTODO: Unimplimented 
}
func (self *ChosenEncryptionAlgorithm) Unpack(st *Stream){ // ut3 'EncryptionAlgorithm'
    self.EncryptionAlgorithm.Unpack(st) // ut32
}

func (self *ChosenEncryptionAlgorithm) Pack(st *Stream){
    self.EncryptionAlgorithm.Pack(st) // ut32
}

type ChosenIntegrityProtectionAlgorithm struct{ //{'type': 'IntegrityProtectionAlgorithm'}
    IntegrityProtectionAlgorithm //UserTypeTODO: Unimplimented 
}
func (self *ChosenIntegrityProtectionAlgorithm) Unpack(st *Stream){ // ut3 'IntegrityProtectionAlgorithm'
    self.IntegrityProtectionAlgorithm.Unpack(st) // ut32
}

func (self *ChosenIntegrityProtectionAlgorithm) Pack(st *Stream){
    self.IntegrityProtectionAlgorithm.Pack(st) // ut32
}

type CI struct {
  Value HexBytes
}
func (self *CI) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(2)
}
func (self *CI) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 2)
}
type ClassmarkInformation2 struct {
  Value HexBytes
}
func (self *ClassmarkInformation2) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *ClassmarkInformation2) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type ClassmarkInformation3 struct {
  Value HexBytes
}
func (self *ClassmarkInformation3) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *ClassmarkInformation3) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
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
type CNID struct {
  Value uint64
}
func (self *CNID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4096, 12, 0, 0)
}
func (self * CNID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4096, 12, 0, 0)
}
type DataVolumeReference struct {
  Value uint64
}
func (self *DataVolumeReference) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * DataVolumeReference) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
type DataVolumeReportingIndication struct {
  Value int
}
const (
    DataVolumeReportingIndicationdo_report = 0
    DataVolumeReportingIndicationdo_not_report = 1
)
func (self *DataVolumeReportingIndication) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 2, 0)
}
func (self *DataVolumeReportingIndication) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 2, 0)
}
type DCHID struct {
  Value uint64
}
func (self *DCHID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * DCHID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
type DeliveryOfErroneousSDU struct {
  Value int
}
const (
    DeliveryOfErroneousSDUyes = 0
    DeliveryOfErroneousSDUno = 1
    DeliveryOfErroneousSDUno_error_detection_consideration = 2
)
func (self *DeliveryOfErroneousSDU) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 3, 0)
}
func (self *DeliveryOfErroneousSDU) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 3, 0)
}
type DeliveryOrder struct {
  Value int
}
const (
    DeliveryOrderdelivery_order_requested = 0
    DeliveryOrderdelivery_order_not_requested = 1
)
func (self *DeliveryOrder) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 2, 0)
}
func (self *DeliveryOrder) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 2, 0)
}
type DLGTPPDUSequenceNumber struct {
  Value uint64
}
func (self *DLGTPPDUSequenceNumber) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * DLGTPPDUSequenceNumber) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type DLNPDUSequenceNumber struct {
  Value uint64
}
func (self *DLNPDUSequenceNumber) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * DLNPDUSequenceNumber) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type DRNTI struct {
  Value uint64
}
func (self *DRNTI) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(1048576, 20, 0, 0)
}
func (self * DRNTI) Pack(st *Stream){
    st.formatf_Integer(self.Value, 1048576, 20, 0, 0)
}
type DRXCycleLengthCoefficient struct {
  Value uint64
}
func (self *DRXCycleLengthCoefficient) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4, 2, 0, 6)
}
func (self * DRXCycleLengthCoefficient) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4, 2, 0, 6)
}
type DSCHID struct {
  Value uint64
}
func (self *DSCHID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * DSCHID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
type EncryptionAlgorithm struct {
  Value uint64
}
func (self *EncryptionAlgorithm) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16, 4, 0, 0)
}
func (self * EncryptionAlgorithm) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16, 4, 0, 0)
}
type EncryptionInformation struct { // [{'type': 'PermittedEncryptionAlgorithms', 'name': 'permittedAlgorithms'}, {'type': 'EncryptionKey', 'name': 'key'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['EncryptionInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PermittedAlgorithms PermittedEncryptionAlgorithms
    Key EncryptionKey
    IEExtensions *EncryptionInformationExtIEs
}

func (self * EncryptionInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.PermittedAlgorithms.Unpack(stream)// p8
    self.Key.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &EncryptionInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['EncryptionInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_EncryptionInformationExtIEs, order_EncryptionInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * EncryptionInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PermittedAlgorithms.Pack(stream)
    self.Key.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_EncryptionInformationExtIEs, order_EncryptionInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EncryptionKey struct {
  Len int
  Value HexBytes
}
func (self *EncryptionKey) Unpack(st *Stream){
    self.Value = st.parsef_BitString(128, 128)
}
func (self *EncryptionKey) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 128)
}
func (self *EquipmentsToBeTraced)Unpack(stream *Stream) {
    //coptions := []string{"iMEIlist","iMEISVlist","iMEIgroup","iMEISVgroup"}
    choice := stream.get_choice(2, 1, 4)
    choice_len := 0
    choice_loc := 0
    if choice >= 4 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in EquipmentsToBeTraced\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.IMEIlist = &IMEIList{}//cho6
        self.IMEIlist.Unpack(stream)
    } else if choice == 1 { //ch2
        self.IMEISVlist = &IMEISVList{}//cho6
        self.IMEISVlist.Unpack(stream)
    } else if choice == 2 { //ch2
        self.IMEIgroup = &IMEIGroup{}//cho6
        self.IMEIgroup.Unpack(stream)
    } else if choice == 3 { //ch2
        self.IMEISVgroup = &IMEISVGroup{}//cho6
        self.IMEISVgroup.Unpack(stream)
    }//end of if else

    if choice >= 4 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * EquipmentsToBeTraced) Pack(stream *Stream) {
    if self.IMEIlist != nil {
        stream.set_choice(0, 2, 1, 4)
        self.IMEIlist.Pack(stream)//2
    } else if self.IMEISVlist != nil {
        stream.set_choice(1, 2, 1, 4)
        self.IMEISVlist.Pack(stream)//2
    } else if self.IMEIgroup != nil {
        stream.set_choice(2, 2, 1, 4)
        self.IMEIgroup.Pack(stream)//2
    } else if self.IMEISVgroup != nil {
        stream.set_choice(3, 2, 1, 4)
        self.IMEISVgroup.Pack(stream)//2
    }

}
type EquipmentsToBeTraced struct { //[{'type': 'IMEIList', 'name': 'iMEIlist'}, {'type': 'IMEISVList', 'name': 'iMEISVlist'}, {'type': 'IMEIGroup', 'name': 'iMEIgroup'}, {'type': 'IMEISVGroup', 'name': 'iMEISVgroup'}, None]
    IMEIlist *IMEIList
    IMEISVlist *IMEISVList
    IMEIgroup *IMEIGroup
    IMEISVgroup *IMEISVGroup
} // EquipmentsToBeTraced

type Event struct {
  Value int
}
const (
    Eventstop_change_of_service_area = 0
    Eventdirect = 1
    Eventchange_of_servicearea = 2

    /* Extensions */
    Eventstop_direct = 3
)
func (self *Event) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *Event) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
func (self *GeographicalArea)Unpack(stream *Stream) {
    //coptions := []string{"point","pointWithUnCertainty","polygon","Unknown","pointWithUncertaintyEllipse","pointWithAltitude","pointWithAltitudeAndUncertaintyEllipsoid","ellipsoidArc"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in GeographicalArea\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.Point = &GAPoint{}//cho6
        self.Point.Unpack(stream)
    } else if choice == 1 { //ch2
        self.PointWithUnCertainty = &GAPointWithUnCertainty{}//cho6
        self.PointWithUnCertainty.Unpack(stream)
    } else if choice == 2 { //ch2
        self.Polygon = &GAPolygon{}//cho6
        self.Polygon.Unpack(stream)
    } else if choice == 4 { //ch2
        self.PointWithUncertaintyEllipse = &GAPointWithUnCertaintyEllipse{}//cho6
        self.PointWithUncertaintyEllipse.Unpack(stream)
    } else if choice == 5 { //ch2
        self.PointWithAltitude = &GAPointWithAltitude{}//cho6
        self.PointWithAltitude.Unpack(stream)
    } else if choice == 6 { //ch2
        self.PointWithAltitudeAndUncertaintyEllipsoid = &GAPointWithAltitudeAndUncertaintyEllipsoid{}//cho6
        self.PointWithAltitudeAndUncertaintyEllipsoid.Unpack(stream)
    } else if choice == 7 { //ch2
        self.EllipsoidArc = &GAEllipsoidArc{}//cho6
        self.EllipsoidArc.Unpack(stream)
    }//end of if else

    if choice >= 3 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * GeographicalArea) Pack(stream *Stream) {
    if self.Point != nil {
        stream.set_choice(0, 2, 1, 3)
        self.Point.Pack(stream)//2
    } else if self.PointWithUnCertainty != nil {
        stream.set_choice(1, 2, 1, 3)
        self.PointWithUnCertainty.Pack(stream)//2
    } else if self.Polygon != nil {
        stream.set_choice(2, 2, 1, 3)
        self.Polygon.Pack(stream)//2
    } else if self.PointWithUncertaintyEllipse != nil {
        stream.set_choice(4, 2, 1, 3)
        lenLoc := stream.reserve_len()
        self.PointWithUncertaintyEllipse.Pack(stream)//2
        stream.set_len(lenLoc)
    } else if self.PointWithAltitude != nil {
        stream.set_choice(5, 2, 1, 3)
        lenLoc := stream.reserve_len()
        self.PointWithAltitude.Pack(stream)//2
        stream.set_len(lenLoc)
    } else if self.PointWithAltitudeAndUncertaintyEllipsoid != nil {
        stream.set_choice(6, 2, 1, 3)
        lenLoc := stream.reserve_len()
        self.PointWithAltitudeAndUncertaintyEllipsoid.Pack(stream)//2
        stream.set_len(lenLoc)
    } else if self.EllipsoidArc != nil {
        stream.set_choice(7, 2, 1, 3)
        lenLoc := stream.reserve_len()
        self.EllipsoidArc.Pack(stream)//2
        stream.set_len(lenLoc)
    }

}
type GeographicalArea struct { //[{'type': 'GA-Point', 'name': 'point'}, {'type': 'GA-PointWithUnCertainty', 'name': 'pointWithUnCertainty'}, {'type': 'GA-Polygon', 'name': 'polygon'}, None, {'type': 'GA-PointWithUnCertaintyEllipse', 'name': 'pointWithUncertaintyEllipse'}, {'type': 'GA-PointWithAltitude', 'name': 'pointWithAltitude'}, {'type': 'GA-PointWithAltitudeAndUncertaintyEllipsoid', 'name': 'pointWithAltitudeAndUncertaintyEllipsoid'}, {'type': 'GA-EllipsoidArc', 'name': 'ellipsoidArc'}]
    Point *GAPoint
    PointWithUnCertainty *GAPointWithUnCertainty
    Polygon *GAPolygon
    PointWithUncertaintyEllipse *GAPointWithUnCertaintyEllipse
    PointWithAltitude *GAPointWithAltitude
    PointWithAltitudeAndUncertaintyEllipsoid *GAPointWithAltitudeAndUncertaintyEllipsoid
    EllipsoidArc *GAEllipsoidArc
} // GeographicalArea

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

type GAAltitudeAndDirection struct { // [{'type': 'ENUMERATED', 'values': [('height', 0), ('depth', 1)], 'name': 'directionOfAltitude'}, {'type': 'INTEGER', 'restricted-to': [(0, 32767)], 'name': 'altitude'}, None]
    DirectionOfAltitude ENUMERATED
    Altitude INTEGER
}

func (self * GAAltitudeAndDirection) Unpack(stream *Stream) {
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

func (self * GAAltitudeAndDirection) Pack(stream *Stream) {
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

type GAEllipsoidArc struct { // [{'type': 'GeographicalCoordinates', 'name': 'geographicalCoordinates'}, {'type': 'INTEGER', 'restricted-to': [(0, 65535)], 'name': 'innerRadius'}, {'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'uncertaintyRadius'}, {'type': 'INTEGER', 'restricted-to': [(0, 179)], 'name': 'offsetAngle'}, {'type': 'INTEGER', 'restricted-to': [(0, 179)], 'name': 'includedAngle'}, {'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'confidence'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-EllipsoidArc-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    GeographicalCoordinates GeographicalCoordinates
    InnerRadius INTEGER
    UncertaintyRadius INTEGER
    OffsetAngle INTEGER
    IncludedAngle INTEGER
    Confidence INTEGER
    IEExtensions *GAEllipsoidArcExtIEs
}

func (self * GAEllipsoidArc) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.GeographicalCoordinates.Unpack(stream)// p8
    var Unpack_innerRadius = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(65536, 16, 0, 0)
    }
    Unpack_innerRadius(stream, &self.InnerRadius)// p2
    var Unpack_uncertaintyRadius = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(128, 7, 0, 0)
    }
    Unpack_uncertaintyRadius(stream, &self.UncertaintyRadius)// p2
    var Unpack_offsetAngle = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(180, 8, 0, 0)
    }
    Unpack_offsetAngle(stream, &self.OffsetAngle)// p2
    var Unpack_includedAngle = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(180, 8, 0, 0)
    }
    Unpack_includedAngle(stream, &self.IncludedAngle)// p2
    var Unpack_confidence = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(128, 7, 0, 0)
    }
    Unpack_confidence(stream, &self.Confidence)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GAEllipsoidArcExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-EllipsoidArc-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GAEllipsoidArcExtIEs, order_GAEllipsoidArcExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GAEllipsoidArc) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GeographicalCoordinates.Pack(stream)
    var Pack_innerRadius = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 65536, 16, 0, 0)
    }
    Pack_innerRadius(stream, self.InnerRadius) //f2
    var Pack_uncertaintyRadius = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 128, 7, 0, 0)
    }
    Pack_uncertaintyRadius(stream, self.UncertaintyRadius) //f2
    var Pack_offsetAngle = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 180, 8, 0, 0)
    }
    Pack_offsetAngle(stream, self.OffsetAngle) //f2
    var Pack_includedAngle = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 180, 8, 0, 0)
    }
    Pack_includedAngle(stream, self.IncludedAngle) //f2
    var Pack_confidence = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 128, 7, 0, 0)
    }
    Pack_confidence(stream, self.Confidence) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GAEllipsoidArcExtIEs, order_GAEllipsoidArcExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GAPoint struct { // [{'type': 'GeographicalCoordinates', 'name': 'geographicalCoordinates'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-Point-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    GeographicalCoordinates GeographicalCoordinates
    IEExtensions *GAPointExtIEs
}

func (self * GAPoint) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.GeographicalCoordinates.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GAPointExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-Point-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GAPointExtIEs, order_GAPointExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GAPoint) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GeographicalCoordinates.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GAPointExtIEs, order_GAPointExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GAPointWithAltitude struct { // [{'type': 'GeographicalCoordinates', 'name': 'geographicalCoordinates'}, {'type': 'GA-AltitudeAndDirection', 'name': 'altitudeAndDirection'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-PointWithAltitude-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    GeographicalCoordinates GeographicalCoordinates
    AltitudeAndDirection GAAltitudeAndDirection
    IEExtensions *GAPointWithAltitudeExtIEs
}

func (self * GAPointWithAltitude) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.GeographicalCoordinates.Unpack(stream)// p8
    self.AltitudeAndDirection.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GAPointWithAltitudeExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-PointWithAltitude-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GAPointWithAltitudeExtIEs, order_GAPointWithAltitudeExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GAPointWithAltitude) Pack(stream *Stream) {
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
        IEExtensions := &ProtocolExtensionContainer {table_GAPointWithAltitudeExtIEs, order_GAPointWithAltitudeExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GAPointWithAltitudeAndUncertaintyEllipsoid struct { // [{'type': 'GeographicalCoordinates', 'name': 'geographicalCoordinates'}, {'type': 'GA-AltitudeAndDirection', 'name': 'altitudeAndDirection'}, {'type': 'GA-UncertaintyEllipse', 'name': 'uncertaintyEllipse'}, {'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'uncertaintyAltitude'}, {'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'confidence'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-PointWithAltitudeAndUncertaintyEllipsoid-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    GeographicalCoordinates GeographicalCoordinates
    AltitudeAndDirection GAAltitudeAndDirection
    UncertaintyEllipse GAUncertaintyEllipse
    UncertaintyAltitude INTEGER
    Confidence INTEGER
    IEExtensions *GAPointWithAltitudeAndUncertaintyEllipsoidExtIEs
}

func (self * GAPointWithAltitudeAndUncertaintyEllipsoid) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.GeographicalCoordinates.Unpack(stream)// p8
    self.AltitudeAndDirection.Unpack(stream)// p8
    self.UncertaintyEllipse.Unpack(stream)// p8
    var Unpack_uncertaintyAltitude = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(128, 7, 0, 0)
    }
    Unpack_uncertaintyAltitude(stream, &self.UncertaintyAltitude)// p2
    var Unpack_confidence = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(128, 7, 0, 0)
    }
    Unpack_confidence(stream, &self.Confidence)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GAPointWithAltitudeAndUncertaintyEllipsoidExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-PointWithAltitudeAndUncertaintyEllipsoid-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GAPointWithAltitudeAndUncertaintyEllipsoidExtIEs, order_GAPointWithAltitudeAndUncertaintyEllipsoidExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GAPointWithAltitudeAndUncertaintyEllipsoid) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GeographicalCoordinates.Pack(stream)
    self.AltitudeAndDirection.Pack(stream)
    self.UncertaintyEllipse.Pack(stream)
    var Pack_uncertaintyAltitude = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 128, 7, 0, 0)
    }
    Pack_uncertaintyAltitude(stream, self.UncertaintyAltitude) //f2
    var Pack_confidence = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 128, 7, 0, 0)
    }
    Pack_confidence(stream, self.Confidence) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GAPointWithAltitudeAndUncertaintyEllipsoidExtIEs, order_GAPointWithAltitudeAndUncertaintyEllipsoidExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type GAPointWithUnCertainty struct { // [{'type': 'GeographicalCoordinates', 'name': 'geographicalCoordinates'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-PointWithUnCertainty-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, {'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'uncertaintyCode'}]
    GeographicalCoordinates GeographicalCoordinates
    IEExtensions *GAPointWithUnCertaintyExtIEs
    UncertaintyCode INTEGER
}

func (self * GAPointWithUnCertainty) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.GeographicalCoordinates.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GAPointWithUnCertaintyExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-PointWithUnCertainty-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GAPointWithUnCertaintyExtIEs, order_GAPointWithUnCertaintyExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    var Unpack_uncertaintyCode = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(128, 7, 0, 0)
    }
    Unpack_uncertaintyCode(stream, &self.UncertaintyCode)// p2
    return
}

func (self * GAPointWithUnCertainty) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GeographicalCoordinates.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GAPointWithUnCertaintyExtIEs, order_GAPointWithUnCertaintyExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    var Pack_uncertaintyCode = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 128, 7, 0, 0)
    }
    Pack_uncertaintyCode(stream, self.UncertaintyCode) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GAPointWithUnCertaintyEllipse struct { // [{'type': 'GeographicalCoordinates', 'name': 'geographicalCoordinates'}, {'type': 'GA-UncertaintyEllipse', 'name': 'uncertaintyEllipse'}, {'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'confidence'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-PointWithUnCertaintyEllipse-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    GeographicalCoordinates GeographicalCoordinates
    UncertaintyEllipse GAUncertaintyEllipse
    Confidence INTEGER
    IEExtensions *GAPointWithUnCertaintyEllipseExtIEs
}

func (self * GAPointWithUnCertaintyEllipse) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.GeographicalCoordinates.Unpack(stream)// p8
    self.UncertaintyEllipse.Unpack(stream)// p8
    var Unpack_confidence = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(128, 7, 0, 0)
    }
    Unpack_confidence(stream, &self.Confidence)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GAPointWithUnCertaintyEllipseExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-PointWithUnCertaintyEllipse-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GAPointWithUnCertaintyEllipseExtIEs, order_GAPointWithUnCertaintyEllipseExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GAPointWithUnCertaintyEllipse) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GeographicalCoordinates.Pack(stream)
    self.UncertaintyEllipse.Pack(stream)
    var Pack_confidence = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 128, 7, 0, 0)
    }
    Pack_confidence(stream, self.Confidence) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GAPointWithUnCertaintyEllipseExtIEs, order_GAPointWithUnCertaintyEllipseExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *GAPolygon) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(15)
    _size += 1
    self.Items = make([]GAPolygon_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *GAPolygon_Item) { //[{'type': 'GeographicalCoordinates', 'name': 'geographicalCoordinates'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-Polygon-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        iEExtensions_flag := 0x00000002
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(2)
        self.GeographicalCoordinates.Unpack(stream)// p8
        if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
            self.IEExtensions = &GAPolygonExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-Polygon-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
            IEExtensions := ProtocolExtensionContainer {table_GAPolygonExtIEs, order_GAPolygonExtIEs} // p3
            IEExtensions.Unpack(stream, &self.IEExtensions) // p3
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    for item := 0; item <_size; item +=1 {
        Unpack_Item(stream, &self.Items[item])
    }
}


func (self *GAPolygon) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 15)
    var Pack_Item = func(stream *Stream, self GAPolygon_Item) {//seq
        const iEExtensions_flag uint = 0x00000002
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(2)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        self.GeographicalCoordinates.Pack(stream)
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_GAPolygonExtIEs, order_GAPolygonExtIEs} // p3
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


type GAPolygon_Item struct { // [{'type': 'GeographicalCoordinates', 'name': 'geographicalCoordinates'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-Polygon-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    GeographicalCoordinates GeographicalCoordinates
    IEExtensions *GAPolygonExtIEs
}
type GAPolygon struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'GeographicalCoordinates', 'name': 'geographicalCoordinates'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GA-Polygon-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxNrOfPoints')]}
    Items []GAPolygon_Item
}

type GAUncertaintyEllipse struct { // [{'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'uncertaintySemi-major'}, {'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'uncertaintySemi-minor'}, {'type': 'INTEGER', 'restricted-to': [(0, 179)], 'name': 'orientationOfMajorAxis'}, None]
    UncertaintySemimajor INTEGER
    UncertaintySemiminor INTEGER
    OrientationOfMajorAxis INTEGER
}

func (self * GAUncertaintyEllipse) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_uncertaintySemimajor = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(128, 7, 0, 0)
    }
    Unpack_uncertaintySemimajor(stream, &self.UncertaintySemimajor)// p2
    var Unpack_uncertaintySemiminor = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(128, 7, 0, 0)
    }
    Unpack_uncertaintySemiminor(stream, &self.UncertaintySemiminor)// p2
    var Unpack_orientationOfMajorAxis = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(180, 8, 0, 0)
    }
    Unpack_orientationOfMajorAxis(stream, &self.OrientationOfMajorAxis)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GAUncertaintyEllipse) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_uncertaintySemimajor = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 128, 7, 0, 0)
    }
    Pack_uncertaintySemimajor(stream, self.UncertaintySemimajor) //f2
    var Pack_uncertaintySemiminor = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 128, 7, 0, 0)
    }
    Pack_uncertaintySemiminor(stream, self.UncertaintySemiminor) //f2
    var Pack_orientationOfMajorAxis = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 180, 8, 0, 0)
    }
    Pack_orientationOfMajorAxis(stream, self.OrientationOfMajorAxis) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GERANBSCContainer struct {
  Value HexBytes
}
func (self *GERANBSCContainer) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *GERANBSCContainer) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type GERANCellID struct { // [{'type': 'LAI', 'name': 'lAI'}, {'type': 'RAC', 'name': 'rAC'}, {'type': 'CI', 'name': 'cI'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GERAN-Cell-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    LAI LAI
    RAC RAC
    CI CI
    IEExtensions *GERANCellIDExtIEs
}

func (self * GERANCellID) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.LAI.Unpack(stream)// p8
    self.RAC.Unpack(stream)// p8
    self.CI.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &GERANCellIDExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['GERAN-Cell-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_GERANCellIDExtIEs, order_GERANCellIDExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * GERANCellID) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.LAI.Pack(stream)
    self.RAC.Pack(stream)
    self.CI.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_GERANCellIDExtIEs, order_GERANCellIDExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GERANClassmark struct {
  Value HexBytes
}
func (self *GERANClassmark) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *GERANClassmark) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type GlobalCNID struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'CN-ID', 'name': 'cN-ID'}]
    PLMNidentity PLMNidentity
    CNID CNID
}

func (self * GlobalCNID) Unpack(stream *Stream) {
    self.PLMNidentity.Unpack(stream)// p8
    self.CNID.Unpack(stream)// p8
    return
}

func (self * GlobalCNID) Pack(stream *Stream) {
    self.PLMNidentity.Pack(stream)
    self.CNID.Pack(stream)
}//end

type GlobalRNCID struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'RNC-ID', 'name': 'rNC-ID'}]
    PLMNidentity PLMNidentity
    RNCID RNCID
}

func (self * GlobalRNCID) Unpack(stream *Stream) {
    self.PLMNidentity.Unpack(stream)// p8
    self.RNCID.Unpack(stream)// p8
    return
}

func (self * GlobalRNCID) Pack(stream *Stream) {
    self.PLMNidentity.Pack(stream)
    self.RNCID.Pack(stream)
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
type GuaranteedBitrate struct {
  Value uint64
}
func (self *GuaranteedBitrate) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16000001, 24, 0, 0)
}
func (self * GuaranteedBitrate) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16000001, 24, 0, 0)
}
type HSDSCHMACdFlowID struct {
  Value uint64
}
func (self *HSDSCHMACdFlowID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(8, 3, 0, 0)
}
func (self * HSDSCHMACdFlowID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 8, 3, 0, 0)
}
type IMEI struct {
  Value HexBytes
}
func (self *IMEI) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(8)
}
func (self *IMEI) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 8)
}
type IMEIGroup struct { // [{'type': 'IMEI', 'name': 'iMEI'}, {'type': 'BIT STRING', 'size': [7], 'name': 'iMEIMask'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['IMEIGroup-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    IMEI IMEI
    IMEIMask BITSTRING
    IEExtensions *IMEIGroupExtIEs
}

func (self * IMEIGroup) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.IMEI.Unpack(stream)// p8
    var Unpack_iMEIMask = func(st *Stream, self *BITSTRING){
        self.Value = st.parsef_BitString(7, 7)
    }
    Unpack_iMEIMask(stream, &self.IMEIMask)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &IMEIGroupExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['IMEIGroup-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_IMEIGroupExtIEs, order_IMEIGroupExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * IMEIGroup) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IMEI.Pack(stream)
    var Pack_iMEIMask = func(st *Stream, self BITSTRING) {
        st.formatf_BitString(self.Value, 7)
    }
    Pack_iMEIMask(stream, self.IMEIMask) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_IMEIGroupExtIEs, order_IMEIGroupExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *IMEIList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(64)
    _size += 1
    self.Items = make([]IMEI, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *IMEIList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 64)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type IMEIList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'IMEI'}, 'size': [(1, 'maxNrOfUEsToBeTraced')]}
    Items []IMEI
}

type IMEISV struct {
  Value HexBytes
}
func (self *IMEISV) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(8)
}
func (self *IMEISV) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 8)
}
type IMEISVGroup struct { // [{'type': 'IMEISV', 'name': 'iMEISV'}, {'type': 'BIT STRING', 'size': [7], 'name': 'iMEISVMask'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['IMEISVGroup-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    IMEISV IMEISV
    IMEISVMask BITSTRING
    IEExtensions *IMEISVGroupExtIEs
}

func (self * IMEISVGroup) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.IMEISV.Unpack(stream)// p8
    var Unpack_iMEISVMask = func(st *Stream, self *BITSTRING){
        self.Value = st.parsef_BitString(7, 7)
    }
    Unpack_iMEISVMask(stream, &self.IMEISVMask)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &IMEISVGroupExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['IMEISVGroup-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_IMEISVGroupExtIEs, order_IMEISVGroupExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * IMEISVGroup) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IMEISV.Pack(stream)
    var Pack_iMEISVMask = func(st *Stream, self BITSTRING) {
        st.formatf_BitString(self.Value, 7)
    }
    Pack_iMEISVMask(stream, self.IMEISVMask) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_IMEISVGroupExtIEs, order_IMEISVGroupExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *IMEISVList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(64)
    _size += 1
    self.Items = make([]IMEISV, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *IMEISVList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 64)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type IMEISVList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'IMEISV'}, 'size': [(1, 'maxNrOfUEsToBeTraced')]}
    Items []IMEISV
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
type InformationTransferID struct {
  Value uint64
}
func (self *InformationTransferID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(1048576, 20, 0, 0)
}
func (self * InformationTransferID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 1048576, 20, 0, 0)
}
func (self *InformationTransferType)Unpack(stream *Stream) {
    //coptions := []string{"rNCTraceInformation"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in InformationTransferType\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.RNCTraceInformation = &RNCTraceInformation{}//cho6
        self.RNCTraceInformation.Unpack(stream)
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * InformationTransferType) Pack(stream *Stream) {
    if self.RNCTraceInformation != nil {
        stream.set_choice(0, 0, 1, 1)
        self.RNCTraceInformation.Pack(stream)//2
    }

}
type InformationTransferType struct { //[{'type': 'RNCTraceInformation', 'name': 'rNCTraceInformation'}, None]
    RNCTraceInformation *RNCTraceInformation
} // InformationTransferType

type IntegrityProtectionAlgorithm struct {
  Value uint64
}
func (self *IntegrityProtectionAlgorithm) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16, 4, 0, 0)
}
func (self * IntegrityProtectionAlgorithm) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16, 4, 0, 0)
}
type IntegrityProtectionInformation struct { // [{'type': 'PermittedIntegrityProtectionAlgorithms', 'name': 'permittedAlgorithms'}, {'type': 'IntegrityProtectionKey', 'name': 'key'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['IntegrityProtectionInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PermittedAlgorithms PermittedIntegrityProtectionAlgorithms
    Key IntegrityProtectionKey
    IEExtensions *IntegrityProtectionInformationExtIEs
}

func (self * IntegrityProtectionInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.PermittedAlgorithms.Unpack(stream)// p8
    self.Key.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &IntegrityProtectionInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['IntegrityProtectionInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_IntegrityProtectionInformationExtIEs, order_IntegrityProtectionInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * IntegrityProtectionInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PermittedAlgorithms.Pack(stream)
    self.Key.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_IntegrityProtectionInformationExtIEs, order_IntegrityProtectionInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type IntegrityProtectionKey struct {
  Len int
  Value HexBytes
}
func (self *IntegrityProtectionKey) Unpack(st *Stream){
    self.Value = st.parsef_BitString(128, 128)
}
func (self *IntegrityProtectionKey) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 128)
}
func (self *InterSystemInformationTransferType)Unpack(stream *Stream) {
    //coptions := []string{"rIM-Transfer"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in InterSystemInformationTransferType\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.RIMTransfer = &RIMTransfer{}//cho6
        self.RIMTransfer.Unpack(stream)
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * InterSystemInformationTransferType) Pack(stream *Stream) {
    if self.RIMTransfer != nil {
        stream.set_choice(0, 0, 1, 1)
        self.RIMTransfer.Pack(stream)//2
    }

}
type InterSystemInformationTransferType struct { //[{'type': 'RIM-Transfer', 'name': 'rIM-Transfer'}, None]
    RIMTransfer *RIMTransfer
} // InterSystemInformationTransferType

type InterSystemInformationTransparentContainer struct { // [{'type': 'CellLoadInformation', 'name': 'downlinkCellLoadInformation', 'optional': True}, {'type': 'CellLoadInformation', 'name': 'uplinkCellLoadInformation', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InterSystemInformation-TransparentContainer-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DownlinkCellLoadInformation *CellLoadInformation
    UplinkCellLoadInformation *CellLoadInformation
    IEExtensions *InterSystemInformationTransparentContainerExtIEs
}

func (self * InterSystemInformationTransparentContainer) Unpack(stream *Stream) {
    downlinkCellLoadInformation_flag := 0x00000002
    uplinkCellLoadInformation_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (downlinkCellLoadInformation_flag & _flags) == downlinkCellLoadInformation_flag { //cond2
        self.DownlinkCellLoadInformation = &CellLoadInformation{}//7{'type': 'CellLoadInformation', 'name': 'downlinkCellLoadInformation', 'optional': True}
        self.DownlinkCellLoadInformation.Unpack(stream)// p8
    }
    if (uplinkCellLoadInformation_flag & _flags) == uplinkCellLoadInformation_flag { //cond2
        self.UplinkCellLoadInformation = &CellLoadInformation{}//7{'type': 'CellLoadInformation', 'name': 'uplinkCellLoadInformation', 'optional': True}
        self.UplinkCellLoadInformation.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &InterSystemInformationTransparentContainerExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InterSystemInformation-TransparentContainer-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_InterSystemInformationTransparentContainerExtIEs, order_InterSystemInformationTransparentContainerExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InterSystemInformationTransparentContainer) Pack(stream *Stream) {
    const downlinkCellLoadInformation_flag uint = 0x00000002
    const uplinkCellLoadInformation_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.DownlinkCellLoadInformation != nil { 
        _flags |= downlinkCellLoadInformation_flag
        self.DownlinkCellLoadInformation.Pack(stream)
    }//end of optional
    if self.UplinkCellLoadInformation != nil { 
        _flags |= uplinkCellLoadInformation_flag
        self.UplinkCellLoadInformation.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_InterSystemInformationTransparentContainerExtIEs, order_InterSystemInformationTransparentContainerExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type IuSignallingConnectionIdentifier struct {
  Len int
  Value HexBytes
}
func (self *IuSignallingConnectionIdentifier) Unpack(st *Stream){
    self.Value = st.parsef_BitString(24, 24)
}
func (self *IuSignallingConnectionIdentifier) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 24)
}
func (self *IuTransportAssociation)Unpack(stream *Stream) {
    //coptions := []string{"gTP-TEI","bindingID"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in IuTransportAssociation\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.GTPTEI = &GTPTEI{}//cho6
        self.GTPTEI.Unpack(stream)
    } else if choice == 1 { //ch2
        self.BindingID = &BindingID{}//cho6
        self.BindingID.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * IuTransportAssociation) Pack(stream *Stream) {
    if self.GTPTEI != nil {
        stream.set_choice(0, 1, 1, 2)
        self.GTPTEI.Pack(stream)//2
    } else if self.BindingID != nil {
        stream.set_choice(1, 1, 1, 2)
        self.BindingID.Pack(stream)//2
    }

}
type IuTransportAssociation struct { //[{'type': 'GTP-TEI', 'name': 'gTP-TEI'}, {'type': 'BindingID', 'name': 'bindingID'}, None]
    GTPTEI *GTPTEI
    BindingID *BindingID
} // IuTransportAssociation

type KeyStatus struct {
  Value int
}
const (
    KeyStatusold = 0
    KeyStatusnEw = 1

    /* Extensions */
)
func (self *KeyStatus) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *KeyStatus) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
func (self *LALIST) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65536)
    _size += 1
    self.Items = make([]LALIST_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *LALIST_Item) { //[{'type': 'LAC', 'name': 'lAC'}, {'type': 'ListOF-SNAs', 'name': 'listOF-SNAs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LA-LIST-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        iEExtensions_flag := 0x00000002
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(2)
        self.LAC.Unpack(stream)// p8
        self.ListOFSNAs.Unpack(stream)// p8
        if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
            self.IEExtensions = &LALISTExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LA-LIST-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
            IEExtensions := ProtocolExtensionContainer {table_LALISTExtIEs, order_LALISTExtIEs} // p3
            IEExtensions.Unpack(stream, &self.IEExtensions) // p3
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    for item := 0; item <_size; item +=1 {
        Unpack_Item(stream, &self.Items[item])
    }
}


func (self *LALIST) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65536)
    var Pack_Item = func(stream *Stream, self LALIST_Item) {//seq
        const iEExtensions_flag uint = 0x00000002
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(2)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        self.LAC.Pack(stream)
        self.ListOFSNAs.Pack(stream)
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_LALISTExtIEs, order_LALISTExtIEs} // p3
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


type LALIST_Item struct { // [{'type': 'LAC', 'name': 'lAC'}, {'type': 'ListOF-SNAs', 'name': 'listOF-SNAs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LA-LIST-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    LAC LAC
    ListOFSNAs ListOFSNAs
    IEExtensions *LALISTExtIEs
}
type LALIST struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'LAC', 'name': 'lAC'}, {'type': 'ListOF-SNAs', 'name': 'listOF-SNAs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LA-LIST-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxNrOfLAs')]}
    Items []LALIST_Item
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
type LAI struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'LAC', 'name': 'lAC'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LAI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PLMNidentity PLMNidentity
    LAC LAC
    IEExtensions *LAIExtIEs
}

func (self * LAI) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.PLMNidentity.Unpack(stream)// p8
    self.LAC.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &LAIExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LAI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_LAIExtIEs, order_LAIExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * LAI) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNidentity.Pack(stream)
    self.LAC.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_LAIExtIEs, order_LAIExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type LastKnownServiceArea struct { // [{'type': 'SAI', 'name': 'sAI'}, {'type': 'INTEGER', 'restricted-to': [(0, 32767)], 'name': 'ageOfSAI'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LastKnownServiceArea-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    SAI SAI
    AgeOfSAI INTEGER
    IEExtensions *LastKnownServiceAreaExtIEs
}

func (self * LastKnownServiceArea) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.SAI.Unpack(stream)// p8
    var Unpack_ageOfSAI = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(32768, 15, 0, 0)
    }
    Unpack_ageOfSAI(stream, &self.AgeOfSAI)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &LastKnownServiceAreaExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['LastKnownServiceArea-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_LastKnownServiceAreaExtIEs, order_LastKnownServiceAreaExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * LastKnownServiceArea) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.SAI.Pack(stream)
    var Pack_ageOfSAI = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 32768, 15, 0, 0)
    }
    Pack_ageOfSAI(stream, self.AgeOfSAI) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_LastKnownServiceAreaExtIEs, order_LastKnownServiceAreaExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *ListOFSNAs) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65536)
    _size += 1
    self.Items = make([]SNAC, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *ListOFSNAs) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65536)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type ListOFSNAs struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SNAC'}, 'size': [(1, 'maxNrOfSNAs')]}
    Items []SNAC
}

func (self *ListOfInterfacesToTrace) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(16)
    _size += 1
    self.Items = make([]InterfacesToTraceItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *ListOfInterfacesToTrace) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 16)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type ListOfInterfacesToTrace struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'InterfacesToTraceItem'}, 'size': [(1, 'maxNrOfInterfaces')]}
    Items []InterfacesToTraceItem
}

type InterfacesToTraceItem struct { // [{'type': 'ENUMERATED', 'values': [('iu-cs', 0), ('iu-ps', 1), ('iur', 2), ('iub', 3), ('uu', 4), None], 'name': 'interface'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InterfacesToTraceItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    Interface ENUMERATED
    IEExtensions *InterfacesToTraceItemExtIEs
}

func (self * InterfacesToTraceItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_interface = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(4, 5, 1)
    }
    Unpack_interface(stream, &self.Interface)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &InterfacesToTraceItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['InterfacesToTraceItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_InterfacesToTraceItemExtIEs, order_InterfacesToTraceItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InterfacesToTraceItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_interface = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 4, 5, 1)
    }
    Pack_interface(stream, self.Interface) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_InterfacesToTraceItemExtIEs, order_InterfacesToTraceItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type LoadValue struct {
  Value uint64
}
func (self *LoadValue) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(101, 7, 0, 0)
}
func (self * LoadValue) Pack(st *Stream){
    st.formatf_Integer(self.Value, 101, 7, 0, 0)
}
type LocationRelatedDataRequestType struct { // [{'type': 'RequestedLocationRelatedDataType', 'name': 'requestedLocationRelatedDataType'}, {'type': 'RequestedGPSAssistanceData', 'name': 'requestedGPSAssistanceData', 'optional': True}, None]
    RequestedLocationRelatedDataType RequestedLocationRelatedDataType
    RequestedGPSAssistanceData *RequestedGPSAssistanceData
}

func (self * LocationRelatedDataRequestType) Unpack(stream *Stream) {
    requestedGPSAssistanceData_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RequestedLocationRelatedDataType.Unpack(stream)// p8
    if (requestedGPSAssistanceData_flag & _flags) == requestedGPSAssistanceData_flag { //cond2
        self.RequestedGPSAssistanceData = &RequestedGPSAssistanceData{}//7{'type': 'RequestedGPSAssistanceData', 'name': 'requestedGPSAssistanceData', 'optional': True}
        self.RequestedGPSAssistanceData.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * LocationRelatedDataRequestType) Pack(stream *Stream) {
    const requestedGPSAssistanceData_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RequestedLocationRelatedDataType.Pack(stream)
    if self.RequestedGPSAssistanceData != nil { 
        _flags |= requestedGPSAssistanceData_flag
        self.RequestedGPSAssistanceData.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type LocationRelatedDataRequestTypeSpecificToGERANIuMode struct {
  Value int
}
const (
    LocationRelatedDataRequestTypeSpecificToGERANIuModedecipheringKeysEOTD = 0
    LocationRelatedDataRequestTypeSpecificToGERANIuModededicatedMobileAssistedEOTDAssistanceData = 1
    LocationRelatedDataRequestTypeSpecificToGERANIuModededicatedMobileBasedEOTDAssistanceData = 2

    /* Extensions */
)
func (self *LocationRelatedDataRequestTypeSpecificToGERANIuMode) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *LocationRelatedDataRequestTypeSpecificToGERANIuMode) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type L3Information struct {
  Value HexBytes
}
func (self *L3Information) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *L3Information) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type MaxBitrate struct {
  Value uint64
}
func (self *MaxBitrate) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16000000, 24, 0, 1)
}
func (self * MaxBitrate) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16000000, 24, 0, 1)
}
type MaxSDUSize struct {
  Value uint64
}
func (self *MaxSDUSize) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(32769, 16, 0, 0)
}
func (self * MaxSDUSize) Pack(st *Stream){
    st.formatf_Integer(self.Value, 32769, 16, 0, 0)
}
type NASPDU struct {
  Value HexBytes
}
func (self *NASPDU) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *NASPDU) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type NASSynchronisationIndicator struct {
  Len int
  Value HexBytes
}
func (self *NASSynchronisationIndicator) Unpack(st *Stream){
    self.Value = st.parsef_BitString(4, 4)
}
func (self *NASSynchronisationIndicator) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 4)
}
type NewBSSToOldBSSInformation struct {
  Value HexBytes
}
func (self *NewBSSToOldBSSInformation) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *NewBSSToOldBSSInformation) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type NonSearchingIndication struct {
  Value int
}
const (
    NonSearchingIndicationnon_searching = 0
    NonSearchingIndicationsearching = 1
)
func (self *NonSearchingIndication) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 2, 0)
}
func (self *NonSearchingIndication) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 2, 0)
}
type NRTLoadInformationValue struct {
  Value uint64
}
func (self *NRTLoadInformationValue) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4, 2, 0, 0)
}
func (self * NRTLoadInformationValue) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4, 2, 0, 0)
}
type NumberOfIuInstances struct {
  Value uint64
}
func (self *NumberOfIuInstances) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(2, 1, 0, 1)
}
func (self * NumberOfIuInstances) Pack(st *Stream){
    st.formatf_Integer(self.Value, 2, 1, 0, 1)
}
type NumberOfSteps struct {
  Value uint64
}
func (self *NumberOfSteps) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16, 4, 0, 1)
}
func (self * NumberOfSteps) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16, 4, 0, 1)
}
type OldBSSToNewBSSInformation struct {
  Value HexBytes
}
func (self *OldBSSToNewBSSInformation) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *OldBSSToNewBSSInformation) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type OMCID struct {
  Value HexBytes
}
func (self *OMCID) Unpack(st *Stream) {
    _len := st.parse_olen(5)+3
    if _len < 3 || _len > 22 {
        //fmt.Println ("Invalid len in OMC-ID")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *OMCID) Pack(st *Stream) {
    if len(self.Value) < 3 || len(self.Value) > 22 {
        log.Error ("Invalid len in OMC-ID")
        return
}
    st.format_olen((len(self.Value))-3, 5)
    st.formatf_OctString(self.Value, 0)
}
func (self *PagingAreaID)Unpack(stream *Stream) {
    //coptions := []string{"lAI","rAI"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in PagingAreaID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.LAI = &LAI{}//cho6
        self.LAI.Unpack(stream)
    } else if choice == 1 { //ch2
        self.RAI = &RAI{}//cho6
        self.RAI.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * PagingAreaID) Pack(stream *Stream) {
    if self.LAI != nil {
        stream.set_choice(0, 1, 1, 2)
        self.LAI.Pack(stream)//2
    } else if self.RAI != nil {
        stream.set_choice(1, 1, 1, 2)
        self.RAI.Pack(stream)//2
    }

}
type PagingAreaID struct { //[{'type': 'LAI', 'name': 'lAI'}, {'type': 'RAI', 'name': 'rAI'}, None]
    LAI *LAI
    RAI *RAI
} // PagingAreaID

type PagingCause struct {
  Value int
}
const (
    PagingCauseterminating_conversational_call = 0
    PagingCauseterminating_streaming_call = 1
    PagingCauseterminating_interactive_call = 2
    PagingCauseterminating_background_call = 3
    PagingCauseterminating_low_priority_signalling = 4

    /* Extensions */
    PagingCauseterminating_high_priority_signalling = 5
)
func (self *PagingCause) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 5, 1)
}
func (self *PagingCause) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 5, 1)
}
func (self *PDPTypeInformation) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]PDPType, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PDPTypeInformation) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PDPTypeInformation struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PDP-Type'}, 'size': [(1, 'maxNrOfPDPDirections')]}
    Items []PDPType
}

type PDPType struct {
  Value int
}
const (
    PDPTypeempty = 0
    PDPTypeppp = 1
    PDPTypeosp_ihoss = 2
    PDPTypeipv4 = 3
    PDPTypeipv6 = 4

    /* Extensions */
)
func (self *PDPType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 5, 1)
}
func (self *PDPType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 5, 1)
}
func (self *PermanentNASUEID)Unpack(stream *Stream) {
    //coptions := []string{"iMSI"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in PermanentNASUEID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.IMSI = &IMSI{}//cho6
        self.IMSI.Unpack(stream)
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * PermanentNASUEID) Pack(stream *Stream) {
    if self.IMSI != nil {
        stream.set_choice(0, 0, 1, 1)
        self.IMSI.Pack(stream)//2
    }

}
type PermanentNASUEID struct { //[{'type': 'IMSI', 'name': 'iMSI'}, None]
    IMSI *IMSI
} // PermanentNASUEID

func (self *PermittedEncryptionAlgorithms) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(16)
    _size += 1
    self.Items = make([]EncryptionAlgorithm, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PermittedEncryptionAlgorithms) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 16)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PermittedEncryptionAlgorithms struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'EncryptionAlgorithm'}, 'size': [(1, 16)]}
    Items []EncryptionAlgorithm
}

func (self *PermittedIntegrityProtectionAlgorithms) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(16)
    _size += 1
    self.Items = make([]IntegrityProtectionAlgorithm, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PermittedIntegrityProtectionAlgorithms) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 16)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PermittedIntegrityProtectionAlgorithms struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'IntegrityProtectionAlgorithm'}, 'size': [(1, 16)]}
    Items []IntegrityProtectionAlgorithm
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
func (self *PLMNsinsharednetwork) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]PLMNsinsharednetwork_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *PLMNsinsharednetwork_Item) { //[{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'LA-LIST', 'name': 'lA-LIST'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PLMNs-in-shared-network-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        iEExtensions_flag := 0x00000002
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(2)
        self.PLMNidentity.Unpack(stream)// p8
        self.LALIST.Unpack(stream)// p8
        if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
            self.IEExtensions = &PLMNsinsharednetworkExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PLMNs-in-shared-network-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
            IEExtensions := ProtocolExtensionContainer {table_PLMNsinsharednetworkExtIEs, order_PLMNsinsharednetworkExtIEs} // p3
            IEExtensions.Unpack(stream, &self.IEExtensions) // p3
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    for item := 0; item <_size; item +=1 {
        Unpack_Item(stream, &self.Items[item])
    }
}


func (self *PLMNsinsharednetwork) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    var Pack_Item = func(stream *Stream, self PLMNsinsharednetwork_Item) {//seq
        const iEExtensions_flag uint = 0x00000002
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(2)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        self.PLMNidentity.Pack(stream)
        self.LALIST.Pack(stream)
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_PLMNsinsharednetworkExtIEs, order_PLMNsinsharednetworkExtIEs} // p3
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


type PLMNsinsharednetwork_Item struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'LA-LIST', 'name': 'lA-LIST'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PLMNs-in-shared-network-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNidentity PLMNidentity
    LALIST LALIST
    IEExtensions *PLMNsinsharednetworkExtIEs
}
type PLMNsinsharednetwork struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'LA-LIST', 'name': 'lA-LIST'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PLMNs-in-shared-network-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxNrOfPLMNsSN')]}
    Items []PLMNsinsharednetwork_Item
}

type PositioningDataDiscriminator struct {
  Len int
  Value HexBytes
}
func (self *PositioningDataDiscriminator) Unpack(st *Stream){
    self.Value = st.parsef_BitString(4, 4)
}
func (self *PositioningDataDiscriminator) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 4)
}
func (self *PositioningDataSet) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(9)
    _size += 1
    self.Items = make([]PositioningMethodAndUsage, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *PositioningDataSet) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 9)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type PositioningDataSet struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PositioningMethodAndUsage'}, 'size': [(1, 'maxSet')]}
    Items []PositioningMethodAndUsage
}

type PositioningMethodAndUsage struct {
  Value HexBytes
}
func (self *PositioningMethodAndUsage) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(1)
}
func (self *PositioningMethodAndUsage) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 1)
}
type PositioningPriority struct {
  Value int
}
const (
    PositioningPriorityhigh_Priority = 0
    PositioningPrioritynormal_Priority = 1

    /* Extensions */
)
func (self *PositioningPriority) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *PositioningPriority) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type PositionData struct { // [{'type': 'PositioningDataDiscriminator', 'name': 'positioningDataDiscriminator'}, {'type': 'PositioningDataSet', 'name': 'positioningDataSet', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PositionData-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PositioningDataDiscriminator PositioningDataDiscriminator
    PositioningDataSet *PositioningDataSet
    IEExtensions *PositionDataExtIEs
}

func (self * PositionData) Unpack(stream *Stream) {
    positioningDataSet_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.PositioningDataDiscriminator.Unpack(stream)// p8
    if (positioningDataSet_flag & _flags) == positioningDataSet_flag { //cond2
        self.PositioningDataSet = &PositioningDataSet{}//7{'type': 'PositioningDataSet', 'name': 'positioningDataSet', 'optional': True}
        self.PositioningDataSet.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &PositionDataExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['PositionData-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_PositionDataExtIEs, order_PositionDataExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PositionData) Pack(stream *Stream) {
    const positioningDataSet_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PositioningDataDiscriminator.Pack(stream)
    if self.PositioningDataSet != nil { 
        _flags |= positioningDataSet_flag
        self.PositioningDataSet.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_PositionDataExtIEs, order_PositionDataExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type PositionDataSpecificToGERANIuMode struct {
  Value HexBytes
}
func (self *PositionDataSpecificToGERANIuMode) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *PositionDataSpecificToGERANIuMode) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
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
func (self *ProvidedData)Unpack(stream *Stream) {
    //coptions := []string{"shared-network-information"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in ProvidedData\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.Sharednetworkinformation = &SharedNetworkInformation{}//cho6
        self.Sharednetworkinformation.Unpack(stream)
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ProvidedData) Pack(stream *Stream) {
    if self.Sharednetworkinformation != nil {
        stream.set_choice(0, 0, 1, 1)
        self.Sharednetworkinformation.Pack(stream)//2
    }

}
type ProvidedData struct { //[{'type': 'Shared-Network-Information', 'name': 'shared-network-information'}, None]
    Sharednetworkinformation *SharedNetworkInformation
} // ProvidedData

type PTMSI struct {
  Value HexBytes
}
func (self *PTMSI) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(4)
}
func (self *PTMSI) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 4)
}
type QueuingAllowed struct {
  Value int
}
const (
    QueuingAllowedqueueing_not_allowed = 0
    QueuingAllowedqueueing_allowed = 1
)
func (self *QueuingAllowed) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 2, 0)
}
func (self *QueuingAllowed) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 2, 0)
}
type RABAsymmetryIndicator struct {
  Value int
}
const (
    RABAsymmetryIndicatorsymmetric_bidirectional = 0
    RABAsymmetryIndicatorasymmetric_unidirectional_downlink = 1
    RABAsymmetryIndicatorasymmetric_unidirectional_uplink = 2
    RABAsymmetryIndicatorasymmetric_bidirectional = 3

    /* Extensions */
)
func (self *RABAsymmetryIndicator) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 4, 1)
}
func (self *RABAsymmetryIndicator) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 4, 1)
}
type RABID struct {
  Len int
  Value HexBytes
}
func (self *RABID) Unpack(st *Stream){
    self.Value = st.parsef_BitString(8, 8)
}
func (self *RABID) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 8)
}
func (self *RABParameterGuaranteedBitrateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]GuaranteedBitrate, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *RABParameterGuaranteedBitrateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type RABParameterGuaranteedBitrateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GuaranteedBitrate'}, 'size': [(1, 'maxNrOfSeparateTrafficDirections')]}
    Items []GuaranteedBitrate
}

func (self *RABParameterMaxBitrateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]MaxBitrate, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *RABParameterMaxBitrateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type RABParameterMaxBitrateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MaxBitrate'}, 'size': [(1, 'maxNrOfSeparateTrafficDirections')]}
    Items []MaxBitrate
}

type RABParameters struct { // [{'type': 'TrafficClass', 'name': 'trafficClass'}, {'type': 'RAB-AsymmetryIndicator', 'name': 'rAB-AsymmetryIndicator'}, {'type': 'RAB-Parameter-MaxBitrateList', 'name': 'maxBitrate'}, {'type': 'RAB-Parameter-GuaranteedBitrateList', 'name': 'guaranteedBitRate', 'optional': True}, {'type': 'DeliveryOrder', 'name': 'deliveryOrder'}, {'type': 'MaxSDU-Size', 'name': 'maxSDU-Size'}, {'type': 'SDU-Parameters', 'name': 'sDU-Parameters'}, {'type': 'TransferDelay', 'name': 'transferDelay', 'optional': True}, {'type': 'TrafficHandlingPriority', 'name': 'trafficHandlingPriority', 'optional': True}, {'type': 'AllocationOrRetentionPriority', 'name': 'allocationOrRetentionPriority', 'optional': True}, {'type': 'SourceStatisticsDescriptor', 'name': 'sourceStatisticsDescriptor', 'optional': True}, {'type': 'RelocationRequirement', 'name': 'relocationRequirement', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    TrafficClass TrafficClass
    RABAsymmetryIndicator RABAsymmetryIndicator
    MaxBitrate RABParameterMaxBitrateList
    GuaranteedBitRate *RABParameterGuaranteedBitrateList
    DeliveryOrder DeliveryOrder
    MaxSDUSize MaxSDUSize
    SDUParameters SDUParameters
    TransferDelay *TransferDelay
    TrafficHandlingPriority *TrafficHandlingPriority
    AllocationOrRetentionPriority *AllocationOrRetentionPriority
    SourceStatisticsDescriptor *SourceStatisticsDescriptor
    RelocationRequirement *RelocationRequirement
    IEExtensions *RABParametersExtIEs
}

func (self * RABParameters) Unpack(stream *Stream) {
    guaranteedBitRate_flag := 0x00000002
    transferDelay_flag := 0x00000004
    trafficHandlingPriority_flag := 0x00000008
    allocationOrRetentionPriority_flag := 0x00000010
    sourceStatisticsDescriptor_flag := 0x00000020
    relocationRequirement_flag := 0x00000040
    iEExtensions_flag := 0x00000080
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(8)
    self.TrafficClass.Unpack(stream)// p8
    self.RABAsymmetryIndicator.Unpack(stream)// p8
    self.MaxBitrate.Unpack(stream)// p8
    if (guaranteedBitRate_flag & _flags) == guaranteedBitRate_flag { //cond2
        self.GuaranteedBitRate = &RABParameterGuaranteedBitrateList{}//7{'type': 'RAB-Parameter-GuaranteedBitrateList', 'name': 'guaranteedBitRate', 'optional': True}
        self.GuaranteedBitRate.Unpack(stream)// p8
    }
    self.DeliveryOrder.Unpack(stream)// p8
    self.MaxSDUSize.Unpack(stream)// p8
    self.SDUParameters.Unpack(stream)// p8
    if (transferDelay_flag & _flags) == transferDelay_flag { //cond2
        self.TransferDelay = &TransferDelay{}//7{'type': 'TransferDelay', 'name': 'transferDelay', 'optional': True}
        self.TransferDelay.Unpack(stream)// p8
    }
    if (trafficHandlingPriority_flag & _flags) == trafficHandlingPriority_flag { //cond2
        self.TrafficHandlingPriority = &TrafficHandlingPriority{}//7{'type': 'TrafficHandlingPriority', 'name': 'trafficHandlingPriority', 'optional': True}
        self.TrafficHandlingPriority.Unpack(stream)// p8
    }
    if (allocationOrRetentionPriority_flag & _flags) == allocationOrRetentionPriority_flag { //cond2
        self.AllocationOrRetentionPriority = &AllocationOrRetentionPriority{}//7{'type': 'AllocationOrRetentionPriority', 'name': 'allocationOrRetentionPriority', 'optional': True}
        self.AllocationOrRetentionPriority.Unpack(stream)// p8
    }
    if (sourceStatisticsDescriptor_flag & _flags) == sourceStatisticsDescriptor_flag { //cond2
        self.SourceStatisticsDescriptor = &SourceStatisticsDescriptor{}//7{'type': 'SourceStatisticsDescriptor', 'name': 'sourceStatisticsDescriptor', 'optional': True}
        self.SourceStatisticsDescriptor.Unpack(stream)// p8
    }
    if (relocationRequirement_flag & _flags) == relocationRequirement_flag { //cond2
        self.RelocationRequirement = &RelocationRequirement{}//7{'type': 'RelocationRequirement', 'name': 'relocationRequirement', 'optional': True}
        self.RelocationRequirement.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABParametersExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABParametersExtIEs, order_RABParametersExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABParameters) Pack(stream *Stream) {
    const guaranteedBitRate_flag uint = 0x00000002
    const transferDelay_flag uint = 0x00000004
    const trafficHandlingPriority_flag uint = 0x00000008
    const allocationOrRetentionPriority_flag uint = 0x00000010
    const sourceStatisticsDescriptor_flag uint = 0x00000020
    const relocationRequirement_flag uint = 0x00000040
    const iEExtensions_flag uint = 0x00000080
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(8)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TrafficClass.Pack(stream)
    self.RABAsymmetryIndicator.Pack(stream)
    self.MaxBitrate.Pack(stream)
    if self.GuaranteedBitRate != nil { 
        _flags |= guaranteedBitRate_flag
        self.GuaranteedBitRate.Pack(stream)
    }//end of optional
    self.DeliveryOrder.Pack(stream)
    self.MaxSDUSize.Pack(stream)
    self.SDUParameters.Pack(stream)
    if self.TransferDelay != nil { 
        _flags |= transferDelay_flag
        self.TransferDelay.Pack(stream)
    }//end of optional
    if self.TrafficHandlingPriority != nil { 
        _flags |= trafficHandlingPriority_flag
        self.TrafficHandlingPriority.Pack(stream)
    }//end of optional
    if self.AllocationOrRetentionPriority != nil { 
        _flags |= allocationOrRetentionPriority_flag
        self.AllocationOrRetentionPriority.Pack(stream)
    }//end of optional
    if self.SourceStatisticsDescriptor != nil { 
        _flags |= sourceStatisticsDescriptor_flag
        self.SourceStatisticsDescriptor.Pack(stream)
    }//end of optional
    if self.RelocationRequirement != nil { 
        _flags |= relocationRequirement_flag
        self.RelocationRequirement.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABParametersExtIEs, order_RABParametersExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 8)
}//end

type RABSubflowCombinationBitRate struct {
  Value uint64
}
func (self *RABSubflowCombinationBitRate) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16000001, 24, 0, 0)
}
func (self * RABSubflowCombinationBitRate) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16000001, 24, 0, 0)
}
func (self *RABTrCHMapping) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(256)
    _size += 1
    self.Items = make([]RABTrCHMappingItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *RABTrCHMapping) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 256)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type RABTrCHMapping struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RAB-TrCH-MappingItem'}, 'size': [(1, 'maxNrOfRABs')]}
    Items []RABTrCHMappingItem
}

type RABTrCHMappingItem struct { // [{'type': 'RAB-ID', 'name': 'rAB-ID'}, {'type': 'TrCH-ID-List', 'name': 'trCH-ID-List'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-TrCH-MappingItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RABID RABID
    TrCHIDList TrCHIDList
    IEExtensions *RABTrCHMappingItemExtIEs
}

func (self * RABTrCHMappingItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RABID.Unpack(stream)// p8
    self.TrCHIDList.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RABTrCHMappingItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAB-TrCH-MappingItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RABTrCHMappingItemExtIEs, order_RABTrCHMappingItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RABTrCHMappingItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RABID.Pack(stream)
    self.TrCHIDList.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RABTrCHMappingItemExtIEs, order_RABTrCHMappingItemExtIEs} // p3
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
type RAI struct { // [{'type': 'LAI', 'name': 'lAI'}, {'type': 'RAC', 'name': 'rAC'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    LAI LAI
    RAC RAC
    IEExtensions *RAIExtIEs
}

func (self * RAI) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.LAI.Unpack(stream)// p8
    self.RAC.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RAIExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RAI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RAIExtIEs, order_RAIExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RAI) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.LAI.Pack(stream)
    self.RAC.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RAIExtIEs, order_RAIExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RateControlAllowed struct {
  Value int
}
const (
    RateControlAllowednot_allowed = 0
    RateControlAllowedallowed = 1
)
func (self *RateControlAllowed) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 2, 0)
}
func (self *RateControlAllowed) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 2, 0)
}
type RelocationRequirement struct {
  Value int
}
const (
    RelocationRequirementlossless = 0
    RelocationRequirementnone = 1
)
func (self *RelocationRequirement) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 2, 0)
}
func (self *RelocationRequirement) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 2, 0)
}
type RelocationType struct {
  Value int
}
const (
    RelocationTypeue_not_involved = 0
    RelocationTypeue_involved = 1

    /* Extensions */
)
func (self *RelocationType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *RelocationType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type RepetitionNumber0 struct {
  Value uint64
}
func (self *RepetitionNumber0) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * RepetitionNumber0) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
type RepetitionNumber1 struct {
  Value uint64
}
func (self *RepetitionNumber1) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 1)
}
func (self * RepetitionNumber1) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 1)
}
type ReportArea struct {
  Value int
}
const (
    ReportAreaservice_area = 0
    ReportAreageographical_area = 1

    /* Extensions */
)
func (self *ReportArea) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *ReportArea) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type RequestedGPSAssistanceData struct {
  Value HexBytes
}
func (self *RequestedGPSAssistanceData) Unpack(st *Stream) {
    _len := st.parse_olen(6)+1
    if _len < 1 || _len > 38 {
        //fmt.Println ("Invalid len in RequestedGPSAssistanceData")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *RequestedGPSAssistanceData) Pack(st *Stream) {
    if len(self.Value) < 1 || len(self.Value) > 38 {
        log.Error ("Invalid len in RequestedGPSAssistanceData")
        return
}
    st.format_olen((len(self.Value))-1, 6)
    st.formatf_OctString(self.Value, 0)
}
type RequestedLocationRelatedDataType struct {
  Value int
}
const (
    RequestedLocationRelatedDataTypedecipheringKeysUEBasedOTDOA = 0
    RequestedLocationRelatedDataTypedecipheringKeysAssistedGPS = 1
    RequestedLocationRelatedDataTypededicatedAssistanceDataUEBasedOTDOA = 2
    RequestedLocationRelatedDataTypededicatedAssistanceDataAssistedGPS = 3

    /* Extensions */
)
func (self *RequestedLocationRelatedDataType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 4, 1)
}
func (self *RequestedLocationRelatedDataType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 4, 1)
}
type RequestedRABParameterValues struct { // [{'type': 'Requested-RAB-Parameter-MaxBitrateList', 'name': 'requestedMaxBitrates', 'optional': True}, {'type': 'Requested-RAB-Parameter-GuaranteedBitrateList', 'name': 'requestedGuaranteedBitrates', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Requested-RAB-Parameter-Values-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RequestedMaxBitrates *RequestedRABParameterMaxBitrateList
    RequestedGuaranteedBitrates *RequestedRABParameterGuaranteedBitrateList
    IEExtensions *RequestedRABParameterValuesExtIEs
}

func (self * RequestedRABParameterValues) Unpack(stream *Stream) {
    requestedMaxBitrates_flag := 0x00000002
    requestedGuaranteedBitrates_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (requestedMaxBitrates_flag & _flags) == requestedMaxBitrates_flag { //cond2
        self.RequestedMaxBitrates = &RequestedRABParameterMaxBitrateList{}//7{'type': 'Requested-RAB-Parameter-MaxBitrateList', 'name': 'requestedMaxBitrates', 'optional': True}
        self.RequestedMaxBitrates.Unpack(stream)// p8
    }
    if (requestedGuaranteedBitrates_flag & _flags) == requestedGuaranteedBitrates_flag { //cond2
        self.RequestedGuaranteedBitrates = &RequestedRABParameterGuaranteedBitrateList{}//7{'type': 'Requested-RAB-Parameter-GuaranteedBitrateList', 'name': 'requestedGuaranteedBitrates', 'optional': True}
        self.RequestedGuaranteedBitrates.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RequestedRABParameterValuesExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Requested-RAB-Parameter-Values-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RequestedRABParameterValuesExtIEs, order_RequestedRABParameterValuesExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RequestedRABParameterValues) Pack(stream *Stream) {
    const requestedMaxBitrates_flag uint = 0x00000002
    const requestedGuaranteedBitrates_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.RequestedMaxBitrates != nil { 
        _flags |= requestedMaxBitrates_flag
        self.RequestedMaxBitrates.Pack(stream)
    }//end of optional
    if self.RequestedGuaranteedBitrates != nil { 
        _flags |= requestedGuaranteedBitrates_flag
        self.RequestedGuaranteedBitrates.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RequestedRABParameterValuesExtIEs, order_RequestedRABParameterValuesExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

func (self *RequestedRABParameterMaxBitrateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]MaxBitrate, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *RequestedRABParameterMaxBitrateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type RequestedRABParameterMaxBitrateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MaxBitrate'}, 'size': [(1, 'maxNrOfSeparateTrafficDirections')]}
    Items []MaxBitrate
}

func (self *RequestedRABParameterGuaranteedBitrateList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2)
    _size += 1
    self.Items = make([]GuaranteedBitrate, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *RequestedRABParameterGuaranteedBitrateList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 2)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type RequestedRABParameterGuaranteedBitrateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'GuaranteedBitrate'}, 'size': [(1, 'maxNrOfSeparateTrafficDirections')]}
    Items []GuaranteedBitrate
}

type RequestType struct { // [{'type': 'Event', 'name': 'event'}, {'type': 'ReportArea', 'name': 'reportArea'}, {'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'accuracyCode', 'optional': True}, None]
    Event Event
    ReportArea ReportArea
    AccuracyCode *INTEGER
}

func (self * RequestType) Unpack(stream *Stream) {
    accuracyCode_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.Event.Unpack(stream)// p8
    self.ReportArea.Unpack(stream)// p8
    if (accuracyCode_flag & _flags) == accuracyCode_flag { //cond1
        var Unpack_accuracyCode = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(128, 7, 0, 0)
        }
        self.AccuracyCode = &INTEGER{}//6{'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'accuracyCode', 'optional': True}
        Unpack_accuracyCode(stream, self.AccuracyCode)// p1 {'type': 'INTEGER', 'restricted-to': [(0, 127)], 'name': 'accuracyCode', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RequestType) Pack(stream *Stream) {
    const accuracyCode_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.Event.Pack(stream)
    self.ReportArea.Pack(stream)
    if self.AccuracyCode != nil { //YY
        _flags |= accuracyCode_flag
        var Pack_accuracyCode = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 128, 7, 0, 0)
        }
        Pack_accuracyCode(stream, *self.AccuracyCode) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ResidualBitErrorRatio struct { // [{'type': 'INTEGER', 'restricted-to': [(1, 9)], 'name': 'mantissa'}, {'type': 'INTEGER', 'restricted-to': [(1, 8)], 'name': 'exponent'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResidualBitErrorRatio-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    Mantissa INTEGER
    Exponent INTEGER
    IEExtensions *ResidualBitErrorRatioExtIEs
}

func (self * ResidualBitErrorRatio) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    var Unpack_mantissa = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(9, 4, 0, 1)
    }
    Unpack_mantissa(stream, &self.Mantissa)// p2
    var Unpack_exponent = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(8, 3, 0, 1)
    }
    Unpack_exponent(stream, &self.Exponent)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &ResidualBitErrorRatioExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['ResidualBitErrorRatio-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_ResidualBitErrorRatioExtIEs, order_ResidualBitErrorRatioExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * ResidualBitErrorRatio) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_mantissa = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 9, 4, 0, 1)
    }
    Pack_mantissa(stream, self.Mantissa) //f2
    var Pack_exponent = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 8, 3, 0, 1)
    }
    Pack_exponent(stream, self.Exponent) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_ResidualBitErrorRatioExtIEs, order_ResidualBitErrorRatioExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type ResponseTime struct {
  Value int
}
const (
    ResponseTimelowdelay = 0
    ResponseTimedelaytolerant = 1

    /* Extensions */
)
func (self *ResponseTime) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *ResponseTime) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type RIMInformation struct {
  Value HexBytes
}
func (self *RIMInformation) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RIMInformation) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RIMTransfer struct { // [{'type': 'RIMInformation', 'name': 'rIMInformation'}, {'type': 'RIMRoutingAddress', 'name': 'rIMRoutingAddress', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RIM-Transfer-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    RIMInformation RIMInformation
    RIMRoutingAddress *RIMRoutingAddress
    IEExtensions *RIMTransferExtIEs
}

func (self * RIMTransfer) Unpack(stream *Stream) {
    rIMRoutingAddress_flag := 0x00000001
    iEExtensions_flag := 0x00000002
    _flags := 0
    _flags = stream.get_flags(2)
    self.RIMInformation.Unpack(stream)// p8
    if (rIMRoutingAddress_flag & _flags) == rIMRoutingAddress_flag { //cond2
        self.RIMRoutingAddress = &RIMRoutingAddress{}//7{'type': 'RIMRoutingAddress', 'name': 'rIMRoutingAddress', 'optional': True}
        self.RIMRoutingAddress.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RIMTransferExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RIM-Transfer-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RIMTransferExtIEs, order_RIMTransferExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * RIMTransfer) Pack(stream *Stream) {
    const rIMRoutingAddress_flag uint = 0x00000001
    const iEExtensions_flag uint = 0x00000002
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RIMInformation.Pack(stream)
    if self.RIMRoutingAddress != nil { 
        _flags |= rIMRoutingAddress_flag
        self.RIMRoutingAddress.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RIMTransferExtIEs, order_RIMTransferExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *RIMRoutingAddress)Unpack(stream *Stream) {
    //coptions := []string{"globalRNC-ID","gERAN-Cell-ID"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in RIMRoutingAddress\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.GlobalRNCID = &GlobalRNCID{}//cho6
        self.GlobalRNCID.Unpack(stream)
    } else if choice == 1 { //ch2
        self.GERANCellID = &GERANCellID{}//cho6
        self.GERANCellID.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * RIMRoutingAddress) Pack(stream *Stream) {
    if self.GlobalRNCID != nil {
        stream.set_choice(0, 1, 1, 2)
        self.GlobalRNCID.Pack(stream)//2
    } else if self.GERANCellID != nil {
        stream.set_choice(1, 1, 1, 2)
        self.GERANCellID.Pack(stream)//2
    }

}
type RIMRoutingAddress struct { //[{'type': 'GlobalRNC-ID', 'name': 'globalRNC-ID'}, {'type': 'GERAN-Cell-ID', 'name': 'gERAN-Cell-ID'}, None]
    GlobalRNCID *GlobalRNCID
    GERANCellID *GERANCellID
} // RIMRoutingAddress

type RNCID struct {
  Value uint64
}
func (self *RNCID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4096, 12, 0, 0)
}
func (self * RNCID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4096, 12, 0, 0)
}
type RNCTraceInformation struct { // [{'type': 'TraceReference', 'name': 'traceReference'}, {'type': 'ENUMERATED', 'values': [('activated', 0), ('deactivated', 1)], 'name': 'traceActivationIndicator'}, {'type': 'EquipmentsToBeTraced', 'name': 'equipmentsToBeTraced', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RNCTraceInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    TraceReference TraceReference
    TraceActivationIndicator ENUMERATED
    EquipmentsToBeTraced *EquipmentsToBeTraced
    IEExtensions *RNCTraceInformationExtIEs
}

func (self * RNCTraceInformation) Unpack(stream *Stream) {
    equipmentsToBeTraced_flag := 0x00000001
    iEExtensions_flag := 0x00000002
    _flags := 0
    _flags = stream.get_flags(2)
    self.TraceReference.Unpack(stream)// p8
    var Unpack_traceActivationIndicator = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(1, 2, 0)
    }
    Unpack_traceActivationIndicator(stream, &self.TraceActivationIndicator)// p2
    if (equipmentsToBeTraced_flag & _flags) == equipmentsToBeTraced_flag { //cond2
        self.EquipmentsToBeTraced = &EquipmentsToBeTraced{}//7{'type': 'EquipmentsToBeTraced', 'name': 'equipmentsToBeTraced', 'optional': True}
        self.EquipmentsToBeTraced.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &RNCTraceInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['RNCTraceInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_RNCTraceInformationExtIEs, order_RNCTraceInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * RNCTraceInformation) Pack(stream *Stream) {
    const equipmentsToBeTraced_flag uint = 0x00000001
    const iEExtensions_flag uint = 0x00000002
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TraceReference.Pack(stream)
    var Pack_traceActivationIndicator = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 1, 2, 0)
    }
    Pack_traceActivationIndicator(stream, self.TraceActivationIndicator) //f2
    if self.EquipmentsToBeTraced != nil { 
        _flags |= equipmentsToBeTraced_flag
        self.EquipmentsToBeTraced.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_RNCTraceInformationExtIEs, order_RNCTraceInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RRCContainer struct {
  Value HexBytes
}
func (self *RRCContainer) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *RRCContainer) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
type RTLoadValue struct {
  Value uint64
}
func (self *RTLoadValue) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(101, 7, 0, 0)
}
func (self * RTLoadValue) Pack(st *Stream){
    st.formatf_Integer(self.Value, 101, 7, 0, 0)
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
type SAI struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'LAC', 'name': 'lAC'}, {'type': 'SAC', 'name': 'sAC'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SAI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PLMNidentity PLMNidentity
    LAC LAC
    SAC SAC
    IEExtensions *SAIExtIEs
}

func (self * SAI) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.PLMNidentity.Unpack(stream)// p8
    self.LAC.Unpack(stream)// p8
    self.SAC.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SAIExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SAI-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SAIExtIEs, order_SAIExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * SAI) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNidentity.Pack(stream)
    self.LAC.Pack(stream)
    self.SAC.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SAIExtIEs, order_SAIExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SAPI struct {
  Value int
}
const (
    SAPIsapi_0 = 0
    SAPIsapi_3 = 1

    /* Extensions */
)
func (self *SAPI) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *SAPI) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type SharedNetworkInformation struct { // [{'type': 'PLMNs-in-shared-network', 'name': 'pLMNs-in-shared-network'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Shared-Network-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    PLMNsinsharednetwork PLMNsinsharednetwork
    IEExtensions *SharedNetworkInformationExtIEs
}

func (self * SharedNetworkInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.PLMNsinsharednetwork.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SharedNetworkInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['Shared-Network-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SharedNetworkInformationExtIEs, order_SharedNetworkInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SharedNetworkInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNsinsharednetwork.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SharedNetworkInformationExtIEs, order_SharedNetworkInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SignallingIndication struct {
  Value int
}
const (
    SignallingIndicationsignalling = 0

    /* Extensions */
)
func (self *SignallingIndication) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(1, 1, 1)
}
func (self *SignallingIndication) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 1, 1, 1)
}
type SDUErrorRatio struct { // [{'type': 'INTEGER', 'restricted-to': [(1, 9)], 'name': 'mantissa'}, {'type': 'INTEGER', 'restricted-to': [(1, 6)], 'name': 'exponent'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDU-ErrorRatio-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    Mantissa INTEGER
    Exponent INTEGER
    IEExtensions *SDUErrorRatioExtIEs
}

func (self * SDUErrorRatio) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    var Unpack_mantissa = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(9, 4, 0, 1)
    }
    Unpack_mantissa(stream, &self.Mantissa)// p2
    var Unpack_exponent = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(6, 3, 0, 1)
    }
    Unpack_exponent(stream, &self.Exponent)// p2
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SDUErrorRatioExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDU-ErrorRatio-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SDUErrorRatioExtIEs, order_SDUErrorRatioExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * SDUErrorRatio) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_mantissa = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 9, 4, 0, 1)
    }
    Pack_mantissa(stream, self.Mantissa) //f2
    var Pack_exponent = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 6, 3, 0, 1)
    }
    Pack_exponent(stream, self.Exponent) //f2
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SDUErrorRatioExtIEs, order_SDUErrorRatioExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *SDUFormatInformationParameters) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(64)
    _size += 1
    self.Items = make([]SDUFormatInformationParameters_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *SDUFormatInformationParameters_Item) { //[{'type': 'SubflowSDU-Size', 'name': 'subflowSDU-Size', 'optional': True}, {'type': 'RAB-SubflowCombinationBitRate', 'name': 'rAB-SubflowCombinationBitRate', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDU-FormatInformationParameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        subflowSDUSize_flag := 0x00000002
        rABSubflowCombinationBitRate_flag := 0x00000004
        iEExtensions_flag := 0x00000008
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(4)
        if (subflowSDUSize_flag & _flags) == subflowSDUSize_flag { //cond2
            self.SubflowSDUSize = &SubflowSDUSize{}//7{'type': 'SubflowSDU-Size', 'name': 'subflowSDU-Size', 'optional': True}
            self.SubflowSDUSize.Unpack(stream)// p8
        }
        if (rABSubflowCombinationBitRate_flag & _flags) == rABSubflowCombinationBitRate_flag { //cond2
            self.RABSubflowCombinationBitRate = &RABSubflowCombinationBitRate{}//7{'type': 'RAB-SubflowCombinationBitRate', 'name': 'rAB-SubflowCombinationBitRate', 'optional': True}
            self.RABSubflowCombinationBitRate.Unpack(stream)// p8
        }
        if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
            self.IEExtensions = &SDUFormatInformationParametersExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDU-FormatInformationParameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
            IEExtensions := ProtocolExtensionContainer {table_SDUFormatInformationParametersExtIEs, order_SDUFormatInformationParametersExtIEs} // p3
            IEExtensions.Unpack(stream, &self.IEExtensions) // p3
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    for item := 0; item <_size; item +=1 {
        Unpack_Item(stream, &self.Items[item])
    }
}


func (self *SDUFormatInformationParameters) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 64)
    var Pack_Item = func(stream *Stream, self SDUFormatInformationParameters_Item) {//seq
        const subflowSDUSize_flag uint = 0x00000002
        const rABSubflowCombinationBitRate_flag uint = 0x00000004
        const iEExtensions_flag uint = 0x00000008
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(4)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        if self.SubflowSDUSize != nil { 
            _flags |= subflowSDUSize_flag
            self.SubflowSDUSize.Pack(stream)
        }//end of optional
        if self.RABSubflowCombinationBitRate != nil { 
            _flags |= rABSubflowCombinationBitRate_flag
            self.RABSubflowCombinationBitRate.Pack(stream)
        }//end of optional
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_SDUFormatInformationParametersExtIEs, order_SDUFormatInformationParametersExtIEs} // p3
            IEExtensions.Pack(stream, &self.IEExtensions)
        }//end of optional
        if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
        stream.set_flags(_flags, _flagReserve, 4)
    }//end
    for _, item := range self.Items { // seqof base type
        Pack_Item(stream, item)
    }
    return

}


type SDUFormatInformationParameters_Item struct { // [{'type': 'SubflowSDU-Size', 'name': 'subflowSDU-Size', 'optional': True}, {'type': 'RAB-SubflowCombinationBitRate', 'name': 'rAB-SubflowCombinationBitRate', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDU-FormatInformationParameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    SubflowSDUSize *SubflowSDUSize
    RABSubflowCombinationBitRate *RABSubflowCombinationBitRate
    IEExtensions *SDUFormatInformationParametersExtIEs
}
type SDUFormatInformationParameters struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'SubflowSDU-Size', 'name': 'subflowSDU-Size', 'optional': True}, {'type': 'RAB-SubflowCombinationBitRate', 'name': 'rAB-SubflowCombinationBitRate', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDU-FormatInformationParameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxRAB-SubflowCombination')]}
    Items []SDUFormatInformationParameters_Item
}

func (self *SDUParameters) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(7)
    _size += 1
    self.Items = make([]SDUParameters_Item, _size)//1
    var Unpack_Item = func(stream *Stream, self *SDUParameters_Item) { //[{'type': 'SDU-ErrorRatio', 'name': 'sDU-ErrorRatio', 'optional': True}, {'type': 'ResidualBitErrorRatio', 'name': 'residualBitErrorRatio'}, {'type': 'DeliveryOfErroneousSDU', 'name': 'deliveryOfErroneousSDU'}, {'type': 'SDU-FormatInformationParameters', 'name': 'sDU-FormatInformationParameters', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDU-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
        sDUErrorRatio_flag := 0x00000002
        sDUFormatInformationParameters_flag := 0x00000004
        iEExtensions_flag := 0x00000008
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(4)
        if (sDUErrorRatio_flag & _flags) == sDUErrorRatio_flag { //cond2
            self.SDUErrorRatio = &SDUErrorRatio{}//7{'type': 'SDU-ErrorRatio', 'name': 'sDU-ErrorRatio', 'optional': True}
            self.SDUErrorRatio.Unpack(stream)// p8
        }
        self.ResidualBitErrorRatio.Unpack(stream)// p8
        self.DeliveryOfErroneousSDU.Unpack(stream)// p8
        if (sDUFormatInformationParameters_flag & _flags) == sDUFormatInformationParameters_flag { //cond2
            self.SDUFormatInformationParameters = &SDUFormatInformationParameters{}//7{'type': 'SDU-FormatInformationParameters', 'name': 'sDU-FormatInformationParameters', 'optional': True}
            self.SDUFormatInformationParameters.Unpack(stream)// p8
        }
        if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
            self.IEExtensions = &SDUParametersExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDU-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
            IEExtensions := ProtocolExtensionContainer {table_SDUParametersExtIEs, order_SDUParametersExtIEs} // p3
            IEExtensions.Unpack(stream, &self.IEExtensions) // p3
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    for item := 0; item <_size; item +=1 {
        Unpack_Item(stream, &self.Items[item])
    }
}


func (self *SDUParameters) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 7)
    var Pack_Item = func(stream *Stream, self SDUParameters_Item) {//seq
        const sDUErrorRatio_flag uint = 0x00000002
        const sDUFormatInformationParameters_flag uint = 0x00000004
        const iEExtensions_flag uint = 0x00000008
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(4)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        if self.SDUErrorRatio != nil { 
            _flags |= sDUErrorRatio_flag
            self.SDUErrorRatio.Pack(stream)
        }//end of optional
        self.ResidualBitErrorRatio.Pack(stream)
        self.DeliveryOfErroneousSDU.Pack(stream)
        if self.SDUFormatInformationParameters != nil { 
            _flags |= sDUFormatInformationParameters_flag
            self.SDUFormatInformationParameters.Pack(stream)
        }//end of optional
        if self.IEExtensions != nil { 
            _flags |= iEExtensions_flag
            IEExtensions := &ProtocolExtensionContainer {table_SDUParametersExtIEs, order_SDUParametersExtIEs} // p3
            IEExtensions.Pack(stream, &self.IEExtensions)
        }//end of optional
        if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
        stream.set_flags(_flags, _flagReserve, 4)
    }//end
    for _, item := range self.Items { // seqof base type
        Pack_Item(stream, item)
    }
    return

}


type SDUParameters_Item struct { // [{'type': 'SDU-ErrorRatio', 'name': 'sDU-ErrorRatio', 'optional': True}, {'type': 'ResidualBitErrorRatio', 'name': 'residualBitErrorRatio'}, {'type': 'DeliveryOfErroneousSDU', 'name': 'deliveryOfErroneousSDU'}, {'type': 'SDU-FormatInformationParameters', 'name': 'sDU-FormatInformationParameters', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDU-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    SDUErrorRatio *SDUErrorRatio
    ResidualBitErrorRatio ResidualBitErrorRatio
    DeliveryOfErroneousSDU DeliveryOfErroneousSDU
    SDUFormatInformationParameters *SDUFormatInformationParameters
    IEExtensions *SDUParametersExtIEs
}
type SDUParameters struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SEQUENCE', 'members': [{'type': 'SDU-ErrorRatio', 'name': 'sDU-ErrorRatio', 'optional': True}, {'type': 'ResidualBitErrorRatio', 'name': 'residualBitErrorRatio'}, {'type': 'DeliveryOfErroneousSDU', 'name': 'deliveryOfErroneousSDU'}, {'type': 'SDU-FormatInformationParameters', 'name': 'sDU-FormatInformationParameters', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SDU-Parameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]}, 'size': [(1, 'maxRAB-Subflows')]}
    Items []SDUParameters_Item
}

type SNAAccessInformation struct { // [{'type': 'AuthorisedPLMNs', 'name': 'authorisedPLMNs'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SNA-Access-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    AuthorisedPLMNs AuthorisedPLMNs
    IEExtensions *SNAAccessInformationExtIEs
}

func (self * SNAAccessInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.AuthorisedPLMNs.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SNAAccessInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SNA-Access-Information-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SNAAccessInformationExtIEs, order_SNAAccessInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SNAAccessInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.AuthorisedPLMNs.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SNAAccessInformationExtIEs, order_SNAAccessInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SNAC struct {
  Value uint64
}
func (self *SNAC) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * SNAC) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type ServiceHandover struct {
  Value int
}
const (
    ServiceHandoverhandover_to_GSM_should_be_performed = 0
    ServiceHandoverhandover_to_GSM_should_not_be_performed = 1
    ServiceHandoverhandover_to_GSM_shall_not_be_performed = 2

    /* Extensions */
)
func (self *ServiceHandover) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *ServiceHandover) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
func (self *SourceCellID)Unpack(stream *Stream) {
    //coptions := []string{"sourceUTRANCellID","sourceGERANCellID"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in SourceCellID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.SourceUTRANCellID = &SourceUTRANCellID{}//cho6
        self.SourceUTRANCellID.Unpack(stream)
    } else if choice == 1 { //ch2
        self.SourceGERANCellID = &CGI{}//cho6
        self.SourceGERANCellID.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * SourceCellID) Pack(stream *Stream) {
    if self.SourceUTRANCellID != nil {
        stream.set_choice(0, 1, 1, 2)
        self.SourceUTRANCellID.Pack(stream)//2
    } else if self.SourceGERANCellID != nil {
        stream.set_choice(1, 1, 1, 2)
        self.SourceGERANCellID.Pack(stream)//2
    }

}
type SourceCellID struct { //[{'type': 'SourceUTRANCellID', 'name': 'sourceUTRANCellID'}, {'type': 'CGI', 'name': 'sourceGERANCellID'}, None]
    SourceUTRANCellID *SourceUTRANCellID
    SourceGERANCellID *CGI
} // SourceCellID

func (self *SourceID)Unpack(stream *Stream) {
    //coptions := []string{"sourceRNC-ID","sAI"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in SourceID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.SourceRNCID = &SourceRNCID{}//cho6
        self.SourceRNCID.Unpack(stream)
    } else if choice == 1 { //ch2
        self.SAI = &SAI{}//cho6
        self.SAI.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * SourceID) Pack(stream *Stream) {
    if self.SourceRNCID != nil {
        stream.set_choice(0, 1, 1, 2)
        self.SourceRNCID.Pack(stream)//2
    } else if self.SAI != nil {
        stream.set_choice(1, 1, 1, 2)
        self.SAI.Pack(stream)//2
    }

}
type SourceID struct { //[{'type': 'SourceRNC-ID', 'name': 'sourceRNC-ID'}, {'type': 'SAI', 'name': 'sAI'}, None]
    SourceRNCID *SourceRNCID
    SAI *SAI
} // SourceID

type SourceRNCID struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'RNC-ID', 'name': 'rNC-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SourceRNC-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PLMNidentity PLMNidentity
    RNCID RNCID
    IEExtensions *SourceRNCIDExtIEs
}

func (self * SourceRNCID) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.PLMNidentity.Unpack(stream)// p8
    self.RNCID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SourceRNCIDExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SourceRNC-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SourceRNCIDExtIEs, order_SourceRNCIDExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * SourceRNCID) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNidentity.Pack(stream)
    self.RNCID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SourceRNCIDExtIEs, order_SourceRNCIDExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SourceRNCToTargetRNCTransparentContainer struct { // [{'type': 'RRC-Container', 'name': 'rRC-Container'}, {'type': 'NumberOfIuInstances', 'name': 'numberOfIuInstances'}, {'type': 'RelocationType', 'name': 'relocationType'}, {'type': 'ChosenIntegrityProtectionAlgorithm', 'name': 'chosenIntegrityProtectionAlgorithm', 'optional': True}, {'type': 'IntegrityProtectionKey', 'name': 'integrityProtectionKey', 'optional': True}, {'type': 'ChosenEncryptionAlgorithm', 'name': 'chosenEncryptionAlgorithForSignalling', 'optional': True}, {'type': 'EncryptionKey', 'name': 'cipheringKey', 'optional': True}, {'type': 'ChosenEncryptionAlgorithm', 'name': 'chosenEncryptionAlgorithForCS', 'optional': True}, {'type': 'ChosenEncryptionAlgorithm', 'name': 'chosenEncryptionAlgorithForPS', 'optional': True}, {'type': 'D-RNTI', 'name': 'd-RNTI', 'optional': True}, {'type': 'TargetCellId', 'name': 'targetCellId', 'optional': True}, {'type': 'RAB-TrCH-Mapping', 'name': 'rAB-TrCH-Mapping', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SourceRNC-ToTargetRNC-TransparentContainer-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RRCContainer RRCContainer
    NumberOfIuInstances NumberOfIuInstances
    RelocationType RelocationType
    ChosenIntegrityProtectionAlgorithm *ChosenIntegrityProtectionAlgorithm
    IntegrityProtectionKey *IntegrityProtectionKey
    ChosenEncryptionAlgorithForSignalling *ChosenEncryptionAlgorithm
    CipheringKey *EncryptionKey
    ChosenEncryptionAlgorithForCS *ChosenEncryptionAlgorithm
    ChosenEncryptionAlgorithForPS *ChosenEncryptionAlgorithm
    DRNTI *DRNTI
    TargetCellId *TargetCellId
    RABTrCHMapping *RABTrCHMapping
    IEExtensions *SourceRNCToTargetRNCTransparentContainerExtIEs
}

func (self * SourceRNCToTargetRNCTransparentContainer) Unpack(stream *Stream) {
    chosenIntegrityProtectionAlgorithm_flag := 0x00000002
    integrityProtectionKey_flag := 0x00000004
    chosenEncryptionAlgorithForSignalling_flag := 0x00000008
    cipheringKey_flag := 0x00000010
    chosenEncryptionAlgorithForCS_flag := 0x00000020
    chosenEncryptionAlgorithForPS_flag := 0x00000040
    dRNTI_flag := 0x00000080
    targetCellId_flag := 0x00000100
    rABTrCHMapping_flag := 0x00000200
    iEExtensions_flag := 0x00000400
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(11)
    self.RRCContainer.Unpack(stream)// p8
    self.NumberOfIuInstances.Unpack(stream)// p8
    self.RelocationType.Unpack(stream)// p8
    if (chosenIntegrityProtectionAlgorithm_flag & _flags) == chosenIntegrityProtectionAlgorithm_flag { //cond2
        self.ChosenIntegrityProtectionAlgorithm = &ChosenIntegrityProtectionAlgorithm{}//7{'type': 'ChosenIntegrityProtectionAlgorithm', 'name': 'chosenIntegrityProtectionAlgorithm', 'optional': True}
        self.ChosenIntegrityProtectionAlgorithm.Unpack(stream)// p8
    }
    if (integrityProtectionKey_flag & _flags) == integrityProtectionKey_flag { //cond2
        self.IntegrityProtectionKey = &IntegrityProtectionKey{}//7{'type': 'IntegrityProtectionKey', 'name': 'integrityProtectionKey', 'optional': True}
        self.IntegrityProtectionKey.Unpack(stream)// p8
    }
    if (chosenEncryptionAlgorithForSignalling_flag & _flags) == chosenEncryptionAlgorithForSignalling_flag { //cond2
        self.ChosenEncryptionAlgorithForSignalling = &ChosenEncryptionAlgorithm{}//7{'type': 'ChosenEncryptionAlgorithm', 'name': 'chosenEncryptionAlgorithForSignalling', 'optional': True}
        self.ChosenEncryptionAlgorithForSignalling.Unpack(stream)// p8
    }
    if (cipheringKey_flag & _flags) == cipheringKey_flag { //cond2
        self.CipheringKey = &EncryptionKey{}//7{'type': 'EncryptionKey', 'name': 'cipheringKey', 'optional': True}
        self.CipheringKey.Unpack(stream)// p8
    }
    if (chosenEncryptionAlgorithForCS_flag & _flags) == chosenEncryptionAlgorithForCS_flag { //cond2
        self.ChosenEncryptionAlgorithForCS = &ChosenEncryptionAlgorithm{}//7{'type': 'ChosenEncryptionAlgorithm', 'name': 'chosenEncryptionAlgorithForCS', 'optional': True}
        self.ChosenEncryptionAlgorithForCS.Unpack(stream)// p8
    }
    if (chosenEncryptionAlgorithForPS_flag & _flags) == chosenEncryptionAlgorithForPS_flag { //cond2
        self.ChosenEncryptionAlgorithForPS = &ChosenEncryptionAlgorithm{}//7{'type': 'ChosenEncryptionAlgorithm', 'name': 'chosenEncryptionAlgorithForPS', 'optional': True}
        self.ChosenEncryptionAlgorithForPS.Unpack(stream)// p8
    }
    if (dRNTI_flag & _flags) == dRNTI_flag { //cond2
        self.DRNTI = &DRNTI{}//7{'type': 'D-RNTI', 'name': 'd-RNTI', 'optional': True}
        self.DRNTI.Unpack(stream)// p8
    }
    if (targetCellId_flag & _flags) == targetCellId_flag { //cond2
        self.TargetCellId = &TargetCellId{}//7{'type': 'TargetCellId', 'name': 'targetCellId', 'optional': True}
        self.TargetCellId.Unpack(stream)// p8
    }
    if (rABTrCHMapping_flag & _flags) == rABTrCHMapping_flag { //cond2
        self.RABTrCHMapping = &RABTrCHMapping{}//7{'type': 'RAB-TrCH-Mapping', 'name': 'rAB-TrCH-Mapping', 'optional': True}
        self.RABTrCHMapping.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SourceRNCToTargetRNCTransparentContainerExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SourceRNC-ToTargetRNC-TransparentContainer-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SourceRNCToTargetRNCTransparentContainerExtIEs, order_SourceRNCToTargetRNCTransparentContainerExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SourceRNCToTargetRNCTransparentContainer) Pack(stream *Stream) {
    const chosenIntegrityProtectionAlgorithm_flag uint = 0x00000002
    const integrityProtectionKey_flag uint = 0x00000004
    const chosenEncryptionAlgorithForSignalling_flag uint = 0x00000008
    const cipheringKey_flag uint = 0x00000010
    const chosenEncryptionAlgorithForCS_flag uint = 0x00000020
    const chosenEncryptionAlgorithForPS_flag uint = 0x00000040
    const dRNTI_flag uint = 0x00000080
    const targetCellId_flag uint = 0x00000100
    const rABTrCHMapping_flag uint = 0x00000200
    const iEExtensions_flag uint = 0x00000400
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(11)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RRCContainer.Pack(stream)
    self.NumberOfIuInstances.Pack(stream)
    self.RelocationType.Pack(stream)
    if self.ChosenIntegrityProtectionAlgorithm != nil { 
        _flags |= chosenIntegrityProtectionAlgorithm_flag
        self.ChosenIntegrityProtectionAlgorithm.Pack(stream)
    }//end of optional
    if self.IntegrityProtectionKey != nil { 
        _flags |= integrityProtectionKey_flag
        self.IntegrityProtectionKey.Pack(stream)
    }//end of optional
    if self.ChosenEncryptionAlgorithForSignalling != nil { 
        _flags |= chosenEncryptionAlgorithForSignalling_flag
        self.ChosenEncryptionAlgorithForSignalling.Pack(stream)
    }//end of optional
    if self.CipheringKey != nil { 
        _flags |= cipheringKey_flag
        self.CipheringKey.Pack(stream)
    }//end of optional
    if self.ChosenEncryptionAlgorithForCS != nil { 
        _flags |= chosenEncryptionAlgorithForCS_flag
        self.ChosenEncryptionAlgorithForCS.Pack(stream)
    }//end of optional
    if self.ChosenEncryptionAlgorithForPS != nil { 
        _flags |= chosenEncryptionAlgorithForPS_flag
        self.ChosenEncryptionAlgorithForPS.Pack(stream)
    }//end of optional
    if self.DRNTI != nil { 
        _flags |= dRNTI_flag
        self.DRNTI.Pack(stream)
    }//end of optional
    if self.TargetCellId != nil { 
        _flags |= targetCellId_flag
        self.TargetCellId.Pack(stream)
    }//end of optional
    if self.RABTrCHMapping != nil { 
        _flags |= rABTrCHMapping_flag
        self.RABTrCHMapping.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SourceRNCToTargetRNCTransparentContainerExtIEs, order_SourceRNCToTargetRNCTransparentContainerExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 11)
}//end

type SourceStatisticsDescriptor struct {
  Value int
}
const (
    SourceStatisticsDescriptorspeech = 0
    SourceStatisticsDescriptorunknown = 1

    /* Extensions */
)
func (self *SourceStatisticsDescriptor) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *SourceStatisticsDescriptor) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type SourceUTRANCellID struct { // [{'type': 'PLMNidentity', 'name': 'pLMNidentity'}, {'type': 'TargetCellId', 'name': 'uTRANcellID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SourceUTRANCellID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    PLMNidentity PLMNidentity
    UTRANcellID TargetCellId
    IEExtensions *SourceUTRANCellIDExtIEs
}

func (self * SourceUTRANCellID) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000001
    _flags := 0
    _flags = stream.get_flags(1)
    self.PLMNidentity.Unpack(stream)// p8
    self.UTRANcellID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SourceUTRANCellIDExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SourceUTRANCellID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SourceUTRANCellIDExtIEs, order_SourceUTRANCellIDExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * SourceUTRANCellID) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNidentity.Pack(stream)
    self.UTRANcellID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SourceUTRANCellIDExtIEs, order_SourceUTRANCellIDExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type SRBID struct {
  Value uint64
}
func (self *SRBID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(32, 5, 0, 1)
}
func (self * SRBID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 32, 5, 0, 1)
}
func (self *SRBTrCHMapping) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(8)
    _size += 1
    self.Items = make([]SRBTrCHMappingItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *SRBTrCHMapping) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 8)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type SRBTrCHMapping struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SRB-TrCH-MappingItem'}, 'size': [(1, 'maxNrOfSRBs')]}
    Items []SRBTrCHMappingItem
}

type SRBTrCHMappingItem struct { // [{'type': 'SRB-ID', 'name': 'sRB-ID'}, {'type': 'TrCH-ID', 'name': 'trCH-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SRB-TrCH-MappingItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    SRBID SRBID
    TrCHID TrCHID
    IEExtensions *SRBTrCHMappingItemExtIEs
}

func (self * SRBTrCHMappingItem) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.SRBID.Unpack(stream)// p8
    self.TrCHID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &SRBTrCHMappingItemExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['SRB-TrCH-MappingItem-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_SRBTrCHMappingItemExtIEs, order_SRBTrCHMappingItemExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SRBTrCHMappingItem) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.SRBID.Pack(stream)
    self.TrCHID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_SRBTrCHMappingItemExtIEs, order_SRBTrCHMappingItemExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SubflowSDUSize struct {
  Value uint64
}
func (self *SubflowSDUSize) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4096, 12, 0, 0)
}
func (self * SubflowSDUSize) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4096, 12, 0, 0)
}
type TargetCellId struct {
  Value uint64
}
func (self *TargetCellId) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(268435456, 28, 0, 0)
}
func (self * TargetCellId) Pack(st *Stream){
    st.formatf_Integer(self.Value, 268435456, 28, 0, 0)
}
func (self *TargetID)Unpack(stream *Stream) {
    //coptions := []string{"targetRNC-ID","cGI"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in TargetID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.TargetRNCID = &TargetRNCID{}//cho6
        self.TargetRNCID.Unpack(stream)
    } else if choice == 1 { //ch2
        self.CGI = &CGI{}//cho6
        self.CGI.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * TargetID) Pack(stream *Stream) {
    if self.TargetRNCID != nil {
        stream.set_choice(0, 1, 1, 2)
        self.TargetRNCID.Pack(stream)//2
    } else if self.CGI != nil {
        stream.set_choice(1, 1, 1, 2)
        self.CGI.Pack(stream)//2
    }

}
type TargetID struct { //[{'type': 'TargetRNC-ID', 'name': 'targetRNC-ID'}, {'type': 'CGI', 'name': 'cGI'}, None]
    TargetRNCID *TargetRNCID
    CGI *CGI
} // TargetID

type TargetRNCID struct { // [{'type': 'LAI', 'name': 'lAI'}, {'type': 'RAC', 'name': 'rAC', 'optional': True}, {'type': 'RNC-ID', 'name': 'rNC-ID'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TargetRNC-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}]
    LAI LAI
    RAC *RAC
    RNCID RNCID
    IEExtensions *TargetRNCIDExtIEs
}

func (self * TargetRNCID) Unpack(stream *Stream) {
    rAC_flag := 0x00000001
    iEExtensions_flag := 0x00000002
    _flags := 0
    _flags = stream.get_flags(2)
    self.LAI.Unpack(stream)// p8
    if (rAC_flag & _flags) == rAC_flag { //cond2
        self.RAC = &RAC{}//7{'type': 'RAC', 'name': 'rAC', 'optional': True}
        self.RAC.Unpack(stream)// p8
    }
    self.RNCID.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &TargetRNCIDExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TargetRNC-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_TargetRNCIDExtIEs, order_TargetRNCIDExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    return
}

func (self * TargetRNCID) Pack(stream *Stream) {
    const rAC_flag uint = 0x00000001
    const iEExtensions_flag uint = 0x00000002
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.LAI.Pack(stream)
    if self.RAC != nil { 
        _flags |= rAC_flag
        self.RAC.Pack(stream)
    }//end of optional
    self.RNCID.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_TargetRNCIDExtIEs, order_TargetRNCIDExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type TargetRNCToSourceRNCTransparentContainer struct { // [{'type': 'RRC-Container', 'name': 'rRC-Container'}, {'type': 'D-RNTI', 'name': 'd-RNTI', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TargetRNC-ToSourceRNC-TransparentContainer-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    RRCContainer RRCContainer
    DRNTI *DRNTI
    IEExtensions *TargetRNCToSourceRNCTransparentContainerExtIEs
}

func (self * TargetRNCToSourceRNCTransparentContainer) Unpack(stream *Stream) {
    dRNTI_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RRCContainer.Unpack(stream)// p8
    if (dRNTI_flag & _flags) == dRNTI_flag { //cond2
        self.DRNTI = &DRNTI{}//7{'type': 'D-RNTI', 'name': 'd-RNTI', 'optional': True}
        self.DRNTI.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &TargetRNCToSourceRNCTransparentContainerExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TargetRNC-ToSourceRNC-TransparentContainer-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_TargetRNCToSourceRNCTransparentContainerExtIEs, order_TargetRNCToSourceRNCTransparentContainerExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TargetRNCToSourceRNCTransparentContainer) Pack(stream *Stream) {
    const dRNTI_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RRCContainer.Pack(stream)
    if self.DRNTI != nil { 
        _flags |= dRNTI_flag
        self.DRNTI.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_TargetRNCToSourceRNCTransparentContainerExtIEs, order_TargetRNCToSourceRNCTransparentContainerExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type TBCDSTRING struct {
  Value HexBytes
}
func (self *TBCDSTRING) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}
func (self *TBCDSTRING) Pack(st *Stream) {
    st.format_len(len(self.Value), 0)
    st.formatf_OctString(self.Value, 0)
}
func (self *TemporaryUEID)Unpack(stream *Stream) {
    //coptions := []string{"tMSI","p-TMSI"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in TemporaryUEID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.TMSI = &TMSI{}//cho6
        self.TMSI.Unpack(stream)
    } else if choice == 1 { //ch2
        self.PTMSI = &PTMSI{}//cho6
        self.PTMSI.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * TemporaryUEID) Pack(stream *Stream) {
    if self.TMSI != nil {
        stream.set_choice(0, 1, 1, 2)
        self.TMSI.Pack(stream)//2
    } else if self.PTMSI != nil {
        stream.set_choice(1, 1, 1, 2)
        self.PTMSI.Pack(stream)//2
    }

}
type TemporaryUEID struct { //[{'type': 'TMSI', 'name': 'tMSI'}, {'type': 'P-TMSI', 'name': 'p-TMSI'}, None]
    TMSI *TMSI
    PTMSI *PTMSI
} // TemporaryUEID

type TMSI struct {
  Value HexBytes
}
func (self *TMSI) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(4)
}
func (self *TMSI) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 4)
}
type TraceDepth struct {
  Value int
}
const (
    TraceDepthminimum = 0
    TraceDepthmedium = 1
    TraceDepthmaximum = 2

    /* Extensions */
)
func (self *TraceDepth) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 3, 1)
}
func (self *TraceDepth) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 3, 1)
}
type TracePropagationParameters struct { // [{'type': 'TraceRecordingSessionReference', 'name': 'traceRecordingSessionReference'}, {'type': 'TraceDepth', 'name': 'traceDepth'}, {'type': 'ListOfInterfacesToTrace', 'name': 'listOfInterfacesToTrace', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TracePropagationParameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    TraceRecordingSessionReference TraceRecordingSessionReference
    TraceDepth TraceDepth
    ListOfInterfacesToTrace *ListOfInterfacesToTrace
    IEExtensions *TracePropagationParametersExtIEs
}

func (self * TracePropagationParameters) Unpack(stream *Stream) {
    listOfInterfacesToTrace_flag := 0x00000002
    iEExtensions_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.TraceRecordingSessionReference.Unpack(stream)// p8
    self.TraceDepth.Unpack(stream)// p8
    if (listOfInterfacesToTrace_flag & _flags) == listOfInterfacesToTrace_flag { //cond2
        self.ListOfInterfacesToTrace = &ListOfInterfacesToTrace{}//7{'type': 'ListOfInterfacesToTrace', 'name': 'listOfInterfacesToTrace', 'optional': True}
        self.ListOfInterfacesToTrace.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &TracePropagationParametersExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TracePropagationParameters-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_TracePropagationParametersExtIEs, order_TracePropagationParametersExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TracePropagationParameters) Pack(stream *Stream) {
    const listOfInterfacesToTrace_flag uint = 0x00000002
    const iEExtensions_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TraceRecordingSessionReference.Pack(stream)
    self.TraceDepth.Pack(stream)
    if self.ListOfInterfacesToTrace != nil { 
        _flags |= listOfInterfacesToTrace_flag
        self.ListOfInterfacesToTrace.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_TracePropagationParametersExtIEs, order_TracePropagationParametersExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type TraceRecordingSessionInformation struct { // [{'type': 'TraceReference', 'name': 'traceReference'}, {'type': 'TraceRecordingSessionReference', 'name': 'traceRecordingSessionReference'}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TraceRecordingSessionInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    TraceReference TraceReference
    TraceRecordingSessionReference TraceRecordingSessionReference
    IEExtensions *TraceRecordingSessionInformationExtIEs
}

func (self * TraceRecordingSessionInformation) Unpack(stream *Stream) {
    iEExtensions_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.TraceReference.Unpack(stream)// p8
    self.TraceRecordingSessionReference.Unpack(stream)// p8
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &TraceRecordingSessionInformationExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TraceRecordingSessionInformation-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_TraceRecordingSessionInformationExtIEs, order_TraceRecordingSessionInformationExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TraceRecordingSessionInformation) Pack(stream *Stream) {
    const iEExtensions_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TraceReference.Pack(stream)
    self.TraceRecordingSessionReference.Pack(stream)
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_TraceRecordingSessionInformationExtIEs, order_TraceRecordingSessionInformationExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type TraceRecordingSessionReference struct {
  Value uint64
}
func (self *TraceRecordingSessionReference) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * TraceRecordingSessionReference) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type TraceReference struct {
  Value HexBytes
}
func (self *TraceReference) Unpack(st *Stream) {
    _len := st.parse_olen(1)+2
    if _len < 2 || _len > 3 {
        //fmt.Println ("Invalid len in TraceReference")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *TraceReference) Pack(st *Stream) {
    if len(self.Value) < 2 || len(self.Value) > 3 {
        log.Error ("Invalid len in TraceReference")
        return
}
    st.format_olen((len(self.Value))-2, 1)
    st.formatf_OctString(self.Value, 0)
}
type TraceType struct {
  Value HexBytes
}
func (self *TraceType) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(1)
}
func (self *TraceType) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 1)
}
type TrafficClass struct {
  Value int
}
const (
    TrafficClassconversational = 0
    TrafficClassstreaming = 1
    TrafficClassinteractive = 2
    TrafficClassbackground = 3

    /* Extensions */
)
func (self *TrafficClass) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 4, 1)
}
func (self *TrafficClass) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 4, 1)
}
type TrafficHandlingPriority struct {
  Value uint64
}
func (self *TrafficHandlingPriority) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(16, 4, 0, 0)
}
func (self * TrafficHandlingPriority) Pack(st *Stream){
    st.formatf_Integer(self.Value, 16, 4, 0, 0)
}
type TransferDelay struct {
  Value uint64
}
func (self *TransferDelay) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * TransferDelay) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type UnsuccessfullyTransmittedDataVolume struct {
  Value uint64
}
func (self *UnsuccessfullyTransmittedDataVolume) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967296, 32, 0, 0)
}
func (self * UnsuccessfullyTransmittedDataVolume) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967296, 32, 0, 0)
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
type TrCHID struct { // [{'type': 'DCH-ID', 'name': 'dCH-ID', 'optional': True}, {'type': 'DSCH-ID', 'name': 'dSCH-ID', 'optional': True}, {'type': 'USCH-ID', 'name': 'uSCH-ID', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TrCH-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    DCHID *DCHID
    DSCHID *DSCHID
    USCHID *USCHID
    IEExtensions *TrCHIDExtIEs
}

func (self * TrCHID) Unpack(stream *Stream) {
    dCHID_flag := 0x00000002
    dSCHID_flag := 0x00000004
    uSCHID_flag := 0x00000008
    iEExtensions_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    if (dCHID_flag & _flags) == dCHID_flag { //cond2
        self.DCHID = &DCHID{}//7{'type': 'DCH-ID', 'name': 'dCH-ID', 'optional': True}
        self.DCHID.Unpack(stream)// p8
    }
    if (dSCHID_flag & _flags) == dSCHID_flag { //cond2
        self.DSCHID = &DSCHID{}//7{'type': 'DSCH-ID', 'name': 'dSCH-ID', 'optional': True}
        self.DSCHID.Unpack(stream)// p8
    }
    if (uSCHID_flag & _flags) == uSCHID_flag { //cond2
        self.USCHID = &USCHID{}//7{'type': 'USCH-ID', 'name': 'uSCH-ID', 'optional': True}
        self.USCHID.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &TrCHIDExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['TrCH-ID-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_TrCHIDExtIEs, order_TrCHIDExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TrCHID) Pack(stream *Stream) {
    const dCHID_flag uint = 0x00000002
    const dSCHID_flag uint = 0x00000004
    const uSCHID_flag uint = 0x00000008
    const iEExtensions_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.DCHID != nil { 
        _flags |= dCHID_flag
        self.DCHID.Pack(stream)
    }//end of optional
    if self.DSCHID != nil { 
        _flags |= dSCHID_flag
        self.DSCHID.Pack(stream)
    }//end of optional
    if self.USCHID != nil { 
        _flags |= uSCHID_flag
        self.USCHID.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_TrCHIDExtIEs, order_TrCHIDExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *TrCHIDList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(7)
    _size += 1
    self.Items = make([]TrCHID, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *TrCHIDList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 7)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type TrCHIDList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'TrCH-ID'}, 'size': [(1, 'maxRAB-Subflows')]}
    Items []TrCHID
}

type TriggerID struct {
  Value HexBytes
}
func (self *TriggerID) Unpack(st *Stream) {
    _len := st.parse_olen(5)+3
    if _len < 3 || _len > 22 {
        //fmt.Println ("Invalid len in TriggerID")
        return
    }
    self.Value = st.parsef_OctString(_len)
}
func (self *TriggerID) Pack(st *Stream) {
    if len(self.Value) < 3 || len(self.Value) > 22 {
        log.Error ("Invalid len in TriggerID")
        return
}
    st.format_olen((len(self.Value))-3, 5)
    st.formatf_OctString(self.Value, 0)
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
func (self *UEID)Unpack(stream *Stream) {
    //coptions := []string{"imsi","imei","imeisv"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in UEID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.Imsi = &IMSI{}//cho6
        self.Imsi.Unpack(stream)
    } else if choice == 1 { //ch2
        self.Imei = &IMEI{}//cho6
        self.Imei.Unpack(stream)
    } else if choice == 2 { //ch2
        self.Imeisv = &IMEISV{}//cho6
        self.Imeisv.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * UEID) Pack(stream *Stream) {
    if self.Imsi != nil {
        stream.set_choice(0, 1, 1, 2)
        self.Imsi.Pack(stream)//2
    } else if self.Imei != nil {
        stream.set_choice(1, 1, 1, 2)
        self.Imei.Pack(stream)//2
    } else if self.Imeisv != nil {
        stream.set_choice(2, 1, 1, 2)
        lenLoc := stream.reserve_len()
        self.Imeisv.Pack(stream)//2
        stream.set_len(lenLoc)
    }

}
type UEID struct { //[{'type': 'IMSI', 'name': 'imsi'}, {'type': 'IMEI', 'name': 'imei'}, None, {'type': 'IMEISV', 'name': 'imeisv'}]
    Imsi *IMSI
    Imei *IMEI
    Imeisv *IMEISV
} // UEID

type UESBIIu struct { // [{'type': 'UESBI-IuA', 'name': 'uESBI-IuA', 'optional': True}, {'type': 'UESBI-IuB', 'name': 'uESBI-IuB', 'optional': True}, {'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UESBI-Iu-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}, None]
    UESBIIuA *UESBIIuA
    UESBIIuB *UESBIIuB
    IEExtensions *UESBIIuExtIEs
}

func (self * UESBIIu) Unpack(stream *Stream) {
    uESBIIuA_flag := 0x00000002
    uESBIIuB_flag := 0x00000004
    iEExtensions_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    if (uESBIIuA_flag & _flags) == uESBIIuA_flag { //cond2
        self.UESBIIuA = &UESBIIuA{}//7{'type': 'UESBI-IuA', 'name': 'uESBI-IuA', 'optional': True}
        self.UESBIIuA.Unpack(stream)// p8
    }
    if (uESBIIuB_flag & _flags) == uESBIIuB_flag { //cond2
        self.UESBIIuB = &UESBIIuB{}//7{'type': 'UESBI-IuB', 'name': 'uESBI-IuB', 'optional': True}
        self.UESBIIuB.Unpack(stream)// p8
    }
    if (iEExtensions_flag & _flags) == iEExtensions_flag { //cond2
        self.IEExtensions = &UESBIIuExtIEs{}//7{'type': 'ProtocolExtensionContainer', 'actual-parameters': ['UESBI-Iu-ExtIEs'], 'name': 'iE-Extensions', 'optional': True}
        IEExtensions := ProtocolExtensionContainer {table_UESBIIuExtIEs, order_UESBIIuExtIEs} // p3
        IEExtensions.Unpack(stream, &self.IEExtensions) // p3
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UESBIIu) Pack(stream *Stream) {
    const uESBIIuA_flag uint = 0x00000002
    const uESBIIuB_flag uint = 0x00000004
    const iEExtensions_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.UESBIIuA != nil { 
        _flags |= uESBIIuA_flag
        self.UESBIIuA.Pack(stream)
    }//end of optional
    if self.UESBIIuB != nil { 
        _flags |= uESBIIuB_flag
        self.UESBIIuB.Pack(stream)
    }//end of optional
    if self.IEExtensions != nil { 
        _flags |= iEExtensions_flag
        IEExtensions := &ProtocolExtensionContainer {table_UESBIIuExtIEs, order_UESBIIuExtIEs} // p3
        IEExtensions.Pack(stream, &self.IEExtensions)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type UESBIIuA struct {
  Len int
  Value HexBytes
}
func (self *UESBIIuA) Unpack(st *Stream){
    self.Len = int(st.parse_blen(8, 0)+1)
    self.Value = st.parsef_BitString(128, int(self.Len))
}
func (self *UESBIIuA) Pack(st *Stream) {
    st.format_blen(int(self.Len-1), 8, 0)
    st.formatf_BitString(self.Value, int(self.Len))
}
type UESBIIuB struct {
  Len int
  Value HexBytes
}
func (self *UESBIIuB) Unpack(st *Stream){
    self.Len = int(st.parse_blen(8, 0)+1)
    self.Value = st.parsef_BitString(128, int(self.Len))
}
func (self *UESBIIuB) Pack(st *Stream) {
    st.format_blen(int(self.Len-1), 8, 0)
    st.formatf_BitString(self.Value, int(self.Len))
}
type ULGTPPDUSequenceNumber struct {
  Value uint64
}
func (self *ULGTPPDUSequenceNumber) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * ULGTPPDUSequenceNumber) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type ULNPDUSequenceNumber struct {
  Value uint64
}
func (self *ULNPDUSequenceNumber) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * ULNPDUSequenceNumber) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type UPModeVersions struct {
  Len int
  Value HexBytes
}
func (self *UPModeVersions) Unpack(st *Stream){
    self.Value = st.parsef_BitString(16, 16)
}
func (self *UPModeVersions) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 16)
}
type USCHID struct {
  Value uint64
}
func (self *USCHID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * USCHID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
type UserPlaneMode struct {
  Value int
}
const (
    UserPlaneModetransparent_mode = 0
    UserPlaneModesupport_mode_for_predefined_SDU_sizes = 1

    /* Extensions */
)
func (self *UserPlaneMode) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *UserPlaneMode) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type VerticalAccuracyCode struct {
  Value uint64
}
func (self *VerticalAccuracyCode) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(128, 7, 0, 0)
}
func (self * VerticalAccuracyCode) Pack(st *Stream){
    st.formatf_Integer(self.Value, 128, 7, 0, 0)
}
func (self *RANAPPDU)Unpack(stream *Stream) {
    //coptions := []string{"initiatingMessage","successfulOutcome","unsuccessfulOutcome","outcome"}
    choice := stream.get_choice(2, 1, 4)
    choice_len := 0
    choice_loc := 0
    if choice >= 4 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in RANAPPDU\n", choice, choice_len)
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
    } else if choice == 3 { //ch2
        self.Outcome = &Outcome{}//cho6
        self.Outcome.Unpack(stream)
    }//end of if else

    if choice >= 4 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * RANAPPDU) Pack(stream *Stream) {
    if self.InitiatingMessage != nil {
        stream.set_choice(0, 2, 1, 4)
        self.InitiatingMessage.Pack(stream)//2
    } else if self.SuccessfulOutcome != nil {
        stream.set_choice(1, 2, 1, 4)
        self.SuccessfulOutcome.Pack(stream)//2
    } else if self.UnsuccessfulOutcome != nil {
        stream.set_choice(2, 2, 1, 4)
        self.UnsuccessfulOutcome.Pack(stream)//2
    } else if self.Outcome != nil {
        stream.set_choice(3, 2, 1, 4)
        self.Outcome.Pack(stream)//2
    }

}
type RANAPPDU struct { //[{'type': 'InitiatingMessage', 'name': 'initiatingMessage'}, {'type': 'SuccessfulOutcome', 'name': 'successfulOutcome'}, {'type': 'UnsuccessfulOutcome', 'name': 'unsuccessfulOutcome'}, {'type': 'Outcome', 'name': 'outcome'}, None]
    InitiatingMessage *InitiatingMessage
    SuccessfulOutcome *SuccessfulOutcome
    UnsuccessfulOutcome *UnsuccessfulOutcome
    Outcome *Outcome
} // RANAPPDU

type InitiatingMessage struct { // [{'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'RANAP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'RANAP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode RANAPELEMENTARYPROCEDUREprocedureCode
    Criticality RANAPELEMENTARYPROCEDUREcriticality
    Value RANAPELEMENTARYPROCEDUREInitiatingMessage
}

func (self * InitiatingMessage) Unpack(stream *Stream) {
    //table {'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_RANAPELEMENTARYPROCEDURES[key].INITIATINGMESSAGE
    self.Value = out.(RANAPELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(RANAPELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'RANAP-ELEMENTARY-PROCEDURE.&InitiatingMessage', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * InitiatingMessage) Pack(stream *Stream) {
    //table {'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_RANAPELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(RANAPELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type SuccessfulOutcome struct { // [{'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'RANAP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'RANAP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode RANAPELEMENTARYPROCEDUREprocedureCode
    Criticality RANAPELEMENTARYPROCEDUREcriticality
    Value RANAPELEMENTARYPROCEDURESuccessfulOutcome
}

func (self * SuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_RANAPELEMENTARYPROCEDURES[key].SUCCESSFULOUTCOME
    self.Value = out.(RANAPELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(RANAPELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'RANAP-ELEMENTARY-PROCEDURE.&SuccessfulOutcome', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * SuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_RANAPELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(RANAPELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type UnsuccessfulOutcome struct { // [{'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'RANAP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'RANAP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode RANAPELEMENTARYPROCEDUREprocedureCode
    Criticality RANAPELEMENTARYPROCEDUREcriticality
    Value RANAPELEMENTARYPROCEDUREUnsuccessfulOutcome
}

func (self * UnsuccessfulOutcome) Unpack(stream *Stream) {
    //table {'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_RANAPELEMENTARYPROCEDURES[key].UNSUCCESSFULOUTCOME
    self.Value = out.(RANAPELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(RANAPELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'RANAP-ELEMENTARY-PROCEDURE.&UnsuccessfulOutcome', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * UnsuccessfulOutcome) Pack(stream *Stream) {
    //table {'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_RANAPELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(RANAPELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type Outcome struct { // [{'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}, {'type': 'RANAP-ELEMENTARY-PROCEDURE.&criticality', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'criticality'}, {'type': 'RANAP-ELEMENTARY-PROCEDURE.&Outcome', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}]
    ProcedureCode RANAPELEMENTARYPROCEDUREprocedureCode
    Criticality RANAPELEMENTARYPROCEDUREcriticality
    Value RANAPELEMENTARYPROCEDUREOutcome
}

func (self * Outcome) Unpack(stream *Stream) {
    //table {'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    var _len int = 0
    self.ProcedureCode.Unpack(stream)//p5
    key := self.ProcedureCode
    //log.Debug("%+v\n",key)
    self.Criticality.Unpack(stream)//p9
    _len += stream.parse_len(0)
    location := stream.get_location()
    out := table_RANAPELEMENTARYPROCEDURES[key].OUTCOME
    self.Value = out.(RANAPELEMENTARYPROCEDURE_IF).createOT()
    self.Value.(RANAPELEMENTARYPROCEDURE_IF).UnpackOT(stream, 0)
    //log.Debug("%+v\n",self.Value)
    //{'type': 'RANAP-ELEMENTARY-PROCEDURE.&Outcome', 'table': ['RANAP-ELEMENTARY-PROCEDURES', ['procedureCode']], 'name': 'value'}
    stream.set_location(location, _len)
    return
}

func (self * Outcome) Pack(stream *Stream) {
    //table {'type': 'RANAP-ELEMENTARY-PROCEDURE.&procedureCode', 'table': {'type': 'RANAP-ELEMENTARY-PROCEDURES'}, 'name': 'procedureCode'}
    key := self.ProcedureCode
    //cobj = table_RANAPELEMENTARYPROCEDURES.get_obj(key)
    //fmt.Printf("formating item value %+v\n", self)
    self.ProcedureCode.Pack(stream)
    self.Criticality.Pack(stream)
    location := stream.reserve_len()
    self.Value.(RANAPELEMENTARYPROCEDURE_IF).PackOT(stream, key)
    stream.set_len(location)
}//end

type RANAPPROTOCOLIES struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type RANAPPROTOCOLIESid ProtocolIEID
func (self *RANAPPROTOCOLIESid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *RANAPPROTOCOLIESid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = RANAPPROTOCOLIESid(val)
}
type RANAPPROTOCOLIEScriticality Criticality
func (self *RANAPPROTOCOLIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *RANAPPROTOCOLIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = RANAPPROTOCOLIEScriticality(val)
}
type RANAPPROTOCOLIESValue interface{}
type RANAPPROTOCOLIESpresence Presence
func (self *RANAPPROTOCOLIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *RANAPPROTOCOLIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = RANAPPROTOCOLIESpresence(val)
}

type RANAPPROTOCOLIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type RANAPPROTOCOLIESPAIR struct {//CLASS {'members': [{'type': 'ProtocolIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&firstCriticality'}, {'type': 'OpenType', 'name': '&FirstValue'}, {'type': 'Criticality', 'name': '&secondCriticality'}, {'type': 'OpenType', 'name': '&SecondValue'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'FIRST CRITICALITY', 'FIRST TYPE', 'SECOND CRITICALITY', 'SECOND TYPE', 'PRESENCE'], 'with-type': ['FIRST TYPE', 'SECOND TYPE']}], 'alt-type': {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'FIRST CRITICALITY': {'type': 'Criticality', 'name': '&firstCriticality'}, 'FIRST TYPE': {'type': 'OpenType', 'name': '&FirstValue'}, 'SECOND CRITICALITY': {'type': 'Criticality', 'name': '&secondCriticality'}, 'SECOND TYPE': {'type': 'OpenType', 'name': '&SecondValue'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolIE-ID', 'name': '&id'}}
    ID ProtocolIEID
    FIRSTCRITICALITY Criticality
    FIRSTTYPE interface{}
    SECONDCRITICALITY Criticality
    SECONDTYPE interface{}
    PRESENCE Presence
}
type RANAPPROTOCOLIESPAIRid ProtocolIEID
func (self *RANAPPROTOCOLIESPAIRid) Pack(st *Stream) {
    ieID := ProtocolIEID(*self)
    ieID.Pack(st)
}
func (self *RANAPPROTOCOLIESPAIRid) Unpack(st *Stream) {
    val := ProtocolIEID{}
    val.Unpack(st)
    *self = RANAPPROTOCOLIESPAIRid(val)
}
type RANAPPROTOCOLIESPAIRfirstCriticality Criticality
func (self *RANAPPROTOCOLIESPAIRfirstCriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *RANAPPROTOCOLIESPAIRfirstCriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = RANAPPROTOCOLIESPAIRfirstCriticality(val)
}
type RANAPPROTOCOLIESPAIRFirstValue interface{}
type RANAPPROTOCOLIESPAIRsecondCriticality Criticality
func (self *RANAPPROTOCOLIESPAIRsecondCriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *RANAPPROTOCOLIESPAIRsecondCriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = RANAPPROTOCOLIESPAIRsecondCriticality(val)
}
type RANAPPROTOCOLIESPAIRSecondValue interface{}
type RANAPPROTOCOLIESPAIRpresence Presence
func (self *RANAPPROTOCOLIESPAIRpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *RANAPPROTOCOLIESPAIRpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = RANAPPROTOCOLIESPAIRpresence(val)
}

type RANAPPROTOCOLIESPAIR_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type RANAPPROTOCOLEXTENSION struct {//CLASS {'members': [{'type': 'ProtocolExtensionID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Extension'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'EXTENSION', 'PRESENCE'], 'with-type': ['EXTENSION']}], 'alt-type': {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'ProtocolExtensionID', 'name': '&id'}}
    ID ProtocolExtensionID
    CRITICALITY Criticality
    EXTENSION interface{}
    PRESENCE Presence
}
type RANAPPROTOCOLEXTENSIONid ProtocolExtensionID
func (self *RANAPPROTOCOLEXTENSIONid) Pack(st *Stream) {
    ieID := ProtocolExtensionID(*self)
    ieID.Pack(st)
}
func (self *RANAPPROTOCOLEXTENSIONid) Unpack(st *Stream) {
    val := ProtocolExtensionID{}
    val.Unpack(st)
    *self = RANAPPROTOCOLEXTENSIONid(val)
}
type RANAPPROTOCOLEXTENSIONcriticality Criticality
func (self *RANAPPROTOCOLEXTENSIONcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *RANAPPROTOCOLEXTENSIONcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = RANAPPROTOCOLEXTENSIONcriticality(val)
}
type RANAPPROTOCOLEXTENSIONExtension interface{}
type RANAPPROTOCOLEXTENSIONpresence Presence
func (self *RANAPPROTOCOLEXTENSIONpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *RANAPPROTOCOLEXTENSIONpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = RANAPPROTOCOLEXTENSIONpresence(val)
}

type RANAPPROTOCOLEXTENSION_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type RANAPPRIVATEIES struct {//CLASS {'members': [{'type': 'PrivateIE-ID', 'name': '&id'}, {'type': 'Criticality', 'name': '&criticality'}, {'type': 'OpenType', 'name': '&Value'}, {'type': 'Presence', 'name': '&presence'}], 'with-members': [{'with-order': ['ID', 'CRITICALITY', 'TYPE', 'PRESENCE'], 'with-type': ['TYPE']}], 'alt-type': {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}}, 'id-type': {'type': 'PrivateIE-ID', 'name': '&id'}}
    ID PrivateIEID
    CRITICALITY Criticality
    TYPE interface{}
    PRESENCE Presence
}
type RANAPPRIVATEIESid PrivateIEID
func (self *RANAPPRIVATEIESid) Pack(st *Stream) {
    ieID := PrivateIEID(*self)
    ieID.Pack(st)
}
func (self *RANAPPRIVATEIESid) Unpack(st *Stream) {
    val := PrivateIEID{}
    val.Unpack(st)
    *self = RANAPPRIVATEIESid(val)
}
type RANAPPRIVATEIEScriticality Criticality
func (self *RANAPPRIVATEIEScriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *RANAPPRIVATEIEScriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = RANAPPRIVATEIEScriticality(val)
}
type RANAPPRIVATEIESValue interface{}
type RANAPPRIVATEIESpresence Presence
func (self *RANAPPRIVATEIESpresence) Pack(st *Stream) {
    ieID := Presence(*self)
    ieID.Pack(st)
}
func (self *RANAPPRIVATEIESpresence) Unpack(st *Stream) {
    val := Presence{}
    val.Unpack(st)
    *self = RANAPPRIVATEIESpresence(val)
}

type RANAPPRIVATEIES_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type RANAPELEMENTARYPROCEDURE struct {//CLASS {'members': [{'type': 'OpenType', 'name': '&InitiatingMessage'}, {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, {'type': 'OpenType', 'name': '&Outcome'}, {'type': 'ProcedureCode', 'name': '&procedureCode'}, {'type': 'Criticality', 'name': '&criticality'}], 'with-members': [{'with-order': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME', 'OUTCOME', 'PROCEDURE CODE', 'CRITICALITY'], 'with-type': ['INITIATING MESSAGE', 'SUCCESSFUL OUTCOME', 'UNSUCCESSFUL OUTCOME', 'OUTCOME']}], 'alt-type': {'INITIATING MESSAGE': {'type': 'OpenType', 'name': '&InitiatingMessage'}, 'SUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&SuccessfulOutcome'}, 'UNSUCCESSFUL OUTCOME': {'type': 'OpenType', 'name': '&UnsuccessfulOutcome'}, 'OUTCOME': {'type': 'OpenType', 'name': '&Outcome'}, 'PROCEDURE CODE': {'type': 'ProcedureCode', 'name': '&procedureCode'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}}, 'id-type': {'type': 'ProcedureCode', 'name': '&procedureCode'}}
    INITIATINGMESSAGE interface{}
    SUCCESSFULOUTCOME interface{}
    UNSUCCESSFULOUTCOME interface{}
    OUTCOME interface{}
    PROCEDURECODE ProcedureCode
    CRITICALITY Criticality
}
type RANAPELEMENTARYPROCEDUREInitiatingMessage interface{}
type RANAPELEMENTARYPROCEDURESuccessfulOutcome interface{}
type RANAPELEMENTARYPROCEDUREUnsuccessfulOutcome interface{}
type RANAPELEMENTARYPROCEDUREOutcome interface{}
type RANAPELEMENTARYPROCEDUREprocedureCode ProcedureCode
func (self *RANAPELEMENTARYPROCEDUREprocedureCode) Pack(st *Stream) {
    ieID := ProcedureCode(*self)
    ieID.Pack(st)
}
func (self *RANAPELEMENTARYPROCEDUREprocedureCode) Unpack(st *Stream) {
    val := ProcedureCode{}
    val.Unpack(st)
    *self = RANAPELEMENTARYPROCEDUREprocedureCode(val)
}
type RANAPELEMENTARYPROCEDUREcriticality Criticality
func (self *RANAPELEMENTARYPROCEDUREcriticality) Pack(st *Stream) {
    ieID := Criticality(*self)
    ieID.Pack(st)
}
func (self *RANAPELEMENTARYPROCEDUREcriticality) Unpack(st *Stream) {
    val := Criticality{}
    val.Unpack(st)
    *self = RANAPELEMENTARYPROCEDUREcriticality(val)
}

type RANAPELEMENTARYPROCEDURE_IF interface {
    PackOT(st *Stream, id interface{})
    UnpackOT(st *Stream, id interface{})
    createOT() interface{}
    GetOT(id interface{}) bool
    GetIECount() int
}
type IuReleaseCommandIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, None]}
   Cause  Cause
   list []interface{}
}
func (self *IuReleaseCommandIEs)createOT() interface{}{
    return nil
}
var table_IuReleaseCommandIEs = make(map[int]*RANAPPROTOCOLIES)

var order_IuReleaseCommandIEs = make([]int, 1)

func (self *IuReleaseCommandIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   return count//ObjSet
}
func (self *IuReleaseCommandIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        return true //self.Cause
   }
   return false//ObjSet
}
func (self *IuReleaseCommandIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
   }
}
func (self *IuReleaseCommandIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Pack(st)
      default:
      break
   }
}
func init() {
table_IuReleaseCommandIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_IuReleaseCommandIEs[0] = 4
   }

type IuReleaseCommandExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *IuReleaseCommandExtensions)createOT() interface{}{
    return nil
}
var table_IuReleaseCommandExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_IuReleaseCommandExtensions = make([]int, 0)

type IuReleaseCompleteIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-DataVolumeReportList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-DataVolumeReportList', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-ReleasedList-IuRelComp', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ReleasedList-IuRelComp', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   RABDataVolumeReportList  *RABDataVolumeReportList
   RABReleasedListIuRelComp  *RABReleasedListIuRelComp
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *IuReleaseCompleteIEs)createOT() interface{}{
    return nil
}
var table_IuReleaseCompleteIEs = make(map[int]*RANAPPROTOCOLIES)

var order_IuReleaseCompleteIEs = make([]int, 3)

func (self *IuReleaseCompleteIEs) GetIECount() int{
   count := 0
   if self.RABDataVolumeReportList != nil { count += 1 }
   if self.RABReleasedListIuRelComp != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *IuReleaseCompleteIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 31: //RABDataVolumeReportList
        if self.RABDataVolumeReportList != nil { return true }
      case 44: //RABReleasedListIuRelComp
        if self.RABReleasedListIuRelComp != nil { return true }
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *IuReleaseCompleteIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 31: //RABDataVolumeReportList
        self.RABDataVolumeReportList = &RABDataVolumeReportList{}
        self.RABDataVolumeReportList.Unpack(st)
        self.list = append(self.list, self.RABDataVolumeReportList)
      case 44: //RABReleasedListIuRelComp
        self.RABReleasedListIuRelComp = &RABReleasedListIuRelComp{}
        self.RABReleasedListIuRelComp.Unpack(st)
        self.list = append(self.list, self.RABReleasedListIuRelComp)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *IuReleaseCompleteIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 31: //RABDataVolumeReportList
        if self.RABDataVolumeReportList != nil {self.RABDataVolumeReportList.Pack(st)}
      case 44: //RABReleasedListIuRelComp
        if self.RABReleasedListIuRelComp != nil {self.RABReleasedListIuRelComp.Pack(st)}
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_IuReleaseCompleteIEs[31] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABDataVolumeReportList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABDataVolumeReportList{}, PRESENCE:Presence{Presenceoptional}, }
order_IuReleaseCompleteIEs[0] = 31
table_IuReleaseCompleteIEs[44] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABReleasedListIuRelComp}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABReleasedListIuRelComp{}, PRESENCE:Presence{Presenceoptional}, }
order_IuReleaseCompleteIEs[1] = 44
table_IuReleaseCompleteIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_IuReleaseCompleteIEs[2] = 9
   }

type RABDataVolumeReportItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-DataVolumeReportItem', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-DataVolumeReportItem', 'PRESENCE': 'mandatory'}, None]}
   RABDataVolumeReportItem  RABDataVolumeReportItem
   list []interface{}
}
func (self *RABDataVolumeReportItemIEs)createOT() interface{}{
    return nil
}
var table_RABDataVolumeReportItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABDataVolumeReportItemIEs = make([]int, 1)

func (self *RABDataVolumeReportItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABDataVolumeReportItem
   return count//ObjSet
}
func (self *RABDataVolumeReportItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 30: //RABDataVolumeReportItem
        return true //self.RABDataVolumeReportItem
   }
   return false//ObjSet
}
func (self *RABDataVolumeReportItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 30: //RABDataVolumeReportItem
        self.RABDataVolumeReportItem.Unpack(st)
        self.list = append(self.list, &self.RABDataVolumeReportItem)
   }
}
func (self *RABDataVolumeReportItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 30: //RABDataVolumeReportItem
        self.RABDataVolumeReportItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABDataVolumeReportItemIEs[30] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABDataVolumeReportItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABDataVolumeReportItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABDataVolumeReportItemIEs[0] = 30
   }

type RABDataVolumeReportItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABDataVolumeReportItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABDataVolumeReportItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABDataVolumeReportItemExtIEs = make([]int, 0)

type RABReleasedItemIuRelCompIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ReleasedItem-IuRelComp', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ReleasedItem-IuRelComp', 'PRESENCE': 'mandatory'}, None]}
   RABReleasedItemIuRelComp  RABReleasedItemIuRelComp
   list []interface{}
}
func (self *RABReleasedItemIuRelCompIEs)createOT() interface{}{
    return nil
}
var table_RABReleasedItemIuRelCompIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABReleasedItemIuRelCompIEs = make([]int, 1)

func (self *RABReleasedItemIuRelCompIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABReleasedItemIuRelComp
   return count//ObjSet
}
func (self *RABReleasedItemIuRelCompIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 87: //RABReleasedItemIuRelComp
        return true //self.RABReleasedItemIuRelComp
   }
   return false//ObjSet
}
func (self *RABReleasedItemIuRelCompIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 87: //RABReleasedItemIuRelComp
        self.RABReleasedItemIuRelComp.Unpack(st)
        self.list = append(self.list, &self.RABReleasedItemIuRelComp)
   }
}
func (self *RABReleasedItemIuRelCompIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 87: //RABReleasedItemIuRelComp
        self.RABReleasedItemIuRelComp.Pack(st)
      default:
      break
   }
}
func init() {
table_RABReleasedItemIuRelCompIEs[87] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABReleasedItemIuRelComp}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABReleasedItemIuRelComp{}, PRESENCE:Presence{Presencemandatory}, }
order_RABReleasedItemIuRelCompIEs[0] = 87
   }

type RABReleasedItemIuRelCompExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABReleasedItemIuRelCompExtIEs)createOT() interface{}{
    return nil
}
var table_RABReleasedItemIuRelCompExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABReleasedItemIuRelCompExtIEs = make([]int, 0)

type IuReleaseCompleteExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *IuReleaseCompleteExtensions)createOT() interface{}{
    return nil
}
var table_IuReleaseCompleteExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_IuReleaseCompleteExtensions = make([]int, 0)

type RelocationRequiredIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RelocationType', 'CRITICALITY': 'reject', 'TYPE': 'RelocationType', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-SourceID', 'CRITICALITY': 'ignore', 'TYPE': 'SourceID', 'PRESENCE': 'mandatory'}, {'ID': 'id-TargetID', 'CRITICALITY': 'reject', 'TYPE': 'TargetID', 'PRESENCE': 'mandatory'}, {'ID': 'id-ClassmarkInformation2', 'CRITICALITY': 'reject', 'TYPE': 'ClassmarkInformation2', 'PRESENCE': 'conditional'}, {'ID': 'id-ClassmarkInformation3', 'CRITICALITY': 'ignore', 'TYPE': 'ClassmarkInformation3', 'PRESENCE': 'conditional'}, {'ID': 'id-SourceRNC-ToTargetRNC-TransparentContainer', 'CRITICALITY': 'reject', 'TYPE': 'SourceRNC-ToTargetRNC-TransparentContainer', 'PRESENCE': 'conditional'}, {'ID': 'id-OldBSS-ToNewBSS-Information', 'CRITICALITY': 'ignore', 'TYPE': 'OldBSS-ToNewBSS-Information', 'PRESENCE': 'optional'}, None]}
   RelocationType  RelocationType
   Cause  Cause
   SourceID  SourceID
   TargetID  TargetID
   ClassmarkInformation2  ClassmarkInformation2
   ClassmarkInformation3  ClassmarkInformation3
   SourceRNCToTargetRNCTransparentContainer  SourceRNCToTargetRNCTransparentContainer
   OldBSSToNewBSSInformation  *OldBSSToNewBSSInformation
   list []interface{}
}
func (self *RelocationRequiredIEs)createOT() interface{}{
    return nil
}
var table_RelocationRequiredIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RelocationRequiredIEs = make([]int, 8)

func (self *RelocationRequiredIEs) GetIECount() int{
   count := 0
   count +=1 //self.RelocationType
   count +=1 //self.Cause
   count +=1 //self.SourceID
   count +=1 //self.TargetID
   count +=1 //self.ClassmarkInformation2
   count +=1 //self.ClassmarkInformation3
   count +=1 //self.SourceRNCToTargetRNCTransparentContainer
   if self.OldBSSToNewBSSInformation != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationRequiredIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 56: //RelocationType
        return true //self.RelocationType
      case 4: //Cause
        return true //self.Cause
      case 60: //SourceID
        return true //self.SourceID
      case 62: //TargetID
        return true //self.TargetID
      case 7: //ClassmarkInformation2
        return true //self.ClassmarkInformation2
      case 8: //ClassmarkInformation3
        return true //self.ClassmarkInformation3
      case 61: //SourceRNCToTargetRNCTransparentContainer
        return true //self.SourceRNCToTargetRNCTransparentContainer
      case 20: //OldBSSToNewBSSInformation
        if self.OldBSSToNewBSSInformation != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationRequiredIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 56: //RelocationType
        self.RelocationType.Unpack(st)
        self.list = append(self.list, &self.RelocationType)
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 60: //SourceID
        self.SourceID.Unpack(st)
        self.list = append(self.list, &self.SourceID)
      case 62: //TargetID
        self.TargetID.Unpack(st)
        self.list = append(self.list, &self.TargetID)
      case 7: //ClassmarkInformation2
        self.ClassmarkInformation2.Unpack(st)
        self.list = append(self.list, &self.ClassmarkInformation2)
      case 8: //ClassmarkInformation3
        self.ClassmarkInformation3.Unpack(st)
        self.list = append(self.list, &self.ClassmarkInformation3)
      case 61: //SourceRNCToTargetRNCTransparentContainer
        self.SourceRNCToTargetRNCTransparentContainer.Unpack(st)
        self.list = append(self.list, &self.SourceRNCToTargetRNCTransparentContainer)
      case 20: //OldBSSToNewBSSInformation
        self.OldBSSToNewBSSInformation = &OldBSSToNewBSSInformation{}
        self.OldBSSToNewBSSInformation.Unpack(st)
        self.list = append(self.list, self.OldBSSToNewBSSInformation)
   }
}
func (self *RelocationRequiredIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 56: //RelocationType
        self.RelocationType.Pack(st)
      case 4: //Cause
        self.Cause.Pack(st)
      case 60: //SourceID
        self.SourceID.Pack(st)
      case 62: //TargetID
        self.TargetID.Pack(st)
      case 7: //ClassmarkInformation2
        self.ClassmarkInformation2.Pack(st)
      case 8: //ClassmarkInformation3
        self.ClassmarkInformation3.Pack(st)
      case 61: //SourceRNCToTargetRNCTransparentContainer
        self.SourceRNCToTargetRNCTransparentContainer.Pack(st)
      case 20: //OldBSSToNewBSSInformation
        if self.OldBSSToNewBSSInformation != nil {self.OldBSSToNewBSSInformation.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationRequiredIEs[56] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRelocationType}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RelocationType{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationRequiredIEs[0] = 56
table_RelocationRequiredIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationRequiredIEs[1] = 4
table_RelocationRequiredIEs[60] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idSourceID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&SourceID{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationRequiredIEs[2] = 60
table_RelocationRequiredIEs[62] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idTargetID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TargetID{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationRequiredIEs[3] = 62
table_RelocationRequiredIEs[7] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idClassmarkInformation2}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ClassmarkInformation2{}, PRESENCE:Presence{Presenceconditional}, }
order_RelocationRequiredIEs[4] = 7
table_RelocationRequiredIEs[8] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idClassmarkInformation3}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ClassmarkInformation3{}, PRESENCE:Presence{Presenceconditional}, }
order_RelocationRequiredIEs[5] = 8
table_RelocationRequiredIEs[61] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idSourceRNCToTargetRNCTransparentContainer}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SourceRNCToTargetRNCTransparentContainer{}, PRESENCE:Presence{Presenceconditional}, }
order_RelocationRequiredIEs[6] = 61
table_RelocationRequiredIEs[20] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idOldBSSToNewBSSInformation}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&OldBSSToNewBSSInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequiredIEs[7] = 20
   }

type RelocationRequiredExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-GERAN-Classmark', 'CRITICALITY': 'ignore', 'EXTENSION': 'GERAN-Classmark', 'PRESENCE': 'optional'}, None]}
   GERANClassmark  *GERANClassmark
   list []interface{}
}
func (self *RelocationRequiredExtensions)createOT() interface{}{
    return nil
}
var table_RelocationRequiredExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RelocationRequiredExtensions = make([]int, 1)

func (self *RelocationRequiredExtensions) GetIECount() int{
   count := 0
   if self.GERANClassmark != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationRequiredExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 108: //GERANClassmark
        if self.GERANClassmark != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationRequiredExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 108: //GERANClassmark
        self.GERANClassmark = &GERANClassmark{}
        self.GERANClassmark.Unpack(st)
        self.list = append(self.list, self.GERANClassmark)
   }
}
func (self *RelocationRequiredExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 108: //GERANClassmark
        if self.GERANClassmark != nil {self.GERANClassmark.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationRequiredExtensions[108] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGERANClassmark}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GERANClassmark{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequiredExtensions[0] = 108
   }

type RelocationCommandIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-TargetRNC-ToSourceRNC-TransparentContainer', 'CRITICALITY': 'reject', 'TYPE': 'TargetRNC-ToSourceRNC-TransparentContainer', 'PRESENCE': 'optional'}, {'ID': 'id-L3-Information', 'CRITICALITY': 'ignore', 'TYPE': 'L3-Information', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-RelocationReleaseList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-RelocationReleaseList', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-DataForwardingList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-DataForwardingList', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TargetRNCToSourceRNCTransparentContainer  *TargetRNCToSourceRNCTransparentContainer
   L3Information  *L3Information
   RABRelocationReleaseList  *RABRelocationReleaseList
   RABDataForwardingList  *RABDataForwardingList
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *RelocationCommandIEs)createOT() interface{}{
    return nil
}
var table_RelocationCommandIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RelocationCommandIEs = make([]int, 5)

func (self *RelocationCommandIEs) GetIECount() int{
   count := 0
   if self.TargetRNCToSourceRNCTransparentContainer != nil { count += 1 }
   if self.L3Information != nil { count += 1 }
   if self.RABRelocationReleaseList != nil { count += 1 }
   if self.RABDataForwardingList != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationCommandIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 63: //TargetRNCToSourceRNCTransparentContainer
        if self.TargetRNCToSourceRNCTransparentContainer != nil { return true }
      case 14: //L3Information
        if self.L3Information != nil { return true }
      case 46: //RABRelocationReleaseList
        if self.RABRelocationReleaseList != nil { return true }
      case 28: //RABDataForwardingList
        if self.RABDataForwardingList != nil { return true }
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationCommandIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 63: //TargetRNCToSourceRNCTransparentContainer
        self.TargetRNCToSourceRNCTransparentContainer = &TargetRNCToSourceRNCTransparentContainer{}
        self.TargetRNCToSourceRNCTransparentContainer.Unpack(st)
        self.list = append(self.list, self.TargetRNCToSourceRNCTransparentContainer)
      case 14: //L3Information
        self.L3Information = &L3Information{}
        self.L3Information.Unpack(st)
        self.list = append(self.list, self.L3Information)
      case 46: //RABRelocationReleaseList
        self.RABRelocationReleaseList = &RABRelocationReleaseList{}
        self.RABRelocationReleaseList.Unpack(st)
        self.list = append(self.list, self.RABRelocationReleaseList)
      case 28: //RABDataForwardingList
        self.RABDataForwardingList = &RABDataForwardingList{}
        self.RABDataForwardingList.Unpack(st)
        self.list = append(self.list, self.RABDataForwardingList)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *RelocationCommandIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 63: //TargetRNCToSourceRNCTransparentContainer
        if self.TargetRNCToSourceRNCTransparentContainer != nil {self.TargetRNCToSourceRNCTransparentContainer.Pack(st)}
      case 14: //L3Information
        if self.L3Information != nil {self.L3Information.Pack(st)}
      case 46: //RABRelocationReleaseList
        if self.RABRelocationReleaseList != nil {self.RABRelocationReleaseList.Pack(st)}
      case 28: //RABDataForwardingList
        if self.RABDataForwardingList != nil {self.RABDataForwardingList.Pack(st)}
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationCommandIEs[63] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idTargetRNCToSourceRNCTransparentContainer}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&TargetRNCToSourceRNCTransparentContainer{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationCommandIEs[0] = 63
table_RelocationCommandIEs[14] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idL3Information}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&L3Information{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationCommandIEs[1] = 14
table_RelocationCommandIEs[46] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABRelocationReleaseList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABRelocationReleaseList{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationCommandIEs[2] = 46
table_RelocationCommandIEs[28] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABDataForwardingList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABDataForwardingList{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationCommandIEs[3] = 28
table_RelocationCommandIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationCommandIEs[4] = 9
   }

type RABRelocationReleaseItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-RelocationReleaseItem', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-RelocationReleaseItem', 'PRESENCE': 'mandatory'}, None]}
   RABRelocationReleaseItem  RABRelocationReleaseItem
   list []interface{}
}
func (self *RABRelocationReleaseItemIEs)createOT() interface{}{
    return nil
}
var table_RABRelocationReleaseItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABRelocationReleaseItemIEs = make([]int, 1)

func (self *RABRelocationReleaseItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABRelocationReleaseItem
   return count//ObjSet
}
func (self *RABRelocationReleaseItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 45: //RABRelocationReleaseItem
        return true //self.RABRelocationReleaseItem
   }
   return false//ObjSet
}
func (self *RABRelocationReleaseItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 45: //RABRelocationReleaseItem
        self.RABRelocationReleaseItem.Unpack(st)
        self.list = append(self.list, &self.RABRelocationReleaseItem)
   }
}
func (self *RABRelocationReleaseItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 45: //RABRelocationReleaseItem
        self.RABRelocationReleaseItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABRelocationReleaseItemIEs[45] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABRelocationReleaseItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABRelocationReleaseItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABRelocationReleaseItemIEs[0] = 45
   }

type RABRelocationReleaseItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABRelocationReleaseItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABRelocationReleaseItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABRelocationReleaseItemExtIEs = make([]int, 0)

type RABDataForwardingItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-DataForwardingItem', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-DataForwardingItem', 'PRESENCE': 'mandatory'}, None]}
   RABDataForwardingItem  RABDataForwardingItem
   list []interface{}
}
func (self *RABDataForwardingItemIEs)createOT() interface{}{
    return nil
}
var table_RABDataForwardingItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABDataForwardingItemIEs = make([]int, 1)

func (self *RABDataForwardingItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABDataForwardingItem
   return count//ObjSet
}
func (self *RABDataForwardingItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 26: //RABDataForwardingItem
        return true //self.RABDataForwardingItem
   }
   return false//ObjSet
}
func (self *RABDataForwardingItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 26: //RABDataForwardingItem
        self.RABDataForwardingItem.Unpack(st)
        self.list = append(self.list, &self.RABDataForwardingItem)
   }
}
func (self *RABDataForwardingItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 26: //RABDataForwardingItem
        self.RABDataForwardingItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABDataForwardingItemIEs[26] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABDataForwardingItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABDataForwardingItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABDataForwardingItemIEs[0] = 26
   }

type RABDataForwardingItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-TransportLayerAddress', 'CRITICALITY': 'ignore', 'EXTENSION': 'TransportLayerAddress', 'PRESENCE': 'optional'}, {'ID': 'id-IuTransportAssociation', 'CRITICALITY': 'ignore', 'EXTENSION': 'IuTransportAssociation', 'PRESENCE': 'optional'}, None]}
   TransportLayerAddress  *TransportLayerAddress
   IuTransportAssociation  *IuTransportAssociation
   list []interface{}
}
func (self *RABDataForwardingItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABDataForwardingItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABDataForwardingItemExtIEs = make([]int, 2)

func (self *RABDataForwardingItemExtIEs) GetIECount() int{
   count := 0
   if self.TransportLayerAddress != nil { count += 1 }
   if self.IuTransportAssociation != nil { count += 1 }
   return count//ObjSet
}
func (self *RABDataForwardingItemExtIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 67: //TransportLayerAddress
        if self.TransportLayerAddress != nil { return true }
      case 13: //IuTransportAssociation
        if self.IuTransportAssociation != nil { return true }
   }
   return false//ObjSet
}
func (self *RABDataForwardingItemExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 67: //TransportLayerAddress
        self.TransportLayerAddress = &TransportLayerAddress{}
        self.TransportLayerAddress.Unpack(st)
        self.list = append(self.list, self.TransportLayerAddress)
      case 13: //IuTransportAssociation
        self.IuTransportAssociation = &IuTransportAssociation{}
        self.IuTransportAssociation.Unpack(st)
        self.list = append(self.list, self.IuTransportAssociation)
   }
}
func (self *RABDataForwardingItemExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 67: //TransportLayerAddress
        if self.TransportLayerAddress != nil {self.TransportLayerAddress.Pack(st)}
      case 13: //IuTransportAssociation
        if self.IuTransportAssociation != nil {self.IuTransportAssociation.Pack(st)}
      default:
      break
   }
}
func init() {
table_RABDataForwardingItemExtIEs[67] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idTransportLayerAddress}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&TransportLayerAddress{}, PRESENCE:Presence{Presenceoptional}, }
order_RABDataForwardingItemExtIEs[0] = 67
table_RABDataForwardingItemExtIEs[13] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idIuTransportAssociation}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&IuTransportAssociation{}, PRESENCE:Presence{Presenceoptional}, }
order_RABDataForwardingItemExtIEs[1] = 13
   }

type RelocationCommandExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-InterSystemInformation-TransparentContainer', 'CRITICALITY': 'ignore', 'EXTENSION': 'InterSystemInformation-TransparentContainer', 'PRESENCE': 'optional'}, None]}
   InterSystemInformationTransparentContainer  *InterSystemInformationTransparentContainer
   list []interface{}
}
func (self *RelocationCommandExtensions)createOT() interface{}{
    return nil
}
var table_RelocationCommandExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RelocationCommandExtensions = make([]int, 1)

func (self *RelocationCommandExtensions) GetIECount() int{
   count := 0
   if self.InterSystemInformationTransparentContainer != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationCommandExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 99: //InterSystemInformationTransparentContainer
        if self.InterSystemInformationTransparentContainer != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationCommandExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 99: //InterSystemInformationTransparentContainer
        self.InterSystemInformationTransparentContainer = &InterSystemInformationTransparentContainer{}
        self.InterSystemInformationTransparentContainer.Unpack(st)
        self.list = append(self.list, self.InterSystemInformationTransparentContainer)
   }
}
func (self *RelocationCommandExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 99: //InterSystemInformationTransparentContainer
        if self.InterSystemInformationTransparentContainer != nil {self.InterSystemInformationTransparentContainer.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationCommandExtensions[99] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idInterSystemInformationTransparentContainer}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&InterSystemInformationTransparentContainer{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationCommandExtensions[0] = 99
   }

type RelocationPreparationFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *RelocationPreparationFailureIEs)createOT() interface{}{
    return nil
}
var table_RelocationPreparationFailureIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RelocationPreparationFailureIEs = make([]int, 2)

func (self *RelocationPreparationFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationPreparationFailureIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        return true //self.Cause
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationPreparationFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *RelocationPreparationFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Pack(st)
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationPreparationFailureIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationPreparationFailureIEs[0] = 4
table_RelocationPreparationFailureIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationPreparationFailureIEs[1] = 9
   }

type RelocationPreparationFailureExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-InterSystemInformation-TransparentContainer', 'CRITICALITY': 'ignore', 'EXTENSION': 'InterSystemInformation-TransparentContainer', 'PRESENCE': 'optional'}, None]}
   InterSystemInformationTransparentContainer  *InterSystemInformationTransparentContainer
   list []interface{}
}
func (self *RelocationPreparationFailureExtensions)createOT() interface{}{
    return nil
}
var table_RelocationPreparationFailureExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RelocationPreparationFailureExtensions = make([]int, 1)

func (self *RelocationPreparationFailureExtensions) GetIECount() int{
   count := 0
   if self.InterSystemInformationTransparentContainer != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationPreparationFailureExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 99: //InterSystemInformationTransparentContainer
        if self.InterSystemInformationTransparentContainer != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationPreparationFailureExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 99: //InterSystemInformationTransparentContainer
        self.InterSystemInformationTransparentContainer = &InterSystemInformationTransparentContainer{}
        self.InterSystemInformationTransparentContainer.Unpack(st)
        self.list = append(self.list, self.InterSystemInformationTransparentContainer)
   }
}
func (self *RelocationPreparationFailureExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 99: //InterSystemInformationTransparentContainer
        if self.InterSystemInformationTransparentContainer != nil {self.InterSystemInformationTransparentContainer.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationPreparationFailureExtensions[99] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idInterSystemInformationTransparentContainer}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&InterSystemInformationTransparentContainer{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationPreparationFailureExtensions[0] = 99
   }

type RelocationRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-PermanentNAS-UE-ID', 'CRITICALITY': 'ignore', 'TYPE': 'PermanentNAS-UE-ID', 'PRESENCE': 'optional'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'reject', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-SourceRNC-ToTargetRNC-TransparentContainer', 'CRITICALITY': 'reject', 'TYPE': 'SourceRNC-ToTargetRNC-TransparentContainer', 'PRESENCE': 'mandatory'}, {'ID': 'id-RAB-SetupList-RelocReq', 'CRITICALITY': 'reject', 'TYPE': 'RAB-SetupList-RelocReq', 'PRESENCE': 'optional'}, {'ID': 'id-IntegrityProtectionInformation', 'CRITICALITY': 'ignore', 'TYPE': 'IntegrityProtectionInformation', 'PRESENCE': 'optional'}, {'ID': 'id-EncryptionInformation', 'CRITICALITY': 'ignore', 'TYPE': 'EncryptionInformation', 'PRESENCE': 'optional'}, {'ID': 'id-IuSigConId', 'CRITICALITY': 'ignore', 'TYPE': 'IuSignallingConnectionIdentifier', 'PRESENCE': 'mandatory'}, None]}
   PermanentNASUEID  *PermanentNASUEID
   Cause  Cause
   CNDomainIndicator  CNDomainIndicator
   SourceRNCToTargetRNCTransparentContainer  SourceRNCToTargetRNCTransparentContainer
   RABSetupListRelocReq  *RABSetupListRelocReq
   IntegrityProtectionInformation  *IntegrityProtectionInformation
   EncryptionInformation  *EncryptionInformation
   IuSigConId  IuSignallingConnectionIdentifier
   list []interface{}
}
func (self *RelocationRequestIEs)createOT() interface{}{
    return nil
}
var table_RelocationRequestIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RelocationRequestIEs = make([]int, 8)

func (self *RelocationRequestIEs) GetIECount() int{
   count := 0
   if self.PermanentNASUEID != nil { count += 1 }
   count +=1 //self.Cause
   count +=1 //self.CNDomainIndicator
   count +=1 //self.SourceRNCToTargetRNCTransparentContainer
   if self.RABSetupListRelocReq != nil { count += 1 }
   if self.IntegrityProtectionInformation != nil { count += 1 }
   if self.EncryptionInformation != nil { count += 1 }
   count +=1 //self.IuSigConId
   return count//ObjSet
}
func (self *RelocationRequestIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 23: //PermanentNASUEID
        if self.PermanentNASUEID != nil { return true }
      case 4: //Cause
        return true //self.Cause
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 61: //SourceRNCToTargetRNCTransparentContainer
        return true //self.SourceRNCToTargetRNCTransparentContainer
      case 49: //RABSetupListRelocReq
        if self.RABSetupListRelocReq != nil { return true }
      case 12: //IntegrityProtectionInformation
        if self.IntegrityProtectionInformation != nil { return true }
      case 11: //EncryptionInformation
        if self.EncryptionInformation != nil { return true }
      case 79: //IuSigConId
        return true //self.IuSigConId
   }
   return false//ObjSet
}
func (self *RelocationRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 23: //PermanentNASUEID
        self.PermanentNASUEID = &PermanentNASUEID{}
        self.PermanentNASUEID.Unpack(st)
        self.list = append(self.list, self.PermanentNASUEID)
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 61: //SourceRNCToTargetRNCTransparentContainer
        self.SourceRNCToTargetRNCTransparentContainer.Unpack(st)
        self.list = append(self.list, &self.SourceRNCToTargetRNCTransparentContainer)
      case 49: //RABSetupListRelocReq
        self.RABSetupListRelocReq = &RABSetupListRelocReq{}
        self.RABSetupListRelocReq.Unpack(st)
        self.list = append(self.list, self.RABSetupListRelocReq)
      case 12: //IntegrityProtectionInformation
        self.IntegrityProtectionInformation = &IntegrityProtectionInformation{}
        self.IntegrityProtectionInformation.Unpack(st)
        self.list = append(self.list, self.IntegrityProtectionInformation)
      case 11: //EncryptionInformation
        self.EncryptionInformation = &EncryptionInformation{}
        self.EncryptionInformation.Unpack(st)
        self.list = append(self.list, self.EncryptionInformation)
      case 79: //IuSigConId
        self.IuSigConId.Unpack(st)
        self.list = append(self.list, &self.IuSigConId)
   }
}
func (self *RelocationRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 23: //PermanentNASUEID
        if self.PermanentNASUEID != nil {self.PermanentNASUEID.Pack(st)}
      case 4: //Cause
        self.Cause.Pack(st)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 61: //SourceRNCToTargetRNCTransparentContainer
        self.SourceRNCToTargetRNCTransparentContainer.Pack(st)
      case 49: //RABSetupListRelocReq
        if self.RABSetupListRelocReq != nil {self.RABSetupListRelocReq.Pack(st)}
      case 12: //IntegrityProtectionInformation
        if self.IntegrityProtectionInformation != nil {self.IntegrityProtectionInformation.Pack(st)}
      case 11: //EncryptionInformation
        if self.EncryptionInformation != nil {self.EncryptionInformation.Pack(st)}
      case 79: //IuSigConId
        self.IuSigConId.Pack(st)
      default:
      break
   }
}
func init() {
table_RelocationRequestIEs[23] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idPermanentNASUEID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PermanentNASUEID{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestIEs[0] = 23
table_RelocationRequestIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationRequestIEs[1] = 4
table_RelocationRequestIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationRequestIEs[2] = 3
table_RelocationRequestIEs[61] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idSourceRNCToTargetRNCTransparentContainer}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&SourceRNCToTargetRNCTransparentContainer{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationRequestIEs[3] = 61
table_RelocationRequestIEs[49] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABSetupListRelocReq}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RABSetupListRelocReq{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestIEs[4] = 49
table_RelocationRequestIEs[12] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idIntegrityProtectionInformation}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&IntegrityProtectionInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestIEs[5] = 12
table_RelocationRequestIEs[11] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idEncryptionInformation}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&EncryptionInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestIEs[6] = 11
table_RelocationRequestIEs[79] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idIuSigConId}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&IuSignallingConnectionIdentifier{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationRequestIEs[7] = 79
   }

type RABSetupItemRelocReqIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-SetupItem-RelocReq', 'CRITICALITY': 'reject', 'TYPE': 'RAB-SetupItem-RelocReq', 'PRESENCE': 'mandatory'}, None]}
   RABSetupItemRelocReq  RABSetupItemRelocReq
   list []interface{}
}
func (self *RABSetupItemRelocReqIEs)createOT() interface{}{
    return nil
}
var table_RABSetupItemRelocReqIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABSetupItemRelocReqIEs = make([]int, 1)

func (self *RABSetupItemRelocReqIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABSetupItemRelocReq
   return count//ObjSet
}
func (self *RABSetupItemRelocReqIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 47: //RABSetupItemRelocReq
        return true //self.RABSetupItemRelocReq
   }
   return false//ObjSet
}
func (self *RABSetupItemRelocReqIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 47: //RABSetupItemRelocReq
        self.RABSetupItemRelocReq.Unpack(st)
        self.list = append(self.list, &self.RABSetupItemRelocReq)
   }
}
func (self *RABSetupItemRelocReqIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 47: //RABSetupItemRelocReq
        self.RABSetupItemRelocReq.Pack(st)
      default:
      break
   }
}
func init() {
table_RABSetupItemRelocReqIEs[47] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABSetupItemRelocReq}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RABSetupItemRelocReq{}, PRESENCE:Presence{Presencemandatory}, }
order_RABSetupItemRelocReqIEs[0] = 47
   }

type RABSetupItemRelocReqExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-Alt-RAB-Parameters', 'CRITICALITY': 'ignore', 'EXTENSION': 'Alt-RAB-Parameters', 'PRESENCE': 'optional'}, {'ID': 'id-GERAN-BSC-Container', 'CRITICALITY': 'ignore', 'EXTENSION': 'GERAN-BSC-Container', 'PRESENCE': 'optional'}, None]}
   AltRABParameters  *AltRABParameters
   GERANBSCContainer  *GERANBSCContainer
   list []interface{}
}
func (self *RABSetupItemRelocReqExtIEs)createOT() interface{}{
    return nil
}
var table_RABSetupItemRelocReqExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABSetupItemRelocReqExtIEs = make([]int, 2)

func (self *RABSetupItemRelocReqExtIEs) GetIECount() int{
   count := 0
   if self.AltRABParameters != nil { count += 1 }
   if self.GERANBSCContainer != nil { count += 1 }
   return count//ObjSet
}
func (self *RABSetupItemRelocReqExtIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 89: //AltRABParameters
        if self.AltRABParameters != nil { return true }
      case 107: //GERANBSCContainer
        if self.GERANBSCContainer != nil { return true }
   }
   return false//ObjSet
}
func (self *RABSetupItemRelocReqExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 89: //AltRABParameters
        self.AltRABParameters = &AltRABParameters{}
        self.AltRABParameters.Unpack(st)
        self.list = append(self.list, self.AltRABParameters)
      case 107: //GERANBSCContainer
        self.GERANBSCContainer = &GERANBSCContainer{}
        self.GERANBSCContainer.Unpack(st)
        self.list = append(self.list, self.GERANBSCContainer)
   }
}
func (self *RABSetupItemRelocReqExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 89: //AltRABParameters
        if self.AltRABParameters != nil {self.AltRABParameters.Pack(st)}
      case 107: //GERANBSCContainer
        if self.GERANBSCContainer != nil {self.GERANBSCContainer.Pack(st)}
      default:
      break
   }
}
func init() {
table_RABSetupItemRelocReqExtIEs[89] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idAltRABParameters}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&AltRABParameters{}, PRESENCE:Presence{Presenceoptional}, }
order_RABSetupItemRelocReqExtIEs[0] = 89
table_RABSetupItemRelocReqExtIEs[107] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGERANBSCContainer}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GERANBSCContainer{}, PRESENCE:Presence{Presenceoptional}, }
order_RABSetupItemRelocReqExtIEs[1] = 107
   }

type UserPlaneInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UserPlaneInformationExtIEs)createOT() interface{}{
    return nil
}
var table_UserPlaneInformationExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_UserPlaneInformationExtIEs = make([]int, 0)

type RelocationRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'reject', 'EXTENSION': 'GlobalCN-ID', 'PRESENCE': 'optional'}, {'ID': 'id-SNA-Access-Information', 'CRITICALITY': 'ignore', 'EXTENSION': 'SNA-Access-Information', 'PRESENCE': 'optional'}, {'ID': 'id-UESBI-Iu', 'CRITICALITY': 'ignore', 'EXTENSION': 'UESBI-Iu', 'PRESENCE': 'optional'}, None]}
   GlobalCNID  *GlobalCNID
   SNAAccessInformation  *SNAAccessInformation
   UESBIIu  *UESBIIu
   list []interface{}
}
func (self *RelocationRequestExtensions)createOT() interface{}{
    return nil
}
var table_RelocationRequestExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RelocationRequestExtensions = make([]int, 3)

func (self *RelocationRequestExtensions) GetIECount() int{
   count := 0
   if self.GlobalCNID != nil { count += 1 }
   if self.SNAAccessInformation != nil { count += 1 }
   if self.UESBIIu != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationRequestExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
      case 105: //SNAAccessInformation
        if self.SNAAccessInformation != nil { return true }
      case 118: //UESBIIu
        if self.UESBIIu != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationRequestExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
      case 105: //SNAAccessInformation
        self.SNAAccessInformation = &SNAAccessInformation{}
        self.SNAAccessInformation.Unpack(st)
        self.list = append(self.list, self.SNAAccessInformation)
      case 118: //UESBIIu
        self.UESBIIu = &UESBIIu{}
        self.UESBIIu.Unpack(st)
        self.list = append(self.list, self.UESBIIu)
   }
}
func (self *RelocationRequestExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      case 105: //SNAAccessInformation
        if self.SNAAccessInformation != nil {self.SNAAccessInformation.Pack(st)}
      case 118: //UESBIIu
        if self.UESBIIu != nil {self.UESBIIu.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationRequestExtensions[96] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestExtensions[0] = 96
table_RelocationRequestExtensions[105] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idSNAAccessInformation}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&SNAAccessInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestExtensions[1] = 105
table_RelocationRequestExtensions[118] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idUESBIIu}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&UESBIIu{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestExtensions[2] = 118
   }

type RelocationRequestAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-TargetRNC-ToSourceRNC-TransparentContainer', 'CRITICALITY': 'ignore', 'TYPE': 'TargetRNC-ToSourceRNC-TransparentContainer', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-SetupList-RelocReqAck', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-SetupList-RelocReqAck', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-FailedList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-FailedList', 'PRESENCE': 'optional'}, {'ID': 'id-ChosenIntegrityProtectionAlgorithm', 'CRITICALITY': 'ignore', 'TYPE': 'ChosenIntegrityProtectionAlgorithm', 'PRESENCE': 'optional'}, {'ID': 'id-ChosenEncryptionAlgorithm', 'CRITICALITY': 'ignore', 'TYPE': 'ChosenEncryptionAlgorithm', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   TargetRNCToSourceRNCTransparentContainer  *TargetRNCToSourceRNCTransparentContainer
   RABSetupListRelocReqAck  *RABSetupListRelocReqAck
   RABFailedList  *RABFailedList
   ChosenIntegrityProtectionAlgorithm  *ChosenIntegrityProtectionAlgorithm
   ChosenEncryptionAlgorithm  *ChosenEncryptionAlgorithm
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *RelocationRequestAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_RelocationRequestAcknowledgeIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RelocationRequestAcknowledgeIEs = make([]int, 6)

func (self *RelocationRequestAcknowledgeIEs) GetIECount() int{
   count := 0
   if self.TargetRNCToSourceRNCTransparentContainer != nil { count += 1 }
   if self.RABSetupListRelocReqAck != nil { count += 1 }
   if self.RABFailedList != nil { count += 1 }
   if self.ChosenIntegrityProtectionAlgorithm != nil { count += 1 }
   if self.ChosenEncryptionAlgorithm != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationRequestAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 63: //TargetRNCToSourceRNCTransparentContainer
        if self.TargetRNCToSourceRNCTransparentContainer != nil { return true }
      case 50: //RABSetupListRelocReqAck
        if self.RABSetupListRelocReqAck != nil { return true }
      case 35: //RABFailedList
        if self.RABFailedList != nil { return true }
      case 6: //ChosenIntegrityProtectionAlgorithm
        if self.ChosenIntegrityProtectionAlgorithm != nil { return true }
      case 5: //ChosenEncryptionAlgorithm
        if self.ChosenEncryptionAlgorithm != nil { return true }
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationRequestAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 63: //TargetRNCToSourceRNCTransparentContainer
        self.TargetRNCToSourceRNCTransparentContainer = &TargetRNCToSourceRNCTransparentContainer{}
        self.TargetRNCToSourceRNCTransparentContainer.Unpack(st)
        self.list = append(self.list, self.TargetRNCToSourceRNCTransparentContainer)
      case 50: //RABSetupListRelocReqAck
        self.RABSetupListRelocReqAck = &RABSetupListRelocReqAck{}
        self.RABSetupListRelocReqAck.Unpack(st)
        self.list = append(self.list, self.RABSetupListRelocReqAck)
      case 35: //RABFailedList
        self.RABFailedList = &RABFailedList{}
        self.RABFailedList.Unpack(st)
        self.list = append(self.list, self.RABFailedList)
      case 6: //ChosenIntegrityProtectionAlgorithm
        self.ChosenIntegrityProtectionAlgorithm = &ChosenIntegrityProtectionAlgorithm{}
        self.ChosenIntegrityProtectionAlgorithm.Unpack(st)
        self.list = append(self.list, self.ChosenIntegrityProtectionAlgorithm)
      case 5: //ChosenEncryptionAlgorithm
        self.ChosenEncryptionAlgorithm = &ChosenEncryptionAlgorithm{}
        self.ChosenEncryptionAlgorithm.Unpack(st)
        self.list = append(self.list, self.ChosenEncryptionAlgorithm)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *RelocationRequestAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 63: //TargetRNCToSourceRNCTransparentContainer
        if self.TargetRNCToSourceRNCTransparentContainer != nil {self.TargetRNCToSourceRNCTransparentContainer.Pack(st)}
      case 50: //RABSetupListRelocReqAck
        if self.RABSetupListRelocReqAck != nil {self.RABSetupListRelocReqAck.Pack(st)}
      case 35: //RABFailedList
        if self.RABFailedList != nil {self.RABFailedList.Pack(st)}
      case 6: //ChosenIntegrityProtectionAlgorithm
        if self.ChosenIntegrityProtectionAlgorithm != nil {self.ChosenIntegrityProtectionAlgorithm.Pack(st)}
      case 5: //ChosenEncryptionAlgorithm
        if self.ChosenEncryptionAlgorithm != nil {self.ChosenEncryptionAlgorithm.Pack(st)}
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationRequestAcknowledgeIEs[63] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idTargetRNCToSourceRNCTransparentContainer}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TargetRNCToSourceRNCTransparentContainer{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestAcknowledgeIEs[0] = 63
table_RelocationRequestAcknowledgeIEs[50] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABSetupListRelocReqAck}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABSetupListRelocReqAck{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestAcknowledgeIEs[1] = 50
table_RelocationRequestAcknowledgeIEs[35] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABFailedList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABFailedList{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestAcknowledgeIEs[2] = 35
table_RelocationRequestAcknowledgeIEs[6] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idChosenIntegrityProtectionAlgorithm}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ChosenIntegrityProtectionAlgorithm{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestAcknowledgeIEs[3] = 6
table_RelocationRequestAcknowledgeIEs[5] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idChosenEncryptionAlgorithm}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ChosenEncryptionAlgorithm{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestAcknowledgeIEs[4] = 5
table_RelocationRequestAcknowledgeIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestAcknowledgeIEs[5] = 9
   }

type RABSetupItemRelocReqAckIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-SetupItem-RelocReqAck', 'CRITICALITY': 'reject', 'TYPE': 'RAB-SetupItem-RelocReqAck', 'PRESENCE': 'mandatory'}, None]}
   RABSetupItemRelocReqAck  RABSetupItemRelocReqAck
   list []interface{}
}
func (self *RABSetupItemRelocReqAckIEs)createOT() interface{}{
    return nil
}
var table_RABSetupItemRelocReqAckIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABSetupItemRelocReqAckIEs = make([]int, 1)

func (self *RABSetupItemRelocReqAckIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABSetupItemRelocReqAck
   return count//ObjSet
}
func (self *RABSetupItemRelocReqAckIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 48: //RABSetupItemRelocReqAck
        return true //self.RABSetupItemRelocReqAck
   }
   return false//ObjSet
}
func (self *RABSetupItemRelocReqAckIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 48: //RABSetupItemRelocReqAck
        self.RABSetupItemRelocReqAck.Unpack(st)
        self.list = append(self.list, &self.RABSetupItemRelocReqAck)
   }
}
func (self *RABSetupItemRelocReqAckIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 48: //RABSetupItemRelocReqAck
        self.RABSetupItemRelocReqAck.Pack(st)
      default:
      break
   }
}
func init() {
table_RABSetupItemRelocReqAckIEs[48] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABSetupItemRelocReqAck}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RABSetupItemRelocReqAck{}, PRESENCE:Presence{Presencemandatory}, }
order_RABSetupItemRelocReqAckIEs[0] = 48
   }

type RABSetupItemRelocReqAckExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-Ass-RAB-Parameters', 'CRITICALITY': 'ignore', 'EXTENSION': 'Ass-RAB-Parameters', 'PRESENCE': 'optional'}, {'ID': 'id-TransportLayerAddress', 'CRITICALITY': 'ignore', 'EXTENSION': 'TransportLayerAddress', 'PRESENCE': 'optional'}, {'ID': 'id-IuTransportAssociation', 'CRITICALITY': 'ignore', 'EXTENSION': 'IuTransportAssociation', 'PRESENCE': 'optional'}, None]}
   AssRABParameters  *AssRABParameters
   TransportLayerAddress  *TransportLayerAddress
   IuTransportAssociation  *IuTransportAssociation
   list []interface{}
}
func (self *RABSetupItemRelocReqAckExtIEs)createOT() interface{}{
    return nil
}
var table_RABSetupItemRelocReqAckExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABSetupItemRelocReqAckExtIEs = make([]int, 3)

func (self *RABSetupItemRelocReqAckExtIEs) GetIECount() int{
   count := 0
   if self.AssRABParameters != nil { count += 1 }
   if self.TransportLayerAddress != nil { count += 1 }
   if self.IuTransportAssociation != nil { count += 1 }
   return count//ObjSet
}
func (self *RABSetupItemRelocReqAckExtIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 90: //AssRABParameters
        if self.AssRABParameters != nil { return true }
      case 67: //TransportLayerAddress
        if self.TransportLayerAddress != nil { return true }
      case 13: //IuTransportAssociation
        if self.IuTransportAssociation != nil { return true }
   }
   return false//ObjSet
}
func (self *RABSetupItemRelocReqAckExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 90: //AssRABParameters
        self.AssRABParameters = &AssRABParameters{}
        self.AssRABParameters.Unpack(st)
        self.list = append(self.list, self.AssRABParameters)
      case 67: //TransportLayerAddress
        self.TransportLayerAddress = &TransportLayerAddress{}
        self.TransportLayerAddress.Unpack(st)
        self.list = append(self.list, self.TransportLayerAddress)
      case 13: //IuTransportAssociation
        self.IuTransportAssociation = &IuTransportAssociation{}
        self.IuTransportAssociation.Unpack(st)
        self.list = append(self.list, self.IuTransportAssociation)
   }
}
func (self *RABSetupItemRelocReqAckExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 90: //AssRABParameters
        if self.AssRABParameters != nil {self.AssRABParameters.Pack(st)}
      case 67: //TransportLayerAddress
        if self.TransportLayerAddress != nil {self.TransportLayerAddress.Pack(st)}
      case 13: //IuTransportAssociation
        if self.IuTransportAssociation != nil {self.IuTransportAssociation.Pack(st)}
      default:
      break
   }
}
func init() {
table_RABSetupItemRelocReqAckExtIEs[90] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idAssRABParameters}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&AssRABParameters{}, PRESENCE:Presence{Presenceoptional}, }
order_RABSetupItemRelocReqAckExtIEs[0] = 90
table_RABSetupItemRelocReqAckExtIEs[67] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idTransportLayerAddress}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&TransportLayerAddress{}, PRESENCE:Presence{Presenceoptional}, }
order_RABSetupItemRelocReqAckExtIEs[1] = 67
table_RABSetupItemRelocReqAckExtIEs[13] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idIuTransportAssociation}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&IuTransportAssociation{}, PRESENCE:Presence{Presenceoptional}, }
order_RABSetupItemRelocReqAckExtIEs[2] = 13
   }

type RABFailedItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-FailedItem', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-FailedItem', 'PRESENCE': 'mandatory'}, None]}
   RABFailedItem  RABFailedItem
   list []interface{}
}
func (self *RABFailedItemIEs)createOT() interface{}{
    return nil
}
var table_RABFailedItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABFailedItemIEs = make([]int, 1)

func (self *RABFailedItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABFailedItem
   return count//ObjSet
}
func (self *RABFailedItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 34: //RABFailedItem
        return true //self.RABFailedItem
   }
   return false//ObjSet
}
func (self *RABFailedItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 34: //RABFailedItem
        self.RABFailedItem.Unpack(st)
        self.list = append(self.list, &self.RABFailedItem)
   }
}
func (self *RABFailedItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 34: //RABFailedItem
        self.RABFailedItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABFailedItemIEs[34] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABFailedItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABFailedItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABFailedItemIEs[0] = 34
   }

type RABFailedItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABFailedItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABFailedItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABFailedItemExtIEs = make([]int, 0)

type RelocationRequestAcknowledgeExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-NewBSS-To-OldBSS-Information', 'CRITICALITY': 'ignore', 'EXTENSION': 'NewBSS-To-OldBSS-Information', 'PRESENCE': 'optional'}, None]}
   NewBSSToOldBSSInformation  *NewBSSToOldBSSInformation
   list []interface{}
}
func (self *RelocationRequestAcknowledgeExtensions)createOT() interface{}{
    return nil
}
var table_RelocationRequestAcknowledgeExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RelocationRequestAcknowledgeExtensions = make([]int, 1)

func (self *RelocationRequestAcknowledgeExtensions) GetIECount() int{
   count := 0
   if self.NewBSSToOldBSSInformation != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationRequestAcknowledgeExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 100: //NewBSSToOldBSSInformation
        if self.NewBSSToOldBSSInformation != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationRequestAcknowledgeExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 100: //NewBSSToOldBSSInformation
        self.NewBSSToOldBSSInformation = &NewBSSToOldBSSInformation{}
        self.NewBSSToOldBSSInformation.Unpack(st)
        self.list = append(self.list, self.NewBSSToOldBSSInformation)
   }
}
func (self *RelocationRequestAcknowledgeExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 100: //NewBSSToOldBSSInformation
        if self.NewBSSToOldBSSInformation != nil {self.NewBSSToOldBSSInformation.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationRequestAcknowledgeExtensions[100] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idNewBSSToOldBSSInformation}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&NewBSSToOldBSSInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationRequestAcknowledgeExtensions[0] = 100
   }

type RelocationFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *RelocationFailureIEs)createOT() interface{}{
    return nil
}
var table_RelocationFailureIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RelocationFailureIEs = make([]int, 2)

func (self *RelocationFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationFailureIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        return true //self.Cause
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *RelocationFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Pack(st)
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationFailureIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationFailureIEs[0] = 4
table_RelocationFailureIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationFailureIEs[1] = 9
   }

type RelocationFailureExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-NewBSS-To-OldBSS-Information', 'CRITICALITY': 'ignore', 'EXTENSION': 'NewBSS-To-OldBSS-Information', 'PRESENCE': 'optional'}, {'ID': 'id-GERAN-Classmark', 'CRITICALITY': 'ignore', 'EXTENSION': 'GERAN-Classmark', 'PRESENCE': 'optional'}, None]}
   NewBSSToOldBSSInformation  *NewBSSToOldBSSInformation
   GERANClassmark  *GERANClassmark
   list []interface{}
}
func (self *RelocationFailureExtensions)createOT() interface{}{
    return nil
}
var table_RelocationFailureExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RelocationFailureExtensions = make([]int, 2)

func (self *RelocationFailureExtensions) GetIECount() int{
   count := 0
   if self.NewBSSToOldBSSInformation != nil { count += 1 }
   if self.GERANClassmark != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationFailureExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 100: //NewBSSToOldBSSInformation
        if self.NewBSSToOldBSSInformation != nil { return true }
      case 108: //GERANClassmark
        if self.GERANClassmark != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationFailureExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 100: //NewBSSToOldBSSInformation
        self.NewBSSToOldBSSInformation = &NewBSSToOldBSSInformation{}
        self.NewBSSToOldBSSInformation.Unpack(st)
        self.list = append(self.list, self.NewBSSToOldBSSInformation)
      case 108: //GERANClassmark
        self.GERANClassmark = &GERANClassmark{}
        self.GERANClassmark.Unpack(st)
        self.list = append(self.list, self.GERANClassmark)
   }
}
func (self *RelocationFailureExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 100: //NewBSSToOldBSSInformation
        if self.NewBSSToOldBSSInformation != nil {self.NewBSSToOldBSSInformation.Pack(st)}
      case 108: //GERANClassmark
        if self.GERANClassmark != nil {self.GERANClassmark.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationFailureExtensions[100] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idNewBSSToOldBSSInformation}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&NewBSSToOldBSSInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationFailureExtensions[0] = 100
table_RelocationFailureExtensions[108] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGERANClassmark}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GERANClassmark{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationFailureExtensions[1] = 108
   }

type RelocationCancelIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, None]}
   Cause  Cause
   list []interface{}
}
func (self *RelocationCancelIEs)createOT() interface{}{
    return nil
}
var table_RelocationCancelIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RelocationCancelIEs = make([]int, 1)

func (self *RelocationCancelIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   return count//ObjSet
}
func (self *RelocationCancelIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        return true //self.Cause
   }
   return false//ObjSet
}
func (self *RelocationCancelIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
   }
}
func (self *RelocationCancelIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Pack(st)
      default:
      break
   }
}
func init() {
table_RelocationCancelIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_RelocationCancelIEs[0] = 4
   }

type RelocationCancelExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RelocationCancelExtensions)createOT() interface{}{
    return nil
}
var table_RelocationCancelExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RelocationCancelExtensions = make([]int, 0)

type RelocationCancelAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *RelocationCancelAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_RelocationCancelAcknowledgeIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RelocationCancelAcknowledgeIEs = make([]int, 1)

func (self *RelocationCancelAcknowledgeIEs) GetIECount() int{
   count := 0
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *RelocationCancelAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *RelocationCancelAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *RelocationCancelAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_RelocationCancelAcknowledgeIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_RelocationCancelAcknowledgeIEs[0] = 9
   }

type RelocationCancelAcknowledgeExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RelocationCancelAcknowledgeExtensions)createOT() interface{}{
    return nil
}
var table_RelocationCancelAcknowledgeExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RelocationCancelAcknowledgeExtensions = make([]int, 0)

type SRNSContextRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-DataForwardingList-SRNS-CtxReq', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-DataForwardingList-SRNS-CtxReq', 'PRESENCE': 'mandatory'}, None]}
   RABDataForwardingListSRNSCtxReq  RABDataForwardingListSRNSCtxReq
   list []interface{}
}
func (self *SRNSContextRequestIEs)createOT() interface{}{
    return nil
}
var table_SRNSContextRequestIEs = make(map[int]*RANAPPROTOCOLIES)

var order_SRNSContextRequestIEs = make([]int, 1)

func (self *SRNSContextRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABDataForwardingListSRNSCtxReq
   return count//ObjSet
}
func (self *SRNSContextRequestIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 29: //RABDataForwardingListSRNSCtxReq
        return true //self.RABDataForwardingListSRNSCtxReq
   }
   return false//ObjSet
}
func (self *SRNSContextRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 29: //RABDataForwardingListSRNSCtxReq
        self.RABDataForwardingListSRNSCtxReq.Unpack(st)
        self.list = append(self.list, &self.RABDataForwardingListSRNSCtxReq)
   }
}
func (self *SRNSContextRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 29: //RABDataForwardingListSRNSCtxReq
        self.RABDataForwardingListSRNSCtxReq.Pack(st)
      default:
      break
   }
}
func init() {
table_SRNSContextRequestIEs[29] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABDataForwardingListSRNSCtxReq}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABDataForwardingListSRNSCtxReq{}, PRESENCE:Presence{Presencemandatory}, }
order_SRNSContextRequestIEs[0] = 29
   }

type RABDataForwardingItemSRNSCtxReqIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-DataForwardingItem-SRNS-CtxReq', 'CRITICALITY': 'reject', 'TYPE': 'RAB-DataForwardingItem-SRNS-CtxReq', 'PRESENCE': 'mandatory'}, None]}
   RABDataForwardingItemSRNSCtxReq  RABDataForwardingItemSRNSCtxReq
   list []interface{}
}
func (self *RABDataForwardingItemSRNSCtxReqIEs)createOT() interface{}{
    return nil
}
var table_RABDataForwardingItemSRNSCtxReqIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABDataForwardingItemSRNSCtxReqIEs = make([]int, 1)

func (self *RABDataForwardingItemSRNSCtxReqIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABDataForwardingItemSRNSCtxReq
   return count//ObjSet
}
func (self *RABDataForwardingItemSRNSCtxReqIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 27: //RABDataForwardingItemSRNSCtxReq
        return true //self.RABDataForwardingItemSRNSCtxReq
   }
   return false//ObjSet
}
func (self *RABDataForwardingItemSRNSCtxReqIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 27: //RABDataForwardingItemSRNSCtxReq
        self.RABDataForwardingItemSRNSCtxReq.Unpack(st)
        self.list = append(self.list, &self.RABDataForwardingItemSRNSCtxReq)
   }
}
func (self *RABDataForwardingItemSRNSCtxReqIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 27: //RABDataForwardingItemSRNSCtxReq
        self.RABDataForwardingItemSRNSCtxReq.Pack(st)
      default:
      break
   }
}
func init() {
table_RABDataForwardingItemSRNSCtxReqIEs[27] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABDataForwardingItemSRNSCtxReq}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RABDataForwardingItemSRNSCtxReq{}, PRESENCE:Presence{Presencemandatory}, }
order_RABDataForwardingItemSRNSCtxReqIEs[0] = 27
   }

type RABDataForwardingItemSRNSCtxReqExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABDataForwardingItemSRNSCtxReqExtIEs)createOT() interface{}{
    return nil
}
var table_RABDataForwardingItemSRNSCtxReqExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABDataForwardingItemSRNSCtxReqExtIEs = make([]int, 0)

type SRNSContextRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SRNSContextRequestExtensions)createOT() interface{}{
    return nil
}
var table_SRNSContextRequestExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SRNSContextRequestExtensions = make([]int, 0)

type SRNSContextResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ContextList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ContextList', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-ContextFailedtoTransferList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ContextFailedtoTransferList', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   RABContextList  *RABContextList
   RABContextFailedtoTransferList  *RABContextFailedtoTransferList
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *SRNSContextResponseIEs)createOT() interface{}{
    return nil
}
var table_SRNSContextResponseIEs = make(map[int]*RANAPPROTOCOLIES)

var order_SRNSContextResponseIEs = make([]int, 3)

func (self *SRNSContextResponseIEs) GetIECount() int{
   count := 0
   if self.RABContextList != nil { count += 1 }
   if self.RABContextFailedtoTransferList != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *SRNSContextResponseIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 25: //RABContextList
        if self.RABContextList != nil { return true }
      case 85: //RABContextFailedtoTransferList
        if self.RABContextFailedtoTransferList != nil { return true }
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *SRNSContextResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 25: //RABContextList
        self.RABContextList = &RABContextList{}
        self.RABContextList.Unpack(st)
        self.list = append(self.list, self.RABContextList)
      case 85: //RABContextFailedtoTransferList
        self.RABContextFailedtoTransferList = &RABContextFailedtoTransferList{}
        self.RABContextFailedtoTransferList.Unpack(st)
        self.list = append(self.list, self.RABContextFailedtoTransferList)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *SRNSContextResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 25: //RABContextList
        if self.RABContextList != nil {self.RABContextList.Pack(st)}
      case 85: //RABContextFailedtoTransferList
        if self.RABContextFailedtoTransferList != nil {self.RABContextFailedtoTransferList.Pack(st)}
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_SRNSContextResponseIEs[25] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABContextList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABContextList{}, PRESENCE:Presence{Presenceoptional}, }
order_SRNSContextResponseIEs[0] = 25
table_SRNSContextResponseIEs[85] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABContextFailedtoTransferList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABContextFailedtoTransferList{}, PRESENCE:Presence{Presenceoptional}, }
order_SRNSContextResponseIEs[1] = 85
table_SRNSContextResponseIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_SRNSContextResponseIEs[2] = 9
   }

type RABContextItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ContextItem', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ContextItem', 'PRESENCE': 'mandatory'}, None]}
   RABContextItem  RABContextItem
   list []interface{}
}
func (self *RABContextItemIEs)createOT() interface{}{
    return nil
}
var table_RABContextItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABContextItemIEs = make([]int, 1)

func (self *RABContextItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABContextItem
   return count//ObjSet
}
func (self *RABContextItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 24: //RABContextItem
        return true //self.RABContextItem
   }
   return false//ObjSet
}
func (self *RABContextItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 24: //RABContextItem
        self.RABContextItem.Unpack(st)
        self.list = append(self.list, &self.RABContextItem)
   }
}
func (self *RABContextItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 24: //RABContextItem
        self.RABContextItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABContextItemIEs[24] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABContextItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABContextItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABContextItemIEs[0] = 24
   }

type RABContextItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABContextItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABContextItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABContextItemExtIEs = make([]int, 0)

type RABsContextFailedtoTransferItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ContextFailedtoTransferItem', 'CRITICALITY': 'ignore', 'TYPE': 'RABs-ContextFailedtoTransferItem', 'PRESENCE': 'mandatory'}, None]}
   RABContextFailedtoTransferItem  RABsContextFailedtoTransferItem
   list []interface{}
}
func (self *RABsContextFailedtoTransferItemIEs)createOT() interface{}{
    return nil
}
var table_RABsContextFailedtoTransferItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABsContextFailedtoTransferItemIEs = make([]int, 1)

func (self *RABsContextFailedtoTransferItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABContextFailedtoTransferItem
   return count//ObjSet
}
func (self *RABsContextFailedtoTransferItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 84: //RABContextFailedtoTransferItem
        return true //self.RABContextFailedtoTransferItem
   }
   return false//ObjSet
}
func (self *RABsContextFailedtoTransferItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 84: //RABContextFailedtoTransferItem
        self.RABContextFailedtoTransferItem.Unpack(st)
        self.list = append(self.list, &self.RABContextFailedtoTransferItem)
   }
}
func (self *RABsContextFailedtoTransferItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 84: //RABContextFailedtoTransferItem
        self.RABContextFailedtoTransferItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABsContextFailedtoTransferItemIEs[84] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABContextFailedtoTransferItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABsContextFailedtoTransferItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABsContextFailedtoTransferItemIEs[0] = 84
   }

type RABsContextFailedtoTransferItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABsContextFailedtoTransferItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABsContextFailedtoTransferItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABsContextFailedtoTransferItemExtIEs = make([]int, 0)

type SRNSContextResponseExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SRNSContextResponseExtensions)createOT() interface{}{
    return nil
}
var table_SRNSContextResponseExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SRNSContextResponseExtensions = make([]int, 0)

type SecurityModeCommandIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-IntegrityProtectionInformation', 'CRITICALITY': 'reject', 'TYPE': 'IntegrityProtectionInformation', 'PRESENCE': 'mandatory'}, {'ID': 'id-EncryptionInformation', 'CRITICALITY': 'ignore', 'TYPE': 'EncryptionInformation', 'PRESENCE': 'optional'}, {'ID': 'id-KeyStatus', 'CRITICALITY': 'reject', 'TYPE': 'KeyStatus', 'PRESENCE': 'mandatory'}, None]}
   IntegrityProtectionInformation  IntegrityProtectionInformation
   EncryptionInformation  *EncryptionInformation
   KeyStatus  KeyStatus
   list []interface{}
}
func (self *SecurityModeCommandIEs)createOT() interface{}{
    return nil
}
var table_SecurityModeCommandIEs = make(map[int]*RANAPPROTOCOLIES)

var order_SecurityModeCommandIEs = make([]int, 3)

func (self *SecurityModeCommandIEs) GetIECount() int{
   count := 0
   count +=1 //self.IntegrityProtectionInformation
   if self.EncryptionInformation != nil { count += 1 }
   count +=1 //self.KeyStatus
   return count//ObjSet
}
func (self *SecurityModeCommandIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 12: //IntegrityProtectionInformation
        return true //self.IntegrityProtectionInformation
      case 11: //EncryptionInformation
        if self.EncryptionInformation != nil { return true }
      case 75: //KeyStatus
        return true //self.KeyStatus
   }
   return false//ObjSet
}
func (self *SecurityModeCommandIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 12: //IntegrityProtectionInformation
        self.IntegrityProtectionInformation.Unpack(st)
        self.list = append(self.list, &self.IntegrityProtectionInformation)
      case 11: //EncryptionInformation
        self.EncryptionInformation = &EncryptionInformation{}
        self.EncryptionInformation.Unpack(st)
        self.list = append(self.list, self.EncryptionInformation)
      case 75: //KeyStatus
        self.KeyStatus.Unpack(st)
        self.list = append(self.list, &self.KeyStatus)
   }
}
func (self *SecurityModeCommandIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 12: //IntegrityProtectionInformation
        self.IntegrityProtectionInformation.Pack(st)
      case 11: //EncryptionInformation
        if self.EncryptionInformation != nil {self.EncryptionInformation.Pack(st)}
      case 75: //KeyStatus
        self.KeyStatus.Pack(st)
      default:
      break
   }
}
func init() {
table_SecurityModeCommandIEs[12] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idIntegrityProtectionInformation}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&IntegrityProtectionInformation{}, PRESENCE:Presence{Presencemandatory}, }
order_SecurityModeCommandIEs[0] = 12
table_SecurityModeCommandIEs[11] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idEncryptionInformation}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&EncryptionInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_SecurityModeCommandIEs[1] = 11
table_SecurityModeCommandIEs[75] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idKeyStatus}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&KeyStatus{}, PRESENCE:Presence{Presencemandatory}, }
order_SecurityModeCommandIEs[2] = 75
   }

type SecurityModeCommandExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SecurityModeCommandExtensions)createOT() interface{}{
    return nil
}
var table_SecurityModeCommandExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SecurityModeCommandExtensions = make([]int, 0)

type SecurityModeCompleteIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-ChosenIntegrityProtectionAlgorithm', 'CRITICALITY': 'reject', 'TYPE': 'ChosenIntegrityProtectionAlgorithm', 'PRESENCE': 'mandatory'}, {'ID': 'id-ChosenEncryptionAlgorithm', 'CRITICALITY': 'ignore', 'TYPE': 'ChosenEncryptionAlgorithm', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   ChosenIntegrityProtectionAlgorithm  ChosenIntegrityProtectionAlgorithm
   ChosenEncryptionAlgorithm  *ChosenEncryptionAlgorithm
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *SecurityModeCompleteIEs)createOT() interface{}{
    return nil
}
var table_SecurityModeCompleteIEs = make(map[int]*RANAPPROTOCOLIES)

var order_SecurityModeCompleteIEs = make([]int, 3)

func (self *SecurityModeCompleteIEs) GetIECount() int{
   count := 0
   count +=1 //self.ChosenIntegrityProtectionAlgorithm
   if self.ChosenEncryptionAlgorithm != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *SecurityModeCompleteIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 6: //ChosenIntegrityProtectionAlgorithm
        return true //self.ChosenIntegrityProtectionAlgorithm
      case 5: //ChosenEncryptionAlgorithm
        if self.ChosenEncryptionAlgorithm != nil { return true }
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *SecurityModeCompleteIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 6: //ChosenIntegrityProtectionAlgorithm
        self.ChosenIntegrityProtectionAlgorithm.Unpack(st)
        self.list = append(self.list, &self.ChosenIntegrityProtectionAlgorithm)
      case 5: //ChosenEncryptionAlgorithm
        self.ChosenEncryptionAlgorithm = &ChosenEncryptionAlgorithm{}
        self.ChosenEncryptionAlgorithm.Unpack(st)
        self.list = append(self.list, self.ChosenEncryptionAlgorithm)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *SecurityModeCompleteIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 6: //ChosenIntegrityProtectionAlgorithm
        self.ChosenIntegrityProtectionAlgorithm.Pack(st)
      case 5: //ChosenEncryptionAlgorithm
        if self.ChosenEncryptionAlgorithm != nil {self.ChosenEncryptionAlgorithm.Pack(st)}
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_SecurityModeCompleteIEs[6] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idChosenIntegrityProtectionAlgorithm}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ChosenIntegrityProtectionAlgorithm{}, PRESENCE:Presence{Presencemandatory}, }
order_SecurityModeCompleteIEs[0] = 6
table_SecurityModeCompleteIEs[5] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idChosenEncryptionAlgorithm}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ChosenEncryptionAlgorithm{}, PRESENCE:Presence{Presenceoptional}, }
order_SecurityModeCompleteIEs[1] = 5
table_SecurityModeCompleteIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_SecurityModeCompleteIEs[2] = 9
   }

type SecurityModeCompleteExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SecurityModeCompleteExtensions)createOT() interface{}{
    return nil
}
var table_SecurityModeCompleteExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SecurityModeCompleteExtensions = make([]int, 0)

type SecurityModeRejectIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *SecurityModeRejectIEs)createOT() interface{}{
    return nil
}
var table_SecurityModeRejectIEs = make(map[int]*RANAPPROTOCOLIES)

var order_SecurityModeRejectIEs = make([]int, 2)

func (self *SecurityModeRejectIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *SecurityModeRejectIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        return true //self.Cause
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *SecurityModeRejectIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *SecurityModeRejectIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Pack(st)
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_SecurityModeRejectIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_SecurityModeRejectIEs[0] = 4
table_SecurityModeRejectIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_SecurityModeRejectIEs[1] = 9
   }

type SecurityModeRejectExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SecurityModeRejectExtensions)createOT() interface{}{
    return nil
}
var table_SecurityModeRejectExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SecurityModeRejectExtensions = make([]int, 0)

type DataVolumeReportRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-DataVolumeReportRequestList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-DataVolumeReportRequestList', 'PRESENCE': 'mandatory'}, None]}
   RABDataVolumeReportRequestList  RABDataVolumeReportRequestList
   list []interface{}
}
func (self *DataVolumeReportRequestIEs)createOT() interface{}{
    return nil
}
var table_DataVolumeReportRequestIEs = make(map[int]*RANAPPROTOCOLIES)

var order_DataVolumeReportRequestIEs = make([]int, 1)

func (self *DataVolumeReportRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABDataVolumeReportRequestList
   return count//ObjSet
}
func (self *DataVolumeReportRequestIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 33: //RABDataVolumeReportRequestList
        return true //self.RABDataVolumeReportRequestList
   }
   return false//ObjSet
}
func (self *DataVolumeReportRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 33: //RABDataVolumeReportRequestList
        self.RABDataVolumeReportRequestList.Unpack(st)
        self.list = append(self.list, &self.RABDataVolumeReportRequestList)
   }
}
func (self *DataVolumeReportRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 33: //RABDataVolumeReportRequestList
        self.RABDataVolumeReportRequestList.Pack(st)
      default:
      break
   }
}
func init() {
table_DataVolumeReportRequestIEs[33] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABDataVolumeReportRequestList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABDataVolumeReportRequestList{}, PRESENCE:Presence{Presencemandatory}, }
order_DataVolumeReportRequestIEs[0] = 33
   }

type RABDataVolumeReportRequestItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-DataVolumeReportRequestItem', 'CRITICALITY': 'reject', 'TYPE': 'RAB-DataVolumeReportRequestItem', 'PRESENCE': 'mandatory'}, None]}
   RABDataVolumeReportRequestItem  RABDataVolumeReportRequestItem
   list []interface{}
}
func (self *RABDataVolumeReportRequestItemIEs)createOT() interface{}{
    return nil
}
var table_RABDataVolumeReportRequestItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABDataVolumeReportRequestItemIEs = make([]int, 1)

func (self *RABDataVolumeReportRequestItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABDataVolumeReportRequestItem
   return count//ObjSet
}
func (self *RABDataVolumeReportRequestItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 32: //RABDataVolumeReportRequestItem
        return true //self.RABDataVolumeReportRequestItem
   }
   return false//ObjSet
}
func (self *RABDataVolumeReportRequestItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 32: //RABDataVolumeReportRequestItem
        self.RABDataVolumeReportRequestItem.Unpack(st)
        self.list = append(self.list, &self.RABDataVolumeReportRequestItem)
   }
}
func (self *RABDataVolumeReportRequestItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 32: //RABDataVolumeReportRequestItem
        self.RABDataVolumeReportRequestItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABDataVolumeReportRequestItemIEs[32] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABDataVolumeReportRequestItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&RABDataVolumeReportRequestItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABDataVolumeReportRequestItemIEs[0] = 32
   }

type RABDataVolumeReportRequestItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABDataVolumeReportRequestItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABDataVolumeReportRequestItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABDataVolumeReportRequestItemExtIEs = make([]int, 0)

type DataVolumeReportRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DataVolumeReportRequestExtensions)createOT() interface{}{
    return nil
}
var table_DataVolumeReportRequestExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_DataVolumeReportRequestExtensions = make([]int, 0)

type DataVolumeReportIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-DataVolumeReportList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-DataVolumeReportList', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-FailedtoReportList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-FailedtoReportList', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   RABDataVolumeReportList  *RABDataVolumeReportList
   RABFailedtoReportList  *RABFailedtoReportList
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *DataVolumeReportIEs)createOT() interface{}{
    return nil
}
var table_DataVolumeReportIEs = make(map[int]*RANAPPROTOCOLIES)

var order_DataVolumeReportIEs = make([]int, 3)

func (self *DataVolumeReportIEs) GetIECount() int{
   count := 0
   if self.RABDataVolumeReportList != nil { count += 1 }
   if self.RABFailedtoReportList != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *DataVolumeReportIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 31: //RABDataVolumeReportList
        if self.RABDataVolumeReportList != nil { return true }
      case 72: //RABFailedtoReportList
        if self.RABFailedtoReportList != nil { return true }
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *DataVolumeReportIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 31: //RABDataVolumeReportList
        self.RABDataVolumeReportList = &RABDataVolumeReportList{}
        self.RABDataVolumeReportList.Unpack(st)
        self.list = append(self.list, self.RABDataVolumeReportList)
      case 72: //RABFailedtoReportList
        self.RABFailedtoReportList = &RABFailedtoReportList{}
        self.RABFailedtoReportList.Unpack(st)
        self.list = append(self.list, self.RABFailedtoReportList)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *DataVolumeReportIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 31: //RABDataVolumeReportList
        if self.RABDataVolumeReportList != nil {self.RABDataVolumeReportList.Pack(st)}
      case 72: //RABFailedtoReportList
        if self.RABFailedtoReportList != nil {self.RABFailedtoReportList.Pack(st)}
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_DataVolumeReportIEs[31] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABDataVolumeReportList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABDataVolumeReportList{}, PRESENCE:Presence{Presenceoptional}, }
order_DataVolumeReportIEs[0] = 31
table_DataVolumeReportIEs[72] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABFailedtoReportList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABFailedtoReportList{}, PRESENCE:Presence{Presenceoptional}, }
order_DataVolumeReportIEs[1] = 72
table_DataVolumeReportIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_DataVolumeReportIEs[2] = 9
   }

type DataVolumeReportExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DataVolumeReportExtensions)createOT() interface{}{
    return nil
}
var table_DataVolumeReportExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_DataVolumeReportExtensions = make([]int, 0)

type RABsfailedtoreportItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-FailedtoReportItem', 'CRITICALITY': 'ignore', 'TYPE': 'RABs-failed-to-reportItem', 'PRESENCE': 'mandatory'}, None]}
   RABFailedtoReportItem  RABsfailedtoreportItem
   list []interface{}
}
func (self *RABsfailedtoreportItemIEs)createOT() interface{}{
    return nil
}
var table_RABsfailedtoreportItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABsfailedtoreportItemIEs = make([]int, 1)

func (self *RABsfailedtoreportItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABFailedtoReportItem
   return count//ObjSet
}
func (self *RABsfailedtoreportItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 71: //RABFailedtoReportItem
        return true //self.RABFailedtoReportItem
   }
   return false//ObjSet
}
func (self *RABsfailedtoreportItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 71: //RABFailedtoReportItem
        self.RABFailedtoReportItem.Unpack(st)
        self.list = append(self.list, &self.RABFailedtoReportItem)
   }
}
func (self *RABsfailedtoreportItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 71: //RABFailedtoReportItem
        self.RABFailedtoReportItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABsfailedtoreportItemIEs[71] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABFailedtoReportItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABsfailedtoreportItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABsfailedtoreportItemIEs[0] = 71
   }

type RABsfailedtoreportItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABsfailedtoreportItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABsfailedtoreportItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABsfailedtoreportItemExtIEs = make([]int, 0)

type ResetIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'reject', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'optional'}, None]}
   Cause  Cause
   CNDomainIndicator  CNDomainIndicator
   GlobalRNCID  *GlobalRNCID
   list []interface{}
}
func (self *ResetIEs)createOT() interface{}{
    return nil
}
var table_ResetIEs = make(map[int]*RANAPPROTOCOLIES)

var order_ResetIEs = make([]int, 3)

func (self *ResetIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   count +=1 //self.CNDomainIndicator
   if self.GlobalRNCID != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        return true //self.Cause
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 86: //GlobalRNCID
        self.GlobalRNCID = &GlobalRNCID{}
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, self.GlobalRNCID)
   }
}
func (self *ResetIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Pack(st)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil {self.GlobalRNCID.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetIEs[0] = 4
table_ResetIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetIEs[1] = 3
table_ResetIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetIEs[2] = 86
   }

type ResetExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'EXTENSION': 'GlobalCN-ID', 'PRESENCE': 'optional'}, None]}
   GlobalCNID  *GlobalCNID
   list []interface{}
}
func (self *ResetExtensions)createOT() interface{}{
    return nil
}
var table_ResetExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_ResetExtensions = make([]int, 1)

func (self *ResetExtensions) GetIECount() int{
   count := 0
   if self.GlobalCNID != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
   }
}
func (self *ResetExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetExtensions[96] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetExtensions[0] = 96
   }

type ResetAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'reject', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'optional'}, None]}
   CNDomainIndicator  CNDomainIndicator
   CriticalityDiagnostics  *CriticalityDiagnostics
   GlobalRNCID  *GlobalRNCID
   list []interface{}
}
func (self *ResetAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_ResetAcknowledgeIEs = make(map[int]*RANAPPROTOCOLIES)

var order_ResetAcknowledgeIEs = make([]int, 3)

func (self *ResetAcknowledgeIEs) GetIECount() int{
   count := 0
   count +=1 //self.CNDomainIndicator
   if self.CriticalityDiagnostics != nil { count += 1 }
   if self.GlobalRNCID != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
      case 86: //GlobalRNCID
        self.GlobalRNCID = &GlobalRNCID{}
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, self.GlobalRNCID)
   }
}
func (self *ResetAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil {self.GlobalRNCID.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetAcknowledgeIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetAcknowledgeIEs[0] = 3
table_ResetAcknowledgeIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetAcknowledgeIEs[1] = 9
table_ResetAcknowledgeIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetAcknowledgeIEs[2] = 86
   }

type ResetAcknowledgeExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'EXTENSION': 'GlobalCN-ID', 'PRESENCE': 'optional'}, None]}
   GlobalCNID  *GlobalCNID
   list []interface{}
}
func (self *ResetAcknowledgeExtensions)createOT() interface{}{
    return nil
}
var table_ResetAcknowledgeExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_ResetAcknowledgeExtensions = make([]int, 1)

func (self *ResetAcknowledgeExtensions) GetIECount() int{
   count := 0
   if self.GlobalCNID != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetAcknowledgeExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetAcknowledgeExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
   }
}
func (self *ResetAcknowledgeExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetAcknowledgeExtensions[96] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetAcknowledgeExtensions[0] = 96
   }

type ResetResourceIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'reject', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-IuSigConIdList', 'CRITICALITY': 'ignore', 'TYPE': 'ResetResourceList', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'optional'}, None]}
   CNDomainIndicator  CNDomainIndicator
   Cause  Cause
   IuSigConIdList  ResetResourceList
   GlobalRNCID  *GlobalRNCID
   list []interface{}
}
func (self *ResetResourceIEs)createOT() interface{}{
    return nil
}
var table_ResetResourceIEs = make(map[int]*RANAPPROTOCOLIES)

var order_ResetResourceIEs = make([]int, 4)

func (self *ResetResourceIEs) GetIECount() int{
   count := 0
   count +=1 //self.CNDomainIndicator
   count +=1 //self.Cause
   count +=1 //self.IuSigConIdList
   if self.GlobalRNCID != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetResourceIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 4: //Cause
        return true //self.Cause
      case 77: //IuSigConIdList
        return true //self.IuSigConIdList
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetResourceIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 77: //IuSigConIdList
        self.IuSigConIdList.Unpack(st)
        self.list = append(self.list, &self.IuSigConIdList)
      case 86: //GlobalRNCID
        self.GlobalRNCID = &GlobalRNCID{}
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, self.GlobalRNCID)
   }
}
func (self *ResetResourceIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 4: //Cause
        self.Cause.Pack(st)
      case 77: //IuSigConIdList
        self.IuSigConIdList.Pack(st)
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil {self.GlobalRNCID.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetResourceIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetResourceIEs[0] = 3
table_ResetResourceIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetResourceIEs[1] = 4
table_ResetResourceIEs[77] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idIuSigConIdList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ResetResourceList{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetResourceIEs[2] = 77
table_ResetResourceIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetResourceIEs[3] = 86
   }

type ResetResourceItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-IuSigConIdItem', 'CRITICALITY': 'reject', 'TYPE': 'ResetResourceItem', 'PRESENCE': 'mandatory'}, None]}
   IuSigConIdItem  ResetResourceItem
   list []interface{}
}
func (self *ResetResourceItemIEs)createOT() interface{}{
    return nil
}
var table_ResetResourceItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_ResetResourceItemIEs = make([]int, 1)

func (self *ResetResourceItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.IuSigConIdItem
   return count//ObjSet
}
func (self *ResetResourceItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 78: //IuSigConIdItem
        return true //self.IuSigConIdItem
   }
   return false//ObjSet
}
func (self *ResetResourceItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 78: //IuSigConIdItem
        self.IuSigConIdItem.Unpack(st)
        self.list = append(self.list, &self.IuSigConIdItem)
   }
}
func (self *ResetResourceItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 78: //IuSigConIdItem
        self.IuSigConIdItem.Pack(st)
      default:
      break
   }
}
func init() {
table_ResetResourceItemIEs[78] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idIuSigConIdItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ResetResourceItem{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetResourceItemIEs[0] = 78
   }

type ResetResourceItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *ResetResourceItemExtIEs)createOT() interface{}{
    return nil
}
var table_ResetResourceItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_ResetResourceItemExtIEs = make([]int, 0)

type ResetResourceExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'EXTENSION': 'GlobalCN-ID', 'PRESENCE': 'optional'}, None]}
   GlobalCNID  *GlobalCNID
   list []interface{}
}
func (self *ResetResourceExtensions)createOT() interface{}{
    return nil
}
var table_ResetResourceExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_ResetResourceExtensions = make([]int, 1)

func (self *ResetResourceExtensions) GetIECount() int{
   count := 0
   if self.GlobalCNID != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetResourceExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetResourceExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
   }
}
func (self *ResetResourceExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetResourceExtensions[96] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetResourceExtensions[0] = 96
   }

type ResetResourceAcknowledgeIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'reject', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-IuSigConIdList', 'CRITICALITY': 'ignore', 'TYPE': 'ResetResourceAckList', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   CNDomainIndicator  CNDomainIndicator
   IuSigConIdList  ResetResourceAckList
   GlobalRNCID  *GlobalRNCID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *ResetResourceAcknowledgeIEs)createOT() interface{}{
    return nil
}
var table_ResetResourceAcknowledgeIEs = make(map[int]*RANAPPROTOCOLIES)

var order_ResetResourceAcknowledgeIEs = make([]int, 4)

func (self *ResetResourceAcknowledgeIEs) GetIECount() int{
   count := 0
   count +=1 //self.CNDomainIndicator
   count +=1 //self.IuSigConIdList
   if self.GlobalRNCID != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetResourceAcknowledgeIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 77: //IuSigConIdList
        return true //self.IuSigConIdList
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil { return true }
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetResourceAcknowledgeIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 77: //IuSigConIdList
        self.IuSigConIdList.Unpack(st)
        self.list = append(self.list, &self.IuSigConIdList)
      case 86: //GlobalRNCID
        self.GlobalRNCID = &GlobalRNCID{}
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, self.GlobalRNCID)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *ResetResourceAcknowledgeIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 77: //IuSigConIdList
        self.IuSigConIdList.Pack(st)
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil {self.GlobalRNCID.Pack(st)}
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetResourceAcknowledgeIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetResourceAcknowledgeIEs[0] = 3
table_ResetResourceAcknowledgeIEs[77] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idIuSigConIdList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&ResetResourceAckList{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetResourceAcknowledgeIEs[1] = 77
table_ResetResourceAcknowledgeIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetResourceAcknowledgeIEs[2] = 86
table_ResetResourceAcknowledgeIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetResourceAcknowledgeIEs[3] = 9
   }

type ResetResourceAckItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-IuSigConIdItem', 'CRITICALITY': 'reject', 'TYPE': 'ResetResourceAckItem', 'PRESENCE': 'mandatory'}, None]}
   IuSigConIdItem  ResetResourceAckItem
   list []interface{}
}
func (self *ResetResourceAckItemIEs)createOT() interface{}{
    return nil
}
var table_ResetResourceAckItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_ResetResourceAckItemIEs = make([]int, 1)

func (self *ResetResourceAckItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.IuSigConIdItem
   return count//ObjSet
}
func (self *ResetResourceAckItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 78: //IuSigConIdItem
        return true //self.IuSigConIdItem
   }
   return false//ObjSet
}
func (self *ResetResourceAckItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 78: //IuSigConIdItem
        self.IuSigConIdItem.Unpack(st)
        self.list = append(self.list, &self.IuSigConIdItem)
   }
}
func (self *ResetResourceAckItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 78: //IuSigConIdItem
        self.IuSigConIdItem.Pack(st)
      default:
      break
   }
}
func init() {
table_ResetResourceAckItemIEs[78] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idIuSigConIdItem}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ResetResourceAckItem{}, PRESENCE:Presence{Presencemandatory}, }
order_ResetResourceAckItemIEs[0] = 78
   }

type ResetResourceAckItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *ResetResourceAckItemExtIEs)createOT() interface{}{
    return nil
}
var table_ResetResourceAckItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_ResetResourceAckItemExtIEs = make([]int, 0)

type ResetResourceAcknowledgeExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'EXTENSION': 'GlobalCN-ID', 'PRESENCE': 'optional'}, None]}
   GlobalCNID  *GlobalCNID
   list []interface{}
}
func (self *ResetResourceAcknowledgeExtensions)createOT() interface{}{
    return nil
}
var table_ResetResourceAcknowledgeExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_ResetResourceAcknowledgeExtensions = make([]int, 1)

func (self *ResetResourceAcknowledgeExtensions) GetIECount() int{
   count := 0
   if self.GlobalCNID != nil { count += 1 }
   return count//ObjSet
}
func (self *ResetResourceAcknowledgeExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
   }
   return false//ObjSet
}
func (self *ResetResourceAcknowledgeExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
   }
}
func (self *ResetResourceAcknowledgeExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      default:
      break
   }
}
func init() {
table_ResetResourceAcknowledgeExtensions[96] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_ResetResourceAcknowledgeExtensions[0] = 96
   }

type RABReleaseRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ReleaseList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ReleaseList', 'PRESENCE': 'mandatory'}, None]}
   RABReleaseList  RABReleaseList
   list []interface{}
}
func (self *RABReleaseRequestIEs)createOT() interface{}{
    return nil
}
var table_RABReleaseRequestIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABReleaseRequestIEs = make([]int, 1)

func (self *RABReleaseRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABReleaseList
   return count//ObjSet
}
func (self *RABReleaseRequestIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 41: //RABReleaseList
        return true //self.RABReleaseList
   }
   return false//ObjSet
}
func (self *RABReleaseRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 41: //RABReleaseList
        self.RABReleaseList.Unpack(st)
        self.list = append(self.list, &self.RABReleaseList)
   }
}
func (self *RABReleaseRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 41: //RABReleaseList
        self.RABReleaseList.Pack(st)
      default:
      break
   }
}
func init() {
table_RABReleaseRequestIEs[41] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABReleaseList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABReleaseList{}, PRESENCE:Presence{Presencemandatory}, }
order_RABReleaseRequestIEs[0] = 41
   }

type RABReleaseItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ReleaseItem', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ReleaseItem', 'PRESENCE': 'mandatory'}, None]}
   RABReleaseItem  RABReleaseItem
   list []interface{}
}
func (self *RABReleaseItemIEs)createOT() interface{}{
    return nil
}
var table_RABReleaseItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABReleaseItemIEs = make([]int, 1)

func (self *RABReleaseItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABReleaseItem
   return count//ObjSet
}
func (self *RABReleaseItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 40: //RABReleaseItem
        return true //self.RABReleaseItem
   }
   return false//ObjSet
}
func (self *RABReleaseItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 40: //RABReleaseItem
        self.RABReleaseItem.Unpack(st)
        self.list = append(self.list, &self.RABReleaseItem)
   }
}
func (self *RABReleaseItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 40: //RABReleaseItem
        self.RABReleaseItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABReleaseItemIEs[40] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABReleaseItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABReleaseItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABReleaseItemIEs[0] = 40
   }

type RABReleaseItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABReleaseItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABReleaseItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABReleaseItemExtIEs = make([]int, 0)

type RABReleaseRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABReleaseRequestExtensions)createOT() interface{}{
    return nil
}
var table_RABReleaseRequestExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABReleaseRequestExtensions = make([]int, 0)

type IuReleaseRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, None]}
   Cause  Cause
   list []interface{}
}
func (self *IuReleaseRequestIEs)createOT() interface{}{
    return nil
}
var table_IuReleaseRequestIEs = make(map[int]*RANAPPROTOCOLIES)

var order_IuReleaseRequestIEs = make([]int, 1)

func (self *IuReleaseRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   return count//ObjSet
}
func (self *IuReleaseRequestIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        return true //self.Cause
   }
   return false//ObjSet
}
func (self *IuReleaseRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
   }
}
func (self *IuReleaseRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Pack(st)
      default:
      break
   }
}
func init() {
table_IuReleaseRequestIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_IuReleaseRequestIEs[0] = 4
   }

type IuReleaseRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *IuReleaseRequestExtensions)createOT() interface{}{
    return nil
}
var table_IuReleaseRequestExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_IuReleaseRequestExtensions = make([]int, 0)

type RelocationDetectIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *RelocationDetectIEs)createOT() interface{}{
    return nil
}
var table_RelocationDetectIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RelocationDetectIEs = make([]int, 0)

type RelocationDetectExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RelocationDetectExtensions)createOT() interface{}{
    return nil
}
var table_RelocationDetectExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RelocationDetectExtensions = make([]int, 0)

type RelocationCompleteIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [None]}
   list []interface{}
}
func (self *RelocationCompleteIEs)createOT() interface{}{
    return nil
}
var table_RelocationCompleteIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RelocationCompleteIEs = make([]int, 0)

type RelocationCompleteExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RelocationCompleteExtensions)createOT() interface{}{
    return nil
}
var table_RelocationCompleteExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RelocationCompleteExtensions = make([]int, 0)

type PagingIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'ignore', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-PermanentNAS-UE-ID', 'CRITICALITY': 'ignore', 'TYPE': 'PermanentNAS-UE-ID', 'PRESENCE': 'mandatory'}, {'ID': 'id-TemporaryUE-ID', 'CRITICALITY': 'ignore', 'TYPE': 'TemporaryUE-ID', 'PRESENCE': 'optional'}, {'ID': 'id-PagingAreaID', 'CRITICALITY': 'ignore', 'TYPE': 'PagingAreaID', 'PRESENCE': 'optional'}, {'ID': 'id-PagingCause', 'CRITICALITY': 'ignore', 'TYPE': 'PagingCause', 'PRESENCE': 'optional'}, {'ID': 'id-NonSearchingIndication', 'CRITICALITY': 'ignore', 'TYPE': 'NonSearchingIndication', 'PRESENCE': 'optional'}, {'ID': 'id-DRX-CycleLengthCoefficient', 'CRITICALITY': 'ignore', 'TYPE': 'DRX-CycleLengthCoefficient', 'PRESENCE': 'optional'}, None]}
   CNDomainIndicator  CNDomainIndicator
   PermanentNASUEID  PermanentNASUEID
   TemporaryUEID  *TemporaryUEID
   PagingAreaID  *PagingAreaID
   PagingCause  *PagingCause
   NonSearchingIndication  *NonSearchingIndication
   DRXCycleLengthCoefficient  *DRXCycleLengthCoefficient
   list []interface{}
}
func (self *PagingIEs)createOT() interface{}{
    return nil
}
var table_PagingIEs = make(map[int]*RANAPPROTOCOLIES)

var order_PagingIEs = make([]int, 7)

func (self *PagingIEs) GetIECount() int{
   count := 0
   count +=1 //self.CNDomainIndicator
   count +=1 //self.PermanentNASUEID
   if self.TemporaryUEID != nil { count += 1 }
   if self.PagingAreaID != nil { count += 1 }
   if self.PagingCause != nil { count += 1 }
   if self.NonSearchingIndication != nil { count += 1 }
   if self.DRXCycleLengthCoefficient != nil { count += 1 }
   return count//ObjSet
}
func (self *PagingIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 23: //PermanentNASUEID
        return true //self.PermanentNASUEID
      case 64: //TemporaryUEID
        if self.TemporaryUEID != nil { return true }
      case 21: //PagingAreaID
        if self.PagingAreaID != nil { return true }
      case 22: //PagingCause
        if self.PagingCause != nil { return true }
      case 17: //NonSearchingIndication
        if self.NonSearchingIndication != nil { return true }
      case 76: //DRXCycleLengthCoefficient
        if self.DRXCycleLengthCoefficient != nil { return true }
   }
   return false//ObjSet
}
func (self *PagingIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 23: //PermanentNASUEID
        self.PermanentNASUEID.Unpack(st)
        self.list = append(self.list, &self.PermanentNASUEID)
      case 64: //TemporaryUEID
        self.TemporaryUEID = &TemporaryUEID{}
        self.TemporaryUEID.Unpack(st)
        self.list = append(self.list, self.TemporaryUEID)
      case 21: //PagingAreaID
        self.PagingAreaID = &PagingAreaID{}
        self.PagingAreaID.Unpack(st)
        self.list = append(self.list, self.PagingAreaID)
      case 22: //PagingCause
        self.PagingCause = &PagingCause{}
        self.PagingCause.Unpack(st)
        self.list = append(self.list, self.PagingCause)
      case 17: //NonSearchingIndication
        self.NonSearchingIndication = &NonSearchingIndication{}
        self.NonSearchingIndication.Unpack(st)
        self.list = append(self.list, self.NonSearchingIndication)
      case 76: //DRXCycleLengthCoefficient
        self.DRXCycleLengthCoefficient = &DRXCycleLengthCoefficient{}
        self.DRXCycleLengthCoefficient.Unpack(st)
        self.list = append(self.list, self.DRXCycleLengthCoefficient)
   }
}
func (self *PagingIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 23: //PermanentNASUEID
        self.PermanentNASUEID.Pack(st)
      case 64: //TemporaryUEID
        if self.TemporaryUEID != nil {self.TemporaryUEID.Pack(st)}
      case 21: //PagingAreaID
        if self.PagingAreaID != nil {self.PagingAreaID.Pack(st)}
      case 22: //PagingCause
        if self.PagingCause != nil {self.PagingCause.Pack(st)}
      case 17: //NonSearchingIndication
        if self.NonSearchingIndication != nil {self.NonSearchingIndication.Pack(st)}
      case 76: //DRXCycleLengthCoefficient
        if self.DRXCycleLengthCoefficient != nil {self.DRXCycleLengthCoefficient.Pack(st)}
      default:
      break
   }
}
func init() {
table_PagingIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_PagingIEs[0] = 3
table_PagingIEs[23] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idPermanentNASUEID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PermanentNASUEID{}, PRESENCE:Presence{Presencemandatory}, }
order_PagingIEs[1] = 23
table_PagingIEs[64] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idTemporaryUEID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TemporaryUEID{}, PRESENCE:Presence{Presenceoptional}, }
order_PagingIEs[2] = 64
table_PagingIEs[21] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idPagingAreaID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PagingAreaID{}, PRESENCE:Presence{Presenceoptional}, }
order_PagingIEs[3] = 21
table_PagingIEs[22] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idPagingCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PagingCause{}, PRESENCE:Presence{Presenceoptional}, }
order_PagingIEs[4] = 22
table_PagingIEs[17] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idNonSearchingIndication}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&NonSearchingIndication{}, PRESENCE:Presence{Presenceoptional}, }
order_PagingIEs[5] = 17
table_PagingIEs[76] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idDRXCycleLengthCoefficient}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DRXCycleLengthCoefficient{}, PRESENCE:Presence{Presenceoptional}, }
order_PagingIEs[6] = 76
   }

type PagingExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'EXTENSION': 'GlobalCN-ID', 'PRESENCE': 'optional'}, None]}
   GlobalCNID  *GlobalCNID
   list []interface{}
}
func (self *PagingExtensions)createOT() interface{}{
    return nil
}
var table_PagingExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_PagingExtensions = make([]int, 1)

func (self *PagingExtensions) GetIECount() int{
   count := 0
   if self.GlobalCNID != nil { count += 1 }
   return count//ObjSet
}
func (self *PagingExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
   }
   return false//ObjSet
}
func (self *PagingExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
   }
}
func (self *PagingExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      default:
      break
   }
}
func init() {
table_PagingExtensions[96] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_PagingExtensions[0] = 96
   }

type CommonIDIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-PermanentNAS-UE-ID', 'CRITICALITY': 'ignore', 'TYPE': 'PermanentNAS-UE-ID', 'PRESENCE': 'mandatory'}, None]}
   PermanentNASUEID  PermanentNASUEID
   list []interface{}
}
func (self *CommonIDIEs)createOT() interface{}{
    return nil
}
var table_CommonIDIEs = make(map[int]*RANAPPROTOCOLIES)

var order_CommonIDIEs = make([]int, 1)

func (self *CommonIDIEs) GetIECount() int{
   count := 0
   count +=1 //self.PermanentNASUEID
   return count//ObjSet
}
func (self *CommonIDIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 23: //PermanentNASUEID
        return true //self.PermanentNASUEID
   }
   return false//ObjSet
}
func (self *CommonIDIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 23: //PermanentNASUEID
        self.PermanentNASUEID.Unpack(st)
        self.list = append(self.list, &self.PermanentNASUEID)
   }
}
func (self *CommonIDIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 23: //PermanentNASUEID
        self.PermanentNASUEID.Pack(st)
      default:
      break
   }
}
func init() {
table_CommonIDIEs[23] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idPermanentNASUEID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&PermanentNASUEID{}, PRESENCE:Presence{Presencemandatory}, }
order_CommonIDIEs[0] = 23
   }

type CommonIDExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-SNA-Access-Information', 'CRITICALITY': 'ignore', 'EXTENSION': 'SNA-Access-Information', 'PRESENCE': 'optional'}, {'ID': 'id-UESBI-Iu', 'CRITICALITY': 'ignore', 'EXTENSION': 'UESBI-Iu', 'PRESENCE': 'optional'}, None]}
   SNAAccessInformation  *SNAAccessInformation
   UESBIIu  *UESBIIu
   list []interface{}
}
func (self *CommonIDExtensions)createOT() interface{}{
    return nil
}
var table_CommonIDExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_CommonIDExtensions = make([]int, 2)

func (self *CommonIDExtensions) GetIECount() int{
   count := 0
   if self.SNAAccessInformation != nil { count += 1 }
   if self.UESBIIu != nil { count += 1 }
   return count//ObjSet
}
func (self *CommonIDExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 105: //SNAAccessInformation
        if self.SNAAccessInformation != nil { return true }
      case 118: //UESBIIu
        if self.UESBIIu != nil { return true }
   }
   return false//ObjSet
}
func (self *CommonIDExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 105: //SNAAccessInformation
        self.SNAAccessInformation = &SNAAccessInformation{}
        self.SNAAccessInformation.Unpack(st)
        self.list = append(self.list, self.SNAAccessInformation)
      case 118: //UESBIIu
        self.UESBIIu = &UESBIIu{}
        self.UESBIIu.Unpack(st)
        self.list = append(self.list, self.UESBIIu)
   }
}
func (self *CommonIDExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 105: //SNAAccessInformation
        if self.SNAAccessInformation != nil {self.SNAAccessInformation.Pack(st)}
      case 118: //UESBIIu
        if self.UESBIIu != nil {self.UESBIIu.Pack(st)}
      default:
      break
   }
}
func init() {
table_CommonIDExtensions[105] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idSNAAccessInformation}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&SNAAccessInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_CommonIDExtensions[0] = 105
table_CommonIDExtensions[118] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idUESBIIu}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&UESBIIu{}, PRESENCE:Presence{Presenceoptional}, }
order_CommonIDExtensions[1] = 118
   }

type CNInvokeTraceIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-TraceType', 'CRITICALITY': 'ignore', 'TYPE': 'TraceType', 'PRESENCE': 'optional'}, {'ID': 'id-TraceReference', 'CRITICALITY': 'ignore', 'TYPE': 'TraceReference', 'PRESENCE': 'mandatory'}, {'ID': 'id-TriggerID', 'CRITICALITY': 'ignore', 'TYPE': 'TriggerID', 'PRESENCE': 'optional'}, {'ID': 'id-UE-ID', 'CRITICALITY': 'ignore', 'TYPE': 'UE-ID', 'PRESENCE': 'optional'}, {'ID': 'id-OMC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'OMC-ID', 'PRESENCE': 'optional'}, None]}
   TraceType  *TraceType
   TraceReference  TraceReference
   TriggerID  *TriggerID
   UEID  *UEID
   OMCID  *OMCID
   list []interface{}
}
func (self *CNInvokeTraceIEs)createOT() interface{}{
    return nil
}
var table_CNInvokeTraceIEs = make(map[int]*RANAPPROTOCOLIES)

var order_CNInvokeTraceIEs = make([]int, 5)

func (self *CNInvokeTraceIEs) GetIECount() int{
   count := 0
   if self.TraceType != nil { count += 1 }
   count +=1 //self.TraceReference
   if self.TriggerID != nil { count += 1 }
   if self.UEID != nil { count += 1 }
   if self.OMCID != nil { count += 1 }
   return count//ObjSet
}
func (self *CNInvokeTraceIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 66: //TraceType
        if self.TraceType != nil { return true }
      case 65: //TraceReference
        return true //self.TraceReference
      case 68: //TriggerID
        if self.TriggerID != nil { return true }
      case 69: //UEID
        if self.UEID != nil { return true }
      case 19: //OMCID
        if self.OMCID != nil { return true }
   }
   return false//ObjSet
}
func (self *CNInvokeTraceIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 66: //TraceType
        self.TraceType = &TraceType{}
        self.TraceType.Unpack(st)
        self.list = append(self.list, self.TraceType)
      case 65: //TraceReference
        self.TraceReference.Unpack(st)
        self.list = append(self.list, &self.TraceReference)
      case 68: //TriggerID
        self.TriggerID = &TriggerID{}
        self.TriggerID.Unpack(st)
        self.list = append(self.list, self.TriggerID)
      case 69: //UEID
        self.UEID = &UEID{}
        self.UEID.Unpack(st)
        self.list = append(self.list, self.UEID)
      case 19: //OMCID
        self.OMCID = &OMCID{}
        self.OMCID.Unpack(st)
        self.list = append(self.list, self.OMCID)
   }
}
func (self *CNInvokeTraceIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 66: //TraceType
        if self.TraceType != nil {self.TraceType.Pack(st)}
      case 65: //TraceReference
        self.TraceReference.Pack(st)
      case 68: //TriggerID
        if self.TriggerID != nil {self.TriggerID.Pack(st)}
      case 69: //UEID
        if self.UEID != nil {self.UEID.Pack(st)}
      case 19: //OMCID
        if self.OMCID != nil {self.OMCID.Pack(st)}
      default:
      break
   }
}
func init() {
table_CNInvokeTraceIEs[66] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idTraceType}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TraceType{}, PRESENCE:Presence{Presenceoptional}, }
order_CNInvokeTraceIEs[0] = 66
table_CNInvokeTraceIEs[65] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idTraceReference}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TraceReference{}, PRESENCE:Presence{Presencemandatory}, }
order_CNInvokeTraceIEs[1] = 65
table_CNInvokeTraceIEs[68] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idTriggerID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TriggerID{}, PRESENCE:Presence{Presenceoptional}, }
order_CNInvokeTraceIEs[2] = 68
table_CNInvokeTraceIEs[69] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idUEID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&UEID{}, PRESENCE:Presence{Presenceoptional}, }
order_CNInvokeTraceIEs[3] = 69
table_CNInvokeTraceIEs[19] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idOMCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&OMCID{}, PRESENCE:Presence{Presenceoptional}, }
order_CNInvokeTraceIEs[4] = 19
   }

type CNInvokeTraceExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-TracePropagationParameters', 'CRITICALITY': 'ignore', 'EXTENSION': 'TracePropagationParameters', 'PRESENCE': 'optional'}, None]}
   TracePropagationParameters  *TracePropagationParameters
   list []interface{}
}
func (self *CNInvokeTraceExtensions)createOT() interface{}{
    return nil
}
var table_CNInvokeTraceExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_CNInvokeTraceExtensions = make([]int, 1)

func (self *CNInvokeTraceExtensions) GetIECount() int{
   count := 0
   if self.TracePropagationParameters != nil { count += 1 }
   return count//ObjSet
}
func (self *CNInvokeTraceExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 125: //TracePropagationParameters
        if self.TracePropagationParameters != nil { return true }
   }
   return false//ObjSet
}
func (self *CNInvokeTraceExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 125: //TracePropagationParameters
        self.TracePropagationParameters = &TracePropagationParameters{}
        self.TracePropagationParameters.Unpack(st)
        self.list = append(self.list, self.TracePropagationParameters)
   }
}
func (self *CNInvokeTraceExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 125: //TracePropagationParameters
        if self.TracePropagationParameters != nil {self.TracePropagationParameters.Pack(st)}
      default:
      break
   }
}
func init() {
table_CNInvokeTraceExtensions[125] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idTracePropagationParameters}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&TracePropagationParameters{}, PRESENCE:Presence{Presenceoptional}, }
order_CNInvokeTraceExtensions[0] = 125
   }

type CNDeactivateTraceIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-TraceReference', 'CRITICALITY': 'ignore', 'TYPE': 'TraceReference', 'PRESENCE': 'mandatory'}, {'ID': 'id-TriggerID', 'CRITICALITY': 'ignore', 'TYPE': 'TriggerID', 'PRESENCE': 'optional'}, None]}
   TraceReference  TraceReference
   TriggerID  *TriggerID
   list []interface{}
}
func (self *CNDeactivateTraceIEs)createOT() interface{}{
    return nil
}
var table_CNDeactivateTraceIEs = make(map[int]*RANAPPROTOCOLIES)

var order_CNDeactivateTraceIEs = make([]int, 2)

func (self *CNDeactivateTraceIEs) GetIECount() int{
   count := 0
   count +=1 //self.TraceReference
   if self.TriggerID != nil { count += 1 }
   return count//ObjSet
}
func (self *CNDeactivateTraceIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 65: //TraceReference
        return true //self.TraceReference
      case 68: //TriggerID
        if self.TriggerID != nil { return true }
   }
   return false//ObjSet
}
func (self *CNDeactivateTraceIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 65: //TraceReference
        self.TraceReference.Unpack(st)
        self.list = append(self.list, &self.TraceReference)
      case 68: //TriggerID
        self.TriggerID = &TriggerID{}
        self.TriggerID.Unpack(st)
        self.list = append(self.list, self.TriggerID)
   }
}
func (self *CNDeactivateTraceIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 65: //TraceReference
        self.TraceReference.Pack(st)
      case 68: //TriggerID
        if self.TriggerID != nil {self.TriggerID.Pack(st)}
      default:
      break
   }
}
func init() {
table_CNDeactivateTraceIEs[65] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idTraceReference}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TraceReference{}, PRESENCE:Presence{Presencemandatory}, }
order_CNDeactivateTraceIEs[0] = 65
table_CNDeactivateTraceIEs[68] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idTriggerID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&TriggerID{}, PRESENCE:Presence{Presenceoptional}, }
order_CNDeactivateTraceIEs[1] = 68
   }

type CNDeactivateTraceExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CNDeactivateTraceExtensions)createOT() interface{}{
    return nil
}
var table_CNDeactivateTraceExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_CNDeactivateTraceExtensions = make([]int, 0)

type LocationReportingControlIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RequestType', 'CRITICALITY': 'ignore', 'TYPE': 'RequestType', 'PRESENCE': 'mandatory'}, None]}
   RequestType  RequestType
   list []interface{}
}
func (self *LocationReportingControlIEs)createOT() interface{}{
    return nil
}
var table_LocationReportingControlIEs = make(map[int]*RANAPPROTOCOLIES)

var order_LocationReportingControlIEs = make([]int, 1)

func (self *LocationReportingControlIEs) GetIECount() int{
   count := 0
   count +=1 //self.RequestType
   return count//ObjSet
}
func (self *LocationReportingControlIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 57: //RequestType
        return true //self.RequestType
   }
   return false//ObjSet
}
func (self *LocationReportingControlIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 57: //RequestType
        self.RequestType.Unpack(st)
        self.list = append(self.list, &self.RequestType)
   }
}
func (self *LocationReportingControlIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 57: //RequestType
        self.RequestType.Pack(st)
      default:
      break
   }
}
func init() {
table_LocationReportingControlIEs[57] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRequestType}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RequestType{}, PRESENCE:Presence{Presencemandatory}, }
order_LocationReportingControlIEs[0] = 57
   }

type LocationReportingControlExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-VerticalAccuracyCode', 'CRITICALITY': 'ignore', 'EXTENSION': 'VerticalAccuracyCode', 'PRESENCE': 'optional'}, {'ID': 'id-ResponseTime', 'CRITICALITY': 'ignore', 'EXTENSION': 'ResponseTime', 'PRESENCE': 'optional'}, {'ID': 'id-PositioningPriority', 'CRITICALITY': 'ignore', 'EXTENSION': 'PositioningPriority', 'PRESENCE': 'optional'}, {'ID': 'id-ClientType', 'CRITICALITY': 'ignore', 'EXTENSION': 'ClientType', 'PRESENCE': 'optional'}, None]}
   VerticalAccuracyCode  *VerticalAccuracyCode
   ResponseTime  *ResponseTime
   PositioningPriority  *PositioningPriority
   ClientType  *ClientType
   list []interface{}
}
func (self *LocationReportingControlExtensions)createOT() interface{}{
    return nil
}
var table_LocationReportingControlExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_LocationReportingControlExtensions = make([]int, 4)

func (self *LocationReportingControlExtensions) GetIECount() int{
   count := 0
   if self.VerticalAccuracyCode != nil { count += 1 }
   if self.ResponseTime != nil { count += 1 }
   if self.PositioningPriority != nil { count += 1 }
   if self.ClientType != nil { count += 1 }
   return count//ObjSet
}
func (self *LocationReportingControlExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 111: //VerticalAccuracyCode
        if self.VerticalAccuracyCode != nil { return true }
      case 112: //ResponseTime
        if self.ResponseTime != nil { return true }
      case 113: //PositioningPriority
        if self.PositioningPriority != nil { return true }
      case 114: //ClientType
        if self.ClientType != nil { return true }
   }
   return false//ObjSet
}
func (self *LocationReportingControlExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 111: //VerticalAccuracyCode
        self.VerticalAccuracyCode = &VerticalAccuracyCode{}
        self.VerticalAccuracyCode.Unpack(st)
        self.list = append(self.list, self.VerticalAccuracyCode)
      case 112: //ResponseTime
        self.ResponseTime = &ResponseTime{}
        self.ResponseTime.Unpack(st)
        self.list = append(self.list, self.ResponseTime)
      case 113: //PositioningPriority
        self.PositioningPriority = &PositioningPriority{}
        self.PositioningPriority.Unpack(st)
        self.list = append(self.list, self.PositioningPriority)
      case 114: //ClientType
        self.ClientType = &ClientType{}
        self.ClientType.Unpack(st)
        self.list = append(self.list, self.ClientType)
   }
}
func (self *LocationReportingControlExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 111: //VerticalAccuracyCode
        if self.VerticalAccuracyCode != nil {self.VerticalAccuracyCode.Pack(st)}
      case 112: //ResponseTime
        if self.ResponseTime != nil {self.ResponseTime.Pack(st)}
      case 113: //PositioningPriority
        if self.PositioningPriority != nil {self.PositioningPriority.Pack(st)}
      case 114: //ClientType
        if self.ClientType != nil {self.ClientType.Pack(st)}
      default:
      break
   }
}
func init() {
table_LocationReportingControlExtensions[111] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idVerticalAccuracyCode}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&VerticalAccuracyCode{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportingControlExtensions[0] = 111
table_LocationReportingControlExtensions[112] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idResponseTime}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&ResponseTime{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportingControlExtensions[1] = 112
table_LocationReportingControlExtensions[113] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idPositioningPriority}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&PositioningPriority{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportingControlExtensions[2] = 113
table_LocationReportingControlExtensions[114] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idClientType}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&ClientType{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportingControlExtensions[3] = 114
   }

type LocationReportIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-AreaIdentity', 'CRITICALITY': 'ignore', 'TYPE': 'AreaIdentity', 'PRESENCE': 'optional'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'optional'}, {'ID': 'id-RequestType', 'CRITICALITY': 'ignore', 'TYPE': 'RequestType', 'PRESENCE': 'optional'}, None]}
   AreaIdentity  *AreaIdentity
   Cause  *Cause
   RequestType  *RequestType
   list []interface{}
}
func (self *LocationReportIEs)createOT() interface{}{
    return nil
}
var table_LocationReportIEs = make(map[int]*RANAPPROTOCOLIES)

var order_LocationReportIEs = make([]int, 3)

func (self *LocationReportIEs) GetIECount() int{
   count := 0
   if self.AreaIdentity != nil { count += 1 }
   if self.Cause != nil { count += 1 }
   if self.RequestType != nil { count += 1 }
   return count//ObjSet
}
func (self *LocationReportIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 0: //AreaIdentity
        if self.AreaIdentity != nil { return true }
      case 4: //Cause
        if self.Cause != nil { return true }
      case 57: //RequestType
        if self.RequestType != nil { return true }
   }
   return false//ObjSet
}
func (self *LocationReportIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 0: //AreaIdentity
        self.AreaIdentity = &AreaIdentity{}
        self.AreaIdentity.Unpack(st)
        self.list = append(self.list, self.AreaIdentity)
      case 4: //Cause
        self.Cause = &Cause{}
        self.Cause.Unpack(st)
        self.list = append(self.list, self.Cause)
      case 57: //RequestType
        self.RequestType = &RequestType{}
        self.RequestType.Unpack(st)
        self.list = append(self.list, self.RequestType)
   }
}
func (self *LocationReportIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 0: //AreaIdentity
        if self.AreaIdentity != nil {self.AreaIdentity.Pack(st)}
      case 4: //Cause
        if self.Cause != nil {self.Cause.Pack(st)}
      case 57: //RequestType
        if self.RequestType != nil {self.RequestType.Pack(st)}
      default:
      break
   }
}
func init() {
table_LocationReportIEs[0] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idAreaIdentity}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&AreaIdentity{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportIEs[0] = 0
table_LocationReportIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportIEs[1] = 4
table_LocationReportIEs[57] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRequestType}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RequestType{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportIEs[2] = 57
   }

type LocationReportExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-LastKnownServiceArea', 'CRITICALITY': 'ignore', 'EXTENSION': 'LastKnownServiceArea', 'PRESENCE': 'optional'}, {'ID': 'id-PositionData', 'CRITICALITY': 'ignore', 'EXTENSION': 'PositionData', 'PRESENCE': 'optional'}, {'ID': 'id-PositionDataSpecificToGERANIuMode', 'CRITICALITY': 'ignore', 'EXTENSION': 'PositionDataSpecificToGERANIuMode', 'PRESENCE': 'optional'}, {'ID': 'id-AccuracyFulfilmentIndicator', 'CRITICALITY': 'ignore', 'EXTENSION': 'AccuracyFulfilmentIndicator', 'PRESENCE': 'optional'}, None]}
   LastKnownServiceArea  *LastKnownServiceArea
   PositionData  *PositionData
   PositionDataSpecificToGERANIuMode  *PositionDataSpecificToGERANIuMode
   AccuracyFulfilmentIndicator  *AccuracyFulfilmentIndicator
   list []interface{}
}
func (self *LocationReportExtensions)createOT() interface{}{
    return nil
}
var table_LocationReportExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_LocationReportExtensions = make([]int, 4)

func (self *LocationReportExtensions) GetIECount() int{
   count := 0
   if self.LastKnownServiceArea != nil { count += 1 }
   if self.PositionData != nil { count += 1 }
   if self.PositionDataSpecificToGERANIuMode != nil { count += 1 }
   if self.AccuracyFulfilmentIndicator != nil { count += 1 }
   return count//ObjSet
}
func (self *LocationReportExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 97: //LastKnownServiceArea
        if self.LastKnownServiceArea != nil { return true }
      case 119: //PositionData
        if self.PositionData != nil { return true }
      case 120: //PositionDataSpecificToGERANIuMode
        if self.PositionDataSpecificToGERANIuMode != nil { return true }
      case 122: //AccuracyFulfilmentIndicator
        if self.AccuracyFulfilmentIndicator != nil { return true }
   }
   return false//ObjSet
}
func (self *LocationReportExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 97: //LastKnownServiceArea
        self.LastKnownServiceArea = &LastKnownServiceArea{}
        self.LastKnownServiceArea.Unpack(st)
        self.list = append(self.list, self.LastKnownServiceArea)
      case 119: //PositionData
        self.PositionData = &PositionData{}
        self.PositionData.Unpack(st)
        self.list = append(self.list, self.PositionData)
      case 120: //PositionDataSpecificToGERANIuMode
        self.PositionDataSpecificToGERANIuMode = &PositionDataSpecificToGERANIuMode{}
        self.PositionDataSpecificToGERANIuMode.Unpack(st)
        self.list = append(self.list, self.PositionDataSpecificToGERANIuMode)
      case 122: //AccuracyFulfilmentIndicator
        self.AccuracyFulfilmentIndicator = &AccuracyFulfilmentIndicator{}
        self.AccuracyFulfilmentIndicator.Unpack(st)
        self.list = append(self.list, self.AccuracyFulfilmentIndicator)
   }
}
func (self *LocationReportExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 97: //LastKnownServiceArea
        if self.LastKnownServiceArea != nil {self.LastKnownServiceArea.Pack(st)}
      case 119: //PositionData
        if self.PositionData != nil {self.PositionData.Pack(st)}
      case 120: //PositionDataSpecificToGERANIuMode
        if self.PositionDataSpecificToGERANIuMode != nil {self.PositionDataSpecificToGERANIuMode.Pack(st)}
      case 122: //AccuracyFulfilmentIndicator
        if self.AccuracyFulfilmentIndicator != nil {self.AccuracyFulfilmentIndicator.Pack(st)}
      default:
      break
   }
}
func init() {
table_LocationReportExtensions[97] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idLastKnownServiceArea}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&LastKnownServiceArea{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportExtensions[0] = 97
table_LocationReportExtensions[119] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idPositionData}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&PositionData{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportExtensions[1] = 119
table_LocationReportExtensions[120] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idPositionDataSpecificToGERANIuMode}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&PositionDataSpecificToGERANIuMode{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportExtensions[2] = 120
table_LocationReportExtensions[122] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idAccuracyFulfilmentIndicator}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&AccuracyFulfilmentIndicator{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationReportExtensions[3] = 122
   }

type InitialUEMessageIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'ignore', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-LAI', 'CRITICALITY': 'ignore', 'TYPE': 'LAI', 'PRESENCE': 'mandatory'}, {'ID': 'id-RAC', 'CRITICALITY': 'ignore', 'TYPE': 'RAC', 'PRESENCE': 'conditional'}, {'ID': 'id-SAI', 'CRITICALITY': 'ignore', 'TYPE': 'SAI', 'PRESENCE': 'mandatory'}, {'ID': 'id-NAS-PDU', 'CRITICALITY': 'ignore', 'TYPE': 'NAS-PDU', 'PRESENCE': 'mandatory'}, {'ID': 'id-IuSigConId', 'CRITICALITY': 'ignore', 'TYPE': 'IuSignallingConnectionIdentifier', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'mandatory'}, None]}
   CNDomainIndicator  CNDomainIndicator
   LAI  LAI
   RAC  RAC
   SAI  SAI
   NASPDU  NASPDU
   IuSigConId  IuSignallingConnectionIdentifier
   GlobalRNCID  GlobalRNCID
   list []interface{}
}
func (self *InitialUEMessageIEs)createOT() interface{}{
    return nil
}
var table_InitialUEMessageIEs = make(map[int]*RANAPPROTOCOLIES)

var order_InitialUEMessageIEs = make([]int, 7)

func (self *InitialUEMessageIEs) GetIECount() int{
   count := 0
   count +=1 //self.CNDomainIndicator
   count +=1 //self.LAI
   count +=1 //self.RAC
   count +=1 //self.SAI
   count +=1 //self.NASPDU
   count +=1 //self.IuSigConId
   count +=1 //self.GlobalRNCID
   return count//ObjSet
}
func (self *InitialUEMessageIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 15: //LAI
        return true //self.LAI
      case 55: //RAC
        return true //self.RAC
      case 58: //SAI
        return true //self.SAI
      case 16: //NASPDU
        return true //self.NASPDU
      case 79: //IuSigConId
        return true //self.IuSigConId
      case 86: //GlobalRNCID
        return true //self.GlobalRNCID
   }
   return false//ObjSet
}
func (self *InitialUEMessageIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 15: //LAI
        self.LAI.Unpack(st)
        self.list = append(self.list, &self.LAI)
      case 55: //RAC
        self.RAC.Unpack(st)
        self.list = append(self.list, &self.RAC)
      case 58: //SAI
        self.SAI.Unpack(st)
        self.list = append(self.list, &self.SAI)
      case 16: //NASPDU
        self.NASPDU.Unpack(st)
        self.list = append(self.list, &self.NASPDU)
      case 79: //IuSigConId
        self.IuSigConId.Unpack(st)
        self.list = append(self.list, &self.IuSigConId)
      case 86: //GlobalRNCID
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, &self.GlobalRNCID)
   }
}
func (self *InitialUEMessageIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 15: //LAI
        self.LAI.Pack(st)
      case 55: //RAC
        self.RAC.Pack(st)
      case 58: //SAI
        self.SAI.Pack(st)
      case 16: //NASPDU
        self.NASPDU.Pack(st)
      case 79: //IuSigConId
        self.IuSigConId.Pack(st)
      case 86: //GlobalRNCID
        self.GlobalRNCID.Pack(st)
      default:
      break
   }
}
func init() {
table_InitialUEMessageIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_InitialUEMessageIEs[0] = 3
table_InitialUEMessageIEs[15] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idLAI}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&LAI{}, PRESENCE:Presence{Presencemandatory}, }
order_InitialUEMessageIEs[1] = 15
table_InitialUEMessageIEs[55] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRAC}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RAC{}, PRESENCE:Presence{Presenceconditional}, }
order_InitialUEMessageIEs[2] = 55
table_InitialUEMessageIEs[58] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idSAI}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&SAI{}, PRESENCE:Presence{Presencemandatory}, }
order_InitialUEMessageIEs[3] = 58
table_InitialUEMessageIEs[16] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idNASPDU}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&NASPDU{}, PRESENCE:Presence{Presencemandatory}, }
order_InitialUEMessageIEs[4] = 16
table_InitialUEMessageIEs[79] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idIuSigConId}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&IuSignallingConnectionIdentifier{}, PRESENCE:Presence{Presencemandatory}, }
order_InitialUEMessageIEs[5] = 79
table_InitialUEMessageIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presencemandatory}, }
order_InitialUEMessageIEs[6] = 86
   }

type InitialUEMessageExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-GERAN-Classmark', 'CRITICALITY': 'ignore', 'EXTENSION': 'GERAN-Classmark', 'PRESENCE': 'optional'}, None]}
   GERANClassmark  *GERANClassmark
   list []interface{}
}
func (self *InitialUEMessageExtensions)createOT() interface{}{
    return nil
}
var table_InitialUEMessageExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_InitialUEMessageExtensions = make([]int, 1)

func (self *InitialUEMessageExtensions) GetIECount() int{
   count := 0
   if self.GERANClassmark != nil { count += 1 }
   return count//ObjSet
}
func (self *InitialUEMessageExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 108: //GERANClassmark
        if self.GERANClassmark != nil { return true }
   }
   return false//ObjSet
}
func (self *InitialUEMessageExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 108: //GERANClassmark
        self.GERANClassmark = &GERANClassmark{}
        self.GERANClassmark.Unpack(st)
        self.list = append(self.list, self.GERANClassmark)
   }
}
func (self *InitialUEMessageExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 108: //GERANClassmark
        if self.GERANClassmark != nil {self.GERANClassmark.Pack(st)}
      default:
      break
   }
}
func init() {
table_InitialUEMessageExtensions[108] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGERANClassmark}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GERANClassmark{}, PRESENCE:Presence{Presenceoptional}, }
order_InitialUEMessageExtensions[0] = 108
   }

type DirectTransferIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-NAS-PDU', 'CRITICALITY': 'ignore', 'TYPE': 'NAS-PDU', 'PRESENCE': 'mandatory'}, {'ID': 'id-LAI', 'CRITICALITY': 'ignore', 'TYPE': 'LAI', 'PRESENCE': 'optional'}, {'ID': 'id-RAC', 'CRITICALITY': 'ignore', 'TYPE': 'RAC', 'PRESENCE': 'optional'}, {'ID': 'id-SAI', 'CRITICALITY': 'ignore', 'TYPE': 'SAI', 'PRESENCE': 'optional'}, {'ID': 'id-SAPI', 'CRITICALITY': 'ignore', 'TYPE': 'SAPI', 'PRESENCE': 'optional'}, None]}
   NASPDU  NASPDU
   LAI  *LAI
   RAC  *RAC
   SAI  *SAI
   SAPI  *SAPI
   list []interface{}
}
func (self *DirectTransferIEs)createOT() interface{}{
    return nil
}
var table_DirectTransferIEs = make(map[int]*RANAPPROTOCOLIES)

var order_DirectTransferIEs = make([]int, 5)

func (self *DirectTransferIEs) GetIECount() int{
   count := 0
   count +=1 //self.NASPDU
   if self.LAI != nil { count += 1 }
   if self.RAC != nil { count += 1 }
   if self.SAI != nil { count += 1 }
   if self.SAPI != nil { count += 1 }
   return count//ObjSet
}
func (self *DirectTransferIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 16: //NASPDU
        return true //self.NASPDU
      case 15: //LAI
        if self.LAI != nil { return true }
      case 55: //RAC
        if self.RAC != nil { return true }
      case 58: //SAI
        if self.SAI != nil { return true }
      case 59: //SAPI
        if self.SAPI != nil { return true }
   }
   return false//ObjSet
}
func (self *DirectTransferIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 16: //NASPDU
        self.NASPDU.Unpack(st)
        self.list = append(self.list, &self.NASPDU)
      case 15: //LAI
        self.LAI = &LAI{}
        self.LAI.Unpack(st)
        self.list = append(self.list, self.LAI)
      case 55: //RAC
        self.RAC = &RAC{}
        self.RAC.Unpack(st)
        self.list = append(self.list, self.RAC)
      case 58: //SAI
        self.SAI = &SAI{}
        self.SAI.Unpack(st)
        self.list = append(self.list, self.SAI)
      case 59: //SAPI
        self.SAPI = &SAPI{}
        self.SAPI.Unpack(st)
        self.list = append(self.list, self.SAPI)
   }
}
func (self *DirectTransferIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 16: //NASPDU
        self.NASPDU.Pack(st)
      case 15: //LAI
        if self.LAI != nil {self.LAI.Pack(st)}
      case 55: //RAC
        if self.RAC != nil {self.RAC.Pack(st)}
      case 58: //SAI
        if self.SAI != nil {self.SAI.Pack(st)}
      case 59: //SAPI
        if self.SAPI != nil {self.SAPI.Pack(st)}
      default:
      break
   }
}
func init() {
table_DirectTransferIEs[16] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idNASPDU}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&NASPDU{}, PRESENCE:Presence{Presencemandatory}, }
order_DirectTransferIEs[0] = 16
table_DirectTransferIEs[15] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idLAI}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&LAI{}, PRESENCE:Presence{Presenceoptional}, }
order_DirectTransferIEs[1] = 15
table_DirectTransferIEs[55] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRAC}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RAC{}, PRESENCE:Presence{Presenceoptional}, }
order_DirectTransferIEs[2] = 55
table_DirectTransferIEs[58] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idSAI}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&SAI{}, PRESENCE:Presence{Presenceoptional}, }
order_DirectTransferIEs[3] = 58
table_DirectTransferIEs[59] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idSAPI}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&SAPI{}, PRESENCE:Presence{Presenceoptional}, }
order_DirectTransferIEs[4] = 59
   }

type DirectTransferExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DirectTransferExtensions)createOT() interface{}{
    return nil
}
var table_DirectTransferExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_DirectTransferExtensions = make([]int, 0)

type OverloadIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-NumberOfSteps', 'CRITICALITY': 'ignore', 'TYPE': 'NumberOfSteps', 'PRESENCE': 'optional'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'optional'}, None]}
   NumberOfSteps  *NumberOfSteps
   GlobalRNCID  *GlobalRNCID
   list []interface{}
}
func (self *OverloadIEs)createOT() interface{}{
    return nil
}
var table_OverloadIEs = make(map[int]*RANAPPROTOCOLIES)

var order_OverloadIEs = make([]int, 2)

func (self *OverloadIEs) GetIECount() int{
   count := 0
   if self.NumberOfSteps != nil { count += 1 }
   if self.GlobalRNCID != nil { count += 1 }
   return count//ObjSet
}
func (self *OverloadIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 18: //NumberOfSteps
        if self.NumberOfSteps != nil { return true }
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil { return true }
   }
   return false//ObjSet
}
func (self *OverloadIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 18: //NumberOfSteps
        self.NumberOfSteps = &NumberOfSteps{}
        self.NumberOfSteps.Unpack(st)
        self.list = append(self.list, self.NumberOfSteps)
      case 86: //GlobalRNCID
        self.GlobalRNCID = &GlobalRNCID{}
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, self.GlobalRNCID)
   }
}
func (self *OverloadIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 18: //NumberOfSteps
        if self.NumberOfSteps != nil {self.NumberOfSteps.Pack(st)}
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil {self.GlobalRNCID.Pack(st)}
      default:
      break
   }
}
func init() {
table_OverloadIEs[18] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idNumberOfSteps}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&NumberOfSteps{}, PRESENCE:Presence{Presenceoptional}, }
order_OverloadIEs[0] = 18
table_OverloadIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presenceoptional}, }
order_OverloadIEs[1] = 86
   }

type OverloadExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'ignore', 'EXTENSION': 'CN-DomainIndicator', 'PRESENCE': 'optional'}, {'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'EXTENSION': 'GlobalCN-ID', 'PRESENCE': 'optional'}, None]}
   CNDomainIndicator  *CNDomainIndicator
   GlobalCNID  *GlobalCNID
   list []interface{}
}
func (self *OverloadExtensions)createOT() interface{}{
    return nil
}
var table_OverloadExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_OverloadExtensions = make([]int, 2)

func (self *OverloadExtensions) GetIECount() int{
   count := 0
   if self.CNDomainIndicator != nil { count += 1 }
   if self.GlobalCNID != nil { count += 1 }
   return count//ObjSet
}
func (self *OverloadExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 3: //CNDomainIndicator
        if self.CNDomainIndicator != nil { return true }
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
   }
   return false//ObjSet
}
func (self *OverloadExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator = &CNDomainIndicator{}
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, self.CNDomainIndicator)
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
   }
}
func (self *OverloadExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 3: //CNDomainIndicator
        if self.CNDomainIndicator != nil {self.CNDomainIndicator.Pack(st)}
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      default:
      break
   }
}
func init() {
table_OverloadExtensions[3] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&CNDomainIndicator{}, PRESENCE:Presence{Presenceoptional}, }
order_OverloadExtensions[0] = 3
table_OverloadExtensions[96] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_OverloadExtensions[1] = 96
   }

type ErrorIndicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, {'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'ignore', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'optional'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'optional'}, None]}
   Cause  *Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   CNDomainIndicator  *CNDomainIndicator
   GlobalRNCID  *GlobalRNCID
   list []interface{}
}
func (self *ErrorIndicationIEs)createOT() interface{}{
    return nil
}
var table_ErrorIndicationIEs = make(map[int]*RANAPPROTOCOLIES)

var order_ErrorIndicationIEs = make([]int, 4)

func (self *ErrorIndicationIEs) GetIECount() int{
   count := 0
   if self.Cause != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   if self.CNDomainIndicator != nil { count += 1 }
   if self.GlobalRNCID != nil { count += 1 }
   return count//ObjSet
}
func (self *ErrorIndicationIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        if self.Cause != nil { return true }
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
      case 3: //CNDomainIndicator
        if self.CNDomainIndicator != nil { return true }
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil { return true }
   }
   return false//ObjSet
}
func (self *ErrorIndicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause = &Cause{}
        self.Cause.Unpack(st)
        self.list = append(self.list, self.Cause)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator = &CNDomainIndicator{}
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, self.CNDomainIndicator)
      case 86: //GlobalRNCID
        self.GlobalRNCID = &GlobalRNCID{}
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, self.GlobalRNCID)
   }
}
func (self *ErrorIndicationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        if self.Cause != nil {self.Cause.Pack(st)}
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      case 3: //CNDomainIndicator
        if self.CNDomainIndicator != nil {self.CNDomainIndicator.Pack(st)}
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil {self.GlobalRNCID.Pack(st)}
      default:
      break
   }
}
func init() {
table_ErrorIndicationIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[0] = 4
table_ErrorIndicationIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[1] = 9
table_ErrorIndicationIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[2] = 3
table_ErrorIndicationIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationIEs[3] = 86
   }

type ErrorIndicationExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'EXTENSION': 'GlobalCN-ID', 'PRESENCE': 'optional'}, None]}
   GlobalCNID  *GlobalCNID
   list []interface{}
}
func (self *ErrorIndicationExtensions)createOT() interface{}{
    return nil
}
var table_ErrorIndicationExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_ErrorIndicationExtensions = make([]int, 1)

func (self *ErrorIndicationExtensions) GetIECount() int{
   count := 0
   if self.GlobalCNID != nil { count += 1 }
   return count//ObjSet
}
func (self *ErrorIndicationExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
   }
   return false//ObjSet
}
func (self *ErrorIndicationExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
   }
}
func (self *ErrorIndicationExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      default:
      break
   }
}
func init() {
table_ErrorIndicationExtensions[96] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_ErrorIndicationExtensions[0] = 96
   }

type SRNSDataForwardCommandIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-DataForwardingList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-DataForwardingList', 'PRESENCE': 'optional'}, None]}
   RABDataForwardingList  *RABDataForwardingList
   list []interface{}
}
func (self *SRNSDataForwardCommandIEs)createOT() interface{}{
    return nil
}
var table_SRNSDataForwardCommandIEs = make(map[int]*RANAPPROTOCOLIES)

var order_SRNSDataForwardCommandIEs = make([]int, 1)

func (self *SRNSDataForwardCommandIEs) GetIECount() int{
   count := 0
   if self.RABDataForwardingList != nil { count += 1 }
   return count//ObjSet
}
func (self *SRNSDataForwardCommandIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 28: //RABDataForwardingList
        if self.RABDataForwardingList != nil { return true }
   }
   return false//ObjSet
}
func (self *SRNSDataForwardCommandIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 28: //RABDataForwardingList
        self.RABDataForwardingList = &RABDataForwardingList{}
        self.RABDataForwardingList.Unpack(st)
        self.list = append(self.list, self.RABDataForwardingList)
   }
}
func (self *SRNSDataForwardCommandIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 28: //RABDataForwardingList
        if self.RABDataForwardingList != nil {self.RABDataForwardingList.Pack(st)}
      default:
      break
   }
}
func init() {
table_SRNSDataForwardCommandIEs[28] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABDataForwardingList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABDataForwardingList{}, PRESENCE:Presence{Presenceoptional}, }
order_SRNSDataForwardCommandIEs[0] = 28
   }

type SRNSDataForwardCommandExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SRNSDataForwardCommandExtensions)createOT() interface{}{
    return nil
}
var table_SRNSDataForwardCommandExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SRNSDataForwardCommandExtensions = make([]int, 0)

type ForwardSRNSContextIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ContextList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ContextList', 'PRESENCE': 'mandatory'}, None]}
   RABContextList  RABContextList
   list []interface{}
}
func (self *ForwardSRNSContextIEs)createOT() interface{}{
    return nil
}
var table_ForwardSRNSContextIEs = make(map[int]*RANAPPROTOCOLIES)

var order_ForwardSRNSContextIEs = make([]int, 1)

func (self *ForwardSRNSContextIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABContextList
   return count//ObjSet
}
func (self *ForwardSRNSContextIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 25: //RABContextList
        return true //self.RABContextList
   }
   return false//ObjSet
}
func (self *ForwardSRNSContextIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 25: //RABContextList
        self.RABContextList.Unpack(st)
        self.list = append(self.list, &self.RABContextList)
   }
}
func (self *ForwardSRNSContextIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 25: //RABContextList
        self.RABContextList.Pack(st)
      default:
      break
   }
}
func init() {
table_ForwardSRNSContextIEs[25] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABContextList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABContextList{}, PRESENCE:Presence{Presencemandatory}, }
order_ForwardSRNSContextIEs[0] = 25
   }

type ForwardSRNSContextExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-SourceRNC-PDCP-context-info', 'CRITICALITY': 'ignore', 'EXTENSION': 'RRC-Container', 'PRESENCE': 'optional'}, None]}
   SourceRNCPDCPcontextinfo  *RRCContainer
   list []interface{}
}
func (self *ForwardSRNSContextExtensions)createOT() interface{}{
    return nil
}
var table_ForwardSRNSContextExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_ForwardSRNSContextExtensions = make([]int, 1)

func (self *ForwardSRNSContextExtensions) GetIECount() int{
   count := 0
   if self.SourceRNCPDCPcontextinfo != nil { count += 1 }
   return count//ObjSet
}
func (self *ForwardSRNSContextExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 103: //SourceRNCPDCPcontextinfo
        if self.SourceRNCPDCPcontextinfo != nil { return true }
   }
   return false//ObjSet
}
func (self *ForwardSRNSContextExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 103: //SourceRNCPDCPcontextinfo
        self.SourceRNCPDCPcontextinfo = &RRCContainer{}
        self.SourceRNCPDCPcontextinfo.Unpack(st)
        self.list = append(self.list, self.SourceRNCPDCPcontextinfo)
   }
}
func (self *ForwardSRNSContextExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 103: //SourceRNCPDCPcontextinfo
        if self.SourceRNCPDCPcontextinfo != nil {self.SourceRNCPDCPcontextinfo.Pack(st)}
      default:
      break
   }
}
func init() {
table_ForwardSRNSContextExtensions[103] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idSourceRNCPDCPcontextinfo}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&RRCContainer{}, PRESENCE:Presence{Presenceoptional}, }
order_ForwardSRNSContextExtensions[0] = 103
   }

type RABAssignmentRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-SetupOrModifyList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-SetupOrModifyList', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-ReleaseList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ReleaseList', 'PRESENCE': 'optional'}, None]}
   RABSetupOrModifyList  *RABSetupOrModifyList
   RABReleaseList  *RABReleaseList
   list []interface{}
}
func (self *RABAssignmentRequestIEs)createOT() interface{}{
    return nil
}
var table_RABAssignmentRequestIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABAssignmentRequestIEs = make([]int, 2)

func (self *RABAssignmentRequestIEs) GetIECount() int{
   count := 0
   if self.RABSetupOrModifyList != nil { count += 1 }
   if self.RABReleaseList != nil { count += 1 }
   return count//ObjSet
}
func (self *RABAssignmentRequestIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 54: //RABSetupOrModifyList
        if self.RABSetupOrModifyList != nil { return true }
      case 41: //RABReleaseList
        if self.RABReleaseList != nil { return true }
   }
   return false//ObjSet
}
func (self *RABAssignmentRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 54: //RABSetupOrModifyList
        self.RABSetupOrModifyList = &RABSetupOrModifyList{}
        self.RABSetupOrModifyList.Unpack(st)
        self.list = append(self.list, self.RABSetupOrModifyList)
      case 41: //RABReleaseList
        self.RABReleaseList = &RABReleaseList{}
        self.RABReleaseList.Unpack(st)
        self.list = append(self.list, self.RABReleaseList)
   }
}
func (self *RABAssignmentRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 54: //RABSetupOrModifyList
        if self.RABSetupOrModifyList != nil {self.RABSetupOrModifyList.Pack(st)}
      case 41: //RABReleaseList
        if self.RABReleaseList != nil {self.RABReleaseList.Pack(st)}
      default:
      break
   }
}
func init() {
table_RABAssignmentRequestIEs[54] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABSetupOrModifyList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABSetupOrModifyList{}, PRESENCE:Presence{Presenceoptional}, }
order_RABAssignmentRequestIEs[0] = 54
table_RABAssignmentRequestIEs[41] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABReleaseList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABReleaseList{}, PRESENCE:Presence{Presenceoptional}, }
order_RABAssignmentRequestIEs[1] = 41
   }

type RABSetupOrModifyItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'FIRST CRITICALITY': {'type': 'Criticality', 'name': '&firstCriticality'}, 'FIRST TYPE': {'type': 'OpenType', 'name': '&FirstValue'}, 'SECOND CRITICALITY': {'type': 'Criticality', 'name': '&secondCriticality'}, 'SECOND TYPE': {'type': 'OpenType', 'name': '&SecondValue'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES-PAIR', 'members': [{'ID': 'id-RAB-SetupOrModifyItem', 'FIRST CRITICALITY': 'reject', 'FIRST TYPE': 'RAB-SetupOrModifyItemFirst', 'SECOND CRITICALITY': 'ignore', 'SECOND TYPE': 'RAB-SetupOrModifyItemSecond', 'PRESENCE': 'mandatory'}, None]}
   RABSetupOrModifyItemFirst
   RABSetupOrModifyItemSecond
   list []interface{}
}
func (self *RABSetupOrModifyItemIEs)createOT() interface{}{
    return nil
}
var table_RABSetupOrModifyItemIEs = make(map[int]*RANAPPROTOCOLIESPAIR)

var order_RABSetupOrModifyItemIEs = make([]int, 1)

func (self *RABSetupOrModifyItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABSetupOrModifyItem
   return count//ObjSet
}
func (self *RABSetupOrModifyItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESPAIRid).Value
   switch cat {
      case 53: //RABSetupOrModifyItem
        return true //self.RABSetupOrModifyItem
   }
   return false//ObjSet
}
func (self *RABSetupOrModifyItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESPAIRid).Value
   switch cat {
      case 53: //RABSetupOrModifyItemFirst
        self.RABSetupOrModifyItemFirst.Unpack(st)
        self.list = append(self.list, &self.RABSetupOrModifyItemFirst)
   }
}
func (self *RABSetupOrModifyItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESPAIRid).Value
   switch cat {
      case 53: //RABSetupOrModifyItemFirst
        self.RABSetupOrModifyItemFirst.Pack(st)
      default:
      break
   }
}
func init() {
table_RABSetupOrModifyItemIEs[53] = &RANAPPROTOCOLIESPAIR{ID:ProtocolIEID{idRABSetupOrModifyItem}, FIRSTCRITICALITY:Criticality{Criticalityreject}, FIRSTTYPE:&RABSetupOrModifyItemFirst{}, SECONDCRITICALITY:Criticality{Criticalityignore}, SECONDTYPE:&RABSetupOrModifyItemSecond{}, PRESENCE:Presence{Presencemandatory}, }
order_RABSetupOrModifyItemIEs[0] = 53
   }

type TransportLayerInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TransportLayerInformationExtIEs)createOT() interface{}{
    return nil
}
var table_TransportLayerInformationExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_TransportLayerInformationExtIEs = make([]int, 0)

type RABSetupOrModifyItemFirstExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABSetupOrModifyItemFirstExtIEs)createOT() interface{}{
    return nil
}
var table_RABSetupOrModifyItemFirstExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABSetupOrModifyItemFirstExtIEs = make([]int, 0)

type RABSetupOrModifyItemSecondExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-Alt-RAB-Parameters', 'CRITICALITY': 'ignore', 'EXTENSION': 'Alt-RAB-Parameters', 'PRESENCE': 'optional'}, {'ID': 'id-GERAN-BSC-Container', 'CRITICALITY': 'ignore', 'EXTENSION': 'GERAN-BSC-Container', 'PRESENCE': 'optional'}, None]}
   AltRABParameters  *AltRABParameters
   GERANBSCContainer  *GERANBSCContainer
   list []interface{}
}
func (self *RABSetupOrModifyItemSecondExtIEs)createOT() interface{}{
    return nil
}
var table_RABSetupOrModifyItemSecondExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABSetupOrModifyItemSecondExtIEs = make([]int, 2)

func (self *RABSetupOrModifyItemSecondExtIEs) GetIECount() int{
   count := 0
   if self.AltRABParameters != nil { count += 1 }
   if self.GERANBSCContainer != nil { count += 1 }
   return count//ObjSet
}
func (self *RABSetupOrModifyItemSecondExtIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 89: //AltRABParameters
        if self.AltRABParameters != nil { return true }
      case 107: //GERANBSCContainer
        if self.GERANBSCContainer != nil { return true }
   }
   return false//ObjSet
}
func (self *RABSetupOrModifyItemSecondExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 89: //AltRABParameters
        self.AltRABParameters = &AltRABParameters{}
        self.AltRABParameters.Unpack(st)
        self.list = append(self.list, self.AltRABParameters)
      case 107: //GERANBSCContainer
        self.GERANBSCContainer = &GERANBSCContainer{}
        self.GERANBSCContainer.Unpack(st)
        self.list = append(self.list, self.GERANBSCContainer)
   }
}
func (self *RABSetupOrModifyItemSecondExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 89: //AltRABParameters
        if self.AltRABParameters != nil {self.AltRABParameters.Pack(st)}
      case 107: //GERANBSCContainer
        if self.GERANBSCContainer != nil {self.GERANBSCContainer.Pack(st)}
      default:
      break
   }
}
func init() {
table_RABSetupOrModifyItemSecondExtIEs[89] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idAltRABParameters}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&AltRABParameters{}, PRESENCE:Presence{Presenceoptional}, }
order_RABSetupOrModifyItemSecondExtIEs[0] = 89
table_RABSetupOrModifyItemSecondExtIEs[107] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGERANBSCContainer}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GERANBSCContainer{}, PRESENCE:Presence{Presenceoptional}, }
order_RABSetupOrModifyItemSecondExtIEs[1] = 107
   }

type RABAssignmentRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABAssignmentRequestExtensions)createOT() interface{}{
    return nil
}
var table_RABAssignmentRequestExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABAssignmentRequestExtensions = make([]int, 0)

type RABAssignmentResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-SetupOrModifiedList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-SetupOrModifiedList', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-ReleasedList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ReleasedList', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-QueuedList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-QueuedList', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-FailedList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-FailedList', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-ReleaseFailedList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ReleaseFailedList', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   RABSetupOrModifiedList  *RABSetupOrModifiedList
   RABReleasedList  *RABReleasedList
   RABQueuedList  *RABQueuedList
   RABFailedList  *RABFailedList
   RABReleaseFailedList  *RABReleaseFailedList
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *RABAssignmentResponseIEs)createOT() interface{}{
    return nil
}
var table_RABAssignmentResponseIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABAssignmentResponseIEs = make([]int, 6)

func (self *RABAssignmentResponseIEs) GetIECount() int{
   count := 0
   if self.RABSetupOrModifiedList != nil { count += 1 }
   if self.RABReleasedList != nil { count += 1 }
   if self.RABQueuedList != nil { count += 1 }
   if self.RABFailedList != nil { count += 1 }
   if self.RABReleaseFailedList != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *RABAssignmentResponseIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 52: //RABSetupOrModifiedList
        if self.RABSetupOrModifiedList != nil { return true }
      case 43: //RABReleasedList
        if self.RABReleasedList != nil { return true }
      case 38: //RABQueuedList
        if self.RABQueuedList != nil { return true }
      case 35: //RABFailedList
        if self.RABFailedList != nil { return true }
      case 39: //RABReleaseFailedList
        if self.RABReleaseFailedList != nil { return true }
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *RABAssignmentResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 52: //RABSetupOrModifiedList
        self.RABSetupOrModifiedList = &RABSetupOrModifiedList{}
        self.RABSetupOrModifiedList.Unpack(st)
        self.list = append(self.list, self.RABSetupOrModifiedList)
      case 43: //RABReleasedList
        self.RABReleasedList = &RABReleasedList{}
        self.RABReleasedList.Unpack(st)
        self.list = append(self.list, self.RABReleasedList)
      case 38: //RABQueuedList
        self.RABQueuedList = &RABQueuedList{}
        self.RABQueuedList.Unpack(st)
        self.list = append(self.list, self.RABQueuedList)
      case 35: //RABFailedList
        self.RABFailedList = &RABFailedList{}
        self.RABFailedList.Unpack(st)
        self.list = append(self.list, self.RABFailedList)
      case 39: //RABReleaseFailedList
        self.RABReleaseFailedList = &RABReleaseFailedList{}
        self.RABReleaseFailedList.Unpack(st)
        self.list = append(self.list, self.RABReleaseFailedList)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *RABAssignmentResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 52: //RABSetupOrModifiedList
        if self.RABSetupOrModifiedList != nil {self.RABSetupOrModifiedList.Pack(st)}
      case 43: //RABReleasedList
        if self.RABReleasedList != nil {self.RABReleasedList.Pack(st)}
      case 38: //RABQueuedList
        if self.RABQueuedList != nil {self.RABQueuedList.Pack(st)}
      case 35: //RABFailedList
        if self.RABFailedList != nil {self.RABFailedList.Pack(st)}
      case 39: //RABReleaseFailedList
        if self.RABReleaseFailedList != nil {self.RABReleaseFailedList.Pack(st)}
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_RABAssignmentResponseIEs[52] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABSetupOrModifiedList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABSetupOrModifiedList{}, PRESENCE:Presence{Presenceoptional}, }
order_RABAssignmentResponseIEs[0] = 52
table_RABAssignmentResponseIEs[43] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABReleasedList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABReleasedList{}, PRESENCE:Presence{Presenceoptional}, }
order_RABAssignmentResponseIEs[1] = 43
table_RABAssignmentResponseIEs[38] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABQueuedList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABQueuedList{}, PRESENCE:Presence{Presenceoptional}, }
order_RABAssignmentResponseIEs[2] = 38
table_RABAssignmentResponseIEs[35] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABFailedList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABFailedList{}, PRESENCE:Presence{Presenceoptional}, }
order_RABAssignmentResponseIEs[3] = 35
table_RABAssignmentResponseIEs[39] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABReleaseFailedList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABReleaseFailedList{}, PRESENCE:Presence{Presenceoptional}, }
order_RABAssignmentResponseIEs[4] = 39
table_RABAssignmentResponseIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_RABAssignmentResponseIEs[5] = 9
   }

type RABSetupOrModifiedItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-SetupOrModifiedItem', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-SetupOrModifiedItem', 'PRESENCE': 'mandatory'}, None]}
   RABSetupOrModifiedItem  RABSetupOrModifiedItem
   list []interface{}
}
func (self *RABSetupOrModifiedItemIEs)createOT() interface{}{
    return nil
}
var table_RABSetupOrModifiedItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABSetupOrModifiedItemIEs = make([]int, 1)

func (self *RABSetupOrModifiedItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABSetupOrModifiedItem
   return count//ObjSet
}
func (self *RABSetupOrModifiedItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 51: //RABSetupOrModifiedItem
        return true //self.RABSetupOrModifiedItem
   }
   return false//ObjSet
}
func (self *RABSetupOrModifiedItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 51: //RABSetupOrModifiedItem
        self.RABSetupOrModifiedItem.Unpack(st)
        self.list = append(self.list, &self.RABSetupOrModifiedItem)
   }
}
func (self *RABSetupOrModifiedItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 51: //RABSetupOrModifiedItem
        self.RABSetupOrModifiedItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABSetupOrModifiedItemIEs[51] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABSetupOrModifiedItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABSetupOrModifiedItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABSetupOrModifiedItemIEs[0] = 51
   }

type RABSetupOrModifiedItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-Ass-RAB-Parameters', 'CRITICALITY': 'ignore', 'EXTENSION': 'Ass-RAB-Parameters', 'PRESENCE': 'optional'}, None]}
   AssRABParameters  *AssRABParameters
   list []interface{}
}
func (self *RABSetupOrModifiedItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABSetupOrModifiedItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABSetupOrModifiedItemExtIEs = make([]int, 1)

func (self *RABSetupOrModifiedItemExtIEs) GetIECount() int{
   count := 0
   if self.AssRABParameters != nil { count += 1 }
   return count//ObjSet
}
func (self *RABSetupOrModifiedItemExtIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 90: //AssRABParameters
        if self.AssRABParameters != nil { return true }
   }
   return false//ObjSet
}
func (self *RABSetupOrModifiedItemExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 90: //AssRABParameters
        self.AssRABParameters = &AssRABParameters{}
        self.AssRABParameters.Unpack(st)
        self.list = append(self.list, self.AssRABParameters)
   }
}
func (self *RABSetupOrModifiedItemExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 90: //AssRABParameters
        if self.AssRABParameters != nil {self.AssRABParameters.Pack(st)}
      default:
      break
   }
}
func init() {
table_RABSetupOrModifiedItemExtIEs[90] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idAssRABParameters}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&AssRABParameters{}, PRESENCE:Presence{Presenceoptional}, }
order_RABSetupOrModifiedItemExtIEs[0] = 90
   }

type RABReleasedItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ReleasedItem', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ReleasedItem', 'PRESENCE': 'mandatory'}, None]}
   RABReleasedItem  RABReleasedItem
   list []interface{}
}
func (self *RABReleasedItemIEs)createOT() interface{}{
    return nil
}
var table_RABReleasedItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABReleasedItemIEs = make([]int, 1)

func (self *RABReleasedItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABReleasedItem
   return count//ObjSet
}
func (self *RABReleasedItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 42: //RABReleasedItem
        return true //self.RABReleasedItem
   }
   return false//ObjSet
}
func (self *RABReleasedItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 42: //RABReleasedItem
        self.RABReleasedItem.Unpack(st)
        self.list = append(self.list, &self.RABReleasedItem)
   }
}
func (self *RABReleasedItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 42: //RABReleasedItem
        self.RABReleasedItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABReleasedItemIEs[42] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABReleasedItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABReleasedItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABReleasedItemIEs[0] = 42
   }

type RABReleasedItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABReleasedItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABReleasedItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABReleasedItemExtIEs = make([]int, 0)

type DataVolumeListExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DataVolumeListExtIEs)createOT() interface{}{
    return nil
}
var table_DataVolumeListExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_DataVolumeListExtIEs = make([]int, 0)

type RABQueuedItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-QueuedItem', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-QueuedItem', 'PRESENCE': 'mandatory'}, None]}
   RABQueuedItem  RABQueuedItem
   list []interface{}
}
func (self *RABQueuedItemIEs)createOT() interface{}{
    return nil
}
var table_RABQueuedItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABQueuedItemIEs = make([]int, 1)

func (self *RABQueuedItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABQueuedItem
   return count//ObjSet
}
func (self *RABQueuedItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 37: //RABQueuedItem
        return true //self.RABQueuedItem
   }
   return false//ObjSet
}
func (self *RABQueuedItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 37: //RABQueuedItem
        self.RABQueuedItem.Unpack(st)
        self.list = append(self.list, &self.RABQueuedItem)
   }
}
func (self *RABQueuedItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 37: //RABQueuedItem
        self.RABQueuedItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABQueuedItemIEs[37] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABQueuedItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABQueuedItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABQueuedItemIEs[0] = 37
   }

type RABQueuedItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABQueuedItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABQueuedItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABQueuedItemExtIEs = make([]int, 0)

type RABAssignmentResponseExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-GERAN-Iumode-RAB-FailedList-RABAssgntResponse', 'CRITICALITY': 'ignore', 'EXTENSION': 'GERAN-Iumode-RAB-FailedList-RABAssgntResponse', 'PRESENCE': 'optional'}, None]}
   GERANIumodeRABFailedListRABAssgntResponse  *GERANIumodeRABFailedListRABAssgntResponse
   list []interface{}
}
func (self *RABAssignmentResponseExtensions)createOT() interface{}{
    return nil
}
var table_RABAssignmentResponseExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABAssignmentResponseExtensions = make([]int, 1)

func (self *RABAssignmentResponseExtensions) GetIECount() int{
   count := 0
   if self.GERANIumodeRABFailedListRABAssgntResponse != nil { count += 1 }
   return count//ObjSet
}
func (self *RABAssignmentResponseExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 110: //GERANIumodeRABFailedListRABAssgntResponse
        if self.GERANIumodeRABFailedListRABAssgntResponse != nil { return true }
   }
   return false//ObjSet
}
func (self *RABAssignmentResponseExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 110: //GERANIumodeRABFailedListRABAssgntResponse
        self.GERANIumodeRABFailedListRABAssgntResponse = &GERANIumodeRABFailedListRABAssgntResponse{}
        self.GERANIumodeRABFailedListRABAssgntResponse.Unpack(st)
        self.list = append(self.list, self.GERANIumodeRABFailedListRABAssgntResponse)
   }
}
func (self *RABAssignmentResponseExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 110: //GERANIumodeRABFailedListRABAssgntResponse
        if self.GERANIumodeRABFailedListRABAssgntResponse != nil {self.GERANIumodeRABFailedListRABAssgntResponse.Pack(st)}
      default:
      break
   }
}
func init() {
table_RABAssignmentResponseExtensions[110] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idGERANIumodeRABFailedListRABAssgntResponse}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&GERANIumodeRABFailedListRABAssgntResponse{}, PRESENCE:Presence{Presenceoptional}, }
order_RABAssignmentResponseExtensions[0] = 110
   }

type GERANIumodeRABFailedRABAssgntResponseItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-GERAN-Iumode-RAB-Failed-RABAssgntResponse-Item', 'CRITICALITY': 'ignore', 'TYPE': 'GERAN-Iumode-RAB-Failed-RABAssgntResponse-Item', 'PRESENCE': 'mandatory'}, None]}
   GERANIumodeRABFailedRABAssgntResponseItem  GERANIumodeRABFailedRABAssgntResponseItem
   list []interface{}
}
func (self *GERANIumodeRABFailedRABAssgntResponseItemIEs)createOT() interface{}{
    return nil
}
var table_GERANIumodeRABFailedRABAssgntResponseItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_GERANIumodeRABFailedRABAssgntResponseItemIEs = make([]int, 1)

func (self *GERANIumodeRABFailedRABAssgntResponseItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.GERANIumodeRABFailedRABAssgntResponseItem
   return count//ObjSet
}
func (self *GERANIumodeRABFailedRABAssgntResponseItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 109: //GERANIumodeRABFailedRABAssgntResponseItem
        return true //self.GERANIumodeRABFailedRABAssgntResponseItem
   }
   return false//ObjSet
}
func (self *GERANIumodeRABFailedRABAssgntResponseItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 109: //GERANIumodeRABFailedRABAssgntResponseItem
        self.GERANIumodeRABFailedRABAssgntResponseItem.Unpack(st)
        self.list = append(self.list, &self.GERANIumodeRABFailedRABAssgntResponseItem)
   }
}
func (self *GERANIumodeRABFailedRABAssgntResponseItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 109: //GERANIumodeRABFailedRABAssgntResponseItem
        self.GERANIumodeRABFailedRABAssgntResponseItem.Pack(st)
      default:
      break
   }
}
func init() {
table_GERANIumodeRABFailedRABAssgntResponseItemIEs[109] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGERANIumodeRABFailedRABAssgntResponseItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GERANIumodeRABFailedRABAssgntResponseItem{}, PRESENCE:Presence{Presencemandatory}, }
order_GERANIumodeRABFailedRABAssgntResponseItemIEs[0] = 109
   }

type GERANIumodeRABFailedRABAssgntResponseItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GERANIumodeRABFailedRABAssgntResponseItemExtIEs)createOT() interface{}{
    return nil
}
var table_GERANIumodeRABFailedRABAssgntResponseItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_GERANIumodeRABFailedRABAssgntResponseItemExtIEs = make([]int, 0)

type PrivateMessageIEs struct { //ObjSet 1 {'ID': {'type': 'PrivateIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PRIVATE-IES', 'members': [None]}
   list []interface{}
}
func (self *PrivateMessageIEs)createOT() interface{}{
    return nil
}
var table_PrivateMessageIEs = make(map[int]*RANAPPRIVATEIES)

var order_PrivateMessageIEs = make([]int, 0)

type RANAPRelocationInformationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-DirectTransferInformationList-RANAP-RelocInf', 'CRITICALITY': 'ignore', 'TYPE': 'DirectTransferInformationList-RANAP-RelocInf', 'PRESENCE': 'optional'}, {'ID': 'id-RAB-ContextList-RANAP-RelocInf', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ContextList-RANAP-RelocInf', 'PRESENCE': 'optional'}, None]}
   DirectTransferInformationListRANAPRelocInf  *DirectTransferInformationListRANAPRelocInf
   RABContextListRANAPRelocInf  *RABContextListRANAPRelocInf
   list []interface{}
}
func (self *RANAPRelocationInformationIEs)createOT() interface{}{
    return nil
}
var table_RANAPRelocationInformationIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RANAPRelocationInformationIEs = make([]int, 2)

func (self *RANAPRelocationInformationIEs) GetIECount() int{
   count := 0
   if self.DirectTransferInformationListRANAPRelocInf != nil { count += 1 }
   if self.RABContextListRANAPRelocInf != nil { count += 1 }
   return count//ObjSet
}
func (self *RANAPRelocationInformationIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 81: //DirectTransferInformationListRANAPRelocInf
        if self.DirectTransferInformationListRANAPRelocInf != nil { return true }
      case 83: //RABContextListRANAPRelocInf
        if self.RABContextListRANAPRelocInf != nil { return true }
   }
   return false//ObjSet
}
func (self *RANAPRelocationInformationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 81: //DirectTransferInformationListRANAPRelocInf
        self.DirectTransferInformationListRANAPRelocInf = &DirectTransferInformationListRANAPRelocInf{}
        self.DirectTransferInformationListRANAPRelocInf.Unpack(st)
        self.list = append(self.list, self.DirectTransferInformationListRANAPRelocInf)
      case 83: //RABContextListRANAPRelocInf
        self.RABContextListRANAPRelocInf = &RABContextListRANAPRelocInf{}
        self.RABContextListRANAPRelocInf.Unpack(st)
        self.list = append(self.list, self.RABContextListRANAPRelocInf)
   }
}
func (self *RANAPRelocationInformationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 81: //DirectTransferInformationListRANAPRelocInf
        if self.DirectTransferInformationListRANAPRelocInf != nil {self.DirectTransferInformationListRANAPRelocInf.Pack(st)}
      case 83: //RABContextListRANAPRelocInf
        if self.RABContextListRANAPRelocInf != nil {self.RABContextListRANAPRelocInf.Pack(st)}
      default:
      break
   }
}
func init() {
table_RANAPRelocationInformationIEs[81] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idDirectTransferInformationListRANAPRelocInf}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DirectTransferInformationListRANAPRelocInf{}, PRESENCE:Presence{Presenceoptional}, }
order_RANAPRelocationInformationIEs[0] = 81
table_RANAPRelocationInformationIEs[83] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABContextListRANAPRelocInf}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABContextListRANAPRelocInf{}, PRESENCE:Presence{Presenceoptional}, }
order_RANAPRelocationInformationIEs[1] = 83
   }

type DirectTransferInformationItemIEsRANAPRelocInf struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-DirectTransferInformationItem-RANAP-RelocInf', 'CRITICALITY': 'ignore', 'TYPE': 'DirectTransferInformationItem-RANAP-RelocInf', 'PRESENCE': 'mandatory'}, None]}
   DirectTransferInformationItemRANAPRelocInf  DirectTransferInformationItemRANAPRelocInf
   list []interface{}
}
func (self *DirectTransferInformationItemIEsRANAPRelocInf)createOT() interface{}{
    return nil
}
var table_DirectTransferInformationItemIEsRANAPRelocInf = make(map[int]*RANAPPROTOCOLIES)

var order_DirectTransferInformationItemIEsRANAPRelocInf = make([]int, 1)

func (self *DirectTransferInformationItemIEsRANAPRelocInf) GetIECount() int{
   count := 0
   count +=1 //self.DirectTransferInformationItemRANAPRelocInf
   return count//ObjSet
}
func (self *DirectTransferInformationItemIEsRANAPRelocInf) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 80: //DirectTransferInformationItemRANAPRelocInf
        return true //self.DirectTransferInformationItemRANAPRelocInf
   }
   return false//ObjSet
}
func (self *DirectTransferInformationItemIEsRANAPRelocInf)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 80: //DirectTransferInformationItemRANAPRelocInf
        self.DirectTransferInformationItemRANAPRelocInf.Unpack(st)
        self.list = append(self.list, &self.DirectTransferInformationItemRANAPRelocInf)
   }
}
func (self *DirectTransferInformationItemIEsRANAPRelocInf)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 80: //DirectTransferInformationItemRANAPRelocInf
        self.DirectTransferInformationItemRANAPRelocInf.Pack(st)
      default:
      break
   }
}
func init() {
table_DirectTransferInformationItemIEsRANAPRelocInf[80] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idDirectTransferInformationItemRANAPRelocInf}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&DirectTransferInformationItemRANAPRelocInf{}, PRESENCE:Presence{Presencemandatory}, }
order_DirectTransferInformationItemIEsRANAPRelocInf[0] = 80
   }

type RANAPDirectTransferInformationItemExtIEsRANAPRelocInf struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RANAPDirectTransferInformationItemExtIEsRANAPRelocInf)createOT() interface{}{
    return nil
}
var table_RANAPDirectTransferInformationItemExtIEsRANAPRelocInf = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RANAPDirectTransferInformationItemExtIEsRANAPRelocInf = make([]int, 0)

type RABContextItemIEsRANAPRelocInf struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ContextItem-RANAP-RelocInf', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ContextItem-RANAP-RelocInf', 'PRESENCE': 'mandatory'}, None]}
   RABContextItemRANAPRelocInf  RABContextItemRANAPRelocInf
   list []interface{}
}
func (self *RABContextItemIEsRANAPRelocInf)createOT() interface{}{
    return nil
}
var table_RABContextItemIEsRANAPRelocInf = make(map[int]*RANAPPROTOCOLIES)

var order_RABContextItemIEsRANAPRelocInf = make([]int, 1)

func (self *RABContextItemIEsRANAPRelocInf) GetIECount() int{
   count := 0
   count +=1 //self.RABContextItemRANAPRelocInf
   return count//ObjSet
}
func (self *RABContextItemIEsRANAPRelocInf) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 82: //RABContextItemRANAPRelocInf
        return true //self.RABContextItemRANAPRelocInf
   }
   return false//ObjSet
}
func (self *RABContextItemIEsRANAPRelocInf)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 82: //RABContextItemRANAPRelocInf
        self.RABContextItemRANAPRelocInf.Unpack(st)
        self.list = append(self.list, &self.RABContextItemRANAPRelocInf)
   }
}
func (self *RABContextItemIEsRANAPRelocInf)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 82: //RABContextItemRANAPRelocInf
        self.RABContextItemRANAPRelocInf.Pack(st)
      default:
      break
   }
}
func init() {
table_RABContextItemIEsRANAPRelocInf[82] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABContextItemRANAPRelocInf}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABContextItemRANAPRelocInf{}, PRESENCE:Presence{Presencemandatory}, }
order_RABContextItemIEsRANAPRelocInf[0] = 82
   }

type RABContextItemExtIEsRANAPRelocInf struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABContextItemExtIEsRANAPRelocInf)createOT() interface{}{
    return nil
}
var table_RABContextItemExtIEsRANAPRelocInf = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABContextItemExtIEsRANAPRelocInf = make([]int, 0)

type RANAPRelocationInformationExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-SourceRNC-PDCP-context-info', 'CRITICALITY': 'ignore', 'EXTENSION': 'RRC-Container', 'PRESENCE': 'optional'}, None]}
   SourceRNCPDCPcontextinfo  *RRCContainer
   list []interface{}
}
func (self *RANAPRelocationInformationExtensions)createOT() interface{}{
    return nil
}
var table_RANAPRelocationInformationExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RANAPRelocationInformationExtensions = make([]int, 1)

func (self *RANAPRelocationInformationExtensions) GetIECount() int{
   count := 0
   if self.SourceRNCPDCPcontextinfo != nil { count += 1 }
   return count//ObjSet
}
func (self *RANAPRelocationInformationExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 103: //SourceRNCPDCPcontextinfo
        if self.SourceRNCPDCPcontextinfo != nil { return true }
   }
   return false//ObjSet
}
func (self *RANAPRelocationInformationExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 103: //SourceRNCPDCPcontextinfo
        self.SourceRNCPDCPcontextinfo = &RRCContainer{}
        self.SourceRNCPDCPcontextinfo.Unpack(st)
        self.list = append(self.list, self.SourceRNCPDCPcontextinfo)
   }
}
func (self *RANAPRelocationInformationExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 103: //SourceRNCPDCPcontextinfo
        if self.SourceRNCPDCPcontextinfo != nil {self.SourceRNCPDCPcontextinfo.Pack(st)}
      default:
      break
   }
}
func init() {
table_RANAPRelocationInformationExtensions[103] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idSourceRNCPDCPcontextinfo}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&RRCContainer{}, PRESENCE:Presence{Presenceoptional}, }
order_RANAPRelocationInformationExtensions[0] = 103
   }

type RABModifyRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ModifyList', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ModifyList', 'PRESENCE': 'mandatory'}, None]}
   RABModifyList  RABModifyList
   list []interface{}
}
func (self *RABModifyRequestIEs)createOT() interface{}{
    return nil
}
var table_RABModifyRequestIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABModifyRequestIEs = make([]int, 1)

func (self *RABModifyRequestIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABModifyList
   return count//ObjSet
}
func (self *RABModifyRequestIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 91: //RABModifyList
        return true //self.RABModifyList
   }
   return false//ObjSet
}
func (self *RABModifyRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 91: //RABModifyList
        self.RABModifyList.Unpack(st)
        self.list = append(self.list, &self.RABModifyList)
   }
}
func (self *RABModifyRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 91: //RABModifyList
        self.RABModifyList.Pack(st)
      default:
      break
   }
}
func init() {
table_RABModifyRequestIEs[91] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABModifyList}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABModifyList{}, PRESENCE:Presence{Presencemandatory}, }
order_RABModifyRequestIEs[0] = 91
   }

type RABModifyItemIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-RAB-ModifyItem', 'CRITICALITY': 'ignore', 'TYPE': 'RAB-ModifyItem', 'PRESENCE': 'mandatory'}, None]}
   RABModifyItem  RABModifyItem
   list []interface{}
}
func (self *RABModifyItemIEs)createOT() interface{}{
    return nil
}
var table_RABModifyItemIEs = make(map[int]*RANAPPROTOCOLIES)

var order_RABModifyItemIEs = make([]int, 1)

func (self *RABModifyItemIEs) GetIECount() int{
   count := 0
   count +=1 //self.RABModifyItem
   return count//ObjSet
}
func (self *RABModifyItemIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 92: //RABModifyItem
        return true //self.RABModifyItem
   }
   return false//ObjSet
}
func (self *RABModifyItemIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 92: //RABModifyItem
        self.RABModifyItem.Unpack(st)
        self.list = append(self.list, &self.RABModifyItem)
   }
}
func (self *RABModifyItemIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 92: //RABModifyItem
        self.RABModifyItem.Pack(st)
      default:
      break
   }
}
func init() {
table_RABModifyItemIEs[92] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idRABModifyItem}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&RABModifyItem{}, PRESENCE:Presence{Presencemandatory}, }
order_RABModifyItemIEs[0] = 92
   }

type RABModifyItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABModifyItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABModifyItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABModifyItemExtIEs = make([]int, 0)

type RABModifyRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RABModifyRequestExtensions)createOT() interface{}{
    return nil
}
var table_RABModifyRequestExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABModifyRequestExtensions = make([]int, 0)

type LocationRelatedDataRequestIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-LocationRelatedDataRequestType', 'CRITICALITY': 'reject', 'TYPE': 'LocationRelatedDataRequestType', 'PRESENCE': 'optional'}, None]}
   LocationRelatedDataRequestType  *LocationRelatedDataRequestType
   list []interface{}
}
func (self *LocationRelatedDataRequestIEs)createOT() interface{}{
    return nil
}
var table_LocationRelatedDataRequestIEs = make(map[int]*RANAPPROTOCOLIES)

var order_LocationRelatedDataRequestIEs = make([]int, 1)

func (self *LocationRelatedDataRequestIEs) GetIECount() int{
   count := 0
   if self.LocationRelatedDataRequestType != nil { count += 1 }
   return count//ObjSet
}
func (self *LocationRelatedDataRequestIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 95: //LocationRelatedDataRequestType
        if self.LocationRelatedDataRequestType != nil { return true }
   }
   return false//ObjSet
}
func (self *LocationRelatedDataRequestIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 95: //LocationRelatedDataRequestType
        self.LocationRelatedDataRequestType = &LocationRelatedDataRequestType{}
        self.LocationRelatedDataRequestType.Unpack(st)
        self.list = append(self.list, self.LocationRelatedDataRequestType)
   }
}
func (self *LocationRelatedDataRequestIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 95: //LocationRelatedDataRequestType
        if self.LocationRelatedDataRequestType != nil {self.LocationRelatedDataRequestType.Pack(st)}
      default:
      break
   }
}
func init() {
table_LocationRelatedDataRequestIEs[95] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idLocationRelatedDataRequestType}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&LocationRelatedDataRequestType{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationRelatedDataRequestIEs[0] = 95
   }

type LocationRelatedDataRequestExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-LocationRelatedDataRequestTypeSpecificToGERANIuMode', 'CRITICALITY': 'reject', 'EXTENSION': 'LocationRelatedDataRequestTypeSpecificToGERANIuMode', 'PRESENCE': 'optional'}, None]}
   LocationRelatedDataRequestTypeSpecificToGERANIuMode  *LocationRelatedDataRequestTypeSpecificToGERANIuMode
   list []interface{}
}
func (self *LocationRelatedDataRequestExtensions)createOT() interface{}{
    return nil
}
var table_LocationRelatedDataRequestExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_LocationRelatedDataRequestExtensions = make([]int, 1)

func (self *LocationRelatedDataRequestExtensions) GetIECount() int{
   count := 0
   if self.LocationRelatedDataRequestTypeSpecificToGERANIuMode != nil { count += 1 }
   return count//ObjSet
}
func (self *LocationRelatedDataRequestExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 115: //LocationRelatedDataRequestTypeSpecificToGERANIuMode
        if self.LocationRelatedDataRequestTypeSpecificToGERANIuMode != nil { return true }
   }
   return false//ObjSet
}
func (self *LocationRelatedDataRequestExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 115: //LocationRelatedDataRequestTypeSpecificToGERANIuMode
        self.LocationRelatedDataRequestTypeSpecificToGERANIuMode = &LocationRelatedDataRequestTypeSpecificToGERANIuMode{}
        self.LocationRelatedDataRequestTypeSpecificToGERANIuMode.Unpack(st)
        self.list = append(self.list, self.LocationRelatedDataRequestTypeSpecificToGERANIuMode)
   }
}
func (self *LocationRelatedDataRequestExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 115: //LocationRelatedDataRequestTypeSpecificToGERANIuMode
        if self.LocationRelatedDataRequestTypeSpecificToGERANIuMode != nil {self.LocationRelatedDataRequestTypeSpecificToGERANIuMode.Pack(st)}
      default:
      break
   }
}
func init() {
table_LocationRelatedDataRequestExtensions[115] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idLocationRelatedDataRequestTypeSpecificToGERANIuMode}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&LocationRelatedDataRequestTypeSpecificToGERANIuMode{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationRelatedDataRequestExtensions[0] = 115
   }

type LocationRelatedDataResponseIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-BroadcastAssistanceDataDecipheringKeys', 'CRITICALITY': 'ignore', 'TYPE': 'BroadcastAssistanceDataDecipheringKeys', 'PRESENCE': 'optional'}, None]}
   BroadcastAssistanceDataDecipheringKeys  *BroadcastAssistanceDataDecipheringKeys
   list []interface{}
}
func (self *LocationRelatedDataResponseIEs)createOT() interface{}{
    return nil
}
var table_LocationRelatedDataResponseIEs = make(map[int]*RANAPPROTOCOLIES)

var order_LocationRelatedDataResponseIEs = make([]int, 1)

func (self *LocationRelatedDataResponseIEs) GetIECount() int{
   count := 0
   if self.BroadcastAssistanceDataDecipheringKeys != nil { count += 1 }
   return count//ObjSet
}
func (self *LocationRelatedDataResponseIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 94: //BroadcastAssistanceDataDecipheringKeys
        if self.BroadcastAssistanceDataDecipheringKeys != nil { return true }
   }
   return false//ObjSet
}
func (self *LocationRelatedDataResponseIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 94: //BroadcastAssistanceDataDecipheringKeys
        self.BroadcastAssistanceDataDecipheringKeys = &BroadcastAssistanceDataDecipheringKeys{}
        self.BroadcastAssistanceDataDecipheringKeys.Unpack(st)
        self.list = append(self.list, self.BroadcastAssistanceDataDecipheringKeys)
   }
}
func (self *LocationRelatedDataResponseIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 94: //BroadcastAssistanceDataDecipheringKeys
        if self.BroadcastAssistanceDataDecipheringKeys != nil {self.BroadcastAssistanceDataDecipheringKeys.Pack(st)}
      default:
      break
   }
}
func init() {
table_LocationRelatedDataResponseIEs[94] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idBroadcastAssistanceDataDecipheringKeys}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&BroadcastAssistanceDataDecipheringKeys{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationRelatedDataResponseIEs[0] = 94
   }

type LocationRelatedDataResponseExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'EXTENSION': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *LocationRelatedDataResponseExtensions)createOT() interface{}{
    return nil
}
var table_LocationRelatedDataResponseExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_LocationRelatedDataResponseExtensions = make([]int, 1)

func (self *LocationRelatedDataResponseExtensions) GetIECount() int{
   count := 0
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *LocationRelatedDataResponseExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *LocationRelatedDataResponseExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *LocationRelatedDataResponseExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_LocationRelatedDataResponseExtensions[9] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationRelatedDataResponseExtensions[0] = 9
   }

type LocationRelatedDataFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, None]}
   Cause  Cause
   list []interface{}
}
func (self *LocationRelatedDataFailureIEs)createOT() interface{}{
    return nil
}
var table_LocationRelatedDataFailureIEs = make(map[int]*RANAPPROTOCOLIES)

var order_LocationRelatedDataFailureIEs = make([]int, 1)

func (self *LocationRelatedDataFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.Cause
   return count//ObjSet
}
func (self *LocationRelatedDataFailureIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        return true //self.Cause
   }
   return false//ObjSet
}
func (self *LocationRelatedDataFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
   }
}
func (self *LocationRelatedDataFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 4: //Cause
        self.Cause.Pack(st)
      default:
      break
   }
}
func init() {
table_LocationRelatedDataFailureIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_LocationRelatedDataFailureIEs[0] = 4
   }

type LocationRelatedDataFailureExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'EXTENSION': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *LocationRelatedDataFailureExtensions)createOT() interface{}{
    return nil
}
var table_LocationRelatedDataFailureExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_LocationRelatedDataFailureExtensions = make([]int, 1)

func (self *LocationRelatedDataFailureExtensions) GetIECount() int{
   count := 0
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *LocationRelatedDataFailureExtensions) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *LocationRelatedDataFailureExtensions)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *LocationRelatedDataFailureExtensions)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_LocationRelatedDataFailureExtensions[9] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_LocationRelatedDataFailureExtensions[0] = 9
   }

type InformationTransferIndicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-InformationTransferID', 'CRITICALITY': 'reject', 'TYPE': 'InformationTransferID', 'PRESENCE': 'mandatory'}, {'ID': 'id-ProvidedData', 'CRITICALITY': 'reject', 'TYPE': 'ProvidedData', 'PRESENCE': 'mandatory'}, {'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'reject', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalCN-ID', 'PRESENCE': 'optional'}, None]}
   InformationTransferID  InformationTransferID
   ProvidedData  ProvidedData
   CNDomainIndicator  CNDomainIndicator
   GlobalCNID  *GlobalCNID
   list []interface{}
}
func (self *InformationTransferIndicationIEs)createOT() interface{}{
    return nil
}
var table_InformationTransferIndicationIEs = make(map[int]*RANAPPROTOCOLIES)

var order_InformationTransferIndicationIEs = make([]int, 4)

func (self *InformationTransferIndicationIEs) GetIECount() int{
   count := 0
   count +=1 //self.InformationTransferID
   count +=1 //self.ProvidedData
   count +=1 //self.CNDomainIndicator
   if self.GlobalCNID != nil { count += 1 }
   return count//ObjSet
}
func (self *InformationTransferIndicationIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        return true //self.InformationTransferID
      case 106: //ProvidedData
        return true //self.ProvidedData
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
   }
   return false//ObjSet
}
func (self *InformationTransferIndicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Unpack(st)
        self.list = append(self.list, &self.InformationTransferID)
      case 106: //ProvidedData
        self.ProvidedData.Unpack(st)
        self.list = append(self.list, &self.ProvidedData)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
   }
}
func (self *InformationTransferIndicationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Pack(st)
      case 106: //ProvidedData
        self.ProvidedData.Pack(st)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      default:
      break
   }
}
func init() {
table_InformationTransferIndicationIEs[104] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idInformationTransferID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&InformationTransferID{}, PRESENCE:Presence{Presencemandatory}, }
order_InformationTransferIndicationIEs[0] = 104
table_InformationTransferIndicationIEs[106] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idProvidedData}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&ProvidedData{}, PRESENCE:Presence{Presencemandatory}, }
order_InformationTransferIndicationIEs[1] = 106
table_InformationTransferIndicationIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_InformationTransferIndicationIEs[2] = 3
table_InformationTransferIndicationIEs[96] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_InformationTransferIndicationIEs[3] = 96
   }

type InformationTransferIndicationExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *InformationTransferIndicationExtensions)createOT() interface{}{
    return nil
}
var table_InformationTransferIndicationExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_InformationTransferIndicationExtensions = make([]int, 0)

type InformationTransferConfirmationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-InformationTransferID', 'CRITICALITY': 'ignore', 'TYPE': 'InformationTransferID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'ignore', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'mandatory'}, None]}
   InformationTransferID  InformationTransferID
   CNDomainIndicator  CNDomainIndicator
   CriticalityDiagnostics  *CriticalityDiagnostics
   GlobalRNCID  GlobalRNCID
   list []interface{}
}
func (self *InformationTransferConfirmationIEs)createOT() interface{}{
    return nil
}
var table_InformationTransferConfirmationIEs = make(map[int]*RANAPPROTOCOLIES)

var order_InformationTransferConfirmationIEs = make([]int, 4)

func (self *InformationTransferConfirmationIEs) GetIECount() int{
   count := 0
   count +=1 //self.InformationTransferID
   count +=1 //self.CNDomainIndicator
   if self.CriticalityDiagnostics != nil { count += 1 }
   count +=1 //self.GlobalRNCID
   return count//ObjSet
}
func (self *InformationTransferConfirmationIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        return true //self.InformationTransferID
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
      case 86: //GlobalRNCID
        return true //self.GlobalRNCID
   }
   return false//ObjSet
}
func (self *InformationTransferConfirmationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Unpack(st)
        self.list = append(self.list, &self.InformationTransferID)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
      case 86: //GlobalRNCID
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, &self.GlobalRNCID)
   }
}
func (self *InformationTransferConfirmationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Pack(st)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      case 86: //GlobalRNCID
        self.GlobalRNCID.Pack(st)
      default:
      break
   }
}
func init() {
table_InformationTransferConfirmationIEs[104] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idInformationTransferID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&InformationTransferID{}, PRESENCE:Presence{Presencemandatory}, }
order_InformationTransferConfirmationIEs[0] = 104
table_InformationTransferConfirmationIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_InformationTransferConfirmationIEs[1] = 3
table_InformationTransferConfirmationIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_InformationTransferConfirmationIEs[2] = 9
table_InformationTransferConfirmationIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presencemandatory}, }
order_InformationTransferConfirmationIEs[3] = 86
   }

type InformationTransferConfirmationExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *InformationTransferConfirmationExtensions)createOT() interface{}{
    return nil
}
var table_InformationTransferConfirmationExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_InformationTransferConfirmationExtensions = make([]int, 0)

type InformationTransferFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-InformationTransferID', 'CRITICALITY': 'ignore', 'TYPE': 'InformationTransferID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'ignore', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'mandatory'}, None]}
   InformationTransferID  InformationTransferID
   CNDomainIndicator  CNDomainIndicator
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   GlobalRNCID  GlobalRNCID
   list []interface{}
}
func (self *InformationTransferFailureIEs)createOT() interface{}{
    return nil
}
var table_InformationTransferFailureIEs = make(map[int]*RANAPPROTOCOLIES)

var order_InformationTransferFailureIEs = make([]int, 5)

func (self *InformationTransferFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.InformationTransferID
   count +=1 //self.CNDomainIndicator
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   count +=1 //self.GlobalRNCID
   return count//ObjSet
}
func (self *InformationTransferFailureIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        return true //self.InformationTransferID
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 4: //Cause
        return true //self.Cause
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
      case 86: //GlobalRNCID
        return true //self.GlobalRNCID
   }
   return false//ObjSet
}
func (self *InformationTransferFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Unpack(st)
        self.list = append(self.list, &self.InformationTransferID)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
      case 86: //GlobalRNCID
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, &self.GlobalRNCID)
   }
}
func (self *InformationTransferFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Pack(st)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 4: //Cause
        self.Cause.Pack(st)
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      case 86: //GlobalRNCID
        self.GlobalRNCID.Pack(st)
      default:
      break
   }
}
func init() {
table_InformationTransferFailureIEs[104] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idInformationTransferID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&InformationTransferID{}, PRESENCE:Presence{Presencemandatory}, }
order_InformationTransferFailureIEs[0] = 104
table_InformationTransferFailureIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_InformationTransferFailureIEs[1] = 3
table_InformationTransferFailureIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_InformationTransferFailureIEs[2] = 4
table_InformationTransferFailureIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_InformationTransferFailureIEs[3] = 9
table_InformationTransferFailureIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presencemandatory}, }
order_InformationTransferFailureIEs[4] = 86
   }

type InformationTransferFailureExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *InformationTransferFailureExtensions)createOT() interface{}{
    return nil
}
var table_InformationTransferFailureExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_InformationTransferFailureExtensions = make([]int, 0)

type UESpecificInformationIndicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-UESBI-Iu', 'CRITICALITY': 'ignore', 'TYPE': 'UESBI-Iu', 'PRESENCE': 'optional'}, None]}
   UESBIIu  *UESBIIu
   list []interface{}
}
func (self *UESpecificInformationIndicationIEs)createOT() interface{}{
    return nil
}
var table_UESpecificInformationIndicationIEs = make(map[int]*RANAPPROTOCOLIES)

var order_UESpecificInformationIndicationIEs = make([]int, 1)

func (self *UESpecificInformationIndicationIEs) GetIECount() int{
   count := 0
   if self.UESBIIu != nil { count += 1 }
   return count//ObjSet
}
func (self *UESpecificInformationIndicationIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 118: //UESBIIu
        if self.UESBIIu != nil { return true }
   }
   return false//ObjSet
}
func (self *UESpecificInformationIndicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 118: //UESBIIu
        self.UESBIIu = &UESBIIu{}
        self.UESBIIu.Unpack(st)
        self.list = append(self.list, self.UESBIIu)
   }
}
func (self *UESpecificInformationIndicationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 118: //UESBIIu
        if self.UESBIIu != nil {self.UESBIIu.Pack(st)}
      default:
      break
   }
}
func init() {
table_UESpecificInformationIndicationIEs[118] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idUESBIIu}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&UESBIIu{}, PRESENCE:Presence{Presenceoptional}, }
order_UESpecificInformationIndicationIEs[0] = 118
   }

type UESpecificInformationIndicationExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UESpecificInformationIndicationExtensions)createOT() interface{}{
    return nil
}
var table_UESpecificInformationIndicationExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_UESpecificInformationIndicationExtensions = make([]int, 0)

type DirectInformationTransferIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-InterSystemInformationTransferType', 'CRITICALITY': 'ignore', 'TYPE': 'InformationTransferType', 'PRESENCE': 'optional'}, {'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'ignore', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'optional'}, {'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalCN-ID', 'PRESENCE': 'optional'}, None]}
   InterSystemInformationTransferType  *InformationTransferType
   CNDomainIndicator  CNDomainIndicator
   GlobalRNCID  *GlobalRNCID
   GlobalCNID  *GlobalCNID
   list []interface{}
}
func (self *DirectInformationTransferIEs)createOT() interface{}{
    return nil
}
var table_DirectInformationTransferIEs = make(map[int]*RANAPPROTOCOLIES)

var order_DirectInformationTransferIEs = make([]int, 4)

func (self *DirectInformationTransferIEs) GetIECount() int{
   count := 0
   if self.InterSystemInformationTransferType != nil { count += 1 }
   count +=1 //self.CNDomainIndicator
   if self.GlobalRNCID != nil { count += 1 }
   if self.GlobalCNID != nil { count += 1 }
   return count//ObjSet
}
func (self *DirectInformationTransferIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 126: //InterSystemInformationTransferType
        if self.InterSystemInformationTransferType != nil { return true }
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil { return true }
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
   }
   return false//ObjSet
}
func (self *DirectInformationTransferIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 126: //InterSystemInformationTransferType
        self.InterSystemInformationTransferType = &InformationTransferType{}
        self.InterSystemInformationTransferType.Unpack(st)
        self.list = append(self.list, self.InterSystemInformationTransferType)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 86: //GlobalRNCID
        self.GlobalRNCID = &GlobalRNCID{}
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, self.GlobalRNCID)
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
   }
}
func (self *DirectInformationTransferIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 126: //InterSystemInformationTransferType
        if self.InterSystemInformationTransferType != nil {self.InterSystemInformationTransferType.Pack(st)}
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 86: //GlobalRNCID
        if self.GlobalRNCID != nil {self.GlobalRNCID.Pack(st)}
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      default:
      break
   }
}
func init() {
table_DirectInformationTransferIEs[126] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idInterSystemInformationTransferType}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&InformationTransferType{}, PRESENCE:Presence{Presenceoptional}, }
order_DirectInformationTransferIEs[0] = 126
table_DirectInformationTransferIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_DirectInformationTransferIEs[1] = 3
table_DirectInformationTransferIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presenceoptional}, }
order_DirectInformationTransferIEs[2] = 86
table_DirectInformationTransferIEs[96] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_DirectInformationTransferIEs[3] = 96
   }

type DirectInformationTransferExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *DirectInformationTransferExtensions)createOT() interface{}{
    return nil
}
var table_DirectInformationTransferExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_DirectInformationTransferExtensions = make([]int, 0)

type UplinkInformationTransferIndicationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-InformationTransferID', 'CRITICALITY': 'reject', 'TYPE': 'InformationTransferID', 'PRESENCE': 'mandatory'}, {'ID': 'id-InformationTransferType', 'CRITICALITY': 'reject', 'TYPE': 'InformationTransferType', 'PRESENCE': 'mandatory'}, {'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'reject', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalRNC-ID', 'CRITICALITY': 'reject', 'TYPE': 'GlobalRNC-ID', 'PRESENCE': 'mandatory'}, None]}
   InformationTransferID  InformationTransferID
   InformationTransferType  InformationTransferType
   CNDomainIndicator  CNDomainIndicator
   GlobalRNCID  GlobalRNCID
   list []interface{}
}
func (self *UplinkInformationTransferIndicationIEs)createOT() interface{}{
    return nil
}
var table_UplinkInformationTransferIndicationIEs = make(map[int]*RANAPPROTOCOLIES)

var order_UplinkInformationTransferIndicationIEs = make([]int, 4)

func (self *UplinkInformationTransferIndicationIEs) GetIECount() int{
   count := 0
   count +=1 //self.InformationTransferID
   count +=1 //self.InformationTransferType
   count +=1 //self.CNDomainIndicator
   count +=1 //self.GlobalRNCID
   return count//ObjSet
}
func (self *UplinkInformationTransferIndicationIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        return true //self.InformationTransferID
      case 123: //InformationTransferType
        return true //self.InformationTransferType
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 86: //GlobalRNCID
        return true //self.GlobalRNCID
   }
   return false//ObjSet
}
func (self *UplinkInformationTransferIndicationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Unpack(st)
        self.list = append(self.list, &self.InformationTransferID)
      case 123: //InformationTransferType
        self.InformationTransferType.Unpack(st)
        self.list = append(self.list, &self.InformationTransferType)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 86: //GlobalRNCID
        self.GlobalRNCID.Unpack(st)
        self.list = append(self.list, &self.GlobalRNCID)
   }
}
func (self *UplinkInformationTransferIndicationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Pack(st)
      case 123: //InformationTransferType
        self.InformationTransferType.Pack(st)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 86: //GlobalRNCID
        self.GlobalRNCID.Pack(st)
      default:
      break
   }
}
func init() {
table_UplinkInformationTransferIndicationIEs[104] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idInformationTransferID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&InformationTransferID{}, PRESENCE:Presence{Presencemandatory}, }
order_UplinkInformationTransferIndicationIEs[0] = 104
table_UplinkInformationTransferIndicationIEs[123] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idInformationTransferType}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&InformationTransferType{}, PRESENCE:Presence{Presencemandatory}, }
order_UplinkInformationTransferIndicationIEs[1] = 123
table_UplinkInformationTransferIndicationIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_UplinkInformationTransferIndicationIEs[2] = 3
table_UplinkInformationTransferIndicationIEs[86] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalRNCID}, CRITICALITY:Criticality{Criticalityreject}, TYPE:&GlobalRNCID{}, PRESENCE:Presence{Presencemandatory}, }
order_UplinkInformationTransferIndicationIEs[3] = 86
   }

type UplinkInformationTransferIndicationExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UplinkInformationTransferIndicationExtensions)createOT() interface{}{
    return nil
}
var table_UplinkInformationTransferIndicationExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_UplinkInformationTransferIndicationExtensions = make([]int, 0)

type UplinkInformationTransferConfirmationIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-InformationTransferID', 'CRITICALITY': 'ignore', 'TYPE': 'InformationTransferID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'ignore', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalCN-ID', 'PRESENCE': 'optional'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   InformationTransferID  InformationTransferID
   CNDomainIndicator  CNDomainIndicator
   GlobalCNID  *GlobalCNID
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *UplinkInformationTransferConfirmationIEs)createOT() interface{}{
    return nil
}
var table_UplinkInformationTransferConfirmationIEs = make(map[int]*RANAPPROTOCOLIES)

var order_UplinkInformationTransferConfirmationIEs = make([]int, 4)

func (self *UplinkInformationTransferConfirmationIEs) GetIECount() int{
   count := 0
   count +=1 //self.InformationTransferID
   count +=1 //self.CNDomainIndicator
   if self.GlobalCNID != nil { count += 1 }
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *UplinkInformationTransferConfirmationIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        return true //self.InformationTransferID
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *UplinkInformationTransferConfirmationIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Unpack(st)
        self.list = append(self.list, &self.InformationTransferID)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *UplinkInformationTransferConfirmationIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Pack(st)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_UplinkInformationTransferConfirmationIEs[104] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idInformationTransferID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&InformationTransferID{}, PRESENCE:Presence{Presencemandatory}, }
order_UplinkInformationTransferConfirmationIEs[0] = 104
table_UplinkInformationTransferConfirmationIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_UplinkInformationTransferConfirmationIEs[1] = 3
table_UplinkInformationTransferConfirmationIEs[96] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_UplinkInformationTransferConfirmationIEs[2] = 96
table_UplinkInformationTransferConfirmationIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_UplinkInformationTransferConfirmationIEs[3] = 9
   }

type UplinkInformationTransferConfirmationExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UplinkInformationTransferConfirmationExtensions)createOT() interface{}{
    return nil
}
var table_UplinkInformationTransferConfirmationExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_UplinkInformationTransferConfirmationExtensions = make([]int, 0)

type UplinkInformationTransferFailureIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolIE-ID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'TYPE': {'type': 'OpenType', 'name': '&Value'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-IES', 'members': [{'ID': 'id-InformationTransferID', 'CRITICALITY': 'ignore', 'TYPE': 'InformationTransferID', 'PRESENCE': 'mandatory'}, {'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'ignore', 'TYPE': 'CN-DomainIndicator', 'PRESENCE': 'mandatory'}, {'ID': 'id-GlobalCN-ID', 'CRITICALITY': 'ignore', 'TYPE': 'GlobalCN-ID', 'PRESENCE': 'optional'}, {'ID': 'id-Cause', 'CRITICALITY': 'ignore', 'TYPE': 'Cause', 'PRESENCE': 'mandatory'}, {'ID': 'id-CriticalityDiagnostics', 'CRITICALITY': 'ignore', 'TYPE': 'CriticalityDiagnostics', 'PRESENCE': 'optional'}, None]}
   InformationTransferID  InformationTransferID
   CNDomainIndicator  CNDomainIndicator
   GlobalCNID  *GlobalCNID
   Cause  Cause
   CriticalityDiagnostics  *CriticalityDiagnostics
   list []interface{}
}
func (self *UplinkInformationTransferFailureIEs)createOT() interface{}{
    return nil
}
var table_UplinkInformationTransferFailureIEs = make(map[int]*RANAPPROTOCOLIES)

var order_UplinkInformationTransferFailureIEs = make([]int, 5)

func (self *UplinkInformationTransferFailureIEs) GetIECount() int{
   count := 0
   count +=1 //self.InformationTransferID
   count +=1 //self.CNDomainIndicator
   if self.GlobalCNID != nil { count += 1 }
   count +=1 //self.Cause
   if self.CriticalityDiagnostics != nil { count += 1 }
   return count//ObjSet
}
func (self *UplinkInformationTransferFailureIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        return true //self.InformationTransferID
      case 3: //CNDomainIndicator
        return true //self.CNDomainIndicator
      case 96: //GlobalCNID
        if self.GlobalCNID != nil { return true }
      case 4: //Cause
        return true //self.Cause
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil { return true }
   }
   return false//ObjSet
}
func (self *UplinkInformationTransferFailureIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Unpack(st)
        self.list = append(self.list, &self.InformationTransferID)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, &self.CNDomainIndicator)
      case 96: //GlobalCNID
        self.GlobalCNID = &GlobalCNID{}
        self.GlobalCNID.Unpack(st)
        self.list = append(self.list, self.GlobalCNID)
      case 4: //Cause
        self.Cause.Unpack(st)
        self.list = append(self.list, &self.Cause)
      case 9: //CriticalityDiagnostics
        self.CriticalityDiagnostics = &CriticalityDiagnostics{}
        self.CriticalityDiagnostics.Unpack(st)
        self.list = append(self.list, self.CriticalityDiagnostics)
   }
}
func (self *UplinkInformationTransferFailureIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLIESid).Value
   switch cat {
      case 104: //InformationTransferID
        self.InformationTransferID.Pack(st)
      case 3: //CNDomainIndicator
        self.CNDomainIndicator.Pack(st)
      case 96: //GlobalCNID
        if self.GlobalCNID != nil {self.GlobalCNID.Pack(st)}
      case 4: //Cause
        self.Cause.Pack(st)
      case 9: //CriticalityDiagnostics
        if self.CriticalityDiagnostics != nil {self.CriticalityDiagnostics.Pack(st)}
      default:
      break
   }
}
func init() {
table_UplinkInformationTransferFailureIEs[104] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idInformationTransferID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&InformationTransferID{}, PRESENCE:Presence{Presencemandatory}, }
order_UplinkInformationTransferFailureIEs[0] = 104
table_UplinkInformationTransferFailureIEs[3] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CNDomainIndicator{}, PRESENCE:Presence{Presencemandatory}, }
order_UplinkInformationTransferFailureIEs[1] = 3
table_UplinkInformationTransferFailureIEs[96] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idGlobalCNID}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&GlobalCNID{}, PRESENCE:Presence{Presenceoptional}, }
order_UplinkInformationTransferFailureIEs[2] = 96
table_UplinkInformationTransferFailureIEs[4] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCause}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&Cause{}, PRESENCE:Presence{Presencemandatory}, }
order_UplinkInformationTransferFailureIEs[3] = 4
table_UplinkInformationTransferFailureIEs[9] = &RANAPPROTOCOLIES{ID:ProtocolIEID{idCriticalityDiagnostics}, CRITICALITY:Criticality{Criticalityignore}, TYPE:&CriticalityDiagnostics{}, PRESENCE:Presence{Presenceoptional}, }
order_UplinkInformationTransferFailureIEs[4] = 9
   }

type UplinkInformationTransferFailureExtensions struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UplinkInformationTransferFailureExtensions)createOT() interface{}{
    return nil
}
var table_UplinkInformationTransferFailureExtensions = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_UplinkInformationTransferFailureExtensions = make([]int, 0)

type AllocationOrRetentionPriorityExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *AllocationOrRetentionPriorityExtIEs)createOT() interface{}{
    return nil
}
var table_AllocationOrRetentionPriorityExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_AllocationOrRetentionPriorityExtIEs = make([]int, 0)

type AltRABParametersExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *AltRABParametersExtIEs)createOT() interface{}{
    return nil
}
var table_AltRABParametersExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_AltRABParametersExtIEs = make([]int, 0)

type AssRABParametersExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *AssRABParametersExtIEs)createOT() interface{}{
    return nil
}
var table_AssRABParametersExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_AssRABParametersExtIEs = make([]int, 0)

type AuthorisedPLMNsExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *AuthorisedPLMNsExtIEs)createOT() interface{}{
    return nil
}
var table_AuthorisedPLMNsExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_AuthorisedPLMNsExtIEs = make([]int, 0)

type CellLoadInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CellLoadInformationExtIEs)createOT() interface{}{
    return nil
}
var table_CellLoadInformationExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_CellLoadInformationExtIEs = make([]int, 0)

type CellLoadInformationGroupExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CellLoadInformationGroupExtIEs)createOT() interface{}{
    return nil
}
var table_CellLoadInformationGroupExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_CellLoadInformationGroupExtIEs = make([]int, 0)

type CriticalityDiagnosticsExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CriticalityDiagnosticsExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsExtIEs = make([]int, 0)

type CriticalityDiagnosticsIEListExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-MessageStructure', 'CRITICALITY': 'ignore', 'EXTENSION': 'MessageStructure', 'PRESENCE': 'optional'}, {'ID': 'id-TypeOfError', 'CRITICALITY': 'ignore', 'EXTENSION': 'TypeOfError', 'PRESENCE': 'mandatory'}, None]}
   MessageStructure  *MessageStructure
   TypeOfError  TypeOfError
   list []interface{}
}
func (self *CriticalityDiagnosticsIEListExtIEs)createOT() interface{}{
    return nil
}
var table_CriticalityDiagnosticsIEListExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_CriticalityDiagnosticsIEListExtIEs = make([]int, 2)

func (self *CriticalityDiagnosticsIEListExtIEs) GetIECount() int{
   count := 0
   if self.MessageStructure != nil { count += 1 }
   count +=1 //self.TypeOfError
   return count//ObjSet
}
func (self *CriticalityDiagnosticsIEListExtIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 88: //MessageStructure
        if self.MessageStructure != nil { return true }
      case 93: //TypeOfError
        return true //self.TypeOfError
   }
   return false//ObjSet
}
func (self *CriticalityDiagnosticsIEListExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 88: //MessageStructure
        self.MessageStructure = &MessageStructure{}
        self.MessageStructure.Unpack(st)
        self.list = append(self.list, self.MessageStructure)
      case 93: //TypeOfError
        self.TypeOfError.Unpack(st)
        self.list = append(self.list, &self.TypeOfError)
   }
}
func (self *CriticalityDiagnosticsIEListExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 88: //MessageStructure
        if self.MessageStructure != nil {self.MessageStructure.Pack(st)}
      case 93: //TypeOfError
        self.TypeOfError.Pack(st)
      default:
      break
   }
}
func init() {
table_CriticalityDiagnosticsIEListExtIEs[88] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idMessageStructure}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&MessageStructure{}, PRESENCE:Presence{Presenceoptional}, }
order_CriticalityDiagnosticsIEListExtIEs[0] = 88
table_CriticalityDiagnosticsIEListExtIEs[93] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idTypeOfError}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&TypeOfError{}, PRESENCE:Presence{Presencemandatory}, }
order_CriticalityDiagnosticsIEListExtIEs[1] = 93
   }

type MessageStructureExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *MessageStructureExtIEs)createOT() interface{}{
    return nil
}
var table_MessageStructureExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_MessageStructureExtIEs = make([]int, 0)

type CGIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *CGIExtIEs)createOT() interface{}{
    return nil
}
var table_CGIExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_CGIExtIEs = make([]int, 0)

type EncryptionInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *EncryptionInformationExtIEs)createOT() interface{}{
    return nil
}
var table_EncryptionInformationExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_EncryptionInformationExtIEs = make([]int, 0)

type GeographicalCoordinatesExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GeographicalCoordinatesExtIEs)createOT() interface{}{
    return nil
}
var table_GeographicalCoordinatesExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_GeographicalCoordinatesExtIEs = make([]int, 0)

type GAEllipsoidArcExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GAEllipsoidArcExtIEs)createOT() interface{}{
    return nil
}
var table_GAEllipsoidArcExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_GAEllipsoidArcExtIEs = make([]int, 0)

type GAPointExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GAPointExtIEs)createOT() interface{}{
    return nil
}
var table_GAPointExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_GAPointExtIEs = make([]int, 0)

type GAPointWithAltitudeExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GAPointWithAltitudeExtIEs)createOT() interface{}{
    return nil
}
var table_GAPointWithAltitudeExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_GAPointWithAltitudeExtIEs = make([]int, 0)

type GAPointWithAltitudeAndUncertaintyEllipsoidExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GAPointWithAltitudeAndUncertaintyEllipsoidExtIEs)createOT() interface{}{
    return nil
}
var table_GAPointWithAltitudeAndUncertaintyEllipsoidExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_GAPointWithAltitudeAndUncertaintyEllipsoidExtIEs = make([]int, 0)

type GAPointWithUnCertaintyExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GAPointWithUnCertaintyExtIEs)createOT() interface{}{
    return nil
}
var table_GAPointWithUnCertaintyExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_GAPointWithUnCertaintyExtIEs = make([]int, 0)

type GAPointWithUnCertaintyEllipseExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GAPointWithUnCertaintyEllipseExtIEs)createOT() interface{}{
    return nil
}
var table_GAPointWithUnCertaintyEllipseExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_GAPointWithUnCertaintyEllipseExtIEs = make([]int, 0)

type GAPolygonExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GAPolygonExtIEs)createOT() interface{}{
    return nil
}
var table_GAPolygonExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_GAPolygonExtIEs = make([]int, 0)

type GERANCellIDExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *GERANCellIDExtIEs)createOT() interface{}{
    return nil
}
var table_GERANCellIDExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_GERANCellIDExtIEs = make([]int, 0)

type IMEIGroupExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *IMEIGroupExtIEs)createOT() interface{}{
    return nil
}
var table_IMEIGroupExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_IMEIGroupExtIEs = make([]int, 0)

type IMEISVGroupExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *IMEISVGroupExtIEs)createOT() interface{}{
    return nil
}
var table_IMEISVGroupExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_IMEISVGroupExtIEs = make([]int, 0)

type IntegrityProtectionInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *IntegrityProtectionInformationExtIEs)createOT() interface{}{
    return nil
}
var table_IntegrityProtectionInformationExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_IntegrityProtectionInformationExtIEs = make([]int, 0)

type InterSystemInformationTransparentContainerExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *InterSystemInformationTransparentContainerExtIEs)createOT() interface{}{
    return nil
}
var table_InterSystemInformationTransparentContainerExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_InterSystemInformationTransparentContainerExtIEs = make([]int, 0)

type LALISTExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *LALISTExtIEs)createOT() interface{}{
    return nil
}
var table_LALISTExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_LALISTExtIEs = make([]int, 0)

type LAIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *LAIExtIEs)createOT() interface{}{
    return nil
}
var table_LAIExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_LAIExtIEs = make([]int, 0)

type LastKnownServiceAreaExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *LastKnownServiceAreaExtIEs)createOT() interface{}{
    return nil
}
var table_LastKnownServiceAreaExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_LastKnownServiceAreaExtIEs = make([]int, 0)

type InterfacesToTraceItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *InterfacesToTraceItemExtIEs)createOT() interface{}{
    return nil
}
var table_InterfacesToTraceItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_InterfacesToTraceItemExtIEs = make([]int, 0)

type PLMNsinsharednetworkExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PLMNsinsharednetworkExtIEs)createOT() interface{}{
    return nil
}
var table_PLMNsinsharednetworkExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_PLMNsinsharednetworkExtIEs = make([]int, 0)

type PositionDataExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *PositionDataExtIEs)createOT() interface{}{
    return nil
}
var table_PositionDataExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_PositionDataExtIEs = make([]int, 0)

type RABParametersExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-SignallingIndication', 'CRITICALITY': 'ignore', 'EXTENSION': 'SignallingIndication', 'PRESENCE': 'optional'}, None]}
   SignallingIndication  *SignallingIndication
   list []interface{}
}
func (self *RABParametersExtIEs)createOT() interface{}{
    return nil
}
var table_RABParametersExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABParametersExtIEs = make([]int, 1)

func (self *RABParametersExtIEs) GetIECount() int{
   count := 0
   if self.SignallingIndication != nil { count += 1 }
   return count//ObjSet
}
func (self *RABParametersExtIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 116: //SignallingIndication
        if self.SignallingIndication != nil { return true }
   }
   return false//ObjSet
}
func (self *RABParametersExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 116: //SignallingIndication
        self.SignallingIndication = &SignallingIndication{}
        self.SignallingIndication.Unpack(st)
        self.list = append(self.list, self.SignallingIndication)
   }
}
func (self *RABParametersExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 116: //SignallingIndication
        if self.SignallingIndication != nil {self.SignallingIndication.Pack(st)}
      default:
      break
   }
}
func init() {
table_RABParametersExtIEs[116] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idSignallingIndication}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&SignallingIndication{}, PRESENCE:Presence{Presenceoptional}, }
order_RABParametersExtIEs[0] = 116
   }

type RABTrCHMappingItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-CN-DomainIndicator', 'CRITICALITY': 'ignore', 'EXTENSION': 'CN-DomainIndicator', 'PRESENCE': 'optional'}, None]}
   CNDomainIndicator  *CNDomainIndicator
   list []interface{}
}
func (self *RABTrCHMappingItemExtIEs)createOT() interface{}{
    return nil
}
var table_RABTrCHMappingItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RABTrCHMappingItemExtIEs = make([]int, 1)

func (self *RABTrCHMappingItemExtIEs) GetIECount() int{
   count := 0
   if self.CNDomainIndicator != nil { count += 1 }
   return count//ObjSet
}
func (self *RABTrCHMappingItemExtIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 3: //CNDomainIndicator
        if self.CNDomainIndicator != nil { return true }
   }
   return false//ObjSet
}
func (self *RABTrCHMappingItemExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 3: //CNDomainIndicator
        self.CNDomainIndicator = &CNDomainIndicator{}
        self.CNDomainIndicator.Unpack(st)
        self.list = append(self.list, self.CNDomainIndicator)
   }
}
func (self *RABTrCHMappingItemExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 3: //CNDomainIndicator
        if self.CNDomainIndicator != nil {self.CNDomainIndicator.Pack(st)}
      default:
      break
   }
}
func init() {
table_RABTrCHMappingItemExtIEs[3] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idCNDomainIndicator}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&CNDomainIndicator{}, PRESENCE:Presence{Presenceoptional}, }
order_RABTrCHMappingItemExtIEs[0] = 3
   }

type RAIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RAIExtIEs)createOT() interface{}{
    return nil
}
var table_RAIExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RAIExtIEs = make([]int, 0)

type RequestedRABParameterValuesExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RequestedRABParameterValuesExtIEs)createOT() interface{}{
    return nil
}
var table_RequestedRABParameterValuesExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RequestedRABParameterValuesExtIEs = make([]int, 0)

type ResidualBitErrorRatioExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *ResidualBitErrorRatioExtIEs)createOT() interface{}{
    return nil
}
var table_ResidualBitErrorRatioExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_ResidualBitErrorRatioExtIEs = make([]int, 0)

type RIMTransferExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RIMTransferExtIEs)createOT() interface{}{
    return nil
}
var table_RIMTransferExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RIMTransferExtIEs = make([]int, 0)

type RNCTraceInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *RNCTraceInformationExtIEs)createOT() interface{}{
    return nil
}
var table_RNCTraceInformationExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_RNCTraceInformationExtIEs = make([]int, 0)

type SAIExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SAIExtIEs)createOT() interface{}{
    return nil
}
var table_SAIExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SAIExtIEs = make([]int, 0)

type SharedNetworkInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SharedNetworkInformationExtIEs)createOT() interface{}{
    return nil
}
var table_SharedNetworkInformationExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SharedNetworkInformationExtIEs = make([]int, 0)

type SDUErrorRatioExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SDUErrorRatioExtIEs)createOT() interface{}{
    return nil
}
var table_SDUErrorRatioExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SDUErrorRatioExtIEs = make([]int, 0)

type SDUFormatInformationParametersExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SDUFormatInformationParametersExtIEs)createOT() interface{}{
    return nil
}
var table_SDUFormatInformationParametersExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SDUFormatInformationParametersExtIEs = make([]int, 0)

type SDUParametersExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SDUParametersExtIEs)createOT() interface{}{
    return nil
}
var table_SDUParametersExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SDUParametersExtIEs = make([]int, 0)

type SNAAccessInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SNAAccessInformationExtIEs)createOT() interface{}{
    return nil
}
var table_SNAAccessInformationExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SNAAccessInformationExtIEs = make([]int, 0)

type SourceRNCIDExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SourceRNCIDExtIEs)createOT() interface{}{
    return nil
}
var table_SourceRNCIDExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SourceRNCIDExtIEs = make([]int, 0)

type SourceRNCToTargetRNCTransparentContainerExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-SRB-TrCH-Mapping', 'CRITICALITY': 'reject', 'EXTENSION': 'SRB-TrCH-Mapping', 'PRESENCE': 'optional'}, {'ID': 'id-CellLoadInformationGroup', 'CRITICALITY': 'ignore', 'EXTENSION': 'CellLoadInformationGroup', 'PRESENCE': 'optional'}, {'ID': 'id-TraceRecordingSessionInformation', 'CRITICALITY': 'ignore', 'EXTENSION': 'TraceRecordingSessionInformation', 'PRESENCE': 'optional'}, None]}
   SRBTrCHMapping  *SRBTrCHMapping
   CellLoadInformationGroup  *CellLoadInformationGroup
   TraceRecordingSessionInformation  *TraceRecordingSessionInformation
   list []interface{}
}
func (self *SourceRNCToTargetRNCTransparentContainerExtIEs)createOT() interface{}{
    return nil
}
var table_SourceRNCToTargetRNCTransparentContainerExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SourceRNCToTargetRNCTransparentContainerExtIEs = make([]int, 3)

func (self *SourceRNCToTargetRNCTransparentContainerExtIEs) GetIECount() int{
   count := 0
   if self.SRBTrCHMapping != nil { count += 1 }
   if self.CellLoadInformationGroup != nil { count += 1 }
   if self.TraceRecordingSessionInformation != nil { count += 1 }
   return count//ObjSet
}
func (self *SourceRNCToTargetRNCTransparentContainerExtIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 98: //SRBTrCHMapping
        if self.SRBTrCHMapping != nil { return true }
      case 121: //CellLoadInformationGroup
        if self.CellLoadInformationGroup != nil { return true }
      case 124: //TraceRecordingSessionInformation
        if self.TraceRecordingSessionInformation != nil { return true }
   }
   return false//ObjSet
}
func (self *SourceRNCToTargetRNCTransparentContainerExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 98: //SRBTrCHMapping
        self.SRBTrCHMapping = &SRBTrCHMapping{}
        self.SRBTrCHMapping.Unpack(st)
        self.list = append(self.list, self.SRBTrCHMapping)
      case 121: //CellLoadInformationGroup
        self.CellLoadInformationGroup = &CellLoadInformationGroup{}
        self.CellLoadInformationGroup.Unpack(st)
        self.list = append(self.list, self.CellLoadInformationGroup)
      case 124: //TraceRecordingSessionInformation
        self.TraceRecordingSessionInformation = &TraceRecordingSessionInformation{}
        self.TraceRecordingSessionInformation.Unpack(st)
        self.list = append(self.list, self.TraceRecordingSessionInformation)
   }
}
func (self *SourceRNCToTargetRNCTransparentContainerExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 98: //SRBTrCHMapping
        if self.SRBTrCHMapping != nil {self.SRBTrCHMapping.Pack(st)}
      case 121: //CellLoadInformationGroup
        if self.CellLoadInformationGroup != nil {self.CellLoadInformationGroup.Pack(st)}
      case 124: //TraceRecordingSessionInformation
        if self.TraceRecordingSessionInformation != nil {self.TraceRecordingSessionInformation.Pack(st)}
      default:
      break
   }
}
func init() {
table_SourceRNCToTargetRNCTransparentContainerExtIEs[98] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idSRBTrCHMapping}, CRITICALITY:Criticality{Criticalityreject}, EXTENSION:&SRBTrCHMapping{}, PRESENCE:Presence{Presenceoptional}, }
order_SourceRNCToTargetRNCTransparentContainerExtIEs[0] = 98
table_SourceRNCToTargetRNCTransparentContainerExtIEs[121] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idCellLoadInformationGroup}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&CellLoadInformationGroup{}, PRESENCE:Presence{Presenceoptional}, }
order_SourceRNCToTargetRNCTransparentContainerExtIEs[1] = 121
table_SourceRNCToTargetRNCTransparentContainerExtIEs[124] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idTraceRecordingSessionInformation}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&TraceRecordingSessionInformation{}, PRESENCE:Presence{Presenceoptional}, }
order_SourceRNCToTargetRNCTransparentContainerExtIEs[2] = 124
   }

type SourceUTRANCellIDExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SourceUTRANCellIDExtIEs)createOT() interface{}{
    return nil
}
var table_SourceUTRANCellIDExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SourceUTRANCellIDExtIEs = make([]int, 0)

type SRBTrCHMappingItemExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *SRBTrCHMappingItemExtIEs)createOT() interface{}{
    return nil
}
var table_SRBTrCHMappingItemExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_SRBTrCHMappingItemExtIEs = make([]int, 0)

type TargetRNCIDExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TargetRNCIDExtIEs)createOT() interface{}{
    return nil
}
var table_TargetRNCIDExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_TargetRNCIDExtIEs = make([]int, 0)

type TargetRNCToSourceRNCTransparentContainerExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TargetRNCToSourceRNCTransparentContainerExtIEs)createOT() interface{}{
    return nil
}
var table_TargetRNCToSourceRNCTransparentContainerExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_TargetRNCToSourceRNCTransparentContainerExtIEs = make([]int, 0)

type TracePropagationParametersExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TracePropagationParametersExtIEs)createOT() interface{}{
    return nil
}
var table_TracePropagationParametersExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_TracePropagationParametersExtIEs = make([]int, 0)

type TraceRecordingSessionInformationExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *TraceRecordingSessionInformationExtIEs)createOT() interface{}{
    return nil
}
var table_TraceRecordingSessionInformationExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_TraceRecordingSessionInformationExtIEs = make([]int, 0)

type TrCHIDExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [{'ID': 'id-hS-DSCH-MAC-d-Flow-ID', 'CRITICALITY': 'ignore', 'EXTENSION': 'HS-DSCH-MAC-d-Flow-ID', 'PRESENCE': 'optional'}, None]}
   HSDSCHMACdFlowID  *HSDSCHMACdFlowID
   list []interface{}
}
func (self *TrCHIDExtIEs)createOT() interface{}{
    return nil
}
var table_TrCHIDExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_TrCHIDExtIEs = make([]int, 1)

func (self *TrCHIDExtIEs) GetIECount() int{
   count := 0
   if self.HSDSCHMACdFlowID != nil { count += 1 }
   return count//ObjSet
}
func (self *TrCHIDExtIEs) GetOT(id interface{}) bool{
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 117: //HSDSCHMACdFlowID
        if self.HSDSCHMACdFlowID != nil { return true }
   }
   return false//ObjSet
}
func (self *TrCHIDExtIEs)UnpackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 117: //HSDSCHMACdFlowID
        self.HSDSCHMACdFlowID = &HSDSCHMACdFlowID{}
        self.HSDSCHMACdFlowID.Unpack(st)
        self.list = append(self.list, self.HSDSCHMACdFlowID)
   }
}
func (self *TrCHIDExtIEs)PackOT(st *Stream, id interface{}){
   cat := id.(RANAPPROTOCOLEXTENSIONid).Value
   switch cat {
      case 117: //HSDSCHMACdFlowID
        if self.HSDSCHMACdFlowID != nil {self.HSDSCHMACdFlowID.Pack(st)}
      default:
      break
   }
}
func init() {
table_TrCHIDExtIEs[117] = &RANAPPROTOCOLEXTENSION{ID:ProtocolExtensionID{idhSDSCHMACdFlowID}, CRITICALITY:Criticality{Criticalityignore}, EXTENSION:&HSDSCHMACdFlowID{}, PRESENCE:Presence{Presenceoptional}, }
order_TrCHIDExtIEs[0] = 117
   }

type UESBIIuExtIEs struct { //ObjSet 1 {'ID': {'type': 'ProtocolExtensionID', 'name': '&id'}, 'CRITICALITY': {'type': 'Criticality', 'name': '&criticality'}, 'EXTENSION': {'type': 'OpenType', 'name': '&Extension'}, 'PRESENCE': {'type': 'Presence', 'name': '&presence'}} {'class': 'RANAP-PROTOCOL-EXTENSION', 'members': [None]}
   list []interface{}
}
func (self *UESBIIuExtIEs)createOT() interface{}{
    return nil
}
var table_UESBIIuExtIEs = make(map[int]*RANAPPROTOCOLEXTENSION)

var order_UESBIIuExtIEs = make([]int, 0)

//class RANAPELEMENTARYPROCEDURES: #OBJSET1 {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'members': [{}, {}, {}, None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_RANAPELEMENTARYPROCEDURES = make(map[RANAPELEMENTARYPROCEDUREprocedureCode]*RANAPELEMENTARYPROCEDURE)


//class RANAPELEMENTARYPROCEDURESCLASS1: #OBJSET1 {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_RANAPELEMENTARYPROCEDURESCLASS1 = make(map[RANAPELEMENTARYPROCEDUREprocedureCode]*RANAPELEMENTARYPROCEDURE)


//class RANAPELEMENTARYPROCEDURESCLASS2: #OBJSET1 {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_RANAPELEMENTARYPROCEDURESCLASS2 = make(map[RANAPELEMENTARYPROCEDUREprocedureCode]*RANAPELEMENTARYPROCEDURE)


//class RANAPELEMENTARYPROCEDURESCLASS3: #OBJSET1 {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'members': [None]} {'type': 'ProcedureCode', 'name': '&procedureCode'}
var table_RANAPELEMENTARYPROCEDURESCLASS3 = make(map[RANAPELEMENTARYPROCEDUREprocedureCode]*RANAPELEMENTARYPROCEDURE)


var idRABAssignment uint64 = 0
const ProcedureCodeRABAssignment = 0
var idIuRelease uint64 = 1
const ProcedureCodeIuRelease = 1
var idRelocationPreparation uint64 = 2
const ProcedureCodeRelocationPreparation = 2
var idRelocationResourceAllocation uint64 = 3
const ProcedureCodeRelocationResourceAllocation = 3
var idRelocationCancel uint64 = 4
const ProcedureCodeRelocationCancel = 4
var idSRNSContextTransfer uint64 = 5
const ProcedureCodeSRNSContextTransfer = 5
var idSecurityModeControl uint64 = 6
const ProcedureCodeSecurityModeControl = 6
var idDataVolumeReport uint64 = 7
const ProcedureCodeDataVolumeReport = 7
var idReset uint64 = 9
const ProcedureCodeReset = 9
var idRABReleaseRequest uint64 = 10
const ProcedureCodeRABReleaseRequest = 10
var idIuReleaseRequest uint64 = 11
const ProcedureCodeIuReleaseRequest = 11
var idRelocationDetect uint64 = 12
const ProcedureCodeRelocationDetect = 12
var idRelocationComplete uint64 = 13
const ProcedureCodeRelocationComplete = 13
var idPaging uint64 = 14
const ProcedureCodePaging = 14
var idCommonID uint64 = 15
const ProcedureCodeCommonID = 15
var idCNInvokeTrace uint64 = 16
const ProcedureCodeCNInvokeTrace = 16
var idLocationReportingControl uint64 = 17
const ProcedureCodeLocationReportingControl = 17
var idLocationReport uint64 = 18
const ProcedureCodeLocationReport = 18
var idInitialUEMessage uint64 = 19
const ProcedureCodeInitialUEMessage = 19
var idDirectTransfer uint64 = 20
const ProcedureCodeDirectTransfer = 20
var idOverloadControl uint64 = 21
const ProcedureCodeOverloadControl = 21
var idErrorIndication uint64 = 22
const ProcedureCodeErrorIndication = 22
var idSRNSDataForward uint64 = 23
const ProcedureCodeSRNSDataForward = 23
var idForwardSRNSContext uint64 = 24
const ProcedureCodeForwardSRNSContext = 24
var idprivateMessage uint64 = 25
const ProcedureCodeprivateMessage = 25
var idCNDeactivateTrace uint64 = 26
const ProcedureCodeCNDeactivateTrace = 26
var idResetResource uint64 = 27
const ProcedureCodeResetResource = 27
var idRANAPRelocation uint64 = 28
const ProcedureCodeRANAPRelocation = 28
var idRABModifyRequest uint64 = 29
const ProcedureCodeRABModifyRequest = 29
var idLocationRelatedData uint64 = 30
const ProcedureCodeLocationRelatedData = 30
var idInformationTransfer uint64 = 31
const ProcedureCodeInformationTransfer = 31
var idUESpecificInformation uint64 = 32
const ProcedureCodeUESpecificInformation = 32
var idUplinkInformationTransfer uint64 = 33
const ProcedureCodeUplinkInformationTransfer = 33
var idDirectInformationTransfer uint64 = 34
const ProcedureCodeDirectInformationTransfer = 34
var maxPrivateIEs uint64 = 65535
var maxProtocolExtensions uint64 = 65535
var maxProtocolIEs uint64 = 65535
var maxNrOfDTs uint64 = 15
var maxNrOfErrors uint64 = 256
var maxNrOfIuSigConIds uint64 = 250
var maxNrOfPDPDirections uint64 = 2
var maxNrOfPoints uint64 = 15
var maxNrOfRABs uint64 = 256
var maxNrOfSeparateTrafficDirections uint64 = 2
var maxNrOfSRBs uint64 = 8
var maxNrOfVol uint64 = 2
var maxNrOfLevels uint64 = 256
var maxNrOfAltValues uint64 = 16
var maxNrOfPLMNsSN uint64 = 32
var maxNrOfLAs uint64 = 65536
var maxNrOfSNAs uint64 = 65536
var maxNrOfUEsToBeTraced uint64 = 64
var maxNrOfInterfaces uint64 = 16
var maxRABSubflows uint64 = 7
var maxRABSubflowCombination uint64 = 64
var maxSet uint64 = 9
var idAreaIdentity uint64 = 0
const INTEGERAreaIdentity = 0
var idCNDomainIndicator uint64 = 3
const INTEGERCNDomainIndicator = 3
var idCause uint64 = 4
const INTEGERCause = 4
var idChosenEncryptionAlgorithm uint64 = 5
const INTEGERChosenEncryptionAlgorithm = 5
var idChosenIntegrityProtectionAlgorithm uint64 = 6
const INTEGERChosenIntegrityProtectionAlgorithm = 6
var idClassmarkInformation2 uint64 = 7
const INTEGERClassmarkInformation2 = 7
var idClassmarkInformation3 uint64 = 8
const INTEGERClassmarkInformation3 = 8
var idCriticalityDiagnostics uint64 = 9
const INTEGERCriticalityDiagnostics = 9
var idDLGTPPDUSequenceNumber uint64 = 10
const INTEGERDLGTPPDUSequenceNumber = 10
var idEncryptionInformation uint64 = 11
const INTEGEREncryptionInformation = 11
var idIntegrityProtectionInformation uint64 = 12
const INTEGERIntegrityProtectionInformation = 12
var idIuTransportAssociation uint64 = 13
const INTEGERIuTransportAssociation = 13
var idL3Information uint64 = 14
const INTEGERL3Information = 14
var idLAI uint64 = 15
const INTEGERLAI = 15
var idNASPDU uint64 = 16
const INTEGERNASPDU = 16
var idNonSearchingIndication uint64 = 17
const INTEGERNonSearchingIndication = 17
var idNumberOfSteps uint64 = 18
const INTEGERNumberOfSteps = 18
var idOMCID uint64 = 19
const INTEGEROMCID = 19
var idOldBSSToNewBSSInformation uint64 = 20
const INTEGEROldBSSToNewBSSInformation = 20
var idPagingAreaID uint64 = 21
const INTEGERPagingAreaID = 21
var idPagingCause uint64 = 22
const INTEGERPagingCause = 22
var idPermanentNASUEID uint64 = 23
const INTEGERPermanentNASUEID = 23
var idRABContextItem uint64 = 24
const INTEGERRABContextItem = 24
var idRABContextList uint64 = 25
const INTEGERRABContextList = 25
var idRABDataForwardingItem uint64 = 26
const INTEGERRABDataForwardingItem = 26
var idRABDataForwardingItemSRNSCtxReq uint64 = 27
const INTEGERRABDataForwardingItemSRNSCtxReq = 27
var idRABDataForwardingList uint64 = 28
const INTEGERRABDataForwardingList = 28
var idRABDataForwardingListSRNSCtxReq uint64 = 29
const INTEGERRABDataForwardingListSRNSCtxReq = 29
var idRABDataVolumeReportItem uint64 = 30
const INTEGERRABDataVolumeReportItem = 30
var idRABDataVolumeReportList uint64 = 31
const INTEGERRABDataVolumeReportList = 31
var idRABDataVolumeReportRequestItem uint64 = 32
const INTEGERRABDataVolumeReportRequestItem = 32
var idRABDataVolumeReportRequestList uint64 = 33
const INTEGERRABDataVolumeReportRequestList = 33
var idRABFailedItem uint64 = 34
const INTEGERRABFailedItem = 34
var idRABFailedList uint64 = 35
const INTEGERRABFailedList = 35
var idRABID uint64 = 36
const INTEGERRABID = 36
var idRABQueuedItem uint64 = 37
const INTEGERRABQueuedItem = 37
var idRABQueuedList uint64 = 38
const INTEGERRABQueuedList = 38
var idRABReleaseFailedList uint64 = 39
const INTEGERRABReleaseFailedList = 39
var idRABReleaseItem uint64 = 40
const INTEGERRABReleaseItem = 40
var idRABReleaseList uint64 = 41
const INTEGERRABReleaseList = 41
var idRABReleasedItem uint64 = 42
const INTEGERRABReleasedItem = 42
var idRABReleasedList uint64 = 43
const INTEGERRABReleasedList = 43
var idRABReleasedListIuRelComp uint64 = 44
const INTEGERRABReleasedListIuRelComp = 44
var idRABRelocationReleaseItem uint64 = 45
const INTEGERRABRelocationReleaseItem = 45
var idRABRelocationReleaseList uint64 = 46
const INTEGERRABRelocationReleaseList = 46
var idRABSetupItemRelocReq uint64 = 47
const INTEGERRABSetupItemRelocReq = 47
var idRABSetupItemRelocReqAck uint64 = 48
const INTEGERRABSetupItemRelocReqAck = 48
var idRABSetupListRelocReq uint64 = 49
const INTEGERRABSetupListRelocReq = 49
var idRABSetupListRelocReqAck uint64 = 50
const INTEGERRABSetupListRelocReqAck = 50
var idRABSetupOrModifiedItem uint64 = 51
const INTEGERRABSetupOrModifiedItem = 51
var idRABSetupOrModifiedList uint64 = 52
const INTEGERRABSetupOrModifiedList = 52
var idRABSetupOrModifyItem uint64 = 53
const INTEGERRABSetupOrModifyItem = 53
var idRABSetupOrModifyList uint64 = 54
const INTEGERRABSetupOrModifyList = 54
var idRAC uint64 = 55
const INTEGERRAC = 55
var idRelocationType uint64 = 56
const INTEGERRelocationType = 56
var idRequestType uint64 = 57
const INTEGERRequestType = 57
var idSAI uint64 = 58
const INTEGERSAI = 58
var idSAPI uint64 = 59
const INTEGERSAPI = 59
var idSourceID uint64 = 60
const INTEGERSourceID = 60
var idSourceRNCToTargetRNCTransparentContainer uint64 = 61
const INTEGERSourceRNCToTargetRNCTransparentContainer = 61
var idTargetID uint64 = 62
const INTEGERTargetID = 62
var idTargetRNCToSourceRNCTransparentContainer uint64 = 63
const INTEGERTargetRNCToSourceRNCTransparentContainer = 63
var idTemporaryUEID uint64 = 64
const INTEGERTemporaryUEID = 64
var idTraceReference uint64 = 65
const INTEGERTraceReference = 65
var idTraceType uint64 = 66
const INTEGERTraceType = 66
var idTransportLayerAddress uint64 = 67
const INTEGERTransportLayerAddress = 67
var idTriggerID uint64 = 68
const INTEGERTriggerID = 68
var idUEID uint64 = 69
const INTEGERUEID = 69
var idULGTPPDUSequenceNumber uint64 = 70
const INTEGERULGTPPDUSequenceNumber = 70
var idRABFailedtoReportItem uint64 = 71
const INTEGERRABFailedtoReportItem = 71
var idRABFailedtoReportList uint64 = 72
const INTEGERRABFailedtoReportList = 72
var idKeyStatus uint64 = 75
const INTEGERKeyStatus = 75
var idDRXCycleLengthCoefficient uint64 = 76
const INTEGERDRXCycleLengthCoefficient = 76
var idIuSigConIdList uint64 = 77
const INTEGERIuSigConIdList = 77
var idIuSigConIdItem uint64 = 78
const INTEGERIuSigConIdItem = 78
var idIuSigConId uint64 = 79
const INTEGERIuSigConId = 79
var idDirectTransferInformationItemRANAPRelocInf uint64 = 80
const INTEGERDirectTransferInformationItemRANAPRelocInf = 80
var idDirectTransferInformationListRANAPRelocInf uint64 = 81
const INTEGERDirectTransferInformationListRANAPRelocInf = 81
var idRABContextItemRANAPRelocInf uint64 = 82
const INTEGERRABContextItemRANAPRelocInf = 82
var idRABContextListRANAPRelocInf uint64 = 83
const INTEGERRABContextListRANAPRelocInf = 83
var idRABContextFailedtoTransferItem uint64 = 84
const INTEGERRABContextFailedtoTransferItem = 84
var idRABContextFailedtoTransferList uint64 = 85
const INTEGERRABContextFailedtoTransferList = 85
var idGlobalRNCID uint64 = 86
const INTEGERGlobalRNCID = 86
var idRABReleasedItemIuRelComp uint64 = 87
const INTEGERRABReleasedItemIuRelComp = 87
var idMessageStructure uint64 = 88
const INTEGERMessageStructure = 88
var idAltRABParameters uint64 = 89
const INTEGERAltRABParameters = 89
var idAssRABParameters uint64 = 90
const INTEGERAssRABParameters = 90
var idRABModifyList uint64 = 91
const INTEGERRABModifyList = 91
var idRABModifyItem uint64 = 92
const INTEGERRABModifyItem = 92
var idTypeOfError uint64 = 93
const INTEGERTypeOfError = 93
var idBroadcastAssistanceDataDecipheringKeys uint64 = 94
const INTEGERBroadcastAssistanceDataDecipheringKeys = 94
var idLocationRelatedDataRequestType uint64 = 95
const INTEGERLocationRelatedDataRequestType = 95
var idGlobalCNID uint64 = 96
const INTEGERGlobalCNID = 96
var idLastKnownServiceArea uint64 = 97
const INTEGERLastKnownServiceArea = 97
var idSRBTrCHMapping uint64 = 98
const INTEGERSRBTrCHMapping = 98
var idInterSystemInformationTransparentContainer uint64 = 99
const INTEGERInterSystemInformationTransparentContainer = 99
var idNewBSSToOldBSSInformation uint64 = 100
const INTEGERNewBSSToOldBSSInformation = 100
var idSourceRNCPDCPcontextinfo uint64 = 103
const INTEGERSourceRNCPDCPcontextinfo = 103
var idInformationTransferID uint64 = 104
const INTEGERInformationTransferID = 104
var idSNAAccessInformation uint64 = 105
const INTEGERSNAAccessInformation = 105
var idProvidedData uint64 = 106
const INTEGERProvidedData = 106
var idGERANBSCContainer uint64 = 107
const INTEGERGERANBSCContainer = 107
var idGERANClassmark uint64 = 108
const INTEGERGERANClassmark = 108
var idGERANIumodeRABFailedRABAssgntResponseItem uint64 = 109
const INTEGERGERANIumodeRABFailedRABAssgntResponseItem = 109
var idGERANIumodeRABFailedListRABAssgntResponse uint64 = 110
const INTEGERGERANIumodeRABFailedListRABAssgntResponse = 110
var idVerticalAccuracyCode uint64 = 111
const INTEGERVerticalAccuracyCode = 111
var idResponseTime uint64 = 112
const INTEGERResponseTime = 112
var idPositioningPriority uint64 = 113
const INTEGERPositioningPriority = 113
var idClientType uint64 = 114
const INTEGERClientType = 114
var idLocationRelatedDataRequestTypeSpecificToGERANIuMode uint64 = 115
const INTEGERLocationRelatedDataRequestTypeSpecificToGERANIuMode = 115
var idSignallingIndication uint64 = 116
const INTEGERSignallingIndication = 116
var idhSDSCHMACdFlowID uint64 = 117
const INTEGERhSDSCHMACdFlowID = 117
var idUESBIIu uint64 = 118
const INTEGERUESBIIu = 118
var idPositionData uint64 = 119
const INTEGERPositionData = 119
var idPositionDataSpecificToGERANIuMode uint64 = 120
const INTEGERPositionDataSpecificToGERANIuMode = 120
var idCellLoadInformationGroup uint64 = 121
const INTEGERCellLoadInformationGroup = 121
var idAccuracyFulfilmentIndicator uint64 = 122
const INTEGERAccuracyFulfilmentIndicator = 122
var idInformationTransferType uint64 = 123
const INTEGERInformationTransferType = 123
var idTraceRecordingSessionInformation uint64 = 124
const INTEGERTraceRecordingSessionInformation = 124
var idTracePropagationParameters uint64 = 125
const INTEGERTracePropagationParameters = 125
var idInterSystemInformationTransferType uint64 = 126
const INTEGERInterSystemInformationTransferType = 126
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'Iu-ReleaseCommand', 'SUCCESSFUL OUTCOME': 'Iu-ReleaseComplete', 'PROCEDURE CODE': 'id-Iu-Release', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idIuRelease}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&IuReleaseCommand{}, SUCCESSFULOUTCOME:&IuReleaseComplete{}, PROCEDURECODE:ProcedureCode{idIuRelease}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetIuReleaseINITIATINGMESSAGE() (*IuReleaseCommand, uint64, int) {/*TYPE, ID, Cricality*/
 return &IuReleaseCommand{}, uint64(idIuRelease), int(Criticalityreject)
}
func GetIuReleaseSUCCESSFULOUTCOME() (*IuReleaseComplete, uint64, int) {/*TYPE, ID, Cricality*/
 return &IuReleaseComplete{}, uint64(idIuRelease), int(Criticalityreject)
}
func (self *IuReleaseCommand) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *IuReleaseCommand) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *IuReleaseCommand) createOT() interface{} {
   return &IuReleaseCommand{}
}
func (self *IuReleaseCommand) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *IuReleaseCommand) GetIECount() int{
    return 0
}
func (self *IuReleaseComplete) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *IuReleaseComplete) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *IuReleaseComplete) createOT() interface{} {
   return &IuReleaseComplete{}
}
func (self *IuReleaseComplete) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *IuReleaseComplete) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RelocationRequired', 'SUCCESSFUL OUTCOME': 'RelocationCommand', 'UNSUCCESSFUL OUTCOME': 'RelocationPreparationFailure', 'PROCEDURE CODE': 'id-RelocationPreparation', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idRelocationPreparation}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RelocationRequired{}, SUCCESSFULOUTCOME:&RelocationCommand{}, UNSUCCESSFULOUTCOME:&RelocationPreparationFailure{}, PROCEDURECODE:ProcedureCode{idRelocationPreparation}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetRelocationPreparationINITIATINGMESSAGE() (*RelocationRequired, uint64, int) {/*TYPE, ID, Cricality*/
 return &RelocationRequired{}, uint64(idRelocationPreparation), int(Criticalityreject)
}
func GetRelocationPreparationSUCCESSFULOUTCOME() (*RelocationCommand, uint64, int) {/*TYPE, ID, Cricality*/
 return &RelocationCommand{}, uint64(idRelocationPreparation), int(Criticalityreject)
}
func GetRelocationPreparationUNSUCCESSFULOUTCOME() (*RelocationPreparationFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &RelocationPreparationFailure{}, uint64(idRelocationPreparation), int(Criticalityreject)
}
func (self *RelocationRequired) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RelocationRequired) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RelocationRequired) createOT() interface{} {
   return &RelocationRequired{}
}
func (self *RelocationRequired) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RelocationRequired) GetIECount() int{
    return 0
}
func (self *RelocationCommand) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RelocationCommand) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RelocationCommand) createOT() interface{} {
   return &RelocationCommand{}
}
func (self *RelocationCommand) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RelocationCommand) GetIECount() int{
    return 0
}
func (self *RelocationPreparationFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RelocationPreparationFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RelocationPreparationFailure) createOT() interface{} {
   return &RelocationPreparationFailure{}
}
func (self *RelocationPreparationFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RelocationPreparationFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RelocationRequest', 'SUCCESSFUL OUTCOME': 'RelocationRequestAcknowledge', 'UNSUCCESSFUL OUTCOME': 'RelocationFailure', 'PROCEDURE CODE': 'id-RelocationResourceAllocation', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idRelocationResourceAllocation}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RelocationRequest{}, SUCCESSFULOUTCOME:&RelocationRequestAcknowledge{}, UNSUCCESSFULOUTCOME:&RelocationFailure{}, PROCEDURECODE:ProcedureCode{idRelocationResourceAllocation}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetRelocationResourceAllocationINITIATINGMESSAGE() (*RelocationRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &RelocationRequest{}, uint64(idRelocationResourceAllocation), int(Criticalityreject)
}
func GetRelocationResourceAllocationSUCCESSFULOUTCOME() (*RelocationRequestAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &RelocationRequestAcknowledge{}, uint64(idRelocationResourceAllocation), int(Criticalityreject)
}
func GetRelocationResourceAllocationUNSUCCESSFULOUTCOME() (*RelocationFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &RelocationFailure{}, uint64(idRelocationResourceAllocation), int(Criticalityreject)
}
func (self *RelocationRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RelocationRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RelocationRequest) createOT() interface{} {
   return &RelocationRequest{}
}
func (self *RelocationRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RelocationRequest) GetIECount() int{
    return 0
}
func (self *RelocationRequestAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RelocationRequestAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RelocationRequestAcknowledge) createOT() interface{} {
   return &RelocationRequestAcknowledge{}
}
func (self *RelocationRequestAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RelocationRequestAcknowledge) GetIECount() int{
    return 0
}
func (self *RelocationFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RelocationFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RelocationFailure) createOT() interface{} {
   return &RelocationFailure{}
}
func (self *RelocationFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RelocationFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RelocationCancel', 'SUCCESSFUL OUTCOME': 'RelocationCancelAcknowledge', 'PROCEDURE CODE': 'id-RelocationCancel', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idRelocationCancel}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RelocationCancel{}, SUCCESSFULOUTCOME:&RelocationCancelAcknowledge{}, PROCEDURECODE:ProcedureCode{idRelocationCancel}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetRelocationCancelINITIATINGMESSAGE() (*RelocationCancel, uint64, int) {/*TYPE, ID, Cricality*/
 return &RelocationCancel{}, uint64(idRelocationCancel), int(Criticalityreject)
}
func GetRelocationCancelSUCCESSFULOUTCOME() (*RelocationCancelAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &RelocationCancelAcknowledge{}, uint64(idRelocationCancel), int(Criticalityreject)
}
func (self *RelocationCancel) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RelocationCancel) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RelocationCancel) createOT() interface{} {
   return &RelocationCancel{}
}
func (self *RelocationCancel) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RelocationCancel) GetIECount() int{
    return 0
}
func (self *RelocationCancelAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RelocationCancelAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RelocationCancelAcknowledge) createOT() interface{} {
   return &RelocationCancelAcknowledge{}
}
func (self *RelocationCancelAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RelocationCancelAcknowledge) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'SRNS-ContextRequest', 'SUCCESSFUL OUTCOME': 'SRNS-ContextResponse', 'PROCEDURE CODE': 'id-SRNS-ContextTransfer', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idSRNSContextTransfer}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&SRNSContextRequest{}, SUCCESSFULOUTCOME:&SRNSContextResponse{}, PROCEDURECODE:ProcedureCode{idSRNSContextTransfer}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetSRNSContextTransferINITIATINGMESSAGE() (*SRNSContextRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &SRNSContextRequest{}, uint64(idSRNSContextTransfer), int(Criticalityreject)
}
func GetSRNSContextTransferSUCCESSFULOUTCOME() (*SRNSContextResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &SRNSContextResponse{}, uint64(idSRNSContextTransfer), int(Criticalityreject)
}
func (self *SRNSContextRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SRNSContextRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SRNSContextRequest) createOT() interface{} {
   return &SRNSContextRequest{}
}
func (self *SRNSContextRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SRNSContextRequest) GetIECount() int{
    return 0
}
func (self *SRNSContextResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SRNSContextResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SRNSContextResponse) createOT() interface{} {
   return &SRNSContextResponse{}
}
func (self *SRNSContextResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SRNSContextResponse) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'SecurityModeCommand', 'SUCCESSFUL OUTCOME': 'SecurityModeComplete', 'UNSUCCESSFUL OUTCOME': 'SecurityModeReject', 'PROCEDURE CODE': 'id-SecurityModeControl', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idSecurityModeControl}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&SecurityModeCommand{}, SUCCESSFULOUTCOME:&SecurityModeComplete{}, UNSUCCESSFULOUTCOME:&SecurityModeReject{}, PROCEDURECODE:ProcedureCode{idSecurityModeControl}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetSecurityModeControlINITIATINGMESSAGE() (*SecurityModeCommand, uint64, int) {/*TYPE, ID, Cricality*/
 return &SecurityModeCommand{}, uint64(idSecurityModeControl), int(Criticalityreject)
}
func GetSecurityModeControlSUCCESSFULOUTCOME() (*SecurityModeComplete, uint64, int) {/*TYPE, ID, Cricality*/
 return &SecurityModeComplete{}, uint64(idSecurityModeControl), int(Criticalityreject)
}
func GetSecurityModeControlUNSUCCESSFULOUTCOME() (*SecurityModeReject, uint64, int) {/*TYPE, ID, Cricality*/
 return &SecurityModeReject{}, uint64(idSecurityModeControl), int(Criticalityreject)
}
func (self *SecurityModeCommand) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SecurityModeCommand) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SecurityModeCommand) createOT() interface{} {
   return &SecurityModeCommand{}
}
func (self *SecurityModeCommand) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SecurityModeCommand) GetIECount() int{
    return 0
}
func (self *SecurityModeComplete) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SecurityModeComplete) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SecurityModeComplete) createOT() interface{} {
   return &SecurityModeComplete{}
}
func (self *SecurityModeComplete) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SecurityModeComplete) GetIECount() int{
    return 0
}
func (self *SecurityModeReject) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SecurityModeReject) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SecurityModeReject) createOT() interface{} {
   return &SecurityModeReject{}
}
func (self *SecurityModeReject) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SecurityModeReject) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'DataVolumeReportRequest', 'SUCCESSFUL OUTCOME': 'DataVolumeReport', 'PROCEDURE CODE': 'id-DataVolumeReport', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idDataVolumeReport}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&DataVolumeReportRequest{}, SUCCESSFULOUTCOME:&DataVolumeReport{}, PROCEDURECODE:ProcedureCode{idDataVolumeReport}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetDataVolumeReportINITIATINGMESSAGE() (*DataVolumeReportRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &DataVolumeReportRequest{}, uint64(idDataVolumeReport), int(Criticalityreject)
}
func GetDataVolumeReportSUCCESSFULOUTCOME() (*DataVolumeReport, uint64, int) {/*TYPE, ID, Cricality*/
 return &DataVolumeReport{}, uint64(idDataVolumeReport), int(Criticalityreject)
}
func (self *DataVolumeReportRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *DataVolumeReportRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *DataVolumeReportRequest) createOT() interface{} {
   return &DataVolumeReportRequest{}
}
func (self *DataVolumeReportRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *DataVolumeReportRequest) GetIECount() int{
    return 0
}
func (self *DataVolumeReport) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *DataVolumeReport) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *DataVolumeReport) createOT() interface{} {
   return &DataVolumeReport{}
}
func (self *DataVolumeReport) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *DataVolumeReport) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'Reset', 'SUCCESSFUL OUTCOME': 'ResetAcknowledge', 'PROCEDURE CODE': 'id-Reset', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idReset}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&Reset{}, SUCCESSFULOUTCOME:&ResetAcknowledge{}, PROCEDURECODE:ProcedureCode{idReset}, CRITICALITY:Criticality{Criticalityreject}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RAB-ReleaseRequest', 'PROCEDURE CODE': 'id-RAB-ReleaseRequest', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idRABReleaseRequest}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RABReleaseRequest{}, PROCEDURECODE:ProcedureCode{idRABReleaseRequest}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetRABReleaseRequestINITIATINGMESSAGE() (*RABReleaseRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &RABReleaseRequest{}, uint64(idRABReleaseRequest), int(Criticalityignore)
}
func (self *RABReleaseRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RABReleaseRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RABReleaseRequest) createOT() interface{} {
   return &RABReleaseRequest{}
}
func (self *RABReleaseRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RABReleaseRequest) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'Iu-ReleaseRequest', 'PROCEDURE CODE': 'id-Iu-ReleaseRequest', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idIuReleaseRequest}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&IuReleaseRequest{}, PROCEDURECODE:ProcedureCode{idIuReleaseRequest}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetIuReleaseRequestINITIATINGMESSAGE() (*IuReleaseRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &IuReleaseRequest{}, uint64(idIuReleaseRequest), int(Criticalityignore)
}
func (self *IuReleaseRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *IuReleaseRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *IuReleaseRequest) createOT() interface{} {
   return &IuReleaseRequest{}
}
func (self *IuReleaseRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *IuReleaseRequest) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RelocationDetect', 'PROCEDURE CODE': 'id-RelocationDetect', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idRelocationDetect}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RelocationDetect{}, PROCEDURECODE:ProcedureCode{idRelocationDetect}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetRelocationDetectINITIATINGMESSAGE() (*RelocationDetect, uint64, int) {/*TYPE, ID, Cricality*/
 return &RelocationDetect{}, uint64(idRelocationDetect), int(Criticalityignore)
}
func (self *RelocationDetect) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RelocationDetect) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RelocationDetect) createOT() interface{} {
   return &RelocationDetect{}
}
func (self *RelocationDetect) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RelocationDetect) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RelocationComplete', 'PROCEDURE CODE': 'id-RelocationComplete', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idRelocationComplete}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RelocationComplete{}, PROCEDURECODE:ProcedureCode{idRelocationComplete}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'Paging', 'PROCEDURE CODE': 'id-Paging', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idPaging}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&Paging{}, PROCEDURECODE:ProcedureCode{idPaging}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetPagingINITIATINGMESSAGE() (*Paging, uint64, int) {/*TYPE, ID, Cricality*/
 return &Paging{}, uint64(idPaging), int(Criticalityignore)
}
func (self *Paging) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *Paging) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *Paging) createOT() interface{} {
   return &Paging{}
}
func (self *Paging) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *Paging) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'CommonID', 'PROCEDURE CODE': 'id-CommonID', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idCommonID}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&CommonID{}, PROCEDURECODE:ProcedureCode{idCommonID}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetCommonIDINITIATINGMESSAGE() (*CommonID, uint64, int) {/*TYPE, ID, Cricality*/
 return &CommonID{}, uint64(idCommonID), int(Criticalityignore)
}
func (self *CommonID) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *CommonID) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *CommonID) createOT() interface{} {
   return &CommonID{}
}
func (self *CommonID) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *CommonID) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'CN-InvokeTrace', 'PROCEDURE CODE': 'id-CN-InvokeTrace', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idCNInvokeTrace}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&CNInvokeTrace{}, PROCEDURECODE:ProcedureCode{idCNInvokeTrace}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetCNInvokeTraceINITIATINGMESSAGE() (*CNInvokeTrace, uint64, int) {/*TYPE, ID, Cricality*/
 return &CNInvokeTrace{}, uint64(idCNInvokeTrace), int(Criticalityignore)
}
func (self *CNInvokeTrace) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *CNInvokeTrace) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *CNInvokeTrace) createOT() interface{} {
   return &CNInvokeTrace{}
}
func (self *CNInvokeTrace) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *CNInvokeTrace) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'CN-DeactivateTrace', 'PROCEDURE CODE': 'id-CN-DeactivateTrace', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idCNDeactivateTrace}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&CNDeactivateTrace{}, PROCEDURECODE:ProcedureCode{idCNDeactivateTrace}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetCNDeactivateTraceINITIATINGMESSAGE() (*CNDeactivateTrace, uint64, int) {/*TYPE, ID, Cricality*/
 return &CNDeactivateTrace{}, uint64(idCNDeactivateTrace), int(Criticalityignore)
}
func (self *CNDeactivateTrace) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *CNDeactivateTrace) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *CNDeactivateTrace) createOT() interface{} {
   return &CNDeactivateTrace{}
}
func (self *CNDeactivateTrace) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *CNDeactivateTrace) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'LocationReportingControl', 'PROCEDURE CODE': 'id-LocationReportingControl', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idLocationReportingControl}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&LocationReportingControl{}, PROCEDURECODE:ProcedureCode{idLocationReportingControl}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetLocationReportingControlINITIATINGMESSAGE() (*LocationReportingControl, uint64, int) {/*TYPE, ID, Cricality*/
 return &LocationReportingControl{}, uint64(idLocationReportingControl), int(Criticalityignore)
}
func (self *LocationReportingControl) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *LocationReportingControl) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *LocationReportingControl) createOT() interface{} {
   return &LocationReportingControl{}
}
func (self *LocationReportingControl) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *LocationReportingControl) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'LocationReport', 'PROCEDURE CODE': 'id-LocationReport', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idLocationReport}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&LocationReport{}, PROCEDURECODE:ProcedureCode{idLocationReport}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetLocationReportINITIATINGMESSAGE() (*LocationReport, uint64, int) {/*TYPE, ID, Cricality*/
 return &LocationReport{}, uint64(idLocationReport), int(Criticalityignore)
}
func (self *LocationReport) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *LocationReport) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *LocationReport) createOT() interface{} {
   return &LocationReport{}
}
func (self *LocationReport) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *LocationReport) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'InitialUE-Message', 'PROCEDURE CODE': 'id-InitialUE-Message', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idInitialUEMessage}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&InitialUEMessage{}, PROCEDURECODE:ProcedureCode{idInitialUEMessage}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetInitialUEMessageINITIATINGMESSAGE() (*InitialUEMessage, uint64, int) {/*TYPE, ID, Cricality*/
 return &InitialUEMessage{}, uint64(idInitialUEMessage), int(Criticalityignore)
}
func (self *InitialUEMessage) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *InitialUEMessage) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *InitialUEMessage) createOT() interface{} {
   return &InitialUEMessage{}
}
func (self *InitialUEMessage) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *InitialUEMessage) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'DirectTransfer', 'PROCEDURE CODE': 'id-DirectTransfer', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idDirectTransfer}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&DirectTransfer{}, PROCEDURECODE:ProcedureCode{idDirectTransfer}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'Overload', 'PROCEDURE CODE': 'id-OverloadControl', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idOverloadControl}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&Overload{}, PROCEDURECODE:ProcedureCode{idOverloadControl}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetOverloadControlINITIATINGMESSAGE() (*Overload, uint64, int) {/*TYPE, ID, Cricality*/
 return &Overload{}, uint64(idOverloadControl), int(Criticalityignore)
}
func (self *Overload) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *Overload) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *Overload) createOT() interface{} {
   return &Overload{}
}
func (self *Overload) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *Overload) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ErrorIndication', 'PROCEDURE CODE': 'id-ErrorIndication', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idErrorIndication}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ErrorIndication{}, PROCEDURECODE:ProcedureCode{idErrorIndication}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'SRNS-DataForwardCommand', 'PROCEDURE CODE': 'id-SRNS-DataForward', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idSRNSDataForward}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&SRNSDataForwardCommand{}, PROCEDURECODE:ProcedureCode{idSRNSDataForward}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetSRNSDataForwardINITIATINGMESSAGE() (*SRNSDataForwardCommand, uint64, int) {/*TYPE, ID, Cricality*/
 return &SRNSDataForwardCommand{}, uint64(idSRNSDataForward), int(Criticalityignore)
}
func (self *SRNSDataForwardCommand) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *SRNSDataForwardCommand) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *SRNSDataForwardCommand) createOT() interface{} {
   return &SRNSDataForwardCommand{}
}
func (self *SRNSDataForwardCommand) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *SRNSDataForwardCommand) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ForwardSRNS-Context', 'PROCEDURE CODE': 'id-ForwardSRNS-Context', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idForwardSRNSContext}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ForwardSRNSContext{}, PROCEDURECODE:ProcedureCode{idForwardSRNSContext}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetForwardSRNSContextINITIATINGMESSAGE() (*ForwardSRNSContext, uint64, int) {/*TYPE, ID, Cricality*/
 return &ForwardSRNSContext{}, uint64(idForwardSRNSContext), int(Criticalityignore)
}
func (self *ForwardSRNSContext) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ForwardSRNSContext) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ForwardSRNSContext) createOT() interface{} {
   return &ForwardSRNSContext{}
}
func (self *ForwardSRNSContext) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ForwardSRNSContext) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RAB-AssignmentRequest', 'OUTCOME': 'RAB-AssignmentResponse', 'PROCEDURE CODE': 'id-RAB-Assignment', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idRABAssignment}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RABAssignmentRequest{}, OUTCOME:&RABAssignmentResponse{}, PROCEDURECODE:ProcedureCode{idRABAssignment}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetRABAssignmentINITIATINGMESSAGE() (*RABAssignmentRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &RABAssignmentRequest{}, uint64(idRABAssignment), int(Criticalityreject)
}
func GetRABAssignmentOUTCOME() (*RABAssignmentResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &RABAssignmentResponse{}, uint64(idRABAssignment), int(Criticalityreject)
}
func (self *RABAssignmentRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RABAssignmentRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RABAssignmentRequest) createOT() interface{} {
   return &RABAssignmentRequest{}
}
func (self *RABAssignmentRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RABAssignmentRequest) GetIECount() int{
    return 0
}
func (self *RABAssignmentResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RABAssignmentResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RABAssignmentResponse) createOT() interface{} {
   return &RABAssignmentResponse{}
}
func (self *RABAssignmentResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RABAssignmentResponse) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'PrivateMessage', 'PROCEDURE CODE': 'id-privateMessage', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idprivateMessage}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&PrivateMessage{}, PROCEDURECODE:ProcedureCode{idprivateMessage}, CRITICALITY:Criticality{Criticalityignore}, }
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
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'ResetResource', 'SUCCESSFUL OUTCOME': 'ResetResourceAcknowledge', 'PROCEDURE CODE': 'id-ResetResource', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idResetResource}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&ResetResource{}, SUCCESSFULOUTCOME:&ResetResourceAcknowledge{}, PROCEDURECODE:ProcedureCode{idResetResource}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetResetResourceINITIATINGMESSAGE() (*ResetResource, uint64, int) {/*TYPE, ID, Cricality*/
 return &ResetResource{}, uint64(idResetResource), int(Criticalityreject)
}
func GetResetResourceSUCCESSFULOUTCOME() (*ResetResourceAcknowledge, uint64, int) {/*TYPE, ID, Cricality*/
 return &ResetResourceAcknowledge{}, uint64(idResetResource), int(Criticalityreject)
}
func (self *ResetResource) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ResetResource) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ResetResource) createOT() interface{} {
   return &ResetResource{}
}
func (self *ResetResource) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ResetResource) GetIECount() int{
    return 0
}
func (self *ResetResourceAcknowledge) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *ResetResourceAcknowledge) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *ResetResourceAcknowledge) createOT() interface{} {
   return &ResetResourceAcknowledge{}
}
func (self *ResetResourceAcknowledge) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *ResetResourceAcknowledge) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RANAP-RelocationInformation', 'PROCEDURE CODE': 'id-RANAP-Relocation', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idRANAPRelocation}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RANAPRelocationInformation{}, PROCEDURECODE:ProcedureCode{idRANAPRelocation}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetRANAPRelocationINITIATINGMESSAGE() (*RANAPRelocationInformation, uint64, int) {/*TYPE, ID, Cricality*/
 return &RANAPRelocationInformation{}, uint64(idRANAPRelocation), int(Criticalityignore)
}
func (self *RANAPRelocationInformation) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RANAPRelocationInformation) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RANAPRelocationInformation) createOT() interface{} {
   return &RANAPRelocationInformation{}
}
func (self *RANAPRelocationInformation) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RANAPRelocationInformation) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'RAB-ModifyRequest', 'PROCEDURE CODE': 'id-RAB-ModifyRequest', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idRABModifyRequest}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&RABModifyRequest{}, PROCEDURECODE:ProcedureCode{idRABModifyRequest}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetRABModifyRequestINITIATINGMESSAGE() (*RABModifyRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &RABModifyRequest{}, uint64(idRABModifyRequest), int(Criticalityignore)
}
func (self *RABModifyRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *RABModifyRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *RABModifyRequest) createOT() interface{} {
   return &RABModifyRequest{}
}
func (self *RABModifyRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *RABModifyRequest) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'LocationRelatedDataRequest', 'SUCCESSFUL OUTCOME': 'LocationRelatedDataResponse', 'UNSUCCESSFUL OUTCOME': 'LocationRelatedDataFailure', 'PROCEDURE CODE': 'id-LocationRelatedData', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idLocationRelatedData}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&LocationRelatedDataRequest{}, SUCCESSFULOUTCOME:&LocationRelatedDataResponse{}, UNSUCCESSFULOUTCOME:&LocationRelatedDataFailure{}, PROCEDURECODE:ProcedureCode{idLocationRelatedData}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetLocationRelatedDataINITIATINGMESSAGE() (*LocationRelatedDataRequest, uint64, int) {/*TYPE, ID, Cricality*/
 return &LocationRelatedDataRequest{}, uint64(idLocationRelatedData), int(Criticalityreject)
}
func GetLocationRelatedDataSUCCESSFULOUTCOME() (*LocationRelatedDataResponse, uint64, int) {/*TYPE, ID, Cricality*/
 return &LocationRelatedDataResponse{}, uint64(idLocationRelatedData), int(Criticalityreject)
}
func GetLocationRelatedDataUNSUCCESSFULOUTCOME() (*LocationRelatedDataFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &LocationRelatedDataFailure{}, uint64(idLocationRelatedData), int(Criticalityreject)
}
func (self *LocationRelatedDataRequest) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *LocationRelatedDataRequest) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *LocationRelatedDataRequest) createOT() interface{} {
   return &LocationRelatedDataRequest{}
}
func (self *LocationRelatedDataRequest) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *LocationRelatedDataRequest) GetIECount() int{
    return 0
}
func (self *LocationRelatedDataResponse) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *LocationRelatedDataResponse) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *LocationRelatedDataResponse) createOT() interface{} {
   return &LocationRelatedDataResponse{}
}
func (self *LocationRelatedDataResponse) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *LocationRelatedDataResponse) GetIECount() int{
    return 0
}
func (self *LocationRelatedDataFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *LocationRelatedDataFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *LocationRelatedDataFailure) createOT() interface{} {
   return &LocationRelatedDataFailure{}
}
func (self *LocationRelatedDataFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *LocationRelatedDataFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'InformationTransferIndication', 'SUCCESSFUL OUTCOME': 'InformationTransferConfirmation', 'UNSUCCESSFUL OUTCOME': 'InformationTransferFailure', 'PROCEDURE CODE': 'id-InformationTransfer', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idInformationTransfer}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&InformationTransferIndication{}, SUCCESSFULOUTCOME:&InformationTransferConfirmation{}, UNSUCCESSFULOUTCOME:&InformationTransferFailure{}, PROCEDURECODE:ProcedureCode{idInformationTransfer}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetInformationTransferINITIATINGMESSAGE() (*InformationTransferIndication, uint64, int) {/*TYPE, ID, Cricality*/
 return &InformationTransferIndication{}, uint64(idInformationTransfer), int(Criticalityreject)
}
func GetInformationTransferSUCCESSFULOUTCOME() (*InformationTransferConfirmation, uint64, int) {/*TYPE, ID, Cricality*/
 return &InformationTransferConfirmation{}, uint64(idInformationTransfer), int(Criticalityreject)
}
func GetInformationTransferUNSUCCESSFULOUTCOME() (*InformationTransferFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &InformationTransferFailure{}, uint64(idInformationTransfer), int(Criticalityreject)
}
func (self *InformationTransferIndication) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *InformationTransferIndication) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *InformationTransferIndication) createOT() interface{} {
   return &InformationTransferIndication{}
}
func (self *InformationTransferIndication) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *InformationTransferIndication) GetIECount() int{
    return 0
}
func (self *InformationTransferConfirmation) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *InformationTransferConfirmation) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *InformationTransferConfirmation) createOT() interface{} {
   return &InformationTransferConfirmation{}
}
func (self *InformationTransferConfirmation) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *InformationTransferConfirmation) GetIECount() int{
    return 0
}
func (self *InformationTransferFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *InformationTransferFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *InformationTransferFailure) createOT() interface{} {
   return &InformationTransferFailure{}
}
func (self *InformationTransferFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *InformationTransferFailure) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'UESpecificInformationIndication', 'PROCEDURE CODE': 'id-UESpecificInformation', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idUESpecificInformation}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&UESpecificInformationIndication{}, PROCEDURECODE:ProcedureCode{idUESpecificInformation}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetUESpecificInformationINITIATINGMESSAGE() (*UESpecificInformationIndication, uint64, int) {/*TYPE, ID, Cricality*/
 return &UESpecificInformationIndication{}, uint64(idUESpecificInformation), int(Criticalityignore)
}
func (self *UESpecificInformationIndication) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *UESpecificInformationIndication) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *UESpecificInformationIndication) createOT() interface{} {
   return &UESpecificInformationIndication{}
}
func (self *UESpecificInformationIndication) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *UESpecificInformationIndication) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'DirectInformationTransfer', 'PROCEDURE CODE': 'id-DirectInformationTransfer', 'CRITICALITY': 'ignore'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idDirectInformationTransfer}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&DirectInformationTransfer{}, PROCEDURECODE:ProcedureCode{idDirectInformationTransfer}, CRITICALITY:Criticality{Criticalityignore}, }
}
func GetDirectInformationTransferINITIATINGMESSAGE() (*DirectInformationTransfer, uint64, int) {/*TYPE, ID, Cricality*/
 return &DirectInformationTransfer{}, uint64(idDirectInformationTransfer), int(Criticalityignore)
}
func (self *DirectInformationTransfer) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *DirectInformationTransfer) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *DirectInformationTransfer) createOT() interface{} {
   return &DirectInformationTransfer{}
}
func (self *DirectInformationTransfer) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *DirectInformationTransfer) GetIECount() int{
    return 0
}
func init() { // ObjVSet:[{'type': 'ProcedureCode', 'name': '&procedureCode'}] {'class': 'RANAP-ELEMENTARY-PROCEDURE', 'value': [{'INITIATING MESSAGE': 'UplinkInformationTransferIndication', 'SUCCESSFUL OUTCOME': 'UplinkInformationTransferConfirmation', 'UNSUCCESSFUL OUTCOME': 'UplinkInformationTransferFailure', 'PROCEDURE CODE': 'id-UplinkInformationTransfer', 'CRITICALITY': 'reject'}]}
table_RANAPELEMENTARYPROCEDURES[RANAPELEMENTARYPROCEDUREprocedureCode{idUplinkInformationTransfer}] = &RANAPELEMENTARYPROCEDURE{INITIATINGMESSAGE:&UplinkInformationTransferIndication{}, SUCCESSFULOUTCOME:&UplinkInformationTransferConfirmation{}, UNSUCCESSFULOUTCOME:&UplinkInformationTransferFailure{}, PROCEDURECODE:ProcedureCode{idUplinkInformationTransfer}, CRITICALITY:Criticality{Criticalityreject}, }
}
func GetUplinkInformationTransferINITIATINGMESSAGE() (*UplinkInformationTransferIndication, uint64, int) {/*TYPE, ID, Cricality*/
 return &UplinkInformationTransferIndication{}, uint64(idUplinkInformationTransfer), int(Criticalityreject)
}
func GetUplinkInformationTransferSUCCESSFULOUTCOME() (*UplinkInformationTransferConfirmation, uint64, int) {/*TYPE, ID, Cricality*/
 return &UplinkInformationTransferConfirmation{}, uint64(idUplinkInformationTransfer), int(Criticalityreject)
}
func GetUplinkInformationTransferUNSUCCESSFULOUTCOME() (*UplinkInformationTransferFailure, uint64, int) {/*TYPE, ID, Cricality*/
 return &UplinkInformationTransferFailure{}, uint64(idUplinkInformationTransfer), int(Criticalityreject)
}
func (self *UplinkInformationTransferIndication) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *UplinkInformationTransferIndication) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *UplinkInformationTransferIndication) createOT() interface{} {
   return &UplinkInformationTransferIndication{}
}
func (self *UplinkInformationTransferIndication) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *UplinkInformationTransferIndication) GetIECount() int{
    return 0
}
func (self *UplinkInformationTransferConfirmation) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *UplinkInformationTransferConfirmation) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *UplinkInformationTransferConfirmation) createOT() interface{} {
   return &UplinkInformationTransferConfirmation{}
}
func (self *UplinkInformationTransferConfirmation) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *UplinkInformationTransferConfirmation) GetIECount() int{
    return 0
}
func (self *UplinkInformationTransferFailure) UnpackOT(st *Stream, id interface{}) {
    self.Unpack(st)
}
func (self *UplinkInformationTransferFailure) PackOT(st *Stream, id interface{}) {
    self.Pack(st)
}
func (self *UplinkInformationTransferFailure) createOT() interface{} {
   return &UplinkInformationTransferFailure{}
}
func (self *UplinkInformationTransferFailure) GetOT(id interface{}) bool{
    return false//ObjVSet
}
func (self *UplinkInformationTransferFailure) GetIECount() int{
    return 0
}
