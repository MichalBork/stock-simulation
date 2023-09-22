package transaction_handler

import (
	"encoding/json"
	"net/http"
	"stock-simulation/pkg/handlers/api_client_handler"
	"stock-simulation/pkg/model"
	"stock-simulation/pkg/repository"
)

type Handler struct {
	WalletRepository repository.WalletRepository
}

func (h *Handler) BuyCrypto(w http.ResponseWriter, r *http.Request) {
	// 1. Dekoduj dane wejściowe (zakładając, że masz odpowiedni model RequestBody dla tej operacji)
	var buyRequest BuyRequest
	err := json.NewDecoder(r.Body).Decode(&buyRequest)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// 2. Pobierz aktualną cenę kryptowaluty
	price, err := api_client_handler.GetCoinPrice(buyRequest.CoinName)
	if err != nil {
		http.Error(w, "Failed to get coin price", http.StatusInternalServerError)
		return
	}

	// 3. Oblicz całkowity koszt transakcji
	totalCost := buyRequest.Amount * price

	userWallet, err := h.WalletRepository.FindByUserID(buyRequest.UserID)
	if err != nil {
		http.Error(w, "Failed to retrieve user's wallet", http.StatusInternalServerError)
		return
	}

	// Sprawdzenie, czy użytkownik ma wystarczająco dużo środków w portfelu
	if userWallet.Currency == "USD" && userWallet.Balance < totalCost {
		http.Error(w, "Insufficient funds", http.StatusBadRequest)
		return
	}

	// Aktualizacja portfela
	userWallet.Balance -= totalCost
	err = h.WalletRepository.UpdateWallet(userWallet)
	if err != nil {
		http.Error(w, "Failed to update wallet", http.StatusInternalServerError)
		return
	}

	// 4. Zaktualizuj portfel użytkownika (musisz sprawdzić, czy użytkownik ma wystarczająco dużo USD, następnie odjąć tę ilość i dodać zakupioną kryptowalutę)
	// Uwaga: Upewnij się, że operacja aktualizacji portfela jest atomowa i bezpieczna w przypadku równoczesnych transakcji.

	// 5. Zapisz transakcję w bazie danych
	transaction := model.Transaction{
		UserID:         buyRequest.UserId,
		BaseCurrency:   "USD",
		TargetCurrency: buyRequest.CoinName,
		Amount:         buyRequest.Amount,
		ExchangeRate:   price,
		Type:           "buy",
	}
	// Zapisz transakcję w bazie danych...

	// 6. Odpowiedz użytkownikowi
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Purchase successful"})
}

type BuyRequest struct {
	CoinName string  `json:"coin_name"`
	Amount   float64 `json:"amount"`
	UserId   int64   `json:"user_id"`
}
