pub fn main() {
    let mut total = 0;

    let real_input = "
";

    let example_input = "
987654321111111
811111111111119
234234234234278
818181911112111
    ";

    let inp = example_input;

    for (index, line) in inp.trim().lines().enumerate() {
        let total_lines = inp.trim().lines().count();
        println!("{} out of {} lines left", total_lines - index, total_lines);

        let num = part_two(line.trim());
        println!("Adding: {}", num);
        total += num;
    }

    println!("Total: {}", total);
}

fn find_index_to_remove(s: &Vec<char>) -> usize {
    for index in 0..s.len() {
        if index == s.len() - 1 {
            return index;
        } else if s[index].to_digit(10).unwrap() < s[index + 1].to_digit(10).unwrap() {
            return index;
        }
    }

    return 0;
}

pub fn part_two(bank: &str) -> i64{
    let mut bank_vec: Vec<char> = bank.chars().into_iter().collect();

    for _ in 0..bank.len() - 12 {
        bank_vec.remove(find_index_to_remove(&bank_vec));
    }

    let number_str: String = bank_vec.iter().collect();
    number_str.parse::<i64>().unwrap()
}

pub fn part_one(bank: &str) -> u32 {
    let mut max_num = 0;

    for (i, digit_one) in bank.chars().enumerate() {
        for (j, digit_two) in bank.chars().enumerate() {
            if i < j {
                let num = (10 * digit_one.to_digit(10).unwrap()) + digit_two.to_digit(10).unwrap();

                if num > max_num {
                    max_num = num;
                }
            }
        }
    }

    max_num
}
