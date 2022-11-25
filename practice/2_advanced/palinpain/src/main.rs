// SPDX-FileCopyrightText: 2022 Google Inc
// SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
//
// SPDX-License-Identifier: Apache-2.0

fn palin_pain(x: i64, y: i64) -> Vec<String> {
    if (x%2 == 1) && (y%2 == 1) {
        // else both are odd - palindrome not possible.
        // Note: this is not ideomaic. Should probably return Result.
        return vec!["-1".to_string()];
    }
    if (x==1) || (y==1) {
        // Only one palindrome possible.
        return vec!["-1".to_string()];
    }

    let mut s1: Vec<char> = Vec::with_capacity((x+y) as usize);
    let mut s2: Vec<char> = Vec::with_capacity((x+y) as usize);
    // Fill a's and b's of the left part
    for _ in 0..x/2 {
        s1.push('a');
    }
    for _ in 0..y/2 {
        s1.push('b');
    }
    for _ in 0..y/2 {
        s2.push('b');
    }
    for _ in 0..x/2 {
        s2.push('a');
    }

    // X is odd - add extra 'A' in the middle.
    if x%2==1 {
        s1.push('a');
        s2.push('a');
    }
    if y%2==1 {
        s1.push('b');
        s2.push('b');
    }

    // Now, mirror to the right side.
    // This is the same as first 4 loops above but in different order
    for _ in 0..y/2 {
        s1.push('b');
    }
    for _ in 0..x/2 {
        s1.push('a');
    }
    for _ in 0..x/2 {
        s2.push('a');
    }
    for _ in 0..y/2 {
        s2.push('b');
    }

    return vec![
        s1.into_iter().collect(),
        s2.into_iter().collect()
    ];
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

        let x = c[0];
        let y = c[1];
        for l in palin_pain(x, y) {
            println!("{}", l);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn t1() {
        assert_eq!(palin_pain(2, 4), vec!["abbbba".to_string(), "bbaabb".to_string()]);
    }

    #[test]
    fn t2() {
        assert_eq!(palin_pain(3, 3), vec!["-1".to_string()]);
    }

    #[test]
    fn t_ones() {
        assert_eq!(palin_pain(1, 3), vec!["-1".to_string()]);
        assert_eq!(palin_pain(3, 1), vec!["-1".to_string()]);
        assert_eq!(palin_pain(1, 1), vec!["-1".to_string()]);
    }
}
