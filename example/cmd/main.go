package main

import (
	"context"

	"github.com/danteay/golog"
	"github.com/danteay/golog/adapters/zerolog"
	"github.com/danteay/golog/internal/contextfields"
	"github.com/danteay/golog/levels"
)

func main() {
	println("---- Using default context fields space ----")
	println("==> Default Context Colored")
	useDefaultContextColored()
	println("==> DefaultContext NoColored")
	useDefaultContextNoColored()
	println("==> After Flush global default fields")
	contextfields.Flush(context.Background())
	useDefaultContextNoColored()

	println()
	println("---- Using Custom exec context ----")
	println("==> Set execution 1")
	useCustomExecContext()
	println("==> Set execution 2")
	useCustomExecContext2()

	println()
	println("---- Using multiple exec contexts ----")
	usingMultipleContext()
}

func useDefaultContextColored() {
	logger := golog.New(
		golog.WithAdapter(
			zerolog.New(
				zerolog.WithLevel(levels.Debug),
				zerolog.Colored(),
			),
		),
	)

	logger.SetContextFields(map[string]any{
		"stage": "test",
	})

	fields := map[string]any{
		"key1": "value1",
	}

	logger.Fields(fields).Debug("Hello %s", "world")
	logger.Fields(fields).Info("Hello %s", "world")
	logger.Fields(fields).Warn("Hello %s", "world")
	logger.Fields(fields).Error("Hello %s", "world")
}

func useDefaultContextNoColored() {
	logger := golog.New(
		golog.WithAdapter(
			zerolog.New(
				zerolog.WithLevel(levels.Debug),
			),
		),
	)

	fields := map[string]any{
		"key1": "value1",
	}

	logger.Fields(fields).Debug("Hello %s", "world")
	logger.Fields(fields).Info("Hello %s", "world")
	logger.Fields(fields).Warn("Hello %s", "world")
	logger.Fields(fields).Error("Hello %s", "world")
}

func useCustomExecContext() {
	ctx := context.WithValue(context.Background(), contextfields.ExecutionContextKey, "execution-1")

	logger := golog.New(
		golog.WithAdapter(
			zerolog.New(
				zerolog.WithLevel(levels.Debug),
				zerolog.Colored(),
			),
		),
	).SetContext(ctx)

	logger.SetContextFields(map[string]any{
		"execution": 1,
	})

	fields := map[string]any{
		"key1": "value1",
	}

	logger.Fields(fields).Debug("Hello %s", "world")
	logger.Fields(fields).Info("Hello %s", "world")
	logger.Fields(fields).Warn("Hello %s", "world")
	logger.Fields(fields).Error("Hello %s", "world")
}

func useCustomExecContext2() {
	ctx := context.WithValue(context.Background(), contextfields.ExecutionContextKey, "execution-2")

	logger := golog.New(
		golog.WithAdapter(
			zerolog.New(
				zerolog.WithLevel(levels.Debug),
				zerolog.Colored(),
			),
		),
	).SetContext(ctx)

	logger.SetContextFields(map[string]any{
		"execution": 2,
	})

	fields := map[string]any{
		"key1": "value1",
	}

	logger.Fields(fields).Debug("Hello %s", "world")
	logger.Fields(fields).Info("Hello %s", "world")
	logger.Fields(fields).Warn("Hello %s", "world")
	logger.Fields(fields).Error("Hello %s", "world")

	println("#> After flushing custom exec context 2 ")

	logger.FlushContextFields()

	logger.Fields(fields).Debug("Hello %s", "world")
	logger.Fields(fields).Info("Hello %s", "world")
	logger.Fields(fields).Warn("Hello %s", "world")
	logger.Fields(fields).Error("Hello %s", "world")
}

func usingMultipleContext() {
	defLog := golog.New(
		golog.WithAdapter(
			zerolog.New(
				zerolog.WithLevel(levels.Debug),
				zerolog.Colored(),
			),
		),
	).SetContext(context.Background())

	defLog.SetContextFields(map[string]any{
		"execution": "default",
	})

	ctx := context.WithValue(context.Background(), contextfields.ExecutionContextKey, "execution")

	customLog := golog.New(
		golog.WithAdapter(
			zerolog.New(
				zerolog.WithLevel(levels.Debug),
				zerolog.Colored(),
			),
		),
	).SetContext(ctx)

	customLog.SetContextFields(map[string]any{
		"execution": "custom",
	})

	defLog.Info("Hello %s", "world")
	customLog.Info("Hello %s", "world")

	println("#> Flush custom")
	customLog.FlushContextFields()
	defLog.Info("Hello %s", "world")
	customLog.Info("Hello %s", "world")

	println("#> Set to custom and flush default")
	defLog.FlushContextFields()
	customLog.SetContextFields(map[string]any{
		"execution": "custom",
	})

	defLog.Info("Hello %s", "world")
	customLog.Info("Hello %s", "world")

	println("#> Set again to default")
	defLog.SetContextFields(map[string]any{
		"execution": "custom",
	})

	defLog.Info("Hello %s", "world")
	customLog.Info("Hello %s", "world")

	println("#> Flush all")
	contextfields.FlushAll()
	defLog.Info("Hello %s", "world")
	customLog.Info("Hello %s", "world")
}
