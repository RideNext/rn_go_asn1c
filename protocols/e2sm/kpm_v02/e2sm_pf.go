
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package e2sm
import (
  log "github.com/sirupsen/logrus"
)
var version = "vkpm_v02"

func fmte2sm() {log.Debug("e2sm")}
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

type InterfaceIDXn struct { // [{'type': 'GlobalRANNodeID', 'name': 'global-NG-RAN-ID'}, None]
    GlobalNGRANID GlobalRANNodeID
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

type InterfaceIDF1 struct { // [{'type': 'GlobalRANNodeID', 'name': 'global-NG-RAN-ID'}, {'type': 'GNB-DU-ID', 'name': 'gNB-DU-ID'}, None]
    GlobalNGRANID GlobalRANNodeID
    GNBDUID GNBDUID
}

func (self * InterfaceIDF1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GlobalNGRANID.Unpack(stream)// p8
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
    self.GlobalNGRANID.Pack(stream)
    self.GNBDUID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type InterfaceIDE1 struct { // [{'type': 'GlobalRANNodeID', 'name': 'global-NG-RAN-ID'}, {'type': 'GNB-CU-UP-ID', 'name': 'gNB-CU-UP-ID'}, None]
    GlobalNGRANID GlobalRANNodeID
    GNBCUUPID GNBCUUPID
}

func (self * InterfaceIDE1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.GlobalNGRANID.Unpack(stream)// p8
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
    self.GlobalNGRANID.Pack(stream)
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

type UEIDGNB struct { // [{'type': 'AMF-UE-NGAP-ID', 'name': 'amf-UE-NGAP-ID'}, {'type': 'GUAMI', 'name': 'guami'}, {'type': 'UEID-GNB-CU-F1AP-ID-List', 'name': 'gNB-CU-UE-F1AP-ID-List', 'optional': True}, {'type': 'UEID-GNB-CU-CP-E1AP-ID-List', 'name': 'gNB-CU-CP-UE-E1AP-ID-List', 'optional': True}, {'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}, {'type': 'NG-RANnodeUEXnAPID', 'name': 'm-NG-RAN-UE-XnAP-ID', 'optional': True}, {'type': 'GlobalGNB-ID', 'name': 'globalGNB-ID', 'optional': True}, None]
    AmfUENGAPID AMFUENGAPID
    Guami GUAMI
    GNBCUUEF1APIDList *UEIDGNBCUF1APIDList
    GNBCUCPUEE1APIDList *UEIDGNBCUCPE1APIDList
    RanUEID *RANUEID
    MNGRANUEXnAPID *NGRANnodeUEXnAPID
    GlobalGNBID *GlobalGNBID
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
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 6)
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

type UEIDGNBDU struct { // [{'type': 'GNB-CU-UE-F1AP-ID', 'name': 'gNB-CU-UE-F1AP-ID'}, {'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}, None]
    GNBCUUEF1APID GNBCUUEF1APID
    RanUEID *RANUEID
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
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDGNBDU) Pack(stream *Stream) {
    const ranUEID_flag uint = 0x00000002
    const ext_flag int = 0x00000001
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
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
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

type UEIDNGENB struct { // [{'type': 'AMF-UE-NGAP-ID', 'name': 'amf-UE-NGAP-ID'}, {'type': 'GUAMI', 'name': 'guami'}, {'type': 'NGENB-CU-UE-W1AP-ID', 'name': 'ng-eNB-CU-UE-W1AP-ID', 'optional': True}, {'type': 'NG-RANnodeUEXnAPID', 'name': 'm-NG-RAN-UE-XnAP-ID', 'optional': True}, {'type': 'GlobalNgENB-ID', 'name': 'globalNgENB-ID', 'optional': True}, None]
    AmfUENGAPID AMFUENGAPID
    Guami GUAMI
    NgeNBCUUEW1APID *NGENBCUUEW1APID
    MNGRANUEXnAPID *NGRANnodeUEXnAPID
    GlobalNgENBID *GlobalNgENBID
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
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDNGENB) Pack(stream *Stream) {
    const ngeNBCUUEW1APID_flag uint = 0x00000002
    const mNGRANUEXnAPID_flag uint = 0x00000004
    const globalNgENBID_flag uint = 0x00000008
    const ext_flag int = 0x00000001
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
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
}//end

type UEIDNGENBDU struct { // [{'type': 'NGENB-CU-UE-W1AP-ID', 'name': 'ng-eNB-CU-UE-W1AP-ID'}, None]
    NgeNBCUUEW1APID NGENBCUUEW1APID
}

func (self * UEIDNGENBDU) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.NgeNBCUUEW1APID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDNGENBDU) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.NgeNBCUUEW1APID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type UEIDENGNB struct { // [{'type': 'ENB-UE-X2AP-ID', 'name': 'm-eNB-UE-X2AP-ID'}, {'type': 'ENB-UE-X2AP-ID-Extension', 'name': 'm-eNB-UE-X2AP-ID-Extension', 'optional': True}, {'type': 'GlobalENB-ID', 'name': 'globalENB-ID'}, {'type': 'GNB-CU-UE-F1AP-ID', 'name': 'gNB-CU-UE-F1AP-ID', 'optional': True}, {'type': 'UEID-GNB-CU-CP-E1AP-ID-List', 'name': 'gNB-CU-CP-UE-E1AP-ID-List', 'optional': True}, {'type': 'RANUEID', 'name': 'ran-UEID', 'optional': True}, None]
    MeNBUEX2APID ENBUEX2APID
    MeNBUEX2APIDExtension *ENBUEX2APIDExtension
    GlobalENBID GlobalENBID
    GNBCUUEF1APID *GNBCUUEF1APID
    GNBCUCPUEE1APIDList *UEIDGNBCUCPE1APIDList
    RanUEID *RANUEID
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
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDENGNB) Pack(stream *Stream) {
    const meNBUEX2APIDExtension_flag uint = 0x00000002
    const gNBCUUEF1APID_flag uint = 0x00000004
    const gNBCUCPUEE1APIDList_flag uint = 0x00000008
    const ranUEID_flag uint = 0x00000010
    const ext_flag int = 0x00000001
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
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

type UEIDENB struct { // [{'type': 'MME-UE-S1AP-ID', 'name': 'mME-UE-S1AP-ID'}, {'type': 'GUMMEI', 'name': 'gUMMEI'}, {'type': 'ENB-UE-X2AP-ID', 'name': 'm-eNB-UE-X2AP-ID', 'optional': True}, {'type': 'ENB-UE-X2AP-ID-Extension', 'name': 'm-eNB-UE-X2AP-ID-Extension', 'optional': True}, {'type': 'GlobalENB-ID', 'name': 'globalENB-ID', 'optional': True}, None]
    MMEUES1APID MMEUES1APID
    GUMMEI GUMMEI
    MeNBUEX2APID *ENBUEX2APID
    MeNBUEX2APIDExtension *ENBUEX2APIDExtension
    GlobalENBID *GlobalENBID
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
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * UEIDENB) Pack(stream *Stream) {
    const meNBUEX2APID_flag uint = 0x00000002
    const meNBUEX2APIDExtension_flag uint = 0x00000004
    const globalENBID_flag uint = 0x00000008
    const ext_flag int = 0x00000001
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
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 4)
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

func (self *GlobalRANNodeID)Unpack(stream *Stream) {
    //coptions := []string{"globalGNB-ID","globalNgENB-ID"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in GlobalRANNodeID\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.GlobalGNBID = &GlobalGNBID{}//cho6
        self.GlobalGNBID.Unpack(stream)
    } else if choice == 1 { //ch2
        self.GlobalNgENBID = &GlobalNgENBID{}//cho6
        self.GlobalNgENBID.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * GlobalRANNodeID) Pack(stream *Stream) {
    if self.GlobalGNBID != nil {
        stream.set_choice(0, 1, 1, 2)
        self.GlobalGNBID.Pack(stream)//2
    } else if self.GlobalNgENBID != nil {
        stream.set_choice(1, 1, 1, 2)
        self.GlobalNgENBID.Pack(stream)//2
    }

}
type GlobalRANNodeID struct { //[{'type': 'GlobalGNB-ID', 'name': 'globalGNB-ID'}, {'type': 'GlobalNgENB-ID', 'name': 'globalNgENB-ID'}, None]
    GlobalGNBID *GlobalGNBID
    GlobalNgENBID *GlobalNgENBID
} // GlobalRANNodeID

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
type NRARFCN_FreqBandListNr struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'FreqBandNrItem'}, 'size': [(1, 'maxnoofNrCellBands')], 'name': 'freqBandListNr'}
    Items []FreqBandNrItem
}
type NRARFCN struct { // [{'type': 'INTEGER', 'restricted-to': [(0, 'maxNRARFCN')], 'name': 'nRARFCN'}, {'type': 'SEQUENCE OF', 'element': {'type': 'FreqBandNrItem'}, 'size': [(1, 'maxnoofNrCellBands')], 'name': 'freqBandListNr'}, None]
    NRARFCN INTEGER
    FreqBandListNr NRARFCN_FreqBandListNr
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
    var Unpack_freqBandListNr = func(stream *Stream, self *NRARFCN_FreqBandListNr){// Seq6 NRARFCN {'type': 'SEQUENCE OF', 'element': {'type': 'FreqBandNrItem'}, 'size': [(1, 'maxnoofNrCellBands')], 'name': 'freqBandListNr'}
        _size := stream.get_listsize(32)
        _size += 1
        self.Items = make([]FreqBandNrItem, _size)//1
        for i := 0; i < _size; i++ {
            self.Items[i].Unpack(stream)
        }
    }

    Unpack_freqBandListNr(stream, &self.FreqBandListNr)// p2
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
    var Pack_freqBandListNr = func(stream *Stream, self NRARFCN_FreqBandListNr) { //seqof 2
        _size := len(self.Items)
        stream.set_listsize(_size-1, 32)
        for _, item := range self.Items {// seqof structure
            item.Pack(stream)
        }
        return

    }

    Pack_freqBandListNr(stream, self.FreqBandListNr) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

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

type RANUEID struct {
  Value HexBytes
}
func (self *RANUEID) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(8)
}
func (self *RANUEID) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 8)
}
type TimeStamp struct {
  Value HexBytes
}
func (self *TimeStamp) Unpack(st *Stream) {
     self.Value = st.parsef_OctString(4)
}
func (self *TimeStamp) Pack(st *Stream) {
    st.formatf_OctString(self.Value, 4)
}
type GranularityPeriod struct {
  Value uint64
}
func (self *GranularityPeriod) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(4294967295, 32, 0, 1)
}
func (self * GranularityPeriod) Pack(st *Stream){
    st.formatf_Integer(self.Value, 4294967295, 32, 0, 1)
}
func (self *MeasurementType)Unpack(stream *Stream) {
    //coptions := []string{"measName","measID"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in MeasurementType\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.MeasName = &MeasurementTypeName{}//cho6
        self.MeasName.Unpack(stream)
    } else if choice == 1 { //ch2
        self.MeasID = &MeasurementTypeID{}//cho6
        self.MeasID.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * MeasurementType) Pack(stream *Stream) {
    if self.MeasName != nil {
        stream.set_choice(0, 1, 1, 2)
        self.MeasName.Pack(stream)//2
    } else if self.MeasID != nil {
        stream.set_choice(1, 1, 1, 2)
        self.MeasID.Pack(stream)//2
    }

}
type MeasurementType struct { //[{'type': 'MeasurementTypeName', 'name': 'measName'}, {'type': 'MeasurementTypeID', 'name': 'measID'}, None]
    MeasName *MeasurementTypeName
    MeasID *MeasurementTypeID
} // MeasurementType

