package main

import (
	"GoCourse/HW-2/accounts/dto"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Port      int
	Host      string
	Cmd       string
	Name      string
	Amount    int
	NewName   string
	NewAmount int
}

func main() {
	portVal := flag.Int("port", 8080, "port to listen on")
	hostVal := flag.String("host", "localhost", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newNameVal := flag.String("newName", "", "new name of account")
	newAmountVal := flag.Int("newAmount", 0, "new balance of account")

	flag.Parse()

	cmd := Command{
		Port:      *portVal,
		Host:      *hostVal,
		Cmd:       *cmdVal,
		Name:      *nameVal,
		Amount:    *amountVal,
		NewName:   *newNameVal,
		NewAmount: *newAmountVal,
	}

	if err := do(cmd); err != nil {
		panic(err)
	}
}

func do(cmd Command) error {
	switch cmd.Cmd {
	case "create":
		if err := create(cmd); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := get(cmd); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "delete":
		if err := deleteAccount(cmd); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}

		return nil
	case "change_name":
		if err := changeAccountName(cmd); err != nil {
			return fmt.Errorf("change account's name failed: %w", err)
		}

		return nil
	case "change_balance":
		if err := changeAccountBalance(cmd); err != nil {
			return fmt.Errorf("change account's balance failed: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unknown command: %s", cmd.Cmd)
	}
}

func get(cmd Command) error {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name))

	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d", response.Name, response.Amount)

	return nil
}

func create(cmd Command) error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data))

	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func deleteAccount(cmd Command) error {
	request := dto.DeleteAccountRequest{
		Name: cmd.Name,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s:%d/account/delete", cmd.Host, cmd.Port), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http delete failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}
		return fmt.Errorf("resp error %s", string(body))
	}

	return nil
}

func changeAccountName(cmd Command) error {
	request := dto.ChangeAccountNameRequest{
		Name:    cmd.Name,
		NewName: cmd.NewName,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("http://%s:%d/account/patch", cmd.Host, cmd.Port), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http patch failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}
		return fmt.Errorf("resp error %s", string(body))
	}

	return nil
}

func changeAccountBalance(cmd Command) error {
	request := dto.ChangeAccountBalanceRequest{
		Name:      cmd.Name,
		NewAmount: cmd.NewAmount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(fmt.Sprintf("http://%s:%d/account/change", cmd.Host, cmd.Port), "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}
		return fmt.Errorf("resp error %s", string(body))
	}

	return nil
}
