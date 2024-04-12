# Обзор Токена и его aрхитектуры

![Token service oriented architecture](token.png "Token SOA")

## Описание

Проект "Token" представляет собой реализацию токена на базе Hyperledger Fabric. Данный проект ставит своей задачей обеспечить понятность, поддерживаемость, тестируемость, расширяемость и отлаживаемость системы токенов.

Система строится на основе трёх основных слоёв:

- Слой Hyperledger Fabric, отвечающий за взаимодействие с блокчейном, хранение и обработку транзакций.

- Сервисный слой токена, где хранятся бизнес сущности, выполняется бизнес-логика токена и обрабатываются данные.

- Слой API контракта, служащий для обработки внешних вызовов и обеспечения взаимодействия с внешними системами.

Реализованная архитектура позволяет проводить быстрое и эффективное тестирование без реальной инфраструктуры блокчейна. Каждый слой проекта тщательно проработан и оптимизирован для обеспечения максимальной производительности и удобства использования.

Проект разработан с учётом лучших практик разработки ПО и архитектуры.

## Содержание

- [Обзор](#обзор-токена-и-его-aрхитектуры)
    - [Описание](#описание)
    - [Содержание](#содержание)
    - [Слой Hyperledger Fabric](#слой-hyperledger-fabric)
        - [ChaincodeStubInterface](#chaincodestubinterface)
        - [Foundation и Батчинг](#foundation-и-батчинг)
        - [Реализация интерфейса Chaincode key-value DB](#реализация-интерфейса-chaincode-key-value-db)
    - [Сервисный слой токена](#сервисный-слой-токена)
        - [Абстракция key-value базы данных](#абстракция-key-value-базы-данных)
        - [Слой хранения и репозиторий](#слой-хранения-и-репозиторий)
        - [Модель данных](#модель-данных)
        - [Сервисный уровень и контроллер](#сервисный-уровень-и-контроллер)
    - [Слой API контракта](#слой-api-контракта)
        - [Объекты передачи данных (DTO)](#объекты-передачи-данных-dto)
        - [Вспомогательные функции проверки доступа](#вспомогательные-функции-проверки-доступа)
        - [Contract Tx и Query API вызовы](#contract-tx-и-query-api-вызовы)
    - [Сборка и инициализация](#сборка-и-инициализация)
    - [Лицензия](#лицензия)
    - [Ссылки](#ссылки)


## Слой Hyperledger Fabric

### ChaincodeStubInterface
ChaincodeStubInterface представляет собой основной интерфейс в Hyperledger Fabric для взаимодействия с ledger (журналом транзакций).

### Foundation и Батчинг
Foundation является обёрткой над ChaincodeStubInterface, которая добавляет функциональность обработки транзакций пакетами, что позволяет оптимизировать и упростить работу с транзакциями.

### Реализация интерфейса Chaincode key-value DB
Chaincode key-value DB это реализация абстракции key-value базы данных, которая непосредственно связана с Foundation и представляет собой адаптацию к специфическим особенностям работы в среде Hyperledger Fabric.


## Сервисный слой токена

### Абстракция key-value базы данных
Вводит абстракцию в виде интерфейса keyvalue.DB, что позволяет унифицировать работу с данными и развязывать зависимость от конкретной реализации базы данных.

### Слой хранения и репозиторий
Пакет storage оперирует над объектами модели, используя keyvalue.DB. Автоматически генерируются интерфейсы пакета repository, что позволяет проводить тестирование других пакетов без привязки к реальному storage.

### Модель данных
В пакете model хранятся бизнес сущности, представляющие основные структуры данных, с которыми работает система.

### Сервисный уровень и контроллер
Сервисный уровень, представленный пакетом service, содержит в себе весь функционал бизнес-логики токена. Работает над объектами model и сохраняет всё в storage через интерфейс repository.


## Слой API контракта

### Объекты передачи данных (DTO)
Представляют собой объекты, предназначенные для обмена данными между различными уровнями архитектуры и внешними вызовами.

### Вспомогательные функции проверки доступа
Функции, помогающие в проверке доступа к различным частям API, что улучшает безопасность и управление правами доступа.

### Contract Tx и Query API вызовы
Contract обрабатывает все внешние вызовы, проверяет права на вызов, парсит входящий DTO, мапит его на model, вызывает service через интерфейс controller, получает ответ, мапит на DTO и отправляет вызывающей стороне.

## Сборка и инициализация
Все зависимости и инициализация компонентов проекта происходят в [main.go](main.go). Обзор этого файла может помочь в понимании общей структуры проекта и связей между его компонентами.

## Лицензия

[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)

## Ссылки

Дополнительные ссылки не предоставлены.