package get_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/get"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/testutils"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Root", func() {

	BeforeEach(func() {
		helpers.UseMemoryClients()
		_, err := helpers.MustKubeClient().CoreV1().Namespaces().Create(&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: defaults.GlooSystem,
			},
		})
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Empty args and flags", func() {
		It("should give clear error message", func() {
			// Ignore the output message since it changes whenever we add flags and it is tested via the cobra lib.
			_, err := testutils.GlooctlOut("get")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(get.EmptyGetError.Error()))
		})
	})
})
