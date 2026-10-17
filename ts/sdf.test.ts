import { expect, test } from "bun:test";
import { getSDF, parseSDF } from "./sdf.ts";

const nop = () => {};
const at = (y: number) => Date.UTC(y, 5, 15);
const yearOf = (code: string) => new Date(parseSDF(code, nop)).getUTCFullYear();

// why: decade and century codes must decode to the period they were made from
test.each([
  [1000, "B0", 1000], [1250, "B2", 1200], [1799, "B7", 1700],
  [1800, "C0", 1800], [1850, "C5", 1850], [1899, "C9", 1890],
  [2100, "X0", 2100], [2150, "X5", 2150], [2199, "X9", 2190],
  [2200, "Y2", 2200], [2750, "Y7", 2700], [2999, "Y9", 2900],
])("year %i formats as %s and parses back to %i", (year, code, start) => {
  expect(getSDF(at(year))).toBe(code);
  expect(yearOf(code)).toBe(start);
});

// why: the formatter emits A0 and Z9, so the parser must accept them
test.each<[string, number]>([["A0", -Infinity], ["Z9", Infinity], ["a", -Infinity], ["z", Infinity]])(
  "%s parses to %d", (code, want) => {
    expect(parseSDF(code, nop)).toBe(want);
  });

test("years outside 1000-2999 round-trip through A0 and Z9", () => {
  expect(parseSDF(getSDF(at(999)) as string, nop)).toBe(-Infinity);
  expect(parseSDF(getSDF(at(3000)) as string, nop)).toBe(Infinity);
});
