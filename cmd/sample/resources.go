package main

import (
	"errors"
	"strings"
)

var resources [][]string

func resourceList() [][]string {
	if resources == nil {
		resources = [][]string{
			{"bindings", "Binding"},                                                  // TODO no example available
			{"componentstatuses", "cs", "ComponentStatus"},                           // TODO no example available
			{"configmaps", "cm", "ConfigMap"},                                        // Done
			{"endpoints", "ep", "Endpoints"},                                         // Done
			{"events", "ev", "Event"},                                                // TODO no example available
			{"limitranges", "limits", "LimitRange"},                                  // Done
			{"namespaces", "ns", "Namespace"},                                        // Done
			{"nodes", "no", "Node"},                                                  // Done
			{"persistentvolumeclaims", "pvc", "PersistentVolumeClaim"},               // Done
			{"persistentvolumes", "pv", "PersistentVolume"},                          // Done
			{"pods", "po", "Pod"},                                                    // Done
			{"podtemplates", "PodTemplate"},                                          // TODO no example available
			{"replicationcontrollers", "rc", "ReplicationController"},                // Done
			{"resourcequotas", "quota", "ResourceQuota"},                             // Done NOTE example is a list of resources
			{"secrets", "Secret"},                                                    // Done
			{"serviceaccounts", "sa", "ServiceAccount"},                              // Done
			{"services", "svc", "Service"},                                           // Done
			{"mutatingwebhookconfigurations", "MutatingWebhookConfiguration"},        // TODO check if it could be improved
			{"validatingwebhookconfigurations", "ValidatingWebhookConfiguration"},    // TODO check if it could be improved
			{"customresourcedefinitions", "crd", "crds", "CustomResourceDefinition"}, // TODO check if it could be improved
			{"apiservices", "APIService"},                                            // TODO no example available
			{"controllerrevisions", "ControllerRevision"},                            // TODO no example available
			{"daemonsets", "ds", "DaemonSet"},                                        // Done
			{"deployments", "deploy", "Deployment"},                                  // Done
			{"replicasets", "rs", "ReplicaSet"},                                      // Done
			{"statefulsets", "sts", "StatefulSet"},                                   // Done NOTE double yaml
			{"tokenreviews", "TokenReview"},                                          // WON'T DO used between services
			{"localsubjectaccessreviews", "LocalSubjectAccessReview"},                // WON'T DO use kubectl auth can-i instead
			{"selfsubjectaccessreviews", "SelfSubjectAccessReview"},                  // WON'T DO use kubectl auth can-i instead
			{"selfsubjectrulesreviews", "SelfSubjectRulesReview"},                    // WON'T DO use kubectl auth can-i instead
			{"subjectaccessreviews", "SubjectAccessReview"},                          // WON'T DO use kubectl auth can-i instead
			{"horizontalpodautoscalers", "hpa", "HorizontalPodAutoscaler"},           // Done
			{"cronjobs", "cj", "CronJob"},                                            // Done
			{"jobs", "Job"},                                                          // Done
			{"certificatesigningrequests", "csr", "CertificateSigningRequest"},       // Done
			{"leases", "Lease"},                                                      // TODO no example available
			{"endpointslices", "EndpointSlice"},                                      // Done
			{"events", "ev", "Event"},                                                // TODO no example available
			{"ingresses", "ing", "Ingress"},                                          // TODO extensions api group
			{"ingressclasses", "IngressClass"},                                       // Done
			{"ingresses", "ing", "Ingress"},                                          // Done // networking.k8s.io api group
			{"networkpolicies", "netpol", "NetworkPolicy"},                           // Done
			{"runtimeclasses", "RuntimeClass"},                                       // Done
			{"poddisruptionbudgets", "pdb", "PodDisruptionBudget"},                   // Done
			{"podsecuritypolicies", "psp", "PodSecurityPolicy"},                      // Done
			{"clusterrolebindings", "ClusterRoleBinding"},                            // Done
			{"clusterroles", "ClusterRole"},                                          // Done
			{"rolebindings", "RoleBinding"},                                          // Done
			{"roles", "Role"},                                                        // Done
			{"priorityclasses", "pc", "PriorityClass"},                               // Done
			{"csidrivers", "CSIDriver"},                                              // TODO no simple example available
			{"csinodes", "CSINode"},                                                  // TODO no simple example available
			{"storageclasses", "sc", "StorageClass"},                                 // Done NOTE multiple examples at https://kubernetes.io/docs/concepts/storage/storage-classes/
			{"volumeattachments", "VolumeAttachment"},                                // TODO no example available
		}
	}

	return resources
}

func resourceName(name string) (string, error) {
	lcaseName := strings.ToLower(name)
	list := resourceList()

	for _, resource := range list {
		for _, alias := range resource {
			if strings.ToLower(alias) == lcaseName {
				// The last item is the resource name
				return resource[len(resource)-1], nil
			}
		}
	}

	return "", errors.New("resource alias not found")
}
