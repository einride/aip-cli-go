package main

import (
	"context"
	"os"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/sgtool"
	"go.einride.tech/sage/tools/sgbuf"
)

type Proto sg.Namespace

func (Proto) All(ctx context.Context) error {
	sg.Deps(ctx, Proto.BufFormat, Proto.BufLint)
	sg.Deps(ctx, Proto.CleanGeneratedProto)
	sg.Deps(ctx, Proto.BufGenerateExample)
	return nil
}

func (Proto) BufLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting proto files...")
	cmd := sgbuf.Command(ctx, "lint")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) BufFormat(ctx context.Context) error {
	sg.Logger(ctx).Println("formating proto files...")
	cmd := sgbuf.Command(ctx, "format", "--write")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) BufPush(ctx context.Context) error {
	sg.Logger(ctx).Println("pushing proto module...")
	cmd := sgbuf.Command(ctx, "push")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) ProtocGenGo(ctx context.Context) error {
	sg.Logger(ctx).Println("installing...")
	_, err := sgtool.GoInstallWithModfile(ctx, "google.golang.org/protobuf/cmd/protoc-gen-go", sg.FromGitRoot("go.mod"))
	return err
}

func (Proto) ProtocGenGoAIPCLI(ctx context.Context) error {
	sg.Logger(ctx).Println("building binary...")
	return sg.Command(
		ctx,
		"go",
		"build",
		"-o",
		sg.FromBinDir("protoc-gen-go-aip-cli"),
		sg.FromGitRoot("cmd", "protoc-gen-go-aip-cli"),
	).Run()
}

func (Proto) CleanGeneratedProto(ctx context.Context) error {
	sg.Logger(ctx).Println("cleaning generated proto files...")
	return os.RemoveAll(sg.FromGitRoot("cmd", "examplectl", "gen"))
}

func (Proto) BufGenerateExample(ctx context.Context) error {
	sg.Deps(ctx, Proto.ProtocGenGo, Proto.ProtocGenGoAIPCLI)
	sg.Logger(ctx).Println("generating example proto stubs...")
	cmd := sgbuf.Command(
		ctx,
		"generate",

		"--output",
		sg.FromGitRoot(),
		"--template",
		"buf.gen.example.yaml",
		"--path",
		"einride/example/freight",
	)
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}
