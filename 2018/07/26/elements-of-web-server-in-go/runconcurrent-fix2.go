import "sync" // OMIT

// RunConcurrent launches funcs,
// and waits for their completion.
func RunConcurrent(funcs ...func()) { // HLdecl
	var wg sync.WaitGroup // HLwg
	wg.Add(len(funcs))    // HLwg
	for _, f := range funcs {
		g := f // HLtrap
		go func() {
			g()       // HLtrap
			wg.Done() // HLwg
		}()
	}
	wg.Wait() // HLwg
} // HLdecl