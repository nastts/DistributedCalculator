# Распределённый вычислитель арифметических выражений Golang
#### Распределённый вычислитель арифметических выражений состоит из 2-ух елементов:
### **Оркестратор** - сервер, который принимает арифметическое выражение, переводит его в набор последовательных задач и обеспечивает порядок их выполнения
### **Агент** - вычислитель, который может получить от оркестратора задачу, выполнить его и вернуть серверу результат
---
>[!TIP]
>### локальная **[ссылка](http://localhost:8080/api/v1/calculate)** сервера 


# Запуск проекта
### 1. **сохраните [проект](https://github.com/nastts/Calculate/archive/refs/heads/main.zip)**
### 2. **откройте терминал и пропишите команду для запуска сервера**
```powershell
go run ./cmd/main.go
```
### 3. **если вы получили сообщение:**
```Go
2025/03/03 03:03:03 сервер запущен
```
### **значит сервер запустился корректно**
>[!IMPORTANT]
>после того, как вы запустили сервер, создайте новый терминал, что отправить запрос

# Использование
>[!IMPORTANT]
>### чтобы отправить выражение на вычисление, необходимо прописать:

```powershell
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"22*2"}' | Select-Object -Expand Content
```
end-point: http://localhost:8080/api/v1/calculate

### в ответ вы получите id

```powershell
{"id":"1"}
```

### код 422❌ - невалидные данные

```powershell
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"22*"}' | Select-Object -Expand Content
```
### результат
```powershell
Invoke-WebRequest : expression is not valid
строка:1 знак:1
+ Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/c ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebR 
   equest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Co 
   mmands.InvokeWebRequestCommand
```
---
>[!IMPORTANT]
>### для получение списка выражений:


```powershell
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/expressions' -ContentType 'application/json' | Select-Object -Expand Content
```
### результат

```powershell
{"expressions":[{"id":"1","status":"done","result":44},{"id":"4","status":"in progress","result":0}]}
```
end-point: http://localhost:8080/api/v1/expressions

### код 404❌ - не найденные данные

```powershell
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/expressions' -ContentType 'application/json' | Select-Object -Expand Content
```
### результат
```powershell
Invoke-WebRequest : expressions is not found
строка:1 знак:1
+ Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/c ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebR 
   equest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Co 
   mmands.InvokeWebRequestCommand
```

---
>[!IMPORTANT]
>### для получение одного выражения:


```powershell
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/expressions/1' -ContentType 'application/json' | Select-Object -Expand 
Content
```
### результат

```powershell
{"expression":{"id":"1","status":"done","result":44}}
```
end-point: http://localhost:8080/api/v1/expressions/1 (id по которому хотите найти выражение)

### код 404❌ - не найденные данные

```powershell
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/expressions/52' -ContentType 'application/json' | Select-Object -Expand 
Content
```
### результат
```powershell
Invoke-WebRequest : id is not found
строка:1 знак:1
+ Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/e ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebR 
   equest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Co 
   mmands.InvokeWebRequestCommand
```
---
>[!IMPORTANT]
>### Получение task


```powershell
Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/internal/task' -ContentType 'application/json' | Select-Object -Expand Content
```
### результат
```powershell
{"task":{"id":"1","arg1":2,"arg2":22,"operation":"*","operationTime":1000}}
```
end-point: http://localhost:8080/internal/task

### код 404❌ - не найденные данные
```powershell
Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/internal/task' -ContentType 'application/json' | Select-Object -Expand Content
```
```powershell
Invoke-WebRequest : tasks is not found
строка:1 знак:1
+ Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/internal/ ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebR  
   equest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Co  
   mmands.InvokeWebRequestCommand
```
---
>[!IMPORTANT]
>### Получение результата task


```powershell
Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/internal/task' -ContentType 'application/json' | Select-Object -Expand Content
```
### результат
```powershell
{"task":{"id":"1","result":"44"}}
```
end-point: http://localhost:8080/internal/task
### код 422❌ - не валидные данные
```powershell
Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/internal/task' -ContentType 'application/json' | Select-Object -Expand Content
```
```powershell
Invoke-WebRequest : expression is not valid
строка:1 знак:1
+ Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/internal/ ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebR  
   equest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Co  
   mmands.InvokeWebRequestCommand
```
### код 404❌ - не найденные данные
```powershell
Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/internal/task' -ContentType 'application/json' | Select-Object -Expand Content
```
```powershell
Invoke-WebRequest : tasks is not found
строка:1 знак:1
+ Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/internal/ ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebR  
   equest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Co  
   mmands.InvokeWebRequestCommand
```