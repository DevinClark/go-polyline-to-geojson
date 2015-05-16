package main

import "testing"
// import "log"

func BenchmarkDecodePolyline1(b *testing.B) {
  for n := 0; n < b.N; n++ {
    decodePolyline([]byte("gyvuE|tuqQ"))
  }
}

func BenchmarkDecodePolyline100(b *testing.B) {
  for n := 0; n < b.N; n++ {
    decodePolyline([]byte("gyvuE|tuqQOrDOvEm@hBeAdAaB^sCAgBI{Bi@kCeAmBiCWqAq@oA{AMgBJeBn@{AbA_ArAw@xBWfBGhC@fBNvBZvAb@`AZ`@RRLLJJDHJJDBB@B@B?B?@?@?B?B?b@?XCLGFEJMFFL?HA?Y?W?S?a@?SA[?U?[AY@[?_@?a@?OL?F?L?TAJ?N?R?R?P?T?L?R?V?X?ZAXA\\?`@B^A`@?pADjCH|BCjB@vGB`EE|DAlAIv@@ZTT^BZc@tCk@jE??u@tGoAnPi@tPQtIs@hKqBjKmCrHsCrG_ExEyD|DaIrI}GpH?{N@eLEiJIoJ@eLEwKBcH@iJCuHEgKDsEE_C"))
  }
}
