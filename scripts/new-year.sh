#!/usr/bin/env bash
#
# Stamp out a fresh Advent of Code year with the same Rust + JS tooling as 2025.
#
#   ./scripts/new-year.sh 2019
#
# Creates <year>/rust (fspoettel template) and <year>/js (aocrunner), both
# configured for that year's puzzles. The 2025/ folders are used as the base,
# so no network is needed.
set -euo pipefail

YEAR="${1:-}"
if ! [[ "$YEAR" =~ ^20[0-9]{2}$ ]]; then
  echo "usage: $0 <year>   (e.g. 2019)" >&2
  exit 1
fi

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

if [[ ! -d 2025/rust || ! -d 2025/js ]]; then
  echo "error: 2025/rust and 2025/js (the base templates) must exist" >&2
  exit 1
fi
if [[ -e "$YEAR" ]]; then
  echo "error: $YEAR already exists — remove it first if you mean to recreate it" >&2
  exit 1
fi

echo "Stamping $YEAR ..."
mkdir -p "$YEAR"

# ---- Rust -------------------------------------------------------------------
cp -r 2025/rust "$YEAR/rust"
rm -rf "$YEAR/rust/target"                       # regenerable build artifacts
rm -f  "$YEAR/rust/src/bin/"[0-9]*.rs            # no days yet
rm -f  "$YEAR/rust/data/inputs/"*.txt "$YEAR/rust/data/examples/"*.txt "$YEAR/rust/data/puzzles/"*.txt 2>/dev/null || true
# point the toolchain at this year's puzzles
sed -i "s/^AOC_YEAR = .*/AOC_YEAR = \"$YEAR\"/" "$YEAR/rust/.cargo/config.toml"

# ---- JavaScript -------------------------------------------------------------
cp -r 2025/js "$YEAR/js"
rm -rf "$YEAR/js/src/day"[0-9]*                  # no days yet
# rewrite year in the runner config and package name (node_modules copied as-is)
node -e '
  const fs = require("fs");
  const year = process.argv[1], dir = process.argv[2];
  const cfgPath = dir + "/.aocrunner.json";
  const cfg = JSON.parse(fs.readFileSync(cfgPath, "utf8"));
  cfg.year = Number(year);
  fs.writeFileSync(cfgPath, JSON.stringify(cfg, null, 2) + "\n");
  const pkgPath = dir + "/package.json";
  const pkg = JSON.parse(fs.readFileSync(pkgPath, "utf8"));
  pkg.name = "aoc" + year;
  pkg.description = "Advent of Code " + year + " - solutions";
  fs.writeFileSync(pkgPath, JSON.stringify(pkg, null, 2) + "\n");
' "$YEAR" "$YEAR/js"

echo "Done."
echo "  Rust: cd $YEAR/rust && cargo scaffold <day>"
echo "  JS:   cd $YEAR/js   && npm start <day>"
