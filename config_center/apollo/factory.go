/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package apollo

import (
	"github.com/apache/dubbo-go/common"
	"github.com/apache/dubbo-go/common/extension"
	"github.com/apache/dubbo-go/config_center"
	"github.com/apache/dubbo-go/config_center/parser"
)

func init() {
	extension.SetConfigCenterFactory("apollo", createDynamicConfigurationFactory)
}

func createDynamicConfigurationFactory() config_center.DynamicConfigurationFactory {
	return &apolloConfigurationFactory{}
}

type apolloConfigurationFactory struct{}

// GetDynamicConfiguration gets the dynamic configuration
func (f *apolloConfigurationFactory) GetDynamicConfiguration(url *common.URL) (config_center.DynamicConfiguration, error) {
	dynamicConfiguration, err := newApolloConfiguration(url)
	if err != nil {
		return nil, err
	}
	dynamicConfiguration.SetParser(&parser.DefaultConfigurationParser{})
	return dynamicConfiguration, err

}
