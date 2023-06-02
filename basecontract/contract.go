package basecontract

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric/common/flogging"
)

// BaseContract defines the utility functions that are provided by the Spydra base smart contract for Hyperledger Fabric.
type SpydraContract struct {
	contractapi.Contract
}

// PaginatedQueryResult is a type that holds the various attributes that will be returned in the response to a query.
type PaginatedQueryResult struct {
	Count    int32       `json:"count,omitempty" metadata:",optional"`
	Bookmark string      `json:"bookmark,omitempty" metadata:",optional"`
	Records  interface{} `json:"records"`
}

// ReadDataFromQueryString executes a Couch DB rich query and provides paginated response.
// The function accepts a Couch DB rich query string, the page size for the results and a bookmark.
func (contract *SpydraContract) ReadDataFromQueryString(ctx contractapi.TransactionContextInterface, queryString string, pageSize int32, bookmark string) (response *PaginatedQueryResult, err error) {
	logger := flogging.MustGetLogger("BaseContract")
	stub := ctx.GetStub()

	var queryData []interface{}

	queryIterator, queryMetadata, err := stub.GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		logger.Errorf("Query failed. Error: %+v", err.Error())
		return
	}

	defer queryIterator.Close()

	for queryIterator.HasNext() {
		query, queryErr := queryIterator.Next()
		if queryErr != nil {
			err = queryErr
			logger.Errorf("Failed to fetch next result. Error: %+v", err.Error())
			return
		}
		var value interface{}
		err = json.Unmarshal(query.Value, &value)
		if err != nil {
			logger.Errorf("Unmarshalling result value to Json failed. Error: %+v", err.Error())
			return
		}

		queryData = append(queryData, value)
	}

	response = new(PaginatedQueryResult)
	response.Count = queryMetadata.FetchedRecordsCount
	if queryMetadata.FetchedRecordsCount >= pageSize {
		response.Bookmark = queryMetadata.Bookmark
	}
	response.Records = queryData

	return
}
