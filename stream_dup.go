package flow

//go:generate go run stream_dupn_gen.go

func (s *Stream) DupN(
	chBuf int,
	numOuts int,
) []*Stream {
	return s.f.DupN(chBuf, numOuts, s)
}
