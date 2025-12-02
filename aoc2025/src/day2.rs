use std::collections::btree_map::Keys;

pub fn main() {
    let input = "9226466333-9226692707,55432-96230,4151-6365,686836-836582,519296-634281,355894-471980,971626-1037744,25107-44804,15139904-15163735,155452-255998,2093-4136,829776608-829880425,4444385616-4444502989,2208288-2231858,261-399,66-119,91876508-91956018,2828255673-2828317078,312330-341840,6464-10967,5489467-5621638,1-18,426-834,3434321102-3434378477,4865070-4972019,54475091-54592515,147-257,48664376-48836792,45-61,1183-1877,24-43";
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
        if is_invalid_id(id) {
            total += i;
        }
    }

    return total;
}

pub fn is_invalid_id(id: String) -> bool {
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
