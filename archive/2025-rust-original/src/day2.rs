use std::collections::btree_map::Keys;

pub fn main() {
    let input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
    let mut total: i64 = 0;

    input.split(',').for_each(|line| {
        let (num_one, num_two) = line.split_once('-').unwrap();
        let num_one = num_one.parse::<i64>().unwrap();
        let num_two = num_two.parse::<i64>().unwrap();
        total += sum_invalid_ids(num_one, num_two);
    });


    println!("Total: {}", total);

    ()
}

pub fn sum_invalid_ids(num_one: i64, num_two: i64) -> i64 {
    let mut total = 0;

    for i in num_one..=num_two {
        let id = i.to_string();
        if is_invalid_id_part_two(id) {
            total += i;
        }
    }

    return total;
}

pub fn is_invalid_id_part_two(id: String) -> bool {
    let midpoint = id.len() / 2;

    for i in 0..midpoint {
        let substring: String = id.chars().take(i + 1).collect();
        let mut new_string = String::new();

        while new_string.len() < id.len() {
            new_string.push_str(&substring);
        }

        if new_string == id {
            return true;
        }
    }

    false
}

pub fn is_invalid_id_part_one(id: String) -> bool {
    if id.len() % 2 != 0 {
        return false;
    }

    let mut midpoint = id.len() / 2;
    let mut start = 0;

    while midpoint < id.len() {
        if id.chars().nth(start) != id.chars().nth(midpoint) {
            return false;
        }

        start += 1;
        midpoint += 1;
    }

    true
}
