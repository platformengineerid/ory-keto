// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"

	"github.com/ory/herodot"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/ory/keto/internal/x/events"
	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

var _ rts.WriteServiceServer = (*handler)(nil)

func protoTuplesWithAction(deltas []*rts.RelationTupleDelta, action rts.RelationTupleDelta_Action) (filtered []*ketoapi.RelationTuple, err error) {
	for _, d := range deltas {
		if d.Action == action {
			it, err := (&ketoapi.RelationTuple{}).FromDataProvider(&ketoapi.OpenAPITupleData{Wrapped: d.RelationTuple})
			if err != nil {
				return nil, err
			}
			filtered = append(filtered, it)
		}
	}
	return
}

func (h *handler) TransactRelationTuples(ctx context.Context, req *rts.TransactRelationTuplesRequest) (*rts.TransactRelationTuplesResponse, error) {
	events.Add(ctx, h.d, events.RelationtuplesChanged)

	insertTuples, err := protoTuplesWithAction(req.RelationTupleDeltas, rts.RelationTupleDelta_ACTION_INSERT)
	if err != nil {
		return nil, err
	}

	deleteTuples, err := protoTuplesWithAction(req.RelationTupleDeltas, rts.RelationTupleDelta_ACTION_DELETE)
	if err != nil {
		return nil, err
	}

	its, err := h.d.Mapper().FromTuple(ctx, append(insertTuples, deleteTuples...)...)
	if err != nil {
		return nil, err
	}

	err = h.d.RelationTupleManager().TransactRelationTuples(ctx, its[:len(insertTuples)], its[len(insertTuples):])
	if err != nil {
		return nil, err
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204"))
	snaptokens := make([]string, len(insertTuples))
	for i := range insertTuples {
		snaptokens[i] = "not yet implemented"
	}
	return &rts.TransactRelationTuplesResponse{
		Snaptokens: snaptokens,
	}, nil
}

func (h *handler) CreateRelationTuple(ctx context.Context, request *rts.CreateRelationTupleRequest) (*rts.CreateRelationTupleResponse, error) {
	tuple, err := (&ketoapi.RelationTuple{}).FromDataProvider(&ketoapi.OpenAPITupleData{Wrapped: request.RelationTuple})
	if err != nil {
		return nil, err
	}

	mapped, err := h.d.Mapper().FromTuple(ctx, tuple)
	if err != nil {
		return nil, err
	}

	if err := h.d.RelationTupleManager().WriteRelationTuples(ctx, mapped...); err != nil {
		return nil, err
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))
	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-location", ReadRouteBase+"?"+tuple.ToURLQuery().Encode()))

	return &rts.CreateRelationTupleResponse{RelationTuple: tuple.ToProto()}, nil
}

func (h *handler) DeleteRelationTuples(ctx context.Context, req *rts.DeleteRelationTuplesRequest) (*rts.DeleteRelationTuplesResponse, error) {
	events.Add(ctx, h.d, events.RelationtuplesDeleted)

	var q ketoapi.RelationQuery

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if hasBody := md["hasbody"]; len(hasBody) > 0 && hasBody[0] == "true" {
			_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "400"))
			return nil, errors.WithStack(herodot.ErrBadRequest.WithReason("body is not allowed for this request"))
		}
	}

	switch {
	case req.RelationQuery != nil:
		q.FromDataProvider(&queryWrapper{req.RelationQuery})
	case req.Query != nil: // nolint
		q.FromDataProvider(&deprecatedQueryWrapper{(*rts.ListRelationTuplesRequest_Query)(req.Query)}) // nolint
	default:
		q.FromDataProvider(&openAPIQueryWrapper{req})
	}

	if q.Namespace == nil || *q.Namespace == "" {
		_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "400"))
		return nil, errors.WithStack(herodot.ErrBadRequest.WithReason("Namespace must be set"))
	}

	iq, err := h.d.ReadOnlyMapper().FromQuery(ctx, &q)
	if err != nil {
		return nil, err
	}
	if err := h.d.RelationTupleManager().DeleteAllRelationTuples(ctx, iq); err != nil {
		return nil, errors.WithStack(herodot.ErrInternalServerError.WithError(err.Error()))
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "204"))
	return &rts.DeleteRelationTuplesResponse{}, nil
}
