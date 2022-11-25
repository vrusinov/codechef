// SPDX-FileCopyrightText: 2022 Google Inc
// SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
//
// SPDX-License-Identifier: Apache-2.0

fn main() {
    use std::io::{stdin};
    let mut s = String::new();
    stdin().read_line(&mut s).expect("No input");
    print!("{}", s);
}
