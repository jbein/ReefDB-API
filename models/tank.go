package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Tank struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Volume       uint      `json:"volume"`
	Location     string    `json:"location"`
	Description  string    `json:"description"`
	Manufacturer string    `json:"manufacturer"`
	Model        string    `json:"model"`
	CreatedAt    time.Time `json:"_"`
	UpdatedAt    time.Time `json:"_"`
}

func GetAllTanks() []Tank {
	var tanks []Tank
	jsonFile, err := os.Open("tanks.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonParser := json.NewDecoder(jsonFile)
	jsonParser.Decode(&tanks)
	return tanks
}

func GetTank(id uint) (*Tank, bool) {
	jsonFile, err := os.Open("tanks.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var tanks []Tank
	jsonParser := json.NewDecoder(jsonFile)
	jsonParser.Decode(&tanks)
	return findTank(tanks, id)
}

func findTank(tanks []Tank, id uint) (*Tank, bool) {
	for _, v := range tanks {
		if v.ID == id {
			return &v, true
		}
	}
	return nil, false
}

func AddNewTank(tank Tank) (*Tank, bool) {
	jsonFile, err := os.OpenFile("tanks.json", os.O_APPEND, 0664)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	encoder := json.NewEncoder(jsonFile)
	encoder.Encode(&tank)

	return GetTank(tank.ID)
}
