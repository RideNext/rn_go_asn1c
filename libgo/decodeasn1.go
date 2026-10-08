package e2ap
import(
   "fmt"
   "encoding/hex"
   )

type Stream struct {
    buff []byte
    c_ind int
    c_bit int
}

type BIT_STRING struct {
  Value []byte
  Len int
}
type OCTET_STRING struct {
  Value []byte
  Len int
}
type INTEGER struct {
  Value int
}
type PRINTABLE_STRING struct {
   Value string
}
type ENUMERATED struct {
   Value int
}

func (self *Stream) Init(buff []byte) {
   self.buff = buff
   self.c_ind = 0
   self.c_bit = 0
}

type StreamIF interface {
   Initialize()
   get_location() int
   set_location(int, int)
   reset_bits()
   read_bit() bool
   parse_ext() bool
   get_bits(bits int) int
   get_rbits(bits int) int
   get_byte() int
   get_listsize(count int) int
   get_choice(bits int) int
   get_flags(bits int) int
   parse_len() int
   parsef_OctString(size int) []byte
   parsef_BitString(values int, size int) []byte
   parsef_Integer(values int, size int, bits int, ext int) int
   parsef_Enumerated(size int, ext int) int
   parsef_PriString(size int) string
   parse_blen(bits int, ext int) int
   parse_olen(size int) int
   parse_Null()
   handle_unkown_ext(int, int, int)
   get_current_location() (int)
   read_flags(int) int 
   write_flags(int, int)
}

func (s *Stream) read_flags(int) (int){
    return 0
}
func (s *Stream) write_flags(int, int){
}
func (s *Stream) get_current_location()( int){
    return 0
}

func (s *Stream) handle_unkown_ext( cnt int,  flg int , ext int) {
}

func (s *Stream)  Initialize (buf []byte) {
    s.buff = buf
    s.c_ind = 0
    s.c_bit = 0
}

func (s *Stream) get_location() int {
    fmt.Println("get_location", s.c_ind)
    return s.c_ind
}

func (s *Stream) set_location(loc int, ch int){
    s.c_ind = loc+ch
    s.c_bit = 0
    fmt.Println("set_location", s.c_ind, loc, ch)
}

func (s *Stream)  reset_bits() {
    if s.c_bit > 0 {
       s.c_ind ++
       s.c_bit = 0
    }
}

func (s *Stream)  parse_ext() bool {
    return s.read_bit()
}

func (s *Stream)  read_bit() bool {
    v := (s.buff[s.c_ind] & (0x1<<uint8(7-s.c_bit))) > 0
    s.c_bit += 1
    if s.c_bit & 8 > 0 {
        s.c_bit = 0
        s.c_ind ++
    }
    return v
}

func (s *Stream)  get_bits(bits int) int {
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
    v := 0
    for i := 0; i < bits; i ++ {
        if s.read_bit() {
            v |= (1 << uint(i))
        }
    }
    return v
}

func (s *Stream)  get_byte() int {
    i := int(s.buff[s.c_ind])
    s.c_ind ++
    return i
}

func (s *Stream)  parse_len(size int) int {
    s.reset_bits()
    v := 0
    fmt.Println("parse_len", size, s.c_ind)
    if size > 8 && size <= 16 {
        v = s.get_byte()
        v <<= 8
        v |= s.get_byte()
        return v
    }
    fl := s.read_bit()
    v = s.get_bits(7)
    if fl {
        v <<= 8
        v |= s.get_byte()
    }
    fmt.Println(v, s.c_ind)
    return v
}

func (s *Stream)  get_listsize(values int) int {
    v := 0
    fmt.Println("\nget_listsize", values, s.c_ind)
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
    fmt.Println(v, s.c_ind)
    return v
}

func (s *Stream)  get_choice(bits int) int {
    return s.get_bits(bits)
}

func (s *Stream)  get_flags(bits int) int {
    return s.get_rbits(bits)
}

func (s *Stream)  parsef_OctString(size int) []byte {
    defer func () {
        if err := recover(); err != nil {
            fmt.Printf("Parse Error %s", err)
        }
     }()

    if size == 1 {
       return s.parsef_BitString(8, 8)
    }
    s.reset_bits()
    ind := s.c_ind
    s.c_ind += size
    v := make([]byte, size)
    copy(v, s.buff[ind:ind+size])
    return v
    //return s.buff[ind:ind+size]
}
func (s *Stream)  parsef_BitString(values int, size int) []byte {
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
            v[i] = byte(s.get_bits(8))
            size -= 8
	} else {
            v[i] = byte(s.get_bits(size)<<uint(8-size))
	}
    }
    return v
}
func (s *Stream)  parsef_Integer(values int, size int, ext int, start int) int {
    v := 0
    fmt.Println("\nparsef_Integer", values, size, s.c_ind)
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
        if size <= 32 {
            fl = s.get_bits(2)
	} else if size <= 64 {
            fl = s.get_bits(3)
	} else {
            fl = s.get_bits(4)
            s.reset_bits()
            for fl>=0 {
                v <<= 8
                v |= s.get_byte()
                fl -= 1
	    }
	}
    }
    fmt.Println("\nparsef_Integer", v, s.c_ind)
    return v+start
}

func (s *Stream)  parsef_Enumerated(size int, ext int) int {
    fmt.Println("\nparsef_Enumerated", size, s.c_ind)
    return s.get_bits(size)
}

func (s *Stream)  parsef_PriString(size int) string {
    s.reset_bits()
    ind := s.c_ind
    s.c_ind += size
    v := make([]byte, size)
    copy(v, s.buff[ind:ind+size])
    return string(v)
    //return string(s.buff[ind:ind+size])
}

func (s *Stream)  parse_blen(bits int, ext int) int {
    return s.get_bits(bits)
}

func (s *Stream)  parse_olen(bits int) int {
    if bits > 0 && bits <= 8 {
        return s.get_bits(bits)
    } else {
        return s.parse_len(bits)
    }
    return 0
}

func (s *Stream)  parse_Null() {
    return
}

type HexBytes []byte

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
