package main

import (
	// "github.com/gin-gonic/gin"

	"gomod/filerw"
	"gomod/invoicedata"
	"gomod/mylib"
	"gomod/ostest"
	"time"
)

func main() {
	println("hello world")
	println("sum is", mylib.Sum(10, 20))
	// mylib.ContainerTest()
	mylib.TimeTest()
	invoice := &invoicedata.Invoice{
		Id:         1,
		CustomerId: 11,
		Raised:     time.Now(),
		Due:        time.Now().Add(time.Hour),
		Paid:       false,
		Note:       "test json",
		Items: []*invoicedata.Item{
			{
				Id:       "22",
				Price:    22,
				Quantity: 100,
				Note:     "this is a item",
			},
		},
	}
	encodeData, _ := invoice.MarshalJSON()
	newInvoice := &invoicedata.Invoice{}
	newInvoice.UnmarshalJSON(encodeData)

	ostest.OsTest()

	filerw.FileRW()

}
