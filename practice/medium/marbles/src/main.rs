// SPDX-FileCopyrightText: 2022 Google Inc
// SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
//
// SPDX-License-Identifier: Apache-2.0

fn marbles(n: i64, k: i64) -> i64 {
    if n == k {
        return 1;
    }

    // first we pick one of each.

    // The rest is n-k multicombination.
    // https://en.wikipedia.org/wiki/Combination#Number_of_combinations_with_repetition

    let nn = n - k;

    // Need at least one of each - no choice here.
    let mut ret: i64 = 1;
    for i in 1..k {
        ret = ret * (nn + i) / i;
    }
    return ret;
}

fn main() {
    use std::io::{self, BufRead};

    let stdin = io::stdin();
    let mut cin_iterator = stdin.lock().lines();

    // Read number of cases.
    // It is actually not used, we simply read until EOF later.
    let buffer = cin_iterator.next().unwrap().unwrap();
    let _: i32 = buffer.trim().parse().expect("input was not an integer");

    // Read cases from stdin, line by line.
    for line in cin_iterator {
        // Shove them into vec of i32
        let c = line
            .unwrap()
            .split_whitespace()
            .map(|x| x.parse::<i64>())
            .collect::<Result<Vec<i64>, _>>()
            .unwrap();

        let n = c[0];
        let k = c[1];
        println!("{}", marbles(n, k));
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(marbles(10, 10), 1);
        assert_eq!(marbles(30, 7), 475020);
    }
}
