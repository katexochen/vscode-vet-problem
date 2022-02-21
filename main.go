package main

import "google.golang.org/genproto/googleapis/cloud/compute/v1"

func main() {
	_ = GetGlobalSynced()
}

func GetGlobalSynced() compute.InsertInstanceTemplateRequest {
	req := compute.InsertInstanceTemplateRequest{}
	req = req
	return req
}
