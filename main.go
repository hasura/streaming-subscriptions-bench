package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hasura/go-graphql-client"
)

// subscription {
//  customer_stream(cursor: {c_custkey: 1}, batch_size: 1) {
//    c_custkey
//    c_acctbal
//  }
//}

type Customer struct {
	CCustkey graphql.Int   `graphql:"c_custkey" json:"c_custkey"`
	CAcctbal graphql.Float `graphql:"c_acctbal" json:"c_acctbal"`
}

type subscription struct {
	CustomerStream []Customer `graphql:"customer_stream(cursor: {c_custkey: 0}, batch_size: 100)" json:"customer_stream"`
}

type CustomerMax struct {
	CustomerAggregate struct {
		Aggregate struct {
			Max struct {
				CCustKey graphql.Int `graphql:"c_custkey" json:"c_custkey"`
			}
		}
	} `graphql:"customer_aggregate"`
}

func main() {

	client := graphql.NewSubscriptionClient("ws://localhost:8080/v1/graphql")
	defer client.Close()

	a := make([]Customer, 100)
	stop := make(chan bool)

	subscriptionId, err := client.Subscribe(&subscription{}, nil, func(message *json.RawMessage, errValue error) error {
		if errValue != nil {
			// handle error
			// if returns error, it will failback to `onError` event
			fmt.Println(errValue)
			return nil
		}
		var data subscription
		if err := json.Unmarshal(*message, &data); err != nil {
			fmt.Println(err)
			return nil
		}

		a = append(a, data.CustomerStream...)
		if len(a) > 100000 {
			stop <- true
		}
		return nil

	})
	if err != nil {
		fmt.Println(err)
	}

	queryClient := graphql.NewClient("http://localhost:8080/v1/graphql", nil)

	var customerMax CustomerMax

	err = queryClient.Query(context.Background(), &customerMax, nil)
	if err != nil {
		fmt.Println(err)
	}

	concurrency := 1009
	initKey := int(customerMax.CustomerAggregate.Aggregate.Max.CCustKey) + 1
	fmt.Println(initKey)

	startConcurrentMutations(concurrency, initKey)

	go client.Run()

	select {
	case <-stop:
		client.Unsubscribe(subscriptionId)
		fmt.Println("done")
	}
}

func startConcurrentMutations(concurrency int, initKey int) {

	for i := 0; i < concurrency; i++ {
		go runMutations(i, initKey, concurrency)
	}

}

func runMutations(seed int, initKey int, step int) {

	client := graphql.NewClient("http://localhost:8080/v1/graphql", nil)

	var m struct {
		InsertCustomerOne struct {
			CAcctbal    graphql.Float  `graphql:"c_acctbal" json:"c_acctbal"`
			CAddress    graphql.String `graphql:"c_address" json:"c_address"`
			CComment    graphql.String `graphql:"c_comment" json:"c_comment"`
			CCustkey    graphql.Int    `graphql:"c_custkey" json:"c_custkey"`
			CMktsegment graphql.String `graphql:"c_mktsegment" json:"c_mktsegment"`
			CName       graphql.String `graphql:"c_name" json:"c_name"`
			CNationkey  graphql.Int    `graphql:"c_nationkey" json:"c_nationkey"`
			CPhone      graphql.String `graphql:"c_phone" json:"c_phone"`
		} `graphql:"insert_customer_one(object: {c_acctbal: 100.0, c_address: \"abcd, new york\", c_comment: \"azsdfg\", c_custkey: $custKey, c_mktsegment: \"BUILDING\", c_name: $custName, c_nationkey: 15, c_phone: \"25-989-741-2988\"})" json:"insert_customer_one"`
	}

	var variables map[string]interface{}
	currentKey := initKey + seed
	for {
		variables = map[string]interface{}{
			"custKey":  graphql.Int(currentKey),
			"custName": graphql.String(fmt.Sprintf("Customer#%d", currentKey)),
		}

		err := client.Mutate(context.Background(), &m, variables)
		if err != nil {
			fmt.Println(err)
		}
		currentKey = currentKey + step
	}
}

// client := graphql.NewClient("http://localhost:8080/v1/graphql", nil)
// err := client.Query(context.Background(), &query, nil)
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// fmt.Println(query.Me)
