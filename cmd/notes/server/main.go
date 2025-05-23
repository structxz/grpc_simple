package main

import (
	"context"
	"log"
	"net"
	"sync"
	"sync/atomic"

	pb "github.com/structxz/grpc_simple/pkg/api/notes"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var idSerial uint64

type server struct {
	pb.UnimplementedNotesServiceServer

	mx sync.RWMutex
	notes map[uint64]*pb.NoteInfo
}

func NewServer() *server {
	return &server{notes: make(map[uint64]*pb.NoteInfo)}
}

func (s *server) SaveNote(_ context.Context, req *pb.SaveNoteRequest) (*pb.SaveNoteResponse, error) {
	info := req.GetInfo()

	log.Printf("SaveNote: received %s", info.GetTitle())

	if err := validateNoteRequest(req); err != nil {
		return nil, err
	}

	id := atomic.AddUint64(&idSerial, 1)
	s.mx.Lock()
	s.notes[id] = info
	s.mx.Unlock()

	return &pb.SaveNoteResponse{
		Id: id,
	}, nil
}

func validateNoteRequest(req *pb.SaveNoteRequest) error {
	info := req.GetInfo()

	var violations []*errdetails.BadRequest_FieldViolation
	if len(info.GetTitle()) == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field: "info.title",
			Description: "empty",
		})
	}
	if len(info.GetContent()) == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field: "info.content",
			Description: "empty",
		})	
	}

	if len(violations) > 0 {
		st, err := status.New(codes.InvalidArgument, codes.InvalidArgument.String()).
			WithDetails(&errdetails.BadRequest{
				FieldViolations: violations,
			})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		return st.Err()
	}

	return nil
}

// ListNotes implements pb.NotesServiceServer
func (s *server) ListNotes(_ context.Context, req *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	log.Println("ListNotes: received")

	s.mx.RLock()
	defer s.mx.RUnlock()

	notes := make([]*pb.Note, 0, len(s.notes))
	for id, note := range s.notes {
		notes = append(notes, &pb.Note{
			Id: id,
			Info: note,
		})
	}

	return &pb.ListNotesResponse{
		Notes: notes,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotesServiceServer(s, NewServer())

	reflection.Register(s) // this needs for json requests

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// SaveNote:
	// grpc_cli call --json_input --json_output localhost:8082 NotesService/SaveNote '{"info:{"title": "note"}"}'
	// ListNotes:
	// grpc_cli call --json_input --json_output localhost:8082 NotesService/ListNotes '{}'

	
	// s.GracefulStop()
	// some timeout
	// s.Stop() -- forced stop, if graceful stop did not stop
}

