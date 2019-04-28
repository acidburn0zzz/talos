/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package basic

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// Credentials implements credentials.PerRPCCredentials. It uses a basic token
// lookup to authenticate users.
type Credentials struct {
	Token string
}

// NewCredentials initializes ClientCredentials with the token.
func NewCredentials(token string) (creds *Credentials) {
	creds = &Credentials{
		Token: token,
	}

	return creds
}

// NewConnection initializes a grpc.ClientConn configured for basic
// authentication.
func NewConnection(address string, port int, creds credentials.PerRPCCredentials) (conn *grpc.ClientConn, err error) {
	grpcOpts := []grpc.DialOption{}

	grpcOpts = append(
		grpcOpts,
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				InsecureSkipVerify: true,
			})),
		grpc.WithPerRPCCredentials(creds),
	)
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", address, port), grpcOpts...)
	if err != nil {
		return
	}

	return conn, nil
}

// GetRequestMetadata sets the value for the "token" key.
func (b *Credentials) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"token": b.Token,
	}, nil
}

// RequireTransportSecurity is set to true in order to encrypt the
// communication.
func (b *Credentials) RequireTransportSecurity() bool {
	return true
}

func (b *Credentials) authorize(ctx context.Context) error {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md["token"]) > 0 && md["token"][0] == b.Token {
			return nil
		}

		return fmt.Errorf("%s", codes.Unauthenticated.String())
	}

	return nil
}

// UnaryInterceptor sets the UnaryServerInterceptor for the server and enforces
// basic authentication.
func (b *Credentials) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	if err := b.authorize(ctx); err != nil {
		return nil, err
	}

	h, err := handler(ctx, req)

	log.Printf("request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err,
	)

	return h, err
}
