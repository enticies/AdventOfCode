use std::collections::{HashMap, HashSet};

type CoordinateYX = (usize, usize);

pub fn main() {
    let real_input = "

";

    let example_input = "
7,1
8,1
8,4
9,4
9,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
";

    let mut coordinates: Vec<CoordinateYX> = vec![];

    let input = example_input;

    input.trim().lines().for_each(|line| {
        let (y, x) = line.trim().split_once(',').unwrap();

        coordinates.push((y.parse::<usize>().unwrap(), x.parse::<usize>().unwrap()));
    });

    part_two(&mut coordinates);
}

pub fn part_two(coordinates: &mut Vec<CoordinateYX>) {
    let max_y = coordinates.iter().map(|&(y, _)| y).max().unwrap();
    let max_x = coordinates.iter().map(|&(_, x)| x).max().unwrap();

    println!("Coordinates: {:?}", coordinates);
    println!("Max Y: {}, Max X: {}", max_y, max_x);

    let mut largest_area = 0;

    let borders_vec = calculate_borders(coordinates);
    let borders: HashSet<CoordinateYX> = borders_vec.iter().cloned().collect();
    println!("Borders ({} total)", borders.len());

    let mut borders_by_row: HashMap<usize, Vec<usize>> = HashMap::new();
    let mut borders_by_col: HashMap<usize, Vec<usize>> = HashMap::new();

    for &(y, x) in &borders_vec {
        borders_by_row.entry(y).or_insert_with(Vec::new).push(x);
        borders_by_col.entry(x).or_insert_with(Vec::new).push(y);
    }

    for (_, v) in borders_by_row.iter_mut() {
        v.sort();
    }
    for (_, v) in borders_by_col.iter_mut() {
        v.sort();
    }

    let mut checked_count = 0;
    let total_pairs = coordinates.len() * (coordinates.len() - 1) / 2;
    println!("Checking {} rectangle pairs...", total_pairs);

    let mut inside_borders_cache: HashMap<CoordinateYX, bool> = HashMap::new();

    let start_time = std::time::Instant::now();
    let mut last_progress_time = start_time;

    for i in 0..coordinates.len() {
        if i % 10 == 0 && i > 0 {
            let elapsed = start_time.elapsed();
            let progress = (i as f64 / coordinates.len() as f64) * 100.0;
            println!(
                "Outer loop: {}/{} ({:.1}%) | Total checked: {} | Elapsed: {:.2}s",
                i,
                coordinates.len(),
                progress,
                checked_count,
                elapsed.as_secs_f64()
            );
        }

        for j in i + 1..coordinates.len() {
            checked_count += 1;

            let time_since_last = last_progress_time.elapsed();
            if checked_count % 100 == 0 || time_since_last.as_secs() >= 2 {
                let elapsed = start_time.elapsed();
                let progress = (checked_count as f64 / total_pairs as f64) * 100.0;
                let rate = checked_count as f64 / elapsed.as_secs_f64();
                let remaining = if rate > 0.0 {
                    (total_pairs - checked_count) as f64 / rate
                } else {
                    0.0
                };

                println!(
                    "Progress: {}/{} ({:.1}%) | Rate: {:.0}/s | Elapsed: {:.1}s | ETA: {:.1}s",
                    checked_count,
                    total_pairs,
                    progress,
                    rate,
                    elapsed.as_secs_f64(),
                    remaining
                );
                last_progress_time = std::time::Instant::now();
            }

            let coord_i = coordinates[i];
            let coord_j = coordinates[j];
            let corner1 = (coord_i.0, coord_j.1);
            let corner2 = (coord_j.0, coord_i.1);

            // coord_i and coord_j are input coordinates (always valid as rectangle corners)
            // We need to check if ALL coordinates inside the rectangle are inside the borders
            let rect_min_y = coord_i.0.min(coord_j.0);
            let rect_max_y = coord_i.0.max(coord_j.0);
            let rect_min_x = coord_i.1.min(coord_j.1);
            let rect_max_x = coord_i.1.max(coord_j.1);

            let is_valid_rectangle = is_rectangle_valid(
                &borders,
                rect_min_y,
                rect_max_y,
                rect_min_x,
                rect_max_x,
                &borders_by_row,
                &borders_by_col,
                &mut inside_borders_cache,
            );

            if is_valid_rectangle {
                let area = calculate_area(coordinates[i], coordinates[j]);
                if checked_count <= 10 || area > 1000000 {
                    println!(
                        "Found valid rectangle: ({:?}, {:?}) corners: {:?}, {:?} with area {}",
                        coord_i, coord_j, corner1, corner2, area
                    );
                    print_grid_with_rectangle(
                        &borders,
                        coordinates[i],
                        coordinates[j],
                        corner1,
                        corner2,
                        max_y,
                        max_x,
                    );
                }

                if area > largest_area {
                    largest_area = area;
                }
            }
        }
    }

    let total_elapsed = start_time.elapsed();
    println!(
        "\nCompleted checking {} rectangle pairs in {:.2}s",
        checked_count,
        total_elapsed.as_secs_f64()
    );
    println!(
        "Cache size: {} coordinates cached",
        inside_borders_cache.len()
    );
    println!("Largest area: {}", largest_area);
}

