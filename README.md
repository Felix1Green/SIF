# SIF
## Бэкенд
### Запуск бэкенда
<pre>
    <code>
  docker-compose -f ./backend/docker-compose.yml up -d --build
    </code>
</pre>

### Swagger
#### Для сервиса UserFacade находится по урлу:
<pre>
    <code>
  http://localhost:8080/swagger/
    </code>
</pre>
#### Для сервиса ChatService находится по урлу:
<pre>
    <code>
  http://localhost:8890/swagger/
    </code>
</pre>

## Фронтенд
### Запуск фронтенда
<pre>
    <code>
  cd frontend
  npm i
  npm start
    </code>
</pre>

### Разработка фронтенда
Чтобы при изменении значений глобальных токенах менялись переменные css, нужно в отдельном потоке запустить themekit
<pre>
    <code>
  themekit build --watch
    </code>
</pre>

Перед тем как залить изменения на гит, нужно запутить линтеры
<pre>
    <code>
  npm run lint
    </code>
</pre>
