package main

import (
  "fmt"
)

func set1_1() {
  const input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
  const output = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
  s := bin_to_base64(hex_to_bin(input))
  fmt.Println(s == output)
}

func set1_2() {
  const p_text = "1c0111001f010100061a024b53535009181c"
  const k_text = "686974207468652062756c6c277320657965"
  const c_text = "746865206b696420646f6e277420706c6179"
  s := bin_to_hex(xor_bin(hex_to_bin(p_text), hex_to_bin(k_text)))
  fmt.Println(s == c_text)
}

func main() {
  set1_1()
  set1_2()
}
