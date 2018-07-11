// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package integtest

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/objectstorage"
	"github.com/oracle/oci-go-sdk/objectstorage/transfer"
	"github.com/stretchr/testify/assert"
)

func getNamespace(t *testing.T) string {
	c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	request := objectstorage.GetNamespaceRequest{}
	r, err := c.GetNamespace(context.Background(), request)
	failIfError(t, err)
	return *r.Value
}

func getObject(t *testing.T, namespace, bucketname, objectname string) (objectstorage.GetObjectResponse, error) {
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	request := objectstorage.GetObjectRequest{
		NamespaceName: &namespace,
		BucketName:    &bucketname,
		ObjectName:    &objectname,
	}

	return c.GetObject(context.Background(), request)
}

func putObject(t *testing.T, namespace, bucketname, objectname string, contentLen int64, content io.ReadCloser, metadata map[string]string) error {
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	request := objectstorage.PutObjectRequest{
		NamespaceName: &namespace,
		BucketName:    &bucketname,
		ObjectName:    &objectname,
		ContentLength: &contentLen,
		PutObjectBody: content,
		OpcMeta:       metadata,
	}
	_, err := c.PutObject(context.Background(), request)
	return err
}

func createBucket(t *testing.T, namespace, compartment, name string) {
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	request := objectstorage.CreateBucketRequest{
		NamespaceName: &namespace,
	}
	request.CompartmentId = &compartment
	request.Name = &name
	request.Metadata = make(map[string]string)
	request.PublicAccessType = objectstorage.CreateBucketDetailsPublicAccessTypeNopublicaccess
	_, err := c.CreateBucket(context.Background(), request)
	failIfError(t, err)
	return
}

func deleteBucket(t *testing.T, namespace, name string) (err error) {
	err = deleteBucketIgnoreError(t, namespace, name)
	failIfError(t, err)
	return
}

func deleteBucketIgnoreError(t *testing.T, namespace, name string) (err error) {
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	request := objectstorage.DeleteBucketRequest{
		NamespaceName: &namespace,
		BucketName:    &name,
	}
	_, err = c.DeleteBucket(context.Background(), request)
	return
}

func deleteObject(t *testing.T, namespace, bucketname, objectname string) (err error) {
	err = deleteObjectIgnoreError(t, namespace, bucketname, objectname)
	failIfError(t, err)
	return
}

func deleteObjectIgnoreError(t *testing.T, namespace, bucketname, objectname string) (err error) {
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	request := objectstorage.DeleteObjectRequest{
		NamespaceName: &namespace,
		BucketName:    &bucketname,
		ObjectName:    &objectname,
	}
	_, err = c.DeleteObject(context.Background(), request)
	return
}

func TestObjectStorageClient_GetNamespace(t *testing.T) {
	namespace := getNamespace(t)
	assert.NotEmpty(t, namespace)
	return
}

func TestObjectStorageClient_BigFile(t *testing.T) {
	bname := getUniqueName("largeBucket")
	namespace := getNamespace(t)

	createBucket(t, getNamespace(t), getTenancyID(), bname)
	defer deleteBucket(t, namespace, bname)

	contentlen := 1024 * 10000
	filepath, filesize, expectedHash := writeTempFileOfSize(int64(contentlen))
	filename := path.Base(filepath)
	fmt.Println("uploading ", filepath)
	defer removeFileFn(filepath)
	file, e := os.Open(filepath)
	defer file.Close()
	failIfError(t, e)

	e = putObject(t, namespace, bname, filename, filesize, file, nil)
	failIfError(t, e)
	fmt.Println(expectedHash)
	rGet, e := getObject(t, namespace, bname, filename)
	failIfError(t, e)
	defer deleteObject(t, namespace, bname, filename)

	h := sha256.New()
	_, e = io.Copy(h, rGet.Content)
	rGet.Content.Close()
	actualHash := hex.EncodeToString(h.Sum(nil))
	assert.NoError(t, e)
	assert.Equal(t, filesize, int64(*rGet.ContentLength))
	assert.Equal(t, "application/octet-stream", *rGet.ContentType)
	assert.Equal(t, expectedHash, actualHash)
}

func TestObjectStorage_GzipFileEncoding(t *testing.T) {
	bname := getUniqueName("BucketWithZip")
	objname := getUniqueName("gzippedcontent")
	namespace := getNamespace(t)

	createBucket(t, getNamespace(t), getTenancyID(), bname)
	defer deleteBucket(t, namespace, bname)

	message := " some random content that will get gzipped"
	zBytes := bytes.Buffer{}
	gz := gzip.NewWriter(&zBytes)
	gz.Write([]byte(message))
	gz.Close()
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	request := objectstorage.PutObjectRequest{
		NamespaceName:   &namespace,
		BucketName:      &bname,
		ObjectName:      &objname,
		ContentLength:   common.Int64(int64(zBytes.Len())),
		PutObjectBody:   ioutil.NopCloser(&zBytes),
		ContentType:     common.String("text/plain"),
		ContentEncoding: common.String("gzip"),
	}
	_, e := c.PutObject(context.Background(), request)
	defer deleteObject(t, namespace, bname, objname)
	failIfError(t, e)

	response, e := getObject(t, namespace, bname, objname)
	failIfError(t, e)

	//The file is served back by its Content-Type, thus it should not be gzipped anymore
	content, e := ioutil.ReadAll(response.Content)
	failIfError(t, e)
	assert.Equal(t, message, string(content))
}

