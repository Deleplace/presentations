r, err := regexp.Compile(pattern)
if err != nil {
	return false, fmt.Errorf("matches failed: %v", err) // HL
}