package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strings"

	pb "github.com/backstopmedia/gRPC-book-example/chapters/streaming/server/proto"
)

type TokenizerService struct{}

func (ts *TokenizerService) Tokenize(s pb.Tokenizer_TokenizeServer) error {
	for {
		req, err := s.Recv()
		if err == io.EOF {
			return nil
		}

		// an error occured when getting the request object
		if err != nil {
			log.Printf("Got an error receiving from the client: %v", err)
			return err
		}

		rdr := bufio.NewReader(bytes.NewReader(req.FileContents))
		scanner := bufio.NewScanner(rdr)
		scanner.Split(bufio.ScanWords)

		results := &pb.TokenizeResponse{
			Words: make(map[string]int64),
		}

		for scanner.Scan() {
			word := strings.TrimSpace(scanner.Text())

			if _, ok := results.Words[word]; ok {
				results.Words[word]++
			} else {
				results.Words[word] = 1
			}
		}

		if err = s.Send(results); err != nil {
			log.Printf("Got an error sending to the client: %v", err)
			return err
		}
	}
}
