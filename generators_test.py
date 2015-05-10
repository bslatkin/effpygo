#!/usr/bin/env python3

# Copyright 2015 Brett Slatkin, Pearson Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import timeit
import unittest

from generators import *


data = '1.0,2.5\n3.5,4.1\n7.5,2.2\n6.9,1.1\n'


class Test(unittest.TestCase):

    def test_single_load(self):
        try:
            rows = load_csv_data(io.StringIO(data))
        except (ValueError, IOError):
            raise Exception('Broke reading CSV')

        for i, row in enumerate(rows):
            print('Row %d is %r' % (i, row))

    def test_streaming(self):
        stream = io.StringIO(data)
        it = load_csv_data_streaming(stream)
        try:
            for i, row in enumerate(it):
                print('Row %d is %r' % (i, row))
        except (ValueError, IOError):
            raise Exception('Broke reading CSV')

    def test_explicit_looping(self):
        stream = io.StringIO(data)
        it = enumerate(load_csv_data_streaming(stream))
        while True:
            try:
                i, row = next(it)
            except StopIteration:
                break
            except (ValueError, IOError):
                raise Exception('Broke after row')
            else:
                print('Row %d is %r' % (i, row))

    def test_distance(self):
        stream = io.StringIO(data)
        it = load_csv_data_streaming(stream)
        for i, distance in enumerate(distance_stream(it)):
            print('Move %d was %f far' % (i, distance))


SETUP = """
import io
import generators
data = "1.5,2.5\\n" * 100000
"""


class Benchmark(unittest.TestCase):

    def test_single_load(self):
        delay = timeit.timeit(
stmt="""
stream = io.StringIO(data)
generators.load_csv_data(stream)
""",
setup=SETUP,
            number=20)
        print('test_single_load: %f per call' % delay)

    def test_streaming(self):
        delay = timeit.timeit(
stmt="""
stream = io.StringIO(data)
for _ in generators.load_csv_data_streaming(stream):
    pass
""",
setup=SETUP,
            number=20)
        print('test_streaming: %f per call' % delay)


if __name__ == '__main__':
    unittest.main()
