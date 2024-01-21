# The_Enricher
Сервис, который будет получать по апи ФИО, из открытых апи обогащать ответ наиболее вероятными возрастом, полом и национальностью и сохранять данные в БД
___
# Оглавление
- Endpoints
- Задание
___
# Endpoints
|Method| URL |Discription| Body / Query params| Result|
|------|-----|--------|-------|------|
|**POST**| localhost:8080/create_user | Creates an entry in DB| name: string <br> surname: string <br> patronymic: string|  JSON {<br>name: string <br> surname: string <br> patronymic: string <br> age: int <br> gender: string <br> nationality: string<br>}|
|**POST**| localhost:8080/update_user | Updates an entry in DB by user_id and field name| user_id: int <br> field_to_update: string <br> new_value: string|  JSON {<br> message: string<br>}|