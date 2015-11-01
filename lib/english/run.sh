#! /bin/bash

TXT=corpus.txt

rm $TXT
cat dickens-christmas-125.txt >> $TXT
cat wilde-ballad-611.txt >> $TXT
cat kant-critique-142.txt >> $TXT

./english.py
