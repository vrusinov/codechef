// SPDX-FileCopyrightText: 2025 Vladimir Rusinov <vladimir.rusinov@gmail.com>
// SPDX-License-Identifier: Apache-2.0

// 404 Not Found
// https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/ERROR404

fn main() {
    let mut line = String::new();
    std::io::stdin().read_line(&mut line).unwrap();
    let x: i32 = line.trim().parse().unwrap();
    if x == 404 {
        println!("NOT FOUND");
    } else {
        println!("FOUND");
    }
}
