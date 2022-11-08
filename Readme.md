# Практика

> Назви гілок відповідають номеру тижня курсу. week_1, week_2 і так далі

Завдання на 1/11/2022 тут:[Readme.md](week_2/interfaces/imports/Readme.md)

Code **here** : [week_2/interfaces/imports](week_2/interfaces/imports)

## Початкова підготовка
1. Налаштовуєте git (*опціонально, якщо не налаштовано*)
```sh
git config --global user.name "Your Name"
git config --global user.email "your.email@gmail.com"
git config --global user.password "YOUR PASSWORD"
````
2. Відправляєте на емейл yevhenii.babich@globallogic.com ваш нікнейм та емейл у github.   


## Git
1. Клонуєте базовий репозиторій себе;
2. Створюєте свій репозитарій
2. Змінюєте ім'я remote;
3. Додаєте свій репозиторій як remote;
4. Відправляєте гілку майстер у свій репозиторій;
5. Перемикаєтеся на нову гілку
### Якщо у вас не настроєна робота через SSH c git:
```sh
git clone https://github.com/yevhenii-babich/go-lessons.git course

cd course
git remote rename origin base

# **_Ваше ім'я_** - слід замінити
git remote add origin https://github.com/**_Ваше ім'я_**/course

git remote -v
#base https://github.com/yevhenii-babich/go-lessons.git (fetch)
#base https://github.com/yevhenii-babich/go-lessons.git (push)
#origin https://github.com/gl-go-cources/course.git (fetch)
#origin https://github.com/gl-go-cources/course.git (push)

git push -u origin main

git checkout -b week_1
````
### Якщо у вас вже налагоджено роботу через SSH c git:
```shell
git clone git@github.com:yevhenii-babich/go-lessons.git course
cd course
git remote rename origin base

# **_Ваше ім'я_** - слід замінити
git remote add origin  git@github.com:**_Ваше ім'я_**/course.git

git remote -v
#base    https://github.com/yevhenii-babich/go-lessons.git (fetch)
#base    https://github.com/yevhenii-babich/go-lessons.git (push)
#origin  git@github.com:vasya/course.git (fetch)
#origin  git@github.com:vasya/course.git (push)
git push -u origin main

git checkout -b week_1

```
## Виконання завдань

У папці з кожним завданням є його опис, файл `main.go` та файли тестів.
Для виконання завдання необхідно відредагувати тіло функції, яка виконуватиме описані дії в завданні.

````
// Тут має бути рішення
// написавши код - необхідно запустити тести
// Ці коментарі можна видаляти
````

Для запуску тесту використовуйте
- або консольну команду `go test . ` - її необхідно запускати в папці конкретного завдання

````
--- FAIL: TestLeapYears (0.00s)
    leap_test.go:9: IsLeapYear(1996) = false, want true (year divisible by 4, not divisible by 100: leap year)
FAIL
FAIL    leap    0.002s
FAIL
````

- або кнопку запуску тестів у редакторі / IDE, якщо вона є

![кнопка запуску тесту у VSCode](run-tests-1.png "Кнопка запуску тестів у VSCod")

![кнопка запуску тесту у VSCode](run-tests-2.png "Кнопка запуску тестів у VSCod")

## Відправлення рішень

Виконавши завдання (коли всі тести проходять успішно) потрібно:
1. Закомітити зміни
2. Відправити їх на гітхаб у свій репозиторій

```sh
git commit -am 'Вирішено завдання першого тижня'
git push -u origin week_1
````

3. Створити мерж-реквест на гітлабі гілки з виконаними завданнями у гілку майстер.
4. Призначити мене ( Yevhenii Babich yevhenii-babich) відповідальним за мерж. Кнопка (Assignee)


## Терміни здачі

Необхідно здати вирішене завдання до 3 години ночі неділі поточного тижня (технічно - це ранок понеділка)

Топ найперших рішень розгляну більш роботально, краще (одне або кілька) розглянемо на лекції
