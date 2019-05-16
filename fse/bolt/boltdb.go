package bolt

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	Db   *bolt.DB
	Path string
}

func InitDB(dbPath string) *BoltDB {
	log.Println("init bolt")
	bdb := &BoltDB{Path: dbPath}
	var err error
	bdb.Db, err = bolt.Open(dbPath, 0777, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatalf("can not init BoltDB - %s\n", err.Error())
	}
	return bdb
}

func (this *BoltDB) ApplyId(table string) (string, error) {
	var id string
	err := this.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			//log.Println(err)
			return err
		}
		seq, err := b.NextSequence()
		id = strconv.Itoa(int(seq))
		if err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func (this *BoltDB) GId(table string) (int, error) {
	var id int
	err := this.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			//log.Println(err)
			return err
		}
		seq, err := b.NextSequence()
		id = int(seq)
		if err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func (this *BoltDB) CloseDB() error {
	if this.Db != nil {
		err := this.Db.Close()
		return err
	}
	return errors.New("The DB is null")
}

func (this *BoltDB) NewTable(name string) error {
	err := this.Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return err
}

func (this *BoltDB) GetValue(table []byte, key []byte) ([]byte, error) {
	var result []byte
	err := this.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return errors.New("The DB is Null")
		}
		result = b.Get([]byte(key))
		return nil
	})
	return result, err
}

func (this *BoltDB) IfTableExist(table []byte) bool {
	var res bool
	_ = this.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(table)
		if b == nil {
			res = false
			return nil
		}
		res = true
		return nil
	})
	return res
}

func (this *BoltDB) IfValueExist(table []byte, key []byte) bool {
	var result []byte
	var res bool
	_ = this.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return errors.New("The DB is Null")
		}
		result = b.Get([]byte(key))
		if len(result) <= 0 || result == nil {
			res = false
			return nil
		}
		res = true
		return nil
	})
	return res

}

func (this *BoltDB) SetValue(table []byte, key []byte, val []byte) error {
	err := this.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			//log.Println(err)
			return err
		}
		err = b.Put([]byte(key), val)
		return err
	})
	return err
}

func (this *BoltDB) DeleteValue(table string, key string) error {
	err := this.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			//log.Println(err)
			return err
		}
		err = b.Delete([]byte(key))
		return err
	})
	return err
}

func (this *BoltDB) GetTable(table string) (map[string][]byte, error) {
	var result = make(map[string][]byte, 0)
	err := this.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return errors.New("The DB is Null")
		}
		b.ForEach(func(k, v []byte) error {
			//fmt.Printf("key=%s, value=%s\n", k, v)
			result[string(k)] = v
			return nil
		})
		return nil
	})
	return result, err
}

func (this *BoltDB) GetTableKeys(table string) ([]string, error) {
	var result = make([]string, 0)
	err := this.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return errors.New("The DB is Null")
		}
		b.ForEach(func(k, v []byte) error {
			result = append(result, string(k))
			return nil
		})

		return nil
	})
	return result, err
}

func (this *BoltDB) GetTableValues(table string) ([]([]byte), error) {
	var result = make([]([]byte), 0)
	err := this.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return errors.New("The DB is Null")
		}
		b.ForEach(func(k, v []byte) error {
			result = append(result, v)
			return nil
		})
		return nil
	})
	return result, err
}

func (this *BoltDB) SetValueAuto(table []byte, val []byte) (string, error) {
	var id = -1
	err := this.Db.Update(func(tx *bolt.Tx) error {
		//println(111)
		b, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			return errors.New("create bucket:" + err.Error())
		}
		seq, err := b.NextSequence()
		id = int(seq)
		if err != nil {
			return err
		}
		err = b.Put([]byte(strconv.Itoa(id)), val)
		if err != nil {
			return errors.New("put into the table:" + err.Error())
		}
		return nil
	})
	return string(strconv.Itoa(id)), err
}

///SetValueAuto
