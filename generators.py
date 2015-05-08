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


def read_and_buffer(stream):
    while True:
        data = stream.read(1024)
        if not data:
            return
        for char in data:
            yield char


def index_words_generator(stream):
    word_start = 0
    current = ''

    for index, char in enumerate(read_and_buffer(stream)):
        if is_letter(char):
            if not current:
                word_start = index
            current += char
        else:
            if current:
                yield word_start, current
                current = ''

    if current:
        yield word_start, current

