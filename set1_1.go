package main

import (
  "fmt"
)

// convert to lower case character
func tolower(x byte) byte {
  if x >= 'A' && x <= 'Z' {
    return x - 'A' + 'a'
  }
  return x
}

// hex to 4 bits
func hex_to_uint4(x byte) byte {
  x = tolower(x)
  if x >= '0' && x <= '9' {
    return x - '0'
  } else if x >= 'a' && x <= 'f' {
    return x - 'a' + 10
  }
  return 0
}

// 4 bits to hex
func uint4_to_hex(x int) byte {
  const hex = "0123456789abcdef"
  return hex[x & 15]
}

func bin_to_hex(b []byte) string {
  s := make([]byte, len(b) << 1)
  for i := 0; i < len(b); i ++ {
    s[i*2] = uint4_to_hex(int(b[i] >> 4))
    s[(i*2) + 1] = uint4_to_hex(int(b[i] & 15))
  }
  return string(s)
}

func hex_to_bin(s string) []byte {
  b := make([]byte, (len(s) + 1) >> 1 )
  for i := 0; i < len(s); i++ {
    b[i/2] <<= 4
    b[i/2] += hex_to_uint4(s[i])
  }
  if len(s) % 2 != 0 {
    b[len(b) - 1] <<= 4
  }
  return b
}

// 6 bits to base64
func uint6_to_base64(x int) byte {
  const base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
  return base64[x & 63]
}

// base64 to 6 bits
func base64_to_uint6(x byte) byte {
  if x >= 'A' && x <= 'Z' {
    return x - 'A'
  } else if x >= 'a' && x <= 'z' {
    return x - 'a' + 26
  } else if x >= '0' && x <= '9' {
    return x - '0' + 52
  } else if x == '+' {
    return 62
  } else if x == '/' {
    return 63
  }
  return 0
}

func bin_to_base64(b []byte) string {

  var s []byte
  j := 0
  x := 0

  for i := 0; i < len(b); i ++ {
    x <<= 8
    x += int(b[i])
    j += 1
    if j == 3 {
      s = append(s, uint6_to_base64(x >> 18))
      s = append(s, uint6_to_base64(x >> 12))
      s = append(s, uint6_to_base64(x >> 6))
      s = append(s, uint6_to_base64(x >> 0))
      j = 0
    }
  }

  if j == 1 {
    x <<= 16
    s = append(s, uint6_to_base64(x >> 18))
    s = append(s, uint6_to_base64(x >> 12))
    s = append(s, '=')
    s = append(s, '=')
  } else if j == 2 {
    x <<= 8
    s = append(s, uint6_to_base64(x >> 18))
    s = append(s, uint6_to_base64(x >> 12))
    s = append(s, uint6_to_base64(x >> 6))
    s = append(s, '=')
  }

  return string(s)
}

func base64_to_bin(s string) []byte {

  var b []byte
  x := 0
  j := 0

  for i := 0; i < len(s); i ++ {
    x <<= 6
    x += int(base64_to_uint6(s[i]))
    j += 1

    if j == 4 {
      b = append(b, byte((x >> 16) & 255))
      b = append(b, byte((x >> 8) & 255))
      b = append(b, byte((x >> 0) & 255))
      j = 0
    }
  }

  if j == 1 {
      b = append(b, byte((x << 2) & 255))
  } else if j == 2 {
      b = append(b, byte((x >> 4) & 255))
      b = append(b, byte((x << 4) & 255))
  } else if j == 3 {
      b = append(b, byte((x >> 10) & 255))
      b = append(b, byte((x >> 2) & 255))
      b = append(b, byte((x << 6) & 255))
  }

  return b
}

const input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
const output = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

const in1 = "pleasure." //  Encodes to: cGxlYXN1cmUu
const in2 = "leasure."  //  Encodes to: bGVhc3VyZS4=
const in3 = "easure."   //  Encodes to: ZWFzdXJlLg==
const in4 = "asure."    //  Encodes to:     YXN1cmUu
const in5 = "sure."     //  Encodes to:     c3VyZS4=

func main() {
  var s string
  for _, s = range []string {in1, in2, in3, in4, in5} {
    var x string
    var y []byte
    x = bin_to_base64([]byte(s))
    y = base64_to_bin(x)
    fmt.Printf("base64 %s bin %s\n", x, y)
  }

  s = bin_to_base64(hex_to_bin(input))
  fmt.Println(s == output)
}
