package utils

import (
	"bytes"
	"errors"
	"sync"
)

var ErrDataPipeClosed = errors.New("closed")

type DataPipe struct {
	dataCh chan *bytes.Buffer
	pool   sync.Pool
	closed chan bool
}

func NewDataPipe(bufSize int) *DataPipe {
	p := &DataPipe{}
	p.dataCh = make(chan *bytes.Buffer, bufSize)
	p.pool = sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
	p.closed = make(chan bool)
	return p
}

func (dp *DataPipe) Write(data []byte) error {
	buf := dp.pool.Get().(*bytes.Buffer)
	_, _ = buf.Write(data)
	select {
	case <-dp.closed:
		return ErrDataPipeClosed
	case dp.dataCh <- buf:
	}
	return nil
}

func (dp *DataPipe) Read() ([]byte, error) {
	var buf *bytes.Buffer
	select {
	case <-dp.closed:
		return nil, ErrDataPipeClosed
	case buf = <-dp.dataCh:
	}

	defer func() {
		buf.Reset()
		dp.pool.Put(buf)
	}()
	return buf.Bytes(), nil
}
