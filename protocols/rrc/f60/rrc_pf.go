
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package rrc
import (
  log "github.com/sirupsen/logrus"
)
var version = "vf60"

func fmtrrc() {log.Debug("rrc")}
type DRBToAddMod_CnAssociation struct { //[{'type': 'INTEGER', 'restricted-to': [(0, 15)], 'name': 'eps-BearerIdentity'}, {'type': 'SDAP-Config', 'name': 'sdap-Config'}]
    EpsBearerIdentity *INTEGER
    SdapConfig *SDAPConfig
} // DRBToAddMod_CnAssociation

type DRBToAddMod struct { // [{'type': 'CHOICE', 'members': [{'type': 'INTEGER', 'restricted-to': [(0, 15)], 'name': 'eps-BearerIdentity'}, {'type': 'SDAP-Config', 'name': 'sdap-Config'}], 'name': 'cnAssociation', 'optional': True}, {'type': 'DRB-Identity', 'name': 'drb-Identity'}, {'type': 'ENUMERATED', 'values': [('tRue', 0)], 'name': 'reestablishPDCP', 'optional': True}, {'type': 'ENUMERATED', 'values': [('tRue', 0)], 'name': 'recoverPDCP', 'optional': True}, {'type': 'PDCP-Config', 'name': 'pdcp-Config', 'optional': True}, None]
    CnAssociation *DRBToAddMod_CnAssociation
    DrbIdentity DRBIdentity
    ReestablishPDCP *ENUMERATED
    RecoverPDCP *ENUMERATED
    PdcpConfig *PDCPConfig
}

func (self * DRBToAddMod) Unpack(stream *Stream) {
    cnAssociation_flag := 0x00000002
    reestablishPDCP_flag := 0x00000004
    recoverPDCP_flag := 0x00000008
    pdcpConfig_flag := 0x00000010
    const ext_flag = 0x00000001
    _flags := 0
    _extflags := 0
    _ecount := 0
    _flags = stream.get_flags(5)
    if (cnAssociation_flag & _flags) == cnAssociation_flag { //cond1
        var Unpack_cnAssociation = func(stream *Stream, self *DRBToAddMod_CnAssociation) {
            //coptions := []string{"eps-BearerIdentity","sdap-Config"}
            choice := stream.get_choice(1, 0, 2)
            if choice == 0 { //ch1
                var Unpack_epsBearerIdentity = func (st *Stream, self *INTEGER) {
                    self.Value = st.parsef_Integer(16, 4, 0, 0)
                }
                self.EpsBearerIdentity = &INTEGER{}//cho5
                Unpack_epsBearerIdentity(stream, self.EpsBearerIdentity);
            } else if choice == 1 { //ch2
                self.SdapConfig = &SDAPConfig{}//cho6
                self.SdapConfig.Unpack(stream)
            }//end of if else

        }
        self.CnAssociation = &DRBToAddMod_CnAssociation{}//5{'type': 'CHOICE', 'members': [{'type': 'INTEGER', 'restricted-to': [(0, 15)], 'name': 'eps-BearerIdentity'}, {'type': 'SDAP-Config', 'name': 'sdap-Config'}], 'name': 'cnAssociation', 'optional': True}
        Unpack_cnAssociation(stream, self.CnAssociation)// p1 {'type': 'CHOICE', 'members': [{'type': 'INTEGER', 'restricted-to': [(0, 15)], 'name': 'eps-BearerIdentity'}, {'type': 'SDAP-Config', 'name': 'sdap-Config'}], 'name': 'cnAssociation', 'optional': True}
    }
    self.DrbIdentity.Unpack(stream)// p8
    if (reestablishPDCP_flag & _flags) == reestablishPDCP_flag { //cond1
        var Unpack_reestablishPDCP = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(0, 1, 0)
        }
        self.ReestablishPDCP = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('tRue', 0)], 'name': 'reestablishPDCP', 'optional': True}
        Unpack_reestablishPDCP(stream, self.ReestablishPDCP)// p1 {'type': 'ENUMERATED', 'values': [('tRue', 0)], 'name': 'reestablishPDCP', 'optional': True}
    }
    if (recoverPDCP_flag & _flags) == recoverPDCP_flag { //cond1
        var Unpack_recoverPDCP = func(st *Stream, self *ENUMERATED) {
            self.Value = st.parsef_Enumerated(0, 1, 0)
        }
        self.RecoverPDCP = &ENUMERATED{}//6{'type': 'ENUMERATED', 'values': [('tRue', 0)], 'name': 'recoverPDCP', 'optional': True}
        Unpack_recoverPDCP(stream, self.RecoverPDCP)// p1 {'type': 'ENUMERATED', 'values': [('tRue', 0)], 'name': 'recoverPDCP', 'optional': True}
    }
    if (pdcpConfig_flag & _flags) == pdcpConfig_flag { //cond2
        self.PdcpConfig = &PDCPConfig{}//7{'type': 'PDCP-Config', 'name': 'pdcp-Config', 'optional': True}
        self.PdcpConfig.Unpack(stream)// p8
    }
    stream.handle_unknown_ext(_ecount, _extflags, 0, _flags)
    return
}

