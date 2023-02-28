package main

/*
Basic Types in Go - https://golangbyexample.com/all-data-types-in-golang-with-examples/
string - string is a basic type, but is a read-only slice of bytes under the hood. Default "".
bool - True/False. accepts and '&&', or '||', negation '!'. Default False.
int - signed, is int64 on 64 bit machines, 32 on 32. Aka 8 bytes/4bytes. All default 0.
int8, int16, int32, int64 - don't use unless you know what you're doing,
	will often be promoted to 'int' and not save any memory+cost time
uint - unsigned version of int. All default 0.
uint8, uint16, uint32, uint64 - as above.
uintptr - stores a pointer, but not linked. corresponding object can be garbage collected. Unsafe.
float32, float64 - 64 inferred. Note there is no "float" w/o a number. Default 0.0
complex64, complex128 - 128 inferred. both parts made of same float type.
byte - alias for uint8. GoLang does NOT have a 'char' datatype. byte for ASCII. and...
rune - alias for uint32. used for Unicode. Unicode can use up to 4 bytes, so uint32.
*/
// func main() {
// }
