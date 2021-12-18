# SIF

## Фронтенд
### Запуск фронтенда
<pre>
    <code>
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
