#!/usr/bin/env python3

import unicodedata
import unittest


def is_letter(char):
    category = unicodedata.category(char)
    if category.startswith('P'):
        return False
    if char.isspace():
        return False
    return True


def index_words(text):
    result = []
    word_start = 0
    current = ''

    for index, char in enumerate(text):
        if is_letter(char):
            if not current:
                word_start = index
            current += char
        else:
            if current:
                result.append((word_start, current))
                current = ''

    if current:
        result.append((word_start, current))

    return result


def index_words_generator(buffer):
    index = 0
    word_start = 0
    current = ''

    while True:
        char = buffer.read(1)
        if not char:
            break

        if is_letter(char):
            if not current:
                word_start = index
            current += char
        else:
            if current:
                yield word_start, current
                current = ''

        index += 1

    if current:
        yield word_start, current

