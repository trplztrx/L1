## 9. Сколько существует способов задать переменную типа slice или map?

### Для `slice`

| Способ                    | Пример                 |
|---------------------------|------------------------|
| Литерал                   | `s1 := []int{1, 2, 3}` |
| make                      | `s2 := make([]int, 5)` |
| Срезка от массива         | `s4 := arr[1:4]`       |
| Срезка от другого слайса  | `s6 := s5[1:4]`        |
| var                       | `var s7 []int`         |

### Для `map`

| Способ | Пример |
|---------|---------------------------------|
| Литерал | `m1 := map[string]int{"a": 1}`  |
| make    | `m2 := make(map[string]int)`    |
| var     | `var m4 map[string]int`         |