type MeasurementTypeName struct {
  Value string
}
func (self *MeasurementTypeName) Unpack(st *Stream) {
    st.parse_ext()
    _len := st.parse_olen(8)+1
    if _len < 1 || _len > 150 {
        print ("Invalid len in MeasurementTypeName")
        return
    }
    self.Value = st.parsef_PriString(_len)
}

func (self *MeasurementTypeName) Pack(st *Stream) {
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
type MeasurementTypeID struct {
  Value uint64
}
func (self *MeasurementTypeID) Unpack(st *Stream) {
    self.Value = st.parsef_Integer(65536, 17, 1, 1)
}
func (self * MeasurementTypeID) Pack(st *Stream){
    st.formatf_Integer(self.Value, 65536, 17, 1, 1)
}
type MeasurementLabel struct { // [{'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'noLabel', 'optional': True}, {'type': 'PLMNIdentity', 'name': 'plmnID', 'optional': True}, {'type': 'S-NSSAI', 'name': 'sliceID', 'optional': True}, {'type': 'FiveQI', 'name': 'fiveQI', 'optional': True}, {'type': 'QosFlowIdentifier', 'name': 'qFI', 'optional': True}, {'type': 'QCI', 'name': 'qCI', 'optional': True}, {'type': 'QCI', 'name': 'qCImax', 'optional': True}, {'type': 'QCI', 'name': 'qCImin', 'optional': True}, {'type': 'INTEGER', 'restricted-to': [(1, 15), None], 'name': 'aRPmax', 'optional': True}, {'type': 'INTEGER', 'restricted-to': [(1, 15), None], 'name': 'aRPmin', 'optional': True}, {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'bitrateRange', 'optional': True}, {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'layerMU-MIMO', 'optional': True}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'sUM', 'optional': True}, {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'distBinX', 'optional': True}, {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'distBinY', 'optional': True}, {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'distBinZ', 'optional': True}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'preLabelOverride', 'optional': True}, {'type': 'ENUMERATED', 'values': [('start', 0), ('end', 1), None], 'name': 'startEndInd', 'optional': True}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'min', 'optional': True}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'max', 'optional': True}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'avg', 'optional': True}, None]
    NoLabel *ENUMERATED
    PlmnID *PLMNIdentity
    SliceID *SNSSAI
    FiveQI *FiveQI
    QFI *QosFlowIdentifier
    QCI *QCI
    QCImax *QCI
    QCImin *QCI
    ARPmax *INTEGER
    ARPmin *INTEGER
    BitrateRange *INTEGER
    LayerMUMIMO *INTEGER
    SUM *ENUMERATED
    DistBinX *INTEGER
    DistBinY *INTEGER
    DistBinZ *INTEGER
    PreLabelOverride *ENUMERATED
    StartEndInd *ENUMERATED
    Min *ENUMERATED
    Max *ENUMERATED
    Avg *ENUMERATED
}

func (self * MeasurementLabel) Unpack(stream *Stream) {
    noLabel_flag := 0x00000002
    plmnID_flag := 0x00000004
    sliceID_flag := 0x00000008
    fiveQI_flag := 0x00000010
    qFI_flag := 0x00000020
    qCI_flag := 0x00000040
    qCImax_flag := 0x00000080
    qCImin_flag := 0x00000100
    aRPmax_flag := 0x00000200
    aRPmin_flag := 0x00000400
    bitrateRange_flag := 0x00000800
    layerMUMIMO_flag := 0x00001000
    sUM_flag := 0x00002000
    distBinX_flag := 0x00004000
    distBinY_flag := 0x00008000
    distBinZ_flag := 0x00010000
    preLabelOverride_flag := 0x00020000
    startEndInd_flag := 0x00040000
    min_flag := 0x00080000
    max_flag := 0x00100000
    avg_flag := 0x00200000
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(22)
    if (noLabel_flag & _flags) == noLabel_flag { //cond1
        var Unpack_noLabel = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.NoLabel = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'noLabel', 'optional': True}
        Unpack_noLabel(stream, self.NoLabel)// p1 {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'noLabel', 'optional': True}
    }
    if (plmnID_flag & _flags) == plmnID_flag { //cond2
        self.PlmnID = &PLMNIdentity{}//7{'type': 'PLMNIdentity', 'name': 'plmnID', 'optional': True}
        self.PlmnID.Unpack(stream)// p8
    }
    if (sliceID_flag & _flags) == sliceID_flag { //cond2
        self.SliceID = &SNSSAI{}//7{'type': 'S-NSSAI', 'name': 'sliceID', 'optional': True}
        self.SliceID.Unpack(stream)// p8
    }
    if (fiveQI_flag & _flags) == fiveQI_flag { //cond2
        self.FiveQI = &FiveQI{}//7{'type': 'FiveQI', 'name': 'fiveQI', 'optional': True}
        self.FiveQI.Unpack(stream)// p8
    }
    if (qFI_flag & _flags) == qFI_flag { //cond2
        self.QFI = &QosFlowIdentifier{}//7{'type': 'QosFlowIdentifier', 'name': 'qFI', 'optional': True}
        self.QFI.Unpack(stream)// p8
    }
    if (qCI_flag & _flags) == qCI_flag { //cond2
        self.QCI = &QCI{}//7{'type': 'QCI', 'name': 'qCI', 'optional': True}
        self.QCI.Unpack(stream)// p8
    }
    if (qCImax_flag & _flags) == qCImax_flag { //cond2
        self.QCImax = &QCI{}//7{'type': 'QCI', 'name': 'qCImax', 'optional': True}
        self.QCImax.Unpack(stream)// p8
    }
    if (qCImin_flag & _flags) == qCImin_flag { //cond2
        self.QCImin = &QCI{}//7{'type': 'QCI', 'name': 'qCImin', 'optional': True}
        self.QCImin.Unpack(stream)// p8
    }
    if (aRPmax_flag & _flags) == aRPmax_flag { //cond1
        var Unpack_aRPmax = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(15, 5, 1, 1)
        }
        self.ARPmax = &INTEGER{}//6{'type': 'INTEGER', 'restricted-to': [(1, 15), None], 'name': 'aRPmax', 'optional': True}
        Unpack_aRPmax(stream, self.ARPmax)// p1 {'type': 'INTEGER', 'restricted-to': [(1, 15), None], 'name': 'aRPmax', 'optional': True}
    }
    if (aRPmin_flag & _flags) == aRPmin_flag { //cond1
        var Unpack_aRPmin = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(15, 5, 1, 1)
        }
        self.ARPmin = &INTEGER{}//6{'type': 'INTEGER', 'restricted-to': [(1, 15), None], 'name': 'aRPmin', 'optional': True}
        Unpack_aRPmin(stream, self.ARPmin)// p1 {'type': 'INTEGER', 'restricted-to': [(1, 15), None], 'name': 'aRPmin', 'optional': True}
    }
    if (bitrateRange_flag & _flags) == bitrateRange_flag { //cond1
        var Unpack_bitrateRange = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(65535, 17, 1, 1)
        }
        self.BitrateRange = &INTEGER{}//6{'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'bitrateRange', 'optional': True}
        Unpack_bitrateRange(stream, self.BitrateRange)// p1 {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'bitrateRange', 'optional': True}
    }
    if (layerMUMIMO_flag & _flags) == layerMUMIMO_flag { //cond1
        var Unpack_layerMUMIMO = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(65535, 17, 1, 1)
        }
        self.LayerMUMIMO = &INTEGER{}//6{'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'layerMU-MIMO', 'optional': True}
        Unpack_layerMUMIMO(stream, self.LayerMUMIMO)// p1 {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'layerMU-MIMO', 'optional': True}
    }
    if (sUM_flag & _flags) == sUM_flag { //cond1
        var Unpack_sUM = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.SUM = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'sUM', 'optional': True}
        Unpack_sUM(stream, self.SUM)// p1 {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'sUM', 'optional': True}
    }
    if (distBinX_flag & _flags) == distBinX_flag { //cond1
        var Unpack_distBinX = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(65535, 17, 1, 1)
        }
        self.DistBinX = &INTEGER{}//6{'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'distBinX', 'optional': True}
        Unpack_distBinX(stream, self.DistBinX)// p1 {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'distBinX', 'optional': True}
    }
    if (distBinY_flag & _flags) == distBinY_flag { //cond1
        var Unpack_distBinY = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(65535, 17, 1, 1)
        }
        self.DistBinY = &INTEGER{}//6{'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'distBinY', 'optional': True}
        Unpack_distBinY(stream, self.DistBinY)// p1 {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'distBinY', 'optional': True}
    }
    if (distBinZ_flag & _flags) == distBinZ_flag { //cond1
        var Unpack_distBinZ = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(65535, 17, 1, 1)
        }
        self.DistBinZ = &INTEGER{}//6{'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'distBinZ', 'optional': True}
        Unpack_distBinZ(stream, self.DistBinZ)// p1 {'type': 'INTEGER', 'restricted-to': [(1, 65535), None], 'name': 'distBinZ', 'optional': True}
    }
    if (preLabelOverride_flag & _flags) == preLabelOverride_flag { //cond1
        var Unpack_preLabelOverride = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.PreLabelOverride = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'preLabelOverride', 'optional': True}
        Unpack_preLabelOverride(stream, self.PreLabelOverride)// p1 {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'preLabelOverride', 'optional': True}
    }
    if (startEndInd_flag & _flags) == startEndInd_flag { //cond1
        var Unpack_startEndInd = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(2, 2, 1)
        }
        self.StartEndInd = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('start', 0), ('end', 1), None], 'name': 'startEndInd', 'optional': True}
        Unpack_startEndInd(stream, self.StartEndInd)// p1 {'type': 'ENUMERATED', 'values': [('start', 0), ('end', 1), None], 'name': 'startEndInd', 'optional': True}
    }
    if (min_flag & _flags) == min_flag { //cond1
        var Unpack_min = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.Min = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'min', 'optional': True}
        Unpack_min(stream, self.Min)// p1 {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'min', 'optional': True}
    }
    if (max_flag & _flags) == max_flag { //cond1
        var Unpack_max = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.Max = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'max', 'optional': True}
        Unpack_max(stream, self.Max)// p1 {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'max', 'optional': True}
    }
    if (avg_flag & _flags) == avg_flag { //cond1
        var Unpack_avg = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.Avg = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'avg', 'optional': True}
        Unpack_avg(stream, self.Avg)// p1 {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'avg', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MeasurementLabel) Pack(stream *Stream) {
    const noLabel_flag uint = 0x00000002
    const plmnID_flag uint = 0x00000004
    const sliceID_flag uint = 0x00000008
    const fiveQI_flag uint = 0x00000010
    const qFI_flag uint = 0x00000020
    const qCI_flag uint = 0x00000040
    const qCImax_flag uint = 0x00000080
    const qCImin_flag uint = 0x00000100
    const aRPmax_flag uint = 0x00000200
    const aRPmin_flag uint = 0x00000400
    const bitrateRange_flag uint = 0x00000800
    const layerMUMIMO_flag uint = 0x00001000
    const sUM_flag uint = 0x00002000
    const distBinX_flag uint = 0x00004000
    const distBinY_flag uint = 0x00008000
    const distBinZ_flag uint = 0x00010000
    const preLabelOverride_flag uint = 0x00020000
    const startEndInd_flag uint = 0x00040000
    const min_flag uint = 0x00080000
    const max_flag uint = 0x00100000
    const avg_flag uint = 0x00200000
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(22)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.NoLabel != nil { //YY
        _flags |= noLabel_flag
        var Pack_noLabel = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_noLabel(stream, *self.NoLabel) //f1
    }//end of optional
    if self.PlmnID != nil { 
        _flags |= plmnID_flag
        self.PlmnID.Pack(stream)
    }//end of optional
    if self.SliceID != nil { 
        _flags |= sliceID_flag
        self.SliceID.Pack(stream)
    }//end of optional
    if self.FiveQI != nil { 
        _flags |= fiveQI_flag
        self.FiveQI.Pack(stream)
    }//end of optional
    if self.QFI != nil { 
        _flags |= qFI_flag
        self.QFI.Pack(stream)
    }//end of optional
    if self.QCI != nil { 
        _flags |= qCI_flag
        self.QCI.Pack(stream)
    }//end of optional
    if self.QCImax != nil { 
        _flags |= qCImax_flag
        self.QCImax.Pack(stream)
    }//end of optional
    if self.QCImin != nil { 
        _flags |= qCImin_flag
        self.QCImin.Pack(stream)
    }//end of optional
    if self.ARPmax != nil { //YY
        _flags |= aRPmax_flag
        var Pack_aRPmax = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 15, 5, 1, 1)
        }
        Pack_aRPmax(stream, *self.ARPmax) //f1
    }//end of optional
    if self.ARPmin != nil { //YY
        _flags |= aRPmin_flag
        var Pack_aRPmin = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 15, 5, 1, 1)
        }
        Pack_aRPmin(stream, *self.ARPmin) //f1
    }//end of optional
    if self.BitrateRange != nil { //YY
        _flags |= bitrateRange_flag
        var Pack_bitrateRange = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 65535, 17, 1, 1)
        }
        Pack_bitrateRange(stream, *self.BitrateRange) //f1
    }//end of optional
    if self.LayerMUMIMO != nil { //YY
        _flags |= layerMUMIMO_flag
        var Pack_layerMUMIMO = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 65535, 17, 1, 1)
        }
        Pack_layerMUMIMO(stream, *self.LayerMUMIMO) //f1
    }//end of optional
    if self.SUM != nil { //YY
        _flags |= sUM_flag
        var Pack_sUM = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_sUM(stream, *self.SUM) //f1
    }//end of optional
    if self.DistBinX != nil { //YY
        _flags |= distBinX_flag
        var Pack_distBinX = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 65535, 17, 1, 1)
        }
        Pack_distBinX(stream, *self.DistBinX) //f1
    }//end of optional
    if self.DistBinY != nil { //YY
        _flags |= distBinY_flag
        var Pack_distBinY = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 65535, 17, 1, 1)
        }
        Pack_distBinY(stream, *self.DistBinY) //f1
    }//end of optional
    if self.DistBinZ != nil { //YY
        _flags |= distBinZ_flag
        var Pack_distBinZ = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 65535, 17, 1, 1)
        }
        Pack_distBinZ(stream, *self.DistBinZ) //f1
    }//end of optional
    if self.PreLabelOverride != nil { //YY
        _flags |= preLabelOverride_flag
        var Pack_preLabelOverride = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_preLabelOverride(stream, *self.PreLabelOverride) //f1
    }//end of optional
    if self.StartEndInd != nil { //YY
        _flags |= startEndInd_flag
        var Pack_startEndInd = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 2, 2, 1)
        }
        Pack_startEndInd(stream, *self.StartEndInd) //f1
    }//end of optional
    if self.Min != nil { //YY
        _flags |= min_flag
        var Pack_min = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_min(stream, *self.Min) //f1
    }//end of optional
    if self.Max != nil { //YY
        _flags |= max_flag
        var Pack_max = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_max(stream, *self.Max) //f1
    }//end of optional
    if self.Avg != nil { //YY
        _flags |= avg_flag
        var Pack_avg = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_avg(stream, *self.Avg) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 22)
}//end

