package set1_3

import (
  "fmt"
  "github.com/deadsy/matasano/lib"
)

const input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func Run() {
  fmt.Println("set1_3")
  i := lib.Hex2Bin(input)
  var hi_s float32 = 0.0
  hi_c := 0
  for c := 0; c < 256; c ++ {
    o := lib.Xor_Byte(i, byte(c))
    s := lib.English_Score(o)
    if s > hi_s {
      hi_s = s
      hi_c = c
    }
  }

  fmt.Printf("0x%02x: %s\n", hi_c, lib.Bin2Ascii(lib.Xor_Byte(i, byte(hi_c))))
}
