package set1_2

import (
  "fmt"
  "github.com/deadsy/matasano/lib"
)

const input = "1c0111001f010100061a024b53535009181c"
const xor_bits = "686974207468652062756c6c277320657965"
const output = "746865206b696420646f6e277420706c6179"

func Run() {
  fmt.Println("set1_2")
  o := lib.Bin2Hex(lib.Xor(lib.Hex2Bin(input), lib.Hex2Bin(xor_bits)))
  fmt.Printf("%s\n", o)
  fmt.Println(o == output)
}
