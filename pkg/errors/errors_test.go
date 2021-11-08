// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"errors"
	atomixerrors "github.com/atomix/atomix-go-framework/pkg/atomix/errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestFactories(t *testing.T) {
	assert.Equal(t, Unknown, NewUnknown("").(*TypedError).Type)
	assert.Equal(t, "Unknown", NewUnknown("Unknown").Error())
	assert.Equal(t, Canceled, NewCanceled("").(*TypedError).Type)
	assert.Equal(t, "Canceled", NewCanceled("Canceled").Error())
	assert.Equal(t, NotFound, NewNotFound("").(*TypedError).Type)
	assert.Equal(t, "NotFound", NewNotFound("NotFound").Error())
	assert.Equal(t, AlreadyExists, NewAlreadyExists("").(*TypedError).Type)
	assert.Equal(t, "AlreadyExists", NewAlreadyExists("AlreadyExists").Error())
	assert.Equal(t, Unauthorized, NewUnauthorized("").(*TypedError).Type)
	assert.Equal(t, "Unauthorized", NewUnauthorized("Unauthorized").Error())
	assert.Equal(t, Forbidden, NewForbidden("").(*TypedError).Type)
	assert.Equal(t, "Forbidden", NewForbidden("Forbidden").Error())
	assert.Equal(t, Conflict, NewConflict("").(*TypedError).Type)
	assert.Equal(t, "Conflict", NewConflict("Conflict").Error())
	assert.Equal(t, Invalid, NewInvalid("").(*TypedError).Type)
	assert.Equal(t, "Invalid", NewInvalid("Invalid").Error())
	assert.Equal(t, Unavailable, NewUnavailable("").(*TypedError).Type)
	assert.Equal(t, "Unavailable", NewUnavailable("Unavailable").Error())
	assert.Equal(t, NotSupported, NewNotSupported("").(*TypedError).Type)
	assert.Equal(t, "NotSupported", NewNotSupported("NotSupported").Error())
	assert.Equal(t, Timeout, NewTimeout("").(*TypedError).Type)
	assert.Equal(t, "Timeout", NewTimeout("Timeout").Error())
	assert.Equal(t, Internal, NewInternal("").(*TypedError).Type)
	assert.Equal(t, "Internal", NewInternal("Internal").Error())
}

func TestPredicates(t *testing.T) {
	assert.False(t, IsUnknown(errors.New("Unknown")))
	assert.True(t, IsUnknown(NewUnknown("Unknown")))
	assert.False(t, IsCanceled(errors.New("Canceled")))
	assert.True(t, IsCanceled(NewCanceled("Canceled")))
	assert.False(t, IsNotFound(errors.New("NotFound")))
	assert.True(t, IsNotFound(NewNotFound("NotFound")))
	assert.False(t, IsAlreadyExists(errors.New("AlreadyExists")))
	assert.True(t, IsAlreadyExists(NewAlreadyExists("AlreadyExists")))
	assert.False(t, IsUnauthorized(errors.New("Unauthorized")))
	assert.True(t, IsUnauthorized(NewUnauthorized("Unauthorized")))
	assert.False(t, IsForbidden(errors.New("Forbidden")))
	assert.True(t, IsForbidden(NewForbidden("Forbidden")))
	assert.False(t, IsConflict(errors.New("Conflict")))
	assert.True(t, IsConflict(NewConflict("Conflict")))
	assert.False(t, IsInvalid(errors.New("Invalid")))
	assert.True(t, IsInvalid(NewInvalid("Invalid")))
	assert.False(t, IsUnavailable(errors.New("Unavailable")))
	assert.True(t, IsUnavailable(NewUnavailable("Unavailable")))
	assert.False(t, IsNotSupported(errors.New("NotSupported")))
	assert.True(t, IsNotSupported(NewNotSupported("NotSupported")))
	assert.False(t, IsTimeout(errors.New("Timeout")))
	assert.True(t, IsTimeout(NewTimeout("Timeout")))
	assert.False(t, IsInternal(errors.New("Internal")))
	assert.True(t, IsInternal(NewInternal("Internal")))
}

