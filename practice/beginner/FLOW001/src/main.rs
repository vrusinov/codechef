// SPDX-FileCopyrightText: 2022 Google Inc
// SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
//
// SPDX-License-Identifier: Apache-2.0

// Add Two Numbers
//
// https://www.codechef.com/problems/FLOW001

fn add(c: Vec<i32>) -> i32 {
    return c.iter().sum();
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
            .map(|x| x.parse::<i32>())
            .collect::<Result<Vec<i32>, _>>()
            .unwrap();

        let ret = add(c);
        println!("{}", ret);
    }
}
