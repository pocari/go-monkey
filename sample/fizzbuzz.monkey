let fizzBuzz = fn (n) {
  if (n % 15 == 0) {
    return "FizzBuzz"
  } else {
    if (n % 3 == 0) {
      return "Fizz"
    } else {
      if (n % 5 == 0) {
        return "Buzz"
      } else {
        return n
      }
    }
  }
}

let i = 1
while (i < 21) {
  puts(fizzBuzz(i))
  let i = i + 1
}
