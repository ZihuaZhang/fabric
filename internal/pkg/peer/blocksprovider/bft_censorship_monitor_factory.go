/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blocksprovider

import "github.com/ZihuaZhang/fabric/internal/pkg/peer/orderers"

// BFTCensorshipMonitorFactory creates an instance of a BFTCensorshipMonitor. It is an implementation of the
// CensorshipDetectorFactory interface which abstracts the creation of a BFTCensorshipMonitor.
type BFTCensorshipMonitorFactory struct{}

func (f *BFTCensorshipMonitorFactory) Create(chainID string, updatableVerifier UpdatableBlockVerifier, requester DeliverClientRequester, progressReporter BlockProgressReporter, fetchSources []*orderers.Endpoint, blockSourceIndex int, timeoutConf TimeoutConfig) CensorshipDetector {
	return NewBFTCensorshipMonitor(chainID, updatableVerifier, requester, progressReporter, fetchSources, blockSourceIndex, timeoutConf)
}
