// Simple Date Format (c) 2020 Oleksandr Nikitin

import {getSDF} from "./sdf.ts";

export function roundToMonthStartUTC(date) {
  return new Date(Date.UTC(date.getUTCFullYear(), date.getUTCMonth(), 1));
}

export function sdfOneBit9chars(now: number): string {
  // returns YMD + base34 of the rest sans 1 bit (9 chars total)
  let crop = +roundToMonthStartUTC(new Date(now))
  let halfDelta = ((now - crop) / 2) | 0;
  let sdfYMD = (getSDF(now) + '').slice(0, 3).toLowerCase()
  return sdfYMD + halfDelta.toString(34).padStart(6, '0');
}
