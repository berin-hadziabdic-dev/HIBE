package searchcontract

import (
	//"database/sql" // for connecting to local mysql
	//_ "github.com/go-sql-driver/mysql" // for connecting to local mysql
	"encoding/json" //JSON apis
	"fmt"
	"net/http" //http APIs
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql" // Connecting to GC mysql by proxy
)

type contractID struct {
	Name string `json:"contractName"`
}

//ContractVulnerabilityInfo stores
type ContractVulnerabilityInfo struct {
	ContractAddress         string `json:"name"`
	IntegerOverflow         string   `json:"integerOverflow"`
	IntegerUnderflow        string   `json:"integerUnderflow"`
	DOS                     string   `json:"dos"`
	ExceptionState          string   `json:"exceptionState"`
	ExternalCall            string   `json:"externalCall"`
	ExternalCallFixed       string   `json:"externalCallFixed"`
	MultipleCalls           string   `json:"multipleCalls"`
	DelegateCall            string   `json:"delegateCall"`
	PredictableEnv          string   `json:"predictableEnv"`
	TxOrigin                string   `json:"txOrigin"`
	EtherWithdrawal         string   `json:"etherWithdrawal"`
	StateChange             string   `json:"stateChange"`
	UnprotectedSelfdestruct string   `json:"unprotectedSelfdestruct"`
	UncheckedCall           string   `json:"uncheckedCall"`
	Scanned                 bool   //A flag used to communicate whether object contains a valid scanned row.*/
	Error                   error  `json:"error"`
}

//SearchContract is a handler that queries the DB for compromised contracts.
func SearchContract(w http.ResponseWriter, r *http.Request) {

	var contractName contractID                             //Used for parsing in contractName
	err := json.NewDecoder((*r).Body).Decode(&contractName) //decode data from body and store into contract

	if err == nil {
	  queryResult := getRow(&contractName.Name)
		if (*queryResult).Scanned {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(queryResult)
		} else {
			w.WriteHeader(204)
			w.Write(nil)
		}

	} else {
		fmt.Println(err.Error())

		w.WriteHeader(500)
	}
}

//processRow queries the DB for a contract with ID value of name.
func getRow(contractName *string) *ContractVulnerabilityInfo {
	cfg := mysql.Cfg("haveibeenexploited:us-west1:hibe", "root", "root")
	cfg.DBName = "mythril" // TODO -- unhardcode this when you add more
	db, err := mysql.DialCfg(cfg)

	var storage ContractVulnerabilityInfo                               //stores row to encode

	storage.Scanned = false //set scanned to false

	if err == nil {
		scanErr := db.QueryRow("SELECT * FROM contracts WHERE ContractAddress=?;", &contractName).Scan(&storage.ContractAddress, &storage.IntegerOverflow, &storage.IntegerUnderflow, &storage.DOS, &storage.ExceptionState, &storage.ExternalCall, &storage.ExternalCallFixed, &storage.MultipleCalls, &storage.DelegateCall, &storage.PredictableEnv, &storage.TxOrigin, &storage.EtherWithdrawal, &storage.StateChange, &storage.UnprotectedSelfdestruct, &storage.UncheckedCall)

		if scanErr == nil {
			storage.Scanned = true //row scanned in set scanned to true
		}
	} else {
		storage.Error = err
		fmt.Println(err.Error)
	}
	return &storage
}
