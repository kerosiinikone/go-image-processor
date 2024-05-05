package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	c "github.com/kerosiinikone/go-docker-grpc/config"
	img_grpc "github.com/kerosiinikone/go-docker-grpc/grpc"
	"google.golang.org/grpc"
)

type Image struct{}

// Internal Data -> attach metadata to chunks of data later
type ImageChunk struct {
	data      []byte
	height    int32
	width     int32
	completed bool
}

// Augmented gRPC logic for concurrent processing
type apiService struct {
	inch  chan ImageChunk
	outch chan ImageChunk

	img_grpc.UnsafeImageServiceServer
}

func NewImageChunk(d []byte, h int32, w int32, completed bool) ImageChunk {
	return ImageChunk{
		data:      d,
		height:    h,
		width:     w,
		completed: completed,
	}
}

func (svc *apiService) TransferImageBytes(srv img_grpc.ImageService_TransferImageBytesServer) error {
	var (
		ctx = srv.Context()
		wg  sync.WaitGroup
	)

	// Send to client -> listen to processed data
	wg.Add(1)
	go func() {
		for processed := range svc.outch {
			if processed.completed {
				wg.Done()
				return
			} else {
				fmt.Printf("New processed: %+v\n", processed.data)
				resp := img_grpc.Image{
					ImageData:   processed.data,
					ImageHeight: processed.height,
					ImageWidth:  processed.width,
				}
				if err := srv.Send(&resp); err != nil {
					log.Printf("%+v", err)
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			image_data_req, err := srv.Recv()
			if err == io.EOF {
				svc.inch <- NewImageChunk(nil, 0, 0, true)
				wg.Wait()
				return nil
			}
			if err != nil {
				// Retry
				continue
			}
			// Pipe to worker
			svc.inch <- NewImageChunk(image_data_req.ImageData, image_data_req.ImageHeight, image_data_req.ImageWidth, false)
		}
	}
}

func startServerAndListen(cfg *c.Config, in *chan ImageChunk, out *chan ImageChunk) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d1", cfg.Server.Addr, cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	img_grpc.RegisterImageServiceServer(s, &apiService{
		inch:  *in,
		outch: *out,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {

	var (
		cfg   = c.Load()
		inch  = make(chan ImageChunk) // Unprocessed
		outch = make(chan ImageChunk) // Processed
		i     = Image{}
	)

	// Start Workers
	go i.processImageBuffer(&inch, &outch)
	// API
	startServerAndListen(cfg, &inch, &outch)
}
