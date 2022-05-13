# Copyright 2021 The KCL Authors. All rights reserved.

import kcl_plugin.hello as hello

schema TestHello_tolower:
    name = hello.tolower('KCL')
    assert name == 'kcl', "${name}"

schema TestHello_add:
    v = hello.add(1, 2)
    assert v == 3, "${v}"

schema TestHello_update_dict:
    assert hello.update_dict({'name': 123}, 'name', 'kcl')['name'] == 'kcl'

schema TestHello_list_append:
    l = hello.list_append(['abc'], 'name', 123)
    assert len(l) == 3
    assert l[0] == 'abc'
    assert l[1] == 'name'
    assert l[2] == 123


schema TestHello_foo:
    v = hello.foo('aaa', 'bbb', x=123, y=234, abcd=1234)
    assert len(v) == 5
    assert v['a'] == 'aaa'
    assert v['b'] == 'bbb'
    assert v['x'] == 123
    assert v['y'] == 234
    assert v['abcd'] == 1234

    v = hello.foo('aaa', 'bbb', x=123)
    assert len(v) == 3
    assert v['a'] == 'aaa'
    assert v['b'] == 'bbb'
    assert v['x'] == 123
