package set1_1

import (
  "fmt"
  "github.com/deadsy/matasano/lib"
)

const input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
const output = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

func Run() {
  fmt.Println("set1_1")

  x := lib.Hex2Bin(input)
  fmt.Printf("%s\n", lib.Bin2Hex(x))

  y := lib.Bin2Base64(x)
  fmt.Printf("%s\n", y)

  fmt.Println(y == output)
}
