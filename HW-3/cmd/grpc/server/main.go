package main

import (
	"GoCourse/HW-3/proto"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"log"
	"net"
)

const connectionString = "host=0.0.0.0 port=5432 user=postgres dbname=postgres password=postgres"

type server struct {
	proto.AccountServiceServer
	db *sql.DB
}

func (s *server) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {

	_, err := s.db.ExecContext(ctx, "INSERT INTO accounts (name, balance) VALUES ($1, $2)", req.Name, req.Balance)
	if err != nil {
		log.Printf("Failed to create account: %v", err)
		return nil, fmt.Errorf("failed to create account: %v", err)
	}

	return &proto.CreateAccountResponse{Message: "Account created successfully"}, nil
}

func (s *server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	var name string
	var balance int32
	err := s.db.QueryRowContext(ctx, "SELECT name, balance FROM accounts WHERE name = $1", req.Name).Scan(&name, &balance)
	if err != nil {
		log.Printf("Failed to get account: %v", err)
		return nil, fmt.Errorf("failed to get account: %v", err)
	}

	return &proto.GetAccountResponse{Name: name, Balance: balance}, nil
}

func (s *server) UpdateBalance(ctx context.Context, req *proto.UpdateBalanceRequest) (*proto.UpdateBalanceResponse, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE name = $2", req.Balance, req.Name)
	if err != nil {
		log.Printf("Failed to update balance: %v", err)
		return nil, fmt.Errorf("failed to update balance: %v", err)
	}

	return &proto.UpdateBalanceResponse{Message: "Balance updated successfully"}, nil
}

func (s *server) UpdateAccountName(ctx context.Context, req *proto.UpdateAccountNameRequest) (*proto.UpdateAccountNameResponse, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE accounts SET name = $1 WHERE name = $2", req.NewName, req.OldName)
	if err != nil {
		log.Printf("Failed to update account name: %v", err)
		return nil, fmt.Errorf("failed to update account name: %v", err)
	}

	return &proto.UpdateAccountNameResponse{Message: "Account name updated successfully"}, nil
}

func (s *server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.DeleteAccountResponse, error) {
	_, err := s.db.ExecContext(ctx, "DELETE FROM accounts WHERE name = $1", req.Name)
	if err != nil {
		log.Printf("Failed to delete account: %v", err)
		return nil, fmt.Errorf("failed to delete account: %v", err)
	}

	return &proto.DeleteAccountResponse{Message: "Account deleted successfully"}, nil
}

func main() {

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	defer func() {
		_ = db.Close()
	}()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterAccountServiceServer(s, &server{db: db})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
