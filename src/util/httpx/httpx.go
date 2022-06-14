package httpx

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func Init(handler http.Handler) func() {
	Host := viper.GetString("bind.host")
	Port := viper.GetInt("bind.port")
	ReadTimeout := viper.GetInt("bind.read_timeout")
	WriteTimeout := viper.GetInt("bind.write_timeout")
	IdleTimeout := viper.GetInt("bind.idle_timeout")
	CertFile := viper.GetString("bind.cert_file")
	KeyFile := viper.GetString("bind.key_file")
	ShutdownTimeout := viper.GetInt("bind.shutdown_timeout")
	addr := fmt.Sprintf("%s:%d", Host, Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  time.Duration(ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(IdleTimeout) * time.Second,
	}

	go func() {
		fmt.Println("http server listening on:", addr)

		var err error
		if CertFile != "" && KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(CertFile, KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("cannot shutdown http server:", err)
		}

		select {
		case <-ctx.Done():
			fmt.Println("http exiting")
		default:
			fmt.Println("http server stopped")
		}
	}
}
