// Code generated by go generate; DO NOT EDIT.
// This file was generated by stream_pipen_gen.go at 2021-06-18 18:30:33.831323 +0000 UTC
package flow

func (s *Stream) Consume(lane int, k interface{}) {
	_ = s.PipeN(lane, 0, 0, k)
}

func (s *Stream) Pipe(lane, chBuf int, k interface{}) *Stream {
	outs := s.PipeN(lane, chBuf, 1, k)
	return outs[0]
}

func (s *Stream) Pipe2(lane, chBuf int, k interface{}) (*Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 2, k)
	return outs[0], outs[1]
}

func (s *Stream) Pipe3(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 3, k)
	return outs[0], outs[1], outs[2]
}

func (s *Stream) Pipe4(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 4, k)
	return outs[0], outs[1], outs[2], outs[3]
}

func (s *Stream) Pipe5(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 5, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4]
}

func (s *Stream) Pipe6(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 6, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5]
}

func (s *Stream) Pipe7(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 7, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5], outs[6]
}

func (s *Stream) Pipe8(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 8, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5], outs[6], outs[7]
}

func (s *Stream) Pipe9(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 9, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5], outs[6], outs[7], outs[8]
}

func (s *Stream) Pipe10(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 10, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5], outs[6], outs[7], outs[8], outs[9]
}

func (s *Stream) Pipe11(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 11, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5], outs[6], outs[7], outs[8], outs[9], outs[10]
}

func (s *Stream) Pipe12(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 12, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5], outs[6], outs[7], outs[8], outs[9], outs[10], outs[11]
}

func (s *Stream) Pipe13(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 13, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5], outs[6], outs[7], outs[8], outs[9], outs[10], outs[11], outs[12]
}

func (s *Stream) Pipe14(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 14, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5], outs[6], outs[7], outs[8], outs[9], outs[10], outs[11], outs[12], outs[13]
}

func (s *Stream) Pipe15(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 15, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5], outs[6], outs[7], outs[8], outs[9], outs[10], outs[11], outs[12], outs[13], outs[14]
}

func (s *Stream) Pipe16(lane, chBuf int, k interface{}) (*Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream, *Stream) {
	outs := s.PipeN(lane, chBuf, 16, k)
	return outs[0], outs[1], outs[2], outs[3], outs[4], outs[5], outs[6], outs[7], outs[8], outs[9], outs[10], outs[11], outs[12], outs[13], outs[14], outs[15]
}
