# Copyright The KCL Authors. All rights reserved.

# python3 -m pytest

import plugin

def test_http_get():
    assert plugin.get("https://www.kcl-lang.io/")["status"] == 200
