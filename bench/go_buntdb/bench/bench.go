package main

import (
	"log"
	"strconv"
	"math/rand"
    "fmt" 
	"github.com/tidwall/buntdb"
)

func RandString() string {
    return strconv.FormatUint(rand.Uint64(), 16)
}

func main() {
	keys_len := 10000
	var keys = [10000]string{}
	err_num := 0
	put_num := 0
	get_num := 0

	// fmt.Printf("%s\n", RandString())
	// vector.
	// Open the data.db file. It will be created if it doesn't exist.
	db, err := buntdb.Open(":memory:")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i <1000000; i++ {
		if (rand.Int() & 1 == 0) {
			put_num += 3
			db.Update(func(tx *buntdb.Tx) error {
				for count := 0; count <= 2; count++ {
					s := RandString()
					_, _, err := tx.Set(s, RandString(), nil)
					if err != nil {
						return err
					}
					keys[rand.Int() % keys_len] = s
				}
				return err
			})
		} else {
			get_num += 7
			db.View(func(tx *buntdb.Tx) error {
				for count := 0; count <= 6; count++ {
					s := keys[rand.Int() % keys_len]
					_, err := tx.Get(s)
					if err != nil{
						err_num ++;
						// return err
					} else {
						// fmt.Printf("value is %s\n", val)
					}
				}
				return nil
			})
		}
	}

	fmt.Printf("put_num %d\n", put_num);
	fmt.Printf("get_num %d\n", get_num);
	fmt.Printf("err_num %d\n", err_num);

	defer db.Close()

}

