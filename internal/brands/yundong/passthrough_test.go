// Copyright (C) liasica. 2023-present.
//
// Created at 2023-02-16
// Based on yundong by liasica, magicrolan@qq.com.

package yundong

import (
    "encoding/hex"
    "fmt"
    jsoniter "github.com/json-iterator/go"
    "testing"
)

func getPercentageRoundsSubstring(percentage float64) string {
    symbols := "★★★★★★★★★★☆☆☆☆☆☆☆☆☆☆"
    offset := 10 - int(percentage*10.0)
    return symbols[offset*3 : (offset+10)*3]
}

func TestUnpackPassthroughResponse(t *testing.T) {
    var a, b []byte
    var err error
    // b, _ = hex.DecodeString("7b22746d223a313637363536303037313430362c22636f6465223a302c22726573756c74223a227b5c22654d6574657256616c75655c223a31303737337d227d")
    // b, _ = hex.DecodeString("7b22746d223a313637363536303832353533382c22636f6465223a302c22726573756c74223a225b636f6d2e79756e6675747572652e6e6f736d617274636162696e65742e6e657474792e6d6f64656c2e6a736f6e2e5265737036394031306434646330615d227d")
    // b, _ = hex.DecodeString("7b22746d223a2231363736353632333532222c22636f6465223a307d")
    // b, _ = hex.DecodeString("7b2263223a37352c22746d223a313637363536323335322c22706172616d223a7b22616374696f6e223a322c22636162696e6574736e223a22222c226f70656e646f6f72223a357d7d")
    // b, _ = hex.DecodeString("7b22746d223a2231363736353632333532222c22636f6465223a307d")
    // a, _ = hex.DecodeString("7b2263223a36392c22746d223a313637363536313537332c22706172616d223a7b22616374696f6e223a312c226f70656e646f6f72223a322c227461736b746f6b656e223a22227d7d")
    // b, _ = hex.DecodeString("7b22746d223a313637363536313537333736372c22636f6465223a302c22726573756c74223a225b636f6d2e79756e6675747572652e6e6f736d617274636162696e65742e6e657474792e6d6f64656c2e6a736f6e2e5265737036394032343062666335385d227d")
    // a, _ = hex.DecodeString("7b2263223a37392c22746d223a313637363538383631363132382c22706172616d223a7b22636162696e6574736e223a224e434157444642314e37353153313232227d7d")
    // b, _ = hex.DecodeString("7b22746d223a313637363538383631363238342c22636f6465223a302c22726573756c74223a227b5c22654d6574657256616c75655c223a31303737347d227d")
    // a, _ = hex.DecodeString("7b2263223a36392c22746d223a313637363536313537332c22706172616d223a7b22616374696f6e223a312c226f70656e646f6f72223a322c227461736b746f6b656e223a22227d7d")
    // b, _ = hex.DecodeString("7b22746d223a313637363536313537333736372c22636f6465223a302c22726573756c74223a225b636f6d2e79756e6675747572652e6e6f736d617274636162696e65742e6e657474792e6d6f64656c2e6a736f6e2e5265737036394032343062666335385d227d")
    // b, _ = hex.DecodeString("7b22746d223a2231363736363334353330313632222c22636f6465223a307d")
    b, _ = hex.DecodeString("1f017b22746d223a2231363737313433303333353933222c22636f6465223a307d")
    data := new(PassthroughResponse)
    err = jsoniter.Unmarshal(b, data)
    fmt.Println(a, b, err)
    // 1676562352
    // 1676562352
}

func TestNextPassthroughSN(t *testing.T) {
    for i := 0; i < 300; i++ {
        t.Log(NextPassthroughSN())
    }
}
