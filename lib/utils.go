package lib

import (
  "strconv"
)

// convert to lower case character
func tolower(x uint8) uint8 {
  if x >= 'A' && x <= 'Z' {
    return x - 'A' + 'a'
  }
  return x
}

// hex to 4 bits
func hex_to_uint4(x uint8) uint8 {
  x = tolower(x)
  if x >= '0' && x <= '9' {
    return x - '0'
  } else if x >= 'a' && x <= 'f' {
    return x - 'a' + 10
  }
  return 0
}

// 4 bits to hex
func uint4_to_hex(x uint8) uint8 {
  const hex = "0123456789abcdef"
  return hex[x & 15]
}

func Bin2Hex(b []uint8) string {
  s := make([]uint8, len(b) << 1)
  for i := 0; i < len(b); i ++ {
    s[i*2] = uint4_to_hex(b[i] >> 4)
    s[(i*2) + 1] = uint4_to_hex(b[i] & 15)
  }
  return string(s)
}

func Hex2Bin(s string) []byte {
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
func uint6_to_base64(x int) uint8 {
  const base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
  return base64[x & 63]
}

// base64 to 6 bits
func base64_to_uint6(x uint8) uint8 {
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

func Bin2Base64(b []uint8) string {

  var s []uint8

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
      s = append(s, uint6_to_base64(x))
      j = 0
    }
  }

  if j != 0 {
    x <<= (24 - (uint8(j) * 8))
    s = append(s, uint6_to_base64(x >> 18))
    s = append(s, uint6_to_base64(x >> 12))
    s = append(s, uint6_to_base64(x >> 6))
    s = append(s, uint6_to_base64(x))
  }

  return string(s)
}

// xor binary streams
func Xor(a []byte, b []byte) []byte {
  c := make([]byte, len(a))
  for i := 0; i < len(a); i ++ {
    c[i] = a[i] ^ b[i]
  }
  return c
}

func Bin2Ascii(b []uint8) string {
  s := make([]byte, len(b))
  for i := 0; i < len(b); i ++ {
    if strconv.IsPrint(rune(b[i])) {
      s[i] = b[i]
    } else {
      s[i] = '.'
    }
  }
  return string(s)
}

func English_Score(a []byte) float32 {

  english_letter_frequency := map[byte]float32 {
    ' ': 18.74,
    'e': 9.60,
    't': 7.02,
    'a': 6.21,
    'o': 5.84,
    'i': 5.22,
    'n': 5.21,
    'h': 4.87,
    's': 4.77,
    'r': 4.43,
    'd': 3.52,
    'l': 3.20,
    'u': 2.25,
    'm': 1.94,
    'c': 1.88,
    'w': 1.82,
    'g': 1.66,
    'f': 1.62,
    'y': 1.56,
    'p': 1.31,
    ',': 1.24,
    '.': 1.21,
    'b': 1.19,
    'k': 0.74,
    'v': 0.71,
    '"': 0.67,
    '\'': 0.44,
    '-': 0.26,
    '?': 0.12,
    'x': 0.12,
    'j': 0.12,
    ';': 0.08,
    '!': 0.08,
    'q': 0.07,
    'z': 0.07,
    ':': 0.03,
    '1': 0.02,
    '0': 0.01,
    ')': 0.01,
    '*': 0.01,
    '(': 0.01,
    '2': 0.01,
    '`': 0.01,
    '"': 0.01,
    '3': 0.01,
    '9': 0.01,
    '5': 0.01,
    '4': 0.01,
  }

  histogram := map[byte]float32 {}
  for i := 0; i < len(a); i ++ {
    c := tolower(a[i])
    histogram[c] += 1.0
  }
  for k, _ := range histogram {
    histogram[k] /= float32(len(a))
  }

  var score float32 = 0.0
  for k, _ := range histogram {
    score += histogram[k] * english_letter_frequency[k]
  }
  return score
}
