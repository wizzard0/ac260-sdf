import { expect, test } from "bun:test";
import { sdfOneBit9chars } from "./sdf-9chars.ts";

// why: the ID clock is the caller's input, never a hidden Date.now()
test("README examples", () => {
  expect(sdfOneBit9chars(Date.UTC(2020, 2, 4, 5, 6, 7))).toBe("p0331t0vm");
  expect(sdfOneBit9chars(Date.UTC(2020, 2, 4, 5, 6, 7, 2))).toBe("p0331t0vn");
});

test("0 is the epoch, not the current time", () => {
  expect(sdfOneBit9chars(0)).toBe("k01000000");
});

test("a missing or non-finite clock throws", () => {
  for (const bad of [undefined, null, NaN, Infinity, "1"]) {
    expect(() => sdfOneBit9chars(bad as unknown as number)).toThrow(TypeError);
  }
});
