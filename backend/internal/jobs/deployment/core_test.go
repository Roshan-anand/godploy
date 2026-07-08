package deployjob

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestAssignDeployReturnsSubmitError(t *testing.T) {
	d := NewDeploymentService(nil, nil, nil)

	// Empty DeploymentServiceParams should fail validation because
	// DeploymentID, InstanceID, ServiceID, Token, Url, Branch, etc.
	// are all tagged validate:"required".
	//
	// Before the AssignDeploy refactor, submit() returned the
	// validation error and AssignDeploy propagated it so handlers
	// could return HTTP 500.  After the refactor AssignDeploy
	// swallows submit's error and always returns nil, which makes
	// every handler treat a failed queue submit as success.
	err := d.AssignDeploy(context.Background(), &DeploymentServiceParams{}, nil)
	if err == nil {
		t.Error("AssignDeploy with empty params (validation failure) should return an error, got nil")
	}
}

func TestRegisterBuildWorkServiceIsolation(t *testing.T) {
	d := NewDeploymentService(nil, nil, nil)

	serviceA := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	serviceB := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	// Neither service should have an active entry initially.
	if d.HasActiveRebuild(serviceA) {
		t.Error("HasActiveRebuild(serviceA) = true, want false for empty registry")
	}
	if d.HasActiveRebuild(serviceB) {
		t.Error("HasActiveRebuild(serviceB) = true, want false for empty registry")
	}

	// Register work for service A.
	_, _, _, replaced := d.RegisterBuildWork(context.Background(), serviceA, nil)
	if replaced {
		t.Error("RegisterBuildWork(serviceA) replaced = true, want false for first registration")
	}
	if !d.HasActiveRebuild(serviceA) {
		t.Error("HasActiveRebuild(serviceA) = false after registration, want true")
	}
	if d.HasActiveRebuild(serviceB) {
		t.Error("HasActiveRebuild(serviceB) = true, want false when only serviceA is registered")
	}

	// Register work for service B — should not affect service A.
	_, _, _, replaced = d.RegisterBuildWork(context.Background(), serviceB, nil)
	if replaced {
		t.Error("RegisterBuildWork(serviceB) replaced = true, want false for first registration")
	}
	if !d.HasActiveRebuild(serviceA) {
		t.Error("HasActiveRebuild(serviceA) = false after registering serviceB, want true (isolation)")
	}
	if !d.HasActiveRebuild(serviceB) {
		t.Error("HasActiveRebuild(serviceB) = false after registration, want true")
	}
}

func TestRegisterBuildWorkNewestWins(t *testing.T) {
	d := NewDeploymentService(nil, nil, nil)

	serviceID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	// First registration.
	_, ctx1, _, replaced := d.RegisterBuildWork(context.Background(), serviceID, nil)
	if replaced {
		t.Error("first RegisterBuildWork replaced = true, want false")
	}
	if ctx1.Err() != nil {
		t.Error("first context already expired")
	}

	// Second registration — should replace the first, cancel its context.
	_, ctx2, _, replaced2 := d.RegisterBuildWork(context.Background(), serviceID, nil)
	if !replaced2 {
		t.Error("second RegisterBuildWork replaced = false, want true")
	}
	if ctx2.Err() != nil {
		t.Error("second context already expired")
	}
	if ctx1.Err() == nil {
		t.Error("first context should be canceled after replacement")
	}
}

func TestCancelActiveRebuild(t *testing.T) {
	d := NewDeploymentService(nil, nil, nil)

	serviceID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	// Cancel on empty registry is a no-op, returns false.
	if d.CancelActiveRebuild(serviceID) {
		t.Error("CancelActiveRebuild on empty registry returned true, want false")
	}

	_, ctx, _, _ := d.RegisterBuildWork(context.Background(), serviceID, nil)
	if !d.HasActiveRebuild(serviceID) {
		t.Error("HasActiveRebuild = false after registration, want true")
	}

	// Cancel the active entry.
	if !d.CancelActiveRebuild(serviceID) {
		t.Error("CancelActiveRebuild returned false, want true")
	}
	if d.HasActiveRebuild(serviceID) {
		t.Error("HasActiveRebuild = true after cancel, want false")
	}
	if ctx.Err() == nil {
		t.Error("context should be canceled after CancelActiveRebuild")
	}
}

func TestCleanupRebuildIdempotent(t *testing.T) {
	d := NewDeploymentService(nil, nil, nil)

	serviceID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	// Cleanup on empty registry returns false.
	if d.CleanupRebuild(serviceID, 1) {
		t.Error("CleanupRebuild on empty registry returned true, want false")
	}

	jobID, _, _, _ := d.RegisterBuildWork(context.Background(), serviceID, nil)
	if !d.HasActiveRebuild(serviceID) {
		t.Error("HasActiveRebuild = false after registration, want true")
	}

	// Cleanup with wrong jobID returns false.
	if d.CleanupRebuild(serviceID, jobID+1) {
		t.Error("CleanupRebuild with wrong jobID returned true, want false")
	}
	if !d.HasActiveRebuild(serviceID) {
		t.Error("HasActiveRebuild = false after failed cleanup, want true (entry preserved)")
	}

	// Cleanup with correct jobID returns true and removes entry.
	if !d.CleanupRebuild(serviceID, jobID) {
		t.Error("CleanupRebuild with correct jobID returned false, want true")
	}
	if d.HasActiveRebuild(serviceID) {
		t.Error("HasActiveRebuild = true after successful cleanup, want false")
	}

	// Second cleanup is a no-op.
	if d.CleanupRebuild(serviceID, jobID) {
		t.Error("second CleanupRebuild returned true, want false (already removed)")
	}
}
