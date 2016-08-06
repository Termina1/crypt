package main

import (
  "github.com/boltdb/bolt"
  "github.com/satori/go.uuid"
)

const bucket string = "secrets";

func storeAndLink(db *bolt.DB, secret string) (string, error) {

  uid := uuid.NewV4().String()
  err := db.Update(func(tx *bolt.Tx) error {
    tx.CreateBucket([]byte(bucket))
    b := tx.Bucket([]byte(bucket))
    err := b.Put([]byte(uid), []byte(secret))
    return err
  })
  return uid, err
}

func readAndDelete(db *bolt.DB, uid string) (string, error) {
  var result []byte;
  var copyDest []byte;
  err := db.Batch(func(tx *bolt.Tx) error {
    tx.CreateBucket([]byte(bucket))
    b := tx.Bucket([]byte(bucket))
    result = b.Get([]byte(uid))
    err := b.Delete([]byte(uid))
    // bolt will reuse memory after transation, we need to topy it
    copyDest = make([]byte, len(result), (cap(result)+1)*2)
    copy(copyDest, result)
    return err
  })
  return string(copyDest), err
}
