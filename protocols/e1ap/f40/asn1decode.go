
/*********************************************************************************
 * Copyright 2020 RideNext Software Solutions (I) Pvt. Ltd. All rights reserved. *
 *********************************************************************************/


package e1ap


import (
	"fmt"
	"encoding/hex"
	"encoding/json"
)

type Stream struct {
    buff []byte
    c_ind int
    c_bit int
}

type HexBytes []byte

type NULL struct {
}
type REAL struct {
    Value float64
}
type BOOLEAN struct {
  Value bool
}
type BITSTRING struct {
  Value HexBytes
  Len int
}
type OCTETSTRING struct {
  Value HexBytes
  Item interface{}
}

func (self *OCTETSTRING) Pack(st *Stream) {
   st.format_len(len(self.Value), 0)
   st.formatf_OctString(self.Value, 0)
}

func (self *OCTETSTRING) Unpack(st *Stream) {
    _len := st.parse_len(0)
    self.Value = st.parsef_OctString(_len)
}

type INTEGER struct {
  Value uint64
}
type PRINTABLESTRING struct {
   Value string
}
type PrintableString struct {
   Value string
}
type ENUMERATED struct {
   Value int
}
type OBJECTIDENTIFIER struct {
   Value string
}

func (self *Stream) dumpInfo(prompt string) {
    fmt.Printf("%s [%d][%d] = [%02x]\n", prompt, self.c_ind, self.c_bit, self.buff[self.c_ind])
}

func (self *Stream) Init(buff []byte) {
   self.buff = buff
   self.c_ind = 0
   self.c_bit = 0
}

type StreamIF interface {
   Initialize()
   Reset()
   get_location() int
   set_location(int, int)
   reset_bits()
   read_bit() bool
   write_bit() bool
   parse_ext() bool
   format_ext(int)
   get_bits(bits int) int
   set_bits(val uint, bits int)
   get_rbits(bits int) int
   get_byte() int
   set_byte(val int)
   get_listsize(count uint64) int
   set_listsize(val int,count uint64)
   get_choice(bits int) int
   set_choice(val int, bits int)
   get_flags(bits int) int
   set_len(val int, loc uint32)
   parse_len(size int) int
   format_len(lval int, values int)
   parsef_OctString(size int) HexBytes
   formatf_OctString(val HexBytes, sz int)
   parsef_BitString(values int, size int) HexBytes
   formatf_BitString(val HexBytes, bits int)
   parsef_Integer(values uint64, size int, bits int, ext int) uint64
   formatf_Integer(val uint64, values uint64, size int, bits int, ext int)
   parsef_Enumerated(size int, ext int) int
   formatf_Enumerated(val int, size int, ext int)
   parsef_PriString(size int) string
   formatf_PriString(val string, size int)
   parse_blen(bits int, ext int) int
   format_blen(val int, bits int, ext int)
   parse_olen(size int) int
   format_olen(val int, size int)
   parsef_Real(ext int) float64
   formatf_Real(val float64, ext int)
   parsef_bool() bool
   formatf_bool( val bool)
   parsef_Null()
   formatf_Null()
   parsef_ObjectID() string
   handle_unknown_ext(int, int, int, int)
   get_current_location() (int)
   read_flags(int) int
   write_flags(int, int)
   parsef_sequence_ext(int) ( int, int)
   reserve_flags(sz int) uint32
   reserve_len() uint32
   set_flags(flags uint, x uint32, y int)
   Get_buff() []byte
}

func (s *Stream) Reset() {
    s.c_ind = 0
    s.c_bit=0
}

func (s *Stream) set_len(loc uint32){
    //fmt.Printf("set_len %x\n", loc)
    s.reset_bits()
    v := s.c_ind - int(loc+1)
    pos := s.c_ind
    s.c_ind = int(loc)
    pos += s.format_len(v, 0)
    s.c_ind = pos
    s.c_bit = 0
}

func (s *Stream) Get_buff() []byte {
    return s.buff[:s.c_ind]
}
func (s *Stream)  set_flags(v uint, loc uint32 , size int) {
    //fmt.Printf("set_flags %d %x %d\n", v, loc, size)
    ind := s.c_ind
    bit := s.c_bit
    s.c_ind = int(loc>>4)
    s.c_bit = int(loc&0xf)
    s.set_rbits(uint(v), int(size))
    s.c_ind = ind
    s.c_bit = bit

}

