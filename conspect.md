## Лекция 1. Работа с JSON файлами

**JSON** - формат файлов (расширение файлов) который повсеместно используется для реализации передачи данных между серверами на уровне API

**JSON** == **JavaScript Object Notation** (Object - аналог map в Go  только для мира JS)

**JSON** -  простейшее файловое расширение, поддерживающе элементарную структуризацию (выглядит как пара ключ-значение)

## Шаг 1. Создадим простой .json файл
Для этого определим файл ```users.json```.
```
```
"users": [{"name" : "Vasya"}, {"name" : "Vitya"}]}
```
Обратите внимание на то, что ***Что по стандарту в JSON используются двойные кавычки***

### Шаг 2. Создадим чуть более сложный .json
Создадим сразу читаемый

```
{
    "users" : [
        {
          "name" : "Alex",
          "type" : "Admin",
          "age" : 32,
          "social" : {
                "vkontakte" : "https://vk.com/id=123512",
                "facebook": "https://fb.com/id=172835"
            }
        },
        {
          "name" : "Bob",
          "type" : "Regular",
          "age" : 12,
          "social" : {
                "vkontakte" : "https://vk.com/id=123561235",
                "facebook": "https://fb.com/id=19283712"
            }
        },
        {
          "name" : "Alice",
          "type" : "Regular",
          "age" : 19,
          "social" : {
                "vkontakte" : "https://vk.com/id=12123123",
                "facebook": "https://fb.com/id=172123123"
            }
        },
        {
          "name" : "George",
          "type" : "Regular",
          "age" : 42,
          "social" : {
                "vkontakte" : "https://vk.com/id=999999",
                "facebook": "https://fb.com/id=98888888"
            }
        }
    ]
}
```

***ВАЖНО*** - .json не гарантирует упорядоченную выдачу ключей

