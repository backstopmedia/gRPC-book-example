package main

import (
	"archive/zip"
	"bytes"
	"io"
	"log"

	pb "github.com/backstopmedia/gRPC-book-example/chapters/streaming/server/proto"
)

// ArchiverService is an implementation of the Archiver service in archiver.proto
type ArchiverService struct{}

// Zip generates a tar file from the streamed request messages
func (tb *ArchiverService) Zip(stream pb.Archiver_ZipServer) error {
	buf := new(bytes.Buffer)
	zf := zip.NewWriter(buf)

	for {
		// get or wait for the next request object in the stream
		req, err := stream.Recv()

		// we're done, send the zip file
		if err == io.EOF {
			if err = zf.Close(); err != nil {
				log.Printf("Error creating the zip file: %v", err)
				return err
			}

			return stream.SendAndClose(&pb.ZipResponse{
				ZippedContents: buf.Bytes(),
			})
		}

		// an error occured when getting the request object
		if err != nil {
			log.Printf("Error reading the request: %v", err)
			return err
		}

		f, err := zf.Create(req.FileName)
		if err != nil {
			log.Printf("Error creating the zip file entry: %v", err)
			return err
		}

		if _, err := f.Write(req.Contents); err != nil {
			log.Printf("Error writing zip file: %v", err)
			return err
		}
	}
}