func (s *Stream) reserve_len() uint32 {
    //fmt.Printf("reserve_len\n")
    s.reset_bits()
    retval := s.c_ind
    s.c_ind += 1
    return uint32(retval)
}
func (s *Stream) reserve_flags(sz int) uint32 {
    //fmt.Printf("reserve_flags %d\n", sz)
    reserve := uint32((s.c_ind<<4) | s.c_bit)
    s.set_bits(0, sz)
    return reserve
}

func (s*Stream) parsef_sequence_ext(cnt int) ( int, int) {
     //fmt.Printf("parsef_sequence_ext %d %d \n", s.c_ind, s.c_bit)
     count := s.get_bits(7)
     count += 1
     _flags := s.get_flags(cnt)
     return count, _flags
}

func (s* Stream) parsef_Real(ext int) float64{
	fmt.Printf("TODO:parsef_Real %d\n", ext)
    return 0.0
}
func (s* Stream) formatf_Real(val float64, ext int) {
}

func (s *Stream) parsef_ObjectID() string{
	fmt.Printf("TODO:parsef_ObjectID %d\n")
    return "."
}

func (s *Stream) parsef_bool() bool{
     return (s.get_bits(1) == 1)
}

func (s *Stream) formatf_bool(val bool) {
    if val {
	s.set_bits(1, 1)
    } else {
        s.set_bits(0, 1)
    }
}

func (s *Stream) read_flags(int) (int){
    fmt.Printf("TODO:read_flags\n")
    return 0
}
func (s *Stream) write_flags(int, int){
    fmt.Printf("TODO:write_flags\n")
}
func (s *Stream) get_current_location()( int){
    return s.c_ind
}

func (s *Stream) handle_unknown_ext( cnt int,  flg int , ext int, _flg int) {
    //fmt.Printf("TODO:handle_unknown_ext\n")
}

func (s *Stream)  Initialize (buf []byte) {
    s.buff = buf
    s.c_ind = 0
    s.c_bit = 0
}

func (s *Stream) get_location() int {
    //fmt.Println("get_location", s.c_ind)
    return s.c_ind
}

func (s *Stream) set_location(loc int, ch int){
    s.reset_bits()
    s.c_ind = loc+ch
    s.c_bit = 0
    //fmt.Println("set_location", s.c_ind, loc, ch)
}

func (s *Stream)  reset_bits() {
    if s.c_bit > 0 {
       s.c_ind ++
       s.c_bit = 0
    }
}

func (s *Stream)  format_ext(val int) {
    if val > 0 { s.write_bit(1) } else {s.write_bit(0) }
    
}

func (s *Stream)  parse_ext() bool {
    return s.read_bit()
}

func (s *Stream)  write_bit(bv int) {
    if len(s.buff) <= s.c_ind {
       s.buff = append(s.buff, []byte{0, 0, 0, 0}...)
    }
    if bv  != 0 {
        s.buff[s.c_ind] |= (1<<(7-s.c_bit))
    } else {
        s.buff[s.c_ind] &= ^(1<<(7-s.c_bit))
    }
    s.c_bit += 1
    if (s.c_bit&8) > 0 {
       s.c_bit = 0
       s.c_ind += 1
    }
}
func (s *Stream)  read_bit() bool {
    v := (s.buff[s.c_ind] & (0x1<<uint8(7-s.c_bit))) > 0
    s.c_bit += 1
    if s.c_bit & 8 == 8 {
        s.c_bit = 0
        s.c_ind ++
    }
    return v
}

func (s *Stream)  set_rbits(val uint, bits int){
    if bits == 0 { return }
    for i:=0; i < bits; i++ {
        s.write_bit(int(val&1))
        val >>= 1
    }
}
func (s *Stream)  set_bits(val uint, bits int){
    //fmt.Printf("TODO:set_bits\n")
    if bits == 0 { return }
    bmask := uint(1 << (bits-1))
    for i := 0 ; i < bits; i++ {
        if (val&bmask) > 0 {
           s.write_bit(1)
	} else {
           s.write_bit(0)
	}
	bmask >>= 1
    }
    return
}
func (s *Stream)  get_bits(bits int) int {
    if bits == 0 { return 0}
    v := 0
    for i := 0; i < bits; i ++ {
        v <<= 1
        if s.read_bit() {
            v |= 1
        }
    }
    return v

}

