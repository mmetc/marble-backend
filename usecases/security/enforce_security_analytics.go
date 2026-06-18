package security

import (
	"errors"

	"github.com/checkmarble/marble-backend/models"
)

type EnforceSecurityAnalyticsImpl struct {
	EnforceSecurity
}

func (e *EnforceSecurityAnalyticsImpl) ReadAnalytics(analytics models.Analytics) error {
	return errors.Join(
		e.Permission(models.ANALYTICS_READ), e.ReadOrganization(analytics.OrganizationId),
	)
}
