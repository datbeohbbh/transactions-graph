package worker

import (
	"context"
	"errors"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func (bl *Worker) GetAccountType(addr string) (string, error) {
	contractSize, err := bl.ethClient.CodeAt(context.Background(), common.HexToAddress(addr), nil)
	if err != nil {
		return "", err
	}
	if len(contractSize) > 0 {
		return CONTRACT_ACCOUNT, nil
	}
	return USER_ACCOUNT, nil
}

func (bl *Worker) Tracking(ctx context.Context, address string) (bool, error) {
	return bl.dao.ExistedAddress(ctx, "tracking", address)
}

func (bl *Worker) InsertTrackedAddress(ctx context.Context, address, accountType string) (bool, error) {
	return bl.dao.InsertTrackedAddress(ctx, address, accountType)
}

func (bl *Worker) AddAddress(ctx context.Context, request *Address) (*AddResp, error) {
	re, _ := regexp.Compile("0x[0-9a-zA-Z]{40}")

	addr := request.Address
	if !re.MatchString(addr) {
		return nil, errors.New("provided address is not valid ethereum address")
	}

	addr = common.HexToAddress(request.Address).Hex()

	accountType, err := bl.GetAccountType(addr)
	if err != nil {
		return nil, errors.New("failed on check whether address is user or contract account")
	}

	addStatus, err := bl.dao.InsertTrackedAddress(ctx, addr, accountType)
	if err != nil {
		return nil, err
	}

	if addStatus {
		log.Printf("successfully add address %s", addr)
		return &AddResp{
			Msg:    "Address added and now it is being tracked",
			Status: true,
		}, nil
	} else {
		return &AddResp{
			Msg:    addr + " has already been added",
			Status: true,
		}, nil
	}
}
