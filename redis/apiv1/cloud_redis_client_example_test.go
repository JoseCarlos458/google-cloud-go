// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package redis_test

import (
	"context"

	redis "cloud.google.com/go/redis/apiv1"
	"google.golang.org/api/iterator"
	redispb "google.golang.org/genproto/googleapis/cloud/redis/v1"
)

func ExampleNewCloudRedisClient() {
	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleCloudRedisClient_ListInstances() {
	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &redispb.ListInstancesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/redis/v1#ListInstancesRequest.
	}
	it := c.ListInstances(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudRedisClient_GetInstance() {
	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &redispb.GetInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/redis/v1#GetInstanceRequest.
	}
	resp, err := c.GetInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudRedisClient_CreateInstance() {
	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &redispb.CreateInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/redis/v1#CreateInstanceRequest.
	}
	op, err := c.CreateInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudRedisClient_UpdateInstance() {
	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &redispb.UpdateInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/redis/v1#UpdateInstanceRequest.
	}
	op, err := c.UpdateInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudRedisClient_UpgradeInstance() {
	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &redispb.UpgradeInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/redis/v1#UpgradeInstanceRequest.
	}
	op, err := c.UpgradeInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudRedisClient_ImportInstance() {
	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &redispb.ImportInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/redis/v1#ImportInstanceRequest.
	}
	op, err := c.ImportInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudRedisClient_ExportInstance() {
	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &redispb.ExportInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/redis/v1#ExportInstanceRequest.
	}
	op, err := c.ExportInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudRedisClient_FailoverInstance() {
	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &redispb.FailoverInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/redis/v1#FailoverInstanceRequest.
	}
	op, err := c.FailoverInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudRedisClient_DeleteInstance() {
	ctx := context.Background()
	c, err := redis.NewCloudRedisClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &redispb.DeleteInstanceRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/redis/v1#DeleteInstanceRequest.
	}
	op, err := c.DeleteInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
