# Sprawozdanie z zadania
## Wstęp
Ten program napisany w go ma za zadanie "zabawę z liczbami" tzn. stara się odpowiedzieć na pytanie jaka jest silna oraz słaba liczba dla danego imienia i nazwiska, przy czym silną liczbą nazywamy taką, dla której po stworzeniu nicku z imienia i nazwiska silnia z tej liczby będzie zawierać wszystkie kody numeryczne liter z nicku. Słabą liczbą jest wyraz, dla którego liczba wywołań rekurencyjnej funkcji znajdującej trzydziesty wyraz ciągu Fibonacciego jest najbardziej zbliżona do silnej liczby.

## Opis funkcji programu

## `factorial(number int) *big.Int`

Funkcja obliczająca silnię dla zadanego numeru, do obliczeń używa wskaźnika do typu big.Int z biblioteki math/big co umożliwia obliczanie bardzo dużych silni.

## `fib(n int) int`

Funkcja obliczająca nty wyraz ciągu fibonacciego, dodatkowo funkcja sumuje i zapisuje liczbę wywołań w mapie `callCount`. 

## `containsSubstring(container, containee string) int`

Ta funkcja sprawdza, czy ciąg znaków (`containee`) jest podciągiem innego ciągu znaków (`container`). Zwraca 1, jeśli `containee` jest znalezione w `container`, a 0 w przeciwnym przypadku. Zwraca wynik typu int ponieważ jest to później wykorzystane w programie do sumowania.

## `generateNick(name, surname string) string`

Ta funkcja generuje pseudonim, biorąc pierwsze trzy znaki z `name` oraz drugi, trzeci i czwarty znak z `surname`, a następnie je ze sobą łączy.

## `generateAsciiCodesAsStrings(nick string) [6]string`

Ta funkcja konwertuje każdy znak w ciągu znaków `nick` na jego kod ASCII i zwraca tablicę tych kodów jako ciągów znaków.

## `isStrongNumber(asciiCodes [6]string, number *big.Int) bool`

Ta funkcja sprawdza, czy podana liczba jest "silną liczbą". Liczba jest uważana za silną, jeśli zawiera wszystkie kody ASCII pseudonimu. Funkcja zwraca true, jeśli liczba jest silna, a false w przeciwnym razie. Tutaj właśnie wykorzystywane jest to, że `containsSubstring` zwraca int, ta funkcja sprawdza sumę kontrolną tj. czy każdy element tablicy `asciiCodes`, a jest ich 6, zawiera się w `number`.

## `findStrongNumber(name, surname string) int`

Ta funkcja znajduje pierwszą silną liczbę, generując pseudonim z `name` i `surname`, konwertując pseudonim na kody ASCII, a następnie sprawdzając każdą silnię począwszy od 0, aż znajdzie silną liczbę.

## `main()`

Funkcja, która steruje wykonaniem całego programu, inicjalizuje mapę przechowującą liczbę wywołań dla funkcji `fib`, oraz wypisuje silną liczbę dla imienia i nazwiska autora.

## Wyniki

Silna liczba dla imienia i nazwiska `Konrad Kreczko` to `261`, zaś słaba liczba `18`.