# Описание API endpoint

API endpoint предоставляет статистику комментариев для конкретного поста.

## URL

`http://localhost:8000/api/v1/post/:postId/comments/statistics`

## Метод

GET

## Параметры запроса

`:postId` - идентификатор поста для получения статистики комментариев

## Пример ответа

```json
{
  "data":"..."
  "msg": "Server work",
  "result": "Ok"
}
```

В ответе содержится поле `msg`, которое указывает на работоспособность сервера, и поле `result`, содержащее статус выполнения запроса.

## Ошибки

- Если `:postId` не существует, то сервер вернет ошибку с кодом HTTP 404 и сообщением "Пост не найден".
- Если произошла внутренняя ошибка сервера, то сервер вернет ошибку с кодом HTTP 500 и сообщением "Внутренняя ошибка сервера".

## Пример использования

```javascript
fetch('http://localhost:8000/api/v1/post/123/comments/statistics')
  .then(response => response.json())
  .then(data => {
    console.log(data);
  })
  .catch(error => {
    console.error('Ошибка:', error);
  });
```
