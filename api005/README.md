
# Project Title

A brief description of what this project does and who it's for

Api v0.0.5
API działa. Dla uruchomienia api trzeba uruchomić projekt i wpisać w przeglądarce http://localhost:8080/health-check. W odpowiedz musi pojawić się strona z napisem "API is up and running"

Takoż działa Funkcja logowania, dla tego, żeby ją sprawdzić trzeba przejść w przeglądarce za pomocą linku: http://localhost:8080/login. Na danej stronie są dwa pola: „Login” i „Hasło”. Po wpisaniu danych i kliknięciu „send” użytkownik przejdzie na jedną ze stron, z pozytywnym albo negatywnym wynikiem.
Do routingu wykorzystałem gorilla/mux.
Baza danych SQLite podłączona i działa podczas sprawdzania użytkowników.
Niestety nie zdążyłem zbadać dgrijalva/jwt-go dla tokenów i zaimplementować inne funkcje przez brak czasu.


