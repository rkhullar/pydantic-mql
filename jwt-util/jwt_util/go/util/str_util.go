package util

func StrPtr(data string) *string {
	return &data
}

func StrEnc(data string) ByteSlice {
	return ByteSlice(data)
}

func B64StrEnc(data string) string {
	return B64Enc(StrEnc(data))
}
