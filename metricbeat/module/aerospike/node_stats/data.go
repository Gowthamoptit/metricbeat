// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package node_stats

import (
	s "github.com/elastic/beats/v7/libbeat/common/schema"
	c "github.com/elastic/beats/v7/libbeat/common/schema/mapstrstr"
)

var schema = s.Schema{
	"system": s.Object{
		"free": s.Object{
			"memory_kbytes":     c.Int("system_free_mem_kbytes"),
			"memory_pct":	c.Float("system_free_mem_pct"),
			"thp_mem_kbytes":     c.Int("system_thp_mem_kbytes"),
		},
		"cpu": s.Object{
			"kernel_pct":     c.Float("system_kernel_cpu_pct"),
			"total_pct":     c.Float("system_total_cpu_pct"),
			"user_pct":     c.Float("system_user_cpu_pct"),
		},
	},
	"heap": s.Object{
		"active": s.Object{
			"bytes": c.Int("heap_active_kbytes", s.Optional),
		},
		"allocated": s.Object{
			"bytes": c.Int("heap_allocated_kbytes", s.Optional),
		},
		"efficiency": s.Object{
			"pct": c.Float("heap_efficiency_pct", s.Optional),
		},
		"mapped": s.Object{
			"bytes": c.Int("heap_mapped_kbytes", s.Optional),
		},
		"site": s.Object{
                        "count": c.Float("heap_site_count", s.Optional),
                },
	},
	"client_connections": c.Int("client_connections"),
	"heartbeat": s.Object{
		"connections": s.Object{
			"total": c.Int("heartbeat_connections"),
			"closed": c.Int("heartbeat_connections_closed"),
			"opened": c.Int("heartbeat_connections_opened"),
			"received_foreign": c.Int("heartbeat_received_foreign"),
			"received_self": c.Int("heartbeat_received_self"),
		},
	},
	"info": s.Object{
		"complete": c.Int("info_complete"),
		"max_ms":  c.Int("info_max_ms"),
		"queue":  c.Int("info_queue"),
		"threads":  c.Int("info_threads"),
		"timeout":  c.Int("info_timeout"),
	},
	"failed": c.Bool("failed_best_practices"),
        "objects": c.Int("objects"),
}
