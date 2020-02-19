import "sync" // OMIT

// RunConcurrent launches funcs,
// and waits for their completion.
func RunConcurrent(funcs ...func()) { // HLdecl
	var wg sync.WaitGroup // HLwg
	wg.Add(len(funcs))    // HLwg
	for _, f := range funcs {
		go func(g func()) { // HLtrap
			g()
			wg.Done() // HLwg
		}(f) // HLtrap
	}
	wg.Wait() // HLwg
} // HLdecl