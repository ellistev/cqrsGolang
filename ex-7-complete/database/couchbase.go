package database

import (
	"github.com/couchbase/gocb/v2"
	"strings"
	"time"
)

func Connect() *gocb.Cluster{
	var db, err = gocb.Connect(
		"localhost",
		gocb.ClusterOptions{
			Username: "Administrator",
			Password: "password",
		})
	if err != nil {
		panic(err)
	}
	return db
}

func GetCollection(bucketName string) (error, *gocb.Collection) {
	bucket, err := GetBucket(bucketName)
	collection := bucket.DefaultCollection()
	return err, collection
}

func GetCluster() (*gocb.Cluster) {
	var db = Connect()
	db.WaitUntilReady(2*time.Second, nil)

	return db
}

func GetBucket(bucketName string) (*gocb.Bucket, error) {
	var db = Connect()
	db.WaitUntilReady(2*time.Second, nil)

	bucket := db.Bucket(strings.ToLower(bucketName))
	var err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		bucketMgr := db.Buckets()
		err := bucketMgr.CreateBucket(gocb.CreateBucketSettings{
			BucketSettings: gocb.BucketSettings{
				Name:                 strings.ToLower(bucketName),
				FlushEnabled:         true,
				ReplicaIndexDisabled: false,
				RAMQuotaMB:           200,
				NumReplicas:          1,
				BucketType:           gocb.CouchbaseBucketType,
			},
			ConflictResolutionType: gocb.ConflictResolutionTypeSequenceNumber,
		}, nil)
		if err != nil {
			panic(err)
		}

		bucket = db.Bucket(strings.ToLower(bucketName))
		err = bucket.WaitUntilReady(5*time.Second, nil)
	}
	return bucket, err
}
