let sumArray = fn (ary) {
  let sum = 0;
  let i = 0;

  while (i < len(ary)) {
    let sum = sum + ary[i]
    let i = i + 1
  }

  return sum
}

let data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

puts(sumArray(data))

