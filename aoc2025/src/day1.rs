pub fn main() {
    let input = "
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
                ";

    let mut current_point = 50;
    let mut total_click_count = 0;

    for line in input.trim().lines() {
        let mut line_chars = line.trim().chars();

        let direction = line_chars.next().unwrap();
        let distance = line_chars.as_str().parse::<i32>().ok().unwrap();

        if direction == 'L' {
            let (new_point, click_count) = foo(current_point, distance, direction);
            total_click_count += click_count;
            current_point = new_point;
        } else {
            let (new_point, click_count) = foo(current_point, distance, direction);
            total_click_count += click_count;
            current_point = new_point;
        }

        println!(
            "direction: {}, distance: {}, current point: {},  click count {}",
            direction, distance, current_point, total_click_count
        );
    }

    println!("{}", total_click_count);
}

pub fn foo(current_point: i32, distance: i32, direction: char) -> (i32, i32) {
    let mut new_point = current_point;

    let mut i = 0;
    let mut click_count = 0;

    while i < distance {
        if direction == 'L' && new_point == 0 {
            new_point = 100;
        } else if direction == 'R' && new_point == 100 {
            new_point = 0;
        }

        if direction == 'L' {
            new_point -= 1;
        } else {
            new_point += 1;
        }

        if new_point == 0 || new_point == 100 {
            click_count += 1;
        }

        i += 1;
    }

    (new_point, click_count)
}
