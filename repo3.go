package main

import (
	repo1 "github.com/RuniVN/test-repo-1"
	repo2 "github.com/RuniVN/test-repo-2"
)

func main() {
	repo1.DoSomething()
	repo2.DoSomethingRepo2WithRepo1()
}