func (s *Stream)  get_rbits(bits int) int {
    if bits == 0 { return 0}
    v := 0
    for i := 0; i < bits; i ++ {
        if s.read_bit() {
            v |= (1 << uint(i))
        }
    }
    return v
}

func (s *Stream)  set_byte( val int) {
    if len(s.buff) <= s.c_ind { s.buff = append(s.buff, make([]byte, 1024)...)}
    s.buff[s.c_ind] = byte(val&0xff)
}

func (s *Stream)  get_byte() int {
    i := 0
    if s.c_ind < len(s.buff) {
        i = int(s.buff[s.c_ind])
        s.c_ind ++
    }
    return i
}

func (s *Stream)  format_len(v int, values int) int {
    //fmt.Printf("format_len %x %d\n", v, values)
    advance := 0
    s.reset_bits()
    if values > 8 && values <= 16 {
        s.set_byte(int((v>>8)&0xff))
        s.c_ind += 1
        s.set_byte(int(v&0xff))
	//fmt.Printf("format_len %d %d %02x\n", v, s.c_ind-1, s.buff[s.c_ind])
        s.c_ind += 1
        return advance
    }

    fl := 0
    if v > 127 {
	fl = 1
    }
    if len(s.buff) < (s.c_ind+v) { s.buff = append(s.buff, make([]byte, ((s.c_ind+v)-len(s.buff)))...)}
    s.set_bits(uint(fl), int(1))
    if fl == 1 {
        if v >= 16384 {
           s.set_bits(uint(fl), int(1))
	   blocks := v/16384
           s.set_bits(uint(blocks), int(6))
	   remlen := (v % 16384)
	   pos := s.c_ind
           s.c_ind += (blocks*16384)
           s.buff = append(s.buff[:s.c_ind], append([]byte{0}, s.buff[s.c_ind:]...)...)
           advance += 1
           advance += s.format_len(remlen, 0)
           s.c_ind = pos
           return advance
	} else {
           s.set_bits(uint((v&0x7f00)>>8), 7)
           s.buff = append(s.buff[:s.c_ind], append([]byte{0}, s.buff[s.c_ind:]...)...)
           advance += 1
	}
        //s.set_bits(uint((v&0x7f00)>>8), 7)
    } else { s.c_bit = 0}
    s.set_byte(int(v&0xff))
    s.c_ind += 1
    return advance
}
func (s *Stream)  parse_len(size int) int {
    s.reset_bits()
    v := 0
    //fmt.Println("parse_len", size, s.c_ind)
    if size > 8 && size <= 16 {
        v = s.get_byte()
        v <<= 8
        v |= s.get_byte()
	//fmt.Printf("parse_len %d %d %02x\n", v, s.c_ind-2, s.buff[s.c_ind-1])
        return v
    }
    fl := s.read_bit()
    if fl  {
        if s.read_bit() {
           v = s.get_bits(6)
	   v = v * 16384
	   cind := s.c_ind
	   s.c_ind += v
	   v1 := s.parse_len(0)
	   s.c_ind = cind
	   if v1 < 127 {
	      s.buff = append(s.buff[:cind+v], s.buff[cind+v+1:]...)
	   } else {
	      s.buff = append(s.buff[:cind+v], s.buff[cind+v+2:]...)
	   }
	   v += v1
	} else {
           v = s.get_bits(6)
           v <<= 8
           v |= s.get_byte()
	}
    } else {
        v = s.get_bits(7)
    }
    //fmt.Println(v, s.c_ind)
    return v
}

