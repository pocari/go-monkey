let fact = fn(n) {
  if (n == 1) {
    return n
  } else {
    return fact(n - 1) * n
  }
}

let i = 1
while (i < 11) {
  puts(fact(i))
  let i = i + 1
}

