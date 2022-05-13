# Copyright 2020 The KCL Authors. All rights reserved.

import os
import ruamel.yaml as _yaml
from pathlib import Path
from typing import Any, Dict, List

import kclvm.config as kcfg

INFO = {
    'name': 'project_context',
    'describe': 'project_context extract base info from project.yaml&stack.yaml',
    'long_describe': 'project_context extract base info from project.yaml&stack.yaml',
    'version': '0.0.1',
}
BASE_PKG_PATH = "base/pkg"
STACK_YAML_FILE = "stack.yaml"
PROJECT_YAML_FILE = "project.yaml"
CRD_DIR_NAME = "crd"


def get_project_current_path() -> str:
    return kcfg.current_path


def get_project_input_file() -> List[str]:
    return kcfg.input_file


def get_stack_context() -> Dict[str, Any]:
    """get_stack_context returns the current stack context"""
    stackRoot = find_stack_root()
    if stackRoot:
        return loadIfExist(stackRoot / STACK_YAML_FILE)
    return {}


def get_project_context() -> Dict[str, Any]:
    """get_project_context returns the current project context"""
    projectRoot = find_project_root()
    if projectRoot:
        return loadIfExist(projectRoot / PROJECT_YAML_FILE)
    return {}


def get_project_crds() -> List[Any]:
    """get_project_crd returns the crd yaml document array in current project dir"""
    projectStackRoot = find_stack_root()
    if not projectStackRoot:
        return []
    
    projectProjectRoot = projectStackRoot.parent
    projectCrdDir = projectProjectRoot / CRD_DIR_NAME

    crdYamlDocuments = []
    if projectCrdDir.exists():
        for root, dirs, files in sorted(os.walk(projectCrdDir)):
            for file in sorted(files):
                crdFile = (Path(root) / file)
                # safe_load_all() returns a sequence of Python objects corresponding to the documents in the stream
                yamlDocuments = _yaml.safe_load_all(crdFile.read_text())
                crdYamlDocuments += list(yamlDocuments)
    return crdYamlDocuments


def find_stack_root() -> Path:
    """find_stack_root returns the root dir path of current project stack"""
    if not kcfg.input_file or len(kcfg.input_file) == 0 or not kcfg.current_path or len(kcfg.current_path) == 0:
        return None
    inputfileDirListForProject = []
    # filter base dir
    for inputfile in kcfg.input_file:
        if BASE_PKG_PATH not in inputfile:
            inputfileDirListForProject.append(Path(os.path.abspath(inputfile)).parent)
    # return empty dir if no project input file dir found
    if len(inputfileDirListForProject) == 0:
        return None
    for inputfileDir in inputfileDirListForProject:
        # if current dir contains stack.yaml, set current dir as project stack root path
        if (inputfileDir / STACK_YAML_FILE).is_file():
            return inputfileDir
    return None


def find_project_root() -> Path:
    """find_project_root returns the root dir path of current project"""
    stackRoot = find_stack_root()
    return stackRoot.parent if stackRoot else None


def loadIfExist(yamlFile: Path) -> Dict[str, Any]:
    if yamlFile.is_file():
        return _yaml.safe_load(yamlFile.read_text())
    return {}
