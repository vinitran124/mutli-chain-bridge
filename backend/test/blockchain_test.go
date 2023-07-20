package test

import (
	"bridge/app/blockchain"
	"bridge/app/model"
	"fmt"
	"github.com/joho/godotenv"
	"testing"
	"time"
)

func TestTrackingData(t *testing.T) {
	if err := godotenv.Overload(".env"); err != nil {
		fmt.Println("Load env error", err.Error())
	}

	chain, err := blockchain.NewChain("5555")
	if err != nil {
		t.Error(err)
	}

	events := make(chan model.DepositEvent)

	err = chain.TrackDeposit(events)
	if err != nil {
		t.Error(err)
	}

	for {
		select {
		case event := <-events:
			fmt.Println(event)
		default:
			time.Sleep(time.Second)
		}
	}
}

func TestCallWithdrawal(t *testing.T) {
	if err := godotenv.Overload(".env"); err != nil {
		fmt.Println("Load env error", err.Error())
	}

	chain, err := blockchain.NewChain("5555")
	if err != nil {
		t.Error(err)
	}

	VINI_CONTRACT := "0x67bCBB439aeBF339a99166367676F047Cef38BEb"
	err = chain.CallWithdrawal(VINI_CONTRACT, "0xc46cF4Faa3D99EbE5b35bC659C9a01590B5C300B", "10000000000000000000")
	if err != nil {
		t.Error(err)
	}
}

func TestGetTokenAmount(t *testing.T) {
	if err := godotenv.Overload(".env"); err != nil {
		fmt.Println("Load env error", err.Error())
	}

	chain, err := blockchain.NewChain("5555")
	if err != nil {
		t.Error(err)
	}

	VINI := "0x596c14ba2b6e759c73895ce13e64e49054181a7f"
	amount, err := chain.GetTokenInPool(VINI)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(amount)
}
