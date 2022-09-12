var archive []byte
var p int

// ...

if p+4 > len(archive) { // HL
	return fmt.Errorf("corrupt archive: incomplete file offset at %d", p)
}
offset := readBinaryUint64(archive[p : p+4])

if offset+4 > len(archive) { // HL
	return fmt.Errorf("corrupt archive: invalid offset %d", offset)
}
size := readBinaryUint64(archive[offset : offset+4])

if offset+4+size > len(archive) { // HL
	return fmt.Errorf("corrupt archive: invalid size %d starting at %d", size, offset+4)
}
data := archive[offset+4 : offset+4+size]