type TestCondInfo struct { // [{'type': 'TestCond-Type', 'name': 'testType'}, {'type': 'TestCond-Expression', 'name': 'testExpr'}, {'type': 'TestCond-Value', 'name': 'testValue'}, None]
    TestType TestCondType
    TestExpr TestCondExpression
    TestValue TestCondValue
}

func (self * TestCondInfo) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.TestType.Unpack(stream)// p8
    self.TestExpr.Unpack(stream)// p8
    self.TestValue.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * TestCondInfo) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.TestType.Pack(stream)
    self.TestExpr.Pack(stream)
    self.TestValue.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *TestCondType)Unpack(stream *Stream) {
    //coptions := []string{"gBR","aMBR","isStat","isCatM","rSRP","rSRQ","Unknown","Unknown"}
    choice := stream.get_choice(3, 1, 6)
    choice_len := 0
    choice_loc := 0
    if choice >= 6 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in TestCondType\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_gBR = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.GBR = &ENUMERATED{}//cho5
        Unpack_gBR(stream, self.GBR);
    } else if choice == 1 { //ch2
        var Unpack_aMBR = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.AMBR = &ENUMERATED{}//cho5
        Unpack_aMBR(stream, self.AMBR);
    } else if choice == 2 { //ch2
        var Unpack_isStat = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.IsStat = &ENUMERATED{}//cho5
        Unpack_isStat(stream, self.IsStat);
    } else if choice == 3 { //ch2
        var Unpack_isCatM = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.IsCatM = &ENUMERATED{}//cho5
        Unpack_isCatM(stream, self.IsCatM);
    } else if choice == 4 { //ch2
        var Unpack_rSRP = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.RSRP = &ENUMERATED{}//cho5
        Unpack_rSRP(stream, self.RSRP);
    } else if choice == 5 { //ch2
        var Unpack_rSRQ = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.RSRQ = &ENUMERATED{}//cho5
        Unpack_rSRQ(stream, self.RSRQ);
    }//end of if else

    if choice >= 6 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * TestCondType) Pack(stream *Stream) {
    if self.GBR != nil {
        stream.set_choice(0, 3, 1, 6)
        var Pack_gBR = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_gBR(stream, *self.GBR)//3
    } else if self.AMBR != nil {
        stream.set_choice(1, 3, 1, 6)
        var Pack_aMBR = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_aMBR(stream, *self.AMBR)//3
    } else if self.IsStat != nil {
        stream.set_choice(2, 3, 1, 6)
        var Pack_isStat = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_isStat(stream, *self.IsStat)//3
    } else if self.IsCatM != nil {
        stream.set_choice(3, 3, 1, 6)
        var Pack_isCatM = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_isCatM(stream, *self.IsCatM)//3
    } else if self.RSRP != nil {
        stream.set_choice(4, 3, 1, 6)
        var Pack_rSRP = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_rSRP(stream, *self.RSRP)//3
    } else if self.RSRQ != nil {
        stream.set_choice(5, 3, 1, 6)
        var Pack_rSRQ = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_rSRQ(stream, *self.RSRQ)//3
    }

}
type TestCondType struct { //[{'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'gBR'}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'aMBR'}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'isStat'}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'isCatM'}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'rSRP'}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'rSRQ'}, None]
    GBR *ENUMERATED
    AMBR *ENUMERATED
    IsStat *ENUMERATED
    IsCatM *ENUMERATED
    RSRP *ENUMERATED
    RSRQ *ENUMERATED
} // TestCondType

