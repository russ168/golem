/*
   Copyright (C) 2003-2010 Institute for Systems Biology
                           Seattle, Washington, USA.

   This library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   This library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with this library; if not, write to the Free Software
   Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA 02111-1307  USA

*/
package main

import (
	"crypto/x509"
)

//TODO: make these vars not declared when not needed

//buffered channel for use as an incrementer to keep track of submissions
var subidChan = make(chan int, 1)

//buffered channel for creating jobs
var jobChan = make(chan *Job, 1000)


//map of submissions by id
var subMap = map[int]*Submission{}

//tls configurability
var useTls bool = true
var clientCert * x509.Certificate


const (
	//Message type constants
	HELLO   = 1
	DONE    = 2
	START   = 3
	CHECKIN = 4

	COUT   = 5
	CERROR = 6
)