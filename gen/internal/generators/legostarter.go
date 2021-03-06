package generators

import (
	. "github.com/dave/jennifer/jen"

	"github.com/vseinstrumentiru/lego/v2/gen/internal/helpers"
)

func NewLeGoStarter(path string) error {
	if err := helpers.MkDir(path); err != nil {
		return err
	}

	if err := newLegoStarterConfig(path); err != nil {
		return err
	}

	if err := newLegoStarterApp(path); err != nil {
		return err
	}

	f := NewFile("main")
	f.ImportName("github.com/vseinstrumentiru/lego/v2/server", "server")

	f.Func().Id("main").Params().
		Block(
			Qual("github.com/vseinstrumentiru/lego/v2/server", "Run").Call(
				Id("application").Values(),
				Qual("github.com/vseinstrumentiru/lego/v2/app", "WithConfig").Call(
					Op("&").Id("config").Values(),
				),
			),
		)

	return f.Save(helpers.Path(path, "main.go"))
}

func newLegoStarterConfig(path string) error {
	f := NewFile("main")

	f.ImportAlias("github.com/vseinstrumentiru/lego/v2/config", "cfg")

	f.Type().Id("config").Struct(
		Id("App").Qual("github.com/vseinstrumentiru/lego/v2/config", "Application"),
	)

	return f.Save(helpers.Path(path, "config.go"))
}

func newLegoStarterApp(path string) error {
	f := NewFile("main")

	f.Type().Id("application").Struct(
		Comment("Log will injected automatically"),
		Id("Log").Qual("github.com/vseinstrumentiru/lego/v2/multilog", "Logger"),
	)

	f.Func().
		Params(Id("app").Id("application")).Id("Providers").Params().Index().Interface().
		Block(
			Return(
				Index().Interface().Block(
					Comment("add your constructors here..."),
				),
			),
		)

	f.Func().
		Params(Id("app").Id("application")).Id("ConfigureService").Params().Error().
		Block(
			Comment("here you can build your service..."),
			Return(
				Nil(),
			),
		)

	return f.Save(helpers.Path(path, "app.go"))
}
