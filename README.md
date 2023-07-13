# ApolloHW

**ApolloHW** - Telegram бот, использующий веб-интерфейс, для создания расписаний и домашних занятий.

![](https://user-images.githubusercontent.com/83646375/251200663-fc1add4e-1611-47d3-a275-171a1e147349.png)

## Особенности

- Изменение расписаний и добавление примечаний
- Добавление домашних заданий с помощью текста или картинок и их автоматический перенос при изменениях
- Создание выходных дней и каникул
- Удобное управление классом и заявками

## Использование

Создайте бота в [BotFather](https://t.me/botfather) и добавьте кнопку "Меню", содержащую ссылку на будущий сайт.

Клонируйте репозиторий в новую папку (например, apollo):

```bash
git clone https://github.com/MrSterdy/ApolloHW apollo
```

И затем перейдите в неё:

```bash
cd apollo
```

*Не забудьте указать переменные окружения **DB_MYSQL** и **BOT_TOKEN**, означающие ссылку подключения к БД и токен Telegram бота соответственно*

### Go

Установите зависимости и запустите приложение:

```bash
go mod download
go run .
```

### Docker

Создайте образ приложения:

```bash
docker build --tag apollohw
```

И запустите его:

```bash
docker run -rm -p 8080:8080 -e BOT_TOKEN="<TOKEN>" -e DB_MYSQL="<URL>" --name apollo apollohw
```
