# coq

## Planning

* [ ] Message Encoding <-----
  * [ ] Generate golang package from `usp-record.proto` 
  * [ ] Generate golang package from `usp-msg.proto`
* [ ] Messages
  * [ ] Requests, Responses and Errors
  * [ ] CRUDs
  * [ ] State/Capabilities (a.k.a. Profiles)
  * [ ] Notififications/Subsriptions
  * [ ] Operations
  * [ ] Error codes
* [ ] Message Transfer Protocols (MTPs)
  * [ ] CoAP Binding (DEPRECATED => optional)
  * [ ] WebSocket Binding (mandatory) <-----
  * [ ] STOMP Binding (optional)
  * [ ] MQTT Binding (mandatory)
  * [ ] UDSB Binding (optional)
* [ ] End to End Message Exchange
  * [ ] Within E2E Session Context
  * [ ] Without E2E Session Context
* [ ] Controller API:
  * [x] Implement `swagger-usp-controller-v1` stubs
  * [ ] Implement `swagger-usp-controller-v1` carefully
* [ ] Authentication and Authorization
  * [ ] TLS
  * [ ] RBAC
* [ ] Annex A: Bulk Data Collection <-----
  * [ ] Encoding
    * [ ] CSV Bulk Data
    * [ ] JSON Bulk Data
  * [ ] Collectors
    * [ ] HTTP Bulk Data Collection
    * [ ] MQTT Bulk Data Collection
    * [ ] USPEventNotif Bulk Data Collection
* [ ] Appendix I: Software Module Management
  * [ ] TODO reqs: DU/EU/Faults
