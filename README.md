Example for go-qml i18n
=======================

Progress so far:

-   Wrote `example.go` and `ui.qml`.

-   Executed:

        mkdir i18n
        lupdate ui.qml -ts i18n/base.ts
        cp i18n/base.ts i18n/qml_de.ts

-   Ran Qt Linguist to translate `i18n/qml_de.ts`.
