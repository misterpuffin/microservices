package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
  logs.UnimplementedLogServiceServer
  Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
  input := req.GetLogEntry()

  // write the log 
  logEntry := data.LogEntry {
    Name: input.Name,
    Data: input.Data,
  }
  
  err := l.Models.LogEntry.Insert(logEntry)
  if err != nil {
    res := &logs.LogResponse{Result: "failed"}
    return res, err
  }

  // return response
  res := &logs.LogResponse{Result: "logged"}
  return res, nil
}

func (l *LogServer) ListLogs(ctx context.Context, req *logs.ListRequest) (*logs.ListResponse, error) {
  resultsPerPage := req.GetResultPerPage()
  pageNumber := req.GetPageNumber()

  results, err := l.Models.LogEntry.Paginate(resultsPerPage, pageNumber)
  if err != nil {
    res := &logs.ListResponse{Logs: nil, Result: "failed"}
    return res, nil
  }

  payload := []*logs.Log{}
  for _, log := range results {
    payload = append(payload, &logs.Log{Name: log.Name, Data: log.Data})
  }

  res := &logs.ListResponse{Logs: payload, Result: "success"}
  return res, nil
}

func (app *Config) gRPCListen(port string) {
  // Listen to TCP connections on the port
  lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
  if err != nil {
    log.Fatalf("Failed to listen for gRPC: %v", err)
  }

  // Create new grpc server
  s := grpc.NewServer()

  // Register our Log Service with the gRPC server
  logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

  log.Printf("gRPC Server started on port %s", port)

  // Serve the gRPC server
  if err := s.Serve(lis); err != nil {
    log.Fatalf("Failed to listen for gRPC: %v", err)
  }
}
