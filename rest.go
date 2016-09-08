/*
<!--
Copyright (c) 2016 Christoph Berger. Some rights reserved.
Use of this text is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = "Gimme a REST."
description = "The basics of a RESTful Web API, with a tiny REST server in Go."
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2016-09-22"
publishdate = "2016-09-22"
domains = ["Internet and Web"]
tags = ["REST", "Web", "API", "Video"]
categories = ["Tutorial"]
+++

RESTful Web API's are ubiquitous. Time for a minimalistic, five-minutes video tutorial about REST, RESTful API's, and buidling a REST server in Go.

<!--more-->

- - -
*This is the transcript of the video.*
- - -


## The code
*/

// ## Imports and globals
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	data map[string]string
)

func show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	k := p.ByName("key")
	if k == "" {
		fmt.Fprintf(w, "Read list: %v", data)
		return
	}
	fmt.Fprintf(w, "Read entry: data[%s] = %s", k, data[k])
}

func update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	k := p.ByName("key")
	v := p.ByName("value")
	data[k] = v
	fmt.Fprintf(w, "Updated: data[%s] = %s", k, data[k])
}

func main() {
	flag.Parse()
	data = map[string]string{}
	r := httprouter.New()
	r.GET("/entry/:key", show)
	r.GET("/list", show)
	r.PUT("/entry/:key/:value", update)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
