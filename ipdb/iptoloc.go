package ipdb

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*
17mon DB File Format
-------------------
Location Data Offset(4bytes-BigEndian)
-------------------
IP Prefix Index 0 (4 bytes-LittleEndian)
IP Prefix Index 1 (4 bytes-LittleEndian)
.
.
.
IP Prefix Index 255 (4 bytes-LittleEndian)
-------------------
IP Index2 0 (8bytes)
	IP Address (4bytes-BigEndian)
	Location Data Offset in Location Data Area (3bytes-LittleEndian) - 1024
		Location Data Offset is 8-bytes based.
	Location Data Length (1byte)
IP Index2 1 (8bytes)
.
.
.
Location Data Offset Starting Point
-------------------
Location Data Area
-------------------
*/
var ErrInvalidIP = errors.New("the ip string is not a valid ip string")
var ErrNotFound = errors.New("the ip is found in the current ipdb.")

type IPToLocationer interface {
	IPToLocation(ip string) (location string, err error)
}

type ip17monDB struct {
	data                 []byte
	locationDataStartPos uint32
	prefixIndexStartPos  uint32
	ipIndexStartPos      uint32
	max_comp_len         uint32
}

var theDB *ip17monDB

func (db *ip17monDB) Init(dbfilename string) error {

	_, err := os.Stat(dbfilename)
	if err != nil {
		return err
	}

	f, err := os.Open(dbfilename)
	defer f.Close()
	if err != nil {
		return err
	}

	db.data, err = ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	// The First 4 bytes are the offset of the location data area
	db.locationDataStartPos = toU32FromBE(db.data[:4])
	// The next 1024 bytes are the prefix index data for the ip index
	db.prefixIndexStartPos = 4
	// Then the ip index data area offset, the location data index is stored in
	//the corresponding fields.
	db.ipIndexStartPos = db.prefixIndexStartPos + 256*4
	// The ip data index searching routine should not past the boundrary of the
	// following value
	db.max_comp_len = db.locationDataStartPos - db.ipIndexStartPos

	return nil
}

func (db *ip17monDB) IPToLocation(ip string) (location string, err error) {

	ipUint32, err := validateAndConvertIp(ip)
	if err != nil {
		return "", err
	}

	ipPrefix := ipUint32 >> 24
	// Prefix index field is in size 4 bytes , that's why we multiply ipPrefix
	// by 4 when we calculate the ip index data address.
	ipIndex := toU32FromLE(
		db.data[db.prefixIndexStartPos+ipPrefix*4 : db.prefixIndexStartPos+ipPrefix*4+4])

	var startpos uint32

	for i := uint32(0); ; i++ {

		if db.ipIndexStartPos+ipIndex+i*8 > db.max_comp_len {
			return "", ErrNotFound
		}
		// ip index field is in size 8 bytes, that's why we multiply
		// (ipIndex+i) by 8 when we calculate the ip data address.
		startpos = db.ipIndexStartPos + (ipIndex+i)*8
		// The first 4 bytes are ip data
		if ipUint32 <= toU32FromBE(db.data[startpos:startpos+4]) {
			// the original implementation of the ip db substract the *1024*
			// from the location data start position, that's why the *1024*
			// below here.
			locDataOffset := db.locationDataStartPos - 1024 +
				// The next 3 bytes are the location data's offset
				toU32FromLE(db.data[startpos+4:startpos+7])
			return string(
					// The data in the 8th byte is the length of the location data
					db.data[locDataOffset : locDataOffset+uint32(db.data[startpos+7])]),
				nil
		}
	}
}

func IPToLocation(ip string) (location string, err error) {
	if theDB == nil {
		theDB = new(ip17monDB)
		err = theDB.Init("17monipdb.dat")
		if err != nil {
			return "", err
		}
	}
	return theDB.IPToLocation(ip)
}

func toU32FromLE(data []byte) uint32 {
	datalen := len(data)
	if datalen > 4 {
		datalen = 4
	}
	var result uint32
	for i := 0; i < datalen; i++ {
		result = result + uint32(data[i])<<uint32(i*8)
	}
	return result
}

func toU32FromBE(data []byte) uint32 {
	datalen := len(data)
	if datalen > 4 {
		datalen = 4
	}
	var result uint32
	for i := 0; i < datalen; i++ {
		result = result + uint32(data[i])<<uint32((datalen-i-1)*8)
	}
	return result
}

func validateAndConvertIp(ip string) (uint32, error) {

	ips := strings.Split(ip, ".")

	if len(ips) != 4 {
		return 0, ErrInvalidIP
	}

	digits := make([]byte, 4)

	for i, _ := range digits {
		temp, err := strconv.ParseUint(ips[i], 10, 8)
		digits[i] = byte(temp)
		if err != nil {
			return 0, err
		}
		if digits[i] > 0xff {
			return 0, ErrInvalidIP
		}
	}
	return toU32FromBE(digits), nil
}
