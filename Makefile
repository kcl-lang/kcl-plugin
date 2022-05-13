# Copyright 2020 The KCL Authors. All rights reserved.

default:

kcl-test:
	kcl-test ./...

test:
	cd _test && python3 -m pytest

clean:
