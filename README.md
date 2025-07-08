# ac260-sdf
ac260 Simple Date Format

Provides a way to encode any date as a short string, keeps lexicographical ordering.

## Examples

| Example | Precision | Details          |
|---------|-----------|------------------|
| P0      | Year      | 2020             |
| P57     | Month     | July 2025        |
| N3201   | Day       | Feb 1, 2003      |
| P0A12   | Day       | October 12, 2020 |

Also, usable for partially-readable still sortable ID generation.  
The Y-M-D part is exposed directly, the rest is milliseconds with 2ms precision, so the ID fits in the 9 base34 characters total.

## ID examples

| Example   | Date                     |
|-----------|--------------------------|
| p0331t0vm | 2020-03-04T05:06:07Z     |
| p0331t0vn | 2020-03-04T05:06:07.002Z |

## Full semantics

| Notes                  | 1st | Range | 2nd Num | 3rd ch | Example                                  |
|------------------------|-----|-------|---------|--------|------------------------------------------|
| goal: encode yyyy-mm   | A   | -Inf  | ?       | ?      | A0 - Minus Infinity (Before Big Bang?)   |
| human-readable         | B   | 1x__  | Century | ?      | B5 - End of the Middle Ages              |
|                        | C   | 18x_  | Decade  | ?      | С9 - Around year 1890                    |
|                        | D   | 190x  | Year    | Month  |                                          |
| yymm is ambiguous      | E   | 191x  | Year    | Month  |                                          |
| 1 letter yr not enough | F   | 192x  | Year    | Month  |                                          |
|                        | G   | 193x  | Year    | Month  |                                          |
| 200 yrs in 26 letters  | H   | 194x  | Year    | Month  |                                          |
| let's do infinities    | I   | 195x  | Year    | Month  |                                          |
| years 1900-2100 D-W    | J   | 196x  | Year    | Month  |                                          |
| Letter Number [LN]     | K   | 197x  | Year    | Month  |                                          |
| /[a-z]\d[abc\d]/i      | L   | 198x  | Year    | Month  |                                          |
| Months 1-9ABC          | M   | 199x  | Year    | Month  | M49 - Sep 1994, System Shock released    |
|                        | N   | 200x  | Year    | Month  |                                          |
|                        | O   | 201x  | Year    | Month  | O9C - December 2019                      |
|                        | P   | 202x  | Year    | Month  | P01 - January 2020                       |
|                        | Q   | 203x  | Year    | Month  | Q81 - End of Unix Epoch (approx)         |
|                        | R   | 204x  | Year    | Month  |                                          |
|                        | S   | 205x  | Year    | Month  |                                          |
|                        | T   | 206x  | Year    | Month  |                                          |
|                        | U   | 207x  | Year    | Month  | U2 - 2072, where the System Shock is set |
|                        | V   | 208x  | Year    | Month  |                                          |
|                        | W   | 209x  | Year    | Month  |                                          |
|                        | X   | 21x_  | Decade  | ?      |                                          |
|                        | Y   | 2x__  | Century | ?      |                                          |
|                        | Z   | +Inf  | ?       | ?      | Z9 - Infinity (Universe Thermal Death?)  |
