package main

import (
	"context"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/tools/sgconvco"
	"go.einride.tech/sage/tools/sggit"
	"go.einride.tech/sage/tools/sggo"
	"go.einride.tech/sage/tools/sggolangcilintv2"
	"go.einride.tech/sage/tools/sggosemanticrelease"
	"go.einride.tech/sage/tools/sggovulncheck"
	"go.einride.tech/sage/tools/sgmdformat"
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
	sg.Deps(ctx, GoLint, GoTest, GoVulnCheck)
	sg.Deps(ctx, GoModTidy)
	sg.Deps(ctx, GitVerifyNoDiff)
	return nil
}

func ConvcoCheck(ctx context.Context) error {
	sg.Logger(ctx).Println("checking git commits...")
	return sgconvco.Command(ctx, "check", "origin/master..HEAD").Run()
}

func FormatYAML(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting YAML files...")
	return sgyamlfmt.Run(ctx)
}

func FormatMarkdown(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting Markdown files...")
	return sgmdformat.Command(ctx).Run()
}

func GoLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting Go files...")
	return sggolangcilintv2.Run(ctx, sggolangcilintv2.Config{})
}

func GoModTidy(ctx context.Context) error {
	sg.Logger(ctx).Println("tidying Go module files...")
	return sg.Command(ctx, "go", "mod", "tidy", "-v").Run()
}

func GoTest(ctx context.Context) error {
	sg.Logger(ctx).Println("running Go tests...")
	return sggo.TestCommand(ctx).Run()
}

func GoVulnCheck(ctx context.Context) error {
	sg.Logger(ctx).Println("checking Go files for vulnerabilities...")
	return sggovulncheck.RunAll(ctx)
}

func GitVerifyNoDiff(ctx context.Context) error {
	sg.Logger(ctx).Println("verifying that git has no diff...")
	return sggit.VerifyNoDiff(ctx)
}

func SemanticRelease(ctx context.Context, repo string, dry bool) error {
	sg.Logger(ctx).Println("creating semantic release...")
	args := []string{
		"--allow-initial-development-versions",
		"--allow-no-changes",
		"--ci-condition=default",
		"--provider=github",
		"--provider-opt=slug=" + repo,
	}
	if dry {
		args = append(args, "--dry")
	}
	return sggosemanticrelease.Command(ctx, args...).Run()
}
