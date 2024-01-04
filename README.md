## Monkey programming language

from [Writing an interpreter in Go book ](https://interpreterbook.com/)

[<img src="https://interpreterbook.com/img/monkey_logo-d5171d15.png">](https://interpreterbook.com/img/monkey_logo-d5171d15.png)

```
let five = 5;

let ten = 10;

let add = fn (x, y) {
  x + y;
}

let result = add(five, tne);
```


## Macros 


```
>> let unless = macro(condition, consequence, alternative) { quote(if (!(unquote(condition))) { unquote(consequence); } else { unquote(alternative); }); };
>> unless(10 > 5, puts("not greater"), puts("greater"));
greater
null
```