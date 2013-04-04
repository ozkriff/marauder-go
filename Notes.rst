Разделить inputHandler на непосредственно
получение событий и их обработку.
Т.к. последнее можно будет протестировать :)

Показать все названия пакетов, кроме тестовых::

    find . -name "*.go" | xargs grep "^package" --color --no-filename | uniq | grep -v _test | sort