type TestCondExpression struct {
  Value int
}
const (
    TestCondExpressionequal = 0
    TestCondExpressiongreaterthan = 1
    TestCondExpressionlessthan = 2
    TestCondExpressioncontains = 3
    TestCondExpressionpresent = 4

    /* Extensions */
)
func (self *TestCondExpression) Unpack(st *Stream) {
    self.Value = st.parsef_Enumerated(4, 5, 1)
}
func (self *TestCondExpression) Pack(st *Stream) {
    st.formatf_Enumerated(self.Value, 4, 5, 1)
}
func (self *TestCondValue)Unpack(stream *Stream) {
    //coptions := []string{"valueInt","valueEnum","valueBool","valueBitS","valueOctS","valuePrtS","Unknown","Unknown"}
    choice := stream.get_choice(3, 1, 6)
    choice_len := 0
    choice_loc := 0
    if choice >= 6 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in TestCondValue\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_valueInt = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(0, 0, 0, 0)
        }
        self.ValueInt = &INTEGER{}//cho5
        Unpack_valueInt(stream, self.ValueInt);
    } else if choice == 1 { //ch2
        var Unpack_valueEnum = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(0, 0, 0, 0)
        }
        self.ValueEnum = &INTEGER{}//cho5
        Unpack_valueEnum(stream, self.ValueEnum);
    } else if choice == 2 { //ch2
        var  Unpack_valueBool = func (st *Stream, self *BOOLEAN) {
            self.Value = st.parsef_bool()
            }
        self.ValueBool = &BOOLEAN{}//cho5
        Unpack_valueBool(stream, self.ValueBool);
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
        var Unpack_valuePrtS = func(st *Stream, self *PrintableString) {
            _len := st.parse_len(0)
            self.Value = st.parsef_PriString(_len)
        }
        self.ValuePrtS = &PrintableString{}//cho5
        Unpack_valuePrtS(stream, self.ValuePrtS);
    }//end of if else

    if choice >= 6 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * TestCondValue) Pack(stream *Stream) {
    if self.ValueInt != nil {
        stream.set_choice(0, 3, 1, 6)
        var Pack_valueInt = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 0, 0, 0, 0)
        }
        Pack_valueInt(stream, *self.ValueInt)//3
    } else if self.ValueEnum != nil {
        stream.set_choice(1, 3, 1, 6)
        var Pack_valueEnum = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 0, 0, 0, 0)
        }
        Pack_valueEnum(stream, *self.ValueEnum)//3
    } else if self.ValueBool != nil {
        stream.set_choice(2, 3, 1, 6)
        var Pack_valueBool = func(st *Stream, self BOOLEAN) {
            st.formatf_bool(self.Value);
            }
        Pack_valueBool(stream, *self.ValueBool)//3
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
    } else if self.ValuePrtS != nil {
        stream.set_choice(5, 3, 1, 6)
        var Pack_valuePrtS = func(st *Stream, self PrintableString) {
            st.format_len((len(self.Value)), 0)
            st.formatf_PriString(self.Value, len(self.Value))
        }
        Pack_valuePrtS(stream, *self.ValuePrtS)//3
    }

}
type TestCondValue struct { //[{'type': 'INTEGER', 'name': 'valueInt'}, {'type': 'INTEGER', 'name': 'valueEnum'}, {'type': 'BOOLEAN', 'name': 'valueBool'}, {'type': 'BIT STRING', 'name': 'valueBitS'}, {'type': 'OCTET STRING', 'name': 'valueOctS'}, {'type': 'PrintableString', 'name': 'valuePrtS'}, None]
    ValueInt *INTEGER
    ValueEnum *INTEGER
    ValueBool *BOOLEAN
    ValueBitS *BITSTRING
    ValueOctS *OCTETSTRING
    ValuePrtS *PrintableString
} // TestCondValue

func (self *MeasurementInfoList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65535)
    _size += 1
    self.Items = make([]MeasurementInfoItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MeasurementInfoList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65535)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MeasurementInfoList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MeasurementInfoItem'}, 'size': [(1, 'maxnoofMeasurementInfo')]}
    Items []MeasurementInfoItem
}

type MeasurementInfoItem struct { // [{'type': 'MeasurementType', 'name': 'measType'}, {'type': 'LabelInfoList', 'name': 'labelInfoList'}, None]
    MeasType MeasurementType
    LabelInfoList LabelInfoList
}

func (self * MeasurementInfoItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.MeasType.Unpack(stream)// p8
    self.LabelInfoList.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MeasurementInfoItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeasType.Pack(stream)
    self.LabelInfoList.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *LabelInfoList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2147483647)
    self.Items = make([]LabelInfoItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *LabelInfoList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size, 2147483647)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type LabelInfoList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'LabelInfoItem'}, 'size': [(1, 'maxnoofLabelInfo')]}
    Items []LabelInfoItem
}

type LabelInfoItem struct { // [{'type': 'MeasurementLabel', 'name': 'measLabel'}, None]
    MeasLabel MeasurementLabel
}

func (self * LabelInfoItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.MeasLabel.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * LabelInfoItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeasLabel.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *MeasurementData) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65535)
    _size += 1
    self.Items = make([]MeasurementDataItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MeasurementData) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65535)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MeasurementData struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MeasurementDataItem'}, 'size': [(1, 'maxnoofMeasurementRecord')]}
    Items []MeasurementDataItem
}

type MeasurementDataItem struct { // [{'type': 'MeasurementRecord', 'name': 'measRecord'}, {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'incompleteFlag', 'optional': True}, None]
    MeasRecord MeasurementRecord
    IncompleteFlag *ENUMERATED
}

func (self * MeasurementDataItem) Unpack(stream *Stream) {
    incompleteFlag_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.MeasRecord.Unpack(stream)// p8
    if (incompleteFlag_flag & _flags) == incompleteFlag_flag { //cond1
        var Unpack_incompleteFlag = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(1, 1, 1)
        }
        self.IncompleteFlag = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'incompleteFlag', 'optional': True}
        Unpack_incompleteFlag(stream, self.IncompleteFlag)// p1 {'type': 'ENUMERATED', 'values': [('true', 0), None], 'name': 'incompleteFlag', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MeasurementDataItem) Pack(stream *Stream) {
    const incompleteFlag_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeasRecord.Pack(stream)
    if self.IncompleteFlag != nil { //YY
        _flags |= incompleteFlag_flag
        var Pack_incompleteFlag = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 1, 1, 1)
        }
        Pack_incompleteFlag(stream, *self.IncompleteFlag) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *MeasurementRecord) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(2147483647)
    self.Items = make([]MeasurementRecordItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MeasurementRecord) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size, 2147483647)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MeasurementRecord struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MeasurementRecordItem'}, 'size': [(1, 'maxnoofMeasurementValue')]}
    Items []MeasurementRecordItem
}

func (self *MeasurementRecordItem)Unpack(stream *Stream) {
    //coptions := []string{"integer","real","noValue","Unknown"}
    choice := stream.get_choice(2, 1, 3)
    choice_len := 0
    choice_loc := 0
    if choice >= 3 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in MeasurementRecordItem\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        var Unpack_integer = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(4294967296, 32, 0, 0)
        }
        self.Integer = &INTEGER{}//cho5
        Unpack_integer(stream, self.Integer);
    } else if choice == 1 { //ch2
        var Unpack_real = func (st *Stream, self *REAL) {
            self.Value = st.parsef_Real(0)
        }
        self.Real = &REAL{}//cho5
        Unpack_real(stream, self.Real);
    } else if choice == 2 { //ch2
        var Unpack_noValue = func(st *Stream, self *NULL){
            st.parsef_Null()
        }
        self.NoValue = &NULL{}//cho5
        Unpack_noValue(stream, self.NoValue);
    }//end of if else

    if choice >= 3 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * MeasurementRecordItem) Pack(stream *Stream) {
    if self.Integer != nil {
        stream.set_choice(0, 2, 1, 3)
        var Pack_integer = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 4294967296, 32, 0, 0)
        }
        Pack_integer(stream, *self.Integer)//3
    } else if self.Real != nil {
        stream.set_choice(1, 2, 1, 3)
        var Pack_real = func (st *Stream, self REAL) {
            st.formatf_Real(self.Value, 0)

}
        Pack_real(stream, *self.Real)//3
    } else if self.NoValue != nil {
        stream.set_choice(2, 2, 1, 3)
        var Pack_noValue = func(st *Stream, self NULL){
            st.formatf_Null()
        }
        Pack_noValue(stream, *self.NoValue)//3
    }

}
type MeasurementRecordItem struct { //[{'type': 'INTEGER', 'restricted-to': [(0, 4294967295)], 'name': 'integer'}, {'type': 'REAL', 'name': 'real'}, {'type': 'NULL', 'name': 'noValue'}, None]
    Integer *INTEGER
    Real *REAL
    NoValue *NULL
} // MeasurementRecordItem

func (self *MeasurementInfoActionList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65535)
    _size += 1
    self.Items = make([]MeasurementInfoActionItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MeasurementInfoActionList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65535)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MeasurementInfoActionList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MeasurementInfo-Action-Item'}, 'size': [(1, 'maxnoofMeasurementInfo')]}
    Items []MeasurementInfoActionItem
}

type MeasurementInfoActionItem struct { // [{'type': 'MeasurementTypeName', 'name': 'measName'}, {'type': 'MeasurementTypeID', 'name': 'measID', 'optional': True}, None]
    MeasName MeasurementTypeName
    MeasID *MeasurementTypeID
}

