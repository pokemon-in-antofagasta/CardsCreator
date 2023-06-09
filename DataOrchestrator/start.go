package data_orchestrator

import (
	"errors"
	"fmt"
	"os"

	model "example.com/AlejandroWaiz/DataManager/Model"
)

var isAnErrorHappen bool

func (o *DataOrchestratorImplementation) Start() error {

	switch o.excelName {

	case os.Getenv("pokemons_excel_filename"):

		allData, err := o.datareader.ReadExcelAndReturnData()

		if err != nil {
			isAnErrorHappen = true
			return fmt.Errorf("[ReadExcelAndStoreDataItIntoDatabase] => %v", err)
		}

		errors := o.datasaver.SaveAllPokemonsIntoFirestore(allData.([]model.PokemonCard))

		if len(errors) > 0 {
			isAnErrorHappen = true
			return fmt.Errorf("Error saving this pokemons: %v", errors)
		}

	case os.Getenv("movements_excel_filename"):

		allData, err := o.datareader.ReadExcelAndReturnData()

		if err != nil {
			isAnErrorHappen = true
			return fmt.Errorf("[ReadExcelAndStoreDataItIntoDatabase] => %v", err)
		}

		errors := o.datasaver.SaveAllMovementsIntoFirestore(allData.([]model.MovementCard))
		if len(errors) > 0 {
			isAnErrorHappen = true
			return fmt.Errorf("Error saving this movements: %v", errors)
		}

	case os.Getenv("items_excel_filename"):

		allData, err := o.datareader.ReadExcelAndReturnData()

		if err != nil {
			isAnErrorHappen = true
			return fmt.Errorf("[ReadExcelAndStoreDataItIntoDatabase] => %v", err)
		}

		errors := o.datasaver.SaveAllItemsIntoFirestore(allData.([]model.ItemCard))

		if len(errors) > 0 {
			isAnErrorHappen = true
			return fmt.Errorf("Error saving this items: %v", errors)
		}

	case os.Getenv("weathers_excel_filename"):

		allData, err := o.datareader.ReadExcelAndReturnData()

		if err != nil {
			isAnErrorHappen = true
			return fmt.Errorf("[ReadExcelAndStoreDataItIntoDatabase] => %v", err)
		}

		errors := o.datasaver.SaveAllWeathersIntoFirestore(allData.([]model.WeatherCard))

		if len(errors) > 0 {
			isAnErrorHappen = true
			return fmt.Errorf("Error saving this weathers: %v", errors)
		}

	default:
		isAnErrorHappen = true
		return errors.New("Invalid excel name")

	}

	err := o.storeExcelIntoBucket()

	if err != nil {
		return fmt.Errorf("[StoreExcelIntoBucket] err : %v", err)
	}

	return nil

}
