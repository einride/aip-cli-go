//go:build mage
// +build mage

package main

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/magefile/mage/mg"
	"go.einride.tech/mage-tools/mglogr"
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
	ctx = logr.NewContext(ctx, mglogr.New("buf-lint"))
	logr.FromContextOrDiscard(ctx).Info("linting proto files...")
	cmd := mgbuf.Command(ctx, "lint")
	cmd.Dir = mgpath.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) ClangFormatProto(ctx context.Context) error {
	ctx = logr.NewContext(ctx, mglogr.New("clang-format-proto"))
	logr.FromContextOrDiscard(ctx).Info("formatting proto files...")
	protoFiles := mgpath.FindFilesWithExtension(mgpath.FromGitRoot("proto"), ".proto")
	cmd := mgclangformat.FormatProtoCommand(ctx, protoFiles...)
	return cmd.Run()
}

func (Proto) ProtocGenGo(ctx context.Context) error {
	ctx = logr.NewContext(ctx, mglogr.New("protoc-gen-go"))
	logr.FromContextOrDiscard(ctx).Info("installing...")
	_, err := mgtool.GoInstallWithModfile(
		ctx,
		"google.golang.org/protobuf/cmd/protoc-gen-go",
		mgpath.FromGitRoot("go.mod"),
	)
	return err
}

func (Proto) ProtocGenGoGRPC(ctx context.Context) error {
	ctx = logr.NewContext(ctx, mglogr.New("protoc-gen-go-grpc"))
	logr.FromContextOrDiscard(ctx).Info("installing...")
	_, err := mgtool.GoInstall(ctx, "google.golang.org/grpc/cmd/protoc-gen-go-grpc", "v1.2.0")
	return err
}

func (Proto) ProtocGenGoCLI(ctx context.Context) error {
	ctx = logr.NewContext(ctx, mglogr.New("protoc-gen-go-cli"))
	logr.FromContextOrDiscard(ctx).Info("building binary...")
	return mgtool.Command(
		ctx,
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
	ctx = logr.NewContext(ctx, mglogr.New("buf-generate"))
	logr.FromContextOrDiscard(ctx).Info("generating proto stubs...")
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
	ctx = logr.NewContext(ctx, mglogr.New("buf-generate-example"))
	logr.FromContextOrDiscard(ctx).Info("generating example proto stubs...")
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