func (self * MeasurementInfoActionItem) Unpack(stream *Stream) {
    measID_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.MeasName.Unpack(stream)// p8
    if (measID_flag & _flags) == measID_flag { //cond2
        self.MeasID = &MeasurementTypeID{}//7{'type': 'MeasurementTypeID', 'name': 'measID', 'optional': True}
        self.MeasID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MeasurementInfoActionItem) Pack(stream *Stream) {
    const measID_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeasName.Pack(stream)
    if self.MeasID != nil { 
        _flags |= measID_flag
        self.MeasID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *MeasurementCondList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65535)
    _size += 1
    self.Items = make([]MeasurementCondItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MeasurementCondList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65535)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MeasurementCondList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MeasurementCondItem'}, 'size': [(1, 'maxnoofMeasurementInfo')]}
    Items []MeasurementCondItem
}

type MeasurementCondItem struct { // [{'type': 'MeasurementType', 'name': 'measType'}, {'type': 'MatchingCondList', 'name': 'matchingCond'}, None]
    MeasType MeasurementType
    MatchingCond MatchingCondList
}

func (self * MeasurementCondItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.MeasType.Unpack(stream)// p8
    self.MatchingCond.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MeasurementCondItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeasType.Pack(stream)
    self.MatchingCond.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

func (self *MeasurementCondUEidList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65535)
    _size += 1
    self.Items = make([]MeasurementCondUEidItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MeasurementCondUEidList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65535)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MeasurementCondUEidList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MeasurementCondUEidItem'}, 'size': [(1, 'maxnoofMeasurementInfo')]}
    Items []MeasurementCondUEidItem
}

type MeasurementCondUEidItem struct { // [{'type': 'MeasurementType', 'name': 'measType'}, {'type': 'MatchingCondList', 'name': 'matchingCond'}, {'type': 'MatchingUEidList', 'name': 'matchingUEidList', 'optional': True}, None]
    MeasType MeasurementType
    MatchingCond MatchingCondList
    MatchingUEidList *MatchingUEidList
}

func (self * MeasurementCondUEidItem) Unpack(stream *Stream) {
    matchingUEidList_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.MeasType.Unpack(stream)// p8
    self.MatchingCond.Unpack(stream)// p8
    if (matchingUEidList_flag & _flags) == matchingUEidList_flag { //cond2
        self.MatchingUEidList = &MatchingUEidList{}//7{'type': 'MatchingUEidList', 'name': 'matchingUEidList', 'optional': True}
        self.MatchingUEidList.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MeasurementCondUEidItem) Pack(stream *Stream) {
    const matchingUEidList_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeasType.Pack(stream)
    self.MatchingCond.Pack(stream)
    if self.MatchingUEidList != nil { 
        _flags |= matchingUEidList_flag
        self.MatchingUEidList.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

func (self *MatchingCondList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(32768)
    _size += 1
    self.Items = make([]MatchingCondItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MatchingCondList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 32768)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MatchingCondList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MatchingCondItem'}, 'size': [(1, 'maxnoofConditionInfo')]}
    Items []MatchingCondItem
}

func (self *MatchingCondItem)Unpack(stream *Stream) {
    //coptions := []string{"measLabel","testCondInfo"}
    choice := stream.get_choice(1, 1, 2)
    choice_len := 0
    choice_loc := 0
    if choice >= 2 {
        choice_len = stream.parse_len(0)
        choice_loc = stream.get_location()
        log.Info("Extension choice option [%d] len %d in MatchingCondItem\n", choice, choice_len)
    }
    if choice == 0 { //ch1
        self.MeasLabel = &MeasurementLabel{}//cho6
        self.MeasLabel.Unpack(stream)
    } else if choice == 1 { //ch2
        self.TestCondInfo = &TestCondInfo{}//cho6
        self.TestCondInfo.Unpack(stream)
    }//end of if else

    if choice >= 2 {
        stream.set_location(choice_loc, choice_len)
    }
}
func (self * MatchingCondItem) Pack(stream *Stream) {
    if self.MeasLabel != nil {
        stream.set_choice(0, 1, 1, 2)
        self.MeasLabel.Pack(stream)//2
    } else if self.TestCondInfo != nil {
        stream.set_choice(1, 1, 1, 2)
        self.TestCondInfo.Pack(stream)//2
    }

}
type MatchingCondItem struct { //[{'type': 'MeasurementLabel', 'name': 'measLabel'}, {'type': 'TestCondInfo', 'name': 'testCondInfo'}, None]
    MeasLabel *MeasurementLabel
    TestCondInfo *TestCondInfo
} // MatchingCondItem

func (self *MatchingUEidList) Unpack(stream *Stream){ //seq5
    _size := stream.get_listsize(65535)
    _size += 1
    self.Items = make([]MatchingUEidItem, _size)//1
    for i := 0; i < _size; i++ {
        self.Items[i].Unpack(stream)
    }
}


func (self *MatchingUEidList) Pack(stream *Stream) { //seqof 2
    _size := len(self.Items)
    stream.set_listsize(_size-1, 65535)
    for _, item := range self.Items {// seqof structure
        item.Pack(stream)
    }
    return

}


type MatchingUEidList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'MatchingUEidItem'}, 'size': [(1, 'maxnoofUEID')]}
    Items []MatchingUEidItem
}

type MatchingUEidItem struct { // [{'type': 'UEID', 'name': 'ueID'}, None]
    UeID UEID
}

func (self * MatchingUEidItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.UeID.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * MatchingUEidItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UeID.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMKPMEventTriggerDefinition_EventDefinitionformats struct { //[{'type': 'E2SM-KPM-EventTriggerDefinition-Format1', 'name': 'eventDefinition-Format1'}, None]
    EventDefinitionFormat1 *E2SMKPMEventTriggerDefinitionFormat1
} // E2SMKPMEventTriggerDefinition_EventDefinitionformats

type E2SMKPMEventTriggerDefinition struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-KPM-EventTriggerDefinition-Format1', 'name': 'eventDefinition-Format1'}, None], 'name': 'eventDefinition-formats'}, None]
    EventDefinitionformats E2SMKPMEventTriggerDefinition_EventDefinitionformats
}

func (self * E2SMKPMEventTriggerDefinition) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_eventDefinitionformats = func(stream *Stream, self *E2SMKPMEventTriggerDefinition_EventDefinitionformats) {
        //coptions := []string{"eventDefinition-Format1"}
        choice := stream.get_choice(0, 1, 1)
        choice_len := 0
        choice_loc := 0
        if choice >= 1 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMKPMEventTriggerDefinition_EventDefinitionformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.EventDefinitionFormat1 = &E2SMKPMEventTriggerDefinitionFormat1{}//cho6
            self.EventDefinitionFormat1.Unpack(stream)
        }//end of if else

        if choice >= 1 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_eventDefinitionformats(stream, &self.EventDefinitionformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMEventTriggerDefinition) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_eventDefinitionformats = func(stream *Stream, self E2SMKPMEventTriggerDefinition_EventDefinitionformats) {
        if self.EventDefinitionFormat1 != nil {
            stream.set_choice(0, 0, 1, 1)
            self.EventDefinitionFormat1.Pack(stream)//2
        }

    }
    Pack_eventDefinitionformats(stream, self.EventDefinitionformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMKPMEventTriggerDefinitionFormat1 struct { // [{'type': 'INTEGER', 'restricted-to': [(1, 4294967295)], 'name': 'reportingPeriod'}, None]
    ReportingPeriod INTEGER
}

func (self * E2SMKPMEventTriggerDefinitionFormat1) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_reportingPeriod = func (st *Stream, self *INTEGER) {
        self.Value = st.parsef_Integer(4294967295, 32, 0, 1)
    }
    Unpack_reportingPeriod(stream, &self.ReportingPeriod)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMEventTriggerDefinitionFormat1) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_reportingPeriod = func (st *Stream, self INTEGER){
        st.formatf_Integer(self.Value, 4294967295, 32, 0, 1)
    }
    Pack_reportingPeriod(stream, self.ReportingPeriod) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMKPMActionDefinition_ActionDefinitionformats struct { //[{'type': 'E2SM-KPM-ActionDefinition-Format1', 'name': 'actionDefinition-Format1'}, {'type': 'E2SM-KPM-ActionDefinition-Format2', 'name': 'actionDefinition-Format2'}, {'type': 'E2SM-KPM-ActionDefinition-Format3', 'name': 'actionDefinition-Format3'}, None]
    ActionDefinitionFormat1 *E2SMKPMActionDefinitionFormat1
    ActionDefinitionFormat2 *E2SMKPMActionDefinitionFormat2
    ActionDefinitionFormat3 *E2SMKPMActionDefinitionFormat3
} // E2SMKPMActionDefinition_ActionDefinitionformats

type E2SMKPMActionDefinition struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-Style-Type'}, {'type': 'CHOICE', 'members': [{'type': 'E2SM-KPM-ActionDefinition-Format1', 'name': 'actionDefinition-Format1'}, {'type': 'E2SM-KPM-ActionDefinition-Format2', 'name': 'actionDefinition-Format2'}, {'type': 'E2SM-KPM-ActionDefinition-Format3', 'name': 'actionDefinition-Format3'}, None], 'name': 'actionDefinition-formats'}, None]
    RicStyleType RICStyleType
    ActionDefinitionformats E2SMKPMActionDefinition_ActionDefinitionformats
}

