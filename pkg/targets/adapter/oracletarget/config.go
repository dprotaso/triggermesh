/*
Copyright 2021 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oracletarget

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig

	OracleAPIKey            string `envconfig:"ORACLE_API_PRIVATE_KEY" required:"true"`
	OracleAPIKeyPassphrase  string `envconfig:"ORACLE_API_PRIVATE_KEY_PASSPHRASE" required:"true"`
	OracleAPIKeyFingerprint string `envconfig:"ORACLE_API_PRIVATE_KEY_FINGERPRINT" required:"true"`
	UserOCID                string `envconfig:"USER_OCID" required:"true"`
	TenantOCID              string `envconfig:"TENANT_OCID" required:"true"`
	OracleRegion            string `envconfig:"ORACLE_REGION" required:"true"`

	// This is labeled as required for now
	OracleFunction string `envconfig:"ORACLE_FUNCTION" required:"true"`
}
