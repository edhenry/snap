/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Coporation

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

package rest

import (
	"net/http"
	"strconv"

	"github.com/intelsdi-x/pulse/core"
	"github.com/intelsdi-x/pulse/core/cdata"
	"github.com/intelsdi-x/pulse/mgmt/rest/rbody"
	"github.com/julienschmidt/httprouter"
)

func (s *Server) getPluginConfigItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	styp := p.ByName("type")
	if styp == "" {
		cdn := s.mc.GetPluginConfigDataNodeAll()
		item := &rbody.PluginConfigItem{cdn}
		respond(200, item, w)
		return
	}
	var ityp int
	if ityp, err = strconv.Atoi(styp); err != nil {
		respond(400, rbody.FromError(err), w)
		return
	}
	name := p.ByName("name")
	sver := p.ByName("version")
	var iver int
	if sver != "" {
		if iver, err = strconv.Atoi(sver); err != nil {
			respond(400, rbody.FromError(err), w)
			return
		}
	} else {
		iver = -2
	}

	cdn := s.mc.GetPluginConfigDataNode(core.PluginType(ityp), name, iver)
	item := &rbody.PluginConfigItem{cdn}
	respond(200, item, w)
}

func (s *Server) deletePluginConfigItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	styp := p.ByName("type")
	var ityp int
	if styp != "" {
		if ityp, err = strconv.Atoi(styp); err != nil {
			respond(400, rbody.FromError(err), w)
			return
		}
	}
	name := p.ByName("name")
	sver := p.ByName("version")
	var iver int
	if sver != "" {
		if iver, err = strconv.Atoi(sver); err != nil {
			respond(400, rbody.FromError(err), w)
			return
		}
	} else {
		iver = -2
	}

	src := []string{}
	errCode, err := marshalBody(&src, r.Body)
	if errCode != 0 && err != nil {
		respond(400, rbody.FromError(err), w)
		return
	}

	var res cdata.ConfigDataNode
	if styp == "" {
		res = s.mc.DeletePluginConfigDataNodeFieldAll(src...)
	} else {
		res = s.mc.DeletePluginConfigDataNodeField(core.PluginType(ityp), name, iver, src...)
	}

	item := &rbody.DeletePluginConfigItem{res}
	respond(200, item, w)
}

func (s *Server) setPluginConfigItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	styp := p.ByName("type")
	var ityp int
	if styp != "" {
		if ityp, err = strconv.Atoi(styp); err != nil {
			respond(400, rbody.FromError(err), w)
			return
		}
	}
	name := p.ByName("name")
	sver := p.ByName("version")
	var iver int
	if sver != "" {
		if iver, err = strconv.Atoi(sver); err != nil {
			respond(400, rbody.FromError(err), w)
			return
		}
	} else {
		iver = -2
	}

	src := cdata.NewNode()
	errCode, err := marshalBody(src, r.Body)
	if errCode != 0 && err != nil {
		respond(400, rbody.FromError(err), w)
		return
	}

	var res cdata.ConfigDataNode
	if styp == "" {
		res = s.mc.MergePluginConfigDataNodeAll(src)
	} else {
		res = s.mc.MergePluginConfigDataNode(core.PluginType(ityp), name, iver, src)
	}

	item := &rbody.SetPluginConfigItem{res}
	respond(200, item, w)
}
