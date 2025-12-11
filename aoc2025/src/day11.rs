use std::{
    collections::{HashMap, HashSet},
    hash::Hash,
};

pub fn main() {
    let real_input = "

";

    let example_input = "
aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
";

    let mut connections: HashMap<String, Vec<String>> = HashMap::new();

    let input = example_input;

    input.trim().lines().for_each(|line| {
        let parts: Vec<&str> = line.trim().split(' ').collect();
        let key = parts[0].strip_suffix(':').unwrap().to_string();

        let rest_of_parts: Vec<String> = parts.iter().skip(1).map(|s| s.to_string()).collect();
        connections.insert(key, rest_of_parts);
    });

    part_one(connections);
}

pub fn part_one(connections: HashMap<String, Vec<String>>) {
    let mut current = connections.get("you").unwrap();
    let mut already_visited: HashSet<String> = HashSet::new();

    let total = visit_key(connections, already_visited, "you".to_string());

    println!("Total: {}", total);
}

pub fn visit_key(
    connections: HashMap<String, Vec<String>>,
    already_visited: HashSet<String>,
    key: String,
) -> u128 {
    println!();
    println!("Already visited: {:?}", already_visited);
    println!("key: {:?}", key);
    println!();

    if key == "out" {
        return 1;
    }

    if already_visited.contains(&key) {
        return 0;
    }

    let mut total = 0;

    let mut visited = already_visited.clone();
    visited.insert(key.clone());

    for value in connections.get(&key).unwrap() {
        total += visit_key(connections.clone(), visited.clone(), value.clone());
    }

    total 
}
