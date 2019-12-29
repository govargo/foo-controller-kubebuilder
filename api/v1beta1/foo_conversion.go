/*
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

package v1beta1

/*
For imports, we'll need the controller-runtime
[`conversion`](https://godoc.org/sigs.k8s.io/controller-runtime/pkg/conversion)
package, plus the API version for our hub type (v1), and finally some of the
standard packages.
*/
import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	samplecontrollerv1alpha1 "github.com/govargo/foo-controller-kubebuilder/api/v1alpha1"
)

/*
Our "spoke" versions need to implement the
[`Convertible`](https://godoc.org/sigs.k8s.io/controller-runtime/pkg/conversion#Convertible)
interface.  Namely, they'll need `ConvertTo` and `ConvertFrom` methods to convert to/from
the hub version.
*/

/*
ConvertTo is expected to modify its argument to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/
// ConvertTo converts this Foo to the Hub version (v1alpha1).
func (src *Foo) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*samplecontrollerv1alpha1.Foo)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.DeploymentName = src.Spec.DeploymentName
	dst.Spec.Replicas = src.Spec.Replicas

	// Status
	dst.Status.AvailableReplicas = src.Status.AvailableReplicas

	return nil
}

/*
ConvertFrom is expected to modify its receiver to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/

// ConvertFrom converts from the Hub version (v1alpha1) to this version.
func (dst *Foo) ConvertFrom(srcRaw conversion.Hub) error {
	return nil
}