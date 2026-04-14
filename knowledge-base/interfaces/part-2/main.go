package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}

type SimpleAccountNotifier struct{}

func (n *SimpleAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("new account created", " username", account.Username)
	return nil
}

type BetterAccountNotifier struct{}

func (n *BetterAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("new account created by BetterAccountNotifier", " username", account.Username)
	return nil
}

type Account struct {
	Username string
	Email    string
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

func (h *AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

// notify can use email, sms, Discord, telegram or slack <- so its better to use interface
func notifyAccountCreation(account Account) error {
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("Account created: %s\n", account.Username)
	return nil
}

func main() {
	mux := http.NewServeMux()
	accountHandler := &AccountHandler{
		AccountNotifier: &BetterAccountNotifier{},
	}

	mux.HandleFunc("/create-account", accountHandler.handleCreateAccount)
	http.ListenAndServe(":3000", mux)

}
