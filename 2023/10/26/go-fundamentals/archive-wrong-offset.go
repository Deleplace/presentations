var archive []byte
var p int

// ...

offset := readBinaryUint64(archive[p : p+4])
size := readBinaryUint64(archive[offset : offset+4]) // HL
data := archive[offset+4 : offset+4+size]