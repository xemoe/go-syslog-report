package cache

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger"
	types "github.com/xemoe/go-syslog-report/types"
	"io"
	"log"
	"os"
)

var CacheDir = "/tmp"
var CachePrefix = "cache"

func MakeJsonString(item interface{}) string {
	j, _ := json.Marshal(item)
	return string(j)
}

func MakeSha1Hash(item interface{}) string {
	j, _ := json.Marshal(item)
	h := sha1.New()
	h.Write([]byte(j))

	return hex.EncodeToString(h.Sum(nil))
}

func GetDB(prefixstr string, filehash string) *badger.DB {
	opts := badger.DefaultOptions
	opts.Dir = fmt.Sprintf("%s/%s-%s-%s", CacheDir, CachePrefix, prefixstr, filehash)
	opts.ValueDir = fmt.Sprintf("%s/%s-%s-%s", CacheDir, CachePrefix, prefixstr, filehash)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetFileChecksum(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func GetCache(filename string, index types.SyslogLineIndex) (string, string, string, map[string]int, error) {

	result := make(map[string]int)

	indexhash := MakeSha1Hash(index)
	filenamehash := MakeSha1Hash(filename)
	filechecksum := GetFileChecksum(filename)

	//
	// Read cache
	//
	var valCopy []byte
	db := GetDB(indexhash, filenamehash)
	defer db.Close()

	err := db.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(filechecksum))
		if err != nil {
			return err
		}

		val, err := item.Value()
		if err != nil {
			return err
		}

		valCopy = val

		return nil
	})

	if err == nil {
		log.Printf("Get cache %s, %s", filechecksum, valCopy)
		err = json.Unmarshal(valCopy, &result)
	}

	return indexhash, filenamehash, filechecksum, result, err
}

func SaveCache(indexhash string, filenamehash string, filechecksum string, result map[string]int) error {

	db := GetDB(indexhash, filenamehash)
	defer db.Close()
	err := db.Update(func(txn *badger.Txn) error {
		valuejson := MakeJsonString(result)
		err := txn.Set([]byte(filechecksum), []byte(valuejson))
		log.Printf("Set cache %s, %s", filechecksum, valuejson)
		return err
	})

	return err
}
