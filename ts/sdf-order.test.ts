import { expect, test } from "bun:test";
import { getSDF, parseSDF } from "./sdf.ts";

// why: SDF promises that codes sort like the dates they encode, in every range
// given: random dates from year -5000 to 5000 and random forward offsets
// when: x and x + y are both formatted and parsed back
// then: neither the code nor its decoded start ever goes backwards

const nop = () => {};
const MIN = Date.UTC(-5000, 0, 1), MAX = Date.UTC(5000, 0, 1);

// mulberry32: a seeded generator, so a failure always reproduces
const rng = (seed: number) => () => {
  seed = (seed + 0x6d2b79f5) | 0;
  let t = Math.imul(seed ^ (seed >>> 15), 1 | seed);
  t = (t + Math.imul(t ^ (t >>> 7), 61 | t)) ^ t;
  return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
};

test("format(x + y) >= format(x) and parse keeps that order", () => {
  const r = rng(260);
  for (let i = 0; i < 20000; i++) {
    const x = MIN + Math.floor(r() * (MAX - MIN));
    const y = Math.floor(r() ** 4 * (MAX - x));
    const a = getSDF(x) as string, b = getSDF(x + y) as string;
    if (!(a <= b)) throw new Error(`code order: ${new Date(x).toISOString()}=${a} > +${y}=${b}`);
    expect(parseSDF(b, nop)).toBeGreaterThanOrEqual(parseSDF(a, nop));
  }
});
