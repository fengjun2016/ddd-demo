package app

import (
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

var Es *elastic.Client

func InitElastic() {
	host := Config.Elastic.Addr
	fmt.Println(host)
	var err error
	Es, err = elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println("create elastic client error", err)
		return
	}
	fmt.Println("create elastic client success")
}
