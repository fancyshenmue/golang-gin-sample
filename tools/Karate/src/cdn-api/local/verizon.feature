Feature: Test Verizon API
  Background:
    * url 'http://127.0.0.1:8080'

  Scenario: List HTTP Large
    * def data = read('karate_config_data.json')

    Given url data["scenario"]["verizon"]["list_httpLarge"]["url"]
    And headers data["scenario"]["verizon"]["list_httpLarge"]["header"]
    When method get
    Then status 200

  Scenario: Customer Origins Get
    * def data = read('karate_config_data.json')

    Given url data["scenario"]["verizon"]["customer_origins_get"]["url"]
    And headers data["scenario"]["verizon"]["customer_origins_get"]["header"]
    When method get
    Then status 200

#   Scenario: Customer Origins Add
#     * def data = read('karate_config_data.json')

#     Given url data["scenario"]["verizon"]["customer_origins_add_httplarge"]["url"]
#     And headers data["scenario"]["verizon"]["customer_origins_add_httplarge"]["header"]
#     And request data["scenario"]["verizon"]["customer_origins_add_httplarge"]["body"]
#     When method post
#     Then status 200

#   Scenario: CNAMES Add
#     * def data = read('karate_config_data.json')

#     Given url data["scenario"]["verizon"]["cnames_add"]["url"]
#     And headers data["scenario"]["verizon"]["cnames_add"]["header"]
#     And request data["scenario"]["verizon"]["cnames_add"]["body"]
#     When method post
#     Then status 200
