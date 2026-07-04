package deployjob

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestHasActiveRebuild(t *testing.T) {
	d := NewDeploymentService(nil, nil, nil)

	serviceID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	otherServiceID := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	if d.HasActiveRebuild(serviceID) {
		t.Errorf("HasActiveRebuild(%s) = true, want false for empty registry", serviceID)
	}

	_, _, _, _ = d.RegisterRebuild(context.Background(), serviceID, nil)

	if !d.HasActiveRebuild(serviceID) {
		t.Errorf("HasActiveRebuild(%s) = false after registering rebuild, want true", serviceID)
	}

	if d.HasActiveRebuild(otherServiceID) {
		t.Errorf("HasActiveRebuild(%s) = true, want false when only another service has an active rebuild", otherServiceID)
	}

	d.CancelActiveRebuild(serviceID)

	if d.HasActiveRebuild(serviceID) {
		t.Errorf("HasActiveRebuild(%s) = true after canceling rebuild, want false", serviceID)
	}
}
