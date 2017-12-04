#Copyright Uhuchain. All Rights Reserved.
#
#SPDX-License-Identifier: Apache-2.0
#
#


# Tool commands (overridable)
GO_CMD             ?= go
GO_DEP_CMD         ?= dep

dependencies:
	dep ensure
	rm -rf vendor/github.com/hyperledger/fabric

clean:
	$(GO_CMD) clean
