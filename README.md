# AIS
[ 9 семестр: системы искусственного интеллекта ]

## Лабораторная работа 3
### Как запустить
Сборка: 

`go build cmd\main.go`

Запуск:
 
 `go run cmd\main.go`
### Описание данных 

Данные для лабораторной сгенерированы случайным образом по следующему принципу:
* Название статьи является случайным предложением, для генерации используется  faker.Sentence()
* Количество авторов - от 1 до 3, выбирается случайно, для генерации каждого используется  faker.Name()
* Поле исследования - одно случайное поле, выбирается из случайным образом из графа
* Принадлежность базам (РИНЦ, ВАК, WoS) - случайным образом выбирается одна из баз
* Год - случайное число из диапазона \[1970, 2020]
* Количество цитирований - случайное число из диапазона \[0, 40]
* Оценка - случайное число из диапазона \[1, 0)
* Время чтения - случайное число из диапазона \[5, 30)


Список имен находится в файле names.json

Для коллаборативной фильтрации сгенерирована матрица оценок пользователями статьей.
* Имя пользователя является случайным словосочетанием, для генерации используется faker.Name()
* Оценка пользователем статьи - случайное число из диапазона \[0, 5]

Полученная матрица находится в файле users.csv


### Описание алгоритма коллаборативной фильтрации 

* Определяется текущий пользователь 
* Определяется корреляция оценок пользователей с оценками текущего пользователя
* Отсекаются пользователи с отрицательной оценкой корреляции
* Определяется текущая статья
* Определяются оценки статей, для этого суммируются оценки отобранных пользователей, предварительно помноженные 
на значение их корреляции с текущим пользователем

### Пример работы

```
 Введите как вас зовут: 
 Lady Anya Dickinson

 Привет, Lady Anya Dickinson!
 Введите название статьи: 
 Consequatur accusantium voluptatem aut sit perferendis.

 Статья: 
    {Title:Consequatur accusantium voluptatem aut sit perferendis. 
    Authors:[Dr. Estel Jacobi Queen Marlee Wuckert] 
    Fields:[route_problems] 
    RINZ:true 
    WAK:false 
    WOS:false 
    Year:1980 
    Citations:30 
    Score:0.5934263023272138 
    ReadingTime:23} 
     
Наиболее подходящие вам статьи: 
    
    {Title:Perferendis voluptatem accusantium sit consequatur aut. 
    Authors:[Princess Gisselle Hammes Prof. Aurelia Wiza Miss Celine Towne] 
    Fields:[homotopy_groups] 
    RINZ:true 
    WAK:false 
    WOS:false 
    Year:2018 
    Citations:28 
    Score:0.8152128777352514 
    ReadingTime:10} 
     
    {Title:Voluptatem perferendis sit accusantium consequatur aut. 
    Authors:[Ms. Maeve Anderson] 
    Fields:[subsets] 
    RINZ:false 
    WAK:true 
    WOS:false 
    Year:1978 
    Citations:16 
    Score:0.024728093691463882 
    ReadingTime:8} 
     
    {Title:Sit aut voluptatem accusantium perferendis consequatur. 
    Authors:[Princess Claudie Torp Queen Eve Halvorson] 
    Fields:[small_stellated_dodecahedron] 
    RINZ:false 
    WAK:true 
    WOS:false 
    Year:1989 
    Citations:34 
    Score:0.7454155482508904 
    ReadingTime:10} 
     
    {Title:Consequatur accusantium voluptatem aut perferendis sit. 
    Authors:[Princess Daphne Nader] 
    Fields:[mathematical_physics] 
    RINZ:false 
    WAK:false 
    WOS:true 
    Year:1985 
    Citations:6 
    Score:0.7748192532007393 
    ReadingTime:26} 
     
    {Title:Accusantium perferendis aut voluptatem consequatur sit. 
    Authors:[Dr. Winnifred Goldner Ms. Grace Cummings Ms. Beryl Heathcote] 
    Fields:[algebraic_fractals] 
    RINZ:false 
    WAK:false 
    WOS:true 
    Year:2001 
    Citations:5 
    Score:0.9642233503677278 
    ReadingTime:22} 
```