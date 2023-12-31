# 🐙



### API
Референс API можно найти [здесь](./API.md)

### Запуск
```
docker-compose up --build
```

### Чего здесь нет
- Логирование ведется просто в stdout, поднятие полноценного решения вроде Loki, как по мне, черезмерное в таком небольшом задании
- Нету grpc service discovery, просто потому что сервис такой здесь всего один
- Для main сервера использовал фреймворк gin, так как хотелось сделать все побыстрее, но продакшене бы лучше использовал голую http библиотеку
- Я не делал нормальный Makefile для генерации golang кода из protobuf, так как здесь никаких дальнейших телодвижений не придвидится, весь сгенерированный код лежит в папке gen, а команда для генерации, на всякий случай, в gen-proto.sh
- Тесты, проверил все вручную за 10 минут и как мне кажется, на таком масштабе тесты лишь бы раздули проект вдвое, а практического смысла было бы маловато

### Что я изменил
- В сервис auth я передаю не login, а userID, так как login это потенциально изменяемое значение (userID же всегда один). В данном случае это не имеет большого значение, так как нету метода для обновления данных пользователя, но пусть это просто будет хороший тон)
- ID пользователя и продукта сделал в формате UUID v4, это ни на что особо не влияет, просто так привычнее, надеюсь, ничего страшного

### Что бы я сделал иначе
- Сделать коммуникацию с auth сервисом так же по grpc, пусть у сервисов будет единый формат коммуникации без кучи разных коннекторов
- Сделать отдельный gateway, на данный момент main server является и входной точкой и центром всей архитектуры, я бы лучше сделал отдельный gateway, который бы уже перенаправлял запросы к нужным микросервисам
- Понятно, что на данном масштабе это не критично, но если думать про масштабирование, я бы разделил users и products на два отдельных микросервиса, чтобы они между собой не путались
- Шифрование токена (JWE), само собой, ничего страшного не случится, если пользователь узнает свой userID или login, но как говорится, better safe than sorry