func (s *Stream)  set_listsize(v int, values uint64) {
    //fmt.Printf("TODO:set_listsize\n")
    if values > 65536 {
        s.format_len(v, int(values))
    } else if values <= 256 {
	bits := 0
	size := values-1
        for size > 0 {
            bits += 1
            size >>= 1
        }
        s.set_bits(uint(v), bits)
    } else {
        s.reset_bits()
        s.set_byte(int(v&0xff00) >> 8)
        s.c_ind += 1
        s.set_byte(int(v&0xff))
        s.c_ind += 1
    }
}
func (s *Stream)  get_listsize(values uint64) int {
    v := 0
    //fmt.Println("\nget_listsize", values, s.c_ind)
    if values > 65536 {
        v = s.parse_len(0)
    } else if values <= 256 {
	bits := 0
	size := values-1
        for  size > 0 {
            bits += 1
            size >>= 1
	}
        v = s.get_bits(bits)
    } else {
        s.reset_bits()
        v = s.get_byte()
        v <<= 8
        v |= s.get_byte()
    }
    //fmt.Println(v, s.c_ind)
    return v
}

func (s *Stream)  get_choice(bits int, ext int, optcnt int) int {
    if ext > 0 {
        ext = s.get_bits(1)
	if ext > 0 {
	    opt := s.get_bits(7)
            return (1 << bits) + opt
        }
	return s.get_bits(bits)
    }
    if bits > 0 {return s.get_bits(bits)}
    return 0
}

func (s *Stream)  set_choice(choice int, bits int, ext int, optcnt int) {
    //fmt.Printf("TODO:set_choice %x %x\n", choice, bits)
    if ext == 1 {
        if (choice >= (1 << bits)) {
            s.set_bits(1, 1)
	    choice -= (1 << bits)
	    if (choice  < 63 ) {
                s.set_bits(uint(choice), 7)
		return
	    }  else {
	        //TODO: if extensions are more than 63
	    }
        } else {
            s.set_bits(0, 1)
	}
    }
    s.set_bits(uint(choice), bits)
    //s.dumpInfo("Choice")
}

func (s *Stream)  get_flags(bits int) int {
    //fmt.Printf("get_flags %x %d %d\n", s.buff[s.c_ind], s.c_bit, bits)
    return s.get_rbits(bits)
}

func (s *Stream)  parsef_OctString(size int) HexBytes {
    defer func () {
        if err := recover(); err != nil {
            fmt.Printf("Parse Error %s", err)
        }
     }()

    if size <= 2 {
       return s.parsef_BitString(size*8, size*8)
    }
    s.reset_bits()
    ind := s.c_ind
    s.c_ind += size
    v := make([]byte, size)
    copy(v, s.buff[ind:ind+size])
    //fmt.Printf("parsef_OctString %d %02x\n", size, v[size-1])
    return v
    //return s.buff[ind:ind+size]
}
func (s *Stream)  parsef_BitString(values int, size int) HexBytes {
    if values > 16 {
        s.reset_bits()
    }
    sz := int(size/8)
    if size%8 > 0 {
        sz += 1
    }
    v := make([]byte, sz)
    for i := 0; i < sz; i++ {
        if size >=8 {
            //fmt.Printf("%+v %x %x\n",v, s.buff[s.c_ind], s.c_bit)
            v[i] = byte(s.get_bits(8))
            size -= 8
	} else {
            v[i] = byte(s.get_bits(size)<<uint(8-size))
	}
    }
    return v
}
func (s *Stream)  parsef_Integer(values uint64, size int, ext int, start int) uint64 {
    v := 0
    //fmt.Println("\nparsef_Integer", values, size, s.c_ind)
    if ext == 1 { s.get_bits(1) }
    if values == 0 {
        s.reset_bits()
	_ln := s.get_byte()
        for  _ln > 0 {
            v <<= 8
            v |= s.get_byte()
            _ln --
	}
    } else if values < 256 {
        if ext == 1 {
            v = s.get_bits(size-1)
	} else {
            v = s.get_bits(size)
	}
    } else if values == 256 {
        s.reset_bits()
        v=s.get_byte()
    } else if values <= 65536 {
        s.reset_bits()
        v = s.get_byte()
        v <<=8
        v |= s.get_byte()
    } else {
	fl := 0
	size -= ext
        if size <= 32 {
            fl = s.get_bits(2)
	} else if size <= 64 {
            fl = s.get_bits(3)
	} else {
            fl = s.get_bits(4)
	}
        s.reset_bits()
        for fl>=0 {
            v <<= 8
            v |= s.get_byte()
            fl -= 1
	}
    }
    //fmt.Println("\nparsef_Integer", v, s.c_ind)
    return uint64(v+start)
}

