package main

import (
	"context"
	"io/fs"
	"path/filepath"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/sgtool"
	"go.einride.tech/sage/tools/sgbuf"
	"go.einride.tech/sage/tools/sgclangformat"
)

type Proto sg.Namespace

func (Proto) All(ctx context.Context) error {
	sg.Deps(ctx, Proto.ClangFormatProto, Proto.BufLint, Proto.BufGenerate)
	sg.Deps(ctx, Proto.BufGenerateExample)
	return nil
}

func (Proto) BufLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting proto files...")
	cmd := sgbuf.Command(ctx, "lint")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) ClangFormatProto(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting proto files...")
	var protoFiles []string
	if err := filepath.WalkDir(sg.FromGitRoot("proto"), func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && filepath.Ext(path) == ".proto" {
			protoFiles = append(protoFiles, path)
		}
		return nil
	}); err != nil {
		return err
	}
	return sgclangformat.FormatProtoCommand(ctx, protoFiles...).Run()
}

func (Proto) ProtocGenGo(ctx context.Context) error {
	sg.Logger(ctx).Println("installing...")
	_, err := sgtool.GoInstallWithModfile(ctx, "google.golang.org/protobuf/cmd/protoc-gen-go", sg.FromGitRoot("go.mod"))
	return err
}

func (Proto) ProtocGenGoGRPC(ctx context.Context) error {
	sg.Logger(ctx).Println("installing...")
	_, err := sgtool.GoInstall(ctx, "google.golang.org/grpc/cmd/protoc-gen-go-grpc", "v1.2.0")
	return err
}

func (Proto) ProtocGenGoCLI(ctx context.Context) error {
	sg.Logger(ctx).Println("building binary...")
	return sg.Command(ctx, "go", "build", "-o", sg.FromBinDir("protoc-gen-go-cli"), sg.FromGitRoot(".")).Run()
}

func (Proto) BufGenerate(ctx context.Context) error {
	sg.Deps(ctx, Proto.ProtocGenGo)
	sg.Logger(ctx).Println("generating proto stubs...")
	cmd := sgbuf.Command(ctx, "generate", "--template", "buf.gen.yaml", "--path", "einride")
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}

func (Proto) BufGenerateExample(ctx context.Context) error {
	sg.Deps(ctx, Proto.ProtocGenGo, Proto.ProtocGenGoGRPC, Proto.ProtocGenGoCLI)
	sg.Logger(ctx).Println("generating example proto stubs...")
	cmd := sgbuf.Command(
		ctx,
		"generate",
		"buf.build/einride/aip",
		"--template",
		"buf.gen.example.yaml",
		"--path",
		"einride/example/freight",
	)
	cmd.Dir = sg.FromGitRoot("proto")
	return cmd.Run()
}
