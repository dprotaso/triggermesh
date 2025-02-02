/*
Copyright 2022 TriggerMesh Inc.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
)

// Managed event types
const (
	EventTypeXMLToJSONGenericResponse = "io.triggermesh.xmltojsontransformation.error"
)

var xmlToJSONCondSet = apis.NewLivingConditionSet(
	ConditionDeployed,
)

// GetGroupVersionKind implements kmeta.OwnerRefable
func (t *XMLToJSONTransformation) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("XMLToJSONTransformation")
}

// GetConditionSet retrieves the condition set for this resource. Implements the KRShaped interface.
func (t *XMLToJSONTransformation) GetConditionSet() apis.ConditionSet {
	return xmlToJSONCondSet
}

// MarkServiceUnavailable marks XMLToJSONTransformation as not ready with ServiceUnavailable reason.
func (ts *XMLToJSONTransformationStatus) MarkServiceUnavailable(name string) {
	xmlToJSONCondSet.Manage(ts).MarkFalse(
		"ServiceUnavailable",
		"Service %q is not ready.", name)
}

// PropagateKServiceAvailability uses the availability of the provided KService to determine if
// ConditionDeployed should be marked as true or false.
func (ts *XMLToJSONTransformationStatus) PropagateKServiceAvailability(ksvc *servingv1.Service) {
	if ksvc == nil {
		xmlToJSONCondSet.Manage(ts).MarkUnknown(ConditionDeployed, ReasonUnavailable,
			"The status of the adapter Service can not be determined")
		return
	}

	if ts.Address == nil {
		ts.Address = &duckv1.Addressable{}
	}
	ts.Address.URL = ksvc.Status.URL

	if ksvc.IsReady() {
		xmlToJSONCondSet.Manage(ts).MarkTrue(ConditionDeployed)
		return
	}

	msg := "The adapter Service is unavailable"
	readyCond := ksvc.Status.GetCondition(servingv1.ServiceConditionReady)
	if readyCond != nil && readyCond.Message != "" {
		msg += ": " + readyCond.Message
	}

	xmlToJSONCondSet.Manage(ts).MarkFalse(ConditionDeployed, ReasonUnavailable, msg)

}

// GetStatus retrieves the status of the resource. Implements the KRShaped interface.
func (t *XMLToJSONTransformation) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// MarkNoKService sets the condition that the service is not ready
func (ts *XMLToJSONTransformationStatus) MarkNoKService(reason, messageFormat string, messageA ...interface{}) {
	xmlToJSONCondSet.Manage(ts).MarkFalse(ConditionDeployed, reason, messageFormat, messageA...)
}

// IsReady returns true if the resource is ready overall.
func (ts *XMLToJSONTransformationStatus) IsReady() bool {
	return xmlToJSONCondSet.Manage(ts).IsHappy()
}
