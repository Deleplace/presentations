import "sync" // OMIT

// RunConcurrent launches funcs,
// and waits for their completion.
func RunConcurrent(funcs ...func()) { // HLdecl
	var wg sync.WaitGroup // HLwg
	wg.Add(len(funcs))    // HLwg
	for _, f := range funcs {
		f := f // HLtrap
		go func() {
			f()       // HLtrap
			wg.Done() // HLwg
		}()
	}
	wg.Wait() // HLwg
} // HLdecl