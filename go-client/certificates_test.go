/*
 *   Copyright (c) 2023 Intel Corporation
 *   All rights reserved.
 *   SPDX-License-Identifier: BSD-3-Clause
 */
package client

import (
	"net/http"
	"testing"
)

var jwks = `{"keys":[{"alg":"PS384","e":"AQAB","kid":"3fd751f2e0d0f52846c0ecd4972c6e99dfc642051cd339dd9b04381af8c0ddb804514a7a1fee4673ac844fd5db7f15fb","kty":"RSA","n":"qeCH-XC9TqNt8vSF1T5fHTcWyoW6t_TbMCbHh2rvOuaoqpZGNOblVYDmnzkFkrGQwAZ0ra5MrN-PCLxfuodK2OKAYR3sfxx8BiPhfE-rBoAXZLf5-JJRjB34DH8Pm674LX190BVieOmQLiqJafQ0lSArXPQwwRENEgtJr1eAM-wr8o_UhY2_kuQIhu79NPgPor0l5f4jlENNyC_uq84-qg37SCQzNGHEAesdTQIUoDmAMnKaLZfAa4gVIDQn7KZq5PkLM8IuNDoIEq63HkKdOghvB7MTfuX2B9BAYsxmkfoxaUZMG-cV8o2iCe6MxVQUB0zaql1xLo5eSgiKL7vLeJHv_Owv_Vr7PtbwWZe4r5R6RNTABeh7dHyWRfX63EEGJuq2vG67iukxOXgHLvGpdpoC1rhKG9pizffOjzWQsLYV8jxP9b_sM8TsMg9Yq1sa4kRV-2pG39DhjBKgc3Ba3cCiu1GszmXJZ4YPtH30VuPB2e4SlR5VUp9JCDokidLx","x5c":["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"]}]}`

func TestGetAmberCertificates(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/certs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jwks))
	})

	_, err := client.GetAmberCertificates()
	if err != nil {
		t.Errorf("GetAmberCertificates returned unexpected error: %v", err)
	}
}
