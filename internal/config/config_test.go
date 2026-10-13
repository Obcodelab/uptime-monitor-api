package config

import "testing"

func TestLoad(t *testing.T) {
	t.Run("defaults", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://x")
		t.Setenv("JWT_SECRET", "0123456789abcdef0123456789abcdef")
		cfg, err := Load()
		if err != nil {
			t.Fatal(err)
		}
		if cfg.HTTPPort != "8000" || cfg.CheckerWorkers != 10 || cfg.AllowPrivateTargets {
			t.Fatalf("unexpected defaults: %+v", cfg)
		}
	})

	t.Run("short secret fails", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://x")
		t.Setenv("JWT_SECRET", "short")
		if _, err := Load(); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("bad int fails", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "postgres://x")
		t.Setenv("JWT_SECRET", "0123456789abcdef0123456789abcdef")
		t.Setenv("CHECKER_WORKERS", "abc")
		if _, err := Load(); err == nil {
			t.Fatal("expected error")
		}
	})
}