func TestObjectStorageClient_Object(t *testing.T) {
	bname := getUniqueName("bucket")
	data := "some temp data"
	namespace := getNamespace(t)

	createBucket(t, getNamespace(t), getTenancyID(), bname)
	defer deleteBucket(t, namespace, bname)

	contentlen := int64(len([]byte(data)))
	filepath := writeTempFile(data)
	filename := path.Base(filepath)
	defer removeFileFn(filepath)
	file, e := os.Open(filepath)
	defer file.Close()
	failIfError(t, e)
	metadata := make(map[string]string)
	metadata["Test-VERSION"] = "TestOne"
	e = putObject(t, namespace, bname, filename, contentlen, file, metadata)
	failIfError(t, e)

	r, e := getObject(t, namespace, bname, filename)

	failIfError(t, e)
	assert.Equal(t, "TestOne", r.OpcMeta["test-version"])
	defer deleteObject(t, namespace, bname, filename)
	defer r.Content.Close()
	bytes, e := ioutil.ReadAll(r.Content)
	failIfError(t, e)
	assert.Equal(t, contentlen, *r.ContentLength)
	assert.Equal(t, data, string(bytes))
	return
}

func TestObjectStorageClient_AbortUpload(t *testing.T) {
	bname := getUniqueName("abortUpload")
	namespace := getNamespace(t)
	bgCtx := context.Background()

	createBucket(t, namespace, getTenancyID(), bname)
	defer deleteBucket(t, namespace, bname)

	contentlen := 1024 * 1024 * 10
	filepath, filesize, _ := writeTempFileOfSize(int64(contentlen))
	filename := path.Base(filepath)
	defer removeFileFn(filepath)
	file, e := os.Open(filepath)
	defer file.Close()
	failIfError(t, e)

	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())

	// Clean up the object incase it does get writen to the bucket
	deleteObjectFunc := func() {
		request := objectstorage.DeleteObjectRequest{
			NamespaceName: &namespace,
			BucketName:    &bname,
			ObjectName:    &filename,
		}
		_, _ = c.DeleteObject(bgCtx, request)
	}

	request := objectstorage.PutObjectRequest{
		NamespaceName: &namespace,
		BucketName:    &bname,
		ObjectName:    &filename,
		ContentLength: common.Int64(filesize),
		PutObjectBody: file,
	}
	ctx, cancelFn := context.WithTimeout(bgCtx, 5*time.Millisecond)
	defer cancelFn()
	defer deleteObjectFunc()

	_, e = c.PutObject(ctx, request)
	assert.Error(t, e)
}

