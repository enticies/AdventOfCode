type Range = (u64, u64);

pub fn main() {
    let real_input = "

";

    let example_input = "
3-5
10-14
16-20
12-18

1
5
8
11
17
32
";

    let input = real_input;

    let mut sections = input.trim().split("\n\n");

    let ingredient_id_ranges = sections.next().unwrap();
    let available_ingredient_ids = sections.next().unwrap();

    let mut fresh_ingredient_id_ranges: Vec<Range> = vec![];

    let mut available_ids: Vec<u64> = vec![];

    for id_range in ingredient_id_ranges.split('\n') {
        let (start, end) = id_range.split_once('-').unwrap();

        fresh_ingredient_id_ranges
            .push((start.parse::<u64>().unwrap(), end.parse::<u64>().unwrap()));
    }

    for available_id in available_ingredient_ids.split('\n') {
        available_ids.push(available_id.parse::<u64>().unwrap());
    }

    part_two(fresh_ingredient_id_ranges);
}

pub fn part_two(fresh_ingredient_id_ranges: Vec<Range>) {
    let mut total = 0;

    let ranges = merge_overlapped_ranges(fresh_ingredient_id_ranges);

    println!("Merged ranges: {:?}", ranges);

    for range in ranges {
        total += range.1 - range.0 + 1;
    }

    println!("Total: {}", total);
}

pub fn check_range(
    ranges: &Vec<Range>,
    range_one: Range,
    start: usize,
    skip: &Vec<usize>,
) -> Option<usize> {
    println!(
        "ranges: {:?} range_one: {:?}, start: {}, skip: {:?}",
        ranges, range_one, start, skip
    );

    for j in start + 1..ranges.len() {
        let range_two = ranges[j];

        if !skip.contains(&j) && overlaps(&range_one, &range_two) {
            println!("{:?} overlaps {:?}", range_one, range_two);
            return Some(j);
        }
    }

    None
}

pub fn merge_overlapped_ranges(fresh_ingredient_id_ranges: Vec<Range>) -> Vec<Range> {
    let mut merged_ranges = vec![];

    let mut ranges = fresh_ingredient_id_ranges.clone();

    let mut skip: Vec<usize> = vec![];

    for i in 0..ranges.len() {
        if skip.contains(&i) {
            continue;
        }

        let mut range_one = ranges[i];

        while let Some(found_id_index) = check_range(&ranges, range_one, i, &skip) {

            skip.push(found_id_index);
            skip.push(i);

            let merged = merge_ranges(&range_one, &ranges[found_id_index]);
            println!("Merged: {:?}", merged);
            println!("range_one: {:?}", range_one);
            range_one = merged;
        }

        merged_ranges.push(range_one);
    }

    merged_ranges
}

pub fn merge_ranges(range_one: &Range, range_two: &Range) -> Range {
    let start = if range_one.0 <= range_two.0 {
        range_one.0
    } else {
        range_two.0
    };

    let end = if range_one.1 >= range_two.1 {
        range_one.1
    } else {
        range_two.1
    };

    (start, end)
}

pub fn overlaps(a: &Range, b: &Range) -> bool {
    a.0 <= b.1 && b.0 <= a.1
}

pub fn part_one(fresh_ingredient_id_ranges: &Vec<Range>, ids: Vec<u64>) {
    let mut number_of_fresh_ingredients = 0;

    for id in ids {
        'inner: for range in fresh_ingredient_id_ranges {
            if is_inside_range(range, id) {
                number_of_fresh_ingredients += 1;
                break 'inner;
            }
        }
    }

    println!("{number_of_fresh_ingredients}");
}

pub fn is_inside_range(range: &Range, id: u64) -> bool {
    id >= range.0 && id <= range.1
}
