// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: Copyright 2025 Vladimir Rusinov

use std::io::{stdin};

fn main() {
    let mut line = String::new();
    stdin().read_line(&mut line).expect("No input");
    let t = line.trim().parse::<u32>().unwrap();
    for _ in 0..t {
        line.clear();
        stdin().read_line(&mut line).expect("No input");
        let x = line.trim().parse::<u32>().unwrap();
        if x >= 30 {
            println!("YES");
        } else {
            println!("NO");
        }
    }
}