func (s *Stream)  parsef_Enumerated(bits int, rval int, ext int) int {
    //fmt.Println("\nparsef_Enumerated", size, s.c_ind)
    bits -= ext
    if ext > 0 {
        ext = s.get_bits(1)
	if ext > 0 {
	    opt := s.get_bits(7)
            return rval + opt
        }
	return s.get_bits(bits)
    }
    if bits > 1 {return s.get_bits(bits)}
    return 0
}

func (s *Stream)  formatf_PriString(val string, size int) {
    s.reset_bits()
    ind := s.c_ind
    s.c_ind +=  size 
    if len(s.buff) <= s.c_ind {
       s.buff = append(s.buff, make([]byte, s.c_ind-len(s.buff))...)
    }
    //fmt.Printf("TODO:formatf_PriString %s %d %02x %d %d\n", val, ind,  s.buff[ind], s.c_ind, size)
    copy(s.buff[ind:], []byte(val))
    //for  _, v := range []byte(val) {
    //    s.set_byte(int(v))
//	s.c_ind += 1
    //}
    //fmt.Printf("TODO:formatf_PriString %s %d %02x %d %d\n", val, ind,  s.buff[ind], s.c_ind, size)
}

func (s *Stream)  parsef_PriString(size int) string {
    s.reset_bits()
    ind := s.c_ind
    s.c_ind += size
    v := make([]byte, size)
    copy(v, s.buff[ind:ind+size])
    //fmt.Printf("parsef_PriString %s %d %02x %d %d\n", string(v), ind,  s.buff[ind], s.c_ind, size)
    return string(v)
    //return string(s.buff[ind:ind+size])
}

func (s *Stream)  format_blen(val int, bits int, ext int) {
    //fmt.Printf("TODO:format_blen\n")
    s.set_bits(uint(val), bits)
}

func (s *Stream)  parse_blen(bits int, ext int) int {
    return s.get_bits(bits)
}

func (s *Stream)  format_olen(val int, bits int) {
    //fmt.Printf("format_olen %d:%d[ %x %d]\n", s.c_ind, s.c_bit, val, bits)
    if bits > 0 && bits <= 8 {
        s.set_bits(uint(val), bits)
    } else {
        s.format_len(val, bits)
    }

}

func (s *Stream)  parse_olen(bits int) int {
    v := 0
    if bits > 0 && bits <= 8 {
        v = s.get_bits(bits)
    } else {
        v = s.parse_len(bits)
    }
    //fmt.Printf("parse_olen %d:%d [%d] len %x\n", s.c_ind, s.c_bit, bits, v)
    return v
}

func (s *Stream)  parsef_Null() {
    return
}

func (s *Stream)  formatf_Null() {
    return
}

func (s *Stream) formatf_OctString(val HexBytes, sz int) {
    //fmt.Printf("formatf_OctString %d %d %02x %02x\n", sz, len(val), val[0], val[len(val)-1])
    if sz == 1 {
        s.formatf_BitString(val, 8)
	return
    }
    s.reset_bits()
    blocks := len(val)/16384
    ind := 0
    if blocks > 0 {
       ind = blocks*16384
       copy(s.buff[s.c_ind:s.c_ind+ind], val[:ind])
       s.c_ind += ind
       s.c_ind += 1
       if (len(val) - ind) > 127 { s.c_ind += 1}
    }
    for _, v := range val[ind:] {
        s.set_byte(int(v))
        s.c_ind += 1
    }
    //fmt.Printf("formatf_OctString %02x\n", s.buff[s.c_ind-1])
}

