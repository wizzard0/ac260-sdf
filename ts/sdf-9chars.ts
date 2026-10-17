// Simple Date Format (c) 2020 Oleksandr Nikitin

import {getSDF} from "./sdf.ts";

export function roundToMonthStartUTC(date: Date) {
  return new Date(Date.UTC(date.getUTCFullYear(), date.getUTCMonth(), 1));
}

export function sdfOneBit9chars(now: number): string {
  // returns the 3-char year-month code + base34 of 2ms ticks since the month start (9 chars total)
  // the caller passes the clock: 0 is 1970, not "now"
  if (typeof now !== 'number' || !Number.isFinite(now)) throw new TypeError('sdfOneBit9chars: now must be a finite ms timestamp');
  let crop = +roundToMonthStartUTC(new Date(now))
  let halfDelta = ((now - crop) / 2) | 0;
  let sdfYM = (getSDF(now) + '').slice(0, 3).toLowerCase()
  return sdfYM + halfDelta.toString(34).padStart(6, '0');
}

// todo decoder