pub fn is_rectangle_valid(
    borders_set: &HashSet<CoordinateYX>,
    min_y: usize,
    max_y: usize,
    min_x: usize,
    max_x: usize,
    borders_by_row: &HashMap<usize, Vec<usize>>,
    borders_by_col: &HashMap<usize, Vec<usize>>,
    cache: &mut HashMap<CoordinateYX, bool>,
) -> bool {
    for x in min_x..=max_x {
        if !is_inside_borders_cached(
            borders_set,
            (min_y, x),
            borders_by_row,
            borders_by_col,
            cache,
        ) {
            return false;
        }
    }

    for x in min_x..=max_x {
        if !is_inside_borders_cached(
            borders_set,
            (max_y, x),
            borders_by_row,
            borders_by_col,
            cache,
        ) {
            return false;
        }
    }

    for y in min_y + 1..max_y {
        if !is_inside_borders_cached(
            borders_set,
            (y, min_x),
            borders_by_row,
            borders_by_col,
            cache,
        ) {
            return false;
        }
    }

    for y in min_y + 1..max_y {
        if !is_inside_borders_cached(
            borders_set,
            (y, max_x),
            borders_by_row,
            borders_by_col,
            cache,
        ) {
            return false;
        }
    }

    true
}

pub fn is_inside_borders_cached(
    borders_set: &HashSet<CoordinateYX>,
    coordinate: CoordinateYX,
    borders_by_row: &HashMap<usize, Vec<usize>>,
    borders_by_col: &HashMap<usize, Vec<usize>>,
    cache: &mut HashMap<CoordinateYX, bool>,
) -> bool {
    if let Some(&cached_result) = cache.get(&coordinate) {
        return cached_result;
    }

    let result = is_inside_borders(borders_set, coordinate, borders_by_row, borders_by_col);

    cache.insert(coordinate, result);

    result
}

pub fn is_inside_borders(
    borders_set: &HashSet<CoordinateYX>,
    coordinate: CoordinateYX,
    borders_by_row: &HashMap<usize, Vec<usize>>,
    borders_by_col: &HashMap<usize, Vec<usize>>,
) -> bool {
    if borders_set.contains(&coordinate) {
        return true;
    }

    let x_coords = match borders_by_row.get(&coordinate.0) {
        Some(coords) => coords,
        None => return false, 
    };

    if !x_coords.iter().any(|&x| x < coordinate.1) {
        return false;
    }

    if !x_coords.iter().any(|&x| x > coordinate.1) {
        return false;
    }

    let y_coords = match borders_by_col.get(&coordinate.1) {
        Some(coords) => coords,
        None => return false, 
    };

    if !y_coords.iter().any(|&y| y < coordinate.0) {
        return false;
    }

    if !y_coords.iter().any(|&y| y > coordinate.0) {
        return false;
    }

    true
}

