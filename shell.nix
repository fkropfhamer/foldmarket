{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.protobuf
    pkgs.protoc-gen-go
    pkgs.protoc-gen-go-grpc
    pkgs.grpcurl
    pkgs.sqlc
    pkgs.go-migrate
    pkgs.elixir
    pkgs.erlang
    pkgs.nodejs
  ];
}