func (self * DRBToAddMod) Pack(stream *Stream) {
    const cnAssociation_flag uint = 0x00000002
    const reestablishPDCP_flag uint = 0x00000004
    const recoverPDCP_flag uint = 0x00000008
    const pdcpConfig_flag uint = 0x00000010
    const ext_flag int = 0x00000001
    _flagReserve := stream.reserve_flags(5)
    var _extReserve uint32 = 0
    var _extPresent bool = false
    var _flags uint = 0
    var _extflags uint = 0
    if self.CnAssociation != nil { //YY
        _flags |= cnAssociation_flag
        var Pack_cnAssociation = func(stream *Stream, self DRBToAddMod_CnAssociation) {
            if self.EpsBearerIdentity != nil {
                stream.set_choice(0, 1, 0, 2)
                var Pack_epsBearerIdentity = func (st *Stream, self INTEGER){
                    st.formatf_Integer(self.Value, 16, 4, 0, 0)
                }
                Pack_epsBearerIdentity(stream, *self.EpsBearerIdentity)//3
            } else if self.SdapConfig != nil {
                stream.set_choice(1, 1, 0, 2)
                self.SdapConfig.Pack(stream)//2
            }

        }
        Pack_cnAssociation(stream, *self.CnAssociation) //f1
    }//end of optional
    self.DrbIdentity.Pack(stream)
    if self.ReestablishPDCP != nil { //YY
        _flags |= reestablishPDCP_flag
        var Pack_reestablishPDCP = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 0, 1, 0)
        }
        Pack_reestablishPDCP(stream, *self.ReestablishPDCP) //f1
    }//end of optional
    if self.RecoverPDCP != nil { //YY
        _flags |= recoverPDCP_flag
        var Pack_recoverPDCP = func(st *Stream, self ENUMERATED) {
            st.formatf_Enumerated(self.Value, 0, 1, 0)
        }
        Pack_recoverPDCP(stream, *self.RecoverPDCP) //f1
    }//end of optional
    if self.PdcpConfig != nil { 
        _flags |= pdcpConfig_flag
        self.PdcpConfig.Pack(stream)
    }//end of optional
    if _extPresent { stream.set_flags(_extflags, _extReserve, 0) }
    stream.set_flags(_flags, _flagReserve, 5)
}//end

func (self *MeasTriggerQuantityEUTRA)Unpack(stream *Stream) {
    //coptions := []string{"rsrp","rsrq","sinr","Unknown"}
    choice := stream.get_choice(2, 0, 3)
    if choice == 0 { //ch1
        self.Rsrp = &RSRPRangeEUTRA{}//cho6
        self.Rsrp.Unpack(stream)
    } else if choice == 1 { //ch2
        self.Rsrq = &RSRQRangeEUTRA{}//cho6
        self.Rsrq.Unpack(stream)
    } else if choice == 2 { //ch2
        self.Sinr = &SINRRangeEUTRA{}//cho6
        self.Sinr.Unpack(stream)
    }//end of if else

}
func (self * MeasTriggerQuantityEUTRA) Pack(stream *Stream) {
    if self.Rsrp != nil {
        stream.set_choice(0, 2, 0, 3)
        self.Rsrp.Pack(stream)//2
    } else if self.Rsrq != nil {
        stream.set_choice(1, 2, 0, 3)
        self.Rsrq.Pack(stream)//2
    } else if self.Sinr != nil {
        stream.set_choice(2, 2, 0, 3)
        self.Sinr.Pack(stream)//2
    }

}
type MeasTriggerQuantityEUTRA struct { //[{'type': 'RSRP-RangeEUTRA', 'name': 'rsrp'}, {'type': 'RSRQ-RangeEUTRA', 'name': 'rsrq'}, {'type': 'SINR-RangeEUTRA', 'name': 'sinr'}]
    Rsrp *RSRPRangeEUTRA
    Rsrq *RSRQRangeEUTRA
    Sinr *SINRRangeEUTRA
} // MeasTriggerQuantityEUTRA

