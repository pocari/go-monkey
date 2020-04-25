sample


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
```