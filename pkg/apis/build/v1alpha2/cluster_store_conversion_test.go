package v1alpha2

import (
	"context"
	"testing"

	"github.com/sclevine/spec"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

func TestClusterStoreConversion(t *testing.T) {
	spec.Run(t, "testClusterStoreConversion", testClusterStoreConversion)
}

func testClusterStoreConversion(t *testing.T, when spec.G, it spec.S) {
	when("converting to and from v1alpha1", func() {

		store := &ClusterStore{
			ObjectMeta: metav1.ObjectMeta{
				Name: "my-store",
			},
			Spec:       ClusterStoreSpec{
				Sources: []corev1alpha1.StoreImage{
					{
						Image: "image1",
					},
					{
						Image: "image2",
					},
				},
			},
			Status:     ClusterStoreStatus{
				Status: corev1alpha1.Status{
					ObservedGeneration: 1,
					Conditions: []corev1alpha1.Condition{{
						Type:               corev1alpha1.ConditionReady,
						Status:             "True",
						Severity:           "tornado-warning",
						LastTransitionTime: corev1alpha1.VolatileTime{},
						Reason:             "executive-order",
						Message:            "it-is-too-late",
					}},
				},
				Buildpacks: []corev1alpha1.StoreBuildpack{
					{
						BuildpackInfo: corev1alpha1.BuildpackInfo{
							Id:      "",
							Version: "",
						},
						Buildpackage: corev1alpha1.BuildpackageInfo{
							Id:       "",
							Version:  "",
							Homepage: "",
						},
						StoreImage: corev1alpha1.StoreImage{
							Image: "",
						},
						DiffId:   "",
						Digest:   "",
						Size:     0,
						API:      "",
						Homepage: "",
						Order:    []corev1alpha1.OrderEntry{},
						Stacks:   []corev1alpha1.BuildpackStack{},
					},
				},
			},
		}

		it("can convert without data loss", func() {
			v1alpha1ClusterStore := &v1alpha1.ClusterStore{}
			err := store.DeepCopy().ConvertTo(context.TODO(), v1alpha1ClusterStore)
			require.NoError(t, err)

			convertedBackStore := &ClusterStore{}
			err = convertedBackStore.ConvertFrom(context.TODO(), v1alpha1ClusterStore)
			require.NoError(t, err)

			require.Equal(t, store, convertedBackStore)
		})
	})
}
