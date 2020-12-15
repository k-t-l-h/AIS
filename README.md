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


Список имено находится в файле names.json

### Пример работы

```
 Введите название статьи-эталона: 
 Sit consequatur accusantium voluptatem perferendis aut.
 ---------------------
 Введите название второй статьи: 
 Voluptatem accusantium consequatur sit aut perferendis.
 ---------------------
 Статья А: {Title:Sit consequatur accusantium voluptatem perferendis aut. 
            Authors:[Princess Althea Kessler Ms. Estel Hane] 
            Fields:[series] 
            RINZ:true 
            WAK:false 
            WOS:false 
            Year:1990 
            Citations:1 
            Score:0.7436350178287354 
            ReadingTime:20} 

 Статья B: {Title:Voluptatem accusantium consequatur sit aut perferendis. 
            Authors:[Princess Bulah Mills Dr. Shayna Stiedemann] 
            Fields:[multiplication] 
            RINZ:true 
            WAK:false 
            WOS:false 
            Year:2011 
            Citations:12 
            Score:0.26547149784194785 
            ReadingTime:6} 

 Значения метрик: 
 Евклидово расстояние:  27.96835069059036
 Корреляция:  0.9999522090779444
 Расстояние по графу:  8
 Количество различий:  7
```