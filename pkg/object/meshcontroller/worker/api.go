/*
 * Copyright (c) 2017, MegaEase
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://wwwrk.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package worker

import (
	"net/http"

	"github.com/megaease/easegress/pkg/object/meshcontroller/spec"
)

const (
	// meshEurekaPrefix is the mesh eureka registry API url prefix.
	meshEurekaPrefix = "/mesh/eureka"

	// meshNacosPrefix is the mesh nacos registyr API url prefix.
	meshNacosPrefix = "/nacos/v1"
)

func (wrk *Worker) runAPIServer() {
	var apis []*apiEntry
	switch wrk.registryServer.RegistryType {
	case spec.RegistryTypeConsul:
		apis = wrk.consulAPIs()
	case spec.RegistryTypeEureka:
		apis = wrk.eurekaAPIs()
	case spec.RegistryTypeNacos:
		apis = wrk.nacosAPIs()
	default:
		apis = wrk.eurekaAPIs()
	}
	wrk.apiServer.registerAPIs(apis)
}

func (wrk *Worker) emptyHandler(w http.ResponseWriter, r *http.Request) {
	// EaseMesh does not need to implement some APIS like
	// delete, heartbeat of Eureka/Consul/Nacos.
}
