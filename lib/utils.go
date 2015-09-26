package lib

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
