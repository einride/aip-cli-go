package main

import (
	"context"

	"github.com/go-logr/logr"
	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/tools/sgconvco"
	"go.einride.tech/sage/tools/sggit"
	"go.einride.tech/sage/tools/sggo"
	"go.einride.tech/sage/tools/sggolangcilint"
	"go.einride.tech/sage/tools/sggoreview"
	"go.einride.tech/sage/tools/sgmarkdownfmt"
	"go.einride.tech/sage/tools/sgyamlfmt"
)

func main() {
	sg.GenerateMakefiles(
		sg.Makefile{
			Path:          sg.FromGitRoot("Makefile"),
			DefaultTarget: All,
		},
		sg.Makefile{
			Path:          sg.FromGitRoot("proto/Makefile"),
			DefaultTarget: Proto.All,
			Namespace:     Proto{},
		},
	)
}

func All(ctx context.Context) error {
	sg.Deps(ctx, ConvcoCheck, FormatMarkdown, FormatYAML, Proto.All)
	sg.Deps(ctx, GolangciLint, GoReview, GoTest)
	sg.SerialDeps(ctx, GoModTidy, GitVerifyNoDiff)
	return nil
}

func ConvcoCheck(ctx context.Context) error {
	logr.FromContextOrDiscard(ctx).Info("checking git commits...")
	return sgconvco.Command(ctx, "check", "origin/master..HEAD").Run()
}

func FormatYAML(ctx context.Context) error {
	logr.FromContextOrDiscard(ctx).Info("formatting YAML files...")
	return sgyamlfmt.FormatYAML(ctx)
}

func FormatMarkdown(ctx context.Context) error {
	logr.FromContextOrDiscard(ctx).Info("formatting Markdown files...")
	return sgmarkdownfmt.Command(ctx, "-w", ".").Run()
}

func GolangciLint(ctx context.Context) error {
	logr.FromContextOrDiscard(ctx).Info("linting Go files...")
	return sggolangcilint.Run(ctx)
}

func GoReview(ctx context.Context) error {
	logr.FromContextOrDiscard(ctx).Info("reviewing Go files...")
	return sggoreview.Command(ctx, "-c", "1", "./...").Run()
}

func GoModTidy(ctx context.Context) error {
	logr.FromContextOrDiscard(ctx).Info("tidying Go module files...")
	return sg.Command(ctx, "go", "mod", "tidy", "-v").Run()
}

func GoTest(ctx context.Context) error {
	logr.FromContextOrDiscard(ctx).Info("running Go tests...")
	return sggo.TestCommand(ctx).Run()
}

func GitVerifyNoDiff(ctx context.Context) error {
	logr.FromContextOrDiscard(ctx).Info("verifying that git has no diff...")
	return sggit.VerifyNoDiff(ctx)
}
