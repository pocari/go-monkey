sample


## repl

```
% go run ./main.go
Hello pocari! This is the Monkey programming language!
Feel free to type in command
>> 1 + 2 * 3
7
>> let a = 1;
>> let b = 2;
>> a + b * 3
7
>> let f = fn(x, y) { x + y };
>> f(1, 2)
3
>> let fact = fn(x) { if (x == 1) { 1 } else { x * fact(x - 1)} };
>> fact(10)
3628800
>> let adder = fn(delta) { fn (x) { x + delta } }
>> let adder_five = adder(5)
>> adder_five(10)
15
>> adder_five(20)
25
>> let array_a = [1, 2, 3]
>> let array_b = push(array_a, 4)
>> array_b
[1, 2, 3, 4]
>> let second = fn(ary) { ary[1] }
>> second(array_b)
2
>> let foo = "hoge"
>> let foo = "foo"
>> let hoge = "hoge"
>> let foo_hoge = foo + "-" + hoge
>> foo_hoge
foo-hoge
>> let people = [{"name": "Alice", "age": 20}, {"name": "Bob", "age": 21}]
>> puts(people[0]["name"])
Alice
null
>> let i = 0
>> while (i < len(people)) { puts(people[i]["age"]); let i = i + 1; }
20
21
null
```


# monkey

```
% cat sample/sum.monkey
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

% go run ./cmd/monkey/main.go sample/sum.monkey
55
```
