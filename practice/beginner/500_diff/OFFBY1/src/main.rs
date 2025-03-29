// SPDX-FileCopyrightText: 2025 Vladimir Rusinov <vladimir.rusinov@gmail.com>
// SPDX-License-Identifier: Apache-2.0

// Off By One
// https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/OFFBY1

fn main() {
    let mut line = String::new();
    std::io::stdin().read_line(&mut line).unwrap();
    let parts = line.trim().split_whitespace().collect::<Vec<&str>>();
    let a: i32 = parts[0].parse().unwrap();
    let b: i32 = parts[1].parse().unwrap();
    println!("{}", (a + b) * 10 + 1)
}
