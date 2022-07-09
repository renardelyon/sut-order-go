package grpc

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"sut-order-go/lib/helper"
	storagepb "sut-order-go/pb/storage"
	"time"
)

type repo struct {
	storageClient storagepb.StorageServiceClient
}

func NewGrpcRepo(storageClient storagepb.StorageServiceClient) *repo {
	return &repo{
		storageClient: storageClient,
	}
}

func (r *repo) AddFile(path string, userId string) error {
	// TODO: change it so it can receive byte
	file, _ := os.Open(path)

	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := r.storageClient.AddFile(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	req := &storagepb.UploadRequest{
		Data: &storagepb.UploadRequest_Info{
			Info: &storagepb.FileInfo{
				UserId:   userId,
				Filename: path,
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Println(err)
		return err
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println(err)
			return err
		}

		req := &storagepb.UploadRequest{
			Data: &storagepb.UploadRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("image uploaded with id: %s", res.GetId())
	return nil
}

func (r *repo) GetFileByUserId(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &storagepb.GetFileByUserIdRequest{
		UserId: userId,
	}

	stream, err := r.storageClient.GetFileByUserId(ctx, req)
	if err != nil {
		log.Println(err)
		return err
	}

	fileData := bytes.Buffer{}

	for {
		select {
		case <-ctx.Done():
			return helper.ContextError(ctx)
		default:
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("no more stream data")
			break
		}

		chunk := req.GetChunkData()

		_, err = fileData.Write(chunk)
		if err != nil {
			return err
		}
	}

	if err := stream.CloseSend(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
