# Software Engineering – Aufgabe 2

Das Fensterbeispiel aus der Vorlesung zeigt das Composite Pattern in Go. `Glass` und `Frame` sind Blätter. Ein `Window` kann diese Elemente und weitere Fenster enthalten. Alle drei Typen implementieren `Item`.

Der [Prompt](PROMPT.md) beschreibt die Aufgabe für ein KI-Coding-Tool. Die Implementierung steht in [main.go](main.go).

## Ausführen

```sh
go run .
```

Das Programm gibt den Fensterbaum und den Gesamtpreis von 90,00 € aus.
