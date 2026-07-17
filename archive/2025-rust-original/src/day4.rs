type Grid = Vec<Vec<char>>;

const NEIGHBOR_OFFSETS: [(isize, isize); 8] = [
    (-1, -1),
    (0, -1),
    (1, -1),
    (-1, 0),
    (1, 0),
    (-1, 1),
    (0, 1),
    (1, 1),
];

pub fn main() {
    let real_input = "
    ";

    let example_input = "
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
";

    let input = example_input;

    let mut grid: Grid = vec![];

    input.trim().lines().for_each(|line| {
        grid.push(line.trim().chars().collect());
    });

    part_two(&mut grid);
}

pub fn find_accessible_rolls(grid: &Grid) -> Vec<(usize, usize)> {
    let mut marked_rolls = vec![];

    for (row_index, row) in grid.iter().enumerate() {
        for (column_index, _) in row.iter().enumerate() {
            if grid[row_index][column_index] == '@' {
                let neighbour_paper_count =
                    count_paper(grid, column_index as isize, row_index as isize);
                let can_access = neighbour_paper_count < 4;
                if can_access {
                    marked_rolls.push((row_index, column_index));
                }
            }
        }
    }

    marked_rolls
}

pub fn part_two(grid: &mut Grid) {
    let mut total = 0;

    loop {
        let marked_rolls = find_accessible_rolls(grid);
        if marked_rolls.is_empty() {
            break;
        }

        total += marked_rolls.len();

        remove_rolls(grid, marked_rolls);
    }

    println!("{total}");
}

pub fn remove_rolls(grid: &mut Grid, rolls_to_remove: Vec<(usize, usize)>) {
    for roll in rolls_to_remove {
        grid[roll.0][roll.1] = '.';
    }
}

pub fn part_one(grid: &mut Grid) {
    let mut total = 0;

    for (row_index, row) in grid.iter().enumerate() {
        for (column_index, _) in row.iter().enumerate() {
            if grid[row_index][column_index] == '@' {
                let neighbour_paper_count =
                    count_paper(grid, column_index as isize, row_index as isize);
                let can_access = neighbour_paper_count < 4;
                if can_access {
                    total += 1;
                }
            }
        }
    }
    println!("{total}");
}

pub fn count_paper(grid: &Grid, col: isize, row: isize) -> u32 {
    let mut number_of_rolls_of_paper = 0;
    let grid_height = grid.len();
    let grid_width = grid[0].len();

    for (d_col, d_row) in NEIGHBOR_OFFSETS {
        let neighbor_col = d_col + col;
        let neighbor_row = d_row + row;

        if neighbor_col >= 0
            && neighbor_col < grid_width as isize
            && neighbor_row >= 0
            && neighbor_row < grid_height as isize 
            && grid[neighbor_row as usize][neighbor_col as usize] == '@' {
                number_of_rolls_of_paper += 1;
            }
    }

    number_of_rolls_of_paper
}

pub fn print_grid(grid: &Grid) {
    for row in grid.iter() {
        println!("{row:?}");
    }
}
