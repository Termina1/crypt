package main

import (
  "github.com/boltdb/bolt"
  "github.com/satori/go.uuid"
)

const bucket string = "secrets";

func storeAndLink(db *bolt.DB, secret string, salt string) (string, error) {

  uidU := uuid.NewV4()
  uid := uidU.String()
  errbd := db.Update(func(tx *bolt.Tx) error {
    b, berr := tx.CreateBucketIfNotExists([]byte(bucket))
    if berr != nil {
      return berr
    }
    err := b.Put([]byte(uid), []byte(secret))
    if err != nil {
      return err
    }
    err = b.Put([]byte(uid + "_salt"), []byte(salt))
    return err
  })
  return uid, errbd
}

func readAndDelete(db *bolt.DB, uid string) (string, string, error) {
  var copyDest []byte;
  var copySalt []byte;
  err := db.Batch(func(tx *bolt.Tx) error {
    b, berr := tx.CreateBucketIfNotExists([]byte(bucket))
    if berr != nil {
      return berr
    }
    result := b.Get([]byte(uid))
    salt := b.Get([]byte(uid + "_salt"))
    err := b.Delete([]byte(uid))
    // bolt will reuse memory after transation, we need to topy it
    copyDest = make([]byte, len(result), (cap(result)+1)*2)
    copySalt = make([]byte, len(salt), (cap(salt)+1)*2)
    copy(copyDest, result)
    copy(copySalt, salt)
    return err
  })
  return string(copyDest), string(copySalt), err
}
