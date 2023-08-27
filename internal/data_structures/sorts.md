# Insertion Sort Simulation

Starting with the initial array `[7, 3, 5, 1, 9, 2]`:

**Step 1:**

- `i = 1`, `key = 3`
- Compare `key` (3) with `array[0]` (7). Since 3 < 7, shift 7 to the right.
- Array: `[7, 7, 5, 1, 9, 2]`
- Insert `key` (3) at index 0.
- Array: `[3, 7, 5, 1, 9, 2]`

**Step 2:**

- `i = 2`, `key = 5`
- Compare `key` (5) with `array[1]` (7). Since 5 < 7, no shift is needed.
- Insert `key` (5) at index 1.
- Array: `[3, 5, 7, 1, 9, 2]`

**Step 3:**

- `i = 3`, `key = 1`
- Compare `key` (1) with `array[2]` (7). Since 1 < 7, shift 7 to the right.
- Compare `key` (1) with `array[1]` (5). Since 1 < 5, shift 5 to the right.
- Compare `key` (1) with `array[0]` (3). Since 1 < 3, shift 3 to the right.
- Array: `[3, 5, 7, 7, 9, 2]`
- Insert `key` (1) at index 0.
- Array: `[1, 3, 5, 7, 9, 2]`

**Step 4:**

- `i = 4`, `key = 9`
- Compare `key` (9) with `array[3]` (7). Since 9 > 7, no shift is needed.
- Insert `key` (9) at index 4.
- Array: `[1, 3, 5, 7, 9, 2]`

**Step 5:**

- `i = 5`, `key = 2`
- Compare `key` (2) with `array[4]` (9). Since 2 < 9, shift 9 to the right.
- Compare `key` (2) with `array[3]` (7). Since 2 < 7, shift 7 to the right.
- Compare `key` (2) with `array[2]` (5). Since 2 < 5, shift 5 to the right.
- Compare `key` (2) with `array[1]` (3). Since 2 > 3, insert `key` (2) at index 1.
- Array: `[1, 2, 3, 5, 7, 9]`

The final sorted array is `[1, 2, 3, 5, 7, 9]`, and the Insertion Sort algorithm has successfully sorted the input array.
