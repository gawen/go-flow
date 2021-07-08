package flow

// Just start one action, without input neither output.
// Like `flow.Consume`, without inputs.
func (f *Flow) Do(cb interface{}) {
	f.Consume(1, cb)
}
