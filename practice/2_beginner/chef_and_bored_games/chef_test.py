#!/usr/bin/python3

# SPDX-FileCopyrightText: 2022 Google Inc
# SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
#
# SPDX-License-Identifier: Apache-2.0

import chef

import unittest


class TestNumSquares(unittest.TestCase):
    def test_example(self):
        self.assertEqual(chef.num_squares(3), 10)
        self.assertEqual(chef.num_squares(8), 120)
        self.assertEqual(chef.num_squares(1), 1)