func (self * E2SMKPMActionDefinition) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicStyleType.Unpack(stream)// p8
    var Unpack_actionDefinitionformats = func(stream *Stream, self *E2SMKPMActionDefinition_ActionDefinitionformats) {
        //coptions := []string{"actionDefinition-Format1","actionDefinition-Format2","actionDefinition-Format3","Unknown"}
        choice := stream.get_choice(2, 1, 3)
        choice_len := 0
        choice_loc := 0
        if choice >= 3 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMKPMActionDefinition_ActionDefinitionformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.ActionDefinitionFormat1 = &E2SMKPMActionDefinitionFormat1{}//cho6
            self.ActionDefinitionFormat1.Unpack(stream)
        } else if choice == 1 { //ch2
            self.ActionDefinitionFormat2 = &E2SMKPMActionDefinitionFormat2{}//cho6
            self.ActionDefinitionFormat2.Unpack(stream)
        } else if choice == 2 { //ch2
            self.ActionDefinitionFormat3 = &E2SMKPMActionDefinitionFormat3{}//cho6
            self.ActionDefinitionFormat3.Unpack(stream)
        }//end of if else

        if choice >= 3 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_actionDefinitionformats(stream, &self.ActionDefinitionformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMActionDefinition) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicStyleType.Pack(stream)
    var Pack_actionDefinitionformats = func(stream *Stream, self E2SMKPMActionDefinition_ActionDefinitionformats) {
        if self.ActionDefinitionFormat1 != nil {
            stream.set_choice(0, 2, 1, 3)
            self.ActionDefinitionFormat1.Pack(stream)//2
        } else if self.ActionDefinitionFormat2 != nil {
            stream.set_choice(1, 2, 1, 3)
            self.ActionDefinitionFormat2.Pack(stream)//2
        } else if self.ActionDefinitionFormat3 != nil {
            stream.set_choice(2, 2, 1, 3)
            self.ActionDefinitionFormat3.Pack(stream)//2
        }

    }
    Pack_actionDefinitionformats(stream, self.ActionDefinitionformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMKPMActionDefinitionFormat1 struct { // [{'type': 'MeasurementInfoList', 'name': 'measInfoList'}, {'type': 'GranularityPeriod', 'name': 'granulPeriod'}, {'type': 'CGI', 'name': 'cellGlobalID', 'optional': True}, None]
    MeasInfoList MeasurementInfoList
    GranulPeriod GranularityPeriod
    CellGlobalID *CGI
}

func (self * E2SMKPMActionDefinitionFormat1) Unpack(stream *Stream) {
    cellGlobalID_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.MeasInfoList.Unpack(stream)// p8
    self.GranulPeriod.Unpack(stream)// p8
    if (cellGlobalID_flag & _flags) == cellGlobalID_flag { //cond2
        self.CellGlobalID = &CGI{}//7{'type': 'CGI', 'name': 'cellGlobalID', 'optional': True}
        self.CellGlobalID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMActionDefinitionFormat1) Pack(stream *Stream) {
    const cellGlobalID_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeasInfoList.Pack(stream)
    self.GranulPeriod.Pack(stream)
    if self.CellGlobalID != nil { 
        _flags |= cellGlobalID_flag
        self.CellGlobalID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMKPMActionDefinitionFormat2 struct { // [{'type': 'UEID', 'name': 'ueID'}, {'type': 'E2SM-KPM-ActionDefinition-Format1', 'name': 'subscriptInfo'}, None]
    UeID UEID
    SubscriptInfo E2SMKPMActionDefinitionFormat1
}

func (self * E2SMKPMActionDefinitionFormat2) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.UeID.Unpack(stream)// p8
    self.SubscriptInfo.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMActionDefinitionFormat2) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.UeID.Pack(stream)
    self.SubscriptInfo.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMKPMActionDefinitionFormat3 struct { // [{'type': 'MeasurementCondList', 'name': 'measCondList'}, {'type': 'GranularityPeriod', 'name': 'granulPeriod'}, {'type': 'CGI', 'name': 'cellGlobalID', 'optional': True}, None]
    MeasCondList MeasurementCondList
    GranulPeriod GranularityPeriod
    CellGlobalID *CGI
}

func (self * E2SMKPMActionDefinitionFormat3) Unpack(stream *Stream) {
    cellGlobalID_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.MeasCondList.Unpack(stream)// p8
    self.GranulPeriod.Unpack(stream)// p8
    if (cellGlobalID_flag & _flags) == cellGlobalID_flag { //cond2
        self.CellGlobalID = &CGI{}//7{'type': 'CGI', 'name': 'cellGlobalID', 'optional': True}
        self.CellGlobalID.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMActionDefinitionFormat3) Pack(stream *Stream) {
    const cellGlobalID_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeasCondList.Pack(stream)
    self.GranulPeriod.Pack(stream)
    if self.CellGlobalID != nil { 
        _flags |= cellGlobalID_flag
        self.CellGlobalID.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMKPMIndicationHeader_IndicationHeaderformats struct { //[{'type': 'E2SM-KPM-IndicationHeader-Format1', 'name': 'indicationHeader-Format1'}, None]
    IndicationHeaderFormat1 *E2SMKPMIndicationHeaderFormat1
} // E2SMKPMIndicationHeader_IndicationHeaderformats

type E2SMKPMIndicationHeader struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-KPM-IndicationHeader-Format1', 'name': 'indicationHeader-Format1'}, None], 'name': 'indicationHeader-formats'}, None]
    IndicationHeaderformats E2SMKPMIndicationHeader_IndicationHeaderformats
}

func (self * E2SMKPMIndicationHeader) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_indicationHeaderformats = func(stream *Stream, self *E2SMKPMIndicationHeader_IndicationHeaderformats) {
        //coptions := []string{"indicationHeader-Format1"}
        choice := stream.get_choice(0, 1, 1)
        choice_len := 0
        choice_loc := 0
        if choice >= 1 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMKPMIndicationHeader_IndicationHeaderformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.IndicationHeaderFormat1 = &E2SMKPMIndicationHeaderFormat1{}//cho6
            self.IndicationHeaderFormat1.Unpack(stream)
        }//end of if else

        if choice >= 1 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_indicationHeaderformats(stream, &self.IndicationHeaderformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMIndicationHeader) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_indicationHeaderformats = func(stream *Stream, self E2SMKPMIndicationHeader_IndicationHeaderformats) {
        if self.IndicationHeaderFormat1 != nil {
            stream.set_choice(0, 0, 1, 1)
            self.IndicationHeaderFormat1.Pack(stream)//2
        }

    }
    Pack_indicationHeaderformats(stream, self.IndicationHeaderformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMKPMIndicationHeaderFormat1 struct { // [{'type': 'TimeStamp', 'name': 'colletStartTime'}, {'type': 'PrintableString', 'size': [(0, 15), None], 'name': 'fileFormatversion', 'optional': True}, {'type': 'PrintableString', 'size': [(0, 400), None], 'name': 'senderName', 'optional': True}, {'type': 'PrintableString', 'size': [(0, 8), None], 'name': 'senderType', 'optional': True}, {'type': 'PrintableString', 'size': [(0, 32), None], 'name': 'vendorName', 'optional': True}, None]
    ColletStartTime TimeStamp
    FileFormatversion *PrintableString
    SenderName *PrintableString
    SenderType *PrintableString
    VendorName *PrintableString
}

func (self * E2SMKPMIndicationHeaderFormat1) Unpack(stream *Stream) {
    fileFormatversion_flag := 0x00000002
    senderName_flag := 0x00000004
    senderType_flag := 0x00000008
    vendorName_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    self.ColletStartTime.Unpack(stream)// p8
    if (fileFormatversion_flag & _flags) == fileFormatversion_flag { //cond1
        var Unpack_fileFormatversion = func(st *Stream, self *PrintableString) {
            st.parse_ext()
            _len := st.parse_olen(4)
            if _len < 0 || _len > 15 {
                log.Error ("Invalid len in fileFormatversion")
                return
            }
            self.Value = st.parsef_PriString(_len)
        }
        self.FileFormatversion = &PrintableString{}//6{'type': 'PrintableString', 'size': [(0, 15), None], 'name': 'fileFormatversion', 'optional': True}
        Unpack_fileFormatversion(stream, self.FileFormatversion)// p1 {'type': 'PrintableString', 'size': [(0, 15), None], 'name': 'fileFormatversion', 'optional': True}
    }
    if (senderName_flag & _flags) == senderName_flag { //cond1
        var Unpack_senderName = func(st *Stream, self *PrintableString) {
            st.parse_ext()
            _len := st.parse_olen(9)
            if _len < 0 || _len > 400 {
                log.Error ("Invalid len in senderName")
                return
            }
            self.Value = st.parsef_PriString(_len)
        }
        self.SenderName = &PrintableString{}//6{'type': 'PrintableString', 'size': [(0, 400), None], 'name': 'senderName', 'optional': True}
        Unpack_senderName(stream, self.SenderName)// p1 {'type': 'PrintableString', 'size': [(0, 400), None], 'name': 'senderName', 'optional': True}
    }
    if (senderType_flag & _flags) == senderType_flag { //cond1
        var Unpack_senderType = func(st *Stream, self *PrintableString) {
            st.parse_ext()
            _len := st.parse_olen(4)
            if _len < 0 || _len > 8 {
                log.Error ("Invalid len in senderType")
                return
            }
            self.Value = st.parsef_PriString(_len)
        }
        self.SenderType = &PrintableString{}//6{'type': 'PrintableString', 'size': [(0, 8), None], 'name': 'senderType', 'optional': True}
        Unpack_senderType(stream, self.SenderType)// p1 {'type': 'PrintableString', 'size': [(0, 8), None], 'name': 'senderType', 'optional': True}
    }
    if (vendorName_flag & _flags) == vendorName_flag { //cond1
        var Unpack_vendorName = func(st *Stream, self *PrintableString) {
            st.parse_ext()
            _len := st.parse_olen(6)
            if _len < 0 || _len > 32 {
                log.Error ("Invalid len in vendorName")
                return
            }
            self.Value = st.parsef_PriString(_len)
        }
        self.VendorName = &PrintableString{}//6{'type': 'PrintableString', 'size': [(0, 32), None], 'name': 'vendorName', 'optional': True}
        Unpack_vendorName(stream, self.VendorName)// p1 {'type': 'PrintableString', 'size': [(0, 32), None], 'name': 'vendorName', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMIndicationHeaderFormat1) Pack(stream *Stream) {
    const fileFormatversion_flag uint = 0x00000002
    const senderName_flag uint = 0x00000004
    const senderType_flag uint = 0x00000008
    const vendorName_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.ColletStartTime.Pack(stream)
    if self.FileFormatversion != nil { //YY
        _flags |= fileFormatversion_flag
        var Pack_fileFormatversion = func(st *Stream, self PrintableString) {
            _eflag := 0
            if len(self.Value) > 15 {
               _eflag = 1
            }
            st.format_ext(_eflag)
            if len(self.Value) < 0 || len(self.Value) > 15 {
                return;
            }
            st.format_olen((len(self.Value)), 4)
            st.formatf_PriString(self.Value, len(self.Value))
        }
        Pack_fileFormatversion(stream, *self.FileFormatversion) //f1
    }//end of optional
    if self.SenderName != nil { //YY
        _flags |= senderName_flag
        var Pack_senderName = func(st *Stream, self PrintableString) {
            _eflag := 0
            if len(self.Value) > 400 {
               _eflag = 1
            }
            st.format_ext(_eflag)
            if len(self.Value) < 0 || len(self.Value) > 400 {
                return;
            }
            st.format_olen((len(self.Value)), 9)
            st.formatf_PriString(self.Value, len(self.Value))
        }
        Pack_senderName(stream, *self.SenderName) //f1
    }//end of optional
    if self.SenderType != nil { //YY
        _flags |= senderType_flag
        var Pack_senderType = func(st *Stream, self PrintableString) {
            _eflag := 0
            if len(self.Value) > 8 {
               _eflag = 1
            }
            st.format_ext(_eflag)
            if len(self.Value) < 0 || len(self.Value) > 8 {
                return;
            }
            st.format_olen((len(self.Value)), 4)
            st.formatf_PriString(self.Value, len(self.Value))
        }
        Pack_senderType(stream, *self.SenderType) //f1
    }//end of optional
    if self.VendorName != nil { //YY
        _flags |= vendorName_flag
        var Pack_vendorName = func(st *Stream, self PrintableString) {
            _eflag := 0
            if len(self.Value) > 32 {
               _eflag = 1
            }
            st.format_ext(_eflag)
            if len(self.Value) < 0 || len(self.Value) > 32 {
                return;
            }
            st.format_olen((len(self.Value)), 6)
            st.formatf_PriString(self.Value, len(self.Value))
        }
        Pack_vendorName(stream, *self.VendorName) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

type E2SMKPMIndicationMessage_IndicationMessageformats struct { //[{'type': 'E2SM-KPM-IndicationMessage-Format1', 'name': 'indicationMessage-Format1'}, {'type': 'E2SM-KPM-IndicationMessage-Format2', 'name': 'indicationMessage-Format2'}, None]
    IndicationMessageFormat1 *E2SMKPMIndicationMessageFormat1
    IndicationMessageFormat2 *E2SMKPMIndicationMessageFormat2
} // E2SMKPMIndicationMessage_IndicationMessageformats

type E2SMKPMIndicationMessage struct { // [{'type': 'CHOICE', 'members': [{'type': 'E2SM-KPM-IndicationMessage-Format1', 'name': 'indicationMessage-Format1'}, {'type': 'E2SM-KPM-IndicationMessage-Format2', 'name': 'indicationMessage-Format2'}, None], 'name': 'indicationMessage-formats'}, None]
    IndicationMessageformats E2SMKPMIndicationMessage_IndicationMessageformats
}

func (self * E2SMKPMIndicationMessage) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    var Unpack_indicationMessageformats = func(stream *Stream, self *E2SMKPMIndicationMessage_IndicationMessageformats) {
        //coptions := []string{"indicationMessage-Format1","indicationMessage-Format2"}
        choice := stream.get_choice(1, 1, 2)
        choice_len := 0
        choice_loc := 0
        if choice >= 2 {
            choice_len = stream.parse_len(0)
            choice_loc = stream.get_location()
            log.Info("Extension choice option [%d] len %d in E2SMKPMIndicationMessage_IndicationMessageformats\n", choice, choice_len)
        }
        if choice == 0 { //ch1
            self.IndicationMessageFormat1 = &E2SMKPMIndicationMessageFormat1{}//cho6
            self.IndicationMessageFormat1.Unpack(stream)
        } else if choice == 1 { //ch2
            self.IndicationMessageFormat2 = &E2SMKPMIndicationMessageFormat2{}//cho6
            self.IndicationMessageFormat2.Unpack(stream)
        }//end of if else

        if choice >= 2 {
            stream.set_location(choice_loc, choice_len)
        }
    }
    Unpack_indicationMessageformats(stream, &self.IndicationMessageformats)// p2
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMIndicationMessage) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    var Pack_indicationMessageformats = func(stream *Stream, self E2SMKPMIndicationMessage_IndicationMessageformats) {
        if self.IndicationMessageFormat1 != nil {
            stream.set_choice(0, 1, 1, 2)
            self.IndicationMessageFormat1.Pack(stream)//2
        } else if self.IndicationMessageFormat2 != nil {
            stream.set_choice(1, 1, 1, 2)
            self.IndicationMessageFormat2.Pack(stream)//2
        }

    }
    Pack_indicationMessageformats(stream, self.IndicationMessageformats) //f2
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMKPMIndicationMessageFormat1 struct { // [{'type': 'MeasurementData', 'name': 'measData'}, {'type': 'MeasurementInfoList', 'name': 'measInfoList', 'optional': True}, {'type': 'GranularityPeriod', 'name': 'granulPeriod', 'optional': True}, None]
    MeasData MeasurementData
    MeasInfoList *MeasurementInfoList
    GranulPeriod *GranularityPeriod
}

func (self * E2SMKPMIndicationMessageFormat1) Unpack(stream *Stream) {
    measInfoList_flag := 0x00000002
    granulPeriod_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.MeasData.Unpack(stream)// p8
    if (measInfoList_flag & _flags) == measInfoList_flag { //cond2
        self.MeasInfoList = &MeasurementInfoList{}//7{'type': 'MeasurementInfoList', 'name': 'measInfoList', 'optional': True}
        self.MeasInfoList.Unpack(stream)// p8
    }
    if (granulPeriod_flag & _flags) == granulPeriod_flag { //cond2
        self.GranulPeriod = &GranularityPeriod{}//7{'type': 'GranularityPeriod', 'name': 'granulPeriod', 'optional': True}
        self.GranulPeriod.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMIndicationMessageFormat1) Pack(stream *Stream) {
    const measInfoList_flag uint = 0x00000002
    const granulPeriod_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeasData.Pack(stream)
    if self.MeasInfoList != nil { 
        _flags |= measInfoList_flag
        self.MeasInfoList.Pack(stream)
    }//end of optional
    if self.GranulPeriod != nil { 
        _flags |= granulPeriod_flag
        self.GranulPeriod.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type E2SMKPMIndicationMessageFormat2 struct { // [{'type': 'MeasurementData', 'name': 'measData'}, {'type': 'MeasurementCondUEidList', 'name': 'measCondUEidList'}, {'type': 'GranularityPeriod', 'name': 'granulPeriod', 'optional': True}, None]
    MeasData MeasurementData
    MeasCondUEidList MeasurementCondUEidList
    GranulPeriod *GranularityPeriod
}

func (self * E2SMKPMIndicationMessageFormat2) Unpack(stream *Stream) {
    granulPeriod_flag := 0x00000002
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(2)
    self.MeasData.Unpack(stream)// p8
    self.MeasCondUEidList.Unpack(stream)// p8
    if (granulPeriod_flag & _flags) == granulPeriod_flag { //cond2
        self.GranulPeriod = &GranularityPeriod{}//7{'type': 'GranularityPeriod', 'name': 'granulPeriod', 'optional': True}
        self.GranulPeriod.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMIndicationMessageFormat2) Pack(stream *Stream) {
    const granulPeriod_flag uint = 0x00000002
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(2)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.MeasData.Pack(stream)
    self.MeasCondUEidList.Pack(stream)
    if self.GranulPeriod != nil { 
        _flags |= granulPeriod_flag
        self.GranulPeriod.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 2)
}//end

type E2SMKPMRANfunctionDescription_RicReportStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-ReportStyle-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ReportStyle-List', 'optional': True}
    Items []RICReportStyleItem
}
type E2SMKPMRANfunctionDescription_RicEventTriggerStyleList struct { //SEQOF {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-EventTriggerStyle-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-EventTriggerStyle-List', 'optional': True}
    Items []RICEventTriggerStyleItem
}
type E2SMKPMRANfunctionDescription struct { // [{'type': 'RANfunction-Name', 'name': 'ranFunction-Name'}, {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-EventTriggerStyle-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-EventTriggerStyle-List', 'optional': True}, {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-ReportStyle-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ReportStyle-List', 'optional': True}, None]
    RanFunctionName RANfunctionName
    RicEventTriggerStyleList *E2SMKPMRANfunctionDescription_RicEventTriggerStyleList
    RicReportStyleList *E2SMKPMRANfunctionDescription_RicReportStyleList
}

func (self * E2SMKPMRANfunctionDescription) Unpack(stream *Stream) {
    ricEventTriggerStyleList_flag := 0x00000002
    ricReportStyleList_flag := 0x00000004
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(3)
    self.RanFunctionName.Unpack(stream)// p8
    if (ricEventTriggerStyleList_flag & _flags) == ricEventTriggerStyleList_flag { //cond1
        var Unpack_ricEventTriggerStyleList = func(stream *Stream, self *E2SMKPMRANfunctionDescription_RicEventTriggerStyleList){// Seq6 E2SMKPMRANfunctionDescription {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-EventTriggerStyle-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-EventTriggerStyle-List', 'optional': True}
            _size := stream.get_listsize(63)
            _size += 1
            self.Items = make([]RICEventTriggerStyleItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RicEventTriggerStyleList = &E2SMKPMRANfunctionDescription_RicEventTriggerStyleList{}//3
        Unpack_ricEventTriggerStyleList(stream, self.RicEventTriggerStyleList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-EventTriggerStyle-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-EventTriggerStyle-List', 'optional': True}
    }
    if (ricReportStyleList_flag & _flags) == ricReportStyleList_flag { //cond1
        var Unpack_ricReportStyleList = func(stream *Stream, self *E2SMKPMRANfunctionDescription_RicReportStyleList){// Seq6 E2SMKPMRANfunctionDescription {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-ReportStyle-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ReportStyle-List', 'optional': True}
            _size := stream.get_listsize(63)
            _size += 1
            self.Items = make([]RICReportStyleItem, _size)//1
            for i := 0; i < _size; i++ {
                self.Items[i].Unpack(stream)
            }
        }

        self.RicReportStyleList = &E2SMKPMRANfunctionDescription_RicReportStyleList{}//3
        Unpack_ricReportStyleList(stream, self.RicReportStyleList)// p1 {'type': 'SEQUENCE OF', 'element': {'type': 'RIC-ReportStyle-Item'}, 'size': [(1, 'maxnoofRICStyles')], 'name': 'ric-ReportStyle-List', 'optional': True}
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * E2SMKPMRANfunctionDescription) Pack(stream *Stream) {
    const ricEventTriggerStyleList_flag uint = 0x00000002
    const ricReportStyleList_flag uint = 0x00000004
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(3)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RanFunctionName.Pack(stream)
    if self.RicEventTriggerStyleList != nil { //YY
        _flags |= ricEventTriggerStyleList_flag
        var Pack_ricEventTriggerStyleList = func(stream *Stream, self E2SMKPMRANfunctionDescription_RicEventTriggerStyleList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 63)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ricEventTriggerStyleList(stream, *self.RicEventTriggerStyleList) //f1
    }//end of optional
    if self.RicReportStyleList != nil { //YY
        _flags |= ricReportStyleList_flag
        var Pack_ricReportStyleList = func(stream *Stream, self E2SMKPMRANfunctionDescription_RicReportStyleList) { //seqof 2
            _size := len(self.Items)
            stream.set_listsize(_size-1, 63)
            for _, item := range self.Items {// seqof structure
                item.Pack(stream)
            }
            return

        }

        Pack_ricReportStyleList(stream, *self.RicReportStyleList) //f1
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 3)
}//end

type RICEventTriggerStyleItem struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-EventTriggerStyle-Type'}, {'type': 'RIC-Style-Name', 'name': 'ric-EventTriggerStyle-Name'}, {'type': 'RIC-Format-Type', 'name': 'ric-EventTriggerFormat-Type'}, None]
    RicEventTriggerStyleType RICStyleType
    RicEventTriggerStyleName RICStyleName
    RicEventTriggerFormatType RICFormatType
}

func (self * RICEventTriggerStyleItem) Unpack(stream *Stream) {
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

func (self * RICEventTriggerStyleItem) Pack(stream *Stream) {
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

type RICReportStyleItem struct { // [{'type': 'RIC-Style-Type', 'name': 'ric-ReportStyle-Type'}, {'type': 'RIC-Style-Name', 'name': 'ric-ReportStyle-Name'}, {'type': 'RIC-Format-Type', 'name': 'ric-ActionFormat-Type'}, {'type': 'MeasurementInfo-Action-List', 'name': 'measInfo-Action-List'}, {'type': 'RIC-Format-Type', 'name': 'ric-IndicationHeaderFormat-Type'}, {'type': 'RIC-Format-Type', 'name': 'ric-IndicationMessageFormat-Type'}, None]
    RicReportStyleType RICStyleType
    RicReportStyleName RICStyleName
    RicActionFormatType RICFormatType
    MeasInfoActionList MeasurementInfoActionList
    RicIndicationHeaderFormatType RICFormatType
    RicIndicationMessageFormatType RICFormatType
}

func (self * RICReportStyleItem) Unpack(stream *Stream) {
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(1)
    self.RicReportStyleType.Unpack(stream)// p8
    self.RicReportStyleName.Unpack(stream)// p8
    self.RicActionFormatType.Unpack(stream)// p8
    self.MeasInfoActionList.Unpack(stream)// p8
    self.RicIndicationHeaderFormatType.Unpack(stream)// p8
    self.RicIndicationMessageFormatType.Unpack(stream)// p8
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * RICReportStyleItem) Pack(stream *Stream) {
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(1)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    self.RicReportStyleType.Pack(stream)
    self.RicReportStyleName.Pack(stream)
    self.RicActionFormatType.Pack(stream)
    self.MeasInfoActionList.Pack(stream)
    self.RicIndicationHeaderFormatType.Pack(stream)
    self.RicIndicationMessageFormatType.Pack(stream)
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 1)
}//end

type E2SMPDU_Message struct { // [{'type': 'INTEGER', 'restricted-to': [(0, 2)], 'name': 'version'}, {'type': 'E2SM-KPM-EventTriggerDefinition', 'name': 'eventTriggerDefinition', 'optional': True}, {'type': 'E2SM-KPM-ActionDefinition', 'name': 'actionDefinition', 'optional': True}, {'type': 'E2SM-KPM-IndicationHeader', 'name': 'indicationHeader', 'optional': True}, {'type': 'E2SM-KPM-IndicationMessage', 'name': 'indicationMessage', 'optional': True}, {'type': 'E2SM-KPM-RANfunction-Description', 'name': 'function-Description', 'optional': True}, None]
    Version INTEGER
    EventTriggerDefinition *E2SMKPMEventTriggerDefinition
    ActionDefinition *E2SMKPMActionDefinition
    IndicationHeader *E2SMKPMIndicationHeader
    IndicationMessage *E2SMKPMIndicationMessage
    FunctionDescription *E2SMKPMRANfunctionDescription
}
type E2SMPDU struct { // [{'type': 'SEQUENCE', 'members': [{'type': 'INTEGER', 'restricted-to': [(0, 2)], 'name': 'version'}, {'type': 'E2SM-KPM-EventTriggerDefinition', 'name': 'eventTriggerDefinition', 'optional': True}, {'type': 'E2SM-KPM-ActionDefinition', 'name': 'actionDefinition', 'optional': True}, {'type': 'E2SM-KPM-IndicationHeader', 'name': 'indicationHeader', 'optional': True}, {'type': 'E2SM-KPM-IndicationMessage', 'name': 'indicationMessage', 'optional': True}, {'type': 'E2SM-KPM-RANfunction-Description', 'name': 'function-Description', 'optional': True}, None], 'name': 'message'}]
    Message E2SMPDU_Message
}

func (self * E2SMPDU) Unpack(stream *Stream) {
    var Unpack_message = func(stream *Stream, self *E2SMPDU_Message) { //[{'type': 'INTEGER', 'restricted-to': [(0, 2)], 'name': 'version'}, {'type': 'E2SM-KPM-EventTriggerDefinition', 'name': 'eventTriggerDefinition', 'optional': True}, {'type': 'E2SM-KPM-ActionDefinition', 'name': 'actionDefinition', 'optional': True}, {'type': 'E2SM-KPM-IndicationHeader', 'name': 'indicationHeader', 'optional': True}, {'type': 'E2SM-KPM-IndicationMessage', 'name': 'indicationMessage', 'optional': True}, {'type': 'E2SM-KPM-RANfunction-Description', 'name': 'function-Description', 'optional': True}, None]
        eventTriggerDefinition_flag := 0x00000002
        actionDefinition_flag := 0x00000004
        indicationHeader_flag := 0x00000008
        indicationMessage_flag := 0x00000010
        functionDescription_flag := 0x00000020
        const ext_flag = 0x00000001
        _flags := 0
        _extflags := 0
        _ecount := 0
        _flags = stream.get_flags(6)
        var Unpack_version = func (st *Stream, self *INTEGER) {
            self.Value = st.parsef_Integer(3, 2, 0, 0)
        }
        Unpack_version(stream, &self.Version)// p2
        if (eventTriggerDefinition_flag & _flags) == eventTriggerDefinition_flag { //cond2
            self.EventTriggerDefinition = &E2SMKPMEventTriggerDefinition{}//7{'type': 'E2SM-KPM-EventTriggerDefinition', 'name': 'eventTriggerDefinition', 'optional': True}
            self.EventTriggerDefinition.Unpack(stream)// p8
        }
        if (actionDefinition_flag & _flags) == actionDefinition_flag { //cond2
            self.ActionDefinition = &E2SMKPMActionDefinition{}//7{'type': 'E2SM-KPM-ActionDefinition', 'name': 'actionDefinition', 'optional': True}
            self.ActionDefinition.Unpack(stream)// p8
        }
        if (indicationHeader_flag & _flags) == indicationHeader_flag { //cond2
            self.IndicationHeader = &E2SMKPMIndicationHeader{}//7{'type': 'E2SM-KPM-IndicationHeader', 'name': 'indicationHeader', 'optional': True}
            self.IndicationHeader.Unpack(stream)// p8
        }
        if (indicationMessage_flag & _flags) == indicationMessage_flag { //cond2
            self.IndicationMessage = &E2SMKPMIndicationMessage{}//7{'type': 'E2SM-KPM-IndicationMessage', 'name': 'indicationMessage', 'optional': True}
            self.IndicationMessage.Unpack(stream)// p8
        }
        if (functionDescription_flag & _flags) == functionDescription_flag { //cond2
            self.FunctionDescription = &E2SMKPMRANfunctionDescription{}//7{'type': 'E2SM-KPM-RANfunction-Description', 'name': 'function-Description', 'optional': True}
            self.FunctionDescription.Unpack(stream)// p8
        }
        stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
        return
    }
    Unpack_message(stream, &self.Message)// p2
    return
}

func (self * E2SMPDU) Pack(stream *Stream) {
    var Pack_message = func(stream *Stream, self E2SMPDU_Message) {//seq
        const eventTriggerDefinition_flag uint = 0x00000002
        const actionDefinition_flag uint = 0x00000004
        const indicationHeader_flag uint = 0x00000008
        const indicationMessage_flag uint = 0x00000010
        const functionDescription_flag uint = 0x00000020
        const ext_flag int = 0x00000001
        _flagReserve := stream.reserve_flags(6)
        var _extReserve uint32 = 0
        var _extPresent bool = false
        var _flags uint = 0
        var _extflags uint = 0
        var Pack_version = func (st *Stream, self INTEGER){
            st.formatf_Integer(self.Value, 3, 2, 0, 0)
        }
        Pack_version(stream, self.Version) //f2
        if self.EventTriggerDefinition != nil { 
            _flags |= eventTriggerDefinition_flag
            self.EventTriggerDefinition.Pack(stream)
        }//end of optional
        if self.ActionDefinition != nil { 
            _flags |= actionDefinition_flag
            self.ActionDefinition.Pack(stream)
        }//end of optional
        if self.IndicationHeader != nil { 
            _flags |= indicationHeader_flag
            self.IndicationHeader.Pack(stream)
        }//end of optional
        if self.IndicationMessage != nil { 
            _flags |= indicationMessage_flag
            self.IndicationMessage.Pack(stream)
        }//end of optional
        if self.FunctionDescription != nil { 
            _flags |= functionDescription_flag
            self.FunctionDescription.Pack(stream)
        }//end of optional
        if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
        stream.set_flags(_flags, _flagReserve, 6)
    }//end
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
var maxnoofCells uint64 = 16384
var maxnoofRICStyles uint64 = 63
var maxnoofMeasurementInfo uint64 = 65535
var maxnoofLabelInfo uint64 = 2147483647
var maxnoofMeasurementRecord uint64 = 65535
var maxnoofMeasurementValue uint64 = 2147483647
var maxnoofConditionInfo uint64 = 32768
var maxnoofUEID uint64 = 65535
var idEventTriggerDefinition uint64 = 0
const ProcedureCodeEventTriggerDefinition = 0
var idActionDefinition uint64 = 1
const ProcedureCodeActionDefinition = 1
var idIndicationHeader uint64 = 2
const ProcedureCodeIndicationHeader = 2
var idIndicationMessage uint64 = 3
const ProcedureCodeIndicationMessage = 3
var idFunctionDescription uint64 = 4
const ProcedureCodeFunctionDescription = 4
