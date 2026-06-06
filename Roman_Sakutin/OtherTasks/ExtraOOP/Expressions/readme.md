## Домашнее задание 11. Выражения (Go)
### Разработайте структуры Const, Variable, Add, Subtract, Multiply, Divide для вычисления выражений с одной переменной в типе int. Все структуры должны реализовывать интерфейс Expression.
### Структуры должны позволять составлять выражения вида:
### 
### NewSubtract(NewMultiply(NewConst(2), NewVariable("x")),NewConst(3),).Evaluate(5)
### 
### При вычислении такого выражения вместо каждой переменной подставляется значение, переданное в качестве параметра методу Evaluate. Таким образом, результатом вычисления приведенного примера должно стать число 7.
### Требования:
### 1. Подстановка значений переменной
### Метод Evaluate(x int) int должен вычислять выражение с подстановкой значения переменной.
### 2. Метод String
### Должен возвращать выражение в полноскобочной форме:
### 
### NewSubtract(NewMultiply(NewConst(2),NewVariable("x"),),NewConst(3),).String()
### // Должно вывести: "((2 * x) - 3)"
###
### 4. Метод Equals
### Должен проверять, что два выражения идентичны:
###
### fmt.Println(NewMultiply(NewConst(2), NewVariable("x")).Equals(NewMultiply(NewConst(2), NewVariable("x"))),)
### // true
###
### fmt.Println(NewMultiply(NewConst(2), NewVariable("x")).Equals(NewMultiply(NewVariable("x"), NewConst(2))),)
### // false
### 
### 5. Функция main
### Должна вычислять значение выражения
### 
### "x^2 - 2x + 1"
### 
### для x, заданного в аргументах командной строки:
###
### // go run main.go 3
### // Должно вывести: 4
### 
### Советы по реализации:
### • Определите интерфейс Expression с методами:
### • Evaluate(x int) int
### • String() string
### • Equals(Expression) bool
### Создайте структуры Const, Variable, Add, Subtract, Multiply, Divide.
### • Реализуйте конструкторы NewConst, NewVariable, NewAdd, NewSubtract, NewMultiply, NewDivide, которые будут возвращать указатель на созданный объект.
### • Вынесите общую логику бинарных операций в абстрактную структуру (например, BinaryOp), чтобы избежать дублирования кода.