#ifndef PY_CC{XXX}_H
#define PY_CC{XXX}_H

#include "py_Common.h"

namespace py_cocos2d
{

typedef struct Py{XXX}
{
    PyObject_HEAD
    cocos2d::Ref    *ob_body;
}Py{XXX};

extern PyMethodDef Py{XXX}_methods[];
extern PyTypeObject Py{XXX}Type;

}
#endif
