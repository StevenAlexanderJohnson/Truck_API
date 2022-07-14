package convert_handler

import (
	struct_def "Maria_Demo/Structs"
	"database/sql"
	"encoding/json"
)

func Dataset_To_Truck_Data(data *sql.Rows) ([]struct_def.Truck_Data, error) {
	var output []struct_def.Truck_Data = []struct_def.Truck_Data{}
	for data.Next() {
		truck_data := struct_def.Truck_Data{}
		err := data.Scan(&truck_data.Truck_Id, &truck_data.Truck_Number, &truck_data.Truck_Type, &truck_data.Truck_Price)
		if err != nil {
			return nil, err
		}
		output = append(output, truck_data)
	}
	return output, nil
}

func Dataset_To_Roles(data *sql.Rows) ([]string, error) {
	var output []string
	for data.Next() {
		var role string
		err := data.Scan(&role)
		if err != nil {
			return nil, err
		}
		output = append(output, role)
	}
	return output, nil
}

func Struct_ToString(s interface{}) (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		// This should never be hit but if it is there is a problem.
		return "", err
	}
	return string(b), nil
}
