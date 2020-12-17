# AIS
[ 9 семестр: системы искусственного интеллекта ]

## Лабораторная работа 2
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


Список имено находится в файле names.json. 

### Пример работы

```
 Введите название статьи-эталона: 
 Consequatur perferendis sit aut accusantium voluptatem.
 ---------------------
 Введите название второй статьи: 
 Voluptatem accusantium aut perferendis consequatur sit.
 ---------------------
 Статья А: {
        Title:Consequatur perferendis sit aut accusantium voluptatem. 
        Authors:[Prof. Aglae Friesen Queen Guadalupe Morissette Ms. Jazlyn Kulas] 
        Fields:[complex_manifold] 
        RINZ:true 
        WAK:false 
        WOS:false 
        Year:2011 
        Citations:3 
        Score:0.4409283900936391 
        ReadingTime:21
        } 

 Статья B: {
        Title:Voluptatem accusantium aut perferendis consequatur sit. 
        Authors:[Prof. Aglae Friesen Mrs. Elna Schaden Miss Bernita Feeney] 
        Fields:[complex_manifold] 
        RINZ:true 
        WAK:false 
        WOS:false 
        Year:1974 
        Citations:36 
        Score:0.5354199838094965 
        ReadingTime:18
        } 

 Значения метрик: 
 Евклидово расстояние:  4.582807945057589
 Корреляция:  0.6176683984397627
 Расстояние по графу:  0
 Количество различий:  6
```