func TestErrorToStatus(t *testing.T) {
	assert.Equal(t, codes.OK, Status(nil).Code())
	assert.Equal(t, "", Status(nil).Message())
	assert.Equal(t, codes.Unknown, Status(NewUnknown("")).Code())
	assert.Equal(t, "Unknown", Status(NewUnknown("Unknown")).Message())
	assert.Equal(t, codes.Canceled, Status(NewCanceled("")).Code())
	assert.Equal(t, "Canceled", Status(NewCanceled("Canceled")).Message())
	assert.Equal(t, codes.NotFound, Status(NewNotFound("")).Code())
	assert.Equal(t, "NotFound", Status(NewNotFound("NotFound")).Message())
	assert.Equal(t, codes.AlreadyExists, Status(NewAlreadyExists("")).Code())
	assert.Equal(t, "AlreadyExists", Status(NewAlreadyExists("AlreadyExists")).Message())
	assert.Equal(t, codes.Unauthenticated, Status(NewUnauthorized("")).Code())
	assert.Equal(t, "Unauthorized", Status(NewUnauthorized("Unauthorized")).Message())
	assert.Equal(t, codes.PermissionDenied, Status(NewForbidden("")).Code())
	assert.Equal(t, "Forbidden", Status(NewForbidden("Forbidden")).Message())
	assert.Equal(t, codes.FailedPrecondition, Status(NewConflict("")).Code())
	assert.Equal(t, "Conflict", Status(NewConflict("Conflict")).Message())
	assert.Equal(t, codes.InvalidArgument, Status(NewInvalid("")).Code())
	assert.Equal(t, "Invalid", Status(NewInvalid("Invalid")).Message())
	assert.Equal(t, codes.Unavailable, Status(NewUnavailable("")).Code())
	assert.Equal(t, "Unavailable", Status(NewUnavailable("Unavailable")).Message())
	assert.Equal(t, codes.Unimplemented, Status(NewNotSupported("")).Code())
	assert.Equal(t, "NotSupported", Status(NewNotSupported("NotSupported")).Message())
	assert.Equal(t, codes.DeadlineExceeded, Status(NewTimeout("")).Code())
	assert.Equal(t, "Timeout", Status(NewTimeout("Timeout")).Message())
	assert.Equal(t, codes.Internal, Status(NewInternal("")).Code())
	assert.Equal(t, "Internal", Status(NewInternal("Internal")).Message())
}

func TestStatusToError(t *testing.T) {
	assert.Nil(t, FromStatus(status.New(codes.OK, "")))
	assert.True(t, IsUnknown(FromStatus(status.New(codes.Unknown, ""))))
	assert.Equal(t, "Unknown", FromStatus(status.New(codes.Unknown, "Unknown")).Error())
	assert.True(t, IsCanceled(FromStatus(status.New(codes.Canceled, ""))))
	assert.Equal(t, "Canceled", FromStatus(status.New(codes.Canceled, "Canceled")).Error())
	assert.True(t, IsNotFound(FromStatus(status.New(codes.NotFound, ""))))
	assert.Equal(t, "NotFound", FromStatus(status.New(codes.NotFound, "NotFound")).Error())
	assert.True(t, IsAlreadyExists(FromStatus(status.New(codes.AlreadyExists, ""))))
	assert.Equal(t, "AlreadyExists", FromStatus(status.New(codes.AlreadyExists, "AlreadyExists")).Error())
	assert.True(t, IsUnauthorized(FromStatus(status.New(codes.Unauthenticated, ""))))
	assert.Equal(t, "Unauthenticated", FromStatus(status.New(codes.Unauthenticated, "Unauthenticated")).Error())
	assert.True(t, IsForbidden(FromStatus(status.New(codes.PermissionDenied, ""))))
	assert.Equal(t, "PermissionDenied", FromStatus(status.New(codes.PermissionDenied, "PermissionDenied")).Error())
	assert.True(t, IsConflict(FromStatus(status.New(codes.FailedPrecondition, ""))))
	assert.Equal(t, "FailedPrecondition", FromStatus(status.New(codes.FailedPrecondition, "FailedPrecondition")).Error())
	assert.True(t, IsInvalid(FromStatus(status.New(codes.InvalidArgument, ""))))
	assert.Equal(t, "InvalidArgument", FromStatus(status.New(codes.InvalidArgument, "InvalidArgument")).Error())
	assert.True(t, IsUnavailable(FromStatus(status.New(codes.Unavailable, ""))))
	assert.Equal(t, "Unavailable", FromStatus(status.New(codes.Unavailable, "Unavailable")).Error())
	assert.True(t, IsNotSupported(FromStatus(status.New(codes.Unimplemented, ""))))
	assert.Equal(t, "Unimplemented", FromStatus(status.New(codes.Unimplemented, "Unimplemented")).Error())
	assert.True(t, IsTimeout(FromStatus(status.New(codes.DeadlineExceeded, ""))))
	assert.Equal(t, "DeadlineExceeded", FromStatus(status.New(codes.DeadlineExceeded, "DeadlineExceeded")).Error())
	assert.True(t, IsInternal(FromStatus(status.New(codes.Internal, ""))))
	assert.Equal(t, "Internal", FromStatus(status.New(codes.Internal, "Internal")).Error())
}

