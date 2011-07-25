/*
   Copyright (C) 2003-2011 Institute for Systems Biology
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
    "time"
)

type ItemsHandle struct {
	Items         []interface{}
	NumberOfItems int
}

type JobDetailsList struct {
    Items []JobDetails
    NumberOfItems int
}

type WorkerNodeList struct {
    Items []WorkerNode
    NumberOfItems int
}

type WorkerNode struct {
	NodeId   string
	Uri      string
	Hostname string
	MaxJobs  int
	Running  int
}

func NewWorkerNode(nh *NodeHandle) WorkerNode {
	maxJobs, running := nh.Stats()
	return WorkerNode{NodeId: nh.NodeId, Uri: nh.Uri, Hostname: nh.Hostname,
		MaxJobs: maxJobs, Running: running}
}

// reformed entities
type Task struct {
	Count int
	Args  []string
}

type JobDetails struct {
    JobId string
    Uri string

    Owner string
    Label string
    Type string

	FirstCreated *time.Time
	LastModified *time.Time

	Total  int
	Finished  int
	Errored  int

    Running       bool
    Scheduled   bool

    Tasks []Task
}

func (this JobDetails) isComplete() bool {
    return this.Total == (this.Finished + this.Errored)
}

func NewJobDetails(jobId string, owner string, label string, jobtype string, tasks []Task) JobDetails {
    now := time.Time{}

	totalTasks := 0
	for _, task := range tasks {
		totalTasks += task.Count
	}

    return JobDetails{
        JobId: jobId, Uri: "/jobs/" + jobId,
        Owner: owner, Label: label, Type: jobtype,
        Total:totalTasks,Finished:0,Errored:0,
        Running: false, Scheduled: false,
        FirstCreated: &now, LastModified: &now,
        Tasks: tasks }
}
