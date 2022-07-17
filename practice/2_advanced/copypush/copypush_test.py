import unittest

import copypush

class CopyPushTest(unittest.TestCase):

    def test_empty(self):
        self.assertEqual(copypush.copypush(''), 'yes')

    def test_example(self):
        cases = [
            ('ab', 'no'),
            ('oof', 'yes'),
            ('aabaab', 'yes'),
            ('eevee', 'no'),
        ]
        for line, expected in cases:
            print(f'Testing {line}')
            self.assertEqual(copypush.copypush(line), expected)

    def test_failed1(self):
        self.assertEqual(copypush.copypush('jjjj'), 'yes')