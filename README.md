# DistArithExpCalc
Распределённый вычислитель арифметических выражений.
1. Скачайте архив в папку Загрузки(Downloads);

Если у вас Windows:
1.Запустите командную строку (Win+R) и наберите "cmd" или найдите пункт "Командная строка" в главном меню операционной системы. 
2. Напишите go version и нажмите Enter. Ответ должен указать, какая именно версия языка программирования в деле. Если в ответе нет версии - скачайте и установите golang c официального сайта https://go.dev/
3. Скачайте архив и разархивируйте архив.
4. Запустите файл "go-gorilla-api.exe". Если потребуется - дайте разрешение на доступ на вопрос защитника Windows.
5. Работа в командной строке или браузере
5.1 Откройте вторую(новую) командную строку
5.2 Выберите (Скопируйте и вставьте) одну или некоторые команды для просмотра данных:
curl http://127.0.0.1:8080/api/v1/expressions для просмотра всех выражений
curl http://127.0.0.1:8080/api/v1/expressions/:1 для просмотра выражений под номером
curl http://127.0.0.1:8080/api/v1/tasks для просмотра всех задач на вычисление
curl http://127.0.0.1:8080/api/v1/datas для просмотра всех данных

