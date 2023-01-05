// SPDX-FileCopyrightText: 2023 Vladimir Rusinov <vrusinov@google.com>
//
// SPDX-License-Identifier: Apache-2.0
//
// https://www.codechef.com/problems/AVG

fn average_number(n: i32, k: i32, v: i32, a: Vec<i32>) -> i32 {
    let total_len = n + k;
    let sum_of_missing: i32 = total_len * v - a.iter().sum::<i32>();

    if sum_of_missing < 1 {
        return -1;
    }

    if sum_of_missing % k != 0 {
        return -1;
    }
    return sum_of_missing / k;
}

fn main() {
    use std::io::{self, BufRead};

    let stdin = io::stdin();
    let mut cin_iterator = stdin.lock().lines();

    // Read number of cases.
    let buffer = cin_iterator.next().unwrap().unwrap();
    let t: i32 = buffer.trim().parse().expect("input was not an integer");

    // Read cases from stdin, line by line.
    for _ in 0..t {
        // Read vector of ints - for N, K and V.
        let line = cin_iterator.next().unwrap().unwrap();
        let c = line
            .split_whitespace()
            .map(|x| x.parse::<i32>())
            .collect::<Result<Vec<i32>, _>>()
            .unwrap();
        let n = c[0];
        let k = c[1];
        let v = c[2];
        // Read another vector of ints:
        let line = cin_iterator.next().unwrap().unwrap();
        let a = line
            .split_whitespace()
            .map(|x| x.parse::<i32>())
            .collect::<Result<Vec<i32>, _>>()
            .unwrap();
        println!("{}", average_number(n, k, v, a));
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn t1() {
        // Tests from example
        assert_eq!(average_number(3, 3, 4, vec![2, 7, 3]), 4);
        assert_eq!(average_number(3, 1, 4, vec![7, 6, 5]), -1);
        assert_eq!(average_number(3, 3, 4, vec![2, 8, 3]), -1);
    }
}