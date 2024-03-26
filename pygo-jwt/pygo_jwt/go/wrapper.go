package main

import (
	lib "github.com/rkhullar/python-libraries/pygo-jwt/pygo_jwt/go/core"
)

/*
#ifndef WRAPPER_UTIL_H
#define WRAPPER_UTIL_H

#include <stdlib.h>
#include <stdbool.h>

typedef struct string_with_error {char* data; char* error;} StringWithError;
typedef struct bool_with_error {bool data; char* error;} BoolWithError;

#endif
*/
import "C"

// TODO: change to `#include "wrapper_util.h"` once supported

//export NewJWK
func NewJWK(size C.int, id *C.char) *C.StringWithError {
	res, err := lib.NewJWK(int(size), TranslateStrPtr(id))
	return HandleStringWithError(res, err)
}

//export JWKToPEM
func JWKToPEM(json_data *C.char) *C.StringWithError {
	res, err := lib.JWKToPEM(C.GoString(json_data))
	return HandleStringWithError(res, err)
}

//export PEMToJWK
func PEMToJWK(pem *C.char, id *C.char) *C.StringWithError {
	res, err := lib.PEMToJWK(C.GoString(pem), TranslateStrPtr(id))
	return HandleStringWithError(res, err)
}

//export ParseJWKAndSign
func ParseJWKAndSign(key *C.char, data *C.char) *C.StringWithError {
	res, err := lib.ParseJWKAndSign(C.GoString(key), C.GoString(data))
	return HandleStringWithError(res, err)
}

//export ParsePEMAndSign
func ParsePEMAndSign(key *C.char, data *C.char) *C.StringWithError {
	res, err := lib.ParsePEMAndSign(C.GoString(key), C.GoString(data))
	return HandleStringWithError(res, err)
}

//export ExtractPublicJWK
func ExtractPublicJWK(key *C.char) *C.StringWithError {
	res, err := lib.ExtractPublicJWK(C.GoString(key))
	return HandleStringWithError(res, err)
}

//export ExtractPublicPEM
func ExtractPublicPEM(key *C.char) *C.StringWithError {
	res, err := lib.ExtractPublicPEM(C.GoString(key))
	return HandleStringWithError(res, err)
}

//export ParsePublicJWKAndVerify
func ParsePublicJWKAndVerify(key *C.char, data *C.char, signature *C.char) *C.BoolWithError {
	res, err := lib.ParsePublicJWKAndVerify(C.GoString(key), C.GoString(data), C.GoString(signature))
	return HandleBoolWithError(res, err)
}

//export ParsePublicPEMAndVerify
func ParsePublicPEMAndVerify(key *C.char, data *C.char, signature *C.char) *C.BoolWithError {
	res, err := lib.ParsePublicPEMAndVerify(C.GoString(key), C.GoString(data), C.GoString(signature))
	return HandleBoolWithError(res, err)
}

//export ExampleGo
func ExampleGo(n C.int) {
	lib.ExampleGo(int(n))
}

//export MaybeError
func MaybeError(n C.int) *C.StringWithError {
	m := int(n)
	res, err := lib.MaybeError(m)
	return HandleStringWithError(res, err)
}

func main() {}
