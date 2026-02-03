package protocol

// FixedFieldReader provides convenient methods for reading fixed-part fields.
type FixedFieldReader struct {
	buf *Buffer
	err error
}

// NewFixedFieldReader creates a reader for fixed message fields.
func NewFixedFieldReader(data []byte) *FixedFieldReader {
	return &FixedFieldReader{
		buf: NewBuffer(data),
	}
}

// Error returns the first error encountered during reading.
func (r *FixedFieldReader) Error() error {
	return r.err
}

// Remaining returns the number of bytes remaining in the buffer.
func (r *FixedFieldReader) Remaining() int {
	return r.buf.Len()
}

// RemainingBytes returns the remaining bytes in the buffer.
func (r *FixedFieldReader) RemainingBytes() []byte {
	return r.buf.Bytes()
}

// ReadUint8 reads a uint8, storing any error.
func (r *FixedFieldReader) ReadUint8() uint8 {
	if r.err != nil {
		return 0
	}
	val, err := r.buf.ReadUint8()
	if err != nil {
		r.err = err
	}
	return val
}

// ReadInt8 reads an int8, storing any error.
func (r *FixedFieldReader) ReadInt8() int8 {
	if r.err != nil {
		return 0
	}
	val, err := r.buf.ReadInt8()
	if err != nil {
		r.err = err
	}
	return val
}

// ReadUint16 reads a uint16, storing any error.
func (r *FixedFieldReader) ReadUint16() uint16 {
	if r.err != nil {
		return 0
	}
	val, err := r.buf.ReadUint16()
	if err != nil {
		r.err = err
	}
	return val
}

// ReadInt16 reads an int16, storing any error.
func (r *FixedFieldReader) ReadInt16() int16 {
	if r.err != nil {
		return 0
	}
	val, err := r.buf.ReadInt16()
	if err != nil {
		r.err = err
	}
	return val
}

// ReadUint32 reads a uint32, storing any error.
func (r *FixedFieldReader) ReadUint32() uint32 {
	if r.err != nil {
		return 0
	}
	val, err := r.buf.ReadUint32()
	if err != nil {
		r.err = err
	}
	return val
}

// ReadInt32 reads an int32, storing any error.
func (r *FixedFieldReader) ReadInt32() int32 {
	if r.err != nil {
		return 0
	}
	val, err := r.buf.ReadInt32()
	if err != nil {
		r.err = err
	}
	return val
}

// ReadBool reads a boolean (2 bytes per GED-188 spec).
func (r *FixedFieldReader) ReadBool() bool {
	return r.ReadUint16() != 0
}

// ReadFixedString reads a fixed-length string.
func (r *FixedFieldReader) ReadFixedString(n int) string {
	if r.err != nil {
		return ""
	}
	val, err := r.buf.ReadFixedString(n)
	if err != nil {
		r.err = err
	}
	return val
}

// ReadBytes reads n bytes.
func (r *FixedFieldReader) ReadBytes(n int) []byte {
	if r.err != nil {
		return nil
	}
	val, err := r.buf.ReadBytes(n)
	if err != nil {
		r.err = err
	}
	return val
}

// Skip skips n bytes.
func (r *FixedFieldReader) Skip(n int) {
	r.ReadBytes(n)
}

// FixedFieldWriter provides convenient methods for writing fixed-part fields.
type FixedFieldWriter struct {
	buf *Buffer
	err error
}

// NewFixedFieldWriter creates a writer for fixed message fields.
func NewFixedFieldWriter() *FixedFieldWriter {
	return &FixedFieldWriter{
		buf: NewWriteBuffer(),
	}
}

// Error returns the first error encountered during writing.
func (w *FixedFieldWriter) Error() error {
	return w.err
}

// Bytes returns the written bytes.
func (w *FixedFieldWriter) Bytes() []byte {
	return w.buf.Bytes()
}

// WriteUint8 writes a uint8.
func (w *FixedFieldWriter) WriteUint8(v uint8) {
	if w.err != nil {
		return
	}
	w.err = w.buf.WriteUint8(v)
}

// WriteInt8 writes an int8.
func (w *FixedFieldWriter) WriteInt8(v int8) {
	if w.err != nil {
		return
	}
	w.err = w.buf.WriteInt8(v)
}

// WriteUint16 writes a uint16.
func (w *FixedFieldWriter) WriteUint16(v uint16) {
	if w.err != nil {
		return
	}
	w.err = w.buf.WriteUint16(v)
}

// WriteInt16 writes an int16.
func (w *FixedFieldWriter) WriteInt16(v int16) {
	if w.err != nil {
		return
	}
	w.err = w.buf.WriteInt16(v)
}

// WriteUint32 writes a uint32.
func (w *FixedFieldWriter) WriteUint32(v uint32) {
	if w.err != nil {
		return
	}
	w.err = w.buf.WriteUint32(v)
}

// WriteInt32 writes an int32.
func (w *FixedFieldWriter) WriteInt32(v int32) {
	if w.err != nil {
		return
	}
	w.err = w.buf.WriteInt32(v)
}

// WriteBool writes a boolean as 2 bytes per GED-188 spec.
func (w *FixedFieldWriter) WriteBool(v bool) {
	if v {
		w.WriteUint16(1)
	} else {
		w.WriteUint16(0)
	}
}

// WriteFixedString writes a fixed-length null-padded string.
func (w *FixedFieldWriter) WriteFixedString(s string, n int) {
	if w.err != nil {
		return
	}
	w.err = w.buf.WriteFixedString(s, n)
}

// WriteBytes writes raw bytes.
func (w *FixedFieldWriter) WriteBytes(data []byte) {
	if w.err != nil {
		return
	}
	_, w.err = w.buf.Write(data)
}

// WriteZeros writes n zero bytes.
func (w *FixedFieldWriter) WriteZeros(n int) {
	w.WriteBytes(make([]byte, n))
}
