// Copyright (c) 2019 The BFE Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// set available modules in bfe

package bfe_modules

import (
	"github.com/bfenetworks/bfe/bfe_module"
	"github.com/bfenetworks/bfe/bfe_modules/mod_access"
	"github.com/bfenetworks/bfe/bfe_modules/mod_auth_basic"
	"github.com/bfenetworks/bfe/bfe_modules/mod_auth_jwt"
	"github.com/bfenetworks/bfe/bfe_modules/mod_auth_request"
	"github.com/bfenetworks/bfe/bfe_modules/mod_block"
	"github.com/bfenetworks/bfe/bfe_modules/mod_compress"
	"github.com/bfenetworks/bfe/bfe_modules/mod_cors"
	"github.com/bfenetworks/bfe/bfe_modules/mod_doh"
	"github.com/bfenetworks/bfe/bfe_modules/mod_errors"
	"github.com/bfenetworks/bfe/bfe_modules/mod_geo"
	"github.com/bfenetworks/bfe/bfe_modules/mod_header"
	"github.com/bfenetworks/bfe/bfe_modules/mod_http_code"
	"github.com/bfenetworks/bfe/bfe_modules/mod_key_log"
	"github.com/bfenetworks/bfe/bfe_modules/mod_logid"
	"github.com/bfenetworks/bfe/bfe_modules/mod_markdown"
	"github.com/bfenetworks/bfe/bfe_modules/mod_prison"
	"github.com/bfenetworks/bfe/bfe_modules/mod_redirect"
	"github.com/bfenetworks/bfe/bfe_modules/mod_rewrite"
	"github.com/bfenetworks/bfe/bfe_modules/mod_secure_link"
	"github.com/bfenetworks/bfe/bfe_modules/mod_static"
	"github.com/bfenetworks/bfe/bfe_modules/mod_tag"
	"github.com/bfenetworks/bfe/bfe_modules/mod_trace"
	"github.com/bfenetworks/bfe/bfe_modules/mod_trust_clientip"
	"github.com/bfenetworks/bfe/bfe_modules/mod_userid"
	"github.com/bfenetworks/bfe/bfe_modules/mod_waf"
)

// 加载和实现的时候需要对类型实现一致性
// list of all modules, the order is very important
var moduleList = []bfe_module.BfeModule{
	// mod_trust_clientip
	mod_trust_clientip.NewModuleTrustClientIP(),

	// mod_logid
	// Requirement: After mod_trust_clientip
	// 不存在对应的配置项
	mod_logid.NewModuleLogId(),

	// mode_userid
	// 更改返回类型,删除了返回的 string 类型
	mod_userid.NewModuleUserID(),

	// mod_geo
	// Requirement: After mod_logid
	mod_geo.NewModuleGeo(),

	// mod_tag
	mod_tag.NewModuleTag(),

	// mod_trace
	mod_trace.NewModuleTrace(),

	// mod_cors
	mod_cors.NewModuleCors(),

	// mod_block
	// Requirement: After mod_logid
	//  这里存在两个load 是一个需要处理的点
	mod_block.NewModuleBlock(),

	// mod_prison
	// Requirement: After mod_logid
	mod_prison.NewModulePrison(),

	// mod_auth_basic
	// Requirement: before mod_static
	mod_auth_basic.NewModuleAuthBasic(),

	// mod_auth_jwt
	mod_auth_jwt.NewModuleAuthJWT(),

	// mod_secure_link
	mod_secure_link.NewModuleSecureLink(),

	// mod_waf
	mod_waf.NewModuleWaf(),

	// mod_doh
	// 如同logid 都是重新加载配置的方式不太一样
	mod_doh.NewModuleDoh(),

	// mod_redirect
	// Requirement: After mod_logid
	mod_redirect.NewModuleRedirect(),

	// mod_static
	mod_static.NewModuleStatic(),

	// mod_rewrite
	mod_rewrite.NewModuleReWrite(),

	// mod_header
	mod_header.NewModuleHeader(),

	// mod_auth_request
	mod_auth_request.NewModuleAuthRequest(),

	// mod_errors
	mod_errors.NewModuleErrors(),

	// mod_markdown
	mod_markdown.NewModuleMarkdown(),

	// mod_compress
	mod_compress.NewModuleCompress(),

	// mod_key_log
	mod_key_log.NewModuleKeyLog(),

	// mod_http_code
	//  LoadConfData nil
	mod_http_code.NewModuleHttpCode(),

	// mod_access
	// load nil
	mod_access.NewModuleAccess(),
}


var moduleL = []bfe_module.BfeModule{
	// mod_trust_clientip
	//mod_trust_clientip.,

	// mod_logid
	// Requirement: After mod_trust_clientip
	// 不存在对应的配置项
	mod_logid.NewModuleLogId(),

	// mode_userid
	// 更改返回类型,删除了返回的 string 类型
	mod_userid.NewModuleUserID(),

	// mod_geo
	// Requirement: After mod_logid
	mod_geo.NewModuleGeo(),

	// mod_tag
	mod_tag.NewModuleTag(),

	// mod_trace
	mod_trace.NewModuleTrace(),

	// mod_cors
	mod_cors.NewModuleCors(),

	// mod_block
	// Requirement: After mod_logid
	//  这里存在两个load 是一个需要处理的点
	mod_block.NewModuleBlock(),

	// mod_prison
	// Requirement: After mod_logid
	mod_prison.NewModulePrison(),

	// mod_auth_basic
	// Requirement: before mod_static
	mod_auth_basic.NewModuleAuthBasic(),

	// mod_auth_jwt
	mod_auth_jwt.NewModuleAuthJWT(),

	// mod_secure_link
	mod_secure_link.NewModuleSecureLink(),

	// mod_waf
	mod_waf.NewModuleWaf(),

	// mod_doh
	// 如同logid 都是重新加载配置的方式不太一样
	mod_doh.NewModuleDoh(),

	// mod_redirect
	// Requirement: After mod_logid
	mod_redirect.NewModuleRedirect(),

	// mod_static
	mod_static.NewModuleStatic(),

	// mod_rewrite
	mod_rewrite.NewModuleReWrite(),

	// mod_header
	mod_header.NewModuleHeader(),

	// mod_auth_request
	mod_auth_request.NewModuleAuthRequest(),

	// mod_errors
	mod_errors.NewModuleErrors(),

	// mod_markdown
	mod_markdown.NewModuleMarkdown(),

	// mod_compress
	mod_compress.NewModuleCompress(),

	// mod_key_log
	mod_key_log.NewModuleKeyLog(),

	// mod_http_code
	//  LoadConfData nil
	mod_http_code.NewModuleHttpCode(),

	// mod_access
	// load nil
	mod_access.NewModuleAccess(),
}

// init modules list
func InitModuleList(modules []bfe_module.BfeModule) {
	moduleList = modules
}

// add all modules
func SetModules() {
	for _, module := range moduleList {
		bfe_module.AddModule(module)
	}
}
//func LoadconfModules() {
//	for _, name := range bfe_module.ModulesAll {
//		module:= bfe_module.BfeModules.GetModule(name)
//		module.LoadConfData(nil)
//	}
//}
