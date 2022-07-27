package Database

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Initialize_Database(db *sql.DB) error {
	fmt.Println("Executing Schema...")
	err := filepath.Walk("./Database/schema", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		fmt.Printf("Executing: %s\n", info.Name())
		script, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		_, err = db.Exec(string(script))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	fmt.Println("Schema Executed")

	fmt.Println("Initializing Data...")
	data_file, err := os.Open("./Database/load_data.sql")
	if err != nil {
		fmt.Println("ERR: ?", err)
		return err
	}
	defer data_file.Close()

	scanner := bufio.NewScanner(data_file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		_, err := db.Exec(scanner.Text())
		if err != nil {
			return err
		}
	}
	fmt.Println("Data Executed")
	return nil
}
