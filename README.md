# Advent of Code

My Advent of Code solutions.

## Layout

Every year is a live workspace with the same Rust + JS tooling, so I can solve any
year — current or past — the same way:

```
.
├── 2020/  2021/  2023/  2024/  2025/     # each: rust/ + js/, ready to solve
│   ├── rust/          # Rust — fspoettel/advent-of-code-rust template
│   └── js/            # JavaScript — aocrunner
├── scripts/
│   └── new-year.sh    # stamp out any other year (2015–2025) on demand
└── archive/           # my original solutions for past years (reference, untouched)
    ├── 2020-python/
    ├── 2021-python/
    ├── 2023-python/
    ├── 2024-go/
    └── 2025-rust-original/   # original hand-rolled 2025 Rust project (days 1–11)
```

Puzzle inputs are **git-ignored** in every year's projects (AoC asks that inputs not be
redistributed). Paste them in manually — no session token required.

### Adding another year

The five years above are ready to go. For any other year (e.g. 2019):

```bash
./scripts/new-year.sh 2019     # creates 2019/rust + 2019/js, configured for 2019
```

## Rust — `<year>/rust`

Uses the [fspoettel](https://github.com/fspoettel/advent-of-code-rust) template.

```bash
cd 2024/rust               # or any year
cargo scaffold 3            # create src/bin/03.rs + data/{inputs,examples}/03.txt
# paste the example into data/examples/03.txt, the real input into data/inputs/03.txt
cargo test --bin 03        # run the example-based unit tests
cargo solve 3              # run against real input, prints answer + per-part timing
cargo solve 3 --release    # optimized timing
cargo all                  # run every day
cargo time                 # benchmark all days, write a timings table into the README
```

Each day lives in `src/bin/NN.rs` with `part_one`/`part_two` and a `tests` module that
asserts against `data/examples/NN.txt`. Expected example answers go in the `assert_eq!`s.

**Optimizing:** `cargo solve N --release` gives per-part wall-clock time. For rigorous
statistical benchmarks, add [`criterion`](https://github.com/bheisler/criterion.rs) or a
`[profile.dhat]` heap profile (the template already defines a `dhat-heap` feature).

## JavaScript — `<year>/js`

Uses [aocrunner](https://github.com/caderek/aocrunner).

```bash
cd 2024/js                 # or any year
npm start 3                # scaffold src/day03/ and run it in watch mode
# paste the real input into src/day03/input.txt
```

Each `src/dayNN/index.js` exports `part1`/`part2` and inline `tests` (`{ input, expected }`)
that run automatically before the real input, with per-part timing printed each save.

```bash
npm run update:readme      # write a solved-days + timing table into the JS README
npm run format             # prettier
```

Auto-download of inputs is possible by putting your AoC session cookie in a year's
`js/.env` (`AOC_SESSION_KEY=…`) and the Rust aoc-cli config — left disabled by default.
