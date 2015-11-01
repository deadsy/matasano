#! /usr/bin/python

import string

def cleanup1(x):
  # remove any non-printable characters
  x = [c for c in x if c in string.printable]
  # turn all white space into a space
  x = [(c, ' ')[c.isspace()] for c in x]
  # turn repeated spaces into a single space
  y = []
  for i in range(len(x) - 1):
    if x[i] == ' ':
      if x[i + 1] != ' ':
        y.append(' ')
      else:
        # repeated space - skip it
        pass
    else:
      y.append(x[i])
  # add the last character
  y.append(x[-1])
  return ''.join(y)

def cleanup2(x):
  # remove any non-printable characters
  x = [c for c in x if c in string.printable]
  # remove whitespace
  x = [c for c in x if not c.isspace()]
  # capitalise
  x = [c.upper() for c in x]
  return ''.join(x)

def analyze(x):
  """work out the probablility of each bin"""
  n = float(len(x))
  f = [0,] * 256
  for c in x:
    f[ord(c)] += 1
  d = [float(f[i])/n for i in range(256)]
  return d

def distro_nz(d):
  """ordered non-zero elements of the distro"""
  x = []
  for i in range(256):
    if d[i] != 0.0:
      if chr(i) in string.printable:
        name = '\'%c\'' % chr(i)
      else:
        name = '%02x' % i
      x.append((i, name, d[i]))
  return x

def distro_nz_sorted(d):
  """sorted non-zero elements of the distro"""
  x = distro_nz(d)
  x = sorted(x, key = lambda x: x[2], reverse = True)
  return x

def print_distro(x):
  for (i, name, p) in x:
    print('%s: %f' % (name, p))

def main():
  f = open('corpus.txt')
  x = f.read()
  f.close()
  x = cleanup1(x)
  d = analyze(x)
  print_distro(distro_nz_sorted(d))

main()
