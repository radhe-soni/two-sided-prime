package dbclient

import (
        "github.com/radhe-soni/two-sided-prime/product-plan/model/models.go"
)

type IBoltClient interface {
        OpenBoltDb()
        QueryAccount(accountId string) (model.Account, error)
        Seed()
}