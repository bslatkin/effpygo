#!/usr/bin/env python3

import io
import unittest

from generators import *


ADDRESS = (
    'Four score and seven years ago our fathers brought forth on this '
    'continent a new nation, conceived in liberty, and dedicated to the '
    'proposition that all men are created equal.')

ADDRESS_EXPECTED = [
    (0, 'Four'),
    (5, 'score'),
    (11, 'and'),
    (15, 'seven'),
    (21, 'years'),
    (27, 'ago'),
    (31, 'our'),
    (35, 'fathers'),
    (43, 'brought'),
    (51, 'forth'),
    (57, 'on'),
    (60, 'this'),
    (65, 'continent'),
    (75, 'a'),
    (77, 'new'),
    (81, 'nation'),
    (89, 'conceived'),
    (99, 'in'),
    (102, 'liberty'),
    (111, 'and'),
    (115, 'dedicated'),
    (125, 'to'),
    (128, 'the'),
    (132, 'proposition'),
    (144, 'that'),
    (149, 'all'),
    (153, 'men'),
    (157, 'are'),
    (161, 'created'),
    (169, 'equal'),
]

ADDRESS_WITH_SPACES = (
    '  Four score and seven years ago our fathers brought forth on this   '
    'continent a new nation, conceived in liberty, and dedicated to the '
    'proposition that all men are created equal.  ')

ADDRESS_WITH_SPACES_EXPECTED = [
    (2, 'Four'),
    (7, 'score'),
    (13, 'and'),
    (17, 'seven'),
    (23, 'years'),
    (29, 'ago'),
    (33, 'our'),
    (37, 'fathers'),
    (45, 'brought'),
    (53, 'forth'),
    (59, 'on'),
    (62, 'this'),
    (69, 'continent'),
    (79, 'a'),
    (81, 'new'),
    (85, 'nation'),
    (93, 'conceived'),
    (103, 'in'),
    (106, 'liberty'),
    (115, 'and'),
    (119, 'dedicated'),
    (129, 'to'),
    (132, 'the'),
    (136, 'proposition'),
    (148, 'that'),
    (153, 'all'),
    (157, 'men'),
    (161, 'are'),
    (165, 'created'),
    (173, 'equal'),
]

NO_ENDING_LETTER = 'Four    score and    seven'

NO_ENDING_LETTER_EXPECTED = [
    (0, 'Four'),
    (8, 'score'),
    (14, 'and'),
    (21, 'seven'),
]


class TestIndexWords(unittest.TestCase):

    def testBasic(self):
        result = index_words(ADDRESS)
        self.assertListEqual(ADDRESS_EXPECTED, result)

    def testEdgeCases(self):
        result = index_words(ADDRESS_WITH_SPACES)
        self.assertListEqual(ADDRESS_WITH_SPACES_EXPECTED, result)

    def testNoEndingLetter(self):
        result = index_words(NO_ENDING_LETTER)
        self.assertListEqual(NO_ENDING_LETTER_EXPECTED, result)


class TestIndexWordsGenerator(unittest.TestCase):

    def testBasic(self):
        result = list(index_words_generator(io.StringIO(ADDRESS)))
        self.assertListEqual(ADDRESS_EXPECTED, result)

    def testEdgeCases(self):
        result = list(index_words_generator(io.StringIO(ADDRESS_WITH_SPACES)))
        self.assertListEqual(ADDRESS_WITH_SPACES_EXPECTED, result)

    def testNoEndingLetter(self):
        result = list(index_words_generator(io.StringIO(NO_ENDING_LETTER)))
        self.assertListEqual(NO_ENDING_LETTER_EXPECTED, result)


class TestIndexWordsStream(unittest.TestCase):

    def testBasic(self):
        result = list(index_words_stream(io.StringIO(ADDRESS)))
        self.assertListEqual(ADDRESS_EXPECTED, result)

    def testEdgeCases(self):
        result = list(index_words_stream(io.StringIO(ADDRESS_WITH_SPACES)))
        self.assertListEqual(ADDRESS_WITH_SPACES_EXPECTED, result)

    def testNoEndingLetter(self):
        result = list(index_words_stream(io.StringIO(NO_ENDING_LETTER)))
        self.assertListEqual(NO_ENDING_LETTER_EXPECTED, result)


if __name__ == '__main__':
    unittest.main()
