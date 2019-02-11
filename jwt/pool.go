package jwt

import (
	"bytes"
	"encoding/base64"
	"github.com/json-iterator/go"
	"io"
	"reflect"
	"sync"
	"unsafe"
)

var (
	encoderPool = sync.Pool{New: func() interface{} { return new(jwtencoder).init() }}
	decoderPool = sync.Pool{New: func() interface{} { return new(jwtdecoder).init() }}
)

func getEncoder() *jwtencoder { return encoderPool.Get().(*jwtencoder) }
func putEncoder(je *jwtencoder) { encoderPool.Put(je.reset()) }

func getDecoder() *jwtdecoder { return decoderPool.Get().(*jwtdecoder) }
func putDecoder(jd *jwtdecoder) { decoderPool.Put(jd.reset()) }

//

type jwtencoder struct {
	sign    [32]byte
	json    bytes.Buffer
	body    bytes.Buffer
	encoder io.WriteCloser
	jsonenc jsoniter.Encoder
}

func (rcv *jwtencoder) init() *jwtencoder {
	rcv.encoder = base64.NewEncoder(base64.RawURLEncoding, &rcv.body)
	rcv.jsonenc = *json.NewEncoder(&rcv.json)
	return rcv
}

func (rcv *jwtencoder) reset() *jwtencoder {
	rcv.encoder.Close()
	rcv.json.Reset()
	rcv.body.Reset()
	return rcv
}

//

type jwtdecoder struct {
	sign [32]byte
	buf  bytes.Buffer
}

func (rcv *jwtdecoder) init() *jwtdecoder {
	return rcv
}

func (rcv *jwtdecoder) decode(srcBytes []byte) []byte {
	rcv.buf.Reset()

	dl := base64.RawURLEncoding.DecodedLen(len(srcBytes))
	rcv.buf.Grow(dl)

	dst := rcv.buf.Bytes()
	dst = dst[:dl]

	n, _ := base64.RawURLEncoding.Decode(dst, srcBytes)
	return dst[:n]
}

func (rcv *jwtdecoder) reset() *jwtdecoder {
	rcv.buf.Reset()
	return rcv
}

//

func stringToBytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceHeader := reflect.SliceHeader{Data: stringHeader.Data, Len: stringHeader.Len, Cap: stringHeader.Len}
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}
