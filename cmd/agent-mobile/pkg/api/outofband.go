/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api // nolint: dupl

import "github.com/trustbloc/agent-sdk/cmd/agent-mobile/pkg/wrappers/models"

// OutOfBandController defines methods for the out-of-band protocol controller.
type OutOfBandController interface {

	// CreateRequest creates and saves an Out-Of-Band request message.
	CreateRequest(request *models.RequestEnvelope) *models.ResponseEnvelope

	// CreateInvitation creates and saves an out-of-band invitation.
	CreateInvitation(request *models.RequestEnvelope) *models.ResponseEnvelope

	// AcceptRequest from another agent and return the ID of a new connection record.
	AcceptRequest(request *models.RequestEnvelope) *models.ResponseEnvelope

	// AcceptInvitation from another agent and return the ID of the new connection records.
	AcceptInvitation(request *models.RequestEnvelope) *models.ResponseEnvelope

	// Actions returns pending actions that have not yet to be executed or canceled.
	Actions(request *models.RequestEnvelope) *models.ResponseEnvelope

	// ActionContinue allows continuing with the protocol after an action event was triggered.
	ActionContinue(request *models.RequestEnvelope) *models.ResponseEnvelope

	// ActionStop stops the protocol after an action event was triggered.
	ActionStop(request *models.RequestEnvelope) *models.ResponseEnvelope
}
