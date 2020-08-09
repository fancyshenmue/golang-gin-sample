Feature: Test JWT
  Background:
    * url 'http://127.0.0.1:8080'

  Scenario: Login
    * def data = read('karate_config_data.json')

    Given url data["scenario"]["login"]["url"]
    And request data["scenario"]["login"]["body"]
    When method post
    Then status 200
