## Менеджер баланса

Консольное приложение на Go для управления балансом пользователя с базовыми банковскими операциями.

## Зависимости

- Go 1.26+

## Установка

```bash
git clone https://github.com/kreslokatalka/UBM
cd UBM
go run project.go
```
## Описание версии
Для защиты от параллельной обработки в методах Withdraw, Deposit и Get_balance был использован Mutex