package main

import (
	"errors"
	"strings"
)

var resources [][]string

func resourceList() [][]string {
	if resources == nil {
		resources = [][]string{
			{"bindings", "Binding"},                                                  // WON'T DO no example available
			{"componentstatuses", "cs", "ComponentStatus"},                           // WON'T DO no example available
			{"configmaps", "cm", "ConfigMap"},                                        // Done
			{"endpoints", "ep", "Endpoints"},                                         // Done
			{"events", "ev", "Event"},                                                // WON'T DO duplicate
			{"limitranges", "limits", "LimitRange"},                                  // Done
			{"namespaces", "ns", "Namespace"},                                        // Done
			{"nodes", "no", "Node"},                                                  // Done
			{"persistentvolumeclaims", "pvc", "PersistentVolumeClaim"},               // Done
			{"persistentvolumes", "pv", "PersistentVolume"},                          // Done
			{"pods", "po", "Pod"},                                                    // Done
			{"podtemplates", "PodTemplate"},                                          // WON'T DO no example available
			{"replicationcontrollers", "rc", "ReplicationController"},                // Done
			{"resourcequotas", "quota", "ResourceQuota"},                             // Done
			{"secrets", "Secret"},                                                    // Done
			{"serviceaccounts", "sa", "ServiceAccount"},                              // Done
			{"services", "svc", "Service"},                                           // Done NOTE check if it could be improved
			{"mutatingwebhookconfigurations", "MutatingWebhookConfiguration"},        // Done NOTE check if it could be improved
			{"validatingwebhookconfigurations", "ValidatingWebhookConfiguration"},    // Done NOTE check if it could be improved
			{"customresourcedefinitions", "crd", "crds", "CustomResourceDefinition"}, // Done NOTE check if it could be improved
			{"apiservices", "APIService"},                                            // WON'T DO no example available
			{"controllerrevisions", "ControllerRevision"},                            // WON'T DO no example available
			{"daemonsets", "ds", "DaemonSet"},                                        // Done
			{"deployments", "deploy", "Deployment"},                                  // Done
			{"replicasets", "rs", "ReplicaSet"},                                      // Done
			{"statefulsets", "sts", "StatefulSet"},                                   // Done
			{"tokenreviews", "TokenReview"},                                          // WON'T DO used between services
			{"localsubjectaccessreviews", "LocalSubjectAccessReview"},                // WON'T DO use kubectl auth can-i instead
			{"selfsubjectaccessreviews", "SelfSubjectAccessReview"},                  // WON'T DO use kubectl auth can-i instead
			{"selfsubjectrulesreviews", "SelfSubjectRulesReview"},                    // WON'T DO use kubectl auth can-i instead
			{"subjectaccessreviews", "SubjectAccessReview"},                          // WON'T DO use kubectl auth can-i instead
			{"horizontalpodautoscalers", "hpa", "HorizontalPodAutoscaler"},           // Done
			{"cronjobs", "cj", "CronJob"},                                            // Done
			{"jobs", "Job"},                                                          // Done
			{"certificatesigningrequests", "csr", "CertificateSigningRequest"},       // Done
			{"leases", "Lease"},                                                      // WON'T DO read only
			{"endpointslices", "EndpointSlice"},                                      // Done
			{"events", "ev", "Event"},                                                // WON'T DO read only
			{"ingresses", "ing", "Ingress"},                                          // WON'T DO duplicate
			{"ingressclasses", "IngressClass"},                                       // Done
			{"ingresses", "ing", "Ingress"},                                          // Done
			{"networkpolicies", "netpol", "NetworkPolicy"},                           // Done
			{"runtimeclasses", "RuntimeClass"},                                       // Done
			{"poddisruptionbudgets", "pdb", "PodDisruptionBudget"},                   // Done
			{"podsecuritypolicies", "psp", "PodSecurityPolicy"},                      // Done
			{"clusterrolebindings", "ClusterRoleBinding"},                            // Done
			{"clusterroles", "ClusterRole"},                                          // Done
			{"rolebindings", "RoleBinding"},                                          // Done
			{"roles", "Role"},                                                        // Done
			{"priorityclasses", "pc", "PriorityClass"},                               // Done
			{"csidrivers", "CSIDriver"},                                              // WON'T DO no simple example available
			{"csinodes", "CSINode"},                                                  // WON'T DO no simple example available
			{"storageclasses", "sc", "StorageClass"},                                 // Done NOTE multiple examples at https://kubernetes.io/docs/concepts/storage/storage-classes/
			{"volumeattachments", "VolumeAttachment"},                                // WON'T DO no example available
			{"list", "List"},                                                         // Done manually added
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
