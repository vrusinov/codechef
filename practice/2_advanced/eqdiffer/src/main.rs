// Unfortunately external crates are not supported so I could not use counter.
//use counter::Counter;
use std::collections::HashMap;

fn eqdiffer(a: Vec<i32>) -> i32 {
    //println!("vec: {:?}", a);
    if a.len() <= 2 {
        return 0;
    }
    // Since we need difference between _each_ pair to be the same,
    // and all numbers are positive, they all must be the same.
    // Thus, we delete all but the most popular element.

    // Collect frequencies of all elements.
    //let mut frequencies = a.iter().collect::<Counter<_, i32>>();
    let mut frequencies = HashMap::new();
    for i in &a {
        //frequencies.entry(i).and_modify(|n| *n += 1).or_insert(1);
        let n = frequencies.entry(i).or_insert(0);
        *n += 1;
    }

    if frequencies.len() == 1 {
        // They are all the same already
        return 0;
    }

    //let most_common = frequencies.k_most_common_ordered(1)[0].0;
    //println!("most common: {}", most_common);
    let mut most_common = -1;
    let mut most_common_count = 0;
    for (value, count) in &frequencies {
        if count > &most_common_count {
            most_common = **value;
            most_common_count = *count;
        }
    }
    frequencies.remove(&most_common);
    // for (k, v) in &frequencies {
    //     println!("{}: {}", k, v);
    // }
    let to_remove = &frequencies.values().sum();
    if a.len() as i32 - to_remove < 2 {
        return *to_remove - 1;
    } else {
        return *to_remove;
    }
}

fn main() {
    use std::io::{self, BufRead};

    let stdin = io::stdin();
    let mut cin_iterator = stdin.lock().lines();

    // Read number of cases.
    let buffer = cin_iterator.next().unwrap().unwrap();
    let t: i32 = buffer.trim().parse().expect("input was not an integer");

    for _ in 0..t {
        // Read N - not used.
        let buffer = cin_iterator.next().unwrap().unwrap();
        let _: i32 = buffer.trim().parse().expect("input was not an integer");

        // Read vector of ints:
        let line = cin_iterator.next().unwrap().unwrap();
        let a = line
            .split_whitespace()
            .map(|x| x.parse::<i32>())
            .collect::<Result<Vec<i32>, _>>()
            .unwrap();
            println!("{}", eqdiffer(a));
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn t_all_equal() {
        assert_eq!(eqdiffer(vec![2, 2, 2]), 0);
    }

    #[test]
    fn t1() {
        assert_eq!(eqdiffer(vec![1, 2, 3]), 1);
        assert_eq!(eqdiffer(vec![1, 1, 1, 3, 3]), 2);
        assert_eq!(eqdiffer(vec![1, 1, 1, 1, 3]), 1);
    }

    #[test]
    fn t_short() {
        assert_eq!(eqdiffer(vec![1, 2]), 0);
    }

}