func TestGRPCToError(t *testing.T) {
	assert.Nil(t, FromGRPC(status.New(codes.OK, "").Err()))
	assert.True(t, IsUnknown(FromGRPC(status.New(codes.Unknown, "").Err())))
	assert.Equal(t, "Unknown", FromGRPC(status.New(codes.Unknown, "Unknown").Err()).Error())
	assert.True(t, IsCanceled(FromGRPC(status.New(codes.Canceled, "").Err())))
	assert.Equal(t, "Canceled", FromGRPC(status.New(codes.Canceled, "Canceled").Err()).Error())
	assert.True(t, IsNotFound(FromGRPC(status.New(codes.NotFound, "").Err())))
	assert.Equal(t, "NotFound", FromGRPC(status.New(codes.NotFound, "NotFound").Err()).Error())
	assert.True(t, IsAlreadyExists(FromGRPC(status.New(codes.AlreadyExists, "").Err())))
	assert.Equal(t, "AlreadyExists", FromGRPC(status.New(codes.AlreadyExists, "AlreadyExists").Err()).Error())
	assert.True(t, IsUnauthorized(FromGRPC(status.New(codes.Unauthenticated, "").Err())))
	assert.Equal(t, "Unauthenticated", FromGRPC(status.New(codes.Unauthenticated, "Unauthenticated").Err()).Error())
	assert.True(t, IsForbidden(FromGRPC(status.New(codes.PermissionDenied, "").Err())))
	assert.Equal(t, "PermissionDenied", FromGRPC(status.New(codes.PermissionDenied, "PermissionDenied").Err()).Error())
	assert.True(t, IsConflict(FromGRPC(status.New(codes.FailedPrecondition, "").Err())))
	assert.Equal(t, "FailedPrecondition", FromGRPC(status.New(codes.FailedPrecondition, "FailedPrecondition").Err()).Error())
	assert.True(t, IsInvalid(FromGRPC(status.New(codes.InvalidArgument, "").Err())))
	assert.Equal(t, "InvalidArgument", FromGRPC(status.New(codes.InvalidArgument, "InvalidArgument").Err()).Error())
	assert.True(t, IsUnavailable(FromGRPC(status.New(codes.Unavailable, "").Err())))
	assert.Equal(t, "Unavailable", FromGRPC(status.New(codes.Unavailable, "Unavailable").Err()).Error())
	assert.True(t, IsNotSupported(FromGRPC(status.New(codes.Unimplemented, "").Err())))
	assert.Equal(t, "Unimplemented", FromGRPC(status.New(codes.Unimplemented, "Unimplemented").Err()).Error())
	assert.True(t, IsTimeout(FromGRPC(status.New(codes.DeadlineExceeded, "").Err())))
	assert.Equal(t, "DeadlineExceeded", FromGRPC(status.New(codes.DeadlineExceeded, "DeadlineExceeded").Err()).Error())
	assert.True(t, IsInternal(FromGRPC(status.New(codes.Internal, "").Err())))
	assert.Equal(t, "Internal", FromGRPC(status.New(codes.Internal, "Internal").Err()).Error())
}

func TestAtomixToError(t *testing.T) {
	assert.Nil(t, FromAtomix(nil))
	assert.True(t, IsUnknown(FromAtomix(atomixerrors.NewUnknown("Unknown"))))
	assert.Equal(t, "Unknown", FromAtomix(atomixerrors.NewUnknown("Unknown")).Error())
	assert.True(t, IsCanceled(FromAtomix(atomixerrors.NewCanceled("Canceled"))))
	assert.Equal(t, "Canceled", FromAtomix(atomixerrors.NewCanceled("Canceled")).Error())
	assert.True(t, IsNotFound(FromAtomix(atomixerrors.NewNotFound("NotFound"))))
	assert.Equal(t, "NotFound", FromAtomix(atomixerrors.NewNotFound("NotFound")).Error())
	assert.True(t, IsAlreadyExists(FromAtomix(atomixerrors.NewAlreadyExists("AlreadyExists"))))
	assert.Equal(t, "AlreadyExists", FromAtomix(atomixerrors.NewAlreadyExists("AlreadyExists")).Error())
	assert.True(t, IsUnauthorized(FromAtomix(atomixerrors.NewUnauthorized("Unauthorized"))))
	assert.Equal(t, "Unauthorized", FromAtomix(atomixerrors.NewUnauthorized("Unauthorized")).Error())
	assert.True(t, IsForbidden(FromAtomix(atomixerrors.NewForbidden("Forbidden"))))
	assert.Equal(t, "Forbidden", FromAtomix(atomixerrors.NewForbidden("Forbidden")).Error())
	assert.True(t, IsConflict(FromAtomix(atomixerrors.NewConflict("Conflict"))))
	assert.Equal(t, "Conflict", FromAtomix(atomixerrors.NewConflict("Conflict")).Error())
	assert.True(t, IsInvalid(FromAtomix(atomixerrors.NewInvalid("Invalid"))))
	assert.Equal(t, "Invalid", FromAtomix(atomixerrors.NewInvalid("Invalid")).Error())
	assert.True(t, IsUnavailable(FromAtomix(atomixerrors.NewUnavailable("Unavailable"))))
	assert.Equal(t, "Unavailable", FromAtomix(atomixerrors.NewUnavailable("Unavailable")).Error())
	assert.True(t, IsNotSupported(FromAtomix(atomixerrors.NewNotSupported("NotSupported"))))
	assert.Equal(t, "NotSupported", FromAtomix(atomixerrors.NewNotSupported("NotSupported")).Error())
	assert.True(t, IsTimeout(FromAtomix(atomixerrors.NewTimeout("Timeout"))))
	assert.Equal(t, "Timeout", FromAtomix(atomixerrors.NewTimeout("Timeout")).Error())
	assert.True(t, IsInternal(FromAtomix(atomixerrors.NewInternal("Internal"))))
	assert.Equal(t, "Internal", FromAtomix(atomixerrors.NewInternal("Internal")).Error())
}