func (s *Stream) formatf_BitString(v HexBytes, bits int) {
    //fmt.Printf("formatf_BitString %v %d\n", v, bits)
    if bits > 16 {s.reset_bits()}
    blocks := bits/16384
    ind := 0
    if blocks > 0 {
       ind = blocks*2048
       copy(s.buff[s.c_ind:s.c_ind+ind], v[:ind])
       s.c_ind += ind
       bits -= (ind*8)
       s.c_ind += 1
       if bits > 127 { s.c_ind += 1}
    }
    for _, val := range v[ind:] {
        if bits >= 8 {
            s.set_bits(uint(val), 8)
            bits -= 8
        } else {
            s.set_bits(uint(val>>(8-bits)), bits)
        }
    }
}
func (s *Stream) formatf_Integer(val uint64, values uint64, size int, ext int, start int) {
    //fmt.Printf("TODO:formatf_Integer %x %x %d %d %d\n", val, values, size, ext, start)
    val = val - uint64(start)
    if ext != 0 { s.set_bits(0, 1) }
    if values == 0 {
        s.reset_bits()
	_nums := make([]uint8, 0)
	_byte := 1
        _nums = append(_nums, uint8(val&0xff))
	bval := val >> 8
	for bval != 0 {
            _nums = append(_nums, uint8(bval&0xff))
	    bval >>= 8
	    _byte += 1
	}
	s.set_byte(_byte)
	_byte -= 1
	s.c_ind += 1
	for _byte >= 0 {
            s.set_byte(int(_nums[_byte]))
	    _byte -= 1
	    s.c_ind += 1
	}
    } else if values < 256 {
        s.set_bits(uint(val), int(size-ext))
    } else if values  == 256 {
        s.reset_bits()
        s.set_byte(int(val))
        s.c_ind += 1
    } else if values  <= 65536 {
        s.reset_bits()
        s.set_byte(int((val&0xff00) >> 8))
        s.c_ind += 1
        s.set_byte(int(val&0xff))
        s.c_ind += 1
    } else {
	fl := 0
        size -= ext
        if size <= 32 {
            fl = 2
        } else if size <= 64 {
            fl = 3
        } else {
            fl = 4
        }
	_nums := make([]uint8,0)
	_byte := 0
        _nums = append(_nums, uint8(val&0xff))
	bval := val >> 8
        for bval != 0 {
            _nums = append(_nums, uint8(bval&0xff))
            bval >>= 8
            _byte += 1
        }
        s.set_bits(uint(_byte), fl)
        s.reset_bits()
        for _byte > 0 {
            s.set_byte(int(_nums[_byte]))
            s.c_ind += 1
            _byte -= 1
        }
        s.set_byte(int(_nums[_byte]))
        s.c_ind += 1
    }
}
func (s *Stream) formatf_Enumerated(val int, bits int, rval int, ext int){
    //fmt.Printf("formatf_Enumerated %x %d %d\n", val, size, ext)
    bits -= ext
    if ext == 1 {
        if (val >= rval) {
            s.set_bits(1, 1)
	    val -= rval
	    if (val  < 63 ) {
                s.set_bits(uint(val), 7)
		return
	    }  else {
	        //TODO: if extensions are more than 63
	    }
        } else {
            s.set_bits(0, 1)
	}
    }
    s.set_bits(uint(val), bits)
}
func (s *Stream) get_openlen() int {
    c_ind := s.c_ind
    v := s.parse_len(0)
    fmt.Printf("get_openlen %d len : %d\n", c_ind, v)
    return v

}
func (s *Stream) set_openlen(loc uint32, _size int, gtpPresent bool) {
    //oind := s.c_ind
    //obit := s.c_bit
    c_ind := int(loc>>4)
    fmt.Printf("set_openlen %d: len:%d\n", c_ind,  s.c_ind-c_ind)
    s.set_len(uint32(c_ind))
    return
    //s.c_bit = int(loc&0xf)
    //sz := oind - s.c_ind - 1
    //if obit > s.c_bit { sz += 1}
    //s.format_len(sz, 0)
    //s.set_bits(uint(sz), _size)
    //s.c_ind = oind
}

type ObjTypeOf interface{
    Unpack(st *Stream )
    Pack(st *Stream)
}

// MarshalJSON converts the byte slice to a hex string for JSON.
func (h HexBytes) MarshalJSON() ([]byte, error) {
	hexStr := hex.EncodeToString(h)
	return json.Marshal(hexStr)
}
// UnmarshalJSON converts a hex string from JSON back into a byte slice.
func (h *HexBytes) UnmarshalJSON(data []byte) error {
	var hexStr string
	if err := json.Unmarshal(data, &hexStr); err != nil {
		return err
	}

	decodedBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return err
	}

	*h = decodedBytes
	return nil
}
