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

import csv
import collections
import io
import math


def load_csv_data(stream):
    result = []
    for row in csv.reader(stream):
        if len(row) != 2:
            raise ValueError('Rows must have two entries')
        point = float(row[0]), float(row[1])
        result.append(point)
    return result


def load_csv_data_streaming(stream):
    for row in csv.reader(stream):
        if len(row) != 2:
            raise ValueError('Rows must have two entries')
        point = float(row[0]), float(row[1])
        yield point


def distance_stream(it):
    x, y = next(it)
    for a, b in it:
        yield math.sqrt((a - x) ** 2 + (b - y) ** 2)
        x, y = a, b


def main():
    data = '1.0,2.5\n3.5,4.1\n7.5,2.2\n6.9,1.1\n'

    # Single load
    try:
        rows = load_csv_data(io.StringIO(data))
    except (ValueError, IOError):
        raise Exception('Broke reading CSV')

    for i, row in enumerate(rows):
        print('Row %d is %r' % (i, row))

    # Streaming with a nice loop construct
    stream = io.StringIO(data)
    it = load_csv_data_streaming(stream)
    try:
        for i, row in enumerate(it):
            print('Row %d is %r' % (i, row))
    except (ValueError, IOError):
        raise Exception('Broke reading CSV')

    # Streaming with explicit looping; shows which item was bad
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

    # Stream of the delta
    stream = io.StringIO(data)
    it = load_csv_data_streaming(stream)
    for i, distance in enumerate(distance_stream(it)):
        print('Move %d was %f far' % (i, distance))


if __name__ == '__main__':
    main()