func TestObjectStorageClient_AbortMultipartUpload(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.AbortMultipartUploadRequest{}
	_, err := c.AbortMultipartUpload(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_CommitMultipartUpload(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.CommitMultipartUploadRequest{}
	_, err := c.CommitMultipartUpload(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_CreateMultipartUpload(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.CreateMultipartUploadRequest{}
	r, err := c.CreateMultipartUpload(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_CreatePreauthenticatedRequest(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.CreatePreauthenticatedRequestRequest{}
	r, err := c.CreatePreauthenticatedRequest(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_DeletePreauthenticatedRequest(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.DeletePreauthenticatedRequestRequest{}
	_, err := c.DeletePreauthenticatedRequest(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_GetPreauthenticatedRequest(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.GetPreauthenticatedRequestRequest{}
	r, err := c.GetPreauthenticatedRequest(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_HeadBucket(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.HeadBucketRequest{}
	_, err := c.HeadBucket(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_HeadObject(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.HeadObjectRequest{}
	_, err := c.HeadObject(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListBuckets(t *testing.T) {
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.ListBucketsRequest{
		NamespaceName: common.String(getNamespace(t)),
		CompartmentId: common.String(getCompartmentID()),
		Fields:        []objectstorage.ListBucketsFieldsEnum{objectstorage.ListBucketsFieldsTags},
	}
	r, err := c.ListBuckets(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
}

func TestObjectStorageClient_ListMultipartUploadParts(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.ListMultipartUploadPartsRequest{}
	r, err := c.ListMultipartUploadParts(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListMultipartUploads(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.ListMultipartUploadsRequest{}
	r, err := c.ListMultipartUploads(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListObjects(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.ListObjectsRequest{}
	r, err := c.ListObjects(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListPreauthenticatedRequests(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.ListPreauthenticatedRequestsRequest{}
	r, err := c.ListPreauthenticatedRequests(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_UpdateBucket(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.UpdateBucketRequest{}
	r, err := c.UpdateBucket(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_UploadPart(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(configurationProvider())
	failIfError(t, e)
	request := objectstorage.UploadPartRequest{}
	_, err := c.UploadPart(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorage_UploadManager_UploadFile(t *testing.T) {
	ctx := context.Background()
	bname := "uploadManagerFileBucketName"
	objectName := "sampleFileUploadObj"
	namespace := getNamespace(t)

	// clean up
	deleteObjectIgnoreError(t, namespace, bname, objectName)
	deleteBucketIgnoreError(t, namespace, bname)

	createBucket(t, namespace, getCompartmentID(), bname)
	defer deleteBucket(t, namespace, bname)

	contentlen := 1024 * 1000 * 3 // 3MB
	filepath, _, _ := writeTempFileOfSize(int64(contentlen))
	filename := path.Base(filepath)
	defer os.Remove(filename)

	uploadManager := transfer.NewUploadManager()

	req := transfer.UploadFileRequest{
		UploadRequest: transfer.UploadRequest{
			NamespaceName: common.String(namespace),
			BucketName:    common.String(bname),
			ObjectName:    common.String(objectName),
			PartSize:      common.Int64(1024 * 1000 * 1), // 1MB
		},
		FilePath: filepath,
	}

	resp, err := uploadManager.UploadFile(ctx, req)

	if err != nil && resp.IsResumable() {
		resp, err = uploadManager.ResumeUploadFile(ctx, *resp.MultipartUploadResponse.UploadID)
		failIfError(t, err)
	}

	defer deleteObject(t, namespace, bname, objectName)
}

func TestObjectStorage_UploadManager_ResumeUploadFile(t *testing.T) {
	ctx := context.Background()
	bname := "uploadManagerResumeFileBucketName"
	objectName := "sampleFileUploadObj"
	namespace := getNamespace(t)

	// clean up
	deleteObjectIgnoreError(t, namespace, bname, objectName)
	deleteBucketIgnoreError(t, namespace, bname)

	createBucket(t, namespace, getCompartmentID(), bname)
	defer deleteBucket(t, namespace, bname)

	contentlen := 1024 * 1000 * 50 // 50MB
	filepath, _, _ := writeTempFileOfSize(int64(contentlen))
	filename := path.Base(filepath)
	defer os.Remove(filename)

	uploadManager := transfer.NewUploadManager()

	req := transfer.UploadFileRequest{
		UploadRequest: transfer.UploadRequest{
			NamespaceName: common.String(namespace),
			BucketName:    common.String(bname),
			ObjectName:    common.String(objectName),
			PartSize:      common.Int64(1024 * 1000 * 10), // 10MB
		},
		FilePath: filepath,
	}

	// cancel the context after 500ms
	cancelCtx, cancelFn := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancelFn()
	resp, err := uploadManager.UploadFile(cancelCtx, req)

	select {
	case <-cancelCtx.Done():
		if err != nil && resp.IsResumable() {
			fmt.Println("resuming file upload")
			_, err = uploadManager.ResumeUploadFile(context.Background(), *resp.MultipartUploadResponse.UploadID)
			failIfError(t, err)
		}
	default:
		fmt.Println("upload file completed without resume")
	}

	defer deleteObject(t, namespace, bname, objectName)
}

func TestObjectStorage_UploadManager_Stream(t *testing.T) {
	bname := "uploadManagerStreamBucketName"
	objectName := "sampleStreamUploadObj"
	namespace := getNamespace(t)

	// clean up
	deleteObjectIgnoreError(t, namespace, bname, objectName)
	deleteBucketIgnoreError(t, namespace, bname)

	createBucket(t, namespace, getCompartmentID(), bname)
	defer deleteBucket(t, namespace, bname)

	contentlen := 1024 * 1000 * 20 // 20MB
	filepath, _, _ := writeTempFileOfSize(int64(contentlen))
	filename := path.Base(filepath)
	defer func() {
		os.Remove(filename)
	}()

	uploadManager := transfer.NewUploadManager()

	file, _ := os.Open(filepath)
	defer file.Close()

	req := transfer.UploadStreamRequest{
		UploadRequest: transfer.UploadRequest{
			NamespaceName: common.String(namespace),
			BucketName:    common.String(bname),
			ObjectName:    common.String(objectName),
		},
		StreamReader: file, // any struct implements the io.Reader interface
	}

	_, err := uploadManager.UploadStream(context.Background(), req)
	failIfError(t, err)
	defer deleteObject(t, namespace, bname, objectName)
}
