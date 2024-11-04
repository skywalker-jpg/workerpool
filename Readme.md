# Worker Pool

## Описание

Приложение реализует пул воркеров на языке Go, позволяя динамически добавлять и удалять воркеров, обрабатывающих задания. Каждый воркер выполняет задания из очереди и логирует свою активность. Это приложение демонстрирует основы параллельной обработки и управления потоками, что полезно для выполнения долгих операций.

## Функции

- Динамическое добавление и удаление воркеров.
- Обработка заданий из очереди.
- Логирование работы воркеров.
- Поддержка командной строки для управления воркерами.

## Установка

### Шаг 1: Установите Go

Убедитесь, что у вас установлен Go (версии 1.18 и выше). Если Go еще не установлен, выполните следующие шаги:

1. Перейдите на [официальный сайт Go](https://golang.org/dl/).
2. Скачайте версию Go, соответствующую вашей операционной системе.
3. Следуйте инструкциям по установке.

Чтобы проверить установку, выполните в терминале:

```bash
go version
```

### Шаг 2: Клонируйте репозиторий

Склонируйте репозиторий с кодом приложения на ваш компьютер. Для этого выполните следующую команду в терминале:

```bash
git clone <URL_репозитория>
```

### Шаг 3: Перейдите в директорию проекта

Перейдите в директорию проекта:

```bash
cd <имя_директории_проекта>
```

### Шаг 4: Скомпилируйте приложение

В директории проекта выполните команду для компиляции приложения:

```bash
go build -o worker-pool
```

### Шаг 5: Запустите приложение

Запустите приложение, выполнив команду:

```bash
./worker-pool
```

## Использование

После запуска приложения вы можете использовать следующие команды:

- **add**: добавляет нового воркера в пул.
- **remove**: удаляет воркера по ID. Введите команду `remove`, затем введите ID воркера, который вы хотите удалить.
- **exit**: завершает работу программы.

### Пример взаимодействия

После запуска приложения вы увидите список доступных команд. Например:

1. Чтобы добавить воркера, введите команду `add`. Вы увидите сообщение о том, что воркер запущен.

2. Чтобы удалить воркера, введите команду `remove` и затем ID воркера. Вы увидите сообщение о том, что воркер удален.

3. Чтобы выйти из программы, введите команду `exit`. Вы увидите сообщение о завершении работы.

## Логирование

Логи работы воркеров выводятся в стандартный вывод, что позволяет отслеживать их деятельность в реальном времени. Каждый воркер логирует свои действия, включая начало и завершение обработки задания.
