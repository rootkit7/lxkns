```plantuml
hide empty fields
hide empty methods
!define L <size:12><&link-intact></size><i>

interface Namespace {
  L Leaders() []*Process
  L Ealdorman() *Process
}

Namespace ---> "0,1" Process : Ealdorman
Namespace ---> "*" Process : Leaders

class ProcessTable
ProcessTable -> Process : "[PID]"

class Process {
  L Parent *Process
  L Namespaces NamespacesSet
}

Process --> "7" Namespace : Namespaces
Process "*" --> Process : Parent
```
