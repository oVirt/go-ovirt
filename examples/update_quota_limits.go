//
// Copyright (c) 2017 Joey <majunjiev@gmail.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package examples

import (
	"fmt"
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt"
)

func updateQuotaLimits() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("qwer1234").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %v\n", err)
		return
	}
	defer conn.Close()

	// To use `Must` methods, you should recover it if panics
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panics occurs, try the non-Must methods to find the reason")
		}
	}()

	// Find the reference to the root of the tree of services
	systemService := conn.SystemService()

	// Find the data center and the service that manages it
	dcsService := systemService.DataCentersService()
	dc := dcsService.List().
		Search("name=mydc").
		MustSend().
		MustDataCenters().
		Slice()[0]
	dcService := dcsService.DataCenterService(dc.MustId())

	// Find the storage domain and the service that manages it
	sdsService := systemService.StorageDomainsService()
	sd := sdsService.List().
		Search("name=mydata").
		MustSend().
		MustStorageDomains().
		Slice()[0]
	sdService := sdsService.StorageDomainService(sd.MustId())
	sdService.ImagesService()

	// Find the quota and the service that manages it. Note that the service that manages the quota doesn't support
	// search, so we need to retrieve all the quotas and filter explicitly. If the quota doesn't exist, create it.
	quotasService := dcService.QuotasService()
	quotaSlice := quotasService.List().MustSend().MustQuotas()
	var quota *ovirtsdk4.Quota
	for _, q := range quotaSlice.Slice() {
		if q.MustId() == "myquota" {
			quota = q
			break
		}
	}
	if quota == nil {
		quota = quotasService.Add().
			Quota(
				ovirtsdk4.NewQuotaBuilder().
					Name("myquota").
					Description("My quota").
					ClusterHardLimitPct(20).
					ClusterSoftLimitPct(80).
					StorageHardLimitPct(20).
					StorageSoftLimitPct(80).
					MustBuild()).
			MustSend().
			MustQuota()
	}
	quotaService := quotasService.QuotaService(quota.MustId())

	// Find the quota limits for the storage domain that we are interested on
	limitsService := quotaService.QuotaStorageLimitsService()
	limitSlice := limitsService.List().MustSend().MustLimits()
	var limit *ovirtsdk4.QuotaStorageLimit
	for _, l := range limitSlice.Slice() {
		if l.MustId() == sd.MustId() {
			limit = l
			break
		}
	}

	// If that limit exists we will delete it
	if limit != nil {
		limitService := limitsService.LimitService(limit.MustId())
		limitService.Remove().MustSend()
	}

	// Create the limit again with the desired values, in this example it will be 100 GiB
	limitsService.Add().
		Limit(
			ovirtsdk4.NewQuotaStorageLimitBuilder().
				Name("mydatalimit").
				Description("My storage domain limit").
				Limit(100).
				StorageDomain(
					ovirtsdk4.NewStorageDomainBuilder().
						Id(sd.MustId()).
						MustBuild()).
				MustBuild()).
		MustSend()
}
