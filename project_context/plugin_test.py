# Copyright The KCL Authors. All rights reserved.

# python3 -m pytest

import plugin


def test_get_project_current_path():
    assert plugin.get_project_current_path() == None

def test_get_project_input_file():
    assert plugin.get_project_input_file() == None

def test_get_project_context():
    assert plugin.get_project_context() == {}

def test_get_project_crds():
    assert plugin.get_project_crds() == []