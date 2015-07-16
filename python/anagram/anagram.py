#!/usr/bin/env python
# -*- coding: utf-8 -*-

def detect_anagrams(word, anagrams):
    anagram_list = []
    for a in anagrams:
        if word == a:
            # same word we break
            break
        # Lowercase, sort string and join without spaces
        elif ''.join(sorted(word.lower())) == ''.join(sorted(a.lower())):
            anagram_list.append(a)

    return anagram_list
