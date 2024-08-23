# Copyright The KCL Authors. All rights reserved.

# python3 -m pytest

import plugin


def test_plugin_hello_add():
    assert plugin.add(1, 2) == 3


def test_plugin_hello_tolower():
    assert plugin.tolower('KCL') == 'kcl'


def test_plugin_hello_update_dict():
    assert plugin.update_dict({'name': 123}, 'name', 'kcl')['name'] == 'kcl'


def test_plugin_hello_list_append():
    l = plugin.list_append(['abc'], 'name', 123)
    assert len(l) == 3
    assert l[0] == 'abc'
    assert l[1] == 'name'
    assert l[2] == 123


def test_plugin_hello_foo():
    v = plugin.foo('aaa', 'bbb', x=123, y=234, abcd=1234)
    assert len(v) == 5
    assert v['a'] == 'aaa'
    assert v['b'] == 'bbb'
    assert v['x'] == 123
    assert v['y'] == 234
    assert v['abcd'] == 1234

    v = plugin.foo('aaa', 'bbb', x=123)
    assert len(v) == 3
    assert v['a'] == 'aaa'
    assert v['b'] == 'bbb'
    assert v['x'] == 123
