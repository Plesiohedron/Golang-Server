package main

import (
	"GoCourse/HW-3/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	commandName := flag.String("command", "", "Command to execute")
	accountName := flag.String("name", "", "Account's name")
	accountBalance := flag.Int("balance", 0, "Account's balance")
	accountNewName := flag.String("new_name", "", "New account's name")
	accountNewBalance := flag.Int("new_balance", 0, "New account's balance")

	flag.Parse()

	conn, err := grpc.NewClient("0.0.0.0:4567", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch *commandName {
	case "create":
		createResponse, err := c.CreateAccount(ctx, &proto.CreateAccountRequest{Name: *accountName, Balance: int32(*accountBalance)})
		if err != nil {
			log.Fatalf("Failde to create account: %v", err)
		}
		fmt.Println("Create account response:", createResponse.Message)

	case "get":
		getResponse, err := c.GetAccount(ctx, &proto.GetAccountRequest{Name: *accountName})
		if err != nil {
			log.Fatalf("Failed to get account: %v", err)
		}
		fmt.Printf("Account info: Name=%s, Balance=%d\n", getResponse.Name, getResponse.Balance)

	case "patch":
		updateBalanceResponse, err := c.UpdateBalance(ctx, &proto.UpdateBalanceRequest{Name: *accountName, Balance: int32(*accountNewBalance)})
		if err != nil {
			log.Fatalf("Failed to update balance: %v", err)
		}
		fmt.Println("Patch response:", updateBalanceResponse.Message)

	case "change":
		updateNameResponse, err := c.UpdateAccountName(ctx, &proto.UpdateAccountNameRequest{OldName: *accountName, NewName: *accountNewName})
		if err != nil {
			log.Fatalf("Failed to update account name: %v", err)
		}
		fmt.Println("Change account name response:", updateNameResponse.Message)

	case "delete":
		deleteResponse, err := c.DeleteAccount(ctx, &proto.DeleteAccountRequest{Name: *accountName})
		if err != nil {
			log.Fatalf("Failed delete account: %v", err)
		}
		fmt.Println("Delete account response:", deleteResponse.Message)

	default:
		fmt.Println("Invalid operation")
	}
}