pub fn calculate_borders(coordinates: &Vec<CoordinateYX>) -> Vec<CoordinateYX> {
    let mut result: Vec<CoordinateYX> = Vec::new();

    for coordinate in coordinates {
        let horizontal_border_coordinate = coordinates
            .iter()
            .find(|c| c.1 == coordinate.1 && *c != coordinate);
        let vertical_border_coordinate = coordinates
            .iter()
            .find(|c| c.0 == coordinate.0 && *c != coordinate);

        if let Some(horizontal_border_coordinate) = horizontal_border_coordinate {
            let horizontal_start_y = coordinate.0.min(horizontal_border_coordinate.0);
            let horizontal_end_y = coordinate.0.max(horizontal_border_coordinate.0);

            for y in horizontal_start_y..horizontal_end_y + 1 {
                result.push((y, coordinate.1)); 
            }
        }

        if let Some(vertical_border_coordinate) = vertical_border_coordinate {
            let vertical_start_x = coordinate.1.min(vertical_border_coordinate.1);
            let vertical_end_x = coordinate.1.max(vertical_border_coordinate.1);

            for x in vertical_start_x..vertical_end_x + 1 {
                result.push((coordinate.0, x)); 
            }
        }
    }

    for &coord in coordinates {
        result.push(coord);
    }

    result
}

#[allow(dead_code)]
pub fn part_one(coordinates: &Vec<CoordinateYX>) {
    let mut largest_area = 0;

    for i in 0..coordinates.len() {
        for j in i + 1..coordinates.len() {
            if i != j {
                let area = calculate_area(coordinates[i], coordinates[j]);

                if area > largest_area {
                    largest_area = area;
                }
            }
        }
    }

    println!("Largest area: {}", largest_area);
}

pub fn calculate_area(coordinate_one: CoordinateYX, coordinate_two: CoordinateYX) -> u128 {
    ((coordinate_one.0 as i128 - coordinate_two.0 as i128).abs() + 1) as u128
        * ((coordinate_one.1 as i128 - coordinate_two.1 as i128).abs() + 1) as u128
}

pub fn print_grid_with_rectangle(
    borders: &HashSet<CoordinateYX>,
    corner1: CoordinateYX,
    corner2: CoordinateYX,
    computed_corner1: CoordinateYX,
    computed_corner2: CoordinateYX,
    max_y: usize,
    max_x: usize,
) {
    let min_y = 0;
    let min_x = 0;

    let rect_min_y = corner1.0.min(corner2.0);
    let rect_max_y = corner1.0.max(corner2.0);
    let rect_min_x = corner1.1.min(corner2.1);
    let rect_max_x = corner1.1.max(corner2.1);

    println!("\nGrid visualization:");
    println!("Legend: # = rectangle corner, X = border, * = rectangle interior, . = empty");
    print!("   ");
    for x in min_x..=max_x.min(min_x + 50) {
        print!("{}", x % 10);
    }
    println!();

    for y in min_y..=max_y.min(min_y + 30) {
        print!("{:2} ", y);
        for x in min_x..=max_x.min(min_x + 50) {
            let coord = (y, x);

            if coord == corner1
                || coord == corner2
                || coord == computed_corner1
                || coord == computed_corner2
            {
                print!("#");
            }
            else if y >= rect_min_y && y <= rect_max_y && x >= rect_min_x && x <= rect_max_x {
                if borders.contains(&coord) {
                    print!("X");
                } else {
                    print!("*");
                }
            }
            else if borders.contains(&coord) {
                print!("X");
            } else {
                print!(".");
            }
        }
        println!();
    }
    println!();
}
