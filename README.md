## GoDockerDev
Среда для разработки на Go


### Старт разработки с GoDockerDev
1. **У вас должен быть установлен локально интерпретатор на GO**
2. `git clone git@github.com:pavel-one/GoDockerDev.git <yourProjectName>`
3. `cd <yourProjectName> && rm -rf .git`
4. `cp .env.example .env`
5. `make up`

### Команды
- `make up` - запуск приложения
- `make down` - остановка
- `make log` - прослушивание логов приложения
- `make exec` - зайти во внутрь контейнера

### Настройки, если что то не работает
В корнейвой папке есть `.env` - это файл настроек докера:

- `NGINX_PORT` - порт на котором будет работать локальный NGINX
- `DB_PORT` - порт базы данных локально
- `DOCKER_UID` - id текущего пользователя (чтобы его узнать, выполни `id` в консоли)
- `DOCKER_USER` - имя текущего пользователя (чтобы его узнать, выполни `id` в консоли)
- `DB_USER` - пользователь pg
- `DB_PASSWORD` - пароль pg
- `DB_NAME` - название бд в pg