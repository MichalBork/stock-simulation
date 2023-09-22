## Rejestracja
* Implementacja endpointu rejestracyjnego (np. POST /register), który akceptuje dane użytkownika, takie jak nazwa użytkownika, hasło i email.
* Implementacja funkcji, której zadaniem będzie walidacja danych otrzymanych od użytkownika.
* Implementacja funkcji, która będzie hashować hasło przed zapisaniem go do bazy danych.
* Aktualizacja UserRepository z funkcją na dodawanie nowego użytkownika do bazy.
Łączenie z API kryptowalut
* Badanie i wybór odpowiedniego API do obsługi kryptowalut.
* Implementacja klienta API do obsługi połączeń z wybranym API krypto.
* Implementacja funkcji, które będą przetwarzać dane otrzymane z API krypto.
* Pobieranie kursów kryptowalut
Implementacja websocketu, który będzie komunikować się z serwerem w czasie rzeczywistym do otrzymywania danych o kursie kryptowalut.
* Implementacja endpointu, który będzie obsługiwać te dane i dostarczać je do klienta.
* Możliwość zakupu kryptowalut
Implementacja nowej tabeli w bazie danych do przechowywania informacji o transakcjach.
Implementacja endpointu do obsługi żądań zakupu od klienta.
Aktualizacja UserRepository (lub stworzenie nowego TransactionRepository) z funkcjonalnością dodawania i odczytu transakcji.
Ranking użytkowników
Implementacja algorytmu do obliczania rankingu użytkowników.
Implementacja endpointu do obsługi żądań ranglisty od klienta.
Dodanie funkcjonalności do UserRepository (lub stworzenie nowego RankingRepository) do zar