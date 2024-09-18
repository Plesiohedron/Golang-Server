package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type GetAccountRequest struct {
	Name string `json:"name"`
}

type ChangeAccountNameRequest struct {
	Name    string `json:"name"`
	NewName string `json:"newName"`
}

type ChangeAccountBalanceRequest struct {
	Name      string `json:"name"`
	NewAmount int    `json:"newAmount"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
