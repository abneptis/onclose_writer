package onclose_writer

import (
  "bytes"
  "io"
  "os"
)


type lenreadwriter interface {
  io.ReadWriter
  Len()(int)
}

type oncloseWriter struct {
  buff    lenreadwriter
  onClose func(r io.Reader, rlen int64)(os.Error)
}

func (self *oncloseWriter)Write(in []byte)(n int, err os.Error){
  return self.buff.Write(in)
}

func (self *oncloseWriter)Close()(err os.Error){
  return self.onClose(self.buff, int64(self.buff.Len()))
}


// Returns a write-closer whose results (and length) will be fed to
// the named function at the close of the returned descriptor.
//
// You may specify any 'bytes.Buffer' like object that offers Read/Write/Len functions
// or leave it nil, and a default bytes.NewBuffer() will be created
func New(buff lenreadwriter, f func(r io.Reader, len int64)(err os.Error))(io.WriteCloser){
  if buff == nil {
    buff = bytes.NewBuffer(nil)
  }
  return &oncloseWriter{buff:buff, onClose: f}
}

