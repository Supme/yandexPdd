# yandexPdd

Библиотека для работы с API Яндекс почта для домена

// Создаём сессию:

s := pdd.New("you.domain.example", "<ПДД-токен>")

// получаем список всех рассылок нашего домена

ls, err := s.ListGet()

// список всех ящиков входящих в рассылку

subscribers, err := s.ListSubscribers("list@you.domain.example)

...

...

...
