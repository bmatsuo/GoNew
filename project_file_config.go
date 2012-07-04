// Copyright 2012, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*  Filename:    project_file_config.go
 *  Author:      Bryan Matsuo <bryan.matsuo [at] gmail.com>
 *  Created:     2012-07-03 18:13:06.081945 -0700 PDT
 *  Description: 
 */

import ()

type ProjectFileConfig struct {
	Path      string // a template
	Type      string
	Templates []string // template names
}