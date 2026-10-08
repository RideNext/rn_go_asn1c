
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package e2sm
import (
  log "github.com/sirupsen/logrus"
)
var version = "vrc_v04_00"

func fmte2sm() {log.Debug("e2sm")}
func (self *BeamID)Unpack(stream *Stream) {
    //coptions := []string{"nR-Beam-ID"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in BeamID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.NRBeamID = &NRSSBIndex{}//cho6
        self.NRBeamID.Unpack(stream)
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * BeamID) Pack(stream *Stream) {
    if self.NRBeamID != nil {
        stream.set_choice(0, 0, 1, 1)
        self.NRBeamID.Pack(stream)//2
    }

}
type BeamID struct { //[{'type': 'NR-SSB-Index', 'name': 'nR-Beam-ID'}, None]
    NRBeamID *NRSSBIndex
} // BeamID

type CellRNTI struct { // [{'type': 'RNTI-Value', 'name': 'c-RNTI'}, {'type': 'CGI', 'name': 'cell-Global-ID'}, None]
    CRNTI RNTIValue
    CellGlobalID CGI
}

func (self * CellRNTI) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.CRNTI.Unpack(stream)// p8
    self.CellGlobalID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CellRNTI) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CRNTI.Pack(stream)
    self.CellGlobalID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *CGI)Unpack(stream *Stream) {
    //coptions := []string{"nR-CGI","eUTRA-CGI"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in CGI\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.NRCGI = &NRCGI{}//cho6
        self.NRCGI.Unpack(stream)
    } else if choice == 1 { //ch2
        self.EUTRACGI = &EUTRACGI{}//cho6
        self.EUTRACGI.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * CGI) Pack(stream *Stream) {
    if self.NRCGI != nil {
        stream.set_choice(0, 1, 1, 2)
        self.NRCGI.Pack(stream)//2
    } else if self.EUTRACGI != nil {
        stream.set_choice(1, 1, 1, 2)
        self.EUTRACGI.Pack(stream)//2
    }

}
type CGI struct { //[{'type': 'NR-CGI', 'name': 'nR-CGI'}, {'type': 'EUTRA-CGI', 'name': 'eUTRA-CGI'}, None]
    NRCGI *NRCGI
    EUTRACGI *EUTRACGI
} // CGI

func (self *CoreCPID)Unpack(stream *Stream) {
    //coptions := []string{"fiveGC","ePC"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in CoreCPID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.FiveGC = &GUAMI{}//cho6
        self.FiveGC.Unpack(stream)
    } else if choice == 1 { //ch2
        self.EPC = &GUMMEI{}//cho6
        self.EPC.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * CoreCPID) Pack(stream *Stream) {
    if self.FiveGC != nil {
        stream.set_choice(0, 1, 1, 2)
        self.FiveGC.Pack(stream)//2
    } else if self.EPC != nil {
        stream.set_choice(1, 1, 1, 2)
        self.EPC.Pack(stream)//2
    }

}
type CoreCPID struct { //[{'type': 'GUAMI', 'name': 'fiveGC'}, {'type': 'GUMMEI', 'name': 'ePC'}, None]
    FiveGC *GUAMI
    EPC *GUMMEI
} // CoreCPID

func (self *InterfaceIdentifier)Unpack(stream *Stream) {
    //coptions := []string{"nG","xN","f1","e1","s1","x2","w1","Unknown"}
    choice := stream.get_choice(3, 1, 7)
    choice_len := 0
    choice_loc := 0
    if choice >= 7 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in InterfaceIdentifier\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.NG = &InterfaceIDNG{}//cho6
        self.NG.Unpack(stream)
    } else if choice == 1 { //ch2
        self.XN = &InterfaceIDXn{}//cho6
        self.XN.Unpack(stream)
    } else if choice == 2 { //ch2
        self.F1 = &InterfaceIDF1{}//cho6
        self.F1.Unpack(stream)
    } else if choice == 3 { //ch2
        self.E1 = &InterfaceIDE1{}//cho6
        self.E1.Unpack(stream)
    } else if choice == 4 { //ch2
        self.S1 = &InterfaceIDS1{}//cho6
        self.S1.Unpack(stream)
    } else if choice == 5 { //ch2
        self.X2 = &InterfaceIDX2{}//cho6
        self.X2.Unpack(stream)
    } else if choice == 6 { //ch2
        self.W1 = &InterfaceIDW1{}//cho6
        self.W1.Unpack(stream)
    }//end of if else

    if choice >= 7 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * InterfaceIdentifier) Pack(stream *Stream) {
    if self.NG != nil {
        stream.set_choice(0, 3, 1, 7)
        self.NG.Pack(stream)//2
    } else if self.XN != nil {
        stream.set_choice(1, 3, 1, 7)
        self.XN.Pack(stream)//2
    } else if self.F1 != nil {
        stream.set_choice(2, 3, 1, 7)
        self.F1.Pack(stream)//2
    } else if self.E1 != nil {
        stream.set_choice(3, 3, 1, 7)
        self.E1.Pack(stream)//2
    } else if self.S1 != nil {
        stream.set_choice(4, 3, 1, 7)
        self.S1.Pack(stream)//2
    } else if self.X2 != nil {
        stream.set_choice(5, 3, 1, 7)
        self.X2.Pack(stream)//2
    } else if self.W1 != nil {
        stream.set_choice(6, 3, 1, 7)
        self.W1.Pack(stream)//2
    }

}
type InterfaceIdentifier struct { //[{'type': 'InterfaceID-NG', 'name': 'nG'}, {'type': 'InterfaceID-Xn', 'name': 'xN'}, {'type': 'InterfaceID-F1', 'name': 'f1'}, {'type': 'InterfaceID-E1', 'name': 'e1'}, {'type': 'InterfaceID-S1', 'name': 's1'}, {'type': 'InterfaceID-X2', 'name': 'x2'}, {'type': 'InterfaceID-W1', 'name': 'w1'}, None]
    NG *InterfaceIDNG
    XN *InterfaceIDXn
    F1 *InterfaceIDF1
    E1 *InterfaceIDE1
    S1 *InterfaceIDS1
    X2 *InterfaceIDX2
    W1 *InterfaceIDW1
} // InterfaceIdentifier

type InterfaceIDNG struct { // [{'type': 'GUAMI', 'name': 'guami'}, None]
    Guami GUAMI
}

func (self * InterfaceIDNG) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.Guami.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InterfaceIDNG) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.Guami.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type InterfaceIDXn struct { // [{'type': 'GlobalNGRANNodeID', 'name': 'global-NG-RAN-ID'}, None]
    GlobalNGRANID GlobalNGRANNodeID
}

func (self * InterfaceIDXn) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GlobalNGRANID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InterfaceIDXn) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GlobalNGRANID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type InterfaceIDF1 struct { // [{'type': 'GlobalGNB-ID', 'name': 'globalGNB-ID'}, {'type': 'GNB-DU-ID', 'name': 'gNB-DU-ID'}, None]
    GlobalGNBID GlobalGNBID
    GNBDUID GNBDUID
}

func (self * InterfaceIDF1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GlobalGNBID.Unpack(stream)// p8
    self.GNBDUID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InterfaceIDF1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GlobalGNBID.Pack(stream)
    self.GNBDUID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type InterfaceIDE1 struct { // [{'type': 'GlobalGNB-ID', 'name': 'globalGNB-ID'}, {'type': 'GNB-CU-UP-ID', 'name': 'gNB-CU-UP-ID'}, None]
    GlobalGNBID GlobalGNBID
    GNBCUUPID GNBCUUPID
}

func (self * InterfaceIDE1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GlobalGNBID.Unpack(stream)// p8
    self.GNBCUUPID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InterfaceIDE1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GlobalGNBID.Pack(stream)
    self.GNBCUUPID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type InterfaceIDS1 struct { // [{'type': 'GUMMEI', 'name': 'gUMMEI'}, None]
    GUMMEI GUMMEI
}

func (self * InterfaceIDS1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GUMMEI.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InterfaceIDS1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GUMMEI.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type InterfaceIDX2_NodeType struct { //[{'type': 'GlobalENB-ID', 'name': 'global-eNB-ID'}, {'type': 'GlobalenGNB-ID', 'name': 'global-en-gNB-ID'}, None]
    GlobaleNBID *GlobalENBID
    GlobalengNBID *GlobalenGNBID
} // InterfaceIDX2_NodeType

type InterfaceIDX2 struct { // [{'type': 'CHOICE', 'members': [{'type': 'GlobalENB-ID', 'name': 'global-eNB-ID'}, {'type': 'GlobalenGNB-ID', 'name': 'global-en-gNB-ID'}, None], 'name': 'nodeType'}, None]
    NodeType InterfaceIDX2_NodeType
}

func (self * InterfaceIDX2) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_nodeType = func(stream *Stream, self *InterfaceIDX2_NodeType) {
        //coptions := []string{"global-eNB-ID","global-en-gNB-ID"}
        choice := stream.get_choice(1, 1, 2)
        choice_len := 0
        choice_loc := 0
        if choice >= 2 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in InterfaceIDX2_NodeType\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.GlobaleNBID = &GlobalENBID{}//cho6
            self.GlobaleNBID.Unpack(stream)
        } else if choice == 1 { //ch2
            self.GlobalengNBID = &GlobalenGNBID{}//cho6
            self.GlobalengNBID.Unpack(stream)
        }//end of if else

        if choice >= 2 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_nodeType(stream, &self.NodeType)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InterfaceIDX2) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_nodeType = func(stream *Stream, self InterfaceIDX2_NodeType) {
        if self.GlobaleNBID != nil {
            stream.set_choice(0, 1, 1, 2)
            self.GlobaleNBID.Pack(stream)//2
        } else if self.GlobalengNBID != nil {
            stream.set_choice(1, 1, 1, 2)
            self.GlobalengNBID.Pack(stream)//2
        }

    }
    Pack_nodeType(stream, self.NodeType) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type InterfaceIDW1 struct { // [{'type': 'GlobalNgENB-ID', 'name': 'global-ng-eNB-ID'}, {'type': 'NGENB-DU-ID', 'name': 'ng-eNB-DU-ID'}, None]
    GlobalngeNBID GlobalNgENBID
    NgeNBDUID NGENBDUID
}

func (self * InterfaceIDW1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GlobalngeNBID.Unpack(stream)// p8
    self.NgeNBDUID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InterfaceIDW1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GlobalngeNBID.Pack(stream)
    self.NgeNBDUID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type InterfaceMessageID struct { // [{'type': 'INTEGER', 'name': 'interfaceProcedureID'}, {'type': 'ENUMERATED', 'values': [('initiatingMessage', 0), ('successfulOutcome', 1), ('unsuccessfulOutcome', 2), None], 'name': 'messageType'}, None]
    InterfaceProcedureID INTEGER
    MessageType ENUMERATED
}

func (self * InterfaceMessageID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_interfaceProcedureID = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(0, 0, 0, 0)
    }
    Unpack_interfaceProcedureID(stream, &self.InterfaceProcedureID)// p2
    var Unpack_messageType = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(3, 3, 1)
    }
    Unpack_messageType(stream, &self.MessageType)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InterfaceMessageID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_interfaceProcedureID = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 0, 0, 0, 0)
    }
    Pack_interfaceProcedureID(stream, self.InterfaceProcedureID) //f2
    var Pack_messageType = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 3, 3, 1)
    }
    Pack_messageType(stream, self.MessageType) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type InterfaceType struct {
  Value int
}
const (
    InterfaceTypenG = 0
    InterfaceTypexn = 1
    InterfaceTypef1 = 2
    InterfaceTypee1 = 3
    InterfaceTypes1 = 4
    InterfaceTypex2 = 5
    InterfaceTypew1 = 6

    /* Extensions */
)
func (self *InterfaceType) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 7, 1)
}
func (self *InterfaceType) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 7, 1)
}
func (self *GroupID)Unpack(stream *Stream) {
    //coptions := []string{"fiveGC","ePC"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in GroupID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.FiveGC = &FiveQI{}//cho6
        self.FiveGC.Unpack(stream)
    } else if choice == 1 { //ch2
        self.EPC = &QCI{}//cho6
        self.EPC.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * GroupID) Pack(stream *Stream) {
    if self.FiveGC != nil {
        stream.set_choice(0, 1, 1, 2)
        self.FiveGC.Pack(stream)//2
    } else if self.EPC != nil {
        stream.set_choice(1, 1, 1, 2)
        self.EPC.Pack(stream)//2
    }

}
type GroupID struct { //[{'type': 'FiveQI', 'name': 'fiveGC'}, {'type': 'QCI', 'name': 'ePC'}, None]
    FiveGC *FiveQI
    EPC *QCI
} // GroupID

type PartialUEID struct { // [{'type': 'AMF-UE-NGAP-ID', 'name': 'amf-UE-NGAP-ID', 'optional': True}, {'type': 'GUAMI', 'name': 'guami', 'optional': True}, {'type': 'GNB-CU-UE-F1AP-ID', 'name': 'gNB-CU-UE-F1AP-ID', 'optional': True}, {'type': 'GNB-CU-CP-UE-E1AP-ID', 'name': 'gNB-CU-CP-UE-E1AP-ID', 'optional': True}, {'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}, {'type': 'NG-RANnodeUEXnAPID', 'name': 'm-NG-RAN-UE-XnAP-ID', 'optional': True}, {'type': 'GlobalNGRANNodeID', 'name': 'globalNG-RANNode-ID', 'optional': True}, {'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}, {'type': 'NGENB-CU-UE-W1AP-ID', 'name': 'ng-eNB-CU-UE-W1AP-ID', 'optional': True}, {'type': 'ENB-UE-X2AP-ID', 'name': 'm-eNB-UE-X2AP-ID', 'optional': True}, {'type': 'ENB-UE-X2AP-ID-Extension', 'name': 'm-eNB-UE-X2AP-ID-Extension', 'optional': True}, {'type': 'GlobalENB-ID', 'name': 'globalENB-ID', 'optional': True}, {'type': 'MME-UE-S1AP-ID', 'name': 'mME-UE-S1AP-ID', 'optional': True}, {'type': 'GUMMEI', 'name': 'gUMMEI', 'optional': True}, None]
    AmfUENGAPID *AMFUENGAPID
    Guami *GUAMI
    GNBCUUEF1APID *GNBCUUEF1APID
    GNBCUCPUEE1APID *GNBCUCPUEE1APID
    RanUEID *RANUEID
    MNGRANUEXnAPID *NGRANnodeUEXnAPID
    GlobalNGRANNodeID *GlobalNGRANNodeID
    CellRNTI *CellRNTI
    NgeNBCUUEW1APID *NGENBCUUEW1APID
    MeNBUEX2APID *ENBUEX2APID
    MeNBUEX2APIDExtension *ENBUEX2APIDExtension
    GlobalENBID *GlobalENBID
    MMEUES1APID *MMEUES1APID
    GUMMEI *GUMMEI
}

func (self * PartialUEID) Unpack(stream *Stream) {
    amfUENGAPID_flag := 0x00000002
    guami_flag := 0x00000004
    gNBCUUEF1APID_flag := 0x00000008
    gNBCUCPUEE1APID_flag := 0x00000010
    ranUEID_flag := 0x00000020
    mNGRANUEXnAPID_flag := 0x00000040
    globalNGRANNodeID_flag := 0x00000080
    cellRNTI_flag := 0x00000100
    ngeNBCUUEW1APID_flag := 0x00000200
    meNBUEX2APID_flag := 0x00000400
    meNBUEX2APIDExtension_flag := 0x00000800
    globalENBID_flag := 0x00001000
    mMEUES1APID_flag := 0x00002000
    gUMMEI_flag := 0x00004000
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(15)
    if (amfUENGAPID_flag & _flags) == amfUENGAPID_flag { //cond2
        self.AmfUENGAPID = &AMFUENGAPID{}//7{'type': 'AMF-UE-NGAP-ID', 'name': 'amf-UE-NGAP-ID', 'optional': True}
        self.AmfUENGAPID.Unpack(stream)// p8
    }
    if (guami_flag & _flags) == guami_flag { //cond2
        self.Guami = &GUAMI{}//7{'type': 'GUAMI', 'name': 'guami', 'optional': True}
        self.Guami.Unpack(stream)// p8
    }
    if (gNBCUUEF1APID_flag & _flags) == gNBCUUEF1APID_flag { //cond2
        self.GNBCUUEF1APID = &GNBCUUEF1APID{}//7{'type': 'GNB-CU-UE-F1AP-ID', 'name': 'gNB-CU-UE-F1AP-ID', 'optional': True}
        self.GNBCUUEF1APID.Unpack(stream)// p8
    }
    if (gNBCUCPUEE1APID_flag & _flags) == gNBCUCPUEE1APID_flag { //cond2
        self.GNBCUCPUEE1APID = &GNBCUCPUEE1APID{}//7{'type': 'GNB-CU-CP-UE-E1AP-ID', 'name': 'gNB-CU-CP-UE-E1AP-ID', 'optional': True}
        self.GNBCUCPUEE1APID.Unpack(stream)// p8
    }
    if (ranUEID_flag & _flags) == ranUEID_flag { //cond2
        self.RanUEID = &RANUEID{}//7{'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}
        self.RanUEID.Unpack(stream)// p8
    }
    if (mNGRANUEXnAPID_flag & _flags) == mNGRANUEXnAPID_flag { //cond2
        self.MNGRANUEXnAPID = &NGRANnodeUEXnAPID{}//7{'type': 'NG-RANnodeUEXnAPID', 'name': 'm-NG-RAN-UE-XnAP-ID', 'optional': True}
        self.MNGRANUEXnAPID.Unpack(stream)// p8
    }
    if (globalNGRANNodeID_flag & _flags) == globalNGRANNodeID_flag { //cond2
        self.GlobalNGRANNodeID = &GlobalNGRANNodeID{}//7{'type': 'GlobalNGRANNodeID', 'name': 'globalNG-RANNode-ID', 'optional': True}
        self.GlobalNGRANNodeID.Unpack(stream)// p8
    }
    if (cellRNTI_flag & _flags) == cellRNTI_flag { //cond2
        self.CellRNTI = &CellRNTI{}//7{'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}
        self.CellRNTI.Unpack(stream)// p8
    }
    if (ngeNBCUUEW1APID_flag & _flags) == ngeNBCUUEW1APID_flag { //cond2
        self.NgeNBCUUEW1APID = &NGENBCUUEW1APID{}//7{'type': 'NGENB-CU-UE-W1AP-ID', 'name': 'ng-eNB-CU-UE-W1AP-ID', 'optional': True}
        self.NgeNBCUUEW1APID.Unpack(stream)// p8
    }
    if (meNBUEX2APID_flag & _flags) == meNBUEX2APID_flag { //cond2
        self.MeNBUEX2APID = &ENBUEX2APID{}//7{'type': 'ENB-UE-X2AP-ID', 'name': 'm-eNB-UE-X2AP-ID', 'optional': True}
        self.MeNBUEX2APID.Unpack(stream)// p8
    }
    if (meNBUEX2APIDExtension_flag & _flags) == meNBUEX2APIDExtension_flag { //cond2
        self.MeNBUEX2APIDExtension = &ENBUEX2APIDExtension{}//7{'type': 'ENB-UE-X2AP-ID-Extension', 'name': 'm-eNB-UE-X2AP-ID-Extension', 'optional': True}
        self.MeNBUEX2APIDExtension.Unpack(stream)// p8
    }
    if (globalENBID_flag & _flags) == globalENBID_flag { //cond2
        self.GlobalENBID = &GlobalENBID{}//7{'type': 'GlobalENB-ID', 'name': 'globalENB-ID', 'optional': True}
        self.GlobalENBID.Unpack(stream)// p8
    }
    if (mMEUES1APID_flag & _flags) == mMEUES1APID_flag { //cond2
        self.MMEUES1APID = &MMEUES1APID{}//7{'type': 'MME-UE-S1AP-ID', 'name': 'mME-UE-S1AP-ID', 'optional': True}
        self.MMEUES1APID.Unpack(stream)// p8
    }
    if (gUMMEI_flag & _flags) == gUMMEI_flag { //cond2
        self.GUMMEI = &GUMMEI{}//7{'type': 'GUMMEI', 'name': 'gUMMEI', 'optional': True}
        self.GUMMEI.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PartialUEID) Pack(stream *Stream) {
    const amfUENGAPID_flag uint = 0x00000002
    const guami_flag uint = 0x00000004
    const gNBCUUEF1APID_flag uint = 0x00000008
    const gNBCUCPUEE1APID_flag uint = 0x00000010
    const ranUEID_flag uint = 0x00000020
    const mNGRANUEXnAPID_flag uint = 0x00000040
    const globalNGRANNodeID_flag uint = 0x00000080
    const cellRNTI_flag uint = 0x00000100
    const ngeNBCUUEW1APID_flag uint = 0x00000200
    const meNBUEX2APID_flag uint = 0x00000400
    const meNBUEX2APIDExtension_flag uint = 0x00000800
    const globalENBID_flag uint = 0x00001000
    const mMEUES1APID_flag uint = 0x00002000
    const gUMMEI_flag uint = 0x00004000
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(15)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.AmfUENGAPID != nil { 
        _flags |= amfUENGAPID_flag
        self.AmfUENGAPID.Pack(stream)
    }//end of optional
    if self.Guami != nil { 
        _flags |= guami_flag
        self.Guami.Pack(stream)
    }//end of optional
    if self.GNBCUUEF1APID != nil { 
        _flags |= gNBCUUEF1APID_flag
        self.GNBCUUEF1APID.Pack(stream)
    }//end of optional
    if self.GNBCUCPUEE1APID != nil { 
        _flags |= gNBCUCPUEE1APID_flag
        self.GNBCUCPUEE1APID.Pack(stream)
    }//end of optional
    if self.RanUEID != nil { 
        _flags |= ranUEID_flag
        self.RanUEID.Pack(stream)
    }//end of optional
    if self.MNGRANUEXnAPID != nil { 
        _flags |= mNGRANUEXnAPID_flag
        self.MNGRANUEXnAPID.Pack(stream)
    }//end of optional
    if self.GlobalNGRANNodeID != nil { 
        _flags |= globalNGRANNodeID_flag
        self.GlobalNGRANNodeID.Pack(stream)
    }//end of optional
    if self.CellRNTI != nil { 
        _flags |= cellRNTI_flag
        self.CellRNTI.Pack(stream)
    }//end of optional
    if self.NgeNBCUUEW1APID != nil { 
        _flags |= ngeNBCUUEW1APID_flag
        self.NgeNBCUUEW1APID.Pack(stream)
    }//end of optional
    if self.MeNBUEX2APID != nil { 
        _flags |= meNBUEX2APID_flag
        self.MeNBUEX2APID.Pack(stream)
    }//end of optional
    if self.MeNBUEX2APIDExtension != nil { 
        _flags |= meNBUEX2APIDExtension_flag
        self.MeNBUEX2APIDExtension.Pack(stream)
    }//end of optional
    if self.GlobalENBID != nil { 
        _flags |= globalENBID_flag
        self.GlobalENBID.Pack(stream)
    }//end of optional
    if self.MMEUES1APID != nil { 
        _flags |= mMEUES1APID_flag
        self.MMEUES1APID.Pack(stream)
    }//end of optional
    if self.GUMMEI != nil { 
        _flags |= gUMMEI_flag
        self.GUMMEI.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 15)
}//end

func (self *QoSID)Unpack(stream *Stream) {
    //coptions := []string{"fiveGC","ePC"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in QoSID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.FiveGC = &FiveQI{}//cho6
        self.FiveGC.Unpack(stream)
    } else if choice == 1 { //ch2
        self.EPC = &QCI{}//cho6
        self.EPC.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * QoSID) Pack(stream *Stream) {
    if self.FiveGC != nil {
        stream.set_choice(0, 1, 1, 2)
        self.FiveGC.Pack(stream)//2
    } else if self.EPC != nil {
        stream.set_choice(1, 1, 1, 2)
        self.EPC.Pack(stream)//2
    }

}
type QoSID struct { //[{'type': 'FiveQI', 'name': 'fiveGC'}, {'type': 'QCI', 'name': 'ePC'}, None]
    FiveGC *FiveQI
    EPC *QCI
} // QoSID

type RANfunctionName struct { // [{'type': 'PrintableString', 'size': [(1, 150), None], 'name': 'ranFunction-ShortName'}, {'type': 'PrintableString', 'size': [(1, 1000), None], 'name': 'ranFunction-E2SM-OID'}, {'type': 'PrintableString', 'size': [(1, 150), None], 'name': 'ranFunction-Description'}, {'type': 'INTEGER', 'name': 'ranFunction-Instance', 'optional': True}, None]
    RanFunctionShortName PrintableString
    RanFunctionE2SMOID PrintableString
    RanFunctionDescription PrintableString
    RanFunctionInstance *INTEGER
}

func (self * RANfunctionName) Unpack(stream *Stream) {
    ranFunctionInstance_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_ranFunctionShortName = func(st *Stream, self *PrintableString) {
        st.parse_ext()
        _len := st.parse_olen(8)+1
        if _len < 1 || _len > 150 {
            log.Error ("Invalid len in ranFunction-ShortName")
            return
        }
        self.Value = st.parsef_PriString(_len)
    }
    Unpack_ranFunctionShortName(stream, &self.RanFunctionShortName)// p2
    var Unpack_ranFunctionE2SMOID = func(st *Stream, self *PrintableString) {
        st.parse_ext()
        _len := st.parse_olen(10)+1
        if _len < 1 || _len > 1000 {
            log.Error ("Invalid len in ranFunction-E2SM-OID")
            return
        }
        self.Value = st.parsef_PriString(_len)
    }
    Unpack_ranFunctionE2SMOID(stream, &self.RanFunctionE2SMOID)// p2
    var Unpack_ranFunctionDescription = func(st *Stream, self *PrintableString) {
        st.parse_ext()
        _len := st.parse_olen(8)+1
        if _len < 1 || _len > 150 {
            log.Error ("Invalid len in ranFunction-Description")
            return
        }
        self.Value = st.parsef_PriString(_len)
    }
    Unpack_ranFunctionDescription(stream, &self.RanFunctionDescription)// p2
    if (ranFunctionInstance_flag & _flags) == ranFunctionInstance_flag { //cond1
        var Unpack_ranFunctionInstance = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(0, 0, 0, 0)
        }
        self.RanFunctionInstance = &INTEGER{}//6{'type': 'INTEGER', 'name': 'ranFunction-Instance', 'optional': True}
        Unpack_ranFunctionInstance(stream, self.RanFunctionInstance)// p1 {'type': 'INTEGER', 'name': 'ranFunction-Instance', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANfunctionName) Pack(stream *Stream) {
    const ranFunctionInstance_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ranFunctionShortName = func(st *Stream, self PrintableString) {
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
    Pack_ranFunctionShortName(stream, self.RanFunctionShortName) //f2
    var Pack_ranFunctionE2SMOID = func(st *Stream, self PrintableString) {
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
    Pack_ranFunctionE2SMOID(stream, self.RanFunctionE2SMOID) //f2
    var Pack_ranFunctionDescription = func(st *Stream, self PrintableString) {
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
    Pack_ranFunctionDescription(stream, self.RanFunctionDescription) //f2
    if self.RanFunctionInstance != nil { //YY
        _flags |= ranFunctionInstance_flag
        var Pack_ranFunctionInstance = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 0, 0, 0, 0)
        }
        Pack_ranFunctionInstance(stream, *self.RanFunctionInstance) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RICFormatType struct {
  Value uint64
}
func (self *RICFormatType) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(0, 0, 0, 0)
}
func (self * RICFormatType) Pack(st *Stream){
    st.formatf_Integer(self.Value, 0, 0, 0, 0)
}
type RICStyleType struct {
  Value uint64
}
func (self *RICStyleType) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(0, 0, 0, 0)
}
func (self * RICStyleType) Pack(st *Stream){
    st.formatf_Integer(self.Value, 0, 0, 0, 0)
}
type RICStyleName struct {
  Value string
}
func (self *RICStyleName) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in RIC-Style-Name")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *RICStyleName) Pack(st *Stream) {
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
type RRCMessageID_RrcType struct { //[{'type': 'RRCclass-LTE', 'name': 'lTE'}, {'type': 'RRCclass-NR', 'name': 'nR'}, None]
    LTE *RRCclassLTE
    NR *RRCclassNR
} // RRCMessageID_RrcType

type RRCMessageID struct { // [{'type': 'CHOICE', 'members': [{'type': 'RRCclass-LTE', 'name': 'lTE'}, {'type': 'RRCclass-NR', 'name': 'nR'}, None], 'name': 'rrcType'}, {'type': 'INTEGER', 'name': 'messageID'}, None]
    RrcType RRCMessageID_RrcType
    MessageID INTEGER
}

func (self * RRCMessageID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_rrcType = func(stream *Stream, self *RRCMessageID_RrcType) {
        //coptions := []string{"lTE","nR"}
        choice := stream.get_choice(1, 1, 2)
        choice_len := 0
        choice_loc := 0
        if choice >= 2 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in RRCMessageID_RrcType\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.LTE = &RRCclassLTE{}//cho6
            self.LTE.Unpack(stream)
        } else if choice == 1 { //ch2
            self.NR = &RRCclassNR{}//cho6
            self.NR.Unpack(stream)
        }//end of if else

        if choice >= 2 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_rrcType(stream, &self.RrcType)// p2
    var Unpack_messageID = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(0, 0, 0, 0)
    }
    Unpack_messageID(stream, &self.MessageID)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RRCMessageID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_rrcType = func(stream *Stream, self RRCMessageID_RrcType) {
        if self.LTE != nil {
            stream.set_choice(0, 1, 1, 2)
            self.LTE.Pack(stream)//2
        } else if self.NR != nil {
            stream.set_choice(1, 1, 1, 2)
            self.NR.Pack(stream)//2
        }

    }
    Pack_rrcType(stream, self.RrcType) //f2
    var Pack_messageID = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 0, 0, 0, 0)
    }
    Pack_messageID(stream, self.MessageID) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RRCclassLTE struct {
  Value int
}
const (
    RRCclassLTEbCCH_BCH = 0
    RRCclassLTEbCCH_BCH_MBMS = 1
    RRCclassLTEbCCH_DL_SCH = 2
    RRCclassLTEbCCH_DL_SCH_BR = 3
    RRCclassLTEbCCH_DL_SCH_MBMS = 4
    RRCclassLTEmCCH = 5
    RRCclassLTEpCCH = 6
    RRCclassLTEdL_CCCH = 7
    RRCclassLTEdL_DCCH = 8
    RRCclassLTEuL_CCCH = 9
    RRCclassLTEuL_DCCH = 10
    RRCclassLTEsC_MCCH = 11

    /* Extensions */
)
func (self *RRCclassLTE) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(5, 12, 1)
}
func (self *RRCclassLTE) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 5, 12, 1)
}
type RRCclassNR struct {
  Value int
}
const (
    RRCclassNRbCCH_BCH = 0
    RRCclassNRbCCH_DL_SCH = 1
    RRCclassNRdL_CCCH = 2
    RRCclassNRdL_DCCH = 3
    RRCclassNRpCCH = 4
    RRCclassNRuL_CCCH = 5
    RRCclassNRuL_CCCH1 = 6
    RRCclassNRuL_DCCH = 7

    /* Extensions */
)
func (self *RRCclassNR) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 8, 1)
}
func (self *RRCclassNR) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 8, 1)
}
func (self *ServingCellARFCN)Unpack(stream *Stream) {
    //coptions := []string{"nR","eUTRA"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in ServingCellARFCN\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.NR = &NRARFCN{}//cho6
        self.NR.Unpack(stream)
    } else if choice == 1 { //ch2
        self.EUTRA = &EUTRAARFCN{}//cho6
        self.EUTRA.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ServingCellARFCN) Pack(stream *Stream) {
    if self.NR != nil {
        stream.set_choice(0, 1, 1, 2)
        self.NR.Pack(stream)//2
    } else if self.EUTRA != nil {
        stream.set_choice(1, 1, 1, 2)
        self.EUTRA.Pack(stream)//2
    }

}
type ServingCellARFCN struct { //[{'type': 'NR-ARFCN', 'name': 'nR'}, {'type': 'E-UTRA-ARFCN', 'name': 'eUTRA'}, None]
    NR *NRARFCN
    EUTRA *EUTRAARFCN
} // ServingCellARFCN

func (self *ServingCellPCI)Unpack(stream *Stream) {
    //coptions := []string{"nR","eUTRA"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in ServingCellPCI\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.NR = &NRPCI{}//cho6
        self.NR.Unpack(stream)
    } else if choice == 1 { //ch2
        self.EUTRA = &EUTRAPCI{}//cho6
        self.EUTRA.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ServingCellPCI) Pack(stream *Stream) {
    if self.NR != nil {
        stream.set_choice(0, 1, 1, 2)
        self.NR.Pack(stream)//2
    } else if self.EUTRA != nil {
        stream.set_choice(1, 1, 1, 2)
        self.EUTRA.Pack(stream)//2
    }

}
type ServingCellPCI struct { //[{'type': 'NR-PCI', 'name': 'nR'}, {'type': 'E-UTRA-PCI', 'name': 'eUTRA'}, None]
    NR *NRPCI
    EUTRA *EUTRAPCI
} // ServingCellPCI

func (self *UEID)Unpack(stream *Stream) {
    //coptions := []string{"gNB-UEID","gNB-DU-UEID","gNB-CU-UP-UEID","ng-eNB-UEID","ng-eNB-DU-UEID","en-gNB-UEID","eNB-UEID","Unknown"}
    choice := stream.get_choice(3, 1, 7)
    choice_len := 0
    choice_loc := 0
    if choice >= 7 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in UEID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.GNBUEID = &UEIDGNB{}//cho6
        self.GNBUEID.Unpack(stream)
    } else if choice == 1 { //ch2
        self.GNBDUUEID = &UEIDGNBDU{}//cho6
        self.GNBDUUEID.Unpack(stream)
    } else if choice == 2 { //ch2
        self.GNBCUUPUEID = &UEIDGNBCUUP{}//cho6
        self.GNBCUUPUEID.Unpack(stream)
    } else if choice == 3 { //ch2
        self.NgeNBUEID = &UEIDNGENB{}//cho6
        self.NgeNBUEID.Unpack(stream)
    } else if choice == 4 { //ch2
        self.NgeNBDUUEID = &UEIDNGENBDU{}//cho6
        self.NgeNBDUUEID.Unpack(stream)
    } else if choice == 5 { //ch2
        self.EngNBUEID = &UEIDENGNB{}//cho6
        self.EngNBUEID.Unpack(stream)
    } else if choice == 6 { //ch2
        self.ENBUEID = &UEIDENB{}//cho6
        self.ENBUEID.Unpack(stream)
    }//end of if else

    if choice >= 7 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * UEID) Pack(stream *Stream) {
    if self.GNBUEID != nil {
        stream.set_choice(0, 3, 1, 7)
        self.GNBUEID.Pack(stream)//2
    } else if self.GNBDUUEID != nil {
        stream.set_choice(1, 3, 1, 7)
        self.GNBDUUEID.Pack(stream)//2
    } else if self.GNBCUUPUEID != nil {
        stream.set_choice(2, 3, 1, 7)
        self.GNBCUUPUEID.Pack(stream)//2
    } else if self.NgeNBUEID != nil {
        stream.set_choice(3, 3, 1, 7)
        self.NgeNBUEID.Pack(stream)//2
    } else if self.NgeNBDUUEID != nil {
        stream.set_choice(4, 3, 1, 7)
        self.NgeNBDUUEID.Pack(stream)//2
    } else if self.EngNBUEID != nil {
        stream.set_choice(5, 3, 1, 7)
        self.EngNBUEID.Pack(stream)//2
    } else if self.ENBUEID != nil {
        stream.set_choice(6, 3, 1, 7)
        self.ENBUEID.Pack(stream)//2
    }

}
type UEID struct { //[{'type': 'UEID-GNB', 'name': 'gNB-UEID'}, {'type': 'UEID-GNB-DU', 'name': 'gNB-DU-UEID'}, {'type': 'UEID-GNB-CU-UP', 'name': 'gNB-CU-UP-UEID'}, {'type': 'UEID-NG-ENB', 'name': 'ng-eNB-UEID'}, {'type': 'UEID-NG-ENB-DU', 'name': 'ng-eNB-DU-UEID'}, {'type': 'UEID-EN-GNB', 'name': 'en-gNB-UEID'}, {'type': 'UEID-ENB', 'name': 'eNB-UEID'}, None]
    GNBUEID *UEIDGNB
    GNBDUUEID *UEIDGNBDU
    GNBCUUPUEID *UEIDGNBCUUP
    NgeNBUEID *UEIDNGENB
    NgeNBDUUEID *UEIDNGENBDU
    EngNBUEID *UEIDENGNB
    ENBUEID *UEIDENB
} // UEID

type UEIDGNB struct { // [{'type': 'AMF-UE-NGAP-ID', 'name': 'amf-UE-NGAP-ID'}, {'type': 'GUAMI', 'name': 'guami'}, {'type': 'UEID-GNB-CU-F1AP-ID-List', 'name': 'gNB-CU-UE-F1AP-ID-List', 'optional': True}, {'type': 'UEID-GNB-CU-CP-E1AP-ID-List', 'name': 'gNB-CU-CP-UE-E1AP-ID-List', 'optional': True}, {'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}, {'type': 'NG-RANnodeUEXnAPID', 'name': 'm-NG-RAN-UE-XnAP-ID', 'optional': True}, {'type': 'GlobalGNB-ID', 'name': 'globalGNB-ID', 'optional': True}, None, {'type': 'GlobalNGRANNodeID', 'name': 'globalNG-RANNode-ID', 'optional': True}, {'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}]
    AmfUENGAPID AMFUENGAPID
    Guami GUAMI
    GNBCUUEF1APIDList *UEIDGNBCUF1APIDList
    GNBCUCPUEE1APIDList *UEIDGNBCUCPE1APIDList
    RanUEID *RANUEID
    MNGRANUEXnAPID *NGRANnodeUEXnAPID
    GlobalGNBID *GlobalGNBID
    GlobalNGRANNodeID *GlobalNGRANNodeID
    CellRNTI *CellRNTI
}

func (self * UEIDGNB) Unpack(stream *Stream) {
    gNBCUUEF1APIDList_flag := 0x00000002
    gNBCUCPUEE1APIDList_flag := 0x00000004
    ranUEID_flag := 0x00000008
    mNGRANUEXnAPID_flag := 0x00000010
    globalGNBID_flag := 0x00000020
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(6)
    self.AmfUENGAPID.Unpack(stream)// p8
    self.Guami.Unpack(stream)// p8
    if (gNBCUUEF1APIDList_flag & _flags) == gNBCUUEF1APIDList_flag { //cond2
        self.GNBCUUEF1APIDList = &UEIDGNBCUF1APIDList{}//7{'type': 'UEID-GNB-CU-F1AP-ID-List', 'name': 'gNB-CU-UE-F1AP-ID-List', 'optional': True}
        self.GNBCUUEF1APIDList.Unpack(stream)// p8
    }
    if (gNBCUCPUEE1APIDList_flag & _flags) == gNBCUCPUEE1APIDList_flag { //cond2
        self.GNBCUCPUEE1APIDList = &UEIDGNBCUCPE1APIDList{}//7{'type': 'UEID-GNB-CU-CP-E1AP-ID-List', 'name': 'gNB-CU-CP-UE-E1AP-ID-List', 'optional': True}
        self.GNBCUCPUEE1APIDList.Unpack(stream)// p8
    }
    if (ranUEID_flag & _flags) == ranUEID_flag { //cond2
        self.RanUEID = &RANUEID{}//7{'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}
        self.RanUEID.Unpack(stream)// p8
    }
    if (mNGRANUEXnAPID_flag & _flags) == mNGRANUEXnAPID_flag { //cond2
        self.MNGRANUEXnAPID = &NGRANnodeUEXnAPID{}//7{'type': 'NG-RANnodeUEXnAPID', 'name': 'm-NG-RAN-UE-XnAP-ID', 'optional': True}
        self.MNGRANUEXnAPID.Unpack(stream)// p8
    }
    if (globalGNBID_flag & _flags) == globalGNBID_flag { //cond2
        self.GlobalGNBID = &GlobalGNBID{}//7{'type': 'GlobalGNB-ID', 'name': 'globalGNB-ID', 'optional': True}
        self.GlobalGNBID.Unpack(stream)// p8
    }
    if (_flags & ext_flag) > 0 {  //globalNG-RANNode-ID
        _ecount, _extflags = stream.parsef_sequence_ext(2)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.GlobalNGRANNodeID = &GlobalNGRANNodeID{}//7{'type': 'GlobalNGRANNodeID', 'name': 'globalNG-RANNode-ID', 'optional': True}
        self.GlobalNGRANNodeID.Unpack(stream)// p8
    }
    if (_extflags & (1<< 1)) > 0 {
        stream.get_openlen()
        self.CellRNTI = &CellRNTI{}//7{'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}
        self.CellRNTI.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDGNB) Pack(stream *Stream) {
    const gNBCUUEF1APIDList_flag uint = 0x00000002
    const gNBCUCPUEE1APIDList_flag uint = 0x00000004
    const ranUEID_flag uint = 0x00000008
    const mNGRANUEXnAPID_flag uint = 0x00000010
    const globalGNBID_flag uint = 0x00000020
    const ext_flag int = 0x00000001
    const globalNGRANNodeID_flag uint = 0x00000001
    const cellRNTI_flag uint = 0x00000002
    _flagReserve := stream.reserve_flags(6)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.AmfUENGAPID.Pack(stream)
    self.Guami.Pack(stream)
    if self.GNBCUUEF1APIDList != nil { 
        _flags |= gNBCUUEF1APIDList_flag
        self.GNBCUUEF1APIDList.Pack(stream)
    }//end of optional
    if self.GNBCUCPUEE1APIDList != nil { 
        _flags |= gNBCUCPUEE1APIDList_flag
        self.GNBCUCPUEE1APIDList.Pack(stream)
    }//end of optional
    if self.RanUEID != nil { 
        _flags |= ranUEID_flag
        self.RanUEID.Pack(stream)
    }//end of optional
    if self.MNGRANUEXnAPID != nil { 
        _flags |= mNGRANUEXnAPID_flag
        self.MNGRANUEXnAPID.Pack(stream)
    }//end of optional
    if self.GlobalGNBID != nil { 
        _flags |= globalGNBID_flag
        self.GlobalGNBID.Pack(stream)
    }//end of optional
    if self.GlobalNGRANNodeID != nil { 
        //extension Group: globalNG-RANNode-ID 2 
        if _extPresent == false {
            stream.set_bits(1, 7)
            _extReserve = stream.reserve_flags(2)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= globalNGRANNodeID_flag
        self.GlobalNGRANNodeID.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if self.CellRNTI != nil { 
        //extension Group: cell-RNTI 2 
        if _extPresent == false {
            stream.set_bits(1, 7)
            _extReserve = stream.reserve_flags(2)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= cellRNTI_flag
        self.CellRNTI.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 2) }
    stream.set_flags(_flags, _flagReserve, 8)
}//end

func (self *UEIDGNBCUCPE1APIDList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65535)
    _size += 1
    self.Items = make([]UEIDGNBCUCPE1APIDItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *UEIDGNBCUCPE1APIDList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65535)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type UEIDGNBCUCPE1APIDList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'UEID-GNB-CU-CP-E1AP-ID-Item'}, 'size': [(1, 'maxE1APid')]}
    Items []UEIDGNBCUCPE1APIDItem
}

type UEIDGNBCUCPE1APIDItem struct { // [{'type': 'GNB-CU-CP-UE-E1AP-ID', 'name': 'gNB-CU-CP-UE-E1AP-ID'}, None]
    GNBCUCPUEE1APID GNBCUCPUEE1APID
}

func (self * UEIDGNBCUCPE1APIDItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GNBCUCPUEE1APID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDGNBCUCPE1APIDItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GNBCUCPUEE1APID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *UEIDGNBCUF1APIDList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(4)
    _size += 1
    self.Items = make([]UEIDGNBCUCPF1APIDItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *UEIDGNBCUF1APIDList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 4)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type UEIDGNBCUF1APIDList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'UEID-GNB-CU-CP-F1AP-ID-Item'}, 'size': [(1, 'maxF1APid')]}
    Items []UEIDGNBCUCPF1APIDItem
}

type UEIDGNBCUCPF1APIDItem struct { // [{'type': 'GNB-CU-UE-F1AP-ID', 'name': 'gNB-CU-UE-F1AP-ID'}, None]
    GNBCUUEF1APID GNBCUUEF1APID
}

func (self * UEIDGNBCUCPF1APIDItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GNBCUUEF1APID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDGNBCUCPF1APIDItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GNBCUUEF1APID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type UEIDGNBDU struct { // [{'type': 'GNB-CU-UE-F1AP-ID', 'name': 'gNB-CU-UE-F1AP-ID'}, {'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}, None, {'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}]
    GNBCUUEF1APID GNBCUUEF1APID
    RanUEID *RANUEID
    CellRNTI *CellRNTI
}

func (self * UEIDGNBDU) Unpack(stream *Stream) {
    ranUEID_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.GNBCUUEF1APID.Unpack(stream)// p8
    if (ranUEID_flag & _flags) == ranUEID_flag { //cond2
        self.RanUEID = &RANUEID{}//7{'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}
        self.RanUEID.Unpack(stream)// p8
    }
    if (_flags & ext_flag) > 0 {  //cell-RNTI
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.CellRNTI = &CellRNTI{}//7{'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}
        self.CellRNTI.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDGNBDU) Pack(stream *Stream) {
    const ranUEID_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    const cellRNTI_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GNBCUUEF1APID.Pack(stream)
    if self.RanUEID != nil { 
        _flags |= ranUEID_flag
        self.RanUEID.Pack(stream)
    }//end of optional
    if self.CellRNTI != nil { 
        //extension Group: cell-RNTI 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= cellRNTI_flag
        self.CellRNTI.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type UEIDGNBCUUP struct { // [{'type': 'GNB-CU-CP-UE-E1AP-ID', 'name': 'gNB-CU-CP-UE-E1AP-ID'}, {'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}, None]
    GNBCUCPUEE1APID GNBCUCPUEE1APID
    RanUEID *RANUEID
}

func (self * UEIDGNBCUUP) Unpack(stream *Stream) {
    ranUEID_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.GNBCUCPUEE1APID.Unpack(stream)// p8
    if (ranUEID_flag & _flags) == ranUEID_flag { //cond2
        self.RanUEID = &RANUEID{}//7{'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}
        self.RanUEID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDGNBCUUP) Pack(stream *Stream) {
    const ranUEID_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.GNBCUCPUEE1APID.Pack(stream)
    if self.RanUEID != nil { 
        _flags |= ranUEID_flag
        self.RanUEID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UEIDNGENB struct { // [{'type': 'AMF-UE-NGAP-ID', 'name': 'amf-UE-NGAP-ID'}, {'type': 'GUAMI', 'name': 'guami'}, {'type': 'NGENB-CU-UE-W1AP-ID', 'name': 'ng-eNB-CU-UE-W1AP-ID', 'optional': True}, {'type': 'NG-RANnodeUEXnAPID', 'name': 'm-NG-RAN-UE-XnAP-ID', 'optional': True}, {'type': 'GlobalNgENB-ID', 'name': 'globalNgENB-ID', 'optional': True}, None, {'type': 'GlobalNGRANNodeID', 'name': 'globalNG-RANNode-ID', 'optional': True}, {'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}]
    AmfUENGAPID AMFUENGAPID
    Guami GUAMI
    NgeNBCUUEW1APID *NGENBCUUEW1APID
    MNGRANUEXnAPID *NGRANnodeUEXnAPID
    GlobalNgENBID *GlobalNgENBID
    GlobalNGRANNodeID *GlobalNGRANNodeID
    CellRNTI *CellRNTI
}

func (self * UEIDNGENB) Unpack(stream *Stream) {
    ngeNBCUUEW1APID_flag := 0x00000002
    mNGRANUEXnAPID_flag := 0x00000004
    globalNgENBID_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.AmfUENGAPID.Unpack(stream)// p8
    self.Guami.Unpack(stream)// p8
    if (ngeNBCUUEW1APID_flag & _flags) == ngeNBCUUEW1APID_flag { //cond2
        self.NgeNBCUUEW1APID = &NGENBCUUEW1APID{}//7{'type': 'NGENB-CU-UE-W1AP-ID', 'name': 'ng-eNB-CU-UE-W1AP-ID', 'optional': True}
        self.NgeNBCUUEW1APID.Unpack(stream)// p8
    }
    if (mNGRANUEXnAPID_flag & _flags) == mNGRANUEXnAPID_flag { //cond2
        self.MNGRANUEXnAPID = &NGRANnodeUEXnAPID{}//7{'type': 'NG-RANnodeUEXnAPID', 'name': 'm-NG-RAN-UE-XnAP-ID', 'optional': True}
        self.MNGRANUEXnAPID.Unpack(stream)// p8
    }
    if (globalNgENBID_flag & _flags) == globalNgENBID_flag { //cond2
        self.GlobalNgENBID = &GlobalNgENBID{}//7{'type': 'GlobalNgENB-ID', 'name': 'globalNgENB-ID', 'optional': True}
        self.GlobalNgENBID.Unpack(stream)// p8
    }
    if (_flags & ext_flag) > 0 {  //globalNG-RANNode-ID
        _ecount, _extflags = stream.parsef_sequence_ext(2)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.GlobalNGRANNodeID = &GlobalNGRANNodeID{}//7{'type': 'GlobalNGRANNodeID', 'name': 'globalNG-RANNode-ID', 'optional': True}
        self.GlobalNGRANNodeID.Unpack(stream)// p8
    }
    if (_extflags & (1<< 1)) > 0 {
        stream.get_openlen()
        self.CellRNTI = &CellRNTI{}//7{'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}
        self.CellRNTI.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDNGENB) Pack(stream *Stream) {
    const ngeNBCUUEW1APID_flag uint = 0x00000002
    const mNGRANUEXnAPID_flag uint = 0x00000004
    const globalNgENBID_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    const globalNGRANNodeID_flag uint = 0x00000001
    const cellRNTI_flag uint = 0x00000002
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.AmfUENGAPID.Pack(stream)
    self.Guami.Pack(stream)
    if self.NgeNBCUUEW1APID != nil { 
        _flags |= ngeNBCUUEW1APID_flag
        self.NgeNBCUUEW1APID.Pack(stream)
    }//end of optional
    if self.MNGRANUEXnAPID != nil { 
        _flags |= mNGRANUEXnAPID_flag
        self.MNGRANUEXnAPID.Pack(stream)
    }//end of optional
    if self.GlobalNgENBID != nil { 
        _flags |= globalNgENBID_flag
        self.GlobalNgENBID.Pack(stream)
    }//end of optional
    if self.GlobalNGRANNodeID != nil { 
        //extension Group: globalNG-RANNode-ID 2 
        if _extPresent == false {
            stream.set_bits(1, 7)
            _extReserve = stream.reserve_flags(2)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= globalNGRANNodeID_flag
        self.GlobalNGRANNodeID.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if self.CellRNTI != nil { 
        //extension Group: cell-RNTI 2 
        if _extPresent == false {
            stream.set_bits(1, 7)
            _extReserve = stream.reserve_flags(2)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= cellRNTI_flag
        self.CellRNTI.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 2) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

type UEIDNGENBDU struct { // [{'type': 'NGENB-CU-UE-W1AP-ID', 'name': 'ng-eNB-CU-UE-W1AP-ID'}, None, {'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}]
    NgeNBCUUEW1APID NGENBCUUEW1APID
    CellRNTI *CellRNTI
}

func (self * UEIDNGENBDU) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.NgeNBCUUEW1APID.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //cell-RNTI
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.CellRNTI = &CellRNTI{}//7{'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}
        self.CellRNTI.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDNGENBDU) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const cellRNTI_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.NgeNBCUUEW1APID.Pack(stream)
    if self.CellRNTI != nil { 
        //extension Group: cell-RNTI 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= cellRNTI_flag
        self.CellRNTI.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UEIDENGNB struct { // [{'type': 'ENB-UE-X2AP-ID', 'name': 'm-eNB-UE-X2AP-ID'}, {'type': 'ENB-UE-X2AP-ID-Extension', 'name': 'm-eNB-UE-X2AP-ID-Extension', 'optional': True}, {'type': 'GlobalENB-ID', 'name': 'globalENB-ID'}, {'type': 'GNB-CU-UE-F1AP-ID', 'name': 'gNB-CU-UE-F1AP-ID', 'optional': True}, {'type': 'UEID-GNB-CU-CP-E1AP-ID-List', 'name': 'gNB-CU-CP-UE-E1AP-ID-List', 'optional': True}, {'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}, None, {'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}]
    MeNBUEX2APID ENBUEX2APID
    MeNBUEX2APIDExtension *ENBUEX2APIDExtension
    GlobalENBID GlobalENBID
    GNBCUUEF1APID *GNBCUUEF1APID
    GNBCUCPUEE1APIDList *UEIDGNBCUCPE1APIDList
    RanUEID *RANUEID
    CellRNTI *CellRNTI
}

func (self * UEIDENGNB) Unpack(stream *Stream) {
    meNBUEX2APIDExtension_flag := 0x00000002
    gNBCUUEF1APID_flag := 0x00000004
    gNBCUCPUEE1APIDList_flag := 0x00000008
    ranUEID_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.MeNBUEX2APID.Unpack(stream)// p8
    if (meNBUEX2APIDExtension_flag & _flags) == meNBUEX2APIDExtension_flag { //cond2
        self.MeNBUEX2APIDExtension = &ENBUEX2APIDExtension{}//7{'type': 'ENB-UE-X2AP-ID-Extension', 'name': 'm-eNB-UE-X2AP-ID-Extension', 'optional': True}
        self.MeNBUEX2APIDExtension.Unpack(stream)// p8
    }
    self.GlobalENBID.Unpack(stream)// p8
    if (gNBCUUEF1APID_flag & _flags) == gNBCUUEF1APID_flag { //cond2
        self.GNBCUUEF1APID = &GNBCUUEF1APID{}//7{'type': 'GNB-CU-UE-F1AP-ID', 'name': 'gNB-CU-UE-F1AP-ID', 'optional': True}
        self.GNBCUUEF1APID.Unpack(stream)// p8
    }
    if (gNBCUCPUEE1APIDList_flag & _flags) == gNBCUCPUEE1APIDList_flag { //cond2
        self.GNBCUCPUEE1APIDList = &UEIDGNBCUCPE1APIDList{}//7{'type': 'UEID-GNB-CU-CP-E1AP-ID-List', 'name': 'gNB-CU-CP-UE-E1AP-ID-List', 'optional': True}
        self.GNBCUCPUEE1APIDList.Unpack(stream)// p8
    }
    if (ranUEID_flag & _flags) == ranUEID_flag { //cond2
        self.RanUEID = &RANUEID{}//7{'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}
        self.RanUEID.Unpack(stream)// p8
    }
    if (_flags & ext_flag) > 0 {  //cell-RNTI
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.CellRNTI = &CellRNTI{}//7{'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}
        self.CellRNTI.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDENGNB) Pack(stream *Stream) {
    const meNBUEX2APIDExtension_flag uint = 0x00000002
    const gNBCUUEF1APID_flag uint = 0x00000004
    const gNBCUCPUEE1APIDList_flag uint = 0x00000008
    const ranUEID_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    const cellRNTI_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeNBUEX2APID.Pack(stream)
    if self.MeNBUEX2APIDExtension != nil { 
        _flags |= meNBUEX2APIDExtension_flag
        self.MeNBUEX2APIDExtension.Pack(stream)
    }//end of optional
    self.GlobalENBID.Pack(stream)
    if self.GNBCUUEF1APID != nil { 
        _flags |= gNBCUUEF1APID_flag
        self.GNBCUUEF1APID.Pack(stream)
    }//end of optional
    if self.GNBCUCPUEE1APIDList != nil { 
        _flags |= gNBCUCPUEE1APIDList_flag
        self.GNBCUCPUEE1APIDList.Pack(stream)
    }//end of optional
    if self.RanUEID != nil { 
        _flags |= ranUEID_flag
        self.RanUEID.Pack(stream)
    }//end of optional
    if self.CellRNTI != nil { 
        //extension Group: cell-RNTI 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= cellRNTI_flag
        self.CellRNTI.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 6)
}//end

type UEIDENB struct { // [{'type': 'MME-UE-S1AP-ID', 'name': 'mME-UE-S1AP-ID'}, {'type': 'GUMMEI', 'name': 'gUMMEI'}, {'type': 'ENB-UE-X2AP-ID', 'name': 'm-eNB-UE-X2AP-ID', 'optional': True}, {'type': 'ENB-UE-X2AP-ID-Extension', 'name': 'm-eNB-UE-X2AP-ID-Extension', 'optional': True}, {'type': 'GlobalENB-ID', 'name': 'globalENB-ID', 'optional': True}, None, {'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}]
    MMEUES1APID MMEUES1APID
    GUMMEI GUMMEI
    MeNBUEX2APID *ENBUEX2APID
    MeNBUEX2APIDExtension *ENBUEX2APIDExtension
    GlobalENBID *GlobalENBID
    CellRNTI *CellRNTI
}

func (self * UEIDENB) Unpack(stream *Stream) {
    meNBUEX2APID_flag := 0x00000002
    meNBUEX2APIDExtension_flag := 0x00000004
    globalENBID_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.MMEUES1APID.Unpack(stream)// p8
    self.GUMMEI.Unpack(stream)// p8
    if (meNBUEX2APID_flag & _flags) == meNBUEX2APID_flag { //cond2
        self.MeNBUEX2APID = &ENBUEX2APID{}//7{'type': 'ENB-UE-X2AP-ID', 'name': 'm-eNB-UE-X2AP-ID', 'optional': True}
        self.MeNBUEX2APID.Unpack(stream)// p8
    }
    if (meNBUEX2APIDExtension_flag & _flags) == meNBUEX2APIDExtension_flag { //cond2
        self.MeNBUEX2APIDExtension = &ENBUEX2APIDExtension{}//7{'type': 'ENB-UE-X2AP-ID-Extension', 'name': 'm-eNB-UE-X2AP-ID-Extension', 'optional': True}
        self.MeNBUEX2APIDExtension.Unpack(stream)// p8
    }
    if (globalENBID_flag & _flags) == globalENBID_flag { //cond2
        self.GlobalENBID = &GlobalENBID{}//7{'type': 'GlobalENB-ID', 'name': 'globalENB-ID', 'optional': True}
        self.GlobalENBID.Unpack(stream)// p8
    }
    if (_flags & ext_flag) > 0 {  //cell-RNTI
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.CellRNTI = &CellRNTI{}//7{'type': 'Cell-RNTI', 'name': 'cell-RNTI', 'optional': True}
        self.CellRNTI.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDENB) Pack(stream *Stream) {
    const meNBUEX2APID_flag uint = 0x00000002
    const meNBUEX2APIDExtension_flag uint = 0x00000004
    const globalENBID_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    const cellRNTI_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MMEUES1APID.Pack(stream)
    self.GUMMEI.Pack(stream)
    if self.MeNBUEX2APID != nil { 
        _flags |= meNBUEX2APID_flag
        self.MeNBUEX2APID.Pack(stream)
    }//end of optional
    if self.MeNBUEX2APIDExtension != nil { 
        _flags |= meNBUEX2APIDExtension_flag
        self.MeNBUEX2APIDExtension.Pack(stream)
    }//end of optional
    if self.GlobalENBID != nil { 
        _flags |= globalENBID_flag
        self.GlobalENBID.Pack(stream)
    }//end of optional
    if self.CellRNTI != nil { 
        //extension Group: cell-RNTI 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= cellRNTI_flag
        self.CellRNTI.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 5)
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

type GlobalENBID struct { // [{'type': 'PLMNIdentity', 'name': 'pLMNIdentity'}, {'type': 'ENB-ID', 'name': 'eNB-ID'}, None]
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

type GUMMEI struct { // [{'type': 'PLMNIdentity', 'name': 'pLMN-Identity'}, {'type': 'MME-Group-ID', 'name': 'mME-Group-ID'}, {'type': 'MME-Code', 'name': 'mME-Code'}, None]
    PLMNIdentity PLMNIdentity
    MMEGroupID MMEGroupID
    MMECode MMECode
}

func (self * GUMMEI) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PLMNIdentity.Unpack(stream)// p8
    self.MMEGroupID.Unpack(stream)// p8
    self.MMECode.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GUMMEI) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.MMEGroupID.Pack(stream)
    self.MMECode.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type MMEGroupID struct {
  Value HexBytes
}
func (self *MMEGroupID) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(2)
}
func (self *MMEGroupID) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 2)
}
type MMECode struct {
  Value HexBytes
}
func (self *MMECode) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(1)
}
func (self *MMECode) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 1)
}
type MMEUES1APID struct {
  Value uint64
}
func (self *MMEUES1APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967296, 32, 0, 0)
}
func (self * MMEUES1APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967296, 32, 0, 0)
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
type SubscriberProfileIDforRFP struct {
  Value uint64
}
func (self *SubscriberProfileIDforRFP) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 1)
}
func (self * SubscriberProfileIDforRFP) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 1)
}
func (self *ENGNBID)Unpack(stream *Stream) {
    //coptions := []string{"en-gNB-ID"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in ENGNBID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_engNBID = func(st *Stream, self *BITSTRING){
            self.Len = int(st.parse_blen(4, 0)+22)
            self.Value = st.parsef_BitString(32, int(self.Len))
        }
        self.EngNBID = &BITSTRING{}//cho5
        Unpack_engNBID(stream, self.EngNBID);
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * ENGNBID) Pack(stream *Stream) {
    if self.EngNBID != nil {
        stream.set_choice(0, 0, 1, 1)
        var Pack_engNBID = func(st *Stream, self BITSTRING) {
            st.format_blen(int(self.Len-22), 4, 0)
            st.formatf_BitString(self.Value, int(self.Len))
        }
        Pack_engNBID(stream, *self.EngNBID)//3
    }

}
type ENGNBID struct { //[{'type': 'BIT STRING', 'size': [(22, 32)], 'name': 'en-gNB-ID'}, None]
    EngNBID *BITSTRING
} // ENGNBID

type ENBUEX2APID struct {
  Value uint64
}
func (self *ENBUEX2APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4096, 12, 0, 0)
}
func (self * ENBUEX2APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4096, 12, 0, 0)
}
type ENBUEX2APIDExtension struct {
  Value uint64
}
func (self *ENBUEX2APIDExtension) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4096, 13, 1, 0)
}
func (self * ENBUEX2APIDExtension) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4096, 13, 1, 0)
}
type EUTRAARFCN struct {
  Value uint64
}
func (self *EUTRAARFCN) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * EUTRAARFCN) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type EUTRAPCI struct {
  Value uint64
}
func (self *EUTRAPCI) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(504, 10, 1, 0)
}
func (self * EUTRAPCI) Pack(st *Stream){
    st.formatf_Integer(self.Value, 504, 10, 1, 0)
}
type EUTRATAC struct {
  Value HexBytes
}
func (self *EUTRATAC) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(2)
}
func (self *EUTRATAC) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 2)
}
type GlobalenGNBID struct { // [{'type': 'PLMNIdentity', 'name': 'pLMN-Identity'}, {'type': 'EN-GNB-ID', 'name': 'en-gNB-ID'}, None]
    PLMNIdentity PLMNIdentity
    EngNBID ENGNBID
}

func (self * GlobalenGNBID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PLMNIdentity.Unpack(stream)// p8
    self.EngNBID.Unpack(stream)// p8
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
    self.EngNBID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type NGENBCUUEW1APID struct {
  Value uint64
}
func (self *NGENBCUUEW1APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967296, 32, 0, 0)
}
func (self * NGENBCUUEW1APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967296, 32, 0, 0)
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
type NRSSBIndex struct {
  Value uint64
}
func (self *NRSSBIndex) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(64, 6, 0, 0)
}
func (self * NRSSBIndex) Pack(st *Stream){
    st.formatf_Integer(self.Value, 64, 6, 0, 0)
}
type RNTIValue struct {
  Value uint64
}
func (self *RNTIValue) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 16, 0, 0)
}
func (self * RNTIValue) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 16, 0, 0)
}
type AMFPointer struct {
  Len int
  Value HexBytes
}
func (self *AMFPointer) Unpack(st *Stream){
    self.Value = st.parsef_BitString(6, 6)
}
func (self *AMFPointer) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 6)
}
type AMFRegionID struct {
  Len int
  Value HexBytes
}
func (self *AMFRegionID) Unpack(st *Stream){
    self.Value = st.parsef_BitString(8, 8)
}
func (self *AMFRegionID) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 8)
}
type AMFSetID struct {
  Len int
  Value HexBytes
}
func (self *AMFSetID) Unpack(st *Stream){
    self.Value = st.parsef_BitString(10, 10)
}
func (self *AMFSetID) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 10)
}
type AMFUENGAPID struct {
  Value uint64
}
func (self *AMFUENGAPID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(1099511627776, 40, 0, 0)
}
func (self * AMFUENGAPID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 1099511627776, 40, 0, 0)
}
type EUTRACellIdentity struct {
  Len int
  Value HexBytes
}
func (self *EUTRACellIdentity) Unpack(st *Stream){
    self.Value = st.parsef_BitString(28, 28)
}
func (self *EUTRACellIdentity) Pack(st *Stream) {
    st.formatf_BitString(self.Value, 28)
}
type EUTRACGI struct { // [{'type': 'PLMNIdentity', 'name': 'pLMNIdentity'}, {'type': 'EUTRACellIdentity', 'name': 'eUTRACellIdentity'}, None]
    PLMNIdentity PLMNIdentity
    EUTRACellIdentity EUTRACellIdentity
}

func (self * EUTRACGI) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PLMNIdentity.Unpack(stream)// p8
    self.EUTRACellIdentity.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EUTRACGI) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.EUTRACellIdentity.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type FiveQI struct {
  Value uint64
}
func (self *FiveQI) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 9, 1, 0)
}
func (self * FiveQI) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 9, 1, 0)
}
type GlobalGNBID struct { // [{'type': 'PLMNIdentity', 'name': 'pLMNIdentity'}, {'type': 'GNB-ID', 'name': 'gNB-ID'}, None]
    PLMNIdentity PLMNIdentity
    GNBID GNBID
}

func (self * GlobalGNBID) Unpack(stream *Stream) {
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

func (self * GlobalGNBID) Pack(stream *Stream) {
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

type GlobalNgENBID struct { // [{'type': 'PLMNIdentity', 'name': 'pLMNIdentity'}, {'type': 'NgENB-ID', 'name': 'ngENB-ID'}, None]
    PLMNIdentity PLMNIdentity
    NgENBID NgENBID
}

func (self * GlobalNgENBID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PLMNIdentity.Unpack(stream)// p8
    self.NgENBID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GlobalNgENBID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.NgENBID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *GNBID)Unpack(stream *Stream) {
    //coptions := []string{"gNB-ID"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in GNBID\n", choice, choice_len)
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
func (self * GNBID) Pack(stream *Stream) {
    if self.GNBID != nil {
        stream.set_choice(0, 0, 1, 1)
        var Pack_gNBID = func(st *Stream, self BITSTRING) {
            st.format_blen(int(self.Len-22), 4, 0)
            st.formatf_BitString(self.Value, int(self.Len))
        }
        Pack_gNBID(stream, *self.GNBID)//3
    }

}
type GNBID struct { //[{'type': 'BIT STRING', 'size': [(22, 32)], 'name': 'gNB-ID'}, None]
    GNBID *BITSTRING
} // GNBID

type GUAMI struct { // [{'type': 'PLMNIdentity', 'name': 'pLMNIdentity'}, {'type': 'AMFRegionID', 'name': 'aMFRegionID'}, {'type': 'AMFSetID', 'name': 'aMFSetID'}, {'type': 'AMFPointer', 'name': 'aMFPointer'}, None]
    PLMNIdentity PLMNIdentity
    AMFRegionID AMFRegionID
    AMFSetID AMFSetID
    AMFPointer AMFPointer
}

func (self * GUAMI) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PLMNIdentity.Unpack(stream)// p8
    self.AMFRegionID.Unpack(stream)// p8
    self.AMFSetID.Unpack(stream)// p8
    self.AMFPointer.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * GUAMI) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.AMFRegionID.Pack(stream)
    self.AMFSetID.Pack(stream)
    self.AMFPointer.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type IndexToRFSP struct {
  Value uint64
}
func (self *IndexToRFSP) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 9, 1, 1)
}
func (self * IndexToRFSP) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 9, 1, 1)
}
func (self *NgENBID)Unpack(stream *Stream) {
    //coptions := []string{"macroNgENB-ID","shortMacroNgENB-ID","longMacroNgENB-ID","Unknown"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in NgENBID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_macroNgENBID = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(20, 20)
        }
        self.MacroNgENBID = &BITSTRING{}//cho5
        Unpack_macroNgENBID(stream, self.MacroNgENBID);
    } else if choice == 1 { //ch2
        var Unpack_shortMacroNgENBID = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(18, 18)
        }
        self.ShortMacroNgENBID = &BITSTRING{}//cho5
        Unpack_shortMacroNgENBID(stream, self.ShortMacroNgENBID);
    } else if choice == 2 { //ch2
        var Unpack_longMacroNgENBID = func(st *Stream, self *BITSTRING){
            self.Value = st.parsef_BitString(21, 21)
        }
        self.LongMacroNgENBID = &BITSTRING{}//cho5
        Unpack_longMacroNgENBID(stream, self.LongMacroNgENBID);
    }//end of if else

    if choice >= 3 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * NgENBID) Pack(stream *Stream) {
    if self.MacroNgENBID != nil {
        stream.set_choice(0, 2, 1, 3)
        var Pack_macroNgENBID = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 20)
        }
        Pack_macroNgENBID(stream, *self.MacroNgENBID)//3
    } else if self.ShortMacroNgENBID != nil {
        stream.set_choice(1, 2, 1, 3)
        var Pack_shortMacroNgENBID = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 18)
        }
        Pack_shortMacroNgENBID(stream, *self.ShortMacroNgENBID)//3
    } else if self.LongMacroNgENBID != nil {
        stream.set_choice(2, 2, 1, 3)
        var Pack_longMacroNgENBID = func(st *Stream, self BITSTRING) {
            st.formatf_BitString(self.Value, 21)
        }
        Pack_longMacroNgENBID(stream, *self.LongMacroNgENBID)//3
    }

}
type NgENBID struct { //[{'type': 'BIT STRING', 'size': [20], 'name': 'macroNgENB-ID'}, {'type': 'BIT STRING', 'size': [18], 'name': 'shortMacroNgENB-ID'}, {'type': 'BIT STRING', 'size': [21], 'name': 'longMacroNgENB-ID'}, None]
    MacroNgENBID *BITSTRING
    ShortMacroNgENBID *BITSTRING
    LongMacroNgENBID *BITSTRING
} // NgENBID

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
type NRCGI struct { // [{'type': 'PLMNIdentity', 'name': 'pLMNIdentity'}, {'type': 'NRCellIdentity', 'name': 'nRCellIdentity'}, None]
    PLMNIdentity PLMNIdentity
    NRCellIdentity NRCellIdentity
}

func (self * NRCGI) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.PLMNIdentity.Unpack(stream)// p8
    self.NRCellIdentity.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * NRCGI) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.PLMNIdentity.Pack(stream)
    self.NRCellIdentity.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type PLMNIdentity struct {
  Value HexBytes
}
func (self *PLMNIdentity) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(3)
}
func (self *PLMNIdentity) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 3)
}
type QosFlowIdentifier struct {
  Value uint64
}
func (self *QosFlowIdentifier) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(64, 7, 1, 0)
}
func (self * QosFlowIdentifier) Pack(st *Stream){
    st.formatf_Integer(self.Value, 64, 7, 1, 0)
}
type SD struct {
  Value HexBytes
}
func (self *SD) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(3)
}
func (self *SD) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 3)
}
type SNSSAI struct { // [{'type': 'SST', 'name': 'sST'}, {'type': 'SD', 'name': 'sD', 'optional': True}, None]
    SST SST
    SD *SD
}

func (self * SNSSAI) Unpack(stream *Stream) {
    sD_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.SST.Unpack(stream)// p8
    if (sD_flag & _flags) == sD_flag { //cond2
        self.SD = &SD{}//7{'type': 'SD', 'name': 'sD', 'optional': True}
        self.SD.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SNSSAI) Pack(stream *Stream) {
    const sD_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.SST.Pack(stream)
    if self.SD != nil { 
        _flags |= sD_flag
        self.SD.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type SST struct {
  Value HexBytes
}
func (self *SST) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(1)
}
func (self *SST) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 1)
}
type NGRANnodeUEXnAPID struct {
  Value uint64
}
func (self *NGRANnodeUEXnAPID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967296, 32, 0, 0)
}
func (self * NGRANnodeUEXnAPID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967296, 32, 0, 0)
}
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
        self.GNB = &GlobalGNBID{}//cho6
        self.GNB.Unpack(stream)
    } else if choice == 1 { //ch2
        self.NgeNB = &GlobalNgENBID{}//cho6
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
type GlobalNGRANNodeID struct { //[{'type': 'GlobalGNB-ID', 'name': 'gNB'}, {'type': 'GlobalNgENB-ID', 'name': 'ng-eNB'}, None]
    GNB *GlobalGNBID
    NgeNB *GlobalNgENBID
} // GlobalNGRANNodeID

type GNBCUCPUEE1APID struct {
  Value uint64
}
func (self *GNBCUCPUEE1APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967296, 32, 0, 0)
}
func (self * GNBCUCPUEE1APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967296, 32, 0, 0)
}
type GNBCUUPID struct {
  Value uint64
}
func (self *GNBCUUPID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(68719476736, 36, 0, 0)
}
func (self * GNBCUUPID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 68719476736, 36, 0, 0)
}
type FiveGSTAC struct {
  Value HexBytes
}
func (self *FiveGSTAC) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(3)
}
func (self *FiveGSTAC) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 3)
}
type FreqBandNrItem struct { // [{'type': 'INTEGER', 'restricted-to': [(1, 1024), None], 'name': 'freqBandIndicatorNr'}, None]
    FreqBandIndicatorNr INTEGER
}

func (self * FreqBandNrItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_freqBandIndicatorNr = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(1024, 11, 1, 1)
    }
    Unpack_freqBandIndicatorNr(stream, &self.FreqBandIndicatorNr)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * FreqBandNrItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_freqBandIndicatorNr = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 1024, 11, 1, 1)
    }
    Pack_freqBandIndicatorNr(stream, self.FreqBandIndicatorNr) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type GNBCUUEF1APID struct {
  Value uint64
}
func (self *GNBCUUEF1APID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967296, 32, 0, 0)
}
func (self * GNBCUUEF1APID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967296, 32, 0, 0)
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
type NRPCI struct {
  Value uint64
}
func (self *NRPCI) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(1008, 10, 0, 0)
}
func (self * NRPCI) Pack(st *Stream){
    st.formatf_Integer(self.Value, 1008, 10, 0, 0)
}
type NRARFCN struct { // [{'type': 'INTEGER', 'restricted-to': [(0, 'maxNRARFCN')], 'name': 'nRARFCN'}, None]
    NRARFCN INTEGER
}

func (self * NRARFCN) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_nRARFCN = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(3279166, 22, 0, 0)
    }
    Unpack_nRARFCN(stream, &self.NRARFCN)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * NRARFCN) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_nRARFCN = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 3279166, 22, 0, 0)
    }
    Pack_nRARFCN(stream, self.NRARFCN) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *NRFrequencyBandList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32)
    _size += 1
    self.Items = make([]NRFrequencyBandItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *NRFrequencyBandList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type NRFrequencyBandList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'NRFrequencyBandItem'}, 'size': [(1, 'maxnoofNrCellBands')]}
    Items []NRFrequencyBandItem
}

type NRFrequencyBandItem struct { // [{'type': 'INTEGER', 'restricted-to': [(1, 1024), None], 'name': 'freqBandIndicatorNr'}, {'type': 'SupportedSULBandList', 'name': 'supportedSULBandList'}, None]
    FreqBandIndicatorNr INTEGER
    SupportedSULBandList SupportedSULBandList
}

func (self * NRFrequencyBandItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_freqBandIndicatorNr = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(1024, 11, 1, 1)
    }
    Unpack_freqBandIndicatorNr(stream, &self.FreqBandIndicatorNr)// p2
    self.SupportedSULBandList.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * NRFrequencyBandItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_freqBandIndicatorNr = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 1024, 11, 1, 1)
    }
    Pack_freqBandIndicatorNr(stream, self.FreqBandIndicatorNr) //f2
    self.SupportedSULBandList.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type NRFrequencyInfo struct { // [{'type': 'NR-ARFCN', 'name': 'nrARFCN'}, {'type': 'NRFrequencyBand-List', 'name': 'frequencyBand-List'}, {'type': 'NRFrequencyShift7p5khz', 'name': 'frequencyShift7p5khz', 'optional': True}, None]
    NrARFCN NRARFCN
    FrequencyBandList NRFrequencyBandList
    FrequencyShift7p5khz *NRFrequencyShift7p5khz
}

func (self * NRFrequencyInfo) Unpack(stream *Stream) {
    frequencyShift7p5khz_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.NrARFCN.Unpack(stream)// p8
    self.FrequencyBandList.Unpack(stream)// p8
    if (frequencyShift7p5khz_flag & _flags) == frequencyShift7p5khz_flag { //cond2
        self.FrequencyShift7p5khz = &NRFrequencyShift7p5khz{}//7{'type': 'NRFrequencyShift7p5khz', 'name': 'frequencyShift7p5khz', 'optional': True}
        self.FrequencyShift7p5khz.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * NRFrequencyInfo) Pack(stream *Stream) {
    const frequencyShift7p5khz_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.NrARFCN.Pack(stream)
    self.FrequencyBandList.Pack(stream)
    if self.FrequencyShift7p5khz != nil { 
        _flags |= frequencyShift7p5khz_flag
        self.FrequencyShift7p5khz.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type NRFrequencyShift7p5khz struct {
  Value int
}
const (
    NRFrequencyShift7p5khzfalse = 0
    NRFrequencyShift7p5khztrue = 1

    /* Extensions */
)
func (self *NRFrequencyShift7p5khz) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *NRFrequencyShift7p5khz) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
type RANUEID struct {
  Value HexBytes
}
func (self *RANUEID) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(8)
}
func (self *RANUEID) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 8)
}
func (self *SupportedSULBandList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(33)
    _size += 0
    self.Items = make([]SupportedSULFreqBandItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *SupportedSULBandList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-0, 33)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type SupportedSULBandList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'SupportedSULFreqBandItem'}, 'size': [(0, 'maxnoofNrCellBands')]}
    Items []SupportedSULFreqBandItem
}

type SupportedSULFreqBandItem struct { // [{'type': 'INTEGER', 'restricted-to': [(1, 1024), None], 'name': 'freqBandIndicatorNr'}, None]
    FreqBandIndicatorNr INTEGER
}

func (self * SupportedSULFreqBandItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_freqBandIndicatorNr = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(1024, 11, 1, 1)
    }
    Unpack_freqBandIndicatorNr(stream, &self.FreqBandIndicatorNr)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * SupportedSULFreqBandItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_freqBandIndicatorNr = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 1024, 11, 1, 1)
    }
    Pack_freqBandIndicatorNr(stream, self.FreqBandIndicatorNr) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type LogicalOR struct {
  Value int
}
const (
    LogicalORtrue = 0
    LogicalORfalse = 1

    /* Extensions */
)
func (self *LogicalOR) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(2, 2, 1)
}
func (self *LogicalOR) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 2, 2, 1)
}
func (self *NeighborCellList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65535)
    _size += 1
    self.Items = make([]NeighborCellItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *NeighborCellList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65535)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type NeighborCellList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'NeighborCell-Item'}, 'size': [(1, 'maxnoofNeighbourCell')]}
    Items []NeighborCellItem
}

func (self *NeighborCellItem)Unpack(stream *Stream) {
    //coptions := []string{"ranType-Choice-NR","ranType-Choice-EUTRA"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in NeighborCellItem\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.RanTypeChoiceNR = &NeighborCellItemChoiceNR{}//cho6
        self.RanTypeChoiceNR.Unpack(stream)
    } else if choice == 1 { //ch2
        self.RanTypeChoiceEUTRA = &NeighborCellItemChoiceEUTRA{}//cho6
        self.RanTypeChoiceEUTRA.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * NeighborCellItem) Pack(stream *Stream) {
    if self.RanTypeChoiceNR != nil {
        stream.set_choice(0, 1, 1, 2)
        self.RanTypeChoiceNR.Pack(stream)//2
    } else if self.RanTypeChoiceEUTRA != nil {
        stream.set_choice(1, 1, 1, 2)
        self.RanTypeChoiceEUTRA.Pack(stream)//2
    }

}
type NeighborCellItem struct { //[{'type': 'NeighborCell-Item-Choice-NR', 'name': 'ranType-Choice-NR'}, {'type': 'NeighborCell-Item-Choice-E-UTRA', 'name': 'ranType-Choice-EUTRA'}, None]
    RanTypeChoiceNR *NeighborCellItemChoiceNR
    RanTypeChoiceEUTRA *NeighborCellItemChoiceEUTRA
} // NeighborCellItem

type NeighborCellItemChoiceNR struct { // [{'type': 'NR-CGI', 'name': 'nR-CGI'}, {'type': 'NR-PCI', 'name': 'nR-PCI'}, {'type': 'FiveGS-TAC', 'name': 'fiveGS-TAC'}, {'type': 'ENUMERATED', 'values': [('fdd', 0), ('tdd', 1), None], 'name': 'nR-mode-info'}, {'type': 'NRFrequencyInfo', 'name': 'nR-FreqInfo'}, {'type': 'ENUMERATED', 'values': [('true', 0), ('false', 1), None], 'name': 'x2-Xn-established'}, {'type': 'ENUMERATED', 'values': [('true', 0), ('false', 1), None], 'name': 'hO-validated'}, {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'version'}, None]
    NRCGI NRCGI
    NRPCI NRPCI
    FiveGSTAC FiveGSTAC
    NRmodeinfo ENUMERATED
    NRFreqInfo NRFrequencyInfo
    X2Xnestablished ENUMERATED
    HOvalidated ENUMERATED
    Version INTEGER
}

func (self * NeighborCellItemChoiceNR) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.NRCGI.Unpack(stream)// p8
    self.NRPCI.Unpack(stream)// p8
    self.FiveGSTAC.Unpack(stream)// p8
    var Unpack_nRmodeinfo = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_nRmodeinfo(stream, &self.NRmodeinfo)// p2
    self.NRFreqInfo.Unpack(stream)// p8
    var Unpack_x2Xnestablished = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_x2Xnestablished(stream, &self.X2Xnestablished)// p2
    var Unpack_hOvalidated = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_hOvalidated(stream, &self.HOvalidated)// p2
    var Unpack_version = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(65535, 17, 1, 1)
    }
    Unpack_version(stream, &self.Version)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * NeighborCellItemChoiceNR) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.NRCGI.Pack(stream)
    self.NRPCI.Pack(stream)
    self.FiveGSTAC.Pack(stream)
    var Pack_nRmodeinfo = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_nRmodeinfo(stream, self.NRmodeinfo) //f2
    self.NRFreqInfo.Pack(stream)
    var Pack_x2Xnestablished = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_x2Xnestablished(stream, self.X2Xnestablished) //f2
    var Pack_hOvalidated = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_hOvalidated(stream, self.HOvalidated) //f2
    var Pack_version = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 65535, 17, 1, 1)
    }
    Pack_version(stream, self.Version) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type NeighborCellItemChoiceEUTRA struct { // [{'type': 'EUTRA-CGI', 'name': 'eUTRA-CGI'}, {'type': 'E-UTRA-PCI', 'name': 'eUTRA-PCI'}, {'type': 'E-UTRA-ARFCN', 'name': 'eUTRA-ARFCN'}, {'type': 'E-UTRA-TAC', 'name': 'eUTRA-TAC'}, {'type': 'ENUMERATED', 'values': [('true', 0), ('false', 1), None], 'name': 'x2-Xn-established'}, {'type': 'ENUMERATED', 'values': [('true', 0), ('false', 1), None], 'name': 'hO-validated'}, {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'version'}, None]
    EUTRACGI EUTRACGI
    EUTRAPCI EUTRAPCI
    EUTRAARFCN EUTRAARFCN
    EUTRATAC EUTRATAC
    X2Xnestablished ENUMERATED
    HOvalidated ENUMERATED
    Version INTEGER
}

func (self * NeighborCellItemChoiceEUTRA) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.EUTRACGI.Unpack(stream)// p8
    self.EUTRAPCI.Unpack(stream)// p8
    self.EUTRAARFCN.Unpack(stream)// p8
    self.EUTRATAC.Unpack(stream)// p8
    var Unpack_x2Xnestablished = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_x2Xnestablished(stream, &self.X2Xnestablished)// p2
    var Unpack_hOvalidated = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_hOvalidated(stream, &self.HOvalidated)// p2
    var Unpack_version = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(65535, 17, 1, 1)
    }
    Unpack_version(stream, &self.Version)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * NeighborCellItemChoiceEUTRA) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.EUTRACGI.Pack(stream)
    self.EUTRAPCI.Pack(stream)
    self.EUTRAARFCN.Pack(stream)
    self.EUTRATAC.Pack(stream)
    var Pack_x2Xnestablished = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_x2Xnestablished(stream, self.X2Xnestablished) //f2
    var Pack_hOvalidated = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_hOvalidated(stream, self.HOvalidated) //f2
    var Pack_version = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 65535, 17, 1, 1)
    }
    Pack_version(stream, self.Version) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type NeighborRelationInfo struct { // [{'type': 'ServingCell-PCI', 'name': 'servingCellPCI'}, {'type': 'ServingCell-ARFCN', 'name': 'servingCellARFCN'}, {'type': 'NeighborCell-List', 'name': 'neighborCell-List'}, None]
    ServingCellPCI ServingCellPCI
    ServingCellARFCN ServingCellARFCN
    NeighborCellList NeighborCellList
}

func (self * NeighborRelationInfo) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.ServingCellPCI.Unpack(stream)// p8
    self.ServingCellARFCN.Unpack(stream)// p8
    self.NeighborCellList.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * NeighborRelationInfo) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.ServingCellPCI.Pack(stream)
    self.ServingCellARFCN.Pack(stream)
    self.NeighborCellList.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RRCState struct {
  Value int
}
const (
    RRCStaterrc_connected = 0
    RRCStaterrc_inactive = 1
    RRCStaterrc_idle = 2
    RRCStateany = 3

    /* Extensions */
)
func (self *RRCState) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(3, 4, 1)
}
func (self *RRCState) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 3, 4, 1)
}
type EventTriggerCellInfo_CellInfoList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'EventTrigger-Cell-Info-Item'}, 'size': [(1, 'maxnoofCellInfo')], 'name': 'cellInfo-List'}
    Items []EventTriggerCellInfoItem
}
type EventTriggerCellInfo struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'EventTrigger-Cell-Info-Item'}, 'size': [(1, 'maxnoofCellInfo')], 'name': 'cellInfo-List'}, None]
    CellInfoList EventTriggerCellInfo_CellInfoList
}

func (self * EventTriggerCellInfo) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_cellInfoList = func(stream *Stream, self *EventTriggerCellInfo_CellInfoList){// Seq6 EventTriggerCellInfo {'type': 'SEQUENCE OF', 'element': {'type': 'EventTrigger-Cell-Info-Item'}, 'size': [(1, 'maxnoofCellInfo')], 'name': 'cellInfo-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]EventTriggerCellInfoItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_cellInfoList(stream, &self.CellInfoList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EventTriggerCellInfo) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_cellInfoList = func(stream *Stream, self EventTriggerCellInfo_CellInfoList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_cellInfoList(stream, self.CellInfoList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EventTriggerCellInfoItem_CellType struct { //[{'type': 'EventTrigger-Cell-Info-Item-Choice-Individual', 'name': 'cellType-Choice-Individual'}, {'type': 'EventTrigger-Cell-Info-Item-Choice-Group', 'name': 'cellType-Choice-Group'}, None]
    CellTypeChoiceIndividual *EventTriggerCellInfoItemChoiceIndividual
    CellTypeChoiceGroup *EventTriggerCellInfoItemChoiceGroup
} // EventTriggerCellInfoItem_CellType

type EventTriggerCellInfoItem struct { // [{'type': 'RIC-EventTrigger-Cell-ID', 'name': 'eventTriggerCellID'}, {'type': 'CHOICE', 'members': [{'type': 'EventTrigger-Cell-Info-Item-Choice-Individual', 'name': 'cellType-Choice-Individual'}, {'type': 'EventTrigger-Cell-Info-Item-Choice-Group', 'name': 'cellType-Choice-Group'}, None], 'name': 'cellType'}, {'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}, None]
    EventTriggerCellID RICEventTriggerCellID
    CellType EventTriggerCellInfoItem_CellType
    LogicalOR *LogicalOR
}

func (self * EventTriggerCellInfoItem) Unpack(stream *Stream) {
    logicalOR_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.EventTriggerCellID.Unpack(stream)// p8
    var Unpack_cellType = func(stream *Stream, self *EventTriggerCellInfoItem_CellType) {
        //coptions := []string{"cellType-Choice-Individual","cellType-Choice-Group"}
        choice := stream.get_choice(1, 1, 2)
        choice_len := 0
        choice_loc := 0
        if choice >= 2 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in EventTriggerCellInfoItem_CellType\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.CellTypeChoiceIndividual = &EventTriggerCellInfoItemChoiceIndividual{}//cho6
            self.CellTypeChoiceIndividual.Unpack(stream)
        } else if choice == 1 { //ch2
            self.CellTypeChoiceGroup = &EventTriggerCellInfoItemChoiceGroup{}//cho6
            self.CellTypeChoiceGroup.Unpack(stream)
        }//end of if else

        if choice >= 2 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_cellType(stream, &self.CellType)// p2
    if (logicalOR_flag & _flags) == logicalOR_flag { //cond2
        self.LogicalOR = &LogicalOR{}//7{'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}
        self.LogicalOR.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EventTriggerCellInfoItem) Pack(stream *Stream) {
    const logicalOR_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.EventTriggerCellID.Pack(stream)
    var Pack_cellType = func(stream *Stream, self EventTriggerCellInfoItem_CellType) {
        if self.CellTypeChoiceIndividual != nil {
            stream.set_choice(0, 1, 1, 2)
            self.CellTypeChoiceIndividual.Pack(stream)//2
        } else if self.CellTypeChoiceGroup != nil {
            stream.set_choice(1, 1, 1, 2)
            self.CellTypeChoiceGroup.Pack(stream)//2
        }

    }
    Pack_cellType(stream, self.CellType) //f2
    if self.LogicalOR != nil { 
        _flags |= logicalOR_flag
        self.LogicalOR.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type EventTriggerCellInfoItemChoiceIndividual struct { // [{'type': 'CGI', 'name': 'cellGlobalID'}, None]
    CellGlobalID CGI
}

func (self * EventTriggerCellInfoItemChoiceIndividual) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.CellGlobalID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EventTriggerCellInfoItemChoiceIndividual) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CellGlobalID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EventTriggerCellInfoItemChoiceGroup struct { // [{'type': 'RANParameter-Testing', 'name': 'ranParameterTesting'}, None]
    RanParameterTesting RANParameterTesting
}

func (self * EventTriggerCellInfoItemChoiceGroup) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterTesting.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EventTriggerCellInfoItemChoiceGroup) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterTesting.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EventTriggerUEInfo_UeInfoList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'EventTrigger-UE-Info-Item'}, 'size': [(1, 'maxnoofUEInfo')], 'name': 'ueInfo-List'}
    Items []EventTriggerUEInfoItem
}
type EventTriggerUEInfo struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'EventTrigger-UE-Info-Item'}, 'size': [(1, 'maxnoofUEInfo')], 'name': 'ueInfo-List'}, None]
    UeInfoList EventTriggerUEInfo_UeInfoList
}

func (self * EventTriggerUEInfo) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ueInfoList = func(stream *Stream, self *EventTriggerUEInfo_UeInfoList){// Seq6 EventTriggerUEInfo {'type': 'SEQUENCE OF', 'element': {'type': 'EventTrigger-UE-Info-Item'}, 'size': [(1, 'maxnoofUEInfo')], 'name': 'ueInfo-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]EventTriggerUEInfoItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ueInfoList(stream, &self.UeInfoList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EventTriggerUEInfo) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ueInfoList = func(stream *Stream, self EventTriggerUEInfo_UeInfoList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ueInfoList(stream, self.UeInfoList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EventTriggerUEInfoItem_UeType struct { //[{'type': 'EventTrigger-UE-Info-Item-Choice-Individual', 'name': 'ueType-Choice-Individual'}, {'type': 'EventTrigger-UE-Info-Item-Choice-Group', 'name': 'ueType-Choice-Group'}, None]
    UeTypeChoiceIndividual *EventTriggerUEInfoItemChoiceIndividual
    UeTypeChoiceGroup *EventTriggerUEInfoItemChoiceGroup
} // EventTriggerUEInfoItem_UeType

type EventTriggerUEInfoItem struct { // [{'type': 'RIC-EventTrigger-UE-ID', 'name': 'eventTriggerUEID'}, {'type': 'CHOICE', 'members': [{'type': 'EventTrigger-UE-Info-Item-Choice-Individual', 'name': 'ueType-Choice-Individual'}, {'type': 'EventTrigger-UE-Info-Item-Choice-Group', 'name': 'ueType-Choice-Group'}, None], 'name': 'ueType'}, {'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}, None]
    EventTriggerUEID RICEventTriggerUEID
    UeType EventTriggerUEInfoItem_UeType
    LogicalOR *LogicalOR
}

func (self * EventTriggerUEInfoItem) Unpack(stream *Stream) {
    logicalOR_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.EventTriggerUEID.Unpack(stream)// p8
    var Unpack_ueType = func(stream *Stream, self *EventTriggerUEInfoItem_UeType) {
        //coptions := []string{"ueType-Choice-Individual","ueType-Choice-Group"}
        choice := stream.get_choice(1, 1, 2)
        choice_len := 0
        choice_loc := 0
        if choice >= 2 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in EventTriggerUEInfoItem_UeType\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.UeTypeChoiceIndividual = &EventTriggerUEInfoItemChoiceIndividual{}//cho6
            self.UeTypeChoiceIndividual.Unpack(stream)
        } else if choice == 1 { //ch2
            self.UeTypeChoiceGroup = &EventTriggerUEInfoItemChoiceGroup{}//cho6
            self.UeTypeChoiceGroup.Unpack(stream)
        }//end of if else

        if choice >= 2 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_ueType(stream, &self.UeType)// p2
    if (logicalOR_flag & _flags) == logicalOR_flag { //cond2
        self.LogicalOR = &LogicalOR{}//7{'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}
        self.LogicalOR.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EventTriggerUEInfoItem) Pack(stream *Stream) {
    const logicalOR_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.EventTriggerUEID.Pack(stream)
    var Pack_ueType = func(stream *Stream, self EventTriggerUEInfoItem_UeType) {
        if self.UeTypeChoiceIndividual != nil {
            stream.set_choice(0, 1, 1, 2)
            self.UeTypeChoiceIndividual.Pack(stream)//2
        } else if self.UeTypeChoiceGroup != nil {
            stream.set_choice(1, 1, 1, 2)
            self.UeTypeChoiceGroup.Pack(stream)//2
        }

    }
    Pack_ueType(stream, self.UeType) //f2
    if self.LogicalOR != nil { 
        _flags |= logicalOR_flag
        self.LogicalOR.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type EventTriggerUEInfoItemChoiceIndividual struct { // [{'type': 'UEID', 'name': 'ueID'}, {'type': 'RANParameter-Testing', 'name': 'ranParameterTesting', 'optional': True}, None]
    UeID UEID
    RanParameterTesting *RANParameterTesting
}

func (self * EventTriggerUEInfoItemChoiceIndividual) Unpack(stream *Stream) {
    ranParameterTesting_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.UeID.Unpack(stream)// p8
    if (ranParameterTesting_flag & _flags) == ranParameterTesting_flag { //cond2
        self.RanParameterTesting = &RANParameterTesting{}//7{'type': 'RANParameter-Testing', 'name': 'ranParameterTesting', 'optional': True}
        self.RanParameterTesting.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EventTriggerUEInfoItemChoiceIndividual) Pack(stream *Stream) {
    const ranParameterTesting_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UeID.Pack(stream)
    if self.RanParameterTesting != nil { 
        _flags |= ranParameterTesting_flag
        self.RanParameterTesting.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type EventTriggerUEInfoItemChoiceGroup struct { // [{'type': 'RANParameter-Testing', 'name': 'ranParameterTesting'}, None]
    RanParameterTesting RANParameterTesting
}

func (self * EventTriggerUEInfoItemChoiceGroup) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterTesting.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EventTriggerUEInfoItemChoiceGroup) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterTesting.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EventTriggerUEeventInfo_UeEventList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'EventTrigger-UEevent-Info-Item'}, 'size': [(1, 'maxnoofUEeventInfo')], 'name': 'ueEvent-List'}
    Items []EventTriggerUEeventInfoItem
}
type EventTriggerUEeventInfo struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'EventTrigger-UEevent-Info-Item'}, 'size': [(1, 'maxnoofUEeventInfo')], 'name': 'ueEvent-List'}, None]
    UeEventList EventTriggerUEeventInfo_UeEventList
}

func (self * EventTriggerUEeventInfo) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ueEventList = func(stream *Stream, self *EventTriggerUEeventInfo_UeEventList){// Seq6 EventTriggerUEeventInfo {'type': 'SEQUENCE OF', 'element': {'type': 'EventTrigger-UEevent-Info-Item'}, 'size': [(1, 'maxnoofUEeventInfo')], 'name': 'ueEvent-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]EventTriggerUEeventInfoItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ueEventList(stream, &self.UeEventList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EventTriggerUEeventInfo) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ueEventList = func(stream *Stream, self EventTriggerUEeventInfo_UeEventList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ueEventList(stream, self.UeEventList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EventTriggerUEeventInfoItem struct { // [{'type': 'RIC-EventTrigger-UEevent-ID', 'name': 'ueEventID'}, {'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}, None]
    UeEventID RICEventTriggerUEeventID
    LogicalOR *LogicalOR
}

func (self * EventTriggerUEeventInfoItem) Unpack(stream *Stream) {
    logicalOR_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.UeEventID.Unpack(stream)// p8
    if (logicalOR_flag & _flags) == logicalOR_flag { //cond2
        self.LogicalOR = &LogicalOR{}//7{'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}
        self.LogicalOR.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EventTriggerUEeventInfoItem) Pack(stream *Stream) {
    const logicalOR_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UeEventID.Pack(stream)
    if self.LogicalOR != nil { 
        _flags |= logicalOR_flag
        self.LogicalOR.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANParameterID struct {
  Value uint64
}
func (self *RANParameterID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967295, 33, 1, 1)
}
func (self * RANParameterID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967295, 33, 1, 1)
}
type RANParameterName struct {
  Value string
}
func (self *RANParameterName) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in RANParameter-Name")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *RANParameterName) Pack(st *Stream) {
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
type RANParameterDefinition struct { // [{'type': 'RANParameter-Definition-Choice', 'name': 'ranParameter-Definition-Choice'}, None]
    RanParameterDefinitionChoice RANParameterDefinitionChoice
}

func (self * RANParameterDefinition) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterDefinitionChoice.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterDefinition) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterDefinitionChoice.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *RANParameterDefinitionChoice)Unpack(stream *Stream) {
    //coptions := []string{"choiceLIST","choiceSTRUCTURE"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in RANParameterDefinitionChoice\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.ChoiceLIST = &RANParameterDefinitionChoiceLIST{}//cho6
        self.ChoiceLIST.Unpack(stream)
    } else if choice == 1 { //ch2
        self.ChoiceSTRUCTURE = &RANParameterDefinitionChoiceSTRUCTURE{}//cho6
        self.ChoiceSTRUCTURE.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * RANParameterDefinitionChoice) Pack(stream *Stream) {
    if self.ChoiceLIST != nil {
        stream.set_choice(0, 1, 1, 2)
        self.ChoiceLIST.Pack(stream)//2
    } else if self.ChoiceSTRUCTURE != nil {
        stream.set_choice(1, 1, 1, 2)
        self.ChoiceSTRUCTURE.Pack(stream)//2
    }

}
type RANParameterDefinitionChoice struct { //[{'type': 'RANParameter-Definition-Choice-LIST', 'name': 'choiceLIST'}, {'type': 'RANParameter-Definition-Choice-STRUCTURE', 'name': 'choiceSTRUCTURE'}, None]
    ChoiceLIST *RANParameterDefinitionChoiceLIST
    ChoiceSTRUCTURE *RANParameterDefinitionChoiceSTRUCTURE
} // RANParameterDefinitionChoice

type RANParameterDefinitionChoiceLIST_RanParameterList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-Definition-Choice-LIST-Item'}, 'size': [(1, 'maxnoofItemsinList')], 'name': 'ranParameter-List'}
    Items []RANParameterDefinitionChoiceLISTItem
}
type RANParameterDefinitionChoiceLIST struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-Definition-Choice-LIST-Item'}, 'size': [(1, 'maxnoofItemsinList')], 'name': 'ranParameter-List'}, None]
    RanParameterList RANParameterDefinitionChoiceLIST_RanParameterList
}

func (self * RANParameterDefinitionChoiceLIST) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ranParameterList = func(stream *Stream, self *RANParameterDefinitionChoiceLIST_RanParameterList){// Seq6 RANParameterDefinitionChoiceLIST {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-Definition-Choice-LIST-Item'}, 'size': [(1, 'maxnoofItemsinList')], 'name': 'ranParameter-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]RANParameterDefinitionChoiceLISTItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranParameterList(stream, &self.RanParameterList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterDefinitionChoiceLIST) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ranParameterList = func(stream *Stream, self RANParameterDefinitionChoiceLIST_RanParameterList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranParameterList(stream, self.RanParameterList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANParameterDefinitionChoiceLISTItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}, None]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * RANParameterDefinitionChoiceLISTItem) Unpack(stream *Stream) {
    ranParameterDefinition_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (ranParameterDefinition_flag & _flags) == ranParameterDefinition_flag { //cond2
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterDefinitionChoiceLISTItem) Pack(stream *Stream) {
    const ranParameterDefinition_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        _flags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANParameterDefinitionChoiceSTRUCTURE_RanParameterSTRUCTURE struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-Definition-Choice-STRUCTURE-Item'}, 'size': [(1, 'maxnoofParametersinStructure')], 'name': 'ranParameter-STRUCTURE'}
    Items []RANParameterDefinitionChoiceSTRUCTUREItem
}
type RANParameterDefinitionChoiceSTRUCTURE struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-Definition-Choice-STRUCTURE-Item'}, 'size': [(1, 'maxnoofParametersinStructure')], 'name': 'ranParameter-STRUCTURE'}, None]
    RanParameterSTRUCTURE RANParameterDefinitionChoiceSTRUCTURE_RanParameterSTRUCTURE
}

func (self * RANParameterDefinitionChoiceSTRUCTURE) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ranParameterSTRUCTURE = func(stream *Stream, self *RANParameterDefinitionChoiceSTRUCTURE_RanParameterSTRUCTURE){// Seq6 RANParameterDefinitionChoiceSTRUCTURE {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-Definition-Choice-STRUCTURE-Item'}, 'size': [(1, 'maxnoofParametersinStructure')], 'name': 'ranParameter-STRUCTURE'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]RANParameterDefinitionChoiceSTRUCTUREItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranParameterSTRUCTURE(stream, &self.RanParameterSTRUCTURE)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterDefinitionChoiceSTRUCTURE) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ranParameterSTRUCTURE = func(stream *Stream, self RANParameterDefinitionChoiceSTRUCTURE_RanParameterSTRUCTURE) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranParameterSTRUCTURE(stream, self.RanParameterSTRUCTURE) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANParameterDefinitionChoiceSTRUCTUREItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}, None]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * RANParameterDefinitionChoiceSTRUCTUREItem) Unpack(stream *Stream) {
    ranParameterDefinition_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (ranParameterDefinition_flag & _flags) == ranParameterDefinition_flag { //cond2
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterDefinitionChoiceSTRUCTUREItem) Pack(stream *Stream) {
    const ranParameterDefinition_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        _flags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *RANParameterValue)Unpack(stream *Stream) {
    //coptions := []string{"valueBoolean","valueInt","valueReal","valueBitS","valueOctS","valuePrintableString","Unknown","Unknown"}
    choice := stream.get_choice(3, 1, 6)
    choice_len := 0
    choice_loc := 0
    if choice >= 6 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in RANParameterValue\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var  Unpack_valueBoolean = func (st *Stream, self *BOOLEAN) {
            self.Value = st.parsef_bool()
            }
        self.ValueBoolean = &BOOLEAN{}//cho5
        Unpack_valueBoolean(stream, self.ValueBoolean);
    } else if choice == 1 { //ch2
        var Unpack_valueInt = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(0, 0, 0, 0)
        }
        self.ValueInt = &INTEGER{}//cho5
        Unpack_valueInt(stream, self.ValueInt);
    } else if choice == 2 { //ch2
        var Unpack_valueReal = func (st *Stream, self *REAL) {
            self.Value = st.parsef_Real(0)
        }
        self.ValueReal = &REAL{}//cho5
        Unpack_valueReal(stream, self.ValueReal);
    } else if choice == 3 { //ch2
        var Unpack_valueBitS = func(st *Stream, self *BITSTRING){
            self.Len = st.parse_len(0)
            self.Value = st.parsef_BitString(0, self.Len)
        }
        self.ValueBitS = &BITSTRING{}//cho5
        Unpack_valueBitS(stream, self.ValueBitS);
    } else if choice == 4 { //ch2
        var Unpack_valueOctS = func(st *Stream, self *OCTETSTRING) {
            _len := st.parse_len(0)
            self.Value = st.parsef_OctString(_len)
        }
        self.ValueOctS = &OCTETSTRING{}//cho5
        Unpack_valueOctS(stream, self.ValueOctS);
    } else if choice == 5 { //ch2
        var Unpack_valuePrintableString = func(st *Stream, self *PrintableString) {
            _len := st.parse_len(0)
            self.Value = st.parsef_PriString(_len)
        }
        self.ValuePrintableString = &PrintableString{}//cho5
        Unpack_valuePrintableString(stream, self.ValuePrintableString);
    }//end of if else

    if choice >= 6 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * RANParameterValue) Pack(stream *Stream) {
    if self.ValueBoolean != nil {
        stream.set_choice(0, 3, 1, 6)
        var Pack_valueBoolean = func(st *Stream, self BOOLEAN) {
            st.formatf_bool(self.Value);
            }
        Pack_valueBoolean(stream, *self.ValueBoolean)//3
    } else if self.ValueInt != nil {
        stream.set_choice(1, 3, 1, 6)
        var Pack_valueInt = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 0, 0, 0, 0)
        }
        Pack_valueInt(stream, *self.ValueInt)//3
    } else if self.ValueReal != nil {
        stream.set_choice(2, 3, 1, 6)
        var Pack_valueReal = func (st *Stream, self REAL) {
            st.formatf_Real(self.Value, 0)

}
        Pack_valueReal(stream, *self.ValueReal)//3
    } else if self.ValueBitS != nil {
        stream.set_choice(3, 3, 1, 6)
        var Pack_valueBitS = func(st *Stream, self BITSTRING) {
            st.format_len(self.Len, 0)
            st.formatf_BitString(self.Value, self.Len)
        }
        Pack_valueBitS(stream, *self.ValueBitS)//3
    } else if self.ValueOctS != nil {
        stream.set_choice(4, 3, 1, 6)
        var Pack_valueOctS = func(st *Stream, self OCTETSTRING) {
            st.format_len(len(self.Value), 0)
            st.formatf_OctString(self.Value, 0)
        }
        Pack_valueOctS(stream, *self.ValueOctS)//3
    } else if self.ValuePrintableString != nil {
        stream.set_choice(5, 3, 1, 6)
        var Pack_valuePrintableString = func(st *Stream, self PrintableString) {
            st.format_len((len(self.Value)), 0)
            st.formatf_PriString(self.Value, len(self.Value))
        }
        Pack_valuePrintableString(stream, *self.ValuePrintableString)//3
    }

}
type RANParameterValue struct { //[{'type': 'BOOLEAN', 'name': 'valueBoolean'}, {'type': 'INTEGER', 'name': 'valueInt'}, {'type': 'REAL', 'name': 'valueReal'}, {'type': 'BIT STRING', 'name': 'valueBitS'}, {'type': 'OCTET STRING', 'name': 'valueOctS'}, {'type': 'PrintableString', 'name': 'valuePrintableString'}, None]
    ValueBoolean *BOOLEAN
    ValueInt *INTEGER
    ValueReal *REAL
    ValueBitS *BITSTRING
    ValueOctS *OCTETSTRING
    ValuePrintableString *PrintableString
} // RANParameterValue

func (self *RANParameterValueType)Unpack(stream *Stream) {
    //coptions := []string{"ranP-Choice-ElementTrue","ranP-Choice-ElementFalse","ranP-Choice-Structure","ranP-Choice-List"}
    choice := stream.get_choice(2, 1, 4)
    choice_len := 0
    choice_loc := 0
    if choice >= 4 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in RANParameterValueType\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.RanPChoiceElementTrue = &RANParameterValueTypeChoiceElementTrue{}//cho6
        self.RanPChoiceElementTrue.Unpack(stream)
    } else if choice == 1 { //ch2
        self.RanPChoiceElementFalse = &RANParameterValueTypeChoiceElementFalse{}//cho6
        self.RanPChoiceElementFalse.Unpack(stream)
    } else if choice == 2 { //ch2
        self.RanPChoiceStructure = &RANParameterValueTypeChoiceStructure{}//cho6
        self.RanPChoiceStructure.Unpack(stream)
    } else if choice == 3 { //ch2
        self.RanPChoiceList = &RANParameterValueTypeChoiceList{}//cho6
        self.RanPChoiceList.Unpack(stream)
    }//end of if else

    if choice >= 4 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * RANParameterValueType) Pack(stream *Stream) {
    if self.RanPChoiceElementTrue != nil {
        stream.set_choice(0, 2, 1, 4)
        self.RanPChoiceElementTrue.Pack(stream)//2
    } else if self.RanPChoiceElementFalse != nil {
        stream.set_choice(1, 2, 1, 4)
        self.RanPChoiceElementFalse.Pack(stream)//2
    } else if self.RanPChoiceStructure != nil {
        stream.set_choice(2, 2, 1, 4)
        self.RanPChoiceStructure.Pack(stream)//2
    } else if self.RanPChoiceList != nil {
        stream.set_choice(3, 2, 1, 4)
        self.RanPChoiceList.Pack(stream)//2
    }

}
type RANParameterValueType struct { //[{'type': 'RANParameter-ValueType-Choice-ElementTrue', 'name': 'ranP-Choice-ElementTrue'}, {'type': 'RANParameter-ValueType-Choice-ElementFalse', 'name': 'ranP-Choice-ElementFalse'}, {'type': 'RANParameter-ValueType-Choice-Structure', 'name': 'ranP-Choice-Structure'}, {'type': 'RANParameter-ValueType-Choice-List', 'name': 'ranP-Choice-List'}, None]
    RanPChoiceElementTrue *RANParameterValueTypeChoiceElementTrue
    RanPChoiceElementFalse *RANParameterValueTypeChoiceElementFalse
    RanPChoiceStructure *RANParameterValueTypeChoiceStructure
    RanPChoiceList *RANParameterValueTypeChoiceList
} // RANParameterValueType

type RANParameterValueTypeChoiceElementTrue struct { // [{'type': 'RANParameter-Value', 'name': 'ranParameter-value'}, None]
    RanParametervalue RANParameterValue
}

func (self * RANParameterValueTypeChoiceElementTrue) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParametervalue.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterValueTypeChoiceElementTrue) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParametervalue.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANParameterValueTypeChoiceElementFalse struct { // [{'type': 'RANParameter-Value', 'name': 'ranParameter-value', 'optional': True}, None]
    RanParametervalue *RANParameterValue
}

func (self * RANParameterValueTypeChoiceElementFalse) Unpack(stream *Stream) {
    ranParametervalue_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    if (ranParametervalue_flag & _flags) == ranParametervalue_flag { //cond2
        self.RanParametervalue = &RANParameterValue{}//7{'type': 'RANParameter-Value', 'name': 'ranParameter-value', 'optional': True}
        self.RanParametervalue.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterValueTypeChoiceElementFalse) Pack(stream *Stream) {
    const ranParametervalue_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.RanParametervalue != nil { 
        _flags |= ranParametervalue_flag
        self.RanParametervalue.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANParameterValueTypeChoiceStructure struct { // [{'type': 'RANParameter-STRUCTURE', 'name': 'ranParameter-Structure'}, None]
    RanParameterStructure RANParameterSTRUCTURE
}

func (self * RANParameterValueTypeChoiceStructure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterStructure.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterValueTypeChoiceStructure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterStructure.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANParameterValueTypeChoiceList struct { // [{'type': 'RANParameter-LIST', 'name': 'ranParameter-List'}, None]
    RanParameterList RANParameterLIST
}

func (self * RANParameterValueTypeChoiceList) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterList.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterValueTypeChoiceList) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterList.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANParameterSTRUCTURE_SequenceofranParameters struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-STRUCTURE-Item'}, 'size': [(1, 'maxnoofParametersinStructure')], 'name': 'sequence-of-ranParameters', 'optional': True}
    Items []RANParameterSTRUCTUREItem
}
type RANParameterSTRUCTURE struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-STRUCTURE-Item'}, 'size': [(1, 'maxnoofParametersinStructure')], 'name': 'sequence-of-ranParameters', 'optional': True}, None]
    SequenceofranParameters *RANParameterSTRUCTURE_SequenceofranParameters
}

func (self * RANParameterSTRUCTURE) Unpack(stream *Stream) {
    sequenceofranParameters_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    if (sequenceofranParameters_flag & _flags) == sequenceofranParameters_flag { //cond1
        var Unpack_sequenceofranParameters = func(stream *Stream, self *RANParameterSTRUCTURE_SequenceofranParameters){// Seq6 RANParameterSTRUCTURE {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-STRUCTURE-Item'}, 'size': [(1, 'maxnoofParametersinStructure')], 'name': 'sequence-of-ranParameters', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]RANParameterSTRUCTUREItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.SequenceofranParameters = &RANParameterSTRUCTURE_SequenceofranParameters{}//3
        Unpack_sequenceofranParameters(stream, self.SequenceofranParameters)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-STRUCTURE-Item'}, 'size': [(1, 'maxnoofParametersinStructure')], 'name': 'sequence-of-ranParameters', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterSTRUCTURE) Pack(stream *Stream) {
    const sequenceofranParameters_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.SequenceofranParameters != nil { //YY
        _flags |= sequenceofranParameters_flag
        var Pack_sequenceofranParameters = func(stream *Stream, self RANParameterSTRUCTURE_SequenceofranParameters) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_sequenceofranParameters(stream, *self.SequenceofranParameters) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANParameterSTRUCTUREItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
}

func (self * RANParameterSTRUCTUREItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterSTRUCTUREItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANParameterLIST_ListofranParameter struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-STRUCTURE'}, 'size': [(1, 'maxnoofItemsinList')], 'name': 'list-of-ranParameter'}
    Items []RANParameterSTRUCTURE
}
type RANParameterLIST struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-STRUCTURE'}, 'size': [(1, 'maxnoofItemsinList')], 'name': 'list-of-ranParameter'}, None]
    ListofranParameter RANParameterLIST_ListofranParameter
}

func (self * RANParameterLIST) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_listofranParameter = func(stream *Stream, self *RANParameterLIST_ListofranParameter){// Seq6 RANParameterLIST {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-STRUCTURE'}, 'size': [(1, 'maxnoofItemsinList')], 'name': 'list-of-ranParameter'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]RANParameterSTRUCTURE, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_listofranParameter(stream, &self.ListofranParameter)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterLIST) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_listofranParameter = func(stream *Stream, self RANParameterLIST_ListofranParameter) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_listofranParameter(stream, self.ListofranParameter) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *RANParameterTesting) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(255)
    _size += 1
    self.Items = make([]RANParameterTestingItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *RANParameterTesting) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 255)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type RANParameterTesting struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-Testing-Item'}, 'size': [(1, 'maxnoofRANparamTest')]}
    Items []RANParameterTestingItem
}

func (self *RANParameterTestingCondition)Unpack(stream *Stream) {
    //coptions := []string{"ranP-Choice-comparison","ranP-Choice-presence"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in RANParameterTestingCondition\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_ranPChoicecomparison = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(4, 6, 1)
        }
        self.RanPChoicecomparison = &ENUMERATED{}//cho5
        Unpack_ranPChoicecomparison(stream, self.RanPChoicecomparison);
    } else if choice == 1 { //ch2
        var Unpack_ranPChoicepresence = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(3, 4, 1)
        }
        self.RanPChoicepresence = &ENUMERATED{}//cho5
        Unpack_ranPChoicepresence(stream, self.RanPChoicepresence);
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * RANParameterTestingCondition) Pack(stream *Stream) {
    if self.RanPChoicecomparison != nil {
        stream.set_choice(0, 1, 1, 2)
        var Pack_ranPChoicecomparison = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 4, 6, 1)
        }
        Pack_ranPChoicecomparison(stream, *self.RanPChoicecomparison)//3
    } else if self.RanPChoicepresence != nil {
        stream.set_choice(1, 1, 1, 2)
        var Pack_ranPChoicepresence = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 3, 4, 1)
        }
        Pack_ranPChoicepresence(stream, *self.RanPChoicepresence)//3
    }

}
type RANParameterTestingCondition struct { //[{'type': 'ENUMERATED', 'values': [('equal', 0), ('difference', 1), ('greaterthan', 2), ('lessthan', 3), ('contains', 4), ('starts-with', 5), None], 'name': 'ranP-Choice-comparison'}, {'type': 'ENUMERATED', 'values': [('present', 0), ('configured', 1), ('rollover', 2), ('non-zero', 3), None, ('value-change', 4)], 'name': 'ranP-Choice-presence'}, None]
    RanPChoicecomparison *ENUMERATED
    RanPChoicepresence *ENUMERATED
} // RANParameterTestingCondition

type RANParameterTestingItem_RanParameterType struct { //[{'type': 'RANParameter-Testing-Item-Choice-List', 'name': 'ranP-Choice-List'}, {'type': 'RANParameter-Testing-Item-Choice-Structure', 'name': 'ranP-Choice-Structure'}, {'type': 'RANParameter-Testing-Item-Choice-ElementTrue', 'name': 'ranP-Choice-ElementTrue'}, {'type': 'RANParameter-Testing-Item-Choice-ElementFalse', 'name': 'ranP-Choice-ElementFalse'}, None]
    RanPChoiceList *RANParameterTestingItemChoiceList
    RanPChoiceStructure *RANParameterTestingItemChoiceStructure
    RanPChoiceElementTrue *RANParameterTestingItemChoiceElementTrue
    RanPChoiceElementFalse *RANParameterTestingItemChoiceElementFalse
} // RANParameterTestingItem_RanParameterType

type RANParameterTestingItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'CHOICE', 'members': [{'type': 'RANParameter-Testing-Item-Choice-List', 'name': 'ranP-Choice-List'}, {'type': 'RANParameter-Testing-Item-Choice-Structure', 'name': 'ranP-Choice-Structure'}, {'type': 'RANParameter-Testing-Item-Choice-ElementTrue', 'name': 'ranP-Choice-ElementTrue'}, {'type': 'RANParameter-Testing-Item-Choice-ElementFalse', 'name': 'ranP-Choice-ElementFalse'}, None], 'name': 'ranParameter-Type'}, None]
    RanParameterID RANParameterID
    RanParameterType RANParameterTestingItem_RanParameterType
}

func (self * RANParameterTestingItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    var Unpack_ranParameterType = func(stream *Stream, self *RANParameterTestingItem_RanParameterType) {
        //coptions := []string{"ranP-Choice-List","ranP-Choice-Structure","ranP-Choice-ElementTrue","ranP-Choice-ElementFalse"}
        choice := stream.get_choice(2, 1, 4)
        choice_len := 0
        choice_loc := 0
        if choice >= 4 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in RANParameterTestingItem_RanParameterType\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.RanPChoiceList = &RANParameterTestingItemChoiceList{}//cho6
            self.RanPChoiceList.Unpack(stream)
        } else if choice == 1 { //ch2
            self.RanPChoiceStructure = &RANParameterTestingItemChoiceStructure{}//cho6
            self.RanPChoiceStructure.Unpack(stream)
        } else if choice == 2 { //ch2
            self.RanPChoiceElementTrue = &RANParameterTestingItemChoiceElementTrue{}//cho6
            self.RanPChoiceElementTrue.Unpack(stream)
        } else if choice == 3 { //ch2
            self.RanPChoiceElementFalse = &RANParameterTestingItemChoiceElementFalse{}//cho6
            self.RanPChoiceElementFalse.Unpack(stream)
        }//end of if else

        if choice >= 4 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_ranParameterType(stream, &self.RanParameterType)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterTestingItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    var Pack_ranParameterType = func(stream *Stream, self RANParameterTestingItem_RanParameterType) {
        if self.RanPChoiceList != nil {
            stream.set_choice(0, 2, 1, 4)
            self.RanPChoiceList.Pack(stream)//2
        } else if self.RanPChoiceStructure != nil {
            stream.set_choice(1, 2, 1, 4)
            self.RanPChoiceStructure.Pack(stream)//2
        } else if self.RanPChoiceElementTrue != nil {
            stream.set_choice(2, 2, 1, 4)
            self.RanPChoiceElementTrue.Pack(stream)//2
        } else if self.RanPChoiceElementFalse != nil {
            stream.set_choice(3, 2, 1, 4)
            self.RanPChoiceElementFalse.Pack(stream)//2
        }

    }
    Pack_ranParameterType(stream, self.RanParameterType) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANParameterTestingItemChoiceList struct { // [{'type': 'RANParameter-Testing-LIST', 'name': 'ranParameter-List'}, None]
    RanParameterList RANParameterTestingLIST
}

func (self * RANParameterTestingItemChoiceList) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterList.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterTestingItemChoiceList) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterList.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANParameterTestingItemChoiceStructure struct { // [{'type': 'RANParameter-Testing-STRUCTURE', 'name': 'ranParameter-Structure'}, None]
    RanParameterStructure RANParameterTestingSTRUCTURE
}

func (self * RANParameterTestingItemChoiceStructure) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterStructure.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterTestingItemChoiceStructure) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterStructure.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANParameterTestingItemChoiceElementTrue struct { // [{'type': 'RANParameter-Value', 'name': 'ranParameter-value'}, None]
    RanParametervalue RANParameterValue
}

func (self * RANParameterTestingItemChoiceElementTrue) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParametervalue.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterTestingItemChoiceElementTrue) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParametervalue.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANParameterTestingItemChoiceElementFalse struct { // [{'type': 'RANParameter-TestingCondition', 'name': 'ranParameter-TestCondition'}, {'type': 'RANParameter-Value', 'name': 'ranParameter-Value', 'optional': True}, {'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}, None]
    RanParameterTestCondition RANParameterTestingCondition
    RanParameterValue *RANParameterValue
    LogicalOR *LogicalOR
}

func (self * RANParameterTestingItemChoiceElementFalse) Unpack(stream *Stream) {
    ranParameterValue_flag := 0x00000002
    logicalOR_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RanParameterTestCondition.Unpack(stream)// p8
    if (ranParameterValue_flag & _flags) == ranParameterValue_flag { //cond2
        self.RanParameterValue = &RANParameterValue{}//7{'type': 'RANParameter-Value', 'name': 'ranParameter-Value', 'optional': True}
        self.RanParameterValue.Unpack(stream)// p8
    }
    if (logicalOR_flag & _flags) == logicalOR_flag { //cond2
        self.LogicalOR = &LogicalOR{}//7{'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}
        self.LogicalOR.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANParameterTestingItemChoiceElementFalse) Pack(stream *Stream) {
    const ranParameterValue_flag uint = 0x00000002
    const logicalOR_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterTestCondition.Pack(stream)
    if self.RanParameterValue != nil { 
        _flags |= ranParameterValue_flag
        self.RanParameterValue.Pack(stream)
    }//end of optional
    if self.LogicalOR != nil { 
        _flags |= logicalOR_flag
        self.LogicalOR.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

func (self *RANParameterTestingLIST) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65535)
    _size += 1
    self.Items = make([]RANParameterTestingItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *RANParameterTestingLIST) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65535)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type RANParameterTestingLIST struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-Testing-Item'}, 'size': [(1, 'maxnoofItemsinList')]}
    Items []RANParameterTestingItem
}

func (self *RANParameterTestingSTRUCTURE) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65535)
    _size += 1
    self.Items = make([]RANParameterTestingItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *RANParameterTestingSTRUCTURE) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65535)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type RANParameterTestingSTRUCTURE struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANParameter-Testing-Item'}, 'size': [(1, 'maxnoofParametersinStructure')]}
    Items []RANParameterTestingItem
}

type UEGroupDefinition_UeGroupDefinitionIdentifierLIST struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'UEGroupDefinitionIdentifier-Item'}, 'size': [(1, 'maxGroupDefinitionIdentifierParameters')], 'name': 'ueGroupDefinitionIdentifier-LIST'}
    Items []UEGroupDefinitionIdentifierItem
}
type UEGroupDefinition struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'UEGroupDefinitionIdentifier-Item'}, 'size': [(1, 'maxGroupDefinitionIdentifierParameters')], 'name': 'ueGroupDefinitionIdentifier-LIST'}, None]
    UeGroupDefinitionIdentifierLIST UEGroupDefinition_UeGroupDefinitionIdentifierLIST
}

func (self * UEGroupDefinition) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ueGroupDefinitionIdentifierLIST = func(stream *Stream, self *UEGroupDefinition_UeGroupDefinitionIdentifierLIST){// Seq6 UEGroupDefinition {'type': 'SEQUENCE OF', 'element': {'type': 'UEGroupDefinitionIdentifier-Item'}, 'size': [(1, 'maxGroupDefinitionIdentifierParameters')], 'name': 'ueGroupDefinitionIdentifier-LIST'}
        _size := stream.get_listsize(255)
        _size += 1
        self.Items = make([]UEGroupDefinitionIdentifierItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ueGroupDefinitionIdentifierLIST(stream, &self.UeGroupDefinitionIdentifierLIST)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEGroupDefinition) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ueGroupDefinitionIdentifierLIST = func(stream *Stream, self UEGroupDefinition_UeGroupDefinitionIdentifierLIST) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 255)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ueGroupDefinitionIdentifierLIST(stream, self.UeGroupDefinitionIdentifierLIST) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type UEGroupDefinitionIdentifierItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, {'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
    LogicalOR *LogicalOR
}

func (self * UEGroupDefinitionIdentifierItem) Unpack(stream *Stream) {
    logicalOR_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    if (logicalOR_flag & _flags) == logicalOR_flag { //cond2
        self.LogicalOR = &LogicalOR{}//7{'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}
        self.LogicalOR.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEGroupDefinitionIdentifierItem) Pack(stream *Stream) {
    const logicalOR_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if self.LogicalOR != nil { 
        _flags |= logicalOR_flag
        self.LogicalOR.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANCallProcessID struct {
  Value uint64
}
func (self *RANCallProcessID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967295, 33, 1, 1)
}
func (self * RANCallProcessID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967295, 33, 1, 1)
}
type RICCallProcessTypeID struct {
  Value uint64
}
func (self *RICCallProcessTypeID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65535, 17, 1, 1)
}
func (self * RICCallProcessTypeID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65535, 17, 1, 1)
}
type RICCallProcessTypeName struct {
  Value string
}
func (self *RICCallProcessTypeName) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in RIC-CallProcessType-Name")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *RICCallProcessTypeName) Pack(st *Stream) {
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
type RICCallProcessBreakpointID struct {
  Value uint64
}
func (self *RICCallProcessBreakpointID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65535, 17, 1, 1)
}
func (self * RICCallProcessBreakpointID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65535, 17, 1, 1)
}
type RICCallProcessBreakpointName struct {
  Value string
}
func (self *RICCallProcessBreakpointName) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in RIC-CallProcessBreakpoint-Name")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *RICCallProcessBreakpointName) Pack(st *Stream) {
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
type RICControlActionID struct {
  Value uint64
}
func (self *RICControlActionID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65535, 17, 1, 1)
}
func (self * RICControlActionID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65535, 17, 1, 1)
}
type RICControlActionName struct {
  Value string
}
func (self *RICControlActionName) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in RIC-ControlAction-Name")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *RICControlActionName) Pack(st *Stream) {
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
type RICEventTriggerConditionID struct {
  Value uint64
}
func (self *RICEventTriggerConditionID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65535, 17, 1, 1)
}
func (self * RICEventTriggerConditionID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65535, 17, 1, 1)
}
type RICEventTriggerUEID struct {
  Value uint64
}
func (self *RICEventTriggerUEID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65535, 17, 1, 1)
}
func (self * RICEventTriggerUEID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65535, 17, 1, 1)
}
type RICEventTriggerUEeventID struct {
  Value uint64
}
func (self *RICEventTriggerUEeventID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65535, 17, 1, 1)
}
func (self * RICEventTriggerUEeventID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65535, 17, 1, 1)
}
type RICEventTriggerCellID struct {
  Value uint64
}
func (self *RICEventTriggerCellID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65535, 17, 1, 1)
}
func (self * RICEventTriggerCellID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65535, 17, 1, 1)
}
type RICInsertIndicationID struct {
  Value uint64
}
func (self *RICInsertIndicationID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65535, 17, 1, 1)
}
func (self * RICInsertIndicationID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65535, 17, 1, 1)
}
type RICInsertIndicationName struct {
  Value string
}
func (self *RICInsertIndicationName) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in RIC-InsertIndication-Name")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *RICInsertIndicationName) Pack(st *Stream) {
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
type UEGroupID struct {
  Value uint64
}
func (self *UEGroupID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65535, 17, 1, 1)
}
func (self * UEGroupID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65535, 17, 1, 1)
}
type EntityFilterID struct {
  Value uint64
}
func (self *EntityFilterID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(255, 9, 1, 1)
}
func (self * EntityFilterID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 255, 9, 1, 1)
}
type RICPolicyAction_RanParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-PolicyAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranParameters-List', 'optional': True}
    Items []RICPolicyActionRANParameterItem
}
type RICPolicyAction struct { // [{'type': 'RIC-ControlAction-ID', 'name': 'ric-PolicyAction-ID'}, {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-PolicyAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranParameters-List', 'optional': True}, None, {'type': 'ENUMERATED', 'values': [('accept', 0), ('reject', 1), None], 'name': 'ric-PolicyDecision', 'optional': True}]
    RicPolicyActionID RICControlActionID
    RanParametersList *RICPolicyAction_RanParametersList
    RicPolicyDecision *ENUMERATED
}

func (self * RICPolicyAction) Unpack(stream *Stream) {
    ranParametersList_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RicPolicyActionID.Unpack(stream)// p8
    if (ranParametersList_flag & _flags) == ranParametersList_flag { //cond1
        var Unpack_ranParametersList = func(stream *Stream, self *RICPolicyAction_RanParametersList){// Seq6 RICPolicyAction {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-PolicyAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranParameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]RICPolicyActionRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanParametersList = &RICPolicyAction_RanParametersList{}//3
        Unpack_ranParametersList(stream, self.RanParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-PolicyAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranParameters-List', 'optional': True}
    }
    if (_flags & ext_flag) > 0 {  //ric-PolicyDecision
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        var Unpack_ricPolicyDecision = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(2, 2, 1)
        }
        self.RicPolicyDecision = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('accept', 0), ('reject', 1), None], 'name': 'ric-PolicyDecision', 'optional': True}
        Unpack_ricPolicyDecision(stream, self.RicPolicyDecision)// p1 {'type': 'ENUMERATED', 'values': [('accept', 0), ('reject', 1), None], 'name': 'ric-PolicyDecision', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICPolicyAction) Pack(stream *Stream) {
    const ranParametersList_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    const ricPolicyDecision_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicPolicyActionID.Pack(stream)
    if self.RanParametersList != nil { //YY
        _flags |= ranParametersList_flag
        var Pack_ranParametersList = func(stream *Stream, self RICPolicyAction_RanParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranParametersList(stream, *self.RanParametersList) //f1
    }//end of optional
    if self.RicPolicyDecision != nil { //YY
//extension Group:      ric-PolicyDecision 1 
    if _extPresent == false {
        stream.set_bits(0, 7)
        _extReserve = stream.reserve_flags(1)
        _extPresent = true
        _flags |= uint(ext_flag)
    }
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ricPolicyDecision_flag
        var Pack_ricPolicyDecision = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 2, 2, 1)
        }
        Pack_ricPolicyDecision(stream, *self.RicPolicyDecision) //f1
    stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type RICPolicyActionRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
}

func (self * RICPolicyActionRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICPolicyActionRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type UEFilterID struct {
  Value uint64
}
func (self *UEFilterID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65535, 17, 1, 1)
}
func (self * UEFilterID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65535, 17, 1, 1)
}
type AssociatedUEInfo_AssociatedUEInfoList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'Associated-UE-Info-Item'}, 'size': [(1, 'maxnoofUEInfo')], 'name': 'associatedUEInfo-List'}
    Items []AssociatedUEInfoItem
}
type AssociatedUEInfo struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'Associated-UE-Info-Item'}, 'size': [(1, 'maxnoofUEInfo')], 'name': 'associatedUEInfo-List'}, None]
    AssociatedUEInfoList AssociatedUEInfo_AssociatedUEInfoList
}

func (self * AssociatedUEInfo) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_associatedUEInfoList = func(stream *Stream, self *AssociatedUEInfo_AssociatedUEInfoList){// Seq6 AssociatedUEInfo {'type': 'SEQUENCE OF', 'element': {'type': 'Associated-UE-Info-Item'}, 'size': [(1, 'maxnoofUEInfo')], 'name': 'associatedUEInfo-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]AssociatedUEInfoItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_associatedUEInfoList(stream, &self.AssociatedUEInfoList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AssociatedUEInfo) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_associatedUEInfoList = func(stream *Stream, self AssociatedUEInfo_AssociatedUEInfoList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_associatedUEInfoList(stream, self.AssociatedUEInfoList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type AssociatedUEInfoItem_UeType struct { //[{'type': 'EventTrigger-UE-Info-Item-Choice-Individual', 'name': 'ueType-Choice-Individual'}, {'type': 'EventTrigger-UE-Info-Item-Choice-Group', 'name': 'ueType-Choice-Group'}, None]
    UeTypeChoiceIndividual *EventTriggerUEInfoItemChoiceIndividual
    UeTypeChoiceGroup *EventTriggerUEInfoItemChoiceGroup
} // AssociatedUEInfoItem_UeType

type AssociatedUEInfoItem struct { // [{'type': 'UE-Filter-ID', 'name': 'ueFilterID'}, {'type': 'CHOICE', 'members': [{'type': 'EventTrigger-UE-Info-Item-Choice-Individual', 'name': 'ueType-Choice-Individual'}, {'type': 'EventTrigger-UE-Info-Item-Choice-Group', 'name': 'ueType-Choice-Group'}, None], 'name': 'ueType'}, {'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}, None]
    UeFilterID UEFilterID
    UeType AssociatedUEInfoItem_UeType
    LogicalOR *LogicalOR
}

func (self * AssociatedUEInfoItem) Unpack(stream *Stream) {
    logicalOR_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.UeFilterID.Unpack(stream)// p8
    var Unpack_ueType = func(stream *Stream, self *AssociatedUEInfoItem_UeType) {
        //coptions := []string{"ueType-Choice-Individual","ueType-Choice-Group"}
        choice := stream.get_choice(1, 1, 2)
        choice_len := 0
        choice_loc := 0
        if choice >= 2 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in AssociatedUEInfoItem_UeType\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.UeTypeChoiceIndividual = &EventTriggerUEInfoItemChoiceIndividual{}//cho6
            self.UeTypeChoiceIndividual.Unpack(stream)
        } else if choice == 1 { //ch2
            self.UeTypeChoiceGroup = &EventTriggerUEInfoItemChoiceGroup{}//cho6
            self.UeTypeChoiceGroup.Unpack(stream)
        }//end of if else

        if choice >= 2 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_ueType(stream, &self.UeType)// p2
    if (logicalOR_flag & _flags) == logicalOR_flag { //cond2
        self.LogicalOR = &LogicalOR{}//7{'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}
        self.LogicalOR.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AssociatedUEInfoItem) Pack(stream *Stream) {
    const logicalOR_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UeFilterID.Pack(stream)
    var Pack_ueType = func(stream *Stream, self AssociatedUEInfoItem_UeType) {
        if self.UeTypeChoiceIndividual != nil {
            stream.set_choice(0, 1, 1, 2)
            self.UeTypeChoiceIndividual.Pack(stream)//2
        } else if self.UeTypeChoiceGroup != nil {
            stream.set_choice(1, 1, 1, 2)
            self.UeTypeChoiceGroup.Pack(stream)//2
        }

    }
    Pack_ueType(stream, self.UeType) //f2
    if self.LogicalOR != nil { 
        _flags |= logicalOR_flag
        self.LogicalOR.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCEventTrigger_RiceventTriggerformats struct { //[{'type': 'E2SM-RC-EventTrigger-Format1', 'name': 'eventTrigger-Format1'}, {'type': 'E2SM-RC-EventTrigger-Format2', 'name': 'eventTrigger-Format2'}, {'type': 'E2SM-RC-EventTrigger-Format3', 'name': 'eventTrigger-Format3'}, {'type': 'E2SM-RC-EventTrigger-Format4', 'name': 'eventTrigger-Format4'}, {'type': 'NULL', 'name': 'eventTrigger-Format5'}, None]
    EventTriggerFormat1 *E2SMRCEventTriggerFormat1
    EventTriggerFormat2 *E2SMRCEventTriggerFormat2
    EventTriggerFormat3 *E2SMRCEventTriggerFormat3
    EventTriggerFormat4 *E2SMRCEventTriggerFormat4
    EventTriggerFormat5 *NULL
} // E2SMRCEventTrigger_RiceventTriggerformats

type E2SMRCEventTrigger struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-EventTrigger-Format1', 'name': 'eventTrigger-Format1'}, {'type': 'E2SM-RC-EventTrigger-Format2', 'name': 'eventTrigger-Format2'}, {'type': 'E2SM-RC-EventTrigger-Format3', 'name': 'eventTrigger-Format3'}, {'type': 'E2SM-RC-EventTrigger-Format4', 'name': 'eventTrigger-Format4'}, {'type': 'NULL', 'name': 'eventTrigger-Format5'}, None], 'name': 'ric-eventTrigger-formats'}, None]
    RiceventTriggerformats E2SMRCEventTrigger_RiceventTriggerformats
}

func (self * E2SMRCEventTrigger) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_riceventTriggerformats = func(stream *Stream, self *E2SMRCEventTrigger_RiceventTriggerformats) {
        //coptions := []string{"eventTrigger-Format1","eventTrigger-Format2","eventTrigger-Format3","eventTrigger-Format4","eventTrigger-Format5","Unknown","Unknown","Unknown"}
        choice := stream.get_choice(3, 1, 5)
        choice_len := 0
        choice_loc := 0
        if choice >= 5 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCEventTrigger_RiceventTriggerformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.EventTriggerFormat1 = &E2SMRCEventTriggerFormat1{}//cho6
            self.EventTriggerFormat1.Unpack(stream)
        } else if choice == 1 { //ch2
            self.EventTriggerFormat2 = &E2SMRCEventTriggerFormat2{}//cho6
            self.EventTriggerFormat2.Unpack(stream)
        } else if choice == 2 { //ch2
            self.EventTriggerFormat3 = &E2SMRCEventTriggerFormat3{}//cho6
            self.EventTriggerFormat3.Unpack(stream)
        } else if choice == 3 { //ch2
            self.EventTriggerFormat4 = &E2SMRCEventTriggerFormat4{}//cho6
            self.EventTriggerFormat4.Unpack(stream)
        } else if choice == 4 { //ch2
            var Unpack_eventTriggerFormat5 = func(st *Stream, self *NULL){
                st.parsef_Null()
            }
            self.EventTriggerFormat5 = &NULL{}//cho5
            Unpack_eventTriggerFormat5(stream, self.EventTriggerFormat5);
        }//end of if else

        if choice >= 5 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_riceventTriggerformats(stream, &self.RiceventTriggerformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCEventTrigger) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_riceventTriggerformats = func(stream *Stream, self E2SMRCEventTrigger_RiceventTriggerformats) {
        if self.EventTriggerFormat1 != nil {
            stream.set_choice(0, 3, 1, 5)
            self.EventTriggerFormat1.Pack(stream)//2
        } else if self.EventTriggerFormat2 != nil {
            stream.set_choice(1, 3, 1, 5)
            self.EventTriggerFormat2.Pack(stream)//2
        } else if self.EventTriggerFormat3 != nil {
            stream.set_choice(2, 3, 1, 5)
            self.EventTriggerFormat3.Pack(stream)//2
        } else if self.EventTriggerFormat4 != nil {
            stream.set_choice(3, 3, 1, 5)
            self.EventTriggerFormat4.Pack(stream)//2
        } else if self.EventTriggerFormat5 != nil {
            stream.set_choice(4, 3, 1, 5)
            var Pack_eventTriggerFormat5 = func(st *Stream, self NULL){
                st.formatf_Null()
            }
            Pack_eventTriggerFormat5(stream, *self.EventTriggerFormat5)//3
        }

    }
    Pack_riceventTriggerformats(stream, self.RiceventTriggerformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCEventTriggerFormat1_MessageList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EventTrigger-Format1-Item'}, 'size': [(1, 'maxnoofMessages')], 'name': 'message-List'}
    Items []E2SMRCEventTriggerFormat1Item
}
type E2SMRCEventTriggerFormat1 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EventTrigger-Format1-Item'}, 'size': [(1, 'maxnoofMessages')], 'name': 'message-List'}, {'type': 'EventTrigger-UE-Info', 'name': 'globalAssociatedUEInfo', 'optional': True}, None]
    MessageList E2SMRCEventTriggerFormat1_MessageList
    GlobalAssociatedUEInfo *EventTriggerUEInfo
}

func (self * E2SMRCEventTriggerFormat1) Unpack(stream *Stream) {
    globalAssociatedUEInfo_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_messageList = func(stream *Stream, self *E2SMRCEventTriggerFormat1_MessageList){// Seq6 E2SMRCEventTriggerFormat1 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EventTrigger-Format1-Item'}, 'size': [(1, 'maxnoofMessages')], 'name': 'message-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCEventTriggerFormat1Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_messageList(stream, &self.MessageList)// p2
    if (globalAssociatedUEInfo_flag & _flags) == globalAssociatedUEInfo_flag { //cond2
        self.GlobalAssociatedUEInfo = &EventTriggerUEInfo{}//7{'type': 'EventTrigger-UE-Info', 'name': 'globalAssociatedUEInfo', 'optional': True}
        self.GlobalAssociatedUEInfo.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCEventTriggerFormat1) Pack(stream *Stream) {
    const globalAssociatedUEInfo_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_messageList = func(stream *Stream, self E2SMRCEventTriggerFormat1_MessageList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_messageList(stream, self.MessageList) //f2
    if self.GlobalAssociatedUEInfo != nil { 
        _flags |= globalAssociatedUEInfo_flag
        self.GlobalAssociatedUEInfo.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCEventTriggerFormat1Item struct { // [{'type': 'RIC-EventTriggerCondition-ID', 'name': 'ric-eventTriggerCondition-ID'}, {'type': 'MessageType-Choice', 'name': 'messageType'}, {'type': 'ENUMERATED', 'values': [('incoming', 0), ('outgoing', 1), None], 'name': 'messageDirection', 'optional': True}, {'type': 'EventTrigger-UE-Info', 'name': 'associatedUEInfo', 'optional': True}, {'type': 'EventTrigger-UEevent-Info', 'name': 'associatedUEEvent', 'optional': True}, {'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}, None]
    RiceventTriggerConditionID RICEventTriggerConditionID
    MessageType MessageTypeChoice
    MessageDirection *ENUMERATED
    AssociatedUEInfo *EventTriggerUEInfo
    AssociatedUEEvent *EventTriggerUEeventInfo
    LogicalOR *LogicalOR
}

func (self * E2SMRCEventTriggerFormat1Item) Unpack(stream *Stream) {
    messageDirection_flag := 0x00000002
    associatedUEInfo_flag := 0x00000004
    associatedUEEvent_flag := 0x00000008
    logicalOR_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.RiceventTriggerConditionID.Unpack(stream)// p8
    self.MessageType.Unpack(stream)// p8
    if (messageDirection_flag & _flags) == messageDirection_flag { //cond1
        var Unpack_messageDirection = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(2, 2, 1)
        }
        self.MessageDirection = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('incoming', 0), ('outgoing', 1), None], 'name': 'messageDirection', 'optional': True}
        Unpack_messageDirection(stream, self.MessageDirection)// p1 {'type': 'ENUMERATED', 'values': [('incoming', 0), ('outgoing', 1), None], 'name': 'messageDirection', 'optional': True}
    }
    if (associatedUEInfo_flag & _flags) == associatedUEInfo_flag { //cond2
        self.AssociatedUEInfo = &EventTriggerUEInfo{}//7{'type': 'EventTrigger-UE-Info', 'name': 'associatedUEInfo', 'optional': True}
        self.AssociatedUEInfo.Unpack(stream)// p8
    }
    if (associatedUEEvent_flag & _flags) == associatedUEEvent_flag { //cond2
        self.AssociatedUEEvent = &EventTriggerUEeventInfo{}//7{'type': 'EventTrigger-UEevent-Info', 'name': 'associatedUEEvent', 'optional': True}
        self.AssociatedUEEvent.Unpack(stream)// p8
    }
    if (logicalOR_flag & _flags) == logicalOR_flag { //cond2
        self.LogicalOR = &LogicalOR{}//7{'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}
        self.LogicalOR.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCEventTriggerFormat1Item) Pack(stream *Stream) {
    const messageDirection_flag uint = 0x00000002
    const associatedUEInfo_flag uint = 0x00000004
    const associatedUEEvent_flag uint = 0x00000008
    const logicalOR_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RiceventTriggerConditionID.Pack(stream)
    self.MessageType.Pack(stream)
    if self.MessageDirection != nil { //YY
        _flags |= messageDirection_flag
        var Pack_messageDirection = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 2, 2, 1)
        }
        Pack_messageDirection(stream, *self.MessageDirection) //f1
    }//end of optional
    if self.AssociatedUEInfo != nil { 
        _flags |= associatedUEInfo_flag
        self.AssociatedUEInfo.Pack(stream)
    }//end of optional
    if self.AssociatedUEEvent != nil { 
        _flags |= associatedUEEvent_flag
        self.AssociatedUEEvent.Pack(stream)
    }//end of optional
    if self.LogicalOR != nil { 
        _flags |= logicalOR_flag
        self.LogicalOR.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *MessageTypeChoice)Unpack(stream *Stream) {
    //coptions := []string{"messageType-Choice-NI","messageType-Choice-RRC"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in MessageTypeChoice\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.MessageTypeChoiceNI = &MessageTypeChoiceNI{}//cho6
        self.MessageTypeChoiceNI.Unpack(stream)
    } else if choice == 1 { //ch2
        self.MessageTypeChoiceRRC = &MessageTypeChoiceRRC{}//cho6
        self.MessageTypeChoiceRRC.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * MessageTypeChoice) Pack(stream *Stream) {
    if self.MessageTypeChoiceNI != nil {
        stream.set_choice(0, 1, 1, 2)
        self.MessageTypeChoiceNI.Pack(stream)//2
    } else if self.MessageTypeChoiceRRC != nil {
        stream.set_choice(1, 1, 1, 2)
        self.MessageTypeChoiceRRC.Pack(stream)//2
    }

}
type MessageTypeChoice struct { //[{'type': 'MessageType-Choice-NI', 'name': 'messageType-Choice-NI'}, {'type': 'MessageType-Choice-RRC', 'name': 'messageType-Choice-RRC'}, None]
    MessageTypeChoiceNI *MessageTypeChoiceNI
    MessageTypeChoiceRRC *MessageTypeChoiceRRC
} // MessageTypeChoice

type MessageTypeChoiceNI struct { // [{'type': 'InterfaceType', 'name': 'nI-Type'}, {'type': 'InterfaceIdentifier', 'name': 'nI-Identifier', 'optional': True}, {'type': 'Interface-MessageID', 'name': 'nI-Message', 'optional': True}, None]
    NIType InterfaceType
    NIIdentifier *InterfaceIdentifier
    NIMessage *InterfaceMessageID
}

func (self * MessageTypeChoiceNI) Unpack(stream *Stream) {
    nIIdentifier_flag := 0x00000002
    nIMessage_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.NIType.Unpack(stream)// p8
    if (nIIdentifier_flag & _flags) == nIIdentifier_flag { //cond2
        self.NIIdentifier = &InterfaceIdentifier{}//7{'type': 'InterfaceIdentifier', 'name': 'nI-Identifier', 'optional': True}
        self.NIIdentifier.Unpack(stream)// p8
    }
    if (nIMessage_flag & _flags) == nIMessage_flag { //cond2
        self.NIMessage = &InterfaceMessageID{}//7{'type': 'Interface-MessageID', 'name': 'nI-Message', 'optional': True}
        self.NIMessage.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MessageTypeChoiceNI) Pack(stream *Stream) {
    const nIIdentifier_flag uint = 0x00000002
    const nIMessage_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.NIType.Pack(stream)
    if self.NIIdentifier != nil { 
        _flags |= nIIdentifier_flag
        self.NIIdentifier.Pack(stream)
    }//end of optional
    if self.NIMessage != nil { 
        _flags |= nIMessage_flag
        self.NIMessage.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type MessageTypeChoiceRRC struct { // [{'type': 'RRC-MessageID', 'name': 'rRC-Message'}, None]
    RRCMessage RRCMessageID
}

func (self * MessageTypeChoiceRRC) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RRCMessage.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MessageTypeChoiceRRC) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RRCMessage.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCEventTriggerFormat2 struct { // [{'type': 'RIC-CallProcessType-ID', 'name': 'ric-callProcessType-ID'}, {'type': 'RIC-CallProcessBreakpoint-ID', 'name': 'ric-callProcessBreakpoint-ID'}, {'type': 'RANParameter-Testing', 'name': 'associatedE2NodeInfo', 'optional': True}, {'type': 'EventTrigger-UE-Info', 'name': 'associatedUEInfo', 'optional': True}, None]
    RiccallProcessTypeID RICCallProcessTypeID
    RiccallProcessBreakpointID RICCallProcessBreakpointID
    AssociatedE2NodeInfo *RANParameterTesting
    AssociatedUEInfo *EventTriggerUEInfo
}

func (self * E2SMRCEventTriggerFormat2) Unpack(stream *Stream) {
    associatedE2NodeInfo_flag := 0x00000002
    associatedUEInfo_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RiccallProcessTypeID.Unpack(stream)// p8
    self.RiccallProcessBreakpointID.Unpack(stream)// p8
    if (associatedE2NodeInfo_flag & _flags) == associatedE2NodeInfo_flag { //cond2
        self.AssociatedE2NodeInfo = &RANParameterTesting{}//7{'type': 'RANParameter-Testing', 'name': 'associatedE2NodeInfo', 'optional': True}
        self.AssociatedE2NodeInfo.Unpack(stream)// p8
    }
    if (associatedUEInfo_flag & _flags) == associatedUEInfo_flag { //cond2
        self.AssociatedUEInfo = &EventTriggerUEInfo{}//7{'type': 'EventTrigger-UE-Info', 'name': 'associatedUEInfo', 'optional': True}
        self.AssociatedUEInfo.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCEventTriggerFormat2) Pack(stream *Stream) {
    const associatedE2NodeInfo_flag uint = 0x00000002
    const associatedUEInfo_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RiccallProcessTypeID.Pack(stream)
    self.RiccallProcessBreakpointID.Pack(stream)
    if self.AssociatedE2NodeInfo != nil { 
        _flags |= associatedE2NodeInfo_flag
        self.AssociatedE2NodeInfo.Pack(stream)
    }//end of optional
    if self.AssociatedUEInfo != nil { 
        _flags |= associatedUEInfo_flag
        self.AssociatedUEInfo.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type E2SMRCEventTriggerFormat3_E2NodeInfoChangeList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EventTrigger-Format3-Item'}, 'size': [(1, 'maxnoofE2InfoChanges')], 'name': 'e2NodeInfoChange-List'}
    Items []E2SMRCEventTriggerFormat3Item
}
type E2SMRCEventTriggerFormat3 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EventTrigger-Format3-Item'}, 'size': [(1, 'maxnoofE2InfoChanges')], 'name': 'e2NodeInfoChange-List'}, None]
    E2NodeInfoChangeList E2SMRCEventTriggerFormat3_E2NodeInfoChangeList
}

func (self * E2SMRCEventTriggerFormat3) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_e2NodeInfoChangeList = func(stream *Stream, self *E2SMRCEventTriggerFormat3_E2NodeInfoChangeList){// Seq6 E2SMRCEventTriggerFormat3 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EventTrigger-Format3-Item'}, 'size': [(1, 'maxnoofE2InfoChanges')], 'name': 'e2NodeInfoChange-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCEventTriggerFormat3Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_e2NodeInfoChangeList(stream, &self.E2NodeInfoChangeList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCEventTriggerFormat3) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_e2NodeInfoChangeList = func(stream *Stream, self E2SMRCEventTriggerFormat3_E2NodeInfoChangeList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_e2NodeInfoChangeList(stream, self.E2NodeInfoChangeList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCEventTriggerFormat3Item struct { // [{'type': 'RIC-EventTriggerCondition-ID', 'name': 'ric-eventTriggerCondition-ID'}, {'type': 'INTEGER', 'restricted-to': [(1, 512), None], 'name': 'e2NodeInfoChange-ID'}, {'type': 'EventTrigger-Cell-Info', 'name': 'associatedCellInfo', 'optional': True}, {'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}, None]
    RiceventTriggerConditionID RICEventTriggerConditionID
    E2NodeInfoChangeID INTEGER
    AssociatedCellInfo *EventTriggerCellInfo
    LogicalOR *LogicalOR
}

func (self * E2SMRCEventTriggerFormat3Item) Unpack(stream *Stream) {
    associatedCellInfo_flag := 0x00000002
    logicalOR_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RiceventTriggerConditionID.Unpack(stream)// p8
    var Unpack_e2NodeInfoChangeID = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(512, 10, 1, 1)
    }
    Unpack_e2NodeInfoChangeID(stream, &self.E2NodeInfoChangeID)// p2
    if (associatedCellInfo_flag & _flags) == associatedCellInfo_flag { //cond2
        self.AssociatedCellInfo = &EventTriggerCellInfo{}//7{'type': 'EventTrigger-Cell-Info', 'name': 'associatedCellInfo', 'optional': True}
        self.AssociatedCellInfo.Unpack(stream)// p8
    }
    if (logicalOR_flag & _flags) == logicalOR_flag { //cond2
        self.LogicalOR = &LogicalOR{}//7{'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}
        self.LogicalOR.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCEventTriggerFormat3Item) Pack(stream *Stream) {
    const associatedCellInfo_flag uint = 0x00000002
    const logicalOR_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RiceventTriggerConditionID.Pack(stream)
    var Pack_e2NodeInfoChangeID = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 512, 10, 1, 1)
    }
    Pack_e2NodeInfoChangeID(stream, self.E2NodeInfoChangeID) //f2
    if self.AssociatedCellInfo != nil { 
        _flags |= associatedCellInfo_flag
        self.AssociatedCellInfo.Pack(stream)
    }//end of optional
    if self.LogicalOR != nil { 
        _flags |= logicalOR_flag
        self.LogicalOR.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type E2SMRCEventTriggerFormat4_UEInfoChangeList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EventTrigger-Format4-Item'}, 'size': [(1, 'maxnoofUEInfoChanges')], 'name': 'uEInfoChange-List'}
    Items []E2SMRCEventTriggerFormat4Item
}
type E2SMRCEventTriggerFormat4 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EventTrigger-Format4-Item'}, 'size': [(1, 'maxnoofUEInfoChanges')], 'name': 'uEInfoChange-List'}, None]
    UEInfoChangeList E2SMRCEventTriggerFormat4_UEInfoChangeList
}

func (self * E2SMRCEventTriggerFormat4) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_uEInfoChangeList = func(stream *Stream, self *E2SMRCEventTriggerFormat4_UEInfoChangeList){// Seq6 E2SMRCEventTriggerFormat4 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EventTrigger-Format4-Item'}, 'size': [(1, 'maxnoofUEInfoChanges')], 'name': 'uEInfoChange-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCEventTriggerFormat4Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_uEInfoChangeList(stream, &self.UEInfoChangeList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCEventTriggerFormat4) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_uEInfoChangeList = func(stream *Stream, self E2SMRCEventTriggerFormat4_UEInfoChangeList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_uEInfoChangeList(stream, self.UEInfoChangeList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCEventTriggerFormat4Item struct { // [{'type': 'RIC-EventTriggerCondition-ID', 'name': 'ric-eventTriggerCondition-ID'}, {'type': 'TriggerType-Choice', 'name': 'triggerType'}, {'type': 'EventTrigger-UE-Info', 'name': 'associatedUEInfo', 'optional': True}, {'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}, None]
    RiceventTriggerConditionID RICEventTriggerConditionID
    TriggerType TriggerTypeChoice
    AssociatedUEInfo *EventTriggerUEInfo
    LogicalOR *LogicalOR
}

func (self * E2SMRCEventTriggerFormat4Item) Unpack(stream *Stream) {
    associatedUEInfo_flag := 0x00000002
    logicalOR_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RiceventTriggerConditionID.Unpack(stream)// p8
    self.TriggerType.Unpack(stream)// p8
    if (associatedUEInfo_flag & _flags) == associatedUEInfo_flag { //cond2
        self.AssociatedUEInfo = &EventTriggerUEInfo{}//7{'type': 'EventTrigger-UE-Info', 'name': 'associatedUEInfo', 'optional': True}
        self.AssociatedUEInfo.Unpack(stream)// p8
    }
    if (logicalOR_flag & _flags) == logicalOR_flag { //cond2
        self.LogicalOR = &LogicalOR{}//7{'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}
        self.LogicalOR.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCEventTriggerFormat4Item) Pack(stream *Stream) {
    const associatedUEInfo_flag uint = 0x00000002
    const logicalOR_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RiceventTriggerConditionID.Pack(stream)
    self.TriggerType.Pack(stream)
    if self.AssociatedUEInfo != nil { 
        _flags |= associatedUEInfo_flag
        self.AssociatedUEInfo.Pack(stream)
    }//end of optional
    if self.LogicalOR != nil { 
        _flags |= logicalOR_flag
        self.LogicalOR.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

func (self *TriggerTypeChoice)Unpack(stream *Stream) {
    //coptions := []string{"triggerType-Choice-RRCstate","triggerType-Choice-UEID","triggerType-Choice-L2state","Unknown","triggerType-Choice-UEcontext","triggerType-Choice-L2MACschChg"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in TriggerTypeChoice\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.TriggerTypeChoiceRRCstate = &TriggerTypeChoiceRRCstate{}//cho6
        self.TriggerTypeChoiceRRCstate.Unpack(stream)
    } else if choice == 1 { //ch2
        self.TriggerTypeChoiceUEID = &TriggerTypeChoiceUEID{}//cho6
        self.TriggerTypeChoiceUEID.Unpack(stream)
    } else if choice == 2 { //ch2
        self.TriggerTypeChoiceL2state = &TriggerTypeChoiceL2state{}//cho6
        self.TriggerTypeChoiceL2state.Unpack(stream)
    } else if choice == 4 { //ch2
        self.TriggerTypeChoiceUEcontext = &TriggerTypeChoiceUEcontext{}//cho6
        self.TriggerTypeChoiceUEcontext.Unpack(stream)
    } else if choice == 5 { //ch2
        self.TriggerTypeChoiceL2MACschChg = &TriggerTypeChoiceL2MACschChg{}//cho6
        self.TriggerTypeChoiceL2MACschChg.Unpack(stream)
    }//end of if else

    if choice >= 3 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * TriggerTypeChoice) Pack(stream *Stream) {
    if self.TriggerTypeChoiceRRCstate != nil {
        stream.set_choice(0, 2, 1, 3)
        self.TriggerTypeChoiceRRCstate.Pack(stream)//2
    } else if self.TriggerTypeChoiceUEID != nil {
        stream.set_choice(1, 2, 1, 3)
        self.TriggerTypeChoiceUEID.Pack(stream)//2
    } else if self.TriggerTypeChoiceL2state != nil {
        stream.set_choice(2, 2, 1, 3)
        self.TriggerTypeChoiceL2state.Pack(stream)//2
    } else if self.TriggerTypeChoiceUEcontext != nil {
        stream.set_choice(4, 2, 1, 3)
        lenLoc := stream.reserve_len()
        self.TriggerTypeChoiceUEcontext.Pack(stream)//2
        stream.set_len(lenLoc)
    } else if self.TriggerTypeChoiceL2MACschChg != nil {
        stream.set_choice(5, 2, 1, 3)
        lenLoc := stream.reserve_len()
        self.TriggerTypeChoiceL2MACschChg.Pack(stream)//2
        stream.set_len(lenLoc)
    }

}
type TriggerTypeChoice struct { //[{'type': 'TriggerType-Choice-RRCstate', 'name': 'triggerType-Choice-RRCstate'}, {'type': 'TriggerType-Choice-UEID', 'name': 'triggerType-Choice-UEID'}, {'type': 'TriggerType-Choice-L2state', 'name': 'triggerType-Choice-L2state'}, None, {'type': 'TriggerType-Choice-UEcontext', 'name': 'triggerType-Choice-UEcontext'}, {'type': 'TriggerType-Choice-L2MACschChg', 'name': 'triggerType-Choice-L2MACschChg'}]
    TriggerTypeChoiceRRCstate *TriggerTypeChoiceRRCstate
    TriggerTypeChoiceUEID *TriggerTypeChoiceUEID
    TriggerTypeChoiceL2state *TriggerTypeChoiceL2state
    TriggerTypeChoiceUEcontext *TriggerTypeChoiceUEcontext
    TriggerTypeChoiceL2MACschChg *TriggerTypeChoiceL2MACschChg
} // TriggerTypeChoice

type TriggerTypeChoiceRRCstate_RrcStateList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'TriggerType-Choice-RRCstate-Item'}, 'size': [(1, 'maxnoofRRCstate')], 'name': 'rrcState-List'}
    Items []TriggerTypeChoiceRRCstateItem
}
type TriggerTypeChoiceRRCstate struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'TriggerType-Choice-RRCstate-Item'}, 'size': [(1, 'maxnoofRRCstate')], 'name': 'rrcState-List'}, None]
    RrcStateList TriggerTypeChoiceRRCstate_RrcStateList
}

func (self * TriggerTypeChoiceRRCstate) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_rrcStateList = func(stream *Stream, self *TriggerTypeChoiceRRCstate_RrcStateList){// Seq6 TriggerTypeChoiceRRCstate {'type': 'SEQUENCE OF', 'element': {'type': 'TriggerType-Choice-RRCstate-Item'}, 'size': [(1, 'maxnoofRRCstate')], 'name': 'rrcState-List'}
        _size := stream.get_listsize(8)
        _size += 1
        self.Items = make([]TriggerTypeChoiceRRCstateItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_rrcStateList(stream, &self.RrcStateList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TriggerTypeChoiceRRCstate) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_rrcStateList = func(stream *Stream, self TriggerTypeChoiceRRCstate_RrcStateList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 8)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_rrcStateList(stream, self.RrcStateList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type TriggerTypeChoiceRRCstateItem struct { // [{'type': 'RRC-State', 'name': 'stateChangedTo'}, {'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}, None]
    StateChangedTo RRCState
    LogicalOR *LogicalOR
}

func (self * TriggerTypeChoiceRRCstateItem) Unpack(stream *Stream) {
    logicalOR_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.StateChangedTo.Unpack(stream)// p8
    if (logicalOR_flag & _flags) == logicalOR_flag { //cond2
        self.LogicalOR = &LogicalOR{}//7{'type': 'LogicalOR', 'name': 'logicalOR', 'optional': True}
        self.LogicalOR.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TriggerTypeChoiceRRCstateItem) Pack(stream *Stream) {
    const logicalOR_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.StateChangedTo.Pack(stream)
    if self.LogicalOR != nil { 
        _flags |= logicalOR_flag
        self.LogicalOR.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type TriggerTypeChoiceUEID struct { // [{'type': 'INTEGER', 'restricted-to': [(1, 512), None], 'name': 'ueIDchange-ID'}, None]
    UeIDchangeID INTEGER
}

func (self * TriggerTypeChoiceUEID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ueIDchangeID = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(512, 10, 1, 1)
    }
    Unpack_ueIDchangeID(stream, &self.UeIDchangeID)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TriggerTypeChoiceUEID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ueIDchangeID = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 512, 10, 1, 1)
    }
    Pack_ueIDchangeID(stream, self.UeIDchangeID) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type TriggerTypeChoiceL2state struct { // [{'type': 'RANParameter-Testing', 'name': 'associatedL2variables'}, None]
    AssociatedL2variables RANParameterTesting
}

func (self * TriggerTypeChoiceL2state) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.AssociatedL2variables.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TriggerTypeChoiceL2state) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.AssociatedL2variables.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type TriggerTypeChoiceUEcontext struct { // [{'type': 'RANParameter-Testing', 'name': 'associatedUECtxtVariables'}, None]
    AssociatedUECtxtVariables RANParameterTesting
}

func (self * TriggerTypeChoiceUEcontext) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.AssociatedUECtxtVariables.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TriggerTypeChoiceUEcontext) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.AssociatedUECtxtVariables.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type TriggerTypeChoiceL2MACschChg struct { // [{'type': 'L2MACschChgType-Choice', 'name': 'l2MACschChgType'}, None]
    L2MACschChgType L2MACschChgTypeChoice
}

func (self * TriggerTypeChoiceL2MACschChg) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.L2MACschChgType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TriggerTypeChoiceL2MACschChg) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.L2MACschChgType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *L2MACschChgTypeChoice)Unpack(stream *Stream) {
    //coptions := []string{"triggerType-Choice-MIMOandBFconfig"}
    choice := stream.get_choice(0, 1, 1)
    choice_len := 0
    choice_loc := 0
    if choice >= 1 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in L2MACschChgTypeChoice\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.TriggerTypeChoiceMIMOandBFconfig = &TriggerTypeChoiceMIMOandBFconfig{}//cho6
        self.TriggerTypeChoiceMIMOandBFconfig.Unpack(stream)
    }//end of if else

    if choice >= 1 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * L2MACschChgTypeChoice) Pack(stream *Stream) {
    if self.TriggerTypeChoiceMIMOandBFconfig != nil {
        stream.set_choice(0, 0, 1, 1)
        self.TriggerTypeChoiceMIMOandBFconfig.Pack(stream)//2
    }

}
type L2MACschChgTypeChoice struct { //[{'type': 'TriggerType-Choice-MIMOandBFconfig', 'name': 'triggerType-Choice-MIMOandBFconfig'}, None]
    TriggerTypeChoiceMIMOandBFconfig *TriggerTypeChoiceMIMOandBFconfig
} // L2MACschChgTypeChoice

type TriggerTypeChoiceMIMOandBFconfig struct { // [{'type': 'ENUMERATED', 'values': [('enabled', 0), ('disabled', 1), None], 'name': 'mIMOtransModeState'}, None]
    MIMOtransModeState ENUMERATED
}

func (self * TriggerTypeChoiceMIMOandBFconfig) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_mIMOtransModeState = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_mIMOtransModeState(stream, &self.MIMOtransModeState)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TriggerTypeChoiceMIMOandBFconfig) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_mIMOtransModeState = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_mIMOtransModeState(stream, self.MIMOtransModeState) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCActionDefinition_RicactionDefinitionformats struct { //[{'type': 'E2SM-RC-ActionDefinition-Format1', 'name': 'actionDefinition-Format1'}, {'type': 'E2SM-RC-ActionDefinition-Format2', 'name': 'actionDefinition-Format2'}, {'type': 'E2SM-RC-ActionDefinition-Format3', 'name': 'actionDefinition-Format3'}, None, {'type': 'E2SM-RC-ActionDefinition-Format4', 'name': 'actionDefinition-Format4'}]
    ActionDefinitionFormat1 *E2SMRCActionDefinitionFormat1
    ActionDefinitionFormat2 *E2SMRCActionDefinitionFormat2
    ActionDefinitionFormat3 *E2SMRCActionDefinitionFormat3
    ActionDefinitionFormat4 *E2SMRCActionDefinitionFormat4
} // E2SMRCActionDefinition_RicactionDefinitionformats

type E2SMRCActionDefinition struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-Style-Type'}, {'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-ActionDefinition-Format1', 'name': 'actionDefinition-Format1'}, {'type': 'E2SM-RC-ActionDefinition-Format2', 'name': 'actionDefinition-Format2'}, {'type': 'E2SM-RC-ActionDefinition-Format3', 'name': 'actionDefinition-Format3'}, None, {'type': 'E2SM-RC-ActionDefinition-Format4', 'name': 'actionDefinition-Format4'}], 'name': 'ric-actionDefinition-formats'}, None]
    RicStyleType RICStyleType
    RicactionDefinitionformats E2SMRCActionDefinition_RicactionDefinitionformats
}

func (self * E2SMRCActionDefinition) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicStyleType.Unpack(stream)// p8
    var Unpack_ricactionDefinitionformats = func(stream *Stream, self *E2SMRCActionDefinition_RicactionDefinitionformats) {
        //coptions := []string{"actionDefinition-Format1","actionDefinition-Format2","actionDefinition-Format3","Unknown","actionDefinition-Format4"}
        choice := stream.get_choice(2, 1, 3)
        choice_len := 0
        choice_loc := 0
        if choice >= 3 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCActionDefinition_RicactionDefinitionformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.ActionDefinitionFormat1 = &E2SMRCActionDefinitionFormat1{}//cho6
            self.ActionDefinitionFormat1.Unpack(stream)
        } else if choice == 1 { //ch2
            self.ActionDefinitionFormat2 = &E2SMRCActionDefinitionFormat2{}//cho6
            self.ActionDefinitionFormat2.Unpack(stream)
        } else if choice == 2 { //ch2
            self.ActionDefinitionFormat3 = &E2SMRCActionDefinitionFormat3{}//cho6
            self.ActionDefinitionFormat3.Unpack(stream)
        } else if choice == 4 { //ch2
            self.ActionDefinitionFormat4 = &E2SMRCActionDefinitionFormat4{}//cho6
            self.ActionDefinitionFormat4.Unpack(stream)
        }//end of if else

        if choice >= 3 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_ricactionDefinitionformats(stream, &self.RicactionDefinitionformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinition) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicStyleType.Pack(stream)
    var Pack_ricactionDefinitionformats = func(stream *Stream, self E2SMRCActionDefinition_RicactionDefinitionformats) {
        if self.ActionDefinitionFormat1 != nil {
            stream.set_choice(0, 2, 1, 3)
            self.ActionDefinitionFormat1.Pack(stream)//2
        } else if self.ActionDefinitionFormat2 != nil {
            stream.set_choice(1, 2, 1, 3)
            self.ActionDefinitionFormat2.Pack(stream)//2
        } else if self.ActionDefinitionFormat3 != nil {
            stream.set_choice(2, 2, 1, 3)
            self.ActionDefinitionFormat3.Pack(stream)//2
        } else if self.ActionDefinitionFormat4 != nil {
            stream.set_choice(4, 2, 1, 3)
            lenLoc := stream.reserve_len()
            self.ActionDefinitionFormat4.Pack(stream)//2
            stream.set_len(lenLoc)
        }

    }
    Pack_ricactionDefinitionformats(stream, self.RicactionDefinitionformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCActionDefinitionFormat1_RanPToBeReportedList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format1-Item'}, 'size': [(1, 'maxnoofParametersToReport')], 'name': 'ranP-ToBeReported-List'}
    Items []E2SMRCActionDefinitionFormat1Item
}
type E2SMRCActionDefinitionFormat1 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format1-Item'}, 'size': [(1, 'maxnoofParametersToReport')], 'name': 'ranP-ToBeReported-List'}, None]
    RanPToBeReportedList E2SMRCActionDefinitionFormat1_RanPToBeReportedList
}

func (self * E2SMRCActionDefinitionFormat1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ranPToBeReportedList = func(stream *Stream, self *E2SMRCActionDefinitionFormat1_RanPToBeReportedList){// Seq6 E2SMRCActionDefinitionFormat1 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format1-Item'}, 'size': [(1, 'maxnoofParametersToReport')], 'name': 'ranP-ToBeReported-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCActionDefinitionFormat1Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPToBeReportedList(stream, &self.RanPToBeReportedList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinitionFormat1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ranPToBeReportedList = func(stream *Stream, self E2SMRCActionDefinitionFormat1_RanPToBeReportedList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPToBeReportedList(stream, self.RanPToBeReportedList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCActionDefinitionFormat1Item struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParameterDefinition *RANParameterDefinition
}

func (self * E2SMRCActionDefinitionFormat1Item) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinitionFormat1Item) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCActionDefinitionFormat2_RicPolicyConditionsList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format2-Item'}, 'size': [(1, 'maxnoofPolicyConditions')], 'name': 'ric-PolicyConditions-List'}
    Items []E2SMRCActionDefinitionFormat2Item
}
type E2SMRCActionDefinitionFormat2 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format2-Item'}, 'size': [(1, 'maxnoofPolicyConditions')], 'name': 'ric-PolicyConditions-List'}, None]
    RicPolicyConditionsList E2SMRCActionDefinitionFormat2_RicPolicyConditionsList
}

func (self * E2SMRCActionDefinitionFormat2) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricPolicyConditionsList = func(stream *Stream, self *E2SMRCActionDefinitionFormat2_RicPolicyConditionsList){// Seq6 E2SMRCActionDefinitionFormat2 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format2-Item'}, 'size': [(1, 'maxnoofPolicyConditions')], 'name': 'ric-PolicyConditions-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCActionDefinitionFormat2Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricPolicyConditionsList(stream, &self.RicPolicyConditionsList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinitionFormat2) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricPolicyConditionsList = func(stream *Stream, self E2SMRCActionDefinitionFormat2_RicPolicyConditionsList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricPolicyConditionsList(stream, self.RicPolicyConditionsList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCActionDefinitionFormat2Item struct { // [{'type': 'RIC-PolicyAction', 'name': 'ric-PolicyAction'}, {'type': 'RANParameter-Testing', 'name': 'ric-PolicyConditionDefinition', 'optional': True}, None]
    RicPolicyAction RICPolicyAction
    RicPolicyConditionDefinition *RANParameterTesting
}

func (self * E2SMRCActionDefinitionFormat2Item) Unpack(stream *Stream) {
    ricPolicyConditionDefinition_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RicPolicyAction.Unpack(stream)// p8
    if (ricPolicyConditionDefinition_flag & _flags) == ricPolicyConditionDefinition_flag { //cond2
        self.RicPolicyConditionDefinition = &RANParameterTesting{}//7{'type': 'RANParameter-Testing', 'name': 'ric-PolicyConditionDefinition', 'optional': True}
        self.RicPolicyConditionDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinitionFormat2Item) Pack(stream *Stream) {
    const ricPolicyConditionDefinition_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicPolicyAction.Pack(stream)
    if self.RicPolicyConditionDefinition != nil { 
        _flags |= ricPolicyConditionDefinition_flag
        self.RicPolicyConditionDefinition.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCActionDefinitionFormat3_RanPInsertIndicationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format3-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-InsertIndication-List'}
    Items []E2SMRCActionDefinitionFormat3Item
}
type E2SMRCActionDefinitionFormat3 struct { // [{'type': 'RIC-InsertIndication-ID', 'name': 'ric-InsertIndication-ID'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format3-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-InsertIndication-List'}, {'type': 'UEID', 'name': 'ueID', 'optional': True}, None]
    RicInsertIndicationID RICInsertIndicationID
    RanPInsertIndicationList E2SMRCActionDefinitionFormat3_RanPInsertIndicationList
    UeID *UEID
}

func (self * E2SMRCActionDefinitionFormat3) Unpack(stream *Stream) {
    ueID_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RicInsertIndicationID.Unpack(stream)// p8
    var Unpack_ranPInsertIndicationList = func(stream *Stream, self *E2SMRCActionDefinitionFormat3_RanPInsertIndicationList){// Seq6 E2SMRCActionDefinitionFormat3 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format3-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-InsertIndication-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCActionDefinitionFormat3Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPInsertIndicationList(stream, &self.RanPInsertIndicationList)// p2
    if (ueID_flag & _flags) == ueID_flag { //cond2
        self.UeID = &UEID{}//7{'type': 'UEID', 'name': 'ueID', 'optional': True}
        self.UeID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinitionFormat3) Pack(stream *Stream) {
    const ueID_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicInsertIndicationID.Pack(stream)
    var Pack_ranPInsertIndicationList = func(stream *Stream, self E2SMRCActionDefinitionFormat3_RanPInsertIndicationList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPInsertIndicationList(stream, self.RanPInsertIndicationList) //f2
    if self.UeID != nil { 
        _flags |= ueID_flag
        self.UeID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCActionDefinitionFormat3Item struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParameterDefinition *RANParameterDefinition
}

func (self * E2SMRCActionDefinitionFormat3Item) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinitionFormat3Item) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCActionDefinitionFormat4_RicInsertStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format4-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-InsertStyle-List'}
    Items []E2SMRCActionDefinitionFormat4StyleItem
}
type E2SMRCActionDefinitionFormat4 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format4-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-InsertStyle-List'}, {'type': 'UEID', 'name': 'ueID', 'optional': True}, None]
    RicInsertStyleList E2SMRCActionDefinitionFormat4_RicInsertStyleList
    UeID *UEID
}

func (self * E2SMRCActionDefinitionFormat4) Unpack(stream *Stream) {
    ueID_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    var Unpack_ricInsertStyleList = func(stream *Stream, self *E2SMRCActionDefinitionFormat4_RicInsertStyleList){// Seq6 E2SMRCActionDefinitionFormat4 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format4-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-InsertStyle-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]E2SMRCActionDefinitionFormat4StyleItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricInsertStyleList(stream, &self.RicInsertStyleList)// p2
    if (ueID_flag & _flags) == ueID_flag { //cond2
        self.UeID = &UEID{}//7{'type': 'UEID', 'name': 'ueID', 'optional': True}
        self.UeID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinitionFormat4) Pack(stream *Stream) {
    const ueID_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricInsertStyleList = func(stream *Stream, self E2SMRCActionDefinitionFormat4_RicInsertStyleList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricInsertStyleList(stream, self.RicInsertStyleList) //f2
    if self.UeID != nil { 
        _flags |= ueID_flag
        self.UeID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCActionDefinitionFormat4StyleItem_RicInsertIndicationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format4-Indication-Item'}, 'size': [(1, 'maxnoofInsertIndicationActions')], 'name': 'ric-InsertIndication-List'}
    Items []E2SMRCActionDefinitionFormat4IndicationItem
}
type E2SMRCActionDefinitionFormat4StyleItem struct { // [{'type': 'RIC-Style-Type', 'name': 'requested-Insert-Style-Type'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format4-Indication-Item'}, 'size': [(1, 'maxnoofInsertIndicationActions')], 'name': 'ric-InsertIndication-List'}, None]
    RequestedInsertStyleType RICStyleType
    RicInsertIndicationList E2SMRCActionDefinitionFormat4StyleItem_RicInsertIndicationList
}

func (self * E2SMRCActionDefinitionFormat4StyleItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RequestedInsertStyleType.Unpack(stream)// p8
    var Unpack_ricInsertIndicationList = func(stream *Stream, self *E2SMRCActionDefinitionFormat4StyleItem_RicInsertIndicationList){// Seq6 E2SMRCActionDefinitionFormat4StyleItem {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format4-Indication-Item'}, 'size': [(1, 'maxnoofInsertIndicationActions')], 'name': 'ric-InsertIndication-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]E2SMRCActionDefinitionFormat4IndicationItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricInsertIndicationList(stream, &self.RicInsertIndicationList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinitionFormat4StyleItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RequestedInsertStyleType.Pack(stream)
    var Pack_ricInsertIndicationList = func(stream *Stream, self E2SMRCActionDefinitionFormat4StyleItem_RicInsertIndicationList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricInsertIndicationList(stream, self.RicInsertIndicationList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCActionDefinitionFormat4IndicationItem_RanPInsertIndicationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format4-RANP-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-InsertIndication-List'}
    Items []E2SMRCActionDefinitionFormat4RANPItem
}
type E2SMRCActionDefinitionFormat4IndicationItem struct { // [{'type': 'RIC-InsertIndication-ID', 'name': 'ric-InsertIndication-ID'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format4-RANP-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-InsertIndication-List'}, None]
    RicInsertIndicationID RICInsertIndicationID
    RanPInsertIndicationList E2SMRCActionDefinitionFormat4IndicationItem_RanPInsertIndicationList
}

func (self * E2SMRCActionDefinitionFormat4IndicationItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicInsertIndicationID.Unpack(stream)// p8
    var Unpack_ranPInsertIndicationList = func(stream *Stream, self *E2SMRCActionDefinitionFormat4IndicationItem_RanPInsertIndicationList){// Seq6 E2SMRCActionDefinitionFormat4IndicationItem {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ActionDefinition-Format4-RANP-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-InsertIndication-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCActionDefinitionFormat4RANPItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPInsertIndicationList(stream, &self.RanPInsertIndicationList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinitionFormat4IndicationItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicInsertIndicationID.Pack(stream)
    var Pack_ranPInsertIndicationList = func(stream *Stream, self E2SMRCActionDefinitionFormat4IndicationItem_RanPInsertIndicationList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPInsertIndicationList(stream, self.RanPInsertIndicationList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCActionDefinitionFormat4RANPItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParameterDefinition *RANParameterDefinition
}

func (self * E2SMRCActionDefinitionFormat4RANPItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCActionDefinitionFormat4RANPItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCIndicationHeader_RicindicationHeaderformats struct { //[{'type': 'E2SM-RC-IndicationHeader-Format1', 'name': 'indicationHeader-Format1'}, {'type': 'E2SM-RC-IndicationHeader-Format2', 'name': 'indicationHeader-Format2'}, None, {'type': 'E2SM-RC-IndicationHeader-Format3', 'name': 'indicationHeader-Format3'}]
    IndicationHeaderFormat1 *E2SMRCIndicationHeaderFormat1
    IndicationHeaderFormat2 *E2SMRCIndicationHeaderFormat2
    IndicationHeaderFormat3 *E2SMRCIndicationHeaderFormat3
} // E2SMRCIndicationHeader_RicindicationHeaderformats

type E2SMRCIndicationHeader struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-IndicationHeader-Format1', 'name': 'indicationHeader-Format1'}, {'type': 'E2SM-RC-IndicationHeader-Format2', 'name': 'indicationHeader-Format2'}, None, {'type': 'E2SM-RC-IndicationHeader-Format3', 'name': 'indicationHeader-Format3'}], 'name': 'ric-indicationHeader-formats'}, None]
    RicindicationHeaderformats E2SMRCIndicationHeader_RicindicationHeaderformats
}

func (self * E2SMRCIndicationHeader) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricindicationHeaderformats = func(stream *Stream, self *E2SMRCIndicationHeader_RicindicationHeaderformats) {
        //coptions := []string{"indicationHeader-Format1","indicationHeader-Format2","indicationHeader-Format3"}
        choice := stream.get_choice(1, 1, 2)
        choice_len := 0
        choice_loc := 0
        if choice >= 2 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCIndicationHeader_RicindicationHeaderformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.IndicationHeaderFormat1 = &E2SMRCIndicationHeaderFormat1{}//cho6
            self.IndicationHeaderFormat1.Unpack(stream)
        } else if choice == 1 { //ch2
            self.IndicationHeaderFormat2 = &E2SMRCIndicationHeaderFormat2{}//cho6
            self.IndicationHeaderFormat2.Unpack(stream)
        } else if choice == 2 { //ch2
            self.IndicationHeaderFormat3 = &E2SMRCIndicationHeaderFormat3{}//cho6
            self.IndicationHeaderFormat3.Unpack(stream)
        }//end of if else

        if choice >= 2 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_ricindicationHeaderformats(stream, &self.RicindicationHeaderformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationHeader) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricindicationHeaderformats = func(stream *Stream, self E2SMRCIndicationHeader_RicindicationHeaderformats) {
        if self.IndicationHeaderFormat1 != nil {
            stream.set_choice(0, 1, 1, 2)
            self.IndicationHeaderFormat1.Pack(stream)//2
        } else if self.IndicationHeaderFormat2 != nil {
            stream.set_choice(1, 1, 1, 2)
            self.IndicationHeaderFormat2.Pack(stream)//2
        } else if self.IndicationHeaderFormat3 != nil {
            stream.set_choice(2, 1, 1, 2)
            lenLoc := stream.reserve_len()
            self.IndicationHeaderFormat3.Pack(stream)//2
            stream.set_len(lenLoc)
        }

    }
    Pack_ricindicationHeaderformats(stream, self.RicindicationHeaderformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationHeaderFormat1 struct { // [{'type': 'RIC-EventTriggerCondition-ID', 'name': 'ric-eventTriggerCondition-ID', 'optional': True}, None]
    RiceventTriggerConditionID *RICEventTriggerConditionID
}

func (self * E2SMRCIndicationHeaderFormat1) Unpack(stream *Stream) {
    riceventTriggerConditionID_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    if (riceventTriggerConditionID_flag & _flags) == riceventTriggerConditionID_flag { //cond2
        self.RiceventTriggerConditionID = &RICEventTriggerConditionID{}//7{'type': 'RIC-EventTriggerCondition-ID', 'name': 'ric-eventTriggerCondition-ID', 'optional': True}
        self.RiceventTriggerConditionID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationHeaderFormat1) Pack(stream *Stream) {
    const riceventTriggerConditionID_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.RiceventTriggerConditionID != nil { 
        _flags |= riceventTriggerConditionID_flag
        self.RiceventTriggerConditionID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCIndicationHeaderFormat2 struct { // [{'type': 'UEID', 'name': 'ueID'}, {'type': 'RIC-Style-Type', 'name': 'ric-InsertStyle-Type'}, {'type': 'RIC-InsertIndication-ID', 'name': 'ric-InsertIndication-ID'}, None]
    UeID UEID
    RicInsertStyleType RICStyleType
    RicInsertIndicationID RICInsertIndicationID
}

func (self * E2SMRCIndicationHeaderFormat2) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.UeID.Unpack(stream)// p8
    self.RicInsertStyleType.Unpack(stream)// p8
    self.RicInsertIndicationID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationHeaderFormat2) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UeID.Pack(stream)
    self.RicInsertStyleType.Pack(stream)
    self.RicInsertIndicationID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationHeaderFormat3 struct { // [{'type': 'RIC-EventTriggerCondition-ID', 'name': 'ric-eventTriggerCondition-ID', 'optional': True}, {'type': 'UEID', 'name': 'ueID', 'optional': True}, None]
    RiceventTriggerConditionID *RICEventTriggerConditionID
    UeID *UEID
}

func (self * E2SMRCIndicationHeaderFormat3) Unpack(stream *Stream) {
    riceventTriggerConditionID_flag := 0x00000002
    ueID_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    if (riceventTriggerConditionID_flag & _flags) == riceventTriggerConditionID_flag { //cond2
        self.RiceventTriggerConditionID = &RICEventTriggerConditionID{}//7{'type': 'RIC-EventTriggerCondition-ID', 'name': 'ric-eventTriggerCondition-ID', 'optional': True}
        self.RiceventTriggerConditionID.Unpack(stream)// p8
    }
    if (ueID_flag & _flags) == ueID_flag { //cond2
        self.UeID = &UEID{}//7{'type': 'UEID', 'name': 'ueID', 'optional': True}
        self.UeID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationHeaderFormat3) Pack(stream *Stream) {
    const riceventTriggerConditionID_flag uint = 0x00000002
    const ueID_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.RiceventTriggerConditionID != nil { 
        _flags |= riceventTriggerConditionID_flag
        self.RiceventTriggerConditionID.Pack(stream)
    }//end of optional
    if self.UeID != nil { 
        _flags |= ueID_flag
        self.UeID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type E2SMRCIndicationMessage_RicindicationMessageformats struct { //[{'type': 'E2SM-RC-IndicationMessage-Format1', 'name': 'indicationMessage-Format1'}, {'type': 'E2SM-RC-IndicationMessage-Format2', 'name': 'indicationMessage-Format2'}, {'type': 'E2SM-RC-IndicationMessage-Format3', 'name': 'indicationMessage-Format3'}, {'type': 'NULL', 'name': 'indicationMessage-Format4'}, {'type': 'E2SM-RC-IndicationMessage-Format5', 'name': 'indicationMessage-Format5'}, None, {'type': 'E2SM-RC-IndicationMessage-Format6', 'name': 'indicationMessage-Format6'}]
    IndicationMessageFormat1 *E2SMRCIndicationMessageFormat1
    IndicationMessageFormat2 *E2SMRCIndicationMessageFormat2
    IndicationMessageFormat3 *E2SMRCIndicationMessageFormat3
    IndicationMessageFormat4 *NULL
    IndicationMessageFormat5 *E2SMRCIndicationMessageFormat5
    IndicationMessageFormat6 *E2SMRCIndicationMessageFormat6
} // E2SMRCIndicationMessage_RicindicationMessageformats

type E2SMRCIndicationMessage struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-IndicationMessage-Format1', 'name': 'indicationMessage-Format1'}, {'type': 'E2SM-RC-IndicationMessage-Format2', 'name': 'indicationMessage-Format2'}, {'type': 'E2SM-RC-IndicationMessage-Format3', 'name': 'indicationMessage-Format3'}, {'type': 'NULL', 'name': 'indicationMessage-Format4'}, {'type': 'E2SM-RC-IndicationMessage-Format5', 'name': 'indicationMessage-Format5'}, None, {'type': 'E2SM-RC-IndicationMessage-Format6', 'name': 'indicationMessage-Format6'}], 'name': 'ric-indicationMessage-formats'}, None]
    RicindicationMessageformats E2SMRCIndicationMessage_RicindicationMessageformats
}

func (self * E2SMRCIndicationMessage) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricindicationMessageformats = func(stream *Stream, self *E2SMRCIndicationMessage_RicindicationMessageformats) {
        //coptions := []string{"indicationMessage-Format1","indicationMessage-Format2","indicationMessage-Format3","indicationMessage-Format4","indicationMessage-Format5","Unknown","Unknown","Unknown","indicationMessage-Format6"}
        choice := stream.get_choice(3, 1, 5)
        choice_len := 0
        choice_loc := 0
        if choice >= 5 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCIndicationMessage_RicindicationMessageformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.IndicationMessageFormat1 = &E2SMRCIndicationMessageFormat1{}//cho6
            self.IndicationMessageFormat1.Unpack(stream)
        } else if choice == 1 { //ch2
            self.IndicationMessageFormat2 = &E2SMRCIndicationMessageFormat2{}//cho6
            self.IndicationMessageFormat2.Unpack(stream)
        } else if choice == 2 { //ch2
            self.IndicationMessageFormat3 = &E2SMRCIndicationMessageFormat3{}//cho6
            self.IndicationMessageFormat3.Unpack(stream)
        } else if choice == 3 { //ch2
            var Unpack_indicationMessageFormat4 = func(st *Stream, self *NULL){
                st.parsef_Null()
            }
            self.IndicationMessageFormat4 = &NULL{}//cho5
            Unpack_indicationMessageFormat4(stream, self.IndicationMessageFormat4);
        } else if choice == 4 { //ch2
            self.IndicationMessageFormat5 = &E2SMRCIndicationMessageFormat5{}//cho6
            self.IndicationMessageFormat5.Unpack(stream)
        } else if choice == 8 { //ch2
            self.IndicationMessageFormat6 = &E2SMRCIndicationMessageFormat6{}//cho6
            self.IndicationMessageFormat6.Unpack(stream)
        }//end of if else

        if choice >= 5 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_ricindicationMessageformats(stream, &self.RicindicationMessageformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessage) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricindicationMessageformats = func(stream *Stream, self E2SMRCIndicationMessage_RicindicationMessageformats) {
        if self.IndicationMessageFormat1 != nil {
            stream.set_choice(0, 3, 1, 5)
            self.IndicationMessageFormat1.Pack(stream)//2
        } else if self.IndicationMessageFormat2 != nil {
            stream.set_choice(1, 3, 1, 5)
            self.IndicationMessageFormat2.Pack(stream)//2
        } else if self.IndicationMessageFormat3 != nil {
            stream.set_choice(2, 3, 1, 5)
            self.IndicationMessageFormat3.Pack(stream)//2
        } else if self.IndicationMessageFormat4 != nil {
            stream.set_choice(3, 3, 1, 5)
            var Pack_indicationMessageFormat4 = func(st *Stream, self NULL){
                st.formatf_Null()
            }
            Pack_indicationMessageFormat4(stream, *self.IndicationMessageFormat4)//3
        } else if self.IndicationMessageFormat5 != nil {
            stream.set_choice(4, 3, 1, 5)
            self.IndicationMessageFormat5.Pack(stream)//2
        } else if self.IndicationMessageFormat6 != nil {
            stream.set_choice(8, 3, 1, 5)
            lenLoc := stream.reserve_len()
            self.IndicationMessageFormat6.Pack(stream)//2
            stream.set_len(lenLoc)
        }

    }
    Pack_ricindicationMessageformats(stream, self.RicindicationMessageformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat1_RanPReportedList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format1-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-Reported-List'}
    Items []E2SMRCIndicationMessageFormat1Item
}
type E2SMRCIndicationMessageFormat1 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format1-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-Reported-List'}, None]
    RanPReportedList E2SMRCIndicationMessageFormat1_RanPReportedList
}

func (self * E2SMRCIndicationMessageFormat1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ranPReportedList = func(stream *Stream, self *E2SMRCIndicationMessageFormat1_RanPReportedList){// Seq6 E2SMRCIndicationMessageFormat1 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format1-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-Reported-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCIndicationMessageFormat1Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPReportedList(stream, &self.RanPReportedList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ranPReportedList = func(stream *Stream, self E2SMRCIndicationMessageFormat1_RanPReportedList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPReportedList(stream, self.RanPReportedList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat1Item struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
}

func (self * E2SMRCIndicationMessageFormat1Item) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat1Item) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat2_UeParameterList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format2-Item'}, 'size': [(1, 'maxnoofUEID')], 'name': 'ueParameter-List'}
    Items []E2SMRCIndicationMessageFormat2Item
}
type E2SMRCIndicationMessageFormat2 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format2-Item'}, 'size': [(1, 'maxnoofUEID')], 'name': 'ueParameter-List'}, None]
    UeParameterList E2SMRCIndicationMessageFormat2_UeParameterList
}

func (self * E2SMRCIndicationMessageFormat2) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ueParameterList = func(stream *Stream, self *E2SMRCIndicationMessageFormat2_UeParameterList){// Seq6 E2SMRCIndicationMessageFormat2 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format2-Item'}, 'size': [(1, 'maxnoofUEID')], 'name': 'ueParameter-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCIndicationMessageFormat2Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ueParameterList(stream, &self.UeParameterList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat2) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ueParameterList = func(stream *Stream, self E2SMRCIndicationMessageFormat2_UeParameterList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ueParameterList(stream, self.UeParameterList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat2Item_RanPList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format2-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
    Items []E2SMRCIndicationMessageFormat2RANParameterItem
}
type E2SMRCIndicationMessageFormat2Item struct { // [{'type': 'UEID', 'name': 'ueID'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format2-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}, None]
    UeID UEID
    RanPList E2SMRCIndicationMessageFormat2Item_RanPList
}

func (self * E2SMRCIndicationMessageFormat2Item) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.UeID.Unpack(stream)// p8
    var Unpack_ranPList = func(stream *Stream, self *E2SMRCIndicationMessageFormat2Item_RanPList){// Seq6 E2SMRCIndicationMessageFormat2Item {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format2-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCIndicationMessageFormat2RANParameterItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPList(stream, &self.RanPList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat2Item) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UeID.Pack(stream)
    var Pack_ranPList = func(stream *Stream, self E2SMRCIndicationMessageFormat2Item_RanPList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPList(stream, self.RanPList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat2RANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
}

func (self * E2SMRCIndicationMessageFormat2RANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat2RANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat3_CellInfoList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format3-Item'}, 'size': [(1, 'maxnoofCellID')], 'name': 'cellInfo-List'}
    Items []E2SMRCIndicationMessageFormat3Item
}
type E2SMRCIndicationMessageFormat3 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format3-Item'}, 'size': [(1, 'maxnoofCellID')], 'name': 'cellInfo-List'}, None]
    CellInfoList E2SMRCIndicationMessageFormat3_CellInfoList
}

func (self * E2SMRCIndicationMessageFormat3) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_cellInfoList = func(stream *Stream, self *E2SMRCIndicationMessageFormat3_CellInfoList){// Seq6 E2SMRCIndicationMessageFormat3 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format3-Item'}, 'size': [(1, 'maxnoofCellID')], 'name': 'cellInfo-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCIndicationMessageFormat3Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_cellInfoList(stream, &self.CellInfoList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat3) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_cellInfoList = func(stream *Stream, self E2SMRCIndicationMessageFormat3_CellInfoList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_cellInfoList(stream, self.CellInfoList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat3Item struct { // [{'type': 'CGI', 'name': 'cellGlobal-ID'}, {'type': 'OCTET STRING', 'name': 'cellContextInfo', 'optional': True}, {'type': 'BOOLEAN', 'name': 'cellDeleted', 'optional': True}, {'type': 'NeighborRelation-Info', 'name': 'neighborRelation-Table', 'optional': True}, None]
    CellGlobalID CGI
    CellContextInfo *OCTETSTRING
    CellDeleted *BOOLEAN
    NeighborRelationTable *NeighborRelationInfo
}

func (self * E2SMRCIndicationMessageFormat3Item) Unpack(stream *Stream) {
    cellContextInfo_flag := 0x00000002
    cellDeleted_flag := 0x00000004
    neighborRelationTable_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.CellGlobalID.Unpack(stream)// p8
    if (cellContextInfo_flag & _flags) == cellContextInfo_flag { //cond1
        var Unpack_cellContextInfo = func(st *Stream, self *OCTETSTRING) {
            _len := st.parse_len(0)
            self.Value = st.parsef_OctString(_len)
        }
        self.CellContextInfo = &OCTETSTRING{}//6{'type': 'OCTET STRING', 'name': 'cellContextInfo', 'optional': True}
        Unpack_cellContextInfo(stream, self.CellContextInfo)// p1 {'type': 'OCTET STRING', 'name': 'cellContextInfo', 'optional': True}
    }
    if (cellDeleted_flag & _flags) == cellDeleted_flag { //cond1
        var  Unpack_cellDeleted = func (st *Stream, self *BOOLEAN) {
            self.Value = st.parsef_bool()
            }
        self.CellDeleted = &BOOLEAN{}//6{'type': 'BOOLEAN', 'name': 'cellDeleted', 'optional': True}
        Unpack_cellDeleted(stream, self.CellDeleted)// p1 {'type': 'BOOLEAN', 'name': 'cellDeleted', 'optional': True}
    }
    if (neighborRelationTable_flag & _flags) == neighborRelationTable_flag { //cond2
        self.NeighborRelationTable = &NeighborRelationInfo{}//7{'type': 'NeighborRelation-Info', 'name': 'neighborRelation-Table', 'optional': True}
        self.NeighborRelationTable.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat3Item) Pack(stream *Stream) {
    const cellContextInfo_flag uint = 0x00000002
    const cellDeleted_flag uint = 0x00000004
    const neighborRelationTable_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CellGlobalID.Pack(stream)
    if self.CellContextInfo != nil { //YY
        _flags |= cellContextInfo_flag
        var Pack_cellContextInfo = func(st *Stream, self OCTETSTRING) {
            st.format_len(len(self.Value), 0)
            st.formatf_OctString(self.Value, 0)
        }
        Pack_cellContextInfo(stream, *self.CellContextInfo) //f1
    }//end of optional
    if self.CellDeleted != nil { //YY
        _flags |= cellDeleted_flag
        var Pack_cellDeleted = func(st *Stream, self BOOLEAN) {
            st.formatf_bool(self.Value);
            }
        Pack_cellDeleted(stream, *self.CellDeleted) //f1
    }//end of optional
    if self.NeighborRelationTable != nil { 
        _flags |= neighborRelationTable_flag
        self.NeighborRelationTable.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type E2SMRCIndicationMessageFormat5_RanPRequestedList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format5-Item'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-Requested-List'}
    Items []E2SMRCIndicationMessageFormat5Item
}
type E2SMRCIndicationMessageFormat5 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format5-Item'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-Requested-List'}, None]
    RanPRequestedList E2SMRCIndicationMessageFormat5_RanPRequestedList
}

func (self * E2SMRCIndicationMessageFormat5) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ranPRequestedList = func(stream *Stream, self *E2SMRCIndicationMessageFormat5_RanPRequestedList){// Seq6 E2SMRCIndicationMessageFormat5 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format5-Item'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-Requested-List'}
        _size := stream.get_listsize(65536)
        _size += 0
        self.Items = make([]E2SMRCIndicationMessageFormat5Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPRequestedList(stream, &self.RanPRequestedList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat5) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ranPRequestedList = func(stream *Stream, self E2SMRCIndicationMessageFormat5_RanPRequestedList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-0, 65536)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPRequestedList(stream, self.RanPRequestedList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat5Item struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
}

func (self * E2SMRCIndicationMessageFormat5Item) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat5Item) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat6_RicInsertStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format6-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-InsertStyle-List'}
    Items []E2SMRCIndicationMessageFormat6StyleItem
}
type E2SMRCIndicationMessageFormat6 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format6-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-InsertStyle-List'}, None]
    RicInsertStyleList E2SMRCIndicationMessageFormat6_RicInsertStyleList
}

func (self * E2SMRCIndicationMessageFormat6) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricInsertStyleList = func(stream *Stream, self *E2SMRCIndicationMessageFormat6_RicInsertStyleList){// Seq6 E2SMRCIndicationMessageFormat6 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format6-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-InsertStyle-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]E2SMRCIndicationMessageFormat6StyleItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricInsertStyleList(stream, &self.RicInsertStyleList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat6) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricInsertStyleList = func(stream *Stream, self E2SMRCIndicationMessageFormat6_RicInsertStyleList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricInsertStyleList(stream, self.RicInsertStyleList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat6StyleItem_RicInsertIndicationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format6-Indication-Item'}, 'size': [(1, 'maxnoofInsertIndicationActions')], 'name': 'ric-InsertIndication-List'}
    Items []E2SMRCIndicationMessageFormat6IndicationItem
}
type E2SMRCIndicationMessageFormat6StyleItem struct { // [{'type': 'RIC-Style-Type', 'name': 'indicated-Insert-Style-Type'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format6-Indication-Item'}, 'size': [(1, 'maxnoofInsertIndicationActions')], 'name': 'ric-InsertIndication-List'}, None]
    IndicatedInsertStyleType RICStyleType
    RicInsertIndicationList E2SMRCIndicationMessageFormat6StyleItem_RicInsertIndicationList
}

func (self * E2SMRCIndicationMessageFormat6StyleItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.IndicatedInsertStyleType.Unpack(stream)// p8
    var Unpack_ricInsertIndicationList = func(stream *Stream, self *E2SMRCIndicationMessageFormat6StyleItem_RicInsertIndicationList){// Seq6 E2SMRCIndicationMessageFormat6StyleItem {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format6-Indication-Item'}, 'size': [(1, 'maxnoofInsertIndicationActions')], 'name': 'ric-InsertIndication-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]E2SMRCIndicationMessageFormat6IndicationItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricInsertIndicationList(stream, &self.RicInsertIndicationList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat6StyleItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IndicatedInsertStyleType.Pack(stream)
    var Pack_ricInsertIndicationList = func(stream *Stream, self E2SMRCIndicationMessageFormat6StyleItem_RicInsertIndicationList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricInsertIndicationList(stream, self.RicInsertIndicationList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat6IndicationItem_RanPInsertIndicationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format6-RANP-Item'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-InsertIndication-List'}
    Items []E2SMRCIndicationMessageFormat6RANPItem
}
type E2SMRCIndicationMessageFormat6IndicationItem struct { // [{'type': 'RIC-InsertIndication-ID', 'name': 'ric-InsertIndication-ID'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format6-RANP-Item'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-InsertIndication-List'}, None]
    RicInsertIndicationID RICInsertIndicationID
    RanPInsertIndicationList E2SMRCIndicationMessageFormat6IndicationItem_RanPInsertIndicationList
}

func (self * E2SMRCIndicationMessageFormat6IndicationItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicInsertIndicationID.Unpack(stream)// p8
    var Unpack_ranPInsertIndicationList = func(stream *Stream, self *E2SMRCIndicationMessageFormat6IndicationItem_RanPInsertIndicationList){// Seq6 E2SMRCIndicationMessageFormat6IndicationItem {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-IndicationMessage-Format6-RANP-Item'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-InsertIndication-List'}
        _size := stream.get_listsize(65536)
        _size += 0
        self.Items = make([]E2SMRCIndicationMessageFormat6RANPItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPInsertIndicationList(stream, &self.RanPInsertIndicationList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat6IndicationItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicInsertIndicationID.Pack(stream)
    var Pack_ranPInsertIndicationList = func(stream *Stream, self E2SMRCIndicationMessageFormat6IndicationItem_RanPInsertIndicationList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-0, 65536)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPInsertIndicationList(stream, self.RanPInsertIndicationList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCIndicationMessageFormat6RANPItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
}

func (self * E2SMRCIndicationMessageFormat6RANPItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCIndicationMessageFormat6RANPItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCCallProcessID_RiccallProcessIDformats struct { //[{'type': 'E2SM-RC-CallProcessID-Format1', 'name': 'callProcessID-Format1'}, None]
    CallProcessIDFormat1 *E2SMRCCallProcessIDFormat1
} // E2SMRCCallProcessID_RiccallProcessIDformats

type E2SMRCCallProcessID struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-CallProcessID-Format1', 'name': 'callProcessID-Format1'}, None], 'name': 'ric-callProcessID-formats'}, None]
    RiccallProcessIDformats E2SMRCCallProcessID_RiccallProcessIDformats
}

func (self * E2SMRCCallProcessID) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_riccallProcessIDformats = func(stream *Stream, self *E2SMRCCallProcessID_RiccallProcessIDformats) {
        //coptions := []string{"callProcessID-Format1"}
        choice := stream.get_choice(0, 1, 1)
        choice_len := 0
        choice_loc := 0
        if choice >= 1 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCCallProcessID_RiccallProcessIDformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.CallProcessIDFormat1 = &E2SMRCCallProcessIDFormat1{}//cho6
            self.CallProcessIDFormat1.Unpack(stream)
        }//end of if else

        if choice >= 1 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_riccallProcessIDformats(stream, &self.RiccallProcessIDformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCCallProcessID) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_riccallProcessIDformats = func(stream *Stream, self E2SMRCCallProcessID_RiccallProcessIDformats) {
        if self.CallProcessIDFormat1 != nil {
            stream.set_choice(0, 0, 1, 1)
            self.CallProcessIDFormat1.Pack(stream)//2
        }

    }
    Pack_riccallProcessIDformats(stream, self.RiccallProcessIDformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCCallProcessIDFormat1 struct { // [{'type': 'RAN-CallProcess-ID', 'name': 'ric-callProcess-ID'}, None]
    RiccallProcessID RANCallProcessID
}

func (self * E2SMRCCallProcessIDFormat1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RiccallProcessID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCCallProcessIDFormat1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RiccallProcessID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlHeader_RiccontrolHeaderformats struct { //[{'type': 'E2SM-RC-ControlHeader-Format1', 'name': 'controlHeader-Format1'}, None, {'type': 'E2SM-RC-ControlHeader-Format2', 'name': 'controlHeader-Format2'}, {'type': 'E2SM-RC-ControlHeader-Format3', 'name': 'controlHeader-Format3'}]
    ControlHeaderFormat1 *E2SMRCControlHeaderFormat1
    ControlHeaderFormat2 *E2SMRCControlHeaderFormat2
    ControlHeaderFormat3 *E2SMRCControlHeaderFormat3
} // E2SMRCControlHeader_RiccontrolHeaderformats

type E2SMRCControlHeader struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-ControlHeader-Format1', 'name': 'controlHeader-Format1'}, None, {'type': 'E2SM-RC-ControlHeader-Format2', 'name': 'controlHeader-Format2'}, {'type': 'E2SM-RC-ControlHeader-Format3', 'name': 'controlHeader-Format3'}], 'name': 'ric-controlHeader-formats'}, None]
    RiccontrolHeaderformats E2SMRCControlHeader_RiccontrolHeaderformats
}

func (self * E2SMRCControlHeader) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_riccontrolHeaderformats = func(stream *Stream, self *E2SMRCControlHeader_RiccontrolHeaderformats) {
        //coptions := []string{"controlHeader-Format1","controlHeader-Format2","controlHeader-Format3"}
        choice := stream.get_choice(0, 1, 1)
        choice_len := 0
        choice_loc := 0
        if choice >= 1 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCControlHeader_RiccontrolHeaderformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.ControlHeaderFormat1 = &E2SMRCControlHeaderFormat1{}//cho6
            self.ControlHeaderFormat1.Unpack(stream)
        } else if choice == 1 { //ch2
            self.ControlHeaderFormat2 = &E2SMRCControlHeaderFormat2{}//cho6
            self.ControlHeaderFormat2.Unpack(stream)
        } else if choice == 2 { //ch2
            self.ControlHeaderFormat3 = &E2SMRCControlHeaderFormat3{}//cho6
            self.ControlHeaderFormat3.Unpack(stream)
        }//end of if else

        if choice >= 1 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_riccontrolHeaderformats(stream, &self.RiccontrolHeaderformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlHeader) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_riccontrolHeaderformats = func(stream *Stream, self E2SMRCControlHeader_RiccontrolHeaderformats) {
        if self.ControlHeaderFormat1 != nil {
            stream.set_choice(0, 0, 1, 1)
            self.ControlHeaderFormat1.Pack(stream)//2
        } else if self.ControlHeaderFormat2 != nil {
            stream.set_choice(1, 0, 1, 1)
            lenLoc := stream.reserve_len()
            self.ControlHeaderFormat2.Pack(stream)//2
            stream.set_len(lenLoc)
        } else if self.ControlHeaderFormat3 != nil {
            stream.set_choice(2, 0, 1, 1)
            lenLoc := stream.reserve_len()
            self.ControlHeaderFormat3.Pack(stream)//2
            stream.set_len(lenLoc)
        }

    }
    Pack_riccontrolHeaderformats(stream, self.RiccontrolHeaderformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlHeaderFormat1 struct { // [{'type': 'UEID', 'name': 'ueID'}, {'type': 'RIC-Style-Type', 'name': 'ric-Style-Type'}, {'type': 'RIC-ControlAction-ID', 'name': 'ric-ControlAction-ID'}, {'type': 'ENUMERATED', 'values': [('accept', 0), ('reject', 1), None], 'name': 'ric-ControlDecision', 'optional': True}, None]
    UeID UEID
    RicStyleType RICStyleType
    RicControlActionID RICControlActionID
    RicControlDecision *ENUMERATED
}

func (self * E2SMRCControlHeaderFormat1) Unpack(stream *Stream) {
    ricControlDecision_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.UeID.Unpack(stream)// p8
    self.RicStyleType.Unpack(stream)// p8
    self.RicControlActionID.Unpack(stream)// p8
    if (ricControlDecision_flag & _flags) == ricControlDecision_flag { //cond1
        var Unpack_ricControlDecision = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(2, 2, 1)
        }
        self.RicControlDecision = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('accept', 0), ('reject', 1), None], 'name': 'ric-ControlDecision', 'optional': True}
        Unpack_ricControlDecision(stream, self.RicControlDecision)// p1 {'type': 'ENUMERATED', 'values': [('accept', 0), ('reject', 1), None], 'name': 'ric-ControlDecision', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlHeaderFormat1) Pack(stream *Stream) {
    const ricControlDecision_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UeID.Pack(stream)
    self.RicStyleType.Pack(stream)
    self.RicControlActionID.Pack(stream)
    if self.RicControlDecision != nil { //YY
        _flags |= ricControlDecision_flag
        var Pack_ricControlDecision = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 2, 2, 1)
        }
        Pack_ricControlDecision(stream, *self.RicControlDecision) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCControlHeaderFormat2 struct { // [{'type': 'UEID', 'name': 'ueID', 'optional': True}, {'type': 'ENUMERATED', 'values': [('accept', 0), ('reject', 1), None], 'name': 'ric-ControlDecision', 'optional': True}, None]
    UeID *UEID
    RicControlDecision *ENUMERATED
}

func (self * E2SMRCControlHeaderFormat2) Unpack(stream *Stream) {
    ueID_flag := 0x00000002
    ricControlDecision_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    if (ueID_flag & _flags) == ueID_flag { //cond2
        self.UeID = &UEID{}//7{'type': 'UEID', 'name': 'ueID', 'optional': True}
        self.UeID.Unpack(stream)// p8
    }
    if (ricControlDecision_flag & _flags) == ricControlDecision_flag { //cond1
        var Unpack_ricControlDecision = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(2, 2, 1)
        }
        self.RicControlDecision = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('accept', 0), ('reject', 1), None], 'name': 'ric-ControlDecision', 'optional': True}
        Unpack_ricControlDecision(stream, self.RicControlDecision)// p1 {'type': 'ENUMERATED', 'values': [('accept', 0), ('reject', 1), None], 'name': 'ric-ControlDecision', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlHeaderFormat2) Pack(stream *Stream) {
    const ueID_flag uint = 0x00000002
    const ricControlDecision_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.UeID != nil { 
        _flags |= ueID_flag
        self.UeID.Pack(stream)
    }//end of optional
    if self.RicControlDecision != nil { //YY
        _flags |= ricControlDecision_flag
        var Pack_ricControlDecision = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 2, 2, 1)
        }
        Pack_ricControlDecision(stream, *self.RicControlDecision) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type E2SMRCControlHeaderFormat3 struct { // [{'type': 'UE-Group-ID', 'name': 'ue-Group-ID'}, {'type': 'UE-Group-Definition', 'name': 'ue-Group-Definition'}, {'type': 'RIC-Style-Type', 'name': 'ric-Style-Type'}, {'type': 'RIC-ControlAction-ID', 'name': 'ric-ControlAction-ID'}, None]
    UeGroupID UEGroupID
    UeGroupDefinition UEGroupDefinition
    RicStyleType RICStyleType
    RicControlActionID RICControlActionID
}

func (self * E2SMRCControlHeaderFormat3) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.UeGroupID.Unpack(stream)// p8
    self.UeGroupDefinition.Unpack(stream)// p8
    self.RicStyleType.Unpack(stream)// p8
    self.RicControlActionID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlHeaderFormat3) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UeGroupID.Pack(stream)
    self.UeGroupDefinition.Pack(stream)
    self.RicStyleType.Pack(stream)
    self.RicControlActionID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlMessage_RiccontrolMessageformats struct { //[{'type': 'E2SM-RC-ControlMessage-Format1', 'name': 'controlMessage-Format1'}, None, {'type': 'E2SM-RC-ControlMessage-Format2', 'name': 'controlMessage-Format2'}, {'type': 'E2SM-RC-ControlMessage-Format3', 'name': 'controlMessage-Format3'}]
    ControlMessageFormat1 *E2SMRCControlMessageFormat1
    ControlMessageFormat2 *E2SMRCControlMessageFormat2
    ControlMessageFormat3 *E2SMRCControlMessageFormat3
} // E2SMRCControlMessage_RiccontrolMessageformats

type E2SMRCControlMessage struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-ControlMessage-Format1', 'name': 'controlMessage-Format1'}, None, {'type': 'E2SM-RC-ControlMessage-Format2', 'name': 'controlMessage-Format2'}, {'type': 'E2SM-RC-ControlMessage-Format3', 'name': 'controlMessage-Format3'}], 'name': 'ric-controlMessage-formats'}, None]
    RiccontrolMessageformats E2SMRCControlMessage_RiccontrolMessageformats
}

func (self * E2SMRCControlMessage) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_riccontrolMessageformats = func(stream *Stream, self *E2SMRCControlMessage_RiccontrolMessageformats) {
        //coptions := []string{"controlMessage-Format1","controlMessage-Format2","controlMessage-Format3"}
        choice := stream.get_choice(0, 1, 1)
        choice_len := 0
        choice_loc := 0
        if choice >= 1 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCControlMessage_RiccontrolMessageformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.ControlMessageFormat1 = &E2SMRCControlMessageFormat1{}//cho6
            self.ControlMessageFormat1.Unpack(stream)
        } else if choice == 1 { //ch2
            self.ControlMessageFormat2 = &E2SMRCControlMessageFormat2{}//cho6
            self.ControlMessageFormat2.Unpack(stream)
        } else if choice == 2 { //ch2
            self.ControlMessageFormat3 = &E2SMRCControlMessageFormat3{}//cho6
            self.ControlMessageFormat3.Unpack(stream)
        }//end of if else

        if choice >= 1 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_riccontrolMessageformats(stream, &self.RiccontrolMessageformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlMessage) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_riccontrolMessageformats = func(stream *Stream, self E2SMRCControlMessage_RiccontrolMessageformats) {
        if self.ControlMessageFormat1 != nil {
            stream.set_choice(0, 0, 1, 1)
            self.ControlMessageFormat1.Pack(stream)//2
        } else if self.ControlMessageFormat2 != nil {
            stream.set_choice(1, 0, 1, 1)
            lenLoc := stream.reserve_len()
            self.ControlMessageFormat2.Pack(stream)//2
            stream.set_len(lenLoc)
        } else if self.ControlMessageFormat3 != nil {
            stream.set_choice(2, 0, 1, 1)
            lenLoc := stream.reserve_len()
            self.ControlMessageFormat3.Pack(stream)//2
            stream.set_len(lenLoc)
        }

    }
    Pack_riccontrolMessageformats(stream, self.RiccontrolMessageformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlMessageFormat1_RanPList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlMessage-Format1-Item'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
    Items []E2SMRCControlMessageFormat1Item
}
type E2SMRCControlMessageFormat1 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlMessage-Format1-Item'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}, None]
    RanPList E2SMRCControlMessageFormat1_RanPList
}

func (self * E2SMRCControlMessageFormat1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ranPList = func(stream *Stream, self *E2SMRCControlMessageFormat1_RanPList){// Seq6 E2SMRCControlMessageFormat1 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlMessage-Format1-Item'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
        _size := stream.get_listsize(65536)
        _size += 0
        self.Items = make([]E2SMRCControlMessageFormat1Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPList(stream, &self.RanPList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlMessageFormat1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ranPList = func(stream *Stream, self E2SMRCControlMessageFormat1_RanPList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-0, 65536)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPList(stream, self.RanPList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlMessageFormat1Item struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
}

func (self * E2SMRCControlMessageFormat1Item) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlMessageFormat1Item) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlMessageFormat2_RicControlStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlMessage-Format2-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ControlStyle-List'}
    Items []E2SMRCControlMessageFormat2StyleItem
}
type E2SMRCControlMessageFormat2 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlMessage-Format2-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ControlStyle-List'}, None]
    RicControlStyleList E2SMRCControlMessageFormat2_RicControlStyleList
}

func (self * E2SMRCControlMessageFormat2) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricControlStyleList = func(stream *Stream, self *E2SMRCControlMessageFormat2_RicControlStyleList){// Seq6 E2SMRCControlMessageFormat2 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlMessage-Format2-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ControlStyle-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]E2SMRCControlMessageFormat2StyleItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricControlStyleList(stream, &self.RicControlStyleList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlMessageFormat2) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricControlStyleList = func(stream *Stream, self E2SMRCControlMessageFormat2_RicControlStyleList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricControlStyleList(stream, self.RicControlStyleList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlMessageFormat2StyleItem_RicControlActionList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlMessage-Format2-ControlAction-Item'}, 'size': [(1, 'maxnoofMulCtrlActions')], 'name': 'ric-ControlAction-List'}
    Items []E2SMRCControlMessageFormat2ControlActionItem
}
type E2SMRCControlMessageFormat2StyleItem struct { // [{'type': 'RIC-Style-Type', 'name': 'indicated-Control-Style-Type'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlMessage-Format2-ControlAction-Item'}, 'size': [(1, 'maxnoofMulCtrlActions')], 'name': 'ric-ControlAction-List'}, None]
    IndicatedControlStyleType RICStyleType
    RicControlActionList E2SMRCControlMessageFormat2StyleItem_RicControlActionList
}

func (self * E2SMRCControlMessageFormat2StyleItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.IndicatedControlStyleType.Unpack(stream)// p8
    var Unpack_ricControlActionList = func(stream *Stream, self *E2SMRCControlMessageFormat2StyleItem_RicControlActionList){// Seq6 E2SMRCControlMessageFormat2StyleItem {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlMessage-Format2-ControlAction-Item'}, 'size': [(1, 'maxnoofMulCtrlActions')], 'name': 'ric-ControlAction-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]E2SMRCControlMessageFormat2ControlActionItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricControlActionList(stream, &self.RicControlActionList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlMessageFormat2StyleItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IndicatedControlStyleType.Pack(stream)
    var Pack_ricControlActionList = func(stream *Stream, self E2SMRCControlMessageFormat2StyleItem_RicControlActionList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricControlActionList(stream, self.RicControlActionList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlMessageFormat2ControlActionItem struct { // [{'type': 'RIC-ControlAction-ID', 'name': 'ric-ControlAction-ID'}, {'type': 'E2SM-RC-ControlMessage-Format1', 'name': 'ranP-List'}, None]
    RicControlActionID RICControlActionID
    RanPList E2SMRCControlMessageFormat1
}

func (self * E2SMRCControlMessageFormat2ControlActionItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicControlActionID.Unpack(stream)// p8
    self.RanPList.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlMessageFormat2ControlActionItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicControlActionID.Pack(stream)
    self.RanPList.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlMessageFormat3_EntityAgnosticControlRanPList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'EntityAgnostic-ranP-ControlParameters'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'entityAgnosticControlRanP-List', 'optional': True}
    Items []EntityAgnosticranPControlParameters
}
type E2SMRCControlMessageFormat3_ListOfEntityFilters struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EntityFilter'}, 'size': [(0, 'maxnoofAssociatedEntityFilters')], 'name': 'listOfEntityFilters', 'optional': True}
    Items []E2SMRCEntityFilter
}
type E2SMRCControlMessageFormat3 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EntityFilter'}, 'size': [(0, 'maxnoofAssociatedEntityFilters')], 'name': 'listOfEntityFilters', 'optional': True}, {'type': 'SEQUENCE OF', 'element': {'type': 'EntityAgnostic-ranP-ControlParameters'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'entityAgnosticControlRanP-List', 'optional': True}, None]
    ListOfEntityFilters *E2SMRCControlMessageFormat3_ListOfEntityFilters
    EntityAgnosticControlRanPList *E2SMRCControlMessageFormat3_EntityAgnosticControlRanPList
}

func (self * E2SMRCControlMessageFormat3) Unpack(stream *Stream) {
    listOfEntityFilters_flag := 0x00000002
    entityAgnosticControlRanPList_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    if (listOfEntityFilters_flag & _flags) == listOfEntityFilters_flag { //cond1
        var Unpack_listOfEntityFilters = func(stream *Stream, self *E2SMRCControlMessageFormat3_ListOfEntityFilters){// Seq6 E2SMRCControlMessageFormat3 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EntityFilter'}, 'size': [(0, 'maxnoofAssociatedEntityFilters')], 'name': 'listOfEntityFilters', 'optional': True}
            _size := stream.get_listsize(256)
            _size += 0
            self.Items = make([]E2SMRCEntityFilter, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.ListOfEntityFilters = &E2SMRCControlMessageFormat3_ListOfEntityFilters{}//3
        Unpack_listOfEntityFilters(stream, self.ListOfEntityFilters)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-EntityFilter'}, 'size': [(0, 'maxnoofAssociatedEntityFilters')], 'name': 'listOfEntityFilters', 'optional': True}
    }
    if (entityAgnosticControlRanPList_flag & _flags) == entityAgnosticControlRanPList_flag { //cond1
        var Unpack_entityAgnosticControlRanPList = func(stream *Stream, self *E2SMRCControlMessageFormat3_EntityAgnosticControlRanPList){// Seq6 E2SMRCControlMessageFormat3 {'type': 'SEQUENCE OF', 'element': {'type': 'EntityAgnostic-ranP-ControlParameters'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'entityAgnosticControlRanP-List', 'optional': True}
            _size := stream.get_listsize(65536)
            _size += 0
            self.Items = make([]EntityAgnosticranPControlParameters, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.EntityAgnosticControlRanPList = &E2SMRCControlMessageFormat3_EntityAgnosticControlRanPList{}//3
        Unpack_entityAgnosticControlRanPList(stream, self.EntityAgnosticControlRanPList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'EntityAgnostic-ranP-ControlParameters'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'entityAgnosticControlRanP-List', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlMessageFormat3) Pack(stream *Stream) {
    const listOfEntityFilters_flag uint = 0x00000002
    const entityAgnosticControlRanPList_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.ListOfEntityFilters != nil { //YY
        _flags |= listOfEntityFilters_flag
        var Pack_listOfEntityFilters = func(stream *Stream, self E2SMRCControlMessageFormat3_ListOfEntityFilters) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-0, 256)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_listOfEntityFilters(stream, *self.ListOfEntityFilters) //f1
    }//end of optional
    if self.EntityAgnosticControlRanPList != nil { //YY
        _flags |= entityAgnosticControlRanPList_flag
        var Pack_entityAgnosticControlRanPList = func(stream *Stream, self E2SMRCControlMessageFormat3_EntityAgnosticControlRanPList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-0, 65536)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_entityAgnosticControlRanPList(stream, *self.EntityAgnosticControlRanPList) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type E2SMRCEntityFilter_EntitySpecificControlRanPList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'EntitySpecific-ranP-ControlParameters'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'entitySpecificControlRanP-List'}
    Items []EntitySpecificranPControlParameters
}
type E2SMRCEntityFilter struct { // [{'type': 'EntityFilter-ID', 'name': 'entityFilter-ID'}, {'type': 'RANParameter-Testing', 'name': 'entityFilter-Definition'}, {'type': 'SEQUENCE OF', 'element': {'type': 'EntitySpecific-ranP-ControlParameters'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'entitySpecificControlRanP-List'}, None]
    EntityFilterID EntityFilterID
    EntityFilterDefinition RANParameterTesting
    EntitySpecificControlRanPList E2SMRCEntityFilter_EntitySpecificControlRanPList
}

func (self * E2SMRCEntityFilter) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.EntityFilterID.Unpack(stream)// p8
    self.EntityFilterDefinition.Unpack(stream)// p8
    var Unpack_entitySpecificControlRanPList = func(stream *Stream, self *E2SMRCEntityFilter_EntitySpecificControlRanPList){// Seq6 E2SMRCEntityFilter {'type': 'SEQUENCE OF', 'element': {'type': 'EntitySpecific-ranP-ControlParameters'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'entitySpecificControlRanP-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]EntitySpecificranPControlParameters, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_entitySpecificControlRanPList(stream, &self.EntitySpecificControlRanPList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCEntityFilter) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.EntityFilterID.Pack(stream)
    self.EntityFilterDefinition.Pack(stream)
    var Pack_entitySpecificControlRanPList = func(stream *Stream, self E2SMRCEntityFilter_EntitySpecificControlRanPList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_entitySpecificControlRanPList(stream, self.EntitySpecificControlRanPList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EntityAgnosticranPControlParameters struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
}

func (self * EntityAgnosticranPControlParameters) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EntityAgnosticranPControlParameters) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type EntitySpecificranPControlParameters struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
}

func (self * EntitySpecificranPControlParameters) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * EntitySpecificranPControlParameters) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlOutcome_RiccontrolOutcomeformats struct { //[{'type': 'E2SM-RC-ControlOutcome-Format1', 'name': 'controlOutcome-Format1'}, None, {'type': 'E2SM-RC-ControlOutcome-Format2', 'name': 'controlOutcome-Format2'}, {'type': 'E2SM-RC-ControlOutcome-Format3', 'name': 'controlOutcome-Format3'}]
    ControlOutcomeFormat1 *E2SMRCControlOutcomeFormat1
    ControlOutcomeFormat2 *E2SMRCControlOutcomeFormat2
    ControlOutcomeFormat3 *E2SMRCControlOutcomeFormat3
} // E2SMRCControlOutcome_RiccontrolOutcomeformats

type E2SMRCControlOutcome struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-ControlOutcome-Format1', 'name': 'controlOutcome-Format1'}, None, {'type': 'E2SM-RC-ControlOutcome-Format2', 'name': 'controlOutcome-Format2'}, {'type': 'E2SM-RC-ControlOutcome-Format3', 'name': 'controlOutcome-Format3'}], 'name': 'ric-controlOutcome-formats'}, None]
    RiccontrolOutcomeformats E2SMRCControlOutcome_RiccontrolOutcomeformats
}

func (self * E2SMRCControlOutcome) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_riccontrolOutcomeformats = func(stream *Stream, self *E2SMRCControlOutcome_RiccontrolOutcomeformats) {
        //coptions := []string{"controlOutcome-Format1","controlOutcome-Format2","controlOutcome-Format3"}
        choice := stream.get_choice(0, 1, 1)
        choice_len := 0
        choice_loc := 0
        if choice >= 1 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCControlOutcome_RiccontrolOutcomeformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.ControlOutcomeFormat1 = &E2SMRCControlOutcomeFormat1{}//cho6
            self.ControlOutcomeFormat1.Unpack(stream)
        } else if choice == 1 { //ch2
            self.ControlOutcomeFormat2 = &E2SMRCControlOutcomeFormat2{}//cho6
            self.ControlOutcomeFormat2.Unpack(stream)
        } else if choice == 2 { //ch2
            self.ControlOutcomeFormat3 = &E2SMRCControlOutcomeFormat3{}//cho6
            self.ControlOutcomeFormat3.Unpack(stream)
        }//end of if else

        if choice >= 1 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_riccontrolOutcomeformats(stream, &self.RiccontrolOutcomeformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlOutcome) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_riccontrolOutcomeformats = func(stream *Stream, self E2SMRCControlOutcome_RiccontrolOutcomeformats) {
        if self.ControlOutcomeFormat1 != nil {
            stream.set_choice(0, 0, 1, 1)
            self.ControlOutcomeFormat1.Pack(stream)//2
        } else if self.ControlOutcomeFormat2 != nil {
            stream.set_choice(1, 0, 1, 1)
            lenLoc := stream.reserve_len()
            self.ControlOutcomeFormat2.Pack(stream)//2
            stream.set_len(lenLoc)
        } else if self.ControlOutcomeFormat3 != nil {
            stream.set_choice(2, 0, 1, 1)
            lenLoc := stream.reserve_len()
            self.ControlOutcomeFormat3.Pack(stream)//2
            stream.set_len(lenLoc)
        }

    }
    Pack_riccontrolOutcomeformats(stream, self.RiccontrolOutcomeformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlOutcomeFormat1_RanPList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format1-Item'}, 'size': [(0, 'maxnoofRANOutcomeParameters')], 'name': 'ranP-List'}
    Items []E2SMRCControlOutcomeFormat1Item
}
type E2SMRCControlOutcomeFormat1 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format1-Item'}, 'size': [(0, 'maxnoofRANOutcomeParameters')], 'name': 'ranP-List'}, None]
    RanPList E2SMRCControlOutcomeFormat1_RanPList
}

func (self * E2SMRCControlOutcomeFormat1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ranPList = func(stream *Stream, self *E2SMRCControlOutcomeFormat1_RanPList){// Seq6 E2SMRCControlOutcomeFormat1 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format1-Item'}, 'size': [(0, 'maxnoofRANOutcomeParameters')], 'name': 'ranP-List'}
        _size := stream.get_listsize(256)
        _size += 0
        self.Items = make([]E2SMRCControlOutcomeFormat1Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPList(stream, &self.RanPList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlOutcomeFormat1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ranPList = func(stream *Stream, self E2SMRCControlOutcomeFormat1_RanPList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-0, 256)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPList(stream, self.RanPList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlOutcomeFormat1Item struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Value', 'name': 'ranParameter-value'}, None]
    RanParameterID RANParameterID
    RanParametervalue RANParameterValue
}

func (self * E2SMRCControlOutcomeFormat1Item) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalue.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlOutcomeFormat1Item) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalue.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlOutcomeFormat2_RicControlStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format2-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ControlStyle-List'}
    Items []E2SMRCControlOutcomeFormat2StyleItem
}
type E2SMRCControlOutcomeFormat2 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format2-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ControlStyle-List'}, None]
    RicControlStyleList E2SMRCControlOutcomeFormat2_RicControlStyleList
}

func (self * E2SMRCControlOutcomeFormat2) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricControlStyleList = func(stream *Stream, self *E2SMRCControlOutcomeFormat2_RicControlStyleList){// Seq6 E2SMRCControlOutcomeFormat2 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format2-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ControlStyle-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]E2SMRCControlOutcomeFormat2StyleItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricControlStyleList(stream, &self.RicControlStyleList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlOutcomeFormat2) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricControlStyleList = func(stream *Stream, self E2SMRCControlOutcomeFormat2_RicControlStyleList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricControlStyleList(stream, self.RicControlStyleList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlOutcomeFormat2StyleItem_RicControlOutcomeList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format2-ControlOutcome-Item'}, 'size': [(1, 'maxnoofMulCtrlActions')], 'name': 'ric-ControlOutcome-List'}
    Items []E2SMRCControlOutcomeFormat2ControlOutcomeItem
}
type E2SMRCControlOutcomeFormat2StyleItem struct { // [{'type': 'RIC-Style-Type', 'name': 'indicated-Control-Style-Type'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format2-ControlOutcome-Item'}, 'size': [(1, 'maxnoofMulCtrlActions')], 'name': 'ric-ControlOutcome-List'}, None]
    IndicatedControlStyleType RICStyleType
    RicControlOutcomeList E2SMRCControlOutcomeFormat2StyleItem_RicControlOutcomeList
}

func (self * E2SMRCControlOutcomeFormat2StyleItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.IndicatedControlStyleType.Unpack(stream)// p8
    var Unpack_ricControlOutcomeList = func(stream *Stream, self *E2SMRCControlOutcomeFormat2StyleItem_RicControlOutcomeList){// Seq6 E2SMRCControlOutcomeFormat2StyleItem {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format2-ControlOutcome-Item'}, 'size': [(1, 'maxnoofMulCtrlActions')], 'name': 'ric-ControlOutcome-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]E2SMRCControlOutcomeFormat2ControlOutcomeItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricControlOutcomeList(stream, &self.RicControlOutcomeList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlOutcomeFormat2StyleItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.IndicatedControlStyleType.Pack(stream)
    var Pack_ricControlOutcomeList = func(stream *Stream, self E2SMRCControlOutcomeFormat2StyleItem_RicControlOutcomeList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricControlOutcomeList(stream, self.RicControlOutcomeList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlOutcomeFormat2ControlOutcomeItem_RanPList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format2-RANP-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
    Items []E2SMRCControlOutcomeFormat2RANPItem
}
type E2SMRCControlOutcomeFormat2ControlOutcomeItem struct { // [{'type': 'RIC-ControlAction-ID', 'name': 'ric-ControlAction-ID'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format2-RANP-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}, None]
    RicControlActionID RICControlActionID
    RanPList E2SMRCControlOutcomeFormat2ControlOutcomeItem_RanPList
}

func (self * E2SMRCControlOutcomeFormat2ControlOutcomeItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicControlActionID.Unpack(stream)// p8
    var Unpack_ranPList = func(stream *Stream, self *E2SMRCControlOutcomeFormat2ControlOutcomeItem_RanPList){// Seq6 E2SMRCControlOutcomeFormat2ControlOutcomeItem {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format2-RANP-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCControlOutcomeFormat2RANPItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPList(stream, &self.RanPList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlOutcomeFormat2ControlOutcomeItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicControlActionID.Pack(stream)
    var Pack_ranPList = func(stream *Stream, self E2SMRCControlOutcomeFormat2ControlOutcomeItem_RanPList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPList(stream, self.RanPList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlOutcomeFormat2RANPItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Value', 'name': 'ranParameter-value'}, None]
    RanParameterID RANParameterID
    RanParametervalue RANParameterValue
}

func (self * E2SMRCControlOutcomeFormat2RANPItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalue.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlOutcomeFormat2RANPItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalue.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlOutcomeFormat3_RanPList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format3-Item'}, 'size': [(0, 'maxnoofRANOutcomeParameters')], 'name': 'ranP-List'}
    Items []E2SMRCControlOutcomeFormat3Item
}
type E2SMRCControlOutcomeFormat3 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format3-Item'}, 'size': [(0, 'maxnoofRANOutcomeParameters')], 'name': 'ranP-List'}, None]
    RanPList E2SMRCControlOutcomeFormat3_RanPList
}

func (self * E2SMRCControlOutcomeFormat3) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ranPList = func(stream *Stream, self *E2SMRCControlOutcomeFormat3_RanPList){// Seq6 E2SMRCControlOutcomeFormat3 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-ControlOutcome-Format3-Item'}, 'size': [(0, 'maxnoofRANOutcomeParameters')], 'name': 'ranP-List'}
        _size := stream.get_listsize(256)
        _size += 0
        self.Items = make([]E2SMRCControlOutcomeFormat3Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPList(stream, &self.RanPList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlOutcomeFormat3) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ranPList = func(stream *Stream, self E2SMRCControlOutcomeFormat3_RanPList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-0, 256)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPList(stream, self.RanPList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCControlOutcomeFormat3Item struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType'}, None]
    RanParameterID RANParameterID
    RanParametervalueType RANParameterValueType
}

func (self * E2SMRCControlOutcomeFormat3Item) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametervalueType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCControlOutcomeFormat3Item) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametervalueType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCQueryHeader_RicqueryHeaderformats struct { //[{'type': 'E2SM-RC-QueryHeader-Format1', 'name': 'queryHeader-Format1'}, None]
    QueryHeaderFormat1 *E2SMRCQueryHeaderFormat1
} // E2SMRCQueryHeader_RicqueryHeaderformats

type E2SMRCQueryHeader struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-QueryHeader-Format1', 'name': 'queryHeader-Format1'}, None], 'name': 'ric-queryHeader-formats'}, None]
    RicqueryHeaderformats E2SMRCQueryHeader_RicqueryHeaderformats
}

func (self * E2SMRCQueryHeader) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricqueryHeaderformats = func(stream *Stream, self *E2SMRCQueryHeader_RicqueryHeaderformats) {
        //coptions := []string{"queryHeader-Format1"}
        choice := stream.get_choice(0, 1, 1)
        choice_len := 0
        choice_loc := 0
        if choice >= 1 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCQueryHeader_RicqueryHeaderformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.QueryHeaderFormat1 = &E2SMRCQueryHeaderFormat1{}//cho6
            self.QueryHeaderFormat1.Unpack(stream)
        }//end of if else

        if choice >= 1 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_ricqueryHeaderformats(stream, &self.RicqueryHeaderformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryHeader) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricqueryHeaderformats = func(stream *Stream, self E2SMRCQueryHeader_RicqueryHeaderformats) {
        if self.QueryHeaderFormat1 != nil {
            stream.set_choice(0, 0, 1, 1)
            self.QueryHeaderFormat1.Pack(stream)//2
        }

    }
    Pack_ricqueryHeaderformats(stream, self.RicqueryHeaderformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCQueryHeaderFormat1 struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-Style-Type'}, {'type': 'RANParameter-Testing', 'name': 'associatedE2NodeInfo', 'optional': True}, {'type': 'Associated-UE-Info', 'name': 'associatedUEInfo', 'optional': True}, None]
    RicStyleType RICStyleType
    AssociatedE2NodeInfo *RANParameterTesting
    AssociatedUEInfo *AssociatedUEInfo
}

func (self * E2SMRCQueryHeaderFormat1) Unpack(stream *Stream) {
    associatedE2NodeInfo_flag := 0x00000002
    associatedUEInfo_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RicStyleType.Unpack(stream)// p8
    if (associatedE2NodeInfo_flag & _flags) == associatedE2NodeInfo_flag { //cond2
        self.AssociatedE2NodeInfo = &RANParameterTesting{}//7{'type': 'RANParameter-Testing', 'name': 'associatedE2NodeInfo', 'optional': True}
        self.AssociatedE2NodeInfo.Unpack(stream)// p8
    }
    if (associatedUEInfo_flag & _flags) == associatedUEInfo_flag { //cond2
        self.AssociatedUEInfo = &AssociatedUEInfo{}//7{'type': 'Associated-UE-Info', 'name': 'associatedUEInfo', 'optional': True}
        self.AssociatedUEInfo.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryHeaderFormat1) Pack(stream *Stream) {
    const associatedE2NodeInfo_flag uint = 0x00000002
    const associatedUEInfo_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicStyleType.Pack(stream)
    if self.AssociatedE2NodeInfo != nil { 
        _flags |= associatedE2NodeInfo_flag
        self.AssociatedE2NodeInfo.Pack(stream)
    }//end of optional
    if self.AssociatedUEInfo != nil { 
        _flags |= associatedUEInfo_flag
        self.AssociatedUEInfo.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type E2SMRCQueryDefinition_RicqueryDefinitionformats struct { //[{'type': 'E2SM-RC-QueryDefinition-Format1', 'name': 'queryRequest-Format1'}, None]
    QueryRequestFormat1 *E2SMRCQueryDefinitionFormat1
} // E2SMRCQueryDefinition_RicqueryDefinitionformats

type E2SMRCQueryDefinition struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-QueryDefinition-Format1', 'name': 'queryRequest-Format1'}, None], 'name': 'ric-queryDefinition-formats'}, None]
    RicqueryDefinitionformats E2SMRCQueryDefinition_RicqueryDefinitionformats
}

func (self * E2SMRCQueryDefinition) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricqueryDefinitionformats = func(stream *Stream, self *E2SMRCQueryDefinition_RicqueryDefinitionformats) {
        //coptions := []string{"queryRequest-Format1"}
        choice := stream.get_choice(0, 1, 1)
        choice_len := 0
        choice_loc := 0
        if choice >= 1 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCQueryDefinition_RicqueryDefinitionformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.QueryRequestFormat1 = &E2SMRCQueryDefinitionFormat1{}//cho6
            self.QueryRequestFormat1.Unpack(stream)
        }//end of if else

        if choice >= 1 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_ricqueryDefinitionformats(stream, &self.RicqueryDefinitionformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryDefinition) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricqueryDefinitionformats = func(stream *Stream, self E2SMRCQueryDefinition_RicqueryDefinitionformats) {
        if self.QueryRequestFormat1 != nil {
            stream.set_choice(0, 0, 1, 1)
            self.QueryRequestFormat1.Pack(stream)//2
        }

    }
    Pack_ricqueryDefinitionformats(stream, self.RicqueryDefinitionformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCQueryDefinitionFormat1_RanPList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryDefinition-Format1-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
    Items []E2SMRCQueryDefinitionFormat1Item
}
type E2SMRCQueryDefinitionFormat1 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryDefinition-Format1-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}, None]
    RanPList E2SMRCQueryDefinitionFormat1_RanPList
}

func (self * E2SMRCQueryDefinitionFormat1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ranPList = func(stream *Stream, self *E2SMRCQueryDefinitionFormat1_RanPList){// Seq6 E2SMRCQueryDefinitionFormat1 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryDefinition-Format1-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCQueryDefinitionFormat1Item, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPList(stream, &self.RanPList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryDefinitionFormat1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ranPList = func(stream *Stream, self E2SMRCQueryDefinitionFormat1_RanPList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPList(stream, self.RanPList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCQueryDefinitionFormat1Item struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}, None]
    RanParameterID RANParameterID
    RanParameterDefinition *RANParameterDefinition
}

func (self * E2SMRCQueryDefinitionFormat1Item) Unpack(stream *Stream) {
    ranParameterDefinition_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RanParameterID.Unpack(stream)// p8
    if (ranParameterDefinition_flag & _flags) == ranParameterDefinition_flag { //cond2
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryDefinitionFormat1Item) Pack(stream *Stream) {
    const ranParameterDefinition_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    if self.RanParameterDefinition != nil { 
        _flags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCQueryOutcome_RicqueryOutcomeformats struct { //[{'type': 'E2SM-RC-QueryOutcome-Format1', 'name': 'queryOutcome-Format1'}, {'type': 'E2SM-RC-QueryOutcome-Format2', 'name': 'queryOutcome-Format2'}, None]
    QueryOutcomeFormat1 *E2SMRCQueryOutcomeFormat1
    QueryOutcomeFormat2 *E2SMRCQueryOutcomeFormat2
} // E2SMRCQueryOutcome_RicqueryOutcomeformats

type E2SMRCQueryOutcome struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-QueryOutcome-Format1', 'name': 'queryOutcome-Format1'}, {'type': 'E2SM-RC-QueryOutcome-Format2', 'name': 'queryOutcome-Format2'}, None], 'name': 'ric-queryOutcome-formats'}, None]
    RicqueryOutcomeformats E2SMRCQueryOutcome_RicqueryOutcomeformats
}

func (self * E2SMRCQueryOutcome) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricqueryOutcomeformats = func(stream *Stream, self *E2SMRCQueryOutcome_RicqueryOutcomeformats) {
        //coptions := []string{"queryOutcome-Format1","queryOutcome-Format2"}
        choice := stream.get_choice(1, 1, 2)
        choice_len := 0
        choice_loc := 0
        if choice >= 2 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMRCQueryOutcome_RicqueryOutcomeformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.QueryOutcomeFormat1 = &E2SMRCQueryOutcomeFormat1{}//cho6
            self.QueryOutcomeFormat1.Unpack(stream)
        } else if choice == 1 { //ch2
            self.QueryOutcomeFormat2 = &E2SMRCQueryOutcomeFormat2{}//cho6
            self.QueryOutcomeFormat2.Unpack(stream)
        }//end of if else

        if choice >= 2 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_ricqueryOutcomeformats(stream, &self.RicqueryOutcomeformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryOutcome) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricqueryOutcomeformats = func(stream *Stream, self E2SMRCQueryOutcome_RicqueryOutcomeformats) {
        if self.QueryOutcomeFormat1 != nil {
            stream.set_choice(0, 1, 1, 2)
            self.QueryOutcomeFormat1.Pack(stream)//2
        } else if self.QueryOutcomeFormat2 != nil {
            stream.set_choice(1, 1, 1, 2)
            self.QueryOutcomeFormat2.Pack(stream)//2
        }

    }
    Pack_ricqueryOutcomeformats(stream, self.RicqueryOutcomeformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCQueryOutcomeFormat1_CellInfoList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format1-ItemCell'}, 'size': [(1, 'maxnoofCellID')], 'name': 'cellInfo-List'}
    Items []E2SMRCQueryOutcomeFormat1ItemCell
}
type E2SMRCQueryOutcomeFormat1 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format1-ItemCell'}, 'size': [(1, 'maxnoofCellID')], 'name': 'cellInfo-List'}, None]
    CellInfoList E2SMRCQueryOutcomeFormat1_CellInfoList
}

func (self * E2SMRCQueryOutcomeFormat1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_cellInfoList = func(stream *Stream, self *E2SMRCQueryOutcomeFormat1_CellInfoList){// Seq6 E2SMRCQueryOutcomeFormat1 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format1-ItemCell'}, 'size': [(1, 'maxnoofCellID')], 'name': 'cellInfo-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]E2SMRCQueryOutcomeFormat1ItemCell, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_cellInfoList(stream, &self.CellInfoList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryOutcomeFormat1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_cellInfoList = func(stream *Stream, self E2SMRCQueryOutcomeFormat1_CellInfoList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_cellInfoList(stream, self.CellInfoList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCQueryOutcomeFormat1ItemCell_RanPList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format1-ItemParameters'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
    Items []E2SMRCQueryOutcomeFormat1ItemParameters
}
type E2SMRCQueryOutcomeFormat1ItemCell struct { // [{'type': 'CGI', 'name': 'cellGlobal-ID'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format1-ItemParameters'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}, {'type': 'NeighborRelation-Info', 'name': 'neighborRelation-Table', 'optional': True}, None]
    CellGlobalID CGI
    RanPList E2SMRCQueryOutcomeFormat1ItemCell_RanPList
    NeighborRelationTable *NeighborRelationInfo
}

func (self * E2SMRCQueryOutcomeFormat1ItemCell) Unpack(stream *Stream) {
    neighborRelationTable_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.CellGlobalID.Unpack(stream)// p8
    var Unpack_ranPList = func(stream *Stream, self *E2SMRCQueryOutcomeFormat1ItemCell_RanPList){// Seq6 E2SMRCQueryOutcomeFormat1ItemCell {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format1-ItemParameters'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
        _size := stream.get_listsize(65536)
        _size += 0
        self.Items = make([]E2SMRCQueryOutcomeFormat1ItemParameters, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPList(stream, &self.RanPList)// p2
    if (neighborRelationTable_flag & _flags) == neighborRelationTable_flag { //cond2
        self.NeighborRelationTable = &NeighborRelationInfo{}//7{'type': 'NeighborRelation-Info', 'name': 'neighborRelation-Table', 'optional': True}
        self.NeighborRelationTable.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryOutcomeFormat1ItemCell) Pack(stream *Stream) {
    const neighborRelationTable_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CellGlobalID.Pack(stream)
    var Pack_ranPList = func(stream *Stream, self E2SMRCQueryOutcomeFormat1ItemCell_RanPList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-0, 65536)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPList(stream, self.RanPList) //f2
    if self.NeighborRelationTable != nil { 
        _flags |= neighborRelationTable_flag
        self.NeighborRelationTable.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCQueryOutcomeFormat1ItemParameters struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType', 'optional': True}, None]
    RanParameterID RANParameterID
    RanParametervalueType *RANParameterValueType
}

func (self * E2SMRCQueryOutcomeFormat1ItemParameters) Unpack(stream *Stream) {
    ranParametervalueType_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RanParameterID.Unpack(stream)// p8
    if (ranParametervalueType_flag & _flags) == ranParametervalueType_flag { //cond2
        self.RanParametervalueType = &RANParameterValueType{}//7{'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType', 'optional': True}
        self.RanParametervalueType.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryOutcomeFormat1ItemParameters) Pack(stream *Stream) {
    const ranParametervalueType_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    if self.RanParametervalueType != nil { 
        _flags |= ranParametervalueType_flag
        self.RanParametervalueType.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCQueryOutcomeFormat2_UeInfoList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format2-ItemUE'}, 'size': [(0, 'maxnoofUEID')], 'name': 'ueInfo-List'}
    Items []E2SMRCQueryOutcomeFormat2ItemUE
}
type E2SMRCQueryOutcomeFormat2 struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format2-ItemUE'}, 'size': [(0, 'maxnoofUEID')], 'name': 'ueInfo-List'}, None]
    UeInfoList E2SMRCQueryOutcomeFormat2_UeInfoList
}

func (self * E2SMRCQueryOutcomeFormat2) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ueInfoList = func(stream *Stream, self *E2SMRCQueryOutcomeFormat2_UeInfoList){// Seq6 E2SMRCQueryOutcomeFormat2 {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format2-ItemUE'}, 'size': [(0, 'maxnoofUEID')], 'name': 'ueInfo-List'}
        _size := stream.get_listsize(65536)
        _size += 0
        self.Items = make([]E2SMRCQueryOutcomeFormat2ItemUE, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ueInfoList(stream, &self.UeInfoList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryOutcomeFormat2) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ueInfoList = func(stream *Stream, self E2SMRCQueryOutcomeFormat2_UeInfoList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-0, 65536)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ueInfoList(stream, self.UeInfoList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMRCQueryOutcomeFormat2ItemUE_RanPList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format2-ItemParameters'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
    Items []E2SMRCQueryOutcomeFormat2ItemParameters
}
type E2SMRCQueryOutcomeFormat2ItemUE struct { // [{'type': 'UEID', 'name': 'ueID'}, {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format2-ItemParameters'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}, {'type': 'UE-Filter-ID', 'name': 'ueFilterID', 'optional': True}, None]
    UeID UEID
    RanPList E2SMRCQueryOutcomeFormat2ItemUE_RanPList
    UeFilterID *UEFilterID
}

func (self * E2SMRCQueryOutcomeFormat2ItemUE) Unpack(stream *Stream) {
    ueFilterID_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.UeID.Unpack(stream)// p8
    var Unpack_ranPList = func(stream *Stream, self *E2SMRCQueryOutcomeFormat2ItemUE_RanPList){// Seq6 E2SMRCQueryOutcomeFormat2ItemUE {'type': 'SEQUENCE OF', 'element': {'type': 'E2SM-RC-QueryOutcome-Format2-ItemParameters'}, 'size': [(0, 'maxnoofAssociatedRANParameters')], 'name': 'ranP-List'}
        _size := stream.get_listsize(65536)
        _size += 0
        self.Items = make([]E2SMRCQueryOutcomeFormat2ItemParameters, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ranPList(stream, &self.RanPList)// p2
    if (ueFilterID_flag & _flags) == ueFilterID_flag { //cond2
        self.UeFilterID = &UEFilterID{}//7{'type': 'UE-Filter-ID', 'name': 'ueFilterID', 'optional': True}
        self.UeFilterID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryOutcomeFormat2ItemUE) Pack(stream *Stream) {
    const ueFilterID_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UeID.Pack(stream)
    var Pack_ranPList = func(stream *Stream, self E2SMRCQueryOutcomeFormat2ItemUE_RanPList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-0, 65536)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ranPList(stream, self.RanPList) //f2
    if self.UeFilterID != nil { 
        _flags |= ueFilterID_flag
        self.UeFilterID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCQueryOutcomeFormat2ItemParameters struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType', 'optional': True}, None]
    RanParameterID RANParameterID
    RanParametervalueType *RANParameterValueType
}

func (self * E2SMRCQueryOutcomeFormat2ItemParameters) Unpack(stream *Stream) {
    ranParametervalueType_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RanParameterID.Unpack(stream)// p8
    if (ranParametervalueType_flag & _flags) == ranParametervalueType_flag { //cond2
        self.RanParametervalueType = &RANParameterValueType{}//7{'type': 'RANParameter-ValueType', 'name': 'ranParameter-valueType', 'optional': True}
        self.RanParametervalueType.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCQueryOutcomeFormat2ItemParameters) Pack(stream *Stream) {
    const ranParametervalueType_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    if self.RanParametervalueType != nil { 
        _flags |= ranParametervalueType_flag
        self.RanParametervalueType.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMRCRANFunctionDefinition struct { // [{'type': 'RANfunction-Name', 'name': 'ranFunction-Name'}, {'type': 'RANFunctionDefinition-EventTrigger', 'name': 'ranFunctionDefinition-EventTrigger', 'optional': True}, {'type': 'RANFunctionDefinition-Report', 'name': 'ranFunctionDefinition-Report', 'optional': True}, {'type': 'RANFunctionDefinition-Insert', 'name': 'ranFunctionDefinition-Insert', 'optional': True}, {'type': 'RANFunctionDefinition-Control', 'name': 'ranFunctionDefinition-Control', 'optional': True}, {'type': 'RANFunctionDefinition-Policy', 'name': 'ranFunctionDefinition-Policy', 'optional': True}, None, {'type': 'RANFunctionDefinition-Query', 'name': 'ranFunctionDefinition-Query', 'optional': True}]
    RanFunctionName RANfunctionName
    RanFunctionDefinitionEventTrigger *RANFunctionDefinitionEventTrigger
    RanFunctionDefinitionReport *RANFunctionDefinitionReport
    RanFunctionDefinitionInsert *RANFunctionDefinitionInsert
    RanFunctionDefinitionControl *RANFunctionDefinitionControl
    RanFunctionDefinitionPolicy *RANFunctionDefinitionPolicy
    RanFunctionDefinitionQuery *RANFunctionDefinitionQuery
}

func (self * E2SMRCRANFunctionDefinition) Unpack(stream *Stream) {
    ranFunctionDefinitionEventTrigger_flag := 0x00000002
    ranFunctionDefinitionReport_flag := 0x00000004
    ranFunctionDefinitionInsert_flag := 0x00000008
    ranFunctionDefinitionControl_flag := 0x00000010
    ranFunctionDefinitionPolicy_flag := 0x00000020
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(6)
    self.RanFunctionName.Unpack(stream)// p8
    if (ranFunctionDefinitionEventTrigger_flag & _flags) == ranFunctionDefinitionEventTrigger_flag { //cond2
        self.RanFunctionDefinitionEventTrigger = &RANFunctionDefinitionEventTrigger{}//7{'type': 'RANFunctionDefinition-EventTrigger', 'name': 'ranFunctionDefinition-EventTrigger', 'optional': True}
        self.RanFunctionDefinitionEventTrigger.Unpack(stream)// p8
    }
    if (ranFunctionDefinitionReport_flag & _flags) == ranFunctionDefinitionReport_flag { //cond2
        self.RanFunctionDefinitionReport = &RANFunctionDefinitionReport{}//7{'type': 'RANFunctionDefinition-Report', 'name': 'ranFunctionDefinition-Report', 'optional': True}
        self.RanFunctionDefinitionReport.Unpack(stream)// p8
    }
    if (ranFunctionDefinitionInsert_flag & _flags) == ranFunctionDefinitionInsert_flag { //cond2
        self.RanFunctionDefinitionInsert = &RANFunctionDefinitionInsert{}//7{'type': 'RANFunctionDefinition-Insert', 'name': 'ranFunctionDefinition-Insert', 'optional': True}
        self.RanFunctionDefinitionInsert.Unpack(stream)// p8
    }
    if (ranFunctionDefinitionControl_flag & _flags) == ranFunctionDefinitionControl_flag { //cond2
        self.RanFunctionDefinitionControl = &RANFunctionDefinitionControl{}//7{'type': 'RANFunctionDefinition-Control', 'name': 'ranFunctionDefinition-Control', 'optional': True}
        self.RanFunctionDefinitionControl.Unpack(stream)// p8
    }
    if (ranFunctionDefinitionPolicy_flag & _flags) == ranFunctionDefinitionPolicy_flag { //cond2
        self.RanFunctionDefinitionPolicy = &RANFunctionDefinitionPolicy{}//7{'type': 'RANFunctionDefinition-Policy', 'name': 'ranFunctionDefinition-Policy', 'optional': True}
        self.RanFunctionDefinitionPolicy.Unpack(stream)// p8
    }
    if (_flags & ext_flag) > 0 {  //ranFunctionDefinition-Query
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanFunctionDefinitionQuery = &RANFunctionDefinitionQuery{}//7{'type': 'RANFunctionDefinition-Query', 'name': 'ranFunctionDefinition-Query', 'optional': True}
        self.RanFunctionDefinitionQuery.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMRCRANFunctionDefinition) Pack(stream *Stream) {
    const ranFunctionDefinitionEventTrigger_flag uint = 0x00000002
    const ranFunctionDefinitionReport_flag uint = 0x00000004
    const ranFunctionDefinitionInsert_flag uint = 0x00000008
    const ranFunctionDefinitionControl_flag uint = 0x00000010
    const ranFunctionDefinitionPolicy_flag uint = 0x00000020
    const ext_flag int = 0x00000001
    const ranFunctionDefinitionQuery_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(6)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanFunctionName.Pack(stream)
    if self.RanFunctionDefinitionEventTrigger != nil { 
        _flags |= ranFunctionDefinitionEventTrigger_flag
        self.RanFunctionDefinitionEventTrigger.Pack(stream)
    }//end of optional
    if self.RanFunctionDefinitionReport != nil { 
        _flags |= ranFunctionDefinitionReport_flag
        self.RanFunctionDefinitionReport.Pack(stream)
    }//end of optional
    if self.RanFunctionDefinitionInsert != nil { 
        _flags |= ranFunctionDefinitionInsert_flag
        self.RanFunctionDefinitionInsert.Pack(stream)
    }//end of optional
    if self.RanFunctionDefinitionControl != nil { 
        _flags |= ranFunctionDefinitionControl_flag
        self.RanFunctionDefinitionControl.Pack(stream)
    }//end of optional
    if self.RanFunctionDefinitionPolicy != nil { 
        _flags |= ranFunctionDefinitionPolicy_flag
        self.RanFunctionDefinitionPolicy.Pack(stream)
    }//end of optional
    if self.RanFunctionDefinitionQuery != nil { 
        //extension Group: ranFunctionDefinition-Query 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranFunctionDefinitionQuery_flag
        self.RanFunctionDefinitionQuery.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 7)
}//end

type RANFunctionDefinitionEventTrigger_RanCellIdentificationParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'CellIdentification-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-CellIdentificationParameters-List', 'optional': True}
    Items []CellIdentificationRANParameterItem
}
type RANFunctionDefinitionEventTrigger_RanUEIdentificationParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'UEIdentification-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-UEIdentificationParameters-List', 'optional': True}
    Items []UEIdentificationRANParameterItem
}
type RANFunctionDefinitionEventTrigger_RanCallProcessTypesList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-EventTrigger-CallProcess-Item'}, 'size': [(1, 'maxnoofCallProcessTypes')], 'name': 'ran-CallProcessTypes-List', 'optional': True}
    Items []RANFunctionDefinitionEventTriggerCallProcessItem
}
type RANFunctionDefinitionEventTrigger_RanL2ParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'L2Parameters-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-L2Parameters-List', 'optional': True}
    Items []L2ParametersRANParameterItem
}
type RANFunctionDefinitionEventTrigger_RicEventTriggerStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-EventTrigger-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-EventTriggerStyle-List'}
    Items []RANFunctionDefinitionEventTriggerStyleItem
}
type RANFunctionDefinitionEventTrigger struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-EventTrigger-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-EventTriggerStyle-List'}, {'type': 'SEQUENCE OF', 'element': {'type': 'L2Parameters-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-L2Parameters-List', 'optional': True}, {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-EventTrigger-CallProcess-Item'}, 'size': [(1, 'maxnoofCallProcessTypes')], 'name': 'ran-CallProcessTypes-List', 'optional': True}, {'type': 'SEQUENCE OF', 'element': {'type': 'UEIdentification-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-UEIdentificationParameters-List', 'optional': True}, {'type': 'SEQUENCE OF', 'element': {'type': 'CellIdentification-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-CellIdentificationParameters-List', 'optional': True}, None]
    RicEventTriggerStyleList RANFunctionDefinitionEventTrigger_RicEventTriggerStyleList
    RanL2ParametersList *RANFunctionDefinitionEventTrigger_RanL2ParametersList
    RanCallProcessTypesList *RANFunctionDefinitionEventTrigger_RanCallProcessTypesList
    RanUEIdentificationParametersList *RANFunctionDefinitionEventTrigger_RanUEIdentificationParametersList
    RanCellIdentificationParametersList *RANFunctionDefinitionEventTrigger_RanCellIdentificationParametersList
}

func (self * RANFunctionDefinitionEventTrigger) Unpack(stream *Stream) {
    ranL2ParametersList_flag := 0x00000002
    ranCallProcessTypesList_flag := 0x00000004
    ranUEIdentificationParametersList_flag := 0x00000008
    ranCellIdentificationParametersList_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    var Unpack_ricEventTriggerStyleList = func(stream *Stream, self *RANFunctionDefinitionEventTrigger_RicEventTriggerStyleList){// Seq6 RANFunctionDefinitionEventTrigger {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-EventTrigger-Style-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-EventTriggerStyle-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]RANFunctionDefinitionEventTriggerStyleItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricEventTriggerStyleList(stream, &self.RicEventTriggerStyleList)// p2
    if (ranL2ParametersList_flag & _flags) == ranL2ParametersList_flag { //cond1
        var Unpack_ranL2ParametersList = func(stream *Stream, self *RANFunctionDefinitionEventTrigger_RanL2ParametersList){// Seq6 RANFunctionDefinitionEventTrigger {'type': 'SEQUENCE OF', 'element': {'type': 'L2Parameters-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-L2Parameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]L2ParametersRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanL2ParametersList = &RANFunctionDefinitionEventTrigger_RanL2ParametersList{}//3
        Unpack_ranL2ParametersList(stream, self.RanL2ParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'L2Parameters-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-L2Parameters-List', 'optional': True}
    }
    if (ranCallProcessTypesList_flag & _flags) == ranCallProcessTypesList_flag { //cond1
        var Unpack_ranCallProcessTypesList = func(stream *Stream, self *RANFunctionDefinitionEventTrigger_RanCallProcessTypesList){// Seq6 RANFunctionDefinitionEventTrigger {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-EventTrigger-CallProcess-Item'}, 'size': [(1, 'maxnoofCallProcessTypes')], 'name': 'ran-CallProcessTypes-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]RANFunctionDefinitionEventTriggerCallProcessItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanCallProcessTypesList = &RANFunctionDefinitionEventTrigger_RanCallProcessTypesList{}//3
        Unpack_ranCallProcessTypesList(stream, self.RanCallProcessTypesList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-EventTrigger-CallProcess-Item'}, 'size': [(1, 'maxnoofCallProcessTypes')], 'name': 'ran-CallProcessTypes-List', 'optional': True}
    }
    if (ranUEIdentificationParametersList_flag & _flags) == ranUEIdentificationParametersList_flag { //cond1
        var Unpack_ranUEIdentificationParametersList = func(stream *Stream, self *RANFunctionDefinitionEventTrigger_RanUEIdentificationParametersList){// Seq6 RANFunctionDefinitionEventTrigger {'type': 'SEQUENCE OF', 'element': {'type': 'UEIdentification-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-UEIdentificationParameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]UEIdentificationRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanUEIdentificationParametersList = &RANFunctionDefinitionEventTrigger_RanUEIdentificationParametersList{}//3
        Unpack_ranUEIdentificationParametersList(stream, self.RanUEIdentificationParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'UEIdentification-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-UEIdentificationParameters-List', 'optional': True}
    }
    if (ranCellIdentificationParametersList_flag & _flags) == ranCellIdentificationParametersList_flag { //cond1
        var Unpack_ranCellIdentificationParametersList = func(stream *Stream, self *RANFunctionDefinitionEventTrigger_RanCellIdentificationParametersList){// Seq6 RANFunctionDefinitionEventTrigger {'type': 'SEQUENCE OF', 'element': {'type': 'CellIdentification-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-CellIdentificationParameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]CellIdentificationRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanCellIdentificationParametersList = &RANFunctionDefinitionEventTrigger_RanCellIdentificationParametersList{}//3
        Unpack_ranCellIdentificationParametersList(stream, self.RanCellIdentificationParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'CellIdentification-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-CellIdentificationParameters-List', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionEventTrigger) Pack(stream *Stream) {
    const ranL2ParametersList_flag uint = 0x00000002
    const ranCallProcessTypesList_flag uint = 0x00000004
    const ranUEIdentificationParametersList_flag uint = 0x00000008
    const ranCellIdentificationParametersList_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricEventTriggerStyleList = func(stream *Stream, self RANFunctionDefinitionEventTrigger_RicEventTriggerStyleList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricEventTriggerStyleList(stream, self.RicEventTriggerStyleList) //f2
    if self.RanL2ParametersList != nil { //YY
        _flags |= ranL2ParametersList_flag
        var Pack_ranL2ParametersList = func(stream *Stream, self RANFunctionDefinitionEventTrigger_RanL2ParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranL2ParametersList(stream, *self.RanL2ParametersList) //f1
    }//end of optional
    if self.RanCallProcessTypesList != nil { //YY
        _flags |= ranCallProcessTypesList_flag
        var Pack_ranCallProcessTypesList = func(stream *Stream, self RANFunctionDefinitionEventTrigger_RanCallProcessTypesList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranCallProcessTypesList(stream, *self.RanCallProcessTypesList) //f1
    }//end of optional
    if self.RanUEIdentificationParametersList != nil { //YY
        _flags |= ranUEIdentificationParametersList_flag
        var Pack_ranUEIdentificationParametersList = func(stream *Stream, self RANFunctionDefinitionEventTrigger_RanUEIdentificationParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranUEIdentificationParametersList(stream, *self.RanUEIdentificationParametersList) //f1
    }//end of optional
    if self.RanCellIdentificationParametersList != nil { //YY
        _flags |= ranCellIdentificationParametersList_flag
        var Pack_ranCellIdentificationParametersList = func(stream *Stream, self RANFunctionDefinitionEventTrigger_RanCellIdentificationParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranCellIdentificationParametersList(stream, *self.RanCellIdentificationParametersList) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

type RANFunctionDefinitionEventTriggerStyleItem struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-EventTriggerStyle-Type'}, {'type': 'RIC-Style-Name', 'name': 'ric-EventTriggerStyle-Name'}, {'type': 'RIC-Format-Type', 'name': 'ric-EventTriggerFormat-Type'}, None]
    RicEventTriggerStyleType RICStyleType
    RicEventTriggerStyleName RICStyleName
    RicEventTriggerFormatType RICFormatType
}

func (self * RANFunctionDefinitionEventTriggerStyleItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicEventTriggerStyleType.Unpack(stream)// p8
    self.RicEventTriggerStyleName.Unpack(stream)// p8
    self.RicEventTriggerFormatType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionEventTriggerStyleItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicEventTriggerStyleType.Pack(stream)
    self.RicEventTriggerStyleName.Pack(stream)
    self.RicEventTriggerFormatType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type L2ParametersRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * L2ParametersRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * L2ParametersRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type UEIdentificationRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * UEIdentificationRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIdentificationRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type CellIdentificationRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * CellIdentificationRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CellIdentificationRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANFunctionDefinitionEventTriggerCallProcessItem_CallProcessBreakpointsList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-EventTrigger-Breakpoint-Item'}, 'size': [(1, 'maxnoofCallProcessBreakpoints')], 'name': 'callProcessBreakpoints-List'}
    Items []RANFunctionDefinitionEventTriggerBreakpointItem
}
type RANFunctionDefinitionEventTriggerCallProcessItem struct { // [{'type': 'RIC-CallProcessType-ID', 'name': 'callProcessType-ID'}, {'type': 'RIC-CallProcessType-Name', 'name': 'callProcessType-Name'}, {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-EventTrigger-Breakpoint-Item'}, 'size': [(1, 'maxnoofCallProcessBreakpoints')], 'name': 'callProcessBreakpoints-List'}, None]
    CallProcessTypeID RICCallProcessTypeID
    CallProcessTypeName RICCallProcessTypeName
    CallProcessBreakpointsList RANFunctionDefinitionEventTriggerCallProcessItem_CallProcessBreakpointsList
}

func (self * RANFunctionDefinitionEventTriggerCallProcessItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.CallProcessTypeID.Unpack(stream)// p8
    self.CallProcessTypeName.Unpack(stream)// p8
    var Unpack_callProcessBreakpointsList = func(stream *Stream, self *RANFunctionDefinitionEventTriggerCallProcessItem_CallProcessBreakpointsList){// Seq6 RANFunctionDefinitionEventTriggerCallProcessItem {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-EventTrigger-Breakpoint-Item'}, 'size': [(1, 'maxnoofCallProcessBreakpoints')], 'name': 'callProcessBreakpoints-List'}
        _size := stream.get_listsize(65535)
        _size += 1
        self.Items = make([]RANFunctionDefinitionEventTriggerBreakpointItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_callProcessBreakpointsList(stream, &self.CallProcessBreakpointsList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionEventTriggerCallProcessItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CallProcessTypeID.Pack(stream)
    self.CallProcessTypeName.Pack(stream)
    var Pack_callProcessBreakpointsList = func(stream *Stream, self RANFunctionDefinitionEventTriggerCallProcessItem_CallProcessBreakpointsList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 65535)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_callProcessBreakpointsList(stream, self.CallProcessBreakpointsList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANFunctionDefinitionEventTriggerBreakpointItem_RanCallProcessBreakpointParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'CallProcessBreakpoint-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-CallProcessBreakpointParameters-List', 'optional': True}
    Items []CallProcessBreakpointRANParameterItem
}
type RANFunctionDefinitionEventTriggerBreakpointItem struct { // [{'type': 'RIC-CallProcessBreakpoint-ID', 'name': 'callProcessBreakpoint-ID'}, {'type': 'RIC-CallProcessBreakpoint-Name', 'name': 'callProcessBreakpoint-Name'}, {'type': 'SEQUENCE OF', 'element': {'type': 'CallProcessBreakpoint-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-CallProcessBreakpointParameters-List', 'optional': True}, None]
    CallProcessBreakpointID RICCallProcessBreakpointID
    CallProcessBreakpointName RICCallProcessBreakpointName
    RanCallProcessBreakpointParametersList *RANFunctionDefinitionEventTriggerBreakpointItem_RanCallProcessBreakpointParametersList
}

func (self * RANFunctionDefinitionEventTriggerBreakpointItem) Unpack(stream *Stream) {
    ranCallProcessBreakpointParametersList_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.CallProcessBreakpointID.Unpack(stream)// p8
    self.CallProcessBreakpointName.Unpack(stream)// p8
    if (ranCallProcessBreakpointParametersList_flag & _flags) == ranCallProcessBreakpointParametersList_flag { //cond1
        var Unpack_ranCallProcessBreakpointParametersList = func(stream *Stream, self *RANFunctionDefinitionEventTriggerBreakpointItem_RanCallProcessBreakpointParametersList){// Seq6 RANFunctionDefinitionEventTriggerBreakpointItem {'type': 'SEQUENCE OF', 'element': {'type': 'CallProcessBreakpoint-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-CallProcessBreakpointParameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]CallProcessBreakpointRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanCallProcessBreakpointParametersList = &RANFunctionDefinitionEventTriggerBreakpointItem_RanCallProcessBreakpointParametersList{}//3
        Unpack_ranCallProcessBreakpointParametersList(stream, self.RanCallProcessBreakpointParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'CallProcessBreakpoint-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-CallProcessBreakpointParameters-List', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionEventTriggerBreakpointItem) Pack(stream *Stream) {
    const ranCallProcessBreakpointParametersList_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.CallProcessBreakpointID.Pack(stream)
    self.CallProcessBreakpointName.Pack(stream)
    if self.RanCallProcessBreakpointParametersList != nil { //YY
        _flags |= ranCallProcessBreakpointParametersList_flag
        var Pack_ranCallProcessBreakpointParametersList = func(stream *Stream, self RANFunctionDefinitionEventTriggerBreakpointItem_RanCallProcessBreakpointParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranCallProcessBreakpointParametersList(stream, *self.RanCallProcessBreakpointParametersList) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type CallProcessBreakpointRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * CallProcessBreakpointRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * CallProcessBreakpointRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANFunctionDefinitionReport_RicReportStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Report-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ReportStyle-List'}
    Items []RANFunctionDefinitionReportItem
}
type RANFunctionDefinitionReport struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Report-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ReportStyle-List'}, None]
    RicReportStyleList RANFunctionDefinitionReport_RicReportStyleList
}

func (self * RANFunctionDefinitionReport) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricReportStyleList = func(stream *Stream, self *RANFunctionDefinitionReport_RicReportStyleList){// Seq6 RANFunctionDefinitionReport {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Report-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ReportStyle-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]RANFunctionDefinitionReportItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricReportStyleList(stream, &self.RicReportStyleList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionReport) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricReportStyleList = func(stream *Stream, self RANFunctionDefinitionReport_RicReportStyleList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricReportStyleList(stream, self.RicReportStyleList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANFunctionDefinitionReportItem_RanReportParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'Report-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-ReportParameters-List', 'optional': True}
    Items []ReportRANParameterItem
}
type RANFunctionDefinitionReportItem struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-ReportStyle-Type'}, {'type': 'RIC-Style-Name', 'name': 'ric-ReportStyle-Name'}, {'type': 'RIC-Style-Type', 'name': 'ric-SupportedEventTriggerStyle-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-ReportActionFormat-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-IndicationHeaderFormat-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-IndicationMessageFormat-Type'}, {'type': 'SEQUENCE OF', 'element': {'type': 'Report-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-ReportParameters-List', 'optional': True}, None]
    RicReportStyleType RICStyleType
    RicReportStyleName RICStyleName
    RicSupportedEventTriggerStyleType RICStyleType
    RicReportActionFormatType RICFormatType
    RicIndicationHeaderFormatType RICFormatType
    RicIndicationMessageFormatType RICFormatType
    RanReportParametersList *RANFunctionDefinitionReportItem_RanReportParametersList
}

func (self * RANFunctionDefinitionReportItem) Unpack(stream *Stream) {
    ranReportParametersList_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RicReportStyleType.Unpack(stream)// p8
    self.RicReportStyleName.Unpack(stream)// p8
    self.RicSupportedEventTriggerStyleType.Unpack(stream)// p8
    self.RicReportActionFormatType.Unpack(stream)// p8
    self.RicIndicationHeaderFormatType.Unpack(stream)// p8
    self.RicIndicationMessageFormatType.Unpack(stream)// p8
    if (ranReportParametersList_flag & _flags) == ranReportParametersList_flag { //cond1
        var Unpack_ranReportParametersList = func(stream *Stream, self *RANFunctionDefinitionReportItem_RanReportParametersList){// Seq6 RANFunctionDefinitionReportItem {'type': 'SEQUENCE OF', 'element': {'type': 'Report-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-ReportParameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]ReportRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanReportParametersList = &RANFunctionDefinitionReportItem_RanReportParametersList{}//3
        Unpack_ranReportParametersList(stream, self.RanReportParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'Report-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-ReportParameters-List', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionReportItem) Pack(stream *Stream) {
    const ranReportParametersList_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicReportStyleType.Pack(stream)
    self.RicReportStyleName.Pack(stream)
    self.RicSupportedEventTriggerStyleType.Pack(stream)
    self.RicReportActionFormatType.Pack(stream)
    self.RicIndicationHeaderFormatType.Pack(stream)
    self.RicIndicationMessageFormatType.Pack(stream)
    if self.RanReportParametersList != nil { //YY
        _flags |= ranReportParametersList_flag
        var Pack_ranReportParametersList = func(stream *Stream, self RANFunctionDefinitionReportItem_RanReportParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranReportParametersList(stream, *self.RanReportParametersList) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ReportRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * ReportRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ReportRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANFunctionDefinitionInsert_RicInsertStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Insert-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-InsertStyle-List'}
    Items []RANFunctionDefinitionInsertItem
}
type RANFunctionDefinitionInsert struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Insert-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-InsertStyle-List'}, None]
    RicInsertStyleList RANFunctionDefinitionInsert_RicInsertStyleList
}

func (self * RANFunctionDefinitionInsert) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricInsertStyleList = func(stream *Stream, self *RANFunctionDefinitionInsert_RicInsertStyleList){// Seq6 RANFunctionDefinitionInsert {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Insert-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-InsertStyle-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]RANFunctionDefinitionInsertItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricInsertStyleList(stream, &self.RicInsertStyleList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionInsert) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricInsertStyleList = func(stream *Stream, self RANFunctionDefinitionInsert_RicInsertStyleList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricInsertStyleList(stream, self.RicInsertStyleList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANFunctionDefinitionInsertItem_RicInsertIndicationList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Insert-Indication-Item'}, 'size': [(1, 'maxnoofInsertIndication')], 'name': 'ric-InsertIndication-List', 'optional': True}
    Items []RANFunctionDefinitionInsertIndicationItem
}
type RANFunctionDefinitionInsertItem struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-InsertStyle-Type'}, {'type': 'RIC-Style-Name', 'name': 'ric-InsertStyle-Name'}, {'type': 'RIC-Style-Type', 'name': 'ric-SupportedEventTriggerStyle-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-ActionDefinitionFormat-Type'}, {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Insert-Indication-Item'}, 'size': [(1, 'maxnoofInsertIndication')], 'name': 'ric-InsertIndication-List', 'optional': True}, {'type': 'RIC-Format-Type', 'name': 'ric-IndicationHeaderFormat-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-IndicationMessageFormat-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-CallProcessIDFormat-Type'}, None]
    RicInsertStyleType RICStyleType
    RicInsertStyleName RICStyleName
    RicSupportedEventTriggerStyleType RICStyleType
    RicActionDefinitionFormatType RICFormatType
    RicInsertIndicationList *RANFunctionDefinitionInsertItem_RicInsertIndicationList
    RicIndicationHeaderFormatType RICFormatType
    RicIndicationMessageFormatType RICFormatType
    RicCallProcessIDFormatType RICFormatType
}

func (self * RANFunctionDefinitionInsertItem) Unpack(stream *Stream) {
    ricInsertIndicationList_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RicInsertStyleType.Unpack(stream)// p8
    self.RicInsertStyleName.Unpack(stream)// p8
    self.RicSupportedEventTriggerStyleType.Unpack(stream)// p8
    self.RicActionDefinitionFormatType.Unpack(stream)// p8
    if (ricInsertIndicationList_flag & _flags) == ricInsertIndicationList_flag { //cond1
        var Unpack_ricInsertIndicationList = func(stream *Stream, self *RANFunctionDefinitionInsertItem_RicInsertIndicationList){// Seq6 RANFunctionDefinitionInsertItem {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Insert-Indication-Item'}, 'size': [(1, 'maxnoofInsertIndication')], 'name': 'ric-InsertIndication-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]RANFunctionDefinitionInsertIndicationItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RicInsertIndicationList = &RANFunctionDefinitionInsertItem_RicInsertIndicationList{}//3
        Unpack_ricInsertIndicationList(stream, self.RicInsertIndicationList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Insert-Indication-Item'}, 'size': [(1, 'maxnoofInsertIndication')], 'name': 'ric-InsertIndication-List', 'optional': True}
    }
    self.RicIndicationHeaderFormatType.Unpack(stream)// p8
    self.RicIndicationMessageFormatType.Unpack(stream)// p8
    self.RicCallProcessIDFormatType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionInsertItem) Pack(stream *Stream) {
    const ricInsertIndicationList_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicInsertStyleType.Pack(stream)
    self.RicInsertStyleName.Pack(stream)
    self.RicSupportedEventTriggerStyleType.Pack(stream)
    self.RicActionDefinitionFormatType.Pack(stream)
    if self.RicInsertIndicationList != nil { //YY
        _flags |= ricInsertIndicationList_flag
        var Pack_ricInsertIndicationList = func(stream *Stream, self RANFunctionDefinitionInsertItem_RicInsertIndicationList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ricInsertIndicationList(stream, *self.RicInsertIndicationList) //f1
    }//end of optional
    self.RicIndicationHeaderFormatType.Pack(stream)
    self.RicIndicationMessageFormatType.Pack(stream)
    self.RicCallProcessIDFormatType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANFunctionDefinitionInsertIndicationItem_RanInsertIndicationParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'InsertIndication-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-InsertIndicationParameters-List', 'optional': True}
    Items []InsertIndicationRANParameterItem
}
type RANFunctionDefinitionInsertIndicationItem struct { // [{'type': 'RIC-InsertIndication-ID', 'name': 'ric-InsertIndication-ID'}, {'type': 'RIC-InsertIndication-Name', 'name': 'ric-InsertIndication-Name'}, {'type': 'SEQUENCE OF', 'element': {'type': 'InsertIndication-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-InsertIndicationParameters-List', 'optional': True}, None]
    RicInsertIndicationID RICInsertIndicationID
    RicInsertIndicationName RICInsertIndicationName
    RanInsertIndicationParametersList *RANFunctionDefinitionInsertIndicationItem_RanInsertIndicationParametersList
}

func (self * RANFunctionDefinitionInsertIndicationItem) Unpack(stream *Stream) {
    ranInsertIndicationParametersList_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RicInsertIndicationID.Unpack(stream)// p8
    self.RicInsertIndicationName.Unpack(stream)// p8
    if (ranInsertIndicationParametersList_flag & _flags) == ranInsertIndicationParametersList_flag { //cond1
        var Unpack_ranInsertIndicationParametersList = func(stream *Stream, self *RANFunctionDefinitionInsertIndicationItem_RanInsertIndicationParametersList){// Seq6 RANFunctionDefinitionInsertIndicationItem {'type': 'SEQUENCE OF', 'element': {'type': 'InsertIndication-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-InsertIndicationParameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]InsertIndicationRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanInsertIndicationParametersList = &RANFunctionDefinitionInsertIndicationItem_RanInsertIndicationParametersList{}//3
        Unpack_ranInsertIndicationParametersList(stream, self.RanInsertIndicationParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'InsertIndication-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-InsertIndicationParameters-List', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionInsertIndicationItem) Pack(stream *Stream) {
    const ranInsertIndicationParametersList_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicInsertIndicationID.Pack(stream)
    self.RicInsertIndicationName.Pack(stream)
    if self.RanInsertIndicationParametersList != nil { //YY
        _flags |= ranInsertIndicationParametersList_flag
        var Pack_ranInsertIndicationParametersList = func(stream *Stream, self RANFunctionDefinitionInsertIndicationItem_RanInsertIndicationParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranInsertIndicationParametersList(stream, *self.RanInsertIndicationParametersList) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type InsertIndicationRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * InsertIndicationRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * InsertIndicationRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANFunctionDefinitionControl_RicControlStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Control-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ControlStyle-List'}
    Items []RANFunctionDefinitionControlItem
}
type RANFunctionDefinitionControl struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Control-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ControlStyle-List'}, None]
    RicControlStyleList RANFunctionDefinitionControl_RicControlStyleList
}

func (self * RANFunctionDefinitionControl) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricControlStyleList = func(stream *Stream, self *RANFunctionDefinitionControl_RicControlStyleList){// Seq6 RANFunctionDefinitionControl {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Control-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ControlStyle-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]RANFunctionDefinitionControlItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricControlStyleList(stream, &self.RicControlStyleList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionControl) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricControlStyleList = func(stream *Stream, self RANFunctionDefinitionControl_RicControlStyleList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricControlStyleList(stream, self.RicControlStyleList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANFunctionDefinitionControlItem_RanControlOutcomeParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ControlOutcome-RANParameter-Item'}, 'size': [(1, 'maxnoofRANOutcomeParameters')], 'name': 'ran-ControlOutcomeParameters-List', 'optional': True}
    Items []ControlOutcomeRANParameterItem
}
type RANFunctionDefinitionControlItem_RicControlActionList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Control-Action-Item'}, 'size': [(1, 'maxnoofControlAction')], 'name': 'ric-ControlAction-List', 'optional': True}
    Items []RANFunctionDefinitionControlActionItem
}
type RANFunctionDefinitionControlItem struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-ControlStyle-Type'}, {'type': 'RIC-Style-Name', 'name': 'ric-ControlStyle-Name'}, {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Control-Action-Item'}, 'size': [(1, 'maxnoofControlAction')], 'name': 'ric-ControlAction-List', 'optional': True}, {'type': 'RIC-Format-Type', 'name': 'ric-ControlHeaderFormat-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-ControlMessageFormat-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-CallProcessIDFormat-Type', 'optional': True}, {'type': 'RIC-Format-Type', 'name': 'ric-ControlOutcomeFormat-Type'}, {'type': 'SEQUENCE OF', 'element': {'type': 'ControlOutcome-RANParameter-Item'}, 'size': [(1, 'maxnoofRANOutcomeParameters')], 'name': 'ran-ControlOutcomeParameters-List', 'optional': True}, None, {'type': 'ListOfAdditionalSupportedFormats-UEGroupControl', 'name': 'listOfAdditionalSupportedFormats-UEGroupControl', 'optional': True}]
    RicControlStyleType RICStyleType
    RicControlStyleName RICStyleName
    RicControlActionList *RANFunctionDefinitionControlItem_RicControlActionList
    RicControlHeaderFormatType RICFormatType
    RicControlMessageFormatType RICFormatType
    RicCallProcessIDFormatType *RICFormatType
    RicControlOutcomeFormatType RICFormatType
    RanControlOutcomeParametersList *RANFunctionDefinitionControlItem_RanControlOutcomeParametersList
    ListOfAdditionalSupportedFormatsUEGroupControl *ListOfAdditionalSupportedFormatsUEGroupControl
}

func (self * RANFunctionDefinitionControlItem) Unpack(stream *Stream) {
    ricControlActionList_flag := 0x00000002
    ricCallProcessIDFormatType_flag := 0x00000004
    ranControlOutcomeParametersList_flag := 0x00000008
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(4)
    self.RicControlStyleType.Unpack(stream)// p8
    self.RicControlStyleName.Unpack(stream)// p8
    if (ricControlActionList_flag & _flags) == ricControlActionList_flag { //cond1
        var Unpack_ricControlActionList = func(stream *Stream, self *RANFunctionDefinitionControlItem_RicControlActionList){// Seq6 RANFunctionDefinitionControlItem {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Control-Action-Item'}, 'size': [(1, 'maxnoofControlAction')], 'name': 'ric-ControlAction-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]RANFunctionDefinitionControlActionItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RicControlActionList = &RANFunctionDefinitionControlItem_RicControlActionList{}//3
        Unpack_ricControlActionList(stream, self.RicControlActionList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Control-Action-Item'}, 'size': [(1, 'maxnoofControlAction')], 'name': 'ric-ControlAction-List', 'optional': True}
    }
    self.RicControlHeaderFormatType.Unpack(stream)// p8
    self.RicControlMessageFormatType.Unpack(stream)// p8
    if (ricCallProcessIDFormatType_flag & _flags) == ricCallProcessIDFormatType_flag { //cond2
        self.RicCallProcessIDFormatType = &RICFormatType{}//7{'type': 'RIC-Format-Type', 'name': 'ric-CallProcessIDFormat-Type', 'optional': True}
        self.RicCallProcessIDFormatType.Unpack(stream)// p8
    }
    self.RicControlOutcomeFormatType.Unpack(stream)// p8
    if (ranControlOutcomeParametersList_flag & _flags) == ranControlOutcomeParametersList_flag { //cond1
        var Unpack_ranControlOutcomeParametersList = func(stream *Stream, self *RANFunctionDefinitionControlItem_RanControlOutcomeParametersList){// Seq6 RANFunctionDefinitionControlItem {'type': 'SEQUENCE OF', 'element': {'type': 'ControlOutcome-RANParameter-Item'}, 'size': [(1, 'maxnoofRANOutcomeParameters')], 'name': 'ran-ControlOutcomeParameters-List', 'optional': True}
            _size := stream.get_listsize(255)
            _size += 1
            self.Items = make([]ControlOutcomeRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanControlOutcomeParametersList = &RANFunctionDefinitionControlItem_RanControlOutcomeParametersList{}//3
        Unpack_ranControlOutcomeParametersList(stream, self.RanControlOutcomeParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'ControlOutcome-RANParameter-Item'}, 'size': [(1, 'maxnoofRANOutcomeParameters')], 'name': 'ran-ControlOutcomeParameters-List', 'optional': True}
    }
    if (_flags & ext_flag) > 0 {  //listOfAdditionalSupportedFormats-UEGroupControl
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.ListOfAdditionalSupportedFormatsUEGroupControl = &ListOfAdditionalSupportedFormatsUEGroupControl{}//7{'type': 'ListOfAdditionalSupportedFormats-UEGroupControl', 'name': 'listOfAdditionalSupportedFormats-UEGroupControl', 'optional': True}
        self.ListOfAdditionalSupportedFormatsUEGroupControl.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionControlItem) Pack(stream *Stream) {
    const ricControlActionList_flag uint = 0x00000002
    const ricCallProcessIDFormatType_flag uint = 0x00000004
    const ranControlOutcomeParametersList_flag uint = 0x00000008
    const ext_flag int = 0x00000001
    const listOfAdditionalSupportedFormatsUEGroupControl_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(4)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicControlStyleType.Pack(stream)
    self.RicControlStyleName.Pack(stream)
    if self.RicControlActionList != nil { //YY
        _flags |= ricControlActionList_flag
        var Pack_ricControlActionList = func(stream *Stream, self RANFunctionDefinitionControlItem_RicControlActionList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ricControlActionList(stream, *self.RicControlActionList) //f1
    }//end of optional
    self.RicControlHeaderFormatType.Pack(stream)
    self.RicControlMessageFormatType.Pack(stream)
    if self.RicCallProcessIDFormatType != nil { 
        _flags |= ricCallProcessIDFormatType_flag
        self.RicCallProcessIDFormatType.Pack(stream)
    }//end of optional
    self.RicControlOutcomeFormatType.Pack(stream)
    if self.RanControlOutcomeParametersList != nil { //YY
        _flags |= ranControlOutcomeParametersList_flag
        var Pack_ranControlOutcomeParametersList = func(stream *Stream, self RANFunctionDefinitionControlItem_RanControlOutcomeParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 255)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranControlOutcomeParametersList(stream, *self.RanControlOutcomeParametersList) //f1
    }//end of optional
    if self.ListOfAdditionalSupportedFormatsUEGroupControl != nil { 
        //extension Group: listOfAdditionalSupportedFormats-UEGroupControl 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= listOfAdditionalSupportedFormatsUEGroupControl_flag
        self.ListOfAdditionalSupportedFormatsUEGroupControl.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

type ControlOutcomeRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * ControlOutcomeRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ControlOutcomeRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANFunctionDefinitionControlActionItem_RanControlActionParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'ControlAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-ControlActionParameters-List', 'optional': True}
    Items []ControlActionRANParameterItem
}
type RANFunctionDefinitionControlActionItem struct { // [{'type': 'RIC-ControlAction-ID', 'name': 'ric-ControlAction-ID'}, {'type': 'RIC-ControlAction-Name', 'name': 'ric-ControlAction-Name'}, {'type': 'SEQUENCE OF', 'element': {'type': 'ControlAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-ControlActionParameters-List', 'optional': True}, None, {'type': 'ENUMERATED', 'values': [('true', 0), ('false', 1), None], 'name': 'ueGroup-ControlAction-Supported'}]
    RicControlActionID RICControlActionID
    RicControlActionName RICControlActionName
    RanControlActionParametersList *RANFunctionDefinitionControlActionItem_RanControlActionParametersList
    UeGroupControlActionSupported ENUMERATED
}

func (self * RANFunctionDefinitionControlActionItem) Unpack(stream *Stream) {
    ranControlActionParametersList_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RicControlActionID.Unpack(stream)// p8
    self.RicControlActionName.Unpack(stream)// p8
    if (ranControlActionParametersList_flag & _flags) == ranControlActionParametersList_flag { //cond1
        var Unpack_ranControlActionParametersList = func(stream *Stream, self *RANFunctionDefinitionControlActionItem_RanControlActionParametersList){// Seq6 RANFunctionDefinitionControlActionItem {'type': 'SEQUENCE OF', 'element': {'type': 'ControlAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-ControlActionParameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]ControlActionRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanControlActionParametersList = &RANFunctionDefinitionControlActionItem_RanControlActionParametersList{}//3
        Unpack_ranControlActionParametersList(stream, self.RanControlActionParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'ControlAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-ControlActionParameters-List', 'optional': True}
    }
    if (_flags & ext_flag) > 0 {  //ueGroup-ControlAction-Supported
        _ecount, _extflags = stream.parsef_sequence_ext(0)
    }
    var Unpack_ueGroupControlActionSupported = func(st *Stream, self *ENUMERATED) {
        self.Value = st.parsef_Enumerated(2, 2, 1)
    }
    Unpack_ueGroupControlActionSupported(stream, &self.UeGroupControlActionSupported)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionControlActionItem) Pack(stream *Stream) {
    const ranControlActionParametersList_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicControlActionID.Pack(stream)
    self.RicControlActionName.Pack(stream)
    if self.RanControlActionParametersList != nil { //YY
        _flags |= ranControlActionParametersList_flag
        var Pack_ranControlActionParametersList = func(stream *Stream, self RANFunctionDefinitionControlActionItem_RanControlActionParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranControlActionParametersList(stream, *self.RanControlActionParametersList) //f1
    }//end of optional
    //extension Group: ueGroup-ControlAction-Supported 0 
    if _extPresent == false {
        stream.set_bits(-1, 7)
        _extReserve = stream.reserve_flags(0)
        _extPresent = true
        _flags |= uint(ext_flag)
    }//optind
    _lenOReserve := stream.reserve_flags(8)
    var Pack_ueGroupControlActionSupported = func(st *Stream, self ENUMERATED) {
        st.formatf_Enumerated(self.Value, 2, 2, 1)
    }
    Pack_ueGroupControlActionSupported(stream, self.UeGroupControlActionSupported) //f2
    stream.set_openlen(_lenOReserve, 8, true)//item
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type ControlActionRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * ControlActionRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * ControlActionRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *ListOfAdditionalSupportedFormatsUEGroupControl) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(64)
    _size += 0
    self.Items = make([]AdditionalSupportedFormatUEGroupControl, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *ListOfAdditionalSupportedFormatsUEGroupControl) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-0, 64)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type ListOfAdditionalSupportedFormatsUEGroupControl struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'AdditionalSupportedFormat-UEGroupControl'}, 'size': [(0, 'maxnoofFormatTypes')]}
    Items []AdditionalSupportedFormatUEGroupControl
}

type AdditionalSupportedFormatUEGroupControl struct { // [{'type': 'RIC-Format-Type', 'name': 'ric-ControlHeaderFormat-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-ControlMessageFormat-Type'}, None, {'type': 'RIC-Format-Type', 'name': 'ric-ControlOutcomeFormat-Type', 'optional': True}]
    RicControlHeaderFormatType RICFormatType
    RicControlMessageFormatType RICFormatType
    RicControlOutcomeFormatType *RICFormatType
}

func (self * AdditionalSupportedFormatUEGroupControl) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicControlHeaderFormatType.Unpack(stream)// p8
    self.RicControlMessageFormatType.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ric-ControlOutcomeFormat-Type
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RicControlOutcomeFormatType = &RICFormatType{}//7{'type': 'RIC-Format-Type', 'name': 'ric-ControlOutcomeFormat-Type', 'optional': True}
        self.RicControlOutcomeFormatType.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * AdditionalSupportedFormatUEGroupControl) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ricControlOutcomeFormatType_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicControlHeaderFormatType.Pack(stream)
    self.RicControlMessageFormatType.Pack(stream)
    if self.RicControlOutcomeFormatType != nil { 
        //extension Group: ric-ControlOutcomeFormat-Type 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ricControlOutcomeFormatType_flag
        self.RicControlOutcomeFormatType.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANFunctionDefinitionPolicy_RicPolicyStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Policy-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-PolicyStyle-List'}
    Items []RANFunctionDefinitionPolicyItem
}
type RANFunctionDefinitionPolicy struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Policy-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-PolicyStyle-List'}, None]
    RicPolicyStyleList RANFunctionDefinitionPolicy_RicPolicyStyleList
}

func (self * RANFunctionDefinitionPolicy) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricPolicyStyleList = func(stream *Stream, self *RANFunctionDefinitionPolicy_RicPolicyStyleList){// Seq6 RANFunctionDefinitionPolicy {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Policy-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-PolicyStyle-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]RANFunctionDefinitionPolicyItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricPolicyStyleList(stream, &self.RicPolicyStyleList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionPolicy) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricPolicyStyleList = func(stream *Stream, self RANFunctionDefinitionPolicy_RicPolicyStyleList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricPolicyStyleList(stream, self.RicPolicyStyleList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANFunctionDefinitionPolicyItem_RicPolicyActionList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Policy-Action-Item'}, 'size': [(1, 'maxnoofPolicyAction')], 'name': 'ric-PolicyAction-List', 'optional': True}
    Items []RANFunctionDefinitionPolicyActionItem
}
type RANFunctionDefinitionPolicyItem struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-PolicyStyle-Type'}, {'type': 'RIC-Style-Name', 'name': 'ric-PolicyStyle-Name'}, {'type': 'RIC-Style-Type', 'name': 'ric-SupportedEventTriggerStyle-Type'}, {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Policy-Action-Item'}, 'size': [(1, 'maxnoofPolicyAction')], 'name': 'ric-PolicyAction-List', 'optional': True}, None]
    RicPolicyStyleType RICStyleType
    RicPolicyStyleName RICStyleName
    RicSupportedEventTriggerStyleType RICStyleType
    RicPolicyActionList *RANFunctionDefinitionPolicyItem_RicPolicyActionList
}

func (self * RANFunctionDefinitionPolicyItem) Unpack(stream *Stream) {
    ricPolicyActionList_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RicPolicyStyleType.Unpack(stream)// p8
    self.RicPolicyStyleName.Unpack(stream)// p8
    self.RicSupportedEventTriggerStyleType.Unpack(stream)// p8
    if (ricPolicyActionList_flag & _flags) == ricPolicyActionList_flag { //cond1
        var Unpack_ricPolicyActionList = func(stream *Stream, self *RANFunctionDefinitionPolicyItem_RicPolicyActionList){// Seq6 RANFunctionDefinitionPolicyItem {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Policy-Action-Item'}, 'size': [(1, 'maxnoofPolicyAction')], 'name': 'ric-PolicyAction-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]RANFunctionDefinitionPolicyActionItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RicPolicyActionList = &RANFunctionDefinitionPolicyItem_RicPolicyActionList{}//3
        Unpack_ricPolicyActionList(stream, self.RicPolicyActionList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Policy-Action-Item'}, 'size': [(1, 'maxnoofPolicyAction')], 'name': 'ric-PolicyAction-List', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionPolicyItem) Pack(stream *Stream) {
    const ricPolicyActionList_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicPolicyStyleType.Pack(stream)
    self.RicPolicyStyleName.Pack(stream)
    self.RicSupportedEventTriggerStyleType.Pack(stream)
    if self.RicPolicyActionList != nil { //YY
        _flags |= ricPolicyActionList_flag
        var Pack_ricPolicyActionList = func(stream *Stream, self RANFunctionDefinitionPolicyItem_RicPolicyActionList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ricPolicyActionList(stream, *self.RicPolicyActionList) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANFunctionDefinitionPolicyActionItem_RanPolicyConditionParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PolicyCondition-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-PolicyConditionParameters-List', 'optional': True}
    Items []PolicyConditionRANParameterItem
}
type RANFunctionDefinitionPolicyActionItem_RanPolicyActionParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'PolicyAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-PolicyActionParameters-List', 'optional': True}
    Items []PolicyActionRANParameterItem
}
type RANFunctionDefinitionPolicyActionItem struct { // [{'type': 'RIC-ControlAction-ID', 'name': 'ric-PolicyAction-ID'}, {'type': 'RIC-ControlAction-Name', 'name': 'ric-PolicyAction-Name'}, {'type': 'RIC-Format-Type', 'name': 'ric-ActionDefinitionFormat-Type'}, {'type': 'SEQUENCE OF', 'element': {'type': 'PolicyAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-PolicyActionParameters-List', 'optional': True}, {'type': 'SEQUENCE OF', 'element': {'type': 'PolicyCondition-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-PolicyConditionParameters-List', 'optional': True}, None]
    RicPolicyActionID RICControlActionID
    RicPolicyActionName RICControlActionName
    RicActionDefinitionFormatType RICFormatType
    RanPolicyActionParametersList *RANFunctionDefinitionPolicyActionItem_RanPolicyActionParametersList
    RanPolicyConditionParametersList *RANFunctionDefinitionPolicyActionItem_RanPolicyConditionParametersList
}

func (self * RANFunctionDefinitionPolicyActionItem) Unpack(stream *Stream) {
    ranPolicyActionParametersList_flag := 0x00000002
    ranPolicyConditionParametersList_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RicPolicyActionID.Unpack(stream)// p8
    self.RicPolicyActionName.Unpack(stream)// p8
    self.RicActionDefinitionFormatType.Unpack(stream)// p8
    if (ranPolicyActionParametersList_flag & _flags) == ranPolicyActionParametersList_flag { //cond1
        var Unpack_ranPolicyActionParametersList = func(stream *Stream, self *RANFunctionDefinitionPolicyActionItem_RanPolicyActionParametersList){// Seq6 RANFunctionDefinitionPolicyActionItem {'type': 'SEQUENCE OF', 'element': {'type': 'PolicyAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-PolicyActionParameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]PolicyActionRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanPolicyActionParametersList = &RANFunctionDefinitionPolicyActionItem_RanPolicyActionParametersList{}//3
        Unpack_ranPolicyActionParametersList(stream, self.RanPolicyActionParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'PolicyAction-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-PolicyActionParameters-List', 'optional': True}
    }
    if (ranPolicyConditionParametersList_flag & _flags) == ranPolicyConditionParametersList_flag { //cond1
        var Unpack_ranPolicyConditionParametersList = func(stream *Stream, self *RANFunctionDefinitionPolicyActionItem_RanPolicyConditionParametersList){// Seq6 RANFunctionDefinitionPolicyActionItem {'type': 'SEQUENCE OF', 'element': {'type': 'PolicyCondition-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-PolicyConditionParameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]PolicyConditionRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanPolicyConditionParametersList = &RANFunctionDefinitionPolicyActionItem_RanPolicyConditionParametersList{}//3
        Unpack_ranPolicyConditionParametersList(stream, self.RanPolicyConditionParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'PolicyCondition-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-PolicyConditionParameters-List', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionPolicyActionItem) Pack(stream *Stream) {
    const ranPolicyActionParametersList_flag uint = 0x00000002
    const ranPolicyConditionParametersList_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicPolicyActionID.Pack(stream)
    self.RicPolicyActionName.Pack(stream)
    self.RicActionDefinitionFormatType.Pack(stream)
    if self.RanPolicyActionParametersList != nil { //YY
        _flags |= ranPolicyActionParametersList_flag
        var Pack_ranPolicyActionParametersList = func(stream *Stream, self RANFunctionDefinitionPolicyActionItem_RanPolicyActionParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranPolicyActionParametersList(stream, *self.RanPolicyActionParametersList) //f1
    }//end of optional
    if self.RanPolicyConditionParametersList != nil { //YY
        _flags |= ranPolicyConditionParametersList_flag
        var Pack_ranPolicyConditionParametersList = func(stream *Stream, self RANFunctionDefinitionPolicyActionItem_RanPolicyConditionParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranPolicyConditionParametersList(stream, *self.RanPolicyConditionParametersList) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type PolicyActionRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * PolicyActionRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PolicyActionRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type PolicyConditionRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, None, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * PolicyConditionRANParameterItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (_flags & ext_flag) > 0 {  //ranParameter-Definition
        _ecount, _extflags = stream.parsef_sequence_ext(1)
    }
    if (_extflags & (1<< 0)) > 0 {
        stream.get_openlen()
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * PolicyConditionRANParameterItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    const ranParameterDefinition_flag uint = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        //extension Group: ranParameter-Definition 1 
        if _extPresent == false {
            stream.set_bits(0, 7)
            _extReserve = stream.reserve_flags(1)
            _extPresent = true
            _flags |= uint(ext_flag)
        }//optind
        _lenOReserve := stream.reserve_flags(8)
        _extflags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
        stream.set_openlen(_lenOReserve, 8, true)//item
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 1) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type RANFunctionDefinitionQuery_RicQueryStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Query-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-QueryStyle-List'}
    Items []RANFunctionDefinitionQueryItem
}
type RANFunctionDefinitionQuery struct { // [{'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Query-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-QueryStyle-List'}, None]
    RicQueryStyleList RANFunctionDefinitionQuery_RicQueryStyleList
}

func (self * RANFunctionDefinitionQuery) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_ricQueryStyleList = func(stream *Stream, self *RANFunctionDefinitionQuery_RicQueryStyleList){// Seq6 RANFunctionDefinitionQuery {'type': 'SEQUENCE OF', 'element': {'type': 'RANFunctionDefinition-Query-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-QueryStyle-List'}
        _size := stream.get_listsize(63)
        _size += 1
        self.Items = make([]RANFunctionDefinitionQueryItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_ricQueryStyleList(stream, &self.RicQueryStyleList)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionQuery) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_ricQueryStyleList = func(stream *Stream, self RANFunctionDefinitionQuery_RicQueryStyleList) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 63)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_ricQueryStyleList(stream, self.RicQueryStyleList) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type RANFunctionDefinitionQueryItem_RanQueryParametersList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'Query-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-QueryParameters-List', 'optional': True}
    Items []QueryRANParameterItem
}
type RANFunctionDefinitionQueryItem struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-QueryStyle-Type'}, {'type': 'RIC-Style-Name', 'name': 'ric-QueryStyle-Name'}, {'type': 'RIC-Format-Type', 'name': 'ric-QueryHeaderFormat-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-QueryDefinitionFormat-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-QueryOutcomeFormat-Type'}, {'type': 'SEQUENCE OF', 'element': {'type': 'Query-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-QueryParameters-List', 'optional': True}, None]
    RicQueryStyleType RICStyleType
    RicQueryStyleName RICStyleName
    RicQueryHeaderFormatType RICFormatType
    RicQueryDefinitionFormatType RICFormatType
    RicQueryOutcomeFormatType RICFormatType
    RanQueryParametersList *RANFunctionDefinitionQueryItem_RanQueryParametersList
}

func (self * RANFunctionDefinitionQueryItem) Unpack(stream *Stream) {
    ranQueryParametersList_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RicQueryStyleType.Unpack(stream)// p8
    self.RicQueryStyleName.Unpack(stream)// p8
    self.RicQueryHeaderFormatType.Unpack(stream)// p8
    self.RicQueryDefinitionFormatType.Unpack(stream)// p8
    self.RicQueryOutcomeFormatType.Unpack(stream)// p8
    if (ranQueryParametersList_flag & _flags) == ranQueryParametersList_flag { //cond1
        var Unpack_ranQueryParametersList = func(stream *Stream, self *RANFunctionDefinitionQueryItem_RanQueryParametersList){// Seq6 RANFunctionDefinitionQueryItem {'type': 'SEQUENCE OF', 'element': {'type': 'Query-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-QueryParameters-List', 'optional': True}
            _size := stream.get_listsize(65535)
            _size += 1
            self.Items = make([]QueryRANParameterItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RanQueryParametersList = &RANFunctionDefinitionQueryItem_RanQueryParametersList{}//3
        Unpack_ranQueryParametersList(stream, self.RanQueryParametersList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'Query-RANParameter-Item'}, 'size': [(1, 'maxnoofAssociatedRANParameters')], 'name': 'ran-QueryParameters-List', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RANFunctionDefinitionQueryItem) Pack(stream *Stream) {
    const ranQueryParametersList_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicQueryStyleType.Pack(stream)
    self.RicQueryStyleName.Pack(stream)
    self.RicQueryHeaderFormatType.Pack(stream)
    self.RicQueryDefinitionFormatType.Pack(stream)
    self.RicQueryOutcomeFormatType.Pack(stream)
    if self.RanQueryParametersList != nil { //YY
        _flags |= ranQueryParametersList_flag
        var Pack_ranQueryParametersList = func(stream *Stream, self RANFunctionDefinitionQueryItem_RanQueryParametersList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 65535)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ranQueryParametersList(stream, *self.RanQueryParametersList) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type QueryRANParameterItem struct { // [{'type': 'RANParameter-ID', 'name': 'ranParameter-ID'}, {'type': 'RANParameter-Name', 'name': 'ranParameter-name'}, {'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}, None]
    RanParameterID RANParameterID
    RanParametername RANParameterName
    RanParameterDefinition *RANParameterDefinition
}

func (self * QueryRANParameterItem) Unpack(stream *Stream) {
    ranParameterDefinition_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.RanParameterID.Unpack(stream)// p8
    self.RanParametername.Unpack(stream)// p8
    if (ranParameterDefinition_flag & _flags) == ranParameterDefinition_flag { //cond2
        self.RanParameterDefinition = &RANParameterDefinition{}//7{'type': 'RANParameter-Definition', 'name': 'ranParameter-Definition', 'optional': True}
        self.RanParameterDefinition.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * QueryRANParameterItem) Pack(stream *Stream) {
    const ranParameterDefinition_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanParameterID.Pack(stream)
    self.RanParametername.Pack(stream)
    if self.RanParameterDefinition != nil { 
        _flags |= ranParameterDefinition_flag
        self.RanParameterDefinition.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMPDU_Message struct { //[{'type': 'E2SM-RC-EventTrigger', 'name': 'eventTriggerDefinition'}, {'type': 'E2SM-RC-ActionDefinition', 'name': 'actionDefinition'}, {'type': 'E2SM-RC-IndicationHeader', 'name': 'indicationHeader'}, {'type': 'E2SM-RC-IndicationMessage', 'name': 'indicationMessage'}, {'type': 'E2SM-RC-RANFunctionDefinition', 'name': 'ranFuncDefinition'}, {'type': 'E2SM-RC-ControlHeader', 'name': 'controlHeader'}, {'type': 'E2SM-RC-ControlMessage', 'name': 'controlMessage'}, {'type': 'E2SM-RC-ControlOutcome', 'name': 'controlOutcome'}, {'type': 'E2SM-RC-QueryHeader', 'name': 'queryHeader'}, {'type': 'E2SM-RC-QueryDefinition', 'name': 'queryDefinition'}, {'type': 'E2SM-RC-QueryOutcome', 'name': 'queryOutcome'}]
    EventTriggerDefinition *E2SMRCEventTrigger
    ActionDefinition *E2SMRCActionDefinition
    IndicationHeader *E2SMRCIndicationHeader
    IndicationMessage *E2SMRCIndicationMessage
    RanFuncDefinition *E2SMRCRANFunctionDefinition
    ControlHeader *E2SMRCControlHeader
    ControlMessage *E2SMRCControlMessage
    ControlOutcome *E2SMRCControlOutcome
    QueryHeader *E2SMRCQueryHeader
    QueryDefinition *E2SMRCQueryDefinition
    QueryOutcome *E2SMRCQueryOutcome
} // E2SMPDU_Message

type E2SMPDU struct { // [{'type': 'INTEGER', 'restricted-to': [(0, 15)], 'name': 'version'}, {'type': 'CHOICE', 'members': [{'type': 'E2SM-RC-EventTrigger', 'name': 'eventTriggerDefinition'}, {'type': 'E2SM-RC-ActionDefinition', 'name': 'actionDefinition'}, {'type': 'E2SM-RC-IndicationHeader', 'name': 'indicationHeader'}, {'type': 'E2SM-RC-IndicationMessage', 'name': 'indicationMessage'}, {'type': 'E2SM-RC-RANFunctionDefinition', 'name': 'ranFuncDefinition'}, {'type': 'E2SM-RC-ControlHeader', 'name': 'controlHeader'}, {'type': 'E2SM-RC-ControlMessage', 'name': 'controlMessage'}, {'type': 'E2SM-RC-ControlOutcome', 'name': 'controlOutcome'}, {'type': 'E2SM-RC-QueryHeader', 'name': 'queryHeader'}, {'type': 'E2SM-RC-QueryDefinition', 'name': 'queryDefinition'}, {'type': 'E2SM-RC-QueryOutcome', 'name': 'queryOutcome'}], 'name': 'message'}]
    Version INTEGER
    Message E2SMPDU_Message
}

func (self * E2SMPDU) Unpack(stream *Stream) {
    var Unpack_version = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(16, 4, 0, 0)
    }
    Unpack_version(stream, &self.Version)// p2
    var Unpack_message = func(stream *Stream, self *E2SMPDU_Message) {
        //coptions := []string{"eventTriggerDefinition","actionDefinition","indicationHeader","indicationMessage","ranFuncDefinition","controlHeader","controlMessage","controlOutcome","queryHeader","queryDefinition","queryOutcome","Unknown","Unknown","Unknown","Unknown","Unknown"}
        choice := stream.get_choice(4, 0, 11)
        if choice == 0 { //ch1
            self.EventTriggerDefinition = &E2SMRCEventTrigger{}//cho6
            self.EventTriggerDefinition.Unpack(stream)
        } else if choice == 1 { //ch2
            self.ActionDefinition = &E2SMRCActionDefinition{}//cho6
            self.ActionDefinition.Unpack(stream)
        } else if choice == 2 { //ch2
            self.IndicationHeader = &E2SMRCIndicationHeader{}//cho6
            self.IndicationHeader.Unpack(stream)
        } else if choice == 3 { //ch2
            self.IndicationMessage = &E2SMRCIndicationMessage{}//cho6
            self.IndicationMessage.Unpack(stream)
        } else if choice == 4 { //ch2
            self.RanFuncDefinition = &E2SMRCRANFunctionDefinition{}//cho6
            self.RanFuncDefinition.Unpack(stream)
        } else if choice == 5 { //ch2
            self.ControlHeader = &E2SMRCControlHeader{}//cho6
            self.ControlHeader.Unpack(stream)
        } else if choice == 6 { //ch2
            self.ControlMessage = &E2SMRCControlMessage{}//cho6
            self.ControlMessage.Unpack(stream)
        } else if choice == 7 { //ch2
            self.ControlOutcome = &E2SMRCControlOutcome{}//cho6
            self.ControlOutcome.Unpack(stream)
        } else if choice == 8 { //ch2
            self.QueryHeader = &E2SMRCQueryHeader{}//cho6
            self.QueryHeader.Unpack(stream)
        } else if choice == 9 { //ch2
            self.QueryDefinition = &E2SMRCQueryDefinition{}//cho6
            self.QueryDefinition.Unpack(stream)
        } else if choice == 10 { //ch2
            self.QueryOutcome = &E2SMRCQueryOutcome{}//cho6
            self.QueryOutcome.Unpack(stream)
        }//end of if else

    }
    Unpack_message(stream, &self.Message)// p2
    return
}

func (self * E2SMPDU) Pack(stream *Stream) {
    var Pack_version = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 16, 4, 0, 0)
    }
    Pack_version(stream, self.Version) //f2
    var Pack_message = func(stream *Stream, self E2SMPDU_Message) {
        if self.EventTriggerDefinition != nil {
            stream.set_choice(0, 4, 0, 11)
            self.EventTriggerDefinition.Pack(stream)//2
        } else if self.ActionDefinition != nil {
            stream.set_choice(1, 4, 0, 11)
            self.ActionDefinition.Pack(stream)//2
        } else if self.IndicationHeader != nil {
            stream.set_choice(2, 4, 0, 11)
            self.IndicationHeader.Pack(stream)//2
        } else if self.IndicationMessage != nil {
            stream.set_choice(3, 4, 0, 11)
            self.IndicationMessage.Pack(stream)//2
        } else if self.RanFuncDefinition != nil {
            stream.set_choice(4, 4, 0, 11)
            self.RanFuncDefinition.Pack(stream)//2
        } else if self.ControlHeader != nil {
            stream.set_choice(5, 4, 0, 11)
            self.ControlHeader.Pack(stream)//2
        } else if self.ControlMessage != nil {
            stream.set_choice(6, 4, 0, 11)
            self.ControlMessage.Pack(stream)//2
        } else if self.ControlOutcome != nil {
            stream.set_choice(7, 4, 0, 11)
            self.ControlOutcome.Pack(stream)//2
        } else if self.QueryHeader != nil {
            stream.set_choice(8, 4, 0, 11)
            self.QueryHeader.Pack(stream)//2
        } else if self.QueryDefinition != nil {
            stream.set_choice(9, 4, 0, 11)
            self.QueryDefinition.Pack(stream)//2
        } else if self.QueryOutcome != nil {
            stream.set_choice(10, 4, 0, 11)
            self.QueryOutcome.Pack(stream)//2
        }

    }
    Pack_message(stream, self.Message) //f2
}//end

type ProcedureCode struct {
  Value uint64
}
func (self *ProcedureCode) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(256, 8, 0, 0)
}
func (self * ProcedureCode) Pack(st *Stream){
    st.formatf_Integer(self.Value, 256, 8, 0, 0)
}
var maxE1APid uint64 = 65535
var maxF1APid uint64 = 4
var maxEARFCN uint64 = 65535
var maxNRARFCN uint64 = 3279165
var maxnoofNrCellBands uint64 = 32
var maxNrofSSBs1 uint64 = 63
var maxnoofMessages uint64 = 65535
var maxnoofE2InfoChanges uint64 = 65535
var maxnoofUEInfoChanges uint64 = 65535
var maxnoofRRCstate uint64 = 8
var maxnoofParametersToReport uint64 = 65535
var maxnoofPolicyConditions uint64 = 65535
var maxnoofAssociatedRANParameters uint64 = 65535
var maxnoofUEID uint64 = 65535
var maxnoofCellID uint64 = 65535
var maxnoofRANOutcomeParameters uint64 = 255
var maxnoofParametersinStructure uint64 = 65535
var maxnoofItemsinList uint64 = 65535
var maxnoofUEInfo uint64 = 65535
var maxnoofCellInfo uint64 = 65535
var maxnoofUEeventInfo uint64 = 65535
var maxnoofRANparamTest uint64 = 255
var maxnoofNeighbourCell uint64 = 65535
var maxnoofRICStyles uint64 = 63
var maxnoofCallProcessTypes uint64 = 65535
var maxnoofCallProcessBreakpoints uint64 = 65535
var maxnoofInsertIndication uint64 = 65535
var maxnoofControlAction uint64 = 65535
var maxnoofPolicyAction uint64 = 65535
var maxnoofInsertIndicationActions uint64 = 63
var maxnoofMulCtrlActions uint64 = 63
var maxGroupDefinitionIdentifierParameters uint64 = 255
var maxnoofAssociatedEntityFilters uint64 = 255
var maxnoofFormatTypes uint64 = 63
var idEventTriggerDefinition uint64 = 0
const ProcedureCodeEventTriggerDefinition = 0
var idActionDefinition uint64 = 1
const ProcedureCodeActionDefinition = 1
var idIndicationHeader uint64 = 2
const ProcedureCodeIndicationHeader = 2
var idIndicationMessage uint64 = 3
const ProcedureCodeIndicationMessage = 3
var idQueryDefinition uint64 = 4
const ProcedureCodeQueryDefinition = 4
var idControlOutcome uint64 = 5
const ProcedureCodeControlOutcome = 5
var idQueryOutcome uint64 = 6
const ProcedureCodeQueryOutcome = 6
var idQueryHeader uint64 = 7
const ProcedureCodeQueryHeader = 7
var idRANFunctionDefinition uint64 = 8
const ProcedureCodeRANFunctionDefinition = 8
