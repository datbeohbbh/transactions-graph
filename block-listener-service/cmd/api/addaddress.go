package main

import (
	pb "block-listener/listener"
	"context"
	"errors"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func (bl *BlockListener) GetAccountType(addr string) (string, error) {
	contractSize, err := bl.ethClient.CodeAt(context.Background(), common.HexToAddress(addr), nil)
	if err != nil {
		return "", err
	}
	if len(contractSize) > 0 {
		return CONTRACT_ACCOUNT, nil
	}
	return USER_ACCOUNT, nil
}

func (bl *BlockListener) AddAddress(ctx context.Context, request *pb.Address) (*pb.AddResp, error) {
	re, _ := regexp.Compile("0x[0-9a-zA-Z]{40}")

	addr := request.Address
	if !re.MatchString(addr) {
		return nil, errors.New("provided address is not valid ethereum address")
	}

	addr = common.HexToAddress(request.Address).Hex()

	if has, err := bl.HasKey(addr); err != nil {
		log.Println("failed on read level db")
		return nil, errors.New("failed on read level db")
	} else if has {
		return &pb.AddResp{
			Msg:    addr + " has already been added",
			Status: true,
		}, nil
	}

	accountType, err := bl.GetAccountType(addr)
	if err != nil {
		return nil, errors.New("failed on check whether address is user or contract account")
	}

	if err := bl.PutKey(addr, accountType); err != nil {
		log.Println("failed on add new key")
		return nil, errors.New("failed on add new key")
	}

	return &pb.AddResp{
		Msg:    "Address added and now it is being tracked",
		Status: true,
	}, nil
}

func (bl *BlockListener) HasKey(addr string) (bool, error) {
	return bl.levdb.Has([]byte(addr))
}

func (bl *BlockListener) PutKey(addr, accountType string) error {
	return bl.levdb.Put([]byte(addr), []byte(accountType))
}
