//go:build mage
// +build mage

package main

import (
	"context"
	"github.com/magefile/mage/mg"
	"go.einride.tech/mage-tools/mglog"
	"go.einride.tech/mage-tools/mgpath"
	"go.einride.tech/mage-tools/mgtool"
	"go.einride.tech/mage-tools/tools/mgbuf"
	"go.einride.tech/mage-tools/tools/mgclangformat"
)

type Proto mg.Namespace

func (Proto) All() {
	mg.Deps(
		Proto.ClangFormatProto,
		Proto.BufLint,
		Proto.BufGenerate,
	)
	mg.Deps(
		Proto.BufGenerateExample,
	)
}

func (Proto) BufLint(ctx context.Context) error {
	mglog.Logger("buf-lint").Info("linting protos...")
	cmd := mgbuf.Command(ctx, "lint")
	cmd.Dir = mgpath.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) ClangFormatProto() error {
	mglog.Logger("clang-format-proto").Info("formatting protos...")
	cmd := mgclangformat.FormatProtoCommand(mgpath.FromGitRoot("proto"))
	return cmd.Run()
}

func (Proto) ProtocGenGo(ctx context.Context) error {
	_, err := mgtool.GoInstallWithModfile(
		ctx,
		"google.golang.org/protobuf/cmd/protoc-gen-go",
		mgpath.FromGitRoot("go.mod"),
	)
	return err
}

func (Proto) ProtocGenGoGRPC(ctx context.Context) error {
	_, err := mgtool.GoInstall(ctx, "google.golang.org/grpc/cmd/protoc-gen-go-grpc", "v1.2.0")
	return err
}

func (Proto) ProtocGenGoCLI() error {
	mglog.Logger("protoc-gen-go-aip").Info("building binary...")
	return mgtool.Command(
		"go",
		"build",
		"-o",
		mgpath.FromBins("protoc-gen-go-cli"),
		mgpath.FromGitRoot("."),
	).Run()
}

func (Proto) BufGenerate(ctx context.Context) error {
	mg.Deps(
		Proto.ProtocGenGo,
	)
	mglog.Logger("buf-generate").Info("generating protobuf stubs...")
	cmd := mgbuf.Command(ctx, "generate", "--template", "buf.gen.yaml", "--path", "einride")
	cmd.Dir = mgpath.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) BufGenerateExample(ctx context.Context) error {
	mg.Deps(
		Proto.ProtocGenGo,
		Proto.ProtocGenGoGRPC,
		Proto.ProtocGenGoCLI,
	)
	mglog.Logger("buf-generate-example").Info("generating example protobuf stubs...")
	cmd := mgbuf.Command(
		ctx,
		"generate",
		"buf.build/einride/aip",
		"--template",
		"buf.gen.example.yaml",
		"--path",
		"einride/example/freight",
	)
	cmd.Dir = mgpath.FromGitRoot("proto")
	return cmd.Run()
}
