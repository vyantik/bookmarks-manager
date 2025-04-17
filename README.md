# Менеджер закладок (Go)

Это простой консольный менеджер закладок, написанный на языке Go. Он позволяет пользователям добавлять, просматривать и удалять закладки, каждая из которых состоит из имени и URL-адреса.

## Функциональность

-   **Просмотр закладок:** Отображает все сохраненные закладки.
-   **Добавление закладки:** Позволяет добавить новую закладку, запросив у пользователя имя и URL.
-   **Удаление закладки:** Позволяет удалить существующую закладку по ее имени.
-   **Выход:** Завершает работу программы.

## Как запустить

1.  **Установите Go:** Убедитесь, что на вашем компьютере установлен Go (версия 1.x или выше).
2.  **Получите код:** Склонируйте репозиторий или скачайте файлы проекта (`main.go` и `bookmarks_helper.go`) в одну директорию.
3.  **Запустите:** Откройте терминал или командную строку, перейдите в директорию с файлами проекта и выполните команду:
    ```bash
    go run .
    ```
4.  **Используйте меню:** Следуйте инструкциям в консоли для управления закладками:
    -   Введите `1` для просмотра списка закладок.
    -   Введите `2` для добавления новой закладки.
    -   Введите `3` для удаления закладки.
    -   Введите `4` для выхода из программы.

## Файлы проекта

-   <mcfile name="main.go" path="z:\Projects\go\Learning\MAP_LEARNING\main.go"></mcfile>: Содержит основную функцию <mcsymbol name="main" filename="main.go" path="z:\Projects\go\Learning\MAP_LEARNING\main.go" startline="5" type="function"></mcsymbol> и главный цикл обработки пользовательского ввода.
-   <mcfile name="bookmarks_helper.go" path="z:\Projects\go\Learning\MAP_LEARNING\bookmarks_helper.go"></mcfile>: Содержит вспомогательные функции для работы с закладками (<mcsymbol name="addBookmark" filename="bookmarks_helper.go" path="z:\Projects\go\Learning\MAP_LEARNING\bookmarks_helper.go" startline="11" type="function"></mcsymbol>, <mcsymbol name="deleteBookmark" filename="bookmarks_helper.go" path="z:\Projects\go\Learning\MAP_LEARNING\bookmarks_helper.go" startline="23" type="function"></mcsymbol>, <mcsymbol name="printBookmarks" filename="bookmarks_helper.go" path="z:\Projects\go\Learning\MAP_LEARNING\bookmarks_helper.go" startline="32" type="function"></mcsymbol>), получения ввода пользователя (<mcsymbol name="getUserInput" filename="bookmarks_helper.go" path="z:\Projects\go\Learning\MAP_LEARNING\bookmarks_helper.go" startline="41" type="function"></mcsymbol>) и очистки экрана (<mcsymbol name="clearScreen" filename="bookmarks_helper.go" path="z:\Projects\go\Learning\MAP_LEARNING\bookmarks_helper.go" startline="55" type="function"></mcsymbol>).
