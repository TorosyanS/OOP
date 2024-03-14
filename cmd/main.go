package main

import (
	"errors"
	"fmt"
	"os"

	"time"

	"test/internal/embedding/training"
	"test/internal/encapculation/bank"
	"test/internal/polymorphism/storage"
	"test/internal/polymorphism/storage/file_storage"
	"test/internal/polymorphism/storage/map_storage"
	"test/internal/polymorphism/storage/slice_storage"
)

func main() {
	// ИНКАПСУЛЯЦИЯå
	//encapsulation()

	// ВСТРАИВАНИЕ
	//embedding()

	// ПОЛИМОРФИЗМ
	//polymorphism()
}

func encapsulation() {
	transferBank := bank.NewBank()

	transferBank.CreateNewAccount("200")
	err := transferBank.TransferMoney("100", "200", 500)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	account1, err := transferBank.FindAccount("100")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	account2, err := transferBank.FindAccount("200")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Balance %s - %d\n", account1.Number(), account1.Balance())
	fmt.Printf("Balance %s - %d\n", account2.Number(), account2.Balance())
}

func embedding() {
	swimming := training.Swimming{
		Training: training.Training{
			TrainingType: "Плавание",
			Action:       2000,
			LenStep:      training.SwimmingLenStep,
			Duration:     90 * time.Minute,
			Weight:       85,
		},
		LengthPool: 50,
		CountPool:  5,
	}

	fmt.Println(training.ReadData(swimming))

	walking := training.Walking{
		Training: training.Training{
			TrainingType: "Ходьба",
			Action:       20000,
			LenStep:      training.LenStep,
			Duration:     3*time.Hour + 45*time.Minute,
			Weight:       85,
		},
		Height: 185,
	}

	fmt.Println(training.ReadData(walking))

	running := training.Running{
		Training: training.Training{
			TrainingType: "Бег",
			Action:       5000,
			LenStep:      training.LenStep,
			Duration:     30 * time.Minute,
			Weight:       85,
		},
	}

	fmt.Println(training.ReadData(running))
}

func polymorphism() {
	store, err := createStorage(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer store.Close()

	err = store.SavePair("hello", "bye")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	value, err := store.GetValue("hello")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(value)

	value, err = store.GetValue("hello_unknown")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(value)
}

func createStorage(storageType string) (storage storage.Storage, err error) {
	switch storageType {
	case "map":
		storage = map_storage.NewStorage()
	case "slice":
		storage = slice_storage.NewStorage()
	case "file":
		storage, err = file_storage.NewStorage()
	default:
		err = errors.New("unknown storage type")
	}
	return